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
	"sync"

	"github.com/ebitengine/purego"
)

var cairoFloatOnce sync.Once

var (
	cairoSetSourceRGB   func(uintptr, float64, float64, float64)
	cairoSetSourceRGBA  func(uintptr, float64, float64, float64, float64)
	cairoSetLineWidth   func(uintptr, float64)
	cairoGetLineWidth   func(uintptr) float64
	cairoRectangle      func(uintptr, float64, float64, float64, float64)
	cairoArc            func(uintptr, float64, float64, float64, float64, float64)
	cairoLineTo         func(uintptr, float64, float64)
	cairoMoveTo         func(uintptr, float64, float64)
	cairoPaintWithAlpha func(uintptr, float64)
)

func registerCairoFloatFuncs() {
	cairoFloatOnce.Do(func() {
		if cairo == nil || cairo.Dll == 0 {
			return
		}
		lib := uintptr(cairo.Dll)
		purego.RegisterLibFunc(&cairoSetSourceRGB, lib, "cairo_set_source_rgb")
		purego.RegisterLibFunc(&cairoSetSourceRGBA, lib, "cairo_set_source_rgba")
		purego.RegisterLibFunc(&cairoSetLineWidth, lib, "cairo_set_line_width")
		purego.RegisterLibFunc(&cairoGetLineWidth, lib, "cairo_get_line_width")
		purego.RegisterLibFunc(&cairoRectangle, lib, "cairo_rectangle")
		purego.RegisterLibFunc(&cairoArc, lib, "cairo_arc")
		purego.RegisterLibFunc(&cairoLineTo, lib, "cairo_line_to")
		purego.RegisterLibFunc(&cairoMoveTo, lib, "cairo_move_to")
		purego.RegisterLibFunc(&cairoPaintWithAlpha, lib, "cairo_paint_with_alpha")
	})
}
