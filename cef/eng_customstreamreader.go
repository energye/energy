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

// ICefCustomStreamReader Parent: ICefBaseRefCountedOwn
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))
type ICefCustomStreamReader interface {
	ICefBaseRefCountedOwn
}

// TCefCustomStreamReader Parent: TCefBaseRefCountedOwn
//
//	Interface used to read data from a stream. The functions of this interface may be called on any thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_stream_reader_t))
type TCefCustomStreamReader struct {
	TCefBaseRefCountedOwn
}

func NewCefCustomStreamReader(stream IStream, owned bool) ICefCustomStreamReader {
	r1 := CEF().SysCallN(786, GetObjectUintptr(stream), PascalBool(owned))
	return AsCefCustomStreamReader(r1)
}

func NewCefCustomStreamReader1(filename string) ICefCustomStreamReader {
	r1 := CEF().SysCallN(787, PascalStr(filename))
	return AsCefCustomStreamReader(r1)
}
