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
	"github.com/cyber-xxm/energy/v2/cef/process"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/energye/golcl/energy/emfs"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"unsafe"
)

const (
	localDomain = "energy"      // 默认本地资源加载域
	localProto  = "fs"          // 默认本地资源加载协议
	localHome   = "/index.html" //
)

// 本地加载资源
var localLoadRes *LocalLoadResource

// LocalLoadResource 初始化时设置
//
//	本地&内置加载资源
type LocalLoadResource struct {
	LocalLoadConfig
	mimeType    map[string]string
	sourceCache map[string]*source
}

// LocalLoadConfig
//
//	本地&内置资源加载配置
//	然后使用 Build() 函数构建对象
type LocalLoadConfig struct {
	Enable bool   // 设置是否启用本地资源缓存到内存, 默认true: 启用, 禁用时需要调用Disable函数
	Domain string // 自定义域, 格式: xxx | xxx.xx | xxx.xxx.xxx， example, example.com, 默认: energy
	Scheme string // 自定义协议, 不建议使用 HTTP、HTTPS、FILE、FTP、ABOUT和DATA 默认: fs
	// 资源根目录, fs为空时: 本地目录(默认当前程序执行目录), fs不为空时: 默认值 resources, 使用内置加载
	// 本地目录规则: 空("")时当前目录, @当前目录开始(@/to/path)，或绝对目录.
	ResRootDir string        //
	FS         emfs.IEmbedFS // 内置加载资源对象, 不为nil时使用内置加载，默认: nil
	Proxy      IXHRProxy     // 数据请求代理, 在浏览器发送xhr请求时可通过该配置转发, 你可自定义实现该 IXHRProxy 接口
	Home       string        // 默认首页HTML文件名: /index.html , 默认: /index.html
	exePath    string        // 执行文件当前目录
}

// 请求和响应资源
type source struct {
	path         string              // 资源路径, 根据请求URL地址
	fileExt      string              // 资源扩展名, 用于拿到 MimeType
	bytes        []byte              // 资源数据
	err          error               // 获取资源时的错误
	start        int                 // 读取资源时的地址偏移
	statusCode   int32               // 响应状态码
	statusText   string              // 响应状态文本
	mimeType     string              // 响应的资源 MimeType
	header       map[string][]string // 响应头
	resourceType TCefResourceType    // 资源类型
}

// 初始化本地加载配置对象
func localLoadResourceInit(config *LocalLoadConfig) {
	if config == nil && !process.Args.IsMain() {
		return
	}
	localLoadRes = &LocalLoadResource{
		mimeType:    make(map[string]string),
		sourceCache: make(map[string]*source),
	}
	localLoadRes.LocalLoadConfig = *config
}

// Build
//
//	构建本地资源加载配置
//	初始化默认值和默认代理配置
func (m LocalLoadConfig) Build() *LocalLoadConfig {
	if localLoadRes != nil && !process.Args.IsMain() {
		return nil
	}
	var config = &m
	config.Enable = true
	// domain 必须设置
	if config.Domain == "" {
		config.Domain = localDomain
	}
	if config.Scheme == "" {
		config.Scheme = localProto
	}
	// 默认使用 /index.html
	if config.Home == "" {
		config.Home = localHome
	} else if config.Home[0] != '/' {
		config.Home = "/" + config.Home
	}
	m.exePath = ExeDir
	// 默认的资源目录
	if config.ResRootDir == "" {
		if config.FS != nil {
			config.ResRootDir = "resources"
		} else {
			config.ResRootDir = m.exePath
		}
	}
	if m.Proxy != nil {
		if proxy, ok := m.Proxy.(*XHRProxy); ok {
			proxy.init()
		}
	}
	if BrowserWindow.Config.Url == "" || BrowserWindow.Config.Url == defaultAboutBlank {
		defaultURL := new(bytes.Buffer)
		defaultURL.WriteString(m.Scheme)
		defaultURL.WriteString("://")
		defaultURL.WriteString(m.Domain)
		if m.Home != localHome {
			defaultURL.WriteString(m.Home)
		}
		BrowserWindow.Config.Url = defaultURL.String()
	}
	return config
}

// Disable
//
//	如果不想启用该代理配置，需要主动调用该函数，仅在应用出始化时有效
func (m *LocalLoadConfig) Disable() *LocalLoadConfig {
	m.Enable = false
	return m
}

func (m *LocalLoadResource) loadDefaultURL(window IBrowserWindow, browser *ICefBrowser) {
	if m.enable() {
		var homeURL string
		if BrowserWindow.Config.Url != defaultAboutBlank {
			homeURL = window.WindowProperty().Url
		} else {
			defaultURL := new(bytes.Buffer)
			defaultURL.WriteString(m.Scheme)
			defaultURL.WriteString("://")
			defaultURL.WriteString(m.Domain)
			if m.Home != localHome {
				defaultURL.WriteString(m.Home)
			}
			homeURL = defaultURL.String()
		}
		logger.Debug("LocalLoadResource Default-URL:", homeURL)
		window.Chromium().LoadUrl(homeURL)
	}
}

// 方式二 资源处理器默认实现，使用本地资源加载时开启
func (m *LocalLoadResource) getSchemeHandlerFactory(window IBrowserWindow, browser *ICefBrowser) {
	if m.enable() {
		handler := SchemeHandlerFactoryRef.New()
		handler.SetNew(func(browser *ICefBrowser, frame *ICefFrame, schemeName string, request *ICefRequest) *ICefResourceHandler {
			return m.getResourceHandler(browser, frame, request)
		})
		browser.GetRequestContext().RegisterSchemeHandlerFactory(m.Scheme, m.Domain, handler)
	}
	return
}

// 方式一 资源处理器默认实现，使用本地资源加载时开启
func (m *LocalLoadResource) getResourceHandler(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) *ICefResourceHandler {
	if m.enable() {
		if source, ok := m.checkRequest(request); ok {
			handler := ResourceHandlerRef.New(browser, frame, m.Scheme, request)
			//handler.Open(source.open)
			handler.ProcessRequest(source.processRequest)
			handler.GetResponseHeaders(source.response)
			//handler.Read(source.read)
			handler.ReadResponse(source.readResponse)
			return handler
		}
	}
	return nil
}

func (m *LocalLoadResource) enable() bool {
	if m == nil {
		return false
	}
	return m.LocalLoadConfig.Enable
}

// getMimeType
//
//	获取资源mime type
func (m *LocalLoadResource) getMimeType(extension string) string {
	if mimeType, ok := m.mimeType[extension]; ok {
		return mimeType
	} else {
		mimeType = GetMimeType(extension)
		m.mimeType[extension] = mimeType
		return mimeType
	}
}

// return xxx.js = js, xxx.html = html
func (m *LocalLoadResource) ext(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}

// 使用本地资源加载时,先验证每个请求的合法性
//
//	所支持的scheme, domain
//	URL格式, fs://energy/index.html, 文件路径必须包含扩展名
//	这里返回false后不会创建资源处理对象
func (m *LocalLoadResource) checkRequest(request *ICefRequest) (*source, bool) {
	rt := request.ResourceType()
	// 根据资源类型跳过哪些资源不被本地加载
	// TODO: rt_media 类型应该在此去除
	switch rt {
	case /*RT_MEDIA,*/ RT_PING, RT_CSP_REPORT, RT_PLUGIN_RESOURCE:
		return nil, false
	}
	reqUrl, err := url.Parse(request.URL())
	logger.Debug("LocalLoadResource URL:", reqUrl.String(), "RT:", rt)
	if err != nil {
		logger.Error("LocalLoadResource, scheme invalid should file")
		return nil, false
	}
	if reqUrl.Scheme != m.Scheme {
		logger.Error("LocalLoadResource, scheme invalid should file", "current:", reqUrl.Scheme)
		return nil, false
	}
	if reqUrl.Host != m.Domain {
		logger.Error("LocalLoadResource, Incorrect protocol domain should: [fs | file]://energy/index.html", "current:", reqUrl.Host)
		return nil, false
	}
	path := reqUrl.Path
	if path == "" || path == "/" {
		path = m.Home
	}
	ext := m.ext(path)
	return &source{path: path, fileExt: ext, mimeType: m.getMimeType(ext), resourceType: rt}, true
}

// 读取本地或内置资源
func (m *source) readFile() {
	// 必须设置文件根目录, scheme是file时, fileRoot为本地文件目录, scheme是fs时, fileRoot为fs的目录名
	if localLoadRes.FS == nil {
		var path string
		if localLoadRes.ResRootDir[0] == '@' {
			//当前路径
			path = filepath.Join(localLoadRes.exePath, localLoadRes.ResRootDir[1:])
		} else {
			//绝对路径
			path = localLoadRes.ResRootDir
		}
		logger.Debug("LocalLoadResource", "ReadFile Local:", m.path)
		m.bytes, m.err = ioutil.ReadFile(filepath.Join(path, m.path))
		// 在本地读取
		if m.err != nil {
			logger.Error("ReadFile:", m.err.Error())
		}
	} else {
		//在fs读取
		logger.Debug("LocalLoadResource", "ReadFile Embed:", m.path)
		m.bytes, m.err = localLoadRes.FS.ReadFile(localLoadRes.ResRootDir + m.path)
		if m.err != nil {
			logger.Error("ReadFile:", m.err.Error())
		}
	}
}

// checkRequest = true, 打开资源
func (m *source) open(request *ICefRequest, callback *ICefCallback) (handleRequest, ok bool) {
	m.start = 0
	// 当前资源的响应设置默认值
	m.statusCode = 404
	m.statusText = "Not Found"
	m.err = nil
	m.header = nil
	logger.Debug("LocalLoadResource", "ResourceType:", m.resourceType)
	// xhr 请求, 需要通过代理转发出去
	if m.resourceType == RT_XHR && localLoadRes.Proxy != nil {
		if result, err := localLoadRes.Proxy.Send(request); err == nil {
			m.bytes, m.err = result.Data, err
			m.statusCode = result.StatusCode
			m.statusText = result.Status
			m.header = result.Header
		} else {
			m.err = err
			m.statusText = err.Error()
		}
	} else {
		m.readFile()
		if m.err == nil {
			m.statusCode = 200
			m.statusText = "OK"
		} else if localLoadRes.Proxy != nil {
			// 尝试在代理服务请求资源
			if result, err := localLoadRes.Proxy.Send(request); err == nil {
				m.bytes, m.err = result.Data, err
				m.statusCode = result.StatusCode
				m.statusText = result.Status
				m.header = result.Header
				// TODO 需要验证 Content-Type 合法性
				if ct, ok := result.Header["Content-Type"]; ok {
					m.mimeType = ct[0]
				} else {
					m.mimeType = "text/html"
				}
			} else {
				m.bytes = []byte("Invalid resource request")
				m.mimeType = "application/json"
				m.err = err
				m.statusText = err.Error()
			}
		}
	}
	callback.Cont()
	return true, true
}

func (m *source) processRequest(request *ICefRequest, callback *ICefCallback) bool {
	_, _ = m.open(request, callback)
	return true
}

// checkRequest = true, 设置响应信息
func (m *source) response(response *ICefResponse) (responseLength int64, redirectUrl string) {
	response.SetStatus(m.statusCode)
	response.SetStatusText(m.statusText)
	response.SetMimeType(m.mimeType)
	responseLength = int64(len(m.bytes))
	logger.Debug("LocalLoadResource", "Response StatusCode:", m.statusCode, "StatusText:", m.statusText, "MimeType:", m.mimeType)
	if m.header != nil {
		header := response.GetHeaderMap() //StringMultiMapRef.New()
		if header.IsValid() {
			for key, value := range m.header {
				for _, vs := range value {
					header.Append(key, vs)
				}
			}
			response.SetHeaderMap(header)
		}
		header.Free()
	}
	return
}

//func (m *source) out(dataOut uintptr, bytesToRead int32) (bytesRead int32, result bool) {
//	if m.bytes != nil {
//		for bytesRead < bytesToRead && m.readPosition < len(m.bytes) {
//			*(*byte)(unsafe.Pointer(dataOut + uintptr(bytesRead))) = m.bytes[m.readPosition]
//			m.readPosition++
//			bytesRead++
//		}
//		return bytesRead, bytesRead > 0
//	}
//	return
//}

func (m *source) out(dataOut uintptr, bytesToRead int32) (bytesRead int32, result bool) {
	dataSize := len(m.bytes)
	// start 当前读取的开始位置
	// bytes 是空(len=0)没有资源数据
	// start大于dataSize资源读取完成
	if m.start < dataSize {
		var min = func(x, y int) int {
			if x < y {
				return x
			}
			return y
		}
		//把dataOut指针初始化Go类型的切片
		//space切片长度和空间, 使用bytes长度和bytesToRead最小的值
		space := min(dataSize, int(bytesToRead))
		dataOutByteSlice := &reflect.SliceHeader{
			Data: dataOut,
			Len:  space,
			Cap:  space,
		}
		dst := *(*[]byte)(unsafe.Pointer(dataOutByteSlice))
		//end 计算当前读取资源数据的结束位置
		end := m.start
		//拿出最小的数据长度做为结束位置
		//bytesToRead当前最大读取数量一搬最大值是固定
		if dataSize < int(bytesToRead) {
			end += dataSize
		} else {
			end += int(bytesToRead)
		}
		//如果结束位置大于bytes长度,我们使用bytes长度
		end = min(end, dataSize)
		//把每次分块读取的资源数据复制到dataOut
		c := copy(dst, m.bytes[m.start:end])
		m.start += c         //设置下次读取资源开始位置
		bytesRead = int32(c) //读取资源读取字节个数
		return bytesRead, bytesRead > 0
	}
	return
}

// checkRequest = true, 读取bytes, 返回到dataOut
func (m *source) read(dataOut uintptr, bytesToRead int32, callback *ICefResourceReadCallback) (bytesRead int32, result bool) {
	bytesRead, result = m.out(dataOut, bytesToRead)
	logger.Debug("LocalLoadResource", "Read BytesRead:", bytesRead, "Result:", result)
	if result {
		//callback.Cont(int64(bytesRead))
	}
	return
}

func (m *source) readResponse(dataOut uintptr, bytesToRead int32, callback *ICefCallback) (bytesRead int32, result bool) {
	bytesRead, result = m.out(dataOut, bytesToRead)
	logger.Debug("LocalLoadResource", "ReadResponse BytesRead:", bytesRead, "Result:", result)
	if !result {
		callback.Cont()
	}
	return
}
