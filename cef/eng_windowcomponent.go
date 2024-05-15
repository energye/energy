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

// ICEFWindowComponent Parent: ICEFPanelComponent
type ICEFWindowComponent interface {
	ICEFPanelComponent
	// Title
	//  Get the Window title.
	Title() string // property
	// SetTitle Set Title
	SetTitle(AValue string) // property
	// WindowIcon
	//  Get the Window icon.
	WindowIcon() ICefImage // property
	// SetWindowIcon Set WindowIcon
	SetWindowIcon(AValue ICefImage) // property
	// WindowAppIcon
	//  Get or set the Window App icon. This should be a larger icon for use in the host
	//  environment app switching UI. On Windows, this is the ICON_BIG used in
	//  Alt-Tab list and Windows taskbar. The Window icon will be used by default
	//  if no Window App icon is specified.
	WindowAppIcon() ICefImage // property
	// SetWindowAppIcon Set WindowAppIcon
	SetWindowAppIcon(AValue ICefImage) // property
	// Display
	//  Returns the Display that most closely intersects the bounds of this
	//  Window. May return NULL if this Window is not currently displayed.
	Display() ICefDisplay // property
	// ClientAreaBoundsInScreen
	//  Returns the bounds(size and position) of this Window's client area.
	//  Position is in screen coordinates.
	ClientAreaBoundsInScreen() (resultCefRect TCefRect) // property
	// WindowHandle
	//  Retrieve the platform window handle for this Window.
	WindowHandle() TCefWindowHandle // property
	// IsClosed
	//  Returns true(1) if the Window has been closed.
	IsClosed() bool // property
	// IsActive
	//  Returns whether the Window is the currently active Window.
	IsActive() bool // property
	// IsAlwaysOnTop
	//  Returns whether the Window has been set to be on top of other Windows in
	//  the Windowing system.
	IsAlwaysOnTop() bool // property
	// SetIsAlwaysOnTop Set IsAlwaysOnTop
	SetIsAlwaysOnTop(AValue bool) // property
	// IsFullscreen
	//  Returns true(1) if the Window is fullscreen.
	IsFullscreen() bool // property
	// SetIsFullscreen Set IsFullscreen
	SetIsFullscreen(AValue bool) // property
	// IsMaximized
	//  Returns true(1) if the Window is maximized.
	IsMaximized() bool // property
	// IsMinimized
	//  Returns true(1) if the Window is minimized.
	IsMinimized() bool // property
	// AddOverlayView
	//  Add a View that will be overlayed on the Window contents with absolute
	//  positioning and high z-order. Positioning is controlled by |docking_mode|
	//  as described below. The returned cef_overlay_controller_t object is used
	//  to control the overlay. Overlays are hidden by default.
	//  With CEF_DOCKING_MODE_CUSTOM:
	//  <code>
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned in the top-left corner.
	//  2. Optionally change the overlay position and/or size by calling
	//  CefOverlayController methods.
	//  3. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  4. The overlay will be automatically re-sized if |view|'s layout
	//  changes. Optionally change the overlay position and/or size when
	//  OnLayoutChanged is called on the Window's delegate to indicate a
	//  change in Window bounds.</code>
	//  With other docking modes:
	//  <code>
	//  1. The overlay is initially hidden, sized to |view|'s preferred size,
	//  and positioned based on |docking_mode|.
	//  2. Call CefOverlayController::SetVisible(true) to show the overlay.
	//  3. The overlay will be automatically re-sized if |view|'s layout changes
	//  and re-positioned as appropriate when the Window resizes.</code>
	//  Overlays created by this function will receive a higher z-order then any
	//  child Views added previously. It is therefore recommended to call this
	//  function last after all other child Views have been added so that the
	//  overlay displays as the top-most child of the Window.
	AddOverlayView(view ICefView, dockingmode TCefDockingMode) ICefOverlayController // function
	// CreateTopLevelWindow
	//  Create a new Window.
	CreateTopLevelWindow() // procedure
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
	// Maximize
	//  Maximize the Window.
	Maximize() // procedure
	// Minimize
	//  Minimize the Window.
	Minimize() // procedure
	// Restore
	//  Restore the Window.
	Restore() // procedure
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
	//  (screen_x, screen_y) position. This function is exposed primarily for
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
	// SetOnWindowCreated
	//  Called when |window| is created.
	SetOnWindowCreated(fn TOnWindowCreated) // property event
	// SetOnWindowClosing
	//  Called when |window| is closing.
	SetOnWindowClosing(fn TOnWindowClosing) // property event
	// SetOnWindowDestroyed
	//  Called when |window| is destroyed. Release all references to |window| and
	//  do not attempt to execute any functions on |window| after this callback
	//  returns.
	SetOnWindowDestroyed(fn TOnWindowDestroyed) // property event
	// SetOnWindowActivationChanged
	//  Called when |window| is activated or deactivated.
	SetOnWindowActivationChanged(fn TOnWindowActivationChanged) // property event
	// SetOnWindowBoundsChanged
	//  Called when |window| bounds have changed. |new_bounds| will be in DIP
	//  screen coordinates.
	SetOnWindowBoundsChanged(fn TOnWindowBoundsChanged) // property event
	// SetOnGetParentWindow
	//  Return the parent for |window| or NULL if the |window| does not have a
	//  parent. Windows with parents will not get a taskbar button. Set |is_menu|
	//  to true(1) if |window| will be displayed as a menu, in which case it will
	//  not be clipped to the parent window bounds. Set |can_activate_menu| to
	//  false(0) if |is_menu| is true(1) and |window| should not be activated
	//  (given keyboard focus) when displayed.
	SetOnGetParentWindow(fn TOnGetParentWindow) // property event
	// SetOnIsWindowModalDialog
	//  Return true(1) if |window| should be created as a window modal dialog.
	//  Only called when a Window is returned via get_parent_window() with
	//  |is_menu| set to false(0). All controls in the parent Window will be
	//  disabled while |window| is visible. This functionality is not supported by
	//  all Linux window managers. Alternately, use
	//  ICefWindow.ShowAsBrowserModalDialog() for a browser modal dialog
	//  that works on all platforms.
	SetOnIsWindowModalDialog(fn TOnIsWindowModalDialog) // property event
	// SetOnGetInitialBounds
	//  Return the initial bounds for |window| in density independent pixel(DIP)
	//  coordinates. If this function returns an NULL CefRect then
	//  GetPreferredSize() will be called to retrieve the size, and the window
	//  will be placed on the screen with origin(0,0). This function can be used
	//  in combination with ICefView.GetBoundsInScreen() to restore the
	//  previous window bounds.
	SetOnGetInitialBounds(fn TOnGetInitialBounds) // property event
	// SetOnGetInitialShowState
	//  Return the initial show state for |window|.
	SetOnGetInitialShowState(fn TOnGetInitialShowState) // property event
	// SetOnIsFrameless
	//  Return true(1) if |window| should be created without a frame or title
	//  bar. The window will be resizable if can_resize() returns true(1). Use
	//  ICefWindow.SetDraggableRegions() to specify draggable regions.
	SetOnIsFrameless(fn TOnIsFrameless) // property event
	// SetOnWithStandardWindowButtons
	//  Return true(1) if |window| should be created with standard window buttons
	//  like close, minimize and zoom. This function is only supported on macOS.
	SetOnWithStandardWindowButtons(fn TOnWithStandardWindowButtons) // property event
	// SetOnGetTitlebarHeight
	//  Return whether the titlebar height should be overridden, and sets the
	//  height of the titlebar in |titlebar_height|. On macOS, it can also be used
	//  to adjust the vertical position of the traffic light buttons in frameless
	//  windows. The buttons will be positioned halfway down the titlebar at a
	//  height of |titlebar_height| / 2.
	SetOnGetTitlebarHeight(fn TOnGetTitlebarHeight) // property event
	// SetOnCanResize
	//  Return true(1) if |window| can be resized.
	SetOnCanResize(fn TOnCanResize) // property event
	// SetOnCanMaximize
	//  Return true(1) if |window| can be maximized.
	SetOnCanMaximize(fn TOnCanMaximize) // property event
	// SetOnCanMinimize
	//  Return true(1) if |window| can be minimized.
	SetOnCanMinimize(fn TOnCanMinimize) // property event
	// SetOnCanClose
	//  Return true(1) if |window| can be closed. This will be called for user-
	//  initiated window close actions and when ICefWindow.close() is called.
	SetOnCanClose(fn TOnCanClose) // property event
	// SetOnAccelerator
	//  Called when a keyboard accelerator registered with
	//  ICefWindow.SetAccelerator is triggered. Return true(1) if the
	//  accelerator was handled or false(0) otherwise.
	SetOnAccelerator(fn TOnAccelerator) // property event
	// SetOnKeyEvent
	//  Called after all other controls in the window have had a chance to handle
	//  the event. |event| contains information about the keyboard event. Return
	//  true(1) if the keyboard event was handled or false(0) otherwise.
	SetOnKeyEvent(fn TOnWindowKeyEvent) // property event
	// SetOnWindowFullscreenTransition
	//  Called when |window| is transitioning to or from fullscreen mode. On MacOS
	//  the transition occurs asynchronously with |is_competed| set to false(0)
	//  when the transition starts and true(1) after the transition completes. On
	//  other platforms the transition occurs synchronously with |is_completed|
	//  set to true(1) after the transition completes. With the Alloy runtime you
	//  must also implement ICefDisplayHandler.OnFullscreenModeChange to
	//  handle fullscreen transitions initiated by browser content.
	SetOnWindowFullscreenTransition(fn TOnWindowFullscreenTransition) // property event
}

// TCEFWindowComponent Parent: TCEFPanelComponent
type TCEFWindowComponent struct {
	TCEFPanelComponent
	windowCreatedPtr              uintptr
	windowClosingPtr              uintptr
	windowDestroyedPtr            uintptr
	windowActivationChangedPtr    uintptr
	windowBoundsChangedPtr        uintptr
	getParentWindowPtr            uintptr
	isWindowModalDialogPtr        uintptr
	getInitialBoundsPtr           uintptr
	getInitialShowStatePtr        uintptr
	isFramelessPtr                uintptr
	withStandardWindowButtonsPtr  uintptr
	getTitlebarHeightPtr          uintptr
	canResizePtr                  uintptr
	canMaximizePtr                uintptr
	canMinimizePtr                uintptr
	canClosePtr                   uintptr
	acceleratorPtr                uintptr
	keyEventPtr                   uintptr
	windowFullscreenTransitionPtr uintptr
}

func NewCEFWindowComponent(aOwner IComponent) ICEFWindowComponent {
	r1 := CEF().SysCallN(356, GetObjectUintptr(aOwner))
	return AsCEFWindowComponent(r1)
}

func (m *TCEFWindowComponent) Title() string {
	r1 := CEF().SysCallN(399, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFWindowComponent) SetTitle(AValue string) {
	CEF().SysCallN(399, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFWindowComponent) WindowIcon() ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(402, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCEFWindowComponent) SetWindowIcon(AValue ICefImage) {
	CEF().SysCallN(402, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCEFWindowComponent) WindowAppIcon() ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(400, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCEFWindowComponent) SetWindowAppIcon(AValue ICefImage) {
	CEF().SysCallN(400, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCEFWindowComponent) Display() ICefDisplay {
	var resultCefDisplay uintptr
	CEF().SysCallN(359, m.Instance(), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() (resultCefRect TCefRect) {
	CEF().SysCallN(354, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCEFWindowComponent) WindowHandle() TCefWindowHandle {
	r1 := CEF().SysCallN(401, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TCEFWindowComponent) IsClosed() bool {
	r1 := CEF().SysCallN(363, m.Instance())
	return GoBool(r1)
}

func (m *TCEFWindowComponent) IsActive() bool {
	r1 := CEF().SysCallN(361, m.Instance())
	return GoBool(r1)
}

func (m *TCEFWindowComponent) IsAlwaysOnTop() bool {
	r1 := CEF().SysCallN(362, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFWindowComponent) SetIsAlwaysOnTop(AValue bool) {
	CEF().SysCallN(362, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFWindowComponent) IsFullscreen() bool {
	r1 := CEF().SysCallN(364, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFWindowComponent) SetIsFullscreen(AValue bool) {
	CEF().SysCallN(364, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFWindowComponent) IsMaximized() bool {
	r1 := CEF().SysCallN(365, m.Instance())
	return GoBool(r1)
}

func (m *TCEFWindowComponent) IsMinimized() bool {
	r1 := CEF().SysCallN(366, m.Instance())
	return GoBool(r1)
}

func (m *TCEFWindowComponent) AddOverlayView(view ICefView, dockingmode TCefDockingMode) ICefOverlayController {
	var resultCefOverlayController uintptr
	CEF().SysCallN(349, m.Instance(), GetObjectUintptr(view), uintptr(dockingmode), uintptr(unsafePointer(&resultCefOverlayController)))
	return AsCefOverlayController(resultCefOverlayController)
}

func CEFWindowComponentClass() TClass {
	ret := CEF().SysCallN(353)
	return TClass(ret)
}

func (m *TCEFWindowComponent) CreateTopLevelWindow() {
	CEF().SysCallN(357, m.Instance())
}

func (m *TCEFWindowComponent) Show() {
	CEF().SysCallN(396, m.Instance())
}

func (m *TCEFWindowComponent) ShowAsBrowserModalDialog(browserview ICefBrowserView) {
	CEF().SysCallN(397, m.Instance(), GetObjectUintptr(browserview))
}

func (m *TCEFWindowComponent) Hide() {
	CEF().SysCallN(360, m.Instance())
}

func (m *TCEFWindowComponent) CenterWindow(size *TCefSize) {
	CEF().SysCallN(352, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCEFWindowComponent) Close() {
	CEF().SysCallN(355, m.Instance())
}

func (m *TCEFWindowComponent) Activate() {
	CEF().SysCallN(348, m.Instance())
}

func (m *TCEFWindowComponent) Deactivate() {
	CEF().SysCallN(358, m.Instance())
}

func (m *TCEFWindowComponent) BringToTop() {
	CEF().SysCallN(350, m.Instance())
}

func (m *TCEFWindowComponent) Maximize() {
	CEF().SysCallN(367, m.Instance())
}

func (m *TCEFWindowComponent) Minimize() {
	CEF().SysCallN(368, m.Instance())
}

func (m *TCEFWindowComponent) Restore() {
	CEF().SysCallN(371, m.Instance())
}

func (m *TCEFWindowComponent) ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) {
	CEF().SysCallN(398, m.Instance(), GetObjectUintptr(menumodel), uintptr(unsafePointer(screenpoint)), uintptr(anchorposition))
}

func (m *TCEFWindowComponent) CancelMenu() {
	CEF().SysCallN(351, m.Instance())
}

func (m *TCEFWindowComponent) SetDraggableRegions(regionsCount NativeUInt, regions TCefDraggableRegionArray) {
	CEF().SysCallN(376, m.Instance(), uintptr(regionsCount), uintptr(unsafePointer(&regions[0])))
}

func (m *TCEFWindowComponent) SendKeyPress(keycode int32, eventflags uint32) {
	CEF().SysCallN(372, m.Instance(), uintptr(keycode), uintptr(eventflags))
}

func (m *TCEFWindowComponent) SendMouseMove(screenx, screeny int32) {
	CEF().SysCallN(374, m.Instance(), uintptr(screenx), uintptr(screeny))
}

func (m *TCEFWindowComponent) SendMouseEvents(button TCefMouseButtonType, mousedown, mouseup bool) {
	CEF().SysCallN(373, m.Instance(), uintptr(button), PascalBool(mousedown), PascalBool(mouseup))
}

func (m *TCEFWindowComponent) SetAccelerator(commandid, keycode int32, shiftpressed, ctrlpressed, altpressed bool) {
	CEF().SysCallN(375, m.Instance(), uintptr(commandid), uintptr(keycode), PascalBool(shiftpressed), PascalBool(ctrlpressed), PascalBool(altpressed))
}

func (m *TCEFWindowComponent) RemoveAccelerator(commandid int32) {
	CEF().SysCallN(369, m.Instance(), uintptr(commandid))
}

func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	CEF().SysCallN(370, m.Instance())
}

func (m *TCEFWindowComponent) SetOnWindowCreated(fn TOnWindowCreated) {
	if m.windowCreatedPtr != 0 {
		RemoveEventElement(m.windowCreatedPtr)
	}
	m.windowCreatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(392, m.Instance(), m.windowCreatedPtr)
}

func (m *TCEFWindowComponent) SetOnWindowClosing(fn TOnWindowClosing) {
	if m.windowClosingPtr != 0 {
		RemoveEventElement(m.windowClosingPtr)
	}
	m.windowClosingPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(391, m.Instance(), m.windowClosingPtr)
}

func (m *TCEFWindowComponent) SetOnWindowDestroyed(fn TOnWindowDestroyed) {
	if m.windowDestroyedPtr != 0 {
		RemoveEventElement(m.windowDestroyedPtr)
	}
	m.windowDestroyedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(393, m.Instance(), m.windowDestroyedPtr)
}

func (m *TCEFWindowComponent) SetOnWindowActivationChanged(fn TOnWindowActivationChanged) {
	if m.windowActivationChangedPtr != 0 {
		RemoveEventElement(m.windowActivationChangedPtr)
	}
	m.windowActivationChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(389, m.Instance(), m.windowActivationChangedPtr)
}

func (m *TCEFWindowComponent) SetOnWindowBoundsChanged(fn TOnWindowBoundsChanged) {
	if m.windowBoundsChangedPtr != 0 {
		RemoveEventElement(m.windowBoundsChangedPtr)
	}
	m.windowBoundsChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(390, m.Instance(), m.windowBoundsChangedPtr)
}

func (m *TCEFWindowComponent) SetOnGetParentWindow(fn TOnGetParentWindow) {
	if m.getParentWindowPtr != 0 {
		RemoveEventElement(m.getParentWindowPtr)
	}
	m.getParentWindowPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(384, m.Instance(), m.getParentWindowPtr)
}

func (m *TCEFWindowComponent) SetOnIsWindowModalDialog(fn TOnIsWindowModalDialog) {
	if m.isWindowModalDialogPtr != 0 {
		RemoveEventElement(m.isWindowModalDialogPtr)
	}
	m.isWindowModalDialogPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(387, m.Instance(), m.isWindowModalDialogPtr)
}

func (m *TCEFWindowComponent) SetOnGetInitialBounds(fn TOnGetInitialBounds) {
	if m.getInitialBoundsPtr != 0 {
		RemoveEventElement(m.getInitialBoundsPtr)
	}
	m.getInitialBoundsPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(382, m.Instance(), m.getInitialBoundsPtr)
}

func (m *TCEFWindowComponent) SetOnGetInitialShowState(fn TOnGetInitialShowState) {
	if m.getInitialShowStatePtr != 0 {
		RemoveEventElement(m.getInitialShowStatePtr)
	}
	m.getInitialShowStatePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(383, m.Instance(), m.getInitialShowStatePtr)
}

func (m *TCEFWindowComponent) SetOnIsFrameless(fn TOnIsFrameless) {
	if m.isFramelessPtr != 0 {
		RemoveEventElement(m.isFramelessPtr)
	}
	m.isFramelessPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(386, m.Instance(), m.isFramelessPtr)
}

func (m *TCEFWindowComponent) SetOnWithStandardWindowButtons(fn TOnWithStandardWindowButtons) {
	if m.withStandardWindowButtonsPtr != 0 {
		RemoveEventElement(m.withStandardWindowButtonsPtr)
	}
	m.withStandardWindowButtonsPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(395, m.Instance(), m.withStandardWindowButtonsPtr)
}

func (m *TCEFWindowComponent) SetOnGetTitlebarHeight(fn TOnGetTitlebarHeight) {
	if m.getTitlebarHeightPtr != 0 {
		RemoveEventElement(m.getTitlebarHeightPtr)
	}
	m.getTitlebarHeightPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(385, m.Instance(), m.getTitlebarHeightPtr)
}

func (m *TCEFWindowComponent) SetOnCanResize(fn TOnCanResize) {
	if m.canResizePtr != 0 {
		RemoveEventElement(m.canResizePtr)
	}
	m.canResizePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(381, m.Instance(), m.canResizePtr)
}

func (m *TCEFWindowComponent) SetOnCanMaximize(fn TOnCanMaximize) {
	if m.canMaximizePtr != 0 {
		RemoveEventElement(m.canMaximizePtr)
	}
	m.canMaximizePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(379, m.Instance(), m.canMaximizePtr)
}

func (m *TCEFWindowComponent) SetOnCanMinimize(fn TOnCanMinimize) {
	if m.canMinimizePtr != 0 {
		RemoveEventElement(m.canMinimizePtr)
	}
	m.canMinimizePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(380, m.Instance(), m.canMinimizePtr)
}

func (m *TCEFWindowComponent) SetOnCanClose(fn TOnCanClose) {
	if m.canClosePtr != 0 {
		RemoveEventElement(m.canClosePtr)
	}
	m.canClosePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(378, m.Instance(), m.canClosePtr)
}

func (m *TCEFWindowComponent) SetOnAccelerator(fn TOnAccelerator) {
	if m.acceleratorPtr != 0 {
		RemoveEventElement(m.acceleratorPtr)
	}
	m.acceleratorPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(377, m.Instance(), m.acceleratorPtr)
}

func (m *TCEFWindowComponent) SetOnKeyEvent(fn TOnWindowKeyEvent) {
	if m.keyEventPtr != 0 {
		RemoveEventElement(m.keyEventPtr)
	}
	m.keyEventPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(388, m.Instance(), m.keyEventPtr)
}

func (m *TCEFWindowComponent) SetOnWindowFullscreenTransition(fn TOnWindowFullscreenTransition) {
	if m.windowFullscreenTransitionPtr != 0 {
		RemoveEventElement(m.windowFullscreenTransitionPtr)
	}
	m.windowFullscreenTransitionPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(394, m.Instance(), m.windowFullscreenTransitionPtr)
}
