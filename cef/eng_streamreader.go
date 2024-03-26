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

// ICefStreamReader Parent: ICefBaseRefCounted
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))
type ICefStreamReader interface {
	ICefBaseRefCounted
	// Read
	//  Read raw binary data.
	Read(ptr uintptr, size, n NativeUInt) NativeUInt // function
	// Seek
	//  Seek to the specified offset position. |whence| may be any one of SEEK_CUR, SEEK_END or SEEK_SET. Return zero on success and non-zero on failure.
	Seek(offset int64, whence int32) int32 // function
	// Tell
	//  Return the current offset position.
	Tell() (resultInt64 int64) // function
	// Eof
	//  Return non-zero if at end of file.
	Eof() bool // function
	// MayBlock
	//  Return true (1) if this handler performs work like accessing the file system which may block. Used as a hint for determining the thread to access the handler from.
	MayBlock() bool // function
}

// TCefStreamReader Parent: TCefBaseRefCounted
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))
type TCefStreamReader struct {
	TCefBaseRefCounted
}

// StreamReaderRef -> ICefStreamReader
var StreamReaderRef streamReader

// streamReader TCefStreamReader Ref
type streamReader uintptr

func (m *streamReader) UnWrap(data uintptr) ICefStreamReader {
	var resultCefStreamReader uintptr
	CEF().SysCallN(1386, uintptr(data), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForFile(filename string) ICefStreamReader {
	var resultCefStreamReader uintptr
	CEF().SysCallN(1379, PascalStr(filename), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForCustomStream(stream ICefCustomStreamReader) ICefStreamReader {
	var resultCefStreamReader uintptr
	CEF().SysCallN(1377, GetObjectUintptr(stream), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForStream(stream IStream, owned bool) ICefStreamReader {
	var resultCefStreamReader uintptr
	CEF().SysCallN(1380, GetObjectUintptr(stream), PascalBool(owned), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *streamReader) CreateForData(data uintptr, size NativeUInt) ICefStreamReader {
	var resultCefStreamReader uintptr
	CEF().SysCallN(1378, uintptr(data), uintptr(size), uintptr(unsafePointer(&resultCefStreamReader)))
	return AsCefStreamReader(resultCefStreamReader)
}

func (m *TCefStreamReader) Read(ptr uintptr, size, n NativeUInt) NativeUInt {
	r1 := CEF().SysCallN(1383, m.Instance(), uintptr(ptr), uintptr(size), uintptr(n))
	return NativeUInt(r1)
}

func (m *TCefStreamReader) Seek(offset int64, whence int32) int32 {
	r1 := CEF().SysCallN(1384, m.Instance(), uintptr(unsafePointer(&offset)), uintptr(whence))
	return int32(r1)
}

func (m *TCefStreamReader) Tell() (resultInt64 int64) {
	CEF().SysCallN(1385, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefStreamReader) Eof() bool {
	r1 := CEF().SysCallN(1381, m.Instance())
	return GoBool(r1)
}

func (m *TCefStreamReader) MayBlock() bool {
	r1 := CEF().SysCallN(1382, m.Instance())
	return GoBool(r1)
}
