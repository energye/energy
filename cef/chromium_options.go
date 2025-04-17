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
	"github.com/cyber-xxm/energy/v2/types"
)

// The TChromiumOptions properties used to fill the TCefBrowserSettings record which is used during the browser creation.
type TChromiumOptions struct {
	chromium                   IChromium
	javascript                 consts.TCefState
	javascriptCloseWindows     consts.TCefState
	javascriptAccessClipboard  consts.TCefState
	javascriptDomPaste         consts.TCefState
	imageLoading               consts.TCefState
	imageShrinkStandaloneToFit consts.TCefState
	textAreaResize             consts.TCefState
	tabToLinks                 consts.TCefState
	localStorage               consts.TCefState
	databases                  consts.TCefState
	webgl                      consts.TCefState
	backgroundColor            types.TCefColor
	acceptLanguageList         types.String // TODO Remove CEF 118
	windowlessFrameRate        types.Integer
	chromeStatusBubble         consts.TCefState
}

func NewChromiumOptions(chromium IChromium) *TChromiumOptions {
	return &TChromiumOptions{
		javascript:                 consts.STATE_DEFAULT,
		javascriptCloseWindows:     consts.STATE_DEFAULT,
		javascriptAccessClipboard:  consts.STATE_DEFAULT,
		javascriptDomPaste:         consts.STATE_DEFAULT,
		imageLoading:               consts.STATE_DEFAULT,
		imageShrinkStandaloneToFit: consts.STATE_DEFAULT,
		textAreaResize:             consts.STATE_DEFAULT,
		tabToLinks:                 consts.STATE_DEFAULT,
		localStorage:               consts.STATE_DEFAULT,
		databases:                  consts.STATE_DEFAULT,
		webgl:                      consts.STATE_DEFAULT,
		backgroundColor:            0,
		acceptLanguageList:         "", // Remove CEF 118
		windowlessFrameRate:        consts.CEF_OSR_FRAMERATE_DEFAULT,
		chromeStatusBubble:         consts.STATE_DEFAULT,
		chromium:                   chromium,
	}
}

func (m *TChromiumOptions) Javascript() consts.TCefState {
	return m.javascript
}

func (m *TChromiumOptions) JavascriptCloseWindows() consts.TCefState {
	return m.javascriptCloseWindows
}

func (m *TChromiumOptions) JavascriptAccessClipboard() consts.TCefState {
	return m.javascriptAccessClipboard
}

func (m *TChromiumOptions) JavascriptDomPaste() consts.TCefState {
	return m.javascriptDomPaste
}

func (m *TChromiumOptions) ImageLoading() consts.TCefState {
	return m.imageLoading
}

func (m *TChromiumOptions) ImageShrinkStandaloneToFit() consts.TCefState {
	return m.imageShrinkStandaloneToFit
}

func (m *TChromiumOptions) TextAreaResize() consts.TCefState {
	return m.textAreaResize
}

func (m *TChromiumOptions) TabToLinks() consts.TCefState {
	return m.tabToLinks
}

func (m *TChromiumOptions) LocalStorage() consts.TCefState {
	return m.localStorage
}

func (m *TChromiumOptions) Databases() consts.TCefState {
	return m.databases
}

func (m *TChromiumOptions) Webgl() consts.TCefState {
	return m.webgl
}

func (m *TChromiumOptions) BackgroundColor() types.TCefColor {
	return m.backgroundColor
}

// AcceptLanguageList Remove CEF 118
func (m *TChromiumOptions) AcceptLanguageList() types.String {
	return m.acceptLanguageList
}

func (m *TChromiumOptions) WindowlessFrameRate() types.Integer {
	return m.windowlessFrameRate
}

func (m *TChromiumOptions) ChromeStatusBubble() consts.TCefState {
	return m.chromeStatusBubble
}

// setting

func (m *TChromiumOptions) SetJavascript(value consts.TCefState) {
	m.javascript = value
	imports.Proc(def.ChromiumOptions_SetJavascript).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetJavascriptCloseWindows(value consts.TCefState) {
	m.javascriptCloseWindows = value
	imports.Proc(def.ChromiumOptions_SetJavascriptCloseWindows).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetJavascriptAccessClipboard(value consts.TCefState) {
	m.javascriptAccessClipboard = value
	imports.Proc(def.ChromiumOptions_SetJavascriptAccessClipboard).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetJavascriptDomPaste(value consts.TCefState) {
	m.javascriptDomPaste = value
	imports.Proc(def.ChromiumOptions_SetJavascriptDomPaste).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetImageLoading(value consts.TCefState) {
	m.imageLoading = value
	imports.Proc(def.ChromiumOptions_SetImageLoading).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetImageShrinkStandaloneToFit(value consts.TCefState) {
	m.imageShrinkStandaloneToFit = value
	imports.Proc(def.ChromiumOptions_SetImageShrinkStandaloneToFit).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetTextAreaResize(value consts.TCefState) {
	m.textAreaResize = value
	imports.Proc(def.ChromiumOptions_SetTextAreaResize).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetTabToLinks(value consts.TCefState) {
	m.tabToLinks = value
	imports.Proc(def.ChromiumOptions_SetTabToLinks).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetLocalStorage(value consts.TCefState) {
	m.localStorage = value
	imports.Proc(def.ChromiumOptions_SetLocalStorage).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetDatabases(value consts.TCefState) {
	m.databases = value
	imports.Proc(def.ChromiumOptions_SetDatabases).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetWebgl(value consts.TCefState) {
	m.webgl = value
	imports.Proc(def.ChromiumOptions_SetWebgl).Call(m.chromium.Instance(), value.ToPtr())
}

// Background color used for the browser before a document is loaded and when
// no document color is specified. The alpha component must be either fully
// opaque (0xFF) or fully transparent (0x00). If the alpha component is fully
// opaque then the RGB components will be used as the background color. If
// the alpha component is fully transparent for a windowed browser then the
// TCefSettings.background_color value will be used. If the alpha component is
// fully transparent for a windowless (off-screen) browser then transparent
// painting will be enabled.
func (m *TChromiumOptions) SetBackgroundColor(value types.TCefColor) {
	m.backgroundColor = value
	imports.Proc(def.ChromiumOptions_SetBackgroundColor).Call(m.chromium.Instance(), value.ToPtr())
}

// SetAcceptLanguageList Remove CEF 118
func (m *TChromiumOptions) SetAcceptLanguageList(value types.String) {
	m.acceptLanguageList = value
	imports.Proc(def.ChromiumOptions_SetAcceptLanguageList).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetWindowlessFrameRate(value types.Integer) {
	m.windowlessFrameRate = value
	imports.Proc(def.ChromiumOptions_SetWindowlessFrameRate).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetChromeStatusBubble(value consts.TCefState) {
	m.chromeStatusBubble = value
	imports.Proc(def.ChromiumOptions_SetChromeStatusBubble).Call(m.chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) ChromeZoomBubble() consts.TCefState {
	r1, _, _ := imports.Proc(def.ChromiumOptions_ChromeZoomBubble).Call(consts.GetValue, m.chromium.Instance(), 0)
	return consts.TCefState(r1)
}

func (m *TChromiumOptions) SetChromeZoomBubble(value consts.TCefState) {
	imports.Proc(def.ChromiumOptions_ChromeZoomBubble).Call(consts.SetValue, m.chromium.Instance(), uintptr(value))
}
