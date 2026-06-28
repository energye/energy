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
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
)

// IEmbeddedBrowser extension -> IBrowser, LCL Component browser.
type IEmbeddedBrowser interface {
	lcl.IWinControl
	IBrowser
	// SetWindow sets the window instance for webview and initializes relevant callback functions
	// window - Window interface instance used to host webview content
	SetWindow(window window.IWindow)
	SetParent(window lcl.IWinControl)
}

type TEmbeddedBrowser struct {
	cef.ICEFWinControl
	TBrowser
	timer lcl.ITimer
}

// NewEmbeddedBrowser creates a new browser window instance
func NewEmbeddedBrowser(owner lcl.IWinControl) IEmbeddedBrowser {
	m := &TEmbeddedBrowser{}
	m.kind = bkEmbedded
	m.chromium = cef.NewChromium(owner)
	if tool.IsWindows() {
		m.ICEFWinControl = cef.NewWindowParent(owner)
	} else {
		windowParent := cef.NewLinkedWindowParent(owner)
		windowParent.SetChromium(m.chromium)
		m.ICEFWinControl = windowParent
	}

	m.chromium.SetWebRTCIPHandlingPolicy(cefTypes.HpDisableNonProxiedUDP)
	m.chromium.SetWebRTCMultipleRoutes(cefTypes.STATE_DISABLED)
	m.chromium.SetWebRTCNonproxiedUDP(cefTypes.STATE_DISABLED)

	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()

	m.initBrowserDefaultEvent()
	m.initEmbeddedBrowserDefaultEvent()

	m.timer = lcl.NewTimer(owner)
	m.timer.SetEnabled(false)
	m.timer.SetInterval(500)
	m.timer.SetOnTimer(m.createBrowserOnTimer)

	return m
}

// SetWindow sets the window instance for webview and initializes related callback functions
//
//	window - Window interface instance hosting webview content
func (m *TEmbeddedBrowser) SetWindow(window window.IWindow) {
	m.window = window
	if m.window != nil {
		window.AddOnWindowStateChange(m.doOnWindowStateChange)
		window.AddOnWindowResize(m.doOnWindowResize)
		window.AddOnWindowShow(m.doOnWindowShow)
		window.AddOnWindowClose(m.doOnWindowClose)
		window.AddOnWindowCloseQuery(m.doOnWindowCloseQuery)
	}
}

// UpdateBrowserOptions updates browser configuration
func (m *TEmbeddedBrowser) UpdateBrowserOptions() {
	// retrieves global LocalLoad configuration
	if GApplication != nil && GApplication.LocalLoad != nil {
		newLocalLoad := *GApplication.LocalLoad.LocalLoad
		m.SetLocalLoad(newLocalLoad)
	}
	if m.window != nil {
		// sets browser configuration
		options := m.window.Options()
		if options.DefaultURL != "" && m.defaultURL == "" {
			m.SetDefaultURL(options.DefaultURL)
		} else if m.defaultURL != "" {
			m.SetDefaultURL(m.defaultURL)
		}
		if options.BackgroundColor != nil {
			r, g, b := byte(options.BackgroundColor.R), byte(options.BackgroundColor.G), byte(options.BackgroundColor.B)
			color := colors.TColor(colors.RGB(r, g, b))
			m.SetColor(color)
		}
	}
}

func (m *TEmbeddedBrowser) CreateBrowser() {
	logger.Debug("Browser.CreateBrowser")
	m.UpdateBrowserOptions()
	m.createBrowserOnTimer(m.timer)
}

// Close closes the webview window and releases associated resources
func (m *TEmbeddedBrowser) Close() {
	if m.isClose {
		return
	}
	m.isClose = true
	m.chromium.TryCloseBrowser()
	ipc.UnRegisterProcessMessage(m)
}

// SetOnDragOver TBrowser.SetOnDragOver
func (m *TEmbeddedBrowser) SetOnDragOver(fn core.TOnDragOverEvent) {
	m.onDragOver = fn
}

func (m *TEmbeddedBrowser) doOnWindowStateChange(sender lcl.IObject) {
}

func (m *TEmbeddedBrowser) doOnWindowResize(sender lcl.IObject) {
	if m.chromium != nil {
		m.chromium.NotifyMoveOrResizeStarted()
		m.UpdateSize()
	}
}

func (m *TEmbeddedBrowser) doOnWindowShow(sender lcl.IObject) {
	logger.Debug("Browser.doOnWindowShow")
	m.CreateBrowser()
}

func (m *TEmbeddedBrowser) doOnWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {
	logger.Debug("Browser.doOnWindowClose")
	*closeAction = types.CaFree
}

func (m *TEmbeddedBrowser) doOnWindowCloseQuery(sender lcl.IObject, canClose *bool) {
	logger.Debug("Browser.doOnWindowCloseQuery canClose:", m.canClose)
	if tool.IsDarwin() {
		*canClose = m.canClose
	} else {
		*canClose = m.canClose
	}
	if !m.canClose {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.chromium.CloseBrowser(true)
		})
	}
}

func (m *TEmbeddedBrowser) createBrowserOnTimer(sender lcl.IObject) {
	if m.timer == nil {
		return
	}
	m.timer.SetEnabled(false)
	rect := m.ClientRect()
	created := m.chromium.CreateBrowserWithWHandleRectStrRContextDValueBool(m.Handle(), rect, m.windowName,
		m.context, m.extraInfo, false)
	init := m.chromium.Initialized()
	logger.Debug("Browser.createBrowserOnTimer created:", created, "init:", init)
	if !created && !init {
		logger.Debug("Browser.createBrowserOnTimer fail")
		m.timer.SetEnabled(true)
	} else {
		logger.Debug("Browser.createBrowserOnTimer success")
		m.UpdateSize()
		m.timer.SetOnTimer(nil)
		m.timer.Free()
		m.timer = nil
	}
}
