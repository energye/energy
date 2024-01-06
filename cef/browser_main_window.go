//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy 全局窗口管理, 窗口初始化、事件注册、属性配置

package cef

import (
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/ipc/target"
	. "github.com/energye/energy/v2/cef/process"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
)

// 浏览器包装结构体
//
//  主窗口、弹出子进程、默认浏览器实现的事件、窗口维护集合、浏览器配置 仅在主进程中使用
type browserWindow struct {
	mainBrowserWindow IBrowserWindow           // 主浏览器窗口
	popupWindow       IBrowserWindow           // 弹出的子窗口
	browserEvent      *BrowserEvent            // 浏览器全局事件, 已默认实现事件
	Config            *browserConfig           // 浏览器和窗口配置
	windowInfo        map[int32]IBrowserWindow // 窗口信息集合
}

// BrowserEvent 浏览器全局事件监听-已被默认实现事件
//
//	该结构中的对象属性, 是已被默认实现的
type BrowserEvent struct {
	chromium                  IChromiumEvent                           //chromium event
	onBeforePopup             chromiumEventOnBeforePopupEx             //default
	onDragEnter               chromiumEventOnDragEnterEx               //default
	onDraggableRegionsChanged chromiumEventOnDraggableRegionsChangedEx //default
	onLoadEnd                 chromiumEventOnLoadEndEx                 //default
	onAfterCreated            chromiumEventOnAfterCreatedEx            //default
	onBeforeBrowser           chromiumEventOnBeforeBrowserEx           //default
	onBeforeClose             chromiumEventOnBeforeCloseEx             //default
	onClose                   chromiumEventOnCloseEx                   //default
	onMainFrameChanged        chromiumEventOnMainFrameChangedEx        //default
	onBeforeDownload          chromiumEventOnBeforeDownloadEx          //default can cover
	onKeyEvent                chromiumEventOnKeyEventEx                //default can cover
	onProcessMessageReceived  BrowseProcessMessageReceivedEx           //default
	onTitleChange             chromiumEventOnTitleChangeEx             //default
	onContextMenuCommand      chromiumEventOnContextMenuCommandEx      //default can cover
	onBeforeContextMenu       chromiumEventOnBeforeContextMenuEx       //default can cover
	onBeforeResourceLoad      chromiumEventOnBeforeResourceLoadEx      //default
	onRenderCompMsg           chromiumEventOnCompMsg                   //default windows
	onGetResourceHandler      chromiumEventOnGetResourceHandlerEx      //default
}

// LCLBrowserWindow
type lclBrowserWindow struct {
	LCLBrowserWindow
}

// disableMainWindow
// 禁用主窗口使用该窗口结构做为主窗口, 该窗口不被显示
type disableMainWindow struct {
	*lcl.TForm
}

// 创建LCL窗口并运行应用
func (m *browserWindow) createFormAndRun() {
	// LCL窗口
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(BrowserWindow.Config.MainFormOnTaskBar)
	if m.Config.EnableMainWindow {
		// 使用主窗口创建
		lcl.Application.CreateForm(&enableMainWindow, true)
		BrowserWindow.mainBrowserWindow = enableMainWindow
	} else {
		// 使用禁用窗口创建, 它默认不会显示, 在它的 OnFormCreate 中创建主窗口
		lcl.Application.CreateForm(&disabledMainWindow, true)
		lcl.Application.SetShowMainForm(false)
	}
	lcl.Application.Run()
}

// OnFormCreate disableMainWindow
func (m *disableMainWindow) OnFormCreate(sender lcl.IObject) {
	// 禁用主窗口后需要创建一个新的窗口来代替主窗口显示
	lcl.Application.CreateForm(&enableMainWindow, true)
	BrowserWindow.mainBrowserWindow = enableMainWindow
	// 显示窗口，此时的主窗口是默认显示的第一个窗口, 如果将该主窗口关闭，在获取主窗口函数将返回无效的窗口
	BrowserWindow.MainWindow().Show()
}

// OnFormCreate LCL窗口组件窗口创建回调
func (m *lclBrowserWindow) OnFormCreate(sender lcl.IObject) {
	m.windowProperty = &BrowserWindow.Config.WindowProperty
	m.SetWindowType(WT_MAIN_BROWSER)
	m.FormCreate()
	m.defaultWindowEvent()
	m.defaultWindowCloseEvent()
	m.ChromiumCreate(BrowserWindow.Config.ChromiumConfig(), BrowserWindow.Config.Url)
	m.defaultChromiumEvent()
	m.SetProperty()
	m.SetShowInTaskBar()
	if BrowserWindow.Config.browserWindowOnEventCallback != nil {
		BrowserWindow.browserEvent.chromium = m.Chromium()
		BrowserWindow.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, m)
		m.SetProperty() //再次设置可能修改属性
	}
	//browserWindowOnEventCallback 执行完后，注册CompMsgEvent
	m.registerWindowsCompMsgEvent()

	//自定义窗口标题栏
	m.cwcap = &customWindowCaption{
		bw: &m.LCLBrowserWindow,
	}
	//设置 CEF Chromium IPC
	ipc.SetProcessMessage(m)
	// 如果开启了开发者工具，需要在这里初始化开发者工具窗口
	if m.Chromium().Config().EnableDevTools() {
		m.createAuxTools()
		m.GetAuxTools().SetDevTools(createDevtoolsWindow(&m.LCLBrowserWindow))
	}
	if !m.WindowProperty().MainFormOnTaskBar {
		m.mainFormNotInTaskBar()
	}
}

// MainWindow 获取主浏窗口
//
//	返回LCL或VF窗口组件实例
//	Window和MacOS平台LCL窗口组件
//	Linux平台VF窗口组件
func (m *browserWindow) MainWindow() IBrowserWindow {
	return m.mainBrowserWindow
}

// SetBrowserInit 主窗口初始化时回调
//
//	LCL: 可以对主窗体属性设置、创建各种LCL子组件
//	VF : 有很大限制不能使用LCL的组件
//	event	: 浏览器事件
//	window	: 窗口信息对象
func (m *browserWindow) SetBrowserInit(fn browserWindowOnEventCallback) {
	m.Config.setBrowserWindowInitOnEvent(fn)
}

// createNextLCLPopupWindow 预创建下一个弹出的子窗口
func (m *browserWindow) createNextLCLPopupWindow() {
	if m.popupWindow == nil {
		if mw := m.MainWindow(); mw != nil {
			// owner设置nil以防址多窗口时被连带关闭, 如果多窗口时也可指定为真实的主窗口
			//var owner lcl.IComponent
			//if disabledMainWindow != nil {
			//	owner = disabledMainWindow // 真实的主窗口
			//} else if enableMainWindow != nil {
			//	owner = enableMainWindow // 未开启多窗口
			//}
			m.popupWindow = NewLCLWindow(m.Config.WindowProperty, nil)
		}
	}
}

// getNextLCLPopupWindow 拿到预创建的下一个弹出的子窗口
func (m *browserWindow) getNextLCLPopupWindow() IBrowserWindow {
	bw := m.popupWindow
	m.popupWindow = nil
	return bw
}

// GetWindowInfo 根据浏览器窗口ID获取窗口信息
func (m *browserWindow) GetWindowInfo(browserId int32) IBrowserWindow {
	if winInfo, ok := m.windowInfo[browserId]; ok {
		return winInfo
	}
	return nil
}

// GetWindowInfos 获得所有窗口信息
func (m *browserWindow) GetWindowInfos() map[int32]IBrowserWindow {
	return m.windowInfo
}

// PutWindowInfo 创建一个窗口这后我们需要添加到windowInfo中维护列表中
// Chromium 回调函数 SetOnBeforeBrowser 内设置
func (m *browserWindow) PutWindowInfo(browser *ICefBrowser, windowInfo IBrowserWindow) {
	m.windowInfo[browser.BrowserId()] = windowInfo
}

// removeWindowInfo 窗口关闭会从windowInfo移除
func (m *browserWindow) removeWindowInfo(browseId int32) {
	delete(m.windowInfo, browseId)
}

// GetBrowser 获取窗口Browser
func (m *browserWindow) GetBrowser(browseId int32) *ICefBrowser {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Browser()
	}
	return nil
}

// LookForMainWindow
// 找到一个最小的窗口ID做主下一个主窗口
func (m *browserWindow) LookForMainWindow() (window target.IWindow) {
	var (
		browseId     int32 = 0
		browseWindow IBrowserWindow
	)
	// 找到最小浏览器ID做下一个主窗口
	for bid, info := range m.GetWindowInfos() {
		if info.IsClosing() {
			// 已被关闭的窗口忽略
			continue
		} /* else if info.WindowType() == WT_MAIN_BROWSER {
			// 如果是主窗口直接返回
			browseWindow = info
			break
		}*/
		// 找到最小browserID做为主窗口
		if browseId == 0 {
			browseId = bid
			browseWindow = info
		} else if bid < browseId {
			browseId = bid
			browseWindow = info
		}
	}
	if browseWindow != nil {
		// TODO 设置为主窗口，会导致多次关闭进程结束？
		//browseWindow.SetWindowType(WT_MAIN_BROWSER)
		window = browseWindow.AsTargetWindow()
	}
	return
}

// SetOnAfterCreated
func (m *BrowserEvent) SetOnAfterCreated(event chromiumEventOnAfterCreatedEx) {
	if Args.IsMain() {
		m.onAfterCreated = event
	}
}

// SetOnBeforeBrowser
func (m *BrowserEvent) SetOnBeforeBrowser(event chromiumEventOnBeforeBrowserEx) {
	if Args.IsMain() {
		m.onBeforeBrowser = event
	}
}

// SetOnAddressChange
func (m *BrowserEvent) SetOnAddressChange(event chromiumEventOnAddressChange) {
	if Args.IsMain() {
		m.chromium.SetOnAddressChange(event)
	}
}

// SetOnBeforeClose
func (m *BrowserEvent) SetOnBeforeClose(event chromiumEventOnBeforeCloseEx) {
	if Args.IsMain() {
		m.onBeforeClose = event
	}
}

// SetOnClose
func (m *BrowserEvent) SetOnClose(event chromiumEventOnCloseEx) {
	if Args.IsMain() {
		m.onClose = event
	}
}

// SetOnPdfPrintFinished
func (m *BrowserEvent) SetOnPdfPrintFinished(event chromiumEventOnPdfPrintFinished) {
	if Args.IsMain() {
		m.chromium.SetOnPdfPrintFinished(event)
	}
}

// SetOnZoomPctAvailable
func (m *BrowserEvent) SetOnZoomPctAvailable(event chromiumEventOnResultFloat) {
	if Args.IsMain() {
		m.chromium.SetOnZoomPctAvailable(event)
	}
}

// SetOnLoadStart
func (m *BrowserEvent) SetOnLoadStart(event chromiumEventOnLoadStart) {
	if Args.IsMain() {
		m.chromium.SetOnLoadStart(event)
	}
}

// SetOnLoadingStateChange
func (m *BrowserEvent) SetOnLoadingStateChange(event chromiumEventOnLoadingStateChange) {
	if Args.IsMain() {
		m.chromium.SetOnLoadingStateChange(event)
	}
}

// SetOnLoadingProgressChange
func (m *BrowserEvent) SetOnLoadingProgressChange(event chromiumEventOnLoadingProgressChange) {
	if Args.IsMain() {
		m.chromium.SetOnLoadingProgressChange(event)
	}
}

// SetOnLoadError
func (m *BrowserEvent) SetOnLoadError(event chromiumEventOnLoadError) {
	if Args.IsMain() {
		m.chromium.SetOnLoadError(event)
	}
}

// SetOnLoadEnd
func (m *BrowserEvent) SetOnLoadEnd(event chromiumEventOnLoadEndEx) {
	if Args.IsMain() {
		m.onLoadEnd = event
	}
}

// SetOnBeforeDownload
func (m *BrowserEvent) SetOnBeforeDownload(event chromiumEventOnBeforeDownloadEx) {
	if Args.IsMain() {
		m.onBeforeDownload = event
	}
}

// SetOnDownloadUpdated
func (m *BrowserEvent) SetOnDownloadUpdated(event chromiumEventOnDownloadUpdated) {
	if Args.IsMain() {
		m.chromium.SetOnDownloadUpdated(event)
	}
}

// SetOnFullScreenModeChange
func (m *BrowserEvent) SetOnFullScreenModeChange(event chromiumEventOnFullScreenModeChange) {
	if Args.IsMain() {
		m.chromium.SetOnFullScreenModeChange(event)
	}
}

// SetOnKeyEvent
func (m *BrowserEvent) SetOnKeyEvent(event chromiumEventOnKeyEventEx) {
	if Args.IsMain() {
		m.onKeyEvent = event
	}
}

// SetOnTitleChange
func (m *BrowserEvent) SetOnTitleChange(event chromiumEventOnTitleChangeEx) {
	if Args.IsMain() {
		m.onTitleChange = event
	}
}

// SetOnRenderCompMsg windows
func (m *BrowserEvent) SetOnRenderCompMsg(event chromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.onRenderCompMsg = event
	}
}

// SetOnGetResourceHandler
//
//	获取资源处理器，通过该函数自己处理资源获取
//	返回 false 并且设置[本地|内置FS]资源加载时开启并继续执行默认实现
func (m *BrowserEvent) SetOnGetResourceHandler(event chromiumEventOnGetResourceHandlerEx) {
	if Args.IsMain() {
		if localLoadRes.enable() {
			m.onGetResourceHandler = event
		} else {
			m.ChromiumEvent().SetOnGetResourceHandler(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (resourceHandler *ICefResourceHandler) {
				resourceHandler, _ = event(sender, browser, frame, request, nil)
				return
			})
		}
	}
}

func (m *BrowserEvent) ChromiumEvent() IChromiumEvent {
	return m.chromium
}

// SetOnWidgetCompMsg windows
func (m *BrowserEvent) SetOnWidgetCompMsg(event chromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.chromium.SetOnWidgetCompMsg(event)
	}
}

// SetOnBrowserCompMsg windows
func (m *BrowserEvent) SetOnBrowserCompMsg(event chromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.chromium.SetOnBrowserCompMsg(event)
	}
}

// SetOnRenderProcessTerminated
func (m *BrowserEvent) SetOnRenderProcessTerminated(event chromiumEventOnRenderProcessTerminated) {
	if Args.IsMain() {
		m.chromium.SetOnRenderProcessTerminated(event)
	}
}

// SetOnRenderViewReady
func (m *BrowserEvent) SetOnRenderViewReady(event chromiumEventOnRenderViewReady) {
	if Args.IsMain() {
		m.chromium.SetOnRenderViewReady(event)
	}
}

// SetOnScrollOffsetChanged
func (m *BrowserEvent) SetOnScrollOffsetChanged(event chromiumEventOnScrollOffsetChanged) {
	if Args.IsMain() {
		m.chromium.SetOnScrollOffsetChanged(event)
	}
}

// SetOnBrowseProcessMessageReceived
func (m *BrowserEvent) SetOnBrowseProcessMessageReceived(event BrowseProcessMessageReceivedEx) {
	if Args.IsMain() {
		m.onProcessMessageReceived = event
	}
}

// SetOnBeforeResourceLoad
func (m *BrowserEvent) SetOnBeforeResourceLoad(event chromiumEventOnBeforeResourceLoadEx) {
	if Args.IsMain() {
		m.onBeforeResourceLoad = event
	}
}

// SetOnResourceResponse
func (m *BrowserEvent) SetOnResourceResponse(event chromiumEventOnResourceResponse) {
	if Args.IsMain() {
		m.chromium.SetOnResourceResponse(event)
	}
}

// SetOnResourceRedirect
func (m *BrowserEvent) SetOnResourceRedirect(event chromiumEventOnResourceRedirect) {
	if Args.IsMain() {
		m.chromium.SetOnResourceRedirect(event)
	}
}

// SetOnResourceLoadComplete
func (m *BrowserEvent) SetOnResourceLoadComplete(event chromiumEventOnResourceLoadComplete) {
	if Args.IsMain() {
		m.chromium.SetOnResourceLoadComplete(event)
	}
}

// SetOnCookieSet
func (m *BrowserEvent) SetOnCookieSet(event chromiumEventOnCookieSet) {
	if Args.IsMain() {
		m.chromium.SetOnCookieSet(event)
	}
}

// SetOnCookiesDeleted
func (m *BrowserEvent) SetOnCookiesDeleted(event chromiumEventOnCookiesDeleted) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesDeleted(event)
	}
}

// SetOnCookiesFlushed
func (m *BrowserEvent) SetOnCookiesFlushed(event chromiumEventOnCookiesFlushed) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesFlushed(event)
	}
}

// SetOnCookiesVisited
func (m *BrowserEvent) SetOnCookiesVisited(event chromiumEventOnCookiesVisited) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesVisited(event)
	}
}

// SetOnCookieVisitorDestroyed
func (m *BrowserEvent) SetOnCookieVisitorDestroyed(event chromiumEventOnCookieVisitorDestroyed) {
	if Args.IsMain() {
		m.chromium.SetOnCookieVisitorDestroyed(event)
	}
}

// SetOnBeforeContextMenu
func (m *BrowserEvent) SetOnBeforeContextMenu(event chromiumEventOnBeforeContextMenuEx) {
	if Args.IsMain() {
		m.onBeforeContextMenu = event
	}
}

// SetOnContextMenuCommand
func (m *BrowserEvent) SetOnContextMenuCommand(event chromiumEventOnContextMenuCommandEx) {
	if Args.IsMain() {
		m.onContextMenuCommand = event
	}
}

// SetOnContextMenuDismissed
func (m *BrowserEvent) SetOnContextMenuDismissed(event chromiumEventOnContextMenuDismissed) {
	if Args.IsMain() {
		m.chromium.SetOnContextMenuDismissed(event)
	}
}

// SetOnFrameAttached
func (m *BrowserEvent) SetOnFrameAttached(event chromiumEventOnFrameAttached) {
	if Args.IsMain() {
		m.chromium.SetOnFrameAttached(event)
	}
}

// SetOnMainFrameChanged
func (m *BrowserEvent) SetOnMainFrameChanged(event chromiumEventOnMainFrameChangedEx) {
	if Args.IsMain() {
		m.onMainFrameChanged = event
	}
}

// SetOnBeforePopup
//
//	弹出窗口, 已被默认实现的函数
//  此时的chromium还未创建, 如需要获得chromium你需要先自己创建, 可通过 ChromiumCreate 或 自定义实现 NewChromium, 此时应返回true, 否则使用默认行为Chromium应返回false
//  该事件回调会在任意线程中执行，在操作相关窗口(UI)时应在UI线程中操作 RunOnMainThread
//	函数返回值
//	  false: 窗口会以默认行为管理
//	  true: 需要你自己管理窗口行为
func (m *BrowserEvent) SetOnBeforePopup(event chromiumEventOnBeforePopupEx) {
	if Args.IsMain() {
		m.onBeforePopup = event
	}
}

// SetOnOpenUrlFromTab
//func (m *BrowserEvent) SetOnOpenUrlFromTab(event chromiumEventOnOpenUrlFromTab) {
//	if Args.IsMain() {
//		m.chromium.SetOnOpenUrlFromTab(event)
//	}
//}

// SetOnFindResult
func (m *BrowserEvent) SetOnFindResult(event chromiumEventOnFindResult) {
	if Args.IsMain() {
		m.chromium.SetOnFindResult(event)
	}
}

// SetOnDragEnter
func (m *BrowserEvent) SetOnDragEnter(event chromiumEventOnDragEnterEx) {
	if Args.IsMain() {
		m.onDragEnter = event
	}
}

// SetOnDraggableRegionsChanged
func (m *BrowserEvent) SetOnDraggableRegionsChanged(event chromiumEventOnDraggableRegionsChangedEx) {
	if Args.IsMain() {
		m.onDraggableRegionsChanged = event
	}
}
