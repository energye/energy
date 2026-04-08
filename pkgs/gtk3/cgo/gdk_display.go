package cgo

/*
#cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
#include <gdk/gdk.h>
#include "gdk.go.h"

static GdkSeat *toGdkSeat(void *p) { return ((GdkSeat *)p); }
*/
import "C"
import (
	"unsafe"
)

// DeviceManager is a representation of GDK's GdkDeviceManager.
type DeviceManager struct {
	*Object
}

// native returns a pointer to the underlying GdkDeviceManager.
func (v *DeviceManager) native() *C.GdkDeviceManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDeviceManager(p)
}

// Native returns a pointer to the underlying GdkDeviceManager.
func (v *DeviceManager) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDeviceManager(p uintptr) (any, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	return &DeviceManager{obj}, nil
}

// GetDisplay is a wrapper around gdk_device_manager_get_display().
func (v *DeviceManager) GetDisplay() (*Display, error) {
	c := C.gdk_device_manager_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{ToGoObject(unsafe.Pointer(c))}, nil
}

// Display is a representation of GDK's GdkDisplay.
type Display struct {
	*Object
}

// native returns a pointer to the underlying GdkDisplay.
func (v *Display) native() *C.GdkDisplay {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDisplay(p)
}

// Native returns a pointer to the underlying GdkDisplay.
func (v *Display) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDisplay(p uintptr) (any, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	return &Display{obj}, nil
}

func toDisplay(s *C.GdkDisplay) (*Display, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &Object{ToCObject(unsafe.Pointer(s))}
	return &Display{obj}, nil
}

// DisplayOpen is a wrapper around gdk_display_open().
func DisplayOpen(displayName string) (*Display, error) {
	cstr := C.CString(displayName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_display_open((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{ToGoObject(unsafe.Pointer(c))}, nil
}

// DisplayGetDefault is a wrapper around gdk_display_get_default().
func DisplayGetDefault() *Display {
	c := C.gdk_display_get_default()
	if c == nil {
		return nil
	}
	return &Display{ToGoObject(unsafe.Pointer(c))}
}

// GetName is a wrapper around gdk_display_get_name().
func (v *Display) GetName() (string, error) {
	c := C.gdk_display_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetDefaultScreen is a wrapper around gdk_display_get_default_screen().
func (v *Display) GetDefaultScreen() *Screen {
	c := C.gdk_display_get_default_screen(v.native())
	if c == nil {
		return nil
	}
	return &Screen{ToGoObject(unsafe.Pointer(c))}
}

// DeviceIsGrabbed is a wrapper around gdk_display_device_is_grabbed().
func (v *Display) DeviceIsGrabbed(device *Device) bool {
	c := C.gdk_display_device_is_grabbed(v.native(), device.native())
	return GoBool(c)
}

// Beep is a wrapper around gdk_display_beep().
func (v *Display) Beep() {
	C.gdk_display_beep(v.native())
}

// Sync is a wrapper around gdk_display_sync().
func (v *Display) Sync() {
	C.gdk_display_sync(v.native())
}

// Flush is a wrapper around gdk_display_flush().
func (v *Display) Flush() {
	C.gdk_display_flush(v.native())
}

// Close is a wrapper around gdk_display_close().
func (v *Display) Close() {
	C.gdk_display_close(v.native())
}

// IsClosed is a wrapper around gdk_display_is_closed().
func (v *Display) IsClosed() bool {
	c := C.gdk_display_is_closed(v.native())
	return GoBool(c)
}

// GetEvent is a wrapper around gdk_display_get_event().
func (v *Display) GetEvent() (*Event, error) {
	c := C.gdk_display_get_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	//The finalizer is not on the Object but on the event.
	e := &Event{c}
	return e, nil
}

// PeekEvent is a wrapper around gdk_display_peek_event().
func (v *Display) PeekEvent() (*Event, error) {
	c := C.gdk_display_peek_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	//The finalizer is not on the Object but on the event.
	e := &Event{c}
	return e, nil
}

// PutEvent is a wrapper around gdk_display_put_event().
func (v *Display) PutEvent(event *Event) {
	C.gdk_display_put_event(v.native(), event.native())
}

// HasPending is a wrapper around gdk_display_has_pending().
func (v *Display) HasPending() bool {
	c := C.gdk_display_has_pending(v.native())
	return GoBool(c)
}

// SetDoubleClickTime is a wrapper around gdk_display_set_double_click_time().
func (v *Display) SetDoubleClickTime(msec uint) {
	C.gdk_display_set_double_click_time(v.native(), C.guint(msec))
}

// SetDoubleClickDistance is a wrapper around gdk_display_set_double_click_distance().
func (v *Display) SetDoubleClickDistance(distance uint) {
	C.gdk_display_set_double_click_distance(v.native(), C.guint(distance))
}

// SupportsColorCursor is a wrapper around gdk_display_supports_cursor_color().
func (v *Display) SupportsColorCursor() bool {
	c := C.gdk_display_supports_cursor_color(v.native())
	return GoBool(c)
}

// SupportsCursorAlpha is a wrapper around gdk_display_supports_cursor_alpha().
func (v *Display) SupportsCursorAlpha() bool {
	c := C.gdk_display_supports_cursor_alpha(v.native())
	return GoBool(c)
}

// GetDefaultCursorSize is a wrapper around gdk_display_get_default_cursor_size().
func (v *Display) GetDefaultCursorSize() uint {
	c := C.gdk_display_get_default_cursor_size(v.native())
	return uint(c)
}

// GetMaximalCursorSize is a wrapper around gdk_display_get_maximal_cursor_size().
func (v *Display) GetMaximalCursorSize() (width, height uint) {
	var w, h C.guint
	C.gdk_display_get_maximal_cursor_size(v.native(), &w, &h)
	return uint(w), uint(h)
}

// GetDefaultGroup is a wrapper around gdk_display_get_default_group().
func (v *Display) GetDefaultGroup() *Window {
	c := C.gdk_display_get_default_group(v.native())
	if c == nil {
		return nil
	}
	window := new(Window)
	window.Object = ToGoObject(unsafe.Pointer(c))
	return window
}

// SupportsSelectionNotification is a wrapper around gdk_display_supports_selection_notification().
func (v *Display) SupportsSelectionNotification() bool {
	c := C.gdk_display_supports_selection_notification(v.native())
	return GoBool(c)
}

// RequestSelectionNotification is a wrapper around gdk_display_request_selection_notification().
func (v *Display) RequestSelectionNotification(selection Atom) bool {
	c := C.gdk_display_request_selection_notification(v.native(), selection.native())
	return GoBool(c)
}

// SupportsClipboardPersistence is a wrapper around gdk_display_supports_clipboard_persistence().
func (v *Display) SupportsClipboardPersistence() bool {
	c := C.gdk_display_supports_clipboard_persistence(v.native())
	return GoBool(c)
}

// TODO:
// gdk_display_store_clipboard().
// func (v *Display) StoreClipboard(clipboardWindow *Window, time uint32, targets ...Atom) {
// 	panic("Not implemented")
// }

// SupportsShapes is a wrapper around gdk_display_supports_shapes().
func (v *Display) SupportsShapes() bool {
	c := C.gdk_display_supports_shapes(v.native())
	return GoBool(c)
}

// SupportsInputShapes is a wrapper around gdk_display_supports_input_shapes().
func (v *Display) SupportsInputShapes() bool {
	c := C.gdk_display_supports_input_shapes(v.native())
	return GoBool(c)
}

// NotifyStartupComplete is a wrapper around gdk_display_notify_startup_complete().
func (v *Display) NotifyStartupComplete(startupID string) {
	cstr := C.CString(startupID)
	defer C.free(unsafe.Pointer(cstr))
	C.gdk_display_notify_startup_complete(v.native(), (*C.gchar)(cstr))
}

// GetDeviceManager is a wrapper around gdk_display_get_device_manager().
//func (v *Display) GetDeviceManager() *DeviceManager {
//	c := C.gdk_display_get_device_manager(v.native())
//	if c == nil {
//		return nil
//	}
//	return &DeviceManager{ToGoObject(unsafe.Pointer(c))}
//}

// GetScreen is a wrapper around gdk_display_get_screen().
//func (v *Display) GetScreen(screenNum int) (*Screen, error) {
//	c := C.gdk_display_get_screen(v.native(), C.gint(screenNum))
//	if c == nil {
//		return nil, nilPtrErr
//	}
//
//	return &Screen{ToGoObject(unsafe.Pointer(c))}, nil
//}

// DisplayManager is a representation of GDK's GdkDisplayManager.
type DisplayManager struct {
	*Object
}

// native returns a pointer to the underlying GdkDisplayManager.
func (v *DisplayManager) native() *C.GdkDisplayManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDisplayManager(p)
}

// Native returns a pointer to the underlying GdkDisplayManager.
func (v *DisplayManager) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDisplayManager(p uintptr) (any, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	return &DisplayManager{obj}, nil
}

func wrapDisplayManager(obj *Object) *DisplayManager {
	if obj == nil {
		return nil
	}
	return &DisplayManager{obj}
}

// DisplayManagerGet is a wrapper around gdk_display_manager_get().
func DisplayManagerGet() (*DisplayManager, error) {
	c := C.gdk_display_manager_get()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	return &DisplayManager{obj}, nil
}

// GetDefaultDisplay is a wrapper around gdk_display_manager_get_default_display().
func (v *DisplayManager) GetDefaultDisplay() (*Display, error) {
	c := C.gdk_display_manager_get_default_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	return &Display{obj}, nil
}

// SetDefaultDisplay is a wrapper around gdk_display_manager_set_default_display().
func (v *DisplayManager) SetDefaultDisplay(display *Display) {
	C.gdk_display_manager_set_default_display(v.native(), display.native())
}

// OpenDisplay is a representation of gdk_display_manager_open_display().
func (v *DisplayManager) OpenDisplay(name string) (*Display, error) {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))

	c := C.gdk_display_manager_open_display(v.native(), cstr)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	return &Display{obj}, nil
}

// GetClientPointer is a wrapper around gdk_device_manager_get_client_pointer().
//func (v *DeviceManager) GetClientPointer() *Device {
//	c := C.gdk_device_manager_get_client_pointer(v.native())
//	if c == nil {
//		return nil
//	}
//
//	return &Device{ToGoObject(unsafe.Pointer(c))}
//}

func (v *Display) GetDefaultSeat() *Seat {
	return toSeat(C.gdk_display_get_default_seat(v.native()))
}

type Seat struct {
	*Object
}

func (v *Seat) native() *C.GdkSeat {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkSeat(p)
}

// Native returns a pointer to the underlying GdkCursor.
func (v *Seat) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func toSeat(s *C.GdkSeat) *Seat {
	if s == nil {
		return nil
	}
	obj := &Object{ToCObject(unsafe.Pointer(s))}
	return &Seat{obj}
}

func (v *Seat) GetPointer() *Device {
	return toDevice(C.gdk_seat_get_pointer(v.native()))
}
