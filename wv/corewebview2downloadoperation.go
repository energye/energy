//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2DownloadOperation Parent: IObject
//
//	Represents a download operation. Gives access to the download's metadata
//	and supports a user canceling, pausing, or resuming the download.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation">See the ICoreWebView2DownloadOperation article.</a>
type ICoreWebView2DownloadOperation interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DownloadOperation // property
	// DownloadID
	//  Custom ID used to identify this download operation.
	DownloadID() int32 // property
	// URI
	//  The URI of the download.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_uri">See the ICoreWebView2DownloadOperation article.</a>
	URI() string // property
	// ContentDisposition
	//  The Content-Disposition header value from the download's HTTP response.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_contentdisposition">See the ICoreWebView2DownloadOperation article.</a>
	ContentDisposition() string // property
	// MimeType
	//  MIME type of the downloaded content.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_mimetype">See the ICoreWebView2DownloadOperation article.</a>
	MimeType() string // property
	// TotalBytesToReceive
	//  The expected size of the download in total number of bytes based on the
	//  HTTP Content-Length header. Returns -1 if the size is unknown.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_totalbytestoreceive">See the ICoreWebView2DownloadOperation article.</a>
	TotalBytesToReceive() (resultInt64 int64) // property
	// BytesReceived
	//  The number of bytes that have been written to the download file.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_bytesreceived">See the ICoreWebView2DownloadOperation article.</a>
	BytesReceived() (resultInt64 int64) // property
	// EstimatedEndTime
	//  The estimated end time in [ISO 8601 Date and Time Format](https://www.iso.org/iso-8601-date-and-time-format.html).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_estimatedendtime">See the ICoreWebView2DownloadOperation article.</a>
	EstimatedEndTime() TDateTime // property
	// ResultFilePath
	//  The absolute path to the download file, including file name. Host can change
	//  this from `ICoreWebView2DownloadStartingEventArgs`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_resultfilepath">See the ICoreWebView2DownloadOperation article.</a>
	ResultFilePath() string // property
	// State
	//  The state of the download. A download can be in progress, interrupted, or
	//  completed. See `COREWEBVIEW2_DOWNLOAD_STATE` for descriptions of states.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_state">See the ICoreWebView2DownloadOperation article.</a>
	State() TWVDownloadState // property
	// InterruptReason
	//  The reason why connection with file host was broken.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_interruptreason">See the ICoreWebView2DownloadOperation article.</a>
	InterruptReason() TWVDownloadInterruptReason // property
	// CanResume
	//  Returns true if an interrupted download can be resumed. Downloads with
	//  the following interrupt reasons may automatically resume without you
	//  calling any methods:
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE`,
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH`,
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT`.
	//  In these cases download progress may be restarted with `BytesReceived`
	//  reset to 0.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_canresume">See the ICoreWebView2DownloadOperation article.</a>
	CanResume() bool // property
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// Cancel
	//  Cancels the download. If canceled, the default download dialog shows
	//  that the download was canceled. Host should set the `Cancel` property from
	//  `ICoreWebView2SDownloadStartingEventArgs` if the download should be
	//  canceled without displaying the default download dialog.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#cancel">See the ICoreWebView2DownloadOperation article.</a>
	Cancel() bool // function
	// Pause
	//  Pauses the download. If paused, the default download dialog shows that the
	//  download is paused. No effect if download is already paused. Pausing a
	//  download changes the state to `COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED`
	//  with `InterruptReason` set to `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#pause">See the ICoreWebView2DownloadOperation article.</a>
	Pause() bool // function
	// Resume
	//  Resumes a paused download. May also resume a download that was interrupted
	//  for another reason, if `CanResume` returns true. Resuming a download changes
	//  the state from `COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED` to
	//  `COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#resume">See the ICoreWebView2DownloadOperation article.</a>
	Resume() bool // function
}

// TCoreWebView2DownloadOperation Parent: TObject
//
//	Represents a download operation. Gives access to the download's metadata
//	and supports a user canceling, pausing, or resuming the download.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation">See the ICoreWebView2DownloadOperation article.</a>
type TCoreWebView2DownloadOperation struct {
	TObject
}

func NewCoreWebView2DownloadOperation(aBaseIntf ICoreWebView2DownloadOperation, aDownloadID int32) ICoreWebView2DownloadOperation {
	r1 := WV().SysCallN(258, GetObjectUintptr(aBaseIntf), uintptr(aDownloadID))
	return AsCoreWebView2DownloadOperation(r1)
}

func (m *TCoreWebView2DownloadOperation) Initialized() bool {
	r1 := WV().SysCallN(261, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) BaseIntf() ICoreWebView2DownloadOperation {
	var resultCoreWebView2DownloadOperation uintptr
	WV().SysCallN(252, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DownloadOperation)))
	return AsCoreWebView2DownloadOperation(resultCoreWebView2DownloadOperation)
}

func (m *TCoreWebView2DownloadOperation) DownloadID() int32 {
	r1 := WV().SysCallN(259, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2DownloadOperation) URI() string {
	r1 := WV().SysCallN(269, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) ContentDisposition() string {
	r1 := WV().SysCallN(257, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) MimeType() string {
	r1 := WV().SysCallN(263, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) TotalBytesToReceive() (resultInt64 int64) {
	WV().SysCallN(268, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCoreWebView2DownloadOperation) BytesReceived() (resultInt64 int64) {
	WV().SysCallN(253, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCoreWebView2DownloadOperation) EstimatedEndTime() TDateTime {
	r1 := WV().SysCallN(260, m.Instance())
	return TDateTime(r1)
}

func (m *TCoreWebView2DownloadOperation) ResultFilePath() string {
	r1 := WV().SysCallN(265, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) State() TWVDownloadState {
	r1 := WV().SysCallN(267, m.Instance())
	return TWVDownloadState(r1)
}

func (m *TCoreWebView2DownloadOperation) InterruptReason() TWVDownloadInterruptReason {
	r1 := WV().SysCallN(262, m.Instance())
	return TWVDownloadInterruptReason(r1)
}

func (m *TCoreWebView2DownloadOperation) CanResume() bool {
	r1 := WV().SysCallN(254, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(251, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) Cancel() bool {
	r1 := WV().SysCallN(255, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) Pause() bool {
	r1 := WV().SysCallN(264, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) Resume() bool {
	r1 := WV().SysCallN(266, m.Instance())
	return GoBool(r1)
}

func CoreWebView2DownloadOperationClass() TClass {
	ret := WV().SysCallN(256)
	return TClass(ret)
}
