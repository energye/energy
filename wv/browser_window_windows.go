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
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/pkgs/mime"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
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

type TWebview struct {
	lcl.IPanel
	browserId               uint32
	isClose                 bool
	window                  window.IWindow
	windowParent            wv.IWVWindowParent
	browser                 wv.IWVBrowser
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	onAfterCreated          lcl.TNotifyEvent
	onWindowClose           lcl.TCloseEvent
	onWindowShow            lcl.TNotifyEvent
	onWindowDestroy         lcl.TNotifyEvent
	onProcessMessage        TOnProcessMessageEvent
	onResourceRequest       TOnResourceRequestEvent
	onContextMenuRequested  wv.TOnContextMenuRequestedEvent
	onContentLoading        wv.TOnContentLoadingEvent
}

// NewWebview 创建一个新的浏览器窗口实例
func NewWebview(owner lcl.IComponent) IWebview {
	m := &TWebview{browserId: getNextBrowserID()}
	m.IPanel = lcl.NewPanel(owner)
	m.IPanel.SetParentColor(true)
	m.IPanel.SetParentDoubleBuffered(true)
	m.IPanel.SetBevelInner(types.BvNone)
	m.IPanel.SetBevelOuter(types.BvNone)

	m.windowParent = wv.NewWindowParent(owner)
	m.windowParent.SetAlign(types.AlClient)

	m.browser = wv.NewBrowser(owner)

	m.windowParent.SetBrowser(m.browser)
	// ipc message received
	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	ipc.RegisterProcessMessage(m)
	m.initDefaultEvent()
	m.SetBrowserOptions()
	return m
}

func (m *TWebview) SetWindow(window window.IWindow) {
	m.window = window
	if m.window != nil {
		if m.window.BrowserId() == 0 {
			m.window.SetBrowserId(m.browserId)
		}
		m.window.SetOptions()
	}
}

// SetBrowserOptions 设置浏览器窗口的选项配置
func (m *TWebview) SetBrowserOptions() {
	options := gApplication.Options
	if options.DefaultURL != "" {
		m.browser.SetDefaultURL(options.DefaultURL)
	}
}

// SetParent 设置浏览器窗口的父控件
// 该方法会同时设置内部面板的父控件和窗口父控件的引用
func (m *TWebview) SetParent(window lcl.IWinControl) {
	m.IPanel.SetParent(window)
	m.windowParent.SetParent(m.IPanel)
}

func (m *TWebview) navigationStarting() {
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
	env := make(map[string]any)
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
func (m *TWebview) initDefaultEvent() {
	streams := make(map[string]lcl.IMemoryStream)
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
		args = wv.NewCoreWebView2WebMessageReceivedEventArgs(args)
		message := args.WebMessageAsString()
		args.Free()
		if m.messageReceivedDelegate != nil {
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
					handle = m.messageReceivedDelegate.Received(m.BrowserId(), &pMessage)
				case ipc.MT_DRAG_MOVE, ipc.MT_DRAG_DOWN, ipc.MT_DRAG_UP, ipc.MT_DRAG_DBLCLICK:
					// ipc drag window
					if m.window != nil {
						m.drag(pMessage)
						handle = true
					}
				case ipc.MT_DRAG_RESIZE:
					// border drag resize
					if m.window != nil {
						ht := pMessage.Data.(string)
						m.resize(ht)
						handle = true
					}
				case ipc.MT_DRAG_BORDER_WMSZ:
					//fmt.Println("pMessage.Data", pMessage.Data)
					//m._SetCursor(17)
				}
			}
		}
		if !handle && m.onProcessMessage != nil {
			m.onProcessMessage(message)
		}
	})
	m.browser.SetOnWebResourceResponseReceived(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebResourceResponseReceivedEventArgs) {
		tempArgs := wv.NewCoreWebView2WebResourceResponseReceivedEventArgs(args)
		defer tempArgs.Free()
		tempRequest := wv.NewCoreWebView2WebResourceRequestRef(tempArgs.Request())
		defer tempRequest.Free()
		if reqUrl, err := url.Parse(tempRequest.URI()); err == nil {
			if stream, ok := streams[reqUrl.Path]; ok {
				stream.FreeAndNil()
				delete(streams, reqUrl.Path)
			}
		}
	})
	m.browser.SetOnWebResourceRequested(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebResourceRequestedEventArgs) {
		tempArgs := wv.NewCoreWebView2WebResourceRequestedEventArgs(args)
		defer tempArgs.FreeAndNil()
		request := tempArgs.Request()
		tempRequest := wv.NewCoreWebView2WebResourceRequestRef(request)
		defer tempRequest.FreeAndNil()
		var (
			statusCode    int32 = 200
			reasonPhrase        = "OK"
			headers             = ""
			response      wv.ICoreWebView2WebResourceResponse
			stream        lcl.IMemoryStream
			streamAdapter lcl.IStreamAdapter
			uri           = tempRequest.URI()
			reqUrl, err   = url.Parse(uri)
		)
		if err == nil {
			var (
				resource string
				handle   bool
				data     []byte
				path     = reqUrl.Path
				method   = tempRequest.Method()
			)
			if m.onResourceRequest != nil {
				header := make(map[string]string)
				//tempHeaders := wv.NewCoreWebView2HttpRequestHeaders(tempRequest.Headers())
				//defer tempHeaders.Free()
				iterator := tempRequest.Headers().Iterator() //wv.NewCoreWebView2HttpHeadersCollectionIterator(tempHeaders.Iterator())
				if iterator != nil {
					defer iterator.Free()
					var (
						name  string
						value string
					)
					for {
						iterator.GetCurrentHeader(&name, &value)
						if !iterator.MoveNext() {
							break
						}
						header[name] = value
					}
				}
				resource, handle = m.onResourceRequest(uri, path, method, header)
			}
			if handle && resource != "" {
				data = []byte(resource)
			} else if !handle {
				data, err = gApplication.LocalLoad.Read(path)
			}
			if err == nil {
				stream = lcl.NewMemoryStream()
				streamAdapter = lcl.NewStreamAdapter(stream, types.SoOwned)
				defer streamAdapter.Nil()
				lcl.StreamHelper.Write(stream, data)
				// current resource is set temp cache
				// released after the resource processing is complete
				streams[path] = stream
				headers = "Content-Type: " + mime.GetMimeType(reqUrl.Path)
				environment := m.browser.CoreWebView2Environment()
				// success response resource
				environment.CreateWebResourceResponse(lcl.AsStreamAdapter(streamAdapter.AsIntfStream()), statusCode, reasonPhrase, headers, &response)
			}
		}
		if err != nil {
			statusCode = 404
			reasonPhrase = "Not Found"
			environment := m.browser.CoreWebView2Environment()
			// empty response resource
			environment.CreateWebResourceResponse(nil, statusCode, reasonPhrase, headers, &response)
		}
		tempArgs.SetResponse(response)
		if response != nil {
			response.Nil()
			response.Free()
		}
	})
}

// 在窗口显示时调用
func (m *TWebview) CreateBrowser() {
	if gApplication.InitializationError() {
		// Log ???
	} else {
		if gApplication.Initialized() {
			m.browser.CreateBrowserWithHandleBool(m.windowParent.Handle(), true)
		}
	}
}

// BrowserId Window ID, generated by accumulating sequence numbers
func (m *TWebview) BrowserId() uint32 {
	return m.browserId
}

func (m *TWebview) SendMessage(payload []byte) {
	if m.isClose {
		return
	}
	m.browser.PostWebMessageAsString(string(payload))
}

func (m *TWebview) Close() {
	m.isClose = true
	m.windowParent.Free()
	ipc.UnRegisterProcessMessage(m)
}

func (m *TWebview) SetDefaultURL(url string) {
	m.browser.SetDefaultURL(url)
}

func (m *TWebview) LoadURL(url string) {
	m.browser.Navigate(url)
}

func (m *TWebview) SetOnAfterCreated(fn lcl.TNotifyEvent) {
	m.onAfterCreated = fn
}

func (m *TWebview) SetOnWindowClose(fn lcl.TCloseEvent) {
	m.onWindowClose = fn
}

func (m *TWebview) SetOnWindowShow(fn lcl.TNotifyEvent) {
	m.onWindowShow = fn
}

func (m *TWebview) SetOnWindowDestroy(fn lcl.TNotifyEvent) {
	m.onWindowDestroy = fn
}

func (m *TWebview) SetOnResourceRequest(fn TOnResourceRequestEvent) {
	m.onResourceRequest = fn
}

func (m *TWebview) SetOnProcessMessage(fn TOnProcessMessageEvent) {
	m.onProcessMessage = fn
}

func (m *TWebview) drag(message ipc.ProcessMessage) {
	if m.window == nil || m.window.IsFullScreen() {
		return
	}
	switch message.Type {
	case ipc.MT_DRAG_MOVE:
		if m.window.IsFullScreen() {
			return
		}
		if win.ReleaseCapture() {
			win.PostMessage(m.window.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
		m.window.Maximize()
	}
}
func (m *TWebview) resize(ht string) {
	if m.window == nil {
		return
	}
	if m.window.IsFullScreen() || application.GApplication.Options.DisableResize {
		return
	}
	if win.ReleaseCapture() {
		var borderHT uintptr
		switch ht {
		case "n-resize":
			borderHT = messages.HTTOP
		case "ne-resize":
			borderHT = messages.HTTOPRIGHT
		case "e-resize":
			borderHT = messages.HTRIGHT
		case "se-resize":
			borderHT = messages.HTBOTTOMRIGHT
		case "s-resize":
			borderHT = messages.HTBOTTOM
		case "sw-resize":
			borderHT = messages.HTBOTTOMLEFT
		case "w-resize":
			borderHT = messages.HTLEFT
		case "nw-resize":
			borderHT = messages.HTTOPLEFT
		}
		win.PostMessage(m.window.Handle(), messages.WM_NCLBUTTONDOWN, borderHT, 0)
	}
}
