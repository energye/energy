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
	"unsafe"
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
	r1 := LCL().SysCallN(1011, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetAutoSizeHeightIsEditHeight(AValue bool) {
	LCL().SysCallN(1011, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) Alignment() TAlignment {
	r1 := LCL().SysCallN(1010, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomAbstractGroupedEdit) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(1010, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) CanUndo() bool {
	r1 := LCL().SysCallN(1012, m.Instance())
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) CaretPos() (resultPoint TPoint) {
	LCL().SysCallN(1013, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TCustomAbstractGroupedEdit) SetCaretPos(AValue *TPoint) {
	LCL().SysCallN(1013, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TCustomAbstractGroupedEdit) CharCase() TEditCharCase {
	r1 := LCL().SysCallN(1014, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomAbstractGroupedEdit) SetCharCase(AValue TEditCharCase) {
	LCL().SysCallN(1014, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) ParentColor() bool {
	r1 := LCL().SysCallN(1025, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetParentColor(AValue bool) {
	LCL().SysCallN(1025, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) EchoMode() TEchoMode {
	r1 := LCL().SysCallN(1020, 0, m.Instance(), 0)
	return TEchoMode(r1)
}

func (m *TCustomAbstractGroupedEdit) SetEchoMode(AValue TEchoMode) {
	LCL().SysCallN(1020, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) HideSelection() bool {
	r1 := LCL().SysCallN(1021, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetHideSelection(AValue bool) {
	LCL().SysCallN(1021, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) MaxLength() int32 {
	r1 := LCL().SysCallN(1022, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetMaxLength(AValue int32) {
	LCL().SysCallN(1022, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) Modified() bool {
	r1 := LCL().SysCallN(1023, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetModified(AValue bool) {
	LCL().SysCallN(1023, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) NumbersOnly() bool {
	r1 := LCL().SysCallN(1024, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetNumbersOnly(AValue bool) {
	LCL().SysCallN(1024, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) PasswordChar() Char {
	r1 := LCL().SysCallN(1026, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomAbstractGroupedEdit) SetPasswordChar(AValue Char) {
	LCL().SysCallN(1026, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) ReadOnly() bool {
	r1 := LCL().SysCallN(1028, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetReadOnly(AValue bool) {
	LCL().SysCallN(1028, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelLength() int32 {
	r1 := LCL().SysCallN(1029, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelLength(AValue int32) {
	LCL().SysCallN(1029, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelStart() int32 {
	r1 := LCL().SysCallN(1030, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelStart(AValue int32) {
	LCL().SysCallN(1030, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelText() string {
	r1 := LCL().SysCallN(1031, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelText(AValue string) {
	LCL().SysCallN(1031, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAbstractGroupedEdit) Text() string {
	r1 := LCL().SysCallN(1049, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetText(AValue string) {
	LCL().SysCallN(1049, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAbstractGroupedEdit) TextHint() string {
	r1 := LCL().SysCallN(1050, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetTextHint(AValue string) {
	LCL().SysCallN(1050, 1, m.Instance(), PascalStr(AValue))
}

func CustomAbstractGroupedEditClass() TClass {
	ret := LCL().SysCallN(1015)
	return TClass(ret)
}

func (m *TCustomAbstractGroupedEdit) Clear() {
	LCL().SysCallN(1016, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ClearSelection() {
	LCL().SysCallN(1017, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CopyToClipboard() {
	LCL().SysCallN(1018, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CutToClipboard() {
	LCL().SysCallN(1019, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) PasteFromClipboard() {
	LCL().SysCallN(1027, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) SelectAll() {
	LCL().SysCallN(1032, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) Undo() {
	LCL().SysCallN(1051, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ValidateEdit() {
	LCL().SysCallN(1052, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1033, m.Instance(), m.changePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1034, m.Instance(), m.contextPopupPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1035, m.Instance(), m.dblClickPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1036, m.Instance(), m.dragDropPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1037, m.Instance(), m.dragOverPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1038, m.Instance(), m.editingDonePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1039, m.Instance(), m.endDragPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1040, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1041, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1042, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1043, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1045, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1047, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1046, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1044, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1048, m.Instance(), m.startDragPtr)
}
