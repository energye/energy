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

// IControlChildSizing Parent: IPersistent
type IControlChildSizing interface {
	IPersistent
	Control() IWinControl                                 // property
	LeftRightSpacing() int32                              // property
	SetLeftRightSpacing(AValue int32)                     // property
	TopBottomSpacing() int32                              // property
	SetTopBottomSpacing(AValue int32)                     // property
	HorizontalSpacing() int32                             // property
	SetHorizontalSpacing(AValue int32)                    // property
	VerticalSpacing() int32                               // property
	SetVerticalSpacing(AValue int32)                      // property
	EnlargeHorizontal() TChildControlResizeStyle          // property
	SetEnlargeHorizontal(AValue TChildControlResizeStyle) // property
	EnlargeVertical() TChildControlResizeStyle            // property
	SetEnlargeVertical(AValue TChildControlResizeStyle)   // property
	ShrinkHorizontal() TChildControlResizeStyle           // property
	SetShrinkHorizontal(AValue TChildControlResizeStyle)  // property
	ShrinkVertical() TChildControlResizeStyle             // property
	SetShrinkVertical(AValue TChildControlResizeStyle)    // property
	Layout() TControlChildrenLayout                       // property
	SetLayout(AValue TControlChildrenLayout)              // property
	ControlsPerLine() int32                               // property
	SetControlsPerLine(AValue int32)                      // property
	IsEqual(Sizing IControlChildSizing) bool              // function
	AssignTo(Dest IPersistent)                            // procedure
	SetGridSpacing(Spacing int32)                         // procedure
	SetOnChange(fn TNotifyEvent)                          // property event
}

// TControlChildSizing Parent: TPersistent
type TControlChildSizing struct {
	TPersistent
	changePtr uintptr
}

func NewControlChildSizing(OwnerControl IWinControl) IControlChildSizing {
	r1 := LCL().SysCallN(753, GetObjectUintptr(OwnerControl))
	return AsControlChildSizing(r1)
}

func (m *TControlChildSizing) Control() IWinControl {
	r1 := LCL().SysCallN(751, m.Instance())
	return AsWinControl(r1)
}

func (m *TControlChildSizing) LeftRightSpacing() int32 {
	r1 := LCL().SysCallN(759, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetLeftRightSpacing(AValue int32) {
	LCL().SysCallN(759, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) TopBottomSpacing() int32 {
	r1 := LCL().SysCallN(764, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetTopBottomSpacing(AValue int32) {
	LCL().SysCallN(764, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) HorizontalSpacing() int32 {
	r1 := LCL().SysCallN(756, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetHorizontalSpacing(AValue int32) {
	LCL().SysCallN(756, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) VerticalSpacing() int32 {
	r1 := LCL().SysCallN(765, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetVerticalSpacing(AValue int32) {
	LCL().SysCallN(765, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) EnlargeHorizontal() TChildControlResizeStyle {
	r1 := LCL().SysCallN(754, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetEnlargeHorizontal(AValue TChildControlResizeStyle) {
	LCL().SysCallN(754, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) EnlargeVertical() TChildControlResizeStyle {
	r1 := LCL().SysCallN(755, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetEnlargeVertical(AValue TChildControlResizeStyle) {
	LCL().SysCallN(755, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) ShrinkHorizontal() TChildControlResizeStyle {
	r1 := LCL().SysCallN(762, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetShrinkHorizontal(AValue TChildControlResizeStyle) {
	LCL().SysCallN(762, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) ShrinkVertical() TChildControlResizeStyle {
	r1 := LCL().SysCallN(763, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetShrinkVertical(AValue TChildControlResizeStyle) {
	LCL().SysCallN(763, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) Layout() TControlChildrenLayout {
	r1 := LCL().SysCallN(758, 0, m.Instance(), 0)
	return TControlChildrenLayout(r1)
}

func (m *TControlChildSizing) SetLayout(AValue TControlChildrenLayout) {
	LCL().SysCallN(758, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) ControlsPerLine() int32 {
	r1 := LCL().SysCallN(752, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetControlsPerLine(AValue int32) {
	LCL().SysCallN(752, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) IsEqual(Sizing IControlChildSizing) bool {
	r1 := LCL().SysCallN(757, m.Instance(), GetObjectUintptr(Sizing))
	return GoBool(r1)
}

func ControlChildSizingClass() TClass {
	ret := LCL().SysCallN(750)
	return TClass(ret)
}

func (m *TControlChildSizing) AssignTo(Dest IPersistent) {
	LCL().SysCallN(749, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TControlChildSizing) SetGridSpacing(Spacing int32) {
	LCL().SysCallN(760, m.Instance(), uintptr(Spacing))
}

func (m *TControlChildSizing) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(761, m.Instance(), m.changePtr)
}
