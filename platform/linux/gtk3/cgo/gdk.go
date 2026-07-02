package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// RGBA To create a GdkRGBA you have to use NewRGBA function.
type RGBA struct {
	rgba *C.GdkRGBA
}

func marshalRGBA(p uintptr) (any, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRGBA(unsafe.Pointer(c)), nil
}

func WrapRGBA(p unsafe.Pointer) *RGBA {
	return wrapRGBA((*C.GdkRGBA)(p))
}

func wrapRGBA(cRgba *C.GdkRGBA) *RGBA {
	if cRgba == nil {
		return nil
	}
	return &RGBA{cRgba}
}

func NewRGBA(values ...float64) *RGBA {
	cRgba := new(C.GdkRGBA)
	for i, value := range values {
		switch i {
		case 0:
			cRgba.red = C.gdouble(value)
		case 1:
			cRgba.green = C.gdouble(value)
		case 2:
			cRgba.blue = C.gdouble(value)
		case 3:
			cRgba.alpha = C.gdouble(value)
		}
	}
	return wrapRGBA(cRgba)
}

func (c *RGBA) Floats() []float64 {
	return []float64{
		float64(c.rgba.red),
		float64(c.rgba.green),
		float64(c.rgba.blue),
		float64(c.rgba.alpha)}
}

func (c *RGBA) Native() uintptr {
	return uintptr(unsafe.Pointer(c.rgba))
}

// SetColors sets all colors values in the RGBA.
func (c *RGBA) SetColors(r, g, b, a float64) {
	c.rgba.red = C.gdouble(r)
	c.rgba.green = C.gdouble(g)
	c.rgba.blue = C.gdouble(b)
	c.rgba.alpha = C.gdouble(a)
}

/*
GetRed
  The following methods (Get/Set) are made for
  more convenient use of the GdkRGBA object
*/
// GetRed get red value from the RGBA.
func (c *RGBA) GetRed() float64 {
	return float64(c.rgba.red)
}

// GetGreen get green value from the RGBA.
func (c *RGBA) GetGreen() float64 {
	return float64(c.rgba.green)
}

// GetBlue get blue value from the RGBA.
func (c *RGBA) GetBlue() float64 {
	return float64(c.rgba.blue)
}

// GetAlpha get alpha value from the RGBA.
func (c *RGBA) GetAlpha() float64 {
	return float64(c.rgba.alpha)
}

// SetRed set red value in the RGBA.
func (c *RGBA) SetRed(red float64) {
	c.rgba.red = C.gdouble(red)
}

// SetGreen set green value in the RGBA.
func (c *RGBA) SetGreen(green float64) {
	c.rgba.green = C.gdouble(green)
}

// SetBlue set blue value in the RGBA.
func (c *RGBA) SetBlue(blue float64) {
	c.rgba.blue = C.gdouble(blue)
}

// SetAlpha set alpha value in the RGBA.
func (c *RGBA) SetAlpha(alpha float64) {
	c.rgba.alpha = C.gdouble(alpha)
}

// Parse is a representation of gdk_rgba_parse().
func (c *RGBA) Parse(spec string) bool {
	cstr := (*C.gchar)(C.CString(spec))
	defer C.free(unsafe.Pointer(cstr))
	return GoBool(C.gdk_rgba_parse(c.rgba, cstr))
}

// String is a representation of gdk_rgba_to_string().
func (c *RGBA) String() string {
	return C.GoString((*C.char)(C.gdk_rgba_to_string(c.rgba)))
}

// Free is a representation of gdk_rgba_free().
func (c *RGBA) Free() {
	C.gdk_rgba_free(c.rgba)
}

// Equal is a representation of gdk_rgba_equal().
func (c *RGBA) Equal(rgba *RGBA) bool {
	return GoBool(C.gdk_rgba_equal(
		C.gconstpointer(c.rgba),
		C.gconstpointer(rgba.rgba)))
}

// Hash is a representation of gdk_rgba_hash().
func (c *RGBA) Hash() uint {
	return uint(C.gdk_rgba_hash(C.gconstpointer(c.rgba)))
}

func marshalVisualType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return VisualType(c), nil
}

func marshalEventType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EventType(c), nil
}

// added by terrak
// EventMask is now defined in types/gtk3_types.go and shared via import.
