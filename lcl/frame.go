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

// IFrame Parent: ICustomFrame
type IFrame interface {
	ICustomFrame
	AutoScroll() bool                                  // property
	SetAutoScroll(AValue bool)                         // property
	DragCursor() TCursor                               // property
	SetDragCursor(AValue TCursor)                      // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	LCLVersion() string                                // property
	SetLCLVersion(AValue string)                       // property
	ParentColor() bool                                 // property
	SetParentColor(AValue bool)                        // property
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
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

// TFrame Parent: TCustomFrame
type TFrame struct {
	TCustomFrame
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
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
	mouseWheelHorzPtr    uintptr
	mouseWheelLeftPtr    uintptr
	mouseWheelRightPtr   uintptr
	startDockPtr         uintptr
	startDragPtr         uintptr
}

func NewFrame(TheOwner IComponent) IFrame {
	r1 := LCL().SysCallN(3131, GetObjectUintptr(TheOwner))
	return AsFrame(r1)
}

func (m *TFrame) AutoScroll() bool {
	r1 := LCL().SysCallN(3129, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetAutoScroll(AValue bool) {
	LCL().SysCallN(3129, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFrame) DragCursor() TCursor {
	r1 := LCL().SysCallN(3132, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TFrame) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3132, 1, m.Instance(), uintptr(AValue))
}

func (m *TFrame) DragKind() TDragKind {
	r1 := LCL().SysCallN(3133, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TFrame) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3133, 1, m.Instance(), uintptr(AValue))
}

func (m *TFrame) DragMode() TDragMode {
	r1 := LCL().SysCallN(3134, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TFrame) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3134, 1, m.Instance(), uintptr(AValue))
}

func (m *TFrame) LCLVersion() string {
	r1 := LCL().SysCallN(3135, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFrame) SetLCLVersion(AValue string) {
	LCL().SysCallN(3135, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFrame) ParentColor() bool {
	r1 := LCL().SysCallN(3136, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetParentColor(AValue bool) {
	LCL().SysCallN(3136, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFrame) ParentFont() bool {
	r1 := LCL().SysCallN(3137, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetParentFont(AValue bool) {
	LCL().SysCallN(3137, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFrame) ParentShowHint() bool {
	r1 := LCL().SysCallN(3138, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3138, 1, m.Instance(), PascalBool(AValue))
}

func FrameClass() TClass {
	ret := LCL().SysCallN(3130)
	return TClass(ret)
}

func (m *TFrame) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3139, m.Instance(), m.constrainedResizePtr)
}

func (m *TFrame) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3140, m.Instance(), m.contextPopupPtr)
}

func (m *TFrame) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3141, m.Instance(), m.dblClickPtr)
}

func (m *TFrame) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3142, m.Instance(), m.dragDropPtr)
}

func (m *TFrame) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3143, m.Instance(), m.dragOverPtr)
}

func (m *TFrame) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3144, m.Instance(), m.endDockPtr)
}

func (m *TFrame) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3145, m.Instance(), m.endDragPtr)
}

func (m *TFrame) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3146, m.Instance(), m.getSiteInfoPtr)
}

func (m *TFrame) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3147, m.Instance(), m.mouseDownPtr)
}

func (m *TFrame) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3148, m.Instance(), m.mouseEnterPtr)
}

func (m *TFrame) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3149, m.Instance(), m.mouseLeavePtr)
}

func (m *TFrame) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3150, m.Instance(), m.mouseMovePtr)
}

func (m *TFrame) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3151, m.Instance(), m.mouseUpPtr)
}

func (m *TFrame) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3152, m.Instance(), m.mouseWheelPtr)
}

func (m *TFrame) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3153, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TFrame) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3157, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TFrame) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3154, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TFrame) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3155, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TFrame) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3156, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TFrame) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3158, m.Instance(), m.startDockPtr)
}

func (m *TFrame) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3159, m.Instance(), m.startDragPtr)
}
