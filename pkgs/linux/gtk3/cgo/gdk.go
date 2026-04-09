package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

// WindowEdge is a representation of GDK's GdkWindowEdge
type WindowEdge int

const (
	WINDOW_EDGE_NORTH_WEST WindowEdge = C.GDK_WINDOW_EDGE_NORTH_WEST
	WINDOW_EDGE_NORTH      WindowEdge = C.GDK_WINDOW_EDGE_NORTH
	WINDOW_EDGE_NORTH_EAST WindowEdge = C.GDK_WINDOW_EDGE_NORTH_EAST
	WINDOW_EDGE_WEST       WindowEdge = C.GDK_WINDOW_EDGE_WEST
	WINDOW_EDGE_EAST       WindowEdge = C.GDK_WINDOW_EDGE_EAST
	WINDOW_EDGE_SOUTH_WEST WindowEdge = C.GDK_WINDOW_EDGE_SOUTH_WEST
	WINDOW_EDGE_SOUTH      WindowEdge = C.GDK_WINDOW_EDGE_SOUTH
	WINDOW_EDGE_SOUTH_EAST WindowEdge = C.GDK_WINDOW_EDGE_SOUTH_EAST
)

const CURRENT_TIME = C.GDK_CURRENT_TIME

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

// VisualType is a representation of GDK's GdkVisualType.
type VisualType int

const (
	VISUAL_STATIC_GRAY  VisualType = C.GDK_VISUAL_STATIC_GRAY
	VISUAL_GRAYSCALE    VisualType = C.GDK_VISUAL_GRAYSCALE
	VISUAL_STATIC_COLOR VisualType = C.GDK_VISUAL_STATIC_COLOR
	ISUAL_PSEUDO_COLOR  VisualType = C.GDK_VISUAL_PSEUDO_COLOR
	VISUAL_TRUE_COLOR   VisualType = C.GDK_VISUAL_TRUE_COLOR
	VISUAL_DIRECT_COLOR VisualType = C.GDK_VISUAL_DIRECT_COLOR
)

func marshalVisualType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return VisualType(c), nil
}

func marshalEventType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return types.EventType(c), nil
}

// added by terrak
// EventMask is a representation of GDK's GdkEventMask.
type EventMask int

const (
	EXPOSURE_MASK            EventMask = C.GDK_EXPOSURE_MASK
	POINTER_MOTION_MASK      EventMask = C.GDK_POINTER_MOTION_MASK
	POINTER_MOTION_HINT_MASK EventMask = C.GDK_POINTER_MOTION_HINT_MASK
	BUTTON_MOTION_MASK       EventMask = C.GDK_BUTTON_MOTION_MASK
	BUTTON1_MOTION_MASK      EventMask = C.GDK_BUTTON1_MOTION_MASK
	BUTTON2_MOTION_MASK      EventMask = C.GDK_BUTTON2_MOTION_MASK
	BUTTON3_MOTION_MASK      EventMask = C.GDK_BUTTON3_MOTION_MASK
	BUTTON_PRESS_MASK        EventMask = C.GDK_BUTTON_PRESS_MASK
	BUTTON_RELEASE_MASK      EventMask = C.GDK_BUTTON_RELEASE_MASK
	KEY_PRESS_MASK           EventMask = C.GDK_KEY_PRESS_MASK
	KEY_RELEASE_MASK         EventMask = C.GDK_KEY_RELEASE_MASK
	ENTER_NOTIFY_MASK        EventMask = C.GDK_ENTER_NOTIFY_MASK
	LEAVE_NOTIFY_MASK        EventMask = C.GDK_LEAVE_NOTIFY_MASK
	FOCUS_CHANGE_MASK        EventMask = C.GDK_FOCUS_CHANGE_MASK
	STRUCTURE_MASK           EventMask = C.GDK_STRUCTURE_MASK
	PROPERTY_CHANGE_MASK     EventMask = C.GDK_PROPERTY_CHANGE_MASK
	VISIBILITY_NOTIFY_MASK   EventMask = C.GDK_VISIBILITY_NOTIFY_MASK
	PROXIMITY_IN_MASK        EventMask = C.GDK_PROXIMITY_IN_MASK
	PROXIMITY_OUT_MASK       EventMask = C.GDK_PROXIMITY_OUT_MASK
	SUBSTRUCTURE_MASK        EventMask = C.GDK_SUBSTRUCTURE_MASK
	SCROLL_MASK              EventMask = C.GDK_SCROLL_MASK
	TOUCH_MASK               EventMask = C.GDK_TOUCH_MASK
	SMOOTH_SCROLL_MASK       EventMask = C.GDK_SMOOTH_SCROLL_MASK
	ALL_EVENTS_MASK          EventMask = C.GDK_ALL_EVENTS_MASK
)
