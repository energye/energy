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
	"encoding/json"
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"sync"
	"unsafe"
)

// IBrowser extension -> core.IBrowser
type IBrowser interface {
	core.IBrowser
	Chromium() cef.IChromium
	SetBrowserExtraInfo(windowName string, context cef.ICefRequestContext, extraInfo cef.ICefDictionaryValue)
	SendProcessMessageToBrowser(name string, payload []byte)
	SendProcessMessageToRenderer(name string, payload []byte)
}

type TBrowser struct {
	canClose                bool
	browserId               uint32
	isClose                 bool
	defaultURL              string
	loadingTitle            string
	loadingURL              string
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	localLoad               *application.LocalLoadResource
	executeScriptCallback   sync.Map
	executeScriptId         int32
	loadingState            core.TLoadChange
	resourceHandlerList     map[string]*source
	dragFilePathCache       map[string]string

	windowName string
	context    cef.ICefRequestContext
	extraInfo  cef.ICefDictionaryValue

	window   window.IWindow
	chromium cef.IChromium

	schemeHandlerFactory  *tSchemeHandlerFactory
	onBrowserAfterCreated lcl.TNotifyEvent
	onProcessMessage      core.TOnProcessMessageEvent
	onResourceRequest     core.TOnResourceRequestEvent
	onLoadChange          core.TOnLoadChangeEvent
	onContextMenu         core.TOnContextMenuEvent
	onContextMenuCommand  core.TOnContextMenuCommandEvent
	onPopupWindow         core.TOnPopupWindowEvent
	onDragEnter           core.TOnDragEnterEvent
	onDragLeave           core.TOnDragLeaveEvent
	onDragOver            core.TOnDragOverEvent
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
	m.SendProcessMessageToRenderer(internalPostMessageName, payload)
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
	message := &tExecuteScriptMessage{
		Id:     m.executeScriptId,
		Script: script,
	}
	payload, err := json.Marshal(message)
	if err != nil {
		m.executeScriptCallback.Delete(m.executeScriptId)
		return
	}
	m.SendProcessMessageToRenderer(internalExecuteScriptName, payload)
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

func (m *TBrowser) SetBrowserExtraInfo(windowName string, context cef.ICefRequestContext, extraInfo cef.ICefDictionaryValue) {
	m.windowName = windowName
	m.context = context
	m.extraInfo = extraInfo
}
