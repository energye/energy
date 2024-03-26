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
	"unsafe"
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
	r1 := LCL().SysCallN(738, GetObjectUintptr(OwnerControl), uintptr(unsafe.Pointer(ADefault)))
	return AsControlBorderSpacing(r1)
}

func (m *TControlBorderSpacing) Control() IControl {
	r1 := LCL().SysCallN(731, m.Instance())
	return AsControl(r1)
}

func (m *TControlBorderSpacing) Space(Kind TAnchorKind) int32 {
	r1 := LCL().SysCallN(747, 0, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func (m *TControlBorderSpacing) SetSpace(Kind TAnchorKind, AValue int32) {
	LCL().SysCallN(747, 1, m.Instance(), uintptr(Kind), uintptr(AValue))
}

func (m *TControlBorderSpacing) AroundLeft() int32 {
	r1 := LCL().SysCallN(722, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundTop() int32 {
	r1 := LCL().SysCallN(724, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundRight() int32 {
	r1 := LCL().SysCallN(723, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundBottom() int32 {
	r1 := LCL().SysCallN(721, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlLeft() int32 {
	r1 := LCL().SysCallN(734, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlTop() int32 {
	r1 := LCL().SysCallN(736, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlWidth() int32 {
	r1 := LCL().SysCallN(737, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlHeight() int32 {
	r1 := LCL().SysCallN(733, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlRight() int32 {
	r1 := LCL().SysCallN(735, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlBottom() int32 {
	r1 := LCL().SysCallN(732, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) Left() TSpacingSize {
	r1 := LCL().SysCallN(744, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetLeft(AValue TSpacingSize) {
	LCL().SysCallN(744, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Top() TSpacingSize {
	r1 := LCL().SysCallN(748, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetTop(AValue TSpacingSize) {
	LCL().SysCallN(748, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Right() TSpacingSize {
	r1 := LCL().SysCallN(745, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetRight(AValue TSpacingSize) {
	LCL().SysCallN(745, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Bottom() TSpacingSize {
	r1 := LCL().SysCallN(727, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetBottom(AValue TSpacingSize) {
	LCL().SysCallN(727, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Around() TSpacingSize {
	r1 := LCL().SysCallN(720, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetAround(AValue TSpacingSize) {
	LCL().SysCallN(720, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) InnerBorder() int32 {
	r1 := LCL().SysCallN(742, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlBorderSpacing) SetInnerBorder(AValue int32) {
	LCL().SysCallN(742, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) CellAlignHorizontal() TControlCellAlign {
	r1 := LCL().SysCallN(728, 0, m.Instance(), 0)
	return TControlCellAlign(r1)
}

func (m *TControlBorderSpacing) SetCellAlignHorizontal(AValue TControlCellAlign) {
	LCL().SysCallN(728, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) CellAlignVertical() TControlCellAlign {
	r1 := LCL().SysCallN(729, 0, m.Instance(), 0)
	return TControlCellAlign(r1)
}

func (m *TControlBorderSpacing) SetCellAlignVertical(AValue TControlCellAlign) {
	LCL().SysCallN(729, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) IsEqual(Spacing IControlBorderSpacing) bool {
	r1 := LCL().SysCallN(743, m.Instance(), GetObjectUintptr(Spacing))
	return GoBool(r1)
}

func (m *TControlBorderSpacing) GetSideSpace(Kind TAnchorKind) int32 {
	r1 := LCL().SysCallN(739, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func (m *TControlBorderSpacing) GetSpace1(Kind TAnchorKind) int32 {
	r1 := LCL().SysCallN(740, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func ControlBorderSpacingClass() TClass {
	ret := LCL().SysCallN(730)
	return TClass(ret)
}

func (m *TControlBorderSpacing) AssignTo(Dest IPersistent) {
	LCL().SysCallN(725, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TControlBorderSpacing) GetSpaceAround(SpaceAround *TRect) {
	var result0 uintptr
	LCL().SysCallN(741, m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*SpaceAround = *(*TRect)(getPointer(result0))
}

func (m *TControlBorderSpacing) AutoAdjustLayout(AXProportion, AYProportion float64) {
	LCL().SysCallN(726, m.Instance(), uintptr(unsafe.Pointer(&AXProportion)), uintptr(unsafe.Pointer(&AYProportion)))
}

func (m *TControlBorderSpacing) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(746, m.Instance(), m.changePtr)
}
