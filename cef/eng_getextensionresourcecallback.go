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

// ICefGetExtensionResourceCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of ICefExtensionHandler.GetExtensionResource.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h (cef_get_extension_resource_callback_t))</a>
type ICefGetExtensionResourceCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue the request. Read the resource contents from |stream|.
	Cont(stream ICefStreamReader) // procedure
	// Cancel
	//  Cancel the request.
	Cancel() // procedure
}

// TCefGetExtensionResourceCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for asynchronous continuation of ICefExtensionHandler.GetExtensionResource.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h (cef_get_extension_resource_callback_t))</a>
type TCefGetExtensionResourceCallback struct {
	TCefBaseRefCounted
}

// GetExtensionResourceCallbackRef -> ICefGetExtensionResourceCallback
var GetExtensionResourceCallbackRef getExtensionResourceCallback

// getExtensionResourceCallback TCefGetExtensionResourceCallback Ref
type getExtensionResourceCallback uintptr

func (m *getExtensionResourceCallback) UnWrap(data uintptr) ICefGetExtensionResourceCallback {
	var resultCefGetExtensionResourceCallback uintptr
	CEF().SysCallN(985, uintptr(data), uintptr(unsafePointer(&resultCefGetExtensionResourceCallback)))
	return AsCefGetExtensionResourceCallback(resultCefGetExtensionResourceCallback)
}

func (m *TCefGetExtensionResourceCallback) Cont(stream ICefStreamReader) {
	CEF().SysCallN(984, m.Instance(), GetObjectUintptr(stream))
}

func (m *TCefGetExtensionResourceCallback) Cancel() {
	CEF().SysCallN(983, m.Instance())
}
