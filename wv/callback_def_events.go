//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
)

// Loader events

type TOnLoaderNotifyEvent func(sender IObject)
type TOnLoaderBrowserProcessExitedEvent func(sender IObject, environment ICoreWebView2Environment, args ICoreWebView2BrowserProcessExitedEventArgs)
type TOnLoaderNewBrowserVersionAvailableEvent func(sender IObject, environment ICoreWebView2Environment)
type TOnLoaderProcessInfosChangedEvent func(sender IObject, environment ICoreWebView2Environment)
type TOnLoaderGetCustomSchemesEvent func(sender IObject, customSchemes *TWVCustomSchemeInfoArray)

// Browser events

type TNotifyEvent = lcl.TNotifyEvent
type TOnExecuteScriptCompletedEvent func(sender IObject, errorCode int32, resulIObjectAsJson string, executionID int32)
type TOnCapturePreviewCompletedEvent func(sender IObject, errorCode int32)
type TOnWebResourceResponseViewGetContentCompletedEvent func(sender IObject, errorCode int32, contents IStream, resourceID int32)
type TOnGetCookiesCompletedEvent func(sender IObject, result int32, cookieList ICoreWebView2CookieList)
type TOnTrySuspendCompletedEvent func(sender IObject, errorCode int32, isSuccessful bool)
type TOnPrintToPdfCompletedEvent func(sender IObject, errorCode int32, isSuccessful bool)
type TOnCallDevToolsProtocolMethodCompletedEvent func(sender IObject, errorCode int32, returnObjectAsJson string, executionID int32)
type TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent func(sender IObject, errorCode int32, ID string)
type TOnMoveFocusRequestedEvent func(sender IObject, controller ICoreWebView2Controller, args ICoreWebView2MoveFocusRequestedEventArgs)
type TOnAcceleratorKeyPressedEvent func(sender IObject, controller ICoreWebView2Controller, args ICoreWebView2AcceleratorKeyPressedEventArgs)
type TOnBrowserProcessExitedEvent func(sender IObject, environment ICoreWebView2Environment, args ICoreWebView2BrowserProcessExitedEventArgs)
type TOnNavigationStartingEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2NavigationStartingEventArgs)
type TOnNavigationCompletedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2NavigationCompletedEventArgs)
type TOnSourceChangedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2SourceChangedEventArgs)
type TOnContentLoadingEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2ContentLoadingEventArgs)
type TOnNewWindowRequestedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2NewWindowRequestedEventArgs)
type TOnWebResourceRequestedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2WebResourceRequestedEventArgs)
type TOnScriptDialogOpeningEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2ScriptDialogOpeningEventArgs)
type TOnPermissionRequestedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2PermissionRequestedEventArgs)
type TOnProcessFailedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2ProcessFailedEventArgs)
type TOnWebMessageReceivedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2WebMessageReceivedEventArgs)
type TOnDevToolsProtocolEventReceivedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2DevToolsProtocolEventReceivedEventArgs, eventName string, eventID int32)
type TOnWebResourceResponseReceivedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2WebResourceResponseReceivedEventArgs)
type TOnDOMContentLoadedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2DOMContentLoadedEventArgs)
type TOnFrameCreatedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2FrameCreatedEventArgs)
type TOnDownloadStartingEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2DownloadStartingEventArgs)
type TOnClientCertificateRequestedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2ClientCertificateRequestedEventArgs)
type TOnBytesReceivedChangedEvent func(sender IObject, downloadOperation ICoreWebView2DownloadOperation, downloadID int32)
type TOnEstimatedEndTimeChangedEvent func(sender IObject, downloadOperation ICoreWebView2DownloadOperation, downloadID int32)
type TOnDownloadStateChangedEvent func(sender IObject, downloadOperation ICoreWebView2DownloadOperation, downloadID int32)
type TOnFrameNameChangedEvent func(sender IObject, frame ICoreWebView2Frame, frameID uint32)
type TOnFrameDestroyedEvent func(sender IObject, frame ICoreWebView2Frame, frameID uint32)
type TOnInitializationErrorEvent func(sender IObject, errorCode int32, errorMessage string)
type TOnPrintCompletedEvent func(sender IObject, errorCode int32, printStatus TWVPrintStatus)
type TOnRefreshIgnoreCacheCompletedEvent func(sender IObject, errorCode int32, resulIObjectAsJson string)
type TOnRetrieveHTMLCompletedEvent func(sender IObject, result bool, html string)
type TOnRetrieveTextCompletedEvent func(sender IObject, result bool, text string)
type TOnRetrieveMHTMLCompletedEvent func(sender IObject, result bool, mHtml string)
type TOnClearCacheCompletedEvent func(sender IObject, result bool)
type TOnClearDataForOriginCompletedEvent func(sender IObject, result bool)
type TOnOfflineCompletedEvent func(sender IObject, result bool)
type TOnIgnoreCertificateErrorsCompletedEvent func(sender IObject, result bool)
type TOnSimulateKeyEventCompletedEvent func(sender IObject, result bool)
type TOnIsMutedChangedEvent func(sender IObject, webView ICoreWebView2)
type TOnIsDocumentPlayingAudioChangedEvent func(sender IObject, webView ICoreWebView2)
type TOnIsDefaultDownloadDialogOpenChangedEvent func(sender IObject, webView ICoreWebView2)
type TOnProcessInfosChangedEvent func(sender IObject, environment ICoreWebView2Environment)
type TOnFrameNavigationStartingEvent func(sender IObject, frame ICoreWebView2Frame, args ICoreWebView2NavigationStartingEventArgs, frameID uint32)
type TOnFrameNavigationCompletedEvent func(sender IObject, frame ICoreWebView2Frame, args ICoreWebView2NavigationCompletedEventArgs, frameID uint32)
type TOnFrameContentLoadingEvent func(sender IObject, frame ICoreWebView2Frame, args ICoreWebView2ContentLoadingEventArgs, frameID uint32)
type TOnFrameDOMContentLoadedEvent func(sender IObject, frame ICoreWebView2Frame, args ICoreWebView2DOMContentLoadedEventArgs, frameID uint32)
type TOnFrameWebMessageReceivedEvent func(sender IObject, frame ICoreWebView2Frame, args ICoreWebView2WebMessageReceivedEventArgs, frameID uint32)
type TOnBasicAuthenticationRequestedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2BasicAuthenticationRequestedEventArgs)
type TOnContextMenuRequestedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2ContextMenuRequestedEventArgs)
type TOnCustomItemSelectedEvent func(sender IObject, menuItem ICoreWebView2ContextMenuItem)
type TOnStatusBarTextChangedEvent func(sender IObject, webView ICoreWebView2)
type TOnFramePermissionRequestedEvent func(sender IObject, frame ICoreWebView2Frame, args ICoreWebView2PermissionRequestedEventArgs, frameID uint32)
type TOnClearBrowsingDataCompletedEvent func(sender IObject, errorCode int32)
type TOnServerCertificateErrorActionsCompletedEvent func(sender IObject, errorCode int32)
type TOnServerCertificateErrorDetectedEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2ServerCertificateErrorDetectedEventArgs)
type TOnFaviconChangedEvent func(sender IObject, webView ICoreWebView2, args IUnknown)
type TOnGetFaviconCompletedEvent func(sender IObject, errorCode int32, faviconStream IStream)
type TOnPrintToPdfStreamCompletedEvent func(sender IObject, errorCode int32, pdfStream IStream)
type TOnGetCustomSchemesEvent func(sender IObject, customSchemes TWVCustomSchemeInfoArray)
type TOnGetNonDefaultPermissionSettingsCompletedEvent func(sender IObject, errorCode int32, collectionView ICoreWebView2PermissionSettingCollectionView)
type TOnSetPermissionStateCompletedEvent func(sender IObject, errorCode int32)
type TOnLaunchingExternalUriSchemeEvent func(sender IObject, webView ICoreWebView2, args ICoreWebView2LaunchingExternalUriSchemeEventArgs)
type TOnGetProcessExtendedInfosCompletedEvent func(sender IObject, errorCode int32, value ICoreWebView2ProcessExtendedInfoCollection)
type TOnBrowserExtensionRemoveCompletedEvent func(sender IObject, errorCode int32, extensionID string)
type TOnBrowserExtensionEnableCompletedEvent func(sender IObject, errorCode int32, extensionID string)
type TOnProfileAddBrowserExtensionCompletedEvent func(sender IObject, errorCode int32, extension ICoreWebView2BrowserExtension)
type TOnProfileGetBrowserExtensionsCompletedEvent func(sender IObject, errorCode int32, extensionList ICoreWebView2BrowserExtensionList)
type TOnProfileDeletedEvent func(sender IObject, profile ICoreWebView2Profile)
type TOnExecuteScriptWithResultCompletedEvent func(sender IObject, errorCode int32, result ICoreWebView2ExecuteScriptResult, executionID int32)

// Custom events

type TOnCompMsgEvent func(sender IObject, message *types.TMessage, handled bool)

type TDragDropEvent = lcl.TDragDropEvent
type TDragOverEvent = lcl.TDragOverEvent
type TStartDragEvent = lcl.TStartDragEvent
type TEndDragEvent = lcl.TEndDragEvent
