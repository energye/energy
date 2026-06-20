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

package cef

import (
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/window"
)

var (
	frameWidth  = int32(4)
	frameHeight = int32(4)
	frameCorner = frameWidth + frameHeight
)

func (m *TBrowser) drag(message ipc.ProcessMessage) {
	window := m.window.(window.ILinuxWindow)
	_ = window
	// todo 待实现
}

func (m *TBrowser) resize(ht string) {
	// todo 待实现
}
