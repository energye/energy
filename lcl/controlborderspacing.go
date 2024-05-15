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

// IControlBorderSpacing Parent: IPersistent
type IControlBorderSpacing interface {
	IPersistent
	Control() IControl                                   // property
	Space(Kind TAnchorKind) int32                        // property
	SetSpace(Kind TAnchorKind, AValue int32)             // property
	AroundLeft() int32                                   // property
	AroundTop() int32                                    // property
	AroundRight() int32                                  // property
	AroundBottom() int32                                 // property
	ControlLeft() int32                                  // property
	ControlTop() int32                                   // property
	ControlWidth() int32                                 // property
	ControlHeight() int32                                // property
	ControlRight() int32                                 // property
	ControlBottom() int32                                // property
	Left() TSpacingSize                                  // property
	SetLeft(AValue TSpacingSize)                         // property
	Top() TSpacingSize                                   // property
	SetTop(AValue TSpacingSize)                          // property
	Right() TSpacingSize                                 // property
	SetRight(AValue TSpacingSize)                        // property
	Bottom() TSpacingSize                                // property
	SetBottom(AValue TSpacingSize)                       // property
	Around() TSpacingSize                                // property
	SetAround(AValue TSpacingSize)                       // property
	InnerBorder() int32                                  // property
	SetInnerBorder(AValue int32)                         // property
	CellAlignHorizontal() TControlCellAlign              // property
	SetCellAlignHorizontal(AValue TControlCellAlign)     // property
	CellAlignVertical() TControlCellAlign                // property
	SetCellAlignVertical(AValue TControlCellAlign)       // property
	IsEqual(Spacing IControlBorderSpacing) bool          // function
	GetSideSpace(Kind TAnchorKind) int32                 // function
	GetSpace1(Kind TAnchorKind) int32                    // function
	AssignTo(Dest IPersistent)                           // procedure
	GetSpaceAround(SpaceAround *TRect)                   // procedure
	AutoAdjustLayout(AXProportion, AYProportion float64) // procedure
	SetOnChange(fn TNotifyEvent)                         // property event
}

// TControlBorderSpacing Parent: TPersistent
type TControlBorderSpacing struct {
	TPersistent
	changePtr uintptr
}

func NewControlBorderSpacing(OwnerControl IControl, ADefault *TControlBorderSpacingDefault) IControlBorderSpacing {
	r1 := LCL().SysCallN(928, GetObjectUintptr(OwnerControl), uintptr(unsafePointer(ADefault)))
	return AsControlBorderSpacing(r1)
}

func (m *TControlBorderSpacing) Control() IControl {
	r1 := LCL().SysCallN(921, m.Instance())
	return AsControl(r1)
}

func (m *TControlBorderSpacing) Space(Kind TAnchorKind) int32 {
	r1 := LCL().SysCallN(937, 0, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func (m *TControlBorderSpacing) SetSpace(Kind TAnchorKind, AValue int32) {
	LCL().SysCallN(937, 1, m.Instance(), uintptr(Kind), uintptr(AValue))
}

func (m *TControlBorderSpacing) AroundLeft() int32 {
	r1 := LCL().SysCallN(912, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundTop() int32 {
	r1 := LCL().SysCallN(914, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundRight() int32 {
	r1 := LCL().SysCallN(913, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundBottom() int32 {
	r1 := LCL().SysCallN(911, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlLeft() int32 {
	r1 := LCL().SysCallN(924, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlTop() int32 {
	r1 := LCL().SysCallN(926, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlWidth() int32 {
	r1 := LCL().SysCallN(927, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlHeight() int32 {
	r1 := LCL().SysCallN(923, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlRight() int32 {
	r1 := LCL().SysCallN(925, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlBottom() int32 {
	r1 := LCL().SysCallN(922, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) Left() TSpacingSize {
	r1 := LCL().SysCallN(934, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetLeft(AValue TSpacingSize) {
	LCL().SysCallN(934, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Top() TSpacingSize {
	r1 := LCL().SysCallN(938, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetTop(AValue TSpacingSize) {
	LCL().SysCallN(938, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Right() TSpacingSize {
	r1 := LCL().SysCallN(935, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetRight(AValue TSpacingSize) {
	LCL().SysCallN(935, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Bottom() TSpacingSize {
	r1 := LCL().SysCallN(917, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetBottom(AValue TSpacingSize) {
	LCL().SysCallN(917, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Around() TSpacingSize {
	r1 := LCL().SysCallN(910, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetAround(AValue TSpacingSize) {
	LCL().SysCallN(910, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) InnerBorder() int32 {
	r1 := LCL().SysCallN(932, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlBorderSpacing) SetInnerBorder(AValue int32) {
	LCL().SysCallN(932, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) CellAlignHorizontal() TControlCellAlign {
	r1 := LCL().SysCallN(918, 0, m.Instance(), 0)
	return TControlCellAlign(r1)
}

func (m *TControlBorderSpacing) SetCellAlignHorizontal(AValue TControlCellAlign) {
	LCL().SysCallN(918, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) CellAlignVertical() TControlCellAlign {
	r1 := LCL().SysCallN(919, 0, m.Instance(), 0)
	return TControlCellAlign(r1)
}

func (m *TControlBorderSpacing) SetCellAlignVertical(AValue TControlCellAlign) {
	LCL().SysCallN(919, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) IsEqual(Spacing IControlBorderSpacing) bool {
	r1 := LCL().SysCallN(933, m.Instance(), GetObjectUintptr(Spacing))
	return GoBool(r1)
}

func (m *TControlBorderSpacing) GetSideSpace(Kind TAnchorKind) int32 {
	r1 := LCL().SysCallN(929, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func (m *TControlBorderSpacing) GetSpace1(Kind TAnchorKind) int32 {
	r1 := LCL().SysCallN(930, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func ControlBorderSpacingClass() TClass {
	ret := LCL().SysCallN(920)
	return TClass(ret)
}

func (m *TControlBorderSpacing) AssignTo(Dest IPersistent) {
	LCL().SysCallN(915, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TControlBorderSpacing) GetSpaceAround(SpaceAround *TRect) {
	var result0 uintptr
	LCL().SysCallN(931, m.Instance(), uintptr(unsafePointer(&result0)))
	*SpaceAround = *(*TRect)(getPointer(result0))
}

func (m *TControlBorderSpacing) AutoAdjustLayout(AXProportion, AYProportion float64) {
	LCL().SysCallN(916, m.Instance(), uintptr(unsafePointer(&AXProportion)), uintptr(unsafePointer(&AYProportion)))
}

func (m *TControlBorderSpacing) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(936, m.Instance(), m.changePtr)
}
