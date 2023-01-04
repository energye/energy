//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//创建一个新window窗口
func NewWindow(windowProperty *WindowProperty) *LCLBrowserWindow {
	if windowProperty == nil {
		windowProperty = NewWindowProperty()
	}
	var window = &LCLBrowserWindow{}
	window.windowProperty = windowProperty
	//window.TForm = lcl.NewForm(owner)
	lcl.Application.CreateForm(&window)
	window.ParentDoubleBuffered()
	window.FormCreate()
	window.SetNotInTaskBar()
	window.defaultWindowEvent()
	window.SetCaption(windowProperty.Title)
	if windowProperty.IsCenterWindow {
		window.SetPosition(types.PoDesktopCenter)
	} else {
		window.SetPosition(types.PoDefault)
		window.SetBounds(windowProperty.X, windowProperty.Y, windowProperty.Width, windowProperty.Height)
	}
	if windowProperty.IconFS != "" {
		_ = window.Icon().LoadFromFSFile(windowProperty.IconFS)
	} else if windowProperty.Icon != "" {
		window.Icon().LoadFromFile(windowProperty.Icon)
	}
	if windowProperty.AlwaysOnTop {
		window.SetFormStyle(types.FsSystemStayOnTop)
	}
	window.EnabledMinimize(windowProperty.CanMinimize)
	window.EnabledMaximize(windowProperty.CanMaximize)
	if !windowProperty.CanResize {
		window.SetBorderStyle(types.BsSingle)
	}
	return window
}
