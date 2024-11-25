//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// cef -> energy 所有结构类型定义
// 每个结构对象创建 XXXRef.New() 创建并返回CEF对象, 创建后的对象是
// 引用CEF指针在不使用时,使用Free函数合理的释放掉该对象

package cef

import (
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	. "github.com/energye/energy/v2/types"
	"time"
	"unsafe"
)

type unsafePointer = unsafe.Pointer

// TCefCookie CEF Cookie
type TCefCookie struct {
	Url, Name, Value, Domain, Path string
	Secure, Httponly, HasExpires   bool
	Creation, LastAccess, Expires  time.Time
	Count, Total, ID               int32
	SameSite                       consts.TCefCookieSameSite
	Priority                       consts.TCefCookiePriority
	SetImmediately                 bool
}

// TCefKeyEvent CEF 键盘事件
type TCefKeyEvent struct {
	Kind                 consts.TCefKeyEventType // called 'type' in the original CEF source code
	Modifiers            consts.TCefEventFlags
	WindowsKeyCode       Int32
	NativeKeyCode        Int32
	IsSystemKey          Int32
	Character            UInt16
	UnmodifiedCharacter  UInt16
	FocusOnEditableField Int32
}

// TCefRequestContextSettings CEF 请求上下文配置
type TCefRequestContextSettings struct {
	CachePath                        string
	PersistSessionCookies            int32
	AcceptLanguageList               string // Remove CEF 118
	CookieableSchemesList            string
	CookieableSchemesExcludeDefaults int32
}

// TCefBrowserSettings CEF Browser配置
type TCefBrowserSettings struct {
	instance                   *tCefBrowserSettingsPtr
	WindowlessFrameRate        int32
	StandardFontFamily         string
	FixedFontFamily            string
	SerifFontFamily            string
	SansSerifFontFamily        string
	CursiveFontFamily          string
	FantasyFontFamily          string
	DefaultFontSize            int32
	DefaultFixedFontSize       int32
	MinimumFontSize            int32
	MinimumLogicalFontSize     int32
	DefaultEncoding            string
	RemoteFonts                consts.TCefState
	Javascript                 consts.TCefState
	JavascriptCloseWindows     consts.TCefState
	JavascriptAccessClipboard  consts.TCefState
	JavascriptDomPaste         consts.TCefState
	ImageLoading               consts.TCefState
	ImageShrinkStandaLonetoFit consts.TCefState
	TextAreaResize             consts.TCefState
	TabToLinks                 consts.TCefState
	LocalStorage               consts.TCefState
	Databases                  consts.TCefState
	Webgl                      consts.TCefState
	BackgroundColor            TCefColor
	ChromeStatusBubble         consts.TCefState
	ChromeZoomBubble           consts.TCefState
}

// TCefProxy 代理配置
type TCefProxy struct {
	ProxyType              consts.TCefProxyType
	ProxyScheme            consts.TCefProxyScheme
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
	Type          consts.TCefTouchEeventType
	Modifiers     consts.TCefEventFlags
	PointerType   consts.TCefPointerType
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
	Modifiers consts.TCefEventFlags
}

// BeforePopupInfo 弹出子窗口信息
type BeforePopupInfo struct {
	TargetUrl         string
	TargetFrameName   string
	TargetDisposition consts.TCefWindowOpenDisposition
	UserGesture       bool
}

// TCefRect
//
//	/include/internal/cef_types_geometry.h (cef_rect_t)
type TCefRect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type TCefRectArray struct {
	ptr    uintptr
	sizeOf uintptr
	count  uint32
}

type TRGBQuad struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

// NewTCefRectArray
//
//	TCefRect 动态数组结构, 通过指针引用取值
func NewTCefRectArray(ptr uintptr, count uint32) *TCefRectArray {
	return &TCefRectArray{
		ptr:    ptr,
		sizeOf: unsafe.Sizeof(TCefRect{}),
		count:  count,
	}
}

func (m *TCefRectArray) Count() int {
	return int(m.count)
}

func (m *TCefRectArray) Get(index int) *TCefRect {
	if m.count == 0 || index < 0 || index >= int(m.count) {
		return nil
	}
	return (*TCefRect)(common.GetParamPtr(m.ptr, index*int(m.sizeOf)))
}

// TCefSize
//
//	/include/internal/cef_types_geometry.h (cef_size_t)
type TCefSize struct {
	Width  int32
	Height int32
}

// TCefPoint
//
//	/include/internal/cef_types_geometry.h (cef_point_t)
type TCefPoint struct {
	X int32
	Y int32
}

// TCefCursorInfo
//
//	/include/internal/cef_types.h (cef_cursor_info_t)
type TCefCursorInfo struct {
	Hotspot          TCefPoint
	ImageScaleFactor Single
	Buffer           uintptr
	Size             TCefSize
}

// TCefBaseRefCounted
type TCefBaseRefCounted struct {
	instance unsafe.Pointer
}

// TCefResourceHandlerClass
type TCefResourceHandlerClass uintptr

// TCefScreenInfo
//
//	/include/internal/cef_types.h (cef_screen_info_t)
type TCefScreenInfo struct {
	DeviceScaleFactor Single
	Depth             int32
	DepthPerComponent int32
	IsMonochrome      int32
	Rect              TCefRect
	AvailableRect     TCefRect
}

// Touch handle state.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_touch_handle_state_t)</see></para>
type TCefTouchHandleState struct {
	//  Touch handle id. Increments for each new touch handle.
	TouchHandleId int32
	//  Combination of TCefTouchHandleStateFlags values indicating what state is set.
	Flags uint32
	//  Enabled state. Only set if |flags| contains CEF_THS_FLAG_ENABLED.
	Enabled int32
	//  Orientation state. Only set if |flags| contains CEF_THS_FLAG_ORIENTATION.
	Orientation      consts.TCefHorizontalAlignment
	MirrorVertical   int32
	MirrorHorizontal int32
	//  Origin state. Only set if |flags| contains CEF_THS_FLAG_ORIGIN.
	Origin TCefPoint
	//  Alpha state. Only set if |flags| contains CEF_THS_FLAG_ALPHA.
	Alpha float32
}

// TCefRange
//
//	/include/internal/cef_types_geometry.h (cef_range_t)
type TCefRange struct {
	From int32
	To   int32
}

// include/internal/cef_types_geometry.h (cef_insets_t)
type TCefInsets struct {
	Top    int32
	Left   int32
	Bottom int32
	Right  int32
}

type TCefPdfPrintSettings struct {
	Landscape           int32                         // Integer
	PrintBackground     int32                         // Integer
	Scale               float64                       // double
	PaperWidth          float64                       // double
	PaperHeight         float64                       // double
	PreferCssPageSize   int32                         // Integer
	MarginType          consts.TCefPdfPrintMarginType // TCefPdfPrintMarginType
	MarginTop           float64                       // double
	MarginRight         float64                       // double
	MarginBottom        float64                       // double
	MarginLeft          float64                       // double
	PageRanges          string                        // TCefString
	DisplayHeaderFooter int32                         // Integer
	HeaderTemplate      string                        // TCefString
	FooterTemplate      string                        // TCefString
}

// include/internal/cef_types.h (cef_popup_features_t)
type TCefPopupFeatures struct {
	X                  int32
	XSet               int32
	Y                  int32
	YSet               int32
	Width              int32
	WidthSet           int32
	Height             int32
	HeightSet          int32
	MenuBarVisible     int32 // Use-CEF:[49]
	StatusBarVisible   int32 // Use-CEF:[49]
	ToolBarVisible     int32 // Use-CEF:[49]
	LocationBarVisible int32
	ScrollbarsVisible  int32 // Use-CEF:[49]
	IsPopup            int32 // CEF 110 ~ Current :True (1) if browser interface elements should be hidden.
	Resizable          int32
	Fullscreen         int32
	Dialog             int32
	AdditionalFeatures TCefStringList // Use-CEF:[49]
}

// Structure representing IME composition underline information. This is a thin
// wrapper around Blink's WebCompositionUnderline class and should be kept in
// sync with that.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_composition_underline_t)</see></para>
type TCefCompositionUnderline struct {
	//  Underline character range.
	Range TCefRange
	//  Text color.
	Color TCefColor
	//  Background color.
	BackgroundColor TCefColor
	//  Set to true (1) for thick underline.
	Thick int32
	//  Style.
	Style consts.TCefCompositionUnderlineStyle
}

// Initialization settings. Specify NULL or 0 to get the recommended default
// values. Many of these and other settings can also configured using command-
// line switches.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_box_layout_settings_t)</see></para>
type TCefBoxLayoutSettings struct {
	// If true (1) the layout will be horizontal, otherwise the layout will be vertical.
	Horizontal int32
	// Adds additional horizontal space between the child view area and the host view border.
	InsideBorderHorizontalSpacing int32
	// Adds additional vertical space between the child view area and the host view border.
	InsideBorderVerticalSpacing int32
	// Adds additional space around the child view area.
	InsideBorderInsets TCefInsets
	// Adds additional space between child views.
	BetweenChildSpacing int32
	// Specifies where along the main axis the child views should be laid out.
	MainAxisAlignment consts.TCefMainAxisAlignment
	// Specifies where along the cross axis the child views should be laid out.
	CrossAxisAlignment consts.TCefMainAxisAlignment
	// Minimum cross axis size.
	MinimumCrossAxisSize int32
	// Default flex for views when none is specified via CefBoxLayout methods.
	// Using the preferred size as the basis, free space along the main axis is
	// distributed to views in the ratio of their flex weights. Similarly, if the
	// views will overflow the parent, space is subtracted in these ratios. A
	// flex of 0 means this view is not resized. Flex values must not be negative.
	DefaultFlex int32
}

// TLinuxWindowProperties String version
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_linux_window_properties_t)</see></para>
type TLinuxWindowProperties struct {
	instance *tLinuxWindowPropertiesPtr
	// Main window's Wayland's app_id
	WaylandAppId string
	// Main window's WM_CLASS_CLASS in X11
	WmClassClass string
	// Main window's WM_CLASS_NAME in X11
	WmClassName string
	// Main window's WM_WINDOW_ROLE in X11
	WmRoleName string
}

// ResultString 字符串返回值
type ResultString struct {
	value string
}

func (m *ResultString) SetValue(value string) {
	m.value = value
}

func (m *ResultString) Value() string {
	return m.value
}

// ResultBool  bool返回值
type ResultBool struct {
	value bool
}

func (m *ResultBool) SetValue(value bool) {
	m.value = value
}

func (m *ResultBool) Value() bool {
	return m.value
}

// ResultBytes  []byte返回值
type ResultBytes struct {
	value []byte
}

func (m *ResultBytes) SetValue(value []byte) {
	m.value = value
}

func (m *ResultBytes) Value() []byte {
	return m.value
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

func (m *TCefKeyEvent) KeyDown() bool {
	return m.Kind == consts.KEYEVENT_RAW_KEYDOWN || m.Kind == consts.KEYEVENT_KEYDOWN
}

func (m *TCefKeyEvent) KeyUp() bool {
	return m.Kind == consts.KEYEVENT_KEYUP
}
