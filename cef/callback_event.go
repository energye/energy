//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/emfs"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"unsafe"
)

// getPointerOffset 指针偏移地址的指针值
func getParamOf(index int, ptr uintptr) uintptr {
	return *(*uintptr)(getPointerOffset(ptr, uintptr(index)*unsafe.Sizeof(ptr)))
}

// getPointerOffset 指针偏移地址
func getPointerOffset(ptr uintptr, offset uintptr) unsafePointer {
	return unsafePointer(ptr + offset)
}

func removeEventCallbackProc(f uintptr) uintptr {
	RemoveEventElement(f)
	return 0
}

// 回调过程
func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	fn := PtrToElementValue(f)
	if fn != nil {
		// 获取值
		getVal := func(i int) uintptr {
			return getParamOf(i, args)
		}

		// 指针
		getPtr := func(i int) unsafePointer {
			return getPointer(getVal(i))
		}

		// 指针的值
		getPtrVal := func(i int) uintptr {
			return *(*uintptr)(getPtr(i))
		}

		setPtrVal := func(i int, val uintptr) {
			*(*uintptr)(getPtr(i)) = val
		}

		getI32Ptr := func(i int) *int32 {
			return (*int32)(getPtr(i))
		}

		getBoolPtr := func(i int) *bool {
			return (*bool)(getPtr(i))
		}

		getPointPtr := func(i int) *TPoint {
			return (*TPoint)(getPtr(i))
		}

		getRectPtr := func(i int) *TRect {
			return (*TRect)(getPtr(i))
		}

		switch fn.(type) {

		// 通用 browser
		case TOnBrowser:
			browse := AsCefBrowser(getVal(1))
			fn.(TOnBrowser)(lcl.AsObject(getPtr(0)), browse)

		// 通用  Result float
		case TOnFloat:
			value := *(*float64)(getPtr(1))
			fn.(TOnFloat)(lcl.AsObject(getPtr(0)), value)

		// 通用  Result bool
		case TOnBool:
			value := GoBool(getVal(1))
			fn.(TOnBool)(lcl.AsObject(getPtr(0)), value)

		case TOnInt32:
			value := int32(getVal(1))
			fn.(TOnInt32)(lcl.AsObject(getPtr(0)), value)

		case TNotify:
			fn.(TNotify)(lcl.AsObject(getPtr(0)))

		// ===========  ICefRenderHandler  ===========

		case TOnGetAccessibilityHandler:
			fn.(TOnGetAccessibilityHandler)(lcl.AsObject(getPtr(0)), AsAccessibilityHandler(getVal(1)))

		case TOnGetRootScreenRect:
			browser := AsCefBrowser(getVal(1))
			rect := (*TCefRect)(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnGetRootScreenRect)(lcl.AsObject(getPtr(0)), browser, rect, result)

		case TOnGetViewRect:
			browser := AsCefBrowser(getVal(1))
			rect := (*TCefRect)(getPtr(2))
			fn.(TOnGetViewRect)(lcl.AsObject(getPtr(0)), browser, rect)

		case TOnGetScreenPoint:
			browser := AsCefBrowser(getVal(1))
			viewX, viewY := int32(getVal(2)), int32(getVal(3))
			screenX, screenY, result := (*int32)(getPtr(4)), (*int32)(getPtr(5)), (*bool)(getPtr(6))
			fn.(TOnGetScreenPoint)(lcl.AsObject(getPtr(0)), browser, viewX, viewY, screenX, screenY, result)

		case TOnGetScreenInfo:
			browser := AsCefBrowser(getVal(1))
			screenInfo := (*TCefScreenInfo)(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnGetScreenInfo)(lcl.AsObject(getPtr(0)), browser, screenInfo, result)

		case TOnPopupShow:
			browser := AsCefBrowser(getVal(1))
			show := GoBool(getVal(2))
			fn.(TOnPopupShow)(lcl.AsObject(getPtr(0)), browser, show)

		case TOnPopupSize:
			browser := AsCefBrowser(getVal(1))
			rect := *(*TCefRect)(getPtr(2))
			fn.(TOnPopupSize)(lcl.AsObject(getPtr(0)), browser, rect)

		case TOnPaint:
			browser := AsCefBrowser(getVal(1))
			kind := TCefPaintElementType(getVal(2))
			dirtyRectsCount := int(getVal(3))
			dirtyRectsPtr := getVal(4)
			buffer := getVal(5)
			width, height := int32(getVal(6)), int32(getVal(7))
			fn.(TOnPaint)(lcl.AsObject(getPtr(0)), browser, kind, NewCefRectArray(dirtyRectsCount, dirtyRectsPtr), buffer, width, height)

		case TOnAcceleratedPaint:
			browser := AsCefBrowser(getVal(1))
			kind := TCefPaintElementType(getVal(2))
			dirtyRectsCount := int(getVal(3))
			dirtyRectsPtr := getVal(4)
			sharedHandle := getVal(5)
			fn.(TOnAcceleratedPaint)(lcl.AsObject(getPtr(0)), browser, kind, NewCefRectArray(dirtyRectsCount, dirtyRectsPtr), sharedHandle)

		case TOnGetTouchHandleSize:
			browser := AsCefBrowser(getVal(1))
			orientation := TCefHorizontalAlignment(getVal(2))
			sizePtr := (*TCefSize)(getPtr(3))
			fn.(TOnGetTouchHandleSize)(lcl.AsObject(getPtr(0)), browser, orientation, sizePtr)

		case TOnTouchHandleStateChanged:
			browser := AsCefBrowser(getVal(1))
			state := *(*TCefTouchHandleState)(getPtr(2))
			fn.(TOnTouchHandleStateChanged)(lcl.AsObject(getPtr(0)), browser, state)

		case TOnStartDragging:
			browser := AsCefBrowser(getVal(1))
			dragData := AsCefDragData(getPtr(2))
			allowedOps := TCefDragOperations(getVal(3))
			x, y := int32(getVal(4)), int32(getVal(5))
			result := (*bool)(getPtr(6))
			fn.(TOnStartDragging)(lcl.AsObject(getPtr(0)), browser, dragData, allowedOps, x, y, result)

		case TOnUpdateDragCursor:
			browser := AsCefBrowser(getVal(1))
			operation := TCefDragOperation(getVal(2))
			fn.(TOnUpdateDragCursor)(lcl.AsObject(getPtr(0)), browser, operation)

		case TOnScrollOffsetChanged:
			browse := AsCefBrowser(getVal(1))
			fn.(TOnScrollOffsetChanged)(lcl.AsObject(getPtr(0)), browse, float64(getVal(2)), float64(getVal(3)))

		case TOnIMECompositionRangeChanged:
			browser := AsCefBrowser(getVal(1))
			rng := *(*TCefRange)(getPtr(2))
			characterBoundsCount := uint32(getVal(3))
			characterBounds := *(*TCefRect)(getPtr(4))
			fn.(TOnIMECompositionRangeChanged)(lcl.AsObject(getPtr(0)), browser, rng, characterBoundsCount, characterBounds)

		case TOnTextSelectionChanged:
			browser := AsCefBrowser(getVal(1))
			selectedText := GoStr(getVal(2))
			selectedRange := *(*TCefRange)(getPtr(3))
			fn.(TOnTextSelectionChanged)(lcl.AsObject(getPtr(0)), browser, selectedText, selectedRange)

		case TOnVirtualKeyboardRequested:
			browser := AsCefBrowser(getVal(1))
			inputMode := TCefTextInpuMode(getVal(2))
			fn.(TOnVirtualKeyboardRequested)(lcl.AsObject(getPtr(0)), browser, inputMode)

		// ===========  ICefRenderHandler  ===========

		case TOnDragEnter:
			sender := getPtr(0)
			browser := AsCefBrowser(getVal(1))
			dragData := AsCefDragData(getPtr(2))
			mask := TCefDragOperations(getVal(3))
			result := (*bool)(getPtr(4))
			fn.(TOnDragEnter)(lcl.AsObject(sender), browser, dragData, mask, result)

		case TOnDraggableRegionsChanged:
			sender := getPtr(0)
			browser := AsCefBrowser(getVal(1))
			frame := AsCefFrame(getPtr(2))
			regions := NewCefDraggableRegions(int(getVal(3)), getVal(4))
			fn.(TOnDraggableRegionsChanged)(lcl.AsObject(sender), browser, frame, regions)

		// ===========  ICefFindHandler  ===========

		case TOnFindResult:
			browser := AsCefBrowser(getVal(1))
			fn.(TOnFindResult)(lcl.AsObject(getPtr(0)), browser, int32(getVal(2)), int32(getVal(3)),
				*(*TCefRect)(getPtr(4)), int32(getVal(5)), GoBool(getVal(6)))

		// ===========  ICefRequestContextHandler  ===========

		case TOnRequestContextInitialized:
			fn.(TOnRequestContextInitialized)(lcl.AsObject(getPtr(0)), AsCefRequestContext(getPtr(1)))

		// ===========  ICefMediaObserver  ===========

		case TOnSinks:
			ptr := getVal(1)
			count := int(int32(getVal(2)))
			fn.(TOnSinks)(lcl.AsObject(getPtr(0)), MediaSinkArrayRef.New(count, ptr))

		case TOnRoutes:
			ptr := getVal(1)
			count := int(int32(getVal(2)))
			fn.(TOnRoutes)(lcl.AsObject(getPtr(0)), MediaRouteArrayRef.New(count, ptr))

		case TOnRouteStateChanged:
			route := AsCefMediaRoute(getPtr(1))
			state := TCefMediaRouteConnectionState(getVal(2))
			fn.(TOnRouteStateChanged)(lcl.AsObject(getPtr(0)), route, state)

		case TOnRouteMessageReceived:
			route := AsCefMediaRoute(getPtr(1))
			message := GoStr(getVal(2))
			fn.(TOnRouteMessageReceived)(lcl.AsObject(getPtr(0)), route, message)

		// ===========  ICefAudioHandler  ===========

		case TOnGetAudioParameters:
			browse := AsCefBrowser(getVal(1))
			params := (*TCefAudioParameters)(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnGetAudioParameters)(lcl.AsObject(getPtr(0)), browse, params, result)

		case TOnAudioStreamStarted:
			browse := AsCefBrowser(getVal(1))
			params := *(*TCefAudioParameters)(getPtr(2))
			channels := int32(getVal(3))
			fn.(TOnAudioStreamStarted)(lcl.AsObject(getPtr(0)), browse, params, channels)

		case TOnAudioStreamPacket:
			browse := AsCefBrowser(getVal(1))
			data := PPSingle(getPtr(2))
			frames := int32(getVal(3))
			pts := *(*int64)(getPtr(4))
			fn.(TOnAudioStreamPacket)(lcl.AsObject(getPtr(0)), browse, data, frames, pts)
		//case TOnAudioStreamStopped = TOnBrowser

		case TOnAudioStreamError:
			browse := AsCefBrowser(getVal(1))
			fn.(TOnAudioStreamError)(lcl.AsObject(getPtr(0)), browse, GoStr(getVal(2)))

		// ===========  ICefDevToolsMessageObserver  ===========

		case TOnDevToolsMessage:
			browse := AsCefBrowser(getVal(1))
			message := AsCefValue(getPtr(2))
			handled := (*bool)(getPtr(3))
			fn.(TOnDevToolsMessage)(lcl.AsObject(getPtr(0)), browse, message, handled)

		case TOnDevToolsRawMessage:
			browse := AsCefBrowser(getVal(1))
			message := getVal(2)
			messageSize := uint32(getVal(3))
			handled := (*bool)(getPtr(4))
			fn.(TOnDevToolsRawMessage)(lcl.AsObject(getPtr(0)), browse, message, messageSize, handled)

		case TOnDevToolsMethodResult:
			browse := AsCefBrowser(getVal(1))
			messageId := int32(getVal(2))
			success := GoBool(getVal(3))
			result := AsCefValue(getPtr(1))
			fn.(TOnDevToolsMethodResult)(lcl.AsObject(getPtr(0)), browse, messageId, success, result)

		case TOnDevToolsMethodRawResult:
			browse := AsCefBrowser(getVal(1))
			messageId := int32(getVal(2))
			success := GoBool(getVal(3))
			result := getVal(4)
			resultSize := uint32(getVal(5))
			fn.(TOnDevToolsMethodRawResult)(lcl.AsObject(getPtr(0)), browse, messageId, success, result, resultSize)

		case TOnDevToolsEvent:
			browse := AsCefBrowser(getVal(1))
			method := GoStr(getVal(2))
			params := AsCefValue(getPtr(3))
			fn.(TOnDevToolsEvent)(lcl.AsObject(getPtr(0)), browse, method, params)

		case TOnDevToolsRawEvent:
			browse := AsCefBrowser(getVal(1))
			method := GoStr(getVal(2))
			params := getVal(3)
			paramsSize := uint32(getVal(4))
			fn.(TOnDevToolsRawEvent)(lcl.AsObject(getPtr(0)), browse, method, params, paramsSize)

		//case TOnDevToolsAgentAttached = TOnBrowser
		//case TOnDevToolsAgentDetached = TOnBrowser

		// ===========  ICefExtensionHandler  ===========

		case TOnExtensionLoadFailed:
			fn.(TOnExtensionLoadFailed)(lcl.AsObject(getPtr(0)), TCefErrorCode(getVal(1)))

		case TOnExtensionLoaded:
			fn.(TOnExtensionLoaded)(lcl.AsObject(getPtr(0)), AsCefExtension(getVal(1)))

		case TOnExtensionUnloaded:
			fn.(TOnExtensionUnloaded)(lcl.AsObject(getPtr(0)), AsCefExtension(getVal(1)))

		case TOnBeforeBackgroundBrowser:
			extens := AsCefExtension(getVal(1))
			url := GoStr(getVal(2))
			clientPtr := (*uintptr)(getPtr(3))
			browserSettingsPtr := (*tCefBrowserSettings)(getPtr(4))
			result := (*bool)(getPtr(5))
			var client ICefClient
			var browserSettings = browserSettingsPtr.Convert()
			client, *result = fn.(TOnBeforeBackgroundBrowser)(lcl.AsObject(getPtr(0)), extens, url, browserSettings)
			if client != nil {
				*clientPtr = client.Instance()
			}
			browserSettings.SetInstanceValue()

		case TOnBeforeBrowser:
			extens := AsCefExtension(getVal(1))
			browse, activeBrowser := AsCefBrowser(getVal(2)), AsCefBrowser(getVal(3))
			index := int32(getVal(4))
			url := GoStr(getVal(5))
			active := GoBool(getVal(6))
			windowInfoPtr := (*tCefWindowInfo)(getPtr(7))
			clientPtr := (*uintptr)(getPtr(8))
			browserSettingsPtr := (*tCefBrowserSettings)(getPtr(9))
			result := (*bool)(getPtr(10))
			var (
				windowInfo      = windowInfoPtr.Convert()
				client          ICefClient
				browserSettings = browserSettingsPtr.Convert()
			)
			client, *result = fn.(TOnBeforeBrowser)(lcl.AsObject(getPtr(0)), extens, browse, activeBrowser, index, url, active, windowInfo, browserSettings)
			windowInfo.SetInstanceValue()
			if client != nil && client.Instance() > 0 {
				*clientPtr = client.Instance()
			}
			browserSettings.SetInstanceValue()

		case TOnGetActiveBrowser:
			extens := AsCefExtension(getVal(1))
			browse := AsCefBrowser(getVal(2))
			includeIncognito := GoBool(getVal(3))
			resultBrowser := (*uintptr)(getPtr(4))
			var activeBrowser ICefBrowser
			activeBrowser = fn.(TOnGetActiveBrowser)(lcl.AsObject(getPtr(0)), extens, browse, includeIncognito)
			if activeBrowser != nil && activeBrowser.Instance() > 0 {
				*resultBrowser = activeBrowser.Instance()
			}

		case TOnCanAccessBrowser:
			extens := AsCefExtension(getVal(1))
			browse := AsCefBrowser(getVal(2))
			includeIncognito := GoBool(getVal(3))
			targetBrowser := AsCefBrowser(getVal(4))
			result := (*bool)(getPtr(5))
			fn.(TOnCanAccessBrowser)(lcl.AsObject(getPtr(0)), extens, browse, includeIncognito, targetBrowser, result)

		case TOnGetExtensionResource:
			extens := AsCefExtension(getVal(1))
			browse := AsCefBrowser(getVal(2))
			file := GoStr(getVal(3))
			callback := AsCefGetExtensionResourceCallback(getVal(4))
			result := (*bool)(getPtr(5))
			fn.(TOnGetExtensionResource)(lcl.AsObject(getPtr(0)), extens, browse, file, callback, result)

		// ===========  ICefExtensionHandler  ===========

		//case TOnPrintStart = TOnBrowser

		case TOnPrintSettings:
			browse := AsCefBrowser(getVal(1))
			settings := AsCefPrintSettings(getVal(2))
			getDefaults := GoBool(getVal(3))
			fn.(TOnPrintSettings)(lcl.AsObject(getPtr(0)), browse, settings, getDefaults)

		case TOnPrintDialog:
			browse := AsCefBrowser(getVal(1))
			hasSelection := GoBool(getVal(2))
			callback := AsCefPrintDialogCallback(getVal(3))
			result := (*bool)(getPtr(4))
			fn.(TOnPrintDialog)(lcl.AsObject(getPtr(0)), browse, hasSelection, callback, result)

		case TOnPrintJob:
			browse := AsCefBrowser(getVal(1))
			documentName, PDFFilePath := GoStr(getVal(2)), GoStr(getVal(3))
			callback := AsCefPrintJobCallback(getVal(4))
			result := (*bool)(getPtr(5))
			fn.(TOnPrintJob)(lcl.AsObject(getPtr(0)), browse, documentName, PDFFilePath, callback, result)

		//case TOnPrintReset = TOnBrowser

		case TOnGetPDFPaperSize:
			browse := AsCefBrowser(getVal(1))
			deviceUnitsPerInch := int32(getVal(2))
			resultSize := (*TCefSize)(getPtr(3))
			fn.(TOnGetPDFPaperSize)(lcl.AsObject(getPtr(0)), browse, deviceUnitsPerInch, resultSize)

		// ===========  ICefFrameHandler  ===========

		case TOnFrameCreated:
			browse := AsCefBrowser(getVal(1))
			frame := AsCefFrame(getVal(2))
			fn.(TOnFrameCreated)(lcl.AsObject(getPtr(0)), browse, frame)

		case TOnFrameAttached:
			browse := AsCefBrowser(getVal(1))
			frame := AsCefFrame(getVal(2))
			fn.(TOnFrameAttached)(lcl.AsObject(getPtr(0)), browse, frame, GoBool(getVal(3)))

		case TOnFrameDetached:
			browse := AsCefBrowser(getVal(1))
			frame := AsCefFrame(getVal(2))
			fn.(TOnFrameDetached)(lcl.AsObject(getPtr(0)), browse, frame)

		case TOnMainFrameChanged:
			sender := getPtr(0)
			browse := AsCefBrowser(getVal(1))
			oldFrame := AsCefFrame(getVal(2))
			newFrame := AsCefFrame(getVal(3))
			fn.(TOnMainFrameChanged)(lcl.AsObject(sender), browse, oldFrame, newFrame)

		// ===========  ICefCommandHandler  ===========

		case TOnChromeCommand:
			browse := AsCefBrowser(getVal(1))
			params := (int32)(getVal(2))
			disposition := TCefWindowOpenDisposition(getVal(3))
			result := (*bool)(getPtr(4))
			fn.(TOnChromeCommand)(lcl.AsObject(getPtr(0)), browse, params, disposition, result)

		case TOnIsChromeAppMenuItemVisible:
			browser := AsCefBrowser(getVal(1))
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(3))
			fn.(TOnIsChromeAppMenuItemVisible)(lcl.AsObject(getPtr(0)), browser, commandId, result)

		case TOnIsChromeAppMenuItemEnabled:
			browser := AsCefBrowser(getVal(1))
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(3))
			fn.(TOnIsChromeAppMenuItemEnabled)(lcl.AsObject(getPtr(0)), browser, commandId, result)

		case TOnIsChromePageActionIconVisible:
			buttonType := TCefChromePageActionIconType(getVal(1))
			result := (*bool)(getPtr(2))
			fn.(TOnIsChromePageActionIconVisible)(lcl.AsObject(getPtr(0)), buttonType, result)

		case TOnIsChromeToolbarButtonVisible:
			iconType := TCefChromeToolbarButtonType(getVal(1))
			result := (*bool)(getPtr(2))
			fn.(TOnIsChromeToolbarButtonVisible)(lcl.AsObject(getPtr(0)), iconType, result)

		// ===========  ICefPermissionHandler  ===========

		case TOnRequestMediaAccessPermission:
			browse := AsCefBrowser(getVal(1))
			frame := AsCefFrame(getVal(2))
			requestingOrigin := GoStr(getVal(3))
			requestedPermissions := uint32(getVal(4))
			callback := AsCefMediaAccessCallback(getVal(5))
			result := (*bool)(getPtr(6))
			fn.(TOnRequestMediaAccessPermission)(lcl.AsObject(getPtr(0)), browse, frame, requestingOrigin, requestedPermissions, callback, result)

		case TOnShowPermissionPrompt:
			browse := AsCefBrowser(getVal(1))
			promptId := *(*uint64)(getPtr(2))
			requestingOrigin := GoStr(getVal(3))
			requestedPermissions := uint32(getVal(4))
			callback := AsCefPermissionPromptCallback(getVal(5))
			result := (*bool)(getPtr(6))
			fn.(TOnShowPermissionPrompt)(lcl.AsObject(getPtr(0)), browse, promptId, requestingOrigin, requestedPermissions, callback, result)

		case TOnDismissPermissionPrompt:
			browse := AsCefBrowser(getVal(1))
			promptId := *(*uint64)(getPtr(2))
			result := TCefPermissionRequestResult(getVal(3))
			fn.(TOnDismissPermissionPrompt)(lcl.AsObject(getPtr(0)), browse, promptId, result)

		// ===========  TCustomXxx  ===========

		case TOnTextResultAvailable:
			text := GoStr(getVal(1))
			fn.(TOnTextResultAvailable)(lcl.AsObject(getPtr(0)), text)
		//case TOnPdfPrintFinished = TOnBool
		//case TOnPrefsAvailable = TOnBool
		//case TOnCookiesDeleted = TOnInt32

		case TOnResolvedIPsAvailable:
			result := TCefErrorCode(getVal(1))
			resolvedIpsList := lcl.AsStrings(getVal(2))
			fn.(TOnResolvedIPsAvailable)(lcl.AsObject(getPtr(0)), result, resolvedIpsList)

		case TOnNavigationVisitorResultAvailable:
			entry := AsCefNavigationEntry(getVal(1))
			current := GoBool(getVal(2))
			index, total := int32(getVal(3)), int32(getVal(4))
			result := (*bool)(getPtr(5))
			fn.(TOnNavigationVisitorResultAvailable)(lcl.AsObject(getPtr(0)), entry, current, index, total, result)

		case TOnDownloadImageFinished:
			imageUrl := GoStr(getVal(1))
			httpStatusCode := int32(getVal(2))
			image := AsCefImage(getVal(3))
			fn.(TOnDownloadImageFinished)(lcl.AsObject(getPtr(0)), imageUrl, httpStatusCode, image)

		case TOnExecuteTaskOnCefThread:
			taskID := uint32(getVal(1))
			fn.(TOnExecuteTaskOnCefThread)(lcl.AsObject(getPtr(0)), taskID)

		case TOnCookiesVisited:
			cookiePtr := *(*tCookie)(getPtr(1))
			count := int32(getVal(2))
			total := int32(getVal(3))
			id := int32(getVal(4))
			deleteCookie := (*bool)(getPtr(5))
			result := (*bool)(getPtr(6))
			cookie := cookiePtr.Convert()
			fn.(TOnCookiesVisited)(lcl.AsObject(getVal(0)), *cookie, count, total, id, deleteCookie, result)

		//case TOnCookieVisitorDestroyed = TOnInt32

		case TOnCookieSet:
			success := GoBool(getVal(1))
			ID := int32(getVal(2))
			fn.(TOnCookieSet)(lcl.AsObject(getVal(0)), success, ID)

		//case TOnZoomPctAvailable = TOnFloat

		case TOnMediaRouteCreateFinished:
			result := TCefMediaRouterCreateResult(getVal(1))
			error := GoStr(getVal(2))
			route := AsCefMediaRoute(getVal(3))
			fn.(TOnMediaRouteCreateFinished)(lcl.AsObject(getPtr(0)), result, error, route)

		case TOnMediaSinkDeviceInfo:
			ipAddress := GoStr(getVal(1))
			port := int32(getVal(2))
			modelName := GoStr(getVal(3))
			fn.(TOnMediaSinkDeviceInfo)(lcl.AsObject(getPtr(0)), ipAddress, port, modelName)

		// ===========  Windows  ===========

		case TOnCompMsg:
			message := (*types.TMessage)(getPtr(1))
			lResultPtr := (*types.LRESULT)(getPtr(2))
			fn.(TOnCompMsg)(lcl.AsObject(getVal(0)), message, lResultPtr, (*bool)(getPtr(3)))

		// ===========  ICefLoadHandler  ===========

		case TOnLoadStart:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			transitionType := TCefTransitionType(getVal(3))
			fn.(TOnLoadStart)(lcl.AsObject(getPtr(0)), browse, frame, transitionType)

		case TOnLoadEnd:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			fn.(TOnLoadEnd)(lcl.AsObject(getPtr(0)), browse, frame, int32(getVal(3)))

		case TOnLoadError:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			fn.(TOnLoadError)(lcl.AsObject(getPtr(0)), browse, frame, TCefErrorCode(getVal(3)), GoStr(getVal(4)), GoStr(getVal(5)))

		case TOnLoadingStateChange:
			browse := AsCefBrowser(getPtr(1))
			fn.(TOnLoadingStateChange)(lcl.AsObject(getPtr(0)), browse, GoBool(getVal(2)), GoBool(getVal(3)), GoBool(getVal(4)))

		// ===========  ICefFocusHandler  ===========

		case TOnTakeFocus:
			browse := AsCefBrowser(getPtr(1))
			next := GoBool(getVal(2))
			fn.(TOnTakeFocus)(lcl.AsObject(getPtr(0)), browse, next)

		case TOnSetFocus:
			browse := AsCefBrowser(getPtr(1))
			source := TCefFocusSource(getVal(2))
			result := (*bool)(getPtr(3))
			fn.(TOnSetFocus)(lcl.AsObject(getPtr(0)), browse, source, result)

		//case TOnGotFocus = TOnBrowser

		// ===========  ICefContextMenuHandler  ===========

		case TOnBeforeContextMenu:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			params := AsCefContextMenuParams(getPtr(3))
			//KeyAccelerator.clear()
			model := AsCefMenuModel(getPtr(4)) //, CefMis: KeyAccelerator}
			fn.(TOnBeforeContextMenu)(lcl.AsObject(getPtr(0)), browse, frame, params, model)

		case TOnRunContextMenu:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			params := AsCefContextMenuParams(getPtr(3))
			model := AsCefMenuModel(getPtr(4))
			callback := AsCefRunContextMenuCallback(getPtr(5))
			result := (*bool)(getPtr(6))
			fn.(TOnRunContextMenu)(lcl.AsObject(getPtr(0)), browse, frame, params, model, callback, result)

		case TOnContextMenuCommand:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			params := AsCefContextMenuParams(getPtr(3))
			commandId := MenuId(getVal(4))
			eventFlags := uint32(getVal(5))
			result := (*bool)(getPtr(6))
			//if !KeyAccelerator.commandIdEventCallback(browse, commandId, params, eventFlags, result) {
			fn.(TOnContextMenuCommand)(lcl.AsObject(getPtr(0)), browse, frame, params, commandId, eventFlags, result)
			//}
			//params.Free()

		case TOnContextMenuDismissed:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			fn.(TOnContextMenuDismissed)(lcl.AsObject(getPtr(0)), browse, frame)

		case TOnRunQuickMenu:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			location := (*TCefPoint)(getPtr(3))
			size := (*TCefSize)(getPtr(4))
			editStateFlags := TCefQuickMenuEditStateFlags(getVal(5))
			callback := AsCefRunQuickMenuCallback(getPtr(6))
			result := (*bool)(getPtr(7))
			fn.(TOnRunQuickMenu)(lcl.AsObject(getPtr(0)), browse, frame, location, size, editStateFlags, callback, result)

		case TOnQuickMenuCommand:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			commandId := int32(getVal(3))
			eventFlags := TCefEventFlags(getVal(4))
			result := (*bool)(getPtr(5))
			fn.(TOnQuickMenuCommand)(lcl.AsObject(getPtr(0)), browse, frame, commandId, eventFlags, result)

		case TOnQuickMenuDismissed:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			fn.(TOnQuickMenuDismissed)(lcl.AsObject(getPtr(0)), browse, frame)

		// ===========  ICefKeyboardHandler  ===========

		case TOnPreKey:
			browse := AsCefBrowser(getPtr(1))
			event := *(*TCefKeyEvent)(getPtr(2))
			osEvent := *(*TCefEventHandle)(getPtr(3))
			isKeyboardShortcut := (*bool)(getPtr(4))
			result := (*bool)(getPtr(5))
			fn.(TOnPreKey)(lcl.AsObject(getPtr(0)), browse, event, osEvent, isKeyboardShortcut, result)

		case TOnKey:
			browse := AsCefBrowser(getPtr(1))
			keyEvent := *(*TCefKeyEvent)(getPtr(2))
			osEvent := *(*TCefEventHandle)(getPtr(3))
			result := (*bool)(getPtr(4))
			fn.(TOnKey)(lcl.AsObject(getPtr(0)), browse, keyEvent, osEvent, result)

		// ===========  ICefDisplayHandler  ===========

		case TOnAddressChange:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			fn.(TOnAddressChange)(lcl.AsObject(getPtr(0)), browse, frame, GoStr(getVal(3)))

		case TOnTitleChange:
			browse := AsCefBrowser(getPtr(1))
			fn.(TOnTitleChange)(lcl.AsObject(getPtr(0)), browse, GoStr(getVal(2)))

		case TOnFavIconUrlChange:
			browse := AsCefBrowser(getPtr(1))
			iconUrlsList := lcl.AsStrings(getVal(2))
			fn.(TOnFavIconUrlChange)(lcl.AsObject(getPtr(0)), browse, iconUrlsList)

		case TOnFullScreenModeChange:
			browse := AsCefBrowser(getPtr(1))
			fn.(TOnFullScreenModeChange)(lcl.AsObject(getPtr(0)), browse, GoBool(getVal(2)))

		case TOnTooltip:
			browse := AsCefBrowser(getPtr(1))
			textPtr := (*uintptr)(getPtr(2))
			var text = GoStr(*textPtr)
			result := (*bool)(getPtr(3))
			fn.(TOnTooltip)(lcl.AsObject(getPtr(0)), browse, &text, result)
			*textPtr = PascalStr(text)

		case TOnStatusMessage:
			browse := AsCefBrowser(getPtr(1))
			value := GoStr(getVal(2))
			fn.(TOnStatusMessage)(lcl.AsObject(getPtr(0)), browse, value)

		case TOnConsoleMessage:
			browse := AsCefBrowser(getPtr(1))
			level := TCefLogSeverity(getVal(2))
			message, source := GoStr(3), GoStr(4)
			line := int32(getVal(5))
			result := (*bool)(getPtr(6))
			fn.(TOnConsoleMessage)(lcl.AsObject(getPtr(0)), browse, level, message, source, line, result)

		case TOnAutoResize:
			browse := AsCefBrowser(getPtr(1))
			newSize := *(*TCefSize)(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnAutoResize)(lcl.AsObject(getPtr(0)), browse, newSize, result)

		case TOnLoadingProgressChange:
			browse := AsCefBrowser(getPtr(1))
			fn.(TOnLoadingProgressChange)(lcl.AsObject(getPtr(0)), browse, *(*float64)(getPtr(2)))

		case TOnCursorChange:
			browse := AsCefBrowser(getPtr(1))
			cursor := TCefCursorHandle(getVal(2))
			cursorType := TCefCursorType(getVal(3))
			customCursorInfo := *(*TCefCursorInfo)(getPtr(4)) // TODO 复杂类型？
			result := (*bool)(getPtr(5))
			fn.(TOnCursorChange)(lcl.AsObject(getPtr(0)), browse, cursor, cursorType, customCursorInfo, result)

		case TOnMediaAccessChange:
			browse := AsCefBrowser(getPtr(1))
			hasVideoAccess, hasAudioAccess := GoBool(getVal(2)), GoBool(getVal(3))
			fn.(TOnMediaAccessChange)(lcl.AsObject(getPtr(0)), browse, hasVideoAccess, hasAudioAccess)

		// ===========  ICefDownloadHandler  ===========

		case TOnCanDownload:
			browse := AsCefBrowser(getPtr(1))
			url, requestMethod := GoStr(getVal(2)), GoStr(getVal(3))
			result := (*bool)(getPtr(4))
			fn.(TOnCanDownload)(lcl.AsObject(getPtr(0)), browse, url, requestMethod, result)

		case TOnBeforeDownload:
			browse := AsCefBrowser(getPtr(1))
			downItem := AsCefDownloadItem(getPtr(2))
			suggestedName := GoStr(getVal(3))
			callback := AsCefBeforeDownloadCallback(getPtr(4))
			fn.(TOnBeforeDownload)(lcl.AsObject(getPtr(0)), browse, downItem, suggestedName, callback)

		case TOnDownloadUpdated:
			browse := AsCefBrowser(getPtr(1))
			downItem := AsCefDownloadItem(getPtr(2))
			callback := AsCefDownloadItemCallback(getPtr(3))
			fn.(TOnDownloadUpdated)(lcl.AsObject(getPtr(0)), browse, downItem, callback)

		// ===========  ICefJsDialogHandler  ===========

		case TOnJsdialog:
			browse := AsCefBrowser(getPtr(1))
			originUrl := GoStr(getVal(2))
			dialogType := TCefJsDialogType(getVal(3))
			messageText, defaultPromptText := GoStr(getVal(4)), GoStr(getVal(5))
			callback := AsCefJsDialogCallback(getPtr(6))
			suppressMessage := (*bool)(getPtr(7))
			result := (*bool)(getPtr(8))
			fn.(TOnJsdialog)(lcl.AsObject(getPtr(0)), browse, originUrl, dialogType, messageText, defaultPromptText, callback, suppressMessage, result)

		case TOnBeforeUnloadDialog:
			browse := AsCefBrowser(getPtr(1))
			messageText := GoStr(getVal(2))
			isReload := GoBool(getVal(3))
			callback := AsCefJsDialogCallback(getPtr(4))
			result := (*bool)(getPtr(5))
			fn.(TOnBeforeUnloadDialog)(lcl.AsObject(getPtr(0)), browse, messageText, isReload, callback, result)

		//case TOnResetDialogState = TOnBrowser
		//case TOnDialogClosed = TOnBrowser

		// ===========  ICefLifeSpanHandler  ===========

		case TOnBeforePopup:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			var (
				beforePopupPtr     = (*tBeforePopup)(getPtr(3))
				beforePopup        = *(beforePopupPtr.Convert())
				popupFeatures      = *(*TCefPopupFeatures)(getPtr(4))
				windowInfoPtr      = (*tCefWindowInfo)(getPtr(5))
				windowInfo         = windowInfoPtr.Convert()
				clientPtr          = (*uintptr)(getPtr(6))
				client             ICefClient
				settingsPtr        = (*tCefBrowserSettings)(getPtr(7))
				settings           = settingsPtr.Convert()
				extraInfoPtr       = (*uintptr)(getPtr(8)) // CEF49 = nil
				extraInfo          ICefDictionaryValue
				noJavascriptAccess = (*bool)(getPtr(9))
				result             = (*bool)(getPtr(10))
			)
			client, extraInfo, *noJavascriptAccess, *result = fn.(TOnBeforePopup)(lcl.AsObject(getPtr(0)), browse, frame, beforePopup, popupFeatures, windowInfo, settings)
			windowInfo.SetInstanceValue()
			if client != nil && client.Instance() != 0 {
				*clientPtr = client.Instance()
			}
			settings.SetInstanceValue()
			if extraInfo != nil && extraInfo.Instance() != 0 {
				*extraInfoPtr = extraInfo.Instance()
			}

		//case TOnAfterCreated = TOnBrowser
		//case TOnBeforeClose = TOnBrowser

		case TOnClose:
			browse := AsCefBrowser(getPtr(1))
			action := (*TCefCloseBrowserAction)(getPtr(2))
			fn.(TOnClose)(lcl.AsObject(getPtr(0)), browse, action)

		// ===========  ICefRequestHandler  ===========

		case TOnBeforeBrowse:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			userGesture := GoBool(getVal(4))
			isRedirect := GoBool(getVal(5))
			result := (*bool)(getPtr(6))
			fn.(TOnBeforeBrowse)(lcl.AsObject(getPtr(0)), browse, frame, request, userGesture, isRedirect, result)

		case TOnOpenUrlFromTab:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			targetUrl := GoStr(getVal(3))
			targetDisposition := TCefWindowOpenDisposition(getVal(4))
			userGesture := GoBool(getVal(5))
			result := (*bool)(getPtr(6))
			fn.(TOnOpenUrlFromTab)(lcl.AsObject(getPtr(0)), browse, frame, targetUrl, targetDisposition, userGesture, result)

		case TOnGetAuthCredentials:
			browse := AsCefBrowser(getPtr(1))
			originUrl := GoStr(getVal(2))
			isProxy := GoBool(getVal(3))
			host := GoStr(getVal(4))
			port := int32(getVal(5))
			realm := GoStr(getVal(6))
			scheme := GoStr(getVal(7))
			callback := AsCefAuthCallback(getPtr(8))
			result := (*bool)(getPtr(9))
			fn.(TOnGetAuthCredentials)(lcl.AsObject(getPtr(0)), browse, originUrl, isProxy, host, port, realm, scheme, callback, result)

		case TOnCertificateError:
			browser := AsCefBrowser(getPtr(1))
			certError := TCefErrorCode(getVal(2))
			requestUrl := GoStr(getVal(3))
			sslInfo := AsCefSslInfo(getPtr(4))
			callback := AsCefCallback(getPtr(5))
			result := (*bool)(getPtr(6))
			fn.(TOnCertificateError)(lcl.AsObject(getPtr(0)), browser, certError, requestUrl, sslInfo, callback, result)

		case TOnSelectClientCertificate:
			browser := AsCefBrowser(getPtr(1))
			isProxy := GoBool(getVal(2))
			host := GoStr(getVal(3))
			port := int32(getVal(4))
			certificatesCount := int(int32(getVal(5)))
			certificatesPtr := getVal(6)
			certificates := X509CertificateArrayRef.New(certificatesCount, certificatesPtr)
			callback := AsCefSelectClientCertificateCallback(getPtr(7))
			result := (*bool)(getPtr(8))
			fn.(TOnSelectClientCertificate)(lcl.AsObject(getPtr(0)), browser, isProxy, host, port, certificates, callback, result)

		//case TOnRenderViewReady = TOnBrowser

		case TOnRenderProcessTerminated:
			browse := AsCefBrowser(getPtr(1))
			fn.(TOnRenderProcessTerminated)(lcl.AsObject(getPtr(0)), browse, TCefTerminationStatus(getVal(2)))

		case TOnGetResourceRequestHandler:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			isNavigation, isDownload := GoBool(4), GoBool(5)
			requestInitiator := GoStr(getVal(6))
			disableDefaultHandling := (*bool)(getPtr(7))
			resourceRequestHandlerPtr := (*uintptr)(getPtr(8))
			var resourceRequestHandler ICefResourceRequestHandler
			resourceRequestHandler = fn.(TOnGetResourceRequestHandler)(lcl.AsObject(getPtr(0)), browse, frame, request, isNavigation, isDownload, requestInitiator,
				disableDefaultHandling)
			if resourceRequestHandler != nil && resourceRequestHandler.Instance() > 0 {
				*resourceRequestHandlerPtr = resourceRequestHandler.Instance()
			}

		//case TOnDocumentAvailableInMainFrame = TOnBrowser

		// ===========  ICefResourceRequestHandler  ===========

		case TOnBeforeResourceLoad:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			callback := AsCefCallback(getPtr(4))
			result := (*TCefReturnValue)(getPtr(5))
			fn.(TOnBeforeResourceLoad)(lcl.AsObject(getPtr(0)), browse, frame, request, callback, result)

		case TOnGetResourceHandler:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			resourceHandlerPtr := (*uintptr)(getPtr(4))
			var resourceHandler ICefResourceHandler
			resourceHandler = fn.(TOnGetResourceHandler)(lcl.AsObject(getPtr(0)), browse, frame, request)
			if resourceHandler != nil {
				*resourceHandlerPtr = resourceHandler.Instance()
			}

		case TOnResourceRedirect:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			response := AsCefResponse(getPtr(4))
			var newStrPtr = (*uintptr)(getPtr(5))
			var newStr = GoStr(*newStrPtr)
			fn.(TOnResourceRedirect)(lcl.AsObject(getPtr(0)), browse, frame, request, response, &newStr)
			*newStrPtr = PascalStr(newStr)

		case TOnResourceResponse:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			response := AsCefResponse(getPtr(4))
			result := (*bool)(getPtr(5))
			fn.(TOnResourceResponse)(lcl.AsObject(getPtr(0)), browse, frame, request, response, result)

		case TOnGetResourceResponseFilter:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			response := AsCefResponse(getPtr(4))
			responseFilterPtr := (*uintptr)(getPtr(5))
			var responseFilter ICefResponseFilter
			responseFilter = fn.(TOnGetResourceResponseFilter)(lcl.AsObject(getPtr(0)), browse, frame, request, response)
			if responseFilter != nil {
				*responseFilterPtr = responseFilter.Instance()
			}

		case TOnResourceLoadComplete:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			response := AsCefResponse(getPtr(4))
			status := TCefUrlRequestStatus(getVal(5))
			receivedContentLength := *(*int64)(getPtr(6))
			fn.(TOnResourceLoadComplete)(lcl.AsObject(getPtr(0)), browse, frame, request, response, status, receivedContentLength)

		case TOnProtocolExecution:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			allowOsExecution := (*bool)(getPtr(4))
			fn.(TOnProtocolExecution)(lcl.AsObject(getPtr(0)), browse, frame, request, allowOsExecution)

		// ===========  ICefCookieAccessFilter  ===========

		case TOnCanSendCookie:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			cookiePtr := (*tCefCookie)(getPtr(4))
			result := (*bool)(getPtr(5))
			cookie := *(cookiePtr.Convert())
			fn.(TOnCanSendCookie)(lcl.AsObject(getPtr(0)), browse, frame, request, cookie, result)

		case TOnCanSaveCookie:
			browse := AsCefBrowser(getPtr(1))
			frame := AsCefFrame(getPtr(2))
			request := AsCefRequest(getPtr(3))
			response := AsCefResponse(getPtr(4))
			cookiePtr := (*tCefCookie)(getPtr(5))
			result := (*bool)(getPtr(6))
			cookie := *(cookiePtr.Convert())
			fn.(TOnCanSaveCookie)(lcl.AsObject(getPtr(0)), browse, frame, request, response, cookie, result)

		// ===========  ICefDialogHandler  ===========

		case TOnFileDialog:
			browse := AsCefBrowser(getPtr(1))
			mode := TCefFileDialogMode(getVal(2))
			title := GoStr(getVal(3))
			defaultFilePath := GoStr(getVal(4))
			acceptFiltersList := lcl.AsStrings(getVal(5))
			callback := AsCefFileDialogCallback(getPtr(6))
			result := (*bool)(getPtr(7))
			fn.(TOnFileDialog)(lcl.AsObject(getPtr(0)), browse, mode, title, defaultFilePath, acceptFiltersList, callback, result)

		// ===========  ICefClient  ===========

		case TOnProcessMessageReceived:
			browse := AsCefBrowser(getPtr(0))
			frame := AsCefFrame(getPtr(1))
			processId := TCefProcessId(getVal(2))
			message := AsCefProcessMessage(getPtr(3))
			result := (*bool)(getPtr(4))
			fn.(TOnProcessMessageReceived)(browse, frame, processId, message, result)
			//message.Free() // TODO 需要Free?

		// ===========  ICefApp  ===========

		case TOnRegisterCustomSchemes:
			registrar := AsCefSchemeRegistrarRef(getVal(0))
			fn.(TOnRegisterCustomSchemes)(registrar)

		// ===========  ICefBrowserProcessHandler  ===========

		case TOnRegisterCustomPreferences:
			type_ := TCefPreferencesType(getVal(0))
			registrar := AsCefPreferenceRegistrarRef(getVal(1))
			fn.(TOnRegisterCustomPreferences)(type_, registrar)

		case TOnContextInitialized:
			fn.(TOnContextInitialized)()

		case TOnBeforeChildProcessLaunch:
			commandLine := AsCefCommandLine(getVal(0))
			fn.(TOnBeforeChildProcessLaunch)(commandLine)

		case TOnScheduleMessagePumpWork:
			delayMS := *(*int64)(getPtr(0))
			fn.(TOnScheduleMessagePumpWork)(delayMS)

		case TOnGetDefaultClient:
			client := AsCefClient(getVal(0))
			fn.(TOnGetDefaultClient)(client)

		// ===========  ICefResourceBundleHandler  ===========

		case TOnGetLocalizedString:
			stringId := int32(getVal(0))
			outValPtr := (*uintptr)(getPtr(1))
			result := (*bool)(getPtr(2))
			var outVal string
			outVal, *result = fn.(TOnGetLocalizedString)(stringId)
			*outValPtr = PascalStr(outVal)

		case TOnGetDataResource:
			resourceId := int32(getVal(0))
			outDataPtr := (*uintptr)(getPtr(1))
			dataSizePtr := (*uint32)(getPtr(2))
			result := (*bool)(getPtr(3))
			var outData []byte
			outData, *result = fn.(TOnGetDataResource)(resourceId)
			size := len(outData)
			if outData != nil && size > 0 {
				*outDataPtr = uintptr(unsafePointer(&outData[0]))
				*dataSizePtr = uint32(size)
			}

		case TOnGetDataResourceForScale:
			resourceId := int32(getVal(0))
			scaleFactor := TCefScaleFactor(getVal(1))
			outDataPtr := (*uintptr)(getPtr(2))
			dataSizePtr := (*uint32)(getPtr(3))
			result := (*bool)(getPtr(4))
			var outData []byte
			outData, *result = fn.(TOnGetDataResourceForScale)(resourceId, scaleFactor)
			size := len(outData)
			if outData != nil && size > 0 {
				*outDataPtr = uintptr(unsafePointer(&outData[0]))
				*dataSizePtr = uint32(size)
			}

		// ===========  ICefRenderProcessHandler  ===========

		case TOnWebKitInitialized:
			fn.(TOnWebKitInitialized)()

		case TOnBrowserCreated:
			browser := AsCefBrowser(getVal(0))
			extraInfo := AsCefDictionaryValue(getVal(1))
			fn.(TOnBrowserCreated)(browser, extraInfo)

		case TOnBrowserDestroyed:
			browser := AsCefBrowser(getVal(0))
			fn.(TOnBrowserDestroyed)(browser)

		case TOnContextCreated:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			context := AsCefv8Context(getVal(2))
			fn.(TOnContextCreated)(browser, frame, context)

		case TOnContextReleased:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			context := AsCefv8Context(getVal(2))
			fn.(TOnContextReleased)(browser, frame, context)

		case TOnUncaughtException:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			context := AsCefv8Context(getVal(2))
			exception := AsCefV8Exception(getVal(3))
			stackTrace := AsCefV8StackTrace(getVal(4))
			fn.(TOnUncaughtException)(browser, frame, context, exception, stackTrace)

		case TOnFocusedNodeChanged:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			node := AsCefDomNode(getVal(2))
			fn.(TOnFocusedNodeChanged)(browser, frame, node)

		// ===========  ICefLoadHandler  ===========

		case TOnRenderLoadingStateChange:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			isLoading, canGoBack, canGoForward := GoBool(getVal(2)), GoBool(getVal(3)), GoBool(getVal(4))
			fn.(TOnRenderLoadingStateChange)(browser, frame, isLoading, canGoBack, canGoForward)

		case TOnRenderLoadStart:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			transitionType := TCefTransitionType(getVal(2))
			fn.(TOnRenderLoadStart)(browser, frame, transitionType)

		case TOnRenderLoadEnd:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			httpStatusCode := int32(getVal(2))
			fn.(TOnRenderLoadEnd)(browser, frame, httpStatusCode)

		case TOnRenderLoadError:
			browser := AsCefBrowser(getVal(0))
			frame := AsCefFrame(getVal(1))
			errorCode := TCefErrorCode(getVal(2))
			errorText, failedUrl := GoStr(getVal(3)), GoStr(getVal(4))
			fn.(TOnRenderLoadError)(browser, frame, errorCode, errorText, failedUrl)

		// ===========  TCEFWindowComponent  ===========

		case TOnWindow:
			window_ := AsCefWindow(getVal(1))
			fn.(TOnWindow)(lcl.AsObject(getPtr(0)), window_)

		//case TOnWindowCreated = TOnWindow
		//case TOnWindowClosing = TOnWindow
		//case TOnWindowDestroyed = TOnWindow

		case TOnWindowActivationChanged:
			window_ := AsCefWindow(getVal(1))
			active := GoBool(getVal(2))
			fn.(TOnWindowActivationChanged)(lcl.AsObject(getPtr(0)), window_, active)

		case TOnWindowBoundsChanged:
			window_ := AsCefWindow(getVal(1))
			newBounds := *(*TCefRect)(getPtr(2))
			fn.(TOnWindowBoundsChanged)(lcl.AsObject(getPtr(0)), window_, newBounds)

		case TOnGetParentWindow:
			window_ := AsCefWindow(getVal(1))
			isMenu, canActivateMenu := (*bool)(getPtr(2)), (*bool)(getPtr(3))
			resultPtr := (*uintptr)(getPtr(4))
			var result ICefWindow
			result = fn.(TOnGetParentWindow)(lcl.AsObject(getPtr(0)), window_, isMenu, canActivateMenu)
			if result != nil && result.Instance() > 0 {
				*resultPtr = result.Instance()
			}

		case TOnWindowBool:
			window := AsCefWindow(getVal(1))
			result := (*bool)(getPtr(2))
			fn.(TOnWindowBool)(lcl.AsObject(getPtr(0)), window, result)

		case TOnGetInitialBounds:
			window := AsCefWindow(getVal(1))
			result := (*TCefRect)(getPtr(2))
			fn.(TOnGetInitialBounds)(lcl.AsObject(getPtr(0)), window, result)

		case TOnGetInitialShowState:
			window := AsCefWindow(getVal(1))
			result := (*TCefShowState)(getPtr(2))
			fn.(TOnGetInitialShowState)(lcl.AsObject(getPtr(0)), window, result)

		//case TOnIsFrameless = TOnWindowBool
		//case TOnWithStandardWindowButtons = TOnWindowBool

		case TOnGetTitlebarHeight:
			window := AsCefWindow(getVal(1))
			titleBarHeight := (*float32)(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnGetTitlebarHeight)(lcl.AsObject(getPtr(0)), window, titleBarHeight, result)

		//case TOnCanResize = TOnWindowBool
		//case TOnCanMaximize = TOnWindowBool
		//case TOnCanMinimize = TOnWindowBool
		//case TOnCanClose = TOnWindowBool

		case TOnAccelerator:
			window := AsCefWindow(getVal(1))
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(3))
			fn.(TOnAccelerator)(lcl.AsObject(getPtr(0)), window, commandId, result)

		case TOnWindowKeyEvent:
			window := AsCefWindow(getVal(1))
			event := *(*TCefKeyEvent)(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnWindowKeyEvent)(lcl.AsObject(getPtr(0)), window, event, result)

		case TOnWindowFullscreenTransition:
			window := AsCefWindow(getVal(1))
			isCompleted := GoBool(getVal(2))
			fn.(TOnWindowFullscreenTransition)(lcl.AsObject(getPtr(0)), window, isCompleted)

		// ===========  TCEFBrowserViewComponent  ===========

		case TOnBrowserCreatedBvc:
			browserView := AsCefBrowserView(getPtr(1))
			browser := AsCefBrowser(getPtr(2))
			fn.(TOnBrowserCreatedBvc)(lcl.AsObject(getPtr(0)), browserView, browser)

		case TOnBrowserDestroyedBvc:
			browserView := AsCefBrowserView(getPtr(1))
			browser := AsCefBrowser(getPtr(2))
			fn.(TOnBrowserDestroyedBvc)(lcl.AsObject(getPtr(0)), browserView, browser)

		case TOnGetDelegateForPopupBrowserViewBvc:
			browserView := AsCefBrowserView(getPtr(1))
			browserSettingsPtr := (*tCefBrowserSettings)(getPtr(2))
			client := AsCefClient(getPtr(3))
			isDevtools := GoBool(getVal(4))
			resultPtr := (*uintptr)(getPtr(5))
			browserSettings := *(browserSettingsPtr.Convert())
			var result ICefBrowserViewDelegate
			result = fn.(TOnGetDelegateForPopupBrowserViewBvc)(lcl.AsObject(getPtr(0)), browserView, browserSettings, client, isDevtools)
			if result != nil && result.Instance() > 0 {
				*resultPtr = result.Instance()
			}

		case TOnPopupBrowserViewCreatedBvc:
			browserView := AsCefBrowserView(getPtr(1))
			popupBrowserView := AsCefBrowserView(getPtr(2))
			isDevtools := GoBool(getVal(3))
			result := (*bool)(getPtr(4))
			fn.(TOnPopupBrowserViewCreatedBvc)(lcl.AsObject(getPtr(0)), browserView, popupBrowserView, isDevtools, result)

		case TOnGetChromeToolbarTypeBvc:
			fn.(TOnGetChromeToolbarTypeBvc)(lcl.AsObject(getPtr(0)), (*TCefChromeToolbarType)(getPtr(1)))

		case TOnUseFramelessWindowForPictureInPictureBvc:
			browserView := AsCefBrowserView(getPtr(1))
			result := (*bool)(getPtr(2))
			fn.(TOnUseFramelessWindowForPictureInPictureBvc)(lcl.AsObject(getPtr(0)), browserView, result)

		case TOnGestureCommandBvc:
			browserView := AsCefBrowserView(getPtr(1))
			gestureCommand := TCefGestureCommand(getVal(2))
			result := (*bool)(getPtr(3))
			fn.(TOnGestureCommandBvc)(lcl.AsObject(getPtr(0)), browserView, gestureCommand, result)

		// ===========  TAccessibilityHandler  ===========

		case TOnAccessibility:
			value := AsCefValue(getVal(1))
			fn.(TOnAccessibility)(lcl.AsObject(getPtr(0)), value)

		// ===========  TBufferPanel  ===========

		case TOnIMECommitText:
			fn.(TOnIMECommitText)(lcl.AsObject(getPtr(0)), GoStr(getVal(1)), *(*TCefRange)(getPtr(2)), int32(getVal(3)))

		case TOnIMESetComposition:
			text := GoStr(getVal(1))
			underlinesPtr := getVal(2)
			count := int(int32(getVal(3)))
			underlines := NewCefCompositionUnderlineArray(count, underlinesPtr)
			replacementRange := *(*TCefRange)(getPtr(4))
			selectionRange := *(*TCefRange)(getPtr(5))
			fn.(TOnIMESetComposition)(lcl.AsObject(getPtr(0)), text, underlines, replacementRange, selectionRange)

		case TOnHandledMessage:
			message := (*TMessage)(getPtr(1))
			lResult := (*LRESULT)(getPtr(2))
			fn.(TOnHandledMessage)(lcl.AsObject(getVal(0)), message, lResult, (*bool)(getPtr(3)))

		case TConstrainedResize:
			fn.(TConstrainedResize)(lcl.AsObject(getVal(0)), getI32Ptr(1), getI32Ptr(2), getI32Ptr(3), getI32Ptr(4))

		case TContextPopup:
			fn.(TContextPopup)(lcl.AsObject(getVal(0)), *getPointPtr(1), getBoolPtr(2))

		case TDragDrop:
			fn.(TDragDrop)(lcl.AsObject(getVal(0)), lcl.AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		case TEndDrag:
			fn.(TEndDrag)(lcl.AsObject(getVal(0)), lcl.AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		case TDragOver:
			fn.(TDragOver)(lcl.AsObject(getVal(0)), lcl.AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)),
				TDragState(getVal(4)), getBoolPtr(5))

		case TGetSiteInfo:
			fn.(TGetSiteInfo)(lcl.AsObject(getVal(0)), lcl.AsControl(getVal(1)), getRectPtr(2), *getPointPtr(3), getBoolPtr(4))

		case TMouse:
			fn.(TMouse)(lcl.AsObject(getVal(0)), types.TMouseButton(getVal(1)), types.TShiftState(getVal(2)), int32(getVal(3)),
				int32(getVal(4)))

		case TMouseMove:
			fn.(TMouseMove)(lcl.AsObject(getVal(0)), types.TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		case TMouseWheel:
			fn.(TMouseWheel)(lcl.AsObject(getVal(0)), types.TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)), int32(getVal(4)), getBoolPtr(5))

		case TStartDock:
			obj := lcl.AsDragDockObject(getPtrVal(1))
			fn.(TStartDock)(lcl.AsObject(getVal(0)), &obj)
			if obj != nil {
				setPtrVal(1, obj.Instance())
			}

		case TStartDrag:
			obj := lcl.AsDragObject(getVal(1))
			fn.(TStartDrag)(lcl.AsObject(getVal(0)), &obj)
			if obj != nil {
				*(*uintptr)(unsafe.Pointer(getVal(1))) = obj.Instance()
			}

		// ===========  ICefButtonDelegate  ===========

		case TOnButtonPressed:
			fn.(TOnButtonPressed)(lcl.AsObject(getVal(0)), AsCefButton(getVal(1)))

		case TOnButtonStateChanged:
			fn.(TOnButtonStateChanged)(lcl.AsObject(getVal(0)), AsCefButton(getVal(1)))

		// ===========  TCefResponseFilter  ===========

		case TOnInitFilter:
			result := (*bool)(getPtr(1))
			fn.(TOnInitFilter)(lcl.AsObject(getVal(0)), result)

		case TOnFilter:
			dataIn := getVal(1)
			dataInSize := uint32(getVal(2))
			dataInRead := (*uint32)(getPtr(3))
			dataOut := getVal(4)
			dataOutSize := uint32(getVal(5))
			dataOutWritten := (*uint32)(getPtr(6))
			status := (*TCefResponseFilterStatus)(getPtr(7))
			fn.(TOnFilter)(lcl.AsObject(getVal(0)), dataIn, dataInSize, dataInRead, dataOut, dataOutSize, dataOutWritten, status)

		// ===========  TDomVisitor  ===========

		case TOnDomVisitor:
			fn.(TOnDomVisitor)(AsCefDomDocument(getVal(0)))

		// ===========  ICefMenuButtonDelegate  ===========

		case TOnMenuButtonPressed:
			menuButton := AsCefMenuButton(getVal(1))
			screenPoint := *(*TCefPoint)(getPtr(2))
			buttonPressedLock := AsCefMenuButtonPressedLock(getVal(3))
			fn.(TOnMenuButtonPressed)(lcl.AsObject(getVal(0)), menuButton, screenPoint, buttonPressedLock)

		// ===========  TMenuModelDelegate  ===========

		case TOnExecuteCommand:
			menuModel := AsCefMenuModel(getVal(0))
			commandId := int32(getVal(1))
			eventFlags := TCefEventFlags(getVal(1))
			fn.(TOnExecuteCommand)(menuModel, commandId, eventFlags)

		case TOnMouseOutsideMenu:
			menuModel := AsCefMenuModel(getVal(0))
			screenPoint := *(*TCefPoint)(getPtr(1))
			fn.(TOnMouseOutsideMenu)(menuModel, screenPoint)

		case TOnUnhandledOpenSubmenu:
			menuModel := AsCefMenuModel(getVal(0))
			isRTL := GoBool(getVal(1))
			fn.(TOnUnhandledOpenSubmenu)(menuModel, isRTL)

		case TOnUnhandledCloseSubmenu:
			menuModel := AsCefMenuModel(getVal(0))
			isRTL := GoBool(getVal(1))
			fn.(TOnUnhandledCloseSubmenu)(menuModel, isRTL)

		case TOnMenuWillShow:
			menuModel := AsCefMenuModel(getVal(0))
			fn.(TOnMenuWillShow)(menuModel)

		case TOnMenuClosed:
			menuModel := AsCefMenuModel(getVal(0))
			fn.(TOnMenuClosed)(menuModel)

		case TOnFormatLabel:
			menuModel := AsCefMenuModel(getVal(0))
			labelPtr := (*uintptr)(getPtr(1))
			result := (*bool)(getPtr(2))
			var label = GoStr(*labelPtr)
			fn.(TOnFormatLabel)(menuModel, &label, result)
			*labelPtr = PascalStr(label)

		// ===========  TResourceHandler  ===========

		case TOnResourceHandlerOpen:
			request := AsCefRequest(getPtr(0))
			handleRequest := (*bool)(getPtr(1))
			callback := AsCefCallback(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnResourceHandlerOpen)(request, handleRequest, callback, result)

		case TOnResourceHandlerProcessRequest:
			request := AsCefRequest(getPtr(0))
			callback := AsCefCallback(getPtr(1))
			result := (*bool)(getPtr(2))
			fn.(TOnResourceHandlerProcessRequest)(request, callback, result)

		case TOnResourceHandlerGetResponseHeaders:
			response := AsCefResponse(getPtr(0))
			responseLength := (*int64)(getPtr(1))
			redirectUrlPtr := (*uintptr)(getPtr(2))
			var redirectUrl = GoStr(*redirectUrlPtr)
			fn.(TOnResourceHandlerGetResponseHeaders)(response, responseLength, &redirectUrl)
			*redirectUrlPtr = PascalStr(redirectUrl)

		case TOnResourceHandlerSkip:
			bytesToSkip := *(*int64)(getPtr(0))
			bytesSkipped := (*int64)(getPtr(1))
			callback := AsCefResourceSkipCallback(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnResourceHandlerSkip)(bytesToSkip, bytesSkipped, callback, result)

		case TOnResourceHandlerRead:
			dataOut := getVal(0)
			bytesToRead := int32(getVal(1))
			bytesRead := (*int32)(getPtr(2))
			callback := AsCefResourceReadCallback(getPtr(3))
			result := (*bool)(getPtr(4))
			fn.(TOnResourceHandlerRead)(dataOut, bytesToRead, bytesRead, callback, result)

		case TOnResourceHandlerReadResponse:
			dataOut := getVal(0)
			bytesToRead := int32(getVal(1))
			bytesRead := (*int32)(getPtr(2))
			callback := AsCefCallback(getPtr(3))
			result := (*bool)(getPtr(4))
			fn.(TOnResourceHandlerReadResponse)(dataOut, bytesToRead, bytesRead, callback, result)

		case TOnResourceHandlerCancel:
			fn.(TOnResourceHandlerCancel)()

		// ===========  TRunFileDialogCallback  ===========

		case TOnRunFileDialogDismissed:
			fn.(TOnRunFileDialogDismissed)(lcl.AsStrings(getVal(0)))

		// ===========  TSchemeHandlerFactory  ===========

		case TOnSchemeHandlerFactoryNew:
			browse := AsCefBrowser(getPtr(0))
			frame := AsCefFrame(getPtr(1))
			schemeName := GoStr(getVal(2))
			request := AsCefRequest(getVal(3))
			resultPtr := (*uintptr)(getPtr(4))
			var result ICefResourceHandler
			result = fn.(TOnSchemeHandlerFactoryNew)(browse, frame, schemeName, request)
			if result != nil && result.Instance() > 0 {
				*resultPtr = result.Instance()
			}

		// ===========  TCEFServerComponent  ===========

		case TOnServer:
			fn.(TOnServer)(lcl.AsObject(getPtr(0)), AsCEFServer(getVal(1)))

		//case TOnServerCreated = TOnServer
		//case TOnServerDestroyed = TOnServer

		case TOnServerInt32:
			fn.(TOnServerInt32)(lcl.AsObject(getPtr(0)), AsCEFServer(getVal(1)), int32(getVal(2)))

		//case TOnClientConnected = TOnServerInt32
		//case TOnClientDisconnected = TOnServerInt32

		case TOnHttpRequest:
			fn.(TOnHttpRequest)(lcl.AsObject(getPtr(0)), AsCEFServer(getVal(1)), int32(getVal(2)), GoStr(getVal(3)), AsCefRequest(getVal(4)))

		case TOnWebSocketRequest:
			fn.(TOnWebSocketRequest)(lcl.AsObject(getPtr(0)), AsCEFServer(getVal(1)), int32(getVal(2)), GoStr(getVal(3)), AsCefRequest(getVal(4)),
				AsCefCallback(getVal(5)))

		//case TOnWebSocketConnected = TOnServerInt32

		case TOnWebSocketMessage:
			fn.(TOnWebSocketMessage)(lcl.AsObject(getPtr(0)), AsCEFServer(getVal(1)), int32(getVal(2)), Pointer(getVal(3)), uint32(getVal(4)))

		// ===========  ICefTextfieldDelegate  ===========

		case TOnTextfieldKeyEvent:
			event := *(*TCefKeyEvent)(getPtr(2))
			result := (*bool)(getPtr(3))
			fn.(TOnTextfieldKeyEvent)(lcl.AsObject(getPtr(0)), AsCefTextfield(getVal(1)), event, result)

		case TOnAfterUserAction:
			fn.(TOnAfterUserAction)(lcl.AsObject(getPtr(0)), AsCefTextfield(getVal(1)))

		// ===========  TCEFUrlRequestClientComponent  ===========

		case TOnRequestCompleteRcc:
			fn.(TOnRequestCompleteRcc)(lcl.AsObject(getPtr(0)), AsCefUrlRequest(getVal(1)))

		case TOnUploadProgressRcc:
			current, total := *(*int64)(getPtr(2)), *(*int64)(getPtr(3))
			fn.(TOnUploadProgressRcc)(lcl.AsObject(getPtr(0)), AsCefUrlRequest(getVal(1)), current, total)

		case TOnDownloadProgressRcc:
			current, total := *(*int64)(getPtr(2)), *(*int64)(getPtr(3))
			fn.(TOnDownloadProgressRcc)(lcl.AsObject(getPtr(0)), AsCefUrlRequest(getVal(1)), current, total)

		case TOnDownloadDataRcc:
			data := getVal(2)
			dataLength := uint32(getVal(3))
			fn.(TOnDownloadDataRcc)(lcl.AsObject(getPtr(0)), AsCefUrlRequest(getVal(1)), data, dataLength)

		case TOnGetAuthCredentialsRcc:
			isProxy := GoBool(getVal(1))
			host := GoStr(getVal(2))
			port := int32(getVal(3))
			realm, scheme := GoStr(getVal(4)), GoStr(getVal(5))
			callback := AsCefAuthCallback(getVal(6))
			result := (*bool)(getPtr(7))
			fn.(TOnGetAuthCredentialsRcc)(lcl.AsObject(getPtr(0)), isProxy, host, port, realm, scheme, callback, result)

		//case TV8Accessor = TNotify

		// ===========  TV8Accessor  ===========

		case TOnV8AccessorGet:
			name := GoStr(getVal(0))
			object := AsCefv8Value(getPtr(1))
			retValPtr := (*uintptr)(getPtr(2))
			exceptionPtr := (*uintptr)(getPtr(3))
			result := (*bool)(getPtr(4))
			var (
				retVal    ICefv8Value
				exception string
			)
			retVal, exception, *result = fn.(TOnV8AccessorGet)(name, object)
			if retVal != nil && retVal.Instance() > 0 {
				*retValPtr = retVal.Instance()
			}
			*exceptionPtr = PascalStr(exception)

		case TOnV8AccessorSet:
			name := GoStr(getVal(0))
			object := AsCefv8Value(getPtr(1))
			value := AsCefv8Value(getPtr(2))
			exceptionPtr := (*uintptr)(getPtr(3))
			result := (*bool)(getPtr(4))
			var exception string
			exception, *result = fn.(TOnV8AccessorSet)(name, object, value)
			*exceptionPtr = PascalStr(exception)

		// ===========  TV8ArrayBufferReleaseCallback  ===========

		case TOnV8ArrayBufferReleaseBuffer:
			fn.(TOnV8ArrayBufferReleaseBuffer)(getVal(0))

		// ===========  TV8Handler  ===========

		case TOnV8HandlerExecute:
			name := GoStr(getVal(0))
			object := AsCefv8Value(getPtr(1))
			argumentsPtr, argumentsCount := getVal(2), int(int32(getVal(3)))
			retValPtr := (*uintptr)(getPtr(4))
			exceptionPtr := (*uintptr)(getPtr(5))
			result := (*bool)(getPtr(6))
			arguments := V8ValueArrayRef.New(argumentsCount, argumentsPtr)
			var (
				retVal    ICefv8Value
				exception string
			)
			retVal, exception, *result = fn.(TOnV8HandlerExecute)(name, object, arguments)
			if retVal != nil && retVal.Instance() > 0 {
				*retValPtr = retVal.Instance()
			}
			*exceptionPtr = PascalStr(exception)

		// ===========  TV8Interceptor  ===========

		case TOnV8InterceptorGetByName:
			name := GoStr(getVal(0))
			object := AsCefv8Value(getPtr(1))
			retValPtr := (*uintptr)(getPtr(2))
			exceptionPtr := (*uintptr)(getPtr(3))
			result := (*bool)(getPtr(4))
			var (
				retVal    ICefv8Value
				exception string
			)
			retVal, exception, *result = fn.(TOnV8InterceptorGetByName)(name, object)
			if retVal != nil {
				*retValPtr = retVal.Instance()
			}
			*exceptionPtr = PascalStr(exception)

		case TOnV8InterceptorGetByIndex:
			index := int32(getVal(0))
			object := AsCefv8Value(getPtr(1))
			retValPtr := (*uintptr)(getPtr(2))
			exceptionPtr := (*uintptr)(getPtr(3))
			result := (*bool)(getPtr(4))
			var (
				retVal    ICefv8Value
				exception string
			)
			retVal, exception, *result = fn.(TOnV8InterceptorGetByIndex)(index, object)
			if retVal != nil {
				*retValPtr = retVal.Instance()
			}
			*exceptionPtr = PascalStr(exception)

		case TOnV8InterceptorSetByName:
			name := GoStr(getVal(0))
			object := AsCefv8Value(getPtr(1))
			value := AsCefv8Value(getPtr(2))
			exceptionPtr := (*uintptr)(getPtr(3))
			result := (*bool)(getPtr(4))
			var exception string
			exception, *result = fn.(TOnV8InterceptorSetByName)(name, object, value)
			*exceptionPtr = PascalStr(exception)

		case TOnV8InterceptorSetByIndex:
			index := int32(getVal(0))
			object := AsCefv8Value(getPtr(1))
			value := AsCefv8Value(getPtr(2))
			exceptionPtr := (*uintptr)(getPtr(3))
			result := (*bool)(getPtr(4))
			var exception string
			exception, *result = fn.(TOnV8InterceptorSetByIndex)(index, object, value)
			*exceptionPtr = PascalStr(exception)

		// ===========  ICefViewDelegate  ===========

		case TOnViewResultSize:
			view := AsCefView(getVal(1))
			result := (*TCefSize)(getPtr(2))
			fn.(TOnViewResultSize)(lcl.AsObject(getVal(0)), view, result)

		//case TOnGetPreferredSize = TOnViewResultSize
		//case TOnGetMinimumSize = TOnViewResultSize
		//case TOnGetMaximumSize = TOnViewResultSize

		case TOnGetHeightForWidth:
			view := AsCefView(getVal(1))
			width := int32(getVal(2))
			result := (*int32)(getPtr(3))
			fn.(TOnGetHeightForWidth)(lcl.AsObject(getVal(0)), view, width, result)

		case TOnParentViewChanged:
			view := AsCefView(getVal(1))
			added := GoBool(getVal(2))
			parent := AsCefView(getVal(3))
			fn.(TOnParentViewChanged)(lcl.AsObject(getVal(0)), view, added, parent)

		case TOnChildViewChanged:
			view := AsCefView(getVal(1))
			added := GoBool(getVal(2))
			child := AsCefView(getVal(3))
			fn.(TOnChildViewChanged)(lcl.AsObject(getVal(0)), view, added, child)

		case TOnWindowChanged:
			view := AsCefView(getVal(1))
			added := GoBool(getVal(2))
			fn.(TOnWindowChanged)(lcl.AsObject(getVal(0)), view, added)

		case TOnLayoutChanged:
			view := AsCefView(getVal(1))
			newBounds := *(*TCefRect)(getPtr(2))
			fn.(TOnLayoutChanged)(lcl.AsObject(getVal(0)), view, newBounds)

		case TOnView:
			view := AsCefView(getVal(1))
			fn.(TOnView)(lcl.AsObject(getVal(0)), view)

		case TTaskExecute:
			fn.(TTaskExecute)()

		//case TOnFocus = TOnView
		//case TOnBlur = TOnView

		default:
		}
	}
	return 0
}

// GlobalInit
//
//	初始化
func GlobalInit(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	inits.Init(libs, resources)
	SetCEFEventCallback(eventCallback)
	SetCEFRemoveEventCallback(removeEventCallback)
}
