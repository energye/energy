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

// IColorBox Parent: ICustomColorBox
type IColorBox interface {
	ICustomColorBox
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ItemHeight() int32                             // property
	SetItemHeight(AValue int32)                    // property
	ItemWidth() int32                              // property
	SetItemWidth(AValue int32)                     // property
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
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
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

// TColorBox Parent: TCustomColorBox
type TColorBox struct {
	TCustomColorBox
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
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

func NewColorBox(AOwner IComponent) IColorBox {
	r1 := LCL().SysCallN(716, GetObjectUintptr(AOwner))
	return AsColorBox(r1)
}

func (m *TColorBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(717, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TColorBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(717, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(718, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TColorBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(718, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) ItemHeight() int32 {
	r1 := LCL().SysCallN(719, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorBox) SetItemHeight(AValue int32) {
	LCL().SysCallN(719, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) ItemWidth() int32 {
	r1 := LCL().SysCallN(720, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorBox) SetItemWidth(AValue int32) {
	LCL().SysCallN(720, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) ParentColor() bool {
	r1 := LCL().SysCallN(721, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorBox) SetParentColor(AValue bool) {
	LCL().SysCallN(721, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorBox) ParentFont() bool {
	r1 := LCL().SysCallN(722, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorBox) SetParentFont(AValue bool) {
	LCL().SysCallN(722, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(723, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(723, 1, m.Instance(), PascalBool(AValue))
}

func ColorBoxClass() TClass {
	ret := LCL().SysCallN(715)
	return TClass(ret)
}

func (m *TColorBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(724, m.Instance(), m.changePtr)
}

func (m *TColorBox) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(725, m.Instance(), m.closeUpPtr)
}

func (m *TColorBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(726, m.Instance(), m.contextPopupPtr)
}

func (m *TColorBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(727, m.Instance(), m.dblClickPtr)
}

func (m *TColorBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(728, m.Instance(), m.dragDropPtr)
}

func (m *TColorBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(729, m.Instance(), m.dragOverPtr)
}

func (m *TColorBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(732, m.Instance(), m.endDragPtr)
}

func (m *TColorBox) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(730, m.Instance(), m.dropDownPtr)
}

func (m *TColorBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(731, m.Instance(), m.editingDonePtr)
}

func (m *TColorBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(733, m.Instance(), m.mouseDownPtr)
}

func (m *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(734, m.Instance(), m.mouseEnterPtr)
}

func (m *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(735, m.Instance(), m.mouseLeavePtr)
}

func (m *TColorBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(736, m.Instance(), m.mouseMovePtr)
}

func (m *TColorBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(737, m.Instance(), m.mouseUpPtr)
}

func (m *TColorBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(738, m.Instance(), m.mouseWheelPtr)
}

func (m *TColorBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(739, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TColorBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(740, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TColorBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(742, m.Instance(), m.startDragPtr)
}

func (m *TColorBox) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(741, m.Instance(), m.selectPtr)
}
