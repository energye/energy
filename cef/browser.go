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
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types"
	"sync/atomic"
	"unsafe"
)

// global browser id
var globalBrowserID uint32

// return next browser id
func getNextBrowserID() uint32 {
	atomic.AddUint32(&globalBrowserID, 1)
	return globalBrowserID
}

type TBrowser struct {
	windowParentRaw         cef.ICEFWinControl
	chromiumRaw             cef.IChromium
	isClosed                bool
	browserId               uint32
	window                  window.IWindow
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	timer                   lcl.ITimer
	windowName              string
	context                 cef.ICefRequestContext
	extraInfo               cef.ICefDictionaryValue
}

func NewChromium(owner lcl.IWinControl) *TBrowser {
	m := &TBrowser{browserId: getNextBrowserID()}
	m.chromiumRaw = cef.NewChromium(owner)
	if tool.IsWindows() {
		m.windowParentRaw = cef.NewWindowParent(owner)
	} else {
		windowParent := cef.NewLinkedWindowParent(owner)
		windowParent.SetChromium(m.chromiumRaw)
		m.windowParentRaw = windowParent
	}

	const (
		HpDisableNonProxiedUDP = 3
		STATE_DISABLED         = 2
	)
	m.chromiumRaw.SetWebRTCIPHandlingPolicy(HpDisableNonProxiedUDP)
	m.chromiumRaw.SetWebRTCMultipleRoutes(STATE_DISABLED)
	m.chromiumRaw.SetWebRTCNonproxiedUDP(STATE_DISABLED)

	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	ipc.RegisterProcessMessage(m)

	m.initDefaultEvent()

	m.timer = lcl.NewTimer(owner)
	m.timer.SetEnabled(false)
	m.timer.SetInterval(500)
	m.timer.SetOnTimer(m.onTimerCreateBrowser)

	return m
}

func (m *TBrowser) BrowserId() uint32 {
	return m.browserId
}

func (m *TBrowser) SendMessage(payload []byte) {
	if m.isClosed || len(payload) == 0 {
		return
	}
	m.SendProcessMessageToRenderer("", payload)
}

func (m *TBrowser) ExecuteJavaScript(javaScript string) {
	frame := m.chromiumRaw.Browser().GetMainFrame()
	m.chromiumRaw.ExecuteJavaScriptWithStrX2FrameInt(javaScript, "", frame, 0)
}

func (m *TBrowser) SendProcessMessage(name string, targetProcess cefTypes.TCefProcessId, payload []byte) {
	if m.isClosed || len(payload) == 0 {
		return
	}
	processMessage := cef.ProcessMessageRef.New(name)
	messageArgumentList := processMessage.GetArgumentList()
	dataBin := cef.BinaryValueRef.New(uintptr(unsafe.Pointer(&payload[0])), uint32(len(payload)))
	messageArgumentList.SetBinary(0, dataBin)
	frame := m.chromiumRaw.Browser().GetMainFrame()
	defer func() {
		dataBin.Release()
		messageArgumentList.Clear()
		messageArgumentList.Release()
		processMessage.Release()
	}()
	m.chromiumRaw.SendProcessMessageWithPIdPMessageFrame(targetProcess, processMessage, frame)
}

func (m *TBrowser) SendProcessMessageToBrowser(name string, payload []byte) {
	m.SendProcessMessage(name, cefTypes.PID_BROWSER, payload)
}

func (m *TBrowser) SendProcessMessageToRenderer(name string, payload []byte) {
	m.SendProcessMessage(name, cefTypes.PID_RENDERER, payload)
}

func (m *TBrowser) SetWindow(window window.IWindow) {
	m.window = window
	if m.window != nil {
		if m.window.BrowserId() == 0 {
			m.window.SetBrowserId(m.browserId)
		}
		window.AddOnWindowStateChange(m.doOnWindowStateChange)
		window.AddOnWindowResize(m.doOnWindowResize)
		window.AddOnWindowShow(m.doOnWindowShow)
		window.AddOnWindowClose(m.doOnWindowClose)
		window.AddOnWindowCloseQuery(m.doOnWindowCloseQuery)
	}
}

func (m *TBrowser) SetCreateBrowserExtraInfo(windowName string, context cef.ICefRequestContext, extraInfo cef.ICefDictionaryValue) {
	m.windowName = windowName
	m.context = context
	m.extraInfo = extraInfo
}

func (m *TBrowser) CreateBrowser() {
	m.onTimerCreateBrowser(m.timer)
}

func (m *TBrowser) doOnWindowStateChange(sender lcl.IObject) {
}

func (m *TBrowser) doOnWindowResize(sender lcl.IObject) {
}

func (m *TBrowser) doOnWindowShow(sender lcl.IObject) {
	m.CreateBrowser()
}

func (m *TBrowser) doOnWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {

}

func (m *TBrowser) doOnWindowCloseQuery(sender lcl.IObject, canClose *bool) {
}

func (m *TBrowser) onTimerCreateBrowser(sender lcl.IObject) {
	if m.timer == nil {
		return
	}
	m.timer.SetEnabled(false)
	rect := m.windowParentRaw.ClientRect()
	created := m.chromiumRaw.CreateBrowserWithWHandleRectStrRContextDValueBool(m.windowParentRaw.Handle(), rect, m.windowName,
		m.context, m.extraInfo, false)
	init := m.chromiumRaw.Initialized()
	if !created && !init {
		m.timer.SetEnabled(true)
	} else {
		m.windowParentRaw.UpdateSize()
		m.timer.SetOnTimer(nil)
		m.timer.Free()
		m.timer = nil
	}
}

func (m *TBrowser) initDefaultEvent() {

}
