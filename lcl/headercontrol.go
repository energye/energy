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

// IHeaderControl Parent: ICustomHeaderControl
type IHeaderControl interface {
	ICustomHeaderControl
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
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
}

// THeaderControl Parent: TCustomHeaderControl
type THeaderControl struct {
	TCustomHeaderControl
	contextPopupPtr    uintptr
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
}

func NewHeaderControl(AOwner IComponent) IHeaderControl {
	r1 := LCL().SysCallN(3289, GetObjectUintptr(AOwner))
	return AsHeaderControl(r1)
}

func (m *THeaderControl) DragCursor() TCursor {
	r1 := LCL().SysCallN(3290, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *THeaderControl) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3290, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderControl) DragKind() TDragKind {
	r1 := LCL().SysCallN(3291, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *THeaderControl) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3291, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderControl) DragMode() TDragMode {
	r1 := LCL().SysCallN(3292, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *THeaderControl) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3292, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderControl) ParentFont() bool {
	r1 := LCL().SysCallN(3293, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *THeaderControl) SetParentFont(AValue bool) {
	LCL().SysCallN(3293, 1, m.Instance(), PascalBool(AValue))
}

func (m *THeaderControl) ParentShowHint() bool {
	r1 := LCL().SysCallN(3294, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *THeaderControl) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3294, 1, m.Instance(), PascalBool(AValue))
}

func HeaderControlClass() TClass {
	ret := LCL().SysCallN(3288)
	return TClass(ret)
}

func (m *THeaderControl) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3295, m.Instance(), m.contextPopupPtr)
}

func (m *THeaderControl) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3296, m.Instance(), m.dragDropPtr)
}

func (m *THeaderControl) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3297, m.Instance(), m.dragOverPtr)
}

func (m *THeaderControl) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3298, m.Instance(), m.endDockPtr)
}

func (m *THeaderControl) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3299, m.Instance(), m.endDragPtr)
}

func (m *THeaderControl) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3300, m.Instance(), m.mouseDownPtr)
}

func (m *THeaderControl) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3301, m.Instance(), m.mouseEnterPtr)
}

func (m *THeaderControl) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3302, m.Instance(), m.mouseLeavePtr)
}

func (m *THeaderControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3303, m.Instance(), m.mouseMovePtr)
}

func (m *THeaderControl) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3304, m.Instance(), m.mouseUpPtr)
}

func (m *THeaderControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3305, m.Instance(), m.mouseWheelPtr)
}

func (m *THeaderControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3306, m.Instance(), m.mouseWheelDownPtr)
}

func (m *THeaderControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3310, m.Instance(), m.mouseWheelUpPtr)
}

func (m *THeaderControl) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3307, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *THeaderControl) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3308, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *THeaderControl) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3309, m.Instance(), m.mouseWheelRightPtr)
}
