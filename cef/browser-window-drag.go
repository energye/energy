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

const (
	mouseUp   = "mouseUp"
	mouseDown = "mouseDown"
	mouseMove = "mouseMove"
	dragUp    = 0
	dragDown  = 1
	dragMove  = 2
)

type drag struct {
	T      int8           // 0:up, 1:down, 2:move
	X, Y   int32          // data mouse point
	window IBrowserWindow // drag window
	wx, wy int32          // window point
	dx, dy int32          // down mouse point
	mx, my int32          // move mouse point
}
