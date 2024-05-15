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

// ICefBeforeDownloadCallback Parent: ICefBaseRefCounted
//
//	Callback interface used to asynchronously continue a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_before_download_callback_t))</a>
type ICefBeforeDownloadCallback interface {
	ICefBaseRefCounted
	// Cont
	//  Call to continue the download. Set |download_path| to the full file path for the download including the file name or leave blank to use the suggested name and the default temp directory. Set |show_dialog| to true (1) if you do wish to show the default "Save As" dialog.
	Cont(downloadPath string, showDialog bool) // procedure
}

// TCefBeforeDownloadCallback Parent: TCefBaseRefCounted
//
//	Callback interface used to asynchronously continue a download.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_before_download_callback_t))</a>
type TCefBeforeDownloadCallback struct {
	TCefBaseRefCounted
}

// BeforeDownloadCallbackRef -> ICefBeforeDownloadCallback
var BeforeDownloadCallbackRef beforeDownloadCallback

// beforeDownloadCallback TCefBeforeDownloadCallback Ref
type beforeDownloadCallback uintptr

func (m *beforeDownloadCallback) UnWrap(data uintptr) ICefBeforeDownloadCallback {
	var resultCefBeforeDownloadCallback uintptr
	CEF().SysCallN(596, uintptr(data), uintptr(unsafePointer(&resultCefBeforeDownloadCallback)))
	return AsCefBeforeDownloadCallback(resultCefBeforeDownloadCallback)
}

func (m *TCefBeforeDownloadCallback) Cont(downloadPath string, showDialog bool) {
	CEF().SysCallN(595, m.Instance(), PascalStr(downloadPath), PascalBool(showDialog))
}
