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

type TCEFApplication struct {
	cef.ICefApplication
}

type TCEFWorkScheduler struct {
	cef.ICEFWorkScheduler
}

type TCEFWindowParent struct {
	cef.ICEFWinControl
}

type TCEFChromium struct {
	cef.IChromium
}

type ICEFRequestContext interface {
	cef.ICefRequestContext
}

type ICEFDictionaryValue interface {
	cef.ICefDictionaryValue
}

func (m *Application) IsMainProcess() bool {
	return m.ProcessType() == types.PtBrowser
}

func (m *Application) Free() {
	m.TCEFApplication.ICefApplication.Free()
}

func (m *TCEFWorkScheduler) Free() {
	m.ICEFWorkScheduler.Free()
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

func NewCEFApplication() *TCEFApplication {
	m := &TCEFApplication{}
	m.ICefApplication = cef.NewApplication()
	return m
}

func NewCEFWorkScheduler(owner lcl.IComponent) *TCEFWorkScheduler {
	m := &TCEFWorkScheduler{}
	m.ICEFWorkScheduler = cef.NewWorkScheduler(owner)
	return m
}

func NewCEFChromium(owner lcl.IComponent) *TCEFChromium {
	m := &TCEFChromium{}
	m.IChromium = cef.NewChromium(owner)
	return m
}

func NewCEFWindowParent(chromium cef.IChromium, value lcl.IWinControl) *TCEFWindowParent {
	m := &TCEFWindowParent{}
	if tool.IsWindows() {
		m.ICEFWinControl = cef.NewWindowParent(value)
	} else {
		windowParent := cef.NewLinkedWindowParent(value)
		windowParent.SetChromium(chromium)
		m.ICEFWinControl = windowParent
	}
	return m
}
