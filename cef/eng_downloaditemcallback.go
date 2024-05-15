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

// ICefDownloadItemCallback Parent: ICefBaseRefCounted
//
//	Callback interface used to asynchronously cancel a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_download_item_callback_t))</a>
type ICefDownloadItemCallback interface {
	ICefBaseRefCounted
	// Cancel
	//  Call to cancel the download.
	Cancel() // procedure
	// Pause
	//  Call to pause the download.
	Pause() // procedure
	// Resume
	//  Call to resume the download.
	Resume() // procedure
}

// TCefDownloadItemCallback Parent: TCefBaseRefCounted
//
//	Callback interface used to asynchronously cancel a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_download_item_callback_t))</a>
type TCefDownloadItemCallback struct {
	TCefBaseRefCounted
}

// DownloadItemCallbackRef -> ICefDownloadItemCallback
var DownloadItemCallbackRef downloadItemCallback

// downloadItemCallback TCefDownloadItemCallback Ref
type downloadItemCallback uintptr

func (m *downloadItemCallback) UnWrap(data uintptr) ICefDownloadItemCallback {
	var resultCefDownloadItemCallback uintptr
	CEF().SysCallN(882, uintptr(data), uintptr(unsafePointer(&resultCefDownloadItemCallback)))
	return AsCefDownloadItemCallback(resultCefDownloadItemCallback)
}

func (m *TCefDownloadItemCallback) Cancel() {
	CEF().SysCallN(879, m.Instance())
}

func (m *TCefDownloadItemCallback) Pause() {
	CEF().SysCallN(880, m.Instance())
}

func (m *TCefDownloadItemCallback) Resume() {
	CEF().SysCallN(881, m.Instance())
}
