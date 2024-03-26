//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2Controller Parent: IObject
//
//	The owner of the `CoreWebView2` object that provides support for resizing,
//	showing and hiding, focusing, and other functionality related to
//	windowing and composition.  The `CoreWebView2Controller` owns the
//	`CoreWebView2`, and if all references to the `CoreWebView2Controller` go
//	away, the WebView is closed.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller">See the ICoreWebView2Controller article.</a>
type ICoreWebView2Controller interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Controller // property
	// ZoomFactor
	//  The zoom factor for the WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</a>
	ZoomFactor() (resultDouble float64) // property
	// SetZoomFactor Set ZoomFactor
	SetZoomFactor(AValue float64) // property
	// IsVisible
	//  The `IsVisible` property determines whether to show or hide the WebView2.
	//  If `IsVisible` is set to `FALSE`, the WebView2 is transparent and is
	//  not rendered. However, this does not affect the window containing the
	//  WebView2(the `HWND` parameter that was passed to
	//  `CreateCoreWebView2Controller`). If you want that window to disappear
	//  too, run `ShowWindow` on it directly in addition to modifying the
	//  `IsVisible` property. WebView2 as a child window does not get window
	//  messages when the top window is minimized or restored. For performance
	//  reasons, developers should set the `IsVisible` property of the WebView to
	//  `FALSE` when the app window is minimized and back to `TRUE` when the app
	//  window is restored. The app window does this by handling
	//  `SIZE_MINIMIZED and SIZE_RESTORED` command upon receiving `WM_SIZE`
	//  message.
	//  There are CPU and memory benefits when the page is hidden. For instance,
	//  Chromium has code that throttles activities on the page like animations
	//  and some tasks are run less frequently. Similarly, WebView2 will
	//  purge some caches to reduce memory usage.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_isvisible">See the ICoreWebView2Controller article.</a>
	IsVisible() bool // property
	// SetIsVisible Set IsVisible
	SetIsVisible(AValue bool) // property
	// Bounds
	//  The WebView bounds. Bounds are relative to the parent `HWND`. The app
	//  has two ways to position a WebView.
	//  * Create a child `HWND` that is the WebView parent `HWND`. Position
	//  the window where the WebView should be. Use `(0, 0)` for the
	//  top-left corner(the offset) of the `Bounds` of the WebView.
	//  * Use the top-most window of the app as the WebView parent HWND. For
	//  example, to position WebView correctly in the app, set the top-left
	//  corner of the Bound of the WebView.
	//  The values of `Bounds` are limited by the coordinate space of the host.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_bounds">See the ICoreWebView2Controller article.</a>
	Bounds() (resultRect TRect) // property
	// SetBounds Set Bounds
	SetBounds(AValue *TRect) // property
	// ParentWindow
	//  The parent window provided by the app that this WebView is using to
	//  render content. This API initially returns the window passed into
	//  `CreateCoreWebView2Controller`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_parentwindow">See the ICoreWebView2Controller article.</a>
	ParentWindow() THandle // property
	// SetParentWindow Set ParentWindow
	SetParentWindow(AValue THandle) // property
	// DefaultBackgroundColor
	//  The `DefaultBackgroundColor` property is the color WebView renders
	//  underneath all web content. This means WebView renders this color when
	//  there is no web content loaded such as before the initial navigation or
	//  between navigations. This also means web pages with undefined css
	//  background properties or background properties containing transparent
	//  pixels will render their contents over this color. Web pages with defined
	//  and opaque background properties that span the page will obscure the
	//  `DefaultBackgroundColor` and display normally. The default value for this
	//  property is white to resemble the native browser experience.
	//  The Color is specified by the COREWEBVIEW2_COLOR that represents an RGBA
	//  value. The `A` represents an Alpha value, meaning
	//  `DefaultBackgroundColor` can be transparent. In the case of a transparent
	//  `DefaultBackgroundColor` WebView will render hosting app content as the
	//  background. This Alpha value is not supported on Windows 7. Any `A` value
	//  other than 255 will result in E_INVALIDARG on Windows 7.
	//  It is supported on all other WebView compatible platforms.
	//  Semi-transparent colors are not currently supported by this API and
	//  setting `DefaultBackgroundColor` to a semi-transparent color will fail
	//  with E_INVALIDARG. The only supported alpha values are 0 and 255, all
	//  other values will result in E_INVALIDARG.
	//  `DefaultBackgroundColor` can only be an opaque color or transparent.
	//  This value may also be set by using the
	//  `WEBVIEW2_DEFAULT_BACKGROUND_COLOR` environment variable. There is a
	//  known issue with background color where setting the color by API can
	//  still leave the app with a white flicker before the
	//  `DefaultBackgroundColor` takes effect. Setting the color via environment
	//  variable solves this issue. The value must be a hex value that can
	//  optionally prepend a 0x. The value must account for the alpha value
	//  which is represented by the first 2 digits. So any hex value fewer than 8
	//  digits will assume a prepended 00 to the hex value and result in a
	//  transparent color.
	//  `get_DefaultBackgroundColor` will return the result of this environment
	//  variable if used. This environment variable can only set the
	//  `DefaultBackgroundColor` once. Subsequent updates to background color
	//  must be done through API call.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller2#get_defaultbackgroundcolor">See the ICoreWebView2Controller2 article.</a>
	DefaultBackgroundColor() TColor // property
	// SetDefaultBackgroundColor Set DefaultBackgroundColor
	SetDefaultBackgroundColor(AValue TColor) // property
	// RasterizationScale
	//  The rasterization scale for the WebView. The rasterization scale is the
	//  combination of the monitor DPI scale and text scaling set by the user.
	//  This value should be updated when the DPI scale of the app's top level
	//  window changes(i.e. monitor DPI scale changes or window changes monitor)
	//  or when the text scale factor of the system changes.
	//  Rasterization scale applies to the WebView content, as well as
	//  popups, context menus, scroll bars, and so on. Normal app scaling
	//  scenarios should use the ZoomFactor property or SetBoundsAndZoomFactor
	//  API which only scale the rendered HTML content and not popups, context
	//  menus, scroll bars, and so on.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_rasterizationscale">See the ICoreWebView2Controller3 article.</a>
	RasterizationScale() (resultDouble float64) // property
	// SetRasterizationScale Set RasterizationScale
	SetRasterizationScale(AValue float64) // property
	// ShouldDetectMonitorScaleChanges
	//  ShouldDetectMonitorScaleChanges property determines whether the WebView
	//  attempts to track monitor DPI scale changes. When true, the WebView will
	//  track monitor DPI scale changes, update the RasterizationScale property,
	//  and raises RasterizationScaleChanged event. When false, the WebView will
	//  not track monitor DPI scale changes, and the app must update the
	//  RasterizationScale property itself. RasterizationScaleChanged event will
	//  never raise when ShouldDetectMonitorScaleChanges is false. Apps that want
	//  to set their own rasterization scale should set this property to false to
	//  avoid the WebView2 updating the RasterizationScale property to match the
	//  monitor DPI scale.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_shoulddetectmonitorscalechanges">See the ICoreWebView2Controller3 article.</a>
	ShouldDetectMonitorScaleChanges() bool // property
	// SetShouldDetectMonitorScaleChanges Set ShouldDetectMonitorScaleChanges
	SetShouldDetectMonitorScaleChanges(AValue bool) // property
	// BoundsMode
	//  BoundsMode affects how setting the Bounds and RasterizationScale
	//  properties work. Bounds mode can either be in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS
	//  mode or COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE mode.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS, setting the bounds
	//  property will set the size of the WebView in raw screen pixels. Changing
	//  the rasterization scale in this mode won't change the raw pixel size of
	//  the WebView and will only change the rasterization scale.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE, setting the
	//  bounds property will change the logical size of the WebView which can be
	//  described by the following equation: Logical size * rasterization scale = Raw Pixel size
	//  In this case, changing the rasterization scale will keep the logical size
	//  the same and change the raw pixel size.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_boundsmode">See the ICoreWebView2Controller3 article.</a>
	BoundsMode() TWVBoundsMode // property
	// SetBoundsMode Set BoundsMode
	SetBoundsMode(AValue TWVBoundsMode) // property
	// CoreWebView2
	//  Gets the `CoreWebView2` associated with this `CoreWebView2Controller`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_corewebview2">See the ICoreWebView2Controller article.</a>
	CoreWebView2() ICoreWebView2 // property
	// AllowExternalDrop
	//  Gets the `AllowExternalDrop` property which is used to configure the
	//  capability that dragging objects from outside the bounds of webview2 and
	//  dropping into webview2 is allowed or disallowed. The default value is
	//  TRUE.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller4#get_allowexternaldrop">See the ICoreWebView2Controller4 article.</a>
	AllowExternalDrop() bool // property
	// SetAllowExternalDrop Set AllowExternalDrop
	SetAllowExternalDrop(AValue bool) // property
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// MoveFocus
	//  Moves focus into WebView. WebView gets focus and focus is set to
	//  correspondent element in the page hosted in the WebView. For
	//  Programmatic reason, focus is set to previously focused element or the
	//  default element if no previously focused element exists. For `Next`
	//  reason, focus is set to the first element. For `Previous` reason, focus
	//  is set to the last element. WebView changes focus through user
	//  interaction including selecting into a WebView or Tab into it. For
	//  tabbing, the app runs MoveFocus with Next or Previous to align with Tab
	//  and Shift+Tab respectively when it decides the WebView is the next
	//  element that may exist in a tab. Or, the app runs `IsDialogMessage`
	//  as part of the associated message loop to allow the platform to auto
	//  handle tabbing. The platform rotates through all windows with
	//  `WS_TABSTOP`. When the WebView gets focus from `IsDialogMessage`, it is
	//  internally put the focus on the first or last element for tab and
	//  Shift+Tab respectively.
	MoveFocus(aReason TWVMoveFocusReason) bool // function
	// Close
	//  Closes the WebView and cleans up the underlying browser instance.
	//  Cleaning up the browser instance releases the resources powering the
	//  WebView. The browser instance is shut down if no other WebViews are
	//  using it.
	//  After running `Close`, most methods will fail and event handlers stop
	//  running. Specifically, the WebView releases the associated references to
	//  any associated event handlers when `Close` is run.
	//  `Close` is implicitly run when the `CoreWebView2Controller` loses the
	//  final reference and is destructed. But it is best practice to
	//  explicitly run `Close` to avoid any accidental cycle of references
	//  between the WebView and the app code. Specifically, if you capture a
	//  reference to the WebView in an event handler you create a reference cycle
	//  between the WebView and the event handler. Run `Close` to break the
	//  cycle by releasing all event handlers. But to avoid the situation, it is
	//  best to both explicitly run `Close` on the WebView and to not capture a
	//  reference to the WebView to ensure the WebView is cleaned up correctly.
	//  `Close` is synchronous and won't trigger the `beforeunload` event.
	Close() bool // function
	// SetBoundsAndZoomFactor
	//  Updates `aBounds` and `aZoomFactor` properties at the same time. This
	//  operation is atomic from the perspective of the host. After returning
	//  from this function, the `aBounds` and `aZoomFactor` properties are both
	//  updated if the function is successful, or neither is updated if the
	//  function fails. If `aBounds` and `aZoomFactor` are both updated by the
	//  same scale(for example, `aBounds` and `aZoomFactor` are both doubled),
	//  then the page does not display a change in `window.innerWidth` or
	//  `window.innerHeight` and the WebView renders the content at the new size
	//  and zoom without intermediate renderings. This function also updates
	//  just one of `aZoomFactor` or `aBounds` by passing in the new value for one
	//  and the current value for the other.
	SetBoundsAndZoomFactor(aBounds *TRect, aZoomFactor float64) bool // function
	// NotifyParentWindowPositionChanged
	//  This is a notification separate from `Bounds` that tells WebView that the
	//  main WebView parent(or any ancestor) `HWND` moved. This is needed
	//  for accessibility and certain dialogs in WebView to work correctly.
	NotifyParentWindowPositionChanged() bool // function
}

// TCoreWebView2Controller Parent: TObject
//
//	The owner of the `CoreWebView2` object that provides support for resizing,
//	showing and hiding, focusing, and other functionality related to
//	windowing and composition.  The `CoreWebView2Controller` owns the
//	`CoreWebView2`, and if all references to the `CoreWebView2Controller` go
//	away, the WebView is closed.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller">See the ICoreWebView2Controller article.</a>
type TCoreWebView2Controller struct {
	TObject
}

func NewCoreWebView2Controller(aBaseIntf ICoreWebView2Controller) ICoreWebView2Controller {
	r1 := WV().SysCallN(190, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Controller(r1)
}

func (m *TCoreWebView2Controller) Initialized() bool {
	r1 := WV().SysCallN(192, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) BaseIntf() ICoreWebView2Controller {
	var resultCoreWebView2Controller uintptr
	WV().SysCallN(184, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Controller)))
	return AsCoreWebView2Controller(resultCoreWebView2Controller)
}

func (m *TCoreWebView2Controller) ZoomFactor() (resultDouble float64) {
	WV().SysCallN(200, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2Controller) SetZoomFactor(AValue float64) {
	WV().SysCallN(200, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2Controller) IsVisible() bool {
	r1 := WV().SysCallN(193, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) SetIsVisible(AValue bool) {
	WV().SysCallN(193, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Controller) Bounds() (resultRect TRect) {
	WV().SysCallN(185, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2Controller) SetBounds(AValue *TRect) {
	WV().SysCallN(185, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2Controller) ParentWindow() THandle {
	r1 := WV().SysCallN(196, 0, m.Instance(), 0)
	return THandle(r1)
}

func (m *TCoreWebView2Controller) SetParentWindow(AValue THandle) {
	WV().SysCallN(196, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Controller) DefaultBackgroundColor() TColor {
	r1 := WV().SysCallN(191, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCoreWebView2Controller) SetDefaultBackgroundColor(AValue TColor) {
	WV().SysCallN(191, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Controller) RasterizationScale() (resultDouble float64) {
	WV().SysCallN(197, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2Controller) SetRasterizationScale(AValue float64) {
	WV().SysCallN(197, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2Controller) ShouldDetectMonitorScaleChanges() bool {
	r1 := WV().SysCallN(199, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) SetShouldDetectMonitorScaleChanges(AValue bool) {
	WV().SysCallN(199, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Controller) BoundsMode() TWVBoundsMode {
	r1 := WV().SysCallN(186, 0, m.Instance(), 0)
	return TWVBoundsMode(r1)
}

func (m *TCoreWebView2Controller) SetBoundsMode(AValue TWVBoundsMode) {
	WV().SysCallN(186, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Controller) CoreWebView2() ICoreWebView2 {
	var resultCoreWebView2 uintptr
	WV().SysCallN(189, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2)))
	return AsCoreWebView2(resultCoreWebView2)
}

func (m *TCoreWebView2Controller) AllowExternalDrop() bool {
	r1 := WV().SysCallN(183, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) SetAllowExternalDrop(AValue bool) {
	WV().SysCallN(183, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Controller) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(182, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) MoveFocus(aReason TWVMoveFocusReason) bool {
	r1 := WV().SysCallN(194, m.Instance(), uintptr(aReason))
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) Close() bool {
	r1 := WV().SysCallN(188, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) SetBoundsAndZoomFactor(aBounds *TRect, aZoomFactor float64) bool {
	r1 := WV().SysCallN(198, m.Instance(), uintptr(unsafePointer(aBounds)), uintptr(unsafePointer(&aZoomFactor)))
	return GoBool(r1)
}

func (m *TCoreWebView2Controller) NotifyParentWindowPositionChanged() bool {
	r1 := WV().SysCallN(195, m.Instance())
	return GoBool(r1)
}

func CoreWebView2ControllerClass() TClass {
	ret := WV().SysCallN(187)
	return TClass(ret)
}
