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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCEFTextFieldComponent
type TCEFTextFieldComponent struct {
	*TCEFViewComponent
}

// TextFieldComponentRef -> TCEFTextFieldComponent
var TextFieldComponentRef textFieldComponent

type textFieldComponent uintptr

func (*textFieldComponent) New(AOwner lcl.IComponent) *TCEFTextFieldComponent {
	var result uintptr
	imports.Proc(def.TextfieldComponent_Create).Call(lcl.CheckPtr(AOwner), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFTextFieldComponent{&TCEFViewComponent{instance: getInstance(result)}}
	}
	return nil
}

func (m *TCEFTextFieldComponent) CreateTextField() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_CreateTextField).Call(m.Instance())
}

func (m *TCEFTextFieldComponent) SetPasswordInput(passwordInput bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetPasswordInput).Call(m.Instance(), api.PascalBool(passwordInput))
}

func (m *TCEFTextFieldComponent) IsPasswordInput() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_IsPasswordInput).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFTextFieldComponent) SetReadOnly(readOnly bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetReadOnly).Call(m.Instance(), api.PascalBool(readOnly))
}

func (m *TCEFTextFieldComponent) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFTextFieldComponent) GetText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_GetText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFTextFieldComponent) SetText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetText).Call(m.Instance(), api.PascalStr(text))
}

func (m *TCEFTextFieldComponent) AppendText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_AppendText).Call(m.Instance(), api.PascalStr(text))
}

func (m *TCEFTextFieldComponent) InsertOrReplaceText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_InsertOrReplaceText).Call(m.Instance(), api.PascalStr(text))
}

func (m *TCEFTextFieldComponent) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_HasSelection).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFTextFieldComponent) GetSelectedText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_GetSelectedText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFTextFieldComponent) SelectAll(reversed bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SelectAll).Call(m.Instance(), api.PascalBool(reversed))
}

func (m *TCEFTextFieldComponent) ClearSelection() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_ClearSelection).Call(m.Instance())
}

// SelectRange
//
//	CEF 117 Remove
func (m *TCEFTextFieldComponent) SelectRange(range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SelectRange).Call(m.Instance(), uintptr(unsafe.Pointer(&range_)))
}

// SetSelectRange
//
//	CEF 117 ~
func (m *TCEFTextFieldComponent) SetSelectRange(range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SelectedRange).Call(consts.SetValue, m.Instance(), uintptr(unsafe.Pointer(&range_)))
}

// GetSelectRange
//
//	CEF 117 ~
func (m *TCEFTextFieldComponent) GetSelectRange() (result TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SelectedRange).Call(consts.GetValue, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// CursorPosition
//
//	CEF 117 ~
func (m *TCEFTextFieldComponent) CursorPosition() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_CursorPosition).Call(m.Instance())
	return uint32(r1)
}

func (m *TCEFTextFieldComponent) SetTextColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetTextColor).Call(m.Instance(), uintptr(color))
}

func (m *TCEFTextFieldComponent) GetTextColor() (color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_GetTextColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

func (m *TCEFTextFieldComponent) SetSelectionTextColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetSelectionTextColor).Call(m.Instance(), uintptr(color))
}

func (m *TCEFTextFieldComponent) GetSelectionTextColor() (color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_GetSelectionTextColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

func (m *TCEFTextFieldComponent) SetSelectionBackgroundColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetSelectionBackgroundColor).Call(m.Instance(), uintptr(color))
}

func (m *TCEFTextFieldComponent) GetSelectionBackgroundColor() (color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_GetSelectionBackgroundColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

func (m *TCEFTextFieldComponent) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetFontList).Call(m.Instance(), api.PascalStr(fontList))
}

func (m *TCEFTextFieldComponent) ApplyTextColor(color types.TCefColor, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_ApplyTextColor).Call(m.Instance(), uintptr(color), uintptr(unsafe.Pointer(&range_)))
}

func (m *TCEFTextFieldComponent) ApplyTextStyle(style consts.TCefTextStyle, add bool, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_ApplyTextStyle).Call(m.Instance(), uintptr(style), api.PascalBool(add), uintptr(unsafe.Pointer(&range_)))
}

func (m *TCEFTextFieldComponent) IsCommandEnabled(commandId consts.TCefTextFieldCommands) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_IsCommandEnabled).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

func (m *TCEFTextFieldComponent) ExecuteCommand(commandId consts.TCefTextFieldCommands) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_ExecuteCommand).Call(m.Instance(), uintptr(commandId))
}

func (m *TCEFTextFieldComponent) ClearEditHistory() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_ClearEditHistory).Call(m.Instance())
}

func (m *TCEFTextFieldComponent) SetPlaceholderText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetPlaceholderText).Call(m.Instance(), api.PascalStr(text))
}

func (m *TCEFTextFieldComponent) GetPlaceholderText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.TextfieldComponent_GetPlaceholderText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFTextFieldComponent) SetPlaceholderTextColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetPlaceholderTextColor).Call(m.Instance(), uintptr(color))
}

func (m *TCEFTextFieldComponent) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetAccessibleName).Call(m.Instance(), api.PascalStr(name))
}

func (m *TCEFTextFieldComponent) SetOnTextFieldKeyEvent(fn textFieldOnTextFieldKeyEvent) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetOnTextfieldKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFTextFieldComponent) SetOnAfterUserAction(fn textFieldOnAfterUserAction) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.TextfieldComponent_SetOnAfterUserAction).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
