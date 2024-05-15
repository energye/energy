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

// IComboBox Parent: ICustomComboBox
type IComboBox interface {
	ICustomComboBox
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
	SetOnDrawItem(fn TDrawItemEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMeasureItem(fn TMeasureItemEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TComboBox Parent: TCustomComboBox
type TComboBox struct {
	TCustomComboBox
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	drawItemPtr       uintptr
	endDragPtr        uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
	getItemsPtr       uintptr
	measureItemPtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	selectPtr         uintptr
	startDragPtr      uintptr
}

func NewComboBox(TheOwner IComponent) IComboBox {
	r1 := LCL().SysCallN(816, GetObjectUintptr(TheOwner))
	return AsComboBox(r1)
}

func (m *TComboBox) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(814, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TComboBox) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(814, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(817, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TComboBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(817, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(818, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TComboBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(818, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(819, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TComboBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(819, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) ItemHeight() int32 {
	r1 := LCL().SysCallN(820, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBox) SetItemHeight(AValue int32) {
	LCL().SysCallN(820, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) ItemWidth() int32 {
	r1 := LCL().SysCallN(821, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBox) SetItemWidth(AValue int32) {
	LCL().SysCallN(821, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) MaxLength() int32 {
	r1 := LCL().SysCallN(822, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBox) SetMaxLength(AValue int32) {
	LCL().SysCallN(822, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) ParentColor() bool {
	r1 := LCL().SysCallN(823, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetParentColor(AValue bool) {
	LCL().SysCallN(823, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBox) ParentFont() bool {
	r1 := LCL().SysCallN(824, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetParentFont(AValue bool) {
	LCL().SysCallN(824, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(825, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(825, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBox) Sorted() bool {
	r1 := LCL().SysCallN(848, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetSorted(AValue bool) {
	LCL().SysCallN(848, 1, m.Instance(), PascalBool(AValue))
}

func ComboBoxClass() TClass {
	ret := LCL().SysCallN(815)
	return TClass(ret)
}

func (m *TComboBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(826, m.Instance(), m.changePtr)
}

func (m *TComboBox) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(827, m.Instance(), m.closeUpPtr)
}

func (m *TComboBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(828, m.Instance(), m.contextPopupPtr)
}

func (m *TComboBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(829, m.Instance(), m.dblClickPtr)
}

func (m *TComboBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(830, m.Instance(), m.dragDropPtr)
}

func (m *TComboBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(831, m.Instance(), m.dragOverPtr)
}

func (m *TComboBox) SetOnDrawItem(fn TDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(832, m.Instance(), m.drawItemPtr)
}

func (m *TComboBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(835, m.Instance(), m.endDragPtr)
}

func (m *TComboBox) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(833, m.Instance(), m.dropDownPtr)
}

func (m *TComboBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(834, m.Instance(), m.editingDonePtr)
}

func (m *TComboBox) SetOnGetItems(fn TNotifyEvent) {
	if m.getItemsPtr != 0 {
		RemoveEventElement(m.getItemsPtr)
	}
	m.getItemsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(836, m.Instance(), m.getItemsPtr)
}

func (m *TComboBox) SetOnMeasureItem(fn TMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(837, m.Instance(), m.measureItemPtr)
}

func (m *TComboBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(838, m.Instance(), m.mouseDownPtr)
}

func (m *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(839, m.Instance(), m.mouseEnterPtr)
}

func (m *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(840, m.Instance(), m.mouseLeavePtr)
}

func (m *TComboBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(841, m.Instance(), m.mouseMovePtr)
}

func (m *TComboBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(842, m.Instance(), m.mouseUpPtr)
}

func (m *TComboBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(843, m.Instance(), m.mouseWheelPtr)
}

func (m *TComboBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(844, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TComboBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(845, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TComboBox) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(846, m.Instance(), m.selectPtr)
}

func (m *TComboBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(847, m.Instance(), m.startDragPtr)
}
