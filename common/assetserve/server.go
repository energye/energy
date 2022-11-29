package assetserve

import (
	"crypto/tls"
	"embed"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

var contentType = map[string]string{}

//用于保证链接的安全的请求头(header)key名称
var AssetsServerHeaderKeyName = "ASSETS_SERVER_KEY"

//用于保证链接的安全的key值
//
//这里简单的在请求所有资源时增加请求头的判断
//
//不为空时生效
var AssetsServerHeaderKeyValue string

type assetsHttpServer struct {
	LocalAssets  string    //本地静态资源目录 示例: /app/assets/   http://127.0.0.1:8888/demo/demo.html -> /app/assets/demo/demo.html
	AssetsFSName string    //静态资源内置FS目录名 默认值: resources
	Assets       *embed.FS //静态资源内置FS目录对象
	IP           string    //默认值: 127.0.0.1
	PORT         int       //默认值: 80
	SSL          *SSL      //设置后启动https
}

//证书配置，根据 Assets 或 LocalAssets 寻找证书文件位置
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

func NewAssetsHttpServer() *assetsHttpServer {
	return &assetsHttpServer{
		AssetsFSName: "resources",
		IP:           "127.0.0.1",
		PORT:         80,
	}
}

func (m *assetsHttpServer) serveTLS(addr string, handler http.Handler) error {
	server := &http.Server{Addr: addr, Handler: handler}
	if addr == "" {
		addr = ":https"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
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
				certPEMBlock, err = os.ReadFile(m.LocalAssets + certFile)
				if err != nil {
					return tls.Certificate{}, err
				}
				keyPEMBlock, err = os.ReadFile(m.LocalAssets + keyFile)
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
		var err error
		config.Certificates = make([]tls.Certificate, 1)
		config.Certificates[0], err = loadX509KeyPair(m.SSL.SSLCert, m.SSL.SSLKey)
		if err != nil {
			return err
		}
	}
	tlsListener := tls.NewListener(ln, config)
	return server.Serve(tlsListener)
}

func (m *assetsHttpServer) StartHttpServer() error {
	if m.LocalAssets != "" {
		m.LocalAssets = strings.ReplaceAll(m.LocalAssets, "\\", "/")
		if strings.LastIndex(m.LocalAssets, "/") != len(m.LocalAssets)-1 {
			m.LocalAssets = m.LocalAssets + "/"
		}
	}
	addr := fmt.Sprintf("%s:%d", m.IP, m.PORT)
	mux := http.NewServeMux()
	mux.Handle("/", m)
	var err error
	if m.SSL != nil {
		err = m.serveTLS(addr, mux)
	} else {
		err = http.ListenAndServe(addr, mux)
	}
	return err
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
