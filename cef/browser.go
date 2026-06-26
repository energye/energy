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

type browserKind = int32

const (
	bkNone     browserKind = iota // none
	bkViews                       // CEF views framework
	bkEmbedded                    // Embed into window
	bkOSR                         // Off-screen rendering
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

	kind browserKind

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
