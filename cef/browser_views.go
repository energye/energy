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
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types/colors"
)

// IViewsBrowser extension -> IBrowser, CEF views framework Component browser.
type IViewsBrowser interface {
	lcl.IComponent
	IBrowser
	buildViewsBrowser(owner lcl.IComponent, self IViewsWindow)
	SetOptions(options application.Options)
	Options() *application.Options
}

type TViewsBrowser struct {
	lcl.IComponent
	TBrowser

	self IViewsWindow

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

func (m *TViewsBrowser) buildViewsBrowser(owner lcl.IComponent, self IViewsWindow) {
	m.self = self
	m.IComponent = lcl.NewComponent(owner)
	m.kind = bkViews
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

func (m *TViewsBrowser) initViewsWindowDefaultEvent() {
	m.window.SetOnWindowCreated(m.windowOnWindowCreated)
	m.window.SetOnGetInitialShowState(m.windowOnGetInitialShowState)
	m.window.SetOnGetInitialBounds(m.windowOnGetInitialBounds)
	m.window.SetOnCanClose(m.windowOnCanClose)
	m.window.SetOnWindowClosing(m.windowOnWindowClosing)
}

func (m *TViewsBrowser) windowOnWindowCreated(sender lcl.IObject, window cef.ICefWindow) {
	logger.Debug("Window.OnWindowCreated")
	if m.created {
		return
	}
	m.created = true

	m.UpdateBrowserOptions()
	url := m.defaultURL
	if url == "" {
		url = "about:blank"
	}

	created := m.chromium.CreateBrowserWithStrBVComponentRContextDValue(url, m.browserView, m.context, m.extraInfo)
	logger.Debug("Window.OnWindowCreated chromium.CreateBrowser:", created)
	if !created {
		return
	}

	m.window.SetToFillLayout()
	browserView := m.browserView.BrowserView()
	if browserView != nil && browserView.IsValid() {
		m.window.AddChildView(browserView)
	}

	if m.options != nil {
		if m.options.Caption != "" {
			m.window.SetTitle(m.options.Caption)
		}
		if m.options.Width > 0 && m.options.Height > 0 {
			m.window.CenterWindow(cef.TCefSize{
				Width:  m.options.Width,
				Height: m.options.Height,
			})
		}
	}

	lcl.CallFormCreate(m.self, m)

	m.browserView.RequestFocus()
	m.window.Show()
}

func (m *TViewsBrowser) windowOnGetInitialBounds(sender lcl.IObject, window cef.ICefWindow, result *cef.TCefRect) {
	if m.options == nil || result == nil {
		return
	}
	if m.options.X > 0 {
		result.X = m.options.X
	}
	if m.options.Y > 0 {
		result.Y = m.options.Y
	}
	if m.options.Width > 0 {
		result.Width = m.options.Width
	}
	if m.options.Height > 0 {
		result.Height = m.options.Height
	}
}

func (m *TViewsBrowser) windowOnGetInitialShowState(sender lcl.IObject, window cef.ICefWindow, result *cefTypes.TCefShowState) {
	logger.Debug("Window.OnGetInitialShowState")
}

func (m *TViewsBrowser) windowOnWindowClosing(sender lcl.IObject, window cef.ICefWindow) {
	logger.Debug("Window.OnWindowClosing")
}

func (m *TViewsBrowser) windowOnCanClose(sender lcl.IObject, window cef.ICefWindow, result *bool) {
	logger.Debug("Window.OnCanClose")
	if result == nil {
		return
	}
	*result = m.canClose
	if !m.canClose && m.chromium != nil {
		m.chromium.CloseBrowser(true)
	}
}
