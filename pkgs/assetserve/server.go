//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 简易的内置静态资源http服务
//	由于使用go http server该服务可能会引起一些安全软件报毒
//	报毒解决方案
//		1. 配置 ssl 证书, 未测试
//		2. 编译时 使用 -ldflags "-s -w" 去除调试信息和符号
//		3. 使用 upx 工具压缩执行文件
//	最好是 2 和 3 配合使用

package assetserve

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/energye/golcl/energy/emfs"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var contentType = map[string]string{}

// AssetsServerHeaderKeyName
//
//	用于保证链接的安全的请求头(header)key名称
var AssetsServerHeaderKeyName = "ASSETS_SERVER_KEY"

// AssetsServerHeaderKeyValue
//
//	用于保证链接的安全的key值
//	这里简单的在请求所有资源时增加请求头的判断
//	不为空时生效
var AssetsServerHeaderKeyValue string

type assetsHttpServer struct {
	LocalAssets  string        //本地静态资源目录 示例: /app/assets/   http://127.0.0.1:8888/demo/demo.html -> /app/assets/demo/demo.html
	AssetsFSName string        //静态资源内置FS目录名 默认值: resources
	Assets       emfs.IEmbedFS //静态资源内置FS目录对象
	IP           string        //默认值: 127.0.0.1
	PORT         int           //默认值: 80
	SSL          *SSL          //设置后启动https
}

// SSL 证书配置，根据 Assets 或 LocalAssets 寻找证书文件位置
type SSL struct {
	SSLCert string
	SSLKey  string
}

func init() {
	var types = strings.Split(mimeTypes, "\n")
	for _, mime := range types {
		mime = strings.TrimSpace(mime)
		var m = strings.Split(mime, "=")
		if len(m) == 2 {
			contentType["."+m[0]] = m[1]
		}
	}
}

// NewAssetsHttpServer
//
//	创建静态资源http服务
func NewAssetsHttpServer() *assetsHttpServer {
	return &assetsHttpServer{
		AssetsFSName: "resources",
		IP:           "0.0.0.0",
		PORT:         80,
	}
}

func (m *assetsHttpServer) serveTLS(addr string, handler http.Handler) {
	server := &http.Server{Addr: addr, Handler: handler}
	if addr == "" {
		addr = ":https"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		println("serverTLS Listen failed error:", err)
		return
	}
	var config = &tls.Config{}
	config.NextProtos = append(config.NextProtos, "http/1.1")
	configHasCert := len(config.Certificates) > 0 || config.GetCertificate != nil
	if !configHasCert || m.SSL.SSLCert != "" || m.SSL.SSLKey != "" {
		var loadX509KeyPair = func(certFile, keyFile string) (tls.Certificate, error) {
			var (
				certPEMBlock, keyPEMBlock []byte
				err                       error
			)
			if m.Assets != nil {
				certPEMBlock, err = m.Assets.ReadFile(m.AssetsFSName + certFile)
				if err != nil {
					return tls.Certificate{}, err
				}
				keyPEMBlock, err = m.Assets.ReadFile(m.AssetsFSName + keyFile)
				if err != nil {
					return tls.Certificate{}, err
				}
			} else if m.LocalAssets != "" {
				certPEMBlock, err = ioutil.ReadFile(m.LocalAssets + certFile)
				if err != nil {
					return tls.Certificate{}, err
				}
				keyPEMBlock, err = ioutil.ReadFile(m.LocalAssets + keyFile)
				if err != nil {
					return tls.Certificate{}, err
				}
			} else {
				if err != nil {
					return tls.Certificate{}, errors.New("resource directory is not configured")
				}
			}
			return tls.X509KeyPair(certPEMBlock, keyPEMBlock)
		}
		config.Certificates = make([]tls.Certificate, 1)
		config.Certificates[0], err = loadX509KeyPair(m.SSL.SSLCert, m.SSL.SSLKey)
		if err != nil {
			println("serverTLS loadX509KeyPair failed error:", err)
			return
		}
	}
	tlsListener := tls.NewListener(ln, config)
	go func() {
		defer ln.Close()
		if err = server.Serve(tlsListener); err != nil {
			println("tls server listen end", err.Error())
		}
	}()
	//go m.graceShutdown(server)
}

func (m *assetsHttpServer) serve(addr string, handler http.Handler) {
	server := &http.Server{Addr: addr, Handler: handler}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			println("server listen end", err.Error())
		}
	}()
	//go m.graceShutdown(server)
}

func (m *assetsHttpServer) graceShutdown(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c
	println("notify signal", s.String())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		println("http server grace shutdown failed error:", err)
	}
	//os.Exit(1)
}

// StartHttpServer 启动内置Http Server
func (m *assetsHttpServer) StartHttpServer() {
	if m.LocalAssets != "" {
		m.LocalAssets = strings.Replace(m.LocalAssets, "\\", "/", -1) // ReplaceAll
		if strings.LastIndex(m.LocalAssets, "/") != len(m.LocalAssets)-1 {
			m.LocalAssets = m.LocalAssets + "/"
		}
	}
	addr := fmt.Sprintf("%s:%d", m.IP, m.PORT)
	mux := http.NewServeMux()
	mux.Handle("/", m)
	if m.SSL != nil {
		m.serveTLS(addr, mux)
	} else {
		m.serve(addr, mux)
	}
}

func (m *assetsHttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if AssetsServerHeaderKeyValue != "" {
		if AssetsServerHeaderKeyValue != r.Header.Get(AssetsServerHeaderKeyName) {
			return
		}
	}
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	var path = r.URL.Path
	if path == "/" {
		path = "/index.html"
	} else if strings.LastIndex(path, "/") == len(path)-1 {
		path = path + "index.html"
	}
	var (
		byt []byte
		err error
	)
	if m.Assets != nil {
		byt, err = m.Assets.ReadFile(m.AssetsFSName + path)
	} else if m.LocalAssets != "" {
		path = fmt.Sprintf("%s%s", m.LocalAssets, path)
		byt, err = ioutil.ReadFile(path)
	} else {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("resource directory is not configured"))
		return
	}
	if err != nil {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("file not found: " + path))
	} else {
		et := extType(path)
		if et != "" {
			if ct, ok := contentType[et]; ok {
				w.Header().Set("Content-Type", ct)
			}
		}
		w.WriteHeader(200)
		_, _ = w.Write(byt)
	}
}

func extType(path string) string {
	idx := strings.LastIndex(path, ".")
	if idx != -1 {
		return path[idx:]
	}
	return ""
}
