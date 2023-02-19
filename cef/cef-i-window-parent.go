//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type ITCefWindowParent interface {
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
	Handle() types.HWND
	Name() string
	SetName(value string)
	SetParent(value lcl.IWinControl)
	Align() types.TAlign
	SetAlign(value types.TAlign)
	Anchors() types.TAnchors
	SetAnchors(value types.TAnchors)
	Visible() bool
	SetVisible(value bool)
	Enabled() bool
	SetEnabled(value bool)
	Left() int32
	SetLeft(value int32)
	Top() int32
	SetTop(value int32)
	Width() int32
	SetWidth(value int32)
	Height() int32
	SetHeight(value int32)
	BoundsRect() (result types.TRect)
	SetBoundsRect(value types.TRect)
}

func NewCEFWindow(owner lcl.IComponent) ITCefWindowParent {
	if common.IsWindows() {
		return NewCEFWindowParent(owner)
	} else {
		return NewCEFLinkedWindowParent(owner)
	}
}
