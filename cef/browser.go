//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
	"sync"
	"unsafe"
)

// IBrowser extension -> core.IBrowser
type IBrowser interface {
	core.IBrowser
	Chromium() cef.IChromium
}

type TBrowser struct {
	cef.ICEFWinControl
	chromium                cef.IChromium
	canClose                bool
	browserId               uint32
	isClose                 bool
	defaultURL              string
	loadingTitle            string
	loadingURL              string
	loadingState            core.TLoadChange
	resourceHandlerList     map[string]*source
	executeScriptCallback   sync.Map
	executeScriptId         int32
	window                  window.IWindow
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	localLoad               *application.LocalLoadResource
	timer                   lcl.ITimer
	windowName              string
	context                 cef.ICefRequestContext
	extraInfo               cef.ICefDictionaryValue
	schemeHandlerFactory    *tSchemeHandlerFactory
	onBrowserAfterCreated   lcl.TNotifyEvent
	onProcessMessage        core.TOnProcessMessageEvent
	onResourceRequest       core.TOnResourceRequestEvent
	onLoadChange            core.TOnLoadChangeEvent
	onContextMenu           core.TOnContextMenuEvent
	onContextMenuCommand    core.TOnContextMenuCommandEvent
	onPopupWindow           core.TOnPopupWindowEvent
	onDragEnter             core.TOnDragEnterEvent
	onDragLeave             core.TOnDragLeaveEvent
	onDragOver              core.TOnDragOverEvent
}

// NewBrowser creates a new browser window instance
func NewBrowser(owner lcl.IWinControl) IBrowser {
	m := &TBrowser{}
	m.chromium = cef.NewChromium(owner)
	if tool.IsWindows() {
		m.ICEFWinControl = cef.NewWindowParent(owner)
	} else {
		windowParent := cef.NewLinkedWindowParent(owner)
		windowParent.SetChromium(m.chromium)
		m.ICEFWinControl = windowParent
	}

	m.chromium.SetWebRTCIPHandlingPolicy(cefTypes.HpDisableNonProxiedUDP)
	m.chromium.SetWebRTCMultipleRoutes(cefTypes.STATE_DISABLED)
	m.chromium.SetWebRTCNonproxiedUDP(cefTypes.STATE_DISABLED)

	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()

	m.initDefaultEvent()

	m.timer = lcl.NewTimer(owner)
	m.timer.SetEnabled(false)
	m.timer.SetInterval(500)
	m.timer.SetOnTimer(m.createBrowserOnTimer)

	return m
}

func (m *TBrowser) BrowserId() uint32 {
	return m.browserId
}

func (m *TBrowser) Chromium() cef.IChromium {
	return m.chromium
}

func (m *TBrowser) SendMessage(payload []byte) {
	if m.canClose || len(payload) == 0 {
		return
	}
	m.SendProcessMessageToRenderer(core.PostMessageName, payload)
}

func (m *TBrowser) ExecuteScript(javaScript string) {
	frame := m.chromium.Browser().GetMainFrame()
	m.chromium.ExecuteJavaScriptWithStrX2FrameInt(javaScript, "", frame, 0)
}

func (m *TBrowser) ExecuteScriptCallback(script string, callback core.TOnEvaluateScriptCallbackEvent) {
	if script == "" || callback == nil {
		return
	}
	m.executeScriptId++
	m.executeScriptCallback.Store(m.executeScriptId, callback)
	frame := m.chromium.Browser().GetMainFrame()
	m.chromium.ExecuteJavaScriptWithStrX2FrameInt(script, "", frame, 0)
}

func (m *TBrowser) SetLocalLoad(localLoad application.LocalLoad) {
	m.localLoad = application.NewLocalLoadResource(&localLoad)
}

func (m *TBrowser) LocalLoadResource() *application.LocalLoadResource {
	return m.localLoad
}

func (m *TBrowser) SendProcessMessage(name string, targetProcess cefTypes.TCefProcessId, payload []byte) {
	if m.canClose || len(payload) == 0 {
		return
	}
	processMessage := cef.ProcessMessageRef.New(name)
	messageArgumentList := processMessage.GetArgumentList()
	dataBin := cef.BinaryValueRef.New(uintptr(unsafe.Pointer(&payload[0])), uint32(len(payload)))
	messageArgumentList.SetBinary(0, dataBin)
	frame := m.chromium.Browser().GetMainFrame()
	defer func() {
		dataBin.Release()
		messageArgumentList.Clear()
		messageArgumentList.Release()
		processMessage.Release()
	}()
	m.chromium.SendProcessMessageWithPIdPMessageFrame(targetProcess, processMessage, frame)
}

func (m *TBrowser) SendProcessMessageToBrowser(name string, payload []byte) {
	m.SendProcessMessage(name, cefTypes.PID_BROWSER, payload)
}

func (m *TBrowser) SendProcessMessageToRenderer(name string, payload []byte) {
	m.SendProcessMessage(name, cefTypes.PID_RENDERER, payload)
}

// SetWindow sets the window instance for webview and initializes related callback functions
//
//	window - Window interface instance hosting webview content
func (m *TBrowser) SetWindow(window window.IWindow) {
	m.window = window
	if m.window != nil {
		window.AddOnWindowStateChange(m.doOnWindowStateChange)
		window.AddOnWindowResize(m.doOnWindowResize)
		window.AddOnWindowShow(m.doOnWindowShow)
		window.AddOnWindowClose(m.doOnWindowClose)
		window.AddOnWindowCloseQuery(m.doOnWindowCloseQuery)
	}
}

// UpdateBrowserOptions updates browser configuration
func (m *TBrowser) UpdateBrowserOptions() {
	// retrieves global LocalLoad configuration
	if GApplication != nil && GApplication.LocalLoad != nil {
		newLocalLoad := *GApplication.LocalLoad.LocalLoad
		m.SetLocalLoad(newLocalLoad)
	}
	if m.window != nil {
		// sets browser configuration
		options := m.window.Options()
		if options.DefaultURL != "" && m.defaultURL == "" {
			m.SetDefaultURL(options.DefaultURL)
		} else if m.defaultURL != "" {
			m.SetDefaultURL(m.defaultURL)
		}
		if options.BackgroundColor != nil {
			r, g, b := byte(options.BackgroundColor.R), byte(options.BackgroundColor.G), byte(options.BackgroundColor.B)
			color := colors.TColor(colors.RGB(r, g, b))
			m.SetColor(color)
		}
	}
}

func (m *TBrowser) SetBrowserExtraInfo(windowName string, context cef.ICefRequestContext, extraInfo cef.ICefDictionaryValue) {
	m.windowName = windowName
	m.context = context
	m.extraInfo = extraInfo
}

func (m *TBrowser) CreateBrowser() {
	logger.Debug("Browser.CreateBrowser")
	m.UpdateBrowserOptions()
	m.createBrowserOnTimer(m.timer)
}

// Close closes the webview window and releases associated resources
func (m *TBrowser) Close() {
	if m.isClose {
		return
	}
	m.isClose = true
	m.chromium.TryCloseBrowser()
	ipc.UnRegisterProcessMessage(m)
}

// SetDefaultURL sets the default URL for the WebView
func (m *TBrowser) SetDefaultURL(url string) {
	if m.defaultURL != url {
		m.chromium.SetDefaultUrl(url)
	}
	m.defaultURL = url
}

// LoadURL loads the specified URL address into the webview
func (m *TBrowser) LoadURL(url string) {
	m.chromium.LoadURLWithStrFrame(url, m.chromium.Browser().GetMainFrame())
}

// Browser returns the browser object associated with the TWebview instance
func (m *TBrowser) Browser() core.Browser {
	return m.chromium
}

// WindowParent obtains the parent window object associated with the TWebview instance
func (m *TBrowser) WindowParent() core.WindowParent {
	return m
}

// SetOnBrowserAfterCreated sets the callback handler triggered after browser creation completes
func (m *TBrowser) SetOnBrowserAfterCreated(fn lcl.TNotifyEvent) {
	m.onBrowserAfterCreated = fn
}

// SetOnResourceRequest sets the handler for resource request events
// This method registers a callback function that will be triggered when the webview initiates a resource request
func (m *TBrowser) SetOnResourceRequest(fn core.TOnResourceRequestEvent) {
	m.onResourceRequest = fn
}

// SetOnProcessMessage sets the callback function for processing process messages
// This method registers a callback that is triggered when a process message is received
func (m *TBrowser) SetOnProcessMessage(fn core.TOnProcessMessageEvent) {
	m.onProcessMessage = fn
}

func (m *TBrowser) SetOnLoadChange(fn core.TOnLoadChangeEvent) {
	m.onLoadChange = fn
}

func (m *TBrowser) SetOnContextMenu(fn core.TOnContextMenuEvent) {
	m.onContextMenu = fn
}

func (m *TBrowser) SetOnContextMenuCommand(fn core.TOnContextMenuCommandEvent) {
	m.onContextMenuCommand = fn
}

func (m *TBrowser) SetOnPopupWindow(fn core.TOnPopupWindowEvent) {
	m.onPopupWindow = fn
}

func (m *TBrowser) SetOnDragEnter(fn core.TOnDragEnterEvent) {
	m.onDragEnter = fn
}

func (m *TBrowser) SetOnDragLeave(fn core.TOnDragLeaveEvent) {
	m.onDragLeave = fn
}

func (m *TBrowser) SetOnDragOver(fn core.TOnDragOverEvent) {
	m.onDragOver = fn
}

func (m *TBrowser) doOnWindowStateChange(sender lcl.IObject) {
}

func (m *TBrowser) doOnWindowResize(sender lcl.IObject) {
	if m.chromium != nil {
		m.chromium.NotifyMoveOrResizeStarted()
		m.UpdateSize()
	}
}

func (m *TBrowser) doOnWindowShow(sender lcl.IObject) {
	logger.Debug("Browser.doOnWindowShow")
	m.CreateBrowser()
}

func (m *TBrowser) doOnWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {
	logger.Debug("Browser.doOnWindowClose")
	*closeAction = types.CaFree
}

func (m *TBrowser) doOnWindowCloseQuery(sender lcl.IObject, canClose *bool) {
	logger.Debug("Browser.doOnWindowCloseQuery canClose:", m.canClose)
	if tool.IsDarwin() {
		*canClose = m.canClose
	} else {
		*canClose = m.canClose
	}
	if !m.canClose {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.chromium.CloseBrowser(true)
		})
	}
}

func (m *TBrowser) createBrowserOnTimer(sender lcl.IObject) {
	if m.timer == nil {
		return
	}
	m.timer.SetEnabled(false)
	rect := m.ClientRect()
	created := m.chromium.CreateBrowserWithWHandleRectStrRContextDValueBool(m.Handle(), rect, m.windowName,
		m.context, m.extraInfo, false)
	init := m.chromium.Initialized()
	logger.Debug("Browser.createBrowserOnTimer created:", created, "init:", init)
	if !created && !init {
		logger.Debug("Browser.createBrowserOnTimer fail")
		m.timer.SetEnabled(true)
	} else {
		logger.Debug("Browser.createBrowserOnTimer success")
		m.UpdateSize()
		m.timer.SetOnTimer(nil)
		m.timer.Free()
		m.timer = nil
	}
}
