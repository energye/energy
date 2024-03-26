//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefWindow Parent: ICefPanel
//
//	A Window is a top-level Window/widget in the Views hierarchy. By default it
//	will have a non-client area with title bar, icon and buttons that supports
//	moving and resizing. All size and position values are in density independent
//	pixels (DIP) unless otherwise indicated. Methods must be called on the
//	browser process UI thread unless otherwise indicated.
type ICefWindow interface {
	ICefPanel
	// IsClosed
	//  Returns true(1) if the Window has been closed.
	IsClosed() bool // function
	// IsActive
	//  Returns whether the Window is the currently active Window.
	IsActive() bool // function
	// IsAlwaysOnTop
	//  Returns whether the Window has been set to be on top of other Windows in
	//  the Windowing system.
	IsAlwaysOnTop() bool // function
	// IsMaximized
	//  Returns true(1) if the Window is maximized.
	IsMaximized() bool // function
	// IsMinimized
	//  Returns true(1) if the Window is minimized.
	IsMinimized() bool // function
	// IsFullscreen
	//  Returns true(1) if the Window is fullscreen.
	IsFullscreen() bool // function
	// GetTitle
	//  Get the Window title.
	GetTitle() string // function
	// GetWindowIcon
	//  Get the Window icon.
	GetWindowIcon() ICefImage // function
	// GetWindowAppIcon
	//  Get the Window App icon.
	GetWindowAppIcon() ICefImage // function
	// AddOverlayView
	//  Add a View that will be overlayed on the Window contents with absolute
	//  positioning and high z-order. Positioning is controlled by |docking_mode|
	//  as described below. The returned cef_overlay_controller_t object is used
	//  to control the overlay. Overlays are hidden by default.
	//  With CEF_DOCKING_MODE_CUSTOM:
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned in the top-left corner.
	//  2. Optionally change the overlay position and/or size by calling
	//  CefOverlayController methods.
	//  3. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  4. The overlay will be automatically re-sized if |view|'s layout
	//  changes. Optionally change the overlay position and/or size when
	//  OnLayoutChanged is called on the Window's delegate to indicate a
	//  change in Window bounds.
	//  With other docking modes:
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned based on |docking_mode|.
	//  2. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  3. The overlay will be automatically re-sized if |view|'s layout changes
	//  and re-positioned as appropriate when the Window resizes.
	//  Overlays created by this function will receive a higher z-order then any
	//  child Views added previously. It is therefore recommended to call this
	//  function last after all other child Views have been added so that the
	//  overlay displays as the top-most child of the Window.
	AddOverlayView(view ICefView, dockingmode TCefDockingMode) ICefOverlayController // function
	// GetDisplay
	//  Returns the Display that most closely intersects the bounds of this
	//  Window. May return NULL if this Window is not currently displayed.
	GetDisplay() ICefDisplay // function
	// GetClientAreaBoundsInScreen
	//  Returns the bounds(size and position) of this Window's client area.
	//  Position is in screen coordinates.
	GetClientAreaBoundsInScreen() (resultCefRect TCefRect) // function
	// GetWindowHandle
	//  Retrieve the platform window handle for this Window.
	GetWindowHandle() TCefWindowHandle // function
	// Show
	//  Show the Window.
	Show() // procedure
	// ShowAsBrowserModalDialog
	//  Show the Window as a browser modal dialog relative to |browser_view|. A
	//  parent Window must be returned via
	//  cef_window_delegate_t::get_parent_window() and |browser_view| must belong
	//  to that parent Window. While this Window is visible, |browser_view| will
	//  be disabled while other controls in the parent Window remain enabled.
	//  Navigating or destroying the |browser_view| will close this Window
	//  automatically. Alternately, use show() and return true(1) from
	//  cef_window_delegate_t::is_window_modal_dialog() for a window modal dialog
	//  where all controls in the parent Window are disabled.
	ShowAsBrowserModalDialog(browserview ICefBrowserView) // procedure
	// Hide
	//  Hide the Window.
	Hide() // procedure
	// CenterWindow
	//  Sizes the Window to |size| and centers it in the current display.
	CenterWindow(size *TCefSize) // procedure
	// Close
	//  Close the Window.
	Close() // procedure
	// Activate
	//  Activate the Window, assuming it already exists and is visible.
	Activate() // procedure
	// Deactivate
	//  Deactivate the Window, making the next Window in the Z order the active
	//  Window.
	Deactivate() // procedure
	// BringToTop
	//  Bring this Window to the top of other Windows in the Windowing system.
	BringToTop() // procedure
	// SetAlwaysOnTop
	//  Set the Window to be on top of other Windows in the Windowing system.
	SetAlwaysOnTop(ontop bool) // procedure
	// Maximize
	//  Maximize the Window.
	Maximize() // procedure
	// Minimize
	//  Minimize the Window.
	Minimize() // procedure
	// Restore
	//  Restore the Window.
	Restore() // procedure
	// SetFullscreen
	//  Set fullscreen Window state. The
	//  ICefWindowDelegate.OnWindowFullscreenTransition function will be
	//  called during the fullscreen transition for notification purposes.
	SetFullscreen(fullscreen bool) // procedure
	// SetTitle
	//  Set the Window title.
	SetTitle(title string) // procedure
	// SetWindowIcon
	//  Set the Window icon. This should be a 16x16 icon suitable for use in the
	//  Windows's title bar.
	SetWindowIcon(image ICefImage) // procedure
	// SetWindowAppIcon
	//  Set the Window App icon. This should be a larger icon for use in the host
	//  environment app switching UI. On Windows, this is the ICON_BIG used in
	//  Alt-Tab list and Windows taskbar. The Window icon will be used by default
	//  if no Window App icon is specified.
	SetWindowAppIcon(image ICefImage) // procedure
	// ShowMenu
	//  Show a menu with contents |menu_model|. |screen_point| specifies the menu
	//  position in screen coordinates. |anchor_position| specifies how the menu
	//  will be anchored relative to |screen_point|.
	ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) // procedure
	// CancelMenu
	//  Cancel the menu that is currently showing, if any.
	CancelMenu() // procedure
	// SetDraggableRegions
	//  Set the regions where mouse events will be intercepted by this Window to
	//  support drag operations. Call this function with an NULL vector to clear
	//  the draggable regions. The draggable region bounds should be in window
	//  coordinates.
	SetDraggableRegions(regionsCount NativeUInt, regions TCefDraggableRegionArray) // procedure
	// SendKeyPress
	//  Simulate a key press. |key_code| is the VKEY_* value from Chromium's
	//  ui/events/keycodes/keyboard_codes.h header(VK_* values on Windows).
	//  |event_flags| is some combination of EVENTFLAG_SHIFT_DOWN,
	//  EVENTFLAG_CONTROL_DOWN and/or EVENTFLAG_ALT_DOWN. This function is exposed
	//  primarily for testing purposes.
	SendKeyPress(keycode int32, eventflags uint32) // procedure
	// SendMouseMove
	//  Simulate a mouse move. The mouse cursor will be moved to the specified
	// (screen_x, screen_y) position. This function is exposed primarily for
	//  testing purposes.
	SendMouseMove(screenx, screeny int32) // procedure
	// SendMouseEvents
	//  Simulate mouse down and/or mouse up events. |button| is the mouse button
	//  type. If |mouse_down| is true(1) a mouse down event will be sent. If
	//  |mouse_up| is true(1) a mouse up event will be sent. If both are true(1)
	//  a mouse down event will be sent followed by a mouse up event(equivalent
	//  to clicking the mouse button). The events will be sent using the current
	//  cursor position so make sure to call send_mouse_move() first to position
	//  the mouse. This function is exposed primarily for testing purposes.
	SendMouseEvents(button TCefMouseButtonType, mousedown, mouseup bool) // procedure
	// SetAccelerator
	//  Set the keyboard accelerator for the specified |command_id|. |key_code|
	//  can be any virtual key or character value.
	//  cef_window_delegate_t::OnAccelerator will be called if the keyboard
	//  combination is triggered while this window has focus.
	SetAccelerator(commandid, keycode int32, shiftpressed, ctrlpressed, altpressed bool) // procedure
	// RemoveAccelerator
	//  Remove the keyboard accelerator for the specified |command_id|.
	RemoveAccelerator(commandid int32) // procedure
	// RemoveAllAccelerators
	//  Remove all keyboard accelerators.
	RemoveAllAccelerators() // procedure
}

// TCefWindow Parent: TCefPanel
//
//	A Window is a top-level Window/widget in the Views hierarchy. By default it
//	will have a non-client area with title bar, icon and buttons that supports
//	moving and resizing. All size and position values are in density independent
//	pixels (DIP) unless otherwise indicated. Methods must be called on the
//	browser process UI thread unless otherwise indicated.
type TCefWindow struct {
	TCefPanel
}

// WindowRef -> ICefWindow
var WindowRef window

// window TCefWindow Ref
type window uintptr

// UnWrap
//
//	Returns a ICefWindow instance using a PCefWindow data pointer.
func (m *window) UnWrap(data uintptr) ICefWindow {
	var resultCefWindow uintptr
	CEF().SysCallN(1611, uintptr(data), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

// CreateTopLevel
//
//	Create a new Window.
func (m *window) CreateTopLevel(delegate ICefWindowDelegate) ICefWindow {
	var resultCefWindow uintptr
	CEF().SysCallN(1578, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCefWindow) IsClosed() bool {
	r1 := CEF().SysCallN(1589, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsActive() bool {
	r1 := CEF().SysCallN(1587, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsAlwaysOnTop() bool {
	r1 := CEF().SysCallN(1588, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsMaximized() bool {
	r1 := CEF().SysCallN(1591, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsMinimized() bool {
	r1 := CEF().SysCallN(1592, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) IsFullscreen() bool {
	r1 := CEF().SysCallN(1590, m.Instance())
	return GoBool(r1)
}

func (m *TCefWindow) GetTitle() string {
	r1 := CEF().SysCallN(1582, m.Instance())
	return GoStr(r1)
}

func (m *TCefWindow) GetWindowIcon() ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(1585, m.Instance(), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefWindow) GetWindowAppIcon() ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(1583, m.Instance(), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefWindow) AddOverlayView(view ICefView, dockingmode TCefDockingMode) ICefOverlayController {
	var resultCefOverlayController uintptr
	CEF().SysCallN(1573, m.Instance(), GetObjectUintptr(view), uintptr(dockingmode), uintptr(unsafePointer(&resultCefOverlayController)))
	return AsCefOverlayController(resultCefOverlayController)
}

func (m *TCefWindow) GetDisplay() ICefDisplay {
	var resultCefDisplay uintptr
	CEF().SysCallN(1581, m.Instance(), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

func (m *TCefWindow) GetClientAreaBoundsInScreen() (resultCefRect TCefRect) {
	CEF().SysCallN(1580, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefWindow) GetWindowHandle() TCefWindowHandle {
	r1 := CEF().SysCallN(1584, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TCefWindow) Show() {
	CEF().SysCallN(1608, m.Instance())
}

func (m *TCefWindow) ShowAsBrowserModalDialog(browserview ICefBrowserView) {
	CEF().SysCallN(1609, m.Instance(), GetObjectUintptr(browserview))
}

func (m *TCefWindow) Hide() {
	CEF().SysCallN(1586, m.Instance())
}

func (m *TCefWindow) CenterWindow(size *TCefSize) {
	CEF().SysCallN(1576, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCefWindow) Close() {
	CEF().SysCallN(1577, m.Instance())
}

func (m *TCefWindow) Activate() {
	CEF().SysCallN(1572, m.Instance())
}

func (m *TCefWindow) Deactivate() {
	CEF().SysCallN(1579, m.Instance())
}

func (m *TCefWindow) BringToTop() {
	CEF().SysCallN(1574, m.Instance())
}

func (m *TCefWindow) SetAlwaysOnTop(ontop bool) {
	CEF().SysCallN(1602, m.Instance(), PascalBool(ontop))
}

func (m *TCefWindow) Maximize() {
	CEF().SysCallN(1593, m.Instance())
}

func (m *TCefWindow) Minimize() {
	CEF().SysCallN(1594, m.Instance())
}

func (m *TCefWindow) Restore() {
	CEF().SysCallN(1597, m.Instance())
}

func (m *TCefWindow) SetFullscreen(fullscreen bool) {
	CEF().SysCallN(1604, m.Instance(), PascalBool(fullscreen))
}

func (m *TCefWindow) SetTitle(title string) {
	CEF().SysCallN(1605, m.Instance(), PascalStr(title))
}

func (m *TCefWindow) SetWindowIcon(image ICefImage) {
	CEF().SysCallN(1607, m.Instance(), GetObjectUintptr(image))
}

func (m *TCefWindow) SetWindowAppIcon(image ICefImage) {
	CEF().SysCallN(1606, m.Instance(), GetObjectUintptr(image))
}

func (m *TCefWindow) ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) {
	CEF().SysCallN(1610, m.Instance(), GetObjectUintptr(menumodel), uintptr(unsafePointer(screenpoint)), uintptr(anchorposition))
}

func (m *TCefWindow) CancelMenu() {
	CEF().SysCallN(1575, m.Instance())
}

func (m *TCefWindow) SetDraggableRegions(regionsCount NativeUInt, regions TCefDraggableRegionArray) {
	CEF().SysCallN(1603, m.Instance(), uintptr(regionsCount), uintptr(unsafePointer(&regions[0])))
}

func (m *TCefWindow) SendKeyPress(keycode int32, eventflags uint32) {
	CEF().SysCallN(1598, m.Instance(), uintptr(keycode), uintptr(eventflags))
}

func (m *TCefWindow) SendMouseMove(screenx, screeny int32) {
	CEF().SysCallN(1600, m.Instance(), uintptr(screenx), uintptr(screeny))
}

func (m *TCefWindow) SendMouseEvents(button TCefMouseButtonType, mousedown, mouseup bool) {
	CEF().SysCallN(1599, m.Instance(), uintptr(button), PascalBool(mousedown), PascalBool(mouseup))
}

func (m *TCefWindow) SetAccelerator(commandid, keycode int32, shiftpressed, ctrlpressed, altpressed bool) {
	CEF().SysCallN(1601, m.Instance(), uintptr(commandid), uintptr(keycode), PascalBool(shiftpressed), PascalBool(ctrlpressed), PascalBool(altpressed))
}

func (m *TCefWindow) RemoveAccelerator(commandid int32) {
	CEF().SysCallN(1595, m.Instance(), uintptr(commandid))
}

func (m *TCefWindow) RemoveAllAccelerators() {
	CEF().SysCallN(1596, m.Instance())
}
