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

func (m *TWindow) CreateParams(params *types.TCreateParams) {
	if m.options != nil {
		//params.ExStyle = params.ExStyle | win.WS_EX_NOREDIRECTIONBITMAP
		if m.options.WindowTransparent {
			if win32.Windows8629200() {
				// > windows 8
				params.ExStyle = params.ExStyle | win.WS_EX_NOREDIRECTIONBITMAP
			} else {
				// windows 7
				params.ExStyle = params.ExStyle | win.WS_EX_LAYERED
			}
		}
	}
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

// UpdateWindowOption 设置窗口配置选项
func (m *TWindow) UpdateWindowOption() {
	m._HookWndProcMessage()
	if m.options != nil {
		hWnd := m.Handle()
		if m.options.WindowTransparent {
			if win32.Windows1122H2() {
				println(" > Windows11")
				win32.EnableTranslucency(hWnd, int32(m.options.Windows.BackdropType))
			} else if win32.Windows101809() {
				println(" > Windows10")
				win32.SetTranslucentBackground(hWnd)
			} else if !win32.Windows8629200() {
				println(" > Windows7")
				//var enable uint32 = 1
				//win.DwmSetWindowAttribute(hWnd, 0x00000002, unsafe.Pointer(&enable), unsafe.Sizeof(enable))
			}
		}
		if m.options.Windows.WindowProtected {
			win32.SetWindowDisplayAffinity(hWnd, win.WDA_EXCLUDEFROMCAPTURE)
		}
		if m.options.BackgroundColor != nil {
			r, g, b := byte(m.options.BackgroundColor.R), byte(m.options.BackgroundColor.G), byte(m.options.BackgroundColor.B)
			color := colors.TColor(colors.RGB(r, g, b))
			_ = color
			m.SetColor(color)
			win32.SetBackgroundColor(hWnd, r, g, b)
		}
		switch m.options.Windows.Theme {
		case application.SystemDefault:
			win32.ChangeTheme(hWnd, win32.IsCurrentlyDarkMode())
		case application.Light:
			win32.ChangeTheme(hWnd, false)
		case application.Dark:
			win32.ChangeTheme(hWnd, true)
		}
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
		hWnd := m.Handle()
		m.windowsState = types.WsFullScreen
		// save current window rect, use ExitFullScreen
		m.previousWindowPlacement = m.BoundsRect()
		monitorRect := m.Monitor().BoundsRect()
		if !m.options.Frameless {
			// save current window style, use ExitFullScreen
			gwlStyle := win.GetWindowLongPtr(hWnd, win.GWL_STYLE)
			m.oldWindowStyle = gwlStyle
			win.SetWindowLongPtr(hWnd, win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME))
			win.SetWindowPos(hWnd, 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))
			m.SetWindowState(types.WsFullScreen)
		}
		win.SetWindowPos(hWnd, win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			hWnd := m.Handle()
			if !m.options.Frameless {
				win.SetWindowLongPtr(hWnd, win.GWL_STYLE, m.oldWindowStyle)
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
	if m.options == nil {
		return
	}
	options := m.options
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
