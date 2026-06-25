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
	"github.com/energye/cef/cef"
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
)

// IVFBrowser extension -> IBrowser, CEF ViewsFramework Component browser.
type IVFBrowser interface {
	IBrowser
}

type TVFBrowser struct {
	lcl.IComponent
	TBrowser
	owner       lcl.IComponent
	window      cef.ICEFWindowComponent
	browserView cef.ICEFBrowserViewComponent
}

func NewVFBrowser(owner lcl.IComponent) IVFBrowser {
	if owner == nil {
		owner = lcl.NewComponent(nil)
	}
	m := &TVFBrowser{}
	m.owner = owner
	m.chromium = cef.NewChromium(m.owner)
	m.window = cef.NewWindowComponent(m.owner)
	m.browserView = cef.NewBrowserViewComponent(m.owner)

	m.window.SetOnWindowCreated(m.windowOnWindowCreated)

	return m
}

// SetWindow sets the window instance for webview and initializes related callback functions
//
//	window - Window interface instance hosting webview content
//	 VF Browser no impl
func (m *TVFBrowser) SetWindow(_ window.IWindow) {}

// SetParent VF Browser no impl
func (m *TVFBrowser) SetParent(_ lcl.IWinControl) {}

// UpdateBrowserOptions updates browser configuration
func (m *TVFBrowser) UpdateBrowserOptions() {

}

func (m *TVFBrowser) CreateBrowser() {

}

// Close closes the webview window and releases associated resources
func (m *TVFBrowser) Close() {

}

// SetDefaultURL sets the default URL for the WebView
func (m *TVFBrowser) SetDefaultURL(url string) {
	if m.defaultURL != url {
		m.chromium.SetDefaultUrl(url)
	}
	m.defaultURL = url
}

// LoadURL loads the specified URL address into the webview
func (m *TVFBrowser) LoadURL(url string) {
	m.chromium.LoadURLWithStrFrame(url, m.chromium.Browser().GetMainFrame())
}

// Browser returns the browser object associated with the TWebview instance
func (m *TVFBrowser) Browser() core.Browser {
	return m.chromium
}

// WindowParent obtains the parent window object associated with the TWebview instance
func (m *TVFBrowser) WindowParent() core.WindowParent {
	return m
}

// SetOnBrowserAfterCreated sets the callback handler triggered after browser creation completes
func (m *TVFBrowser) SetOnBrowserAfterCreated(fn lcl.TNotifyEvent) {
	m.onBrowserAfterCreated = fn
}

// SetOnResourceRequest sets the handler for resource request events
// This method registers a callback function that will be triggered when the webview initiates a resource request
func (m *TVFBrowser) SetOnResourceRequest(fn core.TOnResourceRequestEvent) {
	m.onResourceRequest = fn
}

// SetOnProcessMessage sets the callback function for processing process messages
// This method registers a callback that is triggered when a process message is received
func (m *TVFBrowser) SetOnProcessMessage(fn core.TOnProcessMessageEvent) {
	m.onProcessMessage = fn
}

func (m *TVFBrowser) SetOnLoadChange(fn core.TOnLoadChangeEvent) {
	m.onLoadChange = fn
}

func (m *TVFBrowser) SetOnContextMenu(fn core.TOnContextMenuEvent) {
	m.onContextMenu = fn
}

func (m *TVFBrowser) SetOnContextMenuCommand(fn core.TOnContextMenuCommandEvent) {
	m.onContextMenuCommand = fn
}

func (m *TVFBrowser) SetOnPopupWindow(fn core.TOnPopupWindowEvent) {
	m.onPopupWindow = fn
}

func (m *TVFBrowser) SetOnDragEnter(fn core.TOnDragEnterEvent) {
	m.onDragEnter = fn
}

func (m *TVFBrowser) SetOnDragLeave(fn core.TOnDragLeaveEvent) {
	m.onDragLeave = fn
}

func (m *TVFBrowser) SetOnDragOver(fn core.TOnDragOverEvent) {
	m.onDragOver = fn
}

func (m *TVFBrowser) windowOnWindowCreated(sender lcl.IObject, window cef.ICefWindow) {

}
