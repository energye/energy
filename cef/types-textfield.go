//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https//www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/types"
	"unsafe"
)

var TextFieldRef textField

type textField uintptr

func (*textField) New(delegate *ICefTextfieldDelegate) *ICefTextfield {
	var result uintptr
	imports.Proc(def.TextfieldRef_CreateTextField).Call(delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefTextfield{&ICefView{
			instance: unsafe.Pointer(result),
		}}
	}
	return nil
}

func (m *ICefTextfield) SetPasswordInput(passwordInput bool) {
	AObj.SetPasswordInput(password_input)
}

func (m *ICefTextfield) IsPasswordInput() bool {
	Result = AObj.IsPasswordInput
}

func (m *ICefTextfield) SetReadOnly(readOnly bool) {
	AObj.SetReadOnly(read_only)
}

func (m *ICefTextfield) IsReadOnly() bool {
	Result = AObj.IsReadOnly
}

func (m *ICefTextfield) GetText() string {
	Result = string(string(AObj.GetText))
}

func (m *ICefTextfield) SetText(text string) {
	AObj.SetText(stringToUStr(text_))
}

func (m *ICefTextfield) AppendText(text string) {
	AObj.AppendText(stringToUStr(text_))
}

func (m *ICefTextfield) InsertOrReplaceText(text string) {
	AObj.InsertOrReplaceText(stringToUStr(text_))
}

func (m *ICefTextfield) HasSelection() bool {
	Result = AObj.HasSelection()
}

func (m *ICefTextfield) GetSelectedText() string {
	Result = string(string(AObj.GetSelectedText))
}

func (m *ICefTextfield) SelectAll(reversed bool) {
	AObj.SelectAll(reversed)
}

func (m *ICefTextfield) ClearSelection() {
	AObj.ClearSelection
}

func (m *ICefTextfield) GetSelectedRange() TCefRange {
	Result = AObj.GetSelectedRange
}

func (m *ICefTextfield) SelectRange(range_ TCefRange) {
	AObj.SelectRange(range)
}

func (m *ICefTextfield) GetCursorPosition() types.NativeUInt {
	Result = AObj.GetCursorPosition
}

func (m *ICefTextfield) SetTextColor(color types.TCefColor) {
	AObj.SetTextColor(color)
}

func (m *ICefTextfield) GetTextColor() {
	TCefColor
	Result = AObj.GetTextColor
}

func (m *ICefTextfield) SetSelectionTextColor(color types.TCefColor) {
	AObj.SetSelectionTextColor(color)
}

func (m *ICefTextfield) GetSelectionTextColor() types.TCefColor {
	Result = AObj.GetSelectionTextColor
}

func (m *ICefTextfield) SetSelectionBackgroundColor(color types.TCefColor) {
	AObj.SetSelectionBackgroundColor(color)
}

func (m *ICefTextfield) GetSelectionBackgroundColor() types.TCefColor {
	Result = AObj.GetSelectionBackgroundColor()
}

func (m *ICefTextfield) SetFontList(fontList string) {
	AObj.SetFontList(stringToUStr(font_list))
}

func (m *ICefTextfield) ApplyTextColor(color types.TCefColor, range_ TCefRange) {
	AObj.ApplyTextColor(color, range)
}

func (m *ICefTextfield) ApplyTextStyle(style consts.TCefTextStyle, add bool, range_ TCefRange) {
	AObj.ApplyTextStyle(style, add, range)
}

func (m *ICefTextfield) IsCommandEnabled(commandId consts.TCefTextFieldCommands) bool {
	Result = AObj.IsCommandEnabled(command_id)
}

func (m *ICefTextfield) ExecuteCommand(commandId consts.TCefTextFieldCommands) {
	AObj.ExecuteCommand(command_id)
}

func (m *ICefTextfield) ClearEditHistory() {
	AObj.ClearEditHistory
}

func (m *ICefTextfield) SetPlaceholderText(text string) {
	AObj.SetPlaceholderText(stringToUStr(text_))
}

func (m *ICefTextfield) GetPlaceholderText() string {
	Result = string(string(AObj.GetPlaceholderText))
}

func (m *ICefTextfield) SetPlaceholderTextColor(color types.TCefColor) {
	AObj.SetPlaceholderTextColor(color)
}

func (m *ICefTextfield) SetAccessibleName(name string) {
	AObj.SetAccessibleName(stringToUStr(name))
}
