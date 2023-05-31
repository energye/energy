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

// DialogHandlerRef -> ICefDialogHandler
var DialogHandlerRef dialogHandler

type dialogHandler uintptr

func (*dialogHandler) New() *ICefDialogHandler {
	var result uintptr
	imports.Proc(internale_CefDialogHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDialogHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// ************************** impl ************************** //

// Instance 实例
func (m *ICefDialogHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDialogHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDialogHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDialogHandler) SetOnFileDialog(fn onFileDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefDialogHandler_OnFileDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onFileDialog func(browser *ICefBrowser, mode consts.TCefFileDialogMode, title, defaultFilePath string, acceptFilters []string, callback *ICefFileDialogCallback) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onFileDialog:
			browse := &ICefBrowser{instance: getPtr(0)}
			mode := consts.TCefFileDialogMode(getVal(1))
			title := api.GoStr(getVal(2))
			defaultFilePath := api.GoStr(getVal(3))
			acceptFiltersList := lcl.AsStrings(getVal(4))
			callback := &ICefFileDialogCallback{instance: getPtr(5)}
			result := (*bool)(getPtr(6))
			var acceptFilters []string
			if acceptFiltersList.IsValid() {
				count := int(acceptFiltersList.Count())
				acceptFilters = make([]string, count, count)
				for i := 0; i < count; i++ {
					acceptFilters[i] = acceptFiltersList.Strings(int32(i))
				}
				acceptFiltersList.Free()
			}
			*result = fn.(onFileDialog)(browse, mode, title, defaultFilePath, acceptFilters, callback)
		default:
			return false
		}
		return true
	})
}
