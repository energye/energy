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

// ICefRunContextMenuCallback Parent: ICefBaseRefCounted
//
//	Callback interface used for continuation of custom context menu display.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_run_context_menu_callback_t))
type ICefRunContextMenuCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Complete context menu display by selecting the specified |command_id| and |event_flags|.
	Cont(commandId int32, eventFlags TCefEventFlags) // procedure
	// Cancel
	//  Cancel context menu display.
	Cancel() // procedure
}

// TCefRunContextMenuCallback Parent: TCefBaseRefCounted
//
//	Callback interface used for continuation of custom context menu display.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_run_context_menu_callback_t))
type TCefRunContextMenuCallback struct {
	TCefBaseRefCounted
}

// RunContextMenuCallbackRef -> ICefRunContextMenuCallback
var RunContextMenuCallbackRef runContextMenuCallback

// runContextMenuCallback TCefRunContextMenuCallback Ref
type runContextMenuCallback uintptr

func (m *runContextMenuCallback) UnWrap(data uintptr) ICefRunContextMenuCallback {
	var resultCefRunContextMenuCallback uintptr
	CEF().SysCallN(1332, uintptr(data), uintptr(unsafePointer(&resultCefRunContextMenuCallback)))
	return AsCefRunContextMenuCallback(resultCefRunContextMenuCallback)
}

func (m *TCefRunContextMenuCallback) Cont(commandId int32, eventFlags TCefEventFlags) {
	CEF().SysCallN(1331, m.Instance(), uintptr(commandId), uintptr(eventFlags))
}

func (m *TCefRunContextMenuCallback) Cancel() {
	CEF().SysCallN(1330, m.Instance())
}
