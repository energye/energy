//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import . "github.com/energye/energy/v2/api"

type tCefBrowserSettings struct {
	WindowlessFrameRate        uintptr //int32
	StandardFontFamily         uintptr //string
	FixedFontFamily            uintptr //string
	SerifFontFamily            uintptr //string
	SansSerifFontFamily        uintptr //string
	CursiveFontFamily          uintptr //string
	FantasyFontFamily          uintptr //string
	DefaultFontSize            uintptr //int32
	DefaultFixedFontSize       uintptr //int32
	MinimumFontSize            uintptr //int32
	MinimumLogicalFontSize     uintptr //int32
	DefaultEncoding            uintptr //string
	RemoteFonts                uintptr //TCefState
	Javascript                 uintptr //TCefState
	JavascriptCloseWindows     uintptr //TCefState
	JavascriptAccessClipboard  uintptr //TCefState
	JavascriptDomPaste         uintptr //TCefState
	ImageLoading               uintptr //TCefState
	ImageShrinkStandaloneToFit uintptr //TCefState
	TextAreaResize             uintptr //TCefState
	TabToLinks                 uintptr //TCefState
	LocalStorage               uintptr //TCefState
	Databases                  uintptr //TCefState
	Webgl                      uintptr //TCefState
	BackgroundColor            uintptr //TCefColor
	ChromeStatusBubble         uintptr //TCefState
	ChromeZoomBubble           uintptr //TCefState
}

func (m *TCefBrowserSettings) Pointer() *tCefBrowserSettings {
	if m == nil {
		return nil
	}
	return &tCefBrowserSettings{
		WindowlessFrameRate:        uintptr(unsafePointer(&m.WindowlessFrameRate)),
		StandardFontFamily:         PascalStr(m.StandardFontFamily),
		FixedFontFamily:            PascalStr(m.FixedFontFamily),
		SerifFontFamily:            PascalStr(m.SerifFontFamily),
		SansSerifFontFamily:        PascalStr(m.SansSerifFontFamily),
		CursiveFontFamily:          PascalStr(m.CursiveFontFamily),
		FantasyFontFamily:          PascalStr(m.FantasyFontFamily),
		DefaultFontSize:            uintptr(unsafePointer(&m.DefaultFontSize)),
		DefaultFixedFontSize:       uintptr(unsafePointer(&m.DefaultFixedFontSize)),
		MinimumFontSize:            uintptr(unsafePointer(&m.MinimumFontSize)),
		MinimumLogicalFontSize:     uintptr(unsafePointer(&m.MinimumLogicalFontSize)),
		DefaultEncoding:            PascalStr(m.DefaultEncoding),
		RemoteFonts:                uintptr(unsafePointer(&m.RemoteFonts)),
		Javascript:                 uintptr(unsafePointer(&m.Javascript)),
		JavascriptCloseWindows:     uintptr(unsafePointer(&m.JavascriptCloseWindows)),
		JavascriptAccessClipboard:  uintptr(unsafePointer(&m.JavascriptAccessClipboard)),
		JavascriptDomPaste:         uintptr(unsafePointer(&m.JavascriptDomPaste)),
		ImageLoading:               uintptr(unsafePointer(&m.ImageLoading)),
		ImageShrinkStandaloneToFit: uintptr(unsafePointer(&m.ImageShrinkStandaloneToFit)),
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

func (m *tCefBrowserSettings) Convert() *TCefBrowserSettings {
	if m == nil {
		return nil
	}
	return &TCefBrowserSettings{
		instance:                   m,
		WindowlessFrameRate:        *(*int32)(unsafePointer(m.WindowlessFrameRate)),
		StandardFontFamily:         GoStr(m.StandardFontFamily),
		FixedFontFamily:            GoStr(m.FixedFontFamily),
		SerifFontFamily:            GoStr(m.SerifFontFamily),
		SansSerifFontFamily:        GoStr(m.SansSerifFontFamily),
		CursiveFontFamily:          GoStr(m.CursiveFontFamily),
		FantasyFontFamily:          GoStr(m.FantasyFontFamily),
		DefaultFontSize:            *(*int32)(unsafePointer(m.DefaultFontSize)),
		DefaultFixedFontSize:       *(*int32)(unsafePointer(m.DefaultFixedFontSize)),
		MinimumFontSize:            *(*int32)(unsafePointer(m.MinimumFontSize)),
		MinimumLogicalFontSize:     *(*int32)(unsafePointer(m.MinimumLogicalFontSize)),
		DefaultEncoding:            GoStr(m.DefaultEncoding),
		RemoteFonts:                *(*TCefState)(unsafePointer(m.RemoteFonts)),
		Javascript:                 *(*TCefState)(unsafePointer(m.Javascript)),
		JavascriptCloseWindows:     *(*TCefState)(unsafePointer(m.JavascriptCloseWindows)),
		JavascriptAccessClipboard:  *(*TCefState)(unsafePointer(m.JavascriptAccessClipboard)),
		JavascriptDomPaste:         *(*TCefState)(unsafePointer(m.JavascriptDomPaste)),
		ImageLoading:               *(*TCefState)(unsafePointer(m.ImageLoading)),
		ImageShrinkStandaloneToFit: *(*TCefState)(unsafePointer(m.ImageShrinkStandaloneToFit)),
		TextAreaResize:             *(*TCefState)(unsafePointer(m.TextAreaResize)),
		TabToLinks:                 *(*TCefState)(unsafePointer(m.TabToLinks)),
		LocalStorage:               *(*TCefState)(unsafePointer(m.LocalStorage)),
		Databases:                  *(*TCefState)(unsafePointer(m.Databases)),
		Webgl:                      *(*TCefState)(unsafePointer(m.Webgl)),
		BackgroundColor:            *(*TCefColor)(unsafePointer(m.BackgroundColor)),
		ChromeStatusBubble:         *(*TCefState)(unsafePointer(m.ChromeStatusBubble)),
		ChromeZoomBubble:           *(*TCefState)(unsafePointer(m.ChromeZoomBubble)),
	}
}

func (m *TCefBrowserSettings) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	*(*int32)(unsafePointer(m.instance.WindowlessFrameRate)) = m.WindowlessFrameRate
	m.instance.StandardFontFamily = PascalStr(m.StandardFontFamily)
	m.instance.FixedFontFamily = PascalStr(m.FixedFontFamily)
	m.instance.SerifFontFamily = PascalStr(m.SerifFontFamily)
	m.instance.SansSerifFontFamily = PascalStr(m.SansSerifFontFamily)
	m.instance.CursiveFontFamily = PascalStr(m.CursiveFontFamily)
	m.instance.FantasyFontFamily = PascalStr(m.FantasyFontFamily)
	*(*int32)(unsafePointer(m.instance.DefaultFontSize)) = m.DefaultFontSize
	*(*int32)(unsafePointer(m.instance.DefaultFixedFontSize)) = m.DefaultFixedFontSize
	*(*int32)(unsafePointer(m.instance.MinimumFontSize)) = m.MinimumFontSize
	*(*int32)(unsafePointer(m.instance.MinimumLogicalFontSize)) = m.MinimumLogicalFontSize
	m.instance.DefaultEncoding = PascalStr(m.DefaultEncoding)
	*(*TCefState)(unsafePointer(m.instance.RemoteFonts)) = m.RemoteFonts
	*(*TCefState)(unsafePointer(m.instance.Javascript)) = m.Javascript
	*(*TCefState)(unsafePointer(m.instance.JavascriptCloseWindows)) = m.JavascriptCloseWindows
	*(*TCefState)(unsafePointer(m.instance.JavascriptAccessClipboard)) = m.JavascriptAccessClipboard
	*(*TCefState)(unsafePointer(m.instance.JavascriptDomPaste)) = m.JavascriptDomPaste
	*(*TCefState)(unsafePointer(m.instance.ImageLoading)) = m.ImageLoading
	*(*TCefState)(unsafePointer(m.instance.ImageShrinkStandaloneToFit)) = m.ImageShrinkStandaloneToFit
	*(*TCefState)(unsafePointer(m.instance.TextAreaResize)) = m.TextAreaResize
	*(*TCefState)(unsafePointer(m.instance.TabToLinks)) = m.TabToLinks
	*(*TCefState)(unsafePointer(m.instance.LocalStorage)) = m.LocalStorage
	*(*TCefState)(unsafePointer(m.instance.Databases)) = m.Databases
	*(*TCefState)(unsafePointer(m.instance.Webgl)) = m.Webgl
	*(*TCefColor)(unsafePointer(m.instance.BackgroundColor)) = m.BackgroundColor
	*(*TCefState)(unsafePointer(m.instance.ChromeStatusBubble)) = m.ChromeStatusBubble
	*(*TCefState)(unsafePointer(m.instance.ChromeZoomBubble)) = m.ChromeZoomBubble
}

type tCefCookie struct {
	Name       uintptr //string
	Value      uintptr //string
	Domain     uintptr //string
	Path       uintptr //string
	Secure     uintptr //int32
	Httponly   uintptr //int32
	Creation   uintptr //TCefBaseTime
	LastAccess uintptr //TCefBaseTime
	HasExpires uintptr //int32
	Expires    uintptr //TCefBaseTime
	SameSite   uintptr //TCefCookieSameSite
	Priority   uintptr //TCefCookiePriority
}

func (m *TCefCookie) Pointer() *tCefCookie {
	if m == nil {
		return nil
	}
	return &tCefCookie{
		Name:       PascalStr(m.Name),
		Value:      PascalStr(m.Value),
		Domain:     PascalStr(m.Domain),
		Path:       PascalStr(m.Path),
		Secure:     uintptr(unsafePointer(&m.Secure)),
		Httponly:   uintptr(unsafePointer(&m.Httponly)),
		Creation:   uintptr(unsafePointer(&m.Creation)),
		LastAccess: uintptr(unsafePointer(&m.LastAccess)),
		HasExpires: uintptr(unsafePointer(&m.HasExpires)),
		Expires:    uintptr(unsafePointer(&m.Expires)),
		SameSite:   uintptr(unsafePointer(&m.SameSite)),
		Priority:   uintptr(unsafePointer(&m.Priority)),
	}
}

func (m *tCefCookie) Convert() *TCefCookie {
	if m == nil {
		return nil
	}
	return &TCefCookie{
		instance:   m,
		Name:       GoStr(m.Name),
		Value:      GoStr(m.Value),
		Domain:     GoStr(m.Domain),
		Path:       GoStr(m.Path),
		Secure:     *(*int32)(unsafePointer(m.Secure)),
		Httponly:   *(*int32)(unsafePointer(m.Httponly)),
		Creation:   *(*TCefBaseTime)(unsafePointer(m.Creation)),
		LastAccess: *(*TCefBaseTime)(unsafePointer(m.LastAccess)),
		HasExpires: *(*int32)(unsafePointer(m.HasExpires)),
		Expires:    *(*TCefBaseTime)(unsafePointer(m.Expires)),
		SameSite:   *(*TCefCookieSameSite)(unsafePointer(m.SameSite)),
		Priority:   *(*TCefCookiePriority)(unsafePointer(m.Priority)),
	}
}

func (m *TCefCookie) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.Name = PascalStr(m.Name)
	m.instance.Value = PascalStr(m.Value)
	m.instance.Domain = PascalStr(m.Domain)
	m.instance.Path = PascalStr(m.Path)
	*(*int32)(unsafePointer(m.instance.Secure)) = m.Secure
	*(*int32)(unsafePointer(m.instance.Httponly)) = m.Httponly
	*(*TCefBaseTime)(unsafePointer(m.instance.Creation)) = m.Creation
	*(*TCefBaseTime)(unsafePointer(m.instance.LastAccess)) = m.LastAccess
	*(*int32)(unsafePointer(m.instance.HasExpires)) = m.HasExpires
	*(*TCefBaseTime)(unsafePointer(m.instance.Expires)) = m.Expires
	*(*TCefCookieSameSite)(unsafePointer(m.instance.SameSite)) = m.SameSite
	*(*TCefCookiePriority)(unsafePointer(m.instance.Priority)) = m.Priority
}

type tCefMediaSinkDeviceInfo struct {
	IpAddress uintptr //string
	Port      uintptr //int32
	ModelName uintptr //string
}

func (m *TCefMediaSinkDeviceInfo) Pointer() *tCefMediaSinkDeviceInfo {
	if m == nil {
		return nil
	}
	return &tCefMediaSinkDeviceInfo{
		IpAddress: PascalStr(m.IpAddress),
		Port:      uintptr(unsafePointer(&m.Port)),
		ModelName: PascalStr(m.ModelName),
	}
}

func (m *tCefMediaSinkDeviceInfo) Convert() *TCefMediaSinkDeviceInfo {
	if m == nil {
		return nil
	}
	return &TCefMediaSinkDeviceInfo{
		instance:  m,
		IpAddress: GoStr(m.IpAddress),
		Port:      *(*int32)(unsafePointer(m.Port)),
		ModelName: GoStr(m.ModelName),
	}
}

func (m *TCefMediaSinkDeviceInfo) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.IpAddress = PascalStr(m.IpAddress)
	*(*int32)(unsafePointer(m.instance.Port)) = m.Port
	m.instance.ModelName = PascalStr(m.ModelName)
}

type tCefPdfPrintSettings struct {
	Landscape           uintptr //int32
	PrintBackground     uintptr //int32
	Scale               uintptr //float64
	PaperWidth          uintptr //float64
	PaperHeight         uintptr //float64
	PreferCssPageSize   uintptr //int32
	MarginType          uintptr //TCefPdfPrintMarginType
	MarginTop           uintptr //float64
	MarginRight         uintptr //float64
	MarginBottom        uintptr //float64
	MarginLeft          uintptr //float64
	PageRanges          uintptr //string
	DisplayHeaderFooter uintptr //int32
	HeaderTemplate      uintptr //string
	FooterTemplate      uintptr //string
	GenerateTaggedPdf   uintptr //int32
}

func (m *TCefPdfPrintSettings) Pointer() *tCefPdfPrintSettings {
	if m == nil {
		return nil
	}
	return &tCefPdfPrintSettings{
		Landscape:           uintptr(unsafePointer(&m.Landscape)),
		PrintBackground:     uintptr(unsafePointer(&m.PrintBackground)),
		Scale:               uintptr(unsafePointer(&m.Scale)),
		PaperWidth:          uintptr(unsafePointer(&m.PaperWidth)),
		PaperHeight:         uintptr(unsafePointer(&m.PaperHeight)),
		PreferCssPageSize:   uintptr(unsafePointer(&m.PreferCssPageSize)),
		MarginType:          uintptr(unsafePointer(&m.MarginType)),
		MarginTop:           uintptr(unsafePointer(&m.MarginTop)),
		MarginRight:         uintptr(unsafePointer(&m.MarginRight)),
		MarginBottom:        uintptr(unsafePointer(&m.MarginBottom)),
		MarginLeft:          uintptr(unsafePointer(&m.MarginLeft)),
		PageRanges:          PascalStr(m.PageRanges),
		DisplayHeaderFooter: uintptr(unsafePointer(&m.DisplayHeaderFooter)),
		HeaderTemplate:      PascalStr(m.HeaderTemplate),
		FooterTemplate:      PascalStr(m.FooterTemplate),
		GenerateTaggedPdf:   uintptr(unsafePointer(&m.GenerateTaggedPdf)),
	}
}

func (m *tCefPdfPrintSettings) Convert() *TCefPdfPrintSettings {
	if m == nil {
		return nil
	}
	return &TCefPdfPrintSettings{
		instance:            m,
		Landscape:           *(*int32)(unsafePointer(m.Landscape)),
		PrintBackground:     *(*int32)(unsafePointer(m.PrintBackground)),
		Scale:               *(*float64)(unsafePointer(m.Scale)),
		PaperWidth:          *(*float64)(unsafePointer(m.PaperWidth)),
		PaperHeight:         *(*float64)(unsafePointer(m.PaperHeight)),
		PreferCssPageSize:   *(*int32)(unsafePointer(m.PreferCssPageSize)),
		MarginType:          *(*TCefPdfPrintMarginType)(unsafePointer(m.MarginType)),
		MarginTop:           *(*float64)(unsafePointer(m.MarginTop)),
		MarginRight:         *(*float64)(unsafePointer(m.MarginRight)),
		MarginBottom:        *(*float64)(unsafePointer(m.MarginBottom)),
		MarginLeft:          *(*float64)(unsafePointer(m.MarginLeft)),
		PageRanges:          GoStr(m.PageRanges),
		DisplayHeaderFooter: *(*int32)(unsafePointer(m.DisplayHeaderFooter)),
		HeaderTemplate:      GoStr(m.HeaderTemplate),
		FooterTemplate:      GoStr(m.FooterTemplate),
		GenerateTaggedPdf:   *(*int32)(unsafePointer(m.GenerateTaggedPdf)),
	}
}

func (m *TCefPdfPrintSettings) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	*(*int32)(unsafePointer(m.instance.Landscape)) = m.Landscape
	*(*int32)(unsafePointer(m.instance.PrintBackground)) = m.PrintBackground
	*(*float64)(unsafePointer(m.instance.Scale)) = m.Scale
	*(*float64)(unsafePointer(m.instance.PaperWidth)) = m.PaperWidth
	*(*float64)(unsafePointer(m.instance.PaperHeight)) = m.PaperHeight
	*(*int32)(unsafePointer(m.instance.PreferCssPageSize)) = m.PreferCssPageSize
	*(*TCefPdfPrintMarginType)(unsafePointer(m.instance.MarginType)) = m.MarginType
	*(*float64)(unsafePointer(m.instance.MarginTop)) = m.MarginTop
	*(*float64)(unsafePointer(m.instance.MarginRight)) = m.MarginRight
	*(*float64)(unsafePointer(m.instance.MarginBottom)) = m.MarginBottom
	*(*float64)(unsafePointer(m.instance.MarginLeft)) = m.MarginLeft
	m.instance.PageRanges = PascalStr(m.PageRanges)
	*(*int32)(unsafePointer(m.instance.DisplayHeaderFooter)) = m.DisplayHeaderFooter
	m.instance.HeaderTemplate = PascalStr(m.HeaderTemplate)
	m.instance.FooterTemplate = PascalStr(m.FooterTemplate)
	*(*int32)(unsafePointer(m.instance.GenerateTaggedPdf)) = m.GenerateTaggedPdf
}

type tCefRequestContextSettings struct {
	CachePath                        uintptr //string
	PersistSessionCookies            uintptr //int32
	PersistUserPreferences           uintptr //int32
	AcceptLanguageList               uintptr //string
	CookieableSchemesList            uintptr //string
	CookieableSchemesExcludeDefaults uintptr //int32
}

func (m *TCefRequestContextSettings) Pointer() *tCefRequestContextSettings {
	if m == nil {
		return nil
	}
	return &tCefRequestContextSettings{
		CachePath:                        PascalStr(m.CachePath),
		PersistSessionCookies:            uintptr(unsafePointer(&m.PersistSessionCookies)),
		PersistUserPreferences:           uintptr(unsafePointer(&m.PersistUserPreferences)),
		AcceptLanguageList:               PascalStr(m.AcceptLanguageList),
		CookieableSchemesList:            PascalStr(m.CookieableSchemesList),
		CookieableSchemesExcludeDefaults: uintptr(unsafePointer(&m.CookieableSchemesExcludeDefaults)),
	}
}

func (m *tCefRequestContextSettings) Convert() *TCefRequestContextSettings {
	if m == nil {
		return nil
	}
	return &TCefRequestContextSettings{
		instance:                         m,
		CachePath:                        GoStr(m.CachePath),
		PersistSessionCookies:            *(*int32)(unsafePointer(m.PersistSessionCookies)),
		PersistUserPreferences:           *(*int32)(unsafePointer(m.PersistUserPreferences)),
		AcceptLanguageList:               GoStr(m.AcceptLanguageList),
		CookieableSchemesList:            GoStr(m.CookieableSchemesList),
		CookieableSchemesExcludeDefaults: *(*int32)(unsafePointer(m.CookieableSchemesExcludeDefaults)),
	}
}

func (m *TCefRequestContextSettings) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.CachePath = PascalStr(m.CachePath)
	*(*int32)(unsafePointer(m.instance.PersistSessionCookies)) = m.PersistSessionCookies
	*(*int32)(unsafePointer(m.instance.PersistUserPreferences)) = m.PersistUserPreferences
	m.instance.AcceptLanguageList = PascalStr(m.AcceptLanguageList)
	m.instance.CookieableSchemesList = PascalStr(m.CookieableSchemesList)
	*(*int32)(unsafePointer(m.instance.CookieableSchemesExcludeDefaults)) = m.CookieableSchemesExcludeDefaults
}

type tCefSettings struct {
	NoSandbox                        uintptr //int32
	BrowserSubprocessPath            uintptr //string
	FrameworkDirPath                 uintptr //string
	MainBundlePath                   uintptr //string
	ChromeRuntime                    uintptr //int32
	MultiThreadedMessageLoop         uintptr //int32
	ExternalMessagePump              uintptr //int32
	WindowlessRenderingEnabled       uintptr //int32
	CommandLineArgsDisabled          uintptr //int32
	CachePath                        uintptr //string
	RootCachePath                    uintptr //string
	PersistSessionCookies            uintptr //int32
	PersistUserPreferences           uintptr //int32
	UserAgent                        uintptr //string
	UserAgentProduct                 uintptr //string
	Locale                           uintptr //string
	LogFile                          uintptr //string
	LogSeverity                      uintptr //TCefLogSeverity
	LogItems                         uintptr //TCefLogItems
	JavascriptFlags                  uintptr //string
	ResourcesDirPath                 uintptr //string
	LocalesDirPath                   uintptr //string
	PackLoadingDisabled              uintptr //int32
	RemoteDebuggingPort              uintptr //int32
	UncaughtExceptionStackSize       uintptr //int32
	BackgroundColor                  uintptr //TCefColor
	AcceptLanguageList               uintptr //string
	CookieableSchemesList            uintptr //string
	CookieableSchemesExcludeDefaults uintptr //int32
	ChromePolicyId                   uintptr //string
}

func (m *TCefSettings) Pointer() *tCefSettings {
	if m == nil {
		return nil
	}
	return &tCefSettings{
		NoSandbox:                        uintptr(unsafePointer(&m.NoSandbox)),
		BrowserSubprocessPath:            PascalStr(m.BrowserSubprocessPath),
		FrameworkDirPath:                 PascalStr(m.FrameworkDirPath),
		MainBundlePath:                   PascalStr(m.MainBundlePath),
		ChromeRuntime:                    uintptr(unsafePointer(&m.ChromeRuntime)),
		MultiThreadedMessageLoop:         uintptr(unsafePointer(&m.MultiThreadedMessageLoop)),
		ExternalMessagePump:              uintptr(unsafePointer(&m.ExternalMessagePump)),
		WindowlessRenderingEnabled:       uintptr(unsafePointer(&m.WindowlessRenderingEnabled)),
		CommandLineArgsDisabled:          uintptr(unsafePointer(&m.CommandLineArgsDisabled)),
		CachePath:                        PascalStr(m.CachePath),
		RootCachePath:                    PascalStr(m.RootCachePath),
		PersistSessionCookies:            uintptr(unsafePointer(&m.PersistSessionCookies)),
		PersistUserPreferences:           uintptr(unsafePointer(&m.PersistUserPreferences)),
		UserAgent:                        PascalStr(m.UserAgent),
		UserAgentProduct:                 PascalStr(m.UserAgentProduct),
		Locale:                           PascalStr(m.Locale),
		LogFile:                          PascalStr(m.LogFile),
		LogSeverity:                      uintptr(unsafePointer(&m.LogSeverity)),
		LogItems:                         uintptr(unsafePointer(&m.LogItems)),
		JavascriptFlags:                  PascalStr(m.JavascriptFlags),
		ResourcesDirPath:                 PascalStr(m.ResourcesDirPath),
		LocalesDirPath:                   PascalStr(m.LocalesDirPath),
		PackLoadingDisabled:              uintptr(unsafePointer(&m.PackLoadingDisabled)),
		RemoteDebuggingPort:              uintptr(unsafePointer(&m.RemoteDebuggingPort)),
		UncaughtExceptionStackSize:       uintptr(unsafePointer(&m.UncaughtExceptionStackSize)),
		BackgroundColor:                  uintptr(unsafePointer(&m.BackgroundColor)),
		AcceptLanguageList:               PascalStr(m.AcceptLanguageList),
		CookieableSchemesList:            PascalStr(m.CookieableSchemesList),
		CookieableSchemesExcludeDefaults: uintptr(unsafePointer(&m.CookieableSchemesExcludeDefaults)),
		ChromePolicyId:                   PascalStr(m.ChromePolicyId),
	}
}

func (m *tCefSettings) Convert() *TCefSettings {
	if m == nil {
		return nil
	}
	return &TCefSettings{
		instance:                         m,
		NoSandbox:                        *(*int32)(unsafePointer(m.NoSandbox)),
		BrowserSubprocessPath:            GoStr(m.BrowserSubprocessPath),
		FrameworkDirPath:                 GoStr(m.FrameworkDirPath),
		MainBundlePath:                   GoStr(m.MainBundlePath),
		ChromeRuntime:                    *(*int32)(unsafePointer(m.ChromeRuntime)),
		MultiThreadedMessageLoop:         *(*int32)(unsafePointer(m.MultiThreadedMessageLoop)),
		ExternalMessagePump:              *(*int32)(unsafePointer(m.ExternalMessagePump)),
		WindowlessRenderingEnabled:       *(*int32)(unsafePointer(m.WindowlessRenderingEnabled)),
		CommandLineArgsDisabled:          *(*int32)(unsafePointer(m.CommandLineArgsDisabled)),
		CachePath:                        GoStr(m.CachePath),
		RootCachePath:                    GoStr(m.RootCachePath),
		PersistSessionCookies:            *(*int32)(unsafePointer(m.PersistSessionCookies)),
		PersistUserPreferences:           *(*int32)(unsafePointer(m.PersistUserPreferences)),
		UserAgent:                        GoStr(m.UserAgent),
		UserAgentProduct:                 GoStr(m.UserAgentProduct),
		Locale:                           GoStr(m.Locale),
		LogFile:                          GoStr(m.LogFile),
		LogSeverity:                      *(*TCefLogSeverity)(unsafePointer(m.LogSeverity)),
		LogItems:                         *(*TCefLogItems)(unsafePointer(m.LogItems)),
		JavascriptFlags:                  GoStr(m.JavascriptFlags),
		ResourcesDirPath:                 GoStr(m.ResourcesDirPath),
		LocalesDirPath:                   GoStr(m.LocalesDirPath),
		PackLoadingDisabled:              *(*int32)(unsafePointer(m.PackLoadingDisabled)),
		RemoteDebuggingPort:              *(*int32)(unsafePointer(m.RemoteDebuggingPort)),
		UncaughtExceptionStackSize:       *(*int32)(unsafePointer(m.UncaughtExceptionStackSize)),
		BackgroundColor:                  *(*TCefColor)(unsafePointer(m.BackgroundColor)),
		AcceptLanguageList:               GoStr(m.AcceptLanguageList),
		CookieableSchemesList:            GoStr(m.CookieableSchemesList),
		CookieableSchemesExcludeDefaults: *(*int32)(unsafePointer(m.CookieableSchemesExcludeDefaults)),
		ChromePolicyId:                   GoStr(m.ChromePolicyId),
	}
}

func (m *TCefSettings) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	*(*int32)(unsafePointer(m.instance.NoSandbox)) = m.NoSandbox
	m.instance.BrowserSubprocessPath = PascalStr(m.BrowserSubprocessPath)
	m.instance.FrameworkDirPath = PascalStr(m.FrameworkDirPath)
	m.instance.MainBundlePath = PascalStr(m.MainBundlePath)
	*(*int32)(unsafePointer(m.instance.ChromeRuntime)) = m.ChromeRuntime
	*(*int32)(unsafePointer(m.instance.MultiThreadedMessageLoop)) = m.MultiThreadedMessageLoop
	*(*int32)(unsafePointer(m.instance.ExternalMessagePump)) = m.ExternalMessagePump
	*(*int32)(unsafePointer(m.instance.WindowlessRenderingEnabled)) = m.WindowlessRenderingEnabled
	*(*int32)(unsafePointer(m.instance.CommandLineArgsDisabled)) = m.CommandLineArgsDisabled
	m.instance.CachePath = PascalStr(m.CachePath)
	m.instance.RootCachePath = PascalStr(m.RootCachePath)
	*(*int32)(unsafePointer(m.instance.PersistSessionCookies)) = m.PersistSessionCookies
	*(*int32)(unsafePointer(m.instance.PersistUserPreferences)) = m.PersistUserPreferences
	m.instance.UserAgent = PascalStr(m.UserAgent)
	m.instance.UserAgentProduct = PascalStr(m.UserAgentProduct)
	m.instance.Locale = PascalStr(m.Locale)
	m.instance.LogFile = PascalStr(m.LogFile)
	*(*TCefLogSeverity)(unsafePointer(m.instance.LogSeverity)) = m.LogSeverity
	*(*TCefLogItems)(unsafePointer(m.instance.LogItems)) = m.LogItems
	m.instance.JavascriptFlags = PascalStr(m.JavascriptFlags)
	m.instance.ResourcesDirPath = PascalStr(m.ResourcesDirPath)
	m.instance.LocalesDirPath = PascalStr(m.LocalesDirPath)
	*(*int32)(unsafePointer(m.instance.PackLoadingDisabled)) = m.PackLoadingDisabled
	*(*int32)(unsafePointer(m.instance.RemoteDebuggingPort)) = m.RemoteDebuggingPort
	*(*int32)(unsafePointer(m.instance.UncaughtExceptionStackSize)) = m.UncaughtExceptionStackSize
	*(*TCefColor)(unsafePointer(m.instance.BackgroundColor)) = m.BackgroundColor
	m.instance.AcceptLanguageList = PascalStr(m.AcceptLanguageList)
	m.instance.CookieableSchemesList = PascalStr(m.CookieableSchemesList)
	*(*int32)(unsafePointer(m.instance.CookieableSchemesExcludeDefaults)) = m.CookieableSchemesExcludeDefaults
	m.instance.ChromePolicyId = PascalStr(m.ChromePolicyId)
}

type tCefUrlParts struct {
	Spec     uintptr //string
	Scheme   uintptr //string
	Username uintptr //string
	Password uintptr //string
	Host     uintptr //string
	Port     uintptr //string
	Origin   uintptr //string
	Path     uintptr //string
	Query    uintptr //string
	Fragment uintptr //string
}

func (m *TCefUrlParts) Pointer() *tCefUrlParts {
	if m == nil {
		return nil
	}
	return &tCefUrlParts{
		Spec:     PascalStr(m.Spec),
		Scheme:   PascalStr(m.Scheme),
		Username: PascalStr(m.Username),
		Password: PascalStr(m.Password),
		Host:     PascalStr(m.Host),
		Port:     PascalStr(m.Port),
		Origin:   PascalStr(m.Origin),
		Path:     PascalStr(m.Path),
		Query:    PascalStr(m.Query),
		Fragment: PascalStr(m.Fragment),
	}
}

func (m *tCefUrlParts) Convert() *TCefUrlParts {
	if m == nil {
		return nil
	}
	return &TCefUrlParts{
		instance: m,
		Spec:     GoStr(m.Spec),
		Scheme:   GoStr(m.Scheme),
		Username: GoStr(m.Username),
		Password: GoStr(m.Password),
		Host:     GoStr(m.Host),
		Port:     GoStr(m.Port),
		Origin:   GoStr(m.Origin),
		Path:     GoStr(m.Path),
		Query:    GoStr(m.Query),
		Fragment: GoStr(m.Fragment),
	}
}

func (m *TCefUrlParts) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.Spec = PascalStr(m.Spec)
	m.instance.Scheme = PascalStr(m.Scheme)
	m.instance.Username = PascalStr(m.Username)
	m.instance.Password = PascalStr(m.Password)
	m.instance.Host = PascalStr(m.Host)
	m.instance.Port = PascalStr(m.Port)
	m.instance.Origin = PascalStr(m.Origin)
	m.instance.Path = PascalStr(m.Path)
	m.instance.Query = PascalStr(m.Query)
	m.instance.Fragment = PascalStr(m.Fragment)
}

type tCookie struct {
	Name       uintptr //string
	Value      uintptr //string
	Domain     uintptr //string
	Path       uintptr //string
	Creation   uintptr //TDateTime
	LastAccess uintptr //TDateTime
	Expires    uintptr //TDateTime
	Secure     uintptr //bool
	Httponly   uintptr //bool
	HasExpires uintptr //bool
	SameSite   uintptr //TCefCookieSameSite
	Priority   uintptr //TCefCookiePriority
}

func (m *TCookie) Pointer() *tCookie {
	if m == nil {
		return nil
	}
	return &tCookie{
		Name:       PascalStr(m.Name),
		Value:      PascalStr(m.Value),
		Domain:     PascalStr(m.Domain),
		Path:       PascalStr(m.Path),
		Creation:   uintptr(unsafePointer(&m.Creation)),
		LastAccess: uintptr(unsafePointer(&m.LastAccess)),
		Expires:    uintptr(unsafePointer(&m.Expires)),
		Secure:     uintptr(unsafePointer(&m.Secure)),
		Httponly:   uintptr(unsafePointer(&m.Httponly)),
		HasExpires: uintptr(unsafePointer(&m.HasExpires)),
		SameSite:   uintptr(unsafePointer(&m.SameSite)),
		Priority:   uintptr(unsafePointer(&m.Priority)),
	}
}

func (m *tCookie) Convert() *TCookie {
	if m == nil {
		return nil
	}
	return &TCookie{
		instance:   m,
		Name:       GoStr(m.Name),
		Value:      GoStr(m.Value),
		Domain:     GoStr(m.Domain),
		Path:       GoStr(m.Path),
		Creation:   *(*TDateTime)(unsafePointer(m.Creation)),
		LastAccess: *(*TDateTime)(unsafePointer(m.LastAccess)),
		Expires:    *(*TDateTime)(unsafePointer(m.Expires)),
		Secure:     *(*bool)(unsafePointer(m.Secure)),
		Httponly:   *(*bool)(unsafePointer(m.Httponly)),
		HasExpires: *(*bool)(unsafePointer(m.HasExpires)),
		SameSite:   *(*TCefCookieSameSite)(unsafePointer(m.SameSite)),
		Priority:   *(*TCefCookiePriority)(unsafePointer(m.Priority)),
	}
}

func (m *TCookie) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.Name = PascalStr(m.Name)
	m.instance.Value = PascalStr(m.Value)
	m.instance.Domain = PascalStr(m.Domain)
	m.instance.Path = PascalStr(m.Path)
	*(*TDateTime)(unsafePointer(m.instance.Creation)) = m.Creation
	*(*TDateTime)(unsafePointer(m.instance.LastAccess)) = m.LastAccess
	*(*TDateTime)(unsafePointer(m.instance.Expires)) = m.Expires
	*(*bool)(unsafePointer(m.instance.Secure)) = m.Secure
	*(*bool)(unsafePointer(m.instance.Httponly)) = m.Httponly
	*(*bool)(unsafePointer(m.instance.HasExpires)) = m.HasExpires
	*(*TCefCookieSameSite)(unsafePointer(m.instance.SameSite)) = m.SameSite
	*(*TCefCookiePriority)(unsafePointer(m.instance.Priority)) = m.Priority
}

type tUrlParts struct {
	Spec     uintptr //string
	Scheme   uintptr //string
	Username uintptr //string
	Password uintptr //string
	Host     uintptr //string
	Port     uintptr //string
	Origin   uintptr //string
	Path     uintptr //string
	Query    uintptr //string
	Fragment uintptr //string
}

func (m *TUrlParts) Pointer() *tUrlParts {
	if m == nil {
		return nil
	}
	return &tUrlParts{
		Spec:     PascalStr(m.Spec),
		Scheme:   PascalStr(m.Scheme),
		Username: PascalStr(m.Username),
		Password: PascalStr(m.Password),
		Host:     PascalStr(m.Host),
		Port:     PascalStr(m.Port),
		Origin:   PascalStr(m.Origin),
		Path:     PascalStr(m.Path),
		Query:    PascalStr(m.Query),
		Fragment: PascalStr(m.Fragment),
	}
}

func (m *tUrlParts) Convert() *TUrlParts {
	if m == nil {
		return nil
	}
	return &TUrlParts{
		instance: m,
		Spec:     GoStr(m.Spec),
		Scheme:   GoStr(m.Scheme),
		Username: GoStr(m.Username),
		Password: GoStr(m.Password),
		Host:     GoStr(m.Host),
		Port:     GoStr(m.Port),
		Origin:   GoStr(m.Origin),
		Path:     GoStr(m.Path),
		Query:    GoStr(m.Query),
		Fragment: GoStr(m.Fragment),
	}
}

func (m *TUrlParts) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.Spec = PascalStr(m.Spec)
	m.instance.Scheme = PascalStr(m.Scheme)
	m.instance.Username = PascalStr(m.Username)
	m.instance.Password = PascalStr(m.Password)
	m.instance.Host = PascalStr(m.Host)
	m.instance.Port = PascalStr(m.Port)
	m.instance.Origin = PascalStr(m.Origin)
	m.instance.Path = PascalStr(m.Path)
	m.instance.Query = PascalStr(m.Query)
	m.instance.Fragment = PascalStr(m.Fragment)
}
