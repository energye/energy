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
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"sync/atomic"
)

// global browser id
var globalBrowserID uint32

// return next browser id
func getNextBrowserID() uint32 {
	atomic.AddUint32(&globalBrowserID, 1)
	return globalBrowserID
}

type TChromium struct {
	*TCEFWindowParent
	*TCEFChromium
	isClosed                bool
	browserId               uint32
	window                  window.IWindow
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	timer                   lcl.ITimer
	windowName              string
	context                 ICEFRequestContext
	extraInfo               ICEFDictionaryValue
}

func NewChromium(owner lcl.IWinControl) *TChromium {
	m := &TChromium{browserId: getNextBrowserID()}
	m.TCEFChromium = NewCEFChromium(owner)
	m.TCEFWindowParent = NewCEFWindowParent(m.TCEFChromium, owner)

	const (
		HpDisableNonProxiedUDP = 3
		STATE_DISABLED         = 2
	)
	m.SetWebRTCIPHandlingPolicy(HpDisableNonProxiedUDP)
	m.SetWebRTCMultipleRoutes(STATE_DISABLED)
	m.SetWebRTCNonproxiedUDP(STATE_DISABLED)

	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	ipc.RegisterProcessMessage(m)

	m.initDefaultEvent()

	m.timer = lcl.NewTimer(m)
	m.timer.SetEnabled(false)
	m.timer.SetInterval(500)
	m.timer.SetOnTimer(m.onTimerCreateBrowser)

	return m
}

func (m *TChromium) BrowserId() uint32 {
	return m.browserId
}

func (m *TChromium) SendMessage(payload []byte) {
	if m.isClosed || len(payload) == 0 {
		return
	}
	m.SendProcessMessageToRenderer("", payload)
}

func (m *TChromium) ExecuteJavaScript(javaScript string) {
	frame := m.Browser().GetMainFrame()
	m.ExecuteJavaScriptWithStrX2FrameInt(javaScript, "", frame, 0)
}

func (m *TChromium) SetWindow(window window.IWindow) {
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

func (m *TChromium) SetCreateBrowserExtraInfo(windowName string, context ICEFRequestContext, extraInfo ICEFDictionaryValue) {
	m.windowName = windowName
	m.context = context
	m.extraInfo = extraInfo
}

func (m *TChromium) CreateBrowser() {
	m.onTimerCreateBrowser(m.timer)
}

func (m *TChromium) doOnWindowStateChange(sender lcl.IObject) {
}

func (m *TChromium) doOnWindowResize(sender lcl.IObject) {
}

func (m *TChromium) doOnWindowShow(sender lcl.IObject) {
	m.CreateBrowser()
}

func (m *TChromium) doOnWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {

}

func (m *TChromium) doOnWindowCloseQuery(sender lcl.IObject, canClose *bool) {
}

func (m *TChromium) onTimerCreateBrowser(sender lcl.IObject) {
	if m.timer == nil {
		return
	}
	m.timer.SetEnabled(false)
	rect := m.ClientRect()
	created := m.CreateBrowserWithWHandleRectStrRContextDValueBool(m.Handle(), rect, m.windowName,
		m.context, m.extraInfo, false)
	init := m.Initialized()
	if !created && !init {
		m.timer.SetEnabled(true)
	} else {
		m.UpdateSize()
		m.timer.SetOnTimer(nil)
		m.timer.Free()
		m.timer = nil
	}
}

func (m *TChromium) initDefaultEvent() {

}
