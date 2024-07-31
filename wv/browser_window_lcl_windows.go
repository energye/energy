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

package wv

import (
	"fmt"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/messages"
)

func (m *BrowserWindow) Drag(message ipc.ProcessMessage) {
	fmt.Println("drag", message.Type)
	switch message.Type {
	case ipc.MT_DRAG_MOVE:
		if m.WindowState() == types.WsFullScreen {
			return
		}
		if win.ReleaseCapture() {
			win.PostMessage(m.Handle(), messages.WM_NCLBUTTONDOWN, messages.HTCAPTION, 0)
		}
	case ipc.MT_DRAG_DOWN:
	case ipc.MT_DRAG_UP:
	case ipc.MT_DRAG_DBLCLICK:
		if win.ReleaseCapture() {
			if m.WindowState() == types.WsNormal {
				win.PostMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_MAXIMIZE, 0)
			} else {
				win.SendMessage(m.Handle(), messages.WM_SYSCOMMAND, messages.SC_RESTORE, 0)
			}
		}
	}
}
