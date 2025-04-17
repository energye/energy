//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Chromium 事件回调

package cef

import (
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

func init() {
	var resourceEventGet = func(fn interface{}, getVal func(idx int) uintptr, resp bool) (sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) {
		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		senderPtr := getPtr(0)
		browser = &ICefBrowser{instance: getPtr(1)}
		frame = &ICefFrame{instance: getPtr(2)}
		request = &ICefRequest{instance: getPtr(3)}
		if resp {
			response = &ICefResponse{instance: getPtr(4)}
		}
		return lcl.AsObject(senderPtr), browser, frame, request, response
	}
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case chromiumEventOnAcceleratedPaint:
			browser := &ICefBrowser{instance: getPtr(1)}
			kind := consts.TCefPaintElementType(getVal(2))
			dirtyRectsCount := *(*uint32)(getPtr(3))
			dirtyRectsPtr := getVal(4)
			info := *(*TCefAcceleratedPaintInfo)(getPtr(5))
			fn.(chromiumEventOnAcceleratedPaint)(lcl.AsObject(getPtr(0)), browser, kind, NewTCefRectArray(dirtyRectsPtr, dirtyRectsCount), info)
		case chromiumEventOnAllConnectionsClosed:
			fn.(chromiumEventOnAllConnectionsClosed)(lcl.AsObject(getPtr(0)))
		case chromiumEventOnAudioStreamError:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnAudioStreamError)(lcl.AsObject(getPtr(0)), browse, api.GoStr(getVal(2)))
		case chromiumEventOnAudioStreamPacket:
			browse := &ICefBrowser{instance: getPtr(1)}
			data := (*uintptr)(getPtr(2))
			frames := int32(getVal(3))
			pts := *(*int64)(getPtr(4))
			fn.(chromiumEventOnAudioStreamPacket)(lcl.AsObject(getPtr(0)), browse, data, frames, pts)
		case chromiumEventOnAudioStreamStarted:
			browse := &ICefBrowser{instance: getPtr(1)}
			params := (*TCefAudioParameters)(getPtr(2))
			fn.(chromiumEventOnAudioStreamStarted)(lcl.AsObject(getPtr(0)), browse, params, int32(getVal(2)))
		case chromiumEventOnAudioStreamStopped:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnAudioStreamStopped)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnAutoResize:
			browse := &ICefBrowser{instance: getPtr(1)}
			newSize := (*TCefSize)(getPtr(2))
			result := (*bool)(getPtr(3))
			*result = fn.(chromiumEventOnAutoResize)(lcl.AsObject(getPtr(0)), browse, newSize)
		case chromiumEventOnBeforeUnloadDialog:
			browse := &ICefBrowser{instance: getPtr(1)}
			messageText := api.GoStr(getVal(2))
			isReload := api.GoBool(getVal(3))
			callback := &ICefJsDialogCallback{instance: getPtr(4)}
			result := (*bool)(getPtr(5))
			*result = fn.(chromiumEventOnBeforeUnloadDialog)(lcl.AsObject(getPtr(0)), browse, messageText, isReload, callback)
		case chromiumEventOnCanDownload:
			browse := &ICefBrowser{instance: getPtr(1)}
			url, requestMethod := api.GoStr(getVal(2)), api.GoStr(getVal(3))
			result := (*bool)(getPtr(4))
			*result = fn.(chromiumEventOnCanDownload)(lcl.AsObject(getPtr(0)), browse, url, requestMethod)
		case chromiumEventOnCanSaveCookie:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			response := &ICefResponse{instance: getPtr(4)}
			cookiePtr := (*tCefCookiePtr)(getInstance(getVal(5)))
			cookie := cookiePtr.convert()
			result := (*bool)(getInstance(getVal(6)))
			*result = fn.(chromiumEventOnCanSaveCookie)(lcl.AsObject(getPtr(0)), browse, frame, request, response, cookie)
		case chromiumEventOnCanSendCookie:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			cookiePtr := (*tCefCookiePtr)(getInstance(getVal(4)))
			cookie := cookiePtr.convert()
			result := (*bool)(getInstance(getVal(5)))
			*result = fn.(chromiumEventOnCanSendCookie)(lcl.AsObject(getPtr(0)), browse, frame, request, cookie)
		case chromiumEventOnCertificateError:
			browser := &ICefBrowser{instance: getPtr(1)}
			certError := consts.TCefErrorCode(getVal(2))
			requestUrl := api.GoStr(getVal(3))
			sslInfo := &ICefSslInfo{instance: getPtr(4)}
			callback := &ICefCallback{instance: getPtr(5)}
			resultPtr := (*bool)(getPtr(6))
			*resultPtr = fn.(chromiumEventOnCertificateError)(lcl.AsObject(getPtr(0)), browser, certError, requestUrl, sslInfo, callback)
		case chromiumEventOnCertificateExceptionsCleared:
			fn.(chromiumEventOnCertificateExceptionsCleared)(lcl.AsObject(getPtr(0)))
		case chromiumEventOnChromeCommand:
			browse := &ICefBrowser{instance: getPtr(1)}
			params := (int32)(getVal(2))
			disposition := consts.TCefWindowOpenDisposition(getVal(3))
			result := (*bool)(getPtr(4))
			*result = fn.(chromiumEventOnChromeCommand)(lcl.AsObject(getPtr(0)), browse, params, disposition)
		case chromiumEventOnConsoleMessage:
			browse := &ICefBrowser{instance: getPtr(1)}
			level := consts.TCefLogSeverity(getVal(2))
			message, source := api.GoStr(getVal(3)), api.GoStr(getVal(4))
			line := int32(getVal(5))
			result := (*bool)(getPtr(6))
			*result = fn.(chromiumEventOnConsoleMessage)(lcl.AsObject(getPtr(0)), browse, level, message, source, line)
		case chromiumEventOnCursorChange:
			browse := &ICefBrowser{instance: getPtr(1)}
			cursor := consts.TCefCursorHandle(getVal(2))
			cursorType := consts.TCefCursorType(getVal(3))
			customCursorInfo := (*TCefCursorInfo)(getPtr(4))
			fn.(chromiumEventOnCursorChange)(lcl.AsObject(getPtr(0)), browse, cursor, cursorType, customCursorInfo)
		case chromiumEventOnDevToolsAgentAttached:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnDevToolsAgentAttached)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnDevToolsAgentDetached:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnDevToolsAgentDetached)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnDevTools:
			browse := &ICefBrowser{instance: getPtr(1)}
			method := api.GoStr(getVal(2))
			params := &ICefValue{instance: getPtr(3)}
			fn.(chromiumEventOnDevTools)(lcl.AsObject(getPtr(0)), browse, method, params)
		case chromiumEventOnDevToolsMessage:
			browse := &ICefBrowser{instance: getPtr(1)}
			message := &ICefValue{instance: getPtr(2)}
			fn.(chromiumEventOnDevToolsMessage)(lcl.AsObject(getPtr(0)), browse, message)
		case chromiumEventOnDevToolsMethodRawResult:
			browse := &ICefBrowser{instance: getPtr(1)}
			messageId := int32(getVal(2))
			success := api.GoBool(getVal(3))
			result := getVal(4)
			resultSize := *(*uint32)(getPtr(5))
			fn.(chromiumEventOnDevToolsMethodRawResult)(lcl.AsObject(getPtr(0)), browse, messageId, success, result, resultSize)
		case chromiumEventOnDevToolsMethodResult:
			browse := &ICefBrowser{instance: getPtr(1)}
			messageId := int32(getVal(2))
			success := api.GoBool(getVal(3))
			result := &ICefValue{instance: getPtr(1)}
			fn.(chromiumEventOnDevToolsMethodResult)(lcl.AsObject(getPtr(0)), browse, messageId, success, result)
		case chromiumEventOnDevToolsRaw:
			browse := &ICefBrowser{instance: getPtr(1)}
			method := api.GoStr(getVal(2))
			params := getVal(3)
			paramsSize := *(*uint32)(getPtr(4))
			fn.(chromiumEventOnDevToolsRaw)(lcl.AsObject(getPtr(0)), browse, method, params, paramsSize)
		case chromiumEventOnDevToolsRawMessage:
			browse := &ICefBrowser{instance: getPtr(1)}
			message := getVal(2)
			messageSize := *(*uint32)(getPtr(3))
			handledPtr := (*bool)(getPtr(4))
			*handledPtr = fn.(chromiumEventOnDevToolsRawMessage)(lcl.AsObject(getPtr(0)), browse, message, messageSize)
		case chromiumEventOnDialogClosed:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnDialogClosed)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnDismissPermissionPrompt:
			browse := &ICefBrowser{instance: getPtr(1)}
			promptId := *(*uint64)(getPtr(2))
			result := consts.TCefPermissionRequestResult(getVal(3))
			fn.(chromiumEventOnDismissPermissionPrompt)(lcl.AsObject(getPtr(0)), browse, promptId, result)
		case chromiumEventOnDocumentAvailableInMainFrame:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnDocumentAvailableInMainFrame)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnDownloadImageFinished:
			imageUrl := api.GoStr(getVal(1))
			httpStatusCode := int32(getVal(2))
			image := &ICefImage{instance: getPtr(3)}
			fn.(chromiumEventOnDownloadImageFinished)(lcl.AsObject(getPtr(0)), imageUrl, httpStatusCode, image)
		case chromiumEventOnExecuteTaskOnCefThread:
			taskID := uint32(getVal(1))
			fn.(chromiumEventOnExecuteTaskOnCefThread)(lcl.AsObject(getPtr(0)), taskID)
		case chromiumEventOnPrintStart:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnPrintStart)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnPrintSettings:
			browse := &ICefBrowser{instance: getPtr(1)}
			settings := &ICefPrintSettings{instance: getPtr(2)}
			getDefaults := api.GoBool(getVal(3))
			fn.(chromiumEventOnPrintSettings)(lcl.AsObject(getPtr(0)), browse, settings, getDefaults)
		case chromiumEventOnPrintDialog:
			browse := &ICefBrowser{instance: getPtr(1)}
			hasSelection := api.GoBool(getVal(2))
			callback := &ICefPrintDialogCallback{instance: getPtr(3)}
			result := (*bool)(getPtr(4))
			*result = fn.(chromiumEventOnPrintDialog)(lcl.AsObject(getPtr(0)), browse, hasSelection, callback)
		case chromiumEventOnPrintJob:
			browse := &ICefBrowser{instance: getPtr(1)}
			documentName, PDFFilePath := api.GoStr(getVal(2)), api.GoStr(getVal(3))
			callback := &ICefPrintJobCallback{instance: getPtr(4)}
			result := (*bool)(getPtr(5))
			*result = fn.(chromiumEventOnPrintJob)(lcl.AsObject(getPtr(0)), browse, documentName, PDFFilePath, callback)
		case chromiumEventOnPrintReset:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnPrintReset)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnGetPDFPaperSize:
			browse := &ICefBrowser{instance: getPtr(1)}
			deviceUnitsPerInch := int32(getVal(2))
			resultSize := (*TCefSize)(getPtr(3))
			fn.(chromiumEventOnGetPDFPaperSize)(lcl.AsObject(getPtr(0)), browse, deviceUnitsPerInch, resultSize)
		case chromiumEventOnFavIconUrlChange:
			browse := &ICefBrowser{instance: getPtr(1)}
			iconUrlsList := lcl.AsStrings(getVal(2))
			var iconUrls []string
			if iconUrlsList.IsValid() {
				count := int(iconUrlsList.Count())
				iconUrls = make([]string, count, count)
				for i := 0; i < count; i++ {
					iconUrls[i] = iconUrlsList.Strings(int32(i))
				}
				iconUrlsList.Free()
			}
			fn.(chromiumEventOnFavIconUrlChange)(lcl.AsObject(getPtr(0)), browse, iconUrls)
		case chromiumEventOnFileDialog:
			browse := &ICefBrowser{instance: getPtr(1)}
			mode := consts.FileDialogMode(getVal(2))
			title := api.GoStr(getVal(3))
			defaultFilePath := api.GoStr(getVal(4))
			acceptFiltersList := lcl.AsStrings(getVal(5))
			acceptExtensions := lcl.AsStrings(getVal(6))
			acceptDescriptions := lcl.AsStrings(getVal(7))
			callback := &ICefFileDialogCallback{instance: getPtr(8)}
			result := (*bool)(getPtr(9))
			*result = fn.(chromiumEventOnFileDialog)(lcl.AsObject(getPtr(0)), browse, mode, title, defaultFilePath, acceptFiltersList, acceptExtensions, acceptDescriptions, callback)
		case chromiumEventOnGetAccessibilityHandler:
			accessibilityHandler := &ICefAccessibilityHandler{instance: getPtr(1)}
			fn.(chromiumEventOnGetAccessibilityHandler)(lcl.AsObject(getPtr(0)), accessibilityHandler)
		case chromiumEventOnGetAudioParameters:
			browse := &ICefBrowser{instance: getPtr(1)}
			params := (*TCefAudioParameters)(getPtr(2))
			result := (*bool)(getPtr(3))
			*result = fn.(chromiumEventOnGetAudioParameters)(lcl.AsObject(getPtr(0)), browse, params)
		case chromiumEventOnGetResourceHandler:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			resourceHandlerPtr := (*uintptr)(getPtr(4))
			resourceHandler := fn.(chromiumEventOnGetResourceHandler)(lcl.AsObject(getPtr(0)), browse, frame, request)
			if resourceHandler != nil && resourceHandler.IsValid() {
				*resourceHandlerPtr = resourceHandler.Instance()
			} else {
				*resourceHandlerPtr = 0
			}
		case chromiumEventOnGetResourceRequestHandlerReqCtxHdlr:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			isNavigation, isDownload := api.GoBool(getVal(4)), api.GoBool(getVal(5))
			requestInitiator := api.GoStr(getVal(6))
			disableDefaultHandlingPtr := (*bool)(getPtr(7))
			resourceRequestHandlerPtr := (*uintptr)(getPtr(8))
			disableDefaultHandling, resourceRequestHandler := fn.(chromiumEventOnGetResourceRequestHandlerReqCtxHdlr)(lcl.AsObject(getPtr(0)), browse, frame, request, isNavigation, isDownload, requestInitiator)
			*disableDefaultHandlingPtr = disableDefaultHandling
			if resourceRequestHandler != nil && resourceRequestHandler.IsValid() {
				*resourceRequestHandlerPtr = resourceRequestHandler.Instance()
			}
		case chromiumEventOnGetResourceRequestHandlerReqHdlr:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			isNavigation, isDownload := api.GoBool(getVal(4)), api.GoBool(getVal(5))
			requestInitiator := api.GoStr(getVal(6))
			disableDefaultHandlingPtr := (*bool)(getPtr(7))
			resourceRequestHandlerPtr := (*uintptr)(getPtr(8))
			disableDefaultHandling, resourceRequestHandler := fn.(chromiumEventOnGetResourceRequestHandlerReqHdlr)(lcl.AsObject(getPtr(0)), browse, frame, request, isNavigation, isDownload, requestInitiator)
			*disableDefaultHandlingPtr = disableDefaultHandling
			if resourceRequestHandler != nil && resourceRequestHandler.IsValid() {
				*resourceRequestHandlerPtr = resourceRequestHandler.Instance()
			}
		case chromiumEventOnGetResourceResponseFilter:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			response := &ICefResponse{instance: getPtr(4)}
			responseFilterPtr := (*uintptr)(getPtr(5))
			responseFilter := fn.(chromiumEventOnGetResourceResponseFilter)(lcl.AsObject(getPtr(0)), browse, frame, request, response)
			if responseFilter != nil && responseFilter.IsValid() {
				*responseFilterPtr = responseFilter.Instance()
			} else {
				*responseFilterPtr = 0
			}
		case chromiumEventOnGetRootScreenRect:
			browser := &ICefBrowser{instance: getPtr(1)}
			rectPtr := (*TCefRect)(getPtr(2))
			resultPtr := (*bool)(getPtr(3))
			rect, result := fn.(chromiumEventOnGetRootScreenRect)(lcl.AsObject(getPtr(0)), browser)
			if rect != nil {
				*rectPtr = *rect
			}
			*resultPtr = result
		case chromiumEventOnGetScreenInfo:
			browser := &ICefBrowser{instance: getPtr(1)}
			screenInfoPtr := (*TCefScreenInfo)(getPtr(2))
			resultPtr := (*bool)(getPtr(3))
			screenInfo, result := fn.(chromiumEventOnGetScreenInfo)(lcl.AsObject(getPtr(0)), browser)
			if screenInfo != nil {
				*screenInfoPtr = *screenInfo
			}
			*resultPtr = result
		case chromiumEventOnGetScreenPoint:
			browser := &ICefBrowser{instance: getPtr(1)}
			viewX, viewY := int32(getVal(2)), int32(getVal(3))
			screenXPtr, screenYPtr, resultPtr := (*int32)(getPtr(4)), (*int32)(getPtr(5)), (*bool)(getPtr(6))
			screenX, screenY, result := fn.(chromiumEventOnGetScreenPoint)(lcl.AsObject(getPtr(0)), browser, viewX, viewY)
			*screenXPtr, *screenYPtr, *resultPtr = screenX, screenY, result
		case chromiumEventOnGetTouchHandleSize:
			browser := &ICefBrowser{instance: getPtr(1)}
			orientation := consts.TCefHorizontalAlignment(getVal(2))
			sizePtr := (*TCefSize)(getPtr(3))
			size := fn.(chromiumEventOnGetTouchHandleSize)(lcl.AsObject(getPtr(0)), browser, orientation)
			if size != nil {
				*sizePtr = *size
			}
		case chromiumEventOnGetViewRect:
			browser := &ICefBrowser{instance: getPtr(1)}
			rectPtr := (*TCefRect)(getPtr(2))
			rect := fn.(chromiumEventOnGetViewRect)(lcl.AsObject(getPtr(0)), browser)
			if rect != nil {
				*rectPtr = *rect
			}
		case chromiumEventOnGotFocus:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnGotFocus)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnHttpAuthCredentialsCleared:
			fn.(chromiumEventOnHttpAuthCredentialsCleared)(lcl.AsObject(getPtr(0)))
		case chromiumEventOnIMECompositionRangeChanged:
			browser := &ICefBrowser{instance: getPtr(1)}
			rng := *(*TCefRange)(getPtr(2))
			characterBoundsCount := *(*uint32)(getPtr(3))
			characterBounds := *(*TCefRect)(getPtr(4))
			fn.(chromiumEventOnIMECompositionRangeChanged)(lcl.AsObject(getPtr(0)), browser, rng, characterBoundsCount, characterBounds)
		case chromiumEventOnJsDialog:
			browse := &ICefBrowser{instance: getPtr(1)}
			originUrl := api.GoStr(getVal(2))
			dialogType := consts.TCefJsDialogType(getVal(3))
			messageText, defaultPromptText := api.GoStr(getVal(4)), api.GoStr(getVal(5))
			callback := &ICefJsDialogCallback{instance: getPtr(6)}
			suppressMessage := (*bool)(getPtr(7))
			result := (*bool)(getPtr(8))
			*suppressMessage, *result = fn.(chromiumEventOnJsDialog)(lcl.AsObject(getPtr(0)), browse, originUrl, dialogType, messageText, defaultPromptText, callback)
		case chromiumEventOnMediaAccessChange:
			browse := &ICefBrowser{instance: getPtr(1)}
			hasVideoAccess, hasAudioAccess := api.GoBool(getVal(2)), api.GoBool(getVal(3))
			fn.(chromiumEventOnMediaAccessChange)(lcl.AsObject(getPtr(0)), browse, hasVideoAccess, hasAudioAccess)
		case chromiumEventOnMediaRouteCreateFinished:
			result := consts.TCefMediaRouterCreateResult(getVal(1))
			error := api.GoStr(getVal(2))
			route := &ICefMediaRoute{instance: getPtr(3)}
			fn.(chromiumEventOnMediaRouteCreateFinished)(lcl.AsObject(getPtr(0)), result, error, route)
		case chromiumEventOnMediaSinkDeviceInfo:
			ipAddress := api.GoStr(getVal(1))
			port := int32(getVal(2))
			modelName := api.GoStr(getVal(3))
			fn.(chromiumEventOnMediaSinkDeviceInfo)(lcl.AsObject(getPtr(0)), ipAddress, port, modelName)
		case chromiumEventOnNavigationVisitorResultAvailable:
			entry := &ICefNavigationEntry{instance: getPtr(1)}
			current := api.GoBool(getVal(2))
			index, total := int32(getVal(3)), int32(getVal(4))
			fn.(chromiumEventOnNavigationVisitorResultAvailable)(lcl.AsObject(getPtr(0)), entry, current, index, total)
		case chromiumEventOnPaint:
			browser := &ICefBrowser{instance: getPtr(1)}
			kind := consts.TCefPaintElementType(getVal(2))
			dirtyRectsCount := *(*uint32)(getPtr(3))
			dirtyRectsPtr := getVal(4)
			buffer := getVal(5)
			width, height := int32(getVal(6)), int32(getVal(7))
			fn.(chromiumEventOnPaint)(lcl.AsObject(getPtr(0)), browser, kind, NewTCefRectArray(dirtyRectsPtr, dirtyRectsCount), buffer, width, height)
		case chromiumEventOnPdfPrintFinished:
			fn.(chromiumEventOnPdfPrintFinished)(lcl.AsObject(getPtr(0)), api.GoBool(getVal(1)))
		case chromiumEventOnPopupShow:
			browser := &ICefBrowser{instance: getPtr(1)}
			show := api.GoBool(getVal(2))
			fn.(chromiumEventOnPopupShow)(lcl.AsObject(getPtr(0)), browser, show)
		case chromiumEventOnPopupSize:
			browser := &ICefBrowser{instance: getPtr(1)}
			rect := (*TCefRect)(getPtr(2))
			fn.(chromiumEventOnPopupSize)(lcl.AsObject(getPtr(0)), browser, rect)
		case chromiumEventOnPrefsAvailable:
			fn.(chromiumEventOnPrefsAvailable)(lcl.AsObject(getPtr(0)), api.GoBool(getVal(1)))
		case chromiumEventOnPrefsUpdated:
			fn.(chromiumEventOnPrefsUpdated)(lcl.AsObject(getPtr(0)))
		case chromiumEventOnPreKey:
			browse := &ICefBrowser{instance: getPtr(1)}
			event := (*TCefKeyEvent)(getPtr(2))
			osEvent := consts.EventHandle(getVal(3))
			isKeyboardShortcut := (*bool)(getPtr(4))
			result := (*bool)(getPtr(5))
			*isKeyboardShortcut, *result = fn.(chromiumEventOnPreKey)(lcl.AsObject(getPtr(0)), browse, event, osEvent)
		case chromiumEventOnProtocolExecution:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			allowOsExecution := (*bool)(getPtr(4))
			*allowOsExecution = fn.(chromiumEventOnProtocolExecution)(lcl.AsObject(getPtr(0)), browse, frame, request)
		case chromiumEventOnQuickMenuCommand:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			commandId := int32(getVal(3))
			eventFlags := consts.TCefEventFlags(getVal(4))
			result := (*bool)(getPtr(5))
			*result = fn.(chromiumEventOnQuickMenuCommand)(lcl.AsObject(getPtr(0)), browse, frame, commandId, eventFlags)
		case chromiumEventOnQuickMenuDismissed:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnQuickMenuDismissed)(lcl.AsObject(getPtr(0)), browse, frame)
		case chromiumEventOnRenderViewReady:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnRenderViewReady)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnRequestContextInitialized:
			requestContext := &ICefRequestContext{instance: getPtr(1)}
			fn.(chromiumEventOnRequestContextInitialized)(lcl.AsObject(getPtr(0)), requestContext)
		case chromiumEventOnRequestMediaAccessPermission:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			requestingOrigin := api.GoStr(getVal(3))
			requestedPermissions := uint32(getVal(4))
			callback := &ICefMediaAccessCallback{instance: getPtr(5)}
			result := (*bool)(getPtr(6))
			*result = fn.(chromiumEventOnRequestMediaAccessPermission)(lcl.AsObject(getPtr(0)), browse, frame, requestingOrigin, requestedPermissions, callback)
		case chromiumEventOnResetDialogState:
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnResetDialogState)(lcl.AsObject(getPtr(0)), browse)
		case chromiumEventOnResolvedHostAvailable:
			result := consts.TCefErrorCode(getVal(1))
			resolvedIpsList := lcl.AsStrings(getVal(2))
			var resolvedIps []string
			if resolvedIpsList.Count() > 0 {
				resolvedIps = make([]string, resolvedIpsList.Count(), resolvedIpsList.Count())
				for i := 0; i < int(resolvedIpsList.Count()); i++ {
					resolvedIps[i] = resolvedIpsList.Strings(int32(i))
				}
				resolvedIpsList.Free()
			}
			fn.(chromiumEventOnResolvedHostAvailable)(lcl.AsObject(getPtr(0)), result, resolvedIps)
		case chromiumEventOnRouteMessageReceived:
			route := &ICefMediaRoute{instance: getPtr(1)}
			message := api.GoStr(getVal(2))
			fn.(chromiumEventOnRouteMessageReceived)(lcl.AsObject(getPtr(0)), route, message)
		case chromiumEventOnRoutes:
			routes := &TCefMediaRouteArray{instance: getPtr(1), count: uint32(int32(getVal(2)))}
			fn.(chromiumEventOnRoutes)(lcl.AsObject(getPtr(0)), routes)
		case chromiumEventOnRouteStateChanged:
			route := &ICefMediaRoute{instance: getPtr(1)}
			state := consts.TCefMediaRouteConnectionState(getVal(2))
			fn.(chromiumEventOnRouteStateChanged)(lcl.AsObject(getPtr(0)), route, state)
		case chromiumEventOnRunContextMenu:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			params := &ICefContextMenuParams{instance: getPtr(3)}
			model := &ICefMenuModel{instance: getPtr(4)}
			callback := &ICefRunContextMenuCallback{instance: getPtr(5)}
			result := (*bool)(getPtr(6))
			*result = fn.(chromiumEventOnRunContextMenu)(lcl.AsObject(getPtr(0)), browse, frame, params, model, callback)
			params.Free()
		case chromiumEventOnRunQuickMenu:
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			location := (*TCefPoint)(getPtr(3))
			size := (*TCefSize)(getPtr(4))
			editStateFlags := consts.TCefQuickMenuEditStateFlags(getVal(5))
			callback := &ICefRunQuickMenuCallback{instance: getPtr(6)}
			result := (*bool)(getPtr(7))
			*result = fn.(chromiumEventOnRunQuickMenu)(lcl.AsObject(getPtr(0)), browse, frame, location, size, editStateFlags, callback)
		case chromiumEventOnSelectClientCertificate:
			browser := &ICefBrowser{instance: getPtr(0)}
			isProxy := api.GoBool(getVal(1))
			host := api.GoStr(getVal(2))
			port := int32(getVal(3))
			certificates := &TCefX509CertificateArray{count: *(*uint32)(getPtr(4)), instance: getPtr(5)}
			callback := &ICefSelectClientCertificateCallback{instance: getPtr(6)}
			resultPtr := (*bool)(getPtr(7))
			*resultPtr = fn.(chromiumEventOnSelectClientCertificate)(lcl.AsObject(getPtr(0)), browser, isProxy, host, port, certificates, callback)
		case chromiumEventOnSetFocus:
			browse := &ICefBrowser{instance: getPtr(1)}
			source := consts.TCefFocusSource(getVal(2))
			result := (*bool)(getPtr(3))
			*result = fn.(chromiumEventOnSetFocus)(lcl.AsObject(getPtr(0)), browse, source)
		case chromiumEventOnShowPermissionPrompt:
			browse := &ICefBrowser{instance: getPtr(1)}
			promptId := *(*uint64)(getPtr(2))
			requestingOrigin := api.GoStr(getVal(3))
			requestedPermissions := uint32(getVal(4))
			callback := &ICefPermissionPromptCallback{instance: getPtr(5)}
			result := (*bool)(getPtr(6))
			*result = fn.(chromiumEventOnShowPermissionPrompt)(lcl.AsObject(getPtr(0)), browse, promptId, requestingOrigin, requestedPermissions, callback)
		case chromiumEventOnSinks:
			sinks := &TCefMediaSinkArray{instance: getPtr(1), count: uint32(int32(getVal(2)))}
			fn.(chromiumEventOnSinks)(lcl.AsObject(getPtr(0)), sinks)
		case chromiumEventOnStartDragging:
			browser := &ICefBrowser{instance: getPtr(1)}
			dragData := &ICefDragData{instance: getPtr(2)}
			allowedOps := consts.TCefDragOperations(getVal(3))
			x, y := int32(getVal(4)), int32(getVal(5))
			resultPtr := (*bool)(getPtr(6))
			*resultPtr = fn.(chromiumEventOnStartDragging)(lcl.AsObject(getPtr(0)), browser, dragData, allowedOps, x, y)
		case chromiumEventOnStatusMessage:
			browse := &ICefBrowser{instance: getPtr(1)}
			value := api.GoStr(getVal(2))
			fn.(chromiumEventOnStatusMessage)(lcl.AsObject(getPtr(0)), browse, value)
		case chromiumEventOnTakeFocus:
			browse := &ICefBrowser{instance: getPtr(1)}
			next := api.GoBool(getVal(2))
			fn.(chromiumEventOnTakeFocus)(lcl.AsObject(getPtr(0)), browse, next)
		case chromiumEventOnTextResultAvailable:
			text := api.GoStr(getVal(1))
			fn.(chromiumEventOnTextResultAvailable)(lcl.AsObject(getPtr(0)), text)
		case chromiumEventOnTextSelectionChanged:
			browser := &ICefBrowser{instance: getPtr(1)}
			selectedText := api.GoStr(getVal(2))
			selectedRange := (*TCefRange)(getPtr(3))
			fn.(chromiumEventOnTextSelectionChanged)(lcl.AsObject(getPtr(0)), browser, selectedText, selectedRange)
		case chromiumEventOnTooltip:
			browse := &ICefBrowser{instance: getPtr(1)}
			textPtr := (*uintptr)(getPtr(2))
			var text = new(string)
			*text = api.GoStr(*textPtr)
			result := (*bool)(getPtr(3))
			ok := fn.(chromiumEventOnTooltip)(lcl.AsObject(getPtr(0)), browse, text)
			*textPtr = api.PascalStr(*text)
			*result = ok
		case chromiumEventOnTouchHandleStateChanged:
			browser := &ICefBrowser{instance: getPtr(1)}
			statePtr := (*tCefTouchHandleStatePtr)(getPtr(2))
			state := statePtr.convert()
			fn.(chromiumEventOnTouchHandleStateChanged)(lcl.AsObject(getPtr(0)), browser, state)
		case chromiumEventOnUpdateDragCursor:
			browser := &ICefBrowser{instance: getPtr(1)}
			operation := consts.TCefDragOperation(getVal(2))
			fn.(chromiumEventOnUpdateDragCursor)(lcl.AsObject(getPtr(0)), browser, operation)
		case chromiumEventOnVirtualKeyboardRequested:
			browser := &ICefBrowser{instance: getPtr(1)}
			inputMode := consts.TCefTextInputMode(getVal(2))
			fn.(chromiumEventOnVirtualKeyboardRequested)(lcl.AsObject(getPtr(0)), browser, inputMode)
		case chromiumEventOnIsChromeAppMenuItemVisible:
			browser := &ICefBrowser{instance: getPtr(1)}
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(3))
			*result = fn.(chromiumEventOnIsChromeAppMenuItemVisible)(lcl.AsObject(getPtr(0)), browser, commandId)
		case chromiumEventOnIsChromeAppMenuItemEnabled:
			browser := &ICefBrowser{instance: getPtr(1)}
			commandId := int32(getVal(2))
			result := (*bool)(getPtr(3))
			*result = fn.(chromiumEventOnIsChromeAppMenuItemEnabled)(lcl.AsObject(getPtr(0)), browser, commandId)
		case chromiumEventOnIsChromePageActionIconVisible:
			buttonType := consts.TCefChromePageActionIconType(getVal(1))
			result := (*bool)(getPtr(2))
			*result = fn.(chromiumEventOnIsChromePageActionIconVisible)(lcl.AsObject(getPtr(0)), buttonType)
		case chromiumEventOnIsChromeToolbarButtonVisible:
			iconType := consts.TCefChromeToolbarButtonType(getVal(1))
			result := (*bool)(getPtr(2))
			*result = fn.(chromiumEventOnIsChromeToolbarButtonVisible)(lcl.AsObject(getPtr(0)), iconType)
		case chromiumEventOnFindResult:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnFindResult)(lcl.AsObject(sender), browse, int32(getVal(2)), int32(getVal(3)), (*TCefRect)(getPtr(4)), int32(getVal(5)), api.GoBool(getVal(6)))
		case BrowseProcessMessageReceived:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			processId := consts.CefProcessId(getVal(3))
			message := &ICefProcessMessage{instance: getPtr(4)}
			var result = (*bool)(getPtr(5))
			*result = browserProcessMessageReceived(browse, frame, message)
			if !*result {
				*result = fn.(BrowseProcessMessageReceived)(lcl.AsObject(sender), browse, frame, processId, message)
			}
			message.Free()
		case chromiumEventOnResourceLoadComplete:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(chromiumEventOnResourceLoadComplete)(sender, browse, frame, request, response, *(*consts.TCefUrlRequestStatus)(getPtr(5)), *(*int64)(getPtr(6)))
		case chromiumEventOnResourceRedirect:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			var newStr = new(string)
			var newStrPtr = (*uintptr)(getPtr(5))
			fn.(chromiumEventOnResourceRedirect)(sender, browse, frame, request, response, newStr)
			*newStrPtr = api.PascalStr(*newStr)
		case chromiumEventOnResourceResponse:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(chromiumEventOnResourceResponse)(sender, browse, frame, request, response, (*bool)(getPtr(5)))
		case chromiumEventOnBeforeResourceLoad:
			sender, browse, frame, req, _ := resourceEventGet(fn, getVal, false)
			callback := &ICefCallback{instance: getPtr(4)}
			fn.(chromiumEventOnBeforeResourceLoad)(sender, browse, frame, req, callback, (*consts.TCefReturnValue)(getPtr(5)))
		case chromiumEventOnBeforeContextMenu:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			params := &ICefContextMenuParams{instance: getPtr(3)}
			KeyAccelerator.clear()
			model := &ICefMenuModel{instance: getPtr(4), CefMis: KeyAccelerator}
			fn.(chromiumEventOnBeforeContextMenu)(lcl.AsObject(sender), browse, frame, params, model)
			params.Free()
		case chromiumEventOnContextMenuCommand:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			params := &ICefContextMenuParams{instance: getPtr(3)}
			commandId := consts.MenuId(getVal(4))
			eventFlags := uint32(getVal(5))
			result := (*bool)(getPtr(6))
			if !KeyAccelerator.commandIdEventCallback(browse, commandId, params, eventFlags, result) {
				*result = fn.(chromiumEventOnContextMenuCommand)(lcl.AsObject(sender), browse, frame, params, commandId, eventFlags)
			}
			params.Free()
		case chromiumEventOnContextMenuDismissed:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnContextMenuDismissed)(lcl.AsObject(sender), browse, frame)
		case chromiumEventOnCookieSet:
			success := api.GoBool(getVal(1))
			ID := int32(getVal(2))
			fn.(chromiumEventOnCookieSet)(lcl.AsObject(getVal(0)), success, ID)
		case chromiumEventOnCookiesDeleted:
			numDeleted := int32(getVal(1))
			fn.(chromiumEventOnCookiesDeleted)(lcl.AsObject(getVal(0)), numDeleted)
		case chromiumEventOnCookiesFlushed:
			fn.(chromiumEventOnCookiesFlushed)(lcl.AsObject(getVal(0)))
		case chromiumEventOnCookiesVisited:
			cookiePtr := (*tCefCookiePtr)(getPtr(1))
			deleteCookiePtr := (*bool)(getPtr(2))
			resultPtr := (*bool)(getPtr(3))
			cookie := cookiePtr.convert()
			fn.(chromiumEventOnCookiesVisited)(lcl.AsObject(getVal(0)), cookie, deleteCookiePtr, resultPtr)
		case chromiumEventOnCookieVisitorDestroyed:
			id := int32(getVal(1))
			fn.(chromiumEventOnCookieVisitorDestroyed)(lcl.AsObject(getVal(0)), id)
		case chromiumEventOnScrollOffsetChanged:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnScrollOffsetChanged)(lcl.AsObject(sender), browse, float64(getVal(2)), float64(getVal(3)))
		case chromiumEventOnRenderProcessTerminated:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			code := int32(getVal(2))
			error_ := api.GoStr(getVal(3))
			fn.(chromiumEventOnRenderProcessTerminated)(lcl.AsObject(sender), browse, consts.TCefTerminationStatus(getVal(2)), code, error_)
		case chromiumEventOnCompMsg:
			message := (*types.TMessage)(getPtr(1))
			lResultPtr := (*types.LRESULT)(getPtr(2))
			fn.(chromiumEventOnCompMsg)(lcl.AsObject(getVal(0)), message, lResultPtr, (*bool)(getPtr(3)))
		case chromiumEventOnTitleChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnTitleChange)(lcl.AsObject(sender), browse, api.GoStr(getVal(2)))
		case chromiumEventOnKey:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			keyEvent := (*TCefKeyEvent)(getPtr(2))
			osEvent := consts.EventHandle(getVal(3))
			fn.(chromiumEventOnKey)(lcl.AsObject(sender), browse, keyEvent, osEvent, (*bool)(getPtr(4)))
		case chromiumEventOnFullScreenModeChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnFullScreenModeChange)(lcl.AsObject(sender), browse, api.GoBool(getVal(2)))
		case chromiumEventOnBeforeBrowser:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			request := &ICefRequest{instance: getPtr(3)}
			var result = (*bool)(getPtr(6))
			*result = fn.(chromiumEventOnBeforeBrowser)(lcl.AsObject(sender), browse, frame, request, api.GoBool(getVal(4)), api.GoBool(getVal(5)))
		case chromiumEventOnAddressChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnAddressChange)(lcl.AsObject(sender), browse, frame, api.GoStr(getVal(3)))
		case chromiumEventOnAfterCreated:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnAfterCreated)(lcl.AsObject(sender), browse)
		case chromiumEventOnBeforeClose:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnBeforeClose)(lcl.AsObject(sender), browse)
		case chromiumEventOnClose:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnClose)(lcl.AsObject(sender), browse, (*consts.TCefCloseBrowserAction)(getPtr(2)))
		case chromiumEventOnResult:
			fn.(chromiumEventOnResult)(lcl.AsObject(getVal(0)), api.GoBool(getVal(1)))
		case chromiumEventOnResultFloat:
			fn.(chromiumEventOnResultFloat)(lcl.AsObject(getVal(0)), *(*float64)(getPtr(1)))
		case chromiumEventOnLoadStart:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			transitionType := consts.TCefTransitionType(getVal(3))
			fn.(chromiumEventOnLoadStart)(lcl.AsObject(sender), browse, frame, transitionType)
		case chromiumEventOnLoadingStateChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnLoadingStateChange)(lcl.AsObject(sender), browse, api.GoBool(getVal(2)), api.GoBool(getVal(3)), api.GoBool(getVal(4)))
		case chromiumEventOnLoadingProgressChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(chromiumEventOnLoadingProgressChange)(lcl.AsObject(sender), browse, *(*float64)(getPtr(2)))
		case chromiumEventOnLoadError:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnLoadError)(lcl.AsObject(sender), browse, frame, consts.CEF_NET_ERROR(getVal(3)), api.GoStr(getVal(4)), api.GoStr(getVal(5)))
		case chromiumEventOnLoadEnd:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnLoadEnd)(lcl.AsObject(sender), browse, frame, int32(getVal(3)))
		case chromiumEventOnBeforeDownload:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			downItem := &ICefDownloadItem{instance: getPtr(2)}
			suggestedName := api.GoStr(getVal(3))
			callback := &ICefBeforeDownloadCallback{instance: getPtr(4)}
			result := (*bool)(getPtr(5))
			*result = fn.(chromiumEventOnBeforeDownload)(lcl.AsObject(sender), browse, downItem, suggestedName, callback)
		case chromiumEventOnDownloadUpdated:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			downItem := &ICefDownloadItem{instance: getPtr(2)}
			callback := &ICefDownloadItemCallback{instance: getPtr(3)}
			fn.(chromiumEventOnDownloadUpdated)(lcl.AsObject(sender), browse, downItem, callback)
		case chromiumEventOnFrameAttached:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnFrameAttached)(lcl.AsObject(sender), browse, frame, api.GoBool(getVal(3)))
		case chromiumEventOnFrameCreated:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnFrameCreated)(lcl.AsObject(sender), browse, frame)
		case chromiumEventOnFrameDetached:
			sender := getVal(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(chromiumEventOnFrameDetached)(lcl.AsObject(sender), browse, frame)
		case chromiumEventOnMainFrameChanged:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			oldFrame := &ICefFrame{instance: getPtr(2)}
			newFrame := &ICefFrame{instance: getPtr(3)}
			fn.(chromiumEventOnMainFrameChanged)(lcl.AsObject(sender), browse, oldFrame, newFrame)
		case chromiumEventOnBeforePopup:
			var (
				browse             = &ICefBrowser{instance: getPtr(1)}
				frame              = &ICefFrame{instance: getPtr(2)}
				beforePInfoPtr     = (*beforePopupInfoPtr)(getPtr(3))
				popupFeaturesPtr   = (*tCefPopupFeaturesPtr)(getPtr(4))
				windowInfoPtr      = (*tCefWindowInfoPtr)(getPtr(5))
				resultClientPtr    = (*uintptr)(getPtr(6))
				browserSettingsPtr = (*tCefBrowserSettingsPtr)(getPtr(7))
				resultExtraInfoPtr = (*uintptr)(getPtr(8)) // CEF49 = nil
				noJavascriptAccess = (*bool)(getPtr(9))
				result             = (*bool)(getPtr(10))
			)
			beforePopupInfo := beforePInfoPtr.convert()
			popupFeatures := popupFeaturesPtr.convert()
			windowInfo := windowInfoPtr.convert()
			resultClient := &ICefClient{}
			browserSettings := browserSettingsPtr.convert()
			resultExtraInfo := &ICefDictionaryValue{}
			*result = fn.(chromiumEventOnBeforePopup)(lcl.AsObject(getPtr(0)), browse, frame, beforePopupInfo, popupFeatures, windowInfo, resultClient, browserSettings, resultExtraInfo, noJavascriptAccess)
			windowInfo.setInstanceValue()
			if resultClient.IsValid() {
				*resultClientPtr = resultClient.Instance()
			}
			browserSettings.setInstanceValue()
			if resultExtraInfo.IsValid() && *resultExtraInfoPtr != 0 {
				*resultExtraInfoPtr = resultExtraInfo.Instance()
			}
		case chromiumEventOnOpenUrlFromTab:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			targetUrl := api.GoStr(getVal(3))
			targetDisposition := consts.TCefWindowOpenDisposition(getVal(4))
			userGesture := api.GoBool(getVal(5))
			result := (*bool)(getPtr(6))
			*result = fn.(chromiumEventOnOpenUrlFromTab)(lcl.AsObject(sender), browse, frame, targetUrl, targetDisposition, userGesture)
		case chromiumEventOnDragEnter:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			dragData := &ICefDragData{instance: getPtr(2)}
			mask := consts.TCefDragOperations(getVal(3))
			result := (*bool)(getPtr(4))
			fn.(chromiumEventOnDragEnter)(lcl.AsObject(sender), browse, dragData, mask, result)
		case chromiumEventOnDraggableRegionsChanged:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			regionsCount := *(*uint32)(getPtr(3))
			regions := NewCefDraggableRegions()
			var region TCefDraggableRegion
			var regionSize = unsafe.Sizeof(region)
			for i := 0; i < int(regionsCount); i++ {
				region = *(*TCefDraggableRegion)(common.GetParamPtr(getVal(4), i*int(regionSize)))
				regions.Append(region)
			}
			fn.(chromiumEventOnDraggableRegionsChanged)(lcl.AsObject(sender), browse, frame, regions)
		case chromiumEventOnGetAuthCredentials:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			originUrl := api.GoStr(getVal(2))
			isProxy := api.GoBool(getVal(3))
			host := api.GoStr(getVal(4))
			port := int32(getVal(5))
			realm := api.GoStr(getVal(6))
			scheme := api.GoStr(getVal(7))
			callback := &ICefAuthCallback{instance: getPtr(8)}
			result := (*bool)(getPtr(9))
			*result = fn.(chromiumEventOnGetAuthCredentials)(lcl.AsObject(sender), browse, originUrl, isProxy, host, port, realm, scheme, callback)
			callback.Free()
		default:
			return false
		}
		return true
	})
}

func getInstance(value interface{}) unsafe.Pointer {
	var ptr uintptr
	switch value.(type) {
	case uintptr:
		ptr = value.(uintptr)
	case unsafe.Pointer:
		ptr = uintptr(value.(unsafe.Pointer))
	case lcl.IObject:
		ptr = lcl.CheckPtr(value)
	default:
		ptr = getUIntPtr(value)
	}
	return unsafe.Pointer(ptr)
}

func getUIntPtr(v interface{}) uintptr {
	switch v.(type) {
	case int:
		return uintptr(v.(int))
	case uint:
		return uintptr(v.(uint))
	case int8:
		return uintptr(v.(int8))
	case uint8:
		return uintptr(v.(uint8))
	case int16:
		return uintptr(v.(int16))
	case uint16:
		return uintptr(v.(uint16))
	case int32:
		return uintptr(v.(int32))
	case uint32:
		return uintptr(v.(uint32))
	case int64:
		return uintptr(v.(int64))
	case uint64:
		return uintptr(v.(uint64))
	}
	return 0
}
