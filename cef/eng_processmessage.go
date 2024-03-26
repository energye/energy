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

// ICefProcessMessage Parent: ICefBaseRefCounted
//
//	Interface representing a message. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_process_message_capi.h">CEF source file: /include/capi/cef_process_message_capi.h (cef_process_message_t))
type ICefProcessMessage interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// Copy
	//  Returns a writable copy of this object. Returns nullptr when message contains a shared memory region.
	Copy() ICefProcessMessage // function
	// GetName
	//  Returns the message name.
	GetName() string // function
	// GetArgumentList
	//  Returns the list of arguments. Returns nullptr when message contains a shared memory region.
	GetArgumentList() ICefListValue // function
	// GetSharedMemoryRegion
	//  Returns the shared memory region. Returns nullptr when message contains an argument list.
	GetSharedMemoryRegion() ICefSharedMemoryRegion // function
}

// TCefProcessMessage Parent: TCefBaseRefCounted
//
//	Interface representing a message. Can be used on any process and thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_process_message_capi.h">CEF source file: /include/capi/cef_process_message_capi.h (cef_process_message_t))
type TCefProcessMessage struct {
	TCefBaseRefCounted
}

// ProcessMessageRef -> ICefProcessMessage
var ProcessMessageRef processMessage

// processMessage TCefProcessMessage Ref
type processMessage uintptr

func (m *processMessage) UnWrap(data uintptr) ICefProcessMessage {
	var resultCefProcessMessage uintptr
	CEF().SysCallN(1248, uintptr(data), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}

func (m *processMessage) New(name string) ICefProcessMessage {
	var resultCefProcessMessage uintptr
	CEF().SysCallN(1247, PascalStr(name), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}

func (m *TCefProcessMessage) IsValid() bool {
	r1 := CEF().SysCallN(1246, m.Instance())
	return GoBool(r1)
}

func (m *TCefProcessMessage) IsReadOnly() bool {
	r1 := CEF().SysCallN(1245, m.Instance())
	return GoBool(r1)
}

func (m *TCefProcessMessage) Copy() ICefProcessMessage {
	var resultCefProcessMessage uintptr
	CEF().SysCallN(1241, m.Instance(), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}

func (m *TCefProcessMessage) GetName() string {
	r1 := CEF().SysCallN(1243, m.Instance())
	return GoStr(r1)
}

func (m *TCefProcessMessage) GetArgumentList() ICefListValue {
	var resultCefListValue uintptr
	CEF().SysCallN(1242, m.Instance(), uintptr(unsafePointer(&resultCefListValue)))
	return AsCefListValue(resultCefListValue)
}

func (m *TCefProcessMessage) GetSharedMemoryRegion() ICefSharedMemoryRegion {
	var resultCefSharedMemoryRegion uintptr
	CEF().SysCallN(1244, m.Instance(), uintptr(unsafePointer(&resultCefSharedMemoryRegion)))
	return AsCefSharedMemoryRegion(resultCefSharedMemoryRegion)
}
