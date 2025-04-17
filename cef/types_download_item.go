//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"time"
	"unsafe"
)

// ICefDownloadItem 下载项
type ICefDownloadItem struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DownloadItemRef -> ICefDownloadItem
var DownloadItemRef downloadItem

// downloadItem
type downloadItem uintptr

func (*downloadItem) UnWrap(data *ICefDownloadItem) *ICefDownloadItem {
	var result uintptr
	imports.Proc(def.CefDownloadItemRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

func (m *ICefDownloadItem) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefDownloadItem) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDownloadItem) IsInProgress() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_IsInProgress).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDownloadItem) IsComplete() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_IsComplete).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDownloadItem) IsCanceled() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_IsCanceled).Call(m.Instance())
	return api.GoBool(r1)
}

// IsInterrupted
//
//	CEF 114 ~
func (m *ICefDownloadItem) IsInterrupted() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_IsInterrupted).Call(m.Instance())
	return api.GoBool(r1)
}

// GetInterruptReason
//
//	CEF 114 ~
func (m *ICefDownloadItem) GetInterruptReason() consts.TCefDownloadInterruptReason {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_GetInterruptReason).Call(m.Instance())
	return consts.TCefDownloadInterruptReason(r1)
}

func (m *ICefDownloadItem) CurrentSpeed() int64 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_CurrentSpeed).Call(m.Instance())
	return int64(r1)
}

func (m *ICefDownloadItem) PercentComplete() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_PercentComplete).Call(m.Instance())
	return int32(r1)
}

func (m *ICefDownloadItem) TotalBytes() int64 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_TotalBytes).Call(m.Instance())
	return int64(r1)
}

func (m *ICefDownloadItem) ReceivedBytes() int64 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_ReceivedBytes).Call(m.Instance())
	return int64(r1)
}

func (m *ICefDownloadItem) StartTime() time.Time {
	if !m.IsValid() {
		return time.Time{}
	}
	var result uintptr
	imports.Proc(def.CefDownloadItem_StartTime).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return common.DDateTimeToGoDateTime(*(*float64)(unsafe.Pointer(result)))
}

func (m *ICefDownloadItem) EndTime() time.Time {
	if !m.IsValid() {
		return time.Time{}
	}
	var result uintptr
	imports.Proc(def.CefDownloadItem_EndTime).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return common.DDateTimeToGoDateTime(*(*float64)(unsafe.Pointer(result)))
}

func (m *ICefDownloadItem) FullPath() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_FullPath).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) Id() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_Id).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefDownloadItem) Url() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_Url).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) OriginalUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_OriginalUrl).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) SuggestedFileName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_SuggestedFileName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) ContentDisposition() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_ContentDisposition).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDownloadItem) MimeType() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDownloadItem_MimeType).Call(m.Instance())
	return api.GoStr(r1)
}

// State 下载状态 -1:下载之前 0:下载中 1:下载取消 2:下载完成
func (m *ICefDownloadItem) State() int32 {
	if !m.IsValid() {
		return 0
	}
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

func (m *ICefDownloadItem) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// ICefDownloadItemCallback
//
// 下载中回调
type ICefDownloadItemCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DownloadItemCallbackRef => ICefDownloadItemCallback
var DownloadItemCallbackRef downloadItemCallback

type downloadItemCallback uintptr

func (*downloadItemCallback) UnWrap(data *ICefDownloadItemCallback) *ICefDownloadItemCallback {
	var result uintptr
	imports.Proc(def.CefDownloadItemCallbackRef_Pause).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefDownloadItemCallback{instance: unsafe.Pointer(result)}
}

func (m *ICefDownloadItemCallback) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefDownloadItemCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

/*
Cancel 取消下载 仅在回调函数中使用
*/
func (m *ICefDownloadItemCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDownloadItemCallback_Cancel).Call(m.Instance())
}

/*
Pause 暂停 仅在回调函数中使用
*/
func (m *ICefDownloadItemCallback) Pause() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDownloadItemCallback_Pause).Call(m.Instance())
}

/*
Resume 恢复下载 仅在回调函数中使用
*/
func (m *ICefDownloadItemCallback) Resume() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDownloadItemCallback_Resume).Call(m.Instance())
}

func (m *ICefDownloadItemCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// ICefBeforeDownloadCallback
//
// 下载之前回调
type ICefBeforeDownloadCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

func (m *ICefBeforeDownloadCallback) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefBeforeDownloadCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Cont
// 设置下载目录 仅在回调函数中使用
//
// downloadPath 设置完整的下载目录, 包含文件名
//
// showDialog 弹出保存目录窗口
func (m *ICefBeforeDownloadCallback) Cont(downloadPath string, showDialog bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefBeforeDownloadCallback_Cont).Call(m.Instance(), api.PascalStr(downloadPath), api.PascalBool(showDialog))
}

func (m *ICefBeforeDownloadCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
