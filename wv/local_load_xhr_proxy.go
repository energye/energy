//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/energye/lcl/emfs"
	"github.com/energye/wv/windows"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// IXHRProxy
//
//	本地资源加载 XHR 请求代理接口
type IXHRProxy interface {
	Send(request wv.ICoreWebView2WebResourceRequestRef) (*XHRProxyResponse, error) // 被动调用，发送请求，在浏览器进程同步执行
}

// XHRProxy
//
//	数据请求代理
type XHRProxy struct {
	Scheme     LocalProxyScheme // http/https/tcp default: http
	IP         string           // default: localhost
	Port       int              // default: 80
	SSL        XHRProxySSL      // https 安全证书配置
	HttpClient *HttpClient      // http/https 客户端, 可自定义配置
}

// XHRProxySSL
//
//	https证书配置，如果其中某一配置为空，则跳过ssl检查, 如果证书配置错误则请求失败
type XHRProxySSL struct {
	FS      emfs.IEmbedFS // 证书到内置执行文件时需要设置
	RootDir string        // 根目录 如果使用 FS 时目录名 root/path, 否则本地目录/to/root/path
	Cert    string        // RootDir/to/path/cert.crt
	Key     string        // RootDir/to/path/key.key
	CARoots []string      // RootDir/to/path/ca.crt
}

// HttpClient
//
//	http/https 客户端
type HttpClient struct {
	Transport *http.Transport
	Client    *http.Client
	Jar       *cookiejar.Jar
	Timeout   time.Duration
}

// XHRProxyResponse
//
//	代理响应数据
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
//
//	被动调用，发送请求，在浏览器进程同步执行
func (m *XHRProxy) Send(request wv.ICoreWebView2WebResourceRequestRef) (*XHRProxyResponse, error) {
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
							path = strings.Replace(filepath.Join(m.SSL.RootDir, path), "\\", "/", -1)
							data, err = m.SSL.FS.ReadFile(path)
						} else {
							path = filepath.Join(m.SSL.RootDir, path)
							data, err = ioutil.ReadFile(path)
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

func (m *XHRProxy) sendHttp(request wv.ICoreWebView2WebResourceRequestRef) (*XHRProxyResponse, error) {
	return m.send("http://", request)
}

func (m *XHRProxy) sendHttps(request wv.ICoreWebView2WebResourceRequestRef) (*XHRProxyResponse, error) {
	return m.send("https://", request)
}

func (m *XHRProxy) send(scheme string, request wv.ICoreWebView2WebResourceRequestRef) (*XHRProxyResponse, error) {
	reqUrl, err := url.Parse(request.URI())
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
	return nil, nil
}

func (m *XHRProxy) sendTcp(request wv.ICoreWebView2WebResourceRequestRef) (*XHRProxyResponse, error) {
	return nil, errors.New("tcp unrealized")
}
