//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package window

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa

*/
import "C"

import (
	"github.com/energye/energy/v3/pkgs/cocoa"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"unsafe"
)

const (
	NSWindowStyleMaskBorderless     = 0 // 窗口没有标题栏和按钮
	NSWindowStyleMaskTitled         = 1 // 窗口具有标题栏
	NSWindowStyleMaskClosable       = 2 // 窗口具有关闭按钮
	NSWindowStyleMaskMiniaturizable = 4 // 窗口可以被最小化
	NSWindowStyleMaskResizable      = 8 // 窗口可以调整大小
)

type IDarwinWindow interface {
	IWindow
	NSInstance() unsafe.Pointer
	NSWindow() lcl.NSWindow
	DragWindow()
	ContentViewFrame() types.TRect
}

type TWindow struct {
	TEnergyWindow
	nsFrostedView unsafe.Pointer
	nsDelegate    unsafe.Pointer
}

func (m *TWindow) NSWindowInstance() unsafe.Pointer {
	return unsafe.Pointer(lcl.PlatformHandle(m.Handle()))
}

func (m *TWindow) NSInstance() unsafe.Pointer {
	return unsafe.Pointer(m.NSWindow())
}

func (m *TWindow) NSWindow() lcl.NSWindow {
	return lcl.PlatformWindow(m.Instance())
}

func (m *TWindow) CreateParams(params *types.TCreateParams) {
}

func (m *TWindow) _BeforeFormCreate() {
}

func (m *TWindow) _BeforeFormShow() {
	if m.flagFirstShow {
		return
	}
	m.flagFirstShow = true
	m.UpdateWindowOption()
}

func (m *TWindow) UpdateWindowOption() {
	if m.options != nil {
		if m.options.Width <= 0 {
			m.options.Width = m.Width()
		}
		if m.options.Height <= 0 {
			m.options.Height = m.Height()
		}
		m.SetCaption(m.options.Caption)
		m.SetBounds(m.options.X, m.options.Y, m.options.Width, m.options.Height)
		if m.options.MacOS.UseWindowDelegate {
			m.nsDelegate = cocoa.CreateWindowDelegate(m.NSInstance())
			m._InitEvent()
		}
		if m.options.WindowTransparent {
			m.SetWindowTransparent()
			if m.options.MacOS.AppearanceName != "" {
				m.SwitchFrostedMaterial(m.options.MacOS.AppearanceName)
			}
		}
		if m.options.BackgroundColor != nil {
			r, g, b, a := uint8(m.options.BackgroundColor.R), uint8(m.options.BackgroundColor.G), uint8(m.options.BackgroundColor.B), uint8(m.options.BackgroundColor.A)
			m.SetBackgroundColor(r, g, b, a)
		}
	}
	m.SetWindowRadius()
	m.Frameless()
	m.TitleBar()
}

func (m *TWindow) SetWindowState(value types.TWindowState) {
	m.windowsState = value
	switch value {
	case types.WsMaximized:
		cocoa.WindowExitMinimized(m.NSInstance())
		cocoa.WindowMaximize(m.NSInstance())
	case types.WsNormal:
		cocoa.WindowRestore(m.NSInstance())
	case types.WsMinimized:
		cocoa.WindowMinimized(m.NSInstance())
	}
}

func (m *TWindow) WindowState() types.TWindowState {
	return m.windowsState
}

func (m *TWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		m.SetWindowState(types.WsFullScreen)
		cocoa.WindowEnterFullScreen(m.NSInstance())
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.SetWindowState(types.WsNormal)
			cocoa.WindowExitFullScreen(m.NSInstance())
		})
	}
}
