//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// VF 窗口组件 BrowserView

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCEFBrowserViewComponent
type TCEFBrowserViewComponent struct {
	lcl.IComponent
	instance unsafe.Pointer
}

// NewBrowserViewComponent
func NewBrowserViewComponent(AOwner lcl.IComponent) *TCEFBrowserViewComponent {
	var result uintptr
	imports.Proc(def.CEFBrowserViewComponent_Create).Call(lcl.CheckPtr(AOwner), uintptr(unsafe.Pointer(&result)))
	return &TCEFBrowserViewComponent{instance: unsafe.Pointer(result)}
}

// CreateBrowserView
func (m *TCEFBrowserViewComponent) CreateBrowserView(client *ICefClient, url string, requestContextSettings *TCefRequestContextSettings, browserSettings *TCefBrowserSettings, extraInfo *ICefDictionaryValue) {
	contextSettingsPtr := requestContextSettings.ToPtr()
	browserSettingsPtr := browserSettings.ToPtr()
	if extraInfo == nil {
		extraInfo = DictionaryValueRef.New()
	}
	imports.Proc(def.CEFBrowserViewComponent_CreateBrowserView).Call(m.Instance(), uintptr(client.instance), api.PascalStr(url), uintptr(unsafe.Pointer(&contextSettingsPtr)), uintptr(unsafe.Pointer(&browserSettingsPtr)), extraInfo.Instance())
}

// Instance
func (m *TCEFBrowserViewComponent) Instance() uintptr {
	return uintptr(m.instance)
}

// GetForBrowser
func (m *TCEFBrowserViewComponent) GetForBrowser(browser *ICefBrowser) {
	imports.Proc(def.CEFBrowserViewComponent_GetForBrowser).Call(m.Instance(), browser.Instance())
}

// SetPreferAccelerators
func (m *TCEFBrowserViewComponent) SetPreferAccelerators(preferAccelerators bool) {
	imports.Proc(def.CEFBrowserViewComponent_SetPreferAccelerators).Call(m.Instance(), api.PascalBool(preferAccelerators))
}

// RequestFocus
func (m *TCEFBrowserViewComponent) RequestFocus() {
	imports.Proc(def.CEFBrowserViewComponent_RequestFocus).Call(m.Instance())
}

// Browser
func (m *TCEFBrowserViewComponent) Browser() *ICefBrowser {
	var result uintptr
	imports.Proc(def.CEFBrowserViewComponent_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBrowser{instance: unsafe.Pointer(result)}
}

//func (m *TCEFBrowserViewComponent) BrowserView() {
// Proc(CEFBrowserViewComponent_BrowserView).Call(m.Instance())
//}

// SetOnBrowserCreated
func (m *TCEFBrowserViewComponent) SetOnBrowserCreated(fn BrowserViewComponentOnBrowserCreated) {
	imports.Proc(def.CEFBrowserViewComponent_SetOnBrowserCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnBrowserDestroyed
func (m *TCEFBrowserViewComponent) SetOnBrowserDestroyed(fn BrowserViewComponentOnBrowserDestroyed) {
	imports.Proc(def.CEFBrowserViewComponent_SetOnBrowserDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetDelegateForPopupBrowserView
func (m *TCEFBrowserViewComponent) SetOnGetDelegateForPopupBrowserView(fn BrowserViewComponentOnGetDelegateForPopupBrowserView) {
	imports.Proc(def.CEFBrowserViewComponent_SetOnGetDelegateForPopupBrowserView).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnPopupBrowserViewCreated
func (m *TCEFBrowserViewComponent) SetOnPopupBrowserViewCreated(fn BrowserViewComponentOnPopupBrowserViewCreated) {
	imports.Proc(def.CEFBrowserViewComponent_SetOnPopupBrowserViewCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetChromeToolbarType
func (m *TCEFBrowserViewComponent) SetOnGetChromeToolbarType(fn BrowserViewComponentOnGetChromeToolbarType) {
	imports.Proc(def.CEFBrowserViewComponent_SetOnGetChromeToolbarType).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case BrowserViewComponentOnBrowserCreated:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(1)}}
			browser := &ICefBrowser{instance: getPtr(2)}
			fn.(BrowserViewComponentOnBrowserCreated)(lcl.AsObject(getPtr(0)), browserView, browser)
		case BrowserViewComponentOnBrowserDestroyed:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(1)}}
			browser := &ICefBrowser{instance: getPtr(2)}
			fn.(BrowserViewComponentOnBrowserDestroyed)(lcl.AsObject(getPtr(0)), browserView, browser)
		case BrowserViewComponentOnGetDelegateForPopupBrowserView:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(1)}}
			browserSettingsPtr := (*tCefBrowserSettingsPtr)(getPtr(2))
			browserSettings := browserSettingsPtr.Convert()
			client := &ICefClient{instance: getPtr(3)}
			result := &ICefBrowserViewDelegate{instance: getPtr(5)}
			fn.(BrowserViewComponentOnGetDelegateForPopupBrowserView)(lcl.AsObject(getPtr(0)), browserView, browserSettings, client, api.GoBool(getVal(4)), result)
		case BrowserViewComponentOnPopupBrowserViewCreated:
			browserView := &ICefBrowserView{&ICefView{instance: getPtr(1)}}
			popupBrowserView := &ICefBrowserView{&ICefView{instance: getPtr(2)}}
			fn.(BrowserViewComponentOnPopupBrowserViewCreated)(lcl.AsObject(getPtr(0)), browserView, popupBrowserView, api.GoBool(getVal(3)), (*bool)(getPtr(4)))
		case BrowserViewComponentOnGetChromeToolbarType:
			fn.(BrowserViewComponentOnGetChromeToolbarType)(lcl.AsObject(getPtr(0)), (*consts.TCefChromeToolbarType)(getPtr(1)))
		default:
			return false
		}
		return true
	})
}
