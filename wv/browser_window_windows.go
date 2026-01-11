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
	created                 bool
	window                  window.IWindow
	windowParent            wv.IWVWindowParent
	browser                 wv.IWVBrowser
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	onBrowserAfterCreated   lcl.TNotifyEvent
	onProcessMessage        TOnProcessMessageEvent
	onResourceRequest       TOnResourceRequestEvent
	onLoadChange            TOnLoadChangeEvent
	onContextMenu           TOnContextMenuEvent
	onContextMenuCommand    TOnContextMenuCommandEvent
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

// SetWindow 设置webview的窗口实例，并初始化相关回调函数
//
//	window - 窗口接口实例，用于承载webview内容
func (m *TWebview) SetWindow(window window.IWindow) {
	m.window = window
	if m.window != nil {
		if m.window.BrowserId() == 0 {
			m.window.SetBrowserId(m.browserId)
		}
		m.window.SetOptions()
	}
	window.SetOnWindowShow(m.onWindowShow)
	window.SetOnWindowClose(m.onWindowClose)
	window.SetOnWindowCloseQuery(m.onWindowCloseQuery)
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
	m.browser.SetOnContextMenuRequested(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2ContextMenuRequestedEventArgs) {
		args = wv.NewCoreWebView2ContextMenuRequestedEventArgs(args)
		defer args.Free()
		menuItemCollection := wv.NewCoreWebView2ContextMenuItemCollection(args.MenuItems())
		defer menuItemCollection.Free()
		menuItemClear := func(menuItems wv.ICoreWebView2ContextMenuItemCollection) {
			menuItems.RemoveAllMenuItems()
		}
		if gApplication.Options.DisableContextMenu {
			menuItemClear(menuItemCollection)
			return
		}
		if m.onContextMenu != nil {
			var (
				environment = m.browser.CoreWebView2Environment()
				add         func(text string, kind TContextMenuKind, menuItems wv.ICoreWebView2ContextMenuItemCollection) (*TContextMenuItem, int32)
			)
			add = func(text string, kind TContextMenuKind, menuItems wv.ICoreWebView2ContextMenuItemCollection) (*TContextMenuItem, int32) {
				var itemKind wvTypes.COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND
				switch kind {
				case CmkCommand:
					itemKind = wvTypes.COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND
				case CmkSub:
					itemKind = wvTypes.COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SUBMENU
				case CmkSeparator:
					itemKind = wvTypes.COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SEPARATOR
				default:
					return nil, 0
				}
				var tempMenuItemBaseIntf wv.ICoreWebView2ContextMenuItem
				if environment.CreateContextMenuItem(text, nil, itemKind, &tempMenuItemBaseIntf) {
					menuItem := wv.NewCoreWebView2ContextMenuItem(tempMenuItemBaseIntf)
					menuItem.AddAllBrowserEvents(m.browser)
					menuItems.AppendValue(menuItem.BaseIntf())
					menuItemId := menuItem.CommandId()
					childMenuItems := wv.NewCoreWebView2ContextMenuItemCollection(menuItem.Children())
					contextMenu := &TContextMenuItem{
						clear: func() {
							menuItemClear(childMenuItems)
						},
						add: func(text string, kind TContextMenuKind) (*TContextMenuItem, int32) {
							newMenuItem, newCommandId := add(text, kind, childMenuItems)
							childMenuItems.Free()
							return newMenuItem, newCommandId
						}}
					return contextMenu, menuItemId
				}
				return nil, 0
			}
			contextMenu := &TContextMenuItem{
				clear: func() {
					menuItemClear(menuItemCollection)
				},
				add: func(text string, kind TContextMenuKind) (*TContextMenuItem, int32) {
					return add(text, kind, menuItemCollection)
				}}
			m.onContextMenu(contextMenu)
		}
	})
	// 代理事件, 自定义菜单项选择事件回调
	m.browser.SetOnCustomItemSelected(func(sender lcl.IObject, menuItem wv.ICoreWebView2ContextMenuItem) {
		if m.onContextMenuCommand != nil {
			menuItem = wv.NewCoreWebView2ContextMenuItem(menuItem)
			defer menuItem.Free()
			m.onContextMenuCommand(menuItem.CommandId())
		}
	})
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
		if m.onBrowserAfterCreated != nil {
			m.onBrowserAfterCreated(sender)
		}
		m.windowParent.UpdateSize()
	})
	m.browser.SetOnContentLoading(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2ContentLoadingEventArgs) {
		if m.onLoadChange != nil {
			webview = wv.NewCoreWebView2(webview)
			defer webview.Free()
			uri := webview.Source()
			title := webview.DocumentTitle()
			m.onLoadChange(uri, title, LcLoading)
		}
	})
	m.browser.SetOnNavigationCompleted(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2NavigationCompletedEventArgs) {
		if m.onLoadChange != nil {
			webview = wv.NewCoreWebView2(webview)
			defer webview.Free()
			uri := webview.Source()
			title := webview.DocumentTitle()
			m.onLoadChange(uri, title, LcFinish)
		}
	})
	m.browser.SetOnNavigationStarting(func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2NavigationStartingEventArgs) {
		m.navigationStarting()
		if m.onLoadChange != nil {
			webview = wv.NewCoreWebView2(webview)
			defer webview.Free()
			uri := webview.Source()
			title := webview.DocumentTitle()
			m.onLoadChange(uri, title, LcStart)
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

// CreateBrowser 创建浏览器实例
// 该方法负责初始化webview浏览器，确保只创建一次，并在应用程序初始化完成后创建浏览器窗口
func (m *TWebview) CreateBrowser() {
	if m.created || m.window == nil {
		return
	}
	m.created = true
	if gApplication.InitializationError() {
		// Log ???
	} else {
		if gApplication.Initialized() {
			m.browser.CreateBrowserWithHandleBool(m.windowParent.Handle(), true)
		}
	}
}

// BrowserId 返回TWebview实例关联的浏览器ID
func (m *TWebview) BrowserId() uint32 {
	return m.browserId
}

// SendMessage 发送消息到webview浏览器
func (m *TWebview) SendMessage(payload []byte) {
	if m.isClose {
		return
	}
	m.browser.PostWebMessageAsString(string(payload))
}

// Close 关闭webview窗口并清理相关资源
func (m *TWebview) Close() {
	if m.isClose {
		return
	}
	m.isClose = true
	m.windowParent.Free()
	ipc.UnRegisterProcessMessage(m)
}

// SetDefaultURL 设置WebView的默认URL
func (m *TWebview) SetDefaultURL(url string) {
	m.browser.SetDefaultURL(url)
}

// LoadURL 加载指定的URL地址到webview中
func (m *TWebview) LoadURL(url string) {
	m.browser.Navigate(url)
}

// Browser 返回TWebview实例关联的浏览器对象
func (m *TWebview) Browser() wv.IWVBrowser {
	return m.browser
}

// onWindowShow 是窗口显示事件的回调函数
// 当窗口显示时触发此函数，用于创建浏览器实例
func (m *TWebview) onWindowShow(sender lcl.IObject) {
	m.CreateBrowser()
}

// onWindowClose 处理窗口关闭事件的回调函数
// 当窗口接收到关闭信号时，该函数会停止浏览器实例以确保资源被正确释放
func (m *TWebview) onWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {
	if m.browser != nil {
		m.browser.Stop()
	}
}

// onWindowCloseQuery 处理窗口关闭查询事件
// 当用户尝试关闭窗口时触发此回调函数
func (m *TWebview) onWindowCloseQuery(sender lcl.IObject, canClose *bool) {
}

// SetOnBrowserAfterCreated 设置浏览器创建后的回调事件处理函数
func (m *TWebview) SetOnBrowserAfterCreated(fn lcl.TNotifyEvent) {
	m.onBrowserAfterCreated = fn
}

// SetOnResourceRequest 设置资源请求事件处理函数
// 该方法用于注册一个回调函数，当webview发起资源请求时会触发此回调
func (m *TWebview) SetOnResourceRequest(fn TOnResourceRequestEvent) {
	m.onResourceRequest = fn
}

// SetOnProcessMessage 设置处理进程消息的回调函数
// 该方法用于注册一个回调函数，当接收到进程消息时会触发该回调
func (m *TWebview) SetOnProcessMessage(fn TOnProcessMessageEvent) {
	m.onProcessMessage = fn
}

func (m *TWebview) SetOnLoadChange(fn TOnLoadChangeEvent) {
	m.onLoadChange = fn
}

func (m *TWebview) SetOnContextMenu(fn TOnContextMenuEvent) {
	m.onContextMenu = fn
}

func (m *TWebview) SetOnContextMenuCommand(fn TOnContextMenuCommandEvent) {
	m.onContextMenuCommand = fn
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
