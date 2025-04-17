//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefBrowserViewDelegate
// include/capi/views/cef_browser_view_delegate_capi.h (cef_browser_view_delegate_t)
type ICefBrowserViewDelegate struct {
	*ICefViewDelegate
}

// BrowserViewDelegateRef -> ICefBrowserViewDelegate
var BrowserViewDelegateRef browserViewDelegate

type browserViewDelegate uintptr

func (*browserViewDelegate) New() *ICefBrowserViewDelegate {
	var result uintptr
	imports.Proc(def.BrowserViewDelegateRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBrowserViewDelegate{&ICefViewDelegate{instance: unsafe.Pointer(result)}}
	}
	return nil
}

func (*browserViewDelegate) NewForCustom(browserViewDelegate *TCEFBrowserViewComponent) *ICefBrowserViewDelegate {
	if !browserViewDelegate.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.BrowserViewDelegateRef_CreateForCustom).Call(browserViewDelegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBrowserViewDelegate{&ICefViewDelegate{instance: unsafe.Pointer(result), ct: consts.CtOther}}
	}
	return nil
}

// Instance 实例
func (m *ICefBrowserViewDelegate) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefBrowserViewDelegate) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefBrowserViewDelegate) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefBrowserViewDelegate) SetOnBrowserCreated(fn browserViewOnBrowserCreated) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnBrowserCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnBrowserDestroyed(fn browserViewOnBrowserDestroyed) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnBrowserDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnGetDelegateForPopupBrowserView(fn browserViewOnGetDelegateForPopupBrowserView) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnGetDelegateForPopupBrowserView).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnPopupBrowserViewCreated(fn browserViewOnPopupBrowserViewCreated) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnPopupBrowserViewCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnGetChromeToolbarType(fn browserViewOnGetChromeToolbarType) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnGetChromeToolbarType).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnUseFramelessWindowForPictureInPicture(fn browserViewOnUseFramelessWindowForPictureInPicture) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnUseFramelessWindowForPictureInPicture).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnGestureCommand(fn browserViewOnGestureCommand) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnGestureCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnGetBrowserRuntimeStyle(fn browserViewOnGetBrowserRuntimeStyle) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnGetBrowserRuntimeStyle).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type browserViewOnBrowserCreated func(browserView *ICefBrowserView, browser *ICefBrowser)
type browserViewOnBrowserDestroyed func(browserView *ICefBrowserView, browser *ICefBrowser)
type browserViewOnGetDelegateForPopupBrowserView func(browserView *ICefBrowserView, browserSettings *TCefBrowserSettings, client *ICefClient, isDevtools bool) *ICefBrowserViewDelegate
type browserViewOnPopupBrowserViewCreated func(browserView, popupBrowserView *ICefBrowserView, isDevtools bool) bool
type browserViewOnGetChromeToolbarType func(browserView *ICefBrowserView, result *consts.TCefChromeToolbarType)
type browserViewOnUseFramelessWindowForPictureInPicture func(browserView *ICefBrowserView) bool
type browserViewOnGestureCommand func(browserView *ICefBrowserView, gestureCommand consts.TCefGestureCommand) bool
type browserViewOnGetBrowserRuntimeStyle func(result *consts.TCefRuntimeStyle)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		var getBrowserView = func(index int) *ICefBrowserView {
			return &ICefBrowserView{&ICefView{instance: getPtr(index)}}
		}
		switch fn.(type) {
		case browserViewOnBrowserCreated:
			browserView := getBrowserView(0)
			browser := &ICefBrowser{instance: getPtr(1)}
			fn.(browserViewOnBrowserCreated)(browserView, browser)
		case browserViewOnBrowserDestroyed:
			browserView := getBrowserView(0)
			browser := &ICefBrowser{instance: getPtr(1)}
			fn.(browserViewOnBrowserDestroyed)(browserView, browser)
		case browserViewOnGetDelegateForPopupBrowserView:
			browserView := getBrowserView(0)
			browserSettingsPtr := (*tCefBrowserSettingsPtr)(getPtr(1))
			browserSettings := browserSettingsPtr.convert()
			client := &ICefClient{instance: getPtr(2)}
			resultPtr := (*uintptr)(getPtr(4))
			result := fn.(browserViewOnGetDelegateForPopupBrowserView)(browserView, browserSettings, client, api.GoBool(getVal(3)))
			if result != nil {
				*resultPtr = result.Instance()
			}
		case browserViewOnPopupBrowserViewCreated:
			browserView := getBrowserView(0)
			popupBrowserView := &ICefBrowserView{&ICefView{instance: getPtr(1)}}
			result := (*bool)(getPtr(3))
			*result = fn.(browserViewOnPopupBrowserViewCreated)(browserView, popupBrowserView, api.GoBool(getVal(2)))
		case browserViewOnGetChromeToolbarType:
			browserView := getBrowserView(0)
			fn.(browserViewOnGetChromeToolbarType)(browserView, (*consts.TCefChromeToolbarType)(getPtr(0)))
		case browserViewOnUseFramelessWindowForPictureInPicture:
			browserView := getBrowserView(0)
			result := (*bool)(getPtr(1))
			*result = fn.(browserViewOnUseFramelessWindowForPictureInPicture)(browserView)
		case browserViewOnGestureCommand:
			browserView := getBrowserView(0)
			gestureCommand := consts.TCefGestureCommand(getVal(1))
			result := (*bool)(getPtr(2))
			*result = fn.(browserViewOnGestureCommand)(browserView, gestureCommand)
		case browserViewOnGetBrowserRuntimeStyle:
			result := (*consts.TCefRuntimeStyle)(getPtr(0))
			fn.(browserViewOnGetBrowserRuntimeStyle)(result)
		default:
			return false
		}
		return true
	})
}
