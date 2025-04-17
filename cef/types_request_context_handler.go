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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefRequestContextHandler
type ICefRequestContextHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// RequestContextHandlerRef -> ICefRequestContextHandler
var RequestContextHandlerRef requestContextHandler

type requestContextHandler uintptr

func (*requestContextHandler) New() *ICefRequestContextHandler {
	var result uintptr
	imports.Proc(def.RequestContextHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContextHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*requestContextHandler) NewForChromium(chromium IChromium) *ICefRequestContextHandler {
	if chromium == nil || chromium.Instance() == 0 {
		return nil
	}
	var result uintptr
	imports.Proc(def.RequestContextHandlerRef_CreateForChromium).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContextHandler{instance: unsafe.Pointer(result), ct: consts.CtOther}
	}
	return nil
}

// Instance 实例
func (m *ICefRequestContextHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRequestContextHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRequestContextHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRequestContextHandler) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefRequestContextHandler) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

func (m *ICefRequestContextHandler) SetOnRequestContextInitialized(fn onRequestContextInitialized) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.RequestContextHandler_OnRequestContextInitialized).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestContextHandler) SetGetResourceRequestHandler(fn getResourceRequestHandler) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.RequestContextHandler_GetResourceRequestHandler).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onRequestContextInitialized func(requestContext *ICefRequestContext)
type getResourceRequestHandler func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, isNavigation, isDownload bool, requestInitiator string) (disableDefaultHandling bool, resourceRequestHandler *ICefResourceRequestHandler)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onRequestContextInitialized:
			requestContext := &ICefRequestContext{instance: getPtr(0)}
			fn.(onRequestContextInitialized)(requestContext)
		case getResourceRequestHandler:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			isNavigation, isDownload := api.GoBool(getVal(3)), api.GoBool(getVal(4))
			requestInitiator := api.GoStr(getVal(5))
			disableDefaultHandlingPtr := (*bool)(getPtr(6))
			resourceRequestHandlerPtr := (*uintptr)(getPtr(7))
			disableDefaultHandling, resourceRequestHandler := fn.(getResourceRequestHandler)(browse, frame, request, isNavigation, isDownload, requestInitiator)
			*disableDefaultHandlingPtr = disableDefaultHandling
			if resourceRequestHandler != nil && resourceRequestHandler.IsValid() {
				*resourceRequestHandlerPtr = resourceRequestHandler.Instance()
			} else {
				*resourceRequestHandlerPtr = 0
			}
		default:
			return false
		}
		return true
	})
}
