//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux
// +build linux

package cef

import (
	"github.com/energye/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/golcl/lcl/types"
	"runtime"
)

// 窗口拖拽JS扩展
// 在这里执行并启用JS拖拽
func dragExtensionJS(frame *ICefFrame) {
	// MacOS只在LCL窗口中使用自定义窗口拖拽, VF窗口默认已实现
	// 在MacOS中LCL窗口没有有效的消息事件
	var executeJS = `energyExtension.drag.setup();energyExtension.drag.os="` + runtime.GOOS + `";`
	frame.ExecuteJavaScript(executeJS, "", 0)
}

// 窗口拖拽JS扩展处理器
//  1. 注册JS扩展到CEF, 注册鼠标事件，通过本地函数在Go里处理鼠标事件
//  2. 通过IPC将鼠标消息发送到主进程，主进程监听到消息处理鼠标事件
//  3. macos 使用窗口坐标实现窗口拖拽
func dragExtensionHandler() {
	energyExtensionHandler := V8HandlerRef.New()
	energyExtensionHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		if name == mouseUp {
			return true
		} else if name == mouseDown {
			var dx, dy int32
			if arguments.Size() > 0 {
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
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragDown, X: dx, Y: dy},
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
			return true
		} else if name == mouseMove {
			var mx, my int32
			if arguments.Size() > 0 {
				point := arguments.Get(0)
				v8ValX := point.getValueByKey("x")
				v8ValY := point.getValueByKey("y")
				mx = v8ValX.GetIntValue()
				my = v8ValY.GetIntValue()
				v8ValX.Free()
				v8ValY.Free()
				point.Free()
			}
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragMove, X: mx, Y: my},
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
			return true
		} else if name == mouseDblClick {
			// caption double click
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragDblClick},
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
			return true
		}
		return false
	})
	RegisterExtension("energyExtension", string(ipc.IPCJS), energyExtensionHandler)
}

func (m *drag) drag() {
	window := m.window.AsLCLBrowserWindow().BrowserWindow()
	switch m.T {
	case dragUp:
	case dragDown:
		m.dx = m.X
		m.dy = m.Y
		point := m.window.Point()
		m.wx = point.X
		m.wy = point.Y
	case dragMove:
		m.mx = m.X
		m.my = m.Y
		if m.window.IsLCL() {
			x := m.wx + (m.mx - m.dx)
			y := m.wy + (m.my - m.dy)
			m.window.SetPoint(x, y)
		}
	case dragDblClick:
		if window.WindowProperty().EnableWebkitAppRegionDClk {
			if window.WindowState() == types.WsNormal {
				window.SetWindowState(types.WsMaximized)
			} else {
				window.SetWindowState(types.WsNormal)
			}
		}
	}
}
