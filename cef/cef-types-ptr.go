//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// cef -> energy 结构指定类型定义
package cef

import (
	. "github.com/energye/energy/consts"
	. "github.com/energye/energy/types"
	"github.com/energye/golcl/lcl/api"
)

type iCefV8ContextPtr struct {
	Browse uintptr //ptr
	Frame  uintptr //ptr
	Global uintptr //ptr
}

type iCefCookiePtr struct {
	url, name, value, domain, path uintptr //string
	secure, httponly, hasExpires   uintptr //bool
	creation, lastAccess, expires  uintptr //float64
	count, total, aID              uintptr //int32
	sameSite                       uintptr //int32 TCefCookieSameSite
	priority                       uintptr //int32 TCefCookiePriority
	aSetImmediately                uintptr //bool
	aDeleteCookie                  uintptr //bool
	aResult                        uintptr //bool
}

type tCefRequestContextSettingsPtr struct {
	Size                             uintptr //uint32
	CachePath                        uintptr //TCefString
	PersistSessionCookies            uintptr //int32
	PersistUserPreferences           uintptr //int32
	AcceptLanguageList               uintptr //uint32
	CookieableSchemesList            uintptr //uint32
	CookieableSchemesExcludeDefaults uintptr //int32
}

type tCefBrowserSettingsPtr struct {
	Size                       uintptr //NativeUInt
	WindowlessFrameRate        uintptr //Integer
	StandardFontFamily         uintptr //TCefString
	FixedFontFamily            uintptr //TCefString
	SerifFontFamily            uintptr //TCefString
	SansSerifFontFamily        uintptr //TCefString
	CursiveFontFamily          uintptr //TCefString
	FantasyFontFamily          uintptr //TCefString
	DefaultFontSize            uintptr //Integer
	DefaultFixedFontSize       uintptr //Integer
	MinimumFontSize            uintptr //Integer
	MinimumLogicalFontSize     uintptr //Integer
	DefaultEncoding            uintptr //TCefString
	RemoteFonts                uintptr //TCefState
	Javascript                 uintptr //TCefState
	JavascriptCloseWindows     uintptr //TCefState
	JavascriptAccessClipboard  uintptr //TCefState
	JavascriptDomPaste         uintptr //TCefState
	ImageLoading               uintptr //TCefState
	ImageShrinkStandaLonetoFit uintptr //TCefState
	TextAreaResize             uintptr //TCefState
	TabToLinks                 uintptr //TCefState
	LocalStorage               uintptr //TCefState
	Databases                  uintptr //TCefState
	Webgl                      uintptr //TCefState
	BackgroundColor            uintptr //TCefColor
	AcceptLanguageList         uintptr //TCefString
	ChromeStatusBubble         uintptr //TCefState
}

// ToPtr 转换为指针
func (m *TCefBrowserSettings) ToPtr() *tCefBrowserSettingsPtr {
	return &tCefBrowserSettingsPtr{
		Size:                       m.Size.ToPtr(),
		WindowlessFrameRate:        m.WindowlessFrameRate.ToPtr(),
		StandardFontFamily:         m.StandardFontFamily.ToPtr(),
		FixedFontFamily:            m.FixedFontFamily.ToPtr(),
		SerifFontFamily:            m.SerifFontFamily.ToPtr(),
		SansSerifFontFamily:        m.SansSerifFontFamily.ToPtr(),
		CursiveFontFamily:          m.CursiveFontFamily.ToPtr(),
		FantasyFontFamily:          m.FantasyFontFamily.ToPtr(),
		DefaultFontSize:            m.DefaultFontSize.ToPtr(),
		DefaultFixedFontSize:       m.DefaultFixedFontSize.ToPtr(),
		MinimumFontSize:            m.MinimumFontSize.ToPtr(),
		MinimumLogicalFontSize:     m.MinimumLogicalFontSize.ToPtr(),
		DefaultEncoding:            m.DefaultEncoding.ToPtr(),
		RemoteFonts:                m.RemoteFonts.ToPtr(),
		Javascript:                 m.Javascript.ToPtr(),
		JavascriptCloseWindows:     m.JavascriptCloseWindows.ToPtr(),
		JavascriptAccessClipboard:  m.JavascriptAccessClipboard.ToPtr(),
		JavascriptDomPaste:         m.JavascriptDomPaste.ToPtr(),
		ImageLoading:               m.ImageLoading.ToPtr(),
		ImageShrinkStandaLonetoFit: m.ImageShrinkStandaLonetoFit.ToPtr(),
		TextAreaResize:             m.TextAreaResize.ToPtr(),
		TabToLinks:                 m.TabToLinks.ToPtr(),
		LocalStorage:               m.LocalStorage.ToPtr(),
		Databases:                  m.Databases.ToPtr(),
		Webgl:                      m.Webgl.ToPtr(),
		BackgroundColor:            m.BackgroundColor.ToPtr(),
		AcceptLanguageList:         m.AcceptLanguageList.ToPtr(),
		ChromeStatusBubble:         m.ChromeStatusBubble.ToPtr(),
	}
}

// Convert 转换为结构
func (m *tCefBrowserSettingsPtr) Convert() *TCefBrowserSettings {
	return &TCefBrowserSettings{
		Size:                       NativeUInt(m.Size),
		WindowlessFrameRate:        Integer(m.WindowlessFrameRate),
		StandardFontFamily:         TCefString(api.GoStr(m.StandardFontFamily)),
		FixedFontFamily:            TCefString(api.GoStr(m.FixedFontFamily)),
		SerifFontFamily:            TCefString(api.GoStr(m.SerifFontFamily)),
		SansSerifFontFamily:        TCefString(api.GoStr(m.SansSerifFontFamily)),
		CursiveFontFamily:          TCefString(api.GoStr(m.CursiveFontFamily)),
		FantasyFontFamily:          TCefString(api.GoStr(m.FantasyFontFamily)),
		DefaultFontSize:            Integer(m.DefaultFontSize),
		DefaultFixedFontSize:       Integer(m.DefaultFixedFontSize),
		MinimumFontSize:            Integer(m.MinimumFontSize),
		MinimumLogicalFontSize:     Integer(m.MinimumLogicalFontSize),
		DefaultEncoding:            TCefString(api.GoStr(m.DefaultEncoding)),
		RemoteFonts:                TCefState(m.RemoteFonts),
		Javascript:                 TCefState(m.Javascript),
		JavascriptCloseWindows:     TCefState(m.JavascriptCloseWindows),
		JavascriptAccessClipboard:  TCefState(m.JavascriptAccessClipboard),
		JavascriptDomPaste:         TCefState(m.JavascriptDomPaste),
		ImageLoading:               TCefState(m.ImageLoading),
		ImageShrinkStandaLonetoFit: TCefState(m.ImageShrinkStandaLonetoFit),
		TextAreaResize:             TCefState(m.TextAreaResize),
		TabToLinks:                 TCefState(m.TabToLinks),
		LocalStorage:               TCefState(m.LocalStorage),
		Databases:                  TCefState(m.Databases),
		Webgl:                      TCefState(m.Webgl),
		BackgroundColor:            TCefColor(m.BackgroundColor),
		AcceptLanguageList:         TCefString(api.GoStr(m.AcceptLanguageList)),
		ChromeStatusBubble:         TCefState(m.ChromeStatusBubble),
	}
}

type tCefProxyPtr struct {
	ProxyType              uintptr //TCefProxyType
	ProxyScheme            uintptr //TCefProxySchem
	ProxyServer            uintptr //string
	ProxyPort              uintptr //int32
	ProxyUsername          uintptr //string
	ProxyPassword          uintptr //string
	ProxyScriptURL         uintptr //string
	ProxyByPassList        uintptr //string
	MaxConnectionsPerProxy uintptr //int32
}

type beforePopupInfoPtr struct {
	TargetUrl         uintptr //string
	TargetFrameName   uintptr //string
	TargetDisposition uintptr //int32
	UserGesture       uintptr //bool
}

type tCefRectPtr struct {
	X      uintptr //int32
	Y      uintptr //int32
	Width  uintptr //int32
	Height uintptr //int32
}

type tCustomHeader struct {
	CustomHeaderName  uintptr //string
	CustomHeaderValue uintptr //string
}
