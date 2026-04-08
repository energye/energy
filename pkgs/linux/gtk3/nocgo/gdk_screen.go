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
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

type Screen struct {
	Object
}

func AsScreen(ptr unsafe.Pointer) IScreen {
	if ptr == nil {
		return nil
	}
	m := new(Screen)
	m.instance = ptr
	return m
}

// GetRGBAVisual is a wrapper around gdk_screen_get_rgba_visual().
func (m *Screen) GetRGBAVisual() IVisual {
	r := gdk3.SysCall("gdk_screen_get_rgba_visual", m.Instance())
	if r == 0 {
		return nil
	}
	v := &Visual{}
	v.instance = unsafe.Pointer(r)
	return v
}

// IsComposited is a wrapper around gdk_screen_is_composited().
func (m *Screen) IsComposited() bool {
	r := gdk3.SysCall("gdk_screen_is_composited", m.Instance())
	return r > 0
}
