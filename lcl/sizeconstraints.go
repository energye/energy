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

// ISizeConstraints Parent: IPersistent
type ISizeConstraints interface {
	IPersistent
	MaxInterfaceHeight() int32                            // property
	MaxInterfaceWidth() int32                             // property
	MinInterfaceHeight() int32                            // property
	MinInterfaceWidth() int32                             // property
	Control() IControl                                    // property
	Options() TSizeConstraintsOptions                     // property
	SetOptions(AValue TSizeConstraintsOptions)            // property
	MaxHeight() TConstraintSize                           // property
	SetMaxHeight(AValue TConstraintSize)                  // property
	MaxWidth() TConstraintSize                            // property
	SetMaxWidth(AValue TConstraintSize)                   // property
	MinHeight() TConstraintSize                           // property
	SetMinHeight(AValue TConstraintSize)                  // property
	MinWidth() TConstraintSize                            // property
	SetMinWidth(AValue TConstraintSize)                   // property
	EffectiveMinWidth() int32                             // function
	EffectiveMinHeight() int32                            // function
	EffectiveMaxWidth() int32                             // function
	EffectiveMaxHeight() int32                            // function
	MinMaxWidth(Width int32) int32                        // function
	MinMaxHeight(Height int32) int32                      // function
	UpdateInterfaceConstraints()                          // procedure
	SetInterfaceConstraints(MinW, MinH, MaxW, MaxH int32) // procedure
	AutoAdjustLayout(AXProportion, AYProportion float64)  // procedure
	SetOnChange(fn TNotifyEvent)                          // property event
}

// TSizeConstraints Parent: TPersistent
type TSizeConstraints struct {
	TPersistent
	changePtr uintptr
}

func NewSizeConstraints(AControl IControl) ISizeConstraints {
	r1 := LCL().SysCallN(4358, GetObjectUintptr(AControl))
	return AsSizeConstraints(r1)
}

func (m *TSizeConstraints) MaxInterfaceHeight() int32 {
	r1 := LCL().SysCallN(4364, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MaxInterfaceWidth() int32 {
	r1 := LCL().SysCallN(4365, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinInterfaceHeight() int32 {
	r1 := LCL().SysCallN(4368, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinInterfaceWidth() int32 {
	r1 := LCL().SysCallN(4369, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) Control() IControl {
	r1 := LCL().SysCallN(4357, m.Instance())
	return AsControl(r1)
}

func (m *TSizeConstraints) Options() TSizeConstraintsOptions {
	r1 := LCL().SysCallN(4373, 0, m.Instance(), 0)
	return TSizeConstraintsOptions(r1)
}

func (m *TSizeConstraints) SetOptions(AValue TSizeConstraintsOptions) {
	LCL().SysCallN(4373, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MaxHeight() TConstraintSize {
	r1 := LCL().SysCallN(4363, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMaxHeight(AValue TConstraintSize) {
	LCL().SysCallN(4363, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MaxWidth() TConstraintSize {
	r1 := LCL().SysCallN(4366, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMaxWidth(AValue TConstraintSize) {
	LCL().SysCallN(4366, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MinHeight() TConstraintSize {
	r1 := LCL().SysCallN(4367, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMinHeight(AValue TConstraintSize) {
	LCL().SysCallN(4367, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MinWidth() TConstraintSize {
	r1 := LCL().SysCallN(4372, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMinWidth(AValue TConstraintSize) {
	LCL().SysCallN(4372, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) EffectiveMinWidth() int32 {
	r1 := LCL().SysCallN(4362, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMinHeight() int32 {
	r1 := LCL().SysCallN(4361, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMaxWidth() int32 {
	r1 := LCL().SysCallN(4360, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMaxHeight() int32 {
	r1 := LCL().SysCallN(4359, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinMaxWidth(Width int32) int32 {
	r1 := LCL().SysCallN(4371, m.Instance(), uintptr(Width))
	return int32(r1)
}

func (m *TSizeConstraints) MinMaxHeight(Height int32) int32 {
	r1 := LCL().SysCallN(4370, m.Instance(), uintptr(Height))
	return int32(r1)
}

func SizeConstraintsClass() TClass {
	ret := LCL().SysCallN(4356)
	return TClass(ret)
}

func (m *TSizeConstraints) UpdateInterfaceConstraints() {
	LCL().SysCallN(4376, m.Instance())
}

func (m *TSizeConstraints) SetInterfaceConstraints(MinW, MinH, MaxW, MaxH int32) {
	LCL().SysCallN(4374, m.Instance(), uintptr(MinW), uintptr(MinH), uintptr(MaxW), uintptr(MaxH))
}

func (m *TSizeConstraints) AutoAdjustLayout(AXProportion, AYProportion float64) {
	LCL().SysCallN(4355, m.Instance(), uintptr(unsafe.Pointer(&AXProportion)), uintptr(unsafe.Pointer(&AYProportion)))
}

func (m *TSizeConstraints) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4375, m.Instance(), m.changePtr)
}
