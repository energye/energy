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
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// auxTools 辅助工具
type auxTools struct {
	devToolsWindow   *devToolsWindow //devTools
	viewSourceWindow IBrowserWindow  //viewSource
}

// WindowProperty
// 窗口属性配置器
//
// 部分属性配置并不支持所有平台
type WindowProperty struct {
	IsShowModel        bool               //是否以模态窗口显示
	windowState        types.TWindowState //窗口 状态
	WindowType         consts.WINDOW_TYPE //窗口 类型 WINDOW_TYPE default: WT_MAIN_BROWSER
	Title              string             //窗口 标题
	Url                string             //默认打开URL
	Icon               string             //窗口图标 加载本地图标 local > /app/resources/icon.ico, linux使用png
	IconFS             string             //窗口图标 加载emfs内置图标 emfs > resources/icon.ico, linux使用png
	EnableHideCaption  bool               //窗口 是否隐藏标题栏, VF窗口组件Linux下不能动态控制
	EnableMinimize     bool               //窗口 是否启用最小化 default: true
	EnableMaximize     bool               //窗口 是否启用最大化 default: true
	EnableResize       bool               //窗口 是否允许调整大小 default: true
	EnableClose        bool               //窗口 关闭时是否关闭窗口 default: true
	EnableCenterWindow bool               //窗口 居中显示 default: true
	EnableDragFile     bool               //窗口 是否允许向窗口内拖拽文件
	AlwaysOnTop        bool               //窗口 窗口置顶
	X                  int32              //窗口 EnableCenterWindow=false X坐标 default: 100
	Y                  int32              //窗口 EnableCenterWindow=false Y坐标 default: 100
	Width              int32              //窗口 宽 default: 1024
	Height             int32              //窗口 高 default: 768
}

// IBrowserWindow
// 浏览器窗口基础接口
//
// 定义了常用函数
type IBrowserWindow interface {
	Id() int32                                                                                                //窗口ID
	Handle() types.HWND                                                                                       //窗口句柄
	Show()                                                                                                    //显示窗口
	Hide()                                                                                                    //隐藏窗口
	WindowState() types.TWindowState                                                                          //返回窗口最小化、最大化、全屏状态
	Maximize()                                                                                                //窗口最大化
	Minimize()                                                                                                //窗口最小化
	Restore()                                                                                                 //窗口还原
	Close()                                                                                                   //关闭窗口 非browser窗口使用
	CloseBrowserWindow()                                                                                      //关闭浏览器窗口 带有browser窗口使用
	WindowType() consts.WINDOW_TYPE                                                                           //窗口类型
	SetWindowType(windowType consts.WINDOW_TYPE)                                                              //设置窗口类型
	Browser() *ICefBrowser                                                                                    //窗口内的Browser对象
	Chromium() IChromium                                                                                      //窗口内的Chromium对象
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
	Point() *TCefPoint                                                                                        //窗口坐标
	Size() *TCefSize                                                                                          //窗口宽高
	Bounds() *TCefRect                                                                                        //窗口坐标和宽高
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
	RunOnMainThread(fn func())                                                                                //在主线程中运行
}

// ILCLBrowserWindow
// 浏览器 LCL 窗口组件接口 继承 IBrowserWindow
//
// 定义了LCL常用函数
type ILCLBrowserWindow interface {
	IBrowserWindow
	BrowserWindow() *LCLBrowserWindow //返回 LCLBrowserWindow 窗口结构
	EnableDefaultCloseEvent()         //启用默认关闭事件
	WindowParent() ICEFWindowParent   //浏览器父窗口组件
	DisableTransparent()              //禁用窗口透明
	EnableTransparent(value uint8)    //启用并设置窗口透明
	DisableSystemMenu()               //禁用标题栏系统菜单
	DisableHelp()                     //禁用标题栏帮助
	EnableSystemMenu()                //启用标题栏系统菜单
	EnableHelp()                      //启用标题栏帮助
	NewTray() ITray                   //创建LCL的系统托盘
	FramelessForLine()                //无边框四边一条细线样式
}

// IViewsFrameworkBrowserWindow
// 浏览器 VF 窗口组件接口 继承 IBrowserWindow
//
// 定义了ViewsFramework常用函数
type IViewsFrameworkBrowserWindow interface {
	IBrowserWindow
	BrowserWindow() *ViewsFrameworkBrowserWindow                                //返回 ViewsFrameworkBrowserWindow 窗口结构
	CreateTopLevelWindow()                                                      //创建窗口, 在窗口组件中需要默认调用Show函数
	CenterWindow(size *TCefSize)                                                //设置窗口居中，同时指定窗口大小
	Component() lcl.IComponent                                                  //窗口父组件
	WindowComponent() *TCEFWindowComponent                                      //窗口组件
	BrowserViewComponent() *TCEFBrowserViewComponent                            //窗口浏览器组件
	SetOnWindowCreated(onWindowCreated WindowComponentOnWindowCreated)          //设置窗口默认的创建回调事件函数
	SetOnGetInitialBounds(onGetInitialBounds WindowComponentOnGetInitialBounds) //设置窗口初始bounds
}

type IAuxTools interface {
	SetDevTools(devToolsWindow *devToolsWindow)
	SetViewSource(viewSourceWindow IBrowserWindow)
	DevTools() *devToolsWindow
	ViewSource() IBrowserWindow
}

// NewWindowProperty
// 创建一个属性配置器，带有窗口默认属性值
func NewWindowProperty() WindowProperty {
	return WindowProperty{
		Title:              "ENERGY",
		Url:                "about:blank",
		EnableMinimize:     true,
		EnableMaximize:     true,
		EnableResize:       true,
		EnableClose:        true,
		EnableCenterWindow: true,
		X:                  100,
		Y:                  100,
		Width:              1024,
		Height:             768,
	}
}

func (m *auxTools) SetDevTools(devToolsWindow *devToolsWindow) {
	m.devToolsWindow = devToolsWindow
}

func (m *auxTools) SetViewSource(viewSourceWindow IBrowserWindow) {
	m.viewSourceWindow = viewSourceWindow
}

func (m *auxTools) DevTools() *devToolsWindow {
	return m.devToolsWindow
}

func (m *auxTools) ViewSource() IBrowserWindow {
	return m.viewSourceWindow
}
