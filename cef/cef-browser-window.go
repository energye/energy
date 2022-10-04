//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"sync"
)

var (
	BrowserWindow = &browser{
		browserWindow: &browserWindowForm{},
		browserEvent:  &BrowserEvent{},
		Config:        &browserConfig{},
		windowInfo:    make(map[int32]*TCefWindowInfo),
		windowSerial:  1,
		uiLock:        new(sync.Mutex),
	}
)

// 浏览器包装结构体
type browser struct {
	browserWindow *browserWindowForm        //主窗口
	popupWindow   *BaseWindow               //弹出的子窗口
	browserEvent  *BrowserEvent             //浏览器全局事件
	Config        *browserConfig            //浏览器和窗口配置
	windowInfo    map[int32]*TCefWindowInfo //窗口信息集合
	windowSerial  int32                     //窗口序号
	uiLock        *sync.Mutex
}

type browserConfig struct {
	DefaultUrl                        string                                                   //默认URL地址
	Title                             string                                                   //窗口标题
	Width                             int32                                                    //窗口宽
	Height                            int32                                                    //窗口高
	chromiumConfig                    *tCefChromiumConfig                                      //主窗体浏览器配置
	browserWindowOnEventCallback      func(browserEvent *BrowserEvent, window *TCefWindowInfo) //主窗口初始化回调
	browserWindowAfterOnEventCallback func(window *TCefWindowInfo)                             //在初始化子窗体时使用的函数
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
}

type browserWindowForm struct {
	BaseWindow
	isFirstActivate bool
	tray            ITray
}

// 运行应用
//
// 在这里启动浏览器的主进程和子进程
func Run(cefApp *TCEFApplication) {
	if IsDarwin() && !SingleProcess && !Args.IsMain() {
		// mac os 启动子进程
		cefApp.StartSubProcess()
	} else {
		if cefApp.StartMainProcess() {
			lcl.RunApp(&BrowserWindow.browserWindow)
		}
	}
}

func (m *browserWindowForm) OnFormCreate(sender lcl.IObject) {
	if BrowserWindow.Config.chromiumConfig == nil {
		BrowserWindow.Config.chromiumConfig = NewChromiumConfig()
		BrowserWindow.Config.chromiumConfig.SetEnableMenu(true)
		BrowserWindow.Config.chromiumConfig.SetEnableDevTools(true)
		BrowserWindow.Config.chromiumConfig.SetEnableOpenUrlTab(true)
		BrowserWindow.Config.chromiumConfig.SetEnableWindowPopup(true)
		BrowserWindow.Config.Title = ""
		BrowserWindow.Config.Width = 800
		BrowserWindow.Config.Height = 600
	}
	m.isMainWindow = true
	m.FormCreate()
	m.defaultWindowEvent()
	m.defaultWindowCloseEvent()
	m.ChromiumCreate(BrowserWindow.Config.chromiumConfig, BrowserWindow.Config.DefaultUrl)
	m.putChromiumWindowInfo()
	m.defaultChromiumEvent()
	m.AddOnCloseQuery(func(sender lcl.IObject, canClose *bool) bool {
		if m.tray != nil {
			m.tray.close()
		}
		return false
	})
	if BrowserWindow.Config.Title != "" {
		m.SetCaption(BrowserWindow.Config.Title)
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

// 主窗体浏览器初始回调，事件监听和主窗体设置，不适用添加子窗体组件
//
// 主窗体初始化时调用，改变和设置浏览器事件、主窗体、功能
//
// 只适用于主窗口和事件设置
func (m *browser) SetBrowserInit(fn func(event *BrowserEvent, browserWindow *TCefWindowInfo)) {
	m.Config.setBrowserWindowInitOnEvent(fn)
}

func (m *browser) MainWindowForm() *lcl.TForm {
	if m.browserWindow != nil {
		return m.browserWindow.TForm
	}
	return nil
}

// 主窗体浏览器初始之后回调，主窗体设置，和添加子窗体组件
//
// 主窗体初始化之后调用，改变和设置浏览器事件、主窗体、添加子组件、功能
//
// 主窗口和添加子组件使用
func (m *browser) SetBrowserInitAfter(fn func(browserWindow *TCefWindowInfo)) {
	m.Config.setBrowserWindowInitAfterOnEvent(fn)
}

func (m *browserConfig) SetChromiumConfig(chromiumConfig *tCefChromiumConfig) {
	m.chromiumConfig = chromiumConfig
}

func (m *browserConfig) setBrowserWindowInitOnEvent(fn func(event *BrowserEvent, browserWindow *TCefWindowInfo)) {
	if fn != nil && Args.IsMain() {
		m.browserWindowOnEventCallback = fn
	}
}

func (m *browserConfig) setBrowserWindowInitAfterOnEvent(fn func(browserWindow *TCefWindowInfo)) {
	if fn != nil && Args.IsMain() {
		m.browserWindowAfterOnEventCallback = fn
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
		m.chromium.SetOnBeforeResourceLoad(event)
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
