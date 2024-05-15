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

// IVTFixedAreaConstraints Parent: IPersistent
type IVTFixedAreaConstraints interface {
	IPersistent
	MaxHeightPercent() TVTConstraintPercent          // property
	SetMaxHeightPercent(AValue TVTConstraintPercent) // property
	MaxWidthPercent() TVTConstraintPercent           // property
	SetMaxWidthPercent(AValue TVTConstraintPercent)  // property
	MinHeightPercent() TVTConstraintPercent          // property
	SetMinHeightPercent(AValue TVTConstraintPercent) // property
	MinWidthPercent() TVTConstraintPercent           // property
	SetMinWidthPercent(AValue TVTConstraintPercent)  // property
	SetOnChange(fn TNotifyEvent)                     // property event
}

// TVTFixedAreaConstraints Parent: TPersistent
type TVTFixedAreaConstraints struct {
	TPersistent
	changePtr uintptr
}

func NewVTFixedAreaConstraints(AOwner IVTHeader) IVTFixedAreaConstraints {
	r1 := LCL().SysCallN(5850, GetObjectUintptr(AOwner))
	return AsVTFixedAreaConstraints(r1)
}

func (m *TVTFixedAreaConstraints) MaxHeightPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5851, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMaxHeightPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5851, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MaxWidthPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5852, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMaxWidthPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5852, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MinHeightPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5853, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMinHeightPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5853, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MinWidthPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5854, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMinWidthPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5854, 1, m.Instance(), uintptr(AValue))
}

func VTFixedAreaConstraintsClass() TClass {
	ret := LCL().SysCallN(5849)
	return TClass(ret)
}

func (m *TVTFixedAreaConstraints) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5855, m.Instance(), m.changePtr)
}
