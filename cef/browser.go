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
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types"
	"unsafe"
)

type TBrowser struct {
	cef.ICEFWinControl
	chromium                cef.IChromium
	canClose                bool
	browserId               uint32
	window                  window.IWindow
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	timer                   lcl.ITimer
	windowName              string
	context                 cef.ICefRequestContext
	extraInfo               cef.ICefDictionaryValue
	schemeHandlerFactory    *tSchemeHandlerFactory
}

func NewBrowser(owner lcl.IWinControl) *TBrowser {
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
	m.SendProcessMessageToRenderer("ipc", payload)
}

func (m *TBrowser) ExecuteJavaScript(javaScript string) {
	frame := m.chromium.Browser().GetMainFrame()
	m.chromium.ExecuteJavaScriptWithStrX2FrameInt(javaScript, "", frame, 0)
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

func (m *TBrowser) SetBrowserExtraInfo(windowName string, context cef.ICefRequestContext, extraInfo cef.ICefDictionaryValue) {
	m.windowName = windowName
	m.context = context
	m.extraInfo = extraInfo
}

func (m *TBrowser) CreateBrowser() {
	logger.Debug("Browser.CreateBrowser")
	m.createBrowserOnTimer(m.timer)
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
