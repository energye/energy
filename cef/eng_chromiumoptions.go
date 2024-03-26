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

// IChromiumOptions Parent: IPersistent
//
//	The TChromiumOptions properties used to fill the TCefBrowserSettings record which is used during the browser creation.
type IChromiumOptions interface {
	IPersistent
	// Javascript
	//  Controls whether JavaScript can be executed. Also configurable using the
	//  "disable-javascript" command-line switch.
	Javascript() TCefState // property
	// SetJavascript Set Javascript
	SetJavascript(AValue TCefState) // property
	// JavascriptCloseWindows
	//  Controls whether JavaScript can be used to close windows that were not
	//  opened via JavaScript. JavaScript can still be used to close windows that
	//  were opened via JavaScript or that have no back/forward history. Also
	//  configurable using the "disable-javascript-close-windows" command-line
	//  switch.
	JavascriptCloseWindows() TCefState // property
	// SetJavascriptCloseWindows Set JavascriptCloseWindows
	SetJavascriptCloseWindows(AValue TCefState) // property
	// JavascriptAccessClipboard
	//  Controls whether JavaScript can access the clipboard. Also configurable
	//  using the "disable-javascript-access-clipboard" command-line switch.
	JavascriptAccessClipboard() TCefState // property
	// SetJavascriptAccessClipboard Set JavascriptAccessClipboard
	SetJavascriptAccessClipboard(AValue TCefState) // property
	// JavascriptDomPaste
	//  Controls whether DOM pasting is supported in the editor via
	//  execCommand("paste"). The |javascript_access_clipboard| setting must also
	//  be enabled. Also configurable using the "disable-javascript-dom-paste"
	//  command-line switch.
	JavascriptDomPaste() TCefState // property
	// SetJavascriptDomPaste Set JavascriptDomPaste
	SetJavascriptDomPaste(AValue TCefState) // property
	// ImageLoading
	//  Controls whether image URLs will be loaded from the network. A cached
	//  image will still be rendered if requested. Also configurable using the
	//  "disable-image-loading" command-line switch.
	ImageLoading() TCefState // property
	// SetImageLoading Set ImageLoading
	SetImageLoading(AValue TCefState) // property
	// ImageShrinkStandaloneToFit
	//  Controls whether standalone images will be shrunk to fit the page. Also
	//  configurable using the "image-shrink-standalone-to-fit" command-line
	//  switch.
	ImageShrinkStandaloneToFit() TCefState // property
	// SetImageShrinkStandaloneToFit Set ImageShrinkStandaloneToFit
	SetImageShrinkStandaloneToFit(AValue TCefState) // property
	// TextAreaResize
	//  Controls whether text areas can be resized. Also configurable using the
	//  "disable-text-area-resize" command-line switch.
	TextAreaResize() TCefState // property
	// SetTextAreaResize Set TextAreaResize
	SetTextAreaResize(AValue TCefState) // property
	// TabToLinks
	//  Controls whether the tab key can advance focus to links. Also configurable
	//  using the "disable-tab-to-links" command-line switch.
	TabToLinks() TCefState // property
	// SetTabToLinks Set TabToLinks
	SetTabToLinks(AValue TCefState) // property
	// LocalStorage
	//  Controls whether local storage can be used. Also configurable using the
	//  "disable-local-storage" command-line switch.
	LocalStorage() TCefState // property
	// SetLocalStorage Set LocalStorage
	SetLocalStorage(AValue TCefState) // property
	// Databases
	//  Controls whether databases can be used. Also configurable using the
	//  "disable-databases" command-line switch.
	Databases() TCefState // property
	// SetDatabases Set Databases
	SetDatabases(AValue TCefState) // property
	// Webgl
	//  Controls whether WebGL can be used. Note that WebGL requires hardware
	//  support and may not work on all systems even when enabled. Also
	//  configurable using the "disable-webgl" command-line switch.
	Webgl() TCefState // property
	// SetWebgl Set Webgl
	SetWebgl(AValue TCefState) // property
	// BackgroundColor
	//  Background color used for the browser before a document is loaded and when
	//  no document color is specified. The alpha component must be either fully
	//  opaque(0xFF) or fully transparent(0x00). If the alpha component is fully
	//  opaque then the RGB components will be used as the background color. If
	//  the alpha component is fully transparent for a windowed browser then the
	//  TCefSettings.background_color value will be used. If the alpha component is
	//  fully transparent for a windowless(off-screen) browser then transparent
	//  painting will be enabled.
	BackgroundColor() TCefColor // property
	// SetBackgroundColor Set BackgroundColor
	SetBackgroundColor(AValue TCefColor) // property
	// WindowlessFrameRate
	//  The maximum rate in frames per second(fps) that ICefRenderHandler.OnPaint
	//  will be called for a windowless browser. The actual fps may be lower if
	//  the browser cannot generate frames at the requested rate. The minimum
	//  value is 1 and the maximum value is 60(default 30). This value can also
	//  be changed dynamically via ICefBrowserHost.SetWindowlessFrameRate.
	//  Use CEF_OSR_SHARED_TEXTURES_FRAMERATE_DEFAULT as default value if the shared textures are enabled.
	//  Use CEF_OSR_FRAMERATE_DEFAULT as default value if the shared textures are disabled.
	WindowlessFrameRate() int32 // property
	// SetWindowlessFrameRate Set WindowlessFrameRate
	SetWindowlessFrameRate(AValue int32) // property
	// ChromeStatusBubble
	//  Controls whether the Chrome status bubble will be used. Only supported
	//  with the Chrome runtime. For details about the status bubble see
	//  https://www.chromium.org/user-experience/status-bubble/
	ChromeStatusBubble() TCefState // property
	// SetChromeStatusBubble Set ChromeStatusBubble
	SetChromeStatusBubble(AValue TCefState) // property
	// ChromeZoomBubble
	//  Controls whether the Chrome zoom bubble will be shown when zooming. Only
	//  supported with the Chrome runtime.
	ChromeZoomBubble() TCefState // property
	// SetChromeZoomBubble Set ChromeZoomBubble
	SetChromeZoomBubble(AValue TCefState) // property
}

// TChromiumOptions Parent: TPersistent
//
//	The TChromiumOptions properties used to fill the TCefBrowserSettings record which is used during the browser creation.
type TChromiumOptions struct {
	TPersistent
}

func NewChromiumOptions() IChromiumOptions {
	r1 := CEF().SysCallN(2100)
	return AsChromiumOptions(r1)
}

func (m *TChromiumOptions) Javascript() TCefState {
	r1 := CEF().SysCallN(2104, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascript(AValue TCefState) {
	CEF().SysCallN(2104, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) JavascriptCloseWindows() TCefState {
	r1 := CEF().SysCallN(2106, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascriptCloseWindows(AValue TCefState) {
	CEF().SysCallN(2106, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) JavascriptAccessClipboard() TCefState {
	r1 := CEF().SysCallN(2105, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascriptAccessClipboard(AValue TCefState) {
	CEF().SysCallN(2105, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) JavascriptDomPaste() TCefState {
	r1 := CEF().SysCallN(2107, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetJavascriptDomPaste(AValue TCefState) {
	CEF().SysCallN(2107, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ImageLoading() TCefState {
	r1 := CEF().SysCallN(2102, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetImageLoading(AValue TCefState) {
	CEF().SysCallN(2102, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ImageShrinkStandaloneToFit() TCefState {
	r1 := CEF().SysCallN(2103, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetImageShrinkStandaloneToFit(AValue TCefState) {
	CEF().SysCallN(2103, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) TextAreaResize() TCefState {
	r1 := CEF().SysCallN(2110, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetTextAreaResize(AValue TCefState) {
	CEF().SysCallN(2110, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) TabToLinks() TCefState {
	r1 := CEF().SysCallN(2109, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetTabToLinks(AValue TCefState) {
	CEF().SysCallN(2109, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) LocalStorage() TCefState {
	r1 := CEF().SysCallN(2108, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetLocalStorage(AValue TCefState) {
	CEF().SysCallN(2108, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) Databases() TCefState {
	r1 := CEF().SysCallN(2101, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetDatabases(AValue TCefState) {
	CEF().SysCallN(2101, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) Webgl() TCefState {
	r1 := CEF().SysCallN(2111, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetWebgl(AValue TCefState) {
	CEF().SysCallN(2111, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) BackgroundColor() TCefColor {
	r1 := CEF().SysCallN(2096, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TChromiumOptions) SetBackgroundColor(AValue TCefColor) {
	CEF().SysCallN(2096, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) WindowlessFrameRate() int32 {
	r1 := CEF().SysCallN(2112, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumOptions) SetWindowlessFrameRate(AValue int32) {
	CEF().SysCallN(2112, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ChromeStatusBubble() TCefState {
	r1 := CEF().SysCallN(2097, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetChromeStatusBubble(AValue TCefState) {
	CEF().SysCallN(2097, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumOptions) ChromeZoomBubble() TCefState {
	r1 := CEF().SysCallN(2098, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumOptions) SetChromeZoomBubble(AValue TCefState) {
	CEF().SysCallN(2098, 1, m.Instance(), uintptr(AValue))
}

func ChromiumOptionsClass() TClass {
	ret := CEF().SysCallN(2099)
	return TClass(ret)
}
