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

// ITrackBar Parent: ICustomTrackBar
type ITrackBar interface {
	ICustomTrackBar
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
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

// TTrackBar Parent: TCustomTrackBar
type TTrackBar struct {
	TCustomTrackBar
	contextPopupPtr    uintptr
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

func NewTrackBar(AOwner IComponent) ITrackBar {
	r1 := LCL().SysCallN(5548, GetObjectUintptr(AOwner))
	return AsTrackBar(r1)
}

func (m *TTrackBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(5549, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TTrackBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5549, 1, m.Instance(), uintptr(AValue))
}

func (m *TTrackBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(5550, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TTrackBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5550, 1, m.Instance(), uintptr(AValue))
}

func (m *TTrackBar) ParentColor() bool {
	r1 := LCL().SysCallN(5551, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTrackBar) SetParentColor(AValue bool) {
	LCL().SysCallN(5551, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTrackBar) ParentFont() bool {
	r1 := LCL().SysCallN(5552, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTrackBar) SetParentFont(AValue bool) {
	LCL().SysCallN(5552, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTrackBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(5553, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTrackBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5553, 1, m.Instance(), PascalBool(AValue))
}

func TrackBarClass() TClass {
	ret := LCL().SysCallN(5547)
	return TClass(ret)
}

func (m *TTrackBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5554, m.Instance(), m.contextPopupPtr)
}

func (m *TTrackBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5555, m.Instance(), m.dragDropPtr)
}

func (m *TTrackBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5556, m.Instance(), m.dragOverPtr)
}

func (m *TTrackBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5557, m.Instance(), m.endDragPtr)
}

func (m *TTrackBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5558, m.Instance(), m.mouseDownPtr)
}

func (m *TTrackBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5559, m.Instance(), m.mouseEnterPtr)
}

func (m *TTrackBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5560, m.Instance(), m.mouseLeavePtr)
}

func (m *TTrackBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5561, m.Instance(), m.mouseMovePtr)
}

func (m *TTrackBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5562, m.Instance(), m.mouseUpPtr)
}

func (m *TTrackBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5563, m.Instance(), m.mouseWheelPtr)
}

func (m *TTrackBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5564, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TTrackBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5568, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TTrackBar) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5565, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TTrackBar) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5566, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TTrackBar) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5567, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TTrackBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5569, m.Instance(), m.startDragPtr)
}
