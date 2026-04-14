//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cgo

import (
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

type NSVisualEffectView struct {
	NSView
}

func AsNSVisualEffectView(ptr unsafe.Pointer) INSVisualEffectView {
	if ptr == nil {
		return nil
	}
	m := new(NSVisualEffectView)
	m.SetInstance(ptr)
	return m
}

func (m *NSVisualEffectView) SetAutoresizingMask(mask NSAutoresizingMaskOptions) {
	// no impl
}

func (m *NSVisualEffectView) SetBlendingMode(mode NSVisualEffectBlendingMode) {
	// no impl
}

func (m *NSVisualEffectView) SetState(state NSVisualEffectState) {
	// no impl
}
