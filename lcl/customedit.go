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
	r1 := LCL().SysCallN(1620, GetObjectUintptr(AOwner))
	return AsCustomEdit(r1)
}

func (m *TCustomEdit) Alignment() TAlignment {
	r1 := LCL().SysCallN(1611, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomEdit) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(1611, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1612, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomEdit) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1612, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) CanUndo() bool {
	r1 := LCL().SysCallN(1613, m.Instance())
	return GoBool(r1)
}

func (m *TCustomEdit) CaretPos() (resultPoint TPoint) {
	LCL().SysCallN(1614, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomEdit) SetCaretPos(AValue *TPoint) {
	LCL().SysCallN(1614, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomEdit) CharCase() TEditCharCase {
	r1 := LCL().SysCallN(1615, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomEdit) SetCharCase(AValue TEditCharCase) {
	LCL().SysCallN(1615, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) EchoMode() TEchoMode {
	r1 := LCL().SysCallN(1622, 0, m.Instance(), 0)
	return TEchoMode(r1)
}

func (m *TCustomEdit) SetEchoMode(AValue TEchoMode) {
	LCL().SysCallN(1622, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) EmulatedTextHintStatus() TEmulatedTextHintStatus {
	r1 := LCL().SysCallN(1623, m.Instance())
	return TEmulatedTextHintStatus(r1)
}

func (m *TCustomEdit) HideSelection() bool {
	r1 := LCL().SysCallN(1624, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetHideSelection(AValue bool) {
	LCL().SysCallN(1624, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) MaxLength() int32 {
	r1 := LCL().SysCallN(1625, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetMaxLength(AValue int32) {
	LCL().SysCallN(1625, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) Modified() bool {
	r1 := LCL().SysCallN(1626, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetModified(AValue bool) {
	LCL().SysCallN(1626, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) NumbersOnly() bool {
	r1 := LCL().SysCallN(1627, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetNumbersOnly(AValue bool) {
	LCL().SysCallN(1627, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) PasswordChar() Char {
	r1 := LCL().SysCallN(1628, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomEdit) SetPasswordChar(AValue Char) {
	LCL().SysCallN(1628, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) ReadOnly() bool {
	r1 := LCL().SysCallN(1630, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetReadOnly(AValue bool) {
	LCL().SysCallN(1630, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) SelLength() int32 {
	r1 := LCL().SysCallN(1631, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetSelLength(AValue int32) {
	LCL().SysCallN(1631, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) SelStart() int32 {
	r1 := LCL().SysCallN(1632, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetSelStart(AValue int32) {
	LCL().SysCallN(1632, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) SelText() string {
	r1 := LCL().SysCallN(1633, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetSelText(AValue string) {
	LCL().SysCallN(1633, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomEdit) Text() string {
	r1 := LCL().SysCallN(1636, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetText(AValue string) {
	LCL().SysCallN(1636, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomEdit) TextHint() string {
	r1 := LCL().SysCallN(1637, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetTextHint(AValue string) {
	LCL().SysCallN(1637, 1, m.Instance(), PascalStr(AValue))
}

func CustomEditClass() TClass {
	ret := LCL().SysCallN(1616)
	return TClass(ret)
}

func (m *TCustomEdit) Clear() {
	LCL().SysCallN(1617, m.Instance())
}

func (m *TCustomEdit) SelectAll() {
	LCL().SysCallN(1634, m.Instance())
}

func (m *TCustomEdit) ClearSelection() {
	LCL().SysCallN(1618, m.Instance())
}

func (m *TCustomEdit) CopyToClipboard() {
	LCL().SysCallN(1619, m.Instance())
}

func (m *TCustomEdit) CutToClipboard() {
	LCL().SysCallN(1621, m.Instance())
}

func (m *TCustomEdit) PasteFromClipboard() {
	LCL().SysCallN(1629, m.Instance())
}

func (m *TCustomEdit) Undo() {
	LCL().SysCallN(1638, m.Instance())
}

func (m *TCustomEdit) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1635, m.Instance(), m.changePtr)
}
