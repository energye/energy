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
	r1 := LCL().SysCallN(3160, GetObjectUintptr(AOwner))
	return AsItemProp(r1)
}

func (m *TItemProp) EditMask() string {
	r1 := LCL().SysCallN(3161, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TItemProp) SetEditMask(AValue string) {
	LCL().SysCallN(3161, 1, m.Instance(), PascalStr(AValue))
}

func (m *TItemProp) EditStyle() TEditStyle {
	r1 := LCL().SysCallN(3162, 0, m.Instance(), 0)
	return TEditStyle(r1)
}

func (m *TItemProp) SetEditStyle(AValue TEditStyle) {
	LCL().SysCallN(3162, 1, m.Instance(), uintptr(AValue))
}

func (m *TItemProp) KeyDesc() string {
	r1 := LCL().SysCallN(3163, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TItemProp) SetKeyDesc(AValue string) {
	LCL().SysCallN(3163, 1, m.Instance(), PascalStr(AValue))
}

func (m *TItemProp) PickList() IStrings {
	r1 := LCL().SysCallN(3165, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TItemProp) SetPickList(AValue IStrings) {
	LCL().SysCallN(3165, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TItemProp) MaxLength() int32 {
	r1 := LCL().SysCallN(3164, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TItemProp) SetMaxLength(AValue int32) {
	LCL().SysCallN(3164, 1, m.Instance(), uintptr(AValue))
}

func (m *TItemProp) ReadOnly() bool {
	r1 := LCL().SysCallN(3166, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TItemProp) SetReadOnly(AValue bool) {
	LCL().SysCallN(3166, 1, m.Instance(), PascalBool(AValue))
}

func ItemPropClass() TClass {
	ret := LCL().SysCallN(3159)
	return TClass(ret)
}
