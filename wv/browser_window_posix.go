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

package wv

import (
	"github.com/energye/energy/v3/internal/ipc"
)

func (m *BrowserWindow) platformCreate() {

}

func (m *BrowserWindow) navigationStarting() {

}

func (m *BrowserWindow) Resize(ht string) {

}

func (m *BrowserWindow) Drag(message ipc.ProcessMessage) {
	switch message.Type {
	case ipc.MT_DRAG_MOVE:
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
	}
}

func (m *BrowserWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
}

func (m *BrowserWindow) ExitFullScreen() {
	if m.IsFullScreen() {

	}
}
