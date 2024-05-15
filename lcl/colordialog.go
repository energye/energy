//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IColorDialog Parent: ICommonDialog
type IColorDialog interface {
	ICommonDialog
	Color() TColor                   // property
	SetColor(AValue TColor)          // property
	CustomColors() IStrings          // property
	SetCustomColors(AValue IStrings) // property
}

// TColorDialog Parent: TCommonDialog
type TColorDialog struct {
	TCommonDialog
}

func NewColorDialog(TheOwner IComponent) IColorDialog {
	r1 := LCL().SysCallN(765, GetObjectUintptr(TheOwner))
	return AsColorDialog(r1)
}

func (m *TColorDialog) Color() TColor {
	r1 := LCL().SysCallN(764, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TColorDialog) SetColor(AValue TColor) {
	LCL().SysCallN(764, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorDialog) CustomColors() IStrings {
	r1 := LCL().SysCallN(766, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TColorDialog) SetCustomColors(AValue IStrings) {
	LCL().SysCallN(766, 1, m.Instance(), GetObjectUintptr(AValue))
}

func ColorDialogClass() TClass {
	ret := LCL().SysCallN(763)
	return TClass(ret)
}
