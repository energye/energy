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

// ICheckGroup Parent: ICustomCheckGroup
type ICheckGroup interface {
	ICustomCheckGroup
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
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

// TCheckGroup Parent: TCustomCheckGroup
type TCheckGroup struct {
	TCustomCheckGroup
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

func NewCheckGroup(TheOwner IComponent) ICheckGroup {
	r1 := LCL().SysCallN(436, GetObjectUintptr(TheOwner))
	return AsCheckGroup(r1)
}

func (m *TCheckGroup) DragCursor() TCursor {
	r1 := LCL().SysCallN(437, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckGroup) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(437, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckGroup) DragMode() TDragMode {
	r1 := LCL().SysCallN(438, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckGroup) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(438, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckGroup) ParentFont() bool {
	r1 := LCL().SysCallN(440, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckGroup) SetParentFont(AValue bool) {
	LCL().SysCallN(440, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckGroup) ParentColor() bool {
	r1 := LCL().SysCallN(439, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckGroup) SetParentColor(AValue bool) {
	LCL().SysCallN(439, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckGroup) ParentShowHint() bool {
	r1 := LCL().SysCallN(441, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckGroup) SetParentShowHint(AValue bool) {
	LCL().SysCallN(441, 1, m.Instance(), PascalBool(AValue))
}

func CheckGroupClass() TClass {
	ret := LCL().SysCallN(435)
	return TClass(ret)
}

func (m *TCheckGroup) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(442, m.Instance(), m.dblClickPtr)
}

func (m *TCheckGroup) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(443, m.Instance(), m.dragDropPtr)
}

func (m *TCheckGroup) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(444, m.Instance(), m.dragOverPtr)
}

func (m *TCheckGroup) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(445, m.Instance(), m.endDragPtr)
}

func (m *TCheckGroup) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(446, m.Instance(), m.mouseDownPtr)
}

func (m *TCheckGroup) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(447, m.Instance(), m.mouseEnterPtr)
}

func (m *TCheckGroup) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(448, m.Instance(), m.mouseLeavePtr)
}

func (m *TCheckGroup) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(449, m.Instance(), m.mouseMovePtr)
}

func (m *TCheckGroup) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(450, m.Instance(), m.mouseUpPtr)
}

func (m *TCheckGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(451, m.Instance(), m.mouseWheelPtr)
}

func (m *TCheckGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(452, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCheckGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(453, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCheckGroup) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(454, m.Instance(), m.startDragPtr)
}
