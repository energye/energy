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
	"github.com/energye/wv/wv"
)

func (m *BrowserWindow) Drag(message ipc.ProcessMessage) {
	switch message.Type {
	case ipc.MT_DRAG_MOVE:
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
	}
}

func (m *BrowserWindow) _HookWndProcMessage() {

}

func (m *BrowserWindow) _RestoreWndProc() {

}

func (m *BrowserWindow) navigationStarting(webview wv.ICoreWebView2, args wv.ICoreWebView2NavigationStartingEventArgs) {

}
