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
	"github.com/energye/golcl/lcl/types/colors"
)

type WindowDemo struct {
	cef.LCLBrowserWindow
}

func (m *WindowDemo) OnFormCreate(sender lcl.IObject) {
	fmt.Println("LCLBrowserWindow OnFormCreate")
	//m.EnabledMaximize(false)
	//m.EnabledMaximize(false)
	//m.SetBorderStyle(types.BsNone)
	m.ScreenCenter()
	bufferPanel := cef.NewBufferPanel(m)
	bufferPanel.SetParent(m)
	bufferPanel.SetColor(colors.ClAqua)
	rng := cef.TCefRange{}
	rng.To = 123
	rng.From = 321
	characterBounds := make([]cef.TCefRect, 50)
	for i := 0; i < len(characterBounds); i++ {
		characterBounds[i].Width = 1000 + int32(i)
		characterBounds[i].Height = 110 + int32(i)
		characterBounds[i].X = 2000 + int32(i)
		characterBounds[i].Y = 2101 + int32(i)
	}
	bufferPanel.ChangeCompositionRange(rng, characterBounds)
	bufferPanel.SetForcedDeviceScaleFactor(123.11)
	fmt.Println(bufferPanel.GetForcedDeviceScaleFactor())
	bufferPanel.SetOnClick(func(sender lcl.IObject) {
		fmt.Println("SetOnClick")
	})
	monitor := m.Monitor()
	fmt.Println(monitor.WorkareaRect())
}

func main() {
	cef.GlobalInit(nil, nil)
	var window = &WindowDemo{}
	lcl.RunApp(&window)
}
