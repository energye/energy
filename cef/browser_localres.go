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
	"errors"
	"github.com/energye/cef/cef"
	"github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/tool/exec"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"unsafe"
)

type tSchemeHandlerFactory struct {
	factory cef.IEngSchemeHandlerFactory
}

type source struct {
	path         string                 // 资源路径, 根据请求URL地址
	fileExt      string                 // 资源扩展名, 用于拿到 MimeType
	bytes        []byte                 // 资源数据
	err          error                  // 获取资源时的错误
	start        int                    // 读取资源时的地址偏移
	statusCode   int32                  // 响应状态码
	statusText   string                 // 响应状态文本
	mimeType     string                 // 响应的资源 MimeType
	header       map[string][]string    // 响应头
	resourceType types.TCefResourceType // 资源类型
}

func createSchemeHandlerFactory(browser cef.ICefBrowser) *tSchemeHandlerFactory {
	if application.GApplication == nil || application.GApplication.LocalLoad == nil {
		return nil
	}
	localLoad := application.GApplication.LocalLoad
	m := &tSchemeHandlerFactory{}
	m.factory = cef.NewEngSchemeHandlerFactory(0)
	m.factory.SetOnSchemeFactoryNew(m.schemeHandlerFactoryOnSchemeFactoryNew)
	logger.Debug("Chromium.OnAfterCreated > createSchemeHandlerFactory Scheme:", localLoad.Scheme, "Domain:", localLoad.Domain, "factory-IsValid:", m.factory.IsValid())
	intf := cef.AsEngSchemeHandlerFactory(m.factory.AsIntfSchemeHandlerFactory())
	ok := browser.GetHost().GetRequestContext().RegisterSchemeHandlerFactory(localLoad.Scheme, localLoad.Domain, intf)
	logger.Debug("Chromium.OnAfterCreated > createSchemeHandlerFactory RegisterSchemeHandlerFactory:", ok)
	intf.Release()
	return m
}

func (m *tSchemeHandlerFactory) schemeHandlerFactoryOnSchemeFactoryNew(browser cef.ICefBrowser, frame cef.ICefFrame, schemeName string, request cef.ICefRequest) cef.IEngResourceHandler {
	logger.Debug("SchemeHandlerFactory.OnNew schemeName:", schemeName)
	src, err := makeSource(schemeName, request)
	if err != nil {
		logger.Error("SchemeHandlerFactory.OnNew", err.Error())
		return nil
	}
	resourceHandler := cef.NewEngResourceHandler(browser, frame, schemeName, request)
	resourceHandler.SetOnResourceProcessRequest(src.resourceHandlerOnResourceProcessRequest)
	resourceHandler.SetOnResourceGetResponseHeaders(src.resourceHandlerOnResourceGetResponseHeaders)
	resourceHandler.SetOnResourceReadResponse(src.resourceHandlerOnResourceReadResponse)
	resourceHandler.SetOnResourceRead(src.resourceHandlerOnResourceRead)
	resourceHandler = cef.AsEngResourceHandler(resourceHandler.AsIntfResourceHandler())
	return resourceHandler
}

func (m *source) resourceHandlerOnResourceProcessRequest(request cef.ICefRequest, callback cef.ICefCallback) bool {
	logger.Debug("ResourceHandler.OnResourceProcessRequest")
	localLoad := application.GApplication.LocalLoad
	if m.resourceType == types.RT_XHR && localLoad.Proxy != nil {
		if result, err := localLoad.Proxy.Send(request.GetUrl()); err == nil {
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
		} else {
			logger.Error("ResourceHandler.OnResourceProcessRequest", m.err.Error())
			m.statusText = "Invalid resource request"
			m.bytes = []byte(m.statusText)
			m.mimeType = "application/json"
			m.statusCode = 404
		}
	}
	callback.Cont()
	return true
}

func (m *source) resourceHandlerOnResourceGetResponseHeaders(response cef.ICefResponse, outResponseLength *int64, outRedirectUrl *string) {
	logger.Debug("ResourceHandler.OnResourceGetResponseHeaders statusCode:", m.statusCode, "statusText:", m.statusText, "mimeType:", m.mimeType, "dataLen:", len(m.bytes))
	response.SetStatus(m.statusCode)
	response.SetStatusText(m.statusText)
	response.SetMimeType(m.mimeType)
	*outResponseLength = int64(len(m.bytes))
	//if m.header != nil {
	//	header := cef.NewCustomStringMultimap()
	//	intfHeader := cef.AsCefCustomStringMultimap(header.AsIntfStringMultimap())
	//	response.GetHeaderMap(intfHeader)
	//	for key, value := range m.header {
	//		for _, vs := range value {
	//			header.Append(key, vs)
	//		}
	//	}
	//	response.SetHeaderMap(intfHeader)
	//	intfHeader.Release()
	//	header.Free()
	//}
}

func (m *source) resourceHandlerOnResourceReadResponse(dataOut uintptr, bytesToRead int32, bytesRead *int32, callback cef.ICefCallback) bool {
	logger.Debug("ResourceHandler.OnResourceReadResponse > begin.response", "readStart:", m.start, "bytesToRead:", bytesToRead)
	result := m.response(dataOut, bytesToRead, bytesRead)
	if !result {
		callback.Cont()
	}
	logger.Debug("ResourceHandler.OnResourceReadResponse > end.response", "nextReadStart:", m.start, "bytesToRead:", bytesToRead, "bytesRead:", *bytesRead)
	return result
}

func (m *source) resourceHandlerOnResourceRead(dataOut uintptr, bytesToRead int32, bytesRead *int32, callback cef.ICefResourceReadCallback) bool {
	logger.Debug("ResourceHandler.SetOnResourceRead > begin.response", "readStart:", m.start, "bytesToRead:", bytesToRead)
	result := m.response(dataOut, bytesToRead, bytesRead)
	if !result {
		callback.Cont(int64(*bytesRead))
	}
	logger.Debug("ResourceHandler.SetOnResourceRead > end.response", "nextReadStart:", m.start, "bytesToRead:", bytesToRead, "bytesRead:", *bytesRead)
	return true
}

func (m *source) response(dataOut uintptr, bytesToRead int32, bytesRead *int32) bool {
	dataSize := len(m.bytes)
	result := false
	if m.start < dataSize {
		var min = func(x, y int) int {
			if x < y {
				return x
			}
			return y
		}
		space := min(dataSize, int(bytesToRead))
		dataOutByteSlice := &reflect.SliceHeader{
			Data: dataOut,
			Len:  space,
			Cap:  space,
		}
		dst := *(*[]byte)(unsafe.Pointer(dataOutByteSlice))
		end := m.start
		if dataSize < int(bytesToRead) {
			end += dataSize
		} else {
			end += int(bytesToRead)
		}
		end = min(end, dataSize)
		c := copy(dst, m.bytes[m.start:end])
		m.start += c
		*bytesRead = int32(c)
		result = c > 0
	}
	return result
}

func (m *source) readFile() {
	localLoad := application.GApplication.LocalLoad
	if localLoad.FS == nil {
		var path string
		if localLoad.ResRootDir != "" && localLoad.ResRootDir[0] == '@' {
			path = filepath.Join(exec.AppDir(), localLoad.ResRootDir[1:])
		} else {
			path = localLoad.ResRootDir
		}
		logger.Debug("ResourceHandler.OnResourceProcessRequest Local ReadFile:", m.path)
		m.bytes, m.err = os.ReadFile(filepath.Join(path, m.path))
	} else {
		logger.Debug("ResourceHandler.OnResourceProcessRequest Embed ReadFile:", m.path)
		m.bytes, m.err = localLoad.FS.ReadFile(localLoad.ResRootDir + m.path)
	}
}

func makeSource(schemeName string, request cef.ICefRequest) (*source, error) {
	rt := request.GetResourceType()
	switch rt {
	case /*RT_MEDIA,*/ types.RT_PING, types.RT_CSP_REPORT, types.RT_PLUGIN_RESOURCE:
		return nil, errors.New("unsupported resource loading type")
	}
	targetURL := request.GetUrl()
	reqUrl, err := url.Parse(targetURL)
	if err != nil {
		return nil, errors.New("invalid URL: " + targetURL)
	}
	if reqUrl.Scheme != schemeName {
		return nil, errors.New("unsupported scheme: " + reqUrl.Scheme)
	}
	path := reqUrl.Path
	ext := ""
	mimeType := ""
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			ext = path[i+1:]
			break
		}
	}
	if ext != "" {
		mimeType = cef.MiscFunc.CefGetMimeType(ext)
	}
	m := &source{start: 0, statusCode: 404, statusText: "Not Found", err: nil, header: nil,
		path: path, fileExt: ext, mimeType: mimeType, resourceType: rt}
	return m, nil
}
