//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

// TODO no
type customWindowCaption struct {
	bw      *LCLBrowserWindow
	regions *TCefDraggableRegions
}

func (m *customWindowCaption) free() {
	//TODO no
}
