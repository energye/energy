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

// IPanel Parent: ICustomPanel
type IPanel interface {
	ICustomPanel
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
	ShowAccelChar() bool                            // property
	SetShowAccelChar(AValue bool)                   // property
	VerticalAlignment() TVerticalAlignment          // property
	SetVerticalAlignment(AValue TVerticalAlignment) // property
	Wordwrap() bool                                 // property
	SetWordwrap(AValue bool)                        // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)          // property event
	SetOnGetDockCaption(fn TGetDockCaptionEvent)    // property event
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
	SetOnStartDock(fn TStartDockEvent)              // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TPanel Parent: TCustomPanel
type TPanel struct {
	TCustomPanel
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDockPtr         uintptr
	endDragPtr         uintptr
	getSiteInfoPtr     uintptr
	getDockCaptionPtr  uintptr
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
	startDockPtr       uintptr
	startDragPtr       uintptr
}

func NewPanel(TheOwner IComponent) IPanel {
	r1 := LCL().SysCallN(3821, GetObjectUintptr(TheOwner))
	return AsPanel(r1)
}

func (m *TPanel) DragCursor() TCursor {
	r1 := LCL().SysCallN(3822, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TPanel) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3822, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) DragKind() TDragKind {
	r1 := LCL().SysCallN(3823, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TPanel) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3823, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) DragMode() TDragMode {
	r1 := LCL().SysCallN(3824, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TPanel) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3824, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) ParentFont() bool {
	r1 := LCL().SysCallN(3825, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetParentFont(AValue bool) {
	LCL().SysCallN(3825, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPanel) ParentShowHint() bool {
	r1 := LCL().SysCallN(3826, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3826, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPanel) ShowAccelChar() bool {
	r1 := LCL().SysCallN(3848, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetShowAccelChar(AValue bool) {
	LCL().SysCallN(3848, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPanel) VerticalAlignment() TVerticalAlignment {
	r1 := LCL().SysCallN(3849, 0, m.Instance(), 0)
	return TVerticalAlignment(r1)
}

func (m *TPanel) SetVerticalAlignment(AValue TVerticalAlignment) {
	LCL().SysCallN(3849, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) Wordwrap() bool {
	r1 := LCL().SysCallN(3850, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetWordwrap(AValue bool) {
	LCL().SysCallN(3850, 1, m.Instance(), PascalBool(AValue))
}

func PanelClass() TClass {
	ret := LCL().SysCallN(3820)
	return TClass(ret)
}

func (m *TPanel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3827, m.Instance(), m.contextPopupPtr)
}

func (m *TPanel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3828, m.Instance(), m.dblClickPtr)
}

func (m *TPanel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3829, m.Instance(), m.dragDropPtr)
}

func (m *TPanel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3830, m.Instance(), m.dragOverPtr)
}

func (m *TPanel) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3831, m.Instance(), m.endDockPtr)
}

func (m *TPanel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3832, m.Instance(), m.endDragPtr)
}

func (m *TPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3834, m.Instance(), m.getSiteInfoPtr)
}

func (m *TPanel) SetOnGetDockCaption(fn TGetDockCaptionEvent) {
	if m.getDockCaptionPtr != 0 {
		RemoveEventElement(m.getDockCaptionPtr)
	}
	m.getDockCaptionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3833, m.Instance(), m.getDockCaptionPtr)
}

func (m *TPanel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3835, m.Instance(), m.mouseDownPtr)
}

func (m *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3836, m.Instance(), m.mouseEnterPtr)
}

func (m *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3837, m.Instance(), m.mouseLeavePtr)
}

func (m *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3838, m.Instance(), m.mouseMovePtr)
}

func (m *TPanel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3839, m.Instance(), m.mouseUpPtr)
}

func (m *TPanel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3840, m.Instance(), m.mouseWheelPtr)
}

func (m *TPanel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3841, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TPanel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3845, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TPanel) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3842, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TPanel) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3843, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TPanel) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3844, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TPanel) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3846, m.Instance(), m.startDockPtr)
}

func (m *TPanel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3847, m.Instance(), m.startDragPtr)
}
