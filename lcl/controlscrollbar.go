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

// IControlScrollBar Parent: IPersistent
type IControlScrollBar interface {
	IPersistent
	Kind() TScrollBarKind                 // property
	Size() int32                          // property
	Increment() TScrollBarInc             // property
	SetIncrement(AValue TScrollBarInc)    // property
	Page() TScrollBarInc                  // property
	SetPage(AValue TScrollBarInc)         // property
	Smooth() bool                         // property
	SetSmooth(AValue bool)                // property
	Position() int32                      // property
	SetPosition(AValue int32)             // property
	Range() int32                         // property
	SetRange(AValue int32)                // property
	Tracking() bool                       // property
	SetTracking(AValue bool)              // property
	Visible() bool                        // property
	SetVisible(AValue bool)               // property
	IsScrollBarVisible() bool             // function
	ScrollPos() int32                     // function
	GetOtherScrollBar() IControlScrollBar // function
	ControlSize() int32                   // function
	ClientSize() int32                    // function
	ClientSizeWithBar() int32             // function
	ClientSizeWithoutBar() int32          // function
}

// TControlScrollBar Parent: TPersistent
type TControlScrollBar struct {
	TPersistent
}

func NewControlScrollBar(AControl IWinControl, AKind TScrollBarKind) IControlScrollBar {
	r1 := LCL().SysCallN(771, GetObjectUintptr(AControl), uintptr(AKind))
	return AsControlScrollBar(r1)
}

func (m *TControlScrollBar) Kind() TScrollBarKind {
	r1 := LCL().SysCallN(775, m.Instance())
	return TScrollBarKind(r1)
}

func (m *TControlScrollBar) Size() int32 {
	r1 := LCL().SysCallN(780, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) Increment() TScrollBarInc {
	r1 := LCL().SysCallN(773, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TControlScrollBar) SetIncrement(AValue TScrollBarInc) {
	LCL().SysCallN(773, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Page() TScrollBarInc {
	r1 := LCL().SysCallN(776, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TControlScrollBar) SetPage(AValue TScrollBarInc) {
	LCL().SysCallN(776, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Smooth() bool {
	r1 := LCL().SysCallN(781, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetSmooth(AValue bool) {
	LCL().SysCallN(781, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) Position() int32 {
	r1 := LCL().SysCallN(777, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlScrollBar) SetPosition(AValue int32) {
	LCL().SysCallN(777, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Range() int32 {
	r1 := LCL().SysCallN(778, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlScrollBar) SetRange(AValue int32) {
	LCL().SysCallN(778, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Tracking() bool {
	r1 := LCL().SysCallN(782, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetTracking(AValue bool) {
	LCL().SysCallN(782, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) Visible() bool {
	r1 := LCL().SysCallN(783, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetVisible(AValue bool) {
	LCL().SysCallN(783, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) IsScrollBarVisible() bool {
	r1 := LCL().SysCallN(774, m.Instance())
	return GoBool(r1)
}

func (m *TControlScrollBar) ScrollPos() int32 {
	r1 := LCL().SysCallN(779, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) GetOtherScrollBar() IControlScrollBar {
	r1 := LCL().SysCallN(772, m.Instance())
	return AsControlScrollBar(r1)
}

func (m *TControlScrollBar) ControlSize() int32 {
	r1 := LCL().SysCallN(770, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSize() int32 {
	r1 := LCL().SysCallN(767, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSizeWithBar() int32 {
	r1 := LCL().SysCallN(768, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSizeWithoutBar() int32 {
	r1 := LCL().SysCallN(769, m.Instance())
	return int32(r1)
}

func ControlScrollBarClass() TClass {
	ret := LCL().SysCallN(766)
	return TClass(ret)
}
