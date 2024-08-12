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
	code.WriteString("energyExtension.drag.setup();")
	// current os
	code.WriteString("energyExtension.drag.os='" + runtime.GOOS + "';")
	if window != nil {
		// enable resize
		code.WriteString("energyExtension.drag.enableResize=" + strconv.FormatBool(window.WindowProperty().EnableResize) + ";")
	}
	frame.ExecuteJavaScript(code.String(), "", 0)
}
