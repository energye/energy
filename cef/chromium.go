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
	ICEFWindowParent
	ICEFChromium
	isClosed                bool
	browserId               uint32
	window                  window.IWindow
	messageReceivedDelegate ipc.IMessageReceivedDelegate
}

func NewChromium(owner lcl.IWinControl) *TChromium {
	m := &TChromium{browserId: getNextBrowserID()}
	m.ICEFChromium = NewCEFChromium(owner)
	m.ICEFWindowParent = NewCEFWindowParent(m.ICEFChromium, owner)

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

func (m *TChromium) initDefaultEvent() {

}
