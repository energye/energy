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

// IPen Parent: IFPCustomPen
type IPen interface {
	IFPCustomPen
	Color() TColor                                    // property
	SetColor(AValue TColor)                           // property
	Cosmetic() bool                                   // property
	SetCosmetic(AValue bool)                          // property
	EndCapForPenEndCap() TPenEndCap                   // property
	SetEndCapForPenEndCap(AValue TPenEndCap)          // property
	JoinStyleForPenJoinStyle() TPenJoinStyle          // property
	SetJoinStyleForPenJoinStyle(AValue TPenJoinStyle) // property
	GetPatternForUintptr() uintptr                    // function
	SetPatternForPointer(APattern uintptr)            // procedure
}

// TPen Parent: TFPCustomPen
type TPen struct {
	TFPCustomPen
}

func NewPen() IPen {
	r1 := LCL().SysCallN(4515)
	return AsPen(r1)
}

func (m *TPen) Color() TColor {
	r1 := LCL().SysCallN(4513, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TPen) SetColor(AValue TColor) {
	LCL().SysCallN(4513, 1, m.Instance(), uintptr(AValue))
}

func (m *TPen) Cosmetic() bool {
	r1 := LCL().SysCallN(4514, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPen) SetCosmetic(AValue bool) {
	LCL().SysCallN(4514, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPen) EndCapForPenEndCap() TPenEndCap {
	r1 := LCL().SysCallN(4516, 0, m.Instance(), 0)
	return TPenEndCap(r1)
}

func (m *TPen) SetEndCapForPenEndCap(AValue TPenEndCap) {
	LCL().SysCallN(4516, 1, m.Instance(), uintptr(AValue))
}

func (m *TPen) JoinStyleForPenJoinStyle() TPenJoinStyle {
	r1 := LCL().SysCallN(4518, 0, m.Instance(), 0)
	return TPenJoinStyle(r1)
}

func (m *TPen) SetJoinStyleForPenJoinStyle(AValue TPenJoinStyle) {
	LCL().SysCallN(4518, 1, m.Instance(), uintptr(AValue))
}

func (m *TPen) GetPatternForUintptr() uintptr {
	r1 := LCL().SysCallN(4517, m.Instance())
	return uintptr(r1)
}

func PenClass() TClass {
	ret := LCL().SysCallN(4512)
	return TClass(ret)
}

func (m *TPen) SetPatternForPointer(APattern uintptr) {
	LCL().SysCallN(4519, m.Instance(), uintptr(APattern))
}
