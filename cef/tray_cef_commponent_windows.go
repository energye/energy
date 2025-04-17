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
// +build windows

// 基于 LCL 系统托盘 - windows 平台
// Html + CSS + JavaScript实现

package cef

import (
	. "github.com/cyber-xxm/energy/v2/common"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// CEFTray CEF基于LCL组件+html 托盘
type CEFTray struct {
	*lcl.TForm
	owner        lcl.IComponent
	trayIcon     *lcl.TTrayIcon
	chromium     IChromium
	windowParent ICEFWindowParent
	x, y, w, h   int32
	mouseUp      TMouseEvent
	isClosing    bool
	url          string
}

// newLCLTrayWindow 创建LCL系统托盘
func newLCLTrayWindow(owner lcl.IComponent, width, height int32, url string) *CEFTray {
	var trayForm *CEFTray
	lcl.Application.CreateForm(&trayForm)
	trayForm.trayIcon = lcl.NewTrayIcon(owner)
	trayForm.trayIcon.SetVisible(true)
	trayForm.owner = owner
	trayForm.x = -width
	trayForm.y = -height
	trayForm.w = width
	trayForm.h = height
	trayForm.url = url
	trayForm.onMouseEvent()
	trayForm.createTrayWindow()
	return trayForm
}

// OnFormCreate TForm创建
func (m *CEFTray) OnFormCreate(sender lcl.IObject) {
	m.SetShowInTaskBar(types.StNever)
}

// AsSysTray 尝试转换为 SysTray 组件托盘，如果创建的是其它类型托盘返回nil
func (m *CEFTray) AsSysTray() *SysTray {
	return nil
}

// AsViewsFrameTray 尝试转换为 views framework 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *CEFTray) AsViewsFrameTray() *ViewsFrameTray {
	return nil
}

// AsCEFTray 尝试转换为 LCL+CEF 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *CEFTray) AsCEFTray() *CEFTray {
	return m
}

// AsLCLTray 尝试转换为 LCL 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *CEFTray) AsLCLTray() *LCLTray {
	return nil
}

// Show 显示/启动 托盘
func (m *CEFTray) Show() {
	if BrowserWindow.MainWindow().Chromium() == nil || !BrowserWindow.MainWindow().Chromium().Initialized() {
		return
	}
	m.TForm.Show()
}

// Hide 隐藏 托盘
func (m *CEFTray) Hide() {
	m.TForm.Hide()
}

func (m *CEFTray) close() {
	if m.isClosing {
		return
	}
	m.Hide()
	m.trayIcon.SetVisible(false)
	m.TForm.Close()
}

// SetOnDblClick 设置双击事件
func (m *CEFTray) SetOnDblClick(fn TrayICONClick) {
	m.trayIcon.SetOnDblClick(func(sender lcl.IObject) {
		fn()
	})
}

// SetOnClick 设置单击事件
func (m *CEFTray) SetOnClick(fn TrayICONClick) {
	m.trayIcon.SetOnClick(func(sender lcl.IObject) {
		fn()
	})
}

// SetOnMouseUp 鼠标抬起事件
func (m *CEFTray) SetOnMouseUp(fn TMouseEvent) {
	m.mouseUp = fn
}

// SetOnMouseDown 鼠标按下事件
func (m *CEFTray) SetOnMouseDown(fn lcl.TMouseEvent) {
	m.trayIcon.SetOnMouseDown(fn)
}

// SetOnMouseMove 鼠标移动事件
func (m *CEFTray) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	m.trayIcon.SetOnMouseMove(fn)
}

// Visible 显示状态
func (m *CEFTray) Visible() bool {
	return m.TForm.Visible()
}

// SetVisible 设置显示状态
func (m *CEFTray) SetVisible(v bool) {
	m.trayIcon.SetVisible(v)
}

// SetHint 设置提示
func (m *CEFTray) SetHint(value string) {
	m.trayIcon.SetHint(value)
}

// SetTitle 设置标题
func (m *CEFTray) SetTitle(title string) {
	m.TForm.SetCaption(title)
}

func (m *CEFTray) onMouseEvent() {
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

// Notice
// 显示系统通知
//
// title 标题
//
// content 内容
//
// timeout 显示时间(毫秒)
func (m *CEFTray) Notice(title, content string, timeout int32) {
	notification(m.trayIcon, title, content, timeout)
}

func (m *CEFTray) createTrayWindow() {
	m.TForm.SetBorderStyle(types.BsNone)
	m.TForm.SetFormStyle(types.FsStayOnTop)
	m.TForm.SetBounds(-(m.w * 2), -(m.h * 2), m.w, m.h)
	m.TForm.SetOnActivate(func(sender lcl.IObject) {
		m.chromium.Initialized()
		m.chromium.CreateBrowser(m.windowParent, "", nil, nil)
	})
	m.TForm.SetOnWndProc(func(msg *types.TMessage) {
		m.TForm.InheritedWndProc(msg)
		if msg.Msg == 6 && msg.WParam == 0 && msg.LParam == 0 {
			QueueAsyncCall(func(id int) {
				if m.isClosing {
					return
				}
				m.TForm.Hide()
			})
		}
	})
	m.TForm.SetOnDeactivate(func(sender lcl.IObject) {
		if m.isClosing {
			return
		}
		m.TForm.Hide()
	})

	m.TForm.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) {
		*canClose = true
		logger.Debug("tray.window.onCloseQuery canClose:", *canClose)
		if m.isClosing {
			return
		}
		m.isClosing = true
		m.Hide()
		m.chromium.CloseBrowser(true)
		m.trayIcon.Free()
	})
	m.TForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		*action = types.CaFree
		logger.Debug("tray.window.onClose action:", *action)
	})
	m.TForm.SetOnShow(func(sender lcl.IObject) {
		if m.windowParent != nil {
			QueueAsyncCall(func(id int) {
				m.windowParent.UpdateSize()
			})
		}
	})
	m.windowParent = NewCEFWindowParent(m.TForm)
	m.windowParent.SetParent(m.TForm)
	m.windowParent.SetAlign(types.AlClient)
	m.windowParent.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	m.chromium = NewChromium(m.windowParent, nil)
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		model.Clear()
	})
	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {
		m.chromium.LoadUrl(m.url)
	})
	// 移除掉，没什么用了
	//m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool) bool {
	//	chromiumOnBeforeBrowser(nil, browser, frame, request) // default impl
	//	return false
	//})
	m.chromium.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *TCefReturnValue) {
		if assetserve.AssetsServerHeaderKeyValue != "" {
			request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
		}
	})
	m.chromium.SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *TCefCloseBrowserAction) {
		logger.Debug("tray.chromium.onClose")
		if IsDarwin() {
			m.windowParent.DestroyChildWindow()
		}
		*aAction = CbaClose
	})
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		logger.Debug("tray.chromium.onBeforeClose")
	})
	m.chromium.SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, message *ICefProcessMessage) bool {
		return false
	})
	m.windowParent.SetChromium(m.chromium, 0)
	//m.chromium.SetDefaultURL(m.url)
}

func (m *CEFTray) Chromium() IChromium {
	return m.chromium
}

// SetIconFS 设置托盘图标
func (m *CEFTray) SetIconFS(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFSFile(iconResourcePath)
}

// SetIcon 设置托盘图标
func (m *CEFTray) SetIcon(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFile(iconResourcePath)
}
