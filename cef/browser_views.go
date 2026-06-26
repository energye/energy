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
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types/colors"
)

// IViewsBrowser extension -> IBrowser, CEF views framework Component browser.
type IViewsBrowser interface {
	lcl.IComponent
	IBrowser
	SetOptions(options application.Options)
	Options() *application.Options
}

type TViewsBrowser struct {
	lcl.IComponent
	TBrowser

	options *application.Options

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

func (m *TViewsBrowser) SetOptions(options application.Options) {
	m.options = &options
	if m.options.BackgroundColor == nil {
		m.options.BackgroundColor = &colors.TARGB{R: 255, G: 255, B: 255, A: 255}
	}
}

func (m *TViewsBrowser) Options() *application.Options {
	return m.options
}

// UpdateBrowserOptions updates browser configuration
func (m *TViewsBrowser) UpdateBrowserOptions() {
	// retrieves global LocalLoad configuration
	if GApplication != nil && GApplication.LocalLoad != nil {
		newLocalLoad := *GApplication.LocalLoad.LocalLoad
		m.SetLocalLoad(newLocalLoad)
		// sets browser configuration
		options := GApplication.Options
		if m.options == nil {
			m.options = &options
		}
		if m.options.DefaultURL != "" && m.defaultURL == "" {
			m.SetDefaultURL(m.options.DefaultURL)
		} else if m.defaultURL != "" {
			m.SetDefaultURL(m.defaultURL)
		}
	}
}

func (m *TViewsBrowser) CreateBrowser() {
	m.UpdateBrowserOptions()
}

// Close closes the webview window and releases associated resources
func (m *TViewsBrowser) Close() {

}

func (m *TViewsBrowser) windowOnWindowCreated(sender lcl.IObject, window cef.ICefWindow) {

}
