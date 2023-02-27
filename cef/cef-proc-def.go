//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// cef -> energy 导出函数定义
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api/dllimports"
)

func init() {
	var energyImportDefs = []*dllimports.ImportTable{
		//null nil
		dllimports.NewEnergyImport("", 0),
		//CEF
		dllimports.NewEnergyImport("CEFVersion", 0),
		dllimports.NewEnergyImport("LibBuildVersion", 0),
		dllimports.NewEnergyImport("CEF_Win_CreateRectRgn", 0),
		dllimports.NewEnergyImport("CEF_Win_SetRectRgn", 0),
		dllimports.NewEnergyImport("CEF_Win_DeleteObject", 0),
		dllimports.NewEnergyImport("CEF_Win_CombineRgn", 0),
		dllimports.NewEnergyImport("CEF_Win_SetDraggableRegions", 0),
		dllimports.NewEnergyImport("CEF_Win_PtInRegion", 0),
		dllimports.NewEnergyImport("CEF_Win_ScreenToClient", 0),
		dllimports.NewEnergyImport("CEF_Win_ClientToScreen", 0),
		dllimports.NewEnergyImport("CEF_Win_DefWindowProc", 0),
		dllimports.NewEnergyImport("CEF_Win_DefSubclassProc", 0),
		dllimports.NewEnergyImport("CEF_Win_CreateRoundRectRgn", 0),
		dllimports.NewEnergyImport("CEF_Win_SetWindowRgn", 0),
		dllimports.NewEnergyImport("CEF_Win_SetCursor", 0),
		dllimports.NewEnergyImport("CEF_Win_LoadCursor", 0),
		dllimports.NewEnergyImport("CEF_Win_OnPaint", 0),
		//ApplicationQueueAsyncCallFunc
		dllimports.NewEnergyImport("SetApplicationQueueAsyncCallFunc", 0),
		dllimports.NewEnergyImport("CEFApplication_QueueAsyncCall", 0),
		dllimports.NewEnergyImport("SetCEFWindowBindCallbackFunc", 0),
		dllimports.NewEnergyImport("SetCEFIPCCallbackFunc", 0),
		//GoForm
		dllimports.NewEnergyImport("CEF_AddGoForm", 0),
		dllimports.NewEnergyImport("CEF_RemoveGoForm", 0),
		//ICefCallback
		dllimports.NewEnergyImport("cefCallback_Cont", 0),
		dllimports.NewEnergyImport("cefCallback_Cancel", 0),
		//process
		dllimports.NewEnergyImport("SetMacOSXCommandLine", 0),
		dllimports.NewEnergyImport("CEFStartMainProcess", 0),
		dllimports.NewEnergyImport("CEFStartSubProcess", 0),
		dllimports.NewEnergyImport("AddCustomCommandLine", 0),
		//application
		dllimports.NewEnergyImport("CEFApplication_RunMessageLoop", 0),
		dllimports.NewEnergyImport("CEFApplication_QuitMessageLoop", 0),
		dllimports.NewEnergyImport("CEFApplication_Create", 0),
		dllimports.NewEnergyImport("CEFApplication_Destroy", 0),
		dllimports.NewEnergyImport("CEFApplication_Free", 0),
		dllimports.NewEnergyImport("CEFApplication_StopScheduler", 0),
		dllimports.NewEnergyImport("CEFApplication_ExecuteJS", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_SetCommonRootName", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_SetObjectRootName", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_CommonValueBindInfo", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_ObjectValueBindInfo", 0),
		//application - event
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnContextCreated", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnRegCustomSchemes", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnContextInitialized", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnWebKitInitialized", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnBeforeChildProcessLaunch", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnScheduleMessagePumpWork", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnGetDefaultClient", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnGetLocalizedString", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnGetDataResource", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnGetDataResourceForScale", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnProcessMessageReceived", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnBrowserDestroyed", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnRenderLoadStart", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnRenderLoadEnd", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnRenderLoadError", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnRenderLoadingStateChange", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnBrowserCreated", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnContextReleased", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnUncaughtException", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnFocusedNodeChanged", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnLoadingStateChange", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnLoadStart", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnLoadEnd", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnLoadError", 0),
		//CEFParentWindow
		dllimports.NewEnergyImport("CEFWindow_Create", 0),
		dllimports.NewEnergyImport("CEFWindow_GetHandle", 0),
		dllimports.NewEnergyImport("CEFWindow_DestroyChildWindow", 0),
		dllimports.NewEnergyImport("CEFWindow_HandleAllocated", 0),
		dllimports.NewEnergyImport("CEFWindow_CreateHandle", 0),
		dllimports.NewEnergyImport("CEFWindow_Free", 0),
		dllimports.NewEnergyImport("CEFWindow_SetParent", 0),
		dllimports.NewEnergyImport("CEFWindow_GetAlign", 0),
		dllimports.NewEnergyImport("CEFWindow_SetAlign", 0),
		dllimports.NewEnergyImport("CEFWindow_GetAnchors", 0),
		dllimports.NewEnergyImport("CEFWindow_SetAnchors", 0),
		dllimports.NewEnergyImport("CEFWindow_GetVisible", 0),
		dllimports.NewEnergyImport("CEFWindow_SetVisible", 0),
		dllimports.NewEnergyImport("CEFWindow_GetEnabled", 0),
		dllimports.NewEnergyImport("CEFWindow_SetEnabled", 0),
		dllimports.NewEnergyImport("CEFWindow_GetLeft", 0),
		dllimports.NewEnergyImport("CEFWindow_SetLeft", 0),
		dllimports.NewEnergyImport("CEFWindow_GetTop", 0),
		dllimports.NewEnergyImport("CEFWindow_SetTop", 0),
		dllimports.NewEnergyImport("CEFWindow_GetWidth", 0),
		dllimports.NewEnergyImport("CEFWindow_SetWidth", 0),
		dllimports.NewEnergyImport("CEFWindow_GetHeight", 0),
		dllimports.NewEnergyImport("CEFWindow_SetHeight", 0),
		dllimports.NewEnergyImport("CEFWindow_GetBoundsRect", 0),
		dllimports.NewEnergyImport("CEFWindow_SetBoundsRect", 0),
		dllimports.NewEnergyImport("CEFWindow_GetName", 0),
		dllimports.NewEnergyImport("CEFWindow_SetName", 0),
		dllimports.NewEnergyImport("CEFWindow_UpdateSize", 0),
		dllimports.NewEnergyImport("CEFWindow_OnEnter", 0),
		dllimports.NewEnergyImport("CEFWindow_OnExit", 0),
		//CEFLinkedParentWindow
		dllimports.NewEnergyImport("CEFLinkedWindow_Create", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetHandle", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_DestroyChildWindow", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_HandleAllocated", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_CreateHandle", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_Free", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetParent", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetAlign", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetAlign", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetAnchors", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetAnchors", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetVisible", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetVisible", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetEnabled", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetEnabled", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetLeft", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetLeft", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetTop", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetTop", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetWidth", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetWidth", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetHeight", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetHeight", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetBoundsRect", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetBoundsRect", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_GetName", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetName", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_UpdateSize", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_OnEnter", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_OnExit", 0),
		dllimports.NewEnergyImport("CEFLinkedWindow_SetChromium", 0),
		//ICefBrowser
		dllimports.NewEnergyImport("CEFBrowser_ShowDevTools", 0),
		dllimports.NewEnergyImport("CEFBrowser_CloseDevTools", 0),
		dllimports.NewEnergyImport("CEFBrowser_HasDevTools", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetIdentifier", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetFocusedFrame", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetMainFrame", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetFrameById", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetFrameByName", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetHostWindowHandle", 0),
		dllimports.NewEnergyImport("CEFBrowser_CloseBrowser", 0),
		dllimports.NewEnergyImport("CEFBrowser_TryCloseBrowser", 0),
		dllimports.NewEnergyImport("CEFBrowser_SetFocus", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetZoomLevel", 0),
		dllimports.NewEnergyImport("CEFBrowser_SetZoomLevel", 0),
		dllimports.NewEnergyImport("CEFBrowser_RunFileDialog", 0),
		dllimports.NewEnergyImport("CEFBrowser_StartDownload", 0),
		dllimports.NewEnergyImport("CEFBrowser_DownloadImage", 0),
		dllimports.NewEnergyImport("CEFBrowser_Print", 0),
		dllimports.NewEnergyImport("CEFBrowser_ExecuteDevToolsMethod", 0),
		dllimports.NewEnergyImport("CEFBrowser_SendKeyEvent", 0),
		dllimports.NewEnergyImport("CEFBrowser_SetAudioMuted", 0),
		dllimports.NewEnergyImport("CEFBrowser_IsAudioMuted", 0),
		dllimports.NewEnergyImport("CEFBrowser_SetAutoResizeEnabled", 0),
		dllimports.NewEnergyImport("CEFBrowser_SetAccessibilityState", 0),
		dllimports.NewEnergyImport("CEFBrowser_NotifyMoveOrResizeStarted", 0),
		dllimports.NewEnergyImport("CEFBrowser_NotifyScreenInfoChanged", 0),
		dllimports.NewEnergyImport("CEFBrowser_SendCaptureLostEvent", 0),
		dllimports.NewEnergyImport("CEFBrowser_SendTouchEvent", 0),
		dllimports.NewEnergyImport("CEFBrowser_SendMouseWheelEvent", 0),
		dllimports.NewEnergyImport("CEFBrowser_SendMouseMoveEvent", 0),
		dllimports.NewEnergyImport("CEFBrowser_SendMouseClickEvent", 0),
		dllimports.NewEnergyImport("CEFBrowser_CanGoBack", 0),
		dllimports.NewEnergyImport("CEFBrowser_GoBack", 0),
		dllimports.NewEnergyImport("CEFBrowser_CanGoForward", 0),
		dllimports.NewEnergyImport("CEFBrowser_GoForward", 0),
		dllimports.NewEnergyImport("CEFBrowser_IsLoading", 0),
		dllimports.NewEnergyImport("CEFBrowser_Reload", 0),
		dllimports.NewEnergyImport("CEFBrowser_ReloadIgnoreCache", 0),
		dllimports.NewEnergyImport("CEFBrowser_StopLoad", 0),
		dllimports.NewEnergyImport("CEFBrowser_FrameCount", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetFrameNames", 0),
		dllimports.NewEnergyImport("CEFBrowser_Find", 0),
		dllimports.NewEnergyImport("CEFBrowser_StopFinding", 0),
		//TCEFChromium - event
		dllimports.NewEnergyImport("CEFChromium_SetOnAfterCreated", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnBeforeClose", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnClose", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnPdfPrintFinished", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnZoomPctAvailable", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnLoadStart", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnLoadingStateChange", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnLoadingProgressChange", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnLoadError", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnLoadEnd", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnBeforeDownload", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnDownloadUpdated", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnFullScreenModeChange", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnBeforeBrowse", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnAddressChange", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnKeyEvent", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnTitleChange", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnRenderCompMsg", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnWidgetCompMsg", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnBrowserCompMsg", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnRenderProcessTerminated", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnRenderViewReady", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnScrollOffsetChanged", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnProcessMessageReceived", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnFindResult", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnCookieSet", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnCookiesDeleted", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnCookiesFlushed", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnCookiesVisited", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnCookieVisitorDestroyed", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnBeforeContextMenu", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnContextMenuCommand", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnContextMenuDismissed", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnBeforeResourceLoad", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnResourceResponse", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnResourceRedirect", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnResourceLoadComplete", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnFrameAttached", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnFrameCreated", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnFrameDetached", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnMainFrameChanged", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnBeforePopup", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnOpenUrlFromTab", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnDragEnter", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOnDraggableRegionsChanged", 0),
		//TCEFChromium - proc
		dllimports.NewEnergyImport("CEFChromium_Free", 0),
		dllimports.NewEnergyImport("CEFChromium_GetHashCode", 0),
		dllimports.NewEnergyImport("CEFChromium_ClassName", 0),
		dllimports.NewEnergyImport("CEFChromium_Equals", 0),
		dllimports.NewEnergyImport("CEFChromium_ClassType", 0),
		dllimports.NewEnergyImport("CEFChromium_InstanceSize", 0),
		dllimports.NewEnergyImport("CEFChromium_InheritsFrom", 0),
		dllimports.NewEnergyImport("CEFChromium_ToString", 0),
		dllimports.NewEnergyImport("CEFChromium_Create", 0),
		dllimports.NewEnergyImport("CEFChromium_Browser", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserById", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserIdByIndex", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserCount", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserId", 0),
		dllimports.NewEnergyImport("CEFChromium_SetDefaultURL", 0),
		dllimports.NewEnergyImport("CEFChromium_SetMultiBrowserMode", 0),
		dllimports.NewEnergyImport("CEFChromium_LoadURL", 0),
		dllimports.NewEnergyImport("CEFChromium_LoadString", 0),
		dllimports.NewEnergyImport("CEFChromium_StartDownload", 0),
		dllimports.NewEnergyImport("CEFChromium_DownloadImage", 0),
		dllimports.NewEnergyImport("CEFChromium_Reload", 0),
		dllimports.NewEnergyImport("CEFChromium_StopLoad", 0),
		dllimports.NewEnergyImport("CEFChromium_ResetZoomLevel", 0),
		dllimports.NewEnergyImport("CEFChromium_CloseAllBrowsers", 0),
		dllimports.NewEnergyImport("CEFChromium_CreateBrowserByWindow", 0),
		dllimports.NewEnergyImport("CEFChromium_CreateBrowserByLinkedWindow", 0),
		dllimports.NewEnergyImport("CEFChromium_CreateBrowserByBrowserViewComponent", 0),
		dllimports.NewEnergyImport("CEFChromium_Initialized", 0),
		dllimports.NewEnergyImport("CEFChromium_IsSameBrowser", 0),
		dllimports.NewEnergyImport("CEFChromium_PrintToPDF", 0),
		dllimports.NewEnergyImport("CEFChromium_Print", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserDownloadCancel", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserDownloadPause", 0),
		dllimports.NewEnergyImport("CEFChromium_DownloadResume", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserZoom", 0),
		dllimports.NewEnergyImport("CEFChromium_GoBackForward", 0),
		dllimports.NewEnergyImport("CEFChromium_NotifyMoveOrResizeStarted", 0),
		dllimports.NewEnergyImport("CEFChromium_CloseBrowser", 0),
		dllimports.NewEnergyImport("CEFChromium_ExecuteJavaScript", 0),
		dllimports.NewEnergyImport("CEFChromium_ShowDevTools", 0),
		dllimports.NewEnergyImport("CEFChromium_ShowDevToolsByWindowParent", 0),
		dllimports.NewEnergyImport("CEFChromium_CloseDevTools", 0),
		dllimports.NewEnergyImport("CEFChromium_CloseDevToolsByWindowParent", 0),
		dllimports.NewEnergyImport("CEFChromium_VisitAllCookies", 0),
		dllimports.NewEnergyImport("CEFChromium_VisitURLCookies", 0),
		dllimports.NewEnergyImport("CEFChromium_DeleteCookies", 0),
		dllimports.NewEnergyImport("CEFChromium_SetCookie", 0),
		dllimports.NewEnergyImport("CEFChromium_SetProxy", 0),
		dllimports.NewEnergyImport("CEFChromium_UpdatePreferences", 0),
		dllimports.NewEnergyImport("CEFChromium_ExecuteDevToolsMethod", 0),
		dllimports.NewEnergyImport("CEFChromium_CreateClientHandler", 0),
		dllimports.NewEnergyImport("CEFChromium_SetFocus", 0),
		dllimports.NewEnergyImport("CEFChromium_SendCaptureLostEvent", 0),
		dllimports.NewEnergyImport("CEFChromium_FrameIsFocused", 0),
		dllimports.NewEnergyImport("CEFChromium_TryCloseBrowser", 0),
		dllimports.NewEnergyImport("CEFChromium_BrowserHandle", 0),
		dllimports.NewEnergyImport("CEFChromium_WidgetHandle", 0),
		dllimports.NewEnergyImport("CEFChromium_RenderHandle", 0),
		dllimports.NewEnergyImport("CEFChromium_SetCustomHeader", 0), //
		dllimports.NewEnergyImport("CEFChromium_GetCustomHeader", 0),
		dllimports.NewEnergyImport("CEFChromium_SetJavascriptEnabled", 0),
		dllimports.NewEnergyImport("CEFChromium_GetJavascriptEnabled", 0),
		dllimports.NewEnergyImport("CEFChromium_SetWebRTCIPHandlingPolicy", 0),
		dllimports.NewEnergyImport("CEFChromium_GetWebRTCIPHandlingPolicy", 0),
		dllimports.NewEnergyImport("CEFChromium_SetWebRTCMultipleRoutes", 0),
		dllimports.NewEnergyImport("CEFChromium_GetWebRTCMultipleRoutes", 0),
		dllimports.NewEnergyImport("CEFChromium_SetWebRTCNonproxiedUDP", 0),
		dllimports.NewEnergyImport("CEFChromium_GetWebRTCNonproxiedUDP", 0),
		dllimports.NewEnergyImport("CEFChromium_SetBatterySaverModeState", 0),
		dllimports.NewEnergyImport("CEFChromium_GetBatterySaverModeState", 0),
		dllimports.NewEnergyImport("CEFChromium_SetHighEfficiencyMode", 0),
		dllimports.NewEnergyImport("CEFChromium_GetHighEfficiencyMode", 0),
		dllimports.NewEnergyImport("CEFChromium_SetLoadImagesAutomatically", 0),
		dllimports.NewEnergyImport("CEFChromium_GetLoadImagesAutomatically", 0),
		dllimports.NewEnergyImport("CEFChromium_SetQuicAllowed", 0),
		dllimports.NewEnergyImport("CEFChromium_GetQuicAllowed", 0),
		dllimports.NewEnergyImport("CEFChromium_SetOffline", 0),
		dllimports.NewEnergyImport("CEFChromium_GetOffline", 0),
		dllimports.NewEnergyImport("CEFChromium_SetDefaultWindowInfoExStyle", 0),
		dllimports.NewEnergyImport("CEFChromium_GetDefaultWindowInfoExStyle", 0),
		dllimports.NewEnergyImport("CEFChromium_SetBlock3rdPartyCookies", 0),
		dllimports.NewEnergyImport("CEFChromium_GetBlock3rdPartyCookies", 0),
		dllimports.NewEnergyImport("CEFChromium_SetAcceptCookies", 0),
		dllimports.NewEnergyImport("CEFChromium_GetAcceptCookies", 0),
		dllimports.NewEnergyImport("CEFChromium_SetAcceptLanguageList", 0),
		dllimports.NewEnergyImport("CEFChromium_GetAcceptLanguageList", 0),
		dllimports.NewEnergyImport("CEFChromium_SetPrintingEnabled", 0),
		dllimports.NewEnergyImport("CEFChromium_GetPrintingEnabled", 0),
		dllimports.NewEnergyImport("CEFChromium_SetYouTubeRestrict", 0),
		dllimports.NewEnergyImport("CEFChromium_GetYouTubeRestrict", 0),
		dllimports.NewEnergyImport("CEFChromium_SetSafeSearch", 0),
		dllimports.NewEnergyImport("CEFChromium_GetSafeSearch", 0),
		dllimports.NewEnergyImport("CEFChromium_SetAudioMuted", 0),
		dllimports.NewEnergyImport("CEFChromium_GetAudioMuted", 0),
		dllimports.NewEnergyImport("CEFChromium_SetDragOperations", 0),
		dllimports.NewEnergyImport("CEFChromium_GetDragOperations", 0),
		dllimports.NewEnergyImport("CEFChromium_GetFrameCount", 0),
		dllimports.NewEnergyImport("CEFChromium_SetSpellCheckerDicts", 0),
		dllimports.NewEnergyImport("CEFChromium_GetSpellCheckerDicts", 0),
		dllimports.NewEnergyImport("CEFChromium_SetSpellChecking", 0),
		dllimports.NewEnergyImport("CEFChromium_GetSpellChecking", 0),
		dllimports.NewEnergyImport("CEFChromium_SetAlwaysOpenPDFExternally", 0),
		dllimports.NewEnergyImport("CEFChromium_GetAlwaysOpenPDFExternally", 0),
		dllimports.NewEnergyImport("CEFChromium_SetAlwaysAuthorizePlugins", 0),
		dllimports.NewEnergyImport("CEFChromium_GetAlwaysAuthorizePlugins", 0),
		dllimports.NewEnergyImport("CEFChromium_SetAllowOutdatedPlugins", 0),
		dllimports.NewEnergyImport("CEFChromium_GetAllowOutdatedPlugins", 0),
		dllimports.NewEnergyImport("CEFChromium_SetSendReferrer", 0),
		dllimports.NewEnergyImport("CEFChromium_GetSendReferrer", 0),
		dllimports.NewEnergyImport("CEFChromium_SetDoNotTrack", 0),
		dllimports.NewEnergyImport("CEFChromium_GetDoNotTrack", 0),
		dllimports.NewEnergyImport("CEFChromium_SetZoomStep", 0),
		dllimports.NewEnergyImport("CEFChromium_GetZoomStep", 0),
		dllimports.NewEnergyImport("CEFChromium_SetZoomPct", 0),
		dllimports.NewEnergyImport("CEFChromium_GetZoomPct", 0),
		dllimports.NewEnergyImport("CEFChromium_SetZoomLevel", 0),
		dllimports.NewEnergyImport("CEFChromium_GetZoomLevel", 0),
		dllimports.NewEnergyImport("CEFChromium_SetDefaultEncoding", 0),
		dllimports.NewEnergyImport("CEFChromium_GetDefaultEncoding", 0),
		dllimports.NewEnergyImport("CEFChromium_SendProcessMessage", 0),
		//ICefBeforeDownloadCallback
		dllimports.NewEnergyImport("CEFChromium_SetDownloadPath", 0),
		//ICefFrame
		dllimports.NewEnergyImport("CEFFrame_Undo", 0),
		dllimports.NewEnergyImport("CEFFrame_Redo", 0),
		dllimports.NewEnergyImport("CEFFrame_Cut", 0),
		dllimports.NewEnergyImport("CEFFrame_Copy", 0),
		dllimports.NewEnergyImport("CEFFrame_Paste", 0),
		dllimports.NewEnergyImport("CEFFrame_Del", 0),
		dllimports.NewEnergyImport("CEFFrame_SelectAll", 0),
		dllimports.NewEnergyImport("CEFFrame_ViewSource", 0),
		dllimports.NewEnergyImport("CEFFrame_LoadUrl", 0),
		dllimports.NewEnergyImport("CEFFrame_ExecuteJavaScript", 0),
		dllimports.NewEnergyImport("CEFFrame_IsValid", 0),
		dllimports.NewEnergyImport("CEFFrame_IsMain", 0),
		dllimports.NewEnergyImport("CEFFrame_IsFocused", 0),
		dllimports.NewEnergyImport("CEFFrame_SendProcessMessageByIPC", 0),
		dllimports.NewEnergyImport("CEFFrame_SendProcessMessage", 0),
		dllimports.NewEnergyImport("CEFFrame_LoadRequest", 0),
		dllimports.NewEnergyImport("CEFFrame_Browser", 0),
		dllimports.NewEnergyImport("CEFFrame_GetV8Context", 0),
		dllimports.NewEnergyImport("CEFFrame_Identifier", 0),
		dllimports.NewEnergyImport("CEFFrame_Name", 0),
		dllimports.NewEnergyImport("CEFFrame_Url", 0),
		dllimports.NewEnergyImport("CEFFrame_Parent", 0),
		//ICefMenuModel
		dllimports.NewEnergyImport("cefMenuModel_AddSeparator", 0),
		dllimports.NewEnergyImport("cefMenuModel_Clear", 0),
		dllimports.NewEnergyImport("cefMenuModel_IsSubMenu", 0),
		dllimports.NewEnergyImport("cefMenuModel_GetCount", 0),
		dllimports.NewEnergyImport("cefMenuModel_AddItem", 0),
		dllimports.NewEnergyImport("cefMenuModel_AddCheckItem", 0),
		dllimports.NewEnergyImport("cefMenuModel_AddRadioItem", 0),
		dllimports.NewEnergyImport("cefMenuModel_AddSubMenu", 0),
		dllimports.NewEnergyImport("cefMenuModel_Remove", 0),
		dllimports.NewEnergyImport("cefMenuModel_RemoveAt", 0),
		dllimports.NewEnergyImport("cefMenuModel_SetChecked", 0),
		dllimports.NewEnergyImport("cefMenuModel_IsChecked", 0),
		dllimports.NewEnergyImport("cefMenuModel_SetColor", 0),
		dllimports.NewEnergyImport("cefMenuModel_SetFontList", 0),
		dllimports.NewEnergyImport("cefMenuModel_HasAccelerator", 0),
		dllimports.NewEnergyImport("cefMenuModel_SetAccelerator", 0),
		dllimports.NewEnergyImport("cefMenuModel_RemoveAccelerator", 0),
		dllimports.NewEnergyImport("cefMenuModel_IsVisible", 0),
		dllimports.NewEnergyImport("cefMenuModel_SetVisible", 0),
		dllimports.NewEnergyImport("cefMenuModel_IsEnabled", 0),
		dllimports.NewEnergyImport("cefMenuModel_SetEnabled", 0),
		dllimports.NewEnergyImport("cefMenuModel_SetLabel", 0),
		dllimports.NewEnergyImport("cefMenuModel_GetIndexOf", 0),
		//CEFWindowInfo
		dllimports.NewEnergyImport("CEFWindowInfoAsChild", 0),
		dllimports.NewEnergyImport("CEFWindowInfoAsPopUp", 0),
		dllimports.NewEnergyImport("CEFWindowInfoAsWindowless", 0),
		dllimports.NewEnergyImport("CEFJSRegisterExtension", 0),
		//ICefRequest
		dllimports.NewEnergyImport("CefRequest_IsReadOnly", 0),
		dllimports.NewEnergyImport("CefRequest_SetUrl", 0),
		dllimports.NewEnergyImport("CefRequest_SetMethod", 0),
		dllimports.NewEnergyImport("CefRequest_SetReferrer", 0),
		dllimports.NewEnergyImport("CefRequest_SetFlags", 0),
		dllimports.NewEnergyImport("CefRequest_SetFirstPartyForCookies", 0),
		dllimports.NewEnergyImport("CefRequest_GetHeaderByName", 0),
		dllimports.NewEnergyImport("CefRequest_SetHeaderByName", 0),
		dllimports.NewEnergyImport("CefRequest_GetHeaderMap", 0),
		dllimports.NewEnergyImport("CefRequest_SetHeaderMap", 0),
		dllimports.NewEnergyImport("CefRequest_GetPostData", 0),
		dllimports.NewEnergyImport("CefRequest_SetPostData", 0),
		//ICefResponse
		dllimports.NewEnergyImport("CefResponse_IsReadOnly", 0),
		dllimports.NewEnergyImport("CefResponse_SetError", 0),
		dllimports.NewEnergyImport("CefResponse_SetStatus", 0),
		dllimports.NewEnergyImport("CefResponse_SetStatusText", 0),
		dllimports.NewEnergyImport("CefResponse_SetMimeType", 0),
		dllimports.NewEnergyImport("CefResponse_SetCharset", 0),
		dllimports.NewEnergyImport("CefResponse_GetHeaderByName", 0),
		dllimports.NewEnergyImport("CefResponse_SetHeaderByName", 0),
		dllimports.NewEnergyImport("CefResponse_SetURL", 0),
		dllimports.NewEnergyImport("CefResponse_GetHeaderMap", 0),
		dllimports.NewEnergyImport("CefResponse_SetHeaderMap", 0),
		//ICefStringMultiMap
		dllimports.NewEnergyImport("StringMultimapRef_Create", 0),
		dllimports.NewEnergyImport("StringMultimap_GetSize", 0),
		dllimports.NewEnergyImport("StringMultimap_FindCount", 0),
		dllimports.NewEnergyImport("StringMultimap_GetEnumerate", 0),
		dllimports.NewEnergyImport("StringMultimap_GetKey", 0),
		dllimports.NewEnergyImport("StringMultimap_GetValue", 0),
		dllimports.NewEnergyImport("StringMultimap_Append", 0),
		dllimports.NewEnergyImport("StringMultimap_Clear", 0),
		//ICefImage
		dllimports.NewEnergyImport("CEFImage_New", 0),
		dllimports.NewEnergyImport("CEFImage_AddPng", 0),
		dllimports.NewEnergyImport("CEFImage_AddJpeg", 0),
		dllimports.NewEnergyImport("CEFImage_GetWidth", 0),
		dllimports.NewEnergyImport("CEFImage_GetHeight", 0),
		//TCEFWindowComponent
		dllimports.NewEnergyImport("CEFWindowComponent_Create", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_CreateTopLevelWindow", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Show", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Hide", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_CenterWindow", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Close", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Activate", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Deactivate", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_BringToTop", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Maximize", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Minimize", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Restore", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_AddOverlayView", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_ShowMenu", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_CancelMenu", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetDraggableRegions", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SendKeyPress", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SendMouseMove", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SendMouseEvents", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetAccelerator", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_RemoveAccelerator", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_RemoveAllAccelerators", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetAlwaysOnTop", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetFullscreen", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetBackgroundColor", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Bounds", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Size", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Position", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetBounds", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetSize", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetPosition", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetTitle", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Title", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_WindowIcon", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetWindowIcon", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_WindowAppIcon", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetWindowAppIcon", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_Display", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_ClientAreaBoundsInScreen", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_WindowHandle", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_IsClosed", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_IsActive", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_IsAlwaysOnTop", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_IsFullscreen", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_IsMaximized", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_IsMinimized", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_AddChildView", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnWindowCreated", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnWindowDestroyed", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnWindowActivationChanged", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnGetParentWindow", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnGetInitialBounds", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnGetInitialShowState", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnIsFrameless", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnCanResize", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnCanMaximize", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnCanMinimize", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnCanClose", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnAccelerator", 0),
		dllimports.NewEnergyImport("CEFWindowComponent_SetOnKeyEvent", 0),
		//TCEFBrowserViewComponent
		dllimports.NewEnergyImport("CEFBrowserViewComponent_Create", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_CreateBrowserView", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_GetForBrowser", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_SetPreferAccelerators", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_RequestFocus", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_Browser", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_BrowserView", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_SetOnBrowserCreated", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_SetOnBrowserDestroyed", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_SetOnGetDelegateForPopupBrowserView", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_SetOnPopupBrowserViewCreated", 0),
		dllimports.NewEnergyImport("CEFBrowserViewComponent_SetOnGetChromeToolbarType", 0),
		//ICefDisplay
		dllimports.NewEnergyImport("CEFDisplay_ID", 0),
		dllimports.NewEnergyImport("CEFDisplay_DeviceScaleFactor", 0),
		dllimports.NewEnergyImport("CEFDisplay_Rotation", 0),
		dllimports.NewEnergyImport("CEFDisplay_Bounds", 0),
		dllimports.NewEnergyImport("CEFDisplay_WorkArea", 0),
		//ICefWindow
		dllimports.NewEnergyImport("ICEFWindow_Show", 0),
		dllimports.NewEnergyImport("ICEFWindow_Hide", 0),
		dllimports.NewEnergyImport("ICEFWindow_CenterWindow", 0),
		dllimports.NewEnergyImport("ICEFWindow_Close", 0),
		dllimports.NewEnergyImport("ICEFWindow_IsClosed", 0),
		dllimports.NewEnergyImport("ICEFWindow_Activate", 0),
		dllimports.NewEnergyImport("ICEFWindow_Deactivate", 0),
		dllimports.NewEnergyImport("ICEFWindow_IsActive", 0),
		dllimports.NewEnergyImport("ICEFWindow_BringToTop", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetAlwaysOnTop", 0),
		dllimports.NewEnergyImport("ICEFWindow_IsAlwaysOnTop", 0),
		dllimports.NewEnergyImport("ICEFWindow_Maximize", 0),
		dllimports.NewEnergyImport("ICEFWindow_Minimize", 0),
		dllimports.NewEnergyImport("ICEFWindow_Restore", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetFullscreen", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetBackgroundColor", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetBounds", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetSize", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetPosition", 0),
		dllimports.NewEnergyImport("ICEFWindow_IsMaximized", 0),
		dllimports.NewEnergyImport("ICEFWindow_IsMinimized", 0),
		dllimports.NewEnergyImport("ICEFWindow_IsFullscreen", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetTitle", 0),
		dllimports.NewEnergyImport("ICEFWindow_GetTitle", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetWindowIcon", 0),
		dllimports.NewEnergyImport("ICEFWindow_GetWindowIcon", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetWindowAppIcon", 0),
		dllimports.NewEnergyImport("ICEFWindow_GetWindowAppIcon", 0),
		dllimports.NewEnergyImport("ICEFWindow_AddOverlayView", 0),
		dllimports.NewEnergyImport("ICEFWindow_ShowMenu", 0),
		dllimports.NewEnergyImport("ICEFWindow_CancelMenu", 0),
		dllimports.NewEnergyImport("ICEFWindow_GetDisplay", 0),
		dllimports.NewEnergyImport("ICEFWindow_GetClientAreaBoundsInScreen", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetDraggableRegions", 0),
		dllimports.NewEnergyImport("ICEFWindow_GetWindowHandle", 0),
		dllimports.NewEnergyImport("ICEFWindow_SendKeyPress", 0),
		dllimports.NewEnergyImport("ICEFWindow_SendMouseMove", 0),
		dllimports.NewEnergyImport("ICEFWindow_SendMouseEvents", 0),
		dllimports.NewEnergyImport("ICEFWindow_SetAccelerator", 0),
		dllimports.NewEnergyImport("ICEFWindow_RemoveAccelerator", 0),
		dllimports.NewEnergyImport("ICEFWindow_RemoveAllAccelerators", 0),
		//ICefV8Value
		dllimports.NewEnergyImport("CefV8Value_IsValid", 0),
		dllimports.NewEnergyImport("CefV8Value_IsUndefined", 0),
		dllimports.NewEnergyImport("CefV8Value_IsNull", 0),
		dllimports.NewEnergyImport("CefV8Value_IsBool", 0),
		dllimports.NewEnergyImport("CefV8Value_IsInt", 0),
		dllimports.NewEnergyImport("CefV8Value_IsUInt", 0),
		dllimports.NewEnergyImport("CefV8Value_IsDouble", 0),
		dllimports.NewEnergyImport("CefV8Value_IsDate", 0),
		dllimports.NewEnergyImport("CefV8Value_IsString", 0),
		dllimports.NewEnergyImport("CefV8Value_IsObject", 0),
		dllimports.NewEnergyImport("CefV8Value_IsArray", 0),
		dllimports.NewEnergyImport("CefV8Value_IsArrayBuffer", 0),
		dllimports.NewEnergyImport("CefV8Value_IsFunction", 0),
		dllimports.NewEnergyImport("CefV8Value_IsPromise", 0),
		dllimports.NewEnergyImport("CefV8Value_IsSame", 0),
		dllimports.NewEnergyImport("CefV8Value_GetBoolValue", 0),
		dllimports.NewEnergyImport("CefV8Value_GetIntValue", 0),
		dllimports.NewEnergyImport("CefV8Value_GetUIntValue", 0),
		dllimports.NewEnergyImport("CefV8Value_GetDoubleValue", 0),
		dllimports.NewEnergyImport("CefV8Value_GetDateValue", 0),
		dllimports.NewEnergyImport("CefV8Value_GetStringValue", 0),
		dllimports.NewEnergyImport("CefV8Value_IsUserCreated", 0),
		dllimports.NewEnergyImport("CefV8Value_HasException", 0),
		dllimports.NewEnergyImport("CefV8Value_GetException", 0),
		dllimports.NewEnergyImport("CefV8Value_ClearException", 0),
		dllimports.NewEnergyImport("CefV8Value_WillRethrowExceptions", 0),
		dllimports.NewEnergyImport("CefV8Value_SetRethrowExceptions", 0),
		dllimports.NewEnergyImport("CefV8Value_HasValueByKey", 0),
		dllimports.NewEnergyImport("CefV8Value_HasValueByIndex", 0),
		dllimports.NewEnergyImport("CefV8Value_DeleteValueByKey", 0),
		dllimports.NewEnergyImport("CefV8Value_DeleteValueByIndex", 0),
		dllimports.NewEnergyImport("CefV8Value_GetValueByKey", 0),
		dllimports.NewEnergyImport("CefV8Value_GetValueByIndex", 0),
		dllimports.NewEnergyImport("CefV8Value_SetValueByKey", 0),
		dllimports.NewEnergyImport("CefV8Value_SetValueByIndex", 0),
		dllimports.NewEnergyImport("CefV8Value_SetValueByAccessor", 0),
		dllimports.NewEnergyImport("CefV8Value_GetKeys", 0),
		dllimports.NewEnergyImport("CefV8Value_SetUserData", 0),
		dllimports.NewEnergyImport("CefV8Value_GetUserData", 0),
		dllimports.NewEnergyImport("CefV8Value_GetExternallyAllocatedMemory", 0),
		dllimports.NewEnergyImport("CefV8Value_AdjustExternallyAllocatedMemory", 0),
		dllimports.NewEnergyImport("CefV8Value_GetArrayLength", 0),
		dllimports.NewEnergyImport("CefV8Value_GetArrayBufferReleaseCallback", 0),
		dllimports.NewEnergyImport("CefV8Value_NeuterArrayBuffer", 0),
		dllimports.NewEnergyImport("CefV8Value_GetFunctionName", 0),
		dllimports.NewEnergyImport("CefV8Value_GetFunctionHandler", 0),
		dllimports.NewEnergyImport("CefV8Value_ExecuteFunction", 0),
		dllimports.NewEnergyImport("CefV8Value_ExecuteFunctionWithContext", 0),
		dllimports.NewEnergyImport("CefV8Value_ResolvePromise", 0),
		dllimports.NewEnergyImport("CefV8Value_RejectPromise", 0),
		//TCefV8ValueRef
		dllimports.NewEnergyImport("CefV8ValueRef_NewUndefined", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewNull", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewBool", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewInt", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewUInt", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewDouble", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewDate", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewString", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewObject", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewArray", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewArrayBuffer", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewFunction", 0),
		dllimports.NewEnergyImport("CefV8ValueRef_NewPromise", 0),
		//ICefV8Accessor
		dllimports.NewEnergyImport("CefV8Accessor_Create", 0),
		dllimports.NewEnergyImport("CefV8Accessor_Get", 0),
		dllimports.NewEnergyImport("CefV8Accessor_Set", 0),
		dllimports.NewEnergyImport("CefV8Accessor_Destroy", 0),
		//ICefV8Handler
		dllimports.NewEnergyImport("CefV8Handler_Create", 0),
		dllimports.NewEnergyImport("CefV8Handler_Execute", 0),
		dllimports.NewEnergyImport("CefV8Handler_Destroy", 0),
		//ICefV8Interceptor
		dllimports.NewEnergyImport("CefV8InterceptorRef_Create", 0),
		dllimports.NewEnergyImport("CefV8Interceptor_GetByName", 0),
		dllimports.NewEnergyImport("CefV8Interceptor_GetByIndex", 0),
		dllimports.NewEnergyImport("CefV8Interceptor_SetByName", 0),
		dllimports.NewEnergyImport("CefV8Interceptor_SetByIndex", 0),
		dllimports.NewEnergyImport("CefV8Interceptor_Destroy", 0),
		//ICefV8Context
		dllimports.NewEnergyImport("CefV8Context_Eval", 0),
		dllimports.NewEnergyImport("CefV8Context_Enter", 0),
		dllimports.NewEnergyImport("CefV8Context_Exit", 0),
		dllimports.NewEnergyImport("CefV8Context_IsSame", 0),
		dllimports.NewEnergyImport("CefV8Context_Browser", 0),
		dllimports.NewEnergyImport("CefV8Context_Frame", 0),
		dllimports.NewEnergyImport("CefV8Context_Global", 0),
		dllimports.NewEnergyImport("CefV8ContextRef_Current", 0),
		dllimports.NewEnergyImport("CefV8ContextRef_Entered", 0),
		//ICefV8Exception
		dllimports.NewEnergyImport("CefV8Exception_Message", 0),
		dllimports.NewEnergyImport("CefV8Exception_SourceLine", 0),
		dllimports.NewEnergyImport("CefV8Exception_ScriptResourceName", 0),
		dllimports.NewEnergyImport("CefV8Exception_LineNumber", 0),
		dllimports.NewEnergyImport("CefV8Exception_StartPosition", 0),
		dllimports.NewEnergyImport("CefV8Exception_EndPosition", 0),
		dllimports.NewEnergyImport("CefV8Exception_StartColumn", 0),
		dllimports.NewEnergyImport("CefV8Exception_EndColumn", 0),
		//ICefV8ArrayBufferReleaseCallback
		dllimports.NewEnergyImport("CefV8ArrayBufferReleaseCallback_Create", 0),
		dllimports.NewEnergyImport("CefV8ArrayBufferReleaseCallback_ReleaseBuffer", 0),
		//ICefProcessMessage
		dllimports.NewEnergyImport("CefProcessMessageRef_New", 0),
		dllimports.NewEnergyImport("CefProcessMessage_ArgumentList", 0),
		dllimports.NewEnergyImport("CefProcessMessage_IsValid", 0),
		dllimports.NewEnergyImport("CefProcessMessage_Copy", 0),
		dllimports.NewEnergyImport("CefProcessMessage_Name", 0),
		//ICefListValue
		dllimports.NewEnergyImport("CefListValue_New", 0),
		dllimports.NewEnergyImport("CefListValue_IsValid", 0),
		dllimports.NewEnergyImport("CefListValue_IsOwned", 0),
		dllimports.NewEnergyImport("CefListValue_IsReadOnly", 0),
		dllimports.NewEnergyImport("CefListValue_Copy", 0),
		dllimports.NewEnergyImport("CefListValue_SetSize", 0),
		dllimports.NewEnergyImport("CefListValue_GetSize", 0),
		dllimports.NewEnergyImport("CefListValue_Clear", 0),
		dllimports.NewEnergyImport("CefListValue_Remove", 0),
		dllimports.NewEnergyImport("CefListValue_GetType", 0),
		dllimports.NewEnergyImport("CefListValue_GetValue", 0),
		dllimports.NewEnergyImport("CefListValue_GetBool", 0),
		dllimports.NewEnergyImport("CefListValue_GetInt", 0),
		dllimports.NewEnergyImport("CefListValue_GetDouble", 0),
		dllimports.NewEnergyImport("CefListValue_GetString", 0),
		dllimports.NewEnergyImport("CefListValue_GetBinary", 0),
		dllimports.NewEnergyImport("CefListValue_GetDictionary", 0),
		dllimports.NewEnergyImport("CefListValue_GetList", 0),
		dllimports.NewEnergyImport("CefListValue_SetValue", 0),
		dllimports.NewEnergyImport("CefListValue_SetNull", 0),
		dllimports.NewEnergyImport("CefListValue_SetBool", 0),
		dllimports.NewEnergyImport("CefListValue_SetInt", 0),
		dllimports.NewEnergyImport("CefListValue_SetDouble", 0),
		dllimports.NewEnergyImport("CefListValue_SetString", 0),
		dllimports.NewEnergyImport("CefListValue_SetBinary", 0),
		dllimports.NewEnergyImport("CefListValue_SetDictionary", 0),
		dllimports.NewEnergyImport("CefListValue_SetList", 0),
		//ICefValue
		dllimports.NewEnergyImport("CefValueRef_New", 0),
		dllimports.NewEnergyImport("CefValue_IsValid", 0),
		dllimports.NewEnergyImport("CefValue_IsOwned", 0),
		dllimports.NewEnergyImport("CefValue_IsReadOnly", 0),
		dllimports.NewEnergyImport("CefValue_Copy", 0),
		dllimports.NewEnergyImport("CefValue_GetType", 0),
		dllimports.NewEnergyImport("CefValue_GetBool", 0),
		dllimports.NewEnergyImport("CefValue_GetInt", 0),
		dllimports.NewEnergyImport("CefValue_GetDouble", 0),
		dllimports.NewEnergyImport("CefValue_GetString", 0),
		dllimports.NewEnergyImport("CefValue_GetBinary", 0),
		dllimports.NewEnergyImport("CefValue_GetDictionary", 0),
		dllimports.NewEnergyImport("CefValue_GetList", 0),
		dllimports.NewEnergyImport("CefValue_SetNull", 0),
		dllimports.NewEnergyImport("CefValue_SetBool", 0),
		dllimports.NewEnergyImport("CefValue_SetInt", 0),
		dllimports.NewEnergyImport("CefValue_SetDouble", 0),
		dllimports.NewEnergyImport("CefValue_SetString", 0),
		dllimports.NewEnergyImport("CefValue_SetBinary", 0),
		dllimports.NewEnergyImport("CefValue_SetDictionary", 0),
		dllimports.NewEnergyImport("CefValue_SetList", 0),
		//ICefBinaryValue
		dllimports.NewEnergyImport("CefBinaryValueRef_New", 0),
		dllimports.NewEnergyImport("CefBinaryValueRef_Create", 0),
		dllimports.NewEnergyImport("CefBinaryValue_IsValid", 0),
		dllimports.NewEnergyImport("CefBinaryValue_IsOwned", 0),
		dllimports.NewEnergyImport("CefBinaryValue_Copy", 0),
		dllimports.NewEnergyImport("CefBinaryValue_GetSize", 0),
		dllimports.NewEnergyImport("CefBinaryValue_GetData", 0),
		//ICefDictionaryValue
		dllimports.NewEnergyImport("CefDictionaryValueRef_New", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_IsValid", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_IsOwned", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_IsReadOnly", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_Copy", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetSize", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_Clear", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_HasKey", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetKeys", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_Remove", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetType", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetValue", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetBool", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetInt", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetDouble", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetString", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetBinary", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetDictionary", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_GetList", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetValue", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetNull", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetBool", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetInt", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetDouble", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetString", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetBinary", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetDictionary", 0),
		dllimports.NewEnergyImport("CefDictionaryValue_SetList", 0),
		//ICefPostData
		dllimports.NewEnergyImport("CefPostDataRef_Create", 0),
		dllimports.NewEnergyImport("CefPostData_IsReadOnly", 0),
		dllimports.NewEnergyImport("CefPostData_HasExcludedElements", 0),
		dllimports.NewEnergyImport("CefPostData_GetElementCount", 0),
		dllimports.NewEnergyImport("CefPostData_GetElements", 0),
		dllimports.NewEnergyImport("CefPostData_RemoveElement", 0),
		dllimports.NewEnergyImport("CefPostData_AddElement", 0),
		dllimports.NewEnergyImport("CefPostData_RemoveElements", 0),
		//ICefPostDataElement
		dllimports.NewEnergyImport("PostDataElementRef_Create", 0),
		dllimports.NewEnergyImport("PostDataElement_IsReadOnly", 0),
		dllimports.NewEnergyImport("PostDataElement_SetToEmpty", 0),
		dllimports.NewEnergyImport("PostDataElement_SetToFile", 0),
		dllimports.NewEnergyImport("PostDataElement_SetToBytes", 0),
		dllimports.NewEnergyImport("PostDataElement_GetType", 0),
		dllimports.NewEnergyImport("PostDataElement_GetFile", 0),
		dllimports.NewEnergyImport("PostDataElement_GetBytesCount", 0),
		dllimports.NewEnergyImport("PostDataElement_GetBytes", 0),
	}
	imports.SetEnergyImportDefs(energyImportDefs)
}

const (
	//null nil
	internale_null_nil = iota
	//CEF
	internale_CEFVersion
	internale_LibBuildVersion
	internale_CEF_Win_CreateRectRgn
	internale_CEF_Win_SetRectRgn
	internale_CEF_Win_DeleteObject
	internale_CEF_Win_CombineRgn
	internale_CEF_Win_SetDraggableRegions
	internale_CEF_Win_PtInRegion
	internale_CEF_Win_ScreenToClient
	internale_CEF_Win_ClientToScreen
	internale_CEF_Win_DefWindowProc
	internale_CEF_Win_DefSubclassProc
	internale_CEF_Win_CreateRoundRectRgn
	internale_CEF_Win_SetWindowRgn
	internale_CEF_Win_SetCursor
	internale_CEF_Win_LoadCursor
	internale_CEF_Win_OnPaint
	//ApplicationQueueAsyncCallFunc
	internale_SetApplicationQueueAsyncCallFunc
	internale_CEFApplication_QueueAsyncCall
	internale_SetCEFWindowBindCallbackFunc
	internale_SetCEFIPCCallbackFunc
	//GoForm
	internale_CEF_AddGoForm
	internale_CEF_RemoveGoForm
	//ICefCallback
	internale_cefCallback_Cont
	internale_cefCallback_Cancel
	//process
	internale_SetMacOSXCommandLine
	internale_CEFStartMainProcess
	internale_CEFStartSubProcess
	internale_AddCustomCommandLine
	//application
	internale_CEFApplication_RunMessageLoop
	internale_CEFApplication_QuitMessageLoop
	internale_CEFApplication_Create
	internale_CEFApplication_Destroy
	internale_CEFApplication_Free
	internale_CEFApplication_StopScheduler
	internale_CEFApplication_ExecuteJS
	internale_CEFV8ValueRef_SetCommonRootName
	internale_CEFV8ValueRef_SetObjectRootName
	internale_CEFV8ValueRef_CommonValueBindInfo
	internale_CEFV8ValueRef_ObjectValueBindInfo
	//application - event
	internale_CEFGlobalApp_SetOnContextCreated
	internale_CEFGlobalApp_SetOnRegCustomSchemes
	internale_CEFGlobalApp_SetOnContextInitialized
	internale_CEFGlobalApp_SetOnWebKitInitialized
	internale_CEFGlobalApp_SetOnBeforeChildProcessLaunch
	internale_CEFGlobalApp_SetOnScheduleMessagePumpWork
	internale_CEFGlobalApp_SetOnGetDefaultClient
	internale_CEFGlobalApp_SetOnGetLocalizedString
	internale_CEFGlobalApp_SetOnGetDataResource
	internale_CEFGlobalApp_SetOnGetDataResourceForScale
	internale_CEFGlobalApp_SetOnProcessMessageReceived
	internale_CEFGlobalApp_SetOnBrowserDestroyed
	internale_CEFGlobalApp_SetOnRenderLoadStart
	internale_CEFGlobalApp_SetOnRenderLoadEnd
	internale_CEFGlobalApp_SetOnRenderLoadError
	internale_CEFGlobalApp_SetOnRenderLoadingStateChange
	internale_CEFGlobalApp_SetOnBrowserCreated
	internale_CEFGlobalApp_SetOnContextReleased
	internale_CEFGlobalApp_SetOnUncaughtException
	internale_CEFGlobalApp_SetOnFocusedNodeChanged
	internale_CEFGlobalApp_SetOnLoadingStateChange
	internale_CEFGlobalApp_SetOnLoadStart
	internale_CEFGlobalApp_SetOnLoadEnd
	internale_CEFGlobalApp_SetOnLoadError
	//CEFParentWindow
	internale_CEFWindow_Create
	internale_CEFWindow_GetHandle
	internale_CEFWindow_DestroyChildWindow
	internale_CEFWindow_HandleAllocated
	internale_CEFWindow_CreateHandle
	internale_CEFWindow_Free
	internale_CEFWindow_SetParent
	internale_CEFWindow_GetAlign
	internale_CEFWindow_SetAlign
	internale_CEFWindow_GetAnchors
	internale_CEFWindow_SetAnchors
	internale_CEFWindow_GetVisible
	internale_CEFWindow_SetVisible
	internale_CEFWindow_GetEnabled
	internale_CEFWindow_SetEnabled
	internale_CEFWindow_GetLeft
	internale_CEFWindow_SetLeft
	internale_CEFWindow_GetTop
	internale_CEFWindow_SetTop
	internale_CEFWindow_GetWidth
	internale_CEFWindow_SetWidth
	internale_CEFWindow_GetHeight
	internale_CEFWindow_SetHeight
	internale_CEFWindow_GetBoundsRect
	internale_CEFWindow_SetBoundsRect
	internale_CEFWindow_GetName
	internale_CEFWindow_SetName
	internale_CEFWindow_UpdateSize
	internale_CEFWindow_OnEnter
	internale_CEFWindow_OnExit
	//internale_CEFWindow_SetChromium
	//CEFLinkedParentWindow
	internale_CEFLinkedWindow_Create
	internale_CEFLinkedWindow_GetHandle
	internale_CEFLinkedWindow_DestroyChildWindow
	internale_CEFLinkedWindow_HandleAllocated
	internale_CEFLinkedWindow_CreateHandle
	internale_CEFLinkedWindow_Free
	internale_CEFLinkedWindow_SetParent
	internale_CEFLinkedWindow_GetAlign
	internale_CEFLinkedWindow_SetAlign
	internale_CEFLinkedWindow_GetAnchors
	internale_CEFLinkedWindow_SetAnchors
	internale_CEFLinkedWindow_GetVisible
	internale_CEFLinkedWindow_SetVisible
	internale_CEFLinkedWindow_GetEnabled
	internale_CEFLinkedWindow_SetEnabled
	internale_CEFLinkedWindow_GetLeft
	internale_CEFLinkedWindow_SetLeft
	internale_CEFLinkedWindow_GetTop
	internale_CEFLinkedWindow_SetTop
	internale_CEFLinkedWindow_GetWidth
	internale_CEFLinkedWindow_SetWidth
	internale_CEFLinkedWindow_GetHeight
	internale_CEFLinkedWindow_SetHeight
	internale_CEFLinkedWindow_GetBoundsRect
	internale_CEFLinkedWindow_SetBoundsRect
	internale_CEFLinkedWindow_GetName
	internale_CEFLinkedWindow_SetName
	internale_CEFLinkedWindow_UpdateSize
	internale_CEFLinkedWindow_OnEnter
	internale_CEFLinkedWindow_OnExit
	internale_CEFLinkedWindow_SetChromium
	//ICefBrowser
	internale_CEFBrowser_ShowDevTools
	internale_CEFBrowser_CloseDevTools
	internale_CEFBrowser_HasDevTools
	internale_CEFBrowser_GetIdentifier
	internale_CEFBrowser_GetFocusedFrame
	internale_CEFBrowser_GetMainFrame
	internale_CEFBrowser_GetFrameById
	internale_CEFBrowser_GetFrameByName
	internale_CEFBrowser_GetHostWindowHandle
	internale_CEFBrowser_CloseBrowser
	internale_CEFBrowser_TryCloseBrowser
	internale_CEFBrowser_SetFocus
	internale_CEFBrowser_GetZoomLevel
	internale_CEFBrowser_SetZoomLevel
	internale_CEFBrowser_RunFileDialog
	internale_CEFBrowser_StartDownload
	internale_CEFBrowser_DownloadImage
	internale_CEFBrowser_Print
	internale_CEFBrowser_ExecuteDevToolsMethod
	internale_CEFBrowser_SendKeyEvent
	internale_CEFBrowser_SetAudioMuted
	internale_CEFBrowser_IsAudioMuted
	internale_CEFBrowser_SetAutoResizeEnabled
	internale_CEFBrowser_SetAccessibilityState
	internale_CEFBrowser_NotifyMoveOrResizeStarted
	internale_CEFBrowser_NotifyScreenInfoChanged
	internale_CEFBrowser_SendCaptureLostEvent
	internale_CEFBrowser_SendTouchEvent
	internale_CEFBrowser_SendMouseWheelEvent
	internale_CEFBrowser_SendMouseMoveEvent
	internale_CEFBrowser_SendMouseClickEvent
	internale_CEFBrowser_CanGoBack
	internale_CEFBrowser_GoBack
	internale_CEFBrowser_CanGoForward
	internale_CEFBrowser_GoForward
	internale_CEFBrowser_IsLoading
	internale_CEFBrowser_Reload
	internale_CEFBrowser_ReloadIgnoreCache
	internale_CEFBrowser_StopLoad
	internale_CEFBrowser_FrameCount
	internale_CEFBrowser_GetFrameNames
	internale_CEFBrowser_Find
	internale_CEFBrowser_StopFinding
	//TCEFChromium - event
	internale_CEFChromium_SetOnAfterCreated
	internale_CEFChromium_SetOnBeforeClose
	internale_CEFChromium_SetOnClose
	internale_CEFChromium_SetOnPdfPrintFinished
	internale_CEFChromium_SetOnZoomPctAvailable
	internale_CEFChromium_SetOnLoadStart
	internale_CEFChromium_SetOnLoadingStateChange
	internale_CEFChromium_SetOnLoadingProgressChange
	internale_CEFChromium_SetOnLoadError
	internale_CEFChromium_SetOnLoadEnd
	internale_CEFChromium_SetOnBeforeDownload
	internale_CEFChromium_SetOnDownloadUpdated
	internale_CEFChromium_SetOnFullScreenModeChange
	internale_CEFChromium_SetOnBeforeBrowse
	internale_CEFChromium_SetOnAddressChange
	internale_CEFChromium_SetOnKeyEvent
	internale_CEFChromium_SetOnTitleChange
	internale_CEFChromium_SetOnRenderCompMsg
	internale_CEFChromium_SetOnWidgetCompMsg
	internale_CEFChromium_SetOnBrowserCompMsg
	internale_CEFChromium_SetOnRenderProcessTerminated
	internale_CEFChromium_SetOnRenderViewReady
	internale_CEFChromium_SetOnScrollOffsetChanged
	internale_CEFChromium_SetOnProcessMessageReceived
	internale_CEFChromium_SetOnFindResult
	internale_CEFChromium_SetOnCookieSet
	internale_CEFChromium_SetOnCookiesDeleted
	internale_CEFChromium_SetOnCookiesFlushed
	internale_CEFChromium_SetOnCookiesVisited
	internale_CEFChromium_SetOnCookieVisitorDestroyed
	internale_CEFChromium_SetOnBeforeContextMenu
	internale_CEFChromium_SetOnContextMenuCommand
	internale_CEFChromium_SetOnContextMenuDismissed
	internale_CEFChromium_SetOnBeforeResourceLoad
	internale_CEFChromium_SetOnResourceResponse
	internale_CEFChromium_SetOnResourceRedirect
	internale_CEFChromium_SetOnResourceLoadComplete
	internale_CEFChromium_SetOnFrameAttached
	internale_CEFChromium_SetOnFrameCreated
	internale_CEFChromium_SetOnFrameDetached
	internale_CEFChromium_SetOnMainFrameChanged
	internale_CEFChromium_SetOnBeforePopup
	internale_CEFChromium_SetOnOpenUrlFromTab
	internale_CEFChromium_SetOnDragEnter
	internale_CEFChromium_SetOnDraggableRegionsChanged
	//TCEFChromium - proc
	internale_CEFChromium_Free
	internale_CEFChromium_GetHashCode
	internale_CEFChromium_ClassName
	internale_CEFChromium_Equals
	internale_CEFChromium_ClassType
	internale_CEFChromium_InstanceSize
	internale_CEFChromium_InheritsFrom
	internale_CEFChromium_ToString
	internale_CEFChromium_Create
	internale_CEFChromium_Browser
	internale_CEFChromium_BrowserById
	internale_CEFChromium_BrowserIdByIndex
	internale_CEFChromium_BrowserCount
	internale_CEFChromium_BrowserId
	internale_CEFChromium_SetDefaultURL
	internale_CEFChromium_SetMultiBrowserMode
	internale_CEFChromium_LoadURL
	internale_CEFChromium_LoadString
	internale_CEFChromium_StartDownload
	internale_CEFChromium_DownloadImage
	internale_CEFChromium_Reload
	internale_CEFChromium_StopLoad
	internale_CEFChromium_ResetZoomLevel
	internale_CEFChromium_CloseAllBrowsers
	internale_CEFChromium_CreateBrowserByWindow
	internale_CEFChromium_CreateBrowserByLinkedWindow
	internale_CEFChromium_CreateBrowserByBrowserViewComponent
	internale_CEFChromium_Initialized
	internale_CEFChromium_IsSameBrowser
	internale_CEFChromium_PrintToPDF
	internale_CEFChromium_Print
	internale_CEFChromium_BrowserDownloadCancel
	internale_CEFChromium_BrowserDownloadPause
	internale_CEFChromium_DownloadResume
	internale_CEFChromium_BrowserZoom
	internale_CEFChromium_GoBackForward
	internale_CEFChromium_NotifyMoveOrResizeStarted
	internale_CEFChromium_CloseBrowser
	internale_CEFChromium_ExecuteJavaScript
	internale_CEFChromium_ShowDevTools
	internale_CEFChromium_ShowDevToolsByWindowParent
	internale_CEFChromium_CloseDevTools
	internale_CEFChromium_CloseDevToolsByWindowParent
	internale_CEFChromium_VisitAllCookies
	internale_CEFChromium_VisitURLCookies
	internale_CEFChromium_DeleteCookies
	internale_CEFChromium_SetCookie
	internale_CEFChromium_SetProxy
	internale_CEFChromium_UpdatePreferences
	internale_CEFChromium_ExecuteDevToolsMethod
	internale_CEFChromium_CreateClientHandler
	internale_CEFChromium_SetFocus
	internale_CEFChromium_SendCaptureLostEvent
	internale_CEFChromium_FrameIsFocused
	internale_CEFChromium_TryCloseBrowser
	internale_CEFChromium_BrowserHandle
	internale_CEFChromium_WidgetHandle
	internale_CEFChromium_RenderHandle
	internale_CEFChromium_SetCustomHeader
	internale_CEFChromium_GetCustomHeader
	internale_CEFChromium_SetJavascriptEnabled
	internale_CEFChromium_GetJavascriptEnabled
	internale_CEFChromium_SetWebRTCIPHandlingPolicy
	internale_CEFChromium_GetWebRTCIPHandlingPolicy
	internale_CEFChromium_SetWebRTCMultipleRoutes
	internale_CEFChromium_GetWebRTCMultipleRoutes
	internale_CEFChromium_SetWebRTCNonproxiedUDP
	internale_CEFChromium_GetWebRTCNonproxiedUDP
	internale_CEFChromium_SetBatterySaverModeState
	internale_CEFChromium_GetBatterySaverModeState
	internale_CEFChromium_SetHighEfficiencyMode
	internale_CEFChromium_GetHighEfficiencyMode
	internale_CEFChromium_SetLoadImagesAutomatically
	internale_CEFChromium_GetLoadImagesAutomatically
	internale_CEFChromium_SetQuicAllowed
	internale_CEFChromium_GetQuicAllowed
	internale_CEFChromium_SetOffline
	internale_CEFChromium_GetOffline
	internale_CEFChromium_SetDefaultWindowInfoExStyle
	internale_CEFChromium_GetDefaultWindowInfoExStyle
	internale_CEFChromium_SetBlock3rdPartyCookies
	internale_CEFChromium_GetBlock3rdPartyCookies
	internale_CEFChromium_SetAcceptCookies
	internale_CEFChromium_GetAcceptCookies
	internale_CEFChromium_SetAcceptLanguageList
	internale_CEFChromium_GetAcceptLanguageList
	internale_CEFChromium_SetPrintingEnabled
	internale_CEFChromium_GetPrintingEnabled
	internale_CEFChromium_SetYouTubeRestrict
	internale_CEFChromium_GetYouTubeRestrict
	internale_CEFChromium_SetSafeSearch
	internale_CEFChromium_GetSafeSearch
	internale_CEFChromium_SetAudioMuted
	internale_CEFChromium_GetAudioMuted
	internale_CEFChromium_SetDragOperations
	internale_CEFChromium_GetDragOperations
	internale_CEFChromium_GetFrameCount
	internale_CEFChromium_SetSpellCheckerDicts
	internale_CEFChromium_GetSpellCheckerDicts
	internale_CEFChromium_SetSpellChecking
	internale_CEFChromium_GetSpellChecking
	internale_CEFChromium_SetAlwaysOpenPDFExternally
	internale_CEFChromium_GetAlwaysOpenPDFExternally
	internale_CEFChromium_SetAlwaysAuthorizePlugins
	internale_CEFChromium_GetAlwaysAuthorizePlugins
	internale_CEFChromium_SetAllowOutdatedPlugins
	internale_CEFChromium_GetAllowOutdatedPlugins
	internale_CEFChromium_SetSendReferrer
	internale_CEFChromium_GetSendReferrer
	internale_CEFChromium_SetDoNotTrack
	internale_CEFChromium_GetDoNotTrack
	internale_CEFChromium_SetZoomStep
	internale_CEFChromium_GetZoomStep
	internale_CEFChromium_SetZoomPct
	internale_CEFChromium_GetZoomPct
	internale_CEFChromium_SetZoomLevel
	internale_CEFChromium_GetZoomLevel
	internale_CEFChromium_SetDefaultEncoding
	internale_CEFChromium_GetDefaultEncoding
	internale_CEFChromium_SendProcessMessage
	//ICefBeforeDownloadCallback
	internale_CEFChromium_SetDownloadPath
	//ICefFrame
	internale_CEFFrame_Undo
	internale_CEFFrame_Redo
	internale_CEFFrame_Cut
	internale_CEFFrame_Copy
	internale_CEFFrame_Paste
	internale_CEFFrame_Del
	internale_CEFFrame_SelectAll
	internale_CEFFrame_ViewSource
	internale_CEFFrame_LoadUrl
	internale_CEFFrame_ExecuteJavaScript
	internale_CEFFrame_IsValid
	internale_CEFFrame_IsMain
	internale_CEFFrame_IsFocused
	internale_CEFFrame_SendProcessMessageByIPC
	internale_CEFFrame_SendProcessMessage
	internale_CEFFrame_LoadRequest
	internale_CEFFrame_Browser
	internale_CEFFrame_GetV8Context
	internale_CEFFrame_Identifier
	internale_CEFFrame_Name
	internale_CEFFrame_Url
	internale_CEFFrame_Parent
	//ICefMenuModel
	internale_cefMenuModel_AddSeparator
	internale_cefMenuModel_Clear
	internale_cefMenuModel_IsSubMenu
	internale_cefMenuModel_GetCount
	internale_cefMenuModel_AddItem
	internale_cefMenuModel_AddCheckItem
	internale_cefMenuModel_AddRadioItem
	internale_cefMenuModel_AddSubMenu
	internale_cefMenuModel_Remove
	internale_cefMenuModel_RemoveAt
	internale_cefMenuModel_SetChecked
	internale_cefMenuModel_IsChecked
	internale_cefMenuModel_SetColor
	internale_cefMenuModel_SetFontList
	internale_cefMenuModel_HasAccelerator
	internale_cefMenuModel_SetAccelerator
	internale_cefMenuModel_RemoveAccelerator
	internale_cefMenuModel_IsVisible
	internale_cefMenuModel_SetVisible
	internale_cefMenuModel_IsEnabled
	internale_cefMenuModel_SetEnabled
	internale_cefMenuModel_SetLabel
	internale_cefMenuModel_GetIndexOf
	//CEFWindowInfo
	internale_CEFWindowInfoAsChild
	internale_CEFWindowInfoAsPopUp
	internale_CEFWindowInfoAsWindowless
	internale_CEFJSRegisterExtension
	//ICefRequest
	internale_CefRequest_IsReadOnly
	internale_CefRequest_SetUrl
	internale_CefRequest_SetMethod
	internale_CefRequest_SetReferrer
	internale_CefRequest_SetFlags
	internale_CefRequest_SetFirstPartyForCookies
	internale_CefRequest_GetHeaderByName
	internale_CefRequest_SetHeaderByName
	internale_CefRequest_GetHeaderMap
	internale_CefRequest_SetHeaderMap
	internale_CefRequest_GetPostData
	internale_CefRequest_SetPostData
	//ICefResponse
	internale_CefResponse_IsReadOnly
	internale_CefResponse_SetError
	internale_CefResponse_SetStatus
	internale_CefResponse_SetStatusText
	internale_CefResponse_SetMimeType
	internale_CefResponse_SetCharset
	internale_CefResponse_GetHeaderByName
	internale_CefResponse_SetHeaderByName
	internale_CefResponse_SetURL
	internale_CefResponse_GetHeaderMap
	internale_CefResponse_SetHeaderMap
	//ICefStringMultiMap
	internale_StringMultimapRef_Create
	internale_StringMultimap_GetSize
	internale_StringMultimap_FindCount
	internale_StringMultimap_GetEnumerate
	internale_StringMultimap_GetKey
	internale_StringMultimap_GetValue
	internale_StringMultimap_Append
	internale_StringMultimap_Clear
	//ICefImage
	internale_CEFImage_New
	internale_CEFImage_AddPng
	internale_CEFImage_AddJpeg
	internale_CEFImage_GetWidth
	internale_CEFImage_GetHeight
	//TCEFWindowComponent
	internale_CEFWindowComponent_Create
	internale_CEFWindowComponent_CreateTopLevelWindow
	internale_CEFWindowComponent_Show
	internale_CEFWindowComponent_Hide
	internale_CEFWindowComponent_CenterWindow
	internale_CEFWindowComponent_Close
	internale_CEFWindowComponent_Activate
	internale_CEFWindowComponent_Deactivate
	internale_CEFWindowComponent_BringToTop
	internale_CEFWindowComponent_Maximize
	internale_CEFWindowComponent_Minimize
	internale_CEFWindowComponent_Restore
	internale_CEFWindowComponent_AddOverlayView
	internale_CEFWindowComponent_ShowMenu
	internale_CEFWindowComponent_CancelMenu
	internale_CEFWindowComponent_SetDraggableRegions
	internale_CEFWindowComponent_SendKeyPress
	internale_CEFWindowComponent_SendMouseMove
	internale_CEFWindowComponent_SendMouseEvents
	internale_CEFWindowComponent_SetAccelerator
	internale_CEFWindowComponent_RemoveAccelerator
	internale_CEFWindowComponent_RemoveAllAccelerators
	internale_CEFWindowComponent_SetAlwaysOnTop
	internale_CEFWindowComponent_SetFullscreen
	internale_CEFWindowComponent_SetBackgroundColor
	internale_CEFWindowComponent_Bounds
	internale_CEFWindowComponent_Size
	internale_CEFWindowComponent_Position
	internale_CEFWindowComponent_SetBounds
	internale_CEFWindowComponent_SetSize
	internale_CEFWindowComponent_SetPosition
	internale_CEFWindowComponent_SetTitle
	internale_CEFWindowComponent_Title
	internale_CEFWindowComponent_WindowIcon
	internale_CEFWindowComponent_SetWindowIcon
	internale_CEFWindowComponent_WindowAppIcon
	internale_CEFWindowComponent_SetWindowAppIcon
	internale_CEFWindowComponent_Display
	internale_CEFWindowComponent_ClientAreaBoundsInScreen
	internale_CEFWindowComponent_WindowHandle
	internale_CEFWindowComponent_IsClosed
	internale_CEFWindowComponent_IsActive
	internale_CEFWindowComponent_IsAlwaysOnTop
	internale_CEFWindowComponent_IsFullscreen
	internale_CEFWindowComponent_IsMaximized
	internale_CEFWindowComponent_IsMinimized
	internale_CEFWindowComponent_AddChildView
	internale_CEFWindowComponent_SetOnWindowCreated
	internale_CEFWindowComponent_SetOnWindowDestroyed
	internale_CEFWindowComponent_SetOnWindowActivationChanged
	internale_CEFWindowComponent_SetOnGetParentWindow
	internale_CEFWindowComponent_SetOnGetInitialBounds
	internale_CEFWindowComponent_SetOnGetInitialShowState
	internale_CEFWindowComponent_SetOnIsFrameless
	internale_CEFWindowComponent_SetOnCanResize
	internale_CEFWindowComponent_SetOnCanMaximize
	internale_CEFWindowComponent_SetOnCanMinimize
	internale_CEFWindowComponent_SetOnCanClose
	internale_CEFWindowComponent_SetOnAccelerator
	internale_CEFWindowComponent_SetOnKeyEvent
	//TCEFBrowserViewComponent
	internale_CEFBrowserViewComponent_Create
	internale_CEFBrowserViewComponent_CreateBrowserView
	internale_CEFBrowserViewComponent_GetForBrowser
	internale_CEFBrowserViewComponent_SetPreferAccelerators
	internale_CEFBrowserViewComponent_RequestFocus
	internale_CEFBrowserViewComponent_Browser
	internale_CEFBrowserViewComponent_BrowserView
	internale_CEFBrowserViewComponent_SetOnBrowserCreated
	internale_CEFBrowserViewComponent_SetOnBrowserDestroyed
	internale_CEFBrowserViewComponent_SetOnGetDelegateForPopupBrowserView
	internale_CEFBrowserViewComponent_SetOnPopupBrowserViewCreated
	internale_CEFBrowserViewComponent_SetOnGetChromeToolbarType
	//ICefDisplay
	internale_CEFDisplay_ID
	internale_CEFDisplay_DeviceScaleFactor
	internale_CEFDisplay_Rotation
	internale_CEFDisplay_Bounds
	internale_CEFDisplay_WorkArea
	//ICefWindow
	internale_ICEFWindow_Show
	internale_ICEFWindow_Hide
	internale_ICEFWindow_CenterWindow
	internale_ICEFWindow_Close
	internale_ICEFWindow_IsClosed
	internale_ICEFWindow_Activate
	internale_ICEFWindow_Deactivate
	internale_ICEFWindow_IsActive
	internale_ICEFWindow_BringToTop
	internale_ICEFWindow_SetAlwaysOnTop
	internale_ICEFWindow_IsAlwaysOnTop
	internale_ICEFWindow_Maximize
	internale_ICEFWindow_Minimize
	internale_ICEFWindow_Restore
	internale_ICEFWindow_SetFullscreen
	internale_ICEFWindow_SetBackgroundColor
	internale_ICEFWindow_SetBounds
	internale_ICEFWindow_SetSize
	internale_ICEFWindow_SetPosition
	internale_ICEFWindow_IsMaximized
	internale_ICEFWindow_IsMinimized
	internale_ICEFWindow_IsFullscreen
	internale_ICEFWindow_SetTitle
	internale_ICEFWindow_GetTitle
	internale_ICEFWindow_SetWindowIcon
	internale_ICEFWindow_GetWindowIcon
	internale_ICEFWindow_SetWindowAppIcon
	internale_ICEFWindow_GetWindowAppIcon
	internale_ICEFWindow_AddOverlayView
	internale_ICEFWindow_ShowMenu
	internale_ICEFWindow_CancelMenu
	internale_ICEFWindow_GetDisplay
	internale_ICEFWindow_GetClientAreaBoundsInScreen
	internale_ICEFWindow_SetDraggableRegions
	internale_ICEFWindow_GetWindowHandle
	internale_ICEFWindow_SendKeyPress
	internale_ICEFWindow_SendMouseMove
	internale_ICEFWindow_SendMouseEvents
	internale_ICEFWindow_SetAccelerator
	internale_ICEFWindow_RemoveAccelerator
	internale_ICEFWindow_RemoveAllAccelerators
	//ICefV8Value
	internale_CefV8Value_IsValid
	internale_CefV8Value_IsUndefined
	internale_CefV8Value_IsNull
	internale_CefV8Value_IsBool
	internale_CefV8Value_IsInt
	internale_CefV8Value_IsUInt
	internale_CefV8Value_IsDouble
	internale_CefV8Value_IsDate
	internale_CefV8Value_IsString
	internale_CefV8Value_IsObject
	internale_CefV8Value_IsArray
	internale_CefV8Value_IsArrayBuffer
	internale_CefV8Value_IsFunction
	internale_CefV8Value_IsPromise
	internale_CefV8Value_IsSame
	internale_CefV8Value_GetBoolValue
	internale_CefV8Value_GetIntValue
	internale_CefV8Value_GetUIntValue
	internale_CefV8Value_GetDoubleValue
	internale_CefV8Value_GetDateValue
	internale_CefV8Value_GetStringValue
	internale_CefV8Value_IsUserCreated
	internale_CefV8Value_HasException
	internale_CefV8Value_GetException
	internale_CefV8Value_ClearException
	internale_CefV8Value_WillRethrowExceptions
	internale_CefV8Value_SetRethrowExceptions
	internale_CefV8Value_HasValueByKey
	internale_CefV8Value_HasValueByIndex
	internale_CefV8Value_DeleteValueByKey
	internale_CefV8Value_DeleteValueByIndex
	internale_CefV8Value_GetValueByKey
	internale_CefV8Value_GetValueByIndex
	internale_CefV8Value_SetValueByKey
	internale_CefV8Value_SetValueByIndex
	internale_CefV8Value_SetValueByAccessor
	internale_CefV8Value_GetKeys
	internale_CefV8Value_SetUserData
	internale_CefV8Value_GetUserData
	internale_CefV8Value_GetExternallyAllocatedMemory
	internale_CefV8Value_AdjustExternallyAllocatedMemory
	internale_CefV8Value_GetArrayLength
	internale_CefV8Value_GetArrayBufferReleaseCallback
	internale_CefV8Value_NeuterArrayBuffer
	internale_CefV8Value_GetFunctionName
	internale_CefV8Value_GetFunctionHandler
	internale_CefV8Value_ExecuteFunction
	internale_CefV8Value_ExecuteFunctionWithContext
	internale_CefV8Value_ResolvePromise
	internale_CefV8Value_RejectPromise
	//TCefV8ValueRef
	internale_CefV8ValueRef_NewUndefined
	internale_CefV8ValueRef_NewNull
	internale_CefV8ValueRef_NewBool
	internale_CefV8ValueRef_NewInt
	internale_CefV8ValueRef_NewUInt
	internale_CefV8ValueRef_NewDouble
	internale_CefV8ValueRef_NewDate
	internale_CefV8ValueRef_NewString
	internale_CefV8ValueRef_NewObject
	internale_CefV8ValueRef_NewArray
	internale_CefV8ValueRef_NewArrayBuffer
	internale_CefV8ValueRef_NewFunction
	internale_CefV8ValueRef_NewPromise
	//ICefV8Accessor
	internale_CefV8Accessor_Create
	internale_CefV8Accessor_Get
	internale_CefV8Accessor_Set
	internale_CefV8Accessor_Destroy
	//ICefV8Handler
	internale_CefV8Handler_Create
	internale_CefV8Handler_Execute
	internale_CefV8Handler_Destroy
	//ICefV8Interceptor
	internale_CefV8InterceptorRef_Create
	internale_CefV8Interceptor_GetByName
	internale_CefV8Interceptor_GetByIndex
	internale_CefV8Interceptor_SetByName
	internale_CefV8Interceptor_SetByIndex
	internale_CefV8Interceptor_Destroy
	//ICefV8Context
	internale_CefV8Context_Eval
	internale_CefV8Context_Enter
	internale_CefV8Context_Exit
	internale_CefV8Context_IsSame
	internale_CefV8Context_Browser
	internale_CefV8Context_Frame
	internale_CefV8Context_Global
	internale_CefV8ContextRef_Current
	internale_CefV8ContextRef_Entered
	//ICefV8Exception
	internale_CefV8Exception_Message
	internale_CefV8Exception_SourceLine
	internale_CefV8Exception_ScriptResourceName
	internale_CefV8Exception_LineNumber
	internale_CefV8Exception_StartPosition
	internale_CefV8Exception_EndPosition
	internale_CefV8Exception_StartColumn
	internale_CefV8Exception_EndColumn
	//ICefV8ArrayBufferReleaseCallback
	internale_CefV8ArrayBufferReleaseCallback_Create
	internale_CefV8ArrayBufferReleaseCallback_ReleaseBuffer
	//ICefProcessMessage
	internale_CefProcessMessageRef_New
	internale_CefProcessMessage_ArgumentList
	internale_CefProcessMessage_IsValid
	internale_CefProcessMessage_Copy
	internale_CefProcessMessage_Name
	//ICefListValue
	internale_CefListValue_New
	internale_CefListValue_IsValid
	internale_CefListValue_IsOwned
	internale_CefListValue_IsReadOnly
	internale_CefListValue_Copy
	internale_CefListValue_SetSize
	internale_CefListValue_GetSize
	internale_CefListValue_Clear
	internale_CefListValue_Remove
	internale_CefListValue_GetType
	internale_CefListValue_GetValue
	internale_CefListValue_GetBool
	internale_CefListValue_GetInt
	internale_CefListValue_GetDouble
	internale_CefListValue_GetString
	internale_CefListValue_GetBinary
	internale_CefListValue_GetDictionary
	internale_CefListValue_GetList
	internale_CefListValue_SetValue
	internale_CefListValue_SetNull
	internale_CefListValue_SetBool
	internale_CefListValue_SetInt
	internale_CefListValue_SetDouble
	internale_CefListValue_SetString
	internale_CefListValue_SetBinary
	internale_CefListValue_SetDictionary
	internale_CefListValue_SetList
	//ICefValue
	internale_CefValueRef_New
	internale_CefValue_IsValid
	internale_CefValue_IsOwned
	internale_CefValue_IsReadOnly
	internale_CefValue_Copy
	internale_CefValue_GetType
	internale_CefValue_GetBool
	internale_CefValue_GetInt
	internale_CefValue_GetDouble
	internale_CefValue_GetString
	internale_CefValue_GetBinary
	internale_CefValue_GetDictionary
	internale_CefValue_GetList
	internale_CefValue_SetNull
	internale_CefValue_SetBool
	internale_CefValue_SetInt
	internale_CefValue_SetDouble
	internale_CefValue_SetString
	internale_CefValue_SetBinary
	internale_CefValue_SetDictionary
	internale_CefValue_SetList
	//ICefBinaryValue
	internale_CefBinaryValueRef_New
	internale_CefBinaryValueRef_Create
	internale_CefBinaryValue_IsValid
	internale_CefBinaryValue_IsOwned
	internale_CefBinaryValue_Copy
	internale_CefBinaryValue_GetSize
	internale_CefBinaryValue_GetData
	//ICefDictionaryValue
	internale_CefDictionaryValueRef_New
	internale_CefDictionaryValue_IsValid
	internale_CefDictionaryValue_IsOwned
	internale_CefDictionaryValue_IsReadOnly
	internale_CefDictionaryValue_Copy
	internale_CefDictionaryValue_GetSize
	internale_CefDictionaryValue_Clear
	internale_CefDictionaryValue_HasKey
	internale_CefDictionaryValue_GetKeys
	internale_CefDictionaryValue_Remove
	internale_CefDictionaryValue_GetType
	internale_CefDictionaryValue_GetValue
	internale_CefDictionaryValue_GetBool
	internale_CefDictionaryValue_GetInt
	internale_CefDictionaryValue_GetDouble
	internale_CefDictionaryValue_GetString
	internale_CefDictionaryValue_GetBinary
	internale_CefDictionaryValue_GetDictionary
	internale_CefDictionaryValue_GetList
	internale_CefDictionaryValue_SetValue
	internale_CefDictionaryValue_SetNull
	internale_CefDictionaryValue_SetBool
	internale_CefDictionaryValue_SetInt
	internale_CefDictionaryValue_SetDouble
	internale_CefDictionaryValue_SetString
	internale_CefDictionaryValue_SetBinary
	internale_CefDictionaryValue_SetDictionary
	internale_CefDictionaryValue_SetList
	//ICefPostData
	internale_CefPostDataRef_Create
	internale_CefPostData_IsReadOnly
	internale_CefPostData_HasExcludedElements
	internale_CefPostData_GetElementCount
	internale_CefPostData_GetElements
	internale_CefPostData_RemoveElement
	internale_CefPostData_AddElement
	internale_CefPostData_RemoveElements
	//ICefPostDataElement
	internale_PostDataElementRef_Create
	internale_PostDataElement_IsReadOnly
	internale_PostDataElement_SetToEmpty
	internale_PostDataElement_SetToFile
	internale_PostDataElement_SetToBytes
	internale_PostDataElement_GetType
	internale_PostDataElement_GetFile
	internale_PostDataElement_GetBytesCount
	internale_PostDataElement_GetBytes
)
