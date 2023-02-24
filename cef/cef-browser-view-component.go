//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// VF 窗口组件，BrowserView
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
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
	r1, _, _ := imports.Proc(internale_CEFBrowserViewComponent_Create).Call(lcl.CheckPtr(AOwner))
	return &TCEFBrowserViewComponent{
		instance: unsafe.Pointer(r1),
	}
}

// CreateBrowserView
func (m *TCEFBrowserViewComponent) CreateBrowserView(client *ICefClient, url string, requestContextSettings *TCefRequestContextSettings, browserSettings *TCefBrowserSettings, extraInfo *DictionaryValue) {
	contextSettingsPtr := requestContextSettings.ToPtr()
	browserSettingsPtr := browserSettings.ToPtr()
	var dataBytes = []byte{}
	var dataBytesPtr unsafe.Pointer
	var dataBytesLen int = 0
	var argsLen int = 0
	if extraInfo != nil && extraInfo.dataLen > 0 {
		defer func() {
			extraInfo.Clear()
			extraInfo = nil
			dataBytes = nil
			dataBytesPtr = nil
		}()
		dataBytes = extraInfo.Package()
		argsLen = extraInfo.dataLen
		dataBytesPtr = unsafe.Pointer(&dataBytes[0])
		dataBytesLen = len(dataBytes) - 1
	} else {
		dataBytesPtr = unsafe.Pointer(&dataBytes)
	}
	imports.Proc(internale_CEFBrowserViewComponent_CreateBrowserView).Call(uintptr(m.instance), uintptr(client.instance), api.PascalStr(url), uintptr(unsafe.Pointer(&contextSettingsPtr)), uintptr(unsafe.Pointer(&browserSettingsPtr)), uintptr(argsLen), uintptr(dataBytesPtr), uintptr(dataBytesLen))
}

// Instance
func (m *TCEFBrowserViewComponent) Instance() uintptr {
	return uintptr(m.instance)
}

// GetForBrowser
func (m *TCEFBrowserViewComponent) GetForBrowser(browser *ICefBrowser) {
	imports.Proc(internale_CEFBrowserViewComponent_CreateBrowserView).Call(uintptr(m.instance), uintptr(browser.Identifier()))
}

// SetPreferAccelerators
func (m *TCEFBrowserViewComponent) SetPreferAccelerators(preferAccelerators bool) {
	imports.Proc(internale_CEFBrowserViewComponent_SetPreferAccelerators).Call(uintptr(m.instance), api.PascalBool(preferAccelerators))
}

// RequestFocus
func (m *TCEFBrowserViewComponent) RequestFocus() {
	imports.Proc(internale_CEFBrowserViewComponent_RequestFocus).Call(uintptr(m.instance))
}

// Browser
func (m *TCEFBrowserViewComponent) Browser() *ICefBrowser {
	r1, _, _ := imports.Proc(internale_CEFBrowserViewComponent_Browser).Call(uintptr(m.instance))
	browser := &ICefBrowser{
		browseId: int32(r1),
	}
	return browser
}

//func (m *TCEFBrowserViewComponent) BrowserView() {
// Proc(internale_CEFBrowserViewComponent_BrowserView).Call(uintptr(m.instance))
//}

// SetOnBrowserCreated
func (m *TCEFBrowserViewComponent) SetOnBrowserCreated(fn BrowserViewComponentOnBrowserCreated) {
	imports.Proc(internale_CEFBrowserViewComponent_SetOnBrowserCreated).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

// SetOnBrowserDestroyed
func (m *TCEFBrowserViewComponent) SetOnBrowserDestroyed(fn BrowserViewComponentOnBrowserDestroyed) {
	imports.Proc(internale_CEFBrowserViewComponent_SetOnBrowserDestroyed).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

// SetOnGetDelegateForPopupBrowserView
func (m *TCEFBrowserViewComponent) SetOnGetDelegateForPopupBrowserView(fn BrowserViewComponentOnGetDelegateForPopupBrowserView) {
	imports.Proc(internale_CEFBrowserViewComponent_SetOnGetDelegateForPopupBrowserView).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

// SetOnPopupBrowserViewCreated
func (m *TCEFBrowserViewComponent) SetOnPopupBrowserViewCreated(fn BrowserViewComponentOnPopupBrowserViewCreated) {
	imports.Proc(internale_CEFBrowserViewComponent_SetOnPopupBrowserViewCreated).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

// SetOnGetChromeToolbarType
func (m *TCEFBrowserViewComponent) SetOnGetChromeToolbarType(fn BrowserViewComponentOnGetChromeToolbarType) {
	imports.Proc(internale_CEFBrowserViewComponent_SetOnGetChromeToolbarType).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("TCEFBrowserViewComponent Error:", err)
			}
		}()
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case BrowserViewComponentOnBrowserCreated:
			browserView := &ICefBrowserView{
				instance: getPtr(1),
			}
			browser := &ICefBrowser{
				browseId: int32(getVal(2)),
			}
			fn.(BrowserViewComponentOnBrowserCreated)(lcl.AsObject(getPtr(0)), browserView, browser)
		case BrowserViewComponentOnBrowserDestroyed:
			browserView := &ICefBrowserView{
				instance: getPtr(1),
			}
			browser := &ICefBrowser{
				browseId: int32(getVal(2)),
			}
			fn.(BrowserViewComponentOnBrowserDestroyed)(lcl.AsObject(getPtr(0)), browserView, browser)
		case BrowserViewComponentOnGetDelegateForPopupBrowserView:
			browserView := &ICefBrowserView{
				instance: getPtr(1),
			}
			browserSettingsPtr := (*tCefBrowserSettingsPtr)(getPtr(2))
			browserSettings := browserSettingsPtr.Convert()
			client := &ICefClient{
				instance: getPtr(3),
			}
			result := &ICefBrowserViewDelegate{
				instance: getPtr(5),
			}
			fn.(BrowserViewComponentOnGetDelegateForPopupBrowserView)(lcl.AsObject(getPtr(0)), browserView, browserSettings, client, api.GoBool(getVal(4)), result)
		case BrowserViewComponentOnPopupBrowserViewCreated:
			browserView := &ICefBrowserView{
				instance: getPtr(1),
			}
			popupBrowserView := &ICefBrowserView{
				instance: getPtr(2),
			}
			fn.(BrowserViewComponentOnPopupBrowserViewCreated)(lcl.AsObject(getPtr(0)), browserView, popupBrowserView, api.GoBool(getVal(3)), (*bool)(getPtr(4)))
		case BrowserViewComponentOnGetChromeToolbarType:
			fn.(BrowserViewComponentOnGetChromeToolbarType)(lcl.AsObject(getPtr(0)), (*consts.TCefChromeToolbarType)(getPtr(1)))
		default:
			return false
		}
		return true
	})
}
