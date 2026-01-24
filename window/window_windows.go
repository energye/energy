//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows

package window

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/pkgs/win32"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
)

type TWindow struct {
	TEnergyWindow
}

func (m *TWindow) borderFrameless() {
	hWnd := m.Handle()
	gwlStyle := win.GetWindowLong(hWnd, win.GWL_STYLE)
	win.SetWindowLong(hWnd, win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME))
	win.SetWindowPos(hWnd, 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
}

func (m *TWindow) _BeforeFormCreate() {
}

func (m *TWindow) _BeforeFormShow() {
}

// SetOptions 设置webview窗口的选项配置
func (m *TWindow) SetOptions() {
	m._HookWndProcMessage()
	if application.GApplication != nil {
		hWnd := m.Handle()
		options := application.GApplication.Options
		if options.WindowTransparent {
			if win32.Windows1122H2() {
				win32.EnableTranslucency(hWnd, int32(options.Windows.BackdropType))
			} else {
				win32.SetTranslucentBackground(hWnd)
			}
		}
		if options.Windows.WindowProtected {
			win32.SetWindowDisplayAffinity(hWnd, win.WDA_EXCLUDEFROMCAPTURE)
		}
		if options.BackgroundColor != nil {
			r, g, b := byte(options.BackgroundColor.R), byte(options.BackgroundColor.G), byte(options.BackgroundColor.B)
			color := colors.TColor(colors.RGB(r, g, b))
			m.SetColor(color)
			win32.SetBackgroundColor(hWnd, r, g, b)
		}
		switch options.Windows.Theme {
		case application.SystemDefault:
			win32.ChangeTheme(hWnd, win32.IsCurrentlyDarkMode())
		case application.Light:
			win32.ChangeTheme(hWnd, false)
		case application.Dark:
			win32.ChangeTheme(hWnd, true)
		}
		if !application.GApplication.Options.Frameless {
			if application.GApplication.Options.DisableResize {
				m.SetBorderStyleToFormBorderStyle(types.BsSingle)
				m.EnabledMaximize(false)
			}
			if application.GApplication.Options.DisableMinimize {
				m.EnabledMinimize(false)
			}
			if application.GApplication.Options.DisableMaximize {
				m.EnabledMaximize(false)
			}
			if application.GApplication.Options.DisableSystemMenu {
				m.EnabledSystemMenu(false)
			}
		}
		constr := m.Constraints()
		if application.GApplication.Options.MaxWidth > 0 || application.GApplication.Options.MaxHeight > 0 {
			constr.SetMaxWidth(application.GApplication.Options.MaxWidth)
			constr.SetMaxHeight(application.GApplication.Options.MaxHeight)
		}
		if application.GApplication.Options.MinWidth > 0 || application.GApplication.Options.MinHeight > 0 {
			constr.SetMinWidth(application.GApplication.Options.MinWidth)
			constr.SetMinHeight(application.GApplication.Options.MinHeight)
		}
		if options.Width <= 0 {
			options.Width = m.Width()
		}
		if options.Height <= 0 {
			options.Height = m.Height()
		}
		m.SetCaption(options.Caption)
		m.SetBounds(options.X, options.Y, options.Width, options.Height)
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
		hWnd := m.Handle()
		m.windowsState = types.WsFullScreen
		// save current window rect, use ExitFullScreen
		m.previousWindowPlacement = m.BoundsRect()
		monitorRect := m.Monitor().BoundsRect()
		if !application.GApplication.Options.Frameless {
			// save current window style, use ExitFullScreen
			m.oldWindowStyle = uintptr(win.GetWindowLongPtr(hWnd, win.GWL_STYLE))
			m.borderFrameless()
			m.SetWindowState(types.WsFullScreen)
		}
		win.SetWindowPos(hWnd, win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if !application.GApplication.Options.Frameless {
				win.SetWindowLong(m.Handle(), win.GWL_STYLE, m.oldWindowStyle)
			}
			m.windowsState = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBoundsRect(m.previousWindowPlacement)
		})
	}
}

func (m *TWindow) UpdateTheme() {
	if win.IsCurrentlyHighContrastMode() {
		return
	}
	if !win32.Windows101809() {
		return
	}
	options := application.GApplication.Options
	var isDark bool
	switch options.Windows.Theme {
	case application.SystemDefault:
		isDark = win32.IsCurrentlyDarkMode()
	case application.Dark:
		isDark = true
	case application.Light:
		isDark = false
	}
	hWnd := m.Handle()
	win32.ChangeTheme(hWnd, isDark)
	themeSetting := options.Windows.ThemeSetting
	if win32.Windows101809() && themeSetting != nil {
		if m.Active() {
			if isDark {
				win32.SetTitleBarColor(hWnd, themeSetting.DarkTitleBar)
				win32.SetTitleTextColor(hWnd, themeSetting.DarkTitleText)
				win32.SetBorderColor(hWnd, themeSetting.DarkBorder)
			} else {
				win32.SetTitleBarColor(hWnd, themeSetting.LightTitleBar)
				win32.SetTitleTextColor(hWnd, themeSetting.LightTitleText)
				win32.SetBorderColor(hWnd, themeSetting.LightBorder)
			}
		} else {
			if isDark {
				win32.SetTitleBarColor(hWnd, themeSetting.DarkTitleBarInactive)
				win32.SetTitleTextColor(hWnd, themeSetting.DarkTitleTextInactive)
				win32.SetBorderColor(hWnd, themeSetting.DarkBorderInactive)
			} else {
				win32.SetTitleBarColor(hWnd, themeSetting.LightTitleBarInactive)
				win32.SetTitleTextColor(hWnd, themeSetting.LightTitleTextInactive)
				win32.SetBorderColor(hWnd, themeSetting.LightBorderInactive)
			}
		}
	}
}
