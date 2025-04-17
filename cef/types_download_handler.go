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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefDownloadHandler
type ICefDownloadHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DownloadHandlerRef -> ICefDownloadHandler
var DownloadHandlerRef downloadHandler

type downloadHandler uintptr

func (*downloadHandler) New() *ICefDownloadHandler {
	var result uintptr
	imports.Proc(def.CefDownloadHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDownloadHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefDownloadHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDownloadHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDownloadHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDownloadHandler) SetCanDownload(fn canDownload) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDownloadHandler_CanDownload).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDownloadHandler) SetOnBeforeDownload(fn onBeforeDownload) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDownloadHandler_OnBeforeDownload).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDownloadHandler) SetOnDownloadUpdated(fn onDownloadUpdated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDownloadHandler_OnDownloadUpdated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type canDownload func(browser *ICefBrowser, url, requestMethod string) bool
type onBeforeDownload func(browser *ICefBrowser, downloadItem *ICefDownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) bool
type onDownloadUpdated func(browser *ICefBrowser, downloadItem *ICefDownloadItem, callback *ICefDownloadItemCallback)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case canDownload:
			browse := &ICefBrowser{instance: getPtr(0)}
			url, requestMethod := api.GoStr(getVal(1)), api.GoStr(getVal(2))
			result := (*bool)(getPtr(3))
			*result = fn.(canDownload)(browse, url, requestMethod)
		case onBeforeDownload:
			browse := &ICefBrowser{instance: getPtr(0)}
			downlItem := &ICefDownloadItem{instance: getPtr(1)}
			suggestedName := api.GoStr(getVal(2))
			callback := &ICefBeforeDownloadCallback{instance: getPtr(3)}
			result := (*bool)(getPtr(4))
			*result = fn.(onBeforeDownload)(browse, downlItem, suggestedName, callback)
		case onDownloadUpdated:
			browse := &ICefBrowser{instance: getPtr(0)}
			downlItem := &ICefDownloadItem{instance: getPtr(1)}
			callback := &ICefDownloadItemCallback{instance: getPtr(2)}
			fn.(onDownloadUpdated)(browse, downlItem, callback)
		default:
			return false
		}
		return true
	})
}
