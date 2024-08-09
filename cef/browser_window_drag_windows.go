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
	"github.com/energye/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/consts/messages"
	"github.com/energye/golcl/lcl/win"
	"runtime"
)

// 窗口拖拽JS扩展
func dragExtensionJS(frame *ICefFrame) {
	var executeJS = `energyExtension.drag.setup();energyExtension.drag.os="` + runtime.GOOS + `";`
	frame.ExecuteJavaScript(executeJS, "", 0)
}

// 窗口拖拽JS扩展处理器
//  1. 注册JS扩展到CEF, 注册鼠标事件，通过本地函数在Go里处理鼠标事件
//  2. 通过IPC将鼠标消息发送到主进程，主进程监听到消息处理鼠标事件
//  3. windows 使用win.api实现窗口拖拽
func dragExtensionHandler() {
	energyExtensionHandler := V8HandlerRef.New()
	energyExtensionHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		if name == mouseDown || name == mouseUp {
			return true
		} else if name == mouseMove {
			// caption move
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragMove},
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
		} else if name == mouseResize {
			// window border resize
			htValue := arguments.Get(0)
			ht := htValue.GetStringValue()
			htValue.Free()
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragResize, HT: ht},
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
		if window.IsFullScreen() {
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
