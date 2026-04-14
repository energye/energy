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

type Rectangle struct {
	x      int32
	y      int32
	width  int32
	height int32
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
