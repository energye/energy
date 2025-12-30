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
	WindowResize(ht string)
	WindowDrag(message ipc.ProcessMessage)
}

type TWindow struct {
	lcl.TEngForm
	oldWndPrc               uintptr
	oldWindowStyle          uintptr
	windowsState            types.TWindowState
	previousWindowPlacement types.TRect
}

func (m *TWindow) FormCreate(sender lcl.IObject) {

}

func (m *TWindow) OnCloseQuery(sender lcl.IObject, canClose *bool) {

}

func (m *TWindow) OnClose(sender lcl.IObject, closeAction *types.TCloseAction) {
}
