package cef

import (
	"encoding/json"
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/types/messages"
	"unsafe"
)

func (m *TEmbeddedBrowser) iniTEmbeddedBrowserDefaultEvent() {
	logger.Debug("Browser.iniTEmbeddedBrowserDefaultEvent")

	m.chromium.SetOnProcessMessageReceived(m.chromiumOnProcessMessageReceived)

	m.chromium.SetOnAfterCreated(m.chromiumOnAfterCreated)
	m.chromium.SetOnBeforeBrowse(m.chromiumOnBeforeBrowse)

	//m.chromium.SetOnDraggableRegionsChanged(m.chromiumOnDraggableRegionsChanged)

	// close browser
	m.chromium.SetOnBeforeClose(m.chromiumOnBeforeClose)
	m.chromium.SetOnClose(m.chromiumOnClose)

	m.ICEFWinControl.SetOnEnter(m.winControlOnEnter)
	m.ICEFWinControl.SetOnExit(m.winControlOnExit)
}

func (m *TEmbeddedBrowser) winControlOnEnter(sender lcl.IObject) {
	m.chromium.Initialized()
	m.chromium.FrameIsFocused()
	m.chromium.SetFocus(true)
}

func (m *TEmbeddedBrowser) winControlOnExit(sender lcl.IObject) {
	m.chromium.SendCaptureLostEvent()
}

func (m *TEmbeddedBrowser) chromiumOnProcessMessageReceived(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, sourceProcess cefTypes.TCefProcessId,
	message cef.ICefProcessMessage, outResult *bool) {
	name := message.GetName()
	logger.Debug("Chromium.OnProcessMessageReceived name:", name)
	defer func() {
		message.Release()
	}()
	if name == internalPostMessageName {
		var handle bool
		messageData := ""
		args := message.GetArgumentList()
		dataBin := args.GetBinary(0)
		defer func() {
			dataBin.Release()
			args.Release()
		}()
		messageDataBytes := make([]byte, int(dataBin.GetSize()))
		dataBin.GetData(uintptr(unsafe.Pointer(&messageDataBytes[0])), dataBin.GetSize(), 0)
		messageData = string(messageDataBytes)
		if m.messageReceivedDelegate != nil {
			// ipc message
			var pMessage ipc.ProcessMessage
			err := json.Unmarshal(messageDataBytes, &pMessage)
			if err == nil {
				switch pMessage.Type {
				case ipc.MT_READY:
					// ipc ready
					handle = true
				case ipc.MT_EVENT_GO_EMIT, ipc.MT_EVENT_JS_EMIT, ipc.MT_EVENT_GO_EMIT_CALLBACK, ipc.MT_EVENT_JS_EMIT_CALLBACK:
					// ipc on, emit event
					handle = m.messageReceivedDelegate.Received(m.BrowserId(), &pMessage)
				case ipc.MT_DRAG_MOVE, ipc.MT_DRAG_DOWN, ipc.MT_DRAG_UP, ipc.MT_DRAG_DBLCLICK:
					// ipc drag window
					if m.window != nil {
						m.drag(pMessage)
						handle = true
					}
				case ipc.MT_DRAG_RESIZE:
					// border drag resize
					if m.window != nil {
						ht := pMessage.Data.(string)
						m.resize(ht)
						handle = true
					}
				case ipc.MT_DRAG_BORDER_WMSZ:
				case ipc.MT_DRAG_DROP_ENTER, ipc.MT_DRAG_DROP_LEAVE, ipc.MT_DRAG_DROP_OVER:
					m.dragDrop(pMessage, args)
				}
			} else {
				println("MessageReceived-ERROR：", err.Error())
			}
		}
		logger.Debug("Chromium.OnProcessMessageReceived messageData:", messageData)
		if !handle && m.onProcessMessage != nil {
			m.onProcessMessage(messageData)
		}
		*outResult = handle
	} else if name == internalExecuteScriptResultName {
		args := message.GetArgumentList()
		dataBin := args.GetBinary(0)
		defer func() {
			dataBin.Release()
			args.Release()
		}()
		messageDataBytes := make([]byte, int(dataBin.GetSize()))
		dataBin.GetData(uintptr(unsafe.Pointer(&messageDataBytes[0])), dataBin.GetSize(), 0)
		executeScriptResult := tExecuteScriptResultMessage{}
		_ = json.Unmarshal(messageDataBytes, &executeScriptResult)
		executionID := executeScriptResult.Id
		if callback, ok := m.executeScriptCallback.Load(executionID); ok {
			m.executeScriptCallback.Delete(executionID)
			callback.(core.TOnEvaluateScriptCallbackEvent)(executeScriptResult.Data, executeScriptResult.Error)
		}
		*outResult = true
	}
}

func (m *TEmbeddedBrowser) chromiumOnAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnAfterCreated", browser.GetIdentifier())
	if m.window != nil && m.window.BrowserId() == 0 {
		options := m.window.Options()
		m.browserId = uint32(browser.GetIdentifier())
		m.window.SetBrowserId(m.browserId)
		// ipc
		ipc.RegisterProcessMessage(m)
		// local load
		//m.schemeHandlerFactory = createSchemeHandlerFactory(browser)
		// pre-creates a window
		if options.AutoPopupWindow {
			if gPrePopupWindow == nil {
				lcl.RunOnMainThreadAsync(func(id uint32) {
					gPrePopupWindow = NewPopupWindow()
				})
			}
		}
	}
	if m.onBrowserAfterCreated != nil {
		m.onBrowserAfterCreated(sender)
	}
}

func (m *TEmbeddedBrowser) chromiumOnBeforeBrowse(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest, userGesture bool,
	isRedirect bool, outResult *bool) {
	logger.Debug("Chromium.OnBeforeBrowse")
	m.UpdateSize()
}

//func (m *TEmbeddedBrowser) chromiumOnStartDragging(sender lcl.IObject, browser cef.ICefBrowser, dragData cef.ICefDragData, allowedOps cefTypes.TCefDragOperations,
//	X int32, Y int32, outResult *bool) {
//	logger.Debug("Chromium.OnStartDragging", browser.GetIdentifier())
//}

func (m *TEmbeddedBrowser) chromiumOnDraggableRegionsChanged(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, regionsCount cefTypes.NativeUInt,
	regions cef.ICefDraggableRegionArray) {
	logger.Debug("Chromium.OnDraggableRegionsChanged", browser.GetIdentifier())

}

func (m *TEmbeddedBrowser) chromiumOnClose(sender lcl.IObject, browser cef.ICefBrowser, action *cefTypes.TCefCloseBrowserAction) {
	logger.Debug("Chromium.OnClose", browser.GetIdentifier())
	if tool.IsDarwin() {
		ok := m.DestroyChildWindow()
		logger.Debug("Chromium.OnClose => winControl.DestroyChildWindow() Success:", ok)
		*action = cefTypes.CbaClose
	} else if tool.IsLinux() {
		*action = cefTypes.CbaClose
	} else if tool.IsWindows() {
		*action = cefTypes.CbaDelay
	}
	if tool.IsWindows() || tool.IsLinux() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.ICEFWinControl.Free()
		})
	}
}

func (m *TEmbeddedBrowser) chromiumOnBeforeClose(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnBeforeClose", "Current-BrowserID:", m.browserId, "Target-BrowserID:", browser.GetIdentifier())
	if m.browserId != uint32(browser.GetIdentifier()) {
		logger.Debug("Chromium.OnBeforeClose Non-current user browser")
		return
	}
	closeWindow := func() {
		if m.window != nil {
			m.canClose = true
			if tool.IsWindows() {
				rtl.PostMessage(m.window.Handle(), messages.WM_CLOSE, 0, 0)
			} else if tool.IsDarwin() || tool.IsLinux() {
				m.window.Close()
			}
		} else {
			logger.Error("Browser associated window is nil, failed to close browser")
		}
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		closeWindow()
	})
}
