//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import "github.com/energye/golcl/lcl/api/dllimports"

func init() {
	var energyImportDefs = []*dllimports.ImportTable{
		//null nil
		dllimports.NewEnergyImport("", 0),
		//ApplicationQueueAsyncCallFunc
		dllimports.NewEnergyImport("SetApplicationQueueAsyncCallFunc", 0),
		dllimports.NewEnergyImport("CEFApplication_QueueAsyncCall", 0),
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
		dllimports.NewEnergyImport("CEFApplication_Create", 0),
		dllimports.NewEnergyImport("CEFApplication_Free", 0),
		dllimports.NewEnergyImport("CEFApplication_ExecuteJS", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_SetCommonRootName", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_SetObjectRootName", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_CommonValueBindInfo", 0),
		dllimports.NewEnergyImport("CEFV8ValueRef_ObjectValueBindInfo", 0),
		dllimports.NewEnergyImport("CEFGlobalApp_SetOnContextCreated", 0),
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
		//CEFBrowser
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
		dllimports.NewEnergyImport("CEFBrowser_GetFocusedFrame", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetMainFrame", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetFrameById", 0),
		dllimports.NewEnergyImport("CEFBrowser_GetFrameByName", 0),
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
		dllimports.NewEnergyImport("CEFBrowser_CloseDevTools", 0),
		dllimports.NewEnergyImport("CEFBrowser_HasDevTools", 0),
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
		dllimports.NewEnergyImport("CEFBrowser_ShowDevTools", 0),
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
		dllimports.NewEnergyImport("CEFChromium_Initialized", 0),
		dllimports.NewEnergyImport("CEFChromium_GetBrowserId", 0),
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
		dllimports.NewEnergyImport("CEFFrame_SendProcessMessage", 0),
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
		//ICefRequest
		dllimports.NewEnergyImport("cefRequest_IsReadOnly", 0),
		dllimports.NewEnergyImport("cefRequest_SetUrl", 0),
		dllimports.NewEnergyImport("cefRequest_SetMethod", 0),
		dllimports.NewEnergyImport("cefRequest_SetReferrer", 0),
		dllimports.NewEnergyImport("cefRequest_SetFlags", 0),
		dllimports.NewEnergyImport("cefRequest_SetFirstPartyForCookies", 0),
		dllimports.NewEnergyImport("cefRequest_GetHeaderByName", 0),
		dllimports.NewEnergyImport("cefRequest_SetHeaderByName", 0),
		dllimports.NewEnergyImport("cefRequest_GetHeaderMap", 0),
		dllimports.NewEnergyImport("cefRequest_SetHeaderMap", 0),
		//ICefResponse
		dllimports.NewEnergyImport("cefResponse_IsReadOnly", 0),
		dllimports.NewEnergyImport("cefResponse_SetError", 0),
		dllimports.NewEnergyImport("cefResponse_SetStatus", 0),
		dllimports.NewEnergyImport("cefResponse_SetStatusText", 0),
		dllimports.NewEnergyImport("cefResponse_SetMimeType", 0),
		dllimports.NewEnergyImport("cefResponse_SetCharset", 0),
		dllimports.NewEnergyImport("cefResponse_GetHeaderByName", 0),
		dllimports.NewEnergyImport("cefResponse_SetHeaderByName", 0),
		dllimports.NewEnergyImport("cefResponse_SetURL", 0),
		dllimports.NewEnergyImport("cefResponse_GetHeaderMap", 0),
		//ICefStringMultiMap
		dllimports.NewEnergyImport("cefHeaderMap_GetSize", 0),
		dllimports.NewEnergyImport("cefHeaderMap_FindCount", 0),
		dllimports.NewEnergyImport("cefHeaderMap_GetEnumerate", 0),
		dllimports.NewEnergyImport("cefHeaderMap_GetKey", 0),
		dllimports.NewEnergyImport("cefHeaderMap_GetValue", 0),
		dllimports.NewEnergyImport("cefHeaderMap_Append", 0),
		dllimports.NewEnergyImport("cefHeaderMap_Clear", 0),
	}
	dllimports.SetEnergyImportDefs(energyImportDefs)
}

const (
	//null nil
	internale_null_nil = iota
	//ApplicationQueueAsyncCallFunc
	internale_SetApplicationQueueAsyncCallFunc
	internale_CEFApplication_QueueAsyncCall
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
	internale_CEFApplication_Create
	internale_CEFApplication_Free
	internale_CEFApplication_ExecuteJS
	internale_CEFV8ValueRef_SetCommonRootName
	internale_CEFV8ValueRef_SetObjectRootName
	internale_CEFV8ValueRef_CommonValueBindInfo
	internale_CEFV8ValueRef_ObjectValueBindInfo
	internale_CEFGlobalApp_SetOnContextCreated
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
	internale_CEFBrowser_GetFocusedFrame
	internale_CEFBrowser_GetMainFrame
	internale_CEFBrowser_GetFrameById
	internale_CEFBrowser_GetFrameByName
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
	internale_CEFBrowser_CloseDevTools
	internale_CEFBrowser_HasDevTools
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
	internale_CEFBrowser_ShowDevTools
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
	internale_CEFChromium_Initialized
	internale_CEFChromium_GetBrowserId
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
	internale_CEFFrame_SendProcessMessage
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
	//ICefRequest
	internale_cefRequest_IsReadOnly
	internale_cefRequest_SetUrl
	internale_cefRequest_SetMethod
	internale_cefRequest_SetReferrer
	internale_cefRequest_SetFlags
	internale_cefRequest_SetFirstPartyForCookies
	internale_cefRequest_GetHeaderByName
	internale_cefRequest_SetHeaderByName
	internale_cefRequest_GetHeaderMap
	internale_cefRequest_SetHeaderMap
	//ICefResponse
	internale_cefResponse_IsReadOnly
	internale_cefResponse_SetError
	internale_cefResponse_SetStatus
	internale_cefResponse_SetStatusText
	internale_cefResponse_SetMimeType
	internale_cefResponse_SetCharset
	internale_cefResponse_GetHeaderByName
	internale_cefResponse_SetHeaderByName
	internale_cefResponse_SetURL
	internale_cefResponse_GetHeaderMap
	//ICefStringMultiMap
	internale_cefHeaderMap_GetSize
	internale_cefHeaderMap_FindCount
	internale_cefHeaderMap_GetEnumerate
	internale_cefHeaderMap_GetKey
	internale_cefHeaderMap_GetValue
	internale_cefHeaderMap_Append
	internale_cefHeaderMap_Clear
)
