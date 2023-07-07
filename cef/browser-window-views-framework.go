//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// VF窗口组件定义和实现

package cef

import (
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/internal/window"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// ViewsFrameworkBrowserWindow 基于CEF views framework 窗口组件
//
// 该窗口使用CEF内部实现，在linux下107.xx以后版本默认使用GTK3，但无法使用lcl组件集成到窗口中
//
// 当创建应用配置时 MultiThreadedMessageLoop 和 ExternalMessagePump 属性同时为false(linux系统默认强制false)时启用ViewsFramework窗口
type ViewsFrameworkBrowserWindow struct {
	isClosing            bool                              //
	windowType           consts.WINDOW_TYPE                //0:browser 1:devTools 2:viewSource 默认:0
	windowId             int32                             //
	chromium             IChromium                         //
	component            lcl.IComponent                    //
	windowComponent      *TCEFWindowComponent              //
	browserViewComponent *TCEFBrowserViewComponent         //
	windowProperty       *WindowProperty                   //窗口属性
	auxTools             *auxTools                         //辅助工具
	tray                 ITray                             //托盘
	doOnWindowCreated    WindowComponentOnWindowCreated    //窗口创建
	doOnGetInitialBounds WindowComponentOnGetInitialBounds //窗口初始bounds
	regions              *TCefDraggableRegions             //窗口内html拖拽区域
	context              *ICefRequestContext
	extraInfo            *ICefDictionaryValue
}

// NewViewsFrameworkBrowserWindow 创建 ViewsFrameworkBrowserWindow 窗口
func NewViewsFrameworkBrowserWindow(config *TCefChromiumConfig, windowProperty WindowProperty, owner ...lcl.IComponent) *ViewsFrameworkBrowserWindow {
	if config == nil {
		config = NewChromiumConfig()
		config.SetEnableViewSource(false)
		config.SetEnableDevTools(false)
		config.SetEnableMenu(false)
		config.SetEnableWindowPopup(false)
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
		chromium:             NewChromium(component, config),
		windowComponent:      WindowComponentRef.New(component),
		browserViewComponent: BrowserViewComponentRef.New(component),
	}
	m.SetWindowType(windowProperty.WindowType)
	m.chromium.SetEnableMultiBrowserMode(true)
	m.windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *ICefWindow) {
		if m.chromium.CreateBrowserByBrowserViewComponent(windowProperty.Url, m.browserViewComponent, m.context, m.extraInfo) {
			m.windowComponent.AddChildView(m.browserViewComponent)
			if windowProperty.Title != "" {
				m.windowComponent.SetTitle(windowProperty.Title)
			}
			if windowProperty.EnableCenterWindow {
				m.windowComponent.CenterWindow(NewCefSize(windowProperty.Width, windowProperty.Height))
			}
			if windowProperty.IconFS != "" {
				if emfs.IsExist(windowProperty.IconFS) {
					if err := m.windowComponent.SetWindowAppIconByFSFile(1, windowProperty.IconFS); err != nil {
						logger.Error("set window application icon error:", err.Error())
					}
				}
			} else if windowProperty.Icon != "" {
				if tools.IsExist(windowProperty.Icon) {
					if err := m.windowComponent.SetWindowAppIconByFile(1, windowProperty.Icon); err != nil {
						logger.Error("set window application icon error:", err.Error())
					}
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

// ViewsFrameworkBrowserWindow 主窗口初始化
func (m *browserWindow) appContextInitialized(app *TCEFApplication) {
	if !process.Args.IsMain() {
		return
	}
	app.SetOnContextInitialized(func() {
		m.Config.WindowProperty.WindowType = consts.WT_MAIN_BROWSER
		vFrameBrowserWindow := NewViewsFrameworkBrowserWindow(m.Config.ChromiumConfig(), m.Config.WindowProperty)
		vFrameBrowserWindow.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
			chromiumOnBeforeClose(browser)
			if vFrameBrowserWindow.tray != nil {
				vFrameBrowserWindow.tray.close()
			}
		})
		// 重置窗口属性, 注册默认实现事件
		vFrameBrowserWindow.ResetWindowPropertyForEvent()
		vFrameBrowserWindow.registerPopupEvent()
		vFrameBrowserWindow.registerDefaultEvent()
		vFrameBrowserWindow.windowComponent.SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
			*aResult = true
			app.QuitMessageLoop()
		})
		BrowserWindow.mainVFBrowserWindow = vFrameBrowserWindow
		if m.Config.browserWindowOnEventCallback != nil {
			BrowserWindow.browserEvent.chromium = vFrameBrowserWindow.chromium
			m.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, vFrameBrowserWindow)
		}
		ipc.SetProcessMessage(vFrameBrowserWindow.Chromium().(*TCEFChromium))
		vFrameBrowserWindow.CreateTopLevelWindow()
	})
}

// registerPopupEvent 注册弹出子窗口事件
func (m *ViewsFrameworkBrowserWindow) registerPopupEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		chromiumOnBeforeClose(browser)
	})
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
		if !BrowserWindow.Config.ChromiumConfig().EnableWindowPopup() {
			return true
		}
		wp := BrowserWindow.Config.WindowProperty
		wp.Url = beforePopupInfo.TargetUrl
		wp.WindowType = consts.WT_POPUP_SUB_BROWSER
		var vfbw = NewViewsFrameworkBrowserWindow(BrowserWindow.Config.ChromiumConfig(), wp, BrowserWindow.MainWindow().AsViewsFrameworkBrowserWindow().Component())
		var result = false
		if bwEvent.onBeforePopup != nil {
			result = bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, vfbw, noJavascriptAccess)
		}
		if !result {
			vfbw.ResetWindowPropertyForEvent()
			vfbw.registerPopupEvent()
			vfbw.registerDefaultEvent()
			vfbw.CreateTopLevelWindow()
			result = true
		}
		return result
	})
}

// ResetWindowPropertyForEvent 重置窗口属性-通过事件函数
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
	m.windowComponent.SetOnIsFrameless(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableHideCaption
	})
	m.windowComponent.SetAlwaysOnTop(wp.AlwaysOnTop)
	m.windowComponent.SetBounds(NewCefRect(wp.X, wp.Y, wp.Width, wp.Height))
}

// registerDefaultEvent 注册默认事件
func (m *ViewsFrameworkBrowserWindow) registerDefaultEvent() {
	var bwEvent = BrowserWindow.browserEvent
	//默认自定义快捷键
	defaultAcceleratorCustom()
	m.chromium.SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
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
	m.chromium.SetOnBeforeDownload(func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *ICefDownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) {
		if bwEvent.onBeforeDownload != nil {
			bwEvent.onBeforeDownload(sender, browser, beforeDownloadItem, suggestedName, callback)
		} else {
			callback.Cont(consts.ExePath+consts.Separator+suggestedName, true)
		}
	})
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		if bwEvent.onBeforeContextMenu != nil {
			bwEvent.onBeforeContextMenu(sender, browser, frame, params, model)
		} else {
			chromiumOnBeforeContextMenu(sender, browser, frame, params, model)
		}
	})
	m.chromium.SetOnContextMenuCommand(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
		if bwEvent.onContextMenuCommand != nil {
			bwEvent.onContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		} else {
			chromiumOnContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, result)
		}
	})
	m.chromium.SetOnFrameCreated(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) {
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
	m.chromium.SetOnKeyEvent(func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, osEvent *consts.TCefEventHandle, result *bool) {
		if bwEvent.onKeyEvent != nil {
			bwEvent.onKeyEvent(sender, browser, event, osEvent, m, result)
		} else {
			if BrowserWindow.Config.ChromiumConfig().EnableDevTools() {
				if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
					if winInfo.WindowType() == consts.WT_DEV_TOOLS || winInfo.WindowType() == consts.WT_VIEW_SOURCE {
						return
					}
				}
				if event.WindowsKeyCode == consts.VkF12 && event.Kind == consts.KEYEVENT_RAW_KEYDOWN {
					browser.ShowDevTools()
					*result = true
				} else if event.WindowsKeyCode == consts.VkF12 && event.Kind == consts.KEYEVENT_KEYUP {
					*result = true
				}
			}
			if KeyAccelerator.accelerator(browser, event, result) {
				return
			}
		}
	})
	m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool) bool {
		if bwEvent.onBeforeBrowser != nil {
			return bwEvent.onBeforeBrowser(sender, browser, frame, request, userGesture, isRedirect, m)
		}
		return false
	})
	m.chromium.SetOnTitleChange(func(sender lcl.IObject, browser *ICefBrowser, title string) {
		if bwEvent.onTitleChange != nil {
			bwEvent.onTitleChange(sender, browser, title, m)
		}
	})
	m.chromium.SetOnDragEnter(func(sender lcl.IObject, browser *ICefBrowser, dragData *ICefDragData, mask consts.TCefDragOperations, result *bool) {
		*result = !m.WindowProperty().EnableDragFile
		if bwEvent.onDragEnter != nil {
			bwEvent.onDragEnter(sender, browser, dragData, mask, m, result)
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
				bwEvent.onDraggableRegionsChanged(sender, browser, frame, regions, m)
			}
			m.regions = regions
			m.windowComponent.SetDraggableRegions(regions.Regions())
		})
	}
}

// EnableAllDefaultEvent 启用所有默认事件行为
func (m *ViewsFrameworkBrowserWindow) EnableAllDefaultEvent() {
	m.registerPopupEvent()
	m.registerDefaultEvent()
}

// SetOnWindowCreated 窗口创建
func (m *ViewsFrameworkBrowserWindow) SetOnWindowCreated(onWindowCreated WindowComponentOnWindowCreated) {
	m.doOnWindowCreated = onWindowCreated
}

// SetOnGetInitialBounds 窗口初始坐标和大小
func (m *ViewsFrameworkBrowserWindow) SetOnGetInitialBounds(onGetInitialBounds WindowComponentOnGetInitialBounds) {
	m.doOnGetInitialBounds = onGetInitialBounds
}

// SetCreateBrowserExtraInfo
//  设置 Chromium 创建浏览器时设置的扩展信息
func (m *ViewsFrameworkBrowserWindow) SetCreateBrowserExtraInfo(_ string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) {
	m.context = context
	m.extraInfo = extraInfo
}

// IsViewsFramework 返回是否VF窗口组件，这里返回true
func (m *ViewsFrameworkBrowserWindow) IsViewsFramework() bool {
	return true
}

// IsLCL 返回是否LCL窗口组件，这里返回false
func (m *ViewsFrameworkBrowserWindow) IsLCL() bool {
	return false
}

// WindowProperty 部分提供部分窗口属性设置
func (m *ViewsFrameworkBrowserWindow) WindowProperty() *WindowProperty {
	return m.windowProperty
}

// BrowserWindow 返回VF窗口组件实现
func (m *ViewsFrameworkBrowserWindow) BrowserWindow() *ViewsFrameworkBrowserWindow {
	return m
}

// Handle 返回窗口句柄
func (m *ViewsFrameworkBrowserWindow) Handle() types.HWND {
	return types.HWND(m.WindowComponent().WindowHandle().ToPtr())
}

// AsViewsFrameworkBrowserWindow 转换为VF窗口组件，这里返回VF窗口组件
func (m *ViewsFrameworkBrowserWindow) AsViewsFrameworkBrowserWindow() IViewsFrameworkBrowserWindow {
	return m
}

// AsLCLBrowserWindow 转换为LCL窗口组件，这里返回nil
func (m *ViewsFrameworkBrowserWindow) AsLCLBrowserWindow() ILCLBrowserWindow {
	return nil
}

// SetTitle 设置窗口标题
func (m *ViewsFrameworkBrowserWindow) SetTitle(title string) {
	m.WindowProperty().Title = title
}

// SetWidth 设置窗口宽
func (m *ViewsFrameworkBrowserWindow) SetWidth(value int32) {
	m.WindowProperty().Width = value
}

// SetHeight 设置窗口高
func (m *ViewsFrameworkBrowserWindow) SetHeight(value int32) {
	m.WindowProperty().Height = value
}

// Point 返回窗口坐标
func (m *ViewsFrameworkBrowserWindow) Point() *TCefPoint {
	result := m.WindowComponent().Position()
	m.WindowProperty().X = result.X
	m.WindowProperty().Y = result.Y
	return result
}

// Size 返回窗口宽高
func (m *ViewsFrameworkBrowserWindow) Size() *TCefSize {
	result := m.WindowComponent().Size()
	m.WindowProperty().Width = result.Width
	m.WindowProperty().Height = result.Height
	return result
}

// Bounds 返回窗口坐标和宽高
func (m *ViewsFrameworkBrowserWindow) Bounds() *TCefRect {
	result := m.WindowComponent().Bounds()
	m.WindowProperty().X = result.X
	m.WindowProperty().Y = result.Y
	m.WindowProperty().Width = result.Width
	m.WindowProperty().Height = result.Height
	return result
}

// SetPoint 设置窗口坐标
func (m *ViewsFrameworkBrowserWindow) SetPoint(x, y int32) {
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
}

// SetSize 设置窗口宽高
func (m *ViewsFrameworkBrowserWindow) SetSize(width, height int32) {
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
}

// SetBounds 设置窗口坐标和宽高
func (m *ViewsFrameworkBrowserWindow) SetBounds(x, y, width, height int32) {
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
}

// getAuxTools 获取辅助工具-开发者工具
func (m *ViewsFrameworkBrowserWindow) getAuxTools() *auxTools {
	return m.auxTools
}

// createAuxTools 创建辅助工具-开发者工具
func (m *ViewsFrameworkBrowserWindow) createAuxTools() {
	if m.auxTools == nil {
		m.auxTools = &auxTools{}
	}
}

// Browser 返回browser
func (m *ViewsFrameworkBrowserWindow) Browser() *ICefBrowser {
	return m.Chromium().Browser()
}

// Chromium 返回 chromium
func (m *ViewsFrameworkBrowserWindow) Chromium() IChromium {
	return m.chromium
}

// Id 获取窗口ID
func (m *ViewsFrameworkBrowserWindow) Id() int32 {
	if m.windowId == 0 {
		m.windowId = m.Chromium().BrowserId()
	}
	return m.windowId
}

// Show 显示窗口
func (m *ViewsFrameworkBrowserWindow) Show() {
	m.WindowComponent().Show()
}

// Hide 隐藏窗口
func (m *ViewsFrameworkBrowserWindow) Hide() {
	m.WindowComponent().Hide()
}

// Close 关闭窗口,一搬不使用
func (m *ViewsFrameworkBrowserWindow) Close() {
	m.WindowComponent().Close()
}

// WindowState 返回窗口最小化、最大化、全屏状态
func (m *ViewsFrameworkBrowserWindow) WindowState() types.TWindowState {
	if m.windowComponent.IsMinimized() {
		return types.WsMinimized
	} else if m.windowComponent.IsMaximized() {
		return types.WsMaximized
	} else if m.windowComponent.IsFullscreen() {
		return types.WsFullScreen
	}
	return types.WsNormal
}

// Maximize 窗口最大化/还原
func (m *ViewsFrameworkBrowserWindow) Maximize() {
	m.windowProperty.windowState = m.WindowState()
	if m.windowProperty.windowState == types.WsNormal {
		m.WindowComponent().Maximize()
	} else {
		m.Restore()
	}
}

// Minimize 窗口最小化
func (m *ViewsFrameworkBrowserWindow) Minimize() {
	m.WindowComponent().Minimize()
}

// Restore 窗口还原
func (m *ViewsFrameworkBrowserWindow) Restore() {
	m.WindowComponent().Restore()
}

// CloseBrowserWindow 关闭浏览器窗口
func (m *ViewsFrameworkBrowserWindow) CloseBrowserWindow() {
	m.isClosing = true
	m.chromium.CloseBrowser(true)
}

// CreateTopLevelWindow 创建顶层窗口
func (m *ViewsFrameworkBrowserWindow) CreateTopLevelWindow() {
	if m.WindowType() != consts.WT_DEV_TOOLS {
		window.CurrentBrowseWindowCache = m
	}
	m.WindowComponent().CreateTopLevelWindow()
}

// CenterWindow 设置窗口居中，同时指定窗口大小
func (m *ViewsFrameworkBrowserWindow) CenterWindow(size *TCefSize) {
	m.WindowComponent().CenterWindow(size)
}

// SetCenterWindow 设置窗口居中显示
func (m *ViewsFrameworkBrowserWindow) SetCenterWindow(value bool) {
	m.WindowProperty().EnableCenterWindow = value
}

// IsClosing 返回窗口是否正在关闭/或已关闭 true正在或已关闭
func (m *ViewsFrameworkBrowserWindow) IsClosing() bool {
	return m.isClosing
}

// WindowType 返回窗口类型
func (m *ViewsFrameworkBrowserWindow) WindowType() consts.WINDOW_TYPE {
	return m.windowType
}

// SetWindowType 设置窗口类型
func (m *ViewsFrameworkBrowserWindow) SetWindowType(windowType consts.WINDOW_TYPE) {
	m.windowType = windowType
}

// DisableMinimize 禁用最小化按钮
func (m *ViewsFrameworkBrowserWindow) DisableMinimize() {
	m.WindowProperty().EnableMinimize = false
}

// DisableMaximize 禁用最大化按钮
func (m *ViewsFrameworkBrowserWindow) DisableMaximize() {
	m.WindowProperty().EnableMaximize = false
}

// DisableResize 禁用调整窗口大小
func (m *ViewsFrameworkBrowserWindow) DisableResize() {
	m.WindowProperty().EnableResize = false
}

// EnableMinimize 启用最小化按钮
func (m *ViewsFrameworkBrowserWindow) EnableMinimize() {
	m.WindowProperty().EnableMinimize = true
}

// EnableMaximize 启用最大化按钮
func (m *ViewsFrameworkBrowserWindow) EnableMaximize() {
	m.WindowProperty().EnableMaximize = true
}

// EnableResize 启用允许调整窗口大小
func (m *ViewsFrameworkBrowserWindow) EnableResize() {
	m.WindowProperty().EnableResize = true
}

// Component 返回窗口父组件
func (m *ViewsFrameworkBrowserWindow) Component() lcl.IComponent {
	return m.component
}

// WindowComponent 返回窗口组件
func (m *ViewsFrameworkBrowserWindow) WindowComponent() *TCEFWindowComponent {
	return m.windowComponent
}

// BrowserViewComponent 返回浏览器显示组件
func (m *ViewsFrameworkBrowserWindow) BrowserViewComponent() *TCEFBrowserViewComponent {
	return m.browserViewComponent
}
