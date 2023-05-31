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
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ************************** creates ************************** //

// ResourceRequestHandlerRef -> ICefResourceRequestHandler
var ResourceRequestHandlerRef resourceRequestHandler

type resourceRequestHandler uintptr

func (*resourceRequestHandler) New() *ICefResourceRequestHandler {
	var result uintptr
	imports.Proc(internale_ResourceRequestHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefResourceRequestHandler{instance: unsafe.Pointer(result), ct: consts.CtTClient}
	}
	return nil
}

func (*resourceRequestHandler) NewForChromium(chromium IChromium) *ICefResourceRequestHandler {
	var result uintptr
	imports.Proc(internale_ResourceRequestHandlerRef_CreateForChromium).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefResourceRequestHandler{instance: unsafe.Pointer(result), ct: consts.CtChromium}
	}
	return nil
}

// ************************** impl ************************** //

// Instance 实例
func (m *ICefResourceRequestHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefResourceRequestHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefResourceRequestHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefResourceRequestHandler) IsTClientEvent() bool {
	return m.ct == consts.CtTClient
}

func (m *ICefResourceRequestHandler) IsChromiumEvent() bool {
	return m.ct == consts.CtChromium
}

func (m *ICefResourceRequestHandler) SetGetCookieAccessFilter(fn getCookieAccessFilter) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_GetCookieAccessFilter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceRequestHandler) SetOnBeforeResourceLoad(fn onBeforeResourceLoad) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_OnBeforeResourceLoad).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceRequestHandler) SetGetResourceHandler(fn getResourceHandler) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_GetResourceHandler).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceRequestHandler) SetOnResourceRedirect(fn onResourceRedirect) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_OnResourceRedirect).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceRequestHandler) SetOnResourceResponse(fn onResourceResponse) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_OnResourceResponse).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceRequestHandler) SetGetResourceResponseFilter(fn getResourceResponseFilter) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_GetResourceResponseFilter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceRequestHandler) SetOnResourceLoadComplete(fn onResourceLoadComplete) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_OnResourceLoadComplete).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceRequestHandler) SetOnProtocolExecution(fn onProtocolExecution) {
	if !m.IsValid() || m.IsChromiumEvent() {
		return
	}
	imports.Proc(internale_ResourceRequestHandler_OnProtocolExecution).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type getCookieAccessFilter func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (filter *ICefCookieAccessFilter)
type onBeforeResourceLoad func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback) consts.TCefReturnValue
type getResourceHandler func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (resourceHandler *ICefResourceHandler)
type onResourceRedirect func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) (newUrl string)
type onResourceResponse func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) bool
type getResourceResponseFilter func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) (responseFilter *ICefResponseFilter)
type onResourceLoadComplete func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, status consts.TCefUrlRequestStatus, receivedContentLength int64)
type onProtocolExecution func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (allowOsExecution bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case getCookieAccessFilter:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			filterPtr := (*uintptr)(getPtr(3))
			filter := fn.(getCookieAccessFilter)(browse, frame, request)
			if filter != nil && filter.IsValid() {
				*filterPtr = filter.Instance()
			}
		case onBeforeResourceLoad:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			callback := &ICefCallback{instance: getPtr(3)}
			returnValuePtr := (*consts.TCefReturnValue)(getPtr(4))
			returnValue := fn.(onBeforeResourceLoad)(browse, frame, request, callback)
			*returnValuePtr = returnValue
		case getResourceHandler:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			resourceHandlerPtr := (*uintptr)(getPtr(3))
			resourceHandler := fn.(getResourceHandler)(browse, frame, request)
			if resourceHandler != nil && resourceHandler.IsValid() {
				*resourceHandlerPtr = resourceHandler.Instance()
			} else {
				*resourceHandlerPtr = 0
			}
		case onResourceRedirect:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			response := &ICefResponse{instance: getPtr(3)}
			newUrlPtr := (*uintptr)(getPtr(4))
			newUrl := fn.(onResourceRedirect)(browse, frame, request, response)
			*newUrlPtr = api.PascalStr(newUrl)
		case onResourceResponse:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			response := &ICefResponse{instance: getPtr(3)}
			result := (*bool)(getPtr(4))
			*result = fn.(onResourceResponse)(browse, frame, request, response)
		case getResourceResponseFilter:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			response := &ICefResponse{instance: getPtr(3)}
			responseFilterPtr := (*uintptr)(getPtr(4))
			responseFilter := fn.(getResourceResponseFilter)(browse, frame, request, response)
			if responseFilter != nil && responseFilter.IsValid() {
				*responseFilterPtr = responseFilter.Instance()
			} else {
				*responseFilterPtr = 0
			}
		case onResourceLoadComplete:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			status := consts.TCefUrlRequestStatus(getVal(3))
			receivedContentLength := *(*int64)(getPtr(4))
			fn.(onResourceLoadComplete)(browse, frame, request, status, receivedContentLength)
		case onProtocolExecution:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			allowOsExecution := (*bool)(getPtr(3))
			*allowOsExecution = fn.(onProtocolExecution)(browse, frame, request)
		default:
			return false
		}
		return true
	})
}
