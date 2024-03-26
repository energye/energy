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

// IToolWindow Parent: ICustomControl
type IToolWindow interface {
	ICustomControl
	EdgeBorders() TEdgeBorders          // property
	SetEdgeBorders(AValue TEdgeBorders) // property
	EdgeInner() TEdgeStyle              // property
	SetEdgeInner(AValue TEdgeStyle)     // property
	EdgeOuter() TEdgeStyle              // property
	SetEdgeOuter(AValue TEdgeStyle)     // property
	BeginUpdate()                       // procedure
	EndUpdate()                         // procedure
}

// TToolWindow Parent: TCustomControl
type TToolWindow struct {
	TCustomControl
}

func NewToolWindow(TheOwner IComponent) IToolWindow {
	r1 := LCL().SysCallN(4884, GetObjectUintptr(TheOwner))
	return AsToolWindow(r1)
}

func (m *TToolWindow) EdgeBorders() TEdgeBorders {
	r1 := LCL().SysCallN(4885, 0, m.Instance(), 0)
	return TEdgeBorders(r1)
}

func (m *TToolWindow) SetEdgeBorders(AValue TEdgeBorders) {
	LCL().SysCallN(4885, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolWindow) EdgeInner() TEdgeStyle {
	r1 := LCL().SysCallN(4886, 0, m.Instance(), 0)
	return TEdgeStyle(r1)
}

func (m *TToolWindow) SetEdgeInner(AValue TEdgeStyle) {
	LCL().SysCallN(4886, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolWindow) EdgeOuter() TEdgeStyle {
	r1 := LCL().SysCallN(4887, 0, m.Instance(), 0)
	return TEdgeStyle(r1)
}

func (m *TToolWindow) SetEdgeOuter(AValue TEdgeStyle) {
	LCL().SysCallN(4887, 1, m.Instance(), uintptr(AValue))
}

func ToolWindowClass() TClass {
	ret := LCL().SysCallN(4883)
	return TClass(ret)
}

func (m *TToolWindow) BeginUpdate() {
	LCL().SysCallN(4882, m.Instance())
}

func (m *TToolWindow) EndUpdate() {
	LCL().SysCallN(4888, m.Instance())
}
