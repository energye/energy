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
	"embed"
	"github.com/energye/energy/v2/cef/process"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"unsafe"
)

const (
	localDomain = "energy" // default energy
)

// 本地加载资源
var localLoadRes *LocalLoadResource

// LocalLoadResource 初始化时设置
//  本地&内置加载资源
type LocalLoadResource struct {
	LocalLoadConfig
	mimeType    map[string]string
	sourceCache map[string]*source
}

// LocalLoadConfig
//  本地&内置资源加载配置
//  然后使用 Build() 函数构建对象
type LocalLoadConfig struct {
	Enable  bool                // 设置是否启用本地资源缓存到内存, 默认true: 启用, 禁用时需要调用Disable函数
	Domain  string              // 必须设置的域, 格式为: xxx, xxx.xx, xxx.xxx.xxx， example, example.com
	Scheme  LocalCustomerScheme // 自定义协议, local: 本地磁盘目录加载, fs: 内置到执行程序加载, 默认 fs
	exePath string              // 执行文件当前目录
	// 资源根目录, scheme是local时为本地目录(默认当前程序执行目录), scheme是fs时为默认值resources
	// 规则 local: 空(""): 时当前目录, @: 当前目录开始(@/to/path)，或绝对目录
	ResRootDir string
	FS         *embed.FS // 内置加载资源对象, scheme是fs时必须填入
	Proxy      IXHRProxy // 数据请求代理, 在浏览器发送xhr请求时可通过该配置转发, 你可自定义实现该 IXHRProxy 接口
	Home       string    // 默认首页: /index.html, /home.html, /other.html, default: /index.html
}

// 请求和响应资源
type source struct {
	path         string              // 资源路径, 根据请求URL地址
	fileExt      string              // 资源扩展名, 用于拿到 MimeType
	bytes        []byte              // 资源数据
	err          error               // 获取资源时的错误
	readPosition int                 // 读取资源时的地址偏移
	status       int32               // 响应状态码
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
//  构建本地资源加载配置
//  初始化默认值和默认代理配置
func (m LocalLoadConfig) Build() *LocalLoadConfig {
	if localLoadRes != nil && !process.Args.IsMain() {
		return nil
	}
	var config = &m
	// domain 必须设置
	if config.Domain == "" {
		config.Domain = localDomain
	}
	// scheme 必须是 file 或 fs
	if config.Scheme != LcsLocal && config.Scheme != LcsFS {
		config.Scheme = LcsFS
	}
	// 默认使用 /index.html
	if config.Home == "" {
		config.Home = "/index.html"
	} else if config.Home[0] != '/' {
		config.Home = "/" + config.Home
	}
	m.exePath, _ = os.Getwd()
	// 默认的资源目录
	if config.ResRootDir == "" {
		if config.Scheme == LcsFS {
			config.ResRootDir = "resources"
		} else if config.Scheme == LcsLocal {
			config.ResRootDir = m.exePath
		}
	}
	if m.Proxy != nil {
		if proxy, ok := m.Proxy.(*XHRProxy); ok {
			proxy.init()
		}
	}
	return config
}

// Disable
//  如果不想启用该代理配置，需要主动调用该函数，仅在应用出始化时有效
func (m *LocalLoadConfig) Disable() *LocalLoadConfig {
	m.Enable = false
	return m
}

func (m *LocalLoadResource) enable() bool {
	if m == nil {
		return false
	}
	return m.LocalLoadConfig.Enable
}

// getMimeType
//  获取资源mime type
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
//  所支持的scheme, domain
//  URL格式, fs://energy/index.html, 文件路径必须包含扩展名
//  这里返回false后不会创建资源处理对象
func (m *LocalLoadResource) checkRequest(request *ICefRequest) (*source, bool) {
	rt := request.ResourceType()
	// 根据资源类型跳过哪些资源不被本地加载
	// TODO: rt_media 类型应该在此去除
	switch rt {
	case RT_MEDIA, RT_PING, RT_CSP_REPORT, RT_PLUGIN_RESOURCE:
		return nil, false
	}
	reqUrl, err := url.Parse(request.URL())
	logger.Debug("LocalLoadResource URL:", reqUrl.String(), "RT:", rt)
	if err != nil {
		logger.Error("LocalLoadResource, scheme invalid should:", LcsFS, "or", LcsLocal)
		return nil, false
	}
	if reqUrl.Scheme != string(m.Scheme) {
		logger.Error("LocalLoadResource, scheme invalid should:", LcsFS, "or", LcsLocal, "current:", reqUrl.Scheme)
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
	if localLoadRes.Scheme == LcsLocal {
		var path string
		if localLoadRes.ResRootDir[0] == '@' {
			//当前路径
			path = filepath.Join(localLoadRes.exePath, localLoadRes.ResRootDir[1:])
		} else {
			//绝对路径
			path = localLoadRes.ResRootDir
		}
		m.bytes, m.err = ioutil.ReadFile(filepath.Join(path, m.path))
		// 在本地读取
		if m.err != nil {
			logger.Error("ReadFile:", m.err.Error())
		}
	} else if localLoadRes.Scheme == LcsFS && localLoadRes.FS != nil {
		//在fs读取
		m.bytes, m.err = localLoadRes.FS.ReadFile(localLoadRes.ResRootDir + m.path)
		if m.err != nil {
			logger.Error("ReadFile:", m.err.Error())
		}
	}
}

// checkRequest = true, 打开资源
func (m *source) open(request *ICefRequest, callback *ICefCallback) (handleRequest, result bool) {
	m.readPosition = 0
	// 当前资源的响应设置默认值
	m.status = 404
	m.statusText = "Not Found"
	m.err = nil
	m.header = nil
	// xhr 请求, 需要通过代理转发出去
	if m.resourceType == RT_XHR && localLoadRes.Proxy != nil {
		if result, err := localLoadRes.Proxy.Send(request); err == nil {
			m.bytes, m.err = result.Data, err
			m.status = result.StatusCode
			m.header = result.Header
		}
	} else {
		// 如果开启缓存,直接在缓存取
		m.readFile()
		if m.err == nil {
			m.status = 200
			m.statusText = "OK"
		} else {
			// 尝试在代理服务请求资源
			if result, err := localLoadRes.Proxy.Send(request); err == nil {
				m.bytes, m.err = result.Data, err
				m.status = result.StatusCode
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
				m.statusText = err.Error()
			}
		}
	}
	callback.Cont()
	return true, true
}

// checkRequest = true, 设置响应信息
func (m *source) response(response *ICefResponse) (responseLength int64, redirectUrl string) {
	response.SetStatus(m.status)
	response.SetStatusText(m.statusText)
	response.SetMimeType(m.mimeType)
	responseLength = int64(len(m.bytes))
	if m.header != nil {
		header := StringMultiMapRef.New()
		if header.IsValid() {
			for key, value := range m.header {
				for _, vs := range value {
					header.Append(key, vs)
				}
			}
			response.SetHeaderMap(header)
		}
	}
	return
}

// checkRequest = true, 读取bytes, 返回到dataOut
func (m *source) read(dataOut uintptr, bytesToRead int32, callback *ICefResourceReadCallback) (bytesRead int32, result bool) {
	if m.bytes != nil && len(m.bytes) > 0 {
		var i int32 = 0 // 默认 0
		for i < bytesToRead && m.readPosition < len(m.bytes) {
			*(*byte)(unsafe.Pointer(dataOut + uintptr(i))) = m.bytes[m.readPosition]
			m.readPosition++
			i++
		}
		// 读取到最后不缓存时,清空
		if i == 0 {
			m.bytes = nil
		}
		callback.Cont(int64(i))
		return i, i > 0
	}
	return
}
