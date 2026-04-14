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

import (
	"github.com/energye/energy/v3/platform/darwin/cocoa"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"unsafe"
)

type IDarwinWindow interface {
	IWindow
	//NSInstance() unsafe.Pointer
	NSWindow() INSWindow
	//DragWindow()
	//ContentViewFrame() types.TRect
}

type TWindow struct {
	TEnergyWindow
	nsFrostedView INSVisualEffectView
	nsDelegate    INSWindowDelegate
	nsWindow      INSWindow
}

//func (m *TWindow) NSWindowInstance() unsafe.Pointer {
//	return unsafe.Pointer(lcl.PlatformHandle(m.Handle()))
//}

//func (m *TWindow) NSInstance() unsafe.Pointer {
//	return unsafe.Pointer(m.NSWindow())
//}

//func (m *TWindow) NSWindow() lcl.NSWindow {
//	return lcl.PlatformWindow(m.Instance())
//}

func (m *TWindow) NSWindow() INSWindow {
	return m.nsWindow
}
func (m *TWindow) CreateParams(params *types.TCreateParams) {
}

// InternalBeforeFormCreate 在表单创建之前执行的内部初始化方法
// 该方法在 TWindow 实例化过程中被调用
func (m *TWindow) InternalBeforeFormCreate() {
	m.nsWindow = cocoa.AsNSWindow(unsafe.Pointer(lcl.PlatformWindow(m.Instance())))
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
			m.nsDelegate = cocoa.NewWindowDelegate(m.nsWindow)
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
		m.nsWindow.ExitMinimized()
		m.nsWindow.Maximize()
	case types.WsNormal:
		m.nsWindow.Restore()
	case types.WsMinimized:
		m.nsWindow.Minimized()
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
		m.nsWindow.EnterFullScreen()
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.SetWindowState(types.WsNormal)
			m.nsWindow.ExitFullScreen()
		})
	}
}
