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

var gdkFloatOnce sync.Once

var (
	gdkScreenGetResolution func(uintptr) float64
	gdkScreenSetResolution func(uintptr, float64)
)

func registerGdkFloatFuncs() {
	gdkFloatOnce.Do(func() {
		if gdk3 == nil || gdk3.Dll == 0 {
			return
		}
		lib := uintptr(gdk3.Dll)
		purego.RegisterLibFunc(&gdkScreenGetResolution, lib, "gdk_screen_get_resolution")
		purego.RegisterLibFunc(&gdkScreenSetResolution, lib, "gdk_screen_set_resolution")
	})
}
