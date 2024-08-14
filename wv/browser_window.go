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
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/wv/wv"
	"sync/atomic"
)

// IBrowserWindow
//
//	A browser window composed of TForms and webview2 controls
type IBrowserWindow interface {
	lcl.IForm
	IsClosing() bool
	WindowParent() wv.IWVWindowParent
	Browser() wv.IWVBrowser
	WindowId() uint32
	// SetOnBrowserAfterCreated Called after a new browser is created and it's ready to navigate to the default URL.
	SetOnBrowserAfterCreated(fn TNotifyEvent)
	SetOnContextMenuRequestedEvent(fn TOnContextMenuRequestedEvent)
	// SetOnBrowserMessageReceived
	//  runs when the `ICoreWebView2Settings.IsWebMessageEnabled`
	//	setting is set and the top-level document of the WebView runs `window.chrome.webview.postMessage`.
	SetOnBrowserMessageReceived(fn TOnWebMessageReceivedEvent)
	// SetOnNavigationStarting
	//  runs when the WebView main frame is requesting
	//  permission to navigate to a different URI. Redirects trigger this
	//  operation as well, and the navigation id is the same as the original
	//  one. Navigations will be blocked until all `NavigationStarting` event handlers
	SetOnNavigationStarting(fn TOnNavigationStartingEvent)
	SetOnNewWindowRequestedEvent(fn TOnNewWindowRequestedEventEx)
	SetOnShow(fn TNotifyEvent)
	SetOnResize(fn TNotifyEvent)
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
	windowId                    uint32
	isClosing                   bool
	windowParent                wv.IWVWindowParent
	browser                     wv.IWVBrowser
	options                     Options
	newWindowRequestedEventArgs wv.ICoreWebView2NewWindowRequestedEventArgs
	deferral                    wv.ICoreWebView2Deferral
	onAfterCreated              TNotifyEvent
	onWebMessageReceived        TOnWebMessageReceivedEvent
	onNavigationStarting        TOnNavigationStartingEvent
	onContextMenuRequested      TOnContextMenuRequestedEvent
	onNewWindowRequested        TOnNewWindowRequestedEventEx
	onShow                      TNotifyEvent
	onDestroy                   TNotifyEvent
	onClose                     lcl.TCloseEvent
	ipcMessageReceivedDelegate  ipc.IMessageReceivedDelegate
	oldWndPrc                   uintptr
	previousWindowPlacement     types.TRect
	windowsState                types.TWindowState
	preWindowStyle              uintptr
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
	m.defaultOptions()
	// setting and init
	m.SetCaption(m.options.Caption)
	m.SetBounds(m.options.X, m.options.Y, m.options.Width, m.options.Height)
	m.SetDoubleBuffered(true)
	m.SetShowInTaskBar(types.StAlways)
	// background panel
	background := lcl.NewPanel(m)
	background.SetParent(m)
	background.SetParentColor(true)
	background.SetParentDoubleBuffered(true)
	background.SetBevelInner(types.BvNone)
	background.SetBevelOuter(types.BvNone)
	background.SetAlign(types.AlCustom)
	background.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	background.SetBounds(m.options.X, m.options.Y, m.options.Width, m.options.Height)
	// browser window parent component
	m.windowParent = wv.NewWVWindowParent(m)
	m.windowParent.SetParent(background)
	m.windowParent.SetParentDoubleBuffered(true)
	m.windowParent.SetAlign(types.AlClient)

	// webview browser
	m.browser = wv.NewWVBrowser(m)
	m.windowParent.SetBrowser(m.browser)
	if m.options.DefaultURL != "" {
		m.browser.SetDefaultURL(m.options.DefaultURL)
	}
	// Registers the current window to the process message for future use when the specified window handles the message
	ipc.RegisterProcessMessage(m)
	// BrowserWindow Default preset function implementation
	m.defaultEvent()
}

func (m *BrowserWindow) FormAfterCreate(sender lcl.IObject) {
	// add browser window in globalBrowserWindows
	addBrowserWindow(m)
	m._HookWndProcMessage()
}

// FormCreate call afterCreate
func (m *BrowserWindow) afterCreate() { // after
	if m.options.Frameless {

	} else {
		if m.options.DisableResize {
			m.SetBorderStyleForFormBorderStyle(types.BsSingle)
			m.EnabledMaximize(false)
		}
		if m.options.DisableMinimize {
			m.EnabledMinimize(false)
		}
		if m.options.DisableMaximize {
			m.EnabledMaximize(false)
		}
		if m.options.DisableSystemMenu {
			m.EnabledSystemMenu(false)
		}
	}
	constr := m.Constraints()
	if m.options.MaxWidth > 0 || m.options.MaxHeight > 0 {
		constr.SetMaxWidth(m.options.MaxWidth)
		constr.SetMaxHeight(m.options.MaxHeight)
	}
	if m.options.MinWidth > 0 || m.options.MinHeight > 0 {
		constr.SetMinWidth(m.options.MinWidth)
		constr.SetMinHeight(m.options.MinHeight)
	}
	//m.platformCreate()
}

func (m *BrowserWindow) defaultOptions() {
	if m.options.Width <= 0 {
		m.options.Width = 800
	}
	if m.options.Height <= 0 {
		m.options.Height = 600
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

func (m *BrowserWindow) IsClosing() bool {
	return m.isClosing
}

func (m *BrowserWindow) WindowParent() wv.IWVWindowParent {
	return m.windowParent
}

func (m *BrowserWindow) Browser() wv.IWVBrowser {
	return m.browser
}

func (m *BrowserWindow) SetOnShow(fn TNotifyEvent) {
	m.onShow = fn
}

func (m *BrowserWindow) SetOnClose(fn lcl.TCloseEvent) {
	m.onClose = fn
}

func (m *BrowserWindow) SetOnDestroy(fn TNotifyEvent) {
	m.onDestroy = fn
}

func (m *BrowserWindow) SetOnBrowserAfterCreated(fn TNotifyEvent) {
	m.onAfterCreated = fn
}

func (m *BrowserWindow) SetOnContextMenuRequestedEvent(fn TOnContextMenuRequestedEvent) {
	m.onContextMenuRequested = fn
}

func (m *BrowserWindow) SetOnBrowserMessageReceived(fn TOnWebMessageReceivedEvent) {
	m.onWebMessageReceived = fn
}

func (m *BrowserWindow) SetOnNavigationStarting(fn TOnNavigationStartingEvent) {
	m.onNavigationStarting = fn
}

func (m *BrowserWindow) SetOnNewWindowRequestedEvent(fn TOnNewWindowRequestedEventEx) {
	m.onNewWindowRequested = fn
}

func (m *BrowserWindow) Minimize() {
	if m.options.DisableMinimize {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		m.SetWindowState(types.WsMinimized)
	})
}

func (m *BrowserWindow) Maximize() {
	if m.IsFullScreen() || m.options.DisableMaximize {
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
	// In the case of a title bar
	// If the current state is full screen and the extracted state is Ws Maximized,
	// So let's first perform IsFullScreen() judgment here
	if m.IsFullScreen() {
		m.ExitFullScreen()
	} else if m.IsMinimize() || m.IsMaximize() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.SetWindowState(types.WsNormal)
		})
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
