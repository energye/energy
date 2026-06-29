//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

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
	getWindow := func() window.IDarwinWindow {
		if m.window == nil {
			return nil
		}
		return m.window.(window.IDarwinWindow)
	}
	if currentWindow := getWindow(); currentWindow != nil {
		currentWindow.NSWindow().Drag()
	}
}

func (m *TBrowser) resize(ht string) {

}

func (m *TViewsBrowser) startThemeObserver() {

}
