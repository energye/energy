//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// 浏览器包装结构体
type browser struct {
	mainBrowserWindow   *browserWindow               //LCL 主浏览器窗口
	mainVFBrowserWindow *ViewsFrameworkBrowserWindow //Views Frameworks 主浏览器窗口
	popupWindow         IBrowserWindow               //弹出的子窗口
	browserEvent        *BrowserEvent                //浏览器全局事件
	Config              *browserConfig               //浏览器和窗口配置
	windowInfo          map[int32]IBrowserWindow     //窗口信息集合
	windowSerial        int32                        //窗口序号
}

// 浏览器全局事件监听-扩展
//
// 在主浏览器进程调用
type BrowserEvent struct {
	chromium                  IChromiumEvent                          //chromium event
	onBeforePopup             ChromiumEventOnBeforePopupForWindowInfo //default
	onDragEnter               ChromiumEventOnDragEnter                //default
	onDraggableRegionsChanged ChromiumEventOnDraggableRegionsChanged  //default
	onLoadEnd                 ChromiumEventOnLoadEnd                  //default
	onAfterCreated            ChromiumEventOnAfterCreated             //default
	onBeforeBrowser           ChromiumEventOnBeforeBrowser            //default
	onBeforeClose             ChromiumEventOnBeforeClose              //default
	onClose                   ChromiumEventOnClose                    //default
	onFrameCreated            ChromiumEventOnFrameCreated             //default
	onFrameDetached           ChromiumEventOnFrameDetached            //default
	onMainFrameChanged        ChromiumEventOnMainFrameChanged         //default
	onBeforeDownload          ChromiumEventOnBeforeDownload           //default can cover
	onKeyEvent                ChromiumEventOnKeyEvent                 //default can cover
	onProcessMessageReceived  BrowseProcessMessageReceived            //default
	onTitleChange             ChromiumEventOnTitleChange              //default
	onContextMenuCommand      ChromiumEventOnContextMenuCommand       //default can cover
	onBeforeContextMenu       ChromiumEventOnBeforeContextMenu        //default can cover
	onBeforeResourceLoad      ChromiumEventOnBeforeResourceLoad       //default
	onRenderCompMsg           ChromiumEventOnCompMsg                  //default windows
}

type browserWindow struct {
	LCLBrowserWindow
	isFirstActivate bool
}

func (m *browserWindow) OnFormCreate(sender lcl.IObject) {
	m.windowProperty = BrowserWindow.Config.WindowProperty
	m.SetWindowType(WT_MAIN_BROWSER)
	m.FormCreate()
	m.defaultWindowEvent()
	m.defaultWindowCloseEvent()
	m.ChromiumCreate(BrowserWindow.Config.ChromiumConfig(), BrowserWindow.Config.Url)
	m.putChromiumWindowInfo()
	m.defaultChromiumEvent()
	m.SetCaption(m.windowProperty.Title)
	if m.windowProperty.IconFS != "" {
		if emfs.IsExist(m.windowProperty.IconFS) {
			_ = lcl.Application.Icon().LoadFromFSFile(m.windowProperty.IconFS)
		}
	} else if m.windowProperty.Icon != "" {
		if tools.IsExist(m.windowProperty.Icon) {
			lcl.Application.Icon().LoadFromFile(m.windowProperty.Icon)
		}
	}
	if m.windowProperty.CenterWindow {
		m.SetWidth(m.windowProperty.Width)
		m.SetHeight(m.windowProperty.Height)
		m.SetPosition(types.PoDesktopCenter)
	} else {
		m.SetPosition(types.PoDesigned)
		m.SetBounds(m.windowProperty.X, m.windowProperty.Y, m.windowProperty.Width, m.windowProperty.Height)
	}
	if m.windowProperty.AlwaysOnTop {
		m.SetFormStyle(types.FsSystemStayOnTop)
	}
	m.EnabledMinimize(m.windowProperty.CanMinimize)
	m.EnabledMaximize(m.windowProperty.CanMaximize)
	if !m.windowProperty.CanResize {
		m.SetBorderStyle(types.BsSingle)
	}
	if BrowserWindow.Config.browserWindowOnEventCallback != nil {
		BrowserWindow.browserEvent.chromium = m.chromium
		BrowserWindow.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, &m.LCLBrowserWindow)
	}
	//browserWindowOnEventCallback 执行完后，注册CompMsgEvent
	m.registerWindowsCompMsgEvent()

	//主进程（主窗口）启动后回调函数事件
	//主窗体第一次激活之后执行一次
	m.SetOnActivateAfter(func(sender lcl.IObject) {
		if !m.isFirstActivate {
			m.isFirstActivate = true
			if BrowserWindow.Config.browserWindowAfterOnEventCallback != nil {
				BrowserWindow.Config.browserWindowAfterOnEventCallback(&m.LCLBrowserWindow)
			}
		}
	})
}

func (m *browser) MainWindow() IBrowserWindow {
	if m.mainVFBrowserWindow != nil {
		return m.mainVFBrowserWindow
	} else if m.mainBrowserWindow != nil {
		return m.mainBrowserWindow
	}
	return nil
	//if IsMessageLoop {
	//	return m.mainVFBrowserWindow.BrowserWindow()
	//}
	//return m.mainBrowserWindow.BrowserWindow()
}

// 主窗口和chromium初始化时回调
//
// event 			浏览器事件
//
// browserWindow 	窗口信息对象
func (m *browser) SetBrowserInit(fn browserWindowOnEventCallback) {
	m.Config.setBrowserWindowInitOnEvent(fn)
}

// 基于LCL窗口 - 主窗体和chromium初始后回调
//
// 在这里可以对主窗体属性设置、添加子窗口、带有browser的窗口和子组件创建
//
// mainBrowserWindow 窗口信息对象
func (m *browser) SetBrowserInitAfter(fn browserWindowAfterOnEventCallback) {
	m.Config.setBrowserWindowInitAfterOnEvent(fn)
}

// 设置或增加一个窗口序号
func (m *browser) setOrIncNextWindowNum(browserId ...int32) int32 {
	if len(browserId) > 0 {
		m.windowSerial = browserId[0]
	} else {
		m.windowSerial++
	}
	logger.Debug("next window serial:", m.windowSerial)
	return m.windowSerial
}

// 设置或减少一个窗口序号
func (m *browser) setOrDecNextWindowNum(browserId ...int32) int32 {
	if len(browserId) > 0 {
		m.windowSerial = browserId[0]
	} else {
		m.windowSerial--
	}
	return m.windowSerial
}

// 获得窗口序号
func (m *browser) GetNextWindowNum() int32 {
	return m.windowSerial
}

func (m *browser) createNextLCLPopupWindow() {
	m.popupWindow = NewLCLWindow(BrowserWindow.Config.WindowProperty, m.MainWindow().AsLCLBrowserWindow().BrowserWindow())
	m.popupWindow.AsLCLBrowserWindow().BrowserWindow().defaultWindowCloseEvent()
}

// 拿到窗口信息
func (m *browser) GetWindowInfo(browserId int32) IBrowserWindow {
	if winInfo, ok := m.windowInfo[browserId]; ok {
		return winInfo
	}
	return nil
}

func (m *browser) GetWindowsInfo() map[int32]IBrowserWindow {
	return m.windowInfo
}

func (m *browser) putWindowInfo(browserId int32, windowInfo IBrowserWindow) {
	m.windowInfo[browserId] = windowInfo
}

func (m *browser) removeWindowInfo(browseId int32) {
	delete(m.windowInfo, browseId)
	RemoveGoForm(browseId)
}

func (m *browser) GetBrowser(browseId int32) *ICefBrowser {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Browser()
	}
	return nil
}

func (m *browser) putBrowserFrame(browser *ICefBrowser, frame *ICefFrame) {
	if winInfo, ok := m.windowInfo[browser.Identifier()]; ok {
		winInfo.setBrowser(browser)
		winInfo.addFrame(frame)
	}
}

func (m *browser) GetFrames(browseId int32) map[int64]*ICefFrame {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Frames()
	}
	return nil
}

func (m *browser) GetFrame(browseId int32, frameId int64) *ICefFrame {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Frames()[frameId]
	}
	return nil
}

func (m *browser) RemoveFrame(browseId int32, frameId int64) {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		delete(winInfo.Frames(), frameId)
	}
}

func (m *browser) IsSameFrame(browseId int32, frameId int64) bool {
	if frame := m.GetFrame(browseId, frameId); frame != nil {
		return true
	}
	return false
}

func (m *browser) removeNoValidFrames() {
	for _, winInfo := range m.windowInfo {
		for _, frm := range winInfo.Frames() {
			if !frm.IsValid() {
				delete(winInfo.Frames(), frm.Id)
			}
		}
	}
}

// BrowserEvent.SetOnAfterCreated
func (m *BrowserEvent) SetOnAfterCreated(event ChromiumEventOnAfterCreated) {
	if Args.IsMain() {
		m.onAfterCreated = event
	}
}

// BrowserEvent.SetOnBeforeBrowser
func (m *BrowserEvent) SetOnBeforeBrowser(event ChromiumEventOnBeforeBrowser) {
	if Args.IsMain() {
		m.onBeforeBrowser = event
	}
}

// BrowserEvent.SetOnAddressChange
func (m *BrowserEvent) SetOnAddressChange(event ChromiumEventOnAddressChange) {
	if Args.IsMain() {
		m.chromium.SetOnAddressChange(event)
	}
}

// BrowserEvent.SetOnBeforeClose
func (m *BrowserEvent) SetOnBeforeClose(event ChromiumEventOnBeforeClose) {
	if Args.IsMain() {
		m.onBeforeClose = event
	}
}

// BrowserEvent.SetOnClose
func (m *BrowserEvent) SetOnClose(event ChromiumEventOnClose) {
	if Args.IsMain() {
		m.onClose = event
	}
}

// BrowserEvent.SetOnPdfPrintFinished
func (m *BrowserEvent) SetOnPdfPrintFinished(event ChromiumEventOnResult) {
	if Args.IsMain() {
		m.chromium.SetOnPdfPrintFinished(event)
	}
}

// BrowserEvent.SetOnZoomPctAvailable
func (m *BrowserEvent) SetOnZoomPctAvailable(event ChromiumEventOnResultFloat) {
	if Args.IsMain() {
		m.chromium.SetOnZoomPctAvailable(event)
	}
}

// BrowserEvent.SetOnLoadStart
func (m *BrowserEvent) SetOnLoadStart(event ChromiumEventOnLoadStart) {
	if Args.IsMain() {
		m.chromium.SetOnLoadStart(event)
	}
}

// BrowserEvent.SetOnLoadingStateChange
func (m *BrowserEvent) SetOnLoadingStateChange(event ChromiumEventOnLoadingStateChange) {
	if Args.IsMain() {
		m.chromium.SetOnLoadingStateChange(event)
	}
}

// BrowserEvent.SetOnLoadingProgressChange
func (m *BrowserEvent) SetOnLoadingProgressChange(event ChromiumEventOnLoadingProgressChange) {
	if Args.IsMain() {
		m.chromium.SetOnLoadingProgressChange(event)
	}
}

// BrowserEvent.SetOnLoadError
func (m *BrowserEvent) SetOnLoadError(event ChromiumEventOnLoadError) {
	if Args.IsMain() {
		m.chromium.SetOnLoadError(event)
	}
}

// BrowserEvent.SetOnLoadEnd
func (m *BrowserEvent) SetOnLoadEnd(event ChromiumEventOnLoadEnd) {
	if Args.IsMain() {
		m.onLoadEnd = event
	}
}

// BrowserEvent.SetOnBeforeDownload
func (m *BrowserEvent) SetOnBeforeDownload(event ChromiumEventOnBeforeDownload) {
	if Args.IsMain() {
		m.onBeforeDownload = event
	}
}

// BrowserEvent.SetOnDownloadUpdated
func (m *BrowserEvent) SetOnDownloadUpdated(event ChromiumEventOnDownloadUpdated) {
	if Args.IsMain() {
		m.chromium.SetOnDownloadUpdated(event)
	}
}

// BrowserEvent.SetOnFullScreenModeChange
func (m *BrowserEvent) SetOnFullScreenModeChange(event ChromiumEventOnFullScreenModeChange) {
	if Args.IsMain() {
		m.chromium.SetOnFullScreenModeChange(event)
	}
}

// BrowserEvent.SetOnKeyEvent
func (m *BrowserEvent) SetOnKeyEvent(event ChromiumEventOnKeyEvent) {
	if Args.IsMain() {
		m.onKeyEvent = event
	}
}

// BrowserEvent.SetOnTitleChange
func (m *BrowserEvent) SetOnTitleChange(event ChromiumEventOnTitleChange) {
	if Args.IsMain() {
		m.onTitleChange = event
	}
}

// BrowserEvent.SetOnRenderCompMsg windows
func (m *BrowserEvent) SetOnRenderCompMsg(event ChromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.onRenderCompMsg = event
	}
}

// BrowserEvent.SetOnWidgetCompMsg windows
func (m *BrowserEvent) SetOnWidgetCompMsg(event ChromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.chromium.SetOnWidgetCompMsg(event)
	}
}

// BrowserEvent.SetOnBrowserCompMsg windows
func (m *BrowserEvent) SetOnBrowserCompMsg(event ChromiumEventOnCompMsg) {
	if Args.IsMain() {
		m.chromium.SetOnBrowserCompMsg(event)
	}
}

// BrowserEvent.SetOnRenderProcessTerminated
func (m *BrowserEvent) SetOnRenderProcessTerminated(event ChromiumEventOnRenderProcessTerminated) {
	if Args.IsMain() {
		m.chromium.SetOnRenderProcessTerminated(event)
	}
}

// BrowserEvent.SetOnRenderViewReady
func (m *BrowserEvent) SetOnRenderViewReady(event ChromiumEventOnCefBrowser) {
	if Args.IsMain() {
		m.chromium.SetOnRenderViewReady(event)
	}
}

// BrowserEvent.SetOnScrollOffsetChanged
func (m *BrowserEvent) SetOnScrollOffsetChanged(event ChromiumEventOnScrollOffsetChanged) {
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
func (m *BrowserEvent) SetOnBeforeResourceLoad(event ChromiumEventOnBeforeResourceLoad) {
	if Args.IsMain() {
		m.onBeforeResourceLoad = event
	}
}

// BrowserEvent.SetOnResourceResponse
func (m *BrowserEvent) SetOnResourceResponse(event ChromiumEventOnResourceResponse) {
	if Args.IsMain() {
		m.chromium.SetOnResourceResponse(event)
	}
}

// BrowserEvent.SetOnResourceRedirect
func (m *BrowserEvent) SetOnResourceRedirect(event ChromiumEventOnResourceRedirect) {
	if Args.IsMain() {
		m.chromium.SetOnResourceRedirect(event)
	}
}

// BrowserEvent.SetOnResourceLoadComplete
func (m *BrowserEvent) SetOnResourceLoadComplete(event ChromiumEventOnResourceLoadComplete) {
	if Args.IsMain() {
		m.chromium.SetOnResourceLoadComplete(event)
	}
}

// BrowserEvent.SetOnCookieSet
func (m *BrowserEvent) SetOnCookieSet(event ChromiumEventOnCookieSet) {
	if Args.IsMain() {
		m.chromium.SetOnCookieSet(event)
	}
}

// BrowserEvent.SetOnCookiesDeleted
func (m *BrowserEvent) SetOnCookiesDeleted(event ChromiumEventOnCookiesDeleted) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesDeleted(event)
	}
}

// BrowserEvent.SetOnCookiesFlushed
func (m *BrowserEvent) SetOnCookiesFlushed(event ChromiumEventOnCookiesFlushed) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesFlushed(event)
	}
}

// BrowserEvent.SetOnCookiesVisited
func (m *BrowserEvent) SetOnCookiesVisited(event ChromiumEventOnCookiesVisited) {
	if Args.IsMain() {
		m.chromium.SetOnCookiesVisited(event)
	}
}

// BrowserEvent.SetOnCookieVisitorDestroyed
func (m *BrowserEvent) SetOnCookieVisitorDestroyed(event ChromiumEventOnCookieVisitorDestroyed) {
	if Args.IsMain() {
		m.chromium.SetOnCookieVisitorDestroyed(event)
	}
}

// BrowserEvent.SetOnBeforeContextMenu
func (m *BrowserEvent) SetOnBeforeContextMenu(event ChromiumEventOnBeforeContextMenu) {
	if Args.IsMain() {
		m.onBeforeContextMenu = event
	}
}

// BrowserEvent.SetOnContextMenuCommand
func (m *BrowserEvent) SetOnContextMenuCommand(event ChromiumEventOnContextMenuCommand) {
	if Args.IsMain() {
		m.onContextMenuCommand = event
	}
}

// BrowserEvent.SetOnContextMenuDismissed
func (m *BrowserEvent) SetOnContextMenuDismissed(event ChromiumEventOnContextMenuDismissed) {
	if Args.IsMain() {
		m.chromium.SetOnContextMenuDismissed(event)
	}
}

// BrowserEvent.SetOnFrameAttached
func (m *BrowserEvent) SetOnFrameAttached(event ChromiumEventOnFrameAttached) {
	if Args.IsMain() {
		m.chromium.SetOnFrameAttached(event)
	}
}

// BrowserEvent.SetOnFrameCreated
func (m *BrowserEvent) SetOnFrameCreated(event ChromiumEventOnFrameCreated) {
	if Args.IsMain() {
		m.onFrameCreated = event
	}
}

// BrowserEvent.SetOnFrameDetached
func (m *BrowserEvent) SetOnFrameDetached(event ChromiumEventOnFrameDetached) {
	if Args.IsMain() {
		m.onFrameDetached = event
	}
}

// BrowserEvent.SetOnMainFrameChanged
func (m *BrowserEvent) SetOnMainFrameChanged(event ChromiumEventOnMainFrameChanged) {
	if Args.IsMain() {
		m.onMainFrameChanged = event
	}
}

// BrowserEvent.SetOnBeforePopup
//
// 子窗口弹出之前，设置子窗口样式及系统组件和功能
func (m *BrowserEvent) SetOnBeforePopup(event ChromiumEventOnBeforePopupForWindowInfo) {
	if Args.IsMain() {
		m.onBeforePopup = event
	}
}

// BrowserEvent.SetOnOpenUrlFromTab
//func (m *BrowserEvent) SetOnOpenUrlFromTab(event ChromiumEventOnOpenUrlFromTab) {
//	if Args.IsMain() {
//		m.chromium.SetOnOpenUrlFromTab(event)
//	}
//}

// BrowserEvent.SetOnFindResult
func (m *BrowserEvent) SetOnFindResult(event ChromiumEventOnFindResult) {
	if Args.IsMain() {
		m.chromium.SetOnFindResult(event)
	}
}

// BrowserEvent.SetOnDragEnter
func (m *BrowserEvent) SetOnDragEnter(event ChromiumEventOnDragEnter) {
	if Args.IsMain() {
		m.onDragEnter = event
	}
}

// BrowserEvent.SetOnDraggableRegionsChanged
func (m *BrowserEvent) SetOnDraggableRegionsChanged(event ChromiumEventOnDraggableRegionsChanged) {
	if Args.IsMain() {
		m.onDraggableRegionsChanged = event
	}
}
