//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// CEF 事件定义

// 事件方法参数说明
// 1. 参数类型: 带有星(*)号, 指针传递, 一搬情况可通过指针直接修改该入参值, 结构修改字段值, 基础类型赋值
// 2. 参数名: out开头参数名表示该参数只具有返回功能

package cef

import (
	"github.com/energye/energy/v2/lcl"
)

// ==== Events 代理对象接口 ====

type IChromiumEvents = IChromium
type ICefViewDelegateEvents = ICEFViewComponent
type ICefBrowserViewDelegateEvents = ICEFBrowserViewComponent
type ICefButtonDelegateEvents = ICEFButtonComponent
type ICEFUrlRequestClientEvents = ICEFUrlRequestClientComponent
type ICefMenuButtonDelegateEvents = ICEFMenuButtonComponent
type ICefPanelDelegateEvents = ICEFPanelComponent
type IServerEvents = ICEFServerComponent
type ICefTextfieldDelegateEvents = ICEFTextfieldComponent
type ICefWindowDelegateEvents = ICEFWindowComponent

// ==== 事件方法定义 ====

// LCL 事件别名

type TNotify = lcl.TNotifyEvent
type TOnCloseQuery = lcl.TCloseQueryEvent

// 通用 bool

type TOnBool func(sender IObject, value bool)

// 通用 float64

type TOnFloat func(sender IObject, value float64)

// 通用 int32

type TOnInt32 func(sender IObject, value int32)

// 通用 browser

type TOnBrowser func(sender IObject, browser ICefBrowser)

// ICefRenderHandler

type TOnGetAccessibilityHandler func(sender IObject, accessibilityHandler IAccessibilityHandler)
type TOnGetRootScreenRect func(sender IObject, browser ICefBrowser, rect *TCefRect, outResult *bool)
type TOnGetViewRect func(sender IObject, browser ICefBrowser, rect *TCefRect)
type TOnGetScreenPoint func(sender IObject, browser ICefBrowser, viewX, viewY int32, screenX, screenY *int32, outResult *bool)
type TOnGetScreenInfo func(sender IObject, browser ICefBrowser, screenInfo *TCefScreenInfo, outResult *bool)
type TOnPopupShow func(sender IObject, browser ICefBrowser, show bool)
type TOnPopupSize func(sender IObject, browser ICefBrowser, rect TCefRect)
type TOnPaint func(sender IObject, browser ICefBrowser, kind TCefPaintElementType, dirtyRects ICefRectArray, buffer uintptr, width, height int32)
type TOnAcceleratedPaint func(sender IObject, browser ICefBrowser, kind TCefPaintElementType, dirtyRects ICefRectArray, sharedHandle uintptr)
type TOnGetTouchHandleSize func(sender IObject, browser ICefBrowser, orientation TCefHorizontalAlignment, size *TCefSize)
type TOnTouchHandleStateChanged func(sender IObject, browser ICefBrowser, state TCefTouchHandleState)
type TOnStartDragging func(sender IObject, browser ICefBrowser, dragData ICefDragData, allowedOps TCefDragOperations, x, y int32, outResult *bool)
type TOnUpdateDragCursor func(sender IObject, browser ICefBrowser, operation TCefDragOperation)
type TOnScrollOffsetChanged func(sender IObject, browser ICefBrowser, x, y float64)
type TOnIMECompositionRangeChanged func(sender IObject, browser ICefBrowser, selectedRange TCefRange, characterBoundsCount uint32, characterBounds TCefRect)
type TOnTextSelectionChanged func(sender IObject, browser ICefBrowser, selectedText string, selectedRange TCefRange)
type TOnVirtualKeyboardRequested func(sender IObject, browser ICefBrowser, inputMode TCefTextInpuMode)

// ICefDragHandler

type TOnDragEnter func(sender IObject, browser ICefBrowser, dragData ICefDragData, mask TCefDragOperations, outResult *bool)
type TOnDragEnterEx func(window IBrowserWindow, browser ICefBrowser, dragData ICefDragData, mask TCefDragOperations, outResult *bool)
type TOnDraggableRegionsChanged func(sender IObject, browser ICefBrowser, frame ICefFrame, regions ICefDraggableRegion)
type TOnDraggableRegionsChangedEx func(window IBrowserWindow, browser ICefBrowser, frame ICefFrame, regions ICefDraggableRegion)

// ICefFindHandler

type TOnFindResult func(sender IObject, browser ICefBrowser, identifier, count int32, selectionRect TCefRect, activeMatchOrdinal int32, finalUpdate bool)

// ICefRequestContextHandler
//  uses the same TOnGetResourceRequestHandler event type defined for ICefRequestHandler

type TOnRequestContextInitialized func(sender IObject, requestContext ICefRequestContext)

// ICefMediaObserver

type TOnSinks func(sender IObject, sinks ICefMediaSinkArray)
type TOnRoutes func(sender IObject, routes ICefMediaRouteArray)
type TOnRouteStateChanged func(sender IObject, route ICefMediaRoute, state TCefMediaRouteConnectionState)
type TOnRouteMessageReceived func(sender IObject, route ICefMediaRoute, message string)

// ICefAudioHandler

type TOnGetAudioParameters func(sender IObject, browser ICefBrowser, params *TCefAudioParameters, result *bool)
type TOnAudioStreamStarted func(sender IObject, browser ICefBrowser, params TCefAudioParameters, channels int32)
type TOnAudioStreamPacket func(sender IObject, browser ICefBrowser, data PPSingle, frames int32, pts int64)
type TOnAudioStreamStopped = TOnBrowser
type TOnAudioStreamError func(sender IObject, browser ICefBrowser, message string)

// ICefDevToolsMessageObserver

type TOnDevToolsMessage func(sender IObject, browser ICefBrowser, message ICefValue, handled *bool)
type TOnDevToolsRawMessage func(sender IObject, browser ICefBrowser, message uintptr, messageSize uint32, handled *bool)
type TOnDevToolsMethodResult func(sender IObject, browser ICefBrowser, messageId int32, success bool, result ICefValue)
type TOnDevToolsMethodRawResult func(sender IObject, browser ICefBrowser, messageId int32, success bool, result uintptr, resultSize uint32)
type TOnDevToolsEvent func(sender IObject, browser ICefBrowser, method string, params ICefValue)
type TOnDevToolsRawEvent func(sender IObject, browser ICefBrowser, method string, params uintptr, paramsSize uint32)
type TOnDevToolsAgentAttached = TOnBrowser
type TOnDevToolsAgentDetached = TOnBrowser

// ICefExtensionHandler

type TOnExtensionLoadFailed func(sender IObject, result TCefErrorCode)
type TOnExtensionLoaded func(sender IObject, extension ICefExtension)
type TOnExtensionUnloaded func(sender IObject, extension ICefExtension)
type TOnBeforeBackgroundBrowser func(sender IObject, extension ICefExtension, url string, settings *TCefBrowserSettings) (client ICefClient, result bool)
type TOnBeforeBrowser func(sender IObject, extension ICefExtension, browser, activeBrowser ICefBrowser, index int32, url string, active bool,
	windowInfo *TCefWindowInfo, settings *TCefBrowserSettings) (client ICefClient, result bool)
type TOnGetActiveBrowser func(sender IObject, extension ICefExtension, browser ICefBrowser, includeIncognito bool) (activeBrowser ICefBrowser)
type TOnCanAccessBrowser func(sender IObject, extension ICefExtension, browser ICefBrowser, includeIncognito bool, targetBrowser ICefBrowser, result *bool)
type TOnGetExtensionResource func(sender IObject, extension ICefExtension, browser ICefBrowser, file string, callback ICefGetExtensionResourceCallback, result *bool)

// ICefPrintHandler

type TOnPrintStart = TOnBrowser
type TOnPrintSettings func(sender IObject, browser ICefBrowser, settings ICefPrintSettings, getDefaults bool)
type TOnPrintDialog func(sender IObject, browser ICefBrowser, hasSelection bool, callback ICefPrintDialogCallback, result *bool)
type TOnPrintJob func(sender IObject, browser ICefBrowser, documentName, PDFFilePath string, callback ICefPrintJobCallback, result *bool)
type TOnPrintReset = TOnBrowser
type TOnGetPDFPaperSize func(sender IObject, browser ICefBrowser, deviceUnitsPerInch int32, result *TCefSize)

// ICefFrameHandler

type TOnFrameCreated func(sender IObject, browser ICefBrowser, frame ICefFrame)
type TOnFrameAttached func(sender IObject, browser ICefBrowser, frame ICefFrame, reattached bool)
type TOnFrameDetached func(sender IObject, browser ICefBrowser, frame ICefFrame)
type TOnMainFrameChanged func(sender IObject, browser ICefBrowser, oldFrame ICefFrame, newFrame ICefFrame)
type TOnMainFrameChangedEx func(window IBrowserWindow, browser ICefBrowser, oldFrame ICefFrame, newFrame ICefFrame)

// ICefCommandHandler

type TOnChromeCommand func(sender IObject, browser ICefBrowser, commandId int32, disposition TCefWindowOpenDisposition, result *bool)
type TOnIsChromeAppMenuItemVisible func(sender IObject, browser ICefBrowser, commandId int32, result *bool)
type TOnIsChromeAppMenuItemEnabled func(sender IObject, browser ICefBrowser, commandId int32, result *bool)
type TOnIsChromePageActionIconVisible func(sender IObject, iconType TCefChromePageActionIconType, result *bool)
type TOnIsChromeToolbarButtonVisible func(sender IObject, buttonType TCefChromeToolbarButtonType, result *bool)

// ICefPermissionHandler

type TOnRequestMediaAccessPermission func(sender IObject, browser ICefBrowser, frame ICefFrame, requestingOrigin string, requestedPermissions uint32, callback ICefMediaAccessCallback, result *bool)
type TOnShowPermissionPrompt func(sender IObject, browser ICefBrowser, promptId uint64, requestingOrigin string, requestedPermissions uint32, callback ICefPermissionPromptCallback, result *bool)
type TOnDismissPermissionPrompt func(sender IObject, browser ICefBrowser, promptId uint64, result TCefPermissionRequestResult)

// TCustomXxx

type TOnTextResultAvailable func(sender IObject, text string)
type TOnPdfPrintFinished = TOnBool
type TOnPrefsAvailable = TOnBool
type TOnCookiesDeleted = TOnInt32
type TOnResolvedIPsAvailable func(sender IObject, result TCefErrorCode, resolvedIps IStrings)
type TOnNavigationVisitorResultAvailable func(sender IObject, entry ICefNavigationEntry, current bool, index, total int32, result *bool)
type TOnDownloadImageFinished func(sender IObject, imageUrl string, httpStatusCode int32, image ICefImage)
type TOnExecuteTaskOnCefThread func(sender IObject, taskID uint32)
type TOnCookiesVisited func(sender IObject, cookie TCookie, count, total, id int32, deleteCookie, result *bool) //name, value, domain, path string, secure, httponly, hasExpires bool, creation, lastAccess, expires, TDateTime, total, aID int32, same_site TCefCookieSameSite, priority TCefCookiePriority, deleteCookie, result *bool
type TOnCookieVisitorDestroyed = TOnInt32
type TOnCookieSet func(sender IObject, success bool, ID int32)
type TOnZoomPctAvailable = TOnFloat
type TOnMediaRouteCreateFinished func(sender IObject, result TCefMediaRouterCreateResult, error string, route ICefMediaRoute)
type TOnMediaSinkDeviceInfo func(sender IObject, ipAddress string, port int32, modelName string)

// Windows

type TOnCompMsg func(sender IObject, message *TMessage, lResult *LRESULT, handled *bool)

// ICefLoadHandler

type TOnLoadStart func(sender IObject, browser ICefBrowser, frame ICefFrame, transitionType TCefTransitionType)
type TOnLoadEnd func(sender IObject, browser ICefBrowser, frame ICefFrame, httpStatusCode int32)
type TOnLoadEndEx func(window IBrowserWindow, browser ICefBrowser, frame ICefFrame, httpStatusCode int32)
type TOnLoadError func(sender IObject, browser ICefBrowser, frame ICefFrame, errorCode TCefErrorCode, errorText, failedUrl string)
type TOnLoadingStateChange func(sender IObject, browser ICefBrowser, isLoading, canGoBack, canGoForward bool)

// ICefFocusHandler

type TOnTakeFocus func(sender IObject, browser ICefBrowser, next bool)
type TOnSetFocus func(sender IObject, browser ICefBrowser, source TCefFocusSource, outResult *bool)
type TOnGotFocus = TOnBrowser

// ICefContextMenuHandler

type TOnBeforeContextMenu func(sender IObject, browser ICefBrowser, frame ICefFrame, params ICefContextMenuParams, model ICefMenuModel)
type TOnBeforeContextMenuEx func(window IBrowserWindow, browser ICefBrowser, frame ICefFrame, params ICefContextMenuParams, model ICefMenuModel)
type TOnRunContextMenu func(sender IObject, browser ICefBrowser, frame ICefFrame, params ICefContextMenuParams, model ICefMenuModel, callback ICefRunContextMenuCallback, result *bool)
type TOnContextMenuCommand func(sender IObject, browser ICefBrowser, frame ICefFrame, params ICefContextMenuParams, commandId MenuId, eventFlags uint32, outResult *bool)
type TOnContextMenuCommandEx func(window IBrowserWindow, browser ICefBrowser, frame ICefFrame, params ICefContextMenuParams, commandId MenuId, eventFlags uint32, outResult *bool)
type TOnContextMenuDismissed func(sender IObject, browser ICefBrowser, frame ICefFrame)
type TOnRunQuickMenu func(sender IObject, browser ICefBrowser, frame ICefFrame, location *TCefPoint, size *TCefSize, editStateFlags TCefQuickMenuEditStateFlags, callback ICefRunQuickMenuCallback, result *bool)
type TOnQuickMenuCommand func(sender IObject, browser ICefBrowser, frame ICefFrame, commandId int32, eventFlags TCefEventFlags, result *bool)
type TOnQuickMenuDismissed func(sender IObject, browser ICefBrowser, frame ICefFrame)

// ICefKeyboardHandler

type TOnPreKey func(sender IObject, browser ICefBrowser, event TCefKeyEvent, osEvent TCefEventHandle, outIsKeyboardShortcut, outResult *bool)
type TOnKey func(sender IObject, browser ICefBrowser, event TCefKeyEvent, osEvent TCefEventHandle, outResult *bool)
type TOnKeyEx func(window IBrowserWindow, browser ICefBrowser, event TCefKeyEvent, osEvent TCefEventHandle, outResult *bool)

// ICefDisplayHandler

type TOnAddressChange func(sender IObject, browser ICefBrowser, frame ICefFrame, url string)
type TOnTitleChange func(sender IObject, browser ICefBrowser, title string)
type TOnTitleChangeEx func(window IBrowserWindow, browser ICefBrowser, title string)
type TOnFavIconUrlChange func(sender IObject, browser ICefBrowser, iconUrls IStrings)
type TOnFullScreenModeChange func(sender IObject, browser ICefBrowser, fullscreen bool)
type TOnTooltip func(sender IObject, browser ICefBrowser, text *string, outResult *bool)
type TOnStatusMessage func(sender IObject, browser ICefBrowser, value string)
type TOnConsoleMessage func(sender IObject, browser ICefBrowser, level TCefLogSeverity, message, source string, line int32, outResult *bool)
type TOnAutoResize func(sender IObject, browser ICefBrowser, newSize TCefSize, outResult *bool)
type TOnLoadingProgressChange func(sender IObject, browser ICefBrowser, progress float64)
type TOnCursorChange func(sender IObject, browser ICefBrowser, cursor TCefCursorHandle, cursorType TCefCursorType, customCursorInfo TCefCursorInfo, result *bool)
type TOnMediaAccessChange func(sender IObject, browser ICefBrowser, hasVideoAccess, hasAudioAccess bool)

// ICefDownloadHandler

type TOnCanDownload func(sender IObject, browser ICefBrowser, url, requestMethod string, result *bool)
type TOnBeforeDownload func(sender IObject, browser ICefBrowser, downloadItem ICefDownloadItem, suggestedName string, callback ICefBeforeDownloadCallback)
type TOnBeforeDownloadEx func(window IBrowserWindow, browser ICefBrowser, downloadItem ICefDownloadItem, suggestedName string, callback ICefBeforeDownloadCallback)
type TOnDownloadUpdated func(sender IObject, browser ICefBrowser, downloadItem ICefDownloadItem, callback ICefDownloadItemCallback)

// ICefJsDialogHandler

type TOnJsdialog func(sender IObject, browser ICefBrowser, originUrl string, dialogType TCefJsDialogType, messageText, defaultPromptText string, callback ICefJsDialogCallback, outSuppressMessage, outResult *bool)
type TOnBeforeUnloadDialog func(sender IObject, browser ICefBrowser, messageText string, isReload bool, callback ICefJsDialogCallback, outResult *bool)
type TOnResetDialogState = TOnBrowser
type TOnDialogClosed = TOnBrowser

// ICefLifeSpanHandler

type TOnBeforePopup func(sender IObject, browser ICefBrowser, frame ICefFrame, beforePopup TBeforePopup, popupFeatures TCefPopupFeatures, windowInfo *TCefWindowInfo, settings *TCefBrowserSettings) (client ICefClient, extraInfo ICefDictionaryValue, noJavascriptAccess, result bool)
type TOnAfterCreated = TOnBrowser
type TOnBeforeClose = TOnBrowser
type TOnClose func(sender IObject, browser ICefBrowser, action *TCefCloseBrowserAction)

// ICefRequestHandler

type TOnBeforeBrowse func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, userGesture, isRedirect bool, outResult *bool)
type TOnOpenUrlFromTab func(sender IObject, browser ICefBrowser, frame ICefFrame, targetUrl string, targetDisposition TCefWindowOpenDisposition, userGesture bool, outResult *bool)
type TOnGetAuthCredentials func(sender IObject, browser ICefBrowser, originUrl string, isProxy bool, host string, port int32, realm, scheme string, callback ICefAuthCallback, outResult *bool)
type TOnCertificateError func(sender IObject, browser ICefBrowser, certError TCefErrorCode, requestUrl string, sslInfo ICefSslInfo, callback ICefCallback, outResult *bool)
type TOnSelectClientCertificate func(sender IObject, browser ICefBrowser, isProxy bool, host string, port int32, certificates ICefX509CertificateArray, callback ICefSelectClientCertificateCallback, result *bool)
type TOnRenderViewReady = TOnBrowser
type TOnRenderProcessTerminated func(sender IObject, browser ICefBrowser, status TCefTerminationStatus)
type TOnGetResourceRequestHandler func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, isNavigation, isDownload bool, requestInitiator string, disableDefaultHandling *bool) (resourceRequestHandler ICefResourceRequestHandler)
type TOnDocumentAvailableInMainFrame = TOnBrowser

// ICefResourceRequestHandler

type TOnBeforeResourceLoad func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, callback ICefCallback, result *TCefReturnValue)
type TOnBeforeResourceLoadEx func(window IBrowserWindow, browser ICefBrowser, frame ICefFrame, request ICefRequest, callback ICefCallback, result *TCefReturnValue)
type TOnGetResourceHandler func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest) (resourceHandler ICefResourceHandler)
type TOnGetResourceHandlerEx func(window IBrowserWindow, browser ICefBrowser, frame ICefFrame, request ICefRequest) (resourceHandler ICefResourceHandler, flag bool)
type TOnResourceRedirect func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, response ICefResponse, newUrl *string)
type TOnResourceResponse func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, response ICefResponse, outResult *bool)
type TOnGetResourceResponseFilter func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, response ICefResponse) (responseFilter ICefResponseFilter)
type TOnResourceLoadComplete func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, response ICefResponse, status TCefUrlRequestStatus, receivedContentLength int64)
type TOnProtocolExecution func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, allowOsExecution *bool)

// ICefCookieAccessFilter

type TOnCanSendCookie func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, cookie TCefCookie, result *bool)
type TOnCanSaveCookie func(sender IObject, browser ICefBrowser, frame ICefFrame, request ICefRequest, response ICefResponse, cookie TCefCookie, result *bool)

// ICefDialogHandler

type TOnFileDialog func(sender IObject, browser ICefBrowser, mode TCefFileDialogMode, title, defaultFilePath string, acceptFilters IStrings, callback ICefFileDialogCallback, result *bool)

// ICefClient

type TOnProcessMessageReceived func(browser ICefBrowser, frame ICefFrame, sourceProcess TCefProcessId, message ICefProcessMessage, outResult *bool)
type TOnProcessMessageReceivedEx func(window IBrowserWindow, browser ICefBrowser, frame ICefFrame, sourceProcess TCefProcessId, message ICefProcessMessage, outResult *bool)

// ICefApp

type TOnRegisterCustomSchemes func(registrar ICefSchemeRegistrarRef)

// ICefBrowserProcessHandler

type TOnRegisterCustomPreferences func(type_ TCefPreferencesType, registrar ICefPreferenceRegistrarRef)
type TOnContextInitialized func()
type TOnBeforeChildProcessLaunch func(commandLine ICefCommandLine)
type TOnScheduleMessagePumpWork func(delayMS int64)
type TOnGetDefaultClient func(client ICefClient)

// ICefResourceBundleHandler

type TOnGetLocalizedString func(stringId int32) (outVal string, result bool)
type TOnGetDataResource func(resourceId int32) (data []byte, result bool)
type TOnGetDataResourceForScale func(resourceId int32, scaleFactor TCefScaleFactor) (data []byte, result bool)

// ICefRenderProcessHandler

type TOnWebKitInitialized func()
type TOnBrowserCreated func(browser ICefBrowser, extraInfo ICefDictionaryValue)
type TOnBrowserDestroyed func(browser ICefBrowser)
type TOnContextCreated func(browser ICefBrowser, frame ICefFrame, context ICefv8Context)
type TOnContextReleased func(browser ICefBrowser, frame ICefFrame, context ICefv8Context)
type TOnUncaughtException func(browser ICefBrowser, frame ICefFrame, context ICefv8Context, exception ICefV8Exception, stackTrace ICefV8StackTrace)
type TOnFocusedNodeChanged func(browser ICefBrowser, frame ICefFrame, node ICefDomNode)

// ICefLoadHandler

type TOnRenderLoadingStateChange func(browser ICefBrowser, frame ICefFrame, isLoading, canGoBack, canGoForward bool)
type TOnRenderLoadStart func(browser ICefBrowser, frame ICefFrame, transitionType TCefTransitionType)
type TOnRenderLoadEnd func(browser ICefBrowser, frame ICefFrame, httpStatusCode int32)
type TOnRenderLoadError func(browser ICefBrowser, frame ICefFrame, errorCode TCefErrorCode, errorText, failedUrl string)

// TCEFWindowComponent

type TOnWindow func(sender IObject, window ICefWindow)
type TOnWindowBool func(sender IObject, window ICefWindow, result *bool)
type TOnWindowCreated = TOnWindow
type TOnWindowClosing = TOnWindow
type TOnWindowDestroyed = TOnWindow
type TOnWindowActivationChanged func(sender IObject, window ICefWindow, active bool)
type TOnWindowBoundsChanged func(sender IObject, window ICefWindow, newBounds TCefRect)
type TOnGetParentWindow func(sender IObject, window ICefWindow, isMenu, canActivateMenu *bool) ICefWindow
type TOnIsWindowModalDialog = TOnWindowBool
type TOnGetInitialBounds func(sender IObject, window ICefWindow, result *TCefRect)
type TOnGetInitialShowState func(sender IObject, window ICefWindow, result *TCefShowState)
type TOnIsFrameless = TOnWindowBool
type TOnWithStandardWindowButtons = TOnWindowBool
type TOnGetTitlebarHeight func(sender IObject, window ICefWindow, titleBarHeight *float32, result *bool)
type TOnCanResize = TOnWindowBool
type TOnCanMaximize = TOnWindowBool
type TOnCanMinimize = TOnWindowBool
type TOnCanClose = TOnWindowBool
type TOnAccelerator func(sender IObject, window ICefWindow, commandId int32, result *bool)
type TOnWindowKeyEvent func(sender IObject, window ICefWindow, event TCefKeyEvent, result *bool)
type TOnWindowFullscreenTransition func(sender IObject, window ICefWindow, isCompleted bool)

// TCEFBrowserViewComponent

type TOnBrowserCreatedBvc func(sender IObject, browserView ICefBrowserView, browser ICefBrowser)
type TOnBrowserDestroyedBvc func(sender IObject, browserView ICefBrowserView, browser ICefBrowser)
type TOnGetDelegateForPopupBrowserViewBvc func(sender IObject, browserView ICefBrowserView, browserSettings TCefBrowserSettings, client ICefClient, isDevtools bool) (result ICefBrowserViewDelegate)
type TOnPopupBrowserViewCreatedBvc func(sender IObject, browserView, popupBrowserView ICefBrowserView, isDevtools bool, result *bool)
type TOnGetChromeToolbarTypeBvc func(sender IObject, result *TCefChromeToolbarType)
type TOnUseFramelessWindowForPictureInPictureBvc func(sender IObject, browserView ICefBrowserView, result *bool)
type TOnGestureCommandBvc func(sender IObject, browserView ICefBrowserView, gestureCommand TCefGestureCommand, result *bool)

// TAccessibilityHandler

type TOnAccessibility func(sender IObject, value ICefValue)

// TBufferPanel

type TOnIMECommitText func(sender IObject, text string, replacementRange TCefRange, relativeCursorPos int32)
type TOnIMESetComposition func(sender IObject, text string, underlines ICefCompositionUnderlineArray, replacementRange, selectionRange TCefRange)
type TOnHandledMessage func(sender IObject, message *TMessage, lResult *LRESULT, handled *bool) // Windows
type TConstrainedResize = lcl.TConstrainedResizeEvent
type TContextPopup = lcl.TContextPopupEvent
type TDragDrop = lcl.TDragDropEvent
type TEndDrag = lcl.TEndDragEvent
type TDragOver = lcl.TDragOverEvent
type TGetSiteInfo = lcl.TGetSiteInfoEvent
type TMouse = lcl.TMouseEvent
type TMouseMove = lcl.TMouseMoveEvent
type TMouseWheel = lcl.TMouseWheelEvent
type TStartDock = lcl.TStartDockEvent
type TStartDrag = lcl.TStartDragEvent

// ICefButtonDelegate

type TOnButtonPressed func(sender IObject, button ICefButton)
type TOnButtonStateChanged func(sender IObject, button ICefButton)

// TCefResponseFilter

type TOnInitFilter func(sender IObject, result *bool)
type TOnFilter func(sender IObject, dataIn uintptr, dataInSize uint32, dataInRead *uint32, dataOut uintptr, dataOutSize uint32, dataOutWritten *uint32, result *TCefResponseFilterStatus)

// TDomVisitor

type TOnDomVisitor func(document ICefDomDocument)

// ICefMenuButtonDelegate

type TOnMenuButtonPressed func(sender IObject, menuButton ICefMenuButton, screenPoint TCefPoint, buttonPressedLock ICefMenuButtonPressedLock)

// TMenuModelDelegate

type TOnExecuteCommand func(menuModel ICefMenuModel, commandId int32, eventFlags TCefEventFlags)
type TOnMouseOutsideMenu func(menuModel ICefMenuModel, screenPoint TCefPoint)
type TOnUnhandledOpenSubmenu func(menuModel ICefMenuModel, isRTL bool)
type TOnUnhandledCloseSubmenu func(menuModel ICefMenuModel, isRTL bool)
type TOnMenuWillShow func(menuModel ICefMenuModel)
type TOnMenuClosed func(menuModel ICefMenuModel)
type TOnFormatLabel func(menuModel ICefMenuModel, label *string, outResult *bool)

// TResourceHandler

type TOnResourceHandlerOpen func(request ICefRequest, handleRequest *bool, callback ICefCallback, outResult *bool)
type TOnResourceHandlerProcessRequest func(request ICefRequest, callback ICefCallback, outResult *bool) // deprecated
type TOnResourceHandlerGetResponseHeaders func(response ICefResponse, outResponseLength *int64, outRedirectUrl *string)
type TOnResourceHandlerSkip func(bytesToSkip int64, bytesSkipped *int64, callback ICefResourceSkipCallback, outResult *bool)
type TOnResourceHandlerRead func(dataOut uintptr, bytesToRead int32, bytesRead *int32, callback ICefResourceReadCallback, outResult *bool)
type TOnResourceHandlerReadResponse func(dataOut uintptr, bytesToRead int32, bytesRead *int32, callback ICefCallback, outResult *bool) // deprecated
type TOnResourceHandlerCancel func()

// TRunFileDialogCallback

type TOnRunFileDialogDismissed func(filePaths IStrings)

// TSchemeHandlerFactory

type TOnSchemeHandlerFactoryNew func(browser ICefBrowser, frame ICefFrame, schemeName string, request ICefRequest) (result ICefResourceHandler)

// TCEFServerComponent

type TOnServer func(sender IObject, server ICEFServer)
type TOnServerCreated = TOnServer
type TOnServerDestroyed = TOnServer
type TOnServerInt32 = func(sender IObject, server ICEFServer, connectionId int32)
type TOnClientConnected = TOnServerInt32
type TOnClientDisconnected = TOnServerInt32
type TOnHttpRequest func(sender IObject, server ICEFServer, connectionId int32, clientAddress string, request ICefRequest)
type TOnWebSocketRequest func(sender IObject, server ICEFServer, connectionId int32, clientAddress string, request ICefRequest, callback ICefCallback)
type TOnWebSocketConnected = TOnServerInt32
type TOnWebSocketMessage func(sender IObject, server ICEFServer, connectionId int32, data Pointer, dataSize uint32)

// ICefTextfieldDelegate

type TOnTextfieldKeyEvent func(sender IObject, textField ICefTextfield, event TCefKeyEvent, result *bool)
type TOnAfterUserAction func(sender IObject, textField ICefTextfield)

// TCEFUrlRequestClientComponent

type TOnRequestCompleteRcc func(sender IObject, request ICefUrlRequest)
type TOnUploadProgressRcc func(sender IObject, request ICefUrlRequest, current, total int64)
type TOnDownloadProgressRcc func(sender IObject, request ICefUrlRequest, current, total int64)
type TOnDownloadDataRcc func(sender IObject, request ICefUrlRequest, data uintptr, dataLength uint32)
type TOnGetAuthCredentialsRcc func(sender IObject, isProxy bool, host string, port int32, realm, scheme string, callback ICefAuthCallback, result *bool) //重复 TOnGetAuthCredentials
type TNotifyRcc = TNotify

// TV8Accessor

type TOnV8AccessorGet func(name string, object ICefv8Value) (retVal ICefv8Value, exception string, result bool)
type TOnV8AccessorSet func(name string, object, value ICefv8Value) (exception string, result bool)

// TV8ArrayBufferReleaseCallback

type TOnV8ArrayBufferReleaseBuffer func(buffer uintptr)

// TV8Handler

type TOnV8HandlerExecute func(name string, object ICefv8Value, arguments ICefV8ValueArray) (retVal ICefv8Value, exception string, result bool)

// TV8Interceptor

type TOnV8InterceptorGetByName func(name string, object ICefv8Value) (retVal ICefv8Value, exception string, result bool)
type TOnV8InterceptorGetByIndex func(index int32, object ICefv8Value) (retVal ICefv8Value, exception string, result bool)
type TOnV8InterceptorSetByName func(name string, object, value ICefv8Value) (exception string, result bool)
type TOnV8InterceptorSetByIndex func(index int32, object, value ICefv8Value) (exception string, result bool)

// ICefViewDelegate

type TOnViewResultSize func(sender IObject, view ICefView, result *TCefSize)
type TOnGetPreferredSize = TOnViewResultSize
type TOnGetMinimumSize = TOnViewResultSize
type TOnGetMaximumSize = TOnViewResultSize
type TOnGetHeightForWidth func(sender IObject, view ICefView, width int32, result *int32)
type TOnParentViewChanged func(sender IObject, view ICefView, added bool, parent ICefView)
type TOnChildViewChanged func(sender IObject, view ICefView, added bool, child ICefView)
type TOnWindowChanged func(sender IObject, view ICefView, added bool)
type TOnLayoutChanged func(sender IObject, view ICefView, newBounds TCefRect)
type TOnView func(sender IObject, view ICefView)
type TOnFocus = TOnView
type TOnBlur = TOnView

// TTask

type TTaskExecute func()
