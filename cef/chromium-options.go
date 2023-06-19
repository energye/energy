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
	"github.com/energye/energy/v2/types"
)

type IChromiumOptions interface {
	Javascript() consts.TCefState
	JavascriptCloseWindows() consts.TCefState
	JavascriptAccessClipboard() consts.TCefState
	JavascriptDomPaste() consts.TCefState
	ImageLoading() consts.TCefState
	ImageShrinkStandaloneToFit() consts.TCefState
	TextAreaResize() consts.TCefState
	TabToLinks() consts.TCefState
	LocalStorage() consts.TCefState
	Databases() consts.TCefState
	Webgl() consts.TCefState
	BackgroundColor() types.TCefColor
	AcceptLanguageList() types.String
	WindowlessFrameRate() types.Integer
	ChromeStatusBubble() consts.TCefState
	SetJavascript(value consts.TCefState)
	SetJavascriptCloseWindows(value consts.TCefState)
	SetJavascriptAccessClipboard(value consts.TCefState)
	SetJavascriptDomPaste(value consts.TCefState)
	SetImageLoading(value consts.TCefState)
	SetImageShrinkStandaloneToFit(value consts.TCefState)
	SetTextAreaResize(value consts.TCefState)
	SetTabToLinks(value consts.TCefState)
	SetLocalStorage(value consts.TCefState)
	SetDatabases(value consts.TCefState)
	SetWebgl(value consts.TCefState)
	SetBackgroundColor(value types.TCefColor)
	SetAcceptLanguageList(value types.String)
	SetWindowlessFrameRate(value types.Integer)
	SetChromeStatusBubble(value consts.TCefState)
}

func NewChromiumOptions(chromium IChromium) IChromiumOptions {
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
		acceptLanguageList:         "",
		windowlessFrameRate:        consts.CEF_OSR_FRAMERATE_DEFAULT,
		chromeStatusBubble:         consts.STATE_DEFAULT,
		Chromium:                   chromium,
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
	imports.Proc(def.ChromiumOptions_SetJavascript).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetJavascriptCloseWindows(value consts.TCefState) {
	m.javascriptCloseWindows = value
	imports.Proc(def.ChromiumOptions_SetJavascriptCloseWindows).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetJavascriptAccessClipboard(value consts.TCefState) {
	m.javascriptAccessClipboard = value
	imports.Proc(def.ChromiumOptions_SetJavascriptAccessClipboard).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetJavascriptDomPaste(value consts.TCefState) {
	m.javascriptDomPaste = value
	imports.Proc(def.ChromiumOptions_SetJavascriptDomPaste).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetImageLoading(value consts.TCefState) {
	m.imageLoading = value
	imports.Proc(def.ChromiumOptions_SetImageLoading).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetImageShrinkStandaloneToFit(value consts.TCefState) {
	m.imageShrinkStandaloneToFit = value
	imports.Proc(def.ChromiumOptions_SetImageShrinkStandaloneToFit).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetTextAreaResize(value consts.TCefState) {
	m.textAreaResize = value
	imports.Proc(def.ChromiumOptions_SetTextAreaResize).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetTabToLinks(value consts.TCefState) {
	m.tabToLinks = value
	imports.Proc(def.ChromiumOptions_SetTabToLinks).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetLocalStorage(value consts.TCefState) {
	m.localStorage = value
	imports.Proc(def.ChromiumOptions_SetLocalStorage).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetDatabases(value consts.TCefState) {
	m.databases = value
	imports.Proc(def.ChromiumOptions_SetDatabases).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetWebgl(value consts.TCefState) {
	m.webgl = value
	imports.Proc(def.ChromiumOptions_SetWebgl).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetBackgroundColor(value types.TCefColor) {
	m.backgroundColor = value
	imports.Proc(def.ChromiumOptions_SetBackgroundColor).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetAcceptLanguageList(value types.String) {
	m.acceptLanguageList = value
	imports.Proc(def.ChromiumOptions_SetAcceptLanguageList).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetWindowlessFrameRate(value types.Integer) {
	m.windowlessFrameRate = value
	imports.Proc(def.ChromiumOptions_SetWindowlessFrameRate).Call(m.Chromium.Instance(), value.ToPtr())
}

func (m *TChromiumOptions) SetChromeStatusBubble(value consts.TCefState) {
	m.chromeStatusBubble = value
	imports.Proc(def.ChromiumOptions_SetChromeStatusBubble).Call(m.Chromium.Instance(), value.ToPtr())
}
