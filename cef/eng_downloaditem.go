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

// ICefDownloadItem Parent: ICefBaseRefCounted
//
//	Interface used to represent a download item.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_item_capi.h">CEF source file: /include/capi/cef_download_item_capi.h (cef_download_item_t))
type ICefDownloadItem interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsInProgress
	//  Returns true (1) if the download is in progress.
	IsInProgress() bool // function
	// IsComplete
	//  Returns true (1) if the download is complete.
	IsComplete() bool // function
	// IsCanceled
	//  Returns true (1) if the download has been canceled.
	IsCanceled() bool // function
	// IsInterrupted
	//  Returns true (1) if the download has been interrupted.
	IsInterrupted() bool // function
	// GetInterruptReason
	//  Returns the most recent interrupt reason.
	GetInterruptReason() TCefDownloadInterruptReason // function
	// GetCurrentSpeed
	//  Returns a simple speed estimate in bytes/s.
	GetCurrentSpeed() (resultInt64 int64) // function
	// GetPercentComplete
	//  Returns the rough percent complete or -1 if the receive total size is unknown.
	GetPercentComplete() int32 // function
	// GetTotalBytes
	//  Returns the total number of bytes.
	GetTotalBytes() (resultInt64 int64) // function
	// GetReceivedBytes
	//  Returns the number of received bytes.
	GetReceivedBytes() (resultInt64 int64) // function
	// GetStartTime
	//  Returns the time that the download started.
	GetStartTime() (resultDateTime TDateTime) // function
	// GetEndTime
	//  Returns the time that the download ended.
	GetEndTime() (resultDateTime TDateTime) // function
	// GetFullPath
	//  Returns the full path to the downloaded or downloading file.
	GetFullPath() string // function
	// GetId
	//  Returns the unique identifier for this download.
	GetId() uint32 // function
	// GetUrl
	//  Returns the URL.
	GetUrl() string // function
	// GetOriginalUrl
	//  Returns the original URL before any redirections.
	GetOriginalUrl() string // function
	// GetSuggestedFileName
	//  Returns the suggested file name.
	GetSuggestedFileName() string // function
	// GetContentDisposition
	//  Returns the content disposition.
	GetContentDisposition() string // function
	// GetMimeType
	//  Returns the mime type.
	GetMimeType() string // function
}

// TCefDownloadItem Parent: TCefBaseRefCounted
//
//	Interface used to represent a download item.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_item_capi.h">CEF source file: /include/capi/cef_download_item_capi.h (cef_download_item_t))
type TCefDownloadItem struct {
	TCefBaseRefCounted
}

// DownloadItemRef -> ICefDownloadItem
var DownloadItemRef downloadItem

// downloadItem TCefDownloadItem Ref
type downloadItem uintptr

func (m *downloadItem) UnWrap(data uintptr) ICefDownloadItem {
	var resultCefDownloadItem uintptr
	CEF().SysCallN(902, uintptr(data), uintptr(unsafePointer(&resultCefDownloadItem)))
	return AsCefDownloadItem(resultCefDownloadItem)
}

func (m *TCefDownloadItem) IsValid() bool {
	r1 := CEF().SysCallN(901, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsInProgress() bool {
	r1 := CEF().SysCallN(899, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsComplete() bool {
	r1 := CEF().SysCallN(898, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsCanceled() bool {
	r1 := CEF().SysCallN(897, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) IsInterrupted() bool {
	r1 := CEF().SysCallN(900, m.Instance())
	return GoBool(r1)
}

func (m *TCefDownloadItem) GetInterruptReason() TCefDownloadInterruptReason {
	r1 := CEF().SysCallN(888, m.Instance())
	return TCefDownloadInterruptReason(r1)
}

func (m *TCefDownloadItem) GetCurrentSpeed() (resultInt64 int64) {
	CEF().SysCallN(884, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDownloadItem) GetPercentComplete() int32 {
	r1 := CEF().SysCallN(891, m.Instance())
	return int32(r1)
}

func (m *TCefDownloadItem) GetTotalBytes() (resultInt64 int64) {
	CEF().SysCallN(895, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDownloadItem) GetReceivedBytes() (resultInt64 int64) {
	CEF().SysCallN(892, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDownloadItem) GetStartTime() (resultDateTime TDateTime) {
	CEF().SysCallN(893, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefDownloadItem) GetEndTime() (resultDateTime TDateTime) {
	CEF().SysCallN(885, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefDownloadItem) GetFullPath() string {
	r1 := CEF().SysCallN(886, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetId() uint32 {
	r1 := CEF().SysCallN(887, m.Instance())
	return uint32(r1)
}

func (m *TCefDownloadItem) GetUrl() string {
	r1 := CEF().SysCallN(896, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetOriginalUrl() string {
	r1 := CEF().SysCallN(890, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetSuggestedFileName() string {
	r1 := CEF().SysCallN(894, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetContentDisposition() string {
	r1 := CEF().SysCallN(883, m.Instance())
	return GoStr(r1)
}

func (m *TCefDownloadItem) GetMimeType() string {
	r1 := CEF().SysCallN(889, m.Instance())
	return GoStr(r1)
}
