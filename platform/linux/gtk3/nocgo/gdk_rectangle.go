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

import "unsafe"

type Rectangle struct {
	x      int32
	y      int32
	width  int32
	height int32
}

// NewRectangle helper function to create a GdkRectangle.
func NewRectangle(x, y, width, height int) *Rectangle {
	return &Rectangle{
		x:      int32(x),
		y:      int32(y),
		width:  int32(width),
		height: int32(height),
	}
}

// SetRectangleInt helper function to set GdkRectangle values.
func (m *Rectangle) SetRectangleInt(x, y, width, height int) {
	m.x = int32(x)
	m.y = int32(y)
	m.width = int32(width)
	m.height = int32(height)
}

// GetRectangleInt helper function to get GdkRectangle values.
func (m *Rectangle) GetRectangleInt() (x, y, width, height int) {
	return int(m.x), int(m.y), int(m.width), int(m.height)
}

// RectangleIntersect is a wrapper around gdk_rectangle_intersect().
func (m *Rectangle) RectangleIntersect(rect *Rectangle) (*Rectangle, bool) {
	var result Rectangle
	c := gdk3.SysCall("gdk_rectangle_intersect",
		uintptr(unsafe.Pointer(m)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(&result)),
	)
	return &result, ToGoBool(c)
}

// RectangleUnion is a wrapper around gdk_rectangle_union().
func (m *Rectangle) RectangleUnion(rect *Rectangle) *Rectangle {
	var result Rectangle
	gdk3.SysCall("gdk_rectangle_union",
		uintptr(unsafe.Pointer(m)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(&result)),
	)
	return &result
}

// GetX returns x field of the underlying GdkRectangle.
func (m *Rectangle) GetX() int {
	return int(m.x)
}

// SetX sets x field of the underlying GdkRectangle.
func (m *Rectangle) SetX(x int) {
	m.x = int32(x)
}

// GetY returns y field of the underlying GdkRectangle.
func (m *Rectangle) GetY() int {
	return int(m.y)
}

// SetY sets y field of the underlying GdkRectangle.
func (m *Rectangle) SetY(y int) {
	m.y = int32(y)
}

// GetWidth returns width field of the underlying GdkRectangle.
func (m *Rectangle) GetWidth() int {
	return int(m.width)
}

// SetWidth sets width field of the underlying GdkRectangle.
func (m *Rectangle) SetWidth(width int) {
	m.width = int32(width)
}

// GetHeight returns height field of the underlying GdkRectangle.
func (m *Rectangle) GetHeight() int {
	return int(m.height)
}

// SetHeight sets height field of the underlying GdkRectangle.
func (m *Rectangle) SetHeight(height int) {
	m.height = int32(height)
}
