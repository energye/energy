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

// ICefDialogHandler
type ICefDialogHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DialogHandlerRef -> ICefDialogHandler
var DialogHandlerRef dialogHandler

type dialogHandler uintptr

func (*dialogHandler) New() *ICefDialogHandler {
	var result uintptr
	imports.Proc(def.CefDialogHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDialogHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

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
	imports.Proc(def.CefDialogHandler_OnFileDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onFileDialog func(browser *ICefBrowser, mode consts.FileDialogMode, title, defaultFilePath string, acceptFilters, acceptExtensions, acceptDescriptions *lcl.TStrings, callback *ICefFileDialogCallback) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onFileDialog:
			browse := &ICefBrowser{instance: getPtr(0)}
			mode := consts.FileDialogMode(getVal(1))
			title := api.GoStr(getVal(2))
			defaultFilePath := api.GoStr(getVal(3))
			acceptFiltersList := lcl.AsStrings(getVal(4))
			acceptExtensions := lcl.AsStrings(getVal(5))
			acceptDescriptions := lcl.AsStrings(getVal(6))
			callback := &ICefFileDialogCallback{instance: getPtr(7)}
			result := (*bool)(getPtr(7))
			*result = fn.(onFileDialog)(browse, mode, title, defaultFilePath, acceptFiltersList, acceptExtensions, acceptDescriptions, callback)
		default:
			return false
		}
		return true
	})
}
