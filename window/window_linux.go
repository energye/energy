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
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/platform/linux/gtk3"
	gtk3Types "github.com/energye/energy/v3/platform/linux/types"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/godbus/dbus/v5"
	"unsafe"
)

type ILinuxWindow interface {
	IWindow
	GTKWindow() gtk3Types.IWindow
	GTKWindowLayout() gtk3Types.ILayout
	GTKWindowMenuBar() gtk3Types.IMenuBar
	GTKWindowScrolledWindow() gtk3Types.IScrolledWindow
}

type TWindow struct {
	TEnergyWindow
	gtkWindow               gtk3Types.IWindow
	gtkWindowBox            gtk3Types.IBox
	gtkWindowLayout         gtk3Types.ILayout
	gtkWindowMenuBar        gtk3Types.IMenuBar
	gtkWindowScrolledWindow gtk3Types.IScrolledWindow
	gtkCssProvider          gtk3Types.ICssProvider
}

func (m *TWindow) CreateParams(params *types.TCreateParams) {

}

func (m *TWindow) GTKWindow() gtk3Types.IWindow {
	return m.gtkWindow
}

func (m *TWindow) GTKWindowBox() gtk3Types.IBox {
	return m.gtkWindowBox
}

func (m *TWindow) GTKWindowLayout() gtk3Types.ILayout {
	return m.gtkWindowLayout
}

func (m *TWindow) GTKWindowMenuBar() gtk3Types.IMenuBar {
	return m.gtkWindowMenuBar
}

func (m *TWindow) GTKWindowScrolledWindow() gtk3Types.IScrolledWindow {
	return m.gtkWindowScrolledWindow
}

func (m *TWindow) getGtkWidget() {
	var iterate func(list gtk3Types.IList)
	iterate = func(list gtk3Types.IList) {
		if list == nil {
			return
		}
		for i := uint(0); i < list.Length(); i++ {
			data := list.NthDataRaw(i)
			container := gtk3.AsContainer(data)
			widgetName := container.GetName()
			if widgetName == "GtkBox" { // window > level 1
				m.gtkWindowBox = gtk3.AsBox(data)
			} else if widgetName == "GtkMenuBar" { // window > level 2
				m.gtkWindowMenuBar = gtk3.AsMenuBar(data)
			} else if widgetName == "GtkScrolledWindow" || widgetName == "LCLGtkScrolledWindow" { // window > level 2
				m.gtkWindowScrolledWindow = gtk3.AsScrolledWindow(data)
			} else if widgetName == "GtkLayout" { // window > level 3
				m.gtkWindowLayout = gtk3.AsLayout(data)
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
	m.gtkWindow = gtk3.AsWindow(unsafe.Pointer(gtkHandle.Gtk3Window()))
	m.getGtkWidget()
	if m.options != nil {
		if m.options.WindowTransparent {
			screen := m.gtkWindow.GetScreen()
			visual := screen.GetRGBAVisual()
			if visual != nil && screen.IsComposited() {
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
		m.UpdateConfigProperty()
	}
	// 启动系统主题变更监听
	m.startThemeObserver()
}

func (m *TWindow) IsCurrentlyDarkMode() bool {
	conn, err := dbus.SessionBus()
	if err != nil {
		return false
	}
	//defer conn.Close()
	obj := conn.Object("org.freedesktop.portal.Desktop", "/org/freedesktop/portal/desktop")
	call := obj.Call("org.freedesktop.portal.Settings.Read", 0, "org.freedesktop.appearance", "color-scheme")
	if call.Err != nil {
		return false
	}
	var result dbus.Variant
	if err := call.Store(&result); err != nil {
		return false
	}
	innerVariant, ok := result.Value().(dbus.Variant)
	if !ok {
		return false
	}
	colorScheme, ok := innerVariant.Value().(uint32)
	if !ok {
		return false
	}
	return colorScheme == gtk3Types.ColorSchemePreferDark
}

// UpdateTheme 更新主题
func (m *TWindow) UpdateTheme() {
	if m.options == nil {
		return
	}
	isDark := false
	switch m.options.Linux.Theme {
	case application.SystemDefault:
		isDark = m.IsCurrentlyDarkMode()
	case application.Dark:
		isDark = true
	case application.Light:
		isDark = false
	}
	m.doOnThemeChange(isDark)
}

func (m *TWindow) startThemeObserver() {
	settings := gtk3.SettingsGetDefault()
	if settings == nil {
		return
	}
	settings.SetOnThemeChanged(func(sender gtk3Types.PGtkWidget, pspec uintptr, userData gtk3Types.GPointer) {
		m.UpdateTheme()
	})
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
