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

// ITabSheet Parent: ICustomPage
type ITabSheet interface {
	ICustomPage
	PageControl() IPageControl                     // property
	SetPageControl(AValue IPageControl)            // property
	TabIndex() int32                               // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TTabSheet Parent: TCustomPage
type TTabSheet struct {
	TCustomPage
	contextPopupPtr   uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewTabSheet(TheOwner IComponent) ITabSheet {
	r1 := LCL().SysCallN(5326, GetObjectUintptr(TheOwner))
	return AsTabSheet(r1)
}

func (m *TTabSheet) PageControl() IPageControl {
	r1 := LCL().SysCallN(5327, 0, m.Instance(), 0)
	return AsPageControl(r1)
}

func (m *TTabSheet) SetPageControl(AValue IPageControl) {
	LCL().SysCallN(5327, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TTabSheet) TabIndex() int32 {
	r1 := LCL().SysCallN(5343, m.Instance())
	return int32(r1)
}

func (m *TTabSheet) ParentFont() bool {
	r1 := LCL().SysCallN(5328, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTabSheet) SetParentFont(AValue bool) {
	LCL().SysCallN(5328, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTabSheet) ParentShowHint() bool {
	r1 := LCL().SysCallN(5329, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTabSheet) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5329, 1, m.Instance(), PascalBool(AValue))
}

func TabSheetClass() TClass {
	ret := LCL().SysCallN(5325)
	return TClass(ret)
}

func (m *TTabSheet) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5330, m.Instance(), m.contextPopupPtr)
}

func (m *TTabSheet) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5331, m.Instance(), m.dragDropPtr)
}

func (m *TTabSheet) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5332, m.Instance(), m.dragOverPtr)
}

func (m *TTabSheet) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5333, m.Instance(), m.endDragPtr)
}

func (m *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5334, m.Instance(), m.mouseDownPtr)
}

func (m *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5335, m.Instance(), m.mouseEnterPtr)
}

func (m *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5336, m.Instance(), m.mouseLeavePtr)
}

func (m *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5337, m.Instance(), m.mouseMovePtr)
}

func (m *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5338, m.Instance(), m.mouseUpPtr)
}

func (m *TTabSheet) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5339, m.Instance(), m.mouseWheelPtr)
}

func (m *TTabSheet) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5340, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TTabSheet) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5341, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TTabSheet) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5342, m.Instance(), m.startDragPtr)
}
