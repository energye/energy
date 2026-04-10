package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/linux/types"
	"unsafe"
)

// Visual is a representation of GDK's GdkVisual.
type Visual struct {
	*Object
}

func (v *Visual) native() *C.GdkVisual {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkVisual(p)
}

func (v *Visual) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// GetBluePixelDetails is a wrapper around gdk_visual_get_blue_pixel_details().
func (v *Visual) GetBluePixelDetails() (*uint32, *int, *int) {
	var (
		m                *uint32 = nil
		s, p             *int    = nil, nil
		mask             C.guint32
		shift, precision C.gint
	)
	C.gdk_visual_get_blue_pixel_details(v.native(), &mask, &shift, &precision)
	if &mask != nil {
		m = new(uint32)
		*m = uint32(mask)
	}
	if &shift != nil {
		s = new(int)
		*s = int(shift)
	}
	if &precision != nil {
		p = new(int)
		*p = int(precision)
	}
	return m, s, p
}

// GetDepth is a wrapper around gdk_visual_get_depth().
func (v *Visual) GetDepth() int {
	return int(C.gdk_visual_get_depth(v.native()))
}

// GetGreenPixelDetails is a wrapper around gdk_visual_get_green_pixel_details().
func (v *Visual) GetGreenPixelDetails() (*uint32, *int, *int) {
	var (
		m                *uint32 = nil
		s, p             *int    = nil, nil
		mask             C.guint32
		shift, precision C.gint
	)
	C.gdk_visual_get_green_pixel_details(v.native(), &mask, &shift, &precision)
	if &mask != nil {
		m = new(uint32)
		*m = uint32(mask)
	}
	if &shift != nil {
		s = new(int)
		*s = int(shift)
	}
	if &precision != nil {
		p = new(int)
		*p = int(precision)
	}
	return m, s, p
}

// GetRedPixelDetails is a wrapper around gdk_visual_get_red_pixel_details().
func (v *Visual) GetRedPixelDetails() (*uint32, *int, *int) {
	var (
		m                *uint32 = nil
		s, p             *int    = nil, nil
		mask             C.guint32
		shift, precision C.gint
	)
	C.gdk_visual_get_red_pixel_details(v.native(), &mask, &shift, &precision)
	if &mask != nil {
		m = new(uint32)
		*m = uint32(mask)
	}
	if &shift != nil {
		s = new(int)
		*s = int(shift)
	}
	if &precision != nil {
		p = new(int)
		*p = int(precision)
	}
	return m, s, p
}

// GetVisualType is a wrapper around gdk_visual_get_visual_type().
func (v *Visual) GetVisualType() VisualType {
	return VisualType(C.gdk_visual_get_visual_type(v.native()))
}

// GetScreen is a wrapper around gdk_visual_get_screen().
func (v *Visual) GetScreen() IScreen {
	return toScreen(C.gdk_visual_get_screen(v.native()))
}
