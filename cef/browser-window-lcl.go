//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// LCL窗口组件定义和实现

package cef

import (
	"fmt"
	"github.com/energye/energy/v2/cef/internal/assets"
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/cef/internal/window"
	. "github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/consts/messages"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/energy/v2/pkgs/assetserve"
	t "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"runtime"
	"time"
	"unsafe"
)

// LCLBrowserWindow 基于CEF lcl 窗口组件
//
// 该窗口使用CEF和LCL组件实现，CEF<=1.106.xx版本 在windows、MacOSX可正常使用, Linux无法输入中文, CEF>=2.107.xx版本linux强制使用 ViewsFrameworkBrowserWindow 窗口组件
type LCLBrowserWindow struct {
	*lcl.TForm                                     //window form
	isFormCreate              bool                 //是否创建完成 WindowForm
	chromiumBrowser           ICEFChromiumBrowser  //浏览器
	windowProperty            *WindowProperty      //窗口属性
	windowId                  int32                //窗口ID
	windowType                consts.WINDOW_TYPE   //窗口类型
	isClosing                 bool                 //
	canClose                  bool                 //
	onResize                  TNotifyEvent         //扩展事件
	windowResize              TNotifyEvent         //扩展事件
	onActivate                TNotifyEvent         //扩展事件
	onShow                    TNotifyEvent         //扩展事件
	onClose                   TCloseEvent          //扩展事件
	onCloseQuery              TCloseQueryEvent     //扩展事件
	onActivateAfter           lcl.TNotifyEvent     //扩展事件
	auxTools                  IAuxTools            //辅助工具
	tray                      ITray                //托盘
	hWnd                      types.HWND           //
	cwcap                     *customWindowCaption //自定义窗口标题栏
	drag                      *drag                //自定义拖拽
	wmPaintMessage            wmPaint              //
	wmMoveMessage             wmMove               //
	wmSizeMessage             wmSize               //
	wmWindowPosChangedMessage wmWindowPosChanged   //
	screen                    IScreen              //屏幕
}

// NewLCLBrowserWindow 创建一个 LCL 带有 chromium 窗口
//  该窗口默认不具备默认事件处理能力, 通过 EnableDefaultEvent 函数注册事件处理
func NewLCLBrowserWindow(config *TCefChromiumConfig, windowProperty WindowProperty, owner ...lcl.IComponent) *LCLBrowserWindow {
	var browseWindow = NewLCLWindow(windowProperty, owner...)
	browseWindow.ChromiumCreate(config, windowProperty.Url)
	//OnBeforeBrowser 是一个必须的默认事件，在浏览器创建时窗口序号会根据browserId生成
	browseWindow.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool) bool {
		//chromiumOnBeforeBrowser(browser, frame)
		return false
	})
	return browseWindow
}

// NewLCLWindow 创建一个LCL window窗口
func NewLCLWindow(windowProperty WindowProperty, owner ...lcl.IComponent) *LCLBrowserWindow {
	var window = &LCLBrowserWindow{}
	if len(owner) > 0 {
		window.TForm = lcl.NewForm(owner[0])
	} else {
		window.TForm = lcl.NewForm(BrowserWindow.mainBrowserWindow)
		//lcl.Application.CreateForm(&window)
	}
	window.windowProperty = &windowProperty
	window.cwcap = &customWindowCaption{
		bw: window,
	}
	window.SetWindowType(windowProperty.WindowType)
	window.SetDoubleBuffered(true)
	window.FormCreate()
	window.SetShowInTaskBar()
	window.defaultWindowEvent()
	window.setProperty()
	return window
}

// 设置属性
func (m *LCLBrowserWindow) setProperty() {
	wp := m.WindowProperty()
	m.SetTitle(wp.Title)
	if wp.IconFS != "" {
		if emfs.IsExist(wp.IconFS) {
			_ = lcl.Application.Icon().LoadFromFSFile(wp.IconFS)
		}
	} else if wp.Icon != "" {
		if tools.IsExist(wp.Icon) {
			lcl.Application.Icon().LoadFromFile(wp.Icon)
		}
	} else {
		// 默认
		// vf png
		// lcl ico
		if iconData := assets.DefaultICOICON(); iconData != nil {
			lcl.Application.Icon().LoadFromBytes(iconData)
		}
	}
	if wp.EnableCenterWindow {
		m.SetSize(wp.Width, wp.Height)
		m.SetCenterWindow(true)
	} else {
		m.SetPosition(types.PoDesigned)
		m.SetBounds(wp.X, wp.Y, wp.Width, wp.Height)
	}
	if wp.AlwaysOnTop {
		m.SetFormStyle(types.FsSystemStayOnTop)
	}
	if wp.EnableHideCaption {
		m.HideTitle()
	} else {
		if !wp.EnableMinimize {
			m.DisableMinimize()
		}
		if !wp.EnableMaximize {
			m.DisableMaximize()
		}
		if !wp.EnableResize {
			m.SetBorderStyle(types.BsSingle)
		}
	}
	if wp.EnableResize {
		c := m.Constraints()
		if wp.MinWidth > 0 && wp.MinHeight > 0 {
			c.SetMinWidth(wp.MinWidth)
			c.SetMinWidth(wp.MinHeight)
		}
		if wp.MaxWidth > 0 && wp.MaxHeight > 0 {
			c.SetMinWidth(wp.MaxWidth)
			c.SetMinWidth(wp.MaxHeight)
		}
	}
	// 只有隐藏窗口标题时才全屏
	if wp.EnableHideCaption && wp.WindowInitState == types.WsFullScreen {
		m.FullScreen()
	} else {
		m.SetWindowState(wp.WindowInitState)
	}
	// 当前窗口状态
	m.setCurrentProperty()
}

// FullScreen 窗口全屏
func (m *LCLBrowserWindow) FullScreen() {
	if m.WindowProperty().EnableHideCaption {
		m.RunOnMainThread(func() {
			m.WindowProperty().current.ws = types.WsFullScreen
			m.setCurrentProperty()
			m.SetBoundsRect(m.Monitor().BoundsRect())
		})
	}
}

// ExitFullScreen 窗口退出全屏
func (m *LCLBrowserWindow) ExitFullScreen() {
	wp := m.WindowProperty()
	if wp.EnableHideCaption && wp.current.ws == types.WsFullScreen {
		m.RunOnMainThread(func() {
			wp.current.ws = types.WsNormal
			m.SetWindowState(types.WsNormal)
			m.SetBounds(wp.current.x, wp.current.y, wp.current.w, wp.current.h)
		})
	}
}

// IsFullScreen 是否全屏
func (m *LCLBrowserWindow) IsFullScreen() bool {
	return m.WindowProperty().current.ws == types.WsFullScreen
}

// Handle 窗口句柄
func (m *LCLBrowserWindow) Handle() types.HWND {
	if m.hWnd == 0 {
		m.hWnd = m.TForm.Handle()
	}
	return m.hWnd
}

// RunOnMainThread
//	在UI主线程中运行
func (m *LCLBrowserWindow) RunOnMainThread(fn func()) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if api.DMainThreadId() == api.DCurrentThreadId() {
		fn()
	} else {
		QueueAsyncCall(func(id int) {
			fn()
		})
	}
}

// BrowserWindow 返回LCL窗口组件实例对象
func (m *LCLBrowserWindow) BrowserWindow() *LCLBrowserWindow {
	return m
}

// AsViewsFrameworkBrowserWindow 转换为VF窗口组件，这里返回nil
func (m *LCLBrowserWindow) AsViewsFrameworkBrowserWindow() IViewsFrameworkBrowserWindow {
	return nil
}

// AsLCLBrowserWindow 转换为LCL窗口组件，这里返回LCL窗口组件
func (m *LCLBrowserWindow) AsLCLBrowserWindow() ILCLBrowserWindow {
	return m
}

// SetCenterWindow 窗口居中
func (m *LCLBrowserWindow) SetCenterWindow(value bool) {
	if m.TForm == nil {
		return
	}
	if value {
		m.ScreenCenter()
	} else {
		m.SetPosition(types.PoDesigned)
	}
}

// Close 关闭窗口 非browser窗口使用
func (m *LCLBrowserWindow) Close() {
	if m.TForm == nil {
		return
	}
	m.TForm.Close()
}

// SetTitle 设置窗口标题栏标题
func (m *LCLBrowserWindow) SetTitle(title string) {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().Title = title
	m.TForm.SetCaption(title)
}

// SetWidth 设置窗口宽
func (m *LCLBrowserWindow) SetWidth(value int32) {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().Width = value
	m.TForm.SetWidth(value)
}

// SetHeight 设置窗口高
func (m *LCLBrowserWindow) SetHeight(value int32) {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().Height = value
	m.TForm.SetHeight(value)
}

// Point 窗口坐标
func (m *LCLBrowserWindow) Point() *TCefPoint {
	if m.TForm == nil {
		return nil
	}
	result := &TCefPoint{
		X: m.Left(),
		Y: m.Top(),
	}
	m.WindowProperty().X = result.X
	m.WindowProperty().Y = result.Y
	return result
}

// Size 窗口宽高
func (m *LCLBrowserWindow) Size() *TCefSize {
	if m.TForm == nil {
		return nil
	}
	result := &TCefSize{
		Width:  m.Width(),
		Height: m.Height(),
	}
	m.WindowProperty().Width = result.Width
	m.WindowProperty().Height = result.Height
	return result
}

// Bounds 窗口坐标和宽高
func (m *LCLBrowserWindow) Bounds() *TCefRect {
	if m.TForm == nil {
		return nil
	}
	rect := m.BoundsRect()
	result := &TCefRect{
		X:      rect.Left,
		Y:      rect.Top,
		Width:  rect.Width(),
		Height: rect.Height(),
	}
	m.WindowProperty().X = result.X
	m.WindowProperty().Y = result.Y
	m.WindowProperty().Width = result.Width
	m.WindowProperty().Height = result.Height
	return result
}

// SetPoint 设置窗口坐标
func (m *LCLBrowserWindow) SetPoint(x, y int32) {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
	m.TForm.SetLeft(x)
	m.TForm.SetTop(y)
}

// SetSize 设置窗口宽高
func (m *LCLBrowserWindow) SetSize(width, height int32) {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
	m.SetWidth(width)
	m.SetHeight(height)
}

// SetBounds 设置窗口坐标和宽高
func (m *LCLBrowserWindow) SetBounds(x, y, width, height int32) {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().X = x
	m.WindowProperty().Y = y
	m.WindowProperty().Width = width
	m.WindowProperty().Height = height
	m.TForm.SetBounds(x, y, width, height)
}

// GetAuxTools
func (m *LCLBrowserWindow) GetAuxTools() IAuxTools {
	return m.auxTools
}

// createAuxTools
func (m *LCLBrowserWindow) createAuxTools() {
	if m.auxTools == nil {
		m.auxTools = &auxTools{}
	}
}

// Browser
func (m *LCLBrowserWindow) Browser() *ICefBrowser {
	if m == nil || m.Chromium() == nil || !m.Chromium().Initialized() {
		return nil
	}
	return m.Chromium().Browser()
}

// Chromium 返回 chromium
func (m *LCLBrowserWindow) Chromium() IChromium {
	if m.chromiumBrowser == nil {
		return nil
	}
	return m.chromiumBrowser.Chromium()
}

// Id 浏览器窗口ID
func (m *LCLBrowserWindow) Id() int32 {
	if m.windowId == 0 {
		m.windowId = m.Chromium().BrowserId()
	}
	return m.windowId
}

// Show
func (m *LCLBrowserWindow) Show() {
	if m.TForm == nil {
		return
	}
	m.TForm.Show()
}

// Hide
func (m *LCLBrowserWindow) Hide() {
	if m.TForm == nil {
		return
	}
	m.TForm.Hide()
}

// WindowState 返回窗口最小化、最大化、全屏状态
func (m *LCLBrowserWindow) WindowState() types.TWindowState {
	if m.TForm == nil {
		return -1
	}
	return m.TForm.WindowState()
}

// Visible
func (m *LCLBrowserWindow) Visible() bool {
	if m.TForm == nil {
		return false
	}
	return m.TForm.Visible()
}

// SetVisible
func (m *LCLBrowserWindow) SetVisible(value bool) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetVisible(value)
}

// SetDefaultInTaskBar 以默认的方式展示在任务栏上
func (m *LCLBrowserWindow) SetDefaultInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StDefault)
}

// SetShowInTaskBar 展示在任务栏上
func (m *LCLBrowserWindow) SetShowInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StAlways)
}

// SetNotInTaskBar 不会展示在任务栏上
func (m *LCLBrowserWindow) SetNotInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StNever)
}

// WindowParent
//  返回chromium的父组件对象，该对象不是window组件对象,属于window的一个子组件
//  在windows下它是 TCEFWindowParent, linux或macOSx下它是 TCEFLinkedWindowParent
//  通过函数可调整该组件的属性
func (m *LCLBrowserWindow) WindowParent() ICEFWindowParent {
	return m.chromiumBrowser.WindowParent()
}

// IsClosing 返回窗口是否正在关闭/或已关闭 true正在或已关闭
func (m *LCLBrowserWindow) IsClosing() bool {
	return m.isClosing
}

// SetWindowType 设置窗口类型
func (m *LCLBrowserWindow) SetWindowType(windowType consts.WINDOW_TYPE) {
	m.windowType = windowType
}

// WindowType 返回窗口类型
func (m *LCLBrowserWindow) WindowType() consts.WINDOW_TYPE {
	return m.windowType
}

// ChromiumCreate
//  创建window浏览器组件
//  不带有默认事件的chromium
func (m *LCLBrowserWindow) ChromiumCreate(config *TCefChromiumConfig, defaultUrl string) {
	if m.chromiumBrowser != nil {
		return
	}
	if config == nil {
		config = NewChromiumConfig()
	}
	m.chromiumBrowser = NewChromiumBrowser(m, config)
	m.Chromium().SetEnableMultiBrowserMode(true)

	if defaultUrl != "" {
		m.Chromium().SetDefaultURL(defaultUrl)
	}
	//windowParent
	m.WindowParent().DefaultAnchors()
	m.WindowParent().SetOnEnter(func(sender lcl.IObject) {
		if m.isClosing {
			return
		}
		m.Chromium().Initialized()
		m.Chromium().FrameIsFocused()
		m.Chromium().SetFocus(true)
	})
	m.WindowParent().SetOnExit(func(sender lcl.IObject) {
		if m.isClosing {
			return
		}
		m.Chromium().SendCaptureLostEvent()
	})
	// 主窗口、子窗口chromium检测
	if m.WindowProperty().WindowType != consts.WT_DEV_TOOLS {
		window.CurrentBrowseWindowCache = m
	}
}

// WindowProperty 部分提供部分窗口属性设置
func (m *LCLBrowserWindow) WindowProperty() *WindowProperty {
	return m.windowProperty
}

// defaultChromiumEvent 默认的chromium事件
func (m *LCLBrowserWindow) defaultChromiumEvent() {
	if m.WindowType() != consts.WT_DEV_TOOLS {
		m.registerPopupEvent()
		m.registerDefaultEvent()
		m.registerDefaultChromiumCloseEvent()
	}
}

// FormCreate
//  创建窗口
//  不带有默认事件的窗口
func (m *LCLBrowserWindow) FormCreate() {
	if m.isFormCreate {
		return
	}
	m.isFormCreate = true
	m.SetName(fmt.Sprintf("energy_window_name_%d", time.Now().UnixNano()/1e6))
	m.onFormMessages()
}

// defaultWindowEvent 默认窗口活动/关闭处理事件
func (m *LCLBrowserWindow) defaultWindowEvent() {
	if m.WindowType() != consts.WT_DEV_TOOLS {
		m.TForm.SetOnActivate(m.activate)
	}
	m.TForm.SetOnResize(m.resize)
	m.TForm.SetOnShow(m.show)
}

// defaultWindowCloseEvent 默认的窗口关闭事件
func (m *LCLBrowserWindow) defaultWindowCloseEvent() {
	m.TForm.SetOnClose(m.close)
	m.TForm.SetOnCloseQuery(m.closeQuery)
}

// EnableDefaultCloseEvent 启用默认关闭事件
func (m *LCLBrowserWindow) EnableDefaultCloseEvent() {
	m.defaultWindowCloseEvent()
	m.registerDefaultChromiumCloseEvent()
}

// EnableAllDefaultEvent 启用所有默认事件行为
func (m *LCLBrowserWindow) EnableAllDefaultEvent() {
	// 窗口关闭事件，window和chromium将以正确的行为关闭
	m.defaultWindowCloseEvent()
	// chromium事件，在回调事件中实现框架的默认行为
	m.defaultChromiumEvent()
	// 仅windows有的事件，窗口消息事件
	m.registerWindowsCompMsgEvent()
}

// SetOnResize 事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) SetOnResize(fn TNotifyEvent) {
	m.onResize = fn
}

// SetOnActivate 事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) SetOnActivate(fn TNotifyEvent) {
	m.onActivate = fn
}

// SetOnShow 事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) SetOnShow(fn TNotifyEvent) {
	m.onShow = fn
}

// SetOnClose 事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) SetOnClose(fn TCloseEvent) {
	m.onClose = fn
}

// SetOnCloseQuery 事件,不会覆盖默认事件，返回值：false继续执行默认事件, true跳过默认事件
func (m *LCLBrowserWindow) SetOnCloseQuery(fn TCloseQueryEvent) {
	m.onCloseQuery = fn
}

// SetOnActivateAfter 每次激活窗口之后执行一次
func (m *LCLBrowserWindow) SetOnActivateAfter(fn lcl.TNotifyEvent) {
	m.onActivateAfter = fn
}

// Minimize
func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	m.RunOnMainThread(func() {
		m.SetWindowState(types.WsMinimized)
	})
}

// Restore
func (m *LCLBrowserWindow) Restore() {
	m.RunOnMainThread(func() {
		m.SetWindowState(types.WsNormal)
	})
}

// DisableTransparent 禁用口透明
func (m *LCLBrowserWindow) DisableTransparent() {
	if m.TForm == nil {
		return
	}
	m.SetAlphaBlend(false)
	m.SetAlphaBlendValue(255)
}

// EnableTransparent 使窗口透明 value 0 ~ 255
func (m *LCLBrowserWindow) EnableTransparent(value uint8) {
	if m.TForm == nil {
		return
	}
	m.SetAlphaBlend(true)
	m.SetAlphaBlendValue(value)
}

// DisableMinimize 禁用最小化按钮
func (m *LCLBrowserWindow) DisableMinimize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Exclude(types.BiMinimize))
	m.WindowProperty().EnableMinimize = false
	m.EnabledMinimize(m.WindowProperty().EnableMinimize)
}

// DisableMaximize 禁用最大化按钮
func (m *LCLBrowserWindow) DisableMaximize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Exclude(types.BiMaximize))
	m.WindowProperty().EnableMaximize = false
	m.EnabledMaximize(m.WindowProperty().EnableMaximize)
}

// DisableResize 禁用调整窗口大小
func (m *LCLBrowserWindow) DisableResize() {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().EnableResize = false
	if !m.WindowProperty().EnableHideCaption {
		m.TForm.SetBorderStyle(types.BsSingle)
	}
}

// DisableSystemMenu 禁用系统菜单-同时禁用最小化，最大化，关闭按钮
func (m *LCLBrowserWindow) DisableSystemMenu() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Exclude(types.BiSystemMenu))
	m.EnabledSystemMenu(false)
}

// DisableHelp 禁用帮助菜单
func (m *LCLBrowserWindow) DisableHelp() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Exclude(types.BiHelp))
}

// EnableMinimize 启用最小化按钮
func (m *LCLBrowserWindow) EnableMinimize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Include(types.BiMinimize))
	m.WindowProperty().EnableMinimize = true
	m.EnabledMinimize(m.WindowProperty().EnableMinimize)
}

// EnableMaximize 启用最大化按钮
func (m *LCLBrowserWindow) EnableMaximize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Include(types.BiMaximize))
	m.WindowProperty().EnableMaximize = true
	m.EnabledMaximize(m.WindowProperty().EnableMaximize)
}

// EnableResize 启用允许调整窗口大小
func (m *LCLBrowserWindow) EnableResize() {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().EnableResize = true
	if !m.WindowProperty().EnableHideCaption {
		m.TForm.SetBorderStyle(types.BsSizeable)
	}
}

// EnableSystemMenu 启用系统菜单
func (m *LCLBrowserWindow) EnableSystemMenu() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Include(types.BiSystemMenu))
	m.EnabledSystemMenu(true)
}

// EnableHelp  启用帮助菜单
func (m *LCLBrowserWindow) EnableHelp() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Include(types.BiHelp))
}

// IsViewsFramework 返回是否VF窗口组件，这里返回false
func (m *LCLBrowserWindow) IsViewsFramework() bool {
	return false
}

// IsLCL 返回是否LCL窗口组件，这里返回true
func (m *LCLBrowserWindow) IsLCL() bool {
	return true
}

// show 内部调用
func (m *LCLBrowserWindow) show(sender lcl.IObject) {
	if m.onShow != nil {
		m.onShow(sender)
	}
}

// SetCreateBrowserExtraInfo
//  设置 Chromium 创建浏览器时设置的扩展信息
func (m *LCLBrowserWindow) SetCreateBrowserExtraInfo(windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) {
	if m.chromiumBrowser != nil {
		m.chromiumBrowser.SetCreateBrowserExtraInfo(windowName, context, extraInfo)
	}
}

// resize 内部调用
func (m *LCLBrowserWindow) resize(sender lcl.IObject) {
	var ret bool
	if m.onResize != nil {
		if m.onResize(sender) {
			ret = true
		}
	}
	if !ret {
		if m.isClosing {
			return
		}
		m.setCurrentProperty()
		if m.windowResize != nil {
			m.windowResize(sender)
		}
		if m.chromiumBrowser != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
			if m.chromiumBrowser.WindowParent() != nil {
				m.chromiumBrowser.WindowParent().UpdateSize()
			}
		}
	}
}

// 在窗口坐标、大小、全屏时保存当前窗口属性
func (m *LCLBrowserWindow) setCurrentProperty() {
	if m.WindowProperty().current.ws == types.WsFullScreen {
		return
	}
	boundRect := m.BoundsRect()
	m.WindowProperty().current.x = boundRect.Left
	m.WindowProperty().current.y = boundRect.Top
	m.WindowProperty().current.w = boundRect.Width()
	m.WindowProperty().current.h = boundRect.Height()
}

// activate 内部调用
func (m *LCLBrowserWindow) activate(sender lcl.IObject) {
	var ret bool
	if m.onActivate != nil {
		ret = m.onActivate(sender)
	}
	if !ret {
		if m.isClosing {
			return
		}
		if m.chromiumBrowser != nil && !m.chromiumBrowser.IsCreated() {
			if m.WindowProperty().WindowType != consts.WT_DEV_TOOLS {
				m.chromiumBrowser.CreateBrowser()
			}
		}
	}
	if m.onActivateAfter != nil {
		m.onActivateAfter(sender)
	}
}

// registerDefaultEvent 注册默认事件 部分事件允许被覆盖
func (m *LCLBrowserWindow) registerDefaultEvent() {
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
	//事件可以被覆盖
	m.Chromium().SetOnBeforeDownload(func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *ICefDownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) {
		if bwEvent.onBeforeDownload != nil {
			bwEvent.onBeforeDownload(sender, browser, beforeDownloadItem, suggestedName, callback, m)
		} else {
			// 默认保存到当前执行文件所在目录
			callback.Cont(consts.ExeDir+consts.Separator+suggestedName, true)
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
	//事件可以被覆盖
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
		updateBrowserDevTools(m, browser, title)
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
			m.cwcap.regions = regions
			m.setDraggableRegions()
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

// registerPopupEvent 注册弹出子窗口事件
func (m *LCLBrowserWindow) registerPopupEvent() {
	var bwEvent = BrowserWindow.browserEvent
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
		if next := BrowserWindow.getNextLCLPopupWindow(); next != nil {
			bw := next.AsLCLBrowserWindow().BrowserWindow()
			bw.SetWindowType(consts.WT_POPUP_SUB_BROWSER)
			bw.ChromiumCreate(NewChromiumConfig(), beforePopupInfo.TargetUrl)
			var result = false
			if bwEvent.onBeforePopup != nil {
				result = bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, bw, noJavascriptAccess)
			}
			if !result { // true 表示用户自行处理
				bw.defaultWindowCloseEvent()
				bw.defaultChromiumEvent()
				bw.registerWindowsCompMsgEvent()
				bw.setProperty()
				m.RunOnMainThread(func() {
					// show window, run in main thread
					if bw.WindowProperty().IsShowModel {
						bw.ShowModal()
						return
					}
					bw.Show()
				})
				result = true
			}
			return result
		}
		return true
	})
}

// CloseBrowserWindow 关闭带有浏览器的窗口
func (m *LCLBrowserWindow) CloseBrowserWindow() {
	if m == nil || m.TForm == nil {
		return
	}
	m.RunOnMainThread(func() {
		if IsDarwin() {
			//main window close
			if m.WindowType() == consts.WT_MAIN_BROWSER {
				m.Close()
			} else {
				//sub window close
				m.isClosing = true
				m.Hide()
				m.Chromium().CloseBrowser(true)
			}
		} else {
			m.isClosing = true
			m.Hide()
			m.Chromium().CloseBrowser(true)
		}
	})
}

// close 内部调用
func (m *LCLBrowserWindow) close(sender lcl.IObject, action *types.TCloseAction) {
	var ret bool
	if m.onClose != nil {
		ret = m.onClose(sender, action)
	}
	if !ret {
		logger.Debug("window.onClose")
		if m.WindowType() == consts.WT_MAIN_BROWSER || !IsWindows() { // 主窗口 或 非windows
			*action = types.CaFree
		} else if IsWindows() { // windows 子窗口
			*action = types.CaHide
		}
	}
}

// closeQuery 内部调用
func (m *LCLBrowserWindow) closeQuery(sender lcl.IObject, close *bool) {
	var ret bool
	if m.onCloseQuery != nil {
		ret = m.onCloseQuery(sender, close)
	}
	if !ret {
		if m.tray != nil {
			m.tray.close()
		}
		logger.Debug("window.onCloseQuery windowType:", m.WindowType())
		if IsDarwin() {
			//main window close
			if m.WindowType() == consts.WT_MAIN_BROWSER {
				*close = true
				desChildWind := m.WindowParent().DestroyChildWindow()
				logger.Debug("window.onCloseQuery => windowParent.DestroyChildWindow:", desChildWind)
			} else {
				//sub window close
				*close = m.canClose
			}
		} else {
			*close = m.canClose
		}
		m.RunOnMainThread(func() {
			if !m.isClosing {
				m.isClosing = true
				m.Chromium().CloseBrowser(true)
				if IsDarwin() {
					m.Show() // mac 主窗口未得到焦点时应用不退出, 所以show一下
				}
				m.Hide() // 更快的关闭效果, 先隐藏掉
			}
		})
	}
}

// registerDefaultChromiumCloseEvent 注册默认的chromium关闭事件
func (m *LCLBrowserWindow) registerDefaultChromiumCloseEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.Chromium().SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
		logger.Debug("chromium.onClose")
		var flag = false
		if bwEvent.onClose != nil {
			flag = bwEvent.onClose(sender, browser, aAction, m)
		}
		if !flag {
			if IsDarwin() { //MacOSX
				desChildWind := m.WindowParent().DestroyChildWindow()
				logger.Debug("chromium.onClose => windowParent.DestroyChildWindow:", desChildWind)
				*aAction = consts.CbaClose
			} else if IsLinux() {
				*aAction = consts.CbaClose
			} else if IsWindows() {
				*aAction = consts.CbaDelay
			}
			if !IsDarwin() {
				m.RunOnMainThread(func() { //run in main thread
					m.WindowParent().Free()
					logger.Debug("chromium.onClose => windowParent.Free")
				})
			}
			m.cwcap.free() //释放自定义标题栏 rgn
		}
	})
	m.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		logger.Debug("chromium.onBeforeClose")
		var flag = false
		if bwEvent.onBeforeClose != nil {
			flag = bwEvent.onBeforeClose(sender, browser, m)
		}
		if !flag {
			m.canClose = true
			var closeWindow = func() {
				defer func() {
					if err := recover(); err != nil {
						logger.Error("chromium.OnBeforeClose Error:", err)
					}
				}()
				if m.GetAuxTools() != nil {
					m.GetAuxTools().SetDevTools(nil)
				}

				// LCLBrowserWindow 关闭
				if IsWindows() {
					rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
				} else if IsDarwin() || IsLinux() {
					m.Close()
				}
			}
			m.RunOnMainThread(func() { // main thread run
				closeWindow()
			})
			// 最后再移除
			BrowserWindow.removeWindowInfo(m.windowId)
			//chromiumOnBeforeClose(browser)
		}
	})
}

func (m *LCLBrowserWindow) Screen() IScreen {
	if m.screen == nil {
		m.screen = &Screen{window: m}
	}
	return m.screen
}

// wm message event
type messageType int32

const (
	mtMove messageType = iota + 1
	mtSize
	mtWindowPosChanged
	mtPaint
)

type wmPaint func(message *t.TPaint)
type wmMove func(message *t.TMove)
type wmSize func(message *t.TSize)
type wmWindowPosChanged func(message *t.TWindowPosChanged)

func (m *LCLBrowserWindow) onFormMessages() {
	m.setOnWMPaint(func(message *t.TPaint) {
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
		if m.wmPaintMessage != nil {
			m.wmPaintMessage(message)
		}
	})
	m.setOnWMMove(func(message *t.TMove) {
		m.setCurrentProperty()
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
		if m.wmMoveMessage != nil {
			m.wmMoveMessage(message)
		}
	})
	m.setOnWMSize(func(message *t.TSize) {
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
		if m.wmSizeMessage != nil {
			m.wmSizeMessage(message)
		}
	})
	m.setOnWMWindowPosChanged(func(message *t.TWindowPosChanged) {
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
		if m.wmWindowPosChangedMessage != nil {
			m.wmWindowPosChangedMessage(message)
		}
	})
}

func (m *LCLBrowserWindow) setOnWMPaint(fn wmPaint) {
	imports.Proc(def.Form_SetOnMessagesEvent).Call(m.Instance(), uintptr(mtPaint), api.MakeEventDataPtr(fn))
}

func (m *LCLBrowserWindow) setOnWMMove(fn wmMove) {
	imports.Proc(def.Form_SetOnMessagesEvent).Call(m.Instance(), uintptr(mtMove), api.MakeEventDataPtr(fn))
}

func (m *LCLBrowserWindow) setOnWMSize(fn wmSize) {
	imports.Proc(def.Form_SetOnMessagesEvent).Call(m.Instance(), uintptr(mtSize), api.MakeEventDataPtr(fn))
}

func (m *LCLBrowserWindow) setOnWMWindowPosChanged(fn wmWindowPosChanged) {
	imports.Proc(def.Form_SetOnMessagesEvent).Call(m.Instance(), uintptr(mtWindowPosChanged), api.MakeEventDataPtr(fn))
}

func (m *LCLBrowserWindow) SetOnWMPaint(fn wmPaint) {
	m.wmPaintMessage = fn
}

func (m *LCLBrowserWindow) SetOnWMMove(fn wmMove) {
	m.wmMoveMessage = fn
}

func (m *LCLBrowserWindow) SetOnWMSize(fn wmSize) {
	m.wmSizeMessage = fn
}

func (m *LCLBrowserWindow) SetOnWMWindowPosChanged(fn wmWindowPosChanged) {
	m.wmWindowPosChangedMessage = fn
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case wmPaint:
			fn.(wmPaint)((*t.TPaint)(getPtr(0)))
		case wmMove:
			fn.(wmMove)((*t.TMove)(getPtr(0)))
		case wmSize:
			fn.(wmSize)((*t.TSize)(getPtr(0)))
		case wmWindowPosChanged:
			fn.(wmWindowPosChanged)((*t.TWindowPosChanged)(getPtr(0)))
		default:
			return false
		}
		return true
	})
}
