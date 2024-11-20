//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	ipcArgument "github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/cef/ipc/target"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/consts/messages"
	"github.com/energye/golcl/lcl/win"
)

var dragExtension *dragExtensionHandler

type dragExtensionHandler struct {
	energyExtensionObject *ICefV8Value // ipc object
}

func (m *dragExtensionHandler) clear(frameId string) {
	if m.energyExtensionObject != nil {
		m.energyExtensionObject.Free()
		m.energyExtensionObject = nil
	}
}

func (m *dragExtensionHandler) make(frameId string, context *ICefV8Context) {
	if m.energyExtensionObject != nil {
		m.clear(frameId)
	}
	m.energyExtensionObject = V8ValueRef.NewObject(nil)
	handler := V8HandlerRef.New()
	handler.Execute(m.handler)
	m.energyExtensionObject.setValueByKey("mouseUp", V8ValueRef.newFunction("mouseUp", handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.energyExtensionObject.setValueByKey("mouseDown", V8ValueRef.newFunction("mouseDown", handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.energyExtensionObject.setValueByKey("mouseMove", V8ValueRef.newFunction("mouseMove", handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.energyExtensionObject.setValueByKey("mouseResize", V8ValueRef.newFunction("mouseResize", handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.energyExtensionObject.setValueByKey("mouseDblClick", V8ValueRef.newFunction("mouseDblClick", handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	context.Global().setValueByKey(internalEnergyExtension, m.energyExtensionObject, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}

// 窗口拖拽JS扩展处理器
//  1. 注册JS扩展到CEF, 注册鼠标事件，通过本地函数在Go里处理鼠标事件
//  2. 通过IPC将鼠标消息发送到主进程，主进程监听到消息处理鼠标事件
//  3. windows 使用win.api实现窗口拖拽
func (m *dragExtensionHandler) handler(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
	if name == mouseDown || name == mouseUp {
		return true
	} else if name == mouseMove {
		// caption move
		message := &ipcArgument.List{
			Id:   -1,
			Name: internalIPCDRAG,
			Data: &drag{T: dragMove},
		}
		m.sendMessage(message)
		return true
	} else if name == mouseDblClick {
		// caption double click
		message := &ipcArgument.List{
			Id: -1,
			//BId:  ipc.RenderChan().BrowserId(),
			Name: internalIPCDRAG,
			Data: &drag{T: dragDblClick},
		}
		m.sendMessage(message)
		//ipc.RenderChan().IPC().Send(message.Bytes())
		return true
	} else if name == mouseResize {
		// window border resize
		htValue := arguments.Get(0)
		ht := htValue.GetStringValue()
		htValue.Free()
		message := &ipcArgument.List{
			Id: -1,
			//BId:  ipc.RenderChan().BrowserId(),
			Name: internalIPCDRAG,
			Data: &drag{T: dragResize, HT: ht},
		}
		m.sendMessage(message)
		//ipc.RenderChan().IPC().Send(message.Bytes())
		return true
	}
	return false
}

func (m *dragExtensionHandler) sendMessage(message *ipcArgument.List) {
	v8ctx := V8ContextRef.Current()
	defer v8ctx.Free()
	var processMessage target.IProcessMessage
	if application.IsSpecVer49() {
		// CEF49
		processMessage = v8ctx.Browser()
	} else {
		processMessage = v8ctx.Frame()
	}
	//message.BId = v8ctx.Browser().Identifier()
	processMessage.SendProcessMessageForJSONBytes(internalEnergyExtension, consts.PID_BROWSER, message.Bytes())
}

func (m dragExtensionHandler) drag(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) {
	argumentListBytes := message.ArgumentList().GetBinary(0)
	if argumentListBytes == nil {
		return
	}
	var messageDataBytes []byte
	if argumentListBytes.IsValid() {
		size := argumentListBytes.GetSize()
		messageDataBytes = make([]byte, size)
		c := argumentListBytes.GetData(messageDataBytes, 0)
		argumentListBytes.Free()
		if c == 0 {
			return
		}
	}
	var (
		argumentList ipcArgument.IList // json.JSONArray
	)
	if messageDataBytes != nil {
		argumentList = ipcArgument.UnList(messageDataBytes)
		messageDataBytes = nil
	}
	if argumentList != nil && argumentList.GetName() == internalIPCDRAG {
		if wi := BrowserWindow.GetWindowInfo(browser.BrowserId()); wi != nil {
			if wi.IsLCL() {
				window := wi.AsLCLBrowserWindow().BrowserWindow()
				dataJSON := argumentList.JSON()
				if dataJSON != nil {
					if window.drag == nil {
						window.drag = &drag{
							window: wi,
						}
					}
					object := dataJSON.JSONObject()
					window.drag.T = int8(object.GetIntByKey("T"))
					window.drag.X = int32(object.GetIntByKey("X"))
					window.drag.Y = int32(object.GetIntByKey("Y"))
					window.drag.HT = object.GetStringByKey("HT")
					window.drag.TS = int32(object.GetIntByKey("TS"))
				}
				RunOnMainThread(func() {
					window.doDrag()
				})
			}
		}
	}
}

func (m *drag) drag() {
	window := m.window.AsLCLBrowserWindow().BrowserWindow()
	switch m.T {
	case dragUp:
	case dragDown:
	case dragMove:
		// 全屏时不能拖拽窗口
		if window.IsFullScreen() {
			return
		}
		// 此时是 down 事件, 拖拽窗口
		if win.ReleaseCapture() {
			win.PostMessage(m.window.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case dragResize:
		if window.IsFullScreen() || !window.WindowProperty().EnableResize {
			return
		}
		var borderHT uintptr
		switch m.HT {
		case "n-resize":
			borderHT = messages.HTTOP
		case "ne-resize":
			borderHT = messages.HTTOPRIGHT
		case "e-resize":
			borderHT = messages.HTRIGHT
		case "se-resize":
			borderHT = messages.HTBOTTOMRIGHT
		case "s-resize":
			borderHT = messages.HTBOTTOM
		case "sw-resize":
			borderHT = messages.HTBOTTOMLEFT
		case "w-resize":
			borderHT = messages.HTLEFT
		case "nw-resize":
			borderHT = messages.HTTOPLEFT
		}
		if borderHT != 0 {
			if win.ReleaseCapture() {
				win.PostMessage(window.Handle(), messages.WM_NCLBUTTONDOWN, borderHT, 0)
			}
		}
	case dragDblClick:
		if window.WindowProperty().EnableWebkitAppRegionDClk {
			window.Maximize()
		}
	}

}
