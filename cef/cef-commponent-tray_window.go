//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//Cef托盘
type tCefTrayForm struct {
	*lcl.TForm
	owner        lcl.IWinControl
	trayIcon     *lcl.TTrayIcon
	chromium     *TCEFChromium
	windowParent ITCefWindow
	x, y, w, h   int32
	mouseUp      TMouseEvent
	canClose     bool
	isClosing    bool
	url          string
}

func newCefTray(owner lcl.IWinControl, width, height int32, url string) *tCefTrayForm {
	var trayForm *tCefTrayForm
	lcl.Application.CreateForm(&trayForm)
	trayForm.trayIcon = lcl.NewTrayIcon(owner)
	trayForm.trayIcon.SetVisible(true)
	trayForm.owner = owner
	trayForm.x = -width
	trayForm.y = -height
	trayForm.w = width
	trayForm.h = height
	trayForm.url = url
	trayForm.onmMouse()
	trayForm.createCefTrayWindow()
	return trayForm
}

func (m *tCefTrayForm) OnFormCreate(sender lcl.IObject) {
	m.SetShowInTaskBar(types.StNever)
}

func (m *tCefTrayForm) Tray() *Tray {
	return nil
}

func (m *tCefTrayForm) Show() {
	if BrowserWindow.browserWindow.chromium == nil || !BrowserWindow.browserWindow.chromium.Initialized() {
		return
	}
	m.TForm.Show()
}

func (m *tCefTrayForm) Hide() {
	m.TForm.Hide()
}

func (m *tCefTrayForm) close() {
	m.Hide()
	m.TForm.Close()
}

func (m *tCefTrayForm) SetOnDblClick(fn lcl.TNotifyEvent) {
	m.trayIcon.SetOnDblClick(fn)
}

func (m *tCefTrayForm) SetOnClick(fn lcl.TNotifyEvent) {
	m.trayIcon.SetOnClick(fn)
}

func (m *tCefTrayForm) SetOnMouseUp(fn TMouseEvent) {
	m.mouseUp = fn
}
func (m *tCefTrayForm) SetOnMouseDown(fn lcl.TMouseEvent) {
	m.trayIcon.SetOnMouseDown(fn)
}
func (m *tCefTrayForm) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	m.trayIcon.SetOnMouseMove(fn)
}

func (m *tCefTrayForm) Visible() bool {
	return m.TForm.Visible()
}

func (m *tCefTrayForm) SetVisible(v bool) {
	m.trayIcon.SetVisible(v)
}

func (m *tCefTrayForm) SetHint(value string) {
	m.trayIcon.SetHint(value)
}

func (m *tCefTrayForm) SetTitle(title string) {
	m.TForm.SetCaption(title)
}

func (m *tCefTrayForm) onmMouse() {
	QueueAsyncCall(func(id int) {
		m.trayIcon.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
			var monitor = m.TForm.Monitor()
			var monitorWidth = monitor.Width()
			width, height := m.TForm.Width(), m.TForm.Height()
			var mx = x + width
			var my = y + height
			if mx < monitorWidth {
				mx = x
			} else {
				mx = x - width
			}
			if my > m.h {
				my = y
			}
			if my > height {
				my = y - height
			}
			m.TForm.SetBounds(mx, my, width, height)
			var ret bool
			if m.mouseUp != nil {
				ret = m.mouseUp(sender, button, shift, x, y)
			}
			if !ret {
				if button == types.MbRight {
					m.Show()
				}
			}
		})
	})
}

//设置托盘气泡
//title 气泡标题
//content 气泡内容
//timeout 显示时间(毫秒)
func (m *tCefTrayForm) SetBalloon(title, content string, timeout int32) ITray {
	m.trayIcon.SetBalloonTitle(title)
	m.trayIcon.SetBalloonHint(content)
	m.trayIcon.SetBalloonTimeout(timeout)
	return m
}

//显示托盘气泡
func (m *tCefTrayForm) ShowBalloon() {
	m.trayIcon.ShowBalloonHint()
}

func (m *tCefTrayForm) createCefTrayWindow() {
	m.TForm.SetBorderStyle(types.BsNone)
	m.TForm.SetFormStyle(types.FsStayOnTop)
	m.TForm.SetBounds(-(m.w * 2), -(m.h * 2), m.w, m.h)
	m.TForm.SetOnActivate(func(sender lcl.IObject) {
		m.chromium.Initialized()
		m.chromium.CreateBrowser(m.windowParent)
	})
	m.TForm.SetOnWndProc(func(msg *types.TMessage) {
		m.TForm.InheritedWndProc(msg)
		if msg.Msg == 6 && msg.WParam == 0 && msg.LParam == 0 {
			QueueAsyncCall(func(id int) {
				m.TForm.Hide()
			})
		}
	})
	m.TForm.SetOnDeactivate(func(sender lcl.IObject) {
		m.TForm.Hide()
	})

	m.TForm.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) {
		*canClose = m.canClose
		if m.isClosing {
			return
		}
		m.isClosing = true
		m.Hide()
		m.chromium.CloseBrowser(true)
	})
	m.TForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		*action = types.CaFree
	})
	m.TForm.SetOnShow(func(sender lcl.IObject) {
		if m.windowParent != nil {
			QueueAsyncCall(func(id int) {
				m.windowParent.UpdateSize()
			})
		}
	})
	m.windowParent = NewCEFWindow(m.TForm)
	m.windowParent.SetParent(m.TForm)
	m.windowParent.SetAlign(types.AlClient)
	m.windowParent.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	m.chromium = NewChromium(m.windowParent, nil)
	//打开独立出事件
	m.chromium.EnableIndependentEvent()
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		model.Clear()
	})
	m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool {
		BrowserWindow.setOrIncNextWindowNum(browser.Identifier() + 1)
		return false
	})
	m.chromium.SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *TCefCloseBrowsesAction) {
		if IsDarwin() {
			*aAction = CbaDelay
		} else {
			*aAction = CbaClose
		}
		if IsDarwin() {
			m.windowParent.DestroyChildWindow()
		} else {
			QueueAsyncCall(func(id int) { //主进程执行
				m.windowParent.Free()
			})
		}
	})
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		m.canClose = true
	})
	//关闭独立出事件
	m.chromium.DisableIndependentEvent()
	m.chromium.SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, message *ICefProcessMessage) bool {
		return false
	})
	m.windowParent.SetChromium(m.chromium, 0)
	m.chromium.SetDefaultURL(m.url)
}

//设置托盘图标
func (m *tCefTrayForm) SetIcon(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFSFile(iconResourcePath)
}
