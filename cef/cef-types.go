//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"bytes"
	. "github.com/energye/energy/consts"
	. "github.com/energye/energy/types"
	"strings"
	"time"
	"unsafe"
)

type TCefCloseBrowsesAction = CBS

type ICefCookie struct {
	Url, Name, Value, Domain, Path string
	Secure, Httponly, HasExpires   bool
	Creation, LastAccess, Expires  time.Time
	Count, Total, ID               int32
	SameSite                       TCefCookieSameSite
	Priority                       TCefCookiePriority
	SetImmediately                 bool
	DeleteCookie                   bool
	Result                         bool
}

type TCefKeyEvent struct {
	Kind                 TCefKeyEventType // called 'type' in the original CEF source code
	Modifiers            TCefEventFlags
	WindowsKeyCode       Int32
	NativeKeyCode        Int32
	IsSystemKey          Int32
	Character            UInt16
	UnmodifiedCharacter  UInt16
	FocusOnEditableField Int32
}

type TCefRequestContextSettings struct {
	Size                             UInt32
	CachePath                        TCefString
	PersistSessionCookies            Int32
	PersistUserPreferences           Int32
	AcceptLanguageList               TCefString
	CookieableSchemesList            TCefString
	CookieableSchemesExcludeDefaults Int32
}

func (m *TCefRequestContextSettings) ToPtr() *tCefRequestContextSettingsPtr {
	return &tCefRequestContextSettingsPtr{
		Size:                             m.Size.ToPtr(),
		CachePath:                        m.CachePath.ToPtr(),
		PersistSessionCookies:            m.PersistSessionCookies.ToPtr(),
		PersistUserPreferences:           m.PersistUserPreferences.ToPtr(),
		AcceptLanguageList:               m.AcceptLanguageList.ToPtr(),
		CookieableSchemesList:            m.CookieableSchemesList.ToPtr(),
		CookieableSchemesExcludeDefaults: m.CookieableSchemesExcludeDefaults.ToPtr(),
	}
}

type TCefBrowserSettings struct {
	Size                       NativeUInt
	WindowlessFrameRate        Integer
	StandardFontFamily         TCefString
	FixedFontFamily            TCefString
	SerifFontFamily            TCefString
	SansSerifFontFamily        TCefString
	CursiveFontFamily          TCefString
	FantasyFontFamily          TCefString
	DefaultFontSize            Integer
	DefaultFixedFontSize       Integer
	MinimumFontSize            Integer
	MinimumLogicalFontSize     Integer
	DefaultEncoding            TCefString
	RemoteFonts                TCefState
	Javascript                 TCefState
	JavascriptCloseWindows     TCefState
	JavascriptAccessClipboard  TCefState
	JavascriptDomPaste         TCefState
	ImageLoading               TCefState
	ImageShrinkStandaLonetoFit TCefState
	TextAreaResize             TCefState
	TabToLinks                 TCefState
	LocalStorage               TCefState
	Databases                  TCefState
	Webgl                      TCefState
	BackgroundColor            TCefColor
	AcceptLanguageList         TCefString
	ChromeStatusBubble         TCefState
}

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

type TCefCommandLine struct {
	commandLines map[string]string
}

type TCefProxy struct {
	ProxyType              TCefProxyType
	ProxyScheme            TCefProxyScheme
	ProxyServer            string
	ProxyPort              int32
	ProxyUsername          string
	ProxyPassword          string
	ProxyScriptURL         string
	ProxyByPassList        string
	MaxConnectionsPerProxy int32
	CustomHeaderName       string
	CustomHeaderValue      string
}

type TCefTouchEvent struct {
	Id            int32
	X             float32
	Y             float32
	RadiusX       float32
	RadiusY       float32
	RotationAngle float32
	Pressure      float32
	Type          TCefTouchEeventType
	Modifiers     TCefEventFlags
	PointerType   TCefPointerType
}

type TCefMouseEvent struct {
	X         int32
	Y         int32
	Modifiers TCefEventFlags
}

type BeforePopupInfo struct {
	TargetUrl         string
	TargetFrameName   string
	TargetDisposition TCefWindowOpenDisposition
	UserGesture       bool
}

type TCefRect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type TCefSize struct {
	Width  int32
	Height int32
}

type TCefPoint struct {
	X int32
	Y int32
}

type TCefDraggableRegions struct {
	regions      []TCefDraggableRegion
	regionsCount int
}

type TCefDraggableRegion struct {
	Bounds    TCefRect
	Draggable bool
}

type ICefDisplay struct {
	instance unsafe.Pointer
}

type ICefWindow struct {
	instance unsafe.Pointer
}

type ICefView struct {
	instance unsafe.Pointer
}

type ICefClient struct {
	instance unsafe.Pointer
}

type ICefDragData struct {
	instance unsafe.Pointer
}

func NewCefDraggableRegion(rect *TCefRect, draggable bool) TCefDraggableRegion {
	return TCefDraggableRegion{
		Bounds:    *rect,
		Draggable: draggable,
	}
}

func NewCefDraggableRegions() *TCefDraggableRegions {
	return &TCefDraggableRegions{
		regions: make([]TCefDraggableRegion, 0),
	}
}

func (m *TCefDraggableRegions) Regions() []TCefDraggableRegion {
	if m.RegionsCount() == 0 || m.regions == nil || len(m.regions) == 0 {
		m.Append(NewCefDraggableRegion(NewCefRect(0, 0, 0, 0), false))
	}
	return m.regions
}

func (m *TCefDraggableRegions) Append(region TCefDraggableRegion) {
	m.regions = append(m.regions, region)
	m.regionsCount = len(m.regions)
}

func (m *TCefDraggableRegions) RegionsCount() int {
	return m.regionsCount
}

func NewCefRect(x, y, width, height int32) *TCefRect {
	return &TCefRect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func NewCefSize(width, height int32) *TCefSize {
	return &TCefSize{
		Width:  width,
		Height: height,
	}
}

func NewCefPoint(x, y int32) *TCefPoint {
	return &TCefPoint{
		X: x,
		Y: y,
	}
}

func (m *ICefWindow) SetWindow(window *ICefWindow) {
	m.instance = window.instance
}

func (m *ICefClient) SetClient(client *ICefClient) {
	m.instance = client.instance
}

func (m *TCefCommandLine) AppendSwitch(name, value string) {
	m.commandLines[name] = value
}

func (m *TCefCommandLine) AppendArgument(argument string) {
	m.commandLines[argument] = ""
}

func (m *TCefCommandLine) toString() string {
	var str bytes.Buffer
	var i = 0
	var replace = func(s, old, new string) string {
		return strings.ReplaceAll(s, old, new)
	}
	for name, value := range m.commandLines {
		if i > 0 {
			str.WriteString(" ")
		}
		if value != "" {
			str.WriteString(replace(replace(name, " ", ""), "=", ""))
			str.WriteString("=")
			str.WriteString(replace(replace(value, " ", ""), "=", ""))
		} else {
			str.WriteString(replace(name, " ", ""))
		}
		i++
	}
	return str.String()
}

func (m *TCefKeyEvent) KeyDown() bool {
	return m.Kind == KEYEVENT_RAW_KEYDOWN || m.Kind == KEYEVENT_KEYDOWN
}

func (m *TCefKeyEvent) KeyUp() bool {
	return m.Kind == KEYEVENT_KEYUP
}
