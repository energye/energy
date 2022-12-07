//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
)

type ITCefWindow interface {
	lcl.IWinControl
	Type() consts.TCefWindowHandleType
	SetChromium(chromium IChromium, tag int32)
	UpdateSize()
	HandleAllocated() bool
	CreateHandle()
	SetOnEnter(fn lcl.TNotifyEvent)
	SetOnExit(fn lcl.TNotifyEvent)
	DestroyChildWindow() bool
	Free()
}

func NewCEFWindow(owner lcl.IComponent) ITCefWindow {
	if common.IsWindows() {
		return NewCEFWindowParent(owner)
	} else {
		return NewCEFLinkedWindowParent(owner)
	}
}
