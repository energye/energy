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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

// chromium 事件行为
//
// 默认情况所有chromium对象事件行为都在主窗口chromium event中执行
type IChromiumEvent interface {
	lcl.IObject
	SetOnAfterCreated(fn ChromiumEventOnAfterCreated)
	SetOnBeforeBrowser(fn ChromiumEventOnBeforeBrowser)
	SetOnAddressChange(fn ChromiumEventOnAddressChange)
	SetOnBeforeClose(fn ChromiumEventOnBeforeClose)
	SetOnClose(fn ChromiumEventOnClose)
	SetOnPdfPrintFinished(fn ChromiumEventOnResult)
	SetOnZoomPctAvailable(fn ChromiumEventOnResultFloat)
	SetOnLoadStart(fn ChromiumEventOnLoadStart)
	SetOnLoadingStateChange(fn ChromiumEventOnLoadingStateChange)
	SetOnLoadingProgressChange(fn ChromiumEventOnLoadingProgressChange)
	SetOnLoadError(fn ChromiumEventOnLoadError)
	SetOnLoadEnd(fn ChromiumEventOnLoadEnd)
	SetOnBeforeDownload(fn ChromiumEventOnBeforeDownload)
	SetOnDownloadUpdated(fn ChromiumEventOnDownloadUpdated)
	SetOnFullScreenModeChange(fn ChromiumEventOnFullScreenModeChange)
	SetOnKeyEvent(fn ChromiumEventOnKeyEvent)
	SetOnTitleChange(fn ChromiumEventOnTitleChange)
	SetOnRenderCompMsg(fn ChromiumEventOnCompMsg)
	SetOnWidgetCompMsg(fn ChromiumEventOnCompMsg)
	SetOnBrowserCompMsg(fn ChromiumEventOnCompMsg)
	SetOnRenderProcessTerminated(fn ChromiumEventOnRenderProcessTerminated)
	SetOnRenderViewReady(fn ChromiumEventOnCefBrowser)
	SetOnScrollOffsetChanged(fn ChromiumEventOnScrollOffsetChanged)
	SetOnProcessMessageReceived(fn BrowseProcessMessageReceived)
	SetOnFindResult(fn ChromiumEventOnFindResult)
	SetOnBeforeResourceLoad(fn ChromiumEventOnBeforeResourceLoad)
	SetOnResourceResponse(fn ChromiumEventOnResourceResponse)
	SetOnResourceRedirect(fn ChromiumEventOnResourceRedirect)
	SetOnResourceLoadComplete(fn ChromiumEventOnResourceLoadComplete)
	SetOnCookieSet(fn ChromiumEventOnCookieSet)
	SetOnCookiesDeleted(fn ChromiumEventOnCookiesDeleted)
	SetOnCookiesFlushed(fn ChromiumEventOnCookiesFlushed)
	SetOnCookiesVisited(fn ChromiumEventOnCookiesVisited)
	SetOnCookieVisitorDestroyed(fn ChromiumEventOnCookieVisitorDestroyed)
	SetOnBeforeContextMenu(fn ChromiumEventOnBeforeContextMenu)
	SetOnContextMenuCommand(fn ChromiumEventOnContextMenuCommand)
	SetOnContextMenuDismissed(fn ChromiumEventOnContextMenuDismissed)
	SetOnFrameAttached(fn ChromiumEventOnFrameAttached)
	SetOnFrameCreated(fn ChromiumEventOnFrameCreated)
	SetOnFrameDetached(fn ChromiumEventOnFrameDetached)
	SetOnMainFrameChanged(fn ChromiumEventOnMainFrameChanged)
	SetOnBeforePopup(fn ChromiumEventOnBeforePopup)
	SetOnOpenUrlFromTab(fn ChromiumEventOnOpenUrlFromTab)
	SetOnDragEnter(fn ChromiumEventOnDragEnter)
	SetOnDraggableRegionsChanged(fn ChromiumEventOnDraggableRegionsChanged)
}

// Event
func (m *TCEFChromium) SetOnAfterCreated(fn ChromiumEventOnAfterCreated) {
	_CEFChromium_SetOnAfterCreated(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeBrowser(fn ChromiumEventOnBeforeBrowser) {
	_CEFChromium_SetOnBeforeBrowser(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnAddressChange(fn ChromiumEventOnAddressChange) {
	_CEFChromium_SetOnAddressChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeClose(fn ChromiumEventOnBeforeClose) {
	_CEFChromium_SetOnBeforeClose(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnClose(fn ChromiumEventOnClose) {
	_CEFChromium_SetOnClose(m.Instance(), fn)
}

// pdf
func (m *TCEFChromium) SetOnPdfPrintFinished(fn ChromiumEventOnResult) {
	_CEFChromium_SetOnPdfPrintFinished(m.Instance(), fn)
}

// chromiumEvent zoom
func (m *TCEFChromium) SetOnZoomPctAvailable(fn ChromiumEventOnResultFloat) {
	_CEFChromium_SetOnZoomPctAvailable(m.Instance(), fn)
}

// load loading
func (m *TCEFChromium) SetOnLoadStart(fn ChromiumEventOnLoadStart) {
	_CEFChromium_SetOnLoadStart(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadingStateChange(fn ChromiumEventOnLoadingStateChange) {
	_CEFChromium_SetOnLoadingStateChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadingProgressChange(fn ChromiumEventOnLoadingProgressChange) {
	_CEFChromium_SetOnLoadingProgressChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadError(fn ChromiumEventOnLoadError) {
	_CEFChromium_SetOnLoadError(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadEnd(fn ChromiumEventOnLoadEnd) {
	_CEFChromium_SetOnLoadEnd(m.Instance(), fn)
}

// download
func (m *TCEFChromium) SetOnBeforeDownload(fn ChromiumEventOnBeforeDownload) {
	_CEFChromium_SetOnBeforeDownload(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnDownloadUpdated(fn ChromiumEventOnDownloadUpdated) {
	_CEFChromium_SetOnDownloadUpdated(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnFullScreenModeChange(fn ChromiumEventOnFullScreenModeChange) {
	_CEFChromium_SetOnFullScreenModeChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnKeyEvent(fn ChromiumEventOnKeyEvent) {
	_CEFChromium_SetOnKeyEvent(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnTitleChange(fn ChromiumEventOnTitleChange) {
	_CEFChromium_SetOnTitleChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnRenderCompMsg(fn ChromiumEventOnCompMsg) {
	_CEFChromium_SetOnRenderCompMsg(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnWidgetCompMsg(fn ChromiumEventOnCompMsg) {
	_CEFChromium_SetOnWidgetCompMsg(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBrowserCompMsg(fn ChromiumEventOnCompMsg) {
	_CEFChromium_SetOnBrowserCompMsg(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnRenderProcessTerminated(fn ChromiumEventOnRenderProcessTerminated) {
	_CEFChromium_SetOnRenderProcessTerminated(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnRenderViewReady(fn ChromiumEventOnCefBrowser) {
	_CEFChromium_SetOnRenderViewReady(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnScrollOffsetChanged(fn ChromiumEventOnScrollOffsetChanged) {
	_CEFChromium_SetOnScrollOffsetChanged(m.Instance(), fn)
}

// 进程间通信消息接收
func (m *TCEFChromium) SetOnProcessMessageReceived(fn BrowseProcessMessageReceived) {
	_CEFChromium_SetOnProcessMessageReceived(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnFindResult(fn ChromiumEventOnFindResult) {
	_CEFChromium_SetOnFindResult(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeResourceLoad(fn ChromiumEventOnBeforeResourceLoad) {
	_CEFChromium_SetOnBeforeResourceLoad(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnResourceResponse(fn ChromiumEventOnResourceResponse) {
	_CEFChromium_SetOnResourceResponse(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnResourceRedirect(fn ChromiumEventOnResourceRedirect) {
	_CEFChromium_SetOnResourceRedirect(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnResourceLoadComplete(fn ChromiumEventOnResourceLoadComplete) {
	_CEFChromium_SetOnResourceLoadComplete(m.Instance(), fn)
}

// cookie begin
func (m *TCEFChromium) SetOnCookieSet(fn ChromiumEventOnCookieSet) {
	_CEFChromium_SetOnCookieSet(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookiesDeleted(fn ChromiumEventOnCookiesDeleted) {
	_CEFChromium_SetOnCookiesDeleted(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookiesFlushed(fn ChromiumEventOnCookiesFlushed) {
	_CEFChromium_SetOnCookiesFlushed(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookiesVisited(fn ChromiumEventOnCookiesVisited) {
	_CEFChromium_SetOnCookiesVisited(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookieVisitorDestroyed(fn ChromiumEventOnCookieVisitorDestroyed) {
	_CEFChromium_SetOnCookieVisitorDestroyed(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeContextMenu(fn ChromiumEventOnBeforeContextMenu) {
	_CEFChromium_SetOnBeforeContextMenu(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnContextMenuCommand(fn ChromiumEventOnContextMenuCommand) {
	if api.GoBool(m.cfg.enableMenu) {
		_CEFChromium_SetOnContextMenuCommand(m.Instance(), fn)
	}
}

func (m *TCEFChromium) SetOnContextMenuDismissed(fn ChromiumEventOnContextMenuDismissed) {
	if api.GoBool(m.cfg.enableMenu) {
		_CEFChromium_SetOnContextMenuDismissed(m.Instance(), fn)
	}
}

// frame附加连接
func (m *TCEFChromium) SetOnFrameAttached(fn ChromiumEventOnFrameAttached) {
	_CEFChromium_SetOnFrameAttached(m.Instance(), fn)
}

// frame创建
func (m *TCEFChromium) SetOnFrameCreated(fn ChromiumEventOnFrameCreated) {
	_CEFChromium_SetOnFrameCreated(m.Instance(), fn)
}

// 当前frame离开
func (m *TCEFChromium) SetOnFrameDetached(fn ChromiumEventOnFrameDetached) {
	_CEFChromium_SetOnFrameDetached(m.Instance(), fn)
}

// 主frame被改变
func (m *TCEFChromium) SetOnMainFrameChanged(fn ChromiumEventOnMainFrameChanged) {
	_CEFChromium_SetOnMainFrameChanged(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforePopup(fn ChromiumEventOnBeforePopup) {
	_CEFChromium_SetOnBeforePopup(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnOpenUrlFromTab(fn ChromiumEventOnOpenUrlFromTab) {
	_CEFChromium_SetOnOpenUrlFromTab(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnDragEnter(fn ChromiumEventOnDragEnter) {
	_CEFChromium_SetOnDragEnter(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnDraggableRegionsChanged(fn ChromiumEventOnDraggableRegionsChanged) {
	_CEFChromium_SetOnDraggableRegionsChanged(m.Instance(), fn)
}

// --------TCEFChromium Event proc begin--------

func _CEFChromium_SetOnAfterCreated(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnAfterCreated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeClose
func _CEFChromium_SetOnBeforeClose(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnBeforeClose).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnClose
func _CEFChromium_SetOnClose(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnClose).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnPdfPrintFinished
func _CEFChromium_SetOnPdfPrintFinished(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnPdfPrintFinished).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnZoomPctAvailable
func _CEFChromium_SetOnZoomPctAvailable(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnZoomPctAvailable).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadStart
func _CEFChromium_SetOnLoadStart(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnLoadStart).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadingStateChange
func _CEFChromium_SetOnLoadingStateChange(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnLoadingStateChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadingProgressChange
func _CEFChromium_SetOnLoadingProgressChange(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnLoadingProgressChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadError
func _CEFChromium_SetOnLoadError(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnLoadError).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadEnd
func _CEFChromium_SetOnLoadEnd(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnLoadEnd).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeDownload
func _CEFChromium_SetOnBeforeDownload(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnBeforeDownload).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnDownloadUpdated
func _CEFChromium_SetOnDownloadUpdated(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnDownloadUpdated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFullScreenModeChange
func _CEFChromium_SetOnFullScreenModeChange(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnFullScreenModeChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeBrowser
func _CEFChromium_SetOnBeforeBrowser(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnBeforeBrowse).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnAddressChange
func _CEFChromium_SetOnAddressChange(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnAddressChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnKeyEvent
func _CEFChromium_SetOnKeyEvent(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnKeyEvent).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnTitleChange
func _CEFChromium_SetOnTitleChange(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnTitleChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnRenderCompMsg
func _CEFChromium_SetOnRenderCompMsg(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnRenderCompMsg).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnWidgetCompMsg
func _CEFChromium_SetOnWidgetCompMsg(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnWidgetCompMsg).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBrowserCompMsg
func _CEFChromium_SetOnBrowserCompMsg(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnBrowserCompMsg).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnRenderProcessTerminated
func _CEFChromium_SetOnRenderProcessTerminated(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnRenderProcessTerminated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnRenderViewReady
func _CEFChromium_SetOnRenderViewReady(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnRenderViewReady).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnScrollOffsetChanged
func _CEFChromium_SetOnScrollOffsetChanged(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnScrollOffsetChanged).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnProcessMessageReceived
func _CEFChromium_SetOnProcessMessageReceived(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnProcessMessageReceived).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFindResult
func _CEFChromium_SetOnFindResult(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnFindResult).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookieSet
func _CEFChromium_SetOnCookieSet(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnCookieSet).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookiesDeleted
func _CEFChromium_SetOnCookiesDeleted(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnCookiesDeleted).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookiesFlushed
func _CEFChromium_SetOnCookiesFlushed(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnCookiesFlushed).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookiesVisited
func _CEFChromium_SetOnCookiesVisited(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnCookiesVisited).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookieVisitorDestroyed
func _CEFChromium_SetOnCookieVisitorDestroyed(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnCookieVisitorDestroyed).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeContextMenu
func _CEFChromium_SetOnBeforeContextMenu(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnBeforeContextMenu).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnContextMenuCommand
func _CEFChromium_SetOnContextMenuCommand(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnContextMenuCommand).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnContextMenuDismissed
func _CEFChromium_SetOnContextMenuDismissed(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnContextMenuDismissed).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeResourceLoad
func _CEFChromium_SetOnBeforeResourceLoad(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnBeforeResourceLoad).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnResourceResponse
func _CEFChromium_SetOnResourceResponse(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnResourceResponse).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnResourceRedirect
func _CEFChromium_SetOnResourceRedirect(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnResourceRedirect).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnResourceLoadComplete
func _CEFChromium_SetOnResourceLoadComplete(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnResourceLoadComplete).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFrameAttached
func _CEFChromium_SetOnFrameAttached(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnFrameAttached).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFrameCreated
func _CEFChromium_SetOnFrameCreated(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnFrameCreated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFrameDetached
func _CEFChromium_SetOnFrameDetached(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnFrameDetached).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnMainFrameChanged
func _CEFChromium_SetOnMainFrameChanged(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnMainFrameChanged).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforePopup
func _CEFChromium_SetOnBeforePopup(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnBeforePopup).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnOpenUrlFromTab
func _CEFChromium_SetOnOpenUrlFromTab(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnOpenUrlFromTab).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnDragEnter
func _CEFChromium_SetOnDragEnter(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnDragEnter).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnDraggableRegionsChanged
func _CEFChromium_SetOnDraggableRegionsChanged(instance uintptr, fn interface{}) {
	imports.Proc(internale_CEFChromium_SetOnDraggableRegionsChanged).Call(instance, api.MakeEventDataPtr(fn))
}

//--------TCEFChromium Event proc end--------
