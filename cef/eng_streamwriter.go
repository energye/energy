//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefStreamWriter Parent: ICefBaseRefCounted
//
//	Interface used to write data to a stream. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_writer_t))
type ICefStreamWriter interface {
	ICefBaseRefCounted
	// Write
	//  Write raw binary data.
	Write(ptr uintptr, size, n NativeUInt) NativeUInt // function
	// Seek
	//  Seek to the specified offset position. |whence| may be any one of SEEK_CUR, SEEK_END or SEEK_SET. Returns zero on success and non-zero on failure.
	Seek(offset int64, whence int32) int32 // function
	// Tell
	//  Return the current offset position.
	Tell() (resultInt64 int64) // function
	// Flush
	//  Flush the stream.
	Flush() int32 // function
	// MayBlock
	//  Returns true (1) if this writer performs work like accessing the file system which may block. Used as a hint for determining the thread to access the writer from.
	MayBlock() bool // function
}

// TCefStreamWriter Parent: TCefBaseRefCounted
//
//	Interface used to write data to a stream. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_writer_t))
type TCefStreamWriter struct {
	TCefBaseRefCounted
}

// StreamWriterRef -> ICefStreamWriter
var StreamWriterRef streamWriter

// streamWriter TCefStreamWriter Ref
type streamWriter uintptr

func (m *streamWriter) UnWrap(data uintptr) ICefStreamWriter {
	var resultCefStreamWriter uintptr
	CEF().SysCallN(1392, uintptr(data), uintptr(unsafePointer(&resultCefStreamWriter)))
	return AsCefStreamWriter(resultCefStreamWriter)
}

func (m *streamWriter) CreateForFile(fileName string) ICefStreamWriter {
	var resultCefStreamWriter uintptr
	CEF().SysCallN(1387, PascalStr(fileName), uintptr(unsafePointer(&resultCefStreamWriter)))
	return AsCefStreamWriter(resultCefStreamWriter)
}

func (m *TCefStreamWriter) Write(ptr uintptr, size, n NativeUInt) NativeUInt {
	r1 := CEF().SysCallN(1393, m.Instance(), uintptr(ptr), uintptr(size), uintptr(n))
	return NativeUInt(r1)
}

func (m *TCefStreamWriter) Seek(offset int64, whence int32) int32 {
	r1 := CEF().SysCallN(1390, m.Instance(), uintptr(unsafePointer(&offset)), uintptr(whence))
	return int32(r1)
}

func (m *TCefStreamWriter) Tell() (resultInt64 int64) {
	CEF().SysCallN(1391, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefStreamWriter) Flush() int32 {
	r1 := CEF().SysCallN(1388, m.Instance())
	return int32(r1)
}

func (m *TCefStreamWriter) MayBlock() bool {
	r1 := CEF().SysCallN(1389, m.Instance())
	return GoBool(r1)
}
