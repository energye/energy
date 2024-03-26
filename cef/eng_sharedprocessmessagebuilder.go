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

// ICefSharedProcessMessageBuilder Parent: ICefBaseRefCounted
//
//	Interface that builds a ICefProcessMessage containing a shared memory region. This interface is not thread-safe but may be used exclusively on a different thread from the one which constructed it.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_shared_process_message_builder_capi.h">CEF source file: /include/capi/cef_shared_process_message_builder_capi.h (cef_shared_process_message_builder_t))
type ICefSharedProcessMessageBuilder interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the builder is valid.
	IsValid() bool // function
	// Size
	//  Returns the size of the shared memory region in bytes. Returns 0 for invalid instances.
	Size() NativeUInt // function
	// Memory
	//  Returns the pointer to the writable memory. Returns nullptr for invalid instances. The returned pointer is only valid for the life span of this object.
	Memory() uintptr // function
	// Build
	//  Creates a new ICefProcessMessage from the data provided to the builder. Returns nullptr for invalid instances. Invalidates the builder instance.
	Build() ICefProcessMessage // function
}

// TCefSharedProcessMessageBuilder Parent: TCefBaseRefCounted
//
//	Interface that builds a ICefProcessMessage containing a shared memory region. This interface is not thread-safe but may be used exclusively on a different thread from the one which constructed it.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_shared_process_message_builder_capi.h">CEF source file: /include/capi/cef_shared_process_message_builder_capi.h (cef_shared_process_message_builder_t))
type TCefSharedProcessMessageBuilder struct {
	TCefBaseRefCounted
}

// SharedProcessMessageBuilderRef -> ICefSharedProcessMessageBuilder
var SharedProcessMessageBuilderRef sharedProcessMessageBuilder

// sharedProcessMessageBuilder TCefSharedProcessMessageBuilder Ref
type sharedProcessMessageBuilder uintptr

func (m *sharedProcessMessageBuilder) UnWrap(data uintptr) ICefSharedProcessMessageBuilder {
	var resultCefSharedProcessMessageBuilder uintptr
	CEF().SysCallN(1373, uintptr(data), uintptr(unsafePointer(&resultCefSharedProcessMessageBuilder)))
	return AsCefSharedProcessMessageBuilder(resultCefSharedProcessMessageBuilder)
}

func (m *sharedProcessMessageBuilder) CreateBuilder(name string, bytesize NativeUInt) ICefSharedProcessMessageBuilder {
	var resultCefSharedProcessMessageBuilder uintptr
	CEF().SysCallN(1369, PascalStr(name), uintptr(bytesize), uintptr(unsafePointer(&resultCefSharedProcessMessageBuilder)))
	return AsCefSharedProcessMessageBuilder(resultCefSharedProcessMessageBuilder)
}

func (m *TCefSharedProcessMessageBuilder) IsValid() bool {
	r1 := CEF().SysCallN(1370, m.Instance())
	return GoBool(r1)
}

func (m *TCefSharedProcessMessageBuilder) Size() NativeUInt {
	r1 := CEF().SysCallN(1372, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefSharedProcessMessageBuilder) Memory() uintptr {
	r1 := CEF().SysCallN(1371, m.Instance())
	return uintptr(r1)
}

func (m *TCefSharedProcessMessageBuilder) Build() ICefProcessMessage {
	var resultCefProcessMessage uintptr
	CEF().SysCallN(1368, m.Instance(), uintptr(unsafePointer(&resultCefProcessMessage)))
	return AsCefProcessMessage(resultCefProcessMessage)
}
