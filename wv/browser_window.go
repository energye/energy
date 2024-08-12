//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import (
	"encoding/json"
	"github.com/energye/energy/v3/internal/assets"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/wv/wv"
	"sync/atomic"
)

type OnWindowCreate func(window IBrowserWindow)

// IBrowserWindow
//
//	A browser window composed of TForms and webview2 controls
type IBrowserWindow interface {
	lcl.IForm
	IsClosing() bool
	WindowParent() wv.IWVWindowParent
	Browser() wv.IWVBrowser
	WindowId() uint32
	_WndProc(message uint32, wParam, lParam uintptr) uintptr
	// SetOnBrowserAfterCreated Called after a new browser is created and it's ready to navigate to the default URL.
	SetOnBrowserAfterCreated(fn wv.TNotifyEvent)
	// SetOnBrowserMessageReceived
	//  runs when the `ICoreWebView2Settings.IsWebMessageEnabled`
	//	setting is set and the top-level document of the WebView runs `window.chrome.webview.postMessage`.
	SetOnBrowserMessageReceived(fn wv.TOnWebMessageReceivedEvent)
	// SetOnNavigationStarting
	//  runs when the WebView main frame is requesting
	//  permission to navigate to a different URI. Redirects trigger this
	//  operation as well, and the navigation id is the same as the original
	//  one. Navigations will be blocked until all `NavigationStarting` event handlers
	SetOnNavigationStarting(fn wv.TOnNavigationStartingEvent)
	SetOnShow(fn wv.TNotifyEvent)
	SetOnResize(fn wv.TNotifyEvent)
	SetOnClose(fn lcl.TCloseEvent)
	FullScreen()
	ExitFullScreen()
	Minimize()
	Maximize()
	Restore()
	IsFullScreen() bool
	IsMinimize() bool
	IsMaximize() bool
}

// BrowserWindow
//
//	energy webview2 window, It consists of TForm and WVBrowser
type BrowserWindow struct {
	lcl.TForm
	windowId                   uint32
	isClosing                  bool
	windowParent               wv.IWVWindowParent
	browser                    wv.IWVBrowser
	options                    Options
	onWindowCreate             OnWindowCreate
	onAfterCreated             wv.TNotifyEvent
	onWebMessageReceived       wv.TOnWebMessageReceivedEvent
	onNavigationStarting       wv.TOnNavigationStartingEvent
	onContextMenuRequested     wv.TOnContextMenuRequestedEvent
	onShow                     wv.TNotifyEvent
	onDestroy                  wv.TNotifyEvent
	onClose                    lcl.TCloseEvent
	ipcMessageReceivedDelegate ipc.IMessageReceivedDelegate
	oldWndPrc                  uintptr
	previousWindowPlacement    types.TRect
	windowsState               types.TWindowState
}

var windowID uint32

func getWindowID() uint32 {
	atomic.AddUint32(&windowID, 1)
	return windowID
}

// FormCreate
// The function called when creating a window
// Initialize window properties or other tasks here
func (m *BrowserWindow) FormCreate(sender lcl.IObject) {
	m.windowId = getWindowID()
	// 1. add browser window in globalBrowserWindows
	addBrowserWindow(m)
	// 2. windows hook wndProc message
	m._HookWndProcMessage()
	// 3. setting and init
	m.SetCaption(m.options.Caption)
	m.SetBounds(m.options.X, m.options.Y, m.options.Width, m.options.Height)
	m.SetDoubleBuffered(true)
	m.SetShowInTaskBar(types.StAlways)
	//m.SetColor(colors.ClBlack)
	// background panel
	background := lcl.NewPanel(m)
	background.SetParent(m)
	background.SetParentColor(true)
	background.SetParentDoubleBuffered(true)
	background.SetBevelInner(types.BvNone)
	background.SetBevelOuter(types.BvNone)
	//background.SetAlign(types.AlClient)
	background.SetAlign(types.AlCustom)
	background.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	background.SetBounds(m.options.X, m.options.Y, m.options.Width, m.options.Height)
	// browser window parent component
	m.windowParent = wv.NewWVWindowParent(m)
	m.windowParent.SetParent(background)
	m.windowParent.SetParentDoubleBuffered(true)
	m.windowParent.SetAlign(types.AlClient)
	//m.windowParent.SetAlign(types.AlCustom)
	//m.windowParent.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	//var b int32 = 2
	//m.windowParent.SetBounds(b/2, b/2, m.options.Width-b, m.options.Height-b)

	// browser
	m.browser = wv.NewWVBrowser(m)
	m.windowParent.SetBrowser(m.browser)
	if m.options.DefaultURL != "" {
		m.browser.SetDefaultURL(m.options.DefaultURL)
	}
	if m.options.ICON == nil {
		lcl.Application.Icon().LoadFromBytes(assets.ICON.ICO())
	} else {
		lcl.Application.Icon().LoadFromBytes(m.options.ICON)
	}
	// Registers the current window to the process message for future use when the specified window handles the message
	ipc.RegisterProcessMessage(m)
	ipc.SetMainWindowId(m.WindowId())
	// BrowserWindow Default preset function implementation
	m.defaultEvent()
	// call window main form create callback
	if m.onWindowCreate != nil {
		m.onWindowCreate(m)
	}
}

// WindowId Window ID, generated by accumulating sequence numbers
func (m *BrowserWindow) WindowId() uint32 {
	return m.windowId
}

func (m *BrowserWindow) SendMessage(payload []byte) {
	if m.IsClosing() {
		return
	}
	m.browser.PostWebMessageAsString(string(payload))
}

// Default preset function implementation
//
//	Users have two options when implementing event behavior on their own
//	1. Use Browser() to obtain the browser object and remove and override the current specified event
//	2. Specify the event function in the current window and retain the default event behavior
func (m *BrowserWindow) defaultEvent() {
	// ipc message received
	m.ipcMessageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	// webview2 AfterCreated
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		m.windowParent.UpdateSize()
		// local load
		if m.options.LocalLoad != nil {
			m.browser.AddWebResourceRequestedFilter(m.options.LocalLoad.Scheme+"*", wv.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
		}
		// current browser ipc javascript
		m.browser.CoreWebView2().AddScriptToExecuteOnDocumentCreated(string(ipcJS), m.browser)
		// CoreWebView2Settings
		settings := m.browser.CoreWebView2Settings()
		// Global control of devtools account open and clos
		settings.SetAreDevToolsEnabled(!m.options.DisableDevTools)
		if m.onAfterCreated != nil {
			m.onAfterCreated(sender)
		}
	})
	m.browser.SetOnNavigationStarting(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2NavigationStartingEventArgs) {
		m.navigationStarting(webview, args)
		jsCode := `window.energy.drag().enableDrag(true);window.energy.drag().setup();`
		m.browser.ExecuteScript(jsCode, 0)
		if m.onNavigationStarting != nil {
			m.onNavigationStarting(sender, webview, args)
		}
	})
	// context menu
	m.browser.SetOnContextMenuRequested(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2ContextMenuRequestedEventArgs) {
		if m.options.DisableContextMenu {
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
	m.browser.SetOnWebMessageReceived(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebMessageReceivedEventArgs) {
		var handle bool
		if m.ipcMessageReceivedDelegate != nil {
			args = wv.NewCoreWebView2WebMessageReceivedEventArgs(args)
			message := args.WebMessageAsString()
			args.Free()
			// ipc message
			var pMessage ipc.ProcessMessage
			err := json.Unmarshal([]byte(message), &pMessage)
			if err == nil {
				//fmt.Println("message:", pMessage.Type, pMessage.Data)
				switch pMessage.Type {
				case ipc.MT_READY:
					// ipc ready
					handle = true
				case ipc.MT_EVENT_GO_EMIT, ipc.MT_EVENT_JS_EMIT, ipc.MT_EVENT_GO_EMIT_CALLBACK, ipc.MT_EVENT_JS_EMIT_CALLBACK:
					// ipc on, emit event
					handle = m.ipcMessageReceivedDelegate.Received(m.WindowId(), &pMessage)
				case ipc.MT_DRAG_MOVE, ipc.MT_DRAG_DOWN, ipc.MT_DRAG_UP, ipc.MT_DRAG_DBLCLICK:
					// ipc drag window
					m.Drag(pMessage)
					handle = true
				case ipc.MT_DRAG_RESIZE:
					// border drag resize
					ht := pMessage.Data.(string)
					m.Resize(ht)
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
	// window, OnShow
	m.TForm.SetOnShow(func(sender lcl.IObject) {
		if application.InitializationError() {
			// Log ???
		} else {
			if application.Initialized() {
				m.browser.CreateBrowser(m.windowParent.Handle(), true)
			}
		}
		if m.onShow != nil {
			m.onShow(sender)
		}
	})
	// window, OnClose
	m.TForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		if m.onClose != nil {
			m.onClose(sender, action)
		}
		// window close and free
		if *action == types.CaFree {
			m.isClosing = true
			// cancel process message
			ipc.UnRegisterProcessMessage(m)
		}
	})
	m.TForm.SetOnDestroy(func(sender lcl.IObject) {
		m._RestoreWndProc()
		if m.onDestroy != nil {
			m.onDestroy(sender)
		}
	})
}

func (m *BrowserWindow) IsClosing() bool {
	return m.isClosing
}

func (m *BrowserWindow) WindowParent() wv.IWVWindowParent {
	return m.windowParent
}

func (m *BrowserWindow) Browser() wv.IWVBrowser {
	return m.browser
}

func (m *BrowserWindow) SetOnShow(fn wv.TNotifyEvent) {
	m.onShow = fn
}

func (m *BrowserWindow) SetOnClose(fn lcl.TCloseEvent) {
	m.onClose = fn
}

func (m *BrowserWindow) SetOnDestroy(fn lcl.TNotifyEvent) {
	m.onDestroy = fn
}

func (m *BrowserWindow) SetOnBrowserAfterCreated(fn wv.TNotifyEvent) {
	m.onAfterCreated = fn
}

func (m *BrowserWindow) SetOnBrowserMessageReceived(fn wv.TOnWebMessageReceivedEvent) {
	m.onWebMessageReceived = fn
}

func (m *BrowserWindow) SetOnNavigationStarting(fn wv.TOnNavigationStartingEvent) {
	m.onNavigationStarting = fn
}

func (m *BrowserWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.IsMinimize() || m.IsMaximize() {
			m.Restore()
		}
		m.windowsState = types.WsFullScreen
		m.previousWindowPlacement = m.BoundsRect()
		monitorRect := m.Monitor().BoundsRect()
		//m.SetBoundsRect(&monitorRect)
		win.SetWindowPos(m.Handle(), win.HWND_TOP, monitorRect.Left, monitorRect.Top, monitorRect.Width(), monitorRect.Height(), win.SWP_NOOWNERZORDER|win.SWP_FRAMECHANGED)
	})
}

func (m *BrowserWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		m.windowsState = types.WsNormal
		m.SetWindowState(types.WsNormal)
		m.SetBoundsRect(&m.previousWindowPlacement)
		return
	}
}

func (m *BrowserWindow) Minimize() {
	lcl.RunOnMainThreadAsync(func(id uint32) {
		m.SetWindowState(types.WsMinimized)
	})
}

func (m *BrowserWindow) Maximize() {
	if m.IsFullScreen() {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if m.WindowState() == types.WsNormal {
			m.SetWindowState(types.WsMaximized)
		} else {
			m.SetWindowState(types.WsNormal)
		}
	})
}

func (m *BrowserWindow) Restore() {
	if m.IsMinimize() || m.IsMaximize() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.SetWindowState(types.WsNormal)
		})
	} else if m.IsFullScreen() {
		m.ExitFullScreen()
	}
}

func (m *BrowserWindow) IsFullScreen() bool {
	return m.windowsState == types.WsFullScreen
}

func (m *BrowserWindow) IsMinimize() bool {
	return m.WindowState() == types.WsMinimized
}

func (m *BrowserWindow) IsMaximize() bool {
	return m.WindowState() == types.WsMaximized
}
