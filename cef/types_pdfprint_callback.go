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

// ICefPdfPrintCallback
type ICefPdfPrintCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// PdfPrintCallbackRef -> ICefPdfPrintCallback
var PdfPrintCallbackRef pdfPrintCallback

type pdfPrintCallback uintptr

func (*pdfPrintCallback) New() *ICefPdfPrintCallback {
	var result uintptr
	imports.Proc(def.CefPdfPrintCallback_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefPdfPrintCallback{instance: unsafe.Pointer(result)}
}

// Instance 实例
func (m *ICefPdfPrintCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPdfPrintCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPdfPrintCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefPdfPrintCallback) OnPdfPrintFinished(fn onPdfPrintFinished) {
	imports.Proc(def.CefPdfPrintCallback_OnPdfPrintFinished).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onPdfPrintFinished func(path string, ok bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case onPdfPrintFinished:
			fn.(onPdfPrintFinished)(api.GoStr(getVal(0)), api.GoBool(getVal(1)))
		default:
			return false
		}
		return true
	})
}
