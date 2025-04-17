//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Browser View

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefBrowserView
// include/capi/views/cef_browser_view_capi.h (cef_browser_view_t)
type ICefBrowserView struct {
	*ICefView
}

// BrowserViewRef -> ICefBrowserView
var BrowserViewRef browserView

type browserView uintptr

func (*browserView) New(client *ICefClient, url string, browserSettings TCefBrowserSettings, extraInfo *ICefDictionaryValue,
	requestContext *ICefRequestContext, delegate *ICefBrowserViewDelegate) *ICefBrowserView {
	browserSettingsPtr := browserSettings.ToPtr()
	var result uintptr
	imports.Proc(def.CefBrowserViewRef_Create).Call(client.Instance(), api.PascalStr(url), uintptr(unsafe.Pointer(browserSettingsPtr)), extraInfo.Instance(), requestContext.Instance(), delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBrowserView{&ICefView{instance: unsafe.Pointer(result)}}
	}
	return nil
}

func (*browserView) GetForBrowser(browser *ICefBrowser) *ICefBrowserView {
	var result uintptr
	imports.Proc(def.CefBrowserViewRef_GetForBrowser).Call(browser.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBrowserView{&ICefView{instance: unsafe.Pointer(result)}}
	}
	return nil
}

func (m *ICefBrowserView) Browser() *ICefBrowser {
	var result uintptr
	imports.Proc(def.CefBrowserView_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBrowser{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefBrowserView) ChromeToolbar() *ICefView {
	var result uintptr
	imports.Proc(def.CefBrowserView_ChromeToolbar).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefBrowserView) SetPreferAccelerators(preferAccelerators bool) {
	imports.Proc(def.CefBrowserView_SetPreferAccelerators).Call(m.Instance(), api.PascalBool(preferAccelerators))
}

func (m *ICefBrowserView) RuntimeStyle() consts.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefBrowserView_RuntimeStyle).Call(m.Instance())
	return consts.TCefRuntimeStyle(r1)
}
