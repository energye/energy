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

// ILinkLabel Parent: ICustomLabel
type ILinkLabel interface {
	ICustomLabel
	Alignment() TAlignment                   // property
	SetAlignment(AValue TAlignment)          // property
	DragCursor() TCursor                     // property
	SetDragCursor(AValue TCursor)            // property
	DragKind() TDragKind                     // property
	SetDragKind(AValue TDragKind)            // property
	DragMode() TDragMode                     // property
	SetDragMode(AValue TDragMode)            // property
	ParentColor() bool                       // property
	SetParentColor(AValue bool)              // property
	ParentFont() bool                        // property
	SetParentFont(AValue bool)               // property
	ParentShowHint() bool                    // property
	SetParentShowHint(AValue bool)           // property
	SetOnContextPopup(fn TContextPopupEvent) // property event
	SetOnDblClick(fn TNotifyEvent)           // property event
	SetOnDragDrop(fn TDragDropEvent)         // property event
	SetOnDragOver(fn TDragOverEvent)         // property event
	SetOnEndDock(fn TEndDragEvent)           // property event
	SetOnEndDrag(fn TEndDragEvent)           // property event
	SetOnMouseDown(fn TMouseEvent)           // property event
	SetOnMouseEnter(fn TNotifyEvent)         // property event
	SetOnMouseLeave(fn TNotifyEvent)         // property event
	SetOnMouseMove(fn TMouseMoveEvent)       // property event
	SetOnMouseUp(fn TMouseEvent)             // property event
	SetOnStartDock(fn TStartDockEvent)       // property event
	SetOnStartDrag(fn TStartDragEvent)       // property event
	SetOnLinkClick(fn TSysLinkEvent)         // property event
}

// TLinkLabel Parent: TCustomLabel
type TLinkLabel struct {
	TCustomLabel
	contextPopupPtr uintptr
	dblClickPtr     uintptr
	dragDropPtr     uintptr
	dragOverPtr     uintptr
	endDockPtr      uintptr
	endDragPtr      uintptr
	mouseDownPtr    uintptr
	mouseEnterPtr   uintptr
	mouseLeavePtr   uintptr
	mouseMovePtr    uintptr
	mouseUpPtr      uintptr
	startDockPtr    uintptr
	startDragPtr    uintptr
	linkClickPtr    uintptr
}

func NewLinkLabel(AOwner IComponent) ILinkLabel {
	r1 := LCL().SysCallN(3309, GetObjectUintptr(AOwner))
	return AsLinkLabel(r1)
}

func (m *TLinkLabel) Alignment() TAlignment {
	r1 := LCL().SysCallN(3307, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLinkLabel) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3307, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) DragCursor() TCursor {
	r1 := LCL().SysCallN(3310, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLinkLabel) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3310, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) DragKind() TDragKind {
	r1 := LCL().SysCallN(3311, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLinkLabel) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3311, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) DragMode() TDragMode {
	r1 := LCL().SysCallN(3312, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLinkLabel) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3312, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) ParentColor() bool {
	r1 := LCL().SysCallN(3313, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLinkLabel) SetParentColor(AValue bool) {
	LCL().SysCallN(3313, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLinkLabel) ParentFont() bool {
	r1 := LCL().SysCallN(3314, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLinkLabel) SetParentFont(AValue bool) {
	LCL().SysCallN(3314, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLinkLabel) ParentShowHint() bool {
	r1 := LCL().SysCallN(3315, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLinkLabel) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3315, 1, m.Instance(), PascalBool(AValue))
}

func LinkLabelClass() TClass {
	ret := LCL().SysCallN(3308)
	return TClass(ret)
}

func (m *TLinkLabel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3316, m.Instance(), m.contextPopupPtr)
}

func (m *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3317, m.Instance(), m.dblClickPtr)
}

func (m *TLinkLabel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3318, m.Instance(), m.dragDropPtr)
}

func (m *TLinkLabel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3319, m.Instance(), m.dragOverPtr)
}

func (m *TLinkLabel) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3320, m.Instance(), m.endDockPtr)
}

func (m *TLinkLabel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3321, m.Instance(), m.endDragPtr)
}

func (m *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3323, m.Instance(), m.mouseDownPtr)
}

func (m *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3324, m.Instance(), m.mouseEnterPtr)
}

func (m *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3325, m.Instance(), m.mouseLeavePtr)
}

func (m *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3326, m.Instance(), m.mouseMovePtr)
}

func (m *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3327, m.Instance(), m.mouseUpPtr)
}

func (m *TLinkLabel) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3328, m.Instance(), m.startDockPtr)
}

func (m *TLinkLabel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3329, m.Instance(), m.startDragPtr)
}

func (m *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
	if m.linkClickPtr != 0 {
		RemoveEventElement(m.linkClickPtr)
	}
	m.linkClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3322, m.Instance(), m.linkClickPtr)
}
