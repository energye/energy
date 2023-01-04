//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl/types"
)

type TCEFFrame map[int64]*ICefFrame

// 窗口信息
type TCefWindowInfo struct {
	Window         *LCLBrowserWindow `json:"-"` //窗口Form
	Browser        *ICefBrowser      `json:"-"` //浏览器对象,加载完chromium之后创建
	WindowProperty *WindowProperty   `json:"-"` //窗口属性
	Frames         TCEFFrame         `json:"-"` //当前浏览器下的所有frame
	auxTools       *auxTools         //辅助工具
}

type auxTools struct {
	devToolsWindow   *LCLBrowserWindow //devTools
	devToolsX        int32             //上次改变的窗体位置，宽度
	devToolsY        int32             //
	devToolsWidth    int32             //
	devToolsHeight   int32             //
	viewSourceWindow *LCLBrowserWindow //viewSource
	viewSourceUrl    string            //
	viewSourceX      int32             //上次改变的窗体位置，宽度
	viewSourceY      int32             //
	viewSourceWidth  int32             //
	viewSourceHeight int32             //
}

type WindowProperty struct {
	IsShowModel    bool               //是否以模态窗口显示
	WindowState    types.TWindowState //窗口 状态
	Title          string             //窗口 标题
	Url            string             //默认打开URL
	Icon           string             //窗口图标 加载本地图标
	IconFS         string             //窗口图标 加载emfs内置图标
	CanMinimize    bool               //窗口 是否启用最小化功能
	CanMaximize    bool               //窗口 是否启用最大化功能
	CanResize      bool               //窗口 是否允许调整窗口大小
	CanClose       bool               //窗口 关闭时是否关闭窗口
	IsCenterWindow bool               //窗口 是否居中显示
	AlwaysOnTop    bool               //窗口 置顶
	X              int32              //窗口 IsCenterWindow=false X坐标
	Y              int32              //窗口 IsCenterWindow=false Y坐标
	Width          int32              //窗口 宽
	Height         int32              //窗口 高
}

func NewWindowProperty() *WindowProperty {
	return &WindowProperty{
		Title:          "Energy",
		Url:            "about:blank",
		CanMinimize:    true,
		CanMaximize:    true,
		CanResize:      true,
		CanClose:       true,
		IsCenterWindow: true,
		Width:          1024,
		Height:         768,
	}
}

func (m *TCefWindowInfo) Chromium() IChromiumProc {
	return m.Window.chromium
}

func (m *TCefWindowInfo) chromiumEvent() IChromiumEvent {
	return m.Window.chromium
}

func (m *TCefWindowInfo) Minimize() {
	QueueAsyncCall(func(id int) {
		m.Window.SetWindowState(types.WsMinimized)
	})
}

func (m *TCefWindowInfo) Maximize() {
	BrowserWindow.uiLock.Lock()
	defer BrowserWindow.uiLock.Unlock()
	QueueAsyncCall(func(id int) {
		var bs = m.Window.BorderStyle()
		var monitor = m.Window.Monitor().WorkareaRect()
		if bs == types.BsNone {
			var ws = m.Window.WindowState()
			var redWindowState types.TWindowState
			//默认状态0
			if m.WindowProperty.WindowState == types.WsNormal && m.WindowProperty.WindowState == ws {
				redWindowState = types.WsMaximized
			} else {
				if m.WindowProperty.WindowState == types.WsNormal {
					redWindowState = types.WsMaximized
				} else if m.WindowProperty.WindowState == types.WsMaximized {
					redWindowState = types.WsNormal
				}
			}
			m.WindowProperty.WindowState = redWindowState
			if redWindowState == types.WsMaximized {
				m.WindowProperty.X = m.Window.Left()
				m.WindowProperty.Y = m.Window.Top()
				m.WindowProperty.Width = m.Window.Width()
				m.WindowProperty.Height = m.Window.Height()
				m.Window.SetLeft(monitor.Left)
				m.Window.SetTop(monitor.Top)
				m.Window.SetWidth(monitor.Right - monitor.Left - 1)
				m.Window.SetHeight(monitor.Bottom - monitor.Top - 1)
			} else if redWindowState == types.WsNormal {
				m.Window.SetLeft(m.WindowProperty.X)
				m.Window.SetTop(m.WindowProperty.Y)
				m.Window.SetWidth(m.WindowProperty.Width)
				m.Window.SetHeight(m.WindowProperty.Height)
			}
		} else {
			if m.Window.WindowState() == types.WsMaximized {
				m.Window.SetWindowState(types.WsNormal)
				if common.IsDarwin() {
					m.Window.SetWindowState(types.WsMaximized)
					m.Window.SetWindowState(types.WsNormal)
				}
			} else if m.Window.WindowState() == types.WsNormal {
				m.Window.SetWindowState(types.WsMaximized)
			}
			m.WindowProperty.WindowState = m.Window.WindowState()
		}
	})
}

// 关闭窗口-在ui线程中执行
func (m *TCefWindowInfo) Close() {
	BrowserWindow.uiLock.Lock()
	defer BrowserWindow.uiLock.Unlock()
	QueueAsyncCall(func(id int) {
		if m == nil {
			logger.Error("关闭浏览器 WindowInfo 为空")
			return
		}
		if m.Window == nil {
			logger.Error("关闭浏览器 Window 为空")
			return
		}
		if common.IsDarwin() {
			//main window close
			if m.Window.WindowType() == consts.WT_MAIN_BROWSER {
				m.Window.Close()
			} else {
				//sub window close
				m.Window.isClosing = true
				m.Window.Hide()
				m.Window.chromium.CloseBrowser(true)
			}
		} else {
			m.Window.isClosing = true
			m.Window.Hide()
			m.Window.chromium.CloseBrowser(true)
		}
	})
}

//禁用口透明
func (m *TCefWindowInfo) DisableTransparent() {
	m.Window.SetAllowDropFiles(false)
	m.Window.SetAlphaBlend(false)
	m.Window.SetAlphaBlendValue(255)
}

//使窗口透明 value 0 ~ 255
func (m *TCefWindowInfo) EnableTransparent(value uint8) {
	m.Window.SetAllowDropFiles(true)
	m.Window.SetAlphaBlend(true)
	m.Window.SetAlphaBlendValue(value)
}

//禁用最小化按钮
func (m *TCefWindowInfo) DisableMinimize() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Exclude(types.BiMinimize))
}

//禁用最大化按钮
func (m *TCefWindowInfo) DisableMaximize() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Exclude(types.BiMaximize))
}

//禁用系统菜单-同时禁用最小化，最大化，关闭按钮
func (m *TCefWindowInfo) DisableSystemMenu() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Exclude(types.BiSystemMenu))
}

//禁用帮助菜单
func (m *TCefWindowInfo) DisableHelp() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Exclude(types.BiHelp))
}

//启用最小化按钮
func (m *TCefWindowInfo) EnableMinimize() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Include(types.BiMinimize))
}

//启用最大化按钮
func (m *TCefWindowInfo) EnableMaximize() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Include(types.BiMaximize))
}

//启用系统菜单-同时禁用最小化，最大化，关闭按钮
func (m *TCefWindowInfo) EnableSystemMenu() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Include(types.BiSystemMenu))
}

//启用帮助菜单
func (m *TCefWindowInfo) EnableHelp() {
	m.Window.SetBorderIcons(m.Window.BorderIcons().Include(types.BiHelp))
}

func (m TCEFFrame) GetByFrameId(frameId int64) *ICefFrame {
	if m != nil {
		if frame, ok := m[frameId]; ok {
			return frame
		}
	}
	return nil
}
