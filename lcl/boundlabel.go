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

// IBoundLabel Parent: ICustomLabel
type IBoundLabel interface {
	ICustomLabel
	FocusControl() IWinControl                     // property
	SetFocusControl(AValue IWinControl)            // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	ShowAccelChar() bool                           // property
	SetShowAccelChar(AValue bool)                  // property
	Layout() TTextLayout                           // property
	SetLayout(AValue TTextLayout)                  // property
	WordWrap() bool                                // property
	SetWordWrap(AValue bool)                       // property
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TBoundLabel Parent: TCustomLabel
type TBoundLabel struct {
	TCustomLabel
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewBoundLabel(TheOwner IComponent) IBoundLabel {
	r1 := LCL().SysCallN(446, GetObjectUintptr(TheOwner))
	return AsBoundLabel(r1)
}

func (m *TBoundLabel) FocusControl() IWinControl {
	r1 := LCL().SysCallN(449, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TBoundLabel) SetFocusControl(AValue IWinControl) {
	LCL().SysCallN(449, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBoundLabel) DragCursor() TCursor {
	r1 := LCL().SysCallN(447, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TBoundLabel) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(447, 1, m.Instance(), uintptr(AValue))
}

func (m *TBoundLabel) DragMode() TDragMode {
	r1 := LCL().SysCallN(448, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TBoundLabel) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(448, 1, m.Instance(), uintptr(AValue))
}

func (m *TBoundLabel) ParentColor() bool {
	r1 := LCL().SysCallN(451, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetParentColor(AValue bool) {
	LCL().SysCallN(451, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) ParentFont() bool {
	r1 := LCL().SysCallN(452, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetParentFont(AValue bool) {
	LCL().SysCallN(452, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) ParentShowHint() bool {
	r1 := LCL().SysCallN(453, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetParentShowHint(AValue bool) {
	LCL().SysCallN(453, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) ShowAccelChar() bool {
	r1 := LCL().SysCallN(467, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetShowAccelChar(AValue bool) {
	LCL().SysCallN(467, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) Layout() TTextLayout {
	r1 := LCL().SysCallN(450, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TBoundLabel) SetLayout(AValue TTextLayout) {
	LCL().SysCallN(450, 1, m.Instance(), uintptr(AValue))
}

func (m *TBoundLabel) WordWrap() bool {
	r1 := LCL().SysCallN(468, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetWordWrap(AValue bool) {
	LCL().SysCallN(468, 1, m.Instance(), PascalBool(AValue))
}

func BoundLabelClass() TClass {
	ret := LCL().SysCallN(445)
	return TClass(ret)
}

func (m *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(454, m.Instance(), m.dblClickPtr)
}

func (m *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(455, m.Instance(), m.dragDropPtr)
}

func (m *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(456, m.Instance(), m.dragOverPtr)
}

func (m *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(457, m.Instance(), m.endDragPtr)
}

func (m *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(458, m.Instance(), m.mouseDownPtr)
}

func (m *TBoundLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(459, m.Instance(), m.mouseEnterPtr)
}

func (m *TBoundLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(460, m.Instance(), m.mouseLeavePtr)
}

func (m *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(461, m.Instance(), m.mouseMovePtr)
}

func (m *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(462, m.Instance(), m.mouseUpPtr)
}

func (m *TBoundLabel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(463, m.Instance(), m.mouseWheelPtr)
}

func (m *TBoundLabel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(464, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TBoundLabel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(465, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TBoundLabel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(466, m.Instance(), m.startDragPtr)
}
