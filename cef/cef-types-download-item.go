//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF下载项
package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"time"
	"unsafe"
)

func (m *ICefDownloadItem) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefDownloadItem) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefDownloadItem_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDownloadItem) IsInProgress() bool {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_IsInProgress).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDownloadItem) IsComplete() bool {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_IsComplete).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDownloadItem) IsCanceled() bool {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_IsCanceled).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDownloadItem) CurrentSpeed() int64 {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_CurrentSpeed).Call(m.Instance())
	return int64(r1)
}

func (m *ICefDownloadItem) PercentComplete() int32 {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_PercentComplete).Call(m.Instance())
	return int32(r1)
}

func (m *ICefDownloadItem) TotalBytes() int64 {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_TotalBytes).Call(m.Instance())
	return int64(r1)
}

func (m *ICefDownloadItem) ReceivedBytes() int64 {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_ReceivedBytes).Call(m.Instance())
	return int64(r1)
}

func (m *ICefDownloadItem) StartTime() time.Time {
	var result uintptr
	imports.Proc(internale_CefDownloadItem_StartTime).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return common.DDateTimeToGoDateTime(*(*float64)(unsafe.Pointer(result)))
}

func (m *ICefDownloadItem) EndTime() time.Time {
	var result uintptr
	imports.Proc(internale_CefDownloadItem_EndTime).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return common.DDateTimeToGoDateTime(*(*float64)(unsafe.Pointer(result)))
}

func (m *ICefDownloadItem) FullPath() string {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_FullPath).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) Id() uint32 {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_Id).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefDownloadItem) Url() string {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_Url).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) OriginalUrl() string {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_OriginalUrl).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) SuggestedFileName() string {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_SuggestedFileName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) ContentDisposition() string {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_ContentDisposition).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) MimeType() string {
	r1, _, _ := imports.Proc(internale_CefDownloadItem_MimeType).Call(m.Instance())
	return api.GoStr(r1)
}

//State 下载状态 -1:下载之前 0:下载中 1:下载取消 2:下载完成
func (m *ICefDownloadItem) State() int32 {
	if m.IsComplete() {
		return 2
	} else if m.IsCanceled() {
		return 1
	} else if m.IsInProgress() {
		return 0
	} else {
		return -1
	}
}

func (m *ICefDownloadItemCallback) Instance() uintptr {
	return uintptr(m.instance)
}

/*
Cancel 取消下载
*/
func (m *ICefDownloadItemCallback) Cancel() {
	imports.Proc(internale_CefDownloadItemCallback_Cancel).Call(m.Instance())
}

/*
Pause 暂停
*/
func (m *ICefDownloadItemCallback) Pause() {
	imports.Proc(internale_CefDownloadItemCallback_Pause).Call(m.Instance())
}

/*
Resume 恢复下载
*/
func (m *ICefDownloadItemCallback) Resume() {
	imports.Proc(internale_CefDownloadItemCallback_Resume).Call(m.Instance())
}

func (m *ICefBeforeDownloadCallback) Instance() uintptr {
	return uintptr(m.instance)
}

// Cont
// 设置下载目录
//
// downloadPath 设置完整的下载目录, 包含文件名
//
// showDialog 弹出保存目录窗口
func (m *ICefBeforeDownloadCallback) Cont(downloadPath string, showDialog bool) {
	imports.Proc(internale_CefBeforeDownloadCallback_Cont).Call(m.Instance(), api.PascalStr(downloadPath), api.PascalBool(showDialog))
}
