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

// IFlowPanel Parent: ICustomFlowPanel
type IFlowPanel interface {
	ICustomFlowPanel
	DragCursor() TCursor                               // property
	SetDragCursor(AValue TCursor)                      // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	ParentFont() bool                                  // property
	SetParentFont(AValue bool)                         // property
	ParentShowHint() bool                              // property
	SetParentShowHint(AValue bool)                     // property
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnContextPopup(fn TContextPopupEvent)           // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnEndDock(fn TEndDragEvent)                     // property event
	SetOnEndDrag(fn TEndDragEvent)                     // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)             // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

// TFlowPanel Parent: TCustomFlowPanel
type TFlowPanel struct {
	TCustomFlowPanel
	constrainedResizePtr uintptr
	contextPopupPtr      uintptr
	dblClickPtr          uintptr
	dragDropPtr          uintptr
	dragOverPtr          uintptr
	endDockPtr           uintptr
	endDragPtr           uintptr
	getSiteInfoPtr       uintptr
	mouseDownPtr         uintptr
	mouseEnterPtr        uintptr
	mouseLeavePtr        uintptr
	mouseMovePtr         uintptr
	mouseUpPtr           uintptr
	startDockPtr         uintptr
	startDragPtr         uintptr
}

func NewFlowPanel(AOwner IComponent) IFlowPanel {
	r1 := LCL().SysCallN(3050, GetObjectUintptr(AOwner))
	return AsFlowPanel(r1)
}

func (m *TFlowPanel) DragCursor() TCursor {
	r1 := LCL().SysCallN(3051, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TFlowPanel) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3051, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanel) DragKind() TDragKind {
	r1 := LCL().SysCallN(3052, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TFlowPanel) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3052, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanel) DragMode() TDragMode {
	r1 := LCL().SysCallN(3053, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TFlowPanel) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3053, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanel) ParentFont() bool {
	r1 := LCL().SysCallN(3054, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFlowPanel) SetParentFont(AValue bool) {
	LCL().SysCallN(3054, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFlowPanel) ParentShowHint() bool {
	r1 := LCL().SysCallN(3055, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFlowPanel) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3055, 1, m.Instance(), PascalBool(AValue))
}

func FlowPanelClass() TClass {
	ret := LCL().SysCallN(3049)
	return TClass(ret)
}

func (m *TFlowPanel) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3056, m.Instance(), m.constrainedResizePtr)
}

func (m *TFlowPanel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3057, m.Instance(), m.contextPopupPtr)
}

func (m *TFlowPanel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3058, m.Instance(), m.dblClickPtr)
}

func (m *TFlowPanel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3059, m.Instance(), m.dragDropPtr)
}

func (m *TFlowPanel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3060, m.Instance(), m.dragOverPtr)
}

func (m *TFlowPanel) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3061, m.Instance(), m.endDockPtr)
}

func (m *TFlowPanel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3062, m.Instance(), m.endDragPtr)
}

func (m *TFlowPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3063, m.Instance(), m.getSiteInfoPtr)
}

func (m *TFlowPanel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3064, m.Instance(), m.mouseDownPtr)
}

func (m *TFlowPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3065, m.Instance(), m.mouseEnterPtr)
}

func (m *TFlowPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3066, m.Instance(), m.mouseLeavePtr)
}

func (m *TFlowPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3067, m.Instance(), m.mouseMovePtr)
}

func (m *TFlowPanel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3068, m.Instance(), m.mouseUpPtr)
}

func (m *TFlowPanel) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3069, m.Instance(), m.startDockPtr)
}

func (m *TFlowPanel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3070, m.Instance(), m.startDragPtr)
}
