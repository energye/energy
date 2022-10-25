//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/commons"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"reflect"
)

//chromium 事件行为
//
//默认情况所有chromium对象事件行为都在主窗口chromium event中执行
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
	SetOnRenderCompMsg(fn ChromiumEventOnRenderCompMsg)
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
}

// Event
func (m *TCEFChromium) SetOnAfterCreated(fn ChromiumEventOnAfterCreated) {
	_CEFChromium_SetOnAfterCreated(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnBeforeBrowser(fn ChromiumEventOnBeforeBrowser) {
	_CEFChromium_SetOnBeforeBrowser(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnAddressChange(fn ChromiumEventOnAddressChange) {
	_CEFChromium_SetOnAddressChange(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnBeforeClose(fn ChromiumEventOnBeforeClose) {
	_CEFChromium_SetOnBeforeClose(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnClose(fn ChromiumEventOnClose) {
	_CEFChromium_SetOnClose(m.instance, fn, m.independentEvent)
}

// pdf
func (m *TCEFChromium) SetOnPdfPrintFinished(fn ChromiumEventOnResult) {
	_CEFChromium_SetOnPdfPrintFinished(m.instance, fn, m.independentEvent)
}

// chromiumEvent zoom
func (m *TCEFChromium) SetOnZoomPctAvailable(fn ChromiumEventOnResultFloat) {
	_CEFChromium_SetOnZoomPctAvailable(m.instance, fn, m.independentEvent)
}

// load loading
func (m *TCEFChromium) SetOnLoadStart(fn ChromiumEventOnLoadStart) {
	_CEFChromium_SetOnLoadStart(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnLoadingStateChange(fn ChromiumEventOnLoadingStateChange) {
	_CEFChromium_SetOnLoadingStateChange(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnLoadingProgressChange(fn ChromiumEventOnLoadingProgressChange) {
	_CEFChromium_SetOnLoadingProgressChange(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnLoadError(fn ChromiumEventOnLoadError) {
	_CEFChromium_SetOnLoadError(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnLoadEnd(fn ChromiumEventOnLoadEnd) {
	_CEFChromium_SetOnLoadEnd(m.instance, fn, m.independentEvent)
}

// download
func (m *TCEFChromium) SetOnBeforeDownload(fn ChromiumEventOnBeforeDownload) {
	_CEFChromium_SetOnBeforeDownload(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnDownloadUpdated(fn ChromiumEventOnDownloadUpdated) {
	_CEFChromium_SetOnDownloadUpdated(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnFullScreenModeChange(fn ChromiumEventOnFullScreenModeChange) {
	_CEFChromium_SetOnFullScreenModeChange(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnKeyEvent(fn ChromiumEventOnKeyEvent) {
	_CEFChromium_SetOnKeyEvent(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnTitleChange(fn ChromiumEventOnTitleChange) {
	_CEFChromium_SetOnTitleChange(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnRenderCompMsg(fn ChromiumEventOnRenderCompMsg) {
	_CEFChromium_SetOnRenderCompMsg(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnRenderProcessTerminated(fn ChromiumEventOnRenderProcessTerminated) {
	_CEFChromium_SetOnRenderProcessTerminated(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnRenderViewReady(fn ChromiumEventOnCefBrowser) {
	_CEFChromium_SetOnRenderViewReady(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnScrollOffsetChanged(fn ChromiumEventOnScrollOffsetChanged) {
	_CEFChromium_SetOnScrollOffsetChanged(m.instance, fn, m.independentEvent)
}

// 进程间通信消息接收
func (m *TCEFChromium) SetOnProcessMessageReceived(fn BrowseProcessMessageReceived) {
	_CEFChromium_SetOnProcessMessageReceived(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnFindResult(fn ChromiumEventOnFindResult) {
	_CEFChromium_SetOnFindResult(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnBeforeResourceLoad(fn ChromiumEventOnBeforeResourceLoad) {
	_CEFChromium_SetOnBeforeResourceLoad(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnResourceResponse(fn ChromiumEventOnResourceResponse) {
	_CEFChromium_SetOnResourceResponse(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnResourceRedirect(fn ChromiumEventOnResourceRedirect) {
	_CEFChromium_SetOnResourceRedirect(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnResourceLoadComplete(fn ChromiumEventOnResourceLoadComplete) {
	_CEFChromium_SetOnResourceLoadComplete(m.instance, fn, m.independentEvent)
}

// cookie begin
func (m *TCEFChromium) SetOnCookieSet(fn ChromiumEventOnCookieSet) {
	_CEFChromium_SetOnCookieSet(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnCookiesDeleted(fn ChromiumEventOnCookiesDeleted) {
	_CEFChromium_SetOnCookiesDeleted(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnCookiesFlushed(fn ChromiumEventOnCookiesFlushed) {
	_CEFChromium_SetOnCookiesFlushed(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnCookiesVisited(fn ChromiumEventOnCookiesVisited) {
	_CEFChromium_SetOnCookiesVisited(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnCookieVisitorDestroyed(fn ChromiumEventOnCookieVisitorDestroyed) {
	_CEFChromium_SetOnCookieVisitorDestroyed(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnBeforeContextMenu(fn ChromiumEventOnBeforeContextMenu) {
	_CEFChromium_SetOnBeforeContextMenu(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnContextMenuCommand(fn ChromiumEventOnContextMenuCommand) {
	if api.DBoolToGoBool(m.cfg.enableMenu) {
		_CEFChromium_SetOnContextMenuCommand(m.instance, fn, m.independentEvent)
	}
}

func (m *TCEFChromium) SetOnContextMenuDismissed(fn ChromiumEventOnContextMenuDismissed) {
	if api.DBoolToGoBool(m.cfg.enableMenu) {
		_CEFChromium_SetOnContextMenuDismissed(m.instance, fn, m.independentEvent)
	}
}

// frame附加连接
func (m *TCEFChromium) SetOnFrameAttached(fn ChromiumEventOnFrameAttached) {
	_CEFChromium_SetOnFrameAttached(m.instance, fn, m.independentEvent)
}

// frame创建
func (m *TCEFChromium) SetOnFrameCreated(fn ChromiumEventOnFrameCreated) {
	_CEFChromium_SetOnFrameCreated(m.instance, fn, m.independentEvent)
}

// 当前frame离开
func (m *TCEFChromium) SetOnFrameDetached(fn ChromiumEventOnFrameDetached) {
	_CEFChromium_SetOnFrameDetached(m.instance, fn, m.independentEvent)
}

// 主frame被改变
func (m *TCEFChromium) SetOnMainFrameChanged(fn ChromiumEventOnMainFrameChanged) {
	_CEFChromium_SetOnMainFrameChanged(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnBeforePopup(fn ChromiumEventOnBeforePopup) {
	_CEFChromium_SetOnBeforePopup(m.instance, fn, m.independentEvent)
}

func (m *TCEFChromium) SetOnOpenUrlFromTab(fn ChromiumEventOnOpenUrlFromTab) {
	_CEFChromium_SetOnOpenUrlFromTab(m.instance, fn, m.independentEvent)
}

// --------TCEFChromium Event proc begin--------

// 解决事件重复
var chromiumOnEventIdMapping = map[string]uintptr{}

func chromiumOnEventNameToId(instance uintptr, fn interface{}, independentEvent bool) uintptr {
	var eventId uintptr
	if independentEvent {
		eventId = api.GetAddEventToMapFn()(instance, fn)
	} else {
		var (
			name = reflect.ValueOf(fn).Type().Name()
			ok   bool
		)
		if eventId, ok = chromiumOnEventIdMapping[name]; !ok {
			eventId = api.GetAddEventToMapFn()(instance, fn)
			chromiumOnEventIdMapping[name] = eventId
		}
	}
	return eventId
}

func _CEFChromium_SetOnAfterCreated(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnAfterCreated").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnBeforeClose
func _CEFChromium_SetOnBeforeClose(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnBeforeClose").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnClose
func _CEFChromium_SetOnClose(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnClose").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnPdfPrintFinished
func _CEFChromium_SetOnPdfPrintFinished(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnPdfPrintFinished").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnZoomPctAvailable
func _CEFChromium_SetOnZoomPctAvailable(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnZoomPctAvailable").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnLoadStart
func _CEFChromium_SetOnLoadStart(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnLoadStart").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnLoadingStateChange
func _CEFChromium_SetOnLoadingStateChange(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnLoadingStateChange").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnLoadingProgressChange
func _CEFChromium_SetOnLoadingProgressChange(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnLoadingProgressChange").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnLoadError
func _CEFChromium_SetOnLoadError(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnLoadError").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnLoadEnd
func _CEFChromium_SetOnLoadEnd(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnLoadEnd").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnBeforeDownload
func _CEFChromium_SetOnBeforeDownload(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnBeforeDownload").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnDownloadUpdated
func _CEFChromium_SetOnDownloadUpdated(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnDownloadUpdated").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnFullScreenModeChange
func _CEFChromium_SetOnFullScreenModeChange(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnFullScreenModeChange").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnBeforeBrowser
func _CEFChromium_SetOnBeforeBrowser(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnBeforeBrowse").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnAddressChange
func _CEFChromium_SetOnAddressChange(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnAddressChange").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnKeyEvent
func _CEFChromium_SetOnKeyEvent(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnKeyEvent").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnTitleChange
func _CEFChromium_SetOnTitleChange(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnTitleChange").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnRenderCompMsg
func _CEFChromium_SetOnRenderCompMsg(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnRenderCompMsg").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnRenderProcessTerminated
func _CEFChromium_SetOnRenderProcessTerminated(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnRenderProcessTerminated").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnRenderViewReady
func _CEFChromium_SetOnRenderViewReady(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnRenderViewReady").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnScrollOffsetChanged
func _CEFChromium_SetOnScrollOffsetChanged(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnScrollOffsetChanged").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnProcessMessageReceived
func _CEFChromium_SetOnProcessMessageReceived(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnProcessMessageReceived").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnFindResult
func _CEFChromium_SetOnFindResult(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnFindResult").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnCookieSet
func _CEFChromium_SetOnCookieSet(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnCookieSet").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnCookiesDeleted
func _CEFChromium_SetOnCookiesDeleted(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnCookiesDeleted").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnCookiesFlushed
func _CEFChromium_SetOnCookiesFlushed(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnCookiesFlushed").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnCookiesVisited
func _CEFChromium_SetOnCookiesVisited(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnCookiesVisited").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnCookieVisitorDestroyed
func _CEFChromium_SetOnCookieVisitorDestroyed(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnCookieVisitorDestroyed").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnBeforeContextMenu
func _CEFChromium_SetOnBeforeContextMenu(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnBeforeContextMenu").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnContextMenuCommand
func _CEFChromium_SetOnContextMenuCommand(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnContextMenuCommand").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnContextMenuDismissed
func _CEFChromium_SetOnContextMenuDismissed(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnContextMenuDismissed").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnBeforeResourceLoad
func _CEFChromium_SetOnBeforeResourceLoad(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnBeforeResourceLoad").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnResourceResponse
func _CEFChromium_SetOnResourceResponse(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnResourceResponse").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnResourceRedirect
func _CEFChromium_SetOnResourceRedirect(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnResourceRedirect").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnResourceLoadComplete
func _CEFChromium_SetOnResourceLoadComplete(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnResourceLoadComplete").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnFrameAttached
func _CEFChromium_SetOnFrameAttached(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnFrameAttached").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnFrameCreated
func _CEFChromium_SetOnFrameCreated(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnFrameCreated").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnFrameDetached
func _CEFChromium_SetOnFrameDetached(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnFrameDetached").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnMainFrameChanged
func _CEFChromium_SetOnMainFrameChanged(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnMainFrameChanged").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnBeforePopup
func _CEFChromium_SetOnBeforePopup(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnBeforePopup").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

// TCEFChromium _CEFChromium_SetOnOpenUrlFromTab
func _CEFChromium_SetOnOpenUrlFromTab(instance uintptr, fn interface{}, independentEvent bool) {
	Proc("CEFChromium_SetOnOpenUrlFromTab").Call(instance, chromiumOnEventNameToId(instance, fn, independentEvent))
}

//--------TCEFChromium Event proc end--------
