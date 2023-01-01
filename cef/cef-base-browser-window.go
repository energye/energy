//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	. "github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/messages"
	"time"
)

type IBaseWindow interface {
	lcl.IWinControl
	FormCreate()
	WindowParent() ITCefWindow
	Chromium() IChromium
	ChromiumCreate(config *tCefChromiumConfig, defaultUrl string)
	registerEvent()
	registerDefaultEvent()
}

//BaseWindow 是一个基于chromium 和 lcl 的窗口组件
type BaseWindow struct {
	*lcl.TForm                                       //
	chromium            IChromium                    //
	windowParent        ITCefWindow                  //
	vFrameBrowserWindow *ViewsFrameworkBrowserWindow //基于CEF views framework窗口
	windowInfo          *TCefWindowInfo              //基于LCL窗口信息
	windowId            int32                        //
	windowType          WINDOW_TYPE                  //0:browser 1:devTools 2:viewSource 默认:0
	isClosing           bool                         //
	canClose            bool                         //
	onResize            []TNotifyEvent               //
	onActivate          []TNotifyEvent               //
	onShow              []TNotifyEvent               //
	onClose             []TCloseEvent                //
	onCloseQuery        []TCloseQueryEvent           //
	onActivateAfter     lcl.TNotifyEvent             //
	isFormCreate        bool                         //是否创建完成 WindowForm
	isChromiumCreate    bool                         //是否创建完成 Chromium
}

//创建一个带有 chromium 窗口
//
//该窗口默认不具备默认事件处理能力, 通过 EnableDefaultEvent 函数注册事件处理
func NewBrowserWindow(config *tCefChromiumConfig, defaultUrl string) *Window {
	var window = NewWindow()
	window.ChromiumCreate(config, defaultUrl)
	window.putChromiumWindowInfo()
	//BeforeBrowser是一个必须的默认事件，在浏览器创建时窗口序号会根据browserId生成
	window.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool { return false })
	return window
}

func (m *BaseWindow) Id() int32 {
	return m.windowId
}

func (m *BaseWindow) Show() {
	m.TForm.Show()
}

func (m *BaseWindow) Hide() {
	m.TForm.Hide()
}

func (m *BaseWindow) Visible() bool {
	return m.TForm.Visible()
}

func (m *BaseWindow) SetVisible(value bool) {
	m.TForm.SetVisible(value)
}

//返回窗口信息
func (m *BaseWindow) WindowInfo() *TCefWindowInfo {
	return m.windowInfo
}

//以默认的方式展示在任务栏上
func (m *BaseWindow) SetDefaultInTaskBar() {
	m.TForm.SetShowInTaskBar(types.StDefault)
}

//展示在任务栏上
func (m *BaseWindow) SetShowInTaskBar() {
	m.TForm.SetShowInTaskBar(types.StAlways)
}

//不会展示在任务栏上
func (m *BaseWindow) SetNotInTaskBar() {
	m.TForm.SetShowInTaskBar(types.StNever)
}

//返回chromium的父组件对象，该对象不是window组件对象,属于window的一个子组件
//
//在windows下它是 TCEFWindowParent, linux或macOSx下它是 TCEFLinkedWindowParent
//
//通过函数可调整该组件的属性
func (m *BaseWindow) WindowParent() ITCefWindow {
	return m.windowParent
}

//返回窗口关闭状态
func (m *BaseWindow) IsClosing() bool {
	return m.isClosing
}

// 设置窗口类型
func (m *BaseWindow) SetWindowType(windowType WINDOW_TYPE) {
	m.windowType = windowType
}

// 返回窗口类型
func (m *BaseWindow) WindowType() WINDOW_TYPE {
	return m.windowType
}

// 创建window浏览器组件
//
// 不带有默认事件的chromium
func (m *BaseWindow) ChromiumCreate(config *tCefChromiumConfig, defaultUrl string) {
	if m.isChromiumCreate {
		return
	}
	m.isChromiumCreate = true
	m.windowId = BrowserWindow.GetNextWindowNum()
	if config == nil {
		config = NewChromiumConfig()
		config.SetEnableMenu(true)
		config.SetEnableDevTools(true)
		config.SetEnableOpenUrlTab(true)
		config.SetEnableWindowPopup(true)
	}
	m.chromium = NewChromium(m, config)
	m.chromium.SetEnableMultiBrowserMode(true)
	if defaultUrl != "" {
		m.chromium.SetDefaultURL(defaultUrl)
	}
	//windowParent
	m.windowParent = NewCEFWindow(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)
	m.windowParent.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	m.windowParent.SetChromium(m.chromium, 0)
	m.windowParent.SetOnEnter(func(sender lcl.IObject) {
		if m.isClosing {
			return
		}
		m.chromium.Initialized()
		m.chromium.FrameIsFocused()
		m.chromium.SetFocus(true)
	})
	m.windowParent.SetOnExit(func(sender lcl.IObject) {
		if m.isClosing {
			return
		}
		m.chromium.SendCaptureLostEvent()
	})
}

func (m *BaseWindow) putChromiumWindowInfo() {
	m.windowInfo = &TCefWindowInfo{
		Window:         m,
		Browser:        nil,
		Frames:         make(map[int64]*ICefFrame),
		WindowProperty: NewWindowProperty(),
		auxTools:       &auxTools{},
	}
	BrowserWindow.putWindowInfo(m.windowId, m.windowInfo)
}

//默认的chromium事件
func (m *BaseWindow) defaultChromiumEvent() {
	if m.WindowType() != WT_DEV_TOOLS {
		AddGoForm(m.windowId, m.Instance())
		m.registerDefaultEvent()
		m.registerDefaultChromiumCloseEvent()
	}
}

// 创建窗口
//
// 不带有默认事件的窗口
func (m *BaseWindow) FormCreate() {
	if m.isFormCreate {
		return
	}
	m.isFormCreate = true
	m.SetPosition(types.PoDesktopCenter)
	m.SetName(fmt.Sprintf("energy_window_name_%d", time.Now().UnixNano()/1e6))
}

//默认窗口活动/关闭处理事件
func (m *BaseWindow) defaultWindowEvent() {
	if m.WindowType() != WT_DEV_TOOLS {
		m.SetOnResize(m.resize)
		m.SetOnActivate(m.activate)
	}
	m.SetOnShow(m.show)
}

//默认的窗口关闭事件
func (m *BaseWindow) defaultWindowCloseEvent() {
	m.SetOnClose(m.close)
	m.SetOnCloseQuery(m.closeQuery)
}

// 添加OnResize事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *BaseWindow) AddOnResize(fn TNotifyEvent) {
	m.onResize = append(m.onResize, fn)
}

// 添加OnActivate事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *BaseWindow) AddOnActivate(fn TNotifyEvent) {
	m.onActivate = append(m.onActivate, fn)
}

// 添加OnShow事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *BaseWindow) AddOnShow(fn TNotifyEvent) {
	m.onShow = append(m.onShow, fn)
}

// 添加OnClose事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *BaseWindow) AddOnClose(fn TCloseEvent) {
	m.onClose = append(m.onClose, fn)
}

// 添加OnCloseQuery事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *BaseWindow) AddOnCloseQuery(fn TCloseQueryEvent) {
	m.onCloseQuery = append(m.onCloseQuery, fn)
}

//每次激活窗口之后执行一次
func (m *BaseWindow) SetOnActivateAfter(fn lcl.TNotifyEvent) {
	m.onActivateAfter = fn
}

func (m *BaseWindow) show(sender lcl.IObject) {
	var ret bool
	if m.onShow != nil {
		for _, fn := range m.onShow {
			if fn(sender) {
				ret = true
			}
		}
	}
	if !ret {
		if m.windowParent != nil {
			QueueAsyncCall(func(id int) {
				m.windowParent.UpdateSize()
			})
		}
	}
}

func (m *BaseWindow) resize(sender lcl.IObject) {
	var ret bool
	if m.onResize != nil {
		for _, fn := range m.onResize {
			if fn(sender) {
				ret = true
			}
		}
	}
	if !ret {
		if m.isClosing {
			return
		}
		if m.chromium != nil {
			m.chromium.NotifyMoveOrResizeStarted()
		}
		if m.windowParent != nil {
			m.windowParent.UpdateSize()
		}
	}
}

func (m *BaseWindow) activate(sender lcl.IObject) {
	var ret bool
	if m.onActivate != nil {
		for _, fn := range m.onActivate {
			if fn(sender) {
				ret = true
			}
		}
	}
	if !ret {
		if m.isClosing {
			return
		}
		if m.chromium != nil {
			if !m.chromium.Initialized() {
				m.chromium.CreateBrowser(m.windowParent)
			}
		}
	}
	if m.onActivateAfter != nil {
		m.onActivateAfter(sender)
	}
}

// 默认事件注册 部分事件允许被覆盖
func (m *BaseWindow) registerDefaultEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
		if !api.GoBool(BrowserWindow.Config.chromiumConfig.enableWindowPopup) {
			return true
		}
		BrowserWindow.popupWindow.SetWindowType(WT_POPUP_SUB_BROWSER)
		BrowserWindow.popupWindow.ChromiumCreate(BrowserWindow.Config.chromiumConfig, beforePopupInfo.TargetUrl)
		BrowserWindow.popupWindow.chromium.EnableIndependentEvent()
		BrowserWindow.popupWindow.putChromiumWindowInfo()
		BrowserWindow.popupWindow.defaultChromiumEvent()
		var result = false
		defer func() {
			if result {
				QueueAsyncCall(func(id int) {
					BrowserWindow.uiLock.Lock()
					defer BrowserWindow.uiLock.Unlock()
					winProperty := BrowserWindow.popupWindow.windowInfo.WindowProperty
					if winProperty != nil {
						if winProperty.IsShowModel {
							BrowserWindow.popupWindow.ShowModal()
							return
						}
					}
					BrowserWindow.popupWindow.Show()
				})
			}
		}()
		if bwEvent.onBeforePopup != nil {
			result = !bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, BrowserWindow.popupWindow.windowInfo, noJavascriptAccess)
		}
		return result
	})
	m.chromium.SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, message *ipc.ICefProcessMessage) bool {
		if bwEvent.onProcessMessageReceived != nil {
			return bwEvent.onProcessMessageReceived(sender, browser, frame, sourceProcess, message)
		}
		return false
	})
	m.chromium.SetOnFrameCreated(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) {
		QueueAsyncCall(func(id int) {
			BrowserWindow.putBrowserFrame(browser, frame)
		})
		if bwEvent.onFrameCreated != nil {
			bwEvent.onFrameCreated(sender, browser, frame)
		}
	})
	m.chromium.SetOnFrameDetached(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) {
		chromiumOnFrameDetached(browser, frame)
		if bwEvent.onFrameDetached != nil {
			bwEvent.onFrameDetached(sender, browser, frame)
		}
	})
	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {
		if chromiumOnAfterCreate(browser) {
			return
		}
		if bwEvent.onAfterCreated != nil {
			bwEvent.onAfterCreated(sender, browser)
		}
	})
	m.chromium.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *TCefReturnValue) {
		if assetserve.AssetsServerHeaderKeyValue != "" {
			request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
		}
		if bwEvent.onBeforeResourceLoad != nil {
			bwEvent.onBeforeResourceLoad(sender, browser, frame, request, callback, result)
		}
	})
	//事件可以被覆盖
	m.chromium.SetOnBeforeDownload(func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *DownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) {
		if bwEvent.onBeforeDownload != nil {
			bwEvent.onBeforeDownload(sender, browser, beforeDownloadItem, suggestedName, callback)
		} else {
			callback.Cont(ExePath+Separator+suggestedName, true)
		}
	})
	//默认自定义快捷键
	defaultAcceleratorCustom()
	//事件可以被覆盖
	m.chromium.SetOnKeyEvent(func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, result *bool) {
		if api.GoBool(BrowserWindow.Config.chromiumConfig.enableDevTools) {
			if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
				if winInfo.Window.WindowType() == WT_DEV_TOOLS || winInfo.Window.WindowType() == WT_VIEW_SOURCE {
					return
				}
			}
			if event.WindowsKeyCode == VkF12 && event.Kind == KEYEVENT_RAW_KEYDOWN {
				browser.ShowDevTools()
				*result = true
			} else if event.WindowsKeyCode == VkF12 && event.Kind == KEYEVENT_KEYUP {
				*result = true
			}
		}
		if KeyAccelerator.accelerator(browser, event, result) {
			return
		}
		if bwEvent.onKeyEvent != nil {
			bwEvent.onKeyEvent(sender, browser, event, result)
		}
	})
	m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool {
		chromiumOnBeforeBrowser(browser, frame)
		if bwEvent.onBeforeBrowser != nil {
			return bwEvent.onBeforeBrowser(sender, browser, frame)
		}
		return false
	})
	m.chromium.SetOnTitleChange(func(sender lcl.IObject, browser *ICefBrowser, title string) {
		updateBrowserDevTools(browser, title)
		updateBrowserViewSource(browser, title)
		if bwEvent.onTitleChange != nil {
			bwEvent.onTitleChange(sender, browser, title)
		}
	})
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		chromiumOnBeforeContextMenu(sender, browser, frame, params, model)
		if bwEvent.onBeforeContextMenu != nil {
			bwEvent.onBeforeContextMenu(sender, browser, frame, params, model)
		}
	})
	m.chromium.SetOnContextMenuCommand(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId MenuId, eventFlags uint32, result *bool) {
		chromiumOnContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		if bwEvent.onContextMenuCommand != nil {
			bwEvent.onContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		}
	})
	m.chromium.SetOnLoadingStateChange(func(sender lcl.IObject, browser *ICefBrowser, isLoading, canGoBack, canGoForward bool) {
		//BrowserWindow.putBrowserFrame(browser, nil)
		if bwEvent.onLoadingStateChange != nil {
			bwEvent.onLoadingStateChange(sender, browser, isLoading, canGoBack, canGoForward)
		}
	})
}
func (m *BaseWindow) close(sender lcl.IObject, action *types.TCloseAction) {
	var ret bool
	if m.onClose != nil {
		for _, fn := range m.onClose {
			if fn(sender, action) {
				ret = true
			}
		}
	}
	if !ret {
		logger.Debug("window.onClose")
		*action = types.CaFree
	}
}

func (m *BaseWindow) closeQuery(sender lcl.IObject, close *bool) {
	var ret bool
	if m.onCloseQuery != nil {
		for _, fn := range m.onCloseQuery {
			if fn(sender, close) {
				ret = true
			}
		}
	}
	if !ret {
		logger.Debug("window.onCloseQuery windowType:", m.WindowType())
		if IsDarwin() {
			//main window close
			if m.WindowType() == WT_MAIN_BROWSER {
				*close = true
				desChildWind := m.windowParent.DestroyChildWindow()
				logger.Debug("window.onCloseQuery => windowParent.DestroyChildWindow:", desChildWind)
			} else {
				//sub window close
				*close = m.canClose
			}
			if !m.isClosing {
				m.isClosing = true
				m.Hide()
				m.chromium.CloseBrowser(true)
			}
		} else {
			*close = m.canClose
			if !m.isClosing {
				m.isClosing = true
				m.Hide()
				m.chromium.CloseBrowser(true)
			}
		}
	}
}

//默认的chromium关闭事件
func (m *BaseWindow) registerDefaultChromiumCloseEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *TCefCloseBrowsesAction) {
		logger.Debug("chromium.onClose")
		if IsDarwin() { //MacOSX
			desChildWind := m.windowParent.DestroyChildWindow()
			logger.Debug("chromium.onClose => windowParent.DestroyChildWindow:", desChildWind)
			*aAction = CbaClose
		} else if IsLinux() {
			*aAction = CbaClose
		} else if IsWindows() {
			*aAction = CbaDelay
		}
		if !IsDarwin() {
			QueueAsyncCall(func(id int) { //main thread run
				m.windowParent.Free()
				logger.Debug("chromium.onClose => windowParent.Free")
			})
		}
		if bwEvent.onClose != nil {
			bwEvent.onClose(sender, browser, aAction)
		}
	})
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		logger.Debug("chromium.onBeforeClose")
		chromiumOnBeforeClose(browser)
		m.canClose = true
		var tempClose = func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("chromium.OnBeforeClose Error:", err)
				}
			}()
			if m.windowInfo.auxTools.viewSourceWindow != nil {
				m.windowInfo.auxTools.viewSourceWindow = nil
			}
			if m.windowInfo != nil && m.windowInfo.auxTools.devToolsWindow != nil {
				m.windowInfo.auxTools.devToolsWindow.Close()
			}
			BrowserWindow.removeWindowInfo(m.windowId)
			//主窗口关闭
			if m.WindowType() == WT_MAIN_BROWSER {
				if IsWindows() {
					rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
				} else {
					m.Close()
				}
			} else if IsDarwin() {
				m.Close()
			}
		}
		QueueAsyncCall(func(id int) { // main thread run
			tempClose()
		})
		if bwEvent.onBeforeClose != nil {
			bwEvent.onBeforeClose(sender, browser)
		}
	})
}
