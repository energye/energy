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
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
)

// IViewsBrowser extension -> IBrowser, CEF views framework Component browser.
type IViewsBrowser interface {
	lcl.IComponent
	IBrowser
}

type TViewsBrowser struct {
	lcl.IComponent
	TBrowser
	window      cef.ICEFWindowComponent
	browserView cef.ICEFBrowserViewComponent
}

func NewViewsBrowser(owner lcl.IComponent) IViewsBrowser {
	m := &TViewsBrowser{}
	m.IComponent = lcl.NewComponent(owner)
	m.kind = bkViews
	m.chromium = cef.NewChromium(m.IComponent)
	m.window = cef.NewWindowComponent(m.IComponent)
	m.browserView = cef.NewBrowserViewComponent(m.IComponent)

	m.window.SetOnWindowCreated(m.windowOnWindowCreated)

	m.initBrowserDefaultEvent()
	m.initViewsBrowserDefaultEvent()

	return m
}

// SetWindow sets the window instance for webview and initializes related callback functions
//
//	window - Window interface instance hosting webview content
//	 VF Browser no impl
func (m *TViewsBrowser) SetWindow(_ window.IWindow) {}

// SetParent VF Browser no impl
func (m *TViewsBrowser) SetParent(_ lcl.IWinControl) {}

// UpdateBrowserOptions updates browser configuration
func (m *TViewsBrowser) UpdateBrowserOptions() {

}

func (m *TViewsBrowser) CreateBrowser() {

}

// Close closes the webview window and releases associated resources
func (m *TViewsBrowser) Close() {

}

func (m *TViewsBrowser) windowOnWindowCreated(sender lcl.IObject, window cef.ICefWindow) {

}
