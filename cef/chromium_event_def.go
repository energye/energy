//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Chromium 事件接口定义

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
)

// IChromiumEvent
// Chromium 事件行为接口
type IChromiumEvent interface {
	SetOnAfterCreated(fn chromiumEventOnAfterCreated)
	SetOnBeforeBrowser(fn chromiumEventOnBeforeBrowser)
	SetOnAddressChange(fn chromiumEventOnAddressChange)
	SetOnBeforeClose(fn chromiumEventOnBeforeClose)
	SetOnClose(fn chromiumEventOnClose)
	SetOnPdfPrintFinished(fn chromiumEventOnPdfPrintFinished)
	SetOnZoomPctAvailable(fn chromiumEventOnResultFloat)
	SetOnLoadStart(fn chromiumEventOnLoadStart)
	SetOnLoadingStateChange(fn chromiumEventOnLoadingStateChange)
	SetOnLoadingProgressChange(fn chromiumEventOnLoadingProgressChange)
	SetOnLoadError(fn chromiumEventOnLoadError)
	SetOnLoadEnd(fn chromiumEventOnLoadEnd)
	SetOnBeforeDownload(fn chromiumEventOnBeforeDownload)
	SetOnDownloadUpdated(fn chromiumEventOnDownloadUpdated)
	SetOnFullScreenModeChange(fn chromiumEventOnFullScreenModeChange)
	SetOnKeyEvent(fn chromiumEventOnKey)
	SetOnTitleChange(fn chromiumEventOnTitleChange)
	SetOnRenderCompMsg(fn chromiumEventOnCompMsg)
	SetOnBrowserCompMsg(fn chromiumEventOnCompMsg)
	SetOnRenderProcessTerminated(fn chromiumEventOnRenderProcessTerminated)
	SetOnRenderViewReady(fn chromiumEventOnRenderViewReady)
	SetOnScrollOffsetChanged(fn chromiumEventOnScrollOffsetChanged)
	SetOnProcessMessageReceived(fn BrowseProcessMessageReceived)
	SetOnFindResult(fn chromiumEventOnFindResult)
	SetOnBeforeResourceLoad(fn chromiumEventOnBeforeResourceLoad)
	SetOnResourceResponse(fn chromiumEventOnResourceResponse)
	SetOnResourceRedirect(fn chromiumEventOnResourceRedirect)
	SetOnResourceLoadComplete(fn chromiumEventOnResourceLoadComplete)
	SetOnCookieSet(fn chromiumEventOnCookieSet)
	SetOnCookiesDeleted(fn chromiumEventOnCookiesDeleted)
	SetOnCookiesFlushed(fn chromiumEventOnCookiesFlushed)
	SetOnCookiesVisited(fn chromiumEventOnCookiesVisited)
	SetOnCookieVisitorDestroyed(fn chromiumEventOnCookieVisitorDestroyed)
	SetOnBeforeContextMenu(fn chromiumEventOnBeforeContextMenu)
	SetOnContextMenuCommand(fn chromiumEventOnContextMenuCommand)
	SetOnContextMenuDismissed(fn chromiumEventOnContextMenuDismissed)
	SetOnFrameAttached(fn chromiumEventOnFrameAttached)
	SetOnFrameCreated(fn chromiumEventOnFrameCreated)
	SetOnFrameDetached(fn chromiumEventOnFrameDetached)
	SetOnMainFrameChanged(fn chromiumEventOnMainFrameChanged)
	SetOnBeforePopup(fn chromiumEventOnBeforePopup)
	SetOnOpenUrlFromTab(fn chromiumEventOnOpenUrlFromTab)
	SetOnDragEnter(fn chromiumEventOnDragEnter)
	SetOnDraggableRegionsChanged(fn chromiumEventOnDraggableRegionsChanged)
	SetOnGetAuthCredentials(fn chromiumEventOnGetAuthCredentials)
	SetOnAcceleratedPaint(fn chromiumEventOnAcceleratedPaint)
	SetOnAllConnectionsClosed(fn chromiumEventOnAllConnectionsClosed)
	SetOnAudioStreamError(fn chromiumEventOnAudioStreamError)
	SetOnAudioStreamPacket(fn chromiumEventOnAudioStreamPacket)
	SetOnAudioStreamStarted(fn chromiumEventOnAudioStreamStarted)
	SetOnAudioStreamStopped(fn chromiumEventOnAudioStreamStopped)
	SetOnAutoResize(fn chromiumEventOnAutoResize)
	SetOnBeforeUnloadDialog(fn chromiumEventOnBeforeUnloadDialog)
	SetOnCanDownload(fn chromiumEventOnCanDownload)
	SetOnCanSaveCookie(fn chromiumEventOnCanSaveCookie)
	SetOnCanSendCookie(fn chromiumEventOnCanSendCookie)
	SetOnCertificateError(fn chromiumEventOnCertificateError)
	SetOnCertificateExceptionsCleared(fn chromiumEventOnCertificateExceptionsCleared)
	SetOnChromeCommand(fn chromiumEventOnChromeCommand)
	SetOnConsoleMessage(fn chromiumEventOnConsoleMessage)
	SetOnCursorChange(fn chromiumEventOnCursorChange)
	SetOnDevToolsAgentAttached(fn chromiumEventOnDevToolsAgentAttached)
	SetOnDevToolsAgentDetached(fn chromiumEventOnDevToolsAgentDetached)
	SetOnDevToolsEvent(fn chromiumEventOnDevTools)
	SetOnDevToolsMessage(fn chromiumEventOnDevToolsMessage)
	SetOnDevToolsMethodRawResult(fn chromiumEventOnDevToolsMethodRawResult)
	SetOnDevToolsMethodResult(fn chromiumEventOnDevToolsMethodResult)
	SetOnDevToolsRawEvent(fn chromiumEventOnDevToolsRaw)
	SetOnDevToolsRawMessage(fn chromiumEventOnDevToolsRawMessage)
	SetOnDialogClosed(fn chromiumEventOnDialogClosed)
	SetOnDismissPermissionPrompt(fn chromiumEventOnDismissPermissionPrompt)
	SetOnDocumentAvailableInMainFrame(fn chromiumEventOnDocumentAvailableInMainFrame)
	SetOnDownloadImageFinished(fn chromiumEventOnDownloadImageFinished)
	SetOnExecuteTaskOnCefThread(fn chromiumEventOnExecuteTaskOnCefThread)
	SetOnPrintStart(fn chromiumEventOnPrintStart)
	SetOnPrintSettings(fn chromiumEventOnPrintSettings)
	SetOnPrintDialog(fn chromiumEventOnPrintDialog)
	SetOnPrintJob(fn chromiumEventOnPrintJob)
	SetOnPrintReset(fn chromiumEventOnPrintReset)
	SetOnGetPDFPaperSize(fn chromiumEventOnGetPDFPaperSize)
	SetOnFavIconUrlChange(fn chromiumEventOnFavIconUrlChange)
	SetOnFileDialog(fn chromiumEventOnFileDialog)
	SetOnGetAccessibilityHandler(fn chromiumEventOnGetAccessibilityHandler)
	SetOnGetAudioParameters(fn chromiumEventOnGetAudioParameters)
	SetOnGetResourceHandler(fn chromiumEventOnGetResourceHandler)
	SetOnGetResourceRequestHandlerReqCtxHdlr(fn chromiumEventOnGetResourceRequestHandlerReqCtxHdlr)
	SetOnGetResourceRequestHandlerReqHdlr(fn chromiumEventOnGetResourceRequestHandlerReqHdlr)
	SetOnGetResourceResponseFilter(fn chromiumEventOnGetResourceResponseFilter)
	SetOnGetRootScreenRect(fn chromiumEventOnGetRootScreenRect)
	SetOnGetScreenInfo(fn chromiumEventOnGetScreenInfo)
	SetOnGetScreenPoint(fn chromiumEventOnGetScreenPoint)
	SetOnGetTouchHandleSize(fn chromiumEventOnGetTouchHandleSize)
	SetOnGetViewRect(fn chromiumEventOnGetViewRect)
	SetOnGotFocus(fn chromiumEventOnGotFocus)
	SetOnHttpAuthCredentialsCleared(fn chromiumEventOnHttpAuthCredentialsCleared)
	SetOnIMECompositionRangeChanged(fn chromiumEventOnIMECompositionRangeChanged)
	SetOnJsDialog(fn chromiumEventOnJsDialog)
	SetOnMediaAccessChange(fn chromiumEventOnMediaAccessChange)
	SetOnMediaRouteCreateFinished(fn chromiumEventOnMediaRouteCreateFinished)
	SetOnMediaSinkDeviceInfo(fn chromiumEventOnMediaSinkDeviceInfo)
	SetOnNavigationVisitorResultAvailable(fn chromiumEventOnNavigationVisitorResultAvailable)
	SetOnPaint(fn chromiumEventOnPaint)
	SetOnPopupShow(fn chromiumEventOnPopupShow)
	SetOnPopupSize(fn chromiumEventOnPopupSize)
	SetOnPrefsAvailable(fn chromiumEventOnPrefsAvailable)
	SetOnPrefsUpdated(fn chromiumEventOnPrefsUpdated)
	SetOnPreKeyEvent(fn chromiumEventOnPreKey)
	SetOnProtocolExecution(fn chromiumEventOnProtocolExecution)
	SetOnQuickMenuCommand(fn chromiumEventOnQuickMenuCommand)
	SetOnQuickMenuDismissed(fn chromiumEventOnQuickMenuDismissed)
	SetOnRequestContextInitialized(fn chromiumEventOnRequestContextInitialized)
	SetOnRequestMediaAccessPermission(fn chromiumEventOnRequestMediaAccessPermission)
	SetOnResetDialogState(fn chromiumEventOnResetDialogState)
	SetOnResolvedHostAvailable(fn chromiumEventOnResolvedHostAvailable)
	SetOnRouteMessageReceived(fn chromiumEventOnRouteMessageReceived)
	SetOnRoutes(fn chromiumEventOnRoutes)
	SetOnRouteStateChanged(fn chromiumEventOnRouteStateChanged)
	SetOnRunContextMenu(fn chromiumEventOnRunContextMenu)
	SetOnRunQuickMenu(fn chromiumEventOnRunQuickMenu)
	SetOnSelectClientCertificate(fn chromiumEventOnSelectClientCertificate)
	SetOnSetFocus(fn chromiumEventOnSetFocus)
	SetOnShowPermissionPrompt(fn chromiumEventOnShowPermissionPrompt)
	SetOnSinks(fn chromiumEventOnSinks)
	SetOnStartDragging(fn chromiumEventOnStartDragging)
	SetOnStatusMessage(fn chromiumEventOnStatusMessage)
	SetOnTakeFocus(fn chromiumEventOnTakeFocus)
	SetOnTextResultAvailable(fn chromiumEventOnTextResultAvailable)
	SetOnTextSelectionChanged(fn chromiumEventOnTextSelectionChanged)
	SetOnTooltip(fn chromiumEventOnTooltip)
	SetOnTouchHandleStateChanged(fn chromiumEventOnTouchHandleStateChanged)
	SetOnUpdateDragCursor(fn chromiumEventOnUpdateDragCursor)
	SetOnVirtualKeyboardRequested(fn chromiumEventOnVirtualKeyboardRequested)
	SetOnIsChromeAppMenuItemVisible(fn chromiumEventOnIsChromeAppMenuItemVisible)       // CEF 112 ~ , 仅适用于 ChromeRuntime 模式
	SetOnIsChromeAppMenuItemEnabled(fn chromiumEventOnIsChromeAppMenuItemEnabled)       // CEF 112 ~ , 仅适用于 ChromeRuntime 模式
	SetOnIsChromePageActionIconVisible(fn chromiumEventOnIsChromePageActionIconVisible) // CEF 112 ~ , 仅适用于 ChromeRuntime 模式
	SetOnIsChromeToolbarButtonVisible(fn chromiumEventOnIsChromeToolbarButtonVisible)   // CEF 112 ~ , 仅适用于 ChromeRuntime 模式
}

func (m *TCEFChromium) SetOnAfterCreated(fn chromiumEventOnAfterCreated) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnAfterCreated(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeBrowser(fn chromiumEventOnBeforeBrowser) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnBeforeBrowser(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnAddressChange(fn chromiumEventOnAddressChange) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnAddressChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeClose(fn chromiumEventOnBeforeClose) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnBeforeClose(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnClose(fn chromiumEventOnClose) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnClose(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnPdfPrintFinished(fn chromiumEventOnPdfPrintFinished) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnPdfPrintFinished(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnZoomPctAvailable(fn chromiumEventOnResultFloat) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnZoomPctAvailable(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadStart(fn chromiumEventOnLoadStart) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnLoadStart(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadingStateChange(fn chromiumEventOnLoadingStateChange) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnLoadingStateChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadingProgressChange(fn chromiumEventOnLoadingProgressChange) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnLoadingProgressChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadError(fn chromiumEventOnLoadError) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnLoadError(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnLoadEnd(fn chromiumEventOnLoadEnd) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnLoadEnd(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeDownload(fn chromiumEventOnBeforeDownload) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnBeforeDownload(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnDownloadUpdated(fn chromiumEventOnDownloadUpdated) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnDownloadUpdated(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnFullScreenModeChange(fn chromiumEventOnFullScreenModeChange) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnFullScreenModeChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnKeyEvent(fn chromiumEventOnKey) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnKeyEvent(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnTitleChange(fn chromiumEventOnTitleChange) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnTitleChange(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnRenderCompMsg(fn chromiumEventOnCompMsg) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnRenderCompMsg(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBrowserCompMsg(fn chromiumEventOnCompMsg) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnBrowserCompMsg(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnRenderProcessTerminated(fn chromiumEventOnRenderProcessTerminated) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnRenderProcessTerminated(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnRenderViewReady(fn chromiumEventOnRenderViewReady) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnRenderViewReady(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnScrollOffsetChanged(fn chromiumEventOnScrollOffsetChanged) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnScrollOffsetChanged(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnProcessMessageReceived(fn BrowseProcessMessageReceived) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnProcessMessageReceived(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnFindResult(fn chromiumEventOnFindResult) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnFindResult(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeResourceLoad(fn chromiumEventOnBeforeResourceLoad) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnBeforeResourceLoad(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnResourceResponse(fn chromiumEventOnResourceResponse) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnResourceResponse(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnResourceRedirect(fn chromiumEventOnResourceRedirect) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnResourceRedirect(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnResourceLoadComplete(fn chromiumEventOnResourceLoadComplete) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnResourceLoadComplete(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookieSet(fn chromiumEventOnCookieSet) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnCookieSet(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookiesDeleted(fn chromiumEventOnCookiesDeleted) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnCookiesDeleted(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookiesFlushed(fn chromiumEventOnCookiesFlushed) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnCookiesFlushed(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookiesVisited(fn chromiumEventOnCookiesVisited) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnCookiesVisited(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnCookieVisitorDestroyed(fn chromiumEventOnCookieVisitorDestroyed) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnCookieVisitorDestroyed(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforeContextMenu(fn chromiumEventOnBeforeContextMenu) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnBeforeContextMenu(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnContextMenuCommand(fn chromiumEventOnContextMenuCommand) {
	if !m.IsValid() {
		return
	}
	if m.Config().EnableMenu() {
		_CEFChromium_SetOnContextMenuCommand(m.Instance(), fn)
	}
}

func (m *TCEFChromium) SetOnContextMenuDismissed(fn chromiumEventOnContextMenuDismissed) {
	if !m.IsValid() {
		return
	}
	if m.Config().EnableMenu() {
		_CEFChromium_SetOnContextMenuDismissed(m.Instance(), fn)
	}
}

func (m *TCEFChromium) SetOnFrameAttached(fn chromiumEventOnFrameAttached) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnFrameAttached(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnFrameCreated(fn chromiumEventOnFrameCreated) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnFrameCreated(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnFrameDetached(fn chromiumEventOnFrameDetached) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnFrameDetached(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnMainFrameChanged(fn chromiumEventOnMainFrameChanged) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnMainFrameChanged(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnBeforePopup(fn chromiumEventOnBeforePopup) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnBeforePopup(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnOpenUrlFromTab(fn chromiumEventOnOpenUrlFromTab) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnOpenUrlFromTab(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnDragEnter(fn chromiumEventOnDragEnter) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnDragEnter(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnDraggableRegionsChanged(fn chromiumEventOnDraggableRegionsChanged) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnDraggableRegionsChanged(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnGetAuthCredentials(fn chromiumEventOnGetAuthCredentials) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetOnGetAuthCredentials(m.Instance(), fn)
}

func (m *TCEFChromium) SetOnPrintStart(fn chromiumEventOnPrintStart) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPrintStart).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPrintSettings(fn chromiumEventOnPrintSettings) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPrintSettings).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPrintDialog(fn chromiumEventOnPrintDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPrintDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPrintJob(fn chromiumEventOnPrintJob) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPrintJob).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPrintReset(fn chromiumEventOnPrintReset) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPrintReset).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetPDFPaperSize(fn chromiumEventOnGetPDFPaperSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetPDFPaperSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnAcceleratedPaint(fn chromiumEventOnAcceleratedPaint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnAcceleratedPaint).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnAllConnectionsClosed(fn chromiumEventOnAllConnectionsClosed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnAllConnectionsClosed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnAudioStreamError(fn chromiumEventOnAudioStreamError) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnAudioStreamError).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnAudioStreamPacket(fn chromiumEventOnAudioStreamPacket) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnAudioStreamPacket).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnAudioStreamStarted(fn chromiumEventOnAudioStreamStarted) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnAudioStreamStarted).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnAudioStreamStopped(fn chromiumEventOnAudioStreamStopped) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnAudioStreamStopped).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnAutoResize(fn chromiumEventOnAutoResize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnAutoResize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnBeforeUnloadDialog(fn chromiumEventOnBeforeUnloadDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnBeforeUnloadDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnCanDownload(fn chromiumEventOnCanDownload) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnCanDownload).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnCanSaveCookie(fn chromiumEventOnCanSaveCookie) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnCanSaveCookie).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnCanSendCookie(fn chromiumEventOnCanSendCookie) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnCanSendCookie).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnCertificateError(fn chromiumEventOnCertificateError) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnCertificateError).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnCertificateExceptionsCleared(fn chromiumEventOnCertificateExceptionsCleared) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnCertificateExceptionsCleared).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnChromeCommand(fn chromiumEventOnChromeCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnChromeCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnConsoleMessage(fn chromiumEventOnConsoleMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnConsoleMessage).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnCursorChange(fn chromiumEventOnCursorChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnCursorChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsAgentAttached(fn chromiumEventOnDevToolsAgentAttached) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsAgentAttached).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsAgentDetached(fn chromiumEventOnDevToolsAgentDetached) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsAgentDetached).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsEvent(fn chromiumEventOnDevTools) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsMessage(fn chromiumEventOnDevToolsMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsMessage).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsMethodRawResult(fn chromiumEventOnDevToolsMethodRawResult) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsMethodRawResult).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsMethodResult(fn chromiumEventOnDevToolsMethodResult) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsMethodResult).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsRawEvent(fn chromiumEventOnDevToolsRaw) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsRawEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDevToolsRawMessage(fn chromiumEventOnDevToolsRawMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDevToolsRawMessage).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDialogClosed(fn chromiumEventOnDialogClosed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDialogClosed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDismissPermissionPrompt(fn chromiumEventOnDismissPermissionPrompt) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDismissPermissionPrompt).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDocumentAvailableInMainFrame(fn chromiumEventOnDocumentAvailableInMainFrame) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDocumentAvailableInMainFrame).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnDownloadImageFinished(fn chromiumEventOnDownloadImageFinished) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnDownloadImageFinished).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnExecuteTaskOnCefThread(fn chromiumEventOnExecuteTaskOnCefThread) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnExecuteTaskOnCefThread).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnFavIconUrlChange(fn chromiumEventOnFavIconUrlChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnFavIconUrlChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnFileDialog(fn chromiumEventOnFileDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnFileDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetAccessibilityHandler(fn chromiumEventOnGetAccessibilityHandler) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetAccessibilityHandler).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetAudioParameters(fn chromiumEventOnGetAudioParameters) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetAudioParameters).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetResourceHandler(fn chromiumEventOnGetResourceHandler) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetResourceHandler).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetResourceRequestHandlerReqCtxHdlr(fn chromiumEventOnGetResourceRequestHandlerReqCtxHdlr) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetResourceRequestHandler_ReqCtxHdlr).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetResourceRequestHandlerReqHdlr(fn chromiumEventOnGetResourceRequestHandlerReqHdlr) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetResourceRequestHandler_ReqHdlr).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetResourceResponseFilter(fn chromiumEventOnGetResourceResponseFilter) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetResourceResponseFilter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetRootScreenRect(fn chromiumEventOnGetRootScreenRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetRootScreenRect).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetScreenInfo(fn chromiumEventOnGetScreenInfo) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetScreenInfo).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetScreenPoint(fn chromiumEventOnGetScreenPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetScreenPoint).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetTouchHandleSize(fn chromiumEventOnGetTouchHandleSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetTouchHandleSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGetViewRect(fn chromiumEventOnGetViewRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGetViewRect).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnGotFocus(fn chromiumEventOnGotFocus) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnGotFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnHttpAuthCredentialsCleared(fn chromiumEventOnHttpAuthCredentialsCleared) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnHttpAuthCredentialsCleared).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnIMECompositionRangeChanged(fn chromiumEventOnIMECompositionRangeChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnIMECompositionRangeChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnJsDialog(fn chromiumEventOnJsDialog) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnJsDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnMediaAccessChange(fn chromiumEventOnMediaAccessChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnMediaAccessChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnMediaRouteCreateFinished(fn chromiumEventOnMediaRouteCreateFinished) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnMediaRouteCreateFinished).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnMediaSinkDeviceInfo(fn chromiumEventOnMediaSinkDeviceInfo) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnMediaSinkDeviceInfo).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnNavigationVisitorResultAvailable(fn chromiumEventOnNavigationVisitorResultAvailable) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnNavigationVisitorResultAvailable).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPaint(fn chromiumEventOnPaint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPaint).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPopupShow(fn chromiumEventOnPopupShow) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPopupShow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPopupSize(fn chromiumEventOnPopupSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPopupSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPrefsAvailable(fn chromiumEventOnPrefsAvailable) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPrefsAvailable).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPrefsUpdated(fn chromiumEventOnPrefsUpdated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPrefsUpdated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnPreKeyEvent(fn chromiumEventOnPreKey) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnPreKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnProtocolExecution(fn chromiumEventOnProtocolExecution) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnProtocolExecution).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnQuickMenuCommand(fn chromiumEventOnQuickMenuCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnQuickMenuCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnQuickMenuDismissed(fn chromiumEventOnQuickMenuDismissed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnQuickMenuDismissed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnRequestContextInitialized(fn chromiumEventOnRequestContextInitialized) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnRequestContextInitialized).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnRequestMediaAccessPermission(fn chromiumEventOnRequestMediaAccessPermission) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnRequestMediaAccessPermission).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnResetDialogState(fn chromiumEventOnResetDialogState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnResetDialogState).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnResolvedHostAvailable(fn chromiumEventOnResolvedHostAvailable) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnResolvedHostAvailable).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnRouteMessageReceived(fn chromiumEventOnRouteMessageReceived) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnRouteMessageReceived).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnRoutes(fn chromiumEventOnRoutes) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnRoutes).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnRouteStateChanged(fn chromiumEventOnRouteStateChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnRouteStateChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnRunContextMenu(fn chromiumEventOnRunContextMenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnRunContextMenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnRunQuickMenu(fn chromiumEventOnRunQuickMenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnRunQuickMenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnSelectClientCertificate(fn chromiumEventOnSelectClientCertificate) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnSelectClientCertificate).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnSetFocus(fn chromiumEventOnSetFocus) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnSetFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnShowPermissionPrompt(fn chromiumEventOnShowPermissionPrompt) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnShowPermissionPrompt).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnSinks(fn chromiumEventOnSinks) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnSinks).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnStartDragging(fn chromiumEventOnStartDragging) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnStartDragging).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnStatusMessage(fn chromiumEventOnStatusMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnStatusMessage).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnTakeFocus(fn chromiumEventOnTakeFocus) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnTakeFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnTextResultAvailable(fn chromiumEventOnTextResultAvailable) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnTextResultAvailable).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnTextSelectionChanged(fn chromiumEventOnTextSelectionChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnTextSelectionChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnTooltip(fn chromiumEventOnTooltip) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnTooltip).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnTouchHandleStateChanged(fn chromiumEventOnTouchHandleStateChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnTouchHandleStateChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnUpdateDragCursor(fn chromiumEventOnUpdateDragCursor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnUpdateDragCursor).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnVirtualKeyboardRequested(fn chromiumEventOnVirtualKeyboardRequested) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnVirtualKeyboardRequested).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnIsChromeAppMenuItemVisible(fn chromiumEventOnIsChromeAppMenuItemVisible) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnIsChromeAppMenuItemVisible).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnIsChromeAppMenuItemEnabled(fn chromiumEventOnIsChromeAppMenuItemEnabled) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnIsChromeAppMenuItemEnabled).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnIsChromePageActionIconVisible(fn chromiumEventOnIsChromePageActionIconVisible) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnIsChromePageActionIconVisible).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFChromium) SetOnIsChromeToolbarButtonVisible(fn chromiumEventOnIsChromeToolbarButtonVisible) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOnIsChromeToolbarButtonVisible).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// --------TCEFChromium Event proc begin--------

func _CEFChromium_SetOnAfterCreated(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnAfterCreated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeClose
func _CEFChromium_SetOnBeforeClose(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnBeforeClose).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnClose
func _CEFChromium_SetOnClose(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnClose).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnPdfPrintFinished
func _CEFChromium_SetOnPdfPrintFinished(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnPdfPrintFinished).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnZoomPctAvailable
func _CEFChromium_SetOnZoomPctAvailable(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnZoomPctAvailable).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadStart
func _CEFChromium_SetOnLoadStart(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnLoadStart).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadingStateChange
func _CEFChromium_SetOnLoadingStateChange(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnLoadingStateChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadingProgressChange
func _CEFChromium_SetOnLoadingProgressChange(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnLoadingProgressChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadError
func _CEFChromium_SetOnLoadError(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnLoadError).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnLoadEnd
func _CEFChromium_SetOnLoadEnd(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnLoadEnd).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeDownload
func _CEFChromium_SetOnBeforeDownload(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnBeforeDownload).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnDownloadUpdated
func _CEFChromium_SetOnDownloadUpdated(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnDownloadUpdated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFullScreenModeChange
func _CEFChromium_SetOnFullScreenModeChange(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnFullScreenModeChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeBrowser
func _CEFChromium_SetOnBeforeBrowser(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnBeforeBrowse).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnAddressChange
func _CEFChromium_SetOnAddressChange(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnAddressChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnKeyEvent
func _CEFChromium_SetOnKeyEvent(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnKeyEvent).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnTitleChange
func _CEFChromium_SetOnTitleChange(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnTitleChange).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnRenderCompMsg
func _CEFChromium_SetOnRenderCompMsg(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnRenderCompMsg).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBrowserCompMsg
func _CEFChromium_SetOnBrowserCompMsg(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnBrowserCompMsg).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnRenderProcessTerminated
func _CEFChromium_SetOnRenderProcessTerminated(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnRenderProcessTerminated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnRenderViewReady
func _CEFChromium_SetOnRenderViewReady(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnRenderViewReady).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnScrollOffsetChanged
func _CEFChromium_SetOnScrollOffsetChanged(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnScrollOffsetChanged).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnProcessMessageReceived
func _CEFChromium_SetOnProcessMessageReceived(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnProcessMessageReceived).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFindResult
func _CEFChromium_SetOnFindResult(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnFindResult).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookieSet
func _CEFChromium_SetOnCookieSet(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnCookieSet).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookiesDeleted
func _CEFChromium_SetOnCookiesDeleted(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnCookiesDeleted).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookiesFlushed
func _CEFChromium_SetOnCookiesFlushed(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnCookiesFlushed).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookiesVisited
func _CEFChromium_SetOnCookiesVisited(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnCookiesVisited).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnCookieVisitorDestroyed
func _CEFChromium_SetOnCookieVisitorDestroyed(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnCookieVisitorDestroyed).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeContextMenu
func _CEFChromium_SetOnBeforeContextMenu(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnBeforeContextMenu).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnContextMenuCommand
func _CEFChromium_SetOnContextMenuCommand(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnContextMenuCommand).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnContextMenuDismissed
func _CEFChromium_SetOnContextMenuDismissed(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnContextMenuDismissed).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforeResourceLoad
func _CEFChromium_SetOnBeforeResourceLoad(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnBeforeResourceLoad).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnResourceResponse
func _CEFChromium_SetOnResourceResponse(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnResourceResponse).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnResourceRedirect
func _CEFChromium_SetOnResourceRedirect(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnResourceRedirect).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnResourceLoadComplete
func _CEFChromium_SetOnResourceLoadComplete(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnResourceLoadComplete).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFrameAttached
func _CEFChromium_SetOnFrameAttached(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnFrameAttached).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFrameCreated
func _CEFChromium_SetOnFrameCreated(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnFrameCreated).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnFrameDetached
func _CEFChromium_SetOnFrameDetached(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnFrameDetached).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnMainFrameChanged
func _CEFChromium_SetOnMainFrameChanged(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnMainFrameChanged).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnBeforePopup
func _CEFChromium_SetOnBeforePopup(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnBeforePopup).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnOpenUrlFromTab
func _CEFChromium_SetOnOpenUrlFromTab(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnOpenUrlFromTab).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnDragEnter
func _CEFChromium_SetOnDragEnter(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnDragEnter).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnDraggableRegionsChanged
func _CEFChromium_SetOnDraggableRegionsChanged(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnDraggableRegionsChanged).Call(instance, api.MakeEventDataPtr(fn))
}

// TCEFChromium _CEFChromium_SetOnGetAuthCredentials
func _CEFChromium_SetOnGetAuthCredentials(instance uintptr, fn interface{}) {
	imports.Proc(def.CEFChromium_SetOnGetAuthCredentials).Call(instance, api.MakeEventDataPtr(fn))
}

//--------TCEFChromium Event proc end--------
