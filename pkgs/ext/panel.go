//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ext

import (
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type TPanel struct {
	*lcl.TPanel
}

func NewPanel(owner lcl.IComponent) *TPanel {
	m := new(TPanel)
	m.TPanel = lcl.NewPanel(owner)
	return m
}

func (m *TPanel) PanelBevelColor() types.TColor {
	r1, _, _ := imports.LibLCLExt().Proc(Ext_Panel_GetBevelColor).Call(m.Instance())
	return types.TColor(r1)
}

func (m *TPanel) SetPanelBevelColor(colors types.TColor) {
	imports.LibLCLExt().Proc(Ext_Panel_SetBevelColor).Call(m.Instance(), uintptr(colors))
}
