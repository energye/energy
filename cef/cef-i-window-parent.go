package cef

import (
	"github.com/energye/energy/commons"
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
}

func NewCEFWindow(owner lcl.IComponent) ITCefWindow {
	if commons.IsWindows() {
		return NewCEFWindowParent(owner)
	} else {
		return NewCEFLinkedWindowParent(owner)
	}
}
