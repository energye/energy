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
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ************************** creates ************************** //

// LifeSpanHandlerRef -> ICefLifeSpanHandler
var LifeSpanHandlerRef lifeSpanHandler

type lifeSpanHandler uintptr

func (*lifeSpanHandler) New() *ICefLifeSpanHandler {
	var result uintptr
	imports.Proc(internale_CefLifeSpanHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLifeSpanHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// ************************** impl ************************** //

// Instance 实例
func (m *ICefLifeSpanHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefLifeSpanHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefLifeSpanHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefLifeSpanHandler) SetOnBeforePopup(fn onBeforePopup) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefLifeSpanHandler_OnBeforePopup).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLifeSpanHandler) SetOnAfterCreated(fn onAfterCreated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefLifeSpanHandler_OnAfterCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLifeSpanHandler) SetDoClose(fn doClose) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefLifeSpanHandler_DoClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLifeSpanHandler) SetOnBeforeClose(fn onBeforeClose) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefLifeSpanHandler_OnBeforeClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onBeforePopup func(browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool
type onAfterCreated func(browser *ICefBrowser)
type doClose func(browser *ICefBrowser) bool
type onBeforeClose func(browser *ICefBrowser)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onBeforePopup:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			beforePInfoPtr := (*beforePopupInfoPtr)(getPtr(2))
			beforePopupInfo := &BeforePopupInfo{
				TargetUrl:         api.GoStr(beforePInfoPtr.TargetUrl),
				TargetFrameName:   api.GoStr(beforePInfoPtr.TargetFrameName),
				TargetDisposition: consts.TCefWindowOpenDisposition(beforePInfoPtr.TargetDisposition),
				UserGesture:       api.GoBool(beforePInfoPtr.UserGesture),
			}
			var (
				//windowInfo = getPtr(3) // not use
				resultClientPtr = (*uintptr)(getPtr(4))
				client          = &ICefClient{}
				//setting	=  getPtr(5)
				//extra_info =  getPtr(6)
				noJavascriptAccess = (*bool)(getPtr(7))
				result             = (*bool)(getPtr(8))
			)
			//callback
			*result = fn.(onBeforePopup)(browse, frame, beforePopupInfo, client, noJavascriptAccess)
			if client.Instance() != 0 {
				*resultClientPtr = client.Instance()
			}
		case onAfterCreated:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onAfterCreated)(browse)
		case doClose:
			browse := &ICefBrowser{instance: getPtr(0)}
			result := (*bool)(getPtr(1))
			*result = fn.(doClose)(browse)
		case onBeforeClose:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onBeforeClose)(browse)
		default:
			return false
		}
		return true
	})
}
