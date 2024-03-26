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

// ICheckComboBox Parent: ICustomCheckCombo
type ICheckComboBox interface {
	ICustomCheckCombo
	BorderStyle() TBorderStyle                     // property
	SetBorderStyle(AValue TBorderStyle)            // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ItemHeight() int32                             // property
	SetItemHeight(AValue int32)                    // property
	ItemWidth() int32                              // property
	SetItemWidth(AValue int32)                     // property
	MaxLength() int32                              // property
	SetMaxLength(AValue int32)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	Sorted() bool                                  // property
	SetSorted(AValue bool)                         // property
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
}

// TCheckComboBox Parent: TCustomCheckCombo
type TCheckComboBox struct {
	TCustomCheckCombo
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
	getItemsPtr       uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
	selectPtr         uintptr
}

func NewCheckComboBox(AOwner IComponent) ICheckComboBox {
	r1 := LCL().SysCallN(404, GetObjectUintptr(AOwner))
	return AsCheckComboBox(r1)
}

func (m *TCheckComboBox) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(402, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCheckComboBox) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(402, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(405, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckComboBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(405, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(406, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCheckComboBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(406, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(407, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckComboBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(407, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) ItemHeight() int32 {
	r1 := LCL().SysCallN(408, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCheckComboBox) SetItemHeight(AValue int32) {
	LCL().SysCallN(408, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) ItemWidth() int32 {
	r1 := LCL().SysCallN(409, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCheckComboBox) SetItemWidth(AValue int32) {
	LCL().SysCallN(409, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) MaxLength() int32 {
	r1 := LCL().SysCallN(410, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCheckComboBox) SetMaxLength(AValue int32) {
	LCL().SysCallN(410, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) ParentColor() bool {
	r1 := LCL().SysCallN(411, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetParentColor(AValue bool) {
	LCL().SysCallN(411, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckComboBox) ParentFont() bool {
	r1 := LCL().SysCallN(412, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetParentFont(AValue bool) {
	LCL().SysCallN(412, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckComboBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(413, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(413, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckComboBox) Sorted() bool {
	r1 := LCL().SysCallN(434, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetSorted(AValue bool) {
	LCL().SysCallN(434, 1, m.Instance(), PascalBool(AValue))
}

func CheckComboBoxClass() TClass {
	ret := LCL().SysCallN(403)
	return TClass(ret)
}

func (m *TCheckComboBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(414, m.Instance(), m.changePtr)
}

func (m *TCheckComboBox) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(415, m.Instance(), m.closeUpPtr)
}

func (m *TCheckComboBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(416, m.Instance(), m.contextPopupPtr)
}

func (m *TCheckComboBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(417, m.Instance(), m.dblClickPtr)
}

func (m *TCheckComboBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(418, m.Instance(), m.dragDropPtr)
}

func (m *TCheckComboBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(419, m.Instance(), m.dragOverPtr)
}

func (m *TCheckComboBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(422, m.Instance(), m.endDragPtr)
}

func (m *TCheckComboBox) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(420, m.Instance(), m.dropDownPtr)
}

func (m *TCheckComboBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(421, m.Instance(), m.editingDonePtr)
}

func (m *TCheckComboBox) SetOnGetItems(fn TNotifyEvent) {
	if m.getItemsPtr != 0 {
		RemoveEventElement(m.getItemsPtr)
	}
	m.getItemsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(423, m.Instance(), m.getItemsPtr)
}

func (m *TCheckComboBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(424, m.Instance(), m.mouseDownPtr)
}

func (m *TCheckComboBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(425, m.Instance(), m.mouseEnterPtr)
}

func (m *TCheckComboBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(426, m.Instance(), m.mouseLeavePtr)
}

func (m *TCheckComboBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(427, m.Instance(), m.mouseMovePtr)
}

func (m *TCheckComboBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(428, m.Instance(), m.mouseUpPtr)
}

func (m *TCheckComboBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(429, m.Instance(), m.mouseWheelPtr)
}

func (m *TCheckComboBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(430, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCheckComboBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(431, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCheckComboBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(433, m.Instance(), m.startDragPtr)
}

func (m *TCheckComboBox) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(432, m.Instance(), m.selectPtr)
}
