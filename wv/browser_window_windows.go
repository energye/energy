//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows

package wv

import (
	"bytes"
	"encoding/json"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	wvTypes "github.com/energye/wv/types/windows"
	wv "github.com/energye/wv/windows"
	"net/url"
	"runtime"
)

var (
	frameWidth  = win.GetSystemMetrics(32)
	frameHeight = win.GetSystemMetrics(33)
	frameCorner = frameWidth + frameHeight
)

type TBrowserWindow struct {
	browserId                     uint32
	isClose                       bool
	box                           lcl.IPanel
	windowParent                  wv.IWVWindowParent
	browser                       wv.IWVBrowser
	messageReceivedDelegate       ipc.IMessageReceivedDelegate
	onWebMessageReceived          wv.TOnWebMessageReceivedEvent
	onContextMenuRequested        wv.TOnContextMenuRequestedEvent
	onContentLoading              wv.TOnContentLoadingEvent
	onAfterCreated                lcl.TNotifyEvent
	onWindowClose                 lcl.TCloseEvent
	onWindowShow                  lcl.TNotifyEvent
	onWindowDestroy               lcl.TNotifyEvent
	onWebResourceRequested        TOnWebResourceRequestedEvent
	onWebResourceResponseReceived TOnWebResourceResponseReceivedEvent
	onWindowResize                TOnWindowResize
	onWindowDrag                  TOnWindowDrag
}

func NewBrowserWindow(owner lcl.IWinControl) IBrowserWindow {
	m := &TBrowserWindow{browserId: getBrowserID()}
	m.box = lcl.NewPanel(owner)
	m.box.SetParentColor(true)
	m.box.SetParentDoubleBuffered(true)
	m.box.SetBevelInner(types.BvNone)
	m.box.SetBevelOuter(types.BvNone)
	m.box.SetAlign(types.AlCustom)
	m.box.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	m.box.SetBounds(0, 0, gApplication.Options.Width, gApplication.Options.Height)
	m.box.SetParent(owner)

	m.windowParent = wv.NewWindowParent(owner)
	m.windowParent.SetParent(m.box)

	m.browser = wv.NewBrowser(owner)

	m.windowParent.SetBrowser(m.browser)
	// ipc message received
	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	ipc.RegisterProcessMessage(m)
	m.initDefaultEvent()
	return m
}

func (m *TBrowserWindow) BrowserID() uint32 {
	return m.browserId
}

func (m *TBrowserWindow) navigationStarting() {
	jsCode := &bytes.Buffer{}
	var envJS = func(json string) {
		jsCode.WriteString(`window.energy.setOptionsEnv(`)
		jsCode.WriteString(json)
		jsCode.WriteString(`);`)
	}
	optionsJSON, err := json.Marshal(gApplication.Options)
	if err == nil {
		envJS(string(optionsJSON))
	}
	env := make(map[string]interface{})
	env["frameWidth"] = frameWidth
	env["frameHeight"] = frameHeight
	env["frameCorner"] = frameCorner
	env["os"] = runtime.GOOS
	envJSON, err := json.Marshal(env)
	if err == nil {
		envJS(string(envJSON))
	}
	m.browser.ExecuteScript(jsCode.String(), 0)
	m.browser.ExecuteScript(`window.energy.drag().setup();`, 0)
}

// Default preset function implementation
//
//	Users have two options when implementing event behavior on their own
//	1. Use Browser() to obtain the browser object and remove and override the current specified event
//	2. Specify the event function in the current window and retain the default event behavior
func (m *TBrowserWindow) initDefaultEvent() {
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		// local load
		if gApplication.LocalLoad != nil {
			m.browser.AddWebResourceRequestedFilter(gApplication.LocalLoad.Scheme+"*", wvTypes.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
		}
		// current browser ipc javascript
		m.browser.CoreWebView2().AddScriptToExecuteOnDocumentCreated(string(ipcJS), m.browser)
		// CoreWebView2Settings
		settings := m.browser.CoreWebView2Settings()
		// Global control of devtools account open and clos
		settings.SetAreDevToolsEnabled(!gApplication.Options.DisableDevTools)
		if m.onAfterCreated != nil {
			m.onAfterCreated(sender)
		}
		m.windowParent.UpdateSize()
	})
	m.browser.SetOnContentLoading(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2ContentLoadingEventArgs) {
		m.navigationStarting()
		if m.onContentLoading != nil {
			m.onContentLoading(sender, webview, args)
		}
	})
	m.browser.SetOnContextMenuRequested(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2ContextMenuRequestedEventArgs) {
		if gApplication.Options.DisableContextMenu {
			args = wv.NewCoreWebView2ContextMenuRequestedEventArgs(args)
			menuItemCollection := wv.NewCoreWebView2ContextMenuItemCollection(args.MenuItems())
			menuItemCollection.RemoveAllMenuItems()
			menuItemCollection.Free()
			args.Free()
		} else if m.onContextMenuRequested != nil {
			m.onContextMenuRequested(sender, webview, args)
		}
	})
	// process message received
	m.browser.SetOnWebMessageReceived(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebMessageReceivedEventArgs) {
		var handle bool
		if m.messageReceivedDelegate != nil {
			args = wv.NewCoreWebView2WebMessageReceivedEventArgs(args)
			message := args.WebMessageAsString()
			args.Free()
			// ipc message
			var pMessage ipc.ProcessMessage
			err := json.Unmarshal([]byte(message), &pMessage)
			if err == nil {
				switch pMessage.Type {
				case ipc.MT_READY:
					// ipc ready
					handle = true
				case ipc.MT_EVENT_GO_EMIT, ipc.MT_EVENT_JS_EMIT, ipc.MT_EVENT_GO_EMIT_CALLBACK, ipc.MT_EVENT_JS_EMIT_CALLBACK:
					// ipc on, emit event
					handle = m.messageReceivedDelegate.Received(m.BrowserID(), &pMessage)
				case ipc.MT_DRAG_MOVE, ipc.MT_DRAG_DOWN, ipc.MT_DRAG_UP, ipc.MT_DRAG_DBLCLICK:
					// ipc drag window
					if m.onWindowDrag != nil {
						m.onWindowDrag(pMessage)
						handle = true
					}
				case ipc.MT_DRAG_RESIZE:
					// border drag resize
					if m.onWindowResize != nil {
						ht := pMessage.Data.(string)
						m.onWindowResize(ht)
					}
				case ipc.MT_DRAG_BORDER_WMSZ:
					//fmt.Println("pMessage.Data", pMessage.Data)
					//m._SetCursor(17)
				}
			}
		}
		if !handle && m.onWebMessageReceived != nil {
			m.onWebMessageReceived(sender, webview, args)
		}
	})
	m.browser.SetOnWebResourceResponseReceived(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebResourceResponseReceivedEventArgs) {
		var handle bool
		if m.onWebResourceResponseReceived != nil {
			handle = m.onWebResourceResponseReceived(sender, webview, args)
		}
		if !handle && gApplication.LocalLoad != nil {
			tempArgs := wv.NewCoreWebView2WebResourceResponseReceivedEventArgs(args)
			defer tempArgs.Free()
			tempRequest := wv.NewCoreWebView2WebResourceRequestRef(tempArgs.Request())
			defer tempRequest.Free()
			if reqUrl, err := url.Parse(tempRequest.URI()); err == nil {
				gApplication.LocalLoad.ReleaseStream(reqUrl.Path)
			}
		}
	})
	m.browser.SetOnWebResourceRequested(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebResourceRequestedEventArgs) {
		var handle bool
		if m.onWebResourceRequested != nil {
			handle = m.onWebResourceRequested(sender, webview, args)
		}
		if !handle && gApplication.LocalLoad != nil {
			lcl.RunOnMainThreadSync(func() {
				gApplication.LocalLoad.ResourceRequested(m.browser, webview, args)
			})
		}
	})
}

// 在窗口时调用
func (m *TBrowserWindow) CreateBrowser() {
	if gApplication.InitializationError() {
		// Log ???
	} else {
		if gApplication.Initialized() {
			m.browser.CreateBrowserWithHandleBool(m.windowParent.Handle(), true)
		}
	}
}

// BrowserId Window ID, generated by accumulating sequence numbers
func (m *TBrowserWindow) BrowserId() uint32 {
	return m.browserId
}

func (m *TBrowserWindow) SendMessage(payload []byte) {
	if m.isClose {
		return
	}
	m.browser.PostWebMessageAsString(string(payload))
}

func (m *TBrowserWindow) Close() {
	m.isClose = true
	m.windowParent.Free()
	ipc.UnRegisterProcessMessage(m)
}
