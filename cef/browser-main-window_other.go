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

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

func (m *lclBrowserWindow) mainFormNotInTaskBar() {

}

func (m *lclBrowserWindow) SetOnWndProc(fn lcl.TWndProcEvent) {
	m.TForm.SetOnWndProc(func(msg *types.TMessage) {
		m.InheritedWndProc(msg)
		fn(msg)
	})
}
