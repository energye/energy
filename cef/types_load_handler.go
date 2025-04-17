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

// ICefLoadHandler
type ICefLoadHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// LoadHandlerRef -> ICefLoadHandler
var LoadHandlerRef loadHandler

type loadHandler uintptr

func (*loadHandler) New() *ICefLoadHandler {
	var result uintptr
	imports.Proc(def.CefLoadHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLoadHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefLoadHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefLoadHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefLoadHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefLoadHandler) SetOnLoadingStateChange(fn onLoadingStateChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLoadHandler_OnLoadingStateChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLoadHandler) SetOnLoadStart(fn onLoadStart) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLoadHandler_OnLoadStart).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLoadHandler) SetOnLoadEnd(fn onLoadEnd) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLoadHandler_OnLoadEnd).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLoadHandler) SetOnLoadError(fn onLoadError) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLoadHandler_OnLoadError).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onLoadingStateChange func(browser *ICefBrowser, isLoading, canGoBack, canGoForward bool)
type onLoadStart func(browser *ICefBrowser, frame *ICefFrame, transitionType consts.TCefTransitionType)
type onLoadEnd func(browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32)
type onLoadError func(browser *ICefBrowser, frame *ICefFrame, errorCode consts.TCefErrorCode, errorText, failedUrl string)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onLoadingStateChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			isLoading, canGoBack, canGoForward := api.GoBool(getVal(1)), api.GoBool(getVal(2)), api.GoBool(getVal(3))
			fn.(onLoadingStateChange)(browse, isLoading, canGoBack, canGoForward)
		case onLoadStart:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			transitionType := consts.TCefTransitionType(getVal(2))
			fn.(onLoadStart)(browse, frame, transitionType)
		case onLoadEnd:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			httpStatusCode := int32(getVal(2))
			fn.(onLoadEnd)(browse, frame, httpStatusCode)
		case onLoadError:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			errorCode := consts.TCefErrorCode(getVal(2))
			errorText, failedUrl := api.GoStr(3), api.GoStr(4)
			fn.(onLoadError)(browse, frame, errorCode, errorText, failedUrl)
		default:
			return false
		}
		return true
	})
}
