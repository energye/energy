//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Device is a representation of GDK's GdkDevice.
type Device struct {
	Object
}

func AsDevice(ptr unsafe.Pointer) *Device {
	if ptr == nil {
		return nil
	}
	m := new(Device)
	m.instance = ptr
	return m
}

// DeviceManager is a representation of GDK's GdkDeviceManager.
type DeviceManager struct {
	Object
}

func AsDeviceManager(ptr unsafe.Pointer) *DeviceManager {
	if ptr == nil {
		return nil
	}
	m := new(DeviceManager)
	m.instance = ptr
	return m
}

// GetDisplay is a wrapper around gdk_device_manager_get_display().
func (m *DeviceManager) GetDisplay() *Display {
	r := gdk3.SysCall("gdk_device_manager_get_display", m.Instance())
	if r == 0 {
		return nil
	}
	return AsDisplay(unsafe.Pointer(r))
}

// Display is a representation of GDK's GdkDisplay.
type Display struct {
	Object
}

func AsDisplay(ptr unsafe.Pointer) *Display {
	if ptr == nil {
		return nil
	}
	m := new(Display)
	m.instance = ptr
	return m
}

// DisplayOpen is a wrapper around gdk_display_open().
func DisplayOpen(displayName string) (*Display, error) {
	cstr := CStr(displayName)
	r := gdk3.SysCall("gdk_display_open", cstr)
	if r == 0 {
		return nil, errNilPtr
	}
	return AsDisplay(unsafe.Pointer(r)), nil
}

// DisplayGetDefault is a wrapper around gdk_display_get_default().
func DisplayGetDefault() *Display {
	r := gdk3.SysCall("gdk_display_get_default")
	if r == 0 {
		return nil
	}
	return AsDisplay(unsafe.Pointer(r))
}

// GetName is a wrapper around gdk_display_get_name().
func (m *Display) GetName() string {
	r := gdk3.SysCall("gdk_display_get_name", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// GetDefaultScreen is a wrapper around gdk_display_get_default_screen().
func (m *Display) GetDefaultScreen() IScreen {
	r := gdk3.SysCall("gdk_display_get_default_screen", m.Instance())
	if r == 0 {
		return nil
	}
	return AsScreen(unsafe.Pointer(r))
}

// DeviceIsGrabbed is a wrapper around gdk_display_device_is_grabbed().
func (m *Display) DeviceIsGrabbed(device *Device) bool {
	r := gdk3.SysCall("gdk_display_device_is_grabbed", m.Instance(), device.Instance())
	return ToGoBool(r)
}

// Beep is a wrapper around gdk_display_beep().
func (m *Display) Beep() {
	gdk3.SysCall("gdk_display_beep", m.Instance())
}

// Sync is a wrapper around gdk_display_sync().
func (m *Display) Sync() {
	gdk3.SysCall("gdk_display_sync", m.Instance())
}

// Flush is a wrapper around gdk_display_flush().
func (m *Display) Flush() {
	gdk3.SysCall("gdk_display_flush", m.Instance())
}

// Close is a wrapper around gdk_display_close().
func (m *Display) Close() {
	gdk3.SysCall("gdk_display_close", m.Instance())
}

// IsClosed is a wrapper around gdk_display_is_closed().
func (m *Display) IsClosed() bool {
	r := gdk3.SysCall("gdk_display_is_closed", m.Instance())
	return ToGoBool(r)
}

// GetEvent is a wrapper around gdk_display_get_event().
func (m *Display) GetEvent() *Event {
	r := gdk3.SysCall("gdk_display_get_event", m.Instance())
	if r == 0 {
		return nil
	}
	event := new(Event)
	event.instance = unsafe.Pointer(r)
	return event
}

// PeekEvent is a wrapper around gdk_display_peek_event().
func (m *Display) PeekEvent() *Event {
	r := gdk3.SysCall("gdk_display_peek_event", m.Instance())
	if r == 0 {
		return nil
	}
	event := new(Event)
	event.instance = unsafe.Pointer(r)
	return event
}

// PutEvent is a wrapper around gdk_display_put_event().
func (m *Display) PutEvent(event *Event) {
	gdk3.SysCall("gdk_display_put_event", m.Instance(), event.Instance())
}

// HasPending is a wrapper around gdk_display_has_pending().
func (m *Display) HasPending() bool {
	r := gdk3.SysCall("gdk_display_has_pending", m.Instance())
	return ToGoBool(r)
}

// SetDoubleClickTime is a wrapper around gdk_display_set_double_click_time().
func (m *Display) SetDoubleClickTime(msec uint) {
	gdk3.SysCall("gdk_display_set_double_click_time", m.Instance(), uintptr(msec))
}

// SetDoubleClickDistance is a wrapper around gdk_display_set_double_click_distance().
func (m *Display) SetDoubleClickDistance(distance uint) {
	gdk3.SysCall("gdk_display_set_double_click_distance", m.Instance(), uintptr(distance))
}

// SupportsColorCursor is a wrapper around gdk_display_supports_cursor_color().
func (m *Display) SupportsColorCursor() bool {
	r := gdk3.SysCall("gdk_display_supports_cursor_color", m.Instance())
	return ToGoBool(r)
}

// SupportsCursorAlpha is a wrapper around gdk_display_supports_cursor_alpha().
func (m *Display) SupportsCursorAlpha() bool {
	r := gdk3.SysCall("gdk_display_supports_cursor_alpha", m.Instance())
	return ToGoBool(r)
}

// GetDefaultCursorSize is a wrapper around gdk_display_get_default_cursor_size().
func (m *Display) GetDefaultCursorSize() uint {
	r := gdk3.SysCall("gdk_display_get_default_cursor_size", m.Instance())
	return uint(r)
}

// GetMaximalCursorSize is a wrapper around gdk_display_get_maximal_cursor_size().
func (m *Display) GetMaximalCursorSize() (uint, uint) {
	var w, h uint32
	gdk3.SysCall("gdk_display_get_maximal_cursor_size", m.Instance(),
		uintptr(unsafe.Pointer(&w)), uintptr(unsafe.Pointer(&h)))
	return uint(w), uint(h)
}

// GetDefaultGroup is a wrapper around gdk_display_get_default_group().
func (m *Display) GetDefaultGroup() *Window {
	r := gdk3.SysCall("gdk_display_get_default_group", m.Instance())
	if r == 0 {
		return nil
	}
	window := new(Window)
	window.instance = unsafe.Pointer(r)
	return window
}

// SupportsSelectionNotification is a wrapper around gdk_display_supports_selection_notification().
func (m *Display) SupportsSelectionNotification() bool {
	r := gdk3.SysCall("gdk_display_supports_selection_notification", m.Instance())
	return ToGoBool(r)
}

// RequestSelectionNotification is a wrapper around gdk_display_request_selection_notification().
func (m *Display) RequestSelectionNotification(selection TAtom) bool {
	r := gdk3.SysCall("gdk_display_request_selection_notification", m.Instance(), uintptr(selection))
	return ToGoBool(r)
}

// SupportsClipboardPersistence is a wrapper around gdk_display_supports_clipboard_persistence().
func (m *Display) SupportsClipboardPersistence() bool {
	r := gdk3.SysCall("gdk_display_supports_clipboard_persistence", m.Instance())
	return ToGoBool(r)
}

// SupportsShapes is a wrapper around gdk_display_supports_shapes().
func (m *Display) SupportsShapes() bool {
	r := gdk3.SysCall("gdk_display_supports_shapes", m.Instance())
	return ToGoBool(r)
}

// SupportsInputShapes is a wrapper around gdk_display_supports_input_shapes().
func (m *Display) SupportsInputShapes() bool {
	r := gdk3.SysCall("gdk_display_supports_input_shapes", m.Instance())
	return ToGoBool(r)
}

// NotifyStartupComplete is a wrapper around gdk_display_notify_startup_complete().
func (m *Display) NotifyStartupComplete(startupID string) {
	cstr := CStr(startupID)
	gdk3.SysCall("gdk_display_notify_startup_complete", m.Instance(), cstr)
}

// GetDefaultSeat is a wrapper around gdk_display_get_default_seat().
func (m *Display) GetDefaultSeat() *Seat {
	r := gdk3.SysCall("gdk_display_get_default_seat", m.Instance())
	if r == 0 {
		return nil
	}
	return AsSeat(unsafe.Pointer(r))
}

func (m *Display) GetXDisplay() uintptr {
	r := gdk3.SysCall("gdk_x11_display_get_xdisplay", m.Instance())
	if r == 0 {
		return 0
	}
	return r
}

// DisplayManager is a representation of GDK's GdkDisplayManager.
type DisplayManager struct {
	Object
}

func AsDisplayManager(ptr unsafe.Pointer) *DisplayManager {
	if ptr == nil {
		return nil
	}
	m := new(DisplayManager)
	m.instance = ptr
	return m
}

// DisplayManagerGet is a wrapper around gdk_display_manager_get().
func DisplayManagerGet() (*DisplayManager, error) {
	r := gdk3.SysCall("gdk_display_manager_get")
	if r == 0 {
		return nil, errNilPtr
	}
	return AsDisplayManager(unsafe.Pointer(r)), nil
}

// GetDefaultDisplay is a wrapper around gdk_display_manager_get_default_display().
func (m *DisplayManager) GetDefaultDisplay() *Display {
	r := gdk3.SysCall("gdk_display_manager_get_default_display", m.Instance())
	if r == 0 {
		return nil
	}
	return AsDisplay(unsafe.Pointer(r))
}

// SetDefaultDisplay is a wrapper around gdk_display_manager_set_default_display().
func (m *DisplayManager) SetDefaultDisplay(display *Display) {
	gdk3.SysCall("gdk_display_manager_set_default_display", m.Instance(), display.Instance())
}

// OpenDisplay is a wrapper around gdk_display_manager_open_display().
func (m *DisplayManager) OpenDisplay(name string) *Display {
	cstr := CStr(name)
	r := gdk3.SysCall("gdk_display_manager_open_display", m.Instance(), cstr)
	if r == 0 {
		return nil
	}
	return AsDisplay(unsafe.Pointer(r))
}

// Seat is a representation of GDK's GdkSeat.
type Seat struct {
	Object
}

func AsSeat(ptr unsafe.Pointer) *Seat {
	if ptr == nil {
		return nil
	}
	m := new(Seat)
	m.instance = ptr
	return m
}

// GetPointer is a wrapper around gdk_seat_get_pointer().
func (m *Seat) GetPointer() *Device {
	r := gdk3.SysCall("gdk_seat_get_pointer", m.Instance())
	if r == 0 {
		return nil
	}
	return AsDevice(unsafe.Pointer(r))
}
