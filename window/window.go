// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package window

import (
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

type IWindow interface {
	lcl.IEngForm
	Resize(ht string)
	Drag(message ipc.ProcessMessage)
	SetOptions(windowId uint32)
}

type TWindow struct {
	lcl.TEngForm
	windowId                uint32
	oldWndPrc               uintptr
	oldWindowStyle          uintptr
	windowsState            types.TWindowState
	previousWindowPlacement types.TRect
}
