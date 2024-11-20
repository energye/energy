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
	"github.com/energye/energy/v2/cef/internal/ipc"
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

// drag
//
//	custom window drag
//	The second method, Implemented using JavaScript, currently suitable for LCL windows on Windows and Mac OS
//	VF window is already implemented and supported by default
type drag struct {
	T      int8           // data type
	X, Y   int32          // data mouse point
	HT     string         // mouse ht
	TS     int32          // mouse timeStamp
	window IBrowserWindow // drag window
	wx, wy int32          // window point
	dx, dy int32          // down mouse point
	mx, my int32          // move mouse point
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
	RegisterExtension("__drag", string(ipc.IPCJS), nil)
}
