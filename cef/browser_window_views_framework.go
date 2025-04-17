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
	"github.com/cyber-xxm/energy/v2/cef/internal/assets"
	"github.com/cyber-xxm/energy/v2/cef/internal/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"os"
)

// ViewsFrameworkBrowserWindow 基于CEF views framework 窗口组件
//
// 该窗口使用CEF内部实现，在linux下107.xx以后版本默认使用GTK3，但无法使用lcl组件集成到窗口中
//
// 当创建应用配置时 MultiThreadedMessageLoop 和 ExternalMessagePump 属性同时为false(linux系统默认强制false)时启用ViewsFramework窗口
type ViewsFrameworkBrowserWindow struct {
	isClosing             bool                      //
	windowType            consts.WINDOW_TYPE        //窗口类型
	windowId              int32                     //
	chromiumBrowser       ICEFChromiumBrowser       //
	component             lcl.IComponent            //
	windowComponent       *TCEFWindowComponent      //
	browserViewComponent  *TCEFBrowserViewComponent //
	windowProperty        *WindowProperty           //窗口属性
	auxTools              *auxTools                 //辅助工具
	tray                  ITray                     //托盘
	doOnWindowCreated     windowOnWindowCreated     //窗口创建
	doOnGetInitialBounds  windowOnGetInitialBounds  //窗口初始bounds
	doOnCloseQuery        windowOnCanCloseEx        //
	context               *ICefRequestContext       //
	extraInfo             *ICefDictionaryValue      //
	screen                IScreen                   //屏幕
	created               bool                      //创建顶层窗口完成
	canEnableDefaultEvent bool                      //是否启用了默认事件
}

// NewViewsFrameworkBrowserWindow 创建 ViewsFrameworkBrowserWindow 窗口
//
// config: Chromium配置, 提供快捷chromium配置
// windowProperty: 窗口属性
// owner: 被创建组件拥有者
func NewViewsFrameworkBrowserWindow(config *TCefChromiumConfig, windowProperty WindowProperty, owner lcl.IComponent) *ViewsFrameworkBrowserWindow {
	if config == nil {
		config = NewChromiumConfig()
	}
	var component = lcl.NewComponent(owner)
	chromiumBrowser := &TCEFChromiumBrowser{
		chromium: NewChromium(component, config),
	}
	m := &ViewsFrameworkBrowserWindow{
		windowProperty:       &windowProperty,
		component:            component,
		chromiumBrowser:      chromiumBrowser,
		windowComponent:      WindowComponentRef.New(component),
		browserViewComponent: BrowserViewComponentRef.New(component),
	}
	chromiumBrowser.SetSelfWindow(m)
	m.SetWindowType(windowProperty.WindowType)
	m.Chromium().SetEnableMultiBrowserMode(true)
	m.WindowComponent().SetOnWindowCreated(func(window *ICefWindow) {
		if m.Chromium().CreateBrowserByBrowserViewComponent(windowProperty.Url, m.BrowserViewComponent(), m.context, m.extraInfo) {
			m.WindowComponent().AddChildView(m.BrowserViewComponent().BrowserView().AsView())
			if windowProperty.Title != "" {
				m.WindowComponent().SetTitle(windowProperty.Title)
			}
			if windowProperty.EnableCenterWindow {
				m.WindowComponent().CenterWindow(NewCefSize(windowProperty.Width, windowProperty.Height))
			}
			if windowProperty.IconFS != "" {
				if err := m.WindowComponent().SetWindowAppIconByFSFile(1, windowProperty.IconFS); err != nil {
					logger.Error("set window application icon error:", err.Error())
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
				m.doOnWindowCreated(window)
			}
		}
	})
	return m
}

// ViewsFrameworkBrowserWindow 主窗口初始化
func appContextInitialized() {
	// 仅主进程初始化主窗口,
	// 子进程也不会初始， 判断一下省着多调用函数了
	if !process.Args.IsMain() {
		return
	}
	var m = BrowserWindow
	var bwEvent = m.browserEvent
	// VF 主窗口在 application 上下文初始化时创建
	application.SetOnContextInitialized(func() {
		// 主窗口
		m.Config.WindowProperty.WindowType = consts.WT_MAIN_BROWSER
		vfMainWindow := NewViewsFrameworkBrowserWindow(m.Config.ChromiumConfig(), m.Config.WindowProperty, nil)

		// 主窗口关闭流程 before close
		// 如果OnCanClose阻止关闭，该函数不会执行
		vfMainWindow.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
			var flag = false
			if bwEvent.onBeforeClose != nil {
				flag = bwEvent.onBeforeClose(sender, browser, vfMainWindow)
			}
			if !flag {
				chromiumOnBeforeClose(vfMainWindow, browser)
				vfMainWindow.TryCloseWindowAndTerminate()
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
		vfMainWindow.WindowComponent().SetOnCanClose(func(window *ICefWindow, canClose *bool) {
			var flag bool
			if vfMainWindow.doOnCloseQuery != nil {
				flag = vfMainWindow.doOnCloseQuery(window, vfMainWindow, canClose)
			}
			if !flag {
				*canClose = m.Config.WindowProperty.EnableClose
				if m.Config.WindowProperty.EnableClose {
					//*aResult = vfMainWindow.Chromium().TryCloseBrowser()
					vfMainWindow.CloseBrowserWindow()
				}
			}
		})
		// 设置到 MainBrowser, 主窗口有且仅有一个
		BrowserWindow.mainBrowserWindow = vfMainWindow
		if m.Config.browserWindowOnEventCallback != nil {
			BrowserWindow.browserEvent.chromium = vfMainWindow.Chromium()
			m.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, vfMainWindow)
		}
		// IPC
		ipc.SetProcessMessage(vfMainWindow)
		vfMainWindow.CreateTopLevelWindow()
		//创建完窗口之后设置窗口属性
		vfMainWindow.createAfterWindowPropertyForEvent()
	})
}

// TryCloseWindowAndTerminate
// 尝试关闭窗口并退出应用,
// EnableMainWindow
//
//	如果禁用主窗口, 存在多窗口时只在最后一个窗口关闭时才退出整个应用进程
//	如果启用主窗口, 关闭主窗口时退出整个应用进程
func (m *ViewsFrameworkBrowserWindow) TryCloseWindowAndTerminate() {
	var closeWindowAndTerminate = func() {
		if m.tray != nil {
			m.tray.close()
		}
		ui := api.WidgetUI()
		application.QuitMessageLoop()
		// TODO 当前使用 os.Exit(0) 正确退出应用程序
		if ui.IsGTK3() {
			os.Exit(0)
		}
	}
	// 启用主窗口，当前关闭窗口为主浏览器窗口直接退出进程
	if BrowserWindow.Config.EnableMainWindow && m.WindowType() == consts.WT_MAIN_BROWSER {
		closeWindowAndTerminate()
	} else {
		// 禁用主窗口，无窗口列表时退出进程
		count := len(BrowserWindow.GetWindowInfos())
		if count < 1 {
			closeWindowAndTerminate()
		}
	}
}

// Target
//
//	IPC消息接收目标, 当前窗口chromium发送
//	参数: targetType 可选, 接收类型
func (m *ViewsFrameworkBrowserWindow) Target() target.ITarget {
	browse := m.Chromium().Browser()
	if !browse.IsValid() {
		return nil
	}
	return target.NewTarget(m, browse.Identifier(), browse.MainFrame().Identifier())
}

// ProcessMessage
//
//	IPC消息触发当前Chromium
func (m *ViewsFrameworkBrowserWindow) ProcessMessage() target.IProcessMessage {
	if m.chromiumBrowser == nil {
		return nil
	}
	return m.chromiumBrowser.Chromium().(*TCEFChromium)
}

func (m *ViewsFrameworkBrowserWindow) AsTargetWindow() target.IWindow {
	return m
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

// WV 窗口的弹出事件
func (m *ViewsFrameworkBrowserWindow) doBeforePopup(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, popupFeatures *TCefPopupFeatures, windowInfo *TCefWindowInfo, client *ICefClient, settings *TCefBrowserSettings, resultExtraInfo *ICefDictionaryValue, noJavascriptAccess *bool) bool {
	var bwEvent = BrowserWindow.browserEvent

	wp := *m.windowProperty //clone
	wp.Url = beforePopupInfo.TargetUrl
	wp.WindowType = consts.WT_POPUP_SUB_BROWSER
	var vFrameBrowserWindow = NewViewsFrameworkBrowserWindow(NewChromiumConfig(), wp, BrowserWindow.MainWindow().AsViewsFrameworkBrowserWindow().Component())
	var result = false
	if bwEvent.onBeforePopup != nil {
		result = bwEvent.onBeforePopup(sender, vFrameBrowserWindow, browser, frame, beforePopupInfo, popupFeatures, windowInfo, client, settings, resultExtraInfo, noJavascriptAccess)
	}
	if !result {
		vFrameBrowserWindow.ResetWindowPropertyForEvent()
		vFrameBrowserWindow.EnableAllDefaultEvent()
		vFrameBrowserWindow.CreateTopLevelWindow()
		vFrameBrowserWindow.createAfterWindowPropertyForEvent()
		result = true
	}
	return result
}

// ResetWindowPropertyForEvent 重置窗口属性-通过事件函数
//
//	VF窗口初始化时通过回调事件设置一些默认行为，而不像LCL窗口直接通过属性设置
//	在初始化之后部分属性可直接设置
func (m *ViewsFrameworkBrowserWindow) ResetWindowPropertyForEvent() {
	wp := m.WindowProperty()
	m.WindowComponent().SetOnGetInitialShowState(func(window *ICefWindow, aResult *consts.TCefShowState) {
		*aResult = consts.TCefShowState(wp.WindowInitState + 1) // CEF 要 + 1
	})
	m.WindowComponent().SetOnGetInitialBounds(func(window *ICefWindow, aResult *TCefRect) {
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
			m.doOnGetInitialBounds(window, aResult)
		}
	})
	m.WindowComponent().SetOnCanMinimize(func(window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableMinimize
	})
	m.WindowComponent().SetOnCanResize(func(window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableResize
	})
	m.WindowComponent().SetOnCanMaximize(func(window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableMaximize
	})
	m.WindowComponent().SetOnCanClose(func(window *ICefWindow, canClose *bool) {
		var flag bool
		if m.doOnCloseQuery != nil {
			flag = m.doOnCloseQuery(window, m, canClose)
		}
		if !flag {
			*canClose = wp.EnableClose
			if wp.EnableClose {
				//*canClose = m.Chromium().TryCloseBrowser()
				m.CloseBrowserWindow()
			}
		}
	})
	m.WindowComponent().SetOnIsFrameless(func(window *ICefWindow, aResult *bool) {
		*aResult = wp.EnableHideCaption
	})
	m.WindowComponent().SetAlwaysOnTop(wp.AlwaysOnTop)
	m.WindowComponent().SetBounds(NewCefRect(wp.X, wp.Y, wp.Width, wp.Height))
}

// Created 窗口是否创建完, m.CreateTopLevelWindow() 之后
func (m *ViewsFrameworkBrowserWindow) Created() bool {
	return m.created
}

// EnableAllDefaultEvent 启用所有默认事件行为
func (m *ViewsFrameworkBrowserWindow) EnableAllDefaultEvent() {
	m.chromiumBrowser.RegisterDefaultPopupEvent()
	m.chromiumBrowser.RegisterDefaultEvent()
	m.canEnableDefaultEvent = true
}

// SetOnWindowCreated 窗口创建
func (m *ViewsFrameworkBrowserWindow) SetOnWindowCreated(onWindowCreated windowOnWindowCreated) {
	m.doOnWindowCreated = onWindowCreated
}

// SetOnGetInitialBounds 窗口初始坐标和大小
func (m *ViewsFrameworkBrowserWindow) SetOnGetInitialBounds(onGetInitialBounds windowOnGetInitialBounds) {
	m.doOnGetInitialBounds = onGetInitialBounds
}

// SetCreateBrowserExtraInfo
//
//	设置 Chromium 创建浏览器时设置的扩展信息
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
	m.SetSize(value, size.Height)
}

// SetHeight 设置窗口高
func (m *ViewsFrameworkBrowserWindow) SetHeight(value int32) {
	m.WindowProperty().Height = value
	size := m.Size()
	m.SetSize(size.Width, value)
}

// Point 返回窗口坐标
func (m *ViewsFrameworkBrowserWindow) Point() TCefPoint {
	result := m.WindowComponent().Position()
	m.WindowProperty().X = result.X
	m.WindowProperty().Y = result.Y
	return result
}

// Size 返回窗口宽高
func (m *ViewsFrameworkBrowserWindow) Size() TCefSize {
	result := m.WindowComponent().Size()
	m.WindowProperty().Width = result.Width
	m.WindowProperty().Height = result.Height

	return result
}

// Bounds 返回窗口坐标和宽高
func (m *ViewsFrameworkBrowserWindow) Bounds() TCefRect {
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
	m.WindowComponent().SetPosition(TCefPoint{X: x, Y: y})
}

// SetSize 设置窗口宽高
func (m *ViewsFrameworkBrowserWindow) SetSize(width, height int32) {
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
	m.WindowComponent().SetSize(TCefSize{Width: width, Height: height})
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
	return m.chromiumBrowser.Chromium()
}

func (m *ViewsFrameworkBrowserWindow) ChromiumBrowser() ICEFChromiumBrowser {
	return m.chromiumBrowser
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
	m.WindowComponent().CreateTopLevelWindow()
	// 标记已创建
	m.created = true
}

// CenterWindow 设置窗口居中，同时指定窗口大小
func (m *ViewsFrameworkBrowserWindow) CenterWindow(size TCefSize) {
	m.WindowComponent().CenterWindow(size)
}

// SetCenterWindow 设置窗口居中显示
func (m *ViewsFrameworkBrowserWindow) SetCenterWindow(value bool) {
	m.WindowProperty().EnableCenterWindow = value
	if value {
		m.CenterWindow(NewCefSize(m.WindowProperty().Width, m.WindowProperty().Height))
	}
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
//
//	在主线程中运行
func (m *ViewsFrameworkBrowserWindow) RunOnMainThread(fn func()) {
	RunOnMainThread(fn)
}

func (m *ViewsFrameworkBrowserWindow) SetOnCloseQuery(fn windowOnCanCloseEx) {
	m.doOnCloseQuery = fn
}
