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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefStreamWriter
//
//	/include/capi/cef_stream_capi.h (cef_stream_writer_t)
type ICefStreamWriter struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// StreamWriterRef -> ICefStreamWriter
var StreamWriterRef streamWriter

type streamWriter uintptr

func (*streamWriter) UnWrap(data *ICefStreamWriter) *ICefStreamWriter {
	var result uintptr
	imports.Proc(def.CefStreamWriterRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data

}

func (*streamWriter) NewForFile(filename string) *ICefStreamWriter {
	var result uintptr
	imports.Proc(def.CefStreamWriterRef_CreateForFile).Call(api.PascalStr(filename), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefStreamWriter{instance: getInstance(result)}
	}
	return nil
}

//func (*streamWriter) NewForHandler(filename string) *ICefStreamWriter {
//	var result uintptr
//	imports.Proc(CefStreamWriterRef_CreateForHandler).Call(api.PascalStr(filename), uintptr(unsafe.Pointer(&result)))
//	if result != 0 {
//		return &ICefStreamWriter{NewBaseRefCounted(result)}
//	}
//	return nil
//}

func (m *ICefStreamWriter) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefStreamWriter) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefStreamWriter) Write(data []byte, size, n uint32) uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefStreamWriter_Write).Call(m.Instance(), uintptr(unsafe.Pointer(&data[0])), uintptr(size), uintptr(n))
	return uint32(r1)
}

func (m *ICefStreamWriter) Seek(offset int64, whence int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefStreamWriter_Seek).Call(m.Instance(), uintptr(offset), uintptr(whence))
	return int32(r1)
}

func (m *ICefStreamWriter) Tell() (result int64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.CefStreamWriter_Tell).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefStreamWriter) Flush() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefStreamWriter_Flush).Call(m.Instance())
	return int32(r1)
}

func (m *ICefStreamWriter) MayBlock() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefStreamWriter_MayBlock).Call(m.Instance())
	return api.GoBool(r1)
}
