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
	r1 := LCL().SysCallN(961, GetObjectUintptr(AControl), uintptr(AKind))
	return AsControlScrollBar(r1)
}

func (m *TControlScrollBar) Kind() TScrollBarKind {
	r1 := LCL().SysCallN(965, m.Instance())
	return TScrollBarKind(r1)
}

func (m *TControlScrollBar) Size() int32 {
	r1 := LCL().SysCallN(970, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) Increment() TScrollBarInc {
	r1 := LCL().SysCallN(963, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TControlScrollBar) SetIncrement(AValue TScrollBarInc) {
	LCL().SysCallN(963, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Page() TScrollBarInc {
	r1 := LCL().SysCallN(966, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TControlScrollBar) SetPage(AValue TScrollBarInc) {
	LCL().SysCallN(966, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Smooth() bool {
	r1 := LCL().SysCallN(971, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetSmooth(AValue bool) {
	LCL().SysCallN(971, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) Position() int32 {
	r1 := LCL().SysCallN(967, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlScrollBar) SetPosition(AValue int32) {
	LCL().SysCallN(967, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Range() int32 {
	r1 := LCL().SysCallN(968, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlScrollBar) SetRange(AValue int32) {
	LCL().SysCallN(968, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Tracking() bool {
	r1 := LCL().SysCallN(972, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetTracking(AValue bool) {
	LCL().SysCallN(972, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) Visible() bool {
	r1 := LCL().SysCallN(973, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetVisible(AValue bool) {
	LCL().SysCallN(973, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) IsScrollBarVisible() bool {
	r1 := LCL().SysCallN(964, m.Instance())
	return GoBool(r1)
}

func (m *TControlScrollBar) ScrollPos() int32 {
	r1 := LCL().SysCallN(969, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) GetOtherScrollBar() IControlScrollBar {
	r1 := LCL().SysCallN(962, m.Instance())
	return AsControlScrollBar(r1)
}

func (m *TControlScrollBar) ControlSize() int32 {
	r1 := LCL().SysCallN(960, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSize() int32 {
	r1 := LCL().SysCallN(957, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSizeWithBar() int32 {
	r1 := LCL().SysCallN(958, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSizeWithoutBar() int32 {
	r1 := LCL().SysCallN(959, m.Instance())
	return int32(r1)
}

func ControlScrollBarClass() TClass {
	ret := LCL().SysCallN(956)
	return TClass(ret)
}
