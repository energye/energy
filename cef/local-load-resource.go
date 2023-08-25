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
	"github.com/energye/energy/v2/consts"
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

// LocalCustomerScheme 本地资源加载自定义固定协议
//  file, fs
type LocalCustomerScheme string

const (
	LocalCSFile LocalCustomerScheme = "file" // 本地目录 file://energy/index.html
	LocalCSFS   LocalCustomerScheme = "fs"   // 内置 fs://energy/index.html
)

// 本地加载资源
var localLoadRes *LocalLoadResource

// LocalLoadResource 初始化时设置
//  本地或内置加载资源
//  domain: 自定义域名称，默认energy, 不能为空
//  scheme: 自定义协议，默认fs, 可选[file: 在本地目录加载, fs: 在内置exe加载]
//  fileRoot: 资源加载根目录, scheme是file时为本地目录(默认当前程序执行目录), scheme是fs时为资源文件夹名
//  fs: 内置加载资源对象, scheme是fs时必须填入
type LocalLoadResource struct {
	LocalLoadConfig
	mimeType    map[string]string
	sourceCache map[string]*source
}

// LocalLoadConfig
//  本地资源加载配置
//  domain: 自定义域名称，默认energy, 不能为空
//  scheme: 自定义协议，默认fs, 可选[file: 在本地目录加载, fs: 在内置exe加载]
//  fileRoot: 资源加载根目录, scheme是file时为本地目录(默认当前程序执行目录), scheme是fs时为资源文件夹名
//  fs: 内置加载资源对象, scheme是fs时必须填入
//  proxy: 数据请求代理
type LocalLoadConfig struct {
	Enable      bool                // 设置是否启用本地资源缓存到内存, 默认false: 未启用, 提示: 启用该功能, 目前无法加载本地媒体, 媒体资源可http方式加载
	EnableCache bool                // 启用缓存，将加载过的资源存储到内存中
	Domain      string              // 必须设置的域
	Scheme      LocalCustomerScheme // 自定义协议, file: 本地磁盘目录加载, fs: 内置到执行程序加载
	FileRoot    string              // 资源根目录, scheme是file时为本地目录(默认当前程序执行目录) scheme是fs时为资源文件夹名
	FS          *embed.FS           // 内置加载资源对象, scheme是fs时必须填入
	Proxy       XHRProxy            // 数据请求代理, 在浏览器发送xhr请求时可通过该配置转发
}

// XHRProxy
//  数据请求代理
type XHRProxy struct {
	IP   string
	Port int
}

// 请求和响应资源
type source struct {
	path         string                  // 资源路径, 根据请求URL地址
	fileExt      string                  // 资源扩展名, 用于拿到 MimeType
	bytes        []byte                  // 资源数据
	err          error                   // 读取资源的错误
	readPosition int                     // 读取资源时的地址偏移
	status       int32                   // 响应状态码
	statusText   string                  // 响应状态文本
	mimeType     string                  // 响应的资源 MimeType
	resourceType consts.TCefResourceType // 资源类型
}

// 初始化本地加载配置对象
func localLoadResourceInit(config LocalLoadConfig) {
	if localLoadRes != nil {
		return
	}
	localLoadRes = &LocalLoadResource{
		mimeType:    make(map[string]string),
		sourceCache: make(map[string]*source),
	}
	// domain 必须设置
	if config.Domain == "" {
		config.Domain = localDomain
	}
	// scheme 必须是 file 或 fs
	if config.Scheme != LocalCSFile && config.Scheme != LocalCSFS {
		config.Scheme = LocalCSFS
	}
	localLoadRes.LocalLoadConfig = config
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
func (m *LocalLoadResource) checkRequest(request *ICefRequest) (*source, bool) {
	rt := request.ResourceType()
	// 根据资源类型跳过哪些资源不被本地加载
	// TODO: rt_media 类型应该在此去除
	switch rt {
	case consts.RT_XHR, consts.RT_MEDIA, consts.RT_PING, consts.RT_CSP_REPORT, consts.RT_PLUGIN_RESOURCE:
		return nil, false
	}
	reqUrl, err := url.Parse(request.URL())
	logger.Debug("LocalLoadResource URL:", reqUrl.String())
	if err != nil {
		logger.Error("LocalLoadResource, scheme invalid should:", LocalCSFS, "or", LocalCSFile)
		return nil, false
	}
	if reqUrl.Scheme != string(m.Scheme) {
		logger.Error("LocalLoadResource, scheme invalid should:", LocalCSFS, "or", LocalCSFile, "current:", reqUrl.Scheme)
		return nil, false
	}
	if reqUrl.Host != m.Domain {
		logger.Error("LocalLoadResource, Incorrect protocol domain should: [fs | file]://energy/index.html", "current:", reqUrl.Host)
		return nil, false
	}
	path := reqUrl.Path
	ext := m.ext(path)
	if ext == "" {
		logger.Error("LocalLoadResource Incorrect resources should: file.[ext](MimeType)")
		return nil, false
	}
	// 如果开启缓存,直接在缓存拿指定地址的source
	if m.EnableCache {
		if s, ok := m.sourceCache[path]; ok {
			return s, true
		} else {
			s = &source{path: path, fileExt: ext, mimeType: m.getMimeType(ext), resourceType: rt}
			m.sourceCache[path] = s
			return s, true
		}
	}
	return &source{path: path, fileExt: ext, mimeType: m.getMimeType(ext), resourceType: rt}, true
}

// 读取文件
func (m *source) readFile() bool {
	// 必须设置文件根目录, scheme是file时, fileRoot为本地文件目录, scheme是fs时, fileRoot为fs的目录名
	if localLoadRes.FileRoot != "" {
		if localLoadRes.Scheme == LocalCSFile {
			// 在本地读取
			data, err := ioutil.ReadFile(filepath.Join(localLoadRes.FileRoot, m.path))
			if err == nil {
				m.bytes = data
				return true
			}
			logger.Error("ReadFile:", err.Error())
		} else if localLoadRes.Scheme == LocalCSFS && localLoadRes.FS != nil {
			//在fs读取
			data, err := localLoadRes.FS.ReadFile(localLoadRes.FileRoot + m.path)
			if err == nil {
				m.bytes = data
				return true
			}
			logger.Error("ReadFile:", err.Error())
		}
	}
	//失败时,返回404,文件不存在
	return false
}

// 打开资源
func (m *source) open(request *ICefRequest, callback *ICefCallback) (handleRequest, result bool) {
	m.readPosition = 0
	// 当前资源的响应设置默认值
	m.status = 404
	m.statusText = "Not Found"
	// 如果开启缓存,直接在缓存取
	if localLoadRes.EnableCache {
		if m.bytes == nil {
			if !m.readFile() {
				return true, true
			}
		}
	} else {
		if !m.readFile() {
			return true, true
		}
	}
	if m.resourceType == consts.RT_MEDIA {
		m.status = 206
	} else {
		m.status = 200
		m.statusText = "OK"
	}
	callback.Cont()
	return true, true
}

// 设置响应信息
func (m *source) response(response *ICefResponse) (responseLength int64, redirectUrl string) {
	response.SetStatus(m.status)
	response.SetStatusText(m.statusText)
	response.SetMimeType(m.mimeType)
	responseLength = int64(len(m.bytes))
	return
}

// 读取bytes, 返回到dataOut
func (m *source) read(dataOut uintptr, bytesToRead int32, callback *ICefResourceReadCallback) (bytesRead int32, result bool) {
	if m.bytes != nil && len(m.bytes) > 0 {
		var i int32 = 0 // 默认 0
		for i < bytesToRead && m.readPosition < len(m.bytes) {
			*(*byte)(unsafe.Pointer(dataOut + uintptr(i))) = m.bytes[m.readPosition]
			m.readPosition++
			i++
		}
		// 读取到最后不缓存时,清空
		if i == 0 && !localLoadRes.EnableCache {
			m.bytes = nil
		}
		callback.Cont(int64(i))
		return i, i > 0
	}
	return
}
