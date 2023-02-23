//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// cef -> energy 结构类型定义
package cef

import (
	. "github.com/energye/energy/consts"
	. "github.com/energye/energy/types"
	"time"
	"unsafe"
)

// TCefCloseBrowsesAction 浏览器关闭控制
type TCefCloseBrowsesAction = CBS

// ICefCookie CEF Cookie
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

// TCefKeyEvent CEF 键盘事件
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

// TCefRequestContextSettings CEF 请求上下文配置
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

// TCefBrowserSettings CEF Browser配置
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

// TCefCommandLine 进程启动命令行参数设置
type TCefCommandLine struct {
	commandLines map[string]string
}

// TCefProxy 代理配置
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
}

// TCefTouchEvent 触摸事件
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

// TCustomHeader 自定义请求头
type TCustomHeader struct {
	CustomHeaderName  string
	CustomHeaderValue string
}

// TCefMouseEvent 鼠标事件
type TCefMouseEvent struct {
	X         int32
	Y         int32
	Modifiers TCefEventFlags
}

// BeforePopupInfo 弹出子窗口信息
type BeforePopupInfo struct {
	TargetUrl         string
	TargetFrameName   string
	TargetDisposition TCefWindowOpenDisposition
	UserGesture       bool
}

// TCefRect 矩形
type TCefRect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// TCefSize 大小
type TCefSize struct {
	Width  int32
	Height int32
}

// TCefPoint 位置
type TCefPoint struct {
	X int32
	Y int32
}

// TCefDraggableRegions 拖拽区域集合
type TCefDraggableRegions struct {
	regions      []TCefDraggableRegion
	regionsCount int
}

// TCefDraggableRegion 拖拽区域集
type TCefDraggableRegion struct {
	Bounds    TCefRect
	Draggable bool
}

// ICefDisplay
type ICefDisplay struct {
	instance unsafe.Pointer
}

// ICefWindow
type ICefWindow struct {
	instance unsafe.Pointer
}

// ICefView
type ICefView struct {
	instance unsafe.Pointer
}

// ICefClient
type ICefClient struct {
	instance unsafe.Pointer
}

// ICefDragData
type ICefDragData struct {
	instance unsafe.Pointer
}

// ICefV8Context
type ICefV8Context struct {
	instance unsafe.Pointer
	Browser  *ICefBrowser
	Frame    *ICefFrame
	Global   *ICefV8Value
}

// ICefV8Value
type ICefV8Value struct {
	instance unsafe.Pointer
}

// ICefV8Handler
type ICefV8Handler struct {
	instance unsafe.Pointer
}

//ICefV8Interceptor
type ICefV8Interceptor struct {
	instance unsafe.Pointer
}

//ICefV8Accessor
type ICefV8Accessor struct {
	instance unsafe.Pointer
}

//ICefV8ArrayBufferReleaseCallback
type ICefV8ArrayBufferReleaseCallback struct {
	instance unsafe.Pointer
}

// NewCefRect
func NewCefRect(x, y, width, height int32) *TCefRect {
	return &TCefRect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

// NewCefSize
func NewCefSize(width, height int32) *TCefSize {
	return &TCefSize{
		Width:  width,
		Height: height,
	}
}

// NewCefPoint
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

func (m *TCefKeyEvent) KeyDown() bool {
	return m.Kind == KEYEVENT_RAW_KEYDOWN || m.Kind == KEYEVENT_KEYDOWN
}

func (m *TCefKeyEvent) KeyUp() bool {
	return m.Kind == KEYEVENT_KEYUP
}
