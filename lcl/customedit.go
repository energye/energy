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

// ICustomEdit Parent: IWinControl
type ICustomEdit interface {
	IWinControl
	Alignment() TAlignment                           // property
	SetAlignment(AValue TAlignment)                  // property
	BorderStyle() TBorderStyle                       // property
	SetBorderStyle(AValue TBorderStyle)              // property
	CanUndo() bool                                   // property
	CaretPos() (resultPoint TPoint)                  // property
	SetCaretPos(AValue *TPoint)                      // property
	CharCase() TEditCharCase                         // property
	SetCharCase(AValue TEditCharCase)                // property
	EchoMode() TEchoMode                             // property
	SetEchoMode(AValue TEchoMode)                    // property
	EmulatedTextHintStatus() TEmulatedTextHintStatus // property
	HideSelection() bool                             // property
	SetHideSelection(AValue bool)                    // property
	MaxLength() int32                                // property
	SetMaxLength(AValue int32)                       // property
	Modified() bool                                  // property
	SetModified(AValue bool)                         // property
	NumbersOnly() bool                               // property
	SetNumbersOnly(AValue bool)                      // property
	PasswordChar() Char                              // property
	SetPasswordChar(AValue Char)                     // property
	ReadOnly() bool                                  // property
	SetReadOnly(AValue bool)                         // property
	SelLength() int32                                // property
	SetSelLength(AValue int32)                       // property
	SelStart() int32                                 // property
	SetSelStart(AValue int32)                        // property
	SelText() string                                 // property
	SetSelText(AValue string)                        // property
	Text() string                                    // property
	SetText(AValue string)                           // property
	TextHint() string                                // property
	SetTextHint(AValue string)                       // property
	Clear()                                          // procedure
	SelectAll()                                      // procedure
	ClearSelection()                                 // procedure
	CopyToClipboard()                                // procedure
	CutToClipboard()                                 // procedure
	PasteFromClipboard()                             // procedure
	Undo()                                           // procedure
	SetOnChange(fn TNotifyEvent)                     // property event
}

// TCustomEdit Parent: TWinControl
type TCustomEdit struct {
	TWinControl
	changePtr uintptr
}

func NewCustomEdit(AOwner IComponent) ICustomEdit {
	r1 := LCL().SysCallN(1430, GetObjectUintptr(AOwner))
	return AsCustomEdit(r1)
}

func (m *TCustomEdit) Alignment() TAlignment {
	r1 := LCL().SysCallN(1421, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomEdit) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(1421, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1422, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomEdit) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1422, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) CanUndo() bool {
	r1 := LCL().SysCallN(1423, m.Instance())
	return GoBool(r1)
}

func (m *TCustomEdit) CaretPos() (resultPoint TPoint) {
	LCL().SysCallN(1424, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TCustomEdit) SetCaretPos(AValue *TPoint) {
	LCL().SysCallN(1424, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TCustomEdit) CharCase() TEditCharCase {
	r1 := LCL().SysCallN(1425, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomEdit) SetCharCase(AValue TEditCharCase) {
	LCL().SysCallN(1425, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) EchoMode() TEchoMode {
	r1 := LCL().SysCallN(1432, 0, m.Instance(), 0)
	return TEchoMode(r1)
}

func (m *TCustomEdit) SetEchoMode(AValue TEchoMode) {
	LCL().SysCallN(1432, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) EmulatedTextHintStatus() TEmulatedTextHintStatus {
	r1 := LCL().SysCallN(1433, m.Instance())
	return TEmulatedTextHintStatus(r1)
}

func (m *TCustomEdit) HideSelection() bool {
	r1 := LCL().SysCallN(1434, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetHideSelection(AValue bool) {
	LCL().SysCallN(1434, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) MaxLength() int32 {
	r1 := LCL().SysCallN(1435, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetMaxLength(AValue int32) {
	LCL().SysCallN(1435, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) Modified() bool {
	r1 := LCL().SysCallN(1436, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetModified(AValue bool) {
	LCL().SysCallN(1436, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) NumbersOnly() bool {
	r1 := LCL().SysCallN(1437, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetNumbersOnly(AValue bool) {
	LCL().SysCallN(1437, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) PasswordChar() Char {
	r1 := LCL().SysCallN(1438, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomEdit) SetPasswordChar(AValue Char) {
	LCL().SysCallN(1438, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) ReadOnly() bool {
	r1 := LCL().SysCallN(1440, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetReadOnly(AValue bool) {
	LCL().SysCallN(1440, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) SelLength() int32 {
	r1 := LCL().SysCallN(1441, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetSelLength(AValue int32) {
	LCL().SysCallN(1441, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) SelStart() int32 {
	r1 := LCL().SysCallN(1442, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetSelStart(AValue int32) {
	LCL().SysCallN(1442, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) SelText() string {
	r1 := LCL().SysCallN(1443, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetSelText(AValue string) {
	LCL().SysCallN(1443, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomEdit) Text() string {
	r1 := LCL().SysCallN(1446, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetText(AValue string) {
	LCL().SysCallN(1446, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomEdit) TextHint() string {
	r1 := LCL().SysCallN(1447, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetTextHint(AValue string) {
	LCL().SysCallN(1447, 1, m.Instance(), PascalStr(AValue))
}

func CustomEditClass() TClass {
	ret := LCL().SysCallN(1426)
	return TClass(ret)
}

func (m *TCustomEdit) Clear() {
	LCL().SysCallN(1427, m.Instance())
}

func (m *TCustomEdit) SelectAll() {
	LCL().SysCallN(1444, m.Instance())
}

func (m *TCustomEdit) ClearSelection() {
	LCL().SysCallN(1428, m.Instance())
}

func (m *TCustomEdit) CopyToClipboard() {
	LCL().SysCallN(1429, m.Instance())
}

func (m *TCustomEdit) CutToClipboard() {
	LCL().SysCallN(1431, m.Instance())
}

func (m *TCustomEdit) PasteFromClipboard() {
	LCL().SysCallN(1439, m.Instance())
}

func (m *TCustomEdit) Undo() {
	LCL().SysCallN(1448, m.Instance())
}

func (m *TCustomEdit) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1445, m.Instance(), m.changePtr)
}
