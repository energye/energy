//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

import "github.com/energye/lcl/types"

type INSObject interface {
	Instance() uintptr
}

type INSResponder interface {
	INSObject
}

type INSToolBar interface {
	INSObject
}

type INSView interface {
	INSResponder
}

type INSVisualEffectView interface {
	INSView
	SetAutoresizingMask(mask NSAutoresizingMaskOptions)
	SetBlendingMode(mode NSVisualEffectBlendingMode)
	SetState(state NSVisualEffectState)
}

type INSWindow interface {
	INSResponder
	Restore()
	Maximize()
	Minimized()
	ExitMinimized()
	EnterFullScreen()
	ExitFullScreen()
	Drag()
	SetBackgroundColor(r, g, b, alpha uint8)
	SetRadius(radius float32)
	SetTransparent() INSVisualEffectView
	SwitchFrostedMaterial(frostedView INSVisualEffectView, appearanceName string)
	AddSubview(view INSView, x, y, width, height float32)
	ContentViewFrame() (rect types.TRect)
}

type INSWindowDelegate interface {
	INSObject
}
