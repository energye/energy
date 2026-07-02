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

type Visual struct {
	Object
}

func AsVisual(ptr unsafe.Pointer) IVisual {
	if ptr == nil {
		return nil
	}
	m := new(Visual)
	m.instance = ptr
	return m
}

// GetBluePixelDetails is a wrapper around gdk_visual_get_blue_pixel_details().
func (m *Visual) GetBluePixelDetails() (*uint32, *int, *int) {
	var mask uint32
	var shift, precision int32
	gdk3.SysCall("gdk_visual_get_blue_pixel_details", m.Instance(),
		uintptr(unsafe.Pointer(&mask)), uintptr(unsafe.Pointer(&shift)), uintptr(unsafe.Pointer(&precision)))
	s := int(shift)
	p := int(precision)
	return &mask, &s, &p
}

// GetDepth is a wrapper around gdk_visual_get_depth().
func (m *Visual) GetDepth() int {
	r := gdk3.SysCall("gdk_visual_get_depth", m.Instance())
	return int(r)
}

// GetGreenPixelDetails is a wrapper around gdk_visual_get_green_pixel_details().
func (m *Visual) GetGreenPixelDetails() (*uint32, *int, *int) {
	var mask uint32
	var shift, precision int32
	gdk3.SysCall("gdk_visual_get_green_pixel_details", m.Instance(),
		uintptr(unsafe.Pointer(&mask)), uintptr(unsafe.Pointer(&shift)), uintptr(unsafe.Pointer(&precision)))
	s := int(shift)
	p := int(precision)
	return &mask, &s, &p
}

// GetRedPixelDetails is a wrapper around gdk_visual_get_red_pixel_details().
func (m *Visual) GetRedPixelDetails() (*uint32, *int, *int) {
	var mask uint32
	var shift, precision int32
	gdk3.SysCall("gdk_visual_get_red_pixel_details", m.Instance(),
		uintptr(unsafe.Pointer(&mask)), uintptr(unsafe.Pointer(&shift)), uintptr(unsafe.Pointer(&precision)))
	s := int(shift)
	p := int(precision)
	return &mask, &s, &p
}

// GetVisualType is a wrapper around gdk_visual_get_visual_type().
func (m *Visual) GetVisualType() int {
	r := gdk3.SysCall("gdk_visual_get_visual_type", m.Instance())
	return int(r)
}

// GetScreen is a wrapper around gdk_visual_get_screen().
func (m *Visual) GetScreen() IScreen {
	r := gdk3.SysCall("gdk_visual_get_screen", m.Instance())
	if r == 0 {
		return nil
	}
	return AsScreen(unsafe.Pointer(r))
}
