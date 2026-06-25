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
	"encoding/json"
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
	"unsafe"
)

// ILCLBrowser extension -> IBrowser, LCL Component browser.
type ILCLBrowser interface {
	lcl.IWinControl
	IBrowser
}
type TLCLBrowser struct {
	cef.ICEFWinControl
	TBrowser
	timer lcl.ITimer
}

// NewLCLBrowser creates a new browser window instance
func NewLCLBrowser(owner lcl.IWinControl) ILCLBrowser {
	m := &TLCLBrowser{}
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
	m.initLCLBrowserDefaultEvent()

	m.timer = lcl.NewTimer(owner)
	m.timer.SetEnabled(false)
	m.timer.SetInterval(500)
	m.timer.SetOnTimer(m.createBrowserOnTimer)

	return m
}

// SetWindow sets the window instance for webview and initializes related callback functions
//
//	window - Window interface instance hosting webview content
func (m *TLCLBrowser) SetWindow(window window.IWindow) {
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
func (m *TLCLBrowser) UpdateBrowserOptions() {
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

func (m *TLCLBrowser) CreateBrowser() {
	logger.Debug("Browser.CreateBrowser")
	m.UpdateBrowserOptions()
	m.createBrowserOnTimer(m.timer)
}

// Close closes the webview window and releases associated resources
func (m *TLCLBrowser) Close() {
	if m.isClose {
		return
	}
	m.isClose = true
	m.chromium.TryCloseBrowser()
	ipc.UnRegisterProcessMessage(m)
}

// SetDefaultURL sets the default URL for the WebView
func (m *TLCLBrowser) SetDefaultURL(url string) {
	if m.defaultURL != url {
		m.chromium.SetDefaultUrl(url)
	}
	m.defaultURL = url
}

// LoadURL loads the specified URL address into the webview
func (m *TLCLBrowser) LoadURL(url string) {
	m.chromium.LoadURLWithStrFrame(url, m.chromium.Browser().GetMainFrame())
}

// Browser returns the browser object associated with the TWebview instance
func (m *TLCLBrowser) Browser() core.Browser {
	return m.chromium
}

// WindowParent obtains the parent window object associated with the TWebview instance
func (m *TLCLBrowser) WindowParent() core.WindowParent {
	return m
}

// SetOnBrowserAfterCreated sets the callback handler triggered after browser creation completes
func (m *TLCLBrowser) SetOnBrowserAfterCreated(fn lcl.TNotifyEvent) {
	m.onBrowserAfterCreated = fn
}

// SetOnResourceRequest sets the handler for resource request events
// This method registers a callback function that will be triggered when the webview initiates a resource request
func (m *TLCLBrowser) SetOnResourceRequest(fn core.TOnResourceRequestEvent) {
	m.onResourceRequest = fn
}

// SetOnProcessMessage sets the callback function for processing process messages
// This method registers a callback that is triggered when a process message is received
func (m *TLCLBrowser) SetOnProcessMessage(fn core.TOnProcessMessageEvent) {
	m.onProcessMessage = fn
}

func (m *TLCLBrowser) SetOnLoadChange(fn core.TOnLoadChangeEvent) {
	m.onLoadChange = fn
}

func (m *TLCLBrowser) SetOnContextMenu(fn core.TOnContextMenuEvent) {
	m.onContextMenu = fn
}

func (m *TLCLBrowser) SetOnContextMenuCommand(fn core.TOnContextMenuCommandEvent) {
	m.onContextMenuCommand = fn
}

func (m *TLCLBrowser) SetOnPopupWindow(fn core.TOnPopupWindowEvent) {
	m.onPopupWindow = fn
}

func (m *TLCLBrowser) SetOnDragEnter(fn core.TOnDragEnterEvent) {
	m.onDragEnter = fn
}

func (m *TLCLBrowser) SetOnDragLeave(fn core.TOnDragLeaveEvent) {
	m.onDragLeave = fn
}

func (m *TLCLBrowser) SetOnDragOver(fn core.TOnDragOverEvent) {
	m.onDragOver = fn
}

func (m *TLCLBrowser) doOnWindowStateChange(sender lcl.IObject) {
}

func (m *TLCLBrowser) doOnWindowResize(sender lcl.IObject) {
	if m.chromium != nil {
		m.chromium.NotifyMoveOrResizeStarted()
		m.UpdateSize()
	}
}

func (m *TLCLBrowser) doOnWindowShow(sender lcl.IObject) {
	logger.Debug("Browser.doOnWindowShow")
	m.CreateBrowser()
}

func (m *TLCLBrowser) doOnWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {
	logger.Debug("Browser.doOnWindowClose")
	*closeAction = types.CaFree
}

func (m *TLCLBrowser) doOnWindowCloseQuery(sender lcl.IObject, canClose *bool) {
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

func (m *TLCLBrowser) createBrowserOnTimer(sender lcl.IObject) {
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

func (m *TLCLBrowser) dragDrop(message ipc.ProcessMessage, args cef.ICefListValue) {
	data, ok := message.Data.(map[string]any)
	if !ok {
		return
	}
	dataType := core.TDragType(tool.ToInt(data["type"]))
	x := int32(tool.ToInt(data["x"]))
	y := int32(tool.ToInt(data["y"]))
	switch message.Type {
	case ipc.MT_DRAG_DROP_ENTER:
		if m.onDragEnter != nil {
			m.onDragEnter(dataType, x, y)
		}
	case ipc.MT_DRAG_DROP_LEAVE:
		if m.onDragLeave != nil {
			m.onDragLeave()
		}
	case ipc.MT_DRAG_DROP_OVER:
		if m.onDragOver != nil {
			dragData := &core.TDragData{Type: dataType}
			if dataType == core.DragTypeData {
				dragData.Data = []byte(data["text"].(string))
			} else if dataType == core.DragTypeFile {
				objectsDataBin := args.GetBinary(1)
				defer func() {
					objectsDataBin.Release()
				}()
				objectsDataBytes := make([]byte, int(objectsDataBin.GetSize()))
				objectsDataBin.GetData(uintptr(unsafe.Pointer(&objectsDataBytes[0])), objectsDataBin.GetSize(), 0)

				var (
					objectFiles []tObjectFile
					files       []string
				)
				err := json.Unmarshal(objectsDataBytes, &objectFiles)
				if err == nil && m.dragFilePathCache != nil {
					for _, file := range objectFiles {
						if filePath, ok := m.dragFilePathCache[file.Name]; ok {
							files = append(files, filePath)
						}
					}
				}

				dragData.Filenames = files
				m.dragFilePathCache = nil // clear
			}
			m.onDragOver(dragData, x, y)
		}
	}
}
