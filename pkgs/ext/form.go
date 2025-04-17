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
)

type TForm struct {
	*lcl.TForm
}

func NewForm(owner lcl.IComponent) *TForm {
	m := new(TForm)
	m.TForm = lcl.NewForm(owner)
	return m
}

func (m *TForm) FormActivate() {
	imports.LibLCLExt().Proc(Ext_Form_Activate).Call(m.Instance())
}

func (m *TForm) FormDeactivate() {
	imports.LibLCLExt().Proc(Ext_Form_Deactivate).Call(m.Instance())
}
