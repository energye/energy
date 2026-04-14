package cgo

// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Screen is a representation of GDK's GdkScreen.
type Screen struct {
	*Object
}

func toScreen(s *C.GdkScreen) IScreen {
	if s == nil {
		return nil
	}
	return &Screen{ToGoObject(unsafe.Pointer(s))}
}

// ScreenGetDefault is a wrapper around gdk_screen_get_default().
func ScreenGetDefault() IScreen {
	return toScreen(C.gdk_screen_get_default())
}

// native returns a pointer to the underlying GdkScreen.
func (v *Screen) native() *C.GdkScreen {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkScreen(p)
}

// Native returns a pointer to the underlying GdkScreen.
func (v *Screen) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// GetRGBAVisual is a wrapper around gdk_screen_get_rgba_visual().
func (v *Screen) GetRGBAVisual() IVisual {
	c := C.gdk_screen_get_rgba_visual(v.native())
	if c == nil {
		return nil
	}

	return &Visual{ToGoObject(unsafe.Pointer(c))}
}

// GetSystemVisual is a wrapper around gdk_screen_get_system_visual().
func (v *Screen) GetSystemVisual() (*Visual, error) {
	c := C.gdk_screen_get_system_visual(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Visual{ToGoObject(unsafe.Pointer(c))}, nil
}

// IsComposited is a wrapper around gdk_screen_is_composited().
func (v *Screen) IsComposited() bool {
	return GoBool(C.gdk_screen_is_composited(v.native()))
}

// GetRootWindow is a wrapper around gdk_screen_get_root_window().
func (v *Screen) GetRootWindow() *GdkWindow {
	return toGdkWindow(C.gdk_screen_get_root_window(v.native()))
}

// GetDisplay is a wrapper around gdk_screen_get_display().
func (v *Screen) GetDisplay() (*Display, error) {
	return toDisplay(C.gdk_screen_get_display(v.native()))
}

func toString(c *C.gchar) (string, error) {
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetResolution is a wrapper around gdk_screen_get_resolution().
func (v *Screen) GetResolution() float64 {
	return float64(C.gdk_screen_get_resolution(v.native()))
}

// SetResolution is a wrapper around gdk_screen_set_resolution().
func (v *Screen) SetResolution(r float64) {
	C.gdk_screen_set_resolution(v.native(), C.gdouble(r))
}
