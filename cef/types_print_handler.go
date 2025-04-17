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

// ICefPrintHandler
type ICefPrintHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// PrintHandlerRef -> ICefPrintHandler
var PrintHandlerRef printSpanHandler

type printSpanHandler uintptr

func (*printSpanHandler) New() *ICefPrintHandler {
	var result uintptr
	imports.Proc(def.CefPrintHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPrintHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefPrintHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPrintHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPrintHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefPrintHandler) SetOnPrintStart(fn onPrintStart) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintHandler_OnPrintStart).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefPrintHandler) SetOnPrintSettings(fn onPrintSettings) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintHandler_OnPrintSettings).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefPrintHandler) SetOnPrintDialog(fn onPrintDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintHandler_OnPrintDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefPrintHandler) SetOnPrintJob(fn onPrintJob) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintHandler_OnPrintJob).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefPrintHandler) SetOnPrintReset(fn onPrintReset) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintHandler_OnPrintReset).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefPrintHandler) SetGetPDFPaperSize(fn getPDFPaperSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPrintHandler_GetPDFPaperSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onPrintStart func(browser *ICefBrowser)
type onPrintSettings func(browser *ICefBrowser, settings *ICefPrintSettings, getDefaults bool)
type onPrintDialog func(browser *ICefBrowser, hasSelection bool, callback *ICefPrintDialogCallback) bool
type onPrintJob func(browser *ICefBrowser, documentName, PDFFilePath string, callback *ICefPrintJobCallback) bool
type onPrintReset func(browser *ICefBrowser)
type getPDFPaperSize func(browser *ICefBrowser, deviceUnitsPerInch int32) *TCefSize

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onPrintStart:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onPrintStart)(browse)
		case onPrintSettings:
			browse := &ICefBrowser{instance: getPtr(0)}
			settings := &ICefPrintSettings{instance: getPtr(1)}
			getDefaults := api.GoBool(getVal(2))
			fn.(onPrintSettings)(browse, settings, getDefaults)
		case onPrintDialog:
			browse := &ICefBrowser{instance: getPtr(0)}
			hasSelection := api.GoBool(getVal(1))
			callback := &ICefPrintDialogCallback{instance: getPtr(2)}
			result := (*bool)(getPtr(3))
			*result = fn.(onPrintDialog)(browse, hasSelection, callback)
		case onPrintJob:
			browse := &ICefBrowser{instance: getPtr(0)}
			documentName, PDFFilePath := api.GoStr(getVal(1)), api.GoStr(getVal(2))
			callback := &ICefPrintJobCallback{instance: getPtr(3)}
			result := (*bool)(getPtr(4))
			*result = fn.(onPrintJob)(browse, documentName, PDFFilePath, callback)
		case onPrintReset:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onPrintReset)(browse)
		case getPDFPaperSize:
			browse := &ICefBrowser{instance: getPtr(0)}
			deviceUnitsPerInch := int32(getVal(1))
			sizePtr := (*uintptr)(getPtr(2))
			size := fn.(getPDFPaperSize)(browse, deviceUnitsPerInch)
			if size != nil {
				*sizePtr = uintptr(unsafe.Pointer(size))
			}
		default:
			return false
		}
		return true
	})
}
