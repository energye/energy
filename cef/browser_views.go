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
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
	"unsafe"
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
	Restore()
	IsMinimize() bool
	IsMaximize() bool
	IsFullScreen() bool
	FullScreen()
	ExitFullScreen()
	SetIsAlwaysOnTop(value bool)
	SetLeft(v int32)
	SetTop(v int32)
	SetWidth(v int32)
	SetHeight(v int32)
	SetBoundsRect(rect types.TRect)
	BoundsRect() (rect types.TRect)
	SetTitle(title string)
	Title() string
	SetIcon(pngIconData []byte)
	SetWindowState(value types.TWindowState)
	WindowState() types.TWindowState
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

	onActivate lcl.TNotifyEvent
	onResize   lcl.TNotifyEvent
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

	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()

	m.initBrowserDefaultEvent()
	m.initViewsWindowDefaultEvent()
	m.initViewsBrowserDefaultEvent()
	m.initViewsBrowserTagsDefaultEvent()
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
	cef.RunOnMainThread(func() {
		m.window.Show()
	})
}

func (m *TViewsBrowser) Hide() {
	cef.RunOnMainThread(func() {
		m.window.Hide()
	})
}

func (m *TViewsBrowser) IsMain() bool {
	return m.browserType == vbtMain
}

func (m *TViewsBrowser) Minimize() {
	cef.RunOnMainThread(func() {
		m.window.Minimize()
	})
}

func (m *TViewsBrowser) Maximize() {
	cef.RunOnMainThread(func() {
		m.window.Maximize()
	})
}

func (m *TViewsBrowser) Restore() {
	cef.RunOnMainThread(func() {
		m.window.Restore()
	})
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
	cef.RunOnMainThread(func() {
		m.window.SetIsFullscreen(true)
	})
}

func (m *TViewsBrowser) ExitFullScreen() {
	cef.RunOnMainThread(func() {
		m.window.SetIsFullscreen(false)
	})
}

func (m *TViewsBrowser) SetIsAlwaysOnTop(value bool) {
	cef.RunOnMainThread(func() {
		m.window.SetIsAlwaysOnTop(value)
	})
}

func (m *TViewsBrowser) SetLeft(v int32) {
	cef.RunOnMainThread(func() {
		pos := m.window.Position()
		pos.X = v
		m.window.SetPosition(pos)
	})
}

func (m *TViewsBrowser) SetTop(v int32) {
	cef.RunOnMainThread(func() {
		pos := m.window.Position()
		pos.Y = v
		m.window.SetPosition(pos)
	})
}

func (m *TViewsBrowser) SetWidth(v int32) {
	cef.RunOnMainThread(func() {
		bounds := m.window.Bounds()
		bounds.Width = v
		m.window.SetBounds(bounds)
	})
}

func (m *TViewsBrowser) SetHeight(v int32) {
	cef.RunOnMainThread(func() {
		bounds := m.window.Bounds()
		bounds.Height = v
		m.window.SetBounds(bounds)
	})
}

func (m *TViewsBrowser) SetBoundsRect(rect types.TRect) {
	cef.RunOnMainThread(func() {
		bounds := cef.TCefRect{}
		bounds.X = rect.Left
		bounds.Y = rect.Top
		bounds.Width = rect.Width()
		bounds.Height = rect.Height()
		m.window.SetBounds(bounds)
	})
}

func (m *TViewsBrowser) BoundsRect() (rect types.TRect) {
	bounds := m.window.Bounds()
	rect.Left = bounds.X
	rect.Top = bounds.Y
	rect.SetWidth(bounds.Width)
	rect.SetHeight(bounds.Height)
	return
}

func (m *TViewsBrowser) SetTitle(title string) {
	cef.RunOnMainThread(func() {
		m.window.SetTitle(title)
	})
}

func (m *TViewsBrowser) Title() string {
	return m.window.Title()
}

func (m *TViewsBrowser) SetIcon(pngIconData []byte) {
	if !m.created || len(pngIconData) == 0 {
		return
	}
	pngData := uintptr(unsafe.Pointer(&pngIconData[0]))
	pngDataSize := types.NativeUInt(len(pngIconData))
	cefImage := cef.ImageRef.New()
	cefImage.AddPng(1, pngData, pngDataSize)
	m.window.SetWindowAppIcon(cefImage)
}

func (m *TViewsBrowser) SetWindowState(value types.TWindowState) {
	switch m.options.DefaultWindowStatus {
	case types.WsMinimized:
		m.Minimize()
	case types.WsMaximized:
		m.Maximize()
	case types.WsFullScreen:
		if m.IsFullScreen() {
			m.ExitFullScreen()
		} else {
			m.FullScreen()
		}
	default:
		m.Restore()
	}
}

func (m *TViewsBrowser) WindowState() types.TWindowState {
	if m.IsMinimize() {
		return types.WsMinimized
	} else if m.IsMaximize() {
		return types.WsMaximized
	} else if m.IsFullScreen() {
		return types.WsFullScreen
	}
	return types.WsNormal
}

func (m *TViewsBrowser) AsViews() IViewsBrowser {
	return m
}

func (m *TViewsBrowser) AsEmbedded() IEmbeddedBrowser {
	return nil
}

func (m *TViewsBrowser) SetOnActivate(fn lcl.TNotifyEvent) {
	m.onActivate = fn
}

func (m *TViewsBrowser) SetOnResize(fn lcl.TNotifyEvent) {
	m.onResize = fn
}
