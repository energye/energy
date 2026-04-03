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
	GTKWindowMenuBar() *gtk3.MenuBar
	GTKWindowScrolledWindow() *gtk3.ScrolledWindow
}

type TWindow struct {
	TEnergyWindow
	gtkWindow               *gtk3.Window
	gtkWindowBox            *gtk3.Box
	gtkWindowLayout         *gtk3.Layout
	gtkWindowMenuBar        *gtk3.MenuBar
	gtkWindowScrolledWindow *gtk3.ScrolledWindow
	gtkCssProvider          *gtk3.CssProvider
}

func (m *TWindow) CreateParams(params *types.TCreateParams) {

}

func (m *TWindow) GTKWindow() *gtk3.Window {
	return m.gtkWindow
}

func (m *TWindow) GTKWindowBox() *gtk3.Box {
	return m.gtkWindowBox
}

func (m *TWindow) GTKWindowLayout() *gtk3.Layout {
	return m.gtkWindowLayout
}

func (m *TWindow) GTKWindowMenuBar() *gtk3.MenuBar {
	return m.gtkWindowMenuBar
}

func (m *TWindow) GTKWindowScrolledWindow() *gtk3.ScrolledWindow {
	return m.gtkWindowScrolledWindow
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
			if widgetName == "GtkBox" { // window > level 1
				m.gtkWindowBox = gtk3.ToBox(data)
			} else if widgetName == "GtkMenuBar" { // window > level 2
				m.gtkWindowMenuBar = gtk3.ToMenuBar(data)
			} else if widgetName == "GtkScrolledWindow" || widgetName == "LCLGtkScrolledWindow" { // window > level 2
				m.gtkWindowScrolledWindow = gtk3.ToScrolledWindow(data)
			} else if widgetName == "GtkLayout" { // window > level 3
				m.gtkWindowLayout = gtk3.ToLayout(data)
			}
			iterate(container.GetChildren())
		}
	}
	iterate(m.gtkWindow.GetChildren())
	if m.gtkWindowBox == nil {
		println("WARNING: GtkWindow does not have a Box")
	}
	if m.gtkWindowScrolledWindow == nil {
		println("WARNING: GtkWindow does not have a ScrolledWindow")
	}
	if m.gtkWindowLayout == nil {
		println("WARNING: GtkWindow does not have a Layout")
	}

	//options := m.options
	//if options.WebviewTransparent {
	//	m.gtkWindowBox.GetStyleContext().AddClass("webview-box")
	//	m.gtkWindowScrolledWindow.GetStyleContext().AddClass("webview-box")
	//	m.gtkWindowLayout.GetStyleContext().AddClass("webview-box")
	//
	//	r, g, b, a := options.BackgroundColor.R, options.BackgroundColor.G, options.BackgroundColor.B, options.BackgroundColor.A
	//	webviewCss := fmt.Sprintf(".webview-box {background-color: rgba(%d, %d, %d, %1.1f);}", r, g, b, float64(a)/255.0)
	//	if m.gtkCssProvider == nil {
	//		m.gtkCssProvider = gtk3.NewCssProvider()
	//		m.gtkWindowBox.GetStyleContext().AddProvider(m.gtkCssProvider, gtk3.STYLE_PROVIDER_PRIORITY_USER)
	//		m.gtkWindowScrolledWindow.GetStyleContext().AddProvider(m.gtkCssProvider, gtk3.STYLE_PROVIDER_PRIORITY_USER)
	//		m.gtkWindowLayout.GetStyleContext().AddProvider(m.gtkCssProvider, gtk3.STYLE_PROVIDER_PRIORITY_USER)
	//		m.gtkCssProvider.Unref()
	//	}
	//	var err error
	//	err = m.gtkCssProvider.LoadFromData(webviewCss)
	//	if err != nil {
	//		//println("CssProvider.LoadFromData:", err.Error())
	//	}
	//}
}

// InternalBeforeFormCreate 在表单创建之前执行的内部初始化方法
// 该方法在 TWindow 实例化过程中被调用
func (m *TWindow) InternalBeforeFormCreate() {
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
