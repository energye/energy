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
	"os"
	"path/filepath"
	"strings"
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
var localLoadResource = &LocalLoadResource{
	mimeType:    make(map[string]string),
	sourceCache: make(map[string]*source),
	domain:      localDomain,
	scheme:      LocalCSFS,
}

// LocalLoadResource 初始化时设置
//  本地或内置加载资源
//  domain: 自定义域名称，默认energy, 不能为空
//  scheme: 自定义协议，默认fs, 可选[file: 在本地目录加载, fs: 在内置exe加载]
//  fileRoot: 资源加载根目录, scheme是file时为本地目录(默认当前程序执行目录), scheme是fs时为资源文件夹名
//  fs: 内置加载资源对象, scheme是fs时必须填入
type LocalLoadResource struct {
	enable      bool
	enableCache bool
	mimeType    map[string]string
	sourceCache map[string]*source
	domain      string
	scheme      LocalCustomerScheme
	fileRoot    string
	fs          *embed.FS
}

type source struct {
	path         string // file source path
	fileExt      string // file ext
	bytes        []byte // file source byte
	err          error  // file source read byte error
	readPosition int    // file source read byte position
	status       int32  // response code status
	statusText   string // response status text
	mimeType     string // response source mime type
	resourceType consts.TCefResourceType
}

func init() {
	if localLoadResource.enable {
		wd, _ := os.Getwd()
		localLoadResource.fileRoot = wd
	}
}

// SetFileRoot
//  设置资源加载根目录
//  scheme是file时为本地目录(默认当前程序执行目录)
//  scheme是fs时为资源文件夹名
func (m *LocalLoadResource) SetFileRoot(v string) {
	m.fileRoot = v
}

// SetFS
//  设置内置加载资源对象, scheme是fs时必须填入
func (m *LocalLoadResource) SetFS(v *embed.FS) {
	m.fs = v
}

// SetEnableCache
//  设置是否启用本地资源缓存到内存, 默认false: 未启用
//  开启该配置会占用内存
func (m *LocalLoadResource) SetEnableCache(v bool) {
	m.enableCache = v
}

// SetEnable
//  设置是否启用本地资源加载, 默认false: 未启用
//  提示: 启用该功能, 目前无法加载本地媒体, 媒体资源可http方式加载
func (m *LocalLoadResource) SetEnable(v bool) {
	m.enable = v
}

// SetDomain
//  设置本地资源加载自定义域
func (m *LocalLoadResource) SetDomain(domain string) {
	if domain == "" {
		m.domain = localDomain
		return
	}
	m.domain = domain
}

// SetScheme
//  设置本地资源加载自定义协议
func (m *LocalLoadResource) SetScheme(scheme LocalCustomerScheme) {
	if scheme == "" || (scheme != LocalCSFS && scheme != LocalCSFile) {
		m.scheme = LocalCSFS
		return
	}
	m.scheme = scheme
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
	url := request.URL()
	logger.Debug("URL:", url, rt)
	idx := strings.Index(url, ":")
	if idx != -1 {
		scheme := url[:idx]
		if scheme != string(m.scheme) {
			logger.Error("Local load resource, scheme invalid should:", LocalCSFS, "or", LocalCSFile)
			return nil, false
		}
		idx = strings.Index(url, "://")
		if idx == -1 {
			logger.Error("Incorrect protocol domain address should: [fs | file]://energy/index.html")
			return nil, false
		}
		domainPath := url[idx+3:]
		idx = strings.Index(domainPath, "/")
		if idx == -1 {
			logger.Error("Incorrect protocol domain path should: [fs | file]://energy/index.html")
			return nil, false
		}
		domain := domainPath[:idx]
		if domain != m.domain {
			logger.Error("Incorrect protocol domain should: [fs | file]://energy/index.html")
			return nil, false
		}
		idx = strings.Index(domainPath, "/")
		if idx == -1 {
			logger.Error("Incorrect protocol path should: [fs | file]://energy/index.html")
			return nil, false
		}
		path := domainPath[idx:]
		ext := m.ext(path)
		if ext == "" {
			logger.Error("Incorrect resources should: file.[ext](MimeType)")
			return nil, false
		}
		// 如果开启缓存,直接在缓存拿指定地址的source
		if m.enableCache {
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
	return nil, false
}

// 读取文件
func (m *source) readFile() bool {
	// 必须设置文件根目录, scheme是file时, fileRoot为本地文件目录, scheme是fs时, fileRoot为fs的目录名
	if localLoadResource.fileRoot != "" {
		if localLoadResource.scheme == LocalCSFile {
			// 在本地读取
			data, err := ioutil.ReadFile(filepath.Join(localLoadResource.fileRoot, m.path))
			if err == nil {
				m.bytes = data
				return true
			}
			logger.Error("ReadFile:", err.Error())
		} else if localLoadResource.scheme == LocalCSFS {
			//在fs读取
			data, err := localLoadResource.fs.ReadFile(localLoadResource.fileRoot + m.path)
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
	if localLoadResource.enableCache {
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
		if i == 0 && !localLoadResource.enableCache {
			m.bytes = nil
		}
		callback.Cont(int64(i))
		return i, i > 0
	}
	return
}
