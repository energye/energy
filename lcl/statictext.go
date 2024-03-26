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

// IStaticText Parent: ICustomStaticText
type IStaticText interface {
	ICustomStaticText
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
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
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TStaticText Parent: TCustomStaticText
type TStaticText struct {
	TCustomStaticText
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
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
	startDragPtr       uintptr
}

func NewStaticText(AOwner IComponent) IStaticText {
	r1 := LCL().SysCallN(4429, GetObjectUintptr(AOwner))
	return AsStaticText(r1)
}

func (m *TStaticText) DragCursor() TCursor {
	r1 := LCL().SysCallN(4430, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStaticText) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4430, 1, m.Instance(), uintptr(AValue))
}

func (m *TStaticText) DragKind() TDragKind {
	r1 := LCL().SysCallN(4431, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStaticText) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4431, 1, m.Instance(), uintptr(AValue))
}

func (m *TStaticText) DragMode() TDragMode {
	r1 := LCL().SysCallN(4432, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStaticText) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4432, 1, m.Instance(), uintptr(AValue))
}

func (m *TStaticText) ParentFont() bool {
	r1 := LCL().SysCallN(4434, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStaticText) SetParentFont(AValue bool) {
	LCL().SysCallN(4434, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStaticText) ParentColor() bool {
	r1 := LCL().SysCallN(4433, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStaticText) SetParentColor(AValue bool) {
	LCL().SysCallN(4433, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStaticText) ParentShowHint() bool {
	r1 := LCL().SysCallN(4435, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStaticText) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4435, 1, m.Instance(), PascalBool(AValue))
}

func StaticTextClass() TClass {
	ret := LCL().SysCallN(4428)
	return TClass(ret)
}

func (m *TStaticText) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4436, m.Instance(), m.contextPopupPtr)
}

func (m *TStaticText) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4437, m.Instance(), m.dblClickPtr)
}

func (m *TStaticText) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4438, m.Instance(), m.dragDropPtr)
}

func (m *TStaticText) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4439, m.Instance(), m.dragOverPtr)
}

func (m *TStaticText) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4440, m.Instance(), m.endDragPtr)
}

func (m *TStaticText) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4441, m.Instance(), m.mouseDownPtr)
}

func (m *TStaticText) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4442, m.Instance(), m.mouseEnterPtr)
}

func (m *TStaticText) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4443, m.Instance(), m.mouseLeavePtr)
}

func (m *TStaticText) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4444, m.Instance(), m.mouseMovePtr)
}

func (m *TStaticText) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4445, m.Instance(), m.mouseUpPtr)
}

func (m *TStaticText) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4446, m.Instance(), m.mouseWheelPtr)
}

func (m *TStaticText) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4447, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TStaticText) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4451, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TStaticText) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4448, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TStaticText) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4449, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TStaticText) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4450, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TStaticText) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4452, m.Instance(), m.startDragPtr)
}
