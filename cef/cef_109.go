//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !CEF127 && !CEF147

package cef

import (
	"github.com/energye/cef/109/cef"
	"github.com/energye/cef/109/types"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"unsafe"
)

type ICEFApplication interface {
	cef.ICefApplication
}

type ICEFWorkScheduler interface {
	cef.ICEFWorkScheduler
}

type ICEFWindowParent interface {
	cef.ICEFWinControl
}

type ICEFChromium interface {
	cef.IChromium
}

func (m *Application) IsMainProcess() bool {
	return m.ProcessType() == types.PtBrowser
}

func (m *TChromium) SendProcessMessage(name string, targetProcess types.TCefProcessId, payload []byte) {
	if m.isClosed || len(payload) == 0 {
		return
	}
	processMessage := cef.ProcessMessageRef.New(name)
	messageArgumentList := processMessage.GetArgumentList()
	dataBin := cef.BinaryValueRef.New(uintptr(unsafe.Pointer(&payload[0])), uint32(len(payload)))
	messageArgumentList.SetBinary(0, dataBin)
	frame := m.Browser().GetMainFrame()
	defer func() {
		dataBin.Release()
		messageArgumentList.Clear()
		messageArgumentList.Release()
		processMessage.Release()
	}()
	m.SendProcessMessageWithPIdPMessageFrame(targetProcess, processMessage, frame)
}

func (m *TChromium) SendProcessMessageToBrowser(name string, payload []byte) {
	m.SendProcessMessage(name, types.PID_BROWSER, payload)
}

func (m *TChromium) SendProcessMessageToRenderer(name string, payload []byte) {
	m.SendProcessMessage(name, types.PID_RENDERER, payload)
}

func NewCEFApplication() ICEFApplication {
	return cef.NewApplication()
}

func NewCEFWorkScheduler(owner lcl.IComponent) ICEFWorkScheduler {
	return cef.NewWorkScheduler(owner)
}

func NewCEFChromium(owner lcl.IComponent) ICEFChromium {
	return cef.NewChromium(owner)
}

func NewCEFWindowParent(chromium ICEFChromium, value lcl.IWinControl) ICEFWindowParent {
	if tool.IsWindows() {
		return cef.NewWindowParent(value)
	} else {
		windowParent := cef.NewLinkedWindowParent(value)
		windowParent.SetChromium(chromium)
		return windowParent
	}
}
