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

// IPageControl Parent: ICustomTabControl
type IPageControl interface {
	ICustomTabControl
	ActivePageIndex() int32                                                    // property
	SetActivePageIndex(AValue int32)                                           // property
	PagesForTabSheet(Index int32) ITabSheet                                    // property
	ActivePage() ITabSheet                                                     // property
	SetActivePage(AValue ITabSheet)                                            // property
	DragCursor() TCursor                                                       // property
	SetDragCursor(AValue TCursor)                                              // property
	DragKind() TDragKind                                                       // property
	SetDragKind(AValue TDragKind)                                              // property
	DragMode() TDragMode                                                       // property
	SetDragMode(AValue TDragMode)                                              // property
	ParentFont() bool                                                          // property
	SetParentFont(AValue bool)                                                 // property
	ParentShowHint() bool                                                      // property
	SetParentShowHint(AValue bool)                                             // property
	TabIndex() int32                                                           // property
	SetTabIndex(AValue int32)                                                  // property
	FindNextPage(CurPage ITabSheet, GoForward, CheckTabVisible bool) ITabSheet // function
	IndexOfTabAt(X, Y int32) int32                                             // function
	IndexOfTabAt1(P *TPoint) int32                                             // function
	IndexOfPageAt(X, Y int32) int32                                            // function
	IndexOfPageAt1(P *TPoint) int32                                            // function
	AddTabSheet() ITabSheet                                                    // function
	Clear()                                                                    // procedure
	SelectNextPage(GoForward bool)                                             // procedure
	SelectNextPage1(GoForward bool, CheckTabVisible bool)                      // procedure
	SetOnGetDockCaption(fn TGetDockCaptionEvent)                               // property event
	SetOnChange(fn TNotifyEvent)                                               // property event
	SetOnContextPopup(fn TContextPopupEvent)                                   // property event
	SetOnDragDrop(fn TDragDropEvent)                                           // property event
	SetOnDragOver(fn TDragOverEvent)                                           // property event
	SetOnEndDock(fn TEndDragEvent)                                             // property event
	SetOnEndDrag(fn TEndDragEvent)                                             // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)                                     // property event
	SetOnMouseDown(fn TMouseEvent)                                             // property event
	SetOnMouseEnter(fn TNotifyEvent)                                           // property event
	SetOnMouseLeave(fn TNotifyEvent)                                           // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                         // property event
	SetOnMouseUp(fn TMouseEvent)                                               // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                       // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                             // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                               // property event
	SetOnStartDock(fn TStartDockEvent)                                         // property event
	SetOnStartDrag(fn TStartDragEvent)                                         // property event
}

// TPageControl Parent: TCustomTabControl
type TPageControl struct {
	TCustomTabControl
	getDockCaptionPtr uintptr
	changePtr         uintptr
	contextPopupPtr   uintptr
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

func NewPageControl(TheOwner IComponent) IPageControl {
	r1 := LCL().SysCallN(4388, GetObjectUintptr(TheOwner))
	return AsPageControl(r1)
}

func (m *TPageControl) ActivePageIndex() int32 {
	r1 := LCL().SysCallN(4384, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageControl) SetActivePageIndex(AValue int32) {
	LCL().SysCallN(4384, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) PagesForTabSheet(Index int32) ITabSheet {
	r1 := LCL().SysCallN(4397, m.Instance(), uintptr(Index))
	return AsTabSheet(r1)
}

func (m *TPageControl) ActivePage() ITabSheet {
	r1 := LCL().SysCallN(4383, 0, m.Instance(), 0)
	return AsTabSheet(r1)
}

func (m *TPageControl) SetActivePage(AValue ITabSheet) {
	LCL().SysCallN(4383, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPageControl) DragCursor() TCursor {
	r1 := LCL().SysCallN(4389, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TPageControl) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4389, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) DragKind() TDragKind {
	r1 := LCL().SysCallN(4390, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TPageControl) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4390, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) DragMode() TDragMode {
	r1 := LCL().SysCallN(4391, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TPageControl) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4391, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) ParentFont() bool {
	r1 := LCL().SysCallN(4398, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPageControl) SetParentFont(AValue bool) {
	LCL().SysCallN(4398, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPageControl) ParentShowHint() bool {
	r1 := LCL().SysCallN(4399, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPageControl) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4399, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPageControl) TabIndex() int32 {
	r1 := LCL().SysCallN(4420, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageControl) SetTabIndex(AValue int32) {
	LCL().SysCallN(4420, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) FindNextPage(CurPage ITabSheet, GoForward, CheckTabVisible bool) ITabSheet {
	r1 := LCL().SysCallN(4392, m.Instance(), GetObjectUintptr(CurPage), PascalBool(GoForward), PascalBool(CheckTabVisible))
	return AsTabSheet(r1)
}

func (m *TPageControl) IndexOfTabAt(X, Y int32) int32 {
	r1 := LCL().SysCallN(4395, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TPageControl) IndexOfTabAt1(P *TPoint) int32 {
	r1 := LCL().SysCallN(4396, m.Instance(), uintptr(unsafePointer(P)))
	return int32(r1)
}

func (m *TPageControl) IndexOfPageAt(X, Y int32) int32 {
	r1 := LCL().SysCallN(4393, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TPageControl) IndexOfPageAt1(P *TPoint) int32 {
	r1 := LCL().SysCallN(4394, m.Instance(), uintptr(unsafePointer(P)))
	return int32(r1)
}

func (m *TPageControl) AddTabSheet() ITabSheet {
	r1 := LCL().SysCallN(4385, m.Instance())
	return AsTabSheet(r1)
}

func PageControlClass() TClass {
	ret := LCL().SysCallN(4386)
	return TClass(ret)
}

func (m *TPageControl) Clear() {
	LCL().SysCallN(4387, m.Instance())
}

func (m *TPageControl) SelectNextPage(GoForward bool) {
	LCL().SysCallN(4400, m.Instance(), PascalBool(GoForward))
}

func (m *TPageControl) SelectNextPage1(GoForward bool, CheckTabVisible bool) {
	LCL().SysCallN(4401, m.Instance(), PascalBool(GoForward), PascalBool(CheckTabVisible))
}

func (m *TPageControl) SetOnGetDockCaption(fn TGetDockCaptionEvent) {
	if m.getDockCaptionPtr != 0 {
		RemoveEventElement(m.getDockCaptionPtr)
	}
	m.getDockCaptionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4408, m.Instance(), m.getDockCaptionPtr)
}

func (m *TPageControl) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4402, m.Instance(), m.changePtr)
}

func (m *TPageControl) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4403, m.Instance(), m.contextPopupPtr)
}

func (m *TPageControl) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4404, m.Instance(), m.dragDropPtr)
}

func (m *TPageControl) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4405, m.Instance(), m.dragOverPtr)
}

func (m *TPageControl) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4406, m.Instance(), m.endDockPtr)
}

func (m *TPageControl) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4407, m.Instance(), m.endDragPtr)
}

func (m *TPageControl) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4409, m.Instance(), m.getSiteInfoPtr)
}

func (m *TPageControl) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4410, m.Instance(), m.mouseDownPtr)
}

func (m *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4411, m.Instance(), m.mouseEnterPtr)
}

func (m *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4412, m.Instance(), m.mouseLeavePtr)
}

func (m *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4413, m.Instance(), m.mouseMovePtr)
}

func (m *TPageControl) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4414, m.Instance(), m.mouseUpPtr)
}

func (m *TPageControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4415, m.Instance(), m.mouseWheelPtr)
}

func (m *TPageControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4416, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TPageControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4417, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TPageControl) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4418, m.Instance(), m.startDockPtr)
}

func (m *TPageControl) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4419, m.Instance(), m.startDragPtr)
}
