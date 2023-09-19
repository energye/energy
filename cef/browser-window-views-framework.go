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
	"github.com/energye/energy/v2/cef/internal/assets"
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/internal/window"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"runtime"
)

// ViewsFrameworkBrowserWindow 基于CEF views framework 窗口组件
//
// 该窗口使用CEF内部实现，在linux下107.xx以后版本默认使用GTK3，但无法使用lcl组件集成到窗口中
//
// 当创建应用配置时 MultiThreadedMessageLoop 和 ExternalMessagePump 属性同时为false(linux系统默认强制false)时启用ViewsFramework窗口
type ViewsFrameworkBrowserWindow struct {
	isClosing             bool                              //
	windowType            consts.WINDOW_TYPE                //0:browser 1:devTools 2:viewSource 默认:0
	windowId              int32                             //
	chromium              IChromium                         //
	component             lcl.IComponent                    //
	windowComponent       *TCEFWindowComponent              //
	browserViewComponent  *TCEFBrowserViewComponent         //
	windowProperty        *WindowProperty                   //窗口属性
	auxTools              *auxTools                         //辅助工具
	tray                  ITray                             //托盘
	doOnWindowCreated     WindowComponentOnWindowCreated    //窗口创建
	doOnGetInitialBounds  WindowComponentOnGetInitialBounds //窗口初始bounds
	regions               *TCefDraggableRegions             //窗口内html拖拽区域
	context               *ICefRequestContext               //
	extraInfo             *ICefDictionaryValue              //
	screen                IScreen                           //屏幕
	created               bool                              //创建顶层窗口完成
	canEnableDefaultEvent bool                              //是否启用了默认事件
}

// NewViewsFrameworkBrowserWindow 创建 ViewsFrameworkBrowserWindow 窗口
func NewViewsFrameworkBrowserWindow(config *TCefChromiumConfig, windowProperty WindowProperty, owner ...lcl.IComponent) *ViewsFrameworkBrowserWindow {
	if config == nil {
		config = NewChromiumConfig()
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
	m.Chromium().SetEnableMultiBrowserMode(true)
	m.WindowComponent().SetOnWindowCreated(func(sender lcl.IObject, window *ICefWindow) {
		if m.Chromium().CreateBrowserByBrowserViewComponent(windowProperty.Url, m.BrowserViewComponent(), m.context, m.extraInfo) {
			m.WindowComponent().AddChildView(m.BrowserViewComponent())
			if windowProperty.Title != "" {
				m.WindowComponent().SetTitle(windowProperty.Title)
			}
			if windowProperty.EnableCenterWindow {
				m.WindowComponent().CenterWindow(NewCefSize(windowProperty.Width, windowProperty.Height))
			}
			if windowProperty.IconFS != "" {
				if emfs.IsExist(windowProperty.IconFS) {
					if err := m.WindowComponent().SetWindowAppIconByFSFile(1, windowProperty.IconFS); err != nil {
						logger.Error("set window application icon error:", err.Error())
					}
				}
			} else if windowProperty.Icon != "" {
				if tools.IsExist(windowProperty.Icon) {
					if err := m.WindowComponent().SetWindowAppIconByFile(1, windowProperty.Icon); err != nil {
						logger.Error("set window application icon error:", err.Error())
					}
				}
			} else {
				// 默认
				// vf png
				// lcl ico
				if iconData := assets.DefaultPNGICON(); iconData != nil {
					icon := ImageRef.New()
					icon.AddPng(1, assets.DefaultPNGICON())
					m.WindowComponent().SetWindowAppIcon(icon)
				}
			}
			m.BrowserViewComponent().RequestFocus()
			m.WindowComponent().Show()
			if m.doOnWindowCreated != nil {
				m.doOnWindowCreated(sender, window)
			}
		}
	})
	return m
}

// ViewsFrameworkBrowserWindow 主窗口初始化
func (m *browserWindow) appContextInitialized(app *TCEFApplication) {
	// 仅主进程初始化主窗口,
	// 子进程也不会初始， 判断一下省着多调用函数了
	if !process.Args.IsMain() {
		return
	}
	var bwEvent = BrowserWindow.browserEvent
	// VF 主窗口在 application 上下文初始化时创建
	app.SetOnContextInitialized(func() {
		// 主窗口
		m.Config.WindowProperty.WindowType = consts.WT_MAIN_BROWSER
		vfMainWindow := NewViewsFrameworkBrowserWindow(m.Config.ChromiumConfig(), m.Config.WindowProperty)
		// 主窗口关闭流程 before close
		// OnCanClose如果阻止关闭，该函数不会执行
		vfMainWindow.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
			var flag = false
			if bwEvent.onBeforeClose != nil {
				flag = bwEvent.onBeforeClose(sender, browser, vfMainWindow)
			}
			if !flag {
				chromiumOnBeforeClose(browser)
				if vfMainWindow.tray != nil {
					vfMainWindow.tray.close()
				}
				app.QuitMessageLoop()
			}
		})
		// SetOnClose如果阻止关闭，该函数不会执行
		vfMainWindow.Chromium().SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
			if bwEvent.onClose != nil {
				bwEvent.onClose(sender, browser, aAction, vfMainWindow)
			}
		})
		// 重置窗口属性, 使用事件初始窗口属性
		vfMainWindow.ResetWindowPropertyForEvent()
		vfMainWindow.EnableAllDefaultEvent() // 开启默认事件
		// 主窗口关闭时触发该函数
		// EnableClose=true时关闭窗口, false时不关闭窗口
		vfMainWindow.WindowComponent().SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
			*aResult = m.Config.WindowProperty.EnableClose
			if m.Config.WindowProperty.EnableClose {
				*aResult = vfMainWindow.Chromium().TryCloseBrowser()
			}
		})
		// 设置到 MainBrowser, 主窗口有且仅有一个
		BrowserWindow.mainVFBrowserWindow = vfMainWindow
		if m.Config.browserWindowOnEventCallback != nil {
			BrowserWindow.browserEvent.chromium = vfMainWindow.chromium
			m.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, vfMainWindow)
		}
		// IPC
		ipc.SetProcessMessage(vfMainWindow.Chromium().(*TCEFChromium))
		vfMainWindow.CreateTopLevelWindow()
		//创建完窗口之后设置窗口属性
		vfMainWindow.createAfterWindowPropertyForEvent()
	})
}

func (m *ViewsFrameworkBrowserWindow) createAfterWindowPropertyForEvent() {
	wp := m.WindowProperty()
	if wp.EnableResize {
		// VF MinimumSize & MaximumSize 在事件中设置
		// 如果动态设置，需要自己实现该回调函数
		if wp.MinWidth > 0 && wp.MinHeight > 0 {
			m.WindowComponent().SetOnGetMinimumSize(func(view *ICefView, result *TCefSize) {
				result.Width = int32(wp.MinWidth)
				result.Height = int32(wp.MinHeight)
			})
		}
		if wp.MaxWidth > 0 && wp.MaxHeight > 0 {
			m.WindowComponent().SetOnGetMaximumSize(func(view *ICefView, result *TCefSize) {
				result.Width = int32(wp.MaxWidth)
				result.Height = int32(wp.MaxHeight)
			})
		}
	}
}

// registerPopupEvent 注册弹出子窗口事件
func (m *ViewsFrameworkBrowserWindow) registerPopupEvent(isMain bool) {
	var bwEvent = BrowserWindow.browserEvent
	if !isMain {
		// 子窗口关闭流程
		m.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
			var flag bool
			if bwEvent.onBeforeClose != nil {
				flag = bwEvent.onBeforeClose(sender, browser, m)
			}
			if !flag {
				chromiumOnBeforeClose(browser)
			}
		})
		m.Chromium().SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
			if bwEvent.onClose != nil {
				bwEvent.onClose(sender, browser, aAction, m)
			}
		})
	}
	m.Chromium().SetOnOpenUrlFromTab(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool {
		if !m.Chromium().Config().EnableOpenUrlTab() {
			return true
		}
		return false
	})
	m.Chromium().SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
		if !m.Chromium().Config().EnableWindowPopup() {
			return true
		}
		wp := *m.windowProperty //clone
		wp.Url = beforePopupInfo.TargetUrl
		wp.WindowType = consts.WT_POPUP_SUB_BROWSER
		var vFrameBrowserWindow = NewViewsFrameworkBrowserWindow(NewChromiumConfig(), wp, BrowserWindow.MainWindow().AsViewsFrameworkBrowserWindow().Component())
		var result = false
		if bwEvent.onBeforePopup != nil {
			result = bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, vFrameBrowserWindow, noJavascriptAccess)
		}
		if !result {
			vFrameBrowserWindow.ResetWindowPropertyForEvent()
			vFrameBrowserWindow.EnableAllDefaultEvent()
			vFrameBrowserWindow.CreateTopLevelWindow()
			vFrameBrowserWindow.createAfterWindowPropertyForEvent()
			result = true
		}
		return result
	})
}

// ResetWindowPropertyForEvent 重置窗口属性-通过事件函数
//  VF窗口初始化时通过回调事件设置一些默认行为，而不像LCL窗口直接通过属性设置
//  在初始化之后部分属性可直接设置
func (m *ViewsFrameworkBrowserWindow) ResetWindowPropertyForEvent() {
	wp := m.WindowProperty()
	m.WindowComponent().SetOnGetInitialShowState(func(sender lcl.IObject, window *ICefWindow, aResult *consts.TCefShowState) {
		*aResult = consts.TCefShowState(wp.WindowInitState + 1) // CEF 要 + 1
	})
	m.WindowComponent().SetOnGetInitialBounds(func(sender lcl.IObject, window *ICefWindow, aResult *TCefRect) {
		if wp.EnableCenterWindow {
			m.WindowComponent().CenterWindow(NewCefSize(wp.Width, wp.Height))
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
	m.WindowComponent().SetOnCanMinimize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableMinimize
	})
	m.WindowComponent().SetOnCanResize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableResize
	})
	m.WindowComponent().SetOnCanMaximize(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableMaximize
	})
	m.WindowComponent().SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableClose
		if wp.EnableClose {
			*aResult = m.Chromium().TryCloseBrowser()
		}
	})
	m.WindowComponent().SetOnIsFrameless(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableHideCaption
	})
	m.WindowComponent().SetAlwaysOnTop(wp.AlwaysOnTop)
	m.WindowComponent().SetBounds(NewCefRect(wp.X, wp.Y, wp.Width, wp.Height))
}

// registerDefaultEvent 注册默认事件
func (m *ViewsFrameworkBrowserWindow) registerDefaultEvent() {
	var bwEvent = BrowserWindow.browserEvent
	//默认自定义快捷键
	defaultAcceleratorCustom()
	m.Chromium().SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
		if bwEvent.onProcessMessageReceived != nil {
			return bwEvent.onProcessMessageReceived(sender, browser, frame, sourceProcess, message, m)
		}
		return false
	})
	m.Chromium().SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue) {
		if assetserve.AssetsServerHeaderKeyValue != "" {
			request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
		}
		if bwEvent.onBeforeResourceLoad != nil {
			bwEvent.onBeforeResourceLoad(sender, browser, frame, request, callback, result, m)
		}
	})
	m.Chromium().SetOnBeforeDownload(func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *ICefDownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) {
		if bwEvent.onBeforeDownload != nil {
			bwEvent.onBeforeDownload(sender, browser, beforeDownloadItem, suggestedName, callback, m)
		} else {
			callback.Cont(consts.ExePath+consts.Separator+suggestedName, true)
		}
	})
	m.Chromium().SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
		var flag bool
		if bwEvent.onBeforeContextMenu != nil {
			flag = bwEvent.onBeforeContextMenu(sender, browser, frame, params, model, m)
		}
		if !flag {
			chromiumOnBeforeContextMenu(m, browser, frame, params, model)
		}
	})
	m.Chromium().SetOnContextMenuCommand(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32) bool {
		var result bool
		if bwEvent.onContextMenuCommand != nil {
			result = bwEvent.onContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, m)
		}
		if !result {
			result = chromiumOnContextMenuCommand(m, browser, frame, params, commandId, eventFlags)
		}
		return result
	})
	m.Chromium().SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {
		var flag bool
		if bwEvent.onAfterCreated != nil {
			flag = bwEvent.onAfterCreated(sender, browser, m)
		}
		if !flag {
			chromiumOnAfterCreate(m, browser)
		}
	})
	m.Chromium().SetOnKeyEvent(func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, osEvent *consts.TCefEventHandle, result *bool) {
		if bwEvent.onKeyEvent != nil {
			bwEvent.onKeyEvent(sender, browser, event, osEvent, m, result)
		}
		if !*result {
			if m.WindowType() == consts.WT_DEV_TOOLS || m.WindowType() == consts.WT_VIEW_SOURCE {
				return
			}
			if m.Chromium().Config().EnableDevTools() {
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
	m.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool) bool {
		chromiumOnBeforeBrowser(m, browser, frame, request) // default impl
		if bwEvent.onBeforeBrowser != nil {
			return bwEvent.onBeforeBrowser(sender, browser, frame, request, userGesture, isRedirect, m)
		}
		return false
	})
	m.Chromium().SetOnTitleChange(func(sender lcl.IObject, browser *ICefBrowser, title string) {
		if bwEvent.onTitleChange != nil {
			bwEvent.onTitleChange(sender, browser, title, m)
		}
	})
	m.Chromium().SetOnDragEnter(func(sender lcl.IObject, browser *ICefBrowser, dragData *ICefDragData, mask consts.TCefDragOperations, result *bool) {
		*result = !m.WindowProperty().EnableDragFile
		if bwEvent.onDragEnter != nil {
			bwEvent.onDragEnter(sender, browser, dragData, mask, m, result)
		}
	})
	m.Chromium().SetOnLoadEnd(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32) {
		if bwEvent.onLoadEnd != nil {
			bwEvent.onLoadEnd(sender, browser, frame, httpStatusCode, m)
		}
	})
	if m.WindowProperty().EnableWebkitAppRegion {
		m.Chromium().SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, regions *TCefDraggableRegions) {
			if bwEvent.onDraggableRegionsChanged != nil {
				bwEvent.onDraggableRegionsChanged(sender, browser, frame, regions, m)
			}
			m.regions = regions
			m.WindowComponent().SetDraggableRegions(regions.Regions())
		})
	}
	if localLoadRes.enable() {
		m.Chromium().SetOnGetResourceHandler(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (resourceHandler *ICefResourceHandler) {
			//var flag bool
			if bwEvent.onGetResourceHandler != nil {
				resourceHandler, _ = bwEvent.onGetResourceHandler(sender, browser, frame, request, m)
			}
			//if !flag {
			//	resourceHandler = localLoadRes.getResourceHandler(browser, frame, request)
			//}
			return
		})
	}
}

// Created 窗口是否创建完, m.CreateTopLevelWindow() 之后
func (m *ViewsFrameworkBrowserWindow) Created() bool {
	return m.created
}

// EnableAllDefaultEvent 启用所有默认事件行为
func (m *ViewsFrameworkBrowserWindow) EnableAllDefaultEvent() {
	m.registerPopupEvent(m.WindowType() == consts.WT_MAIN_BROWSER)
	m.registerDefaultEvent()
	m.canEnableDefaultEvent = true
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
	m.WindowComponent().SetTitle(title)
}

// SetWidth 设置窗口宽
func (m *ViewsFrameworkBrowserWindow) SetWidth(value int32) {
	m.WindowProperty().Width = value
	size := m.Size()
	if size != nil {
		m.SetSize(value, size.Height)
	}
}

// SetHeight 设置窗口高
func (m *ViewsFrameworkBrowserWindow) SetHeight(value int32) {
	m.WindowProperty().Height = value
	size := m.Size()
	if size != nil {
		m.SetSize(size.Width, value)
	}
}

// Point 返回窗口坐标
func (m *ViewsFrameworkBrowserWindow) Point() *TCefPoint {
	result := m.WindowComponent().Position()
	if result != nil {
		m.WindowProperty().X = result.X
		m.WindowProperty().Y = result.Y
	}
	return result
}

// Size 返回窗口宽高
func (m *ViewsFrameworkBrowserWindow) Size() *TCefSize {
	result := m.WindowComponent().Size()
	if result != nil {
		m.WindowProperty().Width = result.Width
		m.WindowProperty().Height = result.Height
	}
	return result
}

// Bounds 返回窗口坐标和宽高
func (m *ViewsFrameworkBrowserWindow) Bounds() *TCefRect {
	result := m.WindowComponent().Bounds()
	if result != nil {
		m.WindowProperty().X = result.X
		m.WindowProperty().Y = result.Y
		m.WindowProperty().Width = result.Width
		m.WindowProperty().Height = result.Height
	}
	return result
}

// SetPoint 设置窗口坐标
func (m *ViewsFrameworkBrowserWindow) SetPoint(x, y int32) {
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
	m.WindowComponent().SetPosition(&TCefPoint{X: x, Y: y})
}

// SetSize 设置窗口宽高
func (m *ViewsFrameworkBrowserWindow) SetSize(width, height int32) {
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
	m.WindowComponent().SetSize(&TCefSize{Width: width, Height: height})
}

// SetBounds 设置窗口坐标和宽高
func (m *ViewsFrameworkBrowserWindow) SetBounds(x, y, width, height int32) {
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
	m.SetPoint(x, y)
	m.SetSize(width, height)
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
	if m.Created() {
		m.BrowserViewComponent().RequestFocus()
		m.WindowComponent().Show()
	} else {
		if m.canEnableDefaultEvent {
			// 启用了默认事件，窗口属性配置事件在创建顶层窗口之前调用
			m.ResetWindowPropertyForEvent()
		}
		m.CreateTopLevelWindow()
		m.createAfterWindowPropertyForEvent()
	}
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
	if m.WindowComponent().IsMinimized() {
		return types.WsMinimized
	} else if m.WindowComponent().IsMaximized() {
		return types.WsMaximized
	} else if m.WindowComponent().IsFullscreen() {
		return types.WsFullScreen
	}
	return types.WsNormal
}

// Maximize 窗口最大化/还原
func (m *ViewsFrameworkBrowserWindow) Maximize() {
	if m.WindowState() == types.WsNormal {
		m.WindowComponent().Maximize()
	} else {
		m.Restore()
	}
}

func (m *ViewsFrameworkBrowserWindow) FullScreen() {
	m.WindowComponent().SetFullscreen(true)
}

func (m *ViewsFrameworkBrowserWindow) ExitFullScreen() {
	m.WindowComponent().SetFullscreen(false)
}

func (m *ViewsFrameworkBrowserWindow) IsFullScreen() bool {
	return m.WindowComponent().IsFullscreen()
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
	m.Chromium().CloseBrowser(true)
}

// CreateTopLevelWindow 创建顶层窗口
func (m *ViewsFrameworkBrowserWindow) CreateTopLevelWindow() {
	// 非辅助工具类型窗口，不做托管, 辅助工具窗口有自己的事件行为
	if m.WindowType() != consts.WT_DEV_TOOLS {
		window.CurrentBrowseWindowCache = m
	}
	m.WindowComponent().CreateTopLevelWindow()
	// 标记已创建
	m.created = true
}

// CenterWindow 设置窗口居中，同时指定窗口大小
func (m *ViewsFrameworkBrowserWindow) CenterWindow(size *TCefSize) {
	m.WindowComponent().CenterWindow(size)
}

// SetCenterWindow 设置窗口居中显示
func (m *ViewsFrameworkBrowserWindow) SetCenterWindow(value bool) {
	m.WindowProperty().EnableCenterWindow = value
	m.CenterWindow(m.Size())
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

func (m *ViewsFrameworkBrowserWindow) Screen() IScreen {
	if m.screen == nil && m.BrowserViewComponent() != nil {
		m.screen = &Screen{window: m}
	}
	return m.screen
}

// RunOnMainThread
//	在主线程中运行
func (m *ViewsFrameworkBrowserWindow) RunOnMainThread(fn func()) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if api.DMainThreadId() == api.DCurrentThreadId() {
		fn()
	} else {
		lcl.ThreadSync(func() {
			fn()
		})
	}
}
