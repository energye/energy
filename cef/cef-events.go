//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 事件回调
package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	t "github.com/energye/energy/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

func init() {
	var resourceEventGet = func(fn interface{}, getVal func(idx int) uintptr, resp bool) (sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) {
		var (
			instance unsafe.Pointer
		)
		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		senderPtr := getPtr(0)
		browser = &ICefBrowser{instance: getPtr(1)}
		frame = &ICefFrame{instance: getPtr(2)}
		cefRequest := (*iCefRequestPtr)(getPtr(3))
		request = &ICefRequest{
			instance:             common.GetInstancePtr(cefRequest.Instance),
			Url:                  api.GoStr(cefRequest.Url),
			Method:               api.GoStr(cefRequest.Method),
			ReferrerUrl:          api.GoStr(cefRequest.ReferrerUrl),
			ReferrerPolicy:       consts.TCefReferrerPolicy(cefRequest.ReferrerPolicy),
			Flags:                consts.TCefUrlRequestFlags(cefRequest.Flags),
			FirstPartyForCookies: api.GoStr(cefRequest.FirstPartyForCookies),
			ResourceType:         consts.TCefResourceType(cefRequest.ResourceType),
			TransitionType:       consts.TCefTransitionType(cefRequest.TransitionType),
			Identifier:           *(*uint64)(common.GetParamPtr(cefRequest.Identifier, 0)),
		}
		if resp {
			cefResponse := (*iCefResponsePtr)(getPtr(4))
			instance = common.GetInstancePtr(cefResponse.Instance)
			response = &ICefResponse{
				instance:   instance,
				Status:     int32(cefResponse.Status),
				StatusText: api.GoStr(cefResponse.StatusText),
				MimeType:   api.GoStr(cefResponse.MimeType),
				Charset:    api.GoStr(cefResponse.Charset),
				Error:      consts.TCefErrorCode(cefResponse.Error),
				URL:        api.GoStr(cefResponse.URL),
			}
		}
		return lcl.AsObject(senderPtr), browser, frame, request, response
	}
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		var (
			instance unsafe.Pointer
		)
		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case ChromiumEventOnFindResult:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnFindResult)(lcl.AsObject(sender), browse, int32(getVal(2)), int32(getVal(3)), (*TCefRect)(getPtr(4)), int32(getVal(5)), api.GoBool(getVal(6)))
			//browse.Free()
		case BrowseProcessMessageReceived:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			processId := consts.CefProcessId(getVal(3))
			message := &ICefProcessMessage{instance: getPtr(4)}
			var result = (*bool)(getPtr(5))
			*result = fn.(BrowseProcessMessageReceived)(lcl.AsObject(sender), browse, frame, processId, message)
			if !*result {
				*result = browserProcessMessageReceived(browse, frame, message)
			}
			//frame.Free()
			//browse.Free()
			message.Free()
		case ChromiumEventOnResourceLoadComplete:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(ChromiumEventOnResourceLoadComplete)(sender, browse, frame, request, response, *(*consts.TCefUrlRequestStatus)(getPtr(5)), *(*int64)(getPtr(6)))
			//frame.Free()
			//browse.Free()
			//request.Free()
			//response.Free()
		case ChromiumEventOnResourceRedirect:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			var newStr = new(t.TString)
			var newStrPtr = (*uintptr)(getPtr(5))
			fn.(ChromiumEventOnResourceRedirect)(sender, browse, frame, request, response, newStr)
			*newStrPtr = newStr.ToPtr()
			//frame.Free()
			//browse.Free()
			//request.Free()
			//response.Free()
		case ChromiumEventOnResourceResponse:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(ChromiumEventOnResourceResponse)(sender, browse, frame, request, response, (*bool)(getPtr(5)))
			//frame.Free()
			//browse.Free()
			//request.Free()
			//response.Free()
		case ChromiumEventOnBeforeResourceLoad:
			sender, browse, frame, request, _ := resourceEventGet(fn, getVal, false)
			instance = getInstance(getVal(4))
			callback := &ICefCallback{instance: instance}
			fn.(ChromiumEventOnBeforeResourceLoad)(sender, browse, frame, request, callback, (*consts.TCefReturnValue)(getPtr(5)))
			//frame.Free()
			//browse.Free()
			//request.Free()
		//menu begin
		case ChromiumEventOnBeforeContextMenu:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			cefParams := (*iCefContextMenuParams)(getPtr(3))
			params := &ICefContextMenuParams{
				XCoord:            int32(cefParams.XCoord),
				YCoord:            int32(cefParams.YCoord),
				TypeFlags:         consts.TCefContextMenuTypeFlags(cefParams.TypeFlags),
				LinkUrl:           api.GoStr(cefParams.LinkUrl),
				UnfilteredLinkUrl: api.GoStr(cefParams.UnfilteredLinkUrl),
				SourceUrl:         api.GoStr(cefParams.SourceUrl),
				TitleText:         api.GoStr(cefParams.TitleText),
				PageUrl:           api.GoStr(cefParams.PageUrl),
				FrameUrl:          api.GoStr(cefParams.FrameUrl),
				FrameCharset:      api.GoStr(cefParams.FrameCharset),
				MediaType:         consts.TCefContextMenuMediaType(cefParams.MediaType),
				MediaStateFlags:   consts.TCefContextMenuMediaStateFlags(cefParams.MediaStateFlags),
				SelectionText:     api.GoStr(cefParams.SelectionText),
				EditStateFlags:    consts.TCefContextMenuEditStateFlags(cefParams.EditStateFlags),
			}
			KeyAccelerator.clear()
			model := &ICefMenuModel{instance: getPtr(4), CefMis: KeyAccelerator}
			fn.(ChromiumEventOnBeforeContextMenu)(lcl.AsObject(sender), browse, frame, params, model)
			//frame.Free()
			//browse.Free()
			//model.Free()
		case ChromiumEventOnContextMenuCommand:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			cefParams := (*iCefContextMenuParams)(getPtr(3))
			params := &ICefContextMenuParams{
				XCoord:            int32(cefParams.XCoord),
				YCoord:            int32(cefParams.YCoord),
				TypeFlags:         consts.TCefContextMenuTypeFlags(cefParams.TypeFlags),
				LinkUrl:           api.GoStr(cefParams.LinkUrl),
				UnfilteredLinkUrl: api.GoStr(cefParams.UnfilteredLinkUrl),
				SourceUrl:         api.GoStr(cefParams.SourceUrl),
				TitleText:         api.GoStr(cefParams.TitleText),
				PageUrl:           api.GoStr(cefParams.PageUrl),
				FrameUrl:          api.GoStr(cefParams.FrameUrl),
				FrameCharset:      api.GoStr(cefParams.FrameCharset),
				MediaType:         consts.TCefContextMenuMediaType(cefParams.MediaType),
				MediaStateFlags:   consts.TCefContextMenuMediaStateFlags(cefParams.MediaStateFlags),
				SelectionText:     api.GoStr(cefParams.SelectionText),
				EditStateFlags:    consts.TCefContextMenuEditStateFlags(cefParams.EditStateFlags),
			}
			commandId := consts.MenuId(getVal(4))
			eventFlags := uint32(getVal(5))
			result := (*bool)(getPtr(6))
			if !KeyAccelerator.commandIdEventCallback(browse, commandId, params, eventFlags, result) {
				fn.(ChromiumEventOnContextMenuCommand)(lcl.AsObject(sender), browse, frame, params, commandId, eventFlags, result)
			}
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnContextMenuDismissed:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(ChromiumEventOnContextMenuDismissed)(lcl.AsObject(sender), browse, frame)
			//frame.Free()
			//browse.Free()
		//menu end
		//---
		//cookie begin
		case ChromiumEventOnCookieSet:
			success := api.GoBool(getVal(1))
			ID := int32(getVal(2))
			fn.(ChromiumEventOnCookieSet)(lcl.AsObject(getVal(0)), success, ID)
		case ChromiumEventOnCookiesDeleted:
			numDeleted := int32(getVal(1))
			fn.(ChromiumEventOnCookiesDeleted)(lcl.AsObject(getVal(0)), numDeleted)
		case ChromiumEventOnCookiesFlushed:
			fn.(ChromiumEventOnCookiesFlushed)(lcl.AsObject(getVal(0)))
		case ChromiumEventOnCookiesVisited:
			cookie := *(*iCefCookiePtr)(getPtr(1))
			creation := *(*float64)(common.GetParamPtr(cookie.creation, 0))
			lastAccess := *(*float64)(common.GetParamPtr(cookie.lastAccess, 0))
			expires := *(*float64)(common.GetParamPtr(cookie.expires, 0))
			iCookie := &ICefCookie{
				Url:            api.GoStr(cookie.url),
				Name:           api.GoStr(cookie.name),
				Value:          api.GoStr(cookie.value),
				Domain:         api.GoStr(cookie.domain),
				Path:           api.GoStr(cookie.path),
				Secure:         *(*bool)(common.GetParamPtr(cookie.secure, 0)),
				Httponly:       *(*bool)(common.GetParamPtr(cookie.httponly, 0)),
				HasExpires:     *(*bool)(common.GetParamPtr(cookie.hasExpires, 0)),
				Creation:       common.DDateTimeToGoDateTime(creation),
				LastAccess:     common.DDateTimeToGoDateTime(lastAccess),
				Expires:        common.DDateTimeToGoDateTime(expires),
				Count:          int32(cookie.count),
				Total:          int32(cookie.total),
				ID:             int32(cookie.aID),
				SameSite:       consts.TCefCookieSameSite(cookie.sameSite),
				Priority:       consts.TCefCookiePriority(cookie.priority),
				SetImmediately: *(*bool)(common.GetParamPtr(cookie.aSetImmediately, 0)),
				DeleteCookie:   *(*bool)(common.GetParamPtr(cookie.aDeleteCookie, 0)),
				Result:         *(*bool)(common.GetParamPtr(cookie.aResult, 0)),
			}
			fn.(ChromiumEventOnCookiesVisited)(lcl.AsObject(getVal(0)), iCookie)
		case ChromiumEventOnCookieVisitorDestroyed:
			id := int32(getVal(1))
			fn.(ChromiumEventOnCookieVisitorDestroyed)(lcl.AsObject(getVal(0)), id)
		//cookie end
		//--- other
		case ChromiumEventOnScrollOffsetChanged:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnScrollOffsetChanged)(lcl.AsObject(sender), browse, float64(getVal(2)), float64(getVal(3)))
			//browse.Free()
		case ChromiumEventOnRenderProcessTerminated:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnRenderProcessTerminated)(lcl.AsObject(sender), browse, consts.TCefTerminationStatus(getVal(2)))
			//browse.Free()
		case ChromiumEventOnCompMsg:
			message := (*types.TMessage)(getPtr(1))
			lResultPtr := (*types.LRESULT)(getPtr(2))
			fn.(ChromiumEventOnCompMsg)(lcl.AsObject(getVal(0)), message, lResultPtr, (*bool)(getPtr(3)))
		case ChromiumEventOnCefBrowser:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnCefBrowser)(lcl.AsObject(sender), browse)
			//browse.Free()
		case ChromiumEventOnTitleChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnTitleChange)(lcl.AsObject(sender), browse, api.GoStr(getVal(2)))
			//browse.Free()
		case ChromiumEventOnKeyEvent:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			keyEvent := (*TCefKeyEvent)(getPtr(2))
			fn.(ChromiumEventOnKeyEvent)(lcl.AsObject(sender), browse, keyEvent, (*bool)(getPtr(3)))
			//browse.Free()
		case ChromiumEventOnFullScreenModeChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnFullScreenModeChange)(lcl.AsObject(sender), browse, api.GoBool(getVal(2)))
			//browse.Free()
		case ChromiumEventOnBeforeBrowser: //创建浏览器之前
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			var result = (*bool)(getPtr(3))
			*result = fn.(ChromiumEventOnBeforeBrowser)(lcl.AsObject(sender), browse, frame)
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnAddressChange: //创建浏览器之前
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(ChromiumEventOnAddressChange)(lcl.AsObject(sender), browse, frame, api.GoStr(getVal(3)))
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnAfterCreated: //创建浏览器之后
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnAfterCreated)(lcl.AsObject(sender), browse)
			//browse.Free()
		case ChromiumEventOnBeforeClose: //关闭浏览器之前
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnBeforeClose)(lcl.AsObject(sender), browse)
			//browse.Free()
		case ChromiumEventOnClose:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnClose)(lcl.AsObject(sender), browse, (*TCefCloseBrowsesAction)(getPtr(2)))
			//browse.Free()
		case ChromiumEventOnResult: //通用Result bool事件
			fn.(ChromiumEventOnResult)(lcl.AsObject(getVal(0)), api.GoBool(getVal(1)))
		case ChromiumEventOnResultFloat: //通用Result float事件
			fn.(ChromiumEventOnResultFloat)(lcl.AsObject(getVal(0)), *(*float64)(getPtr(1)))
		case ChromiumEventOnLoadStart:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			transitionType := consts.TCefTransitionType(getVal(3))
			fn.(ChromiumEventOnLoadStart)(lcl.AsObject(sender), browse, frame, transitionType)
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnLoadingStateChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnLoadingStateChange)(lcl.AsObject(sender), browse, api.GoBool(getVal(2)), api.GoBool(getVal(3)), api.GoBool(getVal(4)))
			//browse.Free()
		case ChromiumEventOnLoadingProgressChange:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			fn.(ChromiumEventOnLoadingProgressChange)(lcl.AsObject(sender), browse, *(*float64)(getPtr(2)))
			//browse.Free()
		case ChromiumEventOnLoadError:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(ChromiumEventOnLoadError)(lcl.AsObject(sender), browse, frame, consts.CEF_NET_ERROR(getVal(3)), api.GoStr(getVal(4)), api.GoStr(getVal(5)))
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnLoadEnd:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(ChromiumEventOnLoadEnd)(lcl.AsObject(sender), browse, frame, int32(getVal(3)))
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnBeforeDownload: //下载之前
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			downItem := &ICefDownloadItem{instance: getPtr(2)}
			suggestedName := api.GoStr(getVal(3))
			callback := &ICefBeforeDownloadCallback{instance: getPtr(4)}
			fn.(ChromiumEventOnBeforeDownload)(lcl.AsObject(sender), browse, downItem, suggestedName, callback)
			//browse.Free()
		case ChromiumEventOnDownloadUpdated: //下载更新
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			downItem := &ICefDownloadItem{instance: getPtr(2)}
			callback := &ICefDownloadItemCallback{instance: getPtr(3)}
			fn.(ChromiumEventOnDownloadUpdated)(lcl.AsObject(sender), browse, downItem, callback)
			//browse.Free()
		//frame
		case ChromiumEventOnFrameAttached:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(ChromiumEventOnFrameAttached)(lcl.AsObject(sender), browse, frame, api.GoBool(getVal(3)))
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnFrameCreated:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(ChromiumEventOnFrameCreated)(lcl.AsObject(sender), browse, frame)
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnFrameDetached:
			sender := getVal(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			fn.(ChromiumEventOnFrameDetached)(lcl.AsObject(sender), browse, frame)
			//frame.Free()
			//browse.Free()
		case ChromiumEventOnMainFrameChanged:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			oldFrame := &ICefFrame{instance: getPtr(2)}
			newFrame := &ICefFrame{instance: getPtr(3)}
			fn.(ChromiumEventOnMainFrameChanged)(lcl.AsObject(sender), browse, oldFrame, newFrame)
			//oldFrame.Free()
			//newFrame.Free()
			//browse.Free()
		//windowParent popup
		case ChromiumEventOnBeforePopup:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			beforePInfoPtr := (*beforePopupInfoPtr)(getPtr(3))
			beforePInfo := &BeforePopupInfo{
				TargetUrl:         api.GoStr(beforePInfoPtr.TargetUrl),
				TargetFrameName:   api.GoStr(beforePInfoPtr.TargetFrameName),
				TargetDisposition: consts.TCefWindowOpenDisposition(beforePInfoPtr.TargetDisposition),
				UserGesture:       api.GoBool(beforePInfoPtr.UserGesture),
			}
			//windowInfo:=getPtr(4)
			var (
				client             = &ICefClient{instance: getPtr(5)}
				noJavascriptAccess = (*bool)(getPtr(6))
				result             = (*bool)(getPtr(7))
			)
			//callback
			*result = fn.(ChromiumEventOnBeforePopup)(lcl.AsObject(sender), browse, frame, beforePInfo, client, noJavascriptAccess)
			//frame.Free()
			//browse.Free()
		//windowParent open url from tab
		case ChromiumEventOnOpenUrlFromTab:
			//no impl
		case ChromiumEventOnDragEnter:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			dragData := &ICefDragData{instance: getPtr(2)}
			mask := consts.TCefDragOperations(getVal(3))
			result := (*bool)(getPtr(4))
			fn.(ChromiumEventOnDragEnter)(lcl.AsObject(sender), browse, dragData, mask, result)
			//browse.Free()
		case ChromiumEventOnDraggableRegionsChanged:
			sender := getPtr(0)
			browse := &ICefBrowser{instance: getPtr(1)}
			frame := &ICefFrame{instance: getPtr(2)}
			regionsCount := int32(getVal(3))
			regions := NewCefDraggableRegions()
			var region TCefDraggableRegion
			for i := 0; i < int(regionsCount); i++ {
				region = *(*TCefDraggableRegion)(common.GetParamPtr(getVal(4), i*int(unsafe.Sizeof(TCefDraggableRegion{}))))
				regions.Append(region)
			}
			fn.(ChromiumEventOnDraggableRegionsChanged)(lcl.AsObject(sender), browse, frame, regions)
			//frame.Free()
			//browse.Free()
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
