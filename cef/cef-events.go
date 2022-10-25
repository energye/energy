//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

func init() {
	var resourceEventGet = func(fn interface{}, getVal func(idx int) uintptr, resp bool) (sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse) {
		var (
			instance uintptr
			ptr      unsafe.Pointer
		)
		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		senderPtr := getVal(0)
		browser = &ICefBrowser{browseId: int32(getVal(1)), chromium: senderPtr}
		tempFrame := (*cefFrame)(getPtr(2))
		frame = &ICefFrame{
			Browser: browser,
			Name:    api.DStrToGoStr(tempFrame.Name),
			Url:     api.DStrToGoStr(tempFrame.Url),
			Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
		}
		cefRequest := (*rICefRequest)(getPtr(3))
		instance, ptr = GetInstancePtr(cefRequest.Instance)
		request = &ICefRequest{
			instance:             instance,
			ptr:                  ptr,
			Url:                  api.DStrToGoStr(cefRequest.Url),
			Method:               api.DStrToGoStr(cefRequest.Method),
			ReferrerUrl:          api.DStrToGoStr(cefRequest.ReferrerUrl),
			ReferrerPolicy:       int32(cefRequest.ReferrerPolicy),
			Flags:                int(cefRequest.Flags),
			FirstPartyForCookies: api.DStrToGoStr(cefRequest.FirstPartyForCookies),
			ResourceType:         int32(cefRequest.ResourceType),
			TransitionType:       int(cefRequest.TransitionType),
			Identifier:           *(*uint64)(GetParamPtr(cefRequest.Identifier, 0)),
		}
		if resp {
			cefResponse := (*iCefResponse)(getPtr(4))
			instance, ptr = GetInstancePtr(cefResponse.Instance)
			response = &ICefResponse{
				instance:   instance,
				ptr:        ptr,
				Status:     int32(cefResponse.Status),
				StatusText: api.DStrToGoStr(cefResponse.StatusText),
				MimeType:   api.DStrToGoStr(cefResponse.MimeType),
				Charset:    api.DStrToGoStr(cefResponse.Charset),
				Error:      int32(cefResponse.Error),
				URL:        api.DStrToGoStr(cefResponse.URL),
			}
		}
		return lcl.AsObject(senderPtr), browser, frame, request, response
	}
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("CEF Events Error:", err)
			}
		}()
		var (
			instance uintptr
			ptr      unsafe.Pointer
		)
		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case ChromiumEventOnFindResult:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			cefRectPtr := (*tCefRect)(getPtr(4))
			cefRect := &TCefRect{
				X:      int32(cefRectPtr.X),
				Y:      int32(cefRectPtr.Y),
				Width:  int32(cefRectPtr.Width),
				Height: int32(cefRectPtr.Height),
			}
			fn.(ChromiumEventOnFindResult)(lcl.AsObject(sender), browser, int32(getVal(2)), int32(getVal(3)), cefRect, int32(getVal(5)), api.DBoolToGoBool(getVal(6)))
		case BrowseProcessMessageReceived:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			cefProcMsg := (*ipc.CefProcessMessagePtr)(getPtr(4))
			args := ipc.NewArgumentList()
			args.UnPackageBytePtr(cefProcMsg.Data, int32(cefProcMsg.DataLen))
			processMessage := &ipc.ICefProcessMessage{
				Name:         api.DStrToGoStr(cefProcMsg.Name),
				ArgumentList: args,
			}
			var result = (*bool)(getPtr(5))
			*result = fn.(BrowseProcessMessageReceived)(lcl.AsObject(sender), browser, frame, CefProcessId(getVal(3)), processMessage)
			args.Clear()
			cefProcMsg.Data = 0
			cefProcMsg.DataLen = 0
			cefProcMsg.Name = 0
			cefProcMsg = nil
			args = nil
		case ChromiumEventOnResourceLoadComplete:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(ChromiumEventOnResourceLoadComplete)(sender, browse, frame, request, response, *(*TCefUrlRequestStatus)(getPtr(5)), *(*int64)(getPtr(6)))
		case ChromiumEventOnResourceRedirect:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			var newStr = &String{}
			var newStrPtr = (*uintptr)(getPtr(5))
			fn.(ChromiumEventOnResourceRedirect)(sender, browse, frame, request, response, newStr)
			*newStrPtr = api.GoStrToDStr(newStr.GetValue())
		case ChromiumEventOnResourceResponse:
			sender, browse, frame, request, response := resourceEventGet(fn, getVal, true)
			fn.(ChromiumEventOnResourceResponse)(sender, browse, frame, request, response, (*bool)(getPtr(5)))
		case ChromiumEventOnBeforeResourceLoad:
			sender, browse, frame, request, _ := resourceEventGet(fn, getVal, false)
			instance, ptr = getInstance(getVal(4))
			callback := &ICefCallback{instance: instance, ptr: ptr}
			fn.(ChromiumEventOnBeforeResourceLoad)(sender, browse, frame, request, callback, (*TCefReturnValue)(getPtr(5)))
		//menu begin
		case ChromiumEventOnBeforeContextMenu:
			sender := getVal(0)
			instance, ptr = getInstance(getVal(1))
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			cefParams := (*iCefContextMenuParams)(getPtr(3))
			params := &ICefContextMenuParams{
				XCoord:            int32(cefParams.XCoord),
				YCoord:            int32(cefParams.YCoord),
				TypeFlags:         TCefContextMenuTypeFlags(cefParams.TypeFlags),
				LinkUrl:           api.DStrToGoStr(cefParams.LinkUrl),
				UnfilteredLinkUrl: api.DStrToGoStr(cefParams.UnfilteredLinkUrl),
				SourceUrl:         api.DStrToGoStr(cefParams.SourceUrl),
				TitleText:         api.DStrToGoStr(cefParams.TitleText),
				PageUrl:           api.DStrToGoStr(cefParams.PageUrl),
				FrameUrl:          api.DStrToGoStr(cefParams.FrameUrl),
				FrameCharset:      api.DStrToGoStr(cefParams.FrameCharset),
				MediaType:         TCefContextMenuMediaType(cefParams.MediaType),
				MediaStateFlags:   TCefContextMenuMediaStateFlags(cefParams.MediaStateFlags),
				SelectionText:     api.DStrToGoStr(cefParams.SelectionText),
				EditStateFlags:    TCefContextMenuEditStateFlags(cefParams.EditStateFlags),
			}
			instance, ptr = getInstance(getVal(4))
			KeyAccelerator.clear()
			model := &ICefMenuModel{instance: instance, ptr: ptr, CefMis: KeyAccelerator}
			fn.(ChromiumEventOnBeforeContextMenu)(lcl.AsObject(sender), browser, frame, params, model)
		case ChromiumEventOnContextMenuCommand:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			cefParams := (*iCefContextMenuParams)(getPtr(3))
			params := &ICefContextMenuParams{
				XCoord:            int32(cefParams.XCoord),
				YCoord:            int32(cefParams.YCoord),
				TypeFlags:         TCefContextMenuTypeFlags(cefParams.TypeFlags),
				LinkUrl:           api.DStrToGoStr(cefParams.LinkUrl),
				UnfilteredLinkUrl: api.DStrToGoStr(cefParams.UnfilteredLinkUrl),
				SourceUrl:         api.DStrToGoStr(cefParams.SourceUrl),
				TitleText:         api.DStrToGoStr(cefParams.TitleText),
				PageUrl:           api.DStrToGoStr(cefParams.PageUrl),
				FrameUrl:          api.DStrToGoStr(cefParams.FrameUrl),
				FrameCharset:      api.DStrToGoStr(cefParams.FrameCharset),
				MediaType:         TCefContextMenuMediaType(cefParams.MediaType),
				MediaStateFlags:   TCefContextMenuMediaStateFlags(cefParams.MediaStateFlags),
				SelectionText:     api.DStrToGoStr(cefParams.SelectionText),
				EditStateFlags:    TCefContextMenuEditStateFlags(cefParams.EditStateFlags),
			}
			commandId := int32(getVal(4))
			eventFlags := uint32(getVal(5))
			if !KeyAccelerator.commandIdEventCallback(browser, commandId, params, eventFlags, (*bool)(getPtr(5))) {
				fn.(ChromiumEventOnContextMenuCommand)(lcl.AsObject(sender), browser, frame, params, commandId, eventFlags, (*bool)(getPtr(6)))
			}
		case ChromiumEventOnContextMenuDismissed:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnContextMenuDismissed)(lcl.AsObject(sender), browser, frame)
		//menu end
		//---
		//cookie begin
		case ChromiumEventOnCookieSet:
			success := api.DBoolToGoBool(getVal(1))
			ID := int32(getVal(2))
			fn.(ChromiumEventOnCookieSet)(lcl.AsObject(getVal(0)), success, ID)
		case ChromiumEventOnCookiesDeleted:
			numDeleted := int32(getVal(1))
			fn.(ChromiumEventOnCookiesDeleted)(lcl.AsObject(getVal(0)), numDeleted)
		case ChromiumEventOnCookiesFlushed:
			fn.(ChromiumEventOnCookiesFlushed)(lcl.AsObject(getVal(0)))
		case ChromiumEventOnCookiesVisited:
			cookie := *(*cefCookie)(getPtr(1))
			creation := *(*float64)(GetParamPtr(cookie.creation, 0))
			lastAccess := *(*float64)(GetParamPtr(cookie.lastAccess, 0))
			expires := *(*float64)(GetParamPtr(cookie.expires, 0))
			iCookie := &ICefCookie{
				Url:            api.DStrToGoStr(cookie.url),
				Name:           api.DStrToGoStr(cookie.name),
				Value:          api.DStrToGoStr(cookie.value),
				Domain:         api.DStrToGoStr(cookie.domain),
				Path:           api.DStrToGoStr(cookie.path),
				Secure:         *(*bool)(GetParamPtr(cookie.secure, 0)),
				Httponly:       *(*bool)(GetParamPtr(cookie.httponly, 0)),
				HasExpires:     *(*bool)(GetParamPtr(cookie.hasExpires, 0)),
				Creation:       DDateTimeToGoDateTime(creation),
				LastAccess:     DDateTimeToGoDateTime(lastAccess),
				Expires:        DDateTimeToGoDateTime(expires),
				Count:          int32(cookie.count),
				Total:          int32(cookie.total),
				ID:             int32(cookie.aID),
				SameSite:       TCefCookieSameSite(cookie.sameSite),
				Priority:       TCefCookiePriority(cookie.priority),
				SetImmediately: *(*bool)(GetParamPtr(cookie.aSetImmediately, 0)),
				DeleteCookie:   *(*bool)(GetParamPtr(cookie.aDeleteCookie, 0)),
				Result:         *(*bool)(GetParamPtr(cookie.aResult, 0)),
			}
			fn.(ChromiumEventOnCookiesVisited)(lcl.AsObject(getVal(0)), iCookie)
		case ChromiumEventOnCookieVisitorDestroyed:
			id := int32(getVal(1))
			fn.(ChromiumEventOnCookieVisitorDestroyed)(lcl.AsObject(getVal(0)), id)
		//cookie end
		//--- other
		case ChromiumEventOnScrollOffsetChanged:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnScrollOffsetChanged)(lcl.AsObject(sender), browser, float64(getVal(2)), float64(getVal(2)))
		case ChromiumEventOnRenderProcessTerminated:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnRenderProcessTerminated)(lcl.AsObject(sender), browser, TCefTerminationStatus(getVal(2)))
		case ChromiumEventOnRenderCompMsg:
			message := *(*types.TMessage)(getPtr(1))
			fn.(ChromiumEventOnRenderCompMsg)(lcl.AsObject(getVal(0)), message, api.DBoolToGoBool(getVal(2)))
		case ChromiumEventOnCefBrowser:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnCefBrowser)(lcl.AsObject(sender), browser)
		case ChromiumEventOnTitleChange:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnTitleChange)(lcl.AsObject(sender), browser, api.DStrToGoStr(getVal(2)))
		case ChromiumEventOnKeyEvent:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			keyEvent := (*TCefKeyEvent)(getPtr(2))
			fn.(ChromiumEventOnKeyEvent)(lcl.AsObject(sender), browser, keyEvent, (*bool)(getPtr(3)))
		case ChromiumEventOnFullScreenModeChange:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnFullScreenModeChange)(lcl.AsObject(sender), browser, api.DBoolToGoBool(getVal(2)))
		case ChromiumEventOnBeforeBrowser: //创建浏览器之前
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			chromiumOnBeforeBrowser(browser, frame)
			var result = (*bool)(getPtr(3))
			*result = fn.(ChromiumEventOnBeforeBrowser)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnAddressChange: //创建浏览器之前
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnAddressChange)(lcl.AsObject(sender), browser, frame, api.DStrToGoStr(getVal(3)))
		case ChromiumEventOnAfterCreated: //创建浏览器之后
			sender := getVal(0)
			//事件处理函数返回true将不继续执行
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			if chromiumOnAfterCreate(browser) {
				return true
			}
			fn.(ChromiumEventOnAfterCreated)(lcl.AsObject(sender), browser)
		case ChromiumEventOnBeforeClose: //关闭浏览器之前
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			chromiumOnBeforeClose(browser)
			fn.(ChromiumEventOnBeforeClose)(lcl.AsObject(sender), browser)
		case ChromiumEventOnClose:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			chromiumOnClose(browser)
			fn.(ChromiumEventOnClose)(lcl.AsObject(sender), browser, (*TCefCloseBrowsesAction)(getPtr(2)))
		case ChromiumEventOnResult: //通用Result bool事件
			fn.(ChromiumEventOnResult)(lcl.AsObject(getVal(0)), api.DBoolToGoBool(getVal(1)))
		case ChromiumEventOnResultFloat: //通用Result float事件
			fn.(ChromiumEventOnResultFloat)(lcl.AsObject(getVal(0)), *(*float64)(getPtr(1)))
		case ChromiumEventOnLoadStart:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			BrowserWindow.putBrowserFrame(browser, frame)
			fn.(ChromiumEventOnLoadStart)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnLoadingStateChange:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			BrowserWindow.putBrowserFrame(browser, nil)
			fn.(ChromiumEventOnLoadingStateChange)(lcl.AsObject(sender), browser, api.DBoolToGoBool(getVal(2)),
				api.DBoolToGoBool(getVal(3)), api.DBoolToGoBool(getVal(4)))
		case ChromiumEventOnLoadingProgressChange:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			fn.(ChromiumEventOnLoadingProgressChange)(lcl.AsObject(sender), browser, *(*float64)(getPtr(2)))
		case ChromiumEventOnLoadError:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnLoadError)(lcl.AsObject(sender), browser, frame, CEF_NET_ERROR(getVal(3)), api.DStrToGoStr(getVal(4)), api.DStrToGoStr(getVal(5)))
		case ChromiumEventOnLoadEnd:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnLoadEnd)(lcl.AsObject(sender), browser, frame, int32(getVal(3)))
		case ChromiumEventOnBeforeDownload: //下载之前
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			item := (*downloadItem)(getPtr(2))
			downItem := &DownloadItem{
				Id:                 int32(item.Id),
				CurrentSpeed:       int64(item.CurrentSpeed),
				PercentComplete:    int32(item.PercentComplete),
				TotalBytes:         int64(item.TotalBytes),
				ReceivedBytes:      int64(item.ReceivedBytes),
				StartTime:          DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.StartTime, 0))),
				EndTime:            DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.EndTime, 0))),
				FullPath:           api.DStrToGoStr(item.FullPath),
				Url:                api.DStrToGoStr(item.Url),
				OriginalUrl:        api.DStrToGoStr(item.OriginalUrl),
				SuggestedFileName:  api.DStrToGoStr(item.SuggestedFileName),
				ContentDisposition: api.DStrToGoStr(item.ContentDisposition),
				MimeType:           api.DStrToGoStr(item.MimeType),
				IsValid:            *(*bool)(unsafe.Pointer(item.IsValid)),
				State:              int32(item.State),
			}
			suggestedName := api.DStrToGoStr(getVal(3))
			instance, ptr = getInstance(getVal(4))
			callback := &ICefBeforeDownloadCallback{
				instance: instance, ptr: ptr,
				browseId: browser.Identifier(),
				downId:   downItem.Id,
			}
			fn.(ChromiumEventOnBeforeDownload)(lcl.AsObject(sender), browser, downItem, suggestedName, callback)
		case ChromiumEventOnDownloadUpdated: //下载更新
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			item := *(*downloadItem)(getPtr(2))
			downItem := &DownloadItem{
				Id:                 int32(item.Id),
				CurrentSpeed:       int64(item.CurrentSpeed),
				PercentComplete:    int32(item.PercentComplete),
				TotalBytes:         int64(item.TotalBytes),
				ReceivedBytes:      int64(item.ReceivedBytes),
				StartTime:          DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.StartTime, 0))),
				EndTime:            DDateTimeToGoDateTime(*(*float64)(GetParamPtr(item.EndTime, 0))),
				FullPath:           api.DStrToGoStr(item.FullPath),
				Url:                api.DStrToGoStr(item.Url),
				OriginalUrl:        api.DStrToGoStr(item.OriginalUrl),
				SuggestedFileName:  api.DStrToGoStr(item.SuggestedFileName),
				ContentDisposition: api.DStrToGoStr(item.ContentDisposition),
				MimeType:           api.DStrToGoStr(item.MimeType),
				IsValid:            *(*bool)(unsafe.Pointer(item.IsValid)),
				State:              int32(item.State),
			}
			instance, ptr = getInstance(getVal(3))
			callback := &ICefDownloadItemCallback{
				instance: instance, ptr: ptr,
				browseId: browser.Identifier(),
				downId:   downItem.Id,
			}
			fn.(ChromiumEventOnDownloadUpdated)(lcl.AsObject(sender), browser, downItem, callback)
		//frame
		case ChromiumEventOnFrameAttached:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(ChromiumEventOnFrameAttached)(lcl.AsObject(sender), browser, frame, api.DBoolToGoBool(getVal(3)))
		case ChromiumEventOnFrameCreated:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
			}
			fn.(ChromiumEventOnFrameCreated)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnFrameDetached:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1))}
			tempFrame := (*cefFrame)(getPtr(2))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			chromiumOnFrameDetached(browser, frame)
			fn.(ChromiumEventOnFrameDetached)(lcl.AsObject(sender), browser, frame)
		case ChromiumEventOnMainFrameChanged:
			sender := getVal(0)
			browser := &ICefBrowser{browseId: int32(getVal(1)), chromium: sender}
			var (
				oldFrame *ICefFrame = nil
				newFrame *ICefFrame = nil
			)
			tempOldFrame := (*cefFrame)(getPtr(2))
			if tempOldFrame != nil {
				oldFrame = &ICefFrame{
					Browser: browser,
					Name:    api.DStrToGoStr(tempOldFrame.Name),
					Url:     api.DStrToGoStr(tempOldFrame.Url),
					Id:      StrToInt64(api.DStrToGoStr(tempOldFrame.Identifier)),
				}
			}
			tempNewFrame := (*cefFrame)(getPtr(3))
			if tempNewFrame != nil {
				newFrame = &ICefFrame{
					Browser: browser,
					Name:    api.DStrToGoStr(tempNewFrame.Name),
					Url:     api.DStrToGoStr(tempNewFrame.Url),
					Id:      StrToInt64(api.DStrToGoStr(tempNewFrame.Identifier)),
				}
			}
			chromiumOnMainFrameChanged(browser, oldFrame, newFrame)
			fn.(ChromiumEventOnMainFrameChanged)(lcl.AsObject(sender), browser, oldFrame, newFrame)
		//windowParent popup
		case ChromiumEventOnBeforePopup:
			chromiumOnBeforePopup(fn.(ChromiumEventOnBeforePopup), getVal)
		//windowParent open url from tab
		case ChromiumEventOnOpenUrlFromTab:

		default:
			return false
		}
		return true
	})
}

func getInstance(value interface{}) (uintptr, unsafe.Pointer) {
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
	return ptr, unsafe.Pointer(ptr)
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
