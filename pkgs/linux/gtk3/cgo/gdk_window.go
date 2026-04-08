package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"unsafe"
)

// GdkWindow is a representation of GDK's GdkWindow.
type GdkWindow struct {
	*Object
}

// native returns a pointer to the underlying GdkWindow.
func (v *GdkWindow) native() *C.GdkWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkWindow(p)
}

// Native returns a pointer to the underlying GdkWindow.
func (v *GdkWindow) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// WindowGetWidth is a wrapper around gdk_window_get_width()
func (v *GdkWindow) WindowGetWidth() (width int) {
	return int(C.gdk_window_get_width(v.native()))
}

// WindowGetHeight is a wrapper around gdk_window_get_height()
func (v *GdkWindow) WindowGetHeight() (height int) {
	return int(C.gdk_window_get_height(v.native()))
}

// GetRootOrigin is a wrapper around gdk_window_get_root_origin()
func (v *GdkWindow) GetRootOrigin() (x int, y int) {
	var cX C.gint
	var cY C.gint
	C.gdk_window_get_root_origin(v.native(), &cX, &cY)
	return int(cX), int(cY)
}

// GetOrigin is a wrapper around gdk_window_get_origin
func (v *GdkWindow) GetOrigin() (x int, y int) {
	var cX C.gint
	var cY C.gint
	C.gdk_window_get_origin(v.native(), &cX, &cY)
	return int(cX), int(cY)
}

// GetDevicePosition is a wrapper around gdk_window_get_device_position()
func (v *GdkWindow) GetDevicePosition(d *Device) (*GdkWindow, int, int, ModifierType) {
	var x C.gint
	var y C.gint
	var mt C.GdkModifierType
	underneathWindow := C.gdk_window_get_device_position(v.native(), d.native(), &x, &y, &mt)
	obj := &Object{ToCObject(unsafe.Pointer(underneathWindow))}
	rw := &GdkWindow{obj}
	return rw, int(x), int(y), ModifierType(mt)
}

// SetOverrideRedirect is a wrapper around gdk_window_set_override_redirect().
func (v *GdkWindow) SetOverrideRedirect(overrideRedirect bool) {
	C.gdk_window_set_override_redirect(v.native(), CBool(overrideRedirect))
}

func toGdkWindow(s *C.GdkWindow) *GdkWindow {
	if s == nil {
		return nil
	}
	obj := &Object{ToCObject(unsafe.Pointer(s))}
	return &GdkWindow{obj}
}
