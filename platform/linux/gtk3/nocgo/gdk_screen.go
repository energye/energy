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

// GetSystemVisual is a wrapper around gdk_screen_get_system_visual().
func (m *Screen) GetSystemVisual() (*Visual, error) {
	r := gdk3.SysCall("gdk_screen_get_system_visual", m.Instance())
	if r == 0 {
		return nil, errNilPtr
	}
	v := &Visual{}
	v.instance = unsafe.Pointer(r)
	return v, nil
}

// IsComposited is a wrapper around gdk_screen_is_composited().
func (m *Screen) IsComposited() bool {
	r := gdk3.SysCall("gdk_screen_is_composited", m.Instance())
	return r > 0
}

// GetRootWindow is a wrapper around gdk_screen_get_root_window().
func (m *Screen) GetRootWindow() *GdkWindow {
	r := gdk3.SysCall("gdk_screen_get_root_window", m.Instance())
	if r == 0 {
		return nil
	}
	w := &GdkWindow{}
	w.instance = unsafe.Pointer(r)
	return w
}

// GetDisplay is a wrapper around gdk_screen_get_display().
func (m *Screen) GetDisplay() (*Display, error) {
	r := gdk3.SysCall("gdk_screen_get_display", m.Instance())
	if r == 0 {
		return nil, errNilPtr
	}
	d := &Display{}
	d.instance = unsafe.Pointer(r)
	return d, nil
}

// GetResolution is a wrapper around gdk_screen_get_resolution().
func (m *Screen) GetResolution() float64 {
	registerGdkFloatFuncs()
	return gdkScreenGetResolution(m.Instance())
}

// SetResolution is a wrapper around gdk_screen_set_resolution().
func (m *Screen) SetResolution(resolution float64) {
	registerGdkFloatFuncs()
	gdkScreenSetResolution(m.Instance(), resolution)
}

// ScreenGetDefault is a wrapper around gdk_screen_get_default().
func ScreenGetDefault() IScreen {
	r := gdk3.SysCall("gdk_screen_get_default")
	if r == 0 {
		return nil
	}
	s := &Screen{}
	s.instance = unsafe.Pointer(r)
	return s
}
