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

// RGBA is a wrapper around GdkRGBA.
type RGBA struct {
	rgba *GdkRGBA
}

// NewRGBA creates a new RGBA with the given values (r, g, b, a).
func NewRGBA(values ...float64) *RGBA {
	rgba := new(GdkRGBA)
	if len(values) >= 1 {
		rgba.Red = values[0]
	}
	if len(values) >= 2 {
		rgba.Green = values[1]
	}
	if len(values) >= 3 {
		rgba.Blue = values[2]
	}
	if len(values) >= 4 {
		rgba.Alpha = values[3]
	}
	return &RGBA{rgba: rgba}
}

// WrapRGBA wraps a raw pointer as RGBA.
func WrapRGBA(p unsafe.Pointer) *RGBA {
	if p == nil {
		return nil
	}
	return &RGBA{rgba: (*GdkRGBA)(p)}
}

// Native returns the underlying pointer.
func (c *RGBA) Native() uintptr {
	if c == nil || c.rgba == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(c.rgba))
}

// Floats returns [r, g, b, a] as a slice.
func (c *RGBA) Floats() []float64 {
	return []float64{c.rgba.Red, c.rgba.Green, c.rgba.Blue, c.rgba.Alpha}
}

// SetColors sets all four color components.
func (c *RGBA) SetColors(r, g, b, a float64) {
	c.rgba.Red = r
	c.rgba.Green = g
	c.rgba.Blue = b
	c.rgba.Alpha = a
}

// GetRed returns the red component.
func (c *RGBA) GetRed() float64 { return c.rgba.Red }

// GetGreen returns the green component.
func (c *RGBA) GetGreen() float64 { return c.rgba.Green }

// GetBlue returns the blue component.
func (c *RGBA) GetBlue() float64 { return c.rgba.Blue }

// GetAlpha returns the alpha component.
func (c *RGBA) GetAlpha() float64 { return c.rgba.Alpha }

// SetRed sets the red component.
func (c *RGBA) SetRed(red float64) { c.rgba.Red = red }

// SetGreen sets the green component.
func (c *RGBA) SetGreen(green float64) { c.rgba.Green = green }

// SetBlue sets the blue component.
func (c *RGBA) SetBlue(blue float64) { c.rgba.Blue = blue }

// SetAlpha sets the alpha component.
func (c *RGBA) SetAlpha(alpha float64) { c.rgba.Alpha = alpha }

// Parse is a wrapper around gdk_rgba_parse().
func (c *RGBA) Parse(spec string) bool {
	cstr := CStr(spec)
	r := gdk3.SysCall("gdk_rgba_parse", uintptr(unsafe.Pointer(c.rgba)), cstr)
	return ToGoBool(r)
}

// String is a wrapper around gdk_rgba_to_string().
func (c *RGBA) String() string {
	r := gdk3.SysCall("gdk_rgba_to_string", uintptr(unsafe.Pointer(c.rgba)))
	return GoStr(r)
}

// Free is a wrapper around gdk_rgba_free().
func (c *RGBA) Free() {
	gdk3.SysCall("gdk_rgba_free", uintptr(unsafe.Pointer(c.rgba)))
}

// Equal is a wrapper around gdk_rgba_equal().
func (c *RGBA) Equal(rgba *RGBA) bool {
	r := gdk3.SysCall("gdk_rgba_equal",
		uintptr(unsafe.Pointer(c.rgba)), uintptr(unsafe.Pointer(rgba.rgba)))
	return ToGoBool(r)
}

// Hash is a wrapper around gdk_rgba_hash().
func (c *RGBA) Hash() uint {
	r := gdk3.SysCall("gdk_rgba_hash", uintptr(unsafe.Pointer(c.rgba)))
	return uint(r)
}
