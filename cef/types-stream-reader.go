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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// StreamReaderRef -> ICefStreamReader
var StreamReaderRef streamReader

type streamReader uintptr

func (*streamReader) UnWrap(data *ICefStreamReader) *ICefStreamReader {
	var result uintptr
	imports.Proc(internale_CefStreamReaderRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		data.instance = unsafe.Pointer(result)
		return data
	}
	return nil
}

func (*streamReader) NewForFile(filename string) *ICefStreamReader {
	var result uintptr
	imports.Proc(internale_CefStreamReaderRef_CreateForFile).Call(api.PascalStr(filename), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefStreamReader{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*streamReader) NewForStream(stream lcl.IStream, owned bool) *ICefStreamReader {
	var result uintptr
	imports.Proc(internale_CefStreamReaderRef_CreateForStream).Call(stream.Instance(), api.PascalBool(owned), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefStreamReader{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*streamReader) NewForData(data []byte) *ICefStreamReader {
	var result uintptr
	imports.Proc(internale_CefStreamReaderRef_CreateForData).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefStreamReader{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefStreamReader) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefStreamReader) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefStreamReader) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefStreamReader) Read(data []byte, size, n uint32) uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefStreamReader_Read).Call(m.Instance(), uintptr(unsafe.Pointer(&data[0])), uintptr(size), uintptr(n))
	return uint32(r1)
}

func (m *ICefStreamReader) Seek(offset int64, whence int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefStreamReader_Seek).Call(m.Instance(), uintptr(offset), uintptr(whence))
	return int32(r1)
}

func (m *ICefStreamReader) Tell() (result int64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(internale_CefStreamReader_Tell).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefStreamReader) Eof() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefStreamReader_Eof).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefStreamReader) MayBlock() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefStreamReader_MayBlock).Call(m.Instance())
	return api.GoBool(r1)
}
