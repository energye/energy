package main

import (
	"fmt"
	"github.com/energye/energy/cef"
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
