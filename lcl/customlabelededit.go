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

// ICustomLabeledEdit Parent: ICustomEdit
type ICustomLabeledEdit interface {
	ICustomEdit
	EditLabel() IBoundLabel                 // property
	LabelPosition() TLabelPosition          // property
	SetLabelPosition(AValue TLabelPosition) // property
	LabelSpacing() int32                    // property
	SetLabelSpacing(AValue int32)           // property
}

// TCustomLabeledEdit Parent: TCustomEdit
type TCustomLabeledEdit struct {
	TCustomEdit
}

func NewCustomLabeledEdit(TheOwner IComponent) ICustomLabeledEdit {
	r1 := LCL().SysCallN(1985, GetObjectUintptr(TheOwner))
	return AsCustomLabeledEdit(r1)
}

func (m *TCustomLabeledEdit) EditLabel() IBoundLabel {
	r1 := LCL().SysCallN(1986, m.Instance())
	return AsBoundLabel(r1)
}

func (m *TCustomLabeledEdit) LabelPosition() TLabelPosition {
	r1 := LCL().SysCallN(1987, 0, m.Instance(), 0)
	return TLabelPosition(r1)
}

func (m *TCustomLabeledEdit) SetLabelPosition(AValue TLabelPosition) {
	LCL().SysCallN(1987, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomLabeledEdit) LabelSpacing() int32 {
	r1 := LCL().SysCallN(1988, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomLabeledEdit) SetLabelSpacing(AValue int32) {
	LCL().SysCallN(1988, 1, m.Instance(), uintptr(AValue))
}

func CustomLabeledEditClass() TClass {
	ret := LCL().SysCallN(1984)
	return TClass(ret)
}
