package cef

import "github.com/energye/golcl/lcl/api/dllimports"

func init() {
	var energyImportDefs = []*dllimports.ImportTable{
		//
		dllimports.NewEnergyImport("CEFApplication_GetCommonInstance", 0),
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
	}
	dllimports.SetEnergyImportDefs(energyImportDefs)
}

const (
	//CommonInstance
	internale_CEFApplication_GetCommonInstance = iota
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
	internale_CEFWindow_SetChromium
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
	//CEFBrowser
	CEFBrowser_GetHostWindowHandle
	CEFBrowser_CloseBrowser
	CEFBrowser_TryCloseBrowser
	CEFBrowser_SetFocus
	CEFBrowser_GetZoomLevel
	CEFBrowser_SetZoomLevel
	CEFBrowser_RunFileDialog
	CEFBrowser_StartDownload
	CEFBrowser_DownloadImage
	CEFBrowser_Print
	CEFBrowser_GetFocusedFrame
	CEFBrowser_GetMainFrame
	CEFBrowser_GetFrameById
	CEFBrowser_GetFrameByName
	CEFBrowser_ExecuteDevToolsMethod
	CEFBrowser_SendKeyEvent
	CEFBrowser_SetAudioMuted
	CEFBrowser_IsAudioMuted
	CEFBrowser_SetAutoResizeEnabled
	CEFBrowser_SetAccessibilityState
	CEFBrowser_NotifyMoveOrResizeStarted
	CEFBrowser_NotifyScreenInfoChanged
	CEFBrowser_SendCaptureLostEvent
	CEFBrowser_SendTouchEvent
	CEFBrowser_SendMouseWheelEvent
	CEFBrowser_SendMouseMoveEvent
	CEFBrowser_SendMouseClickEvent
	CEFBrowser_CloseDevTools
	CEFBrowser_HasDevTools
	CEFBrowser_CanGoBack
	CEFBrowser_GoBack
	CEFBrowser_CanGoForward
	CEFBrowser_GoForward
	CEFBrowser_IsLoading
	CEFBrowser_Reload
	CEFBrowser_ReloadIgnoreCache
	CEFBrowser_StopLoad
	CEFBrowser_FrameCount
	CEFBrowser_GetFrameNames
	CEFBrowser_Find
	CEFBrowser_StopFinding
	CEFBrowser_ShowDevTools
)
