//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type WindowDemo struct {
	cef.LCLBrowserWindow
}

func (m *WindowDemo) OnFormCreate(sender lcl.IObject) {
	fmt.Println("LCLBrowserWindow OnFormCreate")
	//m.EnabledMaximize(false)
	fmt.Println(m.BorderIcons())
	m.EnabledMaximize(false)
	m.SetBorderStyle(types.BsNone)
	fmt.Println(m.BorderIcons())
	m.ScreenCenter()
}
func main() {
	cef.GlobalInit(nil, nil)
	var window = &WindowDemo{}
	lcl.RunApp(&window)
}
