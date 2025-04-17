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
	"bytes"
	"github.com/cyber-xxm/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/cyber-xxm/energy/v2/cef/ipc/argument"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"runtime"
	"strconv"
)

const (
	mouseUp       = "mouseUp"
	mouseDown     = "mouseDown"
	mouseMove     = "mouseMove"
	mouseResize   = "mouseResize"
	mouseDblClick = "mouseDblClick"
)

const (
	dragUp = iota
	dragDown
	dragMove
	dragResize
	dragDblClick
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
	m.energyExtensionObject = V8ValueRef.NewObject(nil, nil)
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
	switch name {
	case mouseUp:
		return true
	case mouseDown:
		if common.IsDarwin() {
			return true
		}
		var dx, dy int32
		if !common.IsWindows() && arguments.Size() > 0 {
			point := arguments.Get(0)
			v8ValX := point.getValueByKey("x")
			v8ValY := point.getValueByKey("y")
			dx = v8ValX.GetIntValue()
			dy = v8ValY.GetIntValue()
			v8ValX.Free()
			v8ValY.Free()
			point.Free()
		}
		message := &ipcArgument.List{
			Id:   -1,
			Name: internalIPCDRAG,
			Data: &drag{T: dragDown, X: dx, Y: dy},
		}
		m.sendMessage(message)
		return true
	case mouseMove:
		if common.IsDarwin() {
			return true
		}
		var mx, my, timeStamp int32
		if !common.IsWindows() && arguments.Size() > 0 {
			point := arguments.Get(0)
			v8ValX := point.getValueByKey("x")
			v8ValY := point.getValueByKey("y")
			v8TimeStamp := point.getValueByKey("ts")
			mx = v8ValX.GetIntValue()
			my = v8ValY.GetIntValue()
			timeStamp = v8TimeStamp.GetIntValue()
			v8ValX.Free()
			v8ValY.Free()
			v8TimeStamp.Free()
			point.Free()
		}
		// caption move
		message := &ipcArgument.List{
			Id:   -1,
			Name: internalIPCDRAG,
			Data: &drag{T: dragMove, X: mx, Y: my, TS: timeStamp},
		}
		m.sendMessage(message)
		return true
	case mouseDblClick:
		// caption double click
		message := &ipcArgument.List{
			Id:   -1,
			Name: internalIPCDRAG,
			Data: &drag{T: dragDblClick},
		}
		m.sendMessage(message)
		return true
	case mouseResize:
		// window border resize
		htValue := arguments.Get(0)
		ht := htValue.GetStringValue()
		htValue.Free()
		message := &ipcArgument.List{
			Id:   -1,
			Name: internalIPCDRAG,
			Data: &drag{T: dragResize, HT: ht},
		}
		m.sendMessage(message)
		return true
	}
	return false
}

func (m *dragExtensionHandler) sendMessage(message *ipcArgument.List) {
	v8ctx := V8ContextRef.Current()
	defer v8ctx.Free()
	var processMessage target.IProcessMessage
	if application.Is49() {
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
	var argumentList ipcArgument.IList // json.JSONArray
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

// drag
//
//	custom window drag
//	The second method, Implemented using JavaScript, currently suitable for LCL windows on Windows and Mac OS
//	VF window is already implemented and supported by default
type drag struct {
	T       int8           // data type
	X, Y    int32          // data mouse point
	HT      string         // mouse ht
	TS      int32          // mouse timeStamp
	window  IBrowserWindow // window
	wx, wy  int32          // window point
	dx, dy  int32          // down mouse point
	mx, my  int32          // move mouse point
	canDrag bool
}

// Extension JS
func dragExtensionJS(frame *ICefFrame, window IBrowserWindow) {
	var code = bytes.Buffer{}
	// js drag window
	code.WriteString("__drag.setup();")
	// current os
	code.WriteString("__drag.os='" + runtime.GOOS + "';")
	if window != nil {
		// enable resize
		code.WriteString("__drag.enableResize=" + strconv.FormatBool(window.WindowProperty().EnableResize) + ";")
	}
	frame.ExecuteJavaScript(code.String(), "", 0)
}

// 窗口拖拽JS扩展处理器
//  1. 注册JS扩展到CEF, 注册鼠标事件，通过本地函数在Go里处理鼠标事件
//  2. 通过IPC将鼠标消息发送到主进程，主进程监听到消息处理鼠标事件
//  3. macos 使用窗口坐标实现窗口拖拽
func registerDragExtensionHandler() {
	RegisterExtension("__drag", ipc.IPCJS, nil)
}
