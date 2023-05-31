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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// StreamWriterRef -> ICefStreamWriter
var StreamWriterRef streamWriter

type streamWriter uintptr

func (*streamWriter) UnWrap(data *ICefStreamWriter) *ICefStreamWriter {
	var result uintptr
	imports.Proc(internale_CefStreamWriterRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		data.instance = unsafe.Pointer(result)
		return data
	}
	return nil
}

func (*streamWriter) NewForFile(filename string) *ICefStreamWriter {
	var result uintptr
	imports.Proc(internale_CefStreamWriterRef_CreateForFile).Call(api.PascalStr(filename), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefStreamWriter{NewBaseRefCounted(result)}
	}
	return nil
}

//func (*streamWriter) NewForHandler(filename string) *ICefStreamWriter {
//	var result uintptr
//	imports.Proc(internale_CefStreamWriterRef_CreateForHandler).Call(api.PascalStr(filename), uintptr(unsafe.Pointer(&result)))
//	if result != 0 {
//		return &ICefStreamWriter{NewBaseRefCounted(result)}
//	}
//	return nil
//}

func (m *ICefStreamWriter) Write(data []byte, size, n uint32) uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefStreamWriter_Write).Call(m.Instance(), uintptr(unsafe.Pointer(&data[0])), uintptr(size), uintptr(n))
	return uint32(r1)
}

func (m *ICefStreamWriter) Seek(offset int64, whence int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefStreamWriter_Seek).Call(m.Instance(), uintptr(offset), uintptr(whence))
	return int32(r1)
}

func (m *ICefStreamWriter) Tell() (result int64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(internale_CefStreamWriter_Tell).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefStreamWriter) Flush() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefStreamWriter_Flush).Call(m.Instance())
	return int32(r1)
}

func (m *ICefStreamWriter) MayBlock() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefStreamWriter_MayBlock).Call(m.Instance())
	return api.GoBool(r1)
}
