//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/commons"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type TCEFFrame map[int64]*ICefFrame

// 窗口信息
type TCefWindowInfo struct {
	Window         *BaseWindow     `json:"-"` //窗口Form
	Browser        *ICefBrowser    `json:"-"` //浏览器对象,加载完chromium之后创建
	WindowProperty *WindowProperty `json:"-"` //窗口属性
	Frames         TCEFFrame       `json:"-"` //当前浏览器下的所有frame
	auxTools       *auxTools       //辅助工具
}

type auxTools struct {
	devToolsWindow   *BaseWindow //开发者工具窗口
	devToolsX        int32       //上次改变的窗体位置，宽度
	devToolsY        int32       //
	devToolsWidth    int32       //
	devToolsHeight   int32       //
	viewSourceWindow *BaseWindow //viewSource
	viewSourceUrl    string      //
	viewSourceX      int32       //上次改变的窗体位置，宽度
	viewSourceY      int32       //
	viewSourceWidth  int32       //
	viewSourceHeight int32       //
}

type WindowProperty struct {
	IsShowModel    bool //是否以模态窗口显示
	windowState    types.TWindowState
	currentWindowX int32
	currentWindowY int32
	currentWindowW int32
	currentWindowH int32
}

func NewWindowProperty() *WindowProperty {
	return &WindowProperty{}
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
			if m.WindowProperty.windowState == types.WsNormal && m.WindowProperty.windowState == ws {
				redWindowState = types.WsMaximized
			} else {
				if m.WindowProperty.windowState == types.WsNormal {
					redWindowState = types.WsMaximized
				} else if m.WindowProperty.windowState == types.WsMaximized {
					redWindowState = types.WsNormal
				}
			}
			m.WindowProperty.windowState = redWindowState
			if redWindowState == types.WsMaximized {
				m.WindowProperty.currentWindowX = m.Window.Left()
				m.WindowProperty.currentWindowY = m.Window.Top()
				m.WindowProperty.currentWindowW = m.Window.Width()
				m.WindowProperty.currentWindowH = m.Window.Height()
				m.Window.SetLeft(monitor.Left)
				m.Window.SetTop(monitor.Top)
				m.Window.SetWidth(monitor.Right - monitor.Left - 1)
				m.Window.SetHeight(monitor.Bottom - monitor.Top - 1)
			} else if redWindowState == types.WsNormal {
				m.Window.SetLeft(m.WindowProperty.currentWindowX)
				m.Window.SetTop(m.WindowProperty.currentWindowY)
				m.Window.SetWidth(m.WindowProperty.currentWindowW)
				m.Window.SetHeight(m.WindowProperty.currentWindowH)
			}
		} else {
			if m.Window.WindowState() == types.WsMaximized {
				m.Window.SetWindowState(types.WsNormal)
				if commons.IsDarwin() {
					m.Window.SetWindowState(types.WsMaximized)
					m.Window.SetWindowState(types.WsNormal)
				}
			} else if m.Window.WindowState() == types.WsNormal {
				m.Window.SetWindowState(types.WsMaximized)
			}
			m.WindowProperty.windowState = m.Window.WindowState()
		}
	})
}

func (m *TCefWindowInfo) WindowId() int32 {
	return m.Window.windowId
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
			logger.Error("关闭浏览器 Form 为空 WindowId:", m.WindowId())
			return
		}
		m.Window.isClosing = true
		m.Window.Hide()
		m.Window.chromium.CloseBrowser(true)
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

// 设置或增加一个窗口序号
func (m *browser) setOrIncNextWindowNum(browserId ...int32) int32 {
	if len(browserId) > 0 {
		m.windowSerial = browserId[0]
	} else {
		m.windowSerial++
	}
	logger.Debug("next window serial:", m.windowSerial)
	return m.windowSerial
}

// 设置或减少一个窗口序号
func (m *browser) setOrDecNextWindowNum(browserId ...int32) int32 {
	if len(browserId) > 0 {
		m.windowSerial = browserId[0]
	} else {
		m.windowSerial--
	}
	return m.windowSerial
}

// 获得窗口序号
func (m *browser) GetNextWindowNum() int32 {
	return m.windowSerial
}

func (m *browser) createNextPopupWindow() {
	m.popupWindow = &BaseWindow{}
	m.popupWindow.TForm = lcl.NewForm(m.MainWindowForm())
	m.popupWindow.FormCreate()
	m.popupWindow.defaultWindowEvent()
	m.popupWindow.defaultWindowCloseEvent()
}

// 拿到窗口信息
func (m *browser) GetWindowInfo(browserId int32) *TCefWindowInfo {
	if winInfo, ok := m.windowInfo[browserId]; ok {
		return winInfo
	}
	return nil
}

func (m *browser) GetWindowsInfo() map[int32]*TCefWindowInfo {
	return m.windowInfo
}

func (m *browser) putWindowInfo(browserId int32, windowInfo *TCefWindowInfo) {
	if winInfo, ok := m.windowInfo[browserId]; ok {
		winInfo.Window = windowInfo.Window
		//winInfo.chromiumEvent = windowInfo.chromiumEvent
	} else {
		m.windowInfo[browserId] = windowInfo
	}
}

func (m *browser) removeWindowInfo(browseId int32) {
	delete(m.windowInfo, browseId)
	commons.Proc("CEF_RemoveGoForm").Call(uintptr(browseId))
}

func (m *browser) GetBrowser(browseId int32) *ICefBrowser {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Browser
	}
	return nil
}

func (m *browser) putBrowserFrame(browser *ICefBrowser, frame *ICefFrame) {
	if winInfo, ok := m.windowInfo[browser.Identifier()]; !ok {
		winInfo = &TCefWindowInfo{
			Browser: browser,
			Frames:  make(map[int64]*ICefFrame),
		}
		if frame != nil {
			winInfo.Frames[frame.Id] = frame
		}
		m.windowInfo[browser.Identifier()] = winInfo
	} else {
		winInfo.Browser = browser
		if frame != nil {
			winInfo.Frames[frame.Id] = frame
		}
	}
}

func (m *browser) GetFrames(browseId int32) map[int64]*ICefFrame {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Frames
	}
	return nil
}

func (m *browser) GetFrame(browseId int32, frameId int64) *ICefFrame {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Frames[frameId]
	}
	return nil
}

func (m *browser) RemoveFrame(browseId int32, frameId int64) {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		delete(winInfo.Frames, frameId)
	}
}

func (m *browser) IsSameFrame(browseId int32, frameId int64) bool {
	if frame := m.GetFrame(browseId, frameId); frame != nil {
		return true
	}
	return false
}

func (m *browser) removeNoValidFrames() {
	for _, winInfo := range m.windowInfo {
		for _, frm := range winInfo.Frames {
			if !frm.IsValid() {
				delete(winInfo.Frames, frm.Id)
			}
		}
	}
}

func (m TCEFFrame) GetByFrameId(frameId int64) *ICefFrame {
	if m != nil {
		if frame, ok := m[frameId]; ok {
			return frame
		}
	}
	return nil
}
