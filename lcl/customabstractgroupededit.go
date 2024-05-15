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

// ICustomAbstractGroupedEdit Is Abstract Class Parent: ICustomControl
type ICustomAbstractGroupedEdit interface {
	ICustomControl
	AutoSizeHeightIsEditHeight() bool              // property
	SetAutoSizeHeightIsEditHeight(AValue bool)     // property
	Alignment() TAlignment                         // property
	SetAlignment(AValue TAlignment)                // property
	CanUndo() bool                                 // property
	CaretPos() (resultPoint TPoint)                // property
	SetCaretPos(AValue *TPoint)                    // property
	CharCase() TEditCharCase                       // property
	SetCharCase(AValue TEditCharCase)              // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	EchoMode() TEchoMode                           // property
	SetEchoMode(AValue TEchoMode)                  // property
	HideSelection() bool                           // property
	SetHideSelection(AValue bool)                  // property
	MaxLength() int32                              // property
	SetMaxLength(AValue int32)                     // property
	Modified() bool                                // property
	SetModified(AValue bool)                       // property
	NumbersOnly() bool                             // property
	SetNumbersOnly(AValue bool)                    // property
	PasswordChar() Char                            // property
	SetPasswordChar(AValue Char)                   // property
	ReadOnly() bool                                // property
	SetReadOnly(AValue bool)                       // property
	SelLength() int32                              // property
	SetSelLength(AValue int32)                     // property
	SelStart() int32                               // property
	SetSelStart(AValue int32)                      // property
	SelText() string                               // property
	SetSelText(AValue string)                      // property
	Text() string                                  // property
	SetText(AValue string)                         // property
	TextHint() string                              // property
	SetTextHint(AValue string)                     // property
	Clear()                                        // procedure
	ClearSelection()                               // procedure
	CopyToClipboard()                              // procedure
	CutToClipboard()                               // procedure
	PasteFromClipboard()                           // procedure
	SelectAll()                                    // procedure
	Undo()                                         // procedure
	ValidateEdit()                                 // procedure
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TCustomAbstractGroupedEdit Is Abstract Class Parent: TCustomControl
type TCustomAbstractGroupedEdit struct {
	TCustomControl
	changePtr         uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	editingDonePtr    uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseWheelPtr     uintptr
	mouseWheelUpPtr   uintptr
	mouseWheelDownPtr uintptr
	mouseUpPtr        uintptr
	startDragPtr      uintptr
}

func (m *TCustomAbstractGroupedEdit) AutoSizeHeightIsEditHeight() bool {
	r1 := LCL().SysCallN(1201, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetAutoSizeHeightIsEditHeight(AValue bool) {
	LCL().SysCallN(1201, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) Alignment() TAlignment {
	r1 := LCL().SysCallN(1200, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomAbstractGroupedEdit) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(1200, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) CanUndo() bool {
	r1 := LCL().SysCallN(1202, m.Instance())
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) CaretPos() (resultPoint TPoint) {
	LCL().SysCallN(1203, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomAbstractGroupedEdit) SetCaretPos(AValue *TPoint) {
	LCL().SysCallN(1203, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomAbstractGroupedEdit) CharCase() TEditCharCase {
	r1 := LCL().SysCallN(1204, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomAbstractGroupedEdit) SetCharCase(AValue TEditCharCase) {
	LCL().SysCallN(1204, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) ParentColor() bool {
	r1 := LCL().SysCallN(1215, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetParentColor(AValue bool) {
	LCL().SysCallN(1215, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) EchoMode() TEchoMode {
	r1 := LCL().SysCallN(1210, 0, m.Instance(), 0)
	return TEchoMode(r1)
}

func (m *TCustomAbstractGroupedEdit) SetEchoMode(AValue TEchoMode) {
	LCL().SysCallN(1210, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) HideSelection() bool {
	r1 := LCL().SysCallN(1211, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetHideSelection(AValue bool) {
	LCL().SysCallN(1211, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) MaxLength() int32 {
	r1 := LCL().SysCallN(1212, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetMaxLength(AValue int32) {
	LCL().SysCallN(1212, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) Modified() bool {
	r1 := LCL().SysCallN(1213, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetModified(AValue bool) {
	LCL().SysCallN(1213, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) NumbersOnly() bool {
	r1 := LCL().SysCallN(1214, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetNumbersOnly(AValue bool) {
	LCL().SysCallN(1214, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) PasswordChar() Char {
	r1 := LCL().SysCallN(1216, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomAbstractGroupedEdit) SetPasswordChar(AValue Char) {
	LCL().SysCallN(1216, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) ReadOnly() bool {
	r1 := LCL().SysCallN(1218, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetReadOnly(AValue bool) {
	LCL().SysCallN(1218, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelLength() int32 {
	r1 := LCL().SysCallN(1219, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelLength(AValue int32) {
	LCL().SysCallN(1219, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelStart() int32 {
	r1 := LCL().SysCallN(1220, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelStart(AValue int32) {
	LCL().SysCallN(1220, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelText() string {
	r1 := LCL().SysCallN(1221, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelText(AValue string) {
	LCL().SysCallN(1221, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAbstractGroupedEdit) Text() string {
	r1 := LCL().SysCallN(1239, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetText(AValue string) {
	LCL().SysCallN(1239, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAbstractGroupedEdit) TextHint() string {
	r1 := LCL().SysCallN(1240, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetTextHint(AValue string) {
	LCL().SysCallN(1240, 1, m.Instance(), PascalStr(AValue))
}

func CustomAbstractGroupedEditClass() TClass {
	ret := LCL().SysCallN(1205)
	return TClass(ret)
}

func (m *TCustomAbstractGroupedEdit) Clear() {
	LCL().SysCallN(1206, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ClearSelection() {
	LCL().SysCallN(1207, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CopyToClipboard() {
	LCL().SysCallN(1208, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CutToClipboard() {
	LCL().SysCallN(1209, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) PasteFromClipboard() {
	LCL().SysCallN(1217, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) SelectAll() {
	LCL().SysCallN(1222, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) Undo() {
	LCL().SysCallN(1241, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ValidateEdit() {
	LCL().SysCallN(1242, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1223, m.Instance(), m.changePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1224, m.Instance(), m.contextPopupPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1225, m.Instance(), m.dblClickPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1226, m.Instance(), m.dragDropPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1227, m.Instance(), m.dragOverPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1228, m.Instance(), m.editingDonePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1229, m.Instance(), m.endDragPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1230, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1231, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1232, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1233, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1235, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1237, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1236, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1234, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1238, m.Instance(), m.startDragPtr)
}
