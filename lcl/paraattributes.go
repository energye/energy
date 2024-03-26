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

// IParaAttributes Parent: IPersistent
type IParaAttributes interface {
	IPersistent
	Alignment() TAlignment               // property
	SetAlignment(AValue TAlignment)      // property
	FirstIndent() int32                  // property
	SetFirstIndent(AValue int32)         // property
	LeftIndent() int32                   // property
	SetLeftIndent(AValue int32)          // property
	RightIndent() int32                  // property
	SetRightIndent(AValue int32)         // property
	Numbering() TNumberingStyle          // property
	SetNumbering(AValue TNumberingStyle) // property
	TabCount() int32                     // property
	SetTabCount(AValue int32)            // property
	Tab(Index Byte) int32                // property
	SetTab(Index Byte, AValue int32)     // property
}

// TParaAttributes Parent: TPersistent
type TParaAttributes struct {
	TPersistent
}

func NewParaAttributes(AOwner IRichMemo) IParaAttributes {
	r1 := LCL().SysCallN(3863, GetObjectUintptr(AOwner))
	return AsParaAttributes(r1)
}

func (m *TParaAttributes) Alignment() TAlignment {
	r1 := LCL().SysCallN(3861, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TParaAttributes) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3861, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) FirstIndent() int32 {
	r1 := LCL().SysCallN(3864, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetFirstIndent(AValue int32) {
	LCL().SysCallN(3864, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) LeftIndent() int32 {
	r1 := LCL().SysCallN(3865, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetLeftIndent(AValue int32) {
	LCL().SysCallN(3865, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) RightIndent() int32 {
	r1 := LCL().SysCallN(3867, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetRightIndent(AValue int32) {
	LCL().SysCallN(3867, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) Numbering() TNumberingStyle {
	r1 := LCL().SysCallN(3866, 0, m.Instance(), 0)
	return TNumberingStyle(r1)
}

func (m *TParaAttributes) SetNumbering(AValue TNumberingStyle) {
	LCL().SysCallN(3866, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) TabCount() int32 {
	r1 := LCL().SysCallN(3869, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetTabCount(AValue int32) {
	LCL().SysCallN(3869, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) Tab(Index Byte) int32 {
	r1 := LCL().SysCallN(3868, 0, m.Instance(), uintptr(Index))
	return int32(r1)
}

func (m *TParaAttributes) SetTab(Index Byte, AValue int32) {
	LCL().SysCallN(3868, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func ParaAttributesClass() TClass {
	ret := LCL().SysCallN(3862)
	return TClass(ret)
}
