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

// IShape Parent: IGraphicControl
type IShape interface {
	IGraphicControl
	Brush() IBrush                                  // property
	SetBrush(AValue IBrush)                         // property
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	Pen() IPen                                      // property
	SetPen(AValue IPen)                             // property
	Shape() TShapeType                              // property
	SetShape(AValue TShapeType)                     // property
	Paint()                                         // procedure
	StyleChanged(Sender IObject)                    // procedure
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDock(fn TStartDockEvent)              // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TShape Parent: TGraphicControl
type TShape struct {
	TGraphicControl
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDockPtr         uintptr
	endDragPtr         uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	paintPtr           uintptr
	startDockPtr       uintptr
	startDragPtr       uintptr
}

func NewShape(TheOwner IComponent) IShape {
	r1 := LCL().SysCallN(4973, GetObjectUintptr(TheOwner))
	return AsShape(r1)
}

func (m *TShape) Brush() IBrush {
	r1 := LCL().SysCallN(4971, 0, m.Instance(), 0)
	return AsBrush(r1)
}

func (m *TShape) SetBrush(AValue IBrush) {
	LCL().SysCallN(4971, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TShape) DragCursor() TCursor {
	r1 := LCL().SysCallN(4974, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TShape) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4974, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) DragKind() TDragKind {
	r1 := LCL().SysCallN(4975, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TShape) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4975, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) DragMode() TDragMode {
	r1 := LCL().SysCallN(4976, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TShape) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4976, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) ParentShowHint() bool {
	r1 := LCL().SysCallN(4978, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TShape) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4978, 1, m.Instance(), PascalBool(AValue))
}

func (m *TShape) Pen() IPen {
	r1 := LCL().SysCallN(4979, 0, m.Instance(), 0)
	return AsPen(r1)
}

func (m *TShape) SetPen(AValue IPen) {
	LCL().SysCallN(4979, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TShape) Shape() TShapeType {
	r1 := LCL().SysCallN(4998, 0, m.Instance(), 0)
	return TShapeType(r1)
}

func (m *TShape) SetShape(AValue TShapeType) {
	LCL().SysCallN(4998, 1, m.Instance(), uintptr(AValue))
}

func ShapeClass() TClass {
	ret := LCL().SysCallN(4972)
	return TClass(ret)
}

func (m *TShape) Paint() {
	LCL().SysCallN(4977, m.Instance())
}

func (m *TShape) StyleChanged(Sender IObject) {
	LCL().SysCallN(4999, m.Instance(), GetObjectUintptr(Sender))
}

func (m *TShape) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4980, m.Instance(), m.dragDropPtr)
}

func (m *TShape) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4981, m.Instance(), m.dragOverPtr)
}

func (m *TShape) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4982, m.Instance(), m.endDockPtr)
}

func (m *TShape) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4983, m.Instance(), m.endDragPtr)
}

func (m *TShape) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4984, m.Instance(), m.mouseDownPtr)
}

func (m *TShape) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4985, m.Instance(), m.mouseEnterPtr)
}

func (m *TShape) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4986, m.Instance(), m.mouseLeavePtr)
}

func (m *TShape) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4987, m.Instance(), m.mouseMovePtr)
}

func (m *TShape) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4988, m.Instance(), m.mouseUpPtr)
}

func (m *TShape) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4989, m.Instance(), m.mouseWheelPtr)
}

func (m *TShape) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4990, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TShape) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4994, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TShape) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4991, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TShape) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4992, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TShape) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4993, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TShape) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4995, m.Instance(), m.paintPtr)
}

func (m *TShape) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4996, m.Instance(), m.startDockPtr)
}

func (m *TShape) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4997, m.Instance(), m.startDragPtr)
}
