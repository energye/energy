//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Chromium 事件函数定义

package cef

import (
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

/************* IChromium *************/

type chromiumEventOnResult func(sender lcl.IObject, aResultOK bool)      // 通用Result bool
type chromiumEventOnResultFloat func(sender lcl.IObject, result float64) // 通用Result float

type chromiumEventOnAcceleratedPaint func(sender lcl.IObject, browser *ICefBrowser, kind consts.TCefPaintElementType, dirtyRects *TCefRectArray, info TCefAcceleratedPaintInfo)
type chromiumEventOnAllConnectionsClosed func(sender lcl.IObject)
type chromiumEventOnAudioStreamError func(sender lcl.IObject, browser *ICefBrowser, message string)
type chromiumEventOnAudioStreamPacket func(sender lcl.IObject, browser *ICefBrowser, data *uintptr, frames int32, pts int64)
type chromiumEventOnAudioStreamStarted func(sender lcl.IObject, browser *ICefBrowser, params *TCefAudioParameters, channels int32)
type chromiumEventOnAudioStreamStopped func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnAutoResize func(sender lcl.IObject, browser *ICefBrowser, newSize *TCefSize) bool
type chromiumEventOnBeforeUnloadDialog func(sender lcl.IObject, browser *ICefBrowser, messageText string, isReload bool, callback *ICefJsDialogCallback) bool
type chromiumEventOnCanDownload func(sender lcl.IObject, browser *ICefBrowser, url, requestMethod string) bool
type chromiumEventOnCanSaveCookie func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, cookie *TCefCookie) bool
type chromiumEventOnCanSendCookie func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, cookie *TCefCookie) bool
type chromiumEventOnCertificateError func(sender lcl.IObject, browser *ICefBrowser, certError consts.TCefErrorCode, requestUrl string, sslInfo *ICefSslInfo, callback *ICefCallback) bool
type chromiumEventOnCertificateExceptionsCleared func(sender lcl.IObject)
type chromiumEventOnChromeCommand func(sender lcl.IObject, browser *ICefBrowser, commandId int32, disposition consts.TCefWindowOpenDisposition) bool
type chromiumEventOnConsoleMessage func(sender lcl.IObject, browser *ICefBrowser, level consts.TCefLogSeverity, message, source string, line int32) bool
type chromiumEventOnCursorChange func(sender lcl.IObject, browser *ICefBrowser, cursor consts.TCefCursorHandle, cursorType consts.TCefCursorType, customCursorInfo *TCefCursorInfo) bool
type chromiumEventOnDevToolsAgentAttached func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnDevToolsAgentDetached func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnDevTools func(sender lcl.IObject, browser *ICefBrowser, method string, params *ICefValue)
type chromiumEventOnDevToolsMessage func(sender lcl.IObject, browser *ICefBrowser, message *ICefValue) bool
type chromiumEventOnDevToolsMethodRawResult func(sender lcl.IObject, browser *ICefBrowser, messageId int32, success bool, result uintptr, resultSize uint32)
type chromiumEventOnDevToolsMethodResult func(sender lcl.IObject, browser *ICefBrowser, messageId int32, success bool, result *ICefValue)
type chromiumEventOnDevToolsRaw func(sender lcl.IObject, browser *ICefBrowser, method string, params uintptr, paramsSize uint32)
type chromiumEventOnDevToolsRawMessage func(sender lcl.IObject, browser *ICefBrowser, message uintptr, messageSize uint32) (handled bool)
type chromiumEventOnDialogClosed func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnDismissPermissionPrompt func(sender lcl.IObject, browser *ICefBrowser, promptId uint64, result consts.TCefPermissionRequestResult)
type chromiumEventOnDocumentAvailableInMainFrame func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnDownloadImageFinished func(sender lcl.IObject, imageUrl string, httpStatusCode int32, image *ICefImage)
type chromiumEventOnExecuteTaskOnCefThread func(sender lcl.IObject, taskID uint32)
type chromiumEventOnPrintStart func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnPrintSettings func(sender lcl.IObject, browser *ICefBrowser, settings *ICefPrintSettings, getDefaults bool)
type chromiumEventOnPrintDialog func(sender lcl.IObject, browser *ICefBrowser, hasSelection bool, callback *ICefPrintDialogCallback) bool
type chromiumEventOnPrintJob func(sender lcl.IObject, browser *ICefBrowser, documentName, PDFFilePath string, callback *ICefPrintJobCallback) bool
type chromiumEventOnPrintReset func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnGetPDFPaperSize func(sender lcl.IObject, browser *ICefBrowser, deviceUnitsPerInch int32, result *TCefSize)
type chromiumEventOnFavIconUrlChange func(sender lcl.IObject, browser *ICefBrowser, iconUrls []string) // TStrings => []string
type chromiumEventOnFileDialog func(sender lcl.IObject, browser *ICefBrowser, mode consts.FileDialogMode, title, defaultFilePath string, acceptFilters, acceptExtensions, acceptDescriptions *lcl.TStrings, callback *ICefFileDialogCallback) bool
type chromiumEventOnGetAccessibilityHandler func(sender lcl.IObject, accessibilityHandler *ICefAccessibilityHandler)
type chromiumEventOnGetAudioParameters func(sender lcl.IObject, browser *ICefBrowser, params *TCefAudioParameters) bool
type chromiumEventOnGetResourceHandler func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (resourceHandler *ICefResourceHandler)
type chromiumEventOnGetResourceHandlerEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, window IBrowserWindow) (resourceHandler *ICefResourceHandler, result bool)
type chromiumEventOnGetResourceRequestHandlerReqCtxHdlr func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, isNavigation, isDownload bool, requestInitiator string) (disableDefaultHandling bool, resourceRequestHandler *ICefResourceRequestHandler)
type chromiumEventOnGetResourceRequestHandlerReqHdlr func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, isNavigation, isDownload bool, requestInitiator string) (disableDefaultHandling bool, resourceRequestHandler *ICefResourceRequestHandler)
type chromiumEventOnGetResourceResponseFilter func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) (responseFilter *ICefResponseFilter)
type chromiumEventOnGetRootScreenRect func(sender lcl.IObject, browser *ICefBrowser) (rect *TCefRect, result bool)
type chromiumEventOnGetScreenInfo func(sender lcl.IObject, browser *ICefBrowser) (screenInfo *TCefScreenInfo, result bool)
type chromiumEventOnGetScreenPoint func(sender lcl.IObject, browser *ICefBrowser, viewX, viewY int32) (screenX, screenY int32, result bool)
type chromiumEventOnGetTouchHandleSize func(sender lcl.IObject, browser *ICefBrowser, orientation consts.TCefHorizontalAlignment) *TCefSize
type chromiumEventOnGetViewRect func(sender lcl.IObject, browser *ICefBrowser) *TCefRect
type chromiumEventOnGotFocus func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnHttpAuthCredentialsCleared func(sender lcl.IObject)
type chromiumEventOnIMECompositionRangeChanged func(sender lcl.IObject, browser *ICefBrowser, selectedRange TCefRange, characterBoundsCount uint32, characterBounds TCefRect)
type chromiumEventOnJsDialog func(sender lcl.IObject, browser *ICefBrowser, originUrl string, dialogType consts.TCefJsDialogType, messageText, defaultPromptText string, callback *ICefJsDialogCallback) (suppressMessage bool, result bool)
type chromiumEventOnMediaAccessChange func(sender lcl.IObject, browser *ICefBrowser, hasVideoAccess, hasAudioAccess bool)
type chromiumEventOnMediaRouteCreateFinished func(sender lcl.IObject, result consts.TCefMediaRouterCreateResult, error string, route *ICefMediaRoute) // TODO ICefMediaRoute
type chromiumEventOnMediaSinkDeviceInfo func(sender lcl.IObject, ipAddress string, port int32, modelName string)
type chromiumEventOnNavigationVisitorResultAvailable func(sender lcl.IObject, entry *ICefNavigationEntry, current bool, index, total int32) bool // TODO ICefNavigationEntry
type chromiumEventOnPaint func(sender lcl.IObject, browser *ICefBrowser, kind consts.TCefPaintElementType, dirtyRects *TCefRectArray, buffer uintptr, width, height int32)
type chromiumEventOnPdfPrintFinished func(sender lcl.IObject, ok bool)
type chromiumEventOnPopupShow func(sender lcl.IObject, browser *ICefBrowser, show bool)
type chromiumEventOnPopupSize func(sender lcl.IObject, browser *ICefBrowser, rect *TCefRect)
type chromiumEventOnPrefsAvailable func(sender lcl.IObject, resultOK bool)
type chromiumEventOnPrefsUpdated func(sender lcl.IObject)
type chromiumEventOnPreKey func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, osEvent consts.TCefEventHandle) (isKeyboardShortcut, result bool)
type chromiumEventOnProtocolExecution func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (allowOsExecution bool)
type chromiumEventOnQuickMenuCommand func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, commandId int32, eventFlags consts.TCefEventFlags) bool
type chromiumEventOnQuickMenuDismissed func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)
type chromiumEventOnRenderViewReady func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnRequestContextInitialized func(sender lcl.IObject, requestContext *ICefRequestContext)
type chromiumEventOnRequestMediaAccessPermission func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, requestingOrigin string, requestedPermissions uint32, callback *ICefMediaAccessCallback) bool
type chromiumEventOnResetDialogState func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnResolvedHostAvailable func(sender lcl.IObject, result consts.TCefErrorCode, resolvedIps []string) // []string => TStrings
type chromiumEventOnRouteMessageReceived func(sender lcl.IObject, route *ICefMediaRoute, message string)              // TODO ICefMediaRoute
type chromiumEventOnRoutes func(sender lcl.IObject, routes *TCefMediaRouteArray)                                      // TODO TCefMediaRouteArray
type chromiumEventOnRouteStateChanged func(sender lcl.IObject, route *ICefMediaRoute, state consts.TCefMediaRouteConnectionState)
type chromiumEventOnRunContextMenu func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel, callback *ICefRunContextMenuCallback) bool
type chromiumEventOnRunQuickMenu func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, location *TCefPoint, size *TCefSize, editStateFlags consts.TCefQuickMenuEditStateFlags, callback *ICefRunQuickMenuCallback) bool
type chromiumEventOnSelectClientCertificate func(sender lcl.IObject, browser *ICefBrowser, isProxy bool, host string, port int32, certificates *TCefX509CertificateArray, callback *ICefSelectClientCertificateCallback) bool
type chromiumEventOnSetFocus func(sender lcl.IObject, browser *ICefBrowser, source consts.TCefFocusSource) bool
type chromiumEventOnShowPermissionPrompt func(sender lcl.IObject, browser *ICefBrowser, promptId uint64, requestingOrigin string, requestedPermissions uint32, callback *ICefPermissionPromptCallback) bool
type chromiumEventOnSinks func(sender lcl.IObject, sinks *TCefMediaSinkArray) // TODO TCefMediaSinkArray
type chromiumEventOnStartDragging func(sender lcl.IObject, browser *ICefBrowser, dragData *ICefDragData, allowedOps consts.TCefDragOperations, x, y int32) bool
type chromiumEventOnStatusMessage func(sender lcl.IObject, browser *ICefBrowser, value string)
type chromiumEventOnTakeFocus func(sender lcl.IObject, browser *ICefBrowser, next bool)
type chromiumEventOnTextResultAvailable func(sender lcl.IObject, text string)
type chromiumEventOnTextSelectionChanged func(sender lcl.IObject, browser *ICefBrowser, selectedText string, selectedRange *TCefRange)
type chromiumEventOnTooltip func(sender lcl.IObject, browser *ICefBrowser, text *string) (result bool)
type chromiumEventOnTouchHandleStateChanged func(sender lcl.IObject, browser *ICefBrowser, state *TCefTouchHandleState)
type chromiumEventOnUpdateDragCursor func(sender lcl.IObject, browser *ICefBrowser, operation consts.TCefDragOperation)
type chromiumEventOnVirtualKeyboardRequested func(sender lcl.IObject, browser *ICefBrowser, inputMode consts.TCefTextInputMode)
type chromiumEventOnIsChromeAppMenuItemVisible func(sender lcl.IObject, browser *ICefBrowser, commandId int32) bool
type chromiumEventOnIsChromeAppMenuItemEnabled func(sender lcl.IObject, browser *ICefBrowser, commandId int32) bool
type chromiumEventOnIsChromePageActionIconVisible func(sender lcl.IObject, iconType consts.TCefChromePageActionIconType) bool
type chromiumEventOnIsChromeToolbarButtonVisible func(sender lcl.IObject, buttonType consts.TCefChromeToolbarButtonType) bool
type chromiumEventOnBeforeBrowser func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool) bool
type chromiumEventOnBeforeBrowserEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool, window IBrowserWindow) bool
type chromiumEventOnAddressChange func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, url string)
type chromiumEventOnTitleChange func(sender lcl.IObject, browser *ICefBrowser, title string)
type chromiumEventOnTitleChangeEx func(sender lcl.IObject, browser *ICefBrowser, title string, window IBrowserWindow)
type chromiumEventOnRenderProcessTerminated func(sender lcl.IObject, browser *ICefBrowser, status consts.TCefTerminationStatus, errorCode int32, error_ string)
type chromiumEventOnCompMsg func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, aHandled *bool)
type chromiumEventOnAfterCreated func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnAfterCreatedEx func(sender lcl.IObject, browser *ICefBrowser, window IBrowserWindow) bool
type chromiumEventOnBeforeClose func(sender lcl.IObject, browser *ICefBrowser)
type chromiumEventOnBeforeCloseEx func(sender lcl.IObject, browser *ICefBrowser, window IBrowserWindow) bool
type chromiumEventOnClose func(sender lcl.IObject, browser *ICefBrowser, aAction *consts.TCefCloseBrowserAction)
type chromiumEventOnCloseEx func(sender lcl.IObject, browser *ICefBrowser, aAction *consts.TCefCloseBrowserAction, window IBrowserWindow) bool
type chromiumEventOnScrollOffsetChanged func(sender lcl.IObject, browser *ICefBrowser, x, y float64)
type chromiumEventOnLoadStart func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, transitionType consts.TCefTransitionType)
type chromiumEventOnLoadStartEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, transitionType consts.TCefTransitionType, window IBrowserWindow)
type chromiumEventOnLoadingStateChange func(sender lcl.IObject, browser *ICefBrowser, isLoading, canGoBack, canGoForward bool)
type chromiumEventOnLoadingProgressChange func(sender lcl.IObject, browser *ICefBrowser, progress float64)
type chromiumEventOnLoadError func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, errorCode consts.CEF_NET_ERROR, errorText, failedUrl string)
type chromiumEventOnLoadEnd func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32)
type chromiumEventOnLoadEndEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32, window IBrowserWindow)
type chromiumEventOnCookieSet func(sender lcl.IObject, success bool, ID int32)
type chromiumEventOnCookiesDeleted func(sender lcl.IObject, numDeleted int32)
type chromiumEventOnCookiesFlushed func(sender lcl.IObject)
type chromiumEventOnCookiesVisited func(sender lcl.IObject, cookie *TCefCookie, deleteCookie, result *bool)
type chromiumEventOnCookieVisitorDestroyed func(sender lcl.IObject, ID int32)
type chromiumEventOnBeforeContextMenu func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel)
type chromiumEventOnBeforeContextMenuEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel, window IBrowserWindow) bool
type chromiumEventOnContextMenuCommand func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32) bool
type chromiumEventOnContextMenuCommandEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, window IBrowserWindow) bool
type chromiumEventOnContextMenuDismissed func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)
type chromiumEventOnFullScreenModeChange func(sender lcl.IObject, browser *ICefBrowser, fullscreen bool)
type chromiumEventOnBeforeDownload func(sender lcl.IObject, browser *ICefBrowser, downloadItem *ICefDownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) bool
type chromiumEventOnBeforeDownloadEx func(sender lcl.IObject, browser *ICefBrowser, downloadItem *ICefDownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback, window IBrowserWindow) bool
type chromiumEventOnDownloadUpdated func(sender lcl.IObject, browser *ICefBrowser, downloadItem *ICefDownloadItem, callback *ICefDownloadItemCallback)
type chromiumEventOnKey func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, osEvent consts.TCefEventHandle, result *bool)
type chromiumEventOnKeyEventEx func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, osEvent consts.TCefEventHandle, window IBrowserWindow, result *bool)
type chromiumEventOnBeforeResourceLoad func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue)
type chromiumEventOnBeforeResourceLoadEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue, window IBrowserWindow)
type chromiumEventOnResourceResponse func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, result *bool)
type chromiumEventOnResourceRedirect func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, newUrl *string)
type chromiumEventOnResourceLoadComplete func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, status consts.TCefUrlRequestStatus, receivedContentLength int64)
type chromiumEventOnFindResult func(sender lcl.IObject, browser *ICefBrowser, identifier, count int32, selectionRect *TCefRect, activeMatchOrdinal int32, finalUpdate bool)
type chromiumEventOnFrameAttached func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, reattached bool)
type chromiumEventOnFrameCreated func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)
type chromiumEventOnFrameDetached func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)
type chromiumEventOnMainFrameChanged func(sender lcl.IObject, browser *ICefBrowser, oldFrame *ICefFrame, newFrame *ICefFrame)
type chromiumEventOnMainFrameChangedEx func(sender lcl.IObject, browser *ICefBrowser, oldFrame *ICefFrame, newFrame *ICefFrame, window IBrowserWindow)
type chromiumEventOnBeforePopup func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, popupFeatures *TCefPopupFeatures, windowInfo *TCefWindowInfo, resultClient *ICefClient, settings *TCefBrowserSettings, resultExtraInfo *ICefDictionaryValue, noJavascriptAccess *bool) bool
type chromiumEventOnBeforePopupEx func(sender lcl.IObject, popupWindow IBrowserWindow, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, popupFeatures *TCefPopupFeatures, windowInfo *TCefWindowInfo, resultClient *ICefClient, settings *TCefBrowserSettings, resultExtraInfo *ICefDictionaryValue, noJavascriptAccess *bool) bool
type chromiumEventOnOpenUrlFromTab func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool
type chromiumEventOnDragEnter func(sender lcl.IObject, browser *ICefBrowser, dragData *ICefDragData, mask consts.TCefDragOperations, result *bool)
type chromiumEventOnDragEnterEx func(sender lcl.IObject, browser *ICefBrowser, dragData *ICefDragData, mask consts.TCefDragOperations, window IBrowserWindow, result *bool)
type chromiumEventOnDraggableRegionsChanged func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, regions *TCefDraggableRegions)
type chromiumEventOnDraggableRegionsChangedEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, regions *TCefDraggableRegions, window IBrowserWindow)
type chromiumEventOnGetAuthCredentials func(sender lcl.IObject, browser *ICefBrowser, originUrl string, isProxy bool, host string, port int32, realm, scheme string, callback *ICefAuthCallback) bool

/************* ProcessMessageReceived *************/

type BrowseProcessMessageReceived func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool
type BrowseProcessMessageReceivedEx func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage, window IBrowserWindow) bool
type RenderProcessMessageReceived func(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool

/************* TCEFApplication *************/

type GlobalCEFAppEventOnRegCustomSchemes func(registrar *TCefSchemeRegistrarRef)
type GlobalCEFAppEventOnRegisterCustomPreferences func(type_ consts.TCefPreferencesType, registrar *TCefPreferenceRegistrarRef)
type GlobalCEFAppEventOnContextInitialized func()
type GlobalCEFAppEventOnBeforeChildProcessLaunch func( /*commandLine *TCefCommandLine,*/ commandLine *ICefCommandLine)
type GlobalCEFAppEventOnAlreadyRunningAppRelaunchEvent func(commandLine *ICefCommandLine, currentDirectory string) bool
type GlobalCEFAppEventOnGetDefaultClient func(client *ICefClient)
type GlobalCEFAppEventOnGetLocalizedString func(stringId int32, stringVal *ResultString, result *ResultBool)
type GlobalCEFAppEventOnGetDataResource func(resourceId int32, data *ResultBytes, result *ResultBool)
type GlobalCEFAppEventOnGetDataResourceForScale func(resourceId int32, scaleFactor consts.TCefScaleFactor, data *ResultBytes, result *ResultBool)
type GlobalCEFAppEventOnWebKitInitialized func()
type GlobalCEFAppEventOnBrowserCreated func(browser *ICefBrowser, extraInfo *ICefDictionaryValue)
type GlobalCEFAppEventOnBrowserDestroyed func(browser *ICefBrowser)
type GlobalCEFAppEventOnContextCreated func(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool
type GlobalCEFAppEventOnContextReleased func(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context)
type GlobalCEFAppEventOnUncaughtException func(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context, exception *ICefV8Exception, stackTrace *ICefV8StackTrace)
type GlobalCEFAppEventOnFocusedNodeChanged func(browser *ICefBrowser, frame *ICefFrame, node *ICefDomNode)
type GlobalCEFAppEventOnRenderLoadingStateChange func(browser *ICefBrowser, frame *ICefFrame, isLoading, canGoBack, canGoForward bool)
type GlobalCEFAppEventOnRenderLoadStart func(browser *ICefBrowser, frame *ICefFrame, transitionType consts.TCefTransitionType)
type GlobalCEFAppEventOnRenderLoadEnd func(browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32)
type GlobalCEFAppEventOnRenderLoadError func(browser *ICefBrowser, frame *ICefFrame, errorCode consts.TCefErrorCode, errorText, failedUrl string)
type GlobalCEFAppEventOnScheduleMessagePumpWork func(delayMS int64)

/************* LCL Window event *************/

type TCloseEvent func(sender lcl.IObject, action *types.TCloseAction) bool
type TNotifyEvent func(sender lcl.IObject) bool
type TCloseQueryEvent func(sender lcl.IObject, canClose *bool) bool
