//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "github.com/energye/energy/v2/api"

// NewCefRect
func NewCefRect(x, y, width, height int32) *TCefRect {
	return &TCefRect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

// TRGBQuad
type TRGBQuad struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

func (m *TCefKeyEvent) KeyDown() bool {
	return m.Kind == KEYEVENT_RAWKEYDOWN || m.Kind == KEYEVENT_KEYDOWN
}

func (m *TCefKeyEvent) KeyUp() bool {
	return m.Kind == KEYEVENT_KEYUP
}

// TBeforePopup 弹出窗口, 参数字段包装结构
type TBeforePopup struct {
	TargetUrl         string
	TargetFrameName   string
	TargetDisposition TCefWindowOpenDisposition
	UserGesture       bool
}

type tBeforePopup struct {
	TargetUrl         uintptr
	TargetFrameName   uintptr
	TargetDisposition uintptr
	UserGesture       uintptr
}

func (m *tBeforePopup) Convert() *TBeforePopup {
	if m == nil {
		return nil
	}
	return &TBeforePopup{
		TargetUrl:         api.GoStr(m.TargetUrl),
		TargetFrameName:   api.GoStr(m.TargetFrameName),
		TargetDisposition: *(*TCefWindowOpenDisposition)(unsafePointer(m.TargetDisposition)),
		UserGesture:       *(*bool)(unsafePointer(m.UserGesture)),
	}
}
