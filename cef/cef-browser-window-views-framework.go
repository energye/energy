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
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
)

//ViewsFrameworkBrowserWindow 基于CEF views framework 窗口组件
//
//该窗口使用CEF内部实现，在linux下107.xx以后版本默认使用GTK3，但无法使用lcl组件集成到窗口中
//
//当创建应用配置时 MultiThreadedMessageLoop 和 ExternalMessagePump 属性同时为false(linux系统默认强制false)时启用ViewsFramework窗口
type ViewsFrameworkBrowserWindow struct {
	isClosing            bool                              //
	windowType           consts.WINDOW_TYPE                //0:browser 1:devTools 2:viewSource 默认:0
	windowId             int32                             //
	chromium             IChromium                         //
	browser              *ICefBrowser                      //
	component            lcl.IComponent                    //
	windowComponent      *TCEFWindowComponent              //
	browserViewComponent *TCEFBrowserViewComponent         //
	windowProperty       *WindowProperty                   //窗口属性
	frames               TCEFFrame                         //当前浏览器下的所有frame
	auxTools             *auxTools                         //辅助工具
	tray                 ITray                             //托盘
	doOnWindowCreated    WindowComponentOnWindowCreated    //
	doOnGetInitialBounds WindowComponentOnGetInitialBounds //窗口初始bounds
	regions              *TCefDraggableRegions             //窗口内html拖拽区域
}

//创建 ViewsFrameworkBrowserWindow 窗口
func NewViewsFrameworkBrowserWindow(chromiumConfig *tCefChromiumConfig, windowProperty WindowProperty, owner ...lcl.IComponent) *ViewsFrameworkBrowserWindow {
	if chromiumConfig == nil {
		chromiumConfig = NewChromiumConfig()
		chromiumConfig.SetEnableViewSource(false)
		chromiumConfig.SetEnableDevTools(false)
		chromiumConfig.SetEnableMenu(false)
		chromiumConfig.SetEnableWindowPopup(false)
	}

	var component lcl.IComponent
	if len(owner) > 0 {
		component = lcl.NewComponent(owner[0])
	} else {
		component = lcl.NewComponent(nil)
	}
	m := &ViewsFrameworkBrowserWindow{
		windowProperty:       &windowProperty,
		component:            component,
		chromium:             NewChromium(component, chromiumConfig),
		windowComponent:      NewWindowComponent(component),
		browserViewComponent: NewBrowserViewComponent(component),
	}
	m.chromium.SetEnableMultiBrowserMode(true)
	m.windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *ICefWindow) {
		if m.chromium.CreateBrowserByBrowserViewComponent(windowProperty.Url, m.browserViewComponent) {
			m.windowComponent.AddChildView(m.browserViewComponent)
			if windowProperty.Title != "" {
				m.windowComponent.SetTitle(windowProperty.Title)
			}
			if windowProperty.EnableCenterWindow {
				m.windowComponent.CenterWindow(NewCefSize(windowProperty.Width, windowProperty.Height))
			}
			if windowProperty.IconFS != "" {
				if emfs.IsExist(windowProperty.IconFS) {
					_ = m.windowComponent.SetWindowAppIconFS(1, windowProperty.IconFS)
				}
			} else if windowProperty.Icon != "" {
				if tools.IsExist(windowProperty.Icon) {
					_ = m.windowComponent.SetWindowAppIcon(1, windowProperty.Icon)
				}
			}
			m.browserViewComponent.RequestFocus()
			m.windowComponent.Show()
			if m.doOnWindowCreated != nil {
				m.doOnWindowCreated(sender, window)
			}
		}
	})
	return m
}

//ViewsFrameworkBrowserWindow 主窗口初始化
func (m *browser) appContextInitialized(app *TCEFApplication) {
	if !common.Args.IsMain() {
		return
	}
	app.SetOnContextInitialized(func() {
		vFrameBrowserWindow := NewViewsFrameworkBrowserWindow(m.Config.ChromiumConfig(), m.Config.WindowProperty)
		vFrameBrowserWindow.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
			if vFrameBrowserWindow.tray != nil {
				vFrameBrowserWindow.tray.close()
			}
		})
		vFrameBrowserWindow.ResetWindowPropertyForEvent()
		vFrameBrowserWindow.SetWindowType(consts.WT_MAIN_BROWSER)
		vFrameBrowserWindow.windowId = BrowserWindow.GetNextWindowNum()
		vFrameBrowserWindow.putChromiumWindowInfo()
		vFrameBrowserWindow.registerPopupEvent()
		vFrameBrowserWindow.registerDefaultEvent()
		vFrameBrowserWindow.windowComponent.SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
			*aResult = true
			app.QuitMessageLoop()
		})
		vFrameBrowserWindow.doOnWindowCreated = func(sender lcl.IObject, window *ICefWindow) {
			if m.Config.browserWindowAfterOnEventCallback != nil {
				m.Config.browserWindowAfterOnEventCallback(vFrameBrowserWindow)
			}
		}
		BrowserWindow.mainVFBrowserWindow = vFrameBrowserWindow
		if m.Config.browserWindowOnEventCallback != nil {
			BrowserWindow.browserEvent.chromium = vFrameBrowserWindow.chromium
			m.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, vFrameBrowserWindow)
		}
		vFrameBrowserWindow.windowComponent.CreateTopLevelWindow()
	})
}

func (m *ViewsFrameworkBrowserWindow) registerPopupEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
		if !api.GoBool(BrowserWindow.Config.chromiumConfig.enableWindowPopup) {
			return true
		}
		var vfbw = NewViewsFrameworkBrowserWindow(BrowserWindow.Config.ChromiumConfig(), BrowserWindow.Config.WindowProperty, BrowserWindow.MainWindow().AsViewsFrameworkBrowserWindow().Component())
		var result = false
		if bwEvent.onBeforePopup != nil {
			result = bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, vfbw, noJavascriptAccess)
		}
		if !result {
			vfbw.ResetWindowPropertyForEvent()
			vfbw.SetWindowType(consts.WT_POPUP_SUB_BROWSER)
			vfbw.windowId = BrowserWindow.GetNextWindowNum()
			vfbw.putChromiumWindowInfo()
			vfbw.registerPopupEvent()
			vfbw.registerDefaultEvent()
			vfbw.windowComponent.CreateTopLevelWindow()
			result = true
		}
		return result
	})
}

//重置窗口属性-通过事件函数
func (m *ViewsFrameworkBrowserWindow) ResetWindowPropertyForEvent() {
	wp := m.WindowProperty()
	m.windowComponent.SetOnGetInitialBounds(func(sender lcl.IObject, window *ICefWindow, aResult *TCefRect) {
		if wp.EnableCenterWindow {
			m.windowComponent.CenterWindow(NewCefSize(wp.Width, wp.Height))
			aResult.Width = wp.Width
			aResult.Height = wp.Height
		} else {
			aResult.X = wp.X
			aResult.Y = wp.Y
			aResult.Width = wp.Width
			aResult.Height = wp.Height
		}
		if m.doOnGetInitialBounds != nil {
			m.doOnGetInitialBounds(sender, window, aResult)
		}
	})
	m.windowComponent.SetOnCanMinimize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableMinimize
	})
	m.windowComponent.SetOnCanResize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableResize
	})
	m.windowComponent.SetOnCanMaximize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableMaximize
	})
	m.windowComponent.SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableClose
	})
	m.windowComponent.SetAlwaysOnTop(wp.AlwaysOnTop)
	m.windowComponent.SetBounds(NewCefRect(wp.X, wp.Y, wp.Width, wp.Height))
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
	//事件可以覆盖
	m.chromium.SetOnBeforeDownload(func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *DownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) {
		if bwEvent.onBeforeDownload != nil {
			bwEvent.onBeforeDownload(sender, browser, beforeDownloadItem, suggestedName, callback)
		} else {
			callback.Cont(consts.ExePath+consts.Separator+suggestedName, true)
		}
	})
	//事件可以覆盖
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		if bwEvent.onBeforeContextMenu != nil {
			bwEvent.onBeforeContextMenu(sender, browser, frame, params, model)
		} else {
			chromiumOnBeforeContextMenu(sender, browser, frame, params, model)
		}
	})
	//事件可以覆盖
	m.chromium.SetOnContextMenuCommand(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
		if bwEvent.onContextMenuCommand != nil {
			bwEvent.onContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		} else {
			chromiumOnContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		}
	})
	m.chromium.SetOnFrameCreated(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) {
		BrowserWindow.putBrowserFrame(browser, frame)
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
		if bwEvent.onKeyEvent != nil {
			bwEvent.onKeyEvent(sender, browser, event, result)
		} else {
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
		if bwEvent.onTitleChange != nil {
			bwEvent.onTitleChange(sender, browser, title)
		}
	})
	m.chromium.SetOnDragEnter(func(sender lcl.IObject, browser *ICefBrowser, dragData *ICefDragData, mask consts.TCefDragOperations, result *bool) {
		*result = !m.WindowProperty().EnableDragFile
		if bwEvent.onDragEnter != nil {
			bwEvent.onDragEnter(sender, browser, dragData, mask, result)
		}
	})
	m.chromium.SetOnLoadEnd(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32) {
		if bwEvent.onLoadEnd != nil {
			bwEvent.onLoadEnd(sender, browser, frame, httpStatusCode)
		}
	})
	if m.WindowProperty().EnableWebkitAppRegion {
		m.chromium.SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, regions *TCefDraggableRegions) {
			if bwEvent.onDraggableRegionsChanged != nil {
				bwEvent.onDraggableRegionsChanged(sender, browser, frame, regions)
			}
			m.regions = regions
			m.windowComponent.SetDraggableRegions(regions.Regions())
		})
	}
}

//启用所有默认事件行为
func (m *ViewsFrameworkBrowserWindow) EnableAllDefaultEvent() {
	m.registerPopupEvent()
	m.registerDefaultEvent()
}

func (m *ViewsFrameworkBrowserWindow) SetOnWindowCreated(onWindowCreated WindowComponentOnWindowCreated) {
	m.doOnWindowCreated = onWindowCreated
}

func (m *ViewsFrameworkBrowserWindow) SetOnGetInitialBounds(onGetInitialBounds WindowComponentOnGetInitialBounds) {
	m.doOnGetInitialBounds = onGetInitialBounds
}

func (m *ViewsFrameworkBrowserWindow) IsViewsFramework() bool {
	return true
}

func (m *ViewsFrameworkBrowserWindow) IsLCL() bool {
	return false
}

func (m *ViewsFrameworkBrowserWindow) WindowProperty() *WindowProperty {
	return m.windowProperty
}

func (m *ViewsFrameworkBrowserWindow) putChromiumWindowInfo() {
	BrowserWindow.putWindowInfo(m.windowId, m)
}

func (m *ViewsFrameworkBrowserWindow) BrowserWindow() *ViewsFrameworkBrowserWindow {
	return m
}

func (m *ViewsFrameworkBrowserWindow) Handle() types.HWND {
	return types.HWND(m.WindowComponent().WindowHandle().ToPtr())
}

func (m *ViewsFrameworkBrowserWindow) AsViewsFrameworkBrowserWindow() IViewsFrameworkBrowserWindow {
	return m
}

func (m *ViewsFrameworkBrowserWindow) AsLCLBrowserWindow() ILCLBrowserWindow {
	return nil
}

func (m *ViewsFrameworkBrowserWindow) SetTitle(title string) {
	m.WindowProperty().Title = title
}

func (m *ViewsFrameworkBrowserWindow) SetWidth(value int32) {
	m.WindowProperty().Width = value
}

func (m *ViewsFrameworkBrowserWindow) SetHeight(value int32) {
	m.WindowProperty().Height = value
}

func (m *ViewsFrameworkBrowserWindow) Point() *TCefPoint {
	result := m.WindowComponent().Position()
	m.WindowProperty().X = result.X
	m.WindowProperty().Y = result.Y
	return result
}

func (m *ViewsFrameworkBrowserWindow) Size() *TCefSize {
	result := m.WindowComponent().Size()
	m.WindowProperty().Width = result.Width
	m.WindowProperty().Height = result.Height
	return result
}

func (m *ViewsFrameworkBrowserWindow) Bounds() *TCefRect {
	result := m.WindowComponent().Bounds()
	m.WindowProperty().X = result.X
	m.WindowProperty().Y = result.Y
	m.WindowProperty().Width = result.Width
	m.WindowProperty().Height = result.Height
	return result
}

func (m *ViewsFrameworkBrowserWindow) SetPoint(x, y int32) {
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
}

func (m *ViewsFrameworkBrowserWindow) SetSize(width, height int32) {
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
}

func (m *ViewsFrameworkBrowserWindow) SetBounds(x, y, width, height int32) {
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
}

func (m *ViewsFrameworkBrowserWindow) getAuxTools() *auxTools {
	return m.auxTools
}

func (m *ViewsFrameworkBrowserWindow) createAuxTools() {
	if m.auxTools == nil {
		m.auxTools = &auxTools{}
	}
}

func (m *ViewsFrameworkBrowserWindow) Browser() *ICefBrowser {
	return m.browser
}

func (m *ViewsFrameworkBrowserWindow) Frames() TCEFFrame {
	return m.frames
}

func (m *ViewsFrameworkBrowserWindow) createFrames() {
	if m.frames == nil {
		m.frames = make(TCEFFrame)
	}
}

func (m *ViewsFrameworkBrowserWindow) setBrowser(browser *ICefBrowser) {
	m.browser = browser
}

func (m *ViewsFrameworkBrowserWindow) addFrame(frame *ICefFrame) {
	m.createFrames()
	m.frames[frame.Id] = frame
}

func (m *ViewsFrameworkBrowserWindow) Chromium() IChromium {
	return m.chromium
}

func (m *ViewsFrameworkBrowserWindow) Id() int32 {
	return m.windowId
}

func (m *ViewsFrameworkBrowserWindow) Show() {
	m.WindowComponent().Show()
}

func (m *ViewsFrameworkBrowserWindow) Hide() {
	m.WindowComponent().Hide()
}

func (m *ViewsFrameworkBrowserWindow) Close() {
	m.WindowComponent().Close()
}

func (m *ViewsFrameworkBrowserWindow) Maximize() {
	m.WindowComponent().Maximize()
}

func (m *ViewsFrameworkBrowserWindow) Minimize() {
	m.WindowComponent().Minimize()
}

func (m *ViewsFrameworkBrowserWindow) Restore() {
	m.WindowComponent().Restore()
}

func (m *ViewsFrameworkBrowserWindow) CloseBrowserWindow() {
	m.isClosing = true
	m.chromium.CloseBrowser(true)
}

func (m *ViewsFrameworkBrowserWindow) CreateTopLevelWindow() {
	m.WindowComponent().CreateTopLevelWindow()
}

func (m *ViewsFrameworkBrowserWindow) CenterWindow(size *TCefSize) {
	m.WindowComponent().CenterWindow(size)
}

func (m *ViewsFrameworkBrowserWindow) SetCenterWindow(value bool) {
	m.WindowProperty().EnableCenterWindow = value
}

//返回窗口关闭状态
func (m *ViewsFrameworkBrowserWindow) IsClosing() bool {
	return m.isClosing
}

// 返回窗口类型
func (m *ViewsFrameworkBrowserWindow) WindowType() consts.WINDOW_TYPE {
	return m.windowType
}

// 设置窗口类型
func (m *ViewsFrameworkBrowserWindow) SetWindowType(windowType consts.WINDOW_TYPE) {
	m.windowType = windowType
}

//禁用最小化按钮
func (m *ViewsFrameworkBrowserWindow) DisableMinimize() {
	m.WindowProperty().EnableMinimize = false
}

//禁用最大化按钮
func (m *ViewsFrameworkBrowserWindow) DisableMaximize() {
	m.WindowProperty().EnableMaximize = false
}

//禁用调整窗口大小
func (m *ViewsFrameworkBrowserWindow) DisableResize() {
	m.WindowProperty().EnableResize = false
}

//启用最小化按钮
func (m *ViewsFrameworkBrowserWindow) EnableMinimize() {
	m.WindowProperty().EnableMinimize = true
}

//启用最大化按钮
func (m *ViewsFrameworkBrowserWindow) EnableMaximize() {
	m.WindowProperty().EnableMaximize = true
}

//启用调整窗口大小
func (m *ViewsFrameworkBrowserWindow) EnableResize() {
	m.WindowProperty().EnableResize = true
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
