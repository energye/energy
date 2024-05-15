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

// IItemProp Parent: IPersistent
type IItemProp interface {
	IPersistent
	EditMask() string               // property
	SetEditMask(AValue string)      // property
	EditStyle() TEditStyle          // property
	SetEditStyle(AValue TEditStyle) // property
	KeyDesc() string                // property
	SetKeyDesc(AValue string)       // property
	PickList() IStrings             // property
	SetPickList(AValue IStrings)    // property
	MaxLength() int32               // property
	SetMaxLength(AValue int32)      // property
	ReadOnly() bool                 // property
	SetReadOnly(AValue bool)        // property
}

// TItemProp Parent: TPersistent
type TItemProp struct {
	TPersistent
}

func NewItemProp(AOwner IValueListEditor) IItemProp {
	r1 := LCL().SysCallN(3403, GetObjectUintptr(AOwner))
	return AsItemProp(r1)
}

func (m *TItemProp) EditMask() string {
	r1 := LCL().SysCallN(3404, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TItemProp) SetEditMask(AValue string) {
	LCL().SysCallN(3404, 1, m.Instance(), PascalStr(AValue))
}

func (m *TItemProp) EditStyle() TEditStyle {
	r1 := LCL().SysCallN(3405, 0, m.Instance(), 0)
	return TEditStyle(r1)
}

func (m *TItemProp) SetEditStyle(AValue TEditStyle) {
	LCL().SysCallN(3405, 1, m.Instance(), uintptr(AValue))
}

func (m *TItemProp) KeyDesc() string {
	r1 := LCL().SysCallN(3406, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TItemProp) SetKeyDesc(AValue string) {
	LCL().SysCallN(3406, 1, m.Instance(), PascalStr(AValue))
}

func (m *TItemProp) PickList() IStrings {
	r1 := LCL().SysCallN(3408, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TItemProp) SetPickList(AValue IStrings) {
	LCL().SysCallN(3408, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TItemProp) MaxLength() int32 {
	r1 := LCL().SysCallN(3407, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TItemProp) SetMaxLength(AValue int32) {
	LCL().SysCallN(3407, 1, m.Instance(), uintptr(AValue))
}

func (m *TItemProp) ReadOnly() bool {
	r1 := LCL().SysCallN(3409, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TItemProp) SetReadOnly(AValue bool) {
	LCL().SysCallN(3409, 1, m.Instance(), PascalBool(AValue))
}

func ItemPropClass() TClass {
	ret := LCL().SysCallN(3402)
	return TClass(ret)
}
