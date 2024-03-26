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

// ICefCallback Parent: ICefBaseRefCounted
//
//	Generic callback interface used for asynchronous continuation.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_callback_capi.h">CEF source file: /include/capi/cef_callback_capi.h (cef_callback_t))
type ICefCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Continue processing.
	Cont() // procedure
	// Cancel
	//  Cancel processing.
	Cancel() // procedure
}

// TCefCallback Parent: TCefBaseRefCounted
//
//	Generic callback interface used for asynchronous continuation.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_callback_capi.h">CEF source file: /include/capi/cef_callback_capi.h (cef_callback_t))
type TCefCallback struct {
	TCefBaseRefCounted
}

// CallbackRef -> ICefCallback
var CallbackRef callback

// callback TCefCallback Ref
type callback uintptr

func (m *callback) UnWrap(data uintptr) ICefCallback {
	var resultCefCallback uintptr
	CEF().SysCallN(721, uintptr(data), uintptr(unsafePointer(&resultCefCallback)))
	return AsCefCallback(resultCefCallback)
}

func (m *TCefCallback) Cont() {
	CEF().SysCallN(720, m.Instance())
}

func (m *TCefCallback) Cancel() {
	CEF().SysCallN(719, m.Instance())
}
