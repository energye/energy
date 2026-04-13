//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/ebitengine/purego/objc"
	. "github.com/energye/energy/v3/pkgs/cocoa/types"
	"unsafe"
)

type NSVisualEffectView struct {
	NSView
}

func AsNSVisualEffectView(ptr unsafe.Pointer) *NSVisualEffectView {
	if ptr == nil {
		return nil
	}
	m := new(NSVisualEffectView)
	m.SetInstance(ptr)
	return m
}

func NewNSVisualEffectView(bounds CGRect) *NSVisualEffectView {
	visualEffectClass := objc.GetClass("NSVisualEffectView")
	frostedView := objc.ID(visualEffectClass).Send(objc.RegisterName("alloc"))
	frostedView = frostedView.Send(objc.RegisterName("initWithFrame:"), bounds)
	return AsNSVisualEffectView(unsafe.Pointer(frostedView))
}

func (m *NSVisualEffectView) SetAutoresizingMask(mask NSAutoresizingMaskOptions) {
	m.Self().Send(objc.RegisterName("setAutoresizingMask:"), mask)
}

func (m *NSVisualEffectView) SetBlendingMode(mode NSVisualEffectBlendingMode) {
	m.Self().Send(objc.RegisterName("setBlendingMode:"), mode)
}

func (m *NSVisualEffectView) SetState(state NSVisualEffectState) {
	m.Self().Send(objc.RegisterName("setState:"), state)
}
