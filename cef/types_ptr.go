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
	"github.com/energye/energy/v2/consts"
	. "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

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
	AcceptLanguageList               uintptr //uint32  Remove CEF 118
	CookieableSchemesList            uintptr //uint32
	CookieableSchemesExcludeDefaults uintptr //int32
}

type tCefBrowserSettingsPtr struct {
	Size                       UIntptr //NativeUInt
	WindowlessFrameRate        UIntptr //Integer
	StandardFontFamily         UIntptr //TCefString
	FixedFontFamily            UIntptr //TCefString
	SerifFontFamily            UIntptr //TCefString
	SansSerifFontFamily        UIntptr //TCefString
	CursiveFontFamily          UIntptr //TCefString
	FantasyFontFamily          UIntptr //TCefString
	DefaultFontSize            UIntptr //Integer
	DefaultFixedFontSize       UIntptr //Integer
	MinimumFontSize            UIntptr //Integer
	MinimumLogicalFontSize     UIntptr //Integer
	DefaultEncoding            UIntptr //TCefString
	RemoteFonts                UIntptr //TCefState
	Javascript                 UIntptr //TCefState
	JavascriptCloseWindows     UIntptr //TCefState
	JavascriptAccessClipboard  UIntptr //TCefState
	JavascriptDomPaste         UIntptr //TCefState
	ImageLoading               UIntptr //TCefState
	ImageShrinkStandaLonetoFit UIntptr //TCefState
	TextAreaResize             UIntptr //TCefState
	TabToLinks                 UIntptr //TCefState
	LocalStorage               UIntptr //TCefState
	Databases                  UIntptr //TCefState
	Webgl                      UIntptr //TCefState
	BackgroundColor            UIntptr //TCefColor
	AcceptLanguageList         UIntptr //TCefString Remove CEF 118
	ChromeStatusBubble         UIntptr //TCefState
}

type tCefCompositionUnderlinePtr struct {
	Range           uintptr //*TCefRange
	Color           uintptr //TCefColor
	BackgroundColor uintptr // TCefColor
	Thick           uintptr //int32
	Style           uintptr //TCefCompositionUnderlineStyle
}

// SetInstanceValue 为实例指针设置值
func (m *TCefBrowserSettings) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	// 字段指针引用赋值, 如果是字符串类型需直接赋值
	// TODO 需要全部修改
	m.instance.Size.SetValue(uint32(m.Size))
	m.instance.WindowlessFrameRate.SetValue(int32(m.WindowlessFrameRate))
	m.instance.StandardFontFamily = UIntptr(m.StandardFontFamily.ToPtr()) // 字符串赋值
	m.instance.FixedFontFamily = UIntptr(m.FixedFontFamily.ToPtr())
	//m.instance.SerifFontFamily = m.SerifFontFamily.ToPtr()
	//m.instance.SansSerifFontFamily = m.SansSerifFontFamily.ToPtr()
	//m.instance.CursiveFontFamily = m.CursiveFontFamily.ToPtr()
	//m.instance.FantasyFontFamily = m.FantasyFontFamily.ToPtr()
	//m.instance.DefaultFontSize = m.DefaultFontSize.ToPtr()
	//m.instance.DefaultFixedFontSize = m.DefaultFixedFontSize.ToPtr()
	//m.instance.MinimumFontSize = m.MinimumFontSize.ToPtr()
	//m.instance.MinimumLogicalFontSize = m.MinimumLogicalFontSize.ToPtr()
	//m.instance.DefaultEncoding = m.DefaultEncoding.ToPtr()
	//m.instance.RemoteFonts = m.RemoteFonts.ToPtr()
	//m.instance.Javascript = m.Javascript.ToPtr()
	//m.instance.JavascriptCloseWindows = m.JavascriptCloseWindows.ToPtr()
	//m.instance.JavascriptAccessClipboard = m.JavascriptAccessClipboard.ToPtr()
	//m.instance.JavascriptDomPaste = m.JavascriptDomPaste.ToPtr()
	//m.instance.ImageLoading = m.ImageLoading.ToPtr()
	//m.instance.ImageShrinkStandaLonetoFit = m.ImageShrinkStandaLonetoFit.ToPtr()
	//m.instance.TextAreaResize = m.TextAreaResize.ToPtr()
	//m.instance.TabToLinks = m.TabToLinks.ToPtr()
	//m.instance.LocalStorage = m.LocalStorage.ToPtr()
	//m.instance.Databases = m.Databases.ToPtr()
	//m.instance.Webgl = m.Webgl.ToPtr()
	//m.instance.BackgroundColor = m.BackgroundColor.ToPtr()
	//m.instance.AcceptLanguageList = m.AcceptLanguageList.ToPtr() // Remove CEF 118
	//m.instance.ChromeStatusBubble = m.ChromeStatusBubble.ToPtr()
}

// ToPtr 转换为指针
func (m *TCefBrowserSettings) ToPtr() *tCefBrowserSettingsPtr {
	if m == nil {
		return nil
	}
	return &tCefBrowserSettingsPtr{
		Size:                       UIntptr(m.Size.ToPtr()),
		WindowlessFrameRate:        UIntptr(m.WindowlessFrameRate.ToPtr()),
		StandardFontFamily:         UIntptr(m.StandardFontFamily.ToPtr()),
		FixedFontFamily:            UIntptr(m.FixedFontFamily.ToPtr()),
		SerifFontFamily:            UIntptr(m.SerifFontFamily.ToPtr()),
		SansSerifFontFamily:        UIntptr(m.SansSerifFontFamily.ToPtr()),
		CursiveFontFamily:          UIntptr(m.CursiveFontFamily.ToPtr()),
		FantasyFontFamily:          UIntptr(m.FantasyFontFamily.ToPtr()),
		DefaultFontSize:            UIntptr(m.DefaultFontSize.ToPtr()),
		DefaultFixedFontSize:       UIntptr(m.DefaultFixedFontSize.ToPtr()),
		MinimumFontSize:            UIntptr(m.MinimumFontSize.ToPtr()),
		MinimumLogicalFontSize:     UIntptr(m.MinimumLogicalFontSize.ToPtr()),
		DefaultEncoding:            UIntptr(m.DefaultEncoding.ToPtr()),
		RemoteFonts:                UIntptr(m.RemoteFonts.ToPtr()),
		Javascript:                 UIntptr(m.Javascript.ToPtr()),
		JavascriptCloseWindows:     UIntptr(m.JavascriptCloseWindows.ToPtr()),
		JavascriptAccessClipboard:  UIntptr(m.JavascriptAccessClipboard.ToPtr()),
		JavascriptDomPaste:         UIntptr(m.JavascriptDomPaste.ToPtr()),
		ImageLoading:               UIntptr(m.ImageLoading.ToPtr()),
		ImageShrinkStandaLonetoFit: UIntptr(m.ImageShrinkStandaLonetoFit.ToPtr()),
		TextAreaResize:             UIntptr(m.TextAreaResize.ToPtr()),
		TabToLinks:                 UIntptr(m.TabToLinks.ToPtr()),
		LocalStorage:               UIntptr(m.LocalStorage.ToPtr()),
		Databases:                  UIntptr(m.Databases.ToPtr()),
		Webgl:                      UIntptr(m.Webgl.ToPtr()),
		BackgroundColor:            UIntptr(m.BackgroundColor.ToPtr()),
		AcceptLanguageList:         UIntptr(m.AcceptLanguageList.ToPtr()), // Remove CEF 118
		ChromeStatusBubble:         UIntptr(m.ChromeStatusBubble.ToPtr()),
	}
}

// Convert 转换为结构
func (m *tCefBrowserSettingsPtr) Convert() *TCefBrowserSettings {
	getPtr := func(ptr uintptr) unsafe.Pointer {
		return unsafe.Pointer(ptr)
	}
	getCefState := func(ptr uintptr) consts.TCefState {
		// 可以确保字段不为空
		if ptr == 0 {
			return 0
		}
		return *(*consts.TCefState)(getPtr(ptr))
	}
	return &TCefBrowserSettings{
		instance:                   m,
		Size:                       *(*NativeUInt)(getPtr(m.Size.ToPtr())),
		WindowlessFrameRate:        *(*Integer)(getPtr(m.WindowlessFrameRate.ToPtr())),
		StandardFontFamily:         TCefString(api.GoStr(m.StandardFontFamily.ToPtr())),
		FixedFontFamily:            TCefString(api.GoStr(m.FixedFontFamily.ToPtr())),
		SerifFontFamily:            TCefString(api.GoStr(m.SerifFontFamily.ToPtr())),
		SansSerifFontFamily:        TCefString(api.GoStr(m.SansSerifFontFamily.ToPtr())),
		CursiveFontFamily:          TCefString(api.GoStr(m.CursiveFontFamily.ToPtr())),
		FantasyFontFamily:          TCefString(api.GoStr(m.FantasyFontFamily.ToPtr())),
		DefaultFontSize:            *(*Integer)(getPtr(m.DefaultFontSize.ToPtr())),
		DefaultFixedFontSize:       *(*Integer)(getPtr(m.DefaultFixedFontSize.ToPtr())),
		MinimumFontSize:            *(*Integer)(getPtr(m.MinimumFontSize.ToPtr())),
		MinimumLogicalFontSize:     *(*Integer)(getPtr(m.MinimumLogicalFontSize.ToPtr())),
		DefaultEncoding:            TCefString(api.GoStr(m.DefaultEncoding.ToPtr())),
		RemoteFonts:                getCefState(m.RemoteFonts.ToPtr()),
		Javascript:                 getCefState(m.Javascript.ToPtr()),
		JavascriptCloseWindows:     getCefState(m.JavascriptCloseWindows.ToPtr()),
		JavascriptAccessClipboard:  getCefState(m.JavascriptAccessClipboard.ToPtr()),
		JavascriptDomPaste:         getCefState(m.JavascriptDomPaste.ToPtr()),
		ImageLoading:               getCefState(m.ImageLoading.ToPtr()),
		ImageShrinkStandaLonetoFit: getCefState(m.ImageShrinkStandaLonetoFit.ToPtr()),
		TextAreaResize:             getCefState(m.TextAreaResize.ToPtr()),
		TabToLinks:                 getCefState(m.TabToLinks.ToPtr()),
		LocalStorage:               getCefState(m.LocalStorage.ToPtr()),
		Databases:                  getCefState(m.Databases.ToPtr()),
		Webgl:                      getCefState(m.Webgl.ToPtr()),
		BackgroundColor:            *(*TCefColor)(getPtr(m.BackgroundColor.ToPtr())),
		AcceptLanguageList:         TCefString(api.GoStr(m.AcceptLanguageList.ToPtr())), // Remove CEF 118
		ChromeStatusBubble:         getCefState(m.ChromeStatusBubble.ToPtr()),
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

type cefPdfPrintSettingsPtr struct {
	landscape           uintptr //Integer
	printBackground     uintptr //Integer
	scale               uintptr //double
	paperWidth          uintptr //double
	paperHeight         uintptr //double
	preferCssPageSize   uintptr //Integer
	marginType          uintptr //TCefPdfPrintMarginType
	marginTop           uintptr //double
	marginRight         uintptr //double
	marginBottom        uintptr //double
	marginLeft          uintptr //double
	pageRanges          uintptr //TCefString
	displayHeaderFooter uintptr //Integer
	headerTemplate      uintptr //TCefString
	footerTemplate      uintptr //TCefString
}

func (m *CefPdfPrintSettings) ToPtr() *cefPdfPrintSettingsPtr {
	if m == nil {
		return nil
	}
	return &cefPdfPrintSettingsPtr{
		landscape:           uintptr(m.Landscape),
		printBackground:     uintptr(m.PrintBackground),
		scale:               uintptr(unsafe.Pointer(&m.Scale)),
		paperWidth:          uintptr(unsafe.Pointer(&m.PaperWidth)),
		paperHeight:         uintptr(unsafe.Pointer(&m.PaperHeight)),
		preferCssPageSize:   uintptr(m.PreferCssPageSize),
		marginType:          uintptr(m.MarginType),
		marginTop:           uintptr(unsafe.Pointer(&m.MarginTop)), //m.MarginTop,
		marginRight:         uintptr(unsafe.Pointer(&m.MarginRight)),
		marginBottom:        uintptr(unsafe.Pointer(&m.MarginBottom)),
		marginLeft:          uintptr(unsafe.Pointer(&m.MarginLeft)),
		pageRanges:          api.PascalStr(m.PageRanges),
		displayHeaderFooter: uintptr(m.DisplayHeaderFooter),
		headerTemplate:      api.PascalStr(m.HeaderTemplate),
		footerTemplate:      api.PascalStr(m.FooterTemplate),
	}
}
