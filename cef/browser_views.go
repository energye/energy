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

type viewsBrowserType int32

const (
	vbtMain viewsBrowserType = iota // main browser window
	vbtSub                          // sub browser window
)

// IViewsBrowser extension -> IBrowser, CEF views framework Component browser.
type IViewsBrowser interface {
	lcl.IComponent
	IBrowser
	buildViewsBrowser(owner lcl.IComponent, self IViewsBrowser)
	SetOptions(options application.Options)
	Options() *application.Options
	CreateTopLevelWindow()
}

type TViewsBrowser struct {
	lcl.IComponent
	TBrowser

	self IViewsBrowser

	browserType viewsBrowserType

	options *application.Options

	window      cef.ICEFWindowComponent
	browserView cef.ICEFBrowserViewComponent
	created     bool
}

func NewViewsBrowser(owner lcl.IComponent) IViewsBrowser {
	m := &TViewsBrowser{}
	m.buildViewsBrowser(owner, m)
	return m
}

func (m *TViewsBrowser) buildViewsBrowser(owner lcl.IComponent, self IViewsBrowser) {
	m.self = self
	m.IComponent = lcl.NewComponent(owner)
	m.kind = bkViews
	browserType := vbtSub
	if len(GApplication.windowList) == 0 {
		browserType = vbtMain
	}
	m.browserType = browserType
	m.chromium = cef.NewChromium(m.IComponent)
	m.window = cef.NewWindowComponent(m.IComponent)
	m.browserView = cef.NewBrowserViewComponent(m.IComponent)
	m.initBrowserDefaultEvent()
	m.initViewsWindowDefaultEvent()
	m.initViewsBrowserDefaultEvent()
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
	m.CreateTopLevelWindow()
}

func (m *TViewsBrowser) CreateTopLevelWindow() {
	if m.created {
		return
	}
	m.UpdateBrowserOptions()
	m.window.CreateTopLevelWindow()
}

// Close closes the webview window and releases associated resources
func (m *TViewsBrowser) Close() {
	if m.isClose {
		return
	}
	m.isClose = true
	m.chromium.TryCloseBrowser()
	m.window.Close()
}
