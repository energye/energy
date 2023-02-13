//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"time"
	"unsafe"
)

// type ICefBeforeDownloadCallback
//
// 下载之前回调
type ICefBeforeDownloadCallback struct {
	instance unsafe.Pointer
	browseId int32
	downId   int32
}

// type ICefDownloadItemCallback
//
// 下载中回调
type ICefDownloadItemCallback struct {
	instance unsafe.Pointer
	browseId int32
	downId   int32
}

// type BeforeDownloadItem 下载-事件 信息
type DownloadItem struct {
	Id                 int32
	CurrentSpeed       int64
	PercentComplete    int32
	TotalBytes         int64
	ReceivedBytes      int64
	StartTime          time.Time
	EndTime            time.Time
	FullPath           string
	Url                string
	OriginalUrl        string
	SuggestedFileName  string
	ContentDisposition string
	MimeType           string
	IsValid            bool
	State              int32 //下载状态 -1:下载之前 0:下载中 1:下载取消 2:下载完成
}

type downloadItem struct {
	Id                 uintptr //int32
	CurrentSpeed       uintptr //int64
	PercentComplete    uintptr //int32
	TotalBytes         uintptr //int64
	ReceivedBytes      uintptr //int64
	StartTime          uintptr //TDateTime
	EndTime            uintptr //TDateTime
	FullPath           uintptr //string
	Url                uintptr //string
	OriginalUrl        uintptr //string
	SuggestedFileName  uintptr //string
	ContentDisposition uintptr //string
	MimeType           uintptr //string
	IsValid            uintptr //bool
	State              uintptr //int32
}

/*
取消 参数是数组只为了不传参数
*/
func (m *ICefDownloadItemCallback) Cancel(downloadId ...int32) {
	_CEFChromium_BrowserDownloadCancel(uintptr(m.browseId), uintptr(downloadId[0]))
}

/*
暂停 参数是数组只为了不传参数
*/
func (m *ICefDownloadItemCallback) Pause(downloadId ...int32) {
	_CEFChromium_BrowserDownloadPause(uintptr(m.browseId), uintptr(downloadId[0]))
}

/*
恢复 参数是数组只为了不传参数
*/
func (m *ICefDownloadItemCallback) Resume(downloadId ...int32) {
	var did int32 = 0
	if len(downloadId) > 0 {
		did = downloadId[0]
	} else {
		did = m.downId
	}
	if did > 0 {
		_CEFChromium_DownloadResume(uintptr(m.browseId), uintptr(did))
	}
}

// 下载恢复
func (m *ICefDownloadItemCallback) DownloadResume(browseId, downloadId int32) {
	_CEFChromium_DownloadResume(uintptr(browseId), uintptr(downloadId))
}

// 设置下载目录
//
// downloadPath 设置完整的下载目录. 包含文件包
//
// showDialog 显示保存窗口
func (m *ICefBeforeDownloadCallback) Cont(downloadPath string, showDialog bool) {
	imports.Proc(internale_CEFChromium_SetDownloadPath).Call(uintptr(m.instance), api.PascalStr(downloadPath), api.PascalBool(showDialog))
}
