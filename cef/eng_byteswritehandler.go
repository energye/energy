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

// ICefBytesWriteHandler Parent: ICefWriteHandler
type ICefBytesWriteHandler interface {
	ICefWriteHandler
	Write(ptr uintptr, size, n NativeUInt) NativeUInt // function
	Seek(offset int64, whence int32) int32            // function
	Tell() (resultInt64 int64)                        // function
	Flush() int32                                     // function
	MayBlock() bool                                   // function
	GetData() uintptr                                 // function
	GetDataSize() (resultInt64 int64)                 // function
}

// TCefBytesWriteHandler Parent: TCefWriteHandler
type TCefBytesWriteHandler struct {
	TCefWriteHandler
}

func NewCefBytesWriteHandler(aGrow NativeUInt) ICefBytesWriteHandler {
	r1 := CEF().SysCallN(711, uintptr(aGrow))
	return AsCefBytesWriteHandler(r1)
}

func (m *TCefBytesWriteHandler) Write(ptr uintptr, size, n NativeUInt) NativeUInt {
	r1 := CEF().SysCallN(718, m.Instance(), uintptr(ptr), uintptr(size), uintptr(n))
	return NativeUInt(r1)
}

func (m *TCefBytesWriteHandler) Seek(offset int64, whence int32) int32 {
	r1 := CEF().SysCallN(716, m.Instance(), uintptr(unsafePointer(&offset)), uintptr(whence))
	return int32(r1)
}

func (m *TCefBytesWriteHandler) Tell() (resultInt64 int64) {
	CEF().SysCallN(717, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefBytesWriteHandler) Flush() int32 {
	r1 := CEF().SysCallN(712, m.Instance())
	return int32(r1)
}

func (m *TCefBytesWriteHandler) MayBlock() bool {
	r1 := CEF().SysCallN(715, m.Instance())
	return GoBool(r1)
}

func (m *TCefBytesWriteHandler) GetData() uintptr {
	r1 := CEF().SysCallN(713, m.Instance())
	return uintptr(r1)
}

func (m *TCefBytesWriteHandler) GetDataSize() (resultInt64 int64) {
	CEF().SysCallN(714, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func CefBytesWriteHandlerClass() TClass {
	ret := CEF().SysCallN(710)
	return TClass(ret)
}
