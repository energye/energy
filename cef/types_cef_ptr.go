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

// TCefPopupFeatures
type tCefPopupFeaturesPtr struct {
	X                  UIntptr // Integer
	XSet               UIntptr // Integer
	Y                  UIntptr // Integer
	YSet               UIntptr // Integer
	Width              UIntptr // Integer
	WidthSet           UIntptr // Integer
	Height             UIntptr // Integer
	HeightSet          UIntptr // Integer
	MenuBarVisible     UIntptr // Integer // ~ CEF 109
	StatusBarVisible   UIntptr // Integer // ~ CEF 109
	ToolBarVisible     UIntptr // Integer // ~ CEF 109
	LocationBarVisible UIntptr // Integer
	ScrollbarsVisible  UIntptr // Integer // ~ CEF 109
	IsPopup            UIntptr // Integer // CEF 110 ~ Current :True (1) if browser interface elements should be hidden.
	Resizable          UIntptr // Integer
	Fullscreen         UIntptr // Integer
	Dialog             UIntptr // Integer
	AdditionalFeatures UIntptr // TCefStringList // Use-CEF:[49]
}

// TCefBrowserSettings
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
	TargetUrl         UIntptr // string
	TargetFrameName   UIntptr // string
	TargetDisposition UIntptr // int32
	UserGesture       UIntptr // bool
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

// SetInstanceValue 实例指针设置值
func (m *TCefBrowserSettings) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	// 字段指针引用赋值, 如果是字符串类型需直接赋值
	m.instance.Size.SetValue(uint32(m.Size))                                            // NativeUInt
	m.instance.WindowlessFrameRate.SetValue(int32(m.WindowlessFrameRate))               // Integer
	m.instance.StandardFontFamily = UIntptr(m.StandardFontFamily.ToPtr())               // TCefString
	m.instance.FixedFontFamily = UIntptr(m.FixedFontFamily.ToPtr())                     // TCefString
	m.instance.SerifFontFamily = UIntptr(m.SerifFontFamily.ToPtr())                     // TCefString
	m.instance.SansSerifFontFamily = UIntptr(m.SansSerifFontFamily.ToPtr())             // TCefString
	m.instance.CursiveFontFamily = UIntptr(m.CursiveFontFamily.ToPtr())                 // TCefString
	m.instance.FantasyFontFamily = UIntptr(m.FantasyFontFamily.ToPtr())                 // TCefString
	m.instance.DefaultFontSize.SetValue(int32(m.DefaultFontSize))                       // Integer
	m.instance.DefaultFixedFontSize.SetValue(int32(m.DefaultFixedFontSize))             // Integer
	m.instance.MinimumFontSize.SetValue(int32(m.MinimumFontSize))                       // Integer
	m.instance.MinimumLogicalFontSize.SetValue(int32(m.MinimumLogicalFontSize))         // Integer
	m.instance.DefaultEncoding = UIntptr(m.DefaultEncoding.ToPtr())                     // TCefString
	m.instance.RemoteFonts.SetValue(int32(m.RemoteFonts))                               // TCefState
	m.instance.Javascript.SetValue(int32(m.Javascript))                                 // TCefState
	m.instance.JavascriptCloseWindows.SetValue(int32(m.JavascriptCloseWindows))         // TCefState
	m.instance.JavascriptAccessClipboard.SetValue(int32(m.JavascriptAccessClipboard))   // TCefState
	m.instance.JavascriptDomPaste.SetValue(int32(m.JavascriptDomPaste))                 // TCefState
	m.instance.ImageLoading.SetValue(int32(m.ImageLoading))                             // TCefState
	m.instance.ImageShrinkStandaLonetoFit.SetValue(int32(m.ImageShrinkStandaLonetoFit)) // TCefState
	m.instance.TextAreaResize.SetValue(int32(m.TextAreaResize))                         // TCefState
	m.instance.TabToLinks.SetValue(int32(m.TabToLinks))                                 // TCefState
	m.instance.LocalStorage.SetValue(int32(m.LocalStorage))                             // TCefState
	m.instance.Databases.SetValue(int32(m.Databases))                                   // TCefState
	m.instance.Webgl.SetValue(int32(m.Webgl))                                           // TCefState
	m.instance.BackgroundColor.SetValue(uint32(m.BackgroundColor))                      // TCefColor
	m.instance.AcceptLanguageList = UIntptr(m.AcceptLanguageList.ToPtr())               // TCefString Remove CEF 118
	m.instance.ChromeStatusBubble.SetValue(int32(m.ChromeStatusBubble))                 // TCefState
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
		if ptr == 0 {
			return 0
		}
		return *(*consts.TCefState)(getPtr(ptr))
	}
	getInteger := func(ptr uintptr) Integer {
		if ptr == 0 {
			return 0
		}
		return *(*Integer)(getPtr(ptr))
	}
	return &TCefBrowserSettings{
		instance:                   m,
		Size:                       *(*NativeUInt)(getPtr(m.Size.ToPtr())),
		WindowlessFrameRate:        getInteger(m.WindowlessFrameRate.ToPtr()),
		StandardFontFamily:         TCefString(api.GoStr(m.StandardFontFamily.ToPtr())),
		FixedFontFamily:            TCefString(api.GoStr(m.FixedFontFamily.ToPtr())),
		SerifFontFamily:            TCefString(api.GoStr(m.SerifFontFamily.ToPtr())),
		SansSerifFontFamily:        TCefString(api.GoStr(m.SansSerifFontFamily.ToPtr())),
		CursiveFontFamily:          TCefString(api.GoStr(m.CursiveFontFamily.ToPtr())),
		FantasyFontFamily:          TCefString(api.GoStr(m.FantasyFontFamily.ToPtr())),
		DefaultFontSize:            getInteger(m.DefaultFontSize.ToPtr()),
		DefaultFixedFontSize:       getInteger(m.DefaultFixedFontSize.ToPtr()),
		MinimumFontSize:            getInteger(m.MinimumFontSize.ToPtr()),
		MinimumLogicalFontSize:     getInteger(m.MinimumLogicalFontSize.ToPtr()),
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

// Convert 转换为结构
func (m *beforePopupInfoPtr) Convert() *BeforePopupInfo {
	return &BeforePopupInfo{
		TargetUrl:         api.GoStr(m.TargetUrl.ToPtr()),
		TargetFrameName:   api.GoStr(m.TargetFrameName.ToPtr()),
		TargetDisposition: consts.TCefWindowOpenDisposition(m.TargetDisposition),
		UserGesture:       api.GoBool(m.UserGesture.ToPtr()),
	}
}

// Convert 转换为结构
func (m *tCefPopupFeaturesPtr) Convert() *TCefPopupFeatures {
	getPtr := func(ptr uintptr) unsafe.Pointer {
		return unsafe.Pointer(ptr)
	}
	getInteger := func(ptr uintptr) Integer {
		if ptr == 0 {
			return 0
		}
		return *(*Integer)(getPtr(ptr))
	}
	getStringList := func(ptr uintptr) TCefStringList {
		if ptr == 0 {
			return 0
		}
		return *(*TCefStringList)(getPtr(ptr))
	}
	return &TCefPopupFeatures{
		X:                  getInteger(m.X.ToPtr()),
		XSet:               getInteger(m.XSet.ToPtr()),
		Y:                  getInteger(m.Y.ToPtr()),
		YSet:               getInteger(m.YSet.ToPtr()),
		Width:              getInteger(m.Width.ToPtr()),
		WidthSet:           getInteger(m.WidthSet.ToPtr()),
		Height:             getInteger(m.Height.ToPtr()),
		HeightSet:          getInteger(m.HeightSet.ToPtr()),
		MenuBarVisible:     getInteger(m.MenuBarVisible.ToPtr()),
		StatusBarVisible:   getInteger(m.StatusBarVisible.ToPtr()),
		ToolBarVisible:     getInteger(m.ToolBarVisible.ToPtr()),
		LocationBarVisible: getInteger(m.LocationBarVisible.ToPtr()),
		ScrollbarsVisible:  getInteger(m.ScrollbarsVisible.ToPtr()),
		IsPopup:            getInteger(m.IsPopup.ToPtr()),
		Resizable:          getInteger(m.Resizable.ToPtr()),
		Fullscreen:         getInteger(m.Fullscreen.ToPtr()),
		Dialog:             getInteger(m.Dialog.ToPtr()),
		AdditionalFeatures: getStringList(m.AdditionalFeatures.ToPtr()),
	}
}
