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
)

// ICefBrowserView Parent: ICefView
//
//	A View hosting a ICefBrowser instance. Methods must be called on the
//	browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_browser_view_capi.h">CEF source file: /include/capi/views/cef_browser_view_capi.h (cef_browser_view_t)</a>
type ICefBrowserView interface {
	ICefView
	// GetBrowser
	//  Returns the ICefBrowser hosted by this BrowserView. Will return NULL if
	//  the browser has not yet been created or has already been destroyed.
	GetBrowser() ICefBrowser // function
	// GetChromeToolbar
	//  Returns the Chrome toolbar associated with this BrowserView. Only
	//  supported when using the Chrome runtime. The ICefBrowserViewDelegate.GetChromeToolbarType
	//  function must return a value other than
	//  CEF_CTT_NONE and the toolbar will not be available until after this
	//  BrowserView is added to a ICefWindow and
	//  ICefViewDelegate.OnWindowChanged() has been called.
	GetChromeToolbar() ICefView // function
	// SetPreferAccelerators
	//  Sets whether accelerators registered with ICefWindow.SetAccelerator are
	//  triggered before or after the event is sent to the ICefBrowser. If
	//  |prefer_accelerators| is true(1) then the matching accelerator will be
	//  triggered immediately and the event will not be sent to the ICefBrowser.
	//  If |prefer_accelerators| is false(0) then the matching accelerator will
	//  only be triggered if the event is not handled by web content or by
	//  ICefKeyboardHandler. The default value is false(0).
	SetPreferAccelerators(preferaccelerators bool) // procedure
}

// TCefBrowserView Parent: TCefView
//
//	A View hosting a ICefBrowser instance. Methods must be called on the
//	browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_browser_view_capi.h">CEF source file: /include/capi/views/cef_browser_view_capi.h (cef_browser_view_t)</a>
type TCefBrowserView struct {
	TCefView
}

// BrowserViewRef -> ICefBrowserView
var BrowserViewRef browserView

// browserView TCefBrowserView Ref
type browserView uintptr

// UnWrap
//
//	Returns a ICefBrowserView instance using a PCefBrowserView data pointer.
func (m *browserView) UnWrap(data uintptr) ICefBrowserView {
	var resultCefBrowserView uintptr
	CEF().SysCallN(680, uintptr(data), uintptr(unsafePointer(&resultCefBrowserView)))
	return AsCefBrowserView(resultCefBrowserView)
}

// CreateBrowserView
//
//	Create a new BrowserView. The underlying cef_browser_t will not be created
//	until this view is added to the views hierarchy. The optional |extra_info|
//	parameter provides an opportunity to specify extra information specific to
//	the created browser that will be passed to
//	cef_render_process_handler_t::on_browser_created() in the render process.
func (m *browserView) CreateBrowserView(client ICefClient, url string, settings *TCefBrowserSettings, extrainfo ICefDictionaryValue, requestcontext ICefRequestContext, delegate ICefBrowserViewDelegate) ICefBrowserView {
	inArgs2 := settings.Pointer()
	var resultCefBrowserView uintptr
	CEF().SysCallN(675, GetObjectUintptr(client), PascalStr(url), uintptr(unsafePointer(inArgs2)), GetObjectUintptr(extrainfo), GetObjectUintptr(requestcontext), GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefBrowserView)))
	return AsCefBrowserView(resultCefBrowserView)
}

// GetForBrowser
//
//	Returns the BrowserView associated with |browser|.
func (m *browserView) GetForBrowser(browser ICefBrowser) ICefBrowserView {
	var resultCefBrowserView uintptr
	CEF().SysCallN(678, GetObjectUintptr(browser), uintptr(unsafePointer(&resultCefBrowserView)))
	return AsCefBrowserView(resultCefBrowserView)
}

func (m *TCefBrowserView) GetBrowser() ICefBrowser {
	var resultCefBrowser uintptr
	CEF().SysCallN(676, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefBrowserView) GetChromeToolbar() ICefView {
	var resultCefView uintptr
	CEF().SysCallN(677, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefBrowserView) SetPreferAccelerators(preferaccelerators bool) {
	CEF().SysCallN(679, m.Instance(), PascalBool(preferaccelerators))
}
