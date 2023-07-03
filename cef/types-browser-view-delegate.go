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
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

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

func (m *ICefBrowserViewDelegate) SetOnBrowserCreated(fn onBrowserCreated) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnBrowserCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnBrowserDestroyed(fn onBrowserDestroyed) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnBrowserDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnGetDelegateForPopupBrowserView(fn onGetDelegateForPopupBrowserView) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnGetDelegateForPopupBrowserView).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) SetOnPopupBrowserViewCreated(fn onPopupBrowserViewCreated) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.BrowserViewDelegate_SetOnPopupBrowserViewCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefBrowserViewDelegate) GetChromeToolbarType() consts.TCefChromeToolbarType {
	if !m.IsValid() || m.IsOtherEvent() {
		return 0
	}
	r1, _, _ := imports.Proc(def.BrowserViewDelegate_GetChromeToolbarType).Call(m.Instance())
	return consts.TCefChromeToolbarType(r1)
}

type onBrowserCreated func(browserView *ICefBrowserView, browser *ICefBrowser)
type onBrowserDestroyed func(browserView *ICefBrowserView, browser *ICefBrowser)
type onGetDelegateForPopupBrowserView func(browserView *ICefBrowserView, browserSettings *TCefBrowserSettings, client *ICefClient, isDevtools bool) *ICefBrowserViewDelegate
type onPopupBrowserViewCreated func(browserView, popupBrowserView *ICefBrowserView, isDevtools bool, aResult *bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onBrowserCreated:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(0)}}
			browser := &ICefBrowser{instance: getPtr(1)}
			fn.(onBrowserCreated)(browserView, browser)
		case onBrowserDestroyed:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(0)}}
			browser := &ICefBrowser{instance: getPtr(1)}
			fn.(onBrowserDestroyed)(browserView, browser)
		case onGetDelegateForPopupBrowserView:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(0)}}
			browserSettingsPtr := (*tCefBrowserSettingsPtr)(getPtr(1))
			browserSettings := browserSettingsPtr.Convert()
			client := &ICefClient{instance: getPtr(2)}
			resultPtr := (*uintptr)(getPtr(4))
			result := fn.(onGetDelegateForPopupBrowserView)(browserView, browserSettings, client, api.GoBool(getVal(3)))
			if result != nil {
				*resultPtr = result.Instance()
			}
		case onPopupBrowserViewCreated:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(0)}}
			popupBrowserView := &ICefBrowserView{&ICefView{instance: getPtr(1)}}
			fn.(onPopupBrowserViewCreated)(browserView, popupBrowserView, api.GoBool(getVal(2)), (*bool)(getPtr(3)))
		default:
			return false
		}
		return true
	})
}
