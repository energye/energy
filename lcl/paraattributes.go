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
	r1 := LCL().SysCallN(4505, GetObjectUintptr(AOwner))
	return AsParaAttributes(r1)
}

func (m *TParaAttributes) Alignment() TAlignment {
	r1 := LCL().SysCallN(4503, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TParaAttributes) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(4503, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) FirstIndent() int32 {
	r1 := LCL().SysCallN(4506, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetFirstIndent(AValue int32) {
	LCL().SysCallN(4506, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) LeftIndent() int32 {
	r1 := LCL().SysCallN(4507, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetLeftIndent(AValue int32) {
	LCL().SysCallN(4507, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) RightIndent() int32 {
	r1 := LCL().SysCallN(4509, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetRightIndent(AValue int32) {
	LCL().SysCallN(4509, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) Numbering() TNumberingStyle {
	r1 := LCL().SysCallN(4508, 0, m.Instance(), 0)
	return TNumberingStyle(r1)
}

func (m *TParaAttributes) SetNumbering(AValue TNumberingStyle) {
	LCL().SysCallN(4508, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) TabCount() int32 {
	r1 := LCL().SysCallN(4511, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetTabCount(AValue int32) {
	LCL().SysCallN(4511, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) Tab(Index Byte) int32 {
	r1 := LCL().SysCallN(4510, 0, m.Instance(), uintptr(Index))
	return int32(r1)
}

func (m *TParaAttributes) SetTab(Index Byte, AValue int32) {
	LCL().SysCallN(4510, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func ParaAttributesClass() TClass {
	ret := LCL().SysCallN(4504)
	return TClass(ret)
}
