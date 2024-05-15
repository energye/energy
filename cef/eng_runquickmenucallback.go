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

// ICefRunQuickMenuCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for continuation of custom quick menu display.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_run_quick_menu_callback_t))</a>
type ICefRunQuickMenuCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Complete quick menu display by selecting the specified |command_id| and |event_flags|.
	Cont(commandid int32, eventflags TCefEventFlags) // procedure
	// Cancel
	//  Cancel quick menu display.
	Cancel() // procedure
}

// TCefRunQuickMenuCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for continuation of custom quick menu display.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_run_quick_menu_callback_t))</a>
type TCefRunQuickMenuCallback struct {
	TCefBaseRefCounted
}

// RunQuickMenuCallbackRef -> ICefRunQuickMenuCallback
var RunQuickMenuCallbackRef runQuickMenuCallback

// runQuickMenuCallback TCefRunQuickMenuCallback Ref
type runQuickMenuCallback uintptr

func (m *runQuickMenuCallback) UnWrap(data uintptr) ICefRunQuickMenuCallback {
	var resultCefRunQuickMenuCallback uintptr
	CEF().SysCallN(1335, uintptr(data), uintptr(unsafePointer(&resultCefRunQuickMenuCallback)))
	return AsCefRunQuickMenuCallback(resultCefRunQuickMenuCallback)
}

func (m *TCefRunQuickMenuCallback) Cont(commandid int32, eventflags TCefEventFlags) {
	CEF().SysCallN(1334, m.Instance(), uintptr(commandid), uintptr(eventflags))
}

func (m *TCefRunQuickMenuCallback) Cancel() {
	CEF().SysCallN(1333, m.Instance())
}
