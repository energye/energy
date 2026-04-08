package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/gtk3/types"
	"unsafe"
)

// Device is a representation of GDK's GdkDevice.
type Device struct {
	*Object
}

// native returns a pointer to the underlying GdkDevice.
func (v *Device) native() *C.GdkDevice {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDevice(p)
}

// Native returns a pointer to the underlying GdkDevice.
func (v *Device) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDevice(p uintptr) (any, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	return &Device{obj}, nil
}

func toDevice(d *C.GdkDevice) *Device {
	obj := &Object{ToCObject(unsafe.Pointer(d))}
	return &Device{obj}
}

func (v *Device) GetPosition(screen *IScreen, x, y *int) error {
	cs := (**C.GdkScreen)(unsafe.Pointer(uintptr(0)))
	if screen != nil {
		var cval *C.GdkScreen
		cs = &cval
	}

	cx := (*C.gint)(unsafe.Pointer(uintptr(0)))
	if x != nil {
		var cval C.gint
		cx = &cval
	}

	cy := (*C.gint)(unsafe.Pointer(uintptr(0)))
	if y != nil {
		var cval C.gint
		cy = &cval
	}
	C.gdk_device_get_position(v.native(), cs, cx, cy)

	if cs != (**C.GdkScreen)(unsafe.Pointer(uintptr(0))) {
		ms := toScreen(*cs)
		if ms == nil {
			return nilPtrErr
		}
		*screen = ms
	}
	if cx != (*C.gint)(unsafe.Pointer(uintptr(0))) {
		*x = int(*cx)
	}
	if cy != (*C.gint)(unsafe.Pointer(uintptr(0))) {
		*y = int(*cy)
	}
	return nil
}
