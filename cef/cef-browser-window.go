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
	"github.com/energye/golcl/lcl"
	"sync"
)

var (
	//1. BrowserWindow 是基于 BaseWindow 浏览器主窗口
	//
	//2. 可以对窗口的属性设置和事件监听，chromium 的配置和事件监听
	//
	//3. 该窗口是主窗体，因此初始化时必须第一个初始化完成，如果创建子窗口最好在 SetBrowserInitAfter 回调函数中创建
	//你也可以在 SetBrowserInit 回调函数初始化时创建子窗口，但是不能在主窗体显示之前显示子窗口. 尤其是带有chromium的窗口.
	BrowserWindow = &browser{
		mainBrowserWindow: &browserWindow{},
		browserEvent:      &BrowserEvent{},
		Config: &browserConfig{
			Title:  "Energy",
			Width:  1024,
			Height: 768,
		},
		windowInfo:   make(map[int32]*TCefWindowInfo),
		windowSerial: 1,
		uiLock:       new(sync.Mutex),
	}
	browserProcessStartAfterCallback browserProcessStartAfterCallbackFunc
)

type browserProcessStartAfterCallbackFunc func(b bool)

// SetBrowserProcessStartAfterCallback 主进程启动之后回调函数
func SetBrowserProcessStartAfterCallback(callback browserProcessStartAfterCallbackFunc) {
	if Args.IsMain() {
		if browserProcessStartAfterCallback == nil {
			browserProcessStartAfterCallback = callback
		}
	}
}

// 浏览器包装结构体
type browser struct {
	mainBrowserWindow *browserWindow            //主浏览器窗口
	mainWindow        *TCefWindowInfo           //主窗口信息
	popupWindow       *browserWindow            //弹出的子窗口
	browserEvent      *BrowserEvent             //浏览器全局事件
	Config            *browserConfig            //浏览器和窗口配置
	windowInfo        map[int32]*TCefWindowInfo //窗口信息集合
	windowSerial      int32                     //窗口序号
	uiLock            *sync.Mutex
}

// 浏览器全局事件监听
//
// 在主浏览器进程调用
type BrowserEvent struct {
	chromium                 IChromiumEvent                    //chromium event
	onBeforePopup            ChromiumEventOnBeforePopup        //default
	onAfterCreated           ChromiumEventOnAfterCreated       //default
	onBeforeBrowser          ChromiumEventOnBeforeBrowser      //default
	onBeforeClose            ChromiumEventOnBeforeClose        //default
	onClose                  ChromiumEventOnClose              //default
	onFrameCreated           ChromiumEventOnFrameCreated       //default
	onFrameDetached          ChromiumEventOnFrameDetached      //default
	onMainFrameChanged       ChromiumEventOnMainFrameChanged   //default
	onBeforeDownload         ChromiumEventOnBeforeDownload     //default
	onKeyEvent               ChromiumEventOnKeyEvent           //default
	onProcessMessageReceived BrowseProcessMessageReceived      //default
	onTitleChange            ChromiumEventOnTitleChange        //default
	onLoadingStateChange     ChromiumEventOnLoadingStateChange //default
	onContextMenuCommand     ChromiumEventOnContextMenuCommand //default
	onBeforeContextMenu      ChromiumEventOnBeforeContextMenu  //default
	onBeforeResourceLoad     ChromiumEventOnBeforeResourceLoad //default
}

type browserWindow struct {
	BaseWindow
	isFirstActivate bool
	tray            ITray
}

// 运行应用
//
// 多进程方式，启动主进程然后启动子进程，在MacOS下，需要单独调用启动子进程函数，单进程只启动主进程
//
// 主进程启动成功之后，将创建主窗口 mainBrowserWindow 是一个默认的主窗口
//
// 在这里启动浏览器的主进程和子进程
func Run(cefApp *TCEFApplication) {
	if IsDarwin() && !SingleProcess && !Args.IsMain() {
		// mac os 启动子进程
		cefApp.StartSubProcess()
	} else {
		b := cefApp.StartMainProcess()
		if browserProcessStartAfterCallback != nil {
			browserProcessStartAfterCallback(b)
		}
		if b {
			lcl.RunApp(&BrowserWindow.mainBrowserWindow)
		}
	}
}

func (m *browserWindow) OnFormCreate(sender lcl.IObject) {
	if BrowserWindow.Config.chromiumConfig == nil {
		BrowserWindow.Config.chromiumConfig = NewChromiumConfig()
		BrowserWindow.Config.chromiumConfig.SetEnableMenu(true)
		BrowserWindow.Config.chromiumConfig.SetEnableDevTools(true)
		BrowserWindow.Config.chromiumConfig.SetEnableOpenUrlTab(true)
		BrowserWindow.Config.chromiumConfig.SetEnableWindowPopup(true)
	}
	m.SetWindowType(WT_MAIN_BROWSER)
	m.FormCreate()
	m.defaultWindowEvent()
	m.defaultWindowCloseEvent()
	m.ChromiumCreate(BrowserWindow.Config.chromiumConfig, BrowserWindow.Config.DefaultUrl)
	m.putChromiumWindowInfo()
	m.defaultChromiumEvent()
	BrowserWindow.mainWindow = m.windowInfo
	m.AddOnCloseQuery(func(sender lcl.IObject, canClose *bool) bool {
		if m.tray != nil {
			m.tray.close()
		}
		return false
	})
	if BrowserWindow.Config.Title != "" {
		m.SetCaption(BrowserWindow.Config.Title)
	}
	if BrowserWindow.Config.Icon != "" {
		lcl.Application.Icon().LoadFromFSFile(BrowserWindow.Config.Icon)
	}
	m.SetWidth(BrowserWindow.Config.Width)
	m.SetHeight(BrowserWindow.Config.Height)
	if BrowserWindow.Config.browserWindowOnEventCallback != nil {
		BrowserWindow.browserEvent.chromium = m.chromium
		BrowserWindow.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, m.windowInfo)
	}

	//主进程（主窗口）启动后回调函数事件
	//主窗体第一次激活之后执行一次
	m.SetOnActivateAfter(func(sender lcl.IObject) {
		if !m.isFirstActivate {
			m.isFirstActivate = true
			if BrowserWindow.Config.browserWindowAfterOnEventCallback != nil {
				BrowserWindow.Config.browserWindowAfterOnEventCallback(m.windowInfo)
			}
		}
	})
}

func (m *browser) MainWindow() *TCefWindowInfo {
	return m.mainWindow
}

// 主窗口和chromium初始化时回调
//
// 在这里可以对主窗体事件监听和属性设置,和主窗口上的子组件创建
//
// 如果想创建子窗口或带有browser的窗口最好在 SetBrowserInitAfter 回调函数中创建
//
// event 浏览器事件
// mainBrowserWindow 窗口信息对象
func (m *browser) SetBrowserInit(fn func(event *BrowserEvent, browserWindow *TCefWindowInfo)) {
	m.Config.setBrowserWindowInitOnEvent(fn)
}

// 主窗体和chromium初始后回调
//
// 在这里可以对主窗体属性设置、添加子窗口、带有browser的窗口和子组件创建
//
// mainBrowserWindow 窗口信息对象
func (m *browser) SetBrowserInitAfter(fn func(browserWindow *TCefWindowInfo)) {
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

func (m *browser) createNextPopupWindow() {
	m.popupWindow = &browserWindow{}
	m.popupWindow.TForm = lcl.NewForm(m.MainWindow().Window)
	m.popupWindow.FormCreate()
	m.popupWindow.defaultWindowEvent()
	m.popupWindow.defaultWindowCloseEvent()
}

// 拿到窗口信息
func (m *browser) GetWindowInfo(browserId int32) *TCefWindowInfo {
	if winInfo, ok := m.windowInfo[browserId]; ok {
		return winInfo
	}
	return nil
}

func (m *browser) GetWindowsInfo() map[int32]*TCefWindowInfo {
	return m.windowInfo
}

func (m *browser) putWindowInfo(browserId int32, windowInfo *TCefWindowInfo) {
	if winInfo, ok := m.windowInfo[browserId]; ok {
		winInfo.Window = windowInfo.Window
	} else {
		m.windowInfo[browserId] = windowInfo
	}
}

func (m *browser) removeWindowInfo(browseId int32) {
	delete(m.windowInfo, browseId)
	Proc("CEF_RemoveGoForm").Call(uintptr(browseId))
}

func (m *browser) GetBrowser(browseId int32) *ICefBrowser {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Browser
	}
	return nil
}

func (m *browser) putBrowserFrame(browser *ICefBrowser, frame *ICefFrame) {
	if winInfo, ok := m.windowInfo[browser.Identifier()]; !ok {
		winInfo = &TCefWindowInfo{
			Browser: browser,
			Frames:  make(map[int64]*ICefFrame),
		}
		if frame != nil {
			winInfo.Frames[frame.Id] = frame
		}
		m.windowInfo[browser.Identifier()] = winInfo
	} else {
		winInfo.Browser = browser
		if frame != nil {
			winInfo.Frames[frame.Id] = frame
		}
	}
}

func (m *browser) GetFrames(browseId int32) map[int64]*ICefFrame {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Frames
	}
	return nil
}

func (m *browser) GetFrame(browseId int32, frameId int64) *ICefFrame {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		return winInfo.Frames[frameId]
	}
	return nil
}

func (m *browser) RemoveFrame(browseId int32, frameId int64) {
	if winInfo, ok := m.windowInfo[browseId]; ok {
		delete(winInfo.Frames, frameId)
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
		for _, frm := range winInfo.Frames {
			if !frm.IsValid() {
				delete(winInfo.Frames, frm.Id)
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
		m.onLoadingStateChange = event
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
		m.chromium.SetOnLoadEnd(event)
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

// BrowserEvent.SetOnRenderCompMsg
func (m *BrowserEvent) SetOnRenderCompMsg(event ChromiumEventOnRenderCompMsg) {
	if Args.IsMain() {
		m.chromium.SetOnRenderCompMsg(event)
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
func (m *BrowserEvent) SetOnBeforePopup(event ChromiumEventOnBeforePopup) {
	if Args.IsMain() {
		m.onBeforePopup = event
	}
}

// BrowserEvent.SetOnOpenUrlFromTab
func (m *BrowserEvent) SetOnOpenUrlFromTab(event ChromiumEventOnOpenUrlFromTab) {
	if Args.IsMain() {
		m.chromium.SetOnOpenUrlFromTab(event)
	}
}

func (m *BrowserEvent) SetOnFindResult(event ChromiumEventOnFindResult) {
	if Args.IsMain() {
		m.chromium.SetOnFindResult(event)
	}
}
