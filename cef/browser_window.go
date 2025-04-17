//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// LCL + Energy 窗口属性
// 提供了常用属性配置, 如果使用更复杂的配置需要直接使用LCL提供的窗口属性配置

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	et "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"runtime"
)

const (
	defaultTitle      = "ENERGY"
	defaultAboutBlank = "about:blank"
)

// auxTools 辅助工具
type auxTools struct {
	devToolsWindow   *devToolsWindow //devTools
	viewSourceWindow IBrowserWindow  //viewSource
}

// 窗口当前状态属性，仅触发全屏使用
type windowCurrentProperty struct {
	windowState             types.TWindowState
	previousWindowPlacement types.TRect
}

// WindowProperty
//
//	提供部分窗口属性配置，初始化时生效
//	如需更多属性配置或自定义窗口行为请在`SetBrowserInit`回调函数中使用
type WindowProperty struct {
	IsShowModel     bool               // 是否以模态窗口显示
	WindowInitState types.TWindowState // 窗口 初始状态: 最小化、最大化、全屏, 全屏时隐藏标题栏生效
	WindowType      consts.WINDOW_TYPE // 窗口 类型 WINDOW_TYPE default: WT_MAIN_BROWSER
	Title           string             // 窗口 标题
	//Url  默认打开URL, 支持http和LocalLoad(本地资源)加载方式
	//  web服务方式: http's://www.example.com, LocalLoad方式: fs://energy/index.html
	//  LocalLoad: 不需要web服务支持, 如果浏览器调用数据接口需要配置代理转发
	//  LocalLoad: 通过 Config.LocalResource 配置实现
	//  LocalLoad: 地址必须与配置的自定义协议和域相同, 格式 [scheme]://[custom domain]
	Url                       string
	Icon                      string                // 窗口图标 加载本地图标 local > /app/resources/icon.ico, VF窗口linux使用png
	IconFS                    string                // 窗口图标 加载emfs内置图标 emfs > resources/icon.ico, VF窗口linux使用png
	EnableWebkitAppRegionDClk bool                  //
	EnableHideCaption         bool                  // 窗口 是否隐藏标题栏, VF窗口组件Linux下不能动态控制
	EnableMinimize            bool                  // 窗口 是否启用最小化 default: true
	EnableMaximize            bool                  // 窗口 是否启用最大化 default: true
	EnableResize              bool                  // 窗口 是否允许调整大小 default: true
	EnableClose               bool                  // 窗口 关闭时是否关闭窗口 default: true
	EnableCenterWindow        bool                  // 窗口 居中显示 default: true
	EnableDragFile            bool                  // 窗口 是否允许向窗口内拖拽文件
	EnableMainWindow          bool                  // 窗口 是否启用主窗口 default: true, 值为false时不再有主子窗口区分
	AlwaysOnTop               bool                  // 窗口 窗口置顶
	ShowInTaskBar             types.TShowInTaskbar  // 窗口 是否显示在任务栏, 仅适于用自定义窗口, 默认: 始终显示在任务栏
	X                         int32                 // 窗口 EnableCenterWindow=false X坐标 default: 100
	Y                         int32                 // 窗口 EnableCenterWindow=false Y坐标 default: 100
	Width                     int32                 // 窗口 宽 default: 1024
	Height                    int32                 // 窗口 高 default: 768
	MinWidth                  types.TConstraintSize // 窗口 最小宽, EnableResize = true 与 MinHeight > 0 生效
	MinHeight                 types.TConstraintSize // 窗口 最小高, EnableResize = true 与 MinWidth > 0 生效
	MaxWidth                  types.TConstraintSize // 窗口 最大宽, EnableResize = true 与 MaxHeight > 0 生效
	MaxHeight                 types.TConstraintSize // 窗口 最大高, EnableResize = true 与 MaxWidth > 0 生效
	current                   windowCurrentProperty // 窗口 当前属性
}

// IBrowserWindow
//
//	浏览器窗口基础接口
//	定义了常用函数, 更多窗口功能或属性在SetBrowserInit函数中使用
type IBrowserWindow interface {
	Id() int32                                                                                                //窗口ID
	Handle() types.HWND                                                                                       //窗口句柄
	Show()                                                                                                    //显示窗口
	Hide()                                                                                                    //隐藏窗口
	WindowState() types.TWindowState                                                                          //返回窗口最小化、最大化、全屏状态
	Maximize()                                                                                                //窗口最大化
	Minimize()                                                                                                //窗口最小化
	Restore()                                                                                                 //窗口还原
	FullScreen()                                                                                              //全屏模式, 仅隐藏标题栏时有效
	ExitFullScreen()                                                                                          //退出全屏模式
	IsFullScreen() bool                                                                                       //是否全屏模式
	Close()                                                                                                   //关闭窗口 非browser窗口使用
	CloseBrowserWindow()                                                                                      //关闭浏览器窗口 带有browser窗口使用
	WindowType() consts.WINDOW_TYPE                                                                           //窗口类型
	SetWindowType(windowType consts.WINDOW_TYPE)                                                              //设置窗口类型
	Browser() *ICefBrowser                                                                                    //窗口内的Browser对象
	Chromium() IChromium                                                                                      //窗口内的Chromium对象, 返回空时我们可能需要自己创建, ChromiumCreate
	ChromiumBrowser() ICEFChromiumBrowser                                                                     //ChromiumBrowser包装对象
	DisableMaximize()                                                                                         //禁用最大化
	DisableMinimize()                                                                                         //禁用最小化
	DisableResize()                                                                                           //禁用窗口大小调整
	EnableMaximize()                                                                                          //启用最大化
	EnableMinimize()                                                                                          //启用最小化
	EnableResize()                                                                                            //启用允许调整窗口大小
	IsClosing() bool                                                                                          //返回窗口是否正在关闭/或已关闭 true正在或已关闭
	AsViewsFrameworkBrowserWindow() IViewsFrameworkBrowserWindow                                              //转换为ViewsFramework窗口接口, 失败返回nil
	AsLCLBrowserWindow() ILCLBrowserWindow                                                                    //转换为LCL窗口接口, 失败返回nil
	EnableAllDefaultEvent()                                                                                   //启用所有默认事件
	SetTitle(title string)                                                                                    //设置窗口标题栏标题
	IsViewsFramework() bool                                                                                   //是否为 IViewsFrameworkBrowserWindow 窗口，失败返回false
	IsLCL() bool                                                                                              //是否为 ILCLBrowserWindow 窗口，失败返回false
	WindowProperty() *WindowProperty                                                                          //窗口常用属性
	SetWidth(value int32)                                                                                     //设置窗口宽
	SetHeight(value int32)                                                                                    //设置窗口高
	Point() TCefPoint                                                                                         //窗口坐标
	Size() TCefSize                                                                                           //窗口宽高
	Bounds() TCefRect                                                                                         //窗口坐标和宽高
	SetPoint(x, y int32)                                                                                      //设置窗口坐标
	SetSize(width, height int32)                                                                              //设置窗口宽高
	SetBounds(x, y, width, height int32)                                                                      //设置窗口坐标和宽高
	SetCenterWindow(value bool)                                                                               //设置窗口居中
	ShowTitle()                                                                                               //显示窗口标题栏
	HideTitle()                                                                                               //隐藏窗口标题栏
	SetDefaultInTaskBar()                                                                                     //默认窗口在任务栏上显示按钮
	SetShowInTaskBar()                                                                                        //强制窗口在任务栏上显示按钮
	SetNotInTaskBar()                                                                                         //强制窗口不在任务栏上显示按钮
	NewCefTray(width, height int32, url string) ITray                                                         //仅支持windows托盘LCL+[CEF|VF]（使用web端技术自定义实现,如使用LCL窗口组件该托盘实现是LCL+CEF,如使用VF窗口组件该托盘实现是LCL+VF）
	NewSysTray() ITray                                                                                        //systray系统原生
	SetCreateBrowserExtraInfo(windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) //设置 Chromium 创建浏览器时设置的扩展信息
	RunOnMainThread(fn func())                                                                                //在UI主线程中运行
	Screen() IScreen                                                                                          //返回屏幕信息
	Target() target.ITarget                                                                                   //IPC接收目标
	AsTargetWindow() target.IWindow                                                                           //IPC
	doBeforePopup(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo,
		popupFeatures *TCefPopupFeatures, windowInfo *TCefWindowInfo, client *ICefClient, settings *TCefBrowserSettings,
		resultExtraInfo *ICefDictionaryValue, noJavascriptAccess *bool) bool
}

// ILCLBrowserWindow
// 浏览器 LCL 窗口组件接口 继承 IBrowserWindow
//
// 定义了LCL常用函数
type ILCLBrowserWindow interface {
	IBrowserWindow
	BrowserWindow() *LCLBrowserWindow                                //返回 LCLBrowserWindow 窗口结构
	EnableDefaultCloseEvent()                                        //启用默认关闭事件
	WindowParent() ICEFWindowParent                                  //浏览器父窗口组件
	DisableTransparent()                                             //禁用窗口透明
	EnableTransparent(value uint8)                                   //启用并设置窗口透明
	DisableSystemMenu()                                              //禁用标题栏系统菜单
	DisableHelp()                                                    //禁用标题栏帮助
	EnableSystemMenu()                                               //启用标题栏系统菜单
	EnableHelp()                                                     //启用标题栏帮助
	NewTray() ITray                                                  //创建LCL的系统托盘
	SetRoundRectRgn(rgn int)                                         //窗口无边框时圆角设置
	ChromiumCreate(config *TCefChromiumConfig, defaultUrl string)    //chromium实例为空时创建 chromium
	BroderDirectionAdjustments() et.BroderDirectionAdjustments       //返回可以调整窗口大小的边框方向, 默认所有方向
	SetBroderDirectionAdjustments(val et.BroderDirectionAdjustments) // 设置可以调整窗口大小的边框方向, 默认所有方向
}

// IViewsFrameworkBrowserWindow
// 浏览器 VF 窗口组件接口 继承 IBrowserWindow
//
// 定义了ViewsFramework常用函数
type IViewsFrameworkBrowserWindow interface {
	IBrowserWindow
	BrowserWindow() *ViewsFrameworkBrowserWindow                       //返回 ViewsFrameworkBrowserWindow 窗口结构
	CreateTopLevelWindow()                                             //创建窗口, 在窗口组件中需要默认调用Show函数
	CenterWindow(size TCefSize)                                        //设置窗口居中，同时指定窗口大小
	Component() lcl.IComponent                                         //窗口父组件
	WindowComponent() *TCEFWindowComponent                             //窗口组件
	BrowserViewComponent() *TCEFBrowserViewComponent                   //窗口浏览器组件
	SetOnWindowCreated(onWindowCreated windowOnWindowCreated)          //设置窗口默认的创建回调事件函数
	SetOnGetInitialBounds(onGetInitialBounds windowOnGetInitialBounds) //设置窗口初始bounds
}

type IAuxTools interface {
	SetDevTools(devToolsWindow *devToolsWindow)
	DevTools() *devToolsWindow
}

// NewWindowProperty
// 创建一个属性配置器，带有窗口默认属性值
func NewWindowProperty() WindowProperty {
	return WindowProperty{
		Title:                     defaultTitle,
		Url:                       defaultAboutBlank,
		EnableMinimize:            true,
		EnableMaximize:            true,
		EnableResize:              true,
		EnableClose:               true,
		EnableCenterWindow:        true,
		EnableWebkitAppRegionDClk: true,
		EnableMainWindow:          true,
		ShowInTaskBar:             types.StAlways,
		X:                         100,
		Y:                         100,
		Width:                     800,
		Height:                    600,
	}
}

func (m *auxTools) SetDevTools(devToolsWindow *devToolsWindow) {
	m.devToolsWindow = devToolsWindow
}

func (m *auxTools) DevTools() *devToolsWindow {
	return m.devToolsWindow
}

// NewBrowserWindow
//
//	创建浏览器窗口
//	根据当前主窗口类型创建
//	窗口类型
//		  	LCL: 是基于LCL组件库创建的窗口，相比VF有多更的原生小部件使用，更多的窗口操作
//			VF : 是基于CEF ViewFramework 组件创建的窗口, 相比LCL无法使用系统原生小部件，较少的窗口操作
//	config: Chromium配置, 提供快捷chromium配置
//	windowProperty: 窗口属性
//	owner: 被创建组件拥有者
func NewBrowserWindow(config *TCefChromiumConfig, windowProperty WindowProperty, owner lcl.IComponent) IBrowserWindow {
	// 获取当前应用的主窗口
	main := BrowserWindow.MainWindow()
	// 设置为子窗口
	windowProperty.WindowType = consts.WT_POPUP_SUB_BROWSER
	if main.IsLCL() {
		// 创建LCL窗口
		return NewLCLBrowserWindow(config, windowProperty, owner)
	} else if main.IsViewsFramework() {
		// 创建VF窗口
		return NewViewsFrameworkBrowserWindow(config, windowProperty, owner)
	}
	return nil
}

// RunOnMainThread
//
//	在UI主线程中运行
func RunOnMainThread(fn func()) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	logger.Debug("MainThreadId:", api.DMainThreadId(), "CurrentThreadId:", api.DCurrentThreadId(), "IsMessageLoop:", application.IsMessageLoop())
	// MacOS 虽然当前线程是主线程, 但还是需要在UI异步线程中才可正确执行
	if !common.IsDarwin() && api.DMainThreadId() == api.DCurrentThreadId() {
		fn()
	} else {
		// 当前窗口模式是VF时，使用 lcl.ThreadSync, 在运行应用时初始化Application
		if application.IsMessageLoop() {
			lcl.ThreadSync(func() {
				fn()
			})
		} else {
			// 当前窗口模式LCL时，使用 QueueAsyncCall
			QueueAsyncCall(func(id int) {
				fn()
			})
		}
	}
}
