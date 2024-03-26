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

// IComboBoxEx Parent: ICustomComboBoxEx
type IComboBoxEx interface {
	ICustomComboBoxEx
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
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TComboBoxEx Parent: TCustomComboBoxEx
type TComboBoxEx struct {
	TCustomComboBoxEx
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	getItemsPtr       uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	selectPtr         uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewComboBoxEx(TheOwner IComponent) IComboBoxEx {
	r1 := LCL().SysCallN(592, GetObjectUintptr(TheOwner))
	return AsComboBoxEx(r1)
}

func (m *TComboBoxEx) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(590, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TComboBoxEx) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(590, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) DragCursor() TCursor {
	r1 := LCL().SysCallN(593, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TComboBoxEx) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(593, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) DragKind() TDragKind {
	r1 := LCL().SysCallN(594, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TComboBoxEx) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(594, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) DragMode() TDragMode {
	r1 := LCL().SysCallN(595, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TComboBoxEx) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(595, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) ItemHeight() int32 {
	r1 := LCL().SysCallN(596, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBoxEx) SetItemHeight(AValue int32) {
	LCL().SysCallN(596, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) ItemWidth() int32 {
	r1 := LCL().SysCallN(597, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBoxEx) SetItemWidth(AValue int32) {
	LCL().SysCallN(597, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) MaxLength() int32 {
	r1 := LCL().SysCallN(598, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBoxEx) SetMaxLength(AValue int32) {
	LCL().SysCallN(598, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) ParentColor() bool {
	r1 := LCL().SysCallN(599, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBoxEx) SetParentColor(AValue bool) {
	LCL().SysCallN(599, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBoxEx) ParentFont() bool {
	r1 := LCL().SysCallN(600, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBoxEx) SetParentFont(AValue bool) {
	LCL().SysCallN(600, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBoxEx) ParentShowHint() bool {
	r1 := LCL().SysCallN(601, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBoxEx) SetParentShowHint(AValue bool) {
	LCL().SysCallN(601, 1, m.Instance(), PascalBool(AValue))
}

func ComboBoxExClass() TClass {
	ret := LCL().SysCallN(591)
	return TClass(ret)
}

func (m *TComboBoxEx) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(602, m.Instance(), m.changePtr)
}

func (m *TComboBoxEx) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(603, m.Instance(), m.closeUpPtr)
}

func (m *TComboBoxEx) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(604, m.Instance(), m.contextPopupPtr)
}

func (m *TComboBoxEx) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(605, m.Instance(), m.dblClickPtr)
}

func (m *TComboBoxEx) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(606, m.Instance(), m.dragDropPtr)
}

func (m *TComboBoxEx) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(607, m.Instance(), m.dragOverPtr)
}

func (m *TComboBoxEx) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(608, m.Instance(), m.dropDownPtr)
}

func (m *TComboBoxEx) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(609, m.Instance(), m.editingDonePtr)
}

func (m *TComboBoxEx) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(610, m.Instance(), m.endDockPtr)
}

func (m *TComboBoxEx) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(611, m.Instance(), m.endDragPtr)
}

func (m *TComboBoxEx) SetOnGetItems(fn TNotifyEvent) {
	if m.getItemsPtr != 0 {
		RemoveEventElement(m.getItemsPtr)
	}
	m.getItemsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(612, m.Instance(), m.getItemsPtr)
}

func (m *TComboBoxEx) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(613, m.Instance(), m.mouseDownPtr)
}

func (m *TComboBoxEx) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(614, m.Instance(), m.mouseEnterPtr)
}

func (m *TComboBoxEx) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(615, m.Instance(), m.mouseLeavePtr)
}

func (m *TComboBoxEx) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(616, m.Instance(), m.mouseMovePtr)
}

func (m *TComboBoxEx) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(617, m.Instance(), m.mouseUpPtr)
}

func (m *TComboBoxEx) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(618, m.Instance(), m.mouseWheelPtr)
}

func (m *TComboBoxEx) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(619, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TComboBoxEx) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(620, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TComboBoxEx) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(621, m.Instance(), m.selectPtr)
}

func (m *TComboBoxEx) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(622, m.Instance(), m.startDockPtr)
}

func (m *TComboBoxEx) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(623, m.Instance(), m.startDragPtr)
}
