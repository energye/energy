//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package gifanim

import (
	"github.com/energye/golcl/lcl"
)

type Frame struct {
	index      int
	method     byte
	x, y       int32
	w, h       int32
	delay      uint32
	data       []byte
	background *lcl.TBitmap
	image      lcl.IBitmap
}

func (m *Frame) Index() int {
	return m.index
}

func (m *Frame) Point() (x, y int32) {
	x, y = m.x, m.y
	return
}

func (m *Frame) SetPoint(x, y int32) {
	m.x, m.y = x, y
}

func (m *Frame) Rect() (width, height int32) {
	width, height = m.w, m.h
	return
}

func (m *Frame) SetRect(width, height int32) {
	m.w, m.h = width, height
}

func (m *Frame) Delay() uint32 {
	return m.delay
}

func (m *Frame) SetDelay(delay uint32) {
	m.delay = delay
}

func (m *Frame) Data() []byte {
	return m.data
}

func (m *Frame) SetData(data []byte) {
	m.data = data
}

func (m *Frame) Image() lcl.IBitmap {
	return m.image
}

func (m *Frame) scan() {
	if m.image == nil {
		if usePNG {
			m.image = lcl.NewPngImage()
			m.image.(*lcl.TPngImage).SetSize(m.w, m.h)
			m.image.(*lcl.TPngImage).LoadFromBytes(m.data)
		} else {
			m.image = lcl.NewGIFImage()
			m.image.(*lcl.TGIFImage).SetSize(m.w, m.h)
			m.image.(*lcl.TGIFImage).LoadFromBytes(m.data)
		}
	}
}

func (m *Frame) reset() {
	if m.image != nil {
		m.image.Free()
		m.image = nil
	}
}

func (m *Frame) free() {
	m.reset()
	m.data = nil
	m.x, m.y = 0, 0
	m.w, m.h = 0, 0
	m.delay = 0
}
