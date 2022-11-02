package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
)

type ITCefWindow interface {
	lcl.IWinControl
	Type() consts.TCefWindowHandleType
	SetChromium(chromium *TCEFChromium, tag int32)
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
