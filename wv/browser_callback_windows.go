// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

//go:build windows
// +build windows

package wv

import (
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/wv/wv"
)

type TNotifyEvent = wv.TNotifyEvent
type TOnWebMessageReceivedEvent = wv.TOnWebMessageReceivedEvent
type TOnNavigationStartingEvent = wv.TOnNavigationStartingEvent
type TOnNewWindowRequestedEventEx func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2NewWindowRequestedEventArgs, callback *NewWindowCallback)
type TOnContextMenuRequestedEvent = wv.TOnContextMenuRequestedEvent

type NewWindowCallback struct {
	args    wv.ICoreWebView2NewWindowRequestedEventArgs
	handled bool
	window  IBrowserWindow
}

func (m *NewWindowCallback) NewWindow() IBrowserWindow {
	if m.window == nil {
		m.handled = true
		args := wv.NewCoreWebView2NewWindowRequestedEventArgs(m.args)
		var window = &BrowserWindow{options: Options{DefaultURL: args.URI()}}
		window.newWindowRequestedEventArgs = args
		window.deferral = wv.NewCoreWebView2Deferral(args.Deferral())
		lcl.Application.CreateForm(window)
		window.afterCreate()
		m.window = window
	}
	return m.window
}

func (m *NewWindowCallback) SetHandled(v bool) {
	if m.window != nil {
		m.handled = true
		return
	}
	m.handled = v
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
		// new popup window
		if m.newWindowRequestedEventArgs != nil && m.deferral != nil {
			m.newWindowRequestedEventArgs.SetNewWindow(m.Browser().CoreWebView2().BaseIntf())
			m.newWindowRequestedEventArgs.SetHandled(true)
			m.deferral.Complete()
			m.newWindowRequestedEventArgs.Free()
			m.deferral.Free()
			m.newWindowRequestedEventArgs = nil
			m.deferral = nil
		}
		m.windowParent.UpdateSize()
	})
	m.browser.SetOnNavigationStarting(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2NavigationStartingEventArgs) {
		m.navigationStarting()
		jsCode := `window.energy.drag().setup();`
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
				switch pMessage.Type {
				case ipc.MT_READY:
					// ipc ready
					handle = true
				case ipc.MT_EVENT_GO_EMIT, ipc.MT_EVENT_JS_EMIT, ipc.MT_EVENT_GO_EMIT_CALLBACK, ipc.MT_EVENT_JS_EMIT_CALLBACK:
					// ipc on, emit event
					handle = m.ipcMessageReceivedDelegate.Received(m.BrowserId(), &pMessage)
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
	m.browser.SetOnNewWindowRequested(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2NewWindowRequestedEventArgs) {
		callback := &NewWindowCallback{args: args}
		if m.onNewWindowRequested != nil {
			m.onNewWindowRequested(sender, webview, args, callback)
		}
		// true: handle new windows on your own
		if callback.handled {
			if callback.window != nil {
				callback.window.Show()
			} else {
				args = wv.NewCoreWebView2NewWindowRequestedEventArgs(args)
				args.SetHandled(true) // Prevent new windows from popping up
			}
		} else {
			// Default pop-up window
			lcl.RunOnMainThreadAsync(func(id uint32) {
				callback.NewWindow().Show()
			})
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
		if m.options.DefaultWindowStatus == types.WsFullScreen {
			m.FullScreen()
		} else {
			m.SetWindowState(m.options.DefaultWindowStatus)
		}
		if m.onShow != nil {
			m.onShow(sender)
		}
	})
	// window, OnClose
	m.TForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		if m.onClose != nil {
			m.onClose(sender, action)
		} else {
			*action = types.CaFree
			m.WindowParent().Free()
		}
		// window close and free
		if *action == types.CaFree {
			m.isClosing = true
		}
	})
	m.TForm.SetOnDestroy(func(sender lcl.IObject) {
		fmt.Println("SetOnDestroy")
		m._RestoreWndProc()
		deleteBrowserWindow(m)
		// cancel process message
		ipc.UnRegisterProcessMessage(m)
		if m.onDestroy != nil {
			m.onDestroy(sender)
		}
	})
}
