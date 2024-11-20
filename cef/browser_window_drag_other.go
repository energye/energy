//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

func (m *drag) drag() {
	window := m.window.AsLCLBrowserWindow().BrowserWindow()
	switch m.T {
	case dragUp:
	case dragDown:
		m.dx = m.X
		m.dy = m.Y
		point := m.window.Point()
		m.wx = point.X
		m.wy = point.Y
	case dragMove:
		m.mx = m.X
		m.my = m.Y
		if m.window.IsLCL() {
			x := m.wx + (m.mx - m.dx)
			y := m.wy + (m.my - m.dy)
			m.window.SetPoint(x, y)
		}
	case dragDblClick:
		if window.WindowProperty().EnableWebkitAppRegionDClk {
			window.Maximize()
		}
	case dragResize:
	}
}

// locate libgtk-x11-2.0.so
// nm -D --defined-only libgtk-x11-2.0.so | grep gtk_drag_begin
// gtkWidget := lcl.HandleToPlatformHandle(window.Handle())
// gtkWindow := gtkWidget.Window()
// fmt.Println("gtkWidget:", gtkWidget, "gtkWindow:", gtkWindow, m.TimeStamp)
// GDK_ACTION_MOVE := int32(1 << 2) //1 shl 2
// lib, err := dllimports.NewDLL("libgtk-x11-2.0.so")
// fmt.Println("lib:", lib, "err:", err)
// proc, err := lib.GetProcAddr("gtk_window_begin_move_drag")
// proc, err := lib.GetProcAddr("gtk_drag_begin")
// fmt.Println("proc:", proc, "err:", err)
// r1, r2, err := proc.Call(uintptr(gtkWidget), 0, uintptr(GDK_ACTION_MOVE), 1, uintptr(unsafe.Pointer(&event)))
// fmt.Println("r1:", r1, "r2:", r2, "err:", err)
// r1, r2, err := proc.Call(uintptr(gtkWindow), 1, uintptr(m.mx), uintptr(m.my), uintptr(m.TimeStamp))
// fmt.Println("r1:", r1, "r2:", r2, "err:", err)
// gtk_window_begin_resize_drag
