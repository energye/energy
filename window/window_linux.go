//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

package window

import (
	"github.com/energye/energy/v3/pkgs/gtk3"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"unsafe"
)

type ILinuxWindow interface {
	IWindow
	GTKWindow() *gtk3.Window
	GTKWindowLayout() *gtk3.Layout
	GTKMenuBar() *gtk3.MenuBar
	GTKScrolledWindow() *gtk3.ScrolledWindow
}

type TWindow struct {
	TEnergyWindow
	gtkWindow         *gtk3.Window
	gtkWindowLayout   *gtk3.Layout
	gtkMenuBar        *gtk3.MenuBar
	gtkScrolledWindow *gtk3.ScrolledWindow
}

func (m *TWindow) CreateParams(params *types.TCreateParams) {

}

func (m *TWindow) GTKWindow() *gtk3.Window {
	return m.gtkWindow
}

func (m *TWindow) GTKWindowLayout() *gtk3.Layout {
	return m.gtkWindowLayout
}

func (m *TWindow) GTKMenuBar() *gtk3.MenuBar {
	return m.gtkMenuBar
}

func (m *TWindow) GTKScrolledWindow() *gtk3.ScrolledWindow {
	return m.gtkScrolledWindow
}

func (m *TWindow) getGtkWidget() {
	var iterate func(list *gtk3.List)
	iterate = func(list *gtk3.List) {
		if list == nil {
			return
		}
		for i := uint(0); i < list.Length(); i++ {
			data := list.NthDataRaw(i)
			container := gtk3.ToContainer(data)
			widgetName := container.TypeFromInstance().Name()
			if widgetName == "GtkLayout" { // window > level 1
				m.gtkWindowLayout = gtk3.ToLayout(data)
			} else if widgetName == "GtkMenuBar" {
				m.gtkMenuBar = gtk3.ToMenuBar(data) // window > level 2
			} else if widgetName == "GtkScrolledWindow" {
				m.gtkScrolledWindow = gtk3.ToScrolledWindow(data) // window > level 2
			}
			iterate(container.GetChildren())
		}
	}
	iterate(m.gtkWindow.GetChildren())
	if m.gtkWindowLayout == nil {
		println("WARNING: GtkWindow does not have a Layout")
	}
}

func (m *TWindow) _BeforeFormCreate() {
	gtkHandle := lcl.PlatformHandle(m.Handle())
	m.gtkWindow = gtk3.ToGtkWindow(unsafe.Pointer(gtkHandle.Gtk3Window()))
	m.getGtkWidget()
	if m.options != nil {
		if m.options.WindowTransparent {
			screen := m.gtkWindow.GetScreen()
			visual, err := screen.GetRGBAVisual()
			if err == nil && visual != nil && screen.IsComposited() {
				m.gtkWindow.SetVisual(visual)
				m.gtkWindow.SetAppPaintable(true)
			}
		}
	}
	//m.gtkWindow.SetOnConfigure(func(sender *gtk3.Widget, event *gtk3.EventConfigure) bool {
	//	for _, fn := range m.onWindowResizeList {
	//		fn(nil)
	//	}
	//	return false
	//})
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
		m.gtkWindow.SetDecorated(!m.options.Frameless)
		if !m.options.Frameless {
			if m.options.DisableResize {
				m.SetBorderStyleToFormBorderStyle(types.BsSingle)
				m.EnabledMaximize(false)
			}
			if m.options.DisableMinimize {
				m.EnabledMinimize(false)
			}
			if m.options.DisableMaximize {
				m.EnabledMaximize(false)
			}
			if m.options.DisableSystemMenu {
				m.EnabledSystemMenu(false)
			}
		}
		constr := m.Constraints()
		if m.options.MaxWidth > 0 || m.options.MaxHeight > 0 {
			constr.SetMaxWidth(m.options.MaxWidth)
			constr.SetMaxHeight(m.options.MaxHeight)
		}
		if m.options.MinWidth > 0 || m.options.MinHeight > 0 {
			constr.SetMinWidth(m.options.MinWidth)
			constr.SetMinHeight(m.options.MinHeight)
		}
		if m.options.Width <= 0 {
			m.options.Width = m.Width()
		}
		if m.options.Height <= 0 {
			m.options.Height = m.Height()
		}
		m.SetCaption(m.options.Caption)
		m.SetBounds(m.options.X, m.options.Y, m.options.Width, m.options.Height)
	}
}

func (m *TWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.IsMinimize() || m.IsMaximize() {
			m.Restore()
		}
		m.windowsState = types.WsFullScreen
		// save current window rect, use ExitFullScreen
		m.previousWindowPlacement = m.BoundsRect()
		m.SetWindowState(types.WsFullScreen)
		m.gtkWindow.Fullscreen()
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.windowsState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBoundsRect(m.previousWindowPlacement)
			m.gtkWindow.Unfullscreen()
		})
	}
}
