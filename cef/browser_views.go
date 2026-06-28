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
	"github.com/energye/energy/v3/core"
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
	//RawWindowComponent() cef.ICEFWindowComponent
	//RawBrowserViewComponent() cef.ICEFBrowserViewComponent
	Close()
	IsClose() bool
	Show()
	Hide()
	IsMain() bool
	Minimize()
	Maximize()
	IsMinimize() bool
	IsMaximize() bool
	IsFullScreen() bool
	FullScreen()
	ExitFullScreen()
	SetIsAlwaysOnTop(value bool)
	SetOnActivate(fn lcl.TNotifyEvent)
	SetOnResize(fn lcl.TNotifyEvent)
	SetOnThemeChange(fn core.TOnThemeChange)
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
	isFirstShow bool
	bounds      cef.TCefRect

	onActivate    lcl.TNotifyEvent
	onResize      lcl.TNotifyEvent
	onThemeChange core.TOnThemeChange
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

func (m *TViewsBrowser) SetOptions(options application.Options) {
	m.options = &options
	if m.options.BackgroundColor == nil {
		m.options.BackgroundColor = &colors.TARGB{R: 255, G: 255, B: 255, A: 255}
	}
}

func (m *TViewsBrowser) Options() *application.Options {
	return m.options
}

func (m *TViewsBrowser) CenterWindow() {
	if m.bounds.Width > 0 && m.bounds.Height > 0 {
		m.window.CenterWindow(cef.TCefSize{
			Width:  m.bounds.Width,
			Height: m.bounds.Height,
		})
	}
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
		if m.options.Width == 0 {
			m.options.Width = 800
		}
		if m.options.Height == 0 {
			m.options.Height = 600
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

//func (m *TViewsBrowser) RawWindowComponent() cef.ICEFWindowComponent {
//	return m.window
//}
//
//func (m *TViewsBrowser) RawBrowserViewComponent() cef.ICEFBrowserViewComponent {
//	return m.browserView
//}

// Close closes the webview window and releases associated resources
func (m *TViewsBrowser) Close() {
	cef.RunOnMainThread(func() {
		m.window.Close()
	})
}

func (m *TViewsBrowser) IsClose() bool {
	return m.isClose
}

func (m *TViewsBrowser) Show() {
	m.window.Show()
}

func (m *TViewsBrowser) Hide() {
	m.window.Hide()
}

func (m *TViewsBrowser) IsMain() bool {
	return m.browserType == vbtMain
}

func (m *TViewsBrowser) Minimize() {
	m.window.Minimize()
}

func (m *TViewsBrowser) Maximize() {
	m.window.Maximize()
}

func (m *TViewsBrowser) IsMinimize() bool {
	return m.window.IsMinimized()
}

func (m *TViewsBrowser) IsMaximize() bool {
	return m.window.IsMaximized()
}

func (m *TViewsBrowser) IsFullScreen() bool {
	return m.window.IsFullscreen()
}

func (m *TViewsBrowser) FullScreen() {
	m.window.SetIsFullscreen(true)
}

func (m *TViewsBrowser) ExitFullScreen() {
	m.window.SetIsFullscreen(false)
}

func (m *TViewsBrowser) SetIsAlwaysOnTop(value bool) {
	cef.RunOnMainThread(func() {
		m.window.SetIsAlwaysOnTop(value)
	})
}

func (m *TViewsBrowser) SetOnActivate(fn lcl.TNotifyEvent) {
	m.onActivate = fn
}

func (m *TViewsBrowser) SetOnResize(fn lcl.TNotifyEvent) {
	m.onResize = fn
}

func (m *TViewsBrowser) SetOnThemeChange(fn core.TOnThemeChange) {
	m.onThemeChange = fn
}
