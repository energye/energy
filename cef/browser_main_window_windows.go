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

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts/messages"
	t "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

func (m *lclBrowserWindow) mainFormNotInTaskBar() {
	m.SetOnWndProc(func(msg *types.TMessage) {
		if !m.WindowProperty().MainFormOnTaskBar && msg.Msg == messages.WM_SHOWWINDOW {
			m.Hide()
			winapi.SetWindowLong(t.HWND(lcl.Application.Handle()), win.GWL_EXSTYLE, win.WS_EX_TOOLWINDOW)
		}
	})
}
