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
	. "github.com/energye/energy/v2/cef/process"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
)

// 浏览器包装结构体
type browserWindow struct {
	mainBrowserWindow   *lclBrowserWindow            //LCL 主浏览器窗口
	mainVFBrowserWindow *ViewsFrameworkBrowserWindow //Views Frameworks 主浏览器窗口
	popupWindow         IBrowserWindow               //弹出的子窗口
	browserEvent        *BrowserEvent                //浏览器全局事件
	Config              *browserConfig               //浏览器和窗口配置
	windowInfo          map[int32]IBrowserWindow     //窗口信息集合
}

// BrowserEvent 浏览器全局事件监听-已被默认实现事件
type BrowserEvent struct {
	chromium                  IChromiumEvent                          //chromium event
	onBeforePopup             chromiumEventOnBeforePopupForWindowInfo //default
	onDragEnter               chromiumEventOnDragEnter                //default
	onDraggableRegionsChanged chromiumEventOnDraggableRegionsChanged  //default
	onLoadEnd                 chromiumEventOnLoadEnd                  //default
	onAfterCreated            chromiumEventOnAfterCreated             //default
	onBeforeBrowser           chromiumEventOnBeforeBrowser            //default
	onBeforeClose             chromiumEventOnBeforeClose              //default
	onClose                   chromiumEventOnClose                    //default
	onFrameCreated            chromiumEventOnFrameCreated             //default
	onFrameDetached           chromiumEventOnFrameDetached            //default
	onMainFrameChanged        chromiumEventOnMainFrameChanged         //default
	onBeforeDownload          chromiumEventOnBeforeDownload           //default can cover
	onKeyEvent                chromiumEventOnKeyEvent                 //default can cover
	onProcessMessageReceived  BrowseProcessMessageReceived            //default
	onTitleChange             chromiumEventOnTitleChange              //default
	onContextMenuCommand      chromiumEventOnContextMenuCommand       //default can cover
	onBeforeContextMenu       chromiumEventOnBeforeContextMenu        //default can cover
	onBeforeResourceLoad      chromiumEventOnBeforeResourceLoad       //default
	onRenderCompMsg           chromiumEventOnCompMsg                  //default windows
}

// LCLBrowserWindow
type lclBrowserWindow struct {
	LCLBrowserWindow
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
	m.setProperty()
	if BrowserWindow.Config.browserWindowOnEventCallback != nil {
		BrowserWindow.browserEvent.chromium = m.Chromium()
		BrowserWindow.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, &m.LCLBrowserWindow)
		m.setProperty() //再次设置可能修改属性
	}
	//browserWindowOnEventCallback 执行完后，注册CompMsgEvent
	m.registerWindowsCompMsgEvent()

	//自定义窗口标题栏
	m.cwcap = &customWindowCaption{
		bw: &m.LCLBrowserWindow,
	}
	//设置 CEF Chromium IPC
	ipc.SetProcessMessage(m.Chromium().(*TCEFChromium))
}

// MainWindow 获取主浏窗口
//  返回LCL或VF窗口组件实例
//  Window和MacOS平台LCL窗口组件
//  Linux平台VF窗口组件
func (m *browserWindow) MainWindow() IBrowserWindow {
	if m.mainVFBrowserWindow != nil {
		return m.mainVFBrowserWindow
	} else if m.mainBrowserWindow != nil {
		return m.mainBrowserWindow
	}
	return nil
}

// SetBrowserInit 主窗口初始化时回调
//  LCL: 可以对主窗体属性设置、创建各种LCL子组件
//  VF : 有很大限制不能使用LCL的组件
//  event	: 浏览器事件
//  window	: 窗口信息对象
func (m *browserWindow) SetBrowserInit(fn browserWindowOnEventCallback) {
	m.Config.setBrowserWindowInitOnEvent(fn)
}

// createNextLCLPopupWindow 预创建下一个弹出的子窗口
func (m *browserWindow) createNextLCLPopupWindow() {
	if m.popupWindow == nil {
		m.popupWindow = NewLCLWindow(m.Config.WindowProperty, m.MainWindow().AsLCLBrowserWindow().BrowserWindow())
		m.popupWindow.AsLCLBrowserWindow().BrowserWindow().defaultWindowCloseEvent()
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

// putWindowInfo 创建一个窗口这后会添加到windowInfo中
func (m *browserWindow) putWindowInfo(browserId int32, windowInfo IBrowserWindow) {
	m.windowInfo[browserId] = windowInfo
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

// BrowserEvent.SetOnAfterCreated
func (m *BrowserEvent) SetOnAfterCreated(event chromiumEventOnAfterCreated) {
	if Args.IsMain() {
		m.onAfterCreated = event
	}
}

// BrowserEvent.SetOnBeforeBrowser
func (m *BrowserEvent) SetOnBeforeBrowser(event chromiumEventOnBeforeBrowser) {
	if Args.IsMain() {
		m.onBeforeBrowser = event
	}
}

// BrowserEvent.SetOnAddressChange
func (m *BrowserEvent) SetOnAddressChange(event chromiumEventOnAddressChange) {
	if Args.IsMain() {
		m.chromium.SetOnAddressChange(event)
	}
}

// BrowserEvent.SetOnBeforeClose
func (m *BrowserEvent) SetOnBeforeClose(event chromiumEventOnBeforeClose) {
	if Args.IsMain() {
		m.onBeforeClose = event
	}
}

// BrowserEvent.SetOnClose
func (m *BrowserEvent) SetOnClose(event chromiumEventOnClose) {
	if Args.IsMain() {
		m.onClose = event
	}
}

// BrowserEvent.SetOnPdfPrintFinished
func (m *BrowserEvent) SetOnPdfPrintFinished(event chromiumEventOnPdfPrintFinished) {
	if Args.IsMain() {
		m.chromium.SetOnPdfPrintFinished(event)
	}
}

// BrowserEvent.SetOnZoomPctAvailable
func (m *BrowserEvent) SetOnZoomPctAvailable(event chromiumEventOnResultFloat) {
	if Args.IsMain() {
		m.chromium.SetOnZoomPctAvailable(event)
	}
}

// BrowserEvent.SetOnLoadStart
func (m *BrowserEvent) SetOnLoadStart(event chromiumEventOnLoadStart) {
	if Args.IsMain() {
		m.chromium.SetOnLoadStart(event)
	}
}

// BrowserEvent.SetOnLoadingStateChange
func (m *BrowserEvent) SetOnLoadingStateChange(event chromiumEventOnLoadingStateChange) {
	if Args.IsMain() {
		m.chromium.SetOnLoadingStateChange(event)
	}
}

// BrowserEvent.SetOnLoadingProgressChange
func (m *BrowserEvent) SetOnLoadingProgressChange(event chromiumEventOnLoadingProgressChange) {
	if Args.IsMain() {
		m.chromium.SetOnLoadingProgressChange(event)
	}
}

// BrowserEvent.SetOnLoadError
func (m *BrowserEvent) SetOnLoadError(event chromiumEventOnLoadError) {
	if Args.IsMain() {
		m.chromium.SetOnLoadError(event)
	}
}

// BrowserEvent.SetOnLoadEnd
func (m *BrowserEvent) SetOnLoadEnd(event chromiumEventOnLoadEnd) {
	if Args.IsMain() {
		m.onLoadEnd = event
	}
}

// BrowserEvent.SetOnBeforeDownload
func (m *BrowserEvent) SetOnBeforeDownload(event chromiumEventOnBeforeDownload) {
	if Args.IsMain() {
		m.onBeforeDownload = event
	}
}

// BrowserEvent.SetOnDownloadUpdated
func (m *BrowserEvent) SetOnDownloadUpdated(event chromiumEventOnDownloadUpdated) {
	if Args.IsMain() {
		m.chromium.SetOnDownloadUpdated(event)
	}
}

// BrowserEvent.SetOnFullScreenModeChange
func (m *BrowserEvent) SetOnFullScreenModeChange(event chromiumEventOnFullScreenModeChange) {
	if Args.IsMain() {
		m.chromium.SetOnFullScreenModeChange(event)
	}
}

// BrowserEvent.SetOnKeyEvent
func (m *BrowserEvent) SetOnKeyEvent(event chromiumEventOnKeyEvent) {
	if Args.IsMain() {
		m.onKeyEvent = event
	}
}

// BrowserEvent.SetOnTitleChange
func (m *BrowserEvent) SetOnTitleChange(event chromiumEventOnTitleChange) {
	if Args.IsMain() {
		m.onTitleChange = event
	}
}

// BrowserEvent.SetOnRenderCompMsg windows
func (m *BrowserEvent) SetOnRenderCompMsg(event chromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.onRenderCompMsg = event
	}
}

// BrowserEvent.SetOnWidgetCompMsg windows
func (m *BrowserEvent) SetOnWidgetCompMsg(event chromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.chromium.SetOnWidgetCompMsg(event)
	}
}

// BrowserEvent.SetOnBrowserCompMsg windows
func (m *BrowserEvent) SetOnBrowserCompMsg(event chromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.chromium.SetOnBrowserCompMsg(event)
	}
}

// BrowserEvent.SetOnRenderProcessTerminated
func (m *BrowserEvent) SetOnRenderProcessTerminated(event chromiumEventOnRenderProcessTerminated) {
	if Args.IsMain() {
		m.chromium.SetOnRenderProcessTerminated(event)
	}
}

// BrowserEvent.SetOnRenderViewReady
func (m *BrowserEvent) SetOnRenderViewReady(event chromiumEventOnRenderViewReady) {
	if Args.IsMain() {
		m.chromium.SetOnRenderViewReady(event)
	}
}

// BrowserEvent.SetOnScrollOffsetChanged
func (m *BrowserEvent) SetOnScrollOffsetChanged(event chromiumEventOnScrollOffsetChanged) {
	if Args.IsMain() {
		m.chromium.SetOnScrollOffsetChanged(event)
	}
}

// BrowserEvent.SetOnBrowseProcessMessageReceived
func (m *BrowserEvent) SetOnBrowseProcessMessageReceived(event BrowseProcessMessageReceived) {
	if Args.IsMain() {
		m.onProcessMessageReceived = event
	}
}

// BrowserEvent.SetOnBeforeResourceLoad
func (m *BrowserEvent) SetOnBeforeResourceLoad(event chromiumEventOnBeforeResourceLoad) {
	if Args.IsMain() {
		m.onBeforeResourceLoad = event
	}
}

// BrowserEvent.SetOnResourceResponse
func (m *BrowserEvent) SetOnResourceResponse(event chromiumEventOnResourceResponse) {
	if Args.IsMain() {
		m.chromium.SetOnResourceResponse(event)
	}
}

// BrowserEvent.SetOnResourceRedirect
func (m *BrowserEvent) SetOnResourceRedirect(event chromiumEventOnResourceRedirect) {
	if Args.IsMain() {
		m.chromium.SetOnResourceRedirect(event)
	}
}

// BrowserEvent.SetOnResourceLoadComplete
func (m *BrowserEvent) SetOnResourceLoadComplete(event chromiumEventOnResourceLoadComplete) {
	if Args.IsMain() {
		m.chromium.SetOnResourceLoadComplete(event)
	}
}

// BrowserEvent.SetOnCookieSet
func (m *BrowserEvent) SetOnCookieSet(event chromiumEventOnCookieSet) {
	if Args.IsMain() {
		m.chromium.SetOnCookieSet(event)
	}
}

// BrowserEvent.SetOnCookiesDeleted
func (m *BrowserEvent) SetOnCookiesDeleted(event chromiumEventOnCookiesDeleted) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesDeleted(event)
	}
}

// BrowserEvent.SetOnCookiesFlushed
func (m *BrowserEvent) SetOnCookiesFlushed(event chromiumEventOnCookiesFlushed) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesFlushed(event)
	}
}

// BrowserEvent.SetOnCookiesVisited
func (m *BrowserEvent) SetOnCookiesVisited(event chromiumEventOnCookiesVisited) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesVisited(event)
	}
}

// BrowserEvent.SetOnCookieVisitorDestroyed
func (m *BrowserEvent) SetOnCookieVisitorDestroyed(event chromiumEventOnCookieVisitorDestroyed) {
	if Args.IsMain() {
		m.chromium.SetOnCookieVisitorDestroyed(event)
	}
}

// BrowserEvent.SetOnBeforeContextMenu
func (m *BrowserEvent) SetOnBeforeContextMenu(event chromiumEventOnBeforeContextMenu) {
	if Args.IsMain() {
		m.onBeforeContextMenu = event
	}
}

// BrowserEvent.SetOnContextMenuCommand
func (m *BrowserEvent) SetOnContextMenuCommand(event chromiumEventOnContextMenuCommand) {
	if Args.IsMain() {
		m.onContextMenuCommand = event
	}
}

// BrowserEvent.SetOnContextMenuDismissed
func (m *BrowserEvent) SetOnContextMenuDismissed(event chromiumEventOnContextMenuDismissed) {
	if Args.IsMain() {
		m.chromium.SetOnContextMenuDismissed(event)
	}
}

// BrowserEvent.SetOnFrameAttached
func (m *BrowserEvent) SetOnFrameAttached(event chromiumEventOnFrameAttached) {
	if Args.IsMain() {
		m.chromium.SetOnFrameAttached(event)
	}
}

// BrowserEvent.SetOnFrameCreated
func (m *BrowserEvent) SetOnFrameCreated(event chromiumEventOnFrameCreated) {
	if Args.IsMain() {
		m.onFrameCreated = event
	}
}

// BrowserEvent.SetOnFrameDetached
func (m *BrowserEvent) SetOnFrameDetached(event chromiumEventOnFrameDetached) {
	if Args.IsMain() {
		m.onFrameDetached = event
	}
}

// BrowserEvent.SetOnMainFrameChanged
func (m *BrowserEvent) SetOnMainFrameChanged(event chromiumEventOnMainFrameChanged) {
	if Args.IsMain() {
		m.onMainFrameChanged = event
	}
}

// BrowserEvent.SetOnBeforePopup
func (m *BrowserEvent) SetOnBeforePopup(event chromiumEventOnBeforePopupForWindowInfo) {
	if Args.IsMain() {
		m.onBeforePopup = event
	}
}

// BrowserEvent.SetOnOpenUrlFromTab
//func (m *BrowserEvent) SetOnOpenUrlFromTab(event chromiumEventOnOpenUrlFromTab) {
//	if Args.IsMain() {
//		m.chromium.SetOnOpenUrlFromTab(event)
//	}
//}

// BrowserEvent.SetOnFindResult
func (m *BrowserEvent) SetOnFindResult(event chromiumEventOnFindResult) {
	if Args.IsMain() {
		m.chromium.SetOnFindResult(event)
	}
}

// BrowserEvent.SetOnDragEnter
func (m *BrowserEvent) SetOnDragEnter(event chromiumEventOnDragEnter) {
	if Args.IsMain() {
		m.onDragEnter = event
	}
}

// BrowserEvent.SetOnDraggableRegionsChanged
func (m *BrowserEvent) SetOnDraggableRegionsChanged(event chromiumEventOnDraggableRegionsChanged) {
	if Args.IsMain() {
		m.onDraggableRegionsChanged = event
	}
}
