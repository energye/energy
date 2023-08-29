//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"embed"
	"errors"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// IXHRProxy
//  本地资源加载 XHR 请求代理接口
type IXHRProxy interface {
	Send(request *ICefRequest) (*XHRProxyResponse, error) // 被动调用，发送请求，在浏览器进程同步执行
}

// XHRProxy
//  数据请求代理
type XHRProxy struct {
	Scheme     LocalProxyScheme // http/https/tcp default: http
	IP         string           // default: localhost
	Port       int              // default: 80
	SSL        XHRProxySSL      // https 安全证书配置
	HttpClient *HttpClient      // http/https 客户端, 可自定义配置
}

// XHRProxySSL
//  https证书配置，如果其中某一配置为空，则跳过ssl检查, 如果证书配置错误则请求失败
type XHRProxySSL struct {
	FS      *embed.FS // 证书到内置执行文件时需要设置
	RootDir string    // 根目录 如果使用 FS 时目录名 root/path, 否则本地目录/to/root/path
	Cert    string    // RootDir/to/path/cert.crt
	Key     string    // RootDir/to/path/key.key
	CARoots []string  // RootDir/to/path/ca.crt
}

// HttpClient
//  http/https 客户端
type HttpClient struct {
	Transport *http.Transport
	Client    *http.Client
	Jar       *cookiejar.Jar
	Timeout   time.Duration
}

// XHRProxyResponse
//  代理响应数据
type XHRProxyResponse struct {
	Data       []byte              // 响应数据
	DataSize   int                 // 响应数据大小
	StatusCode int32               // 响应状态码
	Status     string              //
	Header     map[string][]string // 响应头
}

func (m *XHRProxySSL) skipVerify() bool {
	return m.RootDir == "" || m.Cert == "" || m.Key == "" || len(m.CARoots) == 0
}

// Send
//  被动调用，发送请求，在浏览器进程同步执行
func (m *XHRProxy) Send(request *ICefRequest) (*XHRProxyResponse, error) {
	if m.Scheme == LpsHttp {
		return m.sendHttp(request)
	} else if m.Scheme == LpsHttps {
		return m.sendHttps(request)
	} /* else if m.Scheme == LpsTcp {
		return m.sendTcp(request)
	}*/
	return nil, errors.New("incorrect scheme")
}

// XHR代理配置
// 如果配置代理，并且是 XHRProxy 时调用
// 否则你可以自己实现代理， 实现 IXHRProxy 接口，自定义代理请求
func (m *XHRProxy) init() {
	if m.Scheme == LpsHttp || m.Scheme == LpsHttps {
		if m.IP == "" {
			m.IP = "localhost"
		}
		if m.HttpClient == nil {
			m.HttpClient = new(HttpClient)
		}
		if m.Scheme == LpsHttps {
			if m.SSL.skipVerify() {
				if m.HttpClient.Transport == nil {
					m.HttpClient.Transport = &http.Transport{
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					}
				}
			} else {
				if m.HttpClient.Transport == nil {
					var (
						certPEMBlock, keyPEMBlock []byte
						err                       error
					)
					var readFile = func(path string) (data []byte, err error) {
						if m.SSL.FS != nil {
							path = strings.ReplaceAll(filepath.Join(m.SSL.RootDir, path), "\\", "/")
							data, err = m.SSL.FS.ReadFile(path)
						} else {
							path = filepath.Join(m.SSL.RootDir, path)
							data, err = os.ReadFile(path)
						}
						return
					}
					certPEMBlock, err = readFile(m.SSL.Cert)
					if err != nil {
						panic(err)
						return
					}
					keyPEMBlock, err = readFile(m.SSL.Key)
					if err != nil {
						panic(err)
						return
					}
					cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
					if err != nil {
						panic(err)
						return
					}
					pool := x509.NewCertPool()
					for _, path := range m.SSL.CARoots {
						if ca, err := readFile(path); err == nil {
							pool.AppendCertsFromPEM(ca)
						} else {
							panic(err)
						}
					}
					m.HttpClient.Transport = &http.Transport{
						TLSClientConfig: &tls.Config{
							Certificates: []tls.Certificate{cert},
							RootCAs:      pool,
						},
					}
				}
			}
		}
		if m.HttpClient.Jar == nil {
			if jar, err := cookiejar.New(nil); err == nil {
				m.HttpClient.Jar = jar
			} else {
				println("[Error] XHRProxy SSL New cookiejar:", err.Error())
			}
		}
		if m.HttpClient.Client == nil {
			if m.HttpClient.Timeout <= 0 {
				m.HttpClient.Timeout = time.Second * 30
			}
			m.HttpClient.Client = &http.Client{
				Jar:     m.HttpClient.Jar,
				Timeout: m.HttpClient.Timeout,
			}
		}
		if m.HttpClient.Client.Transport == nil && m.HttpClient.Transport != nil {
			m.HttpClient.Client.Transport = m.HttpClient.Transport
		}
	}
}

func (m *XHRProxy) sendHttp(request *ICefRequest) (*XHRProxyResponse, error) {
	return m.send("http://", request)
}

func (m *XHRProxy) sendHttps(request *ICefRequest) (*XHRProxyResponse, error) {
	return m.send("https://", request)
}

func (m *XHRProxy) send(scheme string, request *ICefRequest) (*XHRProxyResponse, error) {
	reqUrl, err := url.Parse(request.URL())
	if err != nil {
		return nil, err
	}
	// 构造目标地址
	targetUrl := new(bytes.Buffer)
	targetUrl.WriteString(scheme)
	targetUrl.WriteString(m.IP)
	if m.Port > 0 { // ip:port
		targetUrl.WriteString(":")
		targetUrl.WriteString(strconv.Itoa(m.Port))
	}
	targetUrl.WriteString(reqUrl.Path)
	if reqUrl.RawQuery != "" {
		targetUrl.WriteString("?")
		targetUrl.WriteString(reqUrl.RawQuery)
	}
	// 读取请求数据
	requestData := new(bytes.Buffer)
	data := request.GetPostData()
	if data.IsValid() {
		dataCount := int(data.GetElementCount())
		elements := data.GetElements()
		for i := 0; i < dataCount; i++ {
			element := elements.Get(uint32(i))
			switch element.GetType() {
			case PDE_TYPE_EMPTY:
			case PDE_TYPE_BYTES:
				if byt, c := element.GetBytes(); c > 0 {
					requestData.Write(byt)
				}
			case PDE_TYPE_FILE:
				if f := element.GetFile(); f != "" {
					if byt, err := ioutil.ReadFile(f); err == nil {
						requestData.Write(byt)
					}
				}
			}
			element.Free()
		}
		data.Free()
	}
	tarUrl := targetUrl.String()
	if logger.Enable() {
		logger.Debug("XHRProxy URL:", tarUrl, "method:", request.Method(), "data-size:", requestData.Len())
	}
	httpRequest, err := http.NewRequest(request.Method(), tarUrl, requestData)
	if err != nil {
		return nil, err
	}
	// 设置请求头
	header := request.GetHeaderMap()
	if header.IsValid() {
		size := header.GetSize()
		for i := 0; i < int(size); i++ {
			key := header.GetKey(uint32(i))
			c := header.FindCount(key)
			for j := 0; j < int(c); j++ {
				value := header.GetEnumerate(key, uint32(j))
				httpRequest.Header.Add(key, value)
			}
		}
		header.Free()
	}
	//httpRequest.Header.Add("Host", "www.example.com")
	//httpRequest.Header.Add("Origin", "https://www.example.com")
	//httpRequest.Header.Add("Referer", "https://www.example.com/")
	if m.HttpClient.Client == nil {
		return nil, errors.New("http client is nil")
	}
	httpResponse, err := m.HttpClient.Client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	// 读取响应头
	responseHeader := make(map[string][]string)
	for key, value := range httpResponse.Header {
		for _, vs := range value {
			if header, ok := responseHeader[key]; ok {
				responseHeader[key] = append(header, vs)
			} else {
				responseHeader[key] = []string{vs}
			}
		}
	}
	// 读取响应数据
	buf := new(bytes.Buffer)
	c, err := buf.ReadFrom(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	status := "OK"
	if httpResponse.StatusCode != 200 {
		rs := strings.Split(httpResponse.Status, " ")
		if len(rs) > 1 {
			status = rs[1]
		} else {
			status = httpResponse.Status
		}
	}
	result := &XHRProxyResponse{
		Data:       buf.Bytes(),
		DataSize:   int(c),
		StatusCode: int32(httpResponse.StatusCode),
		Status:     status,
		Header:     responseHeader,
	}
	return result, nil
}

func (m *XHRProxy) sendTcp(request *ICefRequest) (*XHRProxyResponse, error) {
	return nil, errors.New("tcp unrealized")
}
