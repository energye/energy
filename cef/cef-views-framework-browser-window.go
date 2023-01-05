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
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

//基于CEF views framework窗口
//
//该窗口使用CEF内部实现，在linux下107.xx以后版本默认使用GTK3，但无法使用lcl组件集成到窗口中
//
//当创建应用配置时 MultiThreadedMessageLoop 和 ExternalMessagePump 属性同时为false(linux系统默认强制false)时启用ViewsFramework窗口
type ViewsFrameworkBrowserWindow struct {
	chromium             IChromium                 //
	component            lcl.IComponent            //
	windowComponent      *TCEFWindowComponent      //
	browserViewComponent *TCEFBrowserViewComponent //
	windowProperty       *WindowProperty           //窗口属性
}

//创建 ViewsFrameworkBrowserWindow 窗口
func NewViewsFrameworkBrowserWindow(chromiumConfig *tCefChromiumConfig, windowProperty *WindowProperty) *ViewsFrameworkBrowserWindow {
	if chromiumConfig == nil {
		chromiumConfig = BrowserWindow.Config.ChromiumConfig()
	}
	component := lcl.NewComponent(nil)
	m := &ViewsFrameworkBrowserWindow{
		windowProperty:       windowProperty,
		component:            component,
		chromium:             NewChromium(component, chromiumConfig),
		windowComponent:      NewWindowComponent(component),
		browserViewComponent: NewBrowserViewComponent(component),
	}
	m.chromium.SetEnableMultiBrowserMode(true)
	m.registerPopupEvent()
	m.windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *ICefWindow) {
		if m.chromium.CreateBrowserByBrowserViewComponent(windowProperty.Url, m.browserViewComponent) {
			m.windowComponent.AddChildView(m.browserViewComponent)
			m.windowComponent.SetTitle(windowProperty.Title)
			if windowProperty.CenterWindow {
				window.CenterWindow(NewCefSize(windowProperty.Width, windowProperty.Height))
			}
			if windowProperty.IconFS != "" {
				m.windowComponent.SetWindowAppIconFS(1, windowProperty.IconFS)
			} else if windowProperty.Icon != "" {
				m.windowComponent.SetWindowAppIcon(1, windowProperty.Icon)
			}
			m.browserViewComponent.RequestFocus()
			m.windowComponent.Show()
		}
	})
	if !windowProperty.CenterWindow {
		m.windowComponent.SetOnGetInitialBounds(func(sender lcl.IObject, window *ICefWindow, aResult *TCefRect) {
			aResult.X = windowProperty.X
			aResult.Y = windowProperty.Y
			aResult.Width = windowProperty.Width
			aResult.Height = windowProperty.Height
		})
	}
	m.windowComponent.SetOnCanMinimize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = windowProperty.CanMinimize
	})
	m.windowComponent.SetOnCanResize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = windowProperty.CanResize
	})
	m.windowComponent.SetOnCanMaximize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = windowProperty.CanMaximize
	})
	m.windowComponent.SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = windowProperty.CanClose
	})
	m.windowComponent.SetAlwaysOnTop(windowProperty.AlwaysOnTop)
	return m
}

func (m *ViewsFrameworkBrowserWindow) Component() lcl.IComponent {
	return m.component
}

func (m *ViewsFrameworkBrowserWindow) WindowComponent() *TCEFWindowComponent {
	return m.windowComponent
}

func (m *ViewsFrameworkBrowserWindow) BrowserViewComponent() *TCEFBrowserViewComponent {
	return m.browserViewComponent
}

//ViewsFrameworkBrowserWindow 主窗口初始化
func (m *browserWindow) appContextInitialized(app *TCEFApplication) {
	if !common.Args.IsMain() {
		return
	}
	app.SetOnContextInitialized(func() {
		vFrameBrowserWindow := NewViewsFrameworkBrowserWindow(BrowserWindow.Config.ChromiumConfig(), &BrowserWindow.Config.WindowProperty)
		m.chromium = vFrameBrowserWindow.chromium
		m.putChromiumWindowInfo()
		vFrameBrowserWindow.registerPopupEvent()
		vFrameBrowserWindow.registerDefaultEvent()
		vFrameBrowserWindow.windowComponent.SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
			fmt.Println("OnCanClose")
			*aResult = true
			app.QuitMessageLoop()
		})
		if BrowserWindow.Config.viewsFrameBrowserWindowOnEventCallback != nil {
			BrowserWindow.Config.viewsFrameBrowserWindowOnEventCallback(BrowserWindow.browserEvent, vFrameBrowserWindow)
		}
		vFrameBrowserWindow.windowComponent.CreateTopLevelWindow()
		fmt.Println("SetOnContextInitialized IsMessageLoop:", consts.IsMessageLoop)
	})
}

func (m *ViewsFrameworkBrowserWindow) registerPopupEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
		if !api.GoBool(BrowserWindow.Config.chromiumConfig.enableWindowPopup) {
			return true
		}
		fmt.Println("BrowserWindow-TargetUrl:", beforePopupInfo.TargetUrl, "IsMessageLoop:", consts.IsMessageLoop)
		var result = false
		if bwEvent.onBeforePopup != nil {
			result = !bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, BrowserWindow.popupWindow, noJavascriptAccess)
		}
		if !result {
			result = true
			wp := &WindowProperty{
				Title:        BrowserWindow.Config.WindowProperty.Title,
				Url:          beforePopupInfo.TargetUrl,
				CanMinimize:  BrowserWindow.Config.WindowProperty.CanMinimize,
				CanMaximize:  BrowserWindow.Config.WindowProperty.CanMaximize,
				CanResize:    BrowserWindow.Config.WindowProperty.CanResize,
				CanClose:     BrowserWindow.Config.WindowProperty.CanClose,
				CenterWindow: BrowserWindow.Config.WindowProperty.CenterWindow,
				IsShowModel:  BrowserWindow.Config.WindowProperty.IsShowModel,
				WindowState:  BrowserWindow.Config.WindowProperty.WindowState,
				Icon:         BrowserWindow.Config.WindowProperty.Icon,
				IconFS:       BrowserWindow.Config.WindowProperty.IconFS,
				X:            BrowserWindow.Config.WindowProperty.X,
				Y:            BrowserWindow.Config.WindowProperty.Y,
				Width:        BrowserWindow.Config.WindowProperty.Width,
				Height:       BrowserWindow.Config.WindowProperty.Height,
			}
			vFrameBrowserWindow := NewViewsFrameworkBrowserWindow(BrowserWindow.Config.ChromiumConfig(), wp)
			vFrameBrowserWindow.registerPopupEvent()
			vFrameBrowserWindow.registerDefaultEvent()
			vFrameBrowserWindow.windowComponent.CreateTopLevelWindow()

		}
		return result
	})
}

func (m *ViewsFrameworkBrowserWindow) registerDefaultEvent() {
	var bwEvent = BrowserWindow.browserEvent
	//默认自定义快捷键
	defaultAcceleratorCustom()
	m.chromium.SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ipc.ICefProcessMessage) bool {
		if bwEvent.onProcessMessageReceived != nil {
			return bwEvent.onProcessMessageReceived(sender, browser, frame, sourceProcess, message)
		}
		return false
	})
	m.chromium.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue) {
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
			callback.Cont(consts.ExePath+consts.Separator+suggestedName, true)
		}
	})
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		chromiumOnBeforeContextMenu(sender, browser, frame, params, model)
		if bwEvent.onBeforeContextMenu != nil {
			bwEvent.onBeforeContextMenu(sender, browser, frame, params, model)
		}
	})
	m.chromium.SetOnContextMenuCommand(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
		chromiumOnContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		if bwEvent.onContextMenuCommand != nil {
			bwEvent.onContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		}
	})
	m.chromium.SetOnLoadingStateChange(func(sender lcl.IObject, browser *ICefBrowser, isLoading, canGoBack, canGoForward bool) {
		if bwEvent.onLoadingStateChange != nil {
			bwEvent.onLoadingStateChange(sender, browser, isLoading, canGoBack, canGoForward)
		}
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
	//事件可以被覆盖
	m.chromium.SetOnKeyEvent(func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, result *bool) {
		if api.GoBool(BrowserWindow.Config.chromiumConfig.enableDevTools) {
			if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
				if winInfo.WindowType() == consts.WT_DEV_TOOLS || winInfo.WindowType() == consts.WT_VIEW_SOURCE {
					return
				}
			}
			if event.WindowsKeyCode == VkF12 && event.Kind == consts.KEYEVENT_RAW_KEYDOWN {
				browser.ShowDevTools()
				*result = true
			} else if event.WindowsKeyCode == VkF12 && event.Kind == consts.KEYEVENT_KEYUP {
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
}
