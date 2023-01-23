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
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"time"
)

type IBaseWindow interface {
	lcl.IWinControl
	FormCreate()
	WindowParent() ITCefWindowParent
	Chromium() IChromium
	ChromiumCreate(config *tCefChromiumConfig, defaultUrl string)
	registerEvent()
	registerDefaultEvent()
}

//LCLBrowserWindow 基于CEF lcl 窗口组件
//
//该窗口使用CEF和LCL组件实现，CEF<=1.106.xx版本 在windows、MacOSX可正常使用, Linux无法输入中文, CEF>=2.107.xx版本linux强制使用 ViewsFrameworkBrowserWindow 窗口组件
type LCLBrowserWindow struct {
	*lcl.TForm                            //
	chromium         IChromium            //
	browser          *ICefBrowser         //
	windowParent     ITCefWindowParent    //
	windowProperty   *WindowProperty      //
	windowId         int32                //
	windowType       consts.WINDOW_TYPE   //窗口类型
	isClosing        bool                 //
	canClose         bool                 //
	onResize         TNotifyEvent         //
	windowResize     TNotifyEvent         //
	onActivate       TNotifyEvent         //
	onShow           TNotifyEvent         //
	onClose          TCloseEvent          //
	onCloseQuery     TCloseQueryEvent     //
	onActivateAfter  lcl.TNotifyEvent     //
	isFormCreate     bool                 //是否创建完成 WindowForm
	isChromiumCreate bool                 //是否创建完成 Chromium
	frames           TCEFFrame            //当前浏览器下的所有frame
	auxTools         *auxTools            //辅助工具
	tray             ITray                //托盘
	hWnd             types.HWND           //
	windowsState     types.TWindowState   //
	cwcap            *customWindowCaption //自定义窗口标题栏
}

//创建一个 LCL 带有 chromium 窗口
//
//该窗口默认不具备默认事件处理能力, 通过 EnableDefaultEvent 函数注册事件处理
func NewLCLBrowserWindow(config *tCefChromiumConfig, windowProperty WindowProperty) *LCLBrowserWindow {
	var window = NewLCLWindow(windowProperty)
	window.ChromiumCreate(config, windowProperty.Url)
	window.putChromiumWindowInfo()
	//OnBeforeBrowser 是一个必须的默认事件，在浏览器创建时窗口序号会根据browserId生成
	window.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool {
		chromiumOnBeforeBrowser(browser, frame)
		return false
	})
	return window
}

//创建一个LCL window窗口
func NewLCLWindow(windowProperty WindowProperty, owner ...lcl.IComponent) *LCLBrowserWindow {
	var window = &LCLBrowserWindow{}
	if len(owner) > 0 {
		window.TForm = lcl.NewForm(owner[0])
	} else {
		window.TForm = lcl.NewForm(BrowserWindow.mainBrowserWindow)
		//lcl.Application.CreateForm(&window)
	}
	window.windowProperty = &windowProperty
	window.cwcap = new(customWindowCaption)
	window.SetDoubleBuffered(true)
	window.FormCreate()
	window.SetShowInTaskBar()
	window.defaultWindowEvent()
	window.setProperty()
	return window
}

//设置属性
func (m *LCLBrowserWindow) setProperty() {
	m.SetTitle(m.windowProperty.Title)
	if m.windowProperty.IconFS != "" {
		if emfs.IsExist(m.windowProperty.IconFS) {
			_ = lcl.Application.Icon().LoadFromFSFile(m.windowProperty.IconFS)
		}
	} else if m.windowProperty.Icon != "" {
		if tools.IsExist(m.windowProperty.Icon) {
			lcl.Application.Icon().LoadFromFile(m.windowProperty.Icon)
		}
	}
	if m.windowProperty.EnableCenterWindow {
		m.SetSize(m.windowProperty.Width, m.windowProperty.Height)
		m.SetPosition(types.PoDesktopCenter)
	} else {
		m.SetPosition(types.PoDesigned)
		m.SetBounds(m.windowProperty.X, m.windowProperty.Y, m.windowProperty.Width, m.windowProperty.Height)
	}
	if m.windowProperty.AlwaysOnTop {
		m.SetFormStyle(types.FsSystemStayOnTop)
	}
	if m.windowProperty._EnableHideCaption {
		m.HideTitle()
	} else {
		if !m.windowProperty.EnableMinimize {
			m.DisableMinimize()
		}
		if !m.windowProperty.EnableMaximize {
			m.DisableMaximize()
		}
		if !m.windowProperty.EnableResize {
			m.SetBorderStyle(types.BsSingle)
		}
	}
}

func (m *LCLBrowserWindow) Handle() types.HWND {
	if m.hWnd == 0 {
		m.hWnd = m.TForm.Handle()
	}
	return m.hWnd
}

func (m *LCLBrowserWindow) BrowserWindow() *LCLBrowserWindow {
	return m
}

func (m *LCLBrowserWindow) AsViewsFrameworkBrowserWindow() IViewsFrameworkBrowserWindow {
	return nil
}

func (m *LCLBrowserWindow) AsLCLBrowserWindow() ILCLBrowserWindow {
	return m
}

func (m *LCLBrowserWindow) SetCenterWindow(value bool) {
	if m.TForm == nil {
		return
	}
	if value {
		m.SetPosition(types.PoDesktopCenter)
	} else {
		m.SetPosition(types.PoDesigned)
	}
}

func (m *LCLBrowserWindow) Close() {
	if m.TForm == nil {
		return
	}
	m.TForm.Close()
}

func (m *LCLBrowserWindow) SetTitle(title string) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetCaption(title)
}

func (m *LCLBrowserWindow) SetWidth(value int32) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetWidth(value)
}

func (m *LCLBrowserWindow) SetHeight(value int32) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetHeight(value)
}

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

func (m *LCLBrowserWindow) SetPoint(x, y int32) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetLeft(x)
	m.TForm.SetTop(y)
}

func (m *LCLBrowserWindow) SetSize(width, height int32) {
	if m.TForm == nil {
		return
	}
	m.SetWidth(width)
	m.SetHeight(height)
}

func (m *LCLBrowserWindow) SetBounds(x, y, width, height int32) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetBounds(x, y, width, height)
}

func (m *LCLBrowserWindow) getAuxTools() *auxTools {
	return m.auxTools
}

func (m *LCLBrowserWindow) createAuxTools() {
	if m.auxTools == nil {
		m.auxTools = &auxTools{}
	}
}

func (m *LCLBrowserWindow) Browser() *ICefBrowser {
	return m.browser
}

func (m *LCLBrowserWindow) setBrowser(browser *ICefBrowser) {
	m.browser = browser
}

func (m *LCLBrowserWindow) addFrame(frame *ICefFrame) {
	m.createFrames()
	m.frames[frame.Id] = frame
}

func (m *LCLBrowserWindow) Frames() TCEFFrame {
	return m.frames
}

func (m *LCLBrowserWindow) createFrames() {
	if m.frames == nil {
		m.frames = make(TCEFFrame)
	}
}

func (m *LCLBrowserWindow) Chromium() IChromium {
	return m.chromium
}

func (m *LCLBrowserWindow) Id() int32 {
	return m.windowId
}

func (m *LCLBrowserWindow) Show() {
	if m.TForm == nil {
		return
	}
	m.TForm.Show()
}

func (m *LCLBrowserWindow) Hide() {
	if m.TForm == nil {
		return
	}
	m.TForm.Hide()
}

func (m *LCLBrowserWindow) Visible() bool {
	if m.TForm == nil {
		return false
	}
	return m.TForm.Visible()
}

func (m *LCLBrowserWindow) SetVisible(value bool) {
	if m.TForm == nil {
		return
	}
	m.TForm.SetVisible(value)
}

//以默认的方式展示在任务栏上
func (m *LCLBrowserWindow) SetDefaultInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StDefault)
}

//展示在任务栏上
func (m *LCLBrowserWindow) SetShowInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StAlways)
}

//不会展示在任务栏上
func (m *LCLBrowserWindow) SetNotInTaskBar() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetShowInTaskBar(types.StNever)
}

//返回chromium的父组件对象，该对象不是window组件对象,属于window的一个子组件
//
//在windows下它是 TCEFWindowParent, linux或macOSx下它是 TCEFLinkedWindowParent
//
//通过函数可调整该组件的属性
func (m *LCLBrowserWindow) WindowParent() ITCefWindowParent {
	return m.windowParent
}

//返回窗口关闭状态
func (m *LCLBrowserWindow) IsClosing() bool {
	return m.isClosing
}

func (m *LCLBrowserWindow) ITray() {

}

// 设置窗口类型
func (m *LCLBrowserWindow) SetWindowType(windowType consts.WINDOW_TYPE) {
	m.windowType = windowType
}

// 返回窗口类型
func (m *LCLBrowserWindow) WindowType() consts.WINDOW_TYPE {
	return m.windowType
}

// 创建window浏览器组件
//
// 不带有默认事件的chromium
func (m *LCLBrowserWindow) ChromiumCreate(config *tCefChromiumConfig, defaultUrl string) {
	if m.isChromiumCreate {
		return
	}
	m.isChromiumCreate = true
	m.windowId = BrowserWindow.GetNextWindowNum()
	if config == nil {
		config = NewChromiumConfig()
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

func (m *LCLBrowserWindow) WindowProperty() *WindowProperty {
	return m.windowProperty
}

func (m *LCLBrowserWindow) putChromiumWindowInfo() {
	BrowserWindow.putWindowInfo(m.windowId, m)
}

//默认的chromium事件
func (m *LCLBrowserWindow) defaultChromiumEvent() {
	if m.WindowType() != consts.WT_DEV_TOOLS {
		AddGoForm(m.windowId, m.Instance())
		m.registerPopupEvent()
		m.registerDefaultEvent()
		m.registerDefaultChromiumCloseEvent()
	}
}

// 创建窗口
//
// 不带有默认事件的窗口
func (m *LCLBrowserWindow) FormCreate() {
	if m.isFormCreate {
		return
	}
	m.isFormCreate = true
	m.SetName(fmt.Sprintf("energy_window_name_%d", time.Now().UnixNano()/1e6))
}

//默认窗口活动/关闭处理事件
func (m *LCLBrowserWindow) defaultWindowEvent() {
	if m.WindowType() != consts.WT_DEV_TOOLS {
		m.TForm.SetOnResize(m.resize)
		m.TForm.SetOnActivate(m.activate)
	}
	m.TForm.SetOnShow(m.show)
}

//默认的窗口关闭事件
func (m *LCLBrowserWindow) defaultWindowCloseEvent() {
	m.TForm.SetOnClose(m.close)
	m.TForm.SetOnCloseQuery(m.closeQuery)
}

//启用默认关闭事件行为-该窗口将被关闭
func (m *LCLBrowserWindow) EnableDefaultCloseEvent() {
	m.defaultWindowCloseEvent()
	m.registerDefaultChromiumCloseEvent()
}

//启用所有默认事件行为
func (m *LCLBrowserWindow) EnableAllDefaultEvent() {
	m.defaultWindowCloseEvent()
	m.defaultChromiumEvent()
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

//每次激活窗口之后执行一次
func (m *LCLBrowserWindow) SetOnActivateAfter(fn lcl.TNotifyEvent) {
	m.onActivateAfter = fn
}

func (m *LCLBrowserWindow) Minimize() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		m.SetWindowState(types.WsMinimized)
	})
}

// 关闭带有浏览器的窗口
func (m *LCLBrowserWindow) CloseBrowserWindow() {
	if m.TForm == nil {
		return
	}
	QueueAsyncCall(func(id int) {
		if m == nil {
			logger.Error("关闭浏览器 WindowInfo 为空")
			return
		}
		if IsDarwin() {
			//main window close
			if m.WindowType() == consts.WT_MAIN_BROWSER {
				m.Close()
			} else {
				//sub window close
				m.isClosing = true
				m.Hide()
				m.chromium.CloseBrowser(true)
			}
		} else {
			m.isClosing = true
			m.Hide()
			m.chromium.CloseBrowser(true)
		}
	})
}

//禁用口透明
func (m *LCLBrowserWindow) DisableTransparent() {
	if m.TForm == nil {
		return
	}
	m.SetAlphaBlend(false)
	m.SetAlphaBlendValue(255)
}

//使窗口透明 value 0 ~ 255
func (m *LCLBrowserWindow) EnableTransparent(value uint8) {
	if m.TForm == nil {
		return
	}
	m.SetAlphaBlend(true)
	m.SetAlphaBlendValue(value)
}

//禁用最小化按钮
func (m *LCLBrowserWindow) DisableMinimize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Exclude(types.BiMinimize))
	m.WindowProperty().EnableMinimize = false
	m.EnabledMinimize(m.WindowProperty().EnableMinimize)
}

//禁用最大化按钮
func (m *LCLBrowserWindow) DisableMaximize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Exclude(types.BiMaximize))
	m.WindowProperty().EnableMaximize = false
	m.EnabledMaximize(m.WindowProperty().EnableMaximize)
}

//禁用调整窗口大小
func (m *LCLBrowserWindow) DisableResize() {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().EnableResize = false
	if !m.WindowProperty()._EnableHideCaption {
		m.TForm.SetBorderStyle(types.BsSingle)
	}
}

//禁用系统菜单-同时禁用最小化，最大化，关闭按钮
func (m *LCLBrowserWindow) DisableSystemMenu() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Exclude(types.BiSystemMenu))
	m.EnabledSystemMenu(false)
}

//禁用帮助菜单
func (m *LCLBrowserWindow) DisableHelp() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Exclude(types.BiHelp))
}

//启用最小化按钮
func (m *LCLBrowserWindow) EnableMinimize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Include(types.BiMinimize))
	m.WindowProperty().EnableMinimize = true
	m.EnabledMinimize(m.WindowProperty().EnableMinimize)
}

//启用最大化按钮
func (m *LCLBrowserWindow) EnableMaximize() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Include(types.BiMaximize))
	m.WindowProperty().EnableMaximize = true
	m.EnabledMaximize(m.WindowProperty().EnableMaximize)
}

//启用调整窗口大小
func (m *LCLBrowserWindow) EnableResize() {
	if m.TForm == nil {
		return
	}
	m.WindowProperty().EnableResize = true
	if !m.WindowProperty()._EnableHideCaption {
		m.TForm.SetBorderStyle(types.BsSizeable)
	}
}

//启用系统菜单-同时禁用最小化，最大化，关闭按钮
func (m *LCLBrowserWindow) EnableSystemMenu() {
	if m.TForm == nil {
		return
	}
	//m.SetBorderIcons(m.BorderIcons().Include(types.BiSystemMenu))
	m.EnabledSystemMenu(true)
}

//启用帮助菜单
func (m *LCLBrowserWindow) EnableHelp() {
	if m.TForm == nil {
		return
	}
	m.SetBorderIcons(m.BorderIcons().Include(types.BiHelp))
}

func (m *LCLBrowserWindow) IsViewsFramework() bool {
	return false
}

func (m *LCLBrowserWindow) IsLCL() bool {
	return true
}

func (m *LCLBrowserWindow) show(sender lcl.IObject) {
	var ret bool
	if m.onShow != nil {
		ret = m.onShow(sender)
	}
	if !ret {
		if m.windowParent != nil {
			QueueAsyncCall(func(id int) {
				m.windowParent.UpdateSize()
			})
		}
	}
}

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
		if m.windowResize != nil {
			m.windowResize(sender)
		}
		if m.chromium != nil {
			m.chromium.NotifyMoveOrResizeStarted()
		}
		if m.windowParent != nil {
			m.windowParent.UpdateSize()
		}
	}
}

func (m *LCLBrowserWindow) activate(sender lcl.IObject) {
	var ret bool
	if m.onActivate != nil {
		ret = m.onActivate(sender)
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
func (m *LCLBrowserWindow) registerPopupEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
		if !api.GoBool(BrowserWindow.Config.ChromiumConfig().enableWindowPopup) {
			return true
		}
		var bw = BrowserWindow.popupWindow.AsLCLBrowserWindow().BrowserWindow()
		bw.SetWindowType(consts.WT_POPUP_SUB_BROWSER)
		bw.ChromiumCreate(BrowserWindow.Config.ChromiumConfig(), beforePopupInfo.TargetUrl)
		bw.putChromiumWindowInfo()
		bw.defaultChromiumEvent()
		bw.registerWindowsCompMsgEvent()
		var result = false
		if bwEvent.onBeforePopup != nil {
			result = bwEvent.onBeforePopup(sender, browser, frame, beforePopupInfo, bw, noJavascriptAccess)
		}
		bw.setProperty()
		if !result {
			QueueAsyncCall(func(id int) {
				if bw.WindowProperty().IsShowModel {
					bw.ShowModal()
					return
				}
				BrowserWindow.popupWindow.Show()
			})
			result = true
		}
		return result
	})
}

// 默认事件注册 部分事件允许被覆盖
func (m *LCLBrowserWindow) registerDefaultEvent() {
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
			m.cwcap.regions = regions
			m.setDraggableRegions()
		})
	}
}

func (m *LCLBrowserWindow) close(sender lcl.IObject, action *types.TCloseAction) {
	var ret bool
	if m.onClose != nil {
		ret = m.onClose(sender, action)
	}
	if !ret {
		logger.Debug("window.onClose")
		*action = types.CaFree
	}
}

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
func (m *LCLBrowserWindow) registerDefaultChromiumCloseEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.chromium.SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *TCefCloseBrowsesAction) {
		logger.Debug("chromium.onClose")
		if IsDarwin() { //MacOSX
			desChildWind := m.windowParent.DestroyChildWindow()
			logger.Debug("chromium.onClose => windowParent.DestroyChildWindow:", desChildWind)
			*aAction = consts.CbaClose
		} else if IsLinux() {
			*aAction = consts.CbaClose
		} else if IsWindows() {
			*aAction = consts.CbaDelay
		}
		if !IsDarwin() {
			QueueAsyncCall(func(id int) { //main thread run
				m.windowParent.Free()
				logger.Debug("chromium.onClose => windowParent.Free")
			})
		}
		m.cwcap.free()
		if bwEvent.onClose != nil {
			bwEvent.onClose(sender, browser, aAction)
		}
	})
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
		logger.Debug("chromium.onBeforeClose")
		chromiumOnBeforeClose(browser)
		m.canClose = true
		var closeWindow = func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("chromium.OnBeforeClose Error:", err)
				}
			}()
			if m.auxTools != nil {
				if m.auxTools.viewSourceWindow != nil {
					m.auxTools.viewSourceWindow = nil
				}
				if m.auxTools.devToolsWindow != nil {
					m.auxTools.devToolsWindow.Close()
				}
			}
			BrowserWindow.removeWindowInfo(m.windowId)
			//主窗口关闭
			if m.WindowType() == consts.WT_MAIN_BROWSER {
				if IsWindows() {
					rtl.PostMessage(m.Handle(), WM_CLOSE, 0, 0)
				} else {
					m.Close()
				}
			} else if IsDarwin() {
				m.Close()
			}
		}
		QueueAsyncCall(func(id int) { // main thread run
			closeWindow()
		})
		if bwEvent.onBeforeClose != nil {
			bwEvent.onBeforeClose(sender, browser)
		}
	})
}
