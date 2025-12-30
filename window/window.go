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
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

type TWindow struct {
	lcl.TEngForm
	oldWndPrc               uintptr
	oldWindowStyle          uintptr
	windowsState            types.TWindowState
	previousWindowPlacement types.TRect
}
