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
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	. "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
)

func fromPtrInt32(ptr uintptr) int32 {
	if ptr == 0 {
		return 0
	}
	return *(*int32)(unsafePointer(ptr))
}

func fromPtrUInt32(ptr uintptr) uint32 {
	if ptr == 0 {
		return 0
	}
	return *(*uint32)(unsafePointer(ptr))
}

func fromPtrBool(ptr uintptr) bool {
	if ptr == 0 {
		return false
	}
	return *(*bool)(unsafePointer(ptr))
}

func fromPtrStr(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	return api.GoStr(ptr)
}

func fromPtrFloat64(ptr uintptr) float64 {
	if ptr == 0 {
		return 0
	}
	return *(*float64)(unsafePointer(ptr))
}

func fromPtrFloat32(ptr uintptr) float32 {
	if ptr == 0 {
		return 0
	}
	return *(*float32)(unsafePointer(ptr))
}

func fromPtrCefState(ptr uintptr) consts.TCefState {
	if ptr == 0 {
		return 0
	}
	return *(*consts.TCefState)(unsafePointer(ptr))
}

func setCefStateVal(ptr uintptr, val consts.TCefState) {
	if ptr != 0 {
		*(*consts.TCefState)(unsafePointer(ptr)) = val
	}
}

func setInt32Val(ptr uintptr, val int32) {
	if ptr != 0 {
		*(*int32)(unsafePointer(ptr)) = val
	}
}

func setUInt32Val(ptr uintptr, val uint32) {
	if ptr != 0 {
		*(*uint32)(unsafePointer(ptr)) = val
	}
}

type tCefCookiePtr struct {
	url, name, value, domain, path        uintptr //string
	secure, httponly, hasExpires          uintptr //bool
	creation, lastAccess, expires         uintptr //float64
	count, total, aID, sameSite, priority uintptr //int32
	setImmediately                        uintptr //bool
}

func (m *tCefCookiePtr) convert() *TCefCookie {
	return &TCefCookie{
		Url:            fromPtrStr(m.url),
		Name:           fromPtrStr(m.name),
		Value:          fromPtrStr(m.value),
		Domain:         fromPtrStr(m.domain),
		Path:           fromPtrStr(m.path),
		Secure:         fromPtrBool(m.secure),
		Httponly:       fromPtrBool(m.httponly),
		HasExpires:     fromPtrBool(m.hasExpires),
		Creation:       common.DDateTimeToGoDateTime(fromPtrFloat64(m.creation)),
		LastAccess:     common.DDateTimeToGoDateTime(fromPtrFloat64(m.lastAccess)),
		Expires:        common.DDateTimeToGoDateTime(fromPtrFloat64(m.expires)),
		Count:          fromPtrInt32(m.count),
		Total:          fromPtrInt32(m.total),
		ID:             fromPtrInt32(m.aID),
		SameSite:       *(*consts.TCefCookieSameSite)(unsafePointer(m.sameSite)),
		Priority:       *(*consts.TCefCookiePriority)(unsafePointer(m.priority)),
		SetImmediately: fromPtrBool(m.setImmediately),
	}
}

func (m *TCefCookie) ToPtr() *tCefCookiePtr {
	creationPtr := common.GoDateTimeToDDateTime(m.Creation)
	lastAccessPtr := common.GoDateTimeToDDateTime(m.LastAccess)
	expiresPtr := common.GoDateTimeToDDateTime(m.Expires)
	return &tCefCookiePtr{
		url:            api.PascalStr(m.Url),
		name:           api.PascalStr(m.Name),
		value:          api.PascalStr(m.Value),
		domain:         api.PascalStr(m.Domain),
		path:           api.PascalStr(m.Path),
		secure:         uintptr(unsafePointer(&m.Secure)),
		httponly:       uintptr(unsafePointer(&m.Httponly)),
		hasExpires:     uintptr(unsafePointer(&m.HasExpires)),
		creation:       uintptr(unsafePointer(&creationPtr)),
		lastAccess:     uintptr(unsafePointer(&lastAccessPtr)),
		expires:        uintptr(unsafePointer(&expiresPtr)),
		sameSite:       uintptr(unsafePointer(&m.SameSite)),
		priority:       uintptr(unsafePointer(&m.Priority)),
		aID:            uintptr(unsafePointer(&m.ID)),
		count:          uintptr(unsafePointer(&m.Count)),
		total:          uintptr(unsafePointer(&m.Total)),
		setImmediately: uintptr(unsafePointer(&m.SetImmediately)),
	}
}

// ================

type tCefRequestContextSettingsPtr struct {
	CachePath                        uintptr //TCefString
	PersistSessionCookies            uintptr //Int32 // bool
	AcceptLanguageList               uintptr //TCefString  Remove CEF 118
	CookieableSchemesList            uintptr //TCefString
	CookieableSchemesExcludeDefaults uintptr //Int32
}

func (m *TCefRequestContextSettings) ToPtr() *tCefRequestContextSettingsPtr {
	return &tCefRequestContextSettingsPtr{
		CachePath:                        api.PascalStr(m.CachePath),
		PersistSessionCookies:            uintptr(unsafePointer(&m.PersistSessionCookies)),
		AcceptLanguageList:               api.PascalStr(m.AcceptLanguageList), // Remove CEF 118
		CookieableSchemesList:            api.PascalStr(m.CookieableSchemesList),
		CookieableSchemesExcludeDefaults: uintptr(unsafePointer(&m.CookieableSchemesExcludeDefaults)),
	}
}

// ================

type tCefPopupFeaturesPtr struct {
	X                  uintptr // Integer
	XSet               uintptr // Integer
	Y                  uintptr // Integer
	YSet               uintptr // Integer
	Width              uintptr // Integer
	WidthSet           uintptr // Integer
	Height             uintptr // Integer
	HeightSet          uintptr // Integer
	MenuBarVisible     uintptr // Integer // ~ CEF 109
	StatusBarVisible   uintptr // Integer // ~ CEF 109
	ToolBarVisible     uintptr // Integer // ~ CEF 109
	LocationBarVisible uintptr // Integer
	ScrollbarsVisible  uintptr // Integer // ~ CEF 109
	IsPopup            uintptr // Integer // CEF 110 ~ Current :True (1) if browser interface elements should be hidden.
	Resizable          uintptr // Integer
	Fullscreen         uintptr // Integer
	Dialog             uintptr // Integer
	AdditionalFeatures uintptr // TCefStringList // Use-CEF:[49]
}

func (m *tCefPopupFeaturesPtr) convert() *TCefPopupFeatures {
	getStringList := func(ptr uintptr) TCefStringList {
		if ptr == 0 {
			return 0
		}
		return *(*TCefStringList)(unsafePointer(ptr))
	}
	return &TCefPopupFeatures{
		X:                  fromPtrInt32(m.X),
		XSet:               fromPtrInt32(m.XSet),
		Y:                  fromPtrInt32(m.Y),
		YSet:               fromPtrInt32(m.YSet),
		Width:              fromPtrInt32(m.Width),
		WidthSet:           fromPtrInt32(m.WidthSet),
		Height:             fromPtrInt32(m.Height),
		HeightSet:          fromPtrInt32(m.HeightSet),
		MenuBarVisible:     fromPtrInt32(m.MenuBarVisible),
		StatusBarVisible:   fromPtrInt32(m.StatusBarVisible),
		ToolBarVisible:     fromPtrInt32(m.ToolBarVisible),
		LocationBarVisible: fromPtrInt32(m.LocationBarVisible),
		ScrollbarsVisible:  fromPtrInt32(m.ScrollbarsVisible),
		IsPopup:            fromPtrInt32(m.IsPopup),
		Resizable:          fromPtrInt32(m.Resizable),
		Fullscreen:         fromPtrInt32(m.Fullscreen),
		Dialog:             fromPtrInt32(m.Dialog),
		AdditionalFeatures: getStringList(m.AdditionalFeatures),
	}
}

// ================

type tCefRangePtr struct {
	From uintptr //int32
	To   uintptr //int32
}

func (m *TCefRange) ToPtr() *tCefRangePtr {
	return &tCefRangePtr{
		From: uintptr(unsafePointer(&m.From)),
		To:   uintptr(unsafePointer(&m.To)),
	}
}

// ================

type tCefCompositionUnderlinePtr struct {
	Range           uintptr //*TCefRange
	Color           uintptr //TCefColor
	BackgroundColor uintptr // TCefColor
	Thick           uintptr //int32
	Style           uintptr //TCefCompositionUnderlineStyle
}

func (m *TCefCompositionUnderline) ToPtr() *tCefCompositionUnderlinePtr {
	return &tCefCompositionUnderlinePtr{
		Range:           uintptr(unsafePointer(m.Range.ToPtr())),
		Color:           uintptr(unsafePointer(&m.Color)),
		BackgroundColor: uintptr(unsafePointer(&m.BackgroundColor)),
		Thick:           uintptr(unsafePointer(&m.Thick)),
		Style:           uintptr(unsafePointer(&m.Style)),
	}
}

type tCefSimulatedTouchPointPtr struct {
	Id                 uintptr //int32
	X                  uintptr //int32
	Y                  uintptr //int32
	RadiusX            uintptr //float32
	RadiusY            uintptr //float32
	RotationAngle      uintptr //float32
	Force              uintptr //float32
	TangentialPressure uintptr //float32
	TiltX              uintptr //int32
	TiltY              uintptr //int32
	Twist              uintptr //int32
}

func (m *TCefSimulatedTouchPoint) ToPtr() *tCefSimulatedTouchPointPtr {
	return &tCefSimulatedTouchPointPtr{
		Id:                 uintptr(m.Id),
		X:                  uintptr(m.X),
		Y:                  uintptr(m.Y),
		RadiusX:            uintptr(unsafePointer(&m.RadiusX)),
		RadiusY:            uintptr(unsafePointer(&m.RadiusY)),
		RotationAngle:      uintptr(unsafePointer(&m.RotationAngle)),
		Force:              uintptr(unsafePointer(&m.Force)),
		TangentialPressure: uintptr(unsafePointer(&m.TangentialPressure)),
		TiltX:              uintptr(m.TiltX),
		TiltY:              uintptr(m.TiltY),
		Twist:              uintptr(m.Twist),
	}
}

// ================

type tCefProxyPtr struct {
	ProxyType              uintptr //TCefProxyType
	ProxyScheme            uintptr //TCefProxyScheme
	ProxyServer            uintptr //string
	ProxyPort              uintptr //int32
	ProxyUsername          uintptr //string
	ProxyPassword          uintptr //string
	ProxyScriptURL         uintptr //string
	ProxyByPassList        uintptr //string
	MaxConnectionsPerProxy uintptr //int32
}

func (m *TCefProxy) ToPtr() *tCefProxyPtr {
	return &tCefProxyPtr{
		ProxyType:              uintptr(unsafePointer(&m.ProxyType)),
		ProxyScheme:            uintptr(unsafePointer(&m.ProxyScheme)),
		ProxyServer:            api.PascalStr(m.ProxyServer),
		ProxyPort:              uintptr(unsafePointer(&m.ProxyPort)),
		ProxyUsername:          api.PascalStr(m.ProxyUsername),
		ProxyPassword:          api.PascalStr(m.ProxyPassword),
		ProxyScriptURL:         api.PascalStr(m.ProxyScriptURL),
		ProxyByPassList:        api.PascalStr(m.ProxyByPassList),
		MaxConnectionsPerProxy: uintptr(unsafePointer(&m.MaxConnectionsPerProxy)),
	}
}

// ================

type beforePopupInfoPtr struct {
	TargetUrl         uintptr // string
	TargetFrameName   uintptr // string
	TargetDisposition uintptr // int32
	UserGesture       uintptr // bool
}

func (m *beforePopupInfoPtr) convert() *BeforePopupInfo {
	return &BeforePopupInfo{
		TargetUrl:         fromPtrStr(m.TargetUrl),
		TargetFrameName:   fromPtrStr(m.TargetFrameName),
		TargetDisposition: *(*consts.TCefWindowOpenDisposition)(unsafePointer(m.TargetDisposition)),
		UserGesture:       fromPtrBool(m.UserGesture),
	}
}

func (m *BeforePopupInfo) ToPtr() *beforePopupInfoPtr {
	return &beforePopupInfoPtr{
		TargetUrl:         api.PascalStr(m.TargetUrl),
		TargetFrameName:   api.PascalStr(m.TargetFrameName),
		TargetDisposition: uintptr(unsafePointer(&m.TargetDisposition)),
		UserGesture:       uintptr(unsafePointer(&m.UserGesture)),
	}
}

// ================

type tCustomHeader struct {
	CustomHeaderName  uintptr //string
	CustomHeaderValue uintptr //string
}

// ================

type tCefPdfPrintSettingsPtr struct {
	landscape               uintptr //Integer
	printBackground         uintptr //Integer
	scale                   uintptr //double
	paperWidth              uintptr //double
	paperHeight             uintptr //double
	preferCssPageSize       uintptr //Integer
	marginType              uintptr //TCefPdfPrintMarginType
	marginTop               uintptr //double
	marginRight             uintptr //double
	marginBottom            uintptr //double
	marginLeft              uintptr //double
	pageRanges              uintptr //TCefString
	displayHeaderFooter     uintptr //Integer
	headerTemplate          uintptr //TCefString
	footerTemplate          uintptr //TCefString
	generateTaggedPdf       uintptr // Integer
	generateDocumentOutline uintptr // Integer
}

func (m *tCefPdfPrintSettingsPtr) convert() *TCefPdfPrintSettings {
	return &TCefPdfPrintSettings{
		Landscape:               fromPtrInt32(m.landscape),
		PrintBackground:         fromPtrInt32(m.printBackground),
		Scale:                   fromPtrFloat64(m.scale),
		PaperWidth:              fromPtrFloat64(m.paperWidth),
		PaperHeight:             fromPtrFloat64(m.paperHeight),
		PreferCssPageSize:       fromPtrInt32(m.preferCssPageSize),
		MarginType:              *(*consts.TCefPdfPrintMarginType)(unsafePointer(m.marginType)),
		MarginTop:               fromPtrFloat64(m.marginTop),
		MarginRight:             fromPtrFloat64(m.marginRight),
		MarginBottom:            fromPtrFloat64(m.marginBottom),
		MarginLeft:              fromPtrFloat64(m.marginLeft),
		PageRanges:              fromPtrStr(m.pageRanges),
		DisplayHeaderFooter:     fromPtrInt32(m.displayHeaderFooter),
		HeaderTemplate:          fromPtrStr(m.headerTemplate),
		FooterTemplate:          fromPtrStr(m.footerTemplate),
		GenerateTaggedPdf:       fromPtrInt32(m.generateTaggedPdf),
		GenerateDocumentOutline: fromPtrInt32(m.generateDocumentOutline),
	}
}

func (m *TCefPdfPrintSettings) ToPtr() *tCefPdfPrintSettingsPtr {
	if m == nil {
		return nil
	}
	return &tCefPdfPrintSettingsPtr{
		landscape:               uintptr(unsafePointer(&m.Landscape)),
		printBackground:         uintptr(unsafePointer(&m.PrintBackground)),
		scale:                   uintptr(unsafePointer(&m.Scale)),
		paperWidth:              uintptr(unsafePointer(&m.PaperWidth)),
		paperHeight:             uintptr(unsafePointer(&m.PaperHeight)),
		preferCssPageSize:       uintptr(unsafePointer(&m.PreferCssPageSize)),
		marginType:              uintptr(unsafePointer(&m.MarginType)),
		marginTop:               uintptr(unsafePointer(&m.MarginTop)),
		marginRight:             uintptr(unsafePointer(&m.MarginRight)),
		marginBottom:            uintptr(unsafePointer(&m.MarginBottom)),
		marginLeft:              uintptr(unsafePointer(&m.MarginLeft)),
		pageRanges:              api.PascalStr(m.PageRanges),
		displayHeaderFooter:     uintptr(unsafePointer(&m.DisplayHeaderFooter)),
		headerTemplate:          api.PascalStr(m.HeaderTemplate),
		footerTemplate:          api.PascalStr(m.FooterTemplate),
		generateTaggedPdf:       uintptr(unsafePointer(&m.GenerateTaggedPdf)),
		generateDocumentOutline: uintptr(unsafePointer(&m.GenerateDocumentOutline)),
	}
}

// ================

type tCefBrowserSettingsPtr struct {
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
	ChromeStatusBubble         uintptr //TCefState
	ChromeZoomBubble           uintptr //TCefState
}

// SetInstanceValue 实例指针设置值
func (m *TCefBrowserSettings) setInstanceValue() {
	if m.instance == nil {
		return
	}
	setInt32Val(m.instance.WindowlessFrameRate, m.WindowlessFrameRate)
	m.instance.StandardFontFamily = api.PascalStr(m.StandardFontFamily)
	m.instance.FixedFontFamily = api.PascalStr(m.FixedFontFamily)
	m.instance.SerifFontFamily = api.PascalStr(m.SerifFontFamily)
	m.instance.SansSerifFontFamily = api.PascalStr(m.SansSerifFontFamily)
	m.instance.CursiveFontFamily = api.PascalStr(m.CursiveFontFamily)
	m.instance.FantasyFontFamily = api.PascalStr(m.FantasyFontFamily)
	setInt32Val(m.instance.DefaultFontSize, m.DefaultFontSize)
	setInt32Val(m.instance.DefaultFixedFontSize, m.DefaultFixedFontSize)
	setInt32Val(m.instance.MinimumFontSize, m.MinimumFontSize)
	setInt32Val(m.instance.MinimumLogicalFontSize, m.MinimumLogicalFontSize)
	m.instance.DefaultEncoding = api.PascalStr(m.DefaultEncoding)
	setCefStateVal(m.instance.RemoteFonts, m.RemoteFonts)
	setCefStateVal(m.instance.Javascript, m.Javascript)
	setCefStateVal(m.instance.JavascriptCloseWindows, m.JavascriptCloseWindows)
	setCefStateVal(m.instance.JavascriptAccessClipboard, m.JavascriptAccessClipboard)
	setCefStateVal(m.instance.JavascriptDomPaste, m.JavascriptDomPaste)
	setCefStateVal(m.instance.ImageLoading, m.ImageLoading)
	setCefStateVal(m.instance.ImageShrinkStandaLonetoFit, m.ImageShrinkStandaLonetoFit)
	setCefStateVal(m.instance.TextAreaResize, m.TextAreaResize)
	setCefStateVal(m.instance.TabToLinks, m.TabToLinks)
	setCefStateVal(m.instance.LocalStorage, m.LocalStorage)
	setCefStateVal(m.instance.Databases, m.Databases)
	setCefStateVal(m.instance.Webgl, m.Webgl)
	setUInt32Val(m.instance.BackgroundColor, uint32(m.BackgroundColor))
	setCefStateVal(m.instance.ChromeStatusBubble, m.ChromeStatusBubble)
	setCefStateVal(m.instance.ChromeZoomBubble, m.ChromeZoomBubble)
}

func (m *TCefBrowserSettings) ToPtr() *tCefBrowserSettingsPtr {
	if m == nil {
		return nil
	}
	return &tCefBrowserSettingsPtr{
		WindowlessFrameRate:        uintptr(unsafePointer(&m.WindowlessFrameRate)),
		StandardFontFamily:         api.PascalStr(m.StandardFontFamily),
		FixedFontFamily:            api.PascalStr(m.FixedFontFamily),
		SerifFontFamily:            api.PascalStr(m.SerifFontFamily),
		SansSerifFontFamily:        api.PascalStr(m.SansSerifFontFamily),
		CursiveFontFamily:          api.PascalStr(m.CursiveFontFamily),
		FantasyFontFamily:          api.PascalStr(m.FantasyFontFamily),
		DefaultFontSize:            uintptr(unsafePointer(&m.DefaultFontSize)),
		DefaultFixedFontSize:       uintptr(unsafePointer(&m.DefaultFixedFontSize)),
		MinimumFontSize:            uintptr(unsafePointer(&m.MinimumFontSize)),
		MinimumLogicalFontSize:     uintptr(unsafePointer(&m.MinimumLogicalFontSize)),
		DefaultEncoding:            api.PascalStr(m.DefaultEncoding),
		RemoteFonts:                uintptr(unsafePointer(&m.RemoteFonts)),
		Javascript:                 uintptr(unsafePointer(&m.Javascript)),
		JavascriptCloseWindows:     uintptr(unsafePointer(&m.JavascriptCloseWindows)),
		JavascriptAccessClipboard:  uintptr(unsafePointer(&m.JavascriptAccessClipboard)),
		JavascriptDomPaste:         uintptr(unsafePointer(&m.JavascriptDomPaste)),
		ImageLoading:               uintptr(unsafePointer(&m.ImageLoading)),
		ImageShrinkStandaLonetoFit: uintptr(unsafePointer(&m.ImageShrinkStandaLonetoFit)),
		TextAreaResize:             uintptr(unsafePointer(&m.TextAreaResize)),
		TabToLinks:                 uintptr(unsafePointer(&m.TabToLinks)),
		LocalStorage:               uintptr(unsafePointer(&m.LocalStorage)),
		Databases:                  uintptr(unsafePointer(&m.Databases)),
		Webgl:                      uintptr(unsafePointer(&m.Webgl)),
		BackgroundColor:            uintptr(unsafePointer(&m.BackgroundColor)),
		ChromeStatusBubble:         uintptr(unsafePointer(&m.ChromeStatusBubble)),
		ChromeZoomBubble:           uintptr(unsafePointer(&m.ChromeZoomBubble)),
	}
}

func (m *tCefBrowserSettingsPtr) convert() *TCefBrowserSettings {
	if m == nil {
		return nil
	}
	return &TCefBrowserSettings{
		instance:                   m,
		WindowlessFrameRate:        fromPtrInt32(m.WindowlessFrameRate),
		StandardFontFamily:         fromPtrStr(m.StandardFontFamily),
		FixedFontFamily:            fromPtrStr(m.FixedFontFamily),
		SerifFontFamily:            fromPtrStr(m.SerifFontFamily),
		SansSerifFontFamily:        fromPtrStr(m.SansSerifFontFamily),
		CursiveFontFamily:          fromPtrStr(m.CursiveFontFamily),
		FantasyFontFamily:          fromPtrStr(m.FantasyFontFamily),
		DefaultFontSize:            fromPtrInt32(m.DefaultFontSize),
		DefaultFixedFontSize:       fromPtrInt32(m.DefaultFixedFontSize),
		MinimumFontSize:            fromPtrInt32(m.MinimumFontSize),
		MinimumLogicalFontSize:     fromPtrInt32(m.MinimumLogicalFontSize),
		DefaultEncoding:            api.GoStr(m.DefaultEncoding),
		RemoteFonts:                fromPtrCefState(m.RemoteFonts),
		Javascript:                 fromPtrCefState(m.Javascript),
		JavascriptCloseWindows:     fromPtrCefState(m.JavascriptCloseWindows),
		JavascriptAccessClipboard:  fromPtrCefState(m.JavascriptAccessClipboard),
		JavascriptDomPaste:         fromPtrCefState(m.JavascriptDomPaste),
		ImageLoading:               fromPtrCefState(m.ImageLoading),
		ImageShrinkStandaLonetoFit: fromPtrCefState(m.ImageShrinkStandaLonetoFit),
		TextAreaResize:             fromPtrCefState(m.TextAreaResize),
		TabToLinks:                 fromPtrCefState(m.TabToLinks),
		LocalStorage:               fromPtrCefState(m.LocalStorage),
		Databases:                  fromPtrCefState(m.Databases),
		Webgl:                      fromPtrCefState(m.Webgl),
		BackgroundColor:            *(*TCefColor)(unsafePointer(m.BackgroundColor)),
		ChromeStatusBubble:         fromPtrCefState(m.ChromeStatusBubble),
		ChromeZoomBubble:           fromPtrCefState(m.ChromeZoomBubble),
	}
}

// ================

type tLinuxWindowPropertiesPtr struct {
	WaylandAppId uintptr
	WmClassClass uintptr
	WmClassName  uintptr
	WmRoleName   uintptr
}

func (m *TLinuxWindowProperties) setInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.WaylandAppId = api.PascalStr(m.WaylandAppId)
	m.instance.WmClassClass = api.PascalStr(m.WmClassClass)
	m.instance.WmClassName = api.PascalStr(m.WmClassName)
	m.instance.WmRoleName = api.PascalStr(m.WmRoleName)
}

func (m *tLinuxWindowPropertiesPtr) convert() *TLinuxWindowProperties {
	return &TLinuxWindowProperties{
		instance:     m,
		WaylandAppId: fromPtrStr(m.WaylandAppId),
		WmClassClass: fromPtrStr(m.WmClassClass),
		WmClassName:  fromPtrStr(m.WmClassName),
		WmRoleName:   fromPtrStr(m.WmRoleName),
	}
}

func (m *TLinuxWindowProperties) ToPtr() *tLinuxWindowPropertiesPtr {
	return &tLinuxWindowPropertiesPtr{
		WaylandAppId: api.PascalStr(m.WaylandAppId),
		WmClassClass: api.PascalStr(m.WmClassClass),
		WmClassName:  api.PascalStr(m.WmClassName),
		WmRoleName:   api.PascalStr(m.WmRoleName),
	}
}

// ================

type tCefInsetsPtr struct {
	Top    uintptr //int32
	Left   uintptr //int32
	Bottom uintptr //int32
	Right  uintptr //int32
}

func (m *tCefInsetsPtr) convert() *TCefInsets {
	return &TCefInsets{
		Top:    fromPtrInt32(m.Top),
		Left:   fromPtrInt32(m.Left),
		Bottom: fromPtrInt32(m.Bottom),
		Right:  fromPtrInt32(m.Right),
	}
}

func (m *TCefInsets) ToPtr() *tCefInsetsPtr {
	return &tCefInsetsPtr{
		Top:    uintptr(unsafePointer(&m.Top)),
		Left:   uintptr(unsafePointer(&m.Left)),
		Bottom: uintptr(unsafePointer(&m.Bottom)),
		Right:  uintptr(unsafePointer(&m.Right)),
	}
}

// ================

type tCefBoxLayoutSettingsPtr struct {
	Horizontal                    uintptr //Integer
	InsideBorderHorizontalSpacing uintptr //Integer
	InsideBorderVerticalSpacing   uintptr //Integer
	InsideBorderInsets            uintptr //tCefInsetsPtr
	BetweenChildSpacing           uintptr //Integer
	MainAxisAlignment             uintptr //consts.TCefMainAxisAlignment
	CrossAxisAlignment            uintptr //consts.TCefMainAxisAlignment
	MinimumCrossAxisSize          uintptr //Integer
	DefaultFlex                   uintptr //Integer
}

func (m *tCefBoxLayoutSettingsPtr) convert() *TCefBoxLayoutSettings {
	return &TCefBoxLayoutSettings{
		Horizontal:                    fromPtrInt32(m.Horizontal),
		InsideBorderHorizontalSpacing: fromPtrInt32(m.InsideBorderHorizontalSpacing),
		InsideBorderVerticalSpacing:   fromPtrInt32(m.InsideBorderVerticalSpacing),
		InsideBorderInsets:            *((*tCefInsetsPtr)(unsafePointer(m.InsideBorderInsets)).convert()),
		BetweenChildSpacing:           fromPtrInt32(m.BetweenChildSpacing),
		MainAxisAlignment:             *(*consts.TCefMainAxisAlignment)(unsafePointer(m.MainAxisAlignment)),
		CrossAxisAlignment:            *(*consts.TCefMainAxisAlignment)(unsafePointer(m.CrossAxisAlignment)),
		MinimumCrossAxisSize:          fromPtrInt32(m.MinimumCrossAxisSize),
		DefaultFlex:                   fromPtrInt32(m.DefaultFlex),
	}
}

func (m *TCefBoxLayoutSettings) ToPtr() *tCefBoxLayoutSettingsPtr {
	return &tCefBoxLayoutSettingsPtr{
		Horizontal:                    uintptr(unsafePointer(&m.Horizontal)),
		InsideBorderHorizontalSpacing: uintptr(unsafePointer(&m.InsideBorderHorizontalSpacing)),
		InsideBorderVerticalSpacing:   uintptr(unsafePointer(&m.InsideBorderVerticalSpacing)),
		InsideBorderInsets:            uintptr(unsafePointer(m.InsideBorderInsets.ToPtr())),
		BetweenChildSpacing:           uintptr(unsafePointer(&m.BetweenChildSpacing)),
		MainAxisAlignment:             uintptr(unsafePointer(&m.MainAxisAlignment)),
		CrossAxisAlignment:            uintptr(unsafePointer(&m.CrossAxisAlignment)),
		MinimumCrossAxisSize:          uintptr(unsafePointer(&m.MinimumCrossAxisSize)),
		DefaultFlex:                   uintptr(unsafePointer(&m.DefaultFlex)),
	}
}

// ================

type tCefTouchHandleStatePtr struct {
	TouchHandleId    uintptr //int32
	Flags            uintptr //uint32
	Enabled          uintptr //int32
	Orientation      uintptr //consts.TCefHorizontalAlignment
	MirrorVertical   uintptr //int32
	MirrorHorizontal uintptr //int32
	Origin           uintptr //TCefPoint
	Alpha            uintptr //float32
}

func (m *tCefTouchHandleStatePtr) convert() *TCefTouchHandleState {
	return &TCefTouchHandleState{
		TouchHandleId:    fromPtrInt32(m.TouchHandleId),
		Flags:            fromPtrUInt32(m.Flags),
		Enabled:          fromPtrInt32(m.Enabled),
		Orientation:      *(*consts.TCefHorizontalAlignment)(unsafePointer(m.Orientation)),
		MirrorVertical:   fromPtrInt32(m.MirrorVertical),
		MirrorHorizontal: fromPtrInt32(m.MirrorHorizontal),
		Origin:           *(*TCefPoint)(unsafePointer(m.Origin)),
		Alpha:            fromPtrFloat32(m.Alpha),
	}
}

// ================
