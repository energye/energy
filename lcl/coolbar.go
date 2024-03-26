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

// ICoolBar Parent: ICustomCoolBar
type ICoolBar interface {
	ICustomCoolBar
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TCoolBar Parent: TCustomCoolBar
type TCoolBar struct {
	TCustomCoolBar
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	getSiteInfoPtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewCoolBar(AOwner IComponent) ICoolBar {
	r1 := LCL().SysCallN(986, GetObjectUintptr(AOwner))
	return AsCoolBar(r1)
}

func (m *TCoolBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(987, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCoolBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(987, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBar) DragKind() TDragKind {
	r1 := LCL().SysCallN(988, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCoolBar) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(988, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(989, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCoolBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(989, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBar) ParentColor() bool {
	r1 := LCL().SysCallN(990, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBar) SetParentColor(AValue bool) {
	LCL().SysCallN(990, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBar) ParentFont() bool {
	r1 := LCL().SysCallN(991, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBar) SetParentFont(AValue bool) {
	LCL().SysCallN(991, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(992, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(992, 1, m.Instance(), PascalBool(AValue))
}

func CoolBarClass() TClass {
	ret := LCL().SysCallN(985)
	return TClass(ret)
}

func (m *TCoolBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(993, m.Instance(), m.contextPopupPtr)
}

func (m *TCoolBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(994, m.Instance(), m.dblClickPtr)
}

func (m *TCoolBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(995, m.Instance(), m.dragDropPtr)
}

func (m *TCoolBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(996, m.Instance(), m.dragOverPtr)
}

func (m *TCoolBar) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(997, m.Instance(), m.endDockPtr)
}

func (m *TCoolBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(998, m.Instance(), m.endDragPtr)
}

func (m *TCoolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(999, m.Instance(), m.getSiteInfoPtr)
}

func (m *TCoolBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1000, m.Instance(), m.mouseDownPtr)
}

func (m *TCoolBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1001, m.Instance(), m.mouseEnterPtr)
}

func (m *TCoolBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1002, m.Instance(), m.mouseLeavePtr)
}

func (m *TCoolBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1003, m.Instance(), m.mouseMovePtr)
}

func (m *TCoolBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1004, m.Instance(), m.mouseUpPtr)
}

func (m *TCoolBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1005, m.Instance(), m.mouseWheelPtr)
}

func (m *TCoolBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1006, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCoolBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1007, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCoolBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1008, m.Instance(), m.startDockPtr)
}

func (m *TCoolBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1009, m.Instance(), m.startDragPtr)
}
