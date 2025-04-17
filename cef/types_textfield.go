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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefTextfield
// include/capi/views/cef_textfield_capi.h (cef_textfield_t)
type ICefTextfield struct {
	*ICefView
}

// TextFieldRef -> ICefTextfield
var TextFieldRef textField

type textField uintptr

func (*textField) New(delegate *ICefTextFieldDelegate) *ICefTextfield {
	var result uintptr
	imports.Proc(def.TextfieldRef_CreateTextField).Call(delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefTextfield{&ICefView{instance: unsafe.Pointer(result)}}
	}
	return nil
}

func (m *ICefTextfield) SetPasswordInput(passwordInput bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetPasswordInput).Call(m.Instance(), api.PascalBool(passwordInput))
}

func (m *ICefTextfield) IsPasswordInput() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.Textfield_IsPasswordInput).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefTextfield) SetReadOnly(readOnly bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetReadOnly).Call(m.Instance(), api.PascalBool(readOnly))
}

func (m *ICefTextfield) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.Textfield_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefTextfield) GetText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.Textfield_GetText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefTextfield) SetText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetText).Call(m.Instance(), api.PascalStr(text))
}

func (m *ICefTextfield) AppendText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_AppendText).Call(m.Instance(), api.PascalStr(text))
}

func (m *ICefTextfield) InsertOrReplaceText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_InsertOrReplaceText).Call(m.Instance(), api.PascalStr(text))
}

func (m *ICefTextfield) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.Textfield_HasSelection).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefTextfield) GetSelectedText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.Textfield_GetSelectedText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefTextfield) SelectAll(reversed bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SelectAll).Call(m.Instance(), api.PascalBool(reversed))
}

func (m *ICefTextfield) ClearSelection() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_ClearSelection).Call(m.Instance())
}

func (m *ICefTextfield) GetSelectedRange() (result TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_GetSelectedRange).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefTextfield) SelectRange(range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SelectRange).Call(m.Instance(), uintptr(unsafe.Pointer(&range_)))
}

func (m *ICefTextfield) GetCursorPosition() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.Textfield_GetCursorPosition).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefTextfield) SetTextColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetTextColor).Call(m.Instance(), uintptr(color))
}

func (m *ICefTextfield) GetTextColor() (color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_GetTextColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

func (m *ICefTextfield) SetSelectionTextColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetSelectionTextColor).Call(m.Instance(), uintptr(color))
}

func (m *ICefTextfield) GetSelectionTextColor() (color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_GetSelectionTextColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

func (m *ICefTextfield) SetSelectionBackgroundColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetSelectionBackgroundColor).Call(m.Instance(), uintptr(color))
}

func (m *ICefTextfield) GetSelectionBackgroundColor() (color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_GetSelectionBackgroundColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

// Sets the font list. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>",
// where:
//   - FONT_FAMILY_LIST is a comma-separated list of font family names,
//   - STYLES is an optional space-separated list of style names (case-sensitive
//     "Bold" and "Italic" are supported), and
//   - SIZE is an integer font size in pixels with the suffix "px".
//
// Here are examples of valid font description strings:
// - "Arial, Helvetica, Bold Italic 14px"
// - "Arial, 14px"
func (m *ICefTextfield) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetFontList).Call(m.Instance(), api.PascalStr(fontList))
}

func (m *ICefTextfield) ApplyTextColor(color types.TCefColor, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_ApplyTextColor).Call(m.Instance(), uintptr(color), uintptr(unsafe.Pointer(&range_)))
}

func (m *ICefTextfield) ApplyTextStyle(style consts.TCefTextStyle, add bool, range_ TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_ApplyTextStyle).Call(m.Instance(), uintptr(style), api.PascalBool(add), uintptr(unsafe.Pointer(&range_)))
}

func (m *ICefTextfield) IsCommandEnabled(commandId consts.TCefTextFieldCommands) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.Textfield_IsCommandEnabled).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

func (m *ICefTextfield) ExecuteCommand(commandId consts.TCefTextFieldCommands) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_ExecuteCommand).Call(m.Instance(), uintptr(commandId))
}

func (m *ICefTextfield) ClearEditHistory() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_ClearEditHistory).Call(m.Instance())
}

func (m *ICefTextfield) SetPlaceholderText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetPlaceholderText).Call(m.Instance(), api.PascalStr(text))
}

func (m *ICefTextfield) GetPlaceholderText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.Textfield_GetPlaceholderText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefTextfield) SetPlaceholderTextColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetPlaceholderTextColor).Call(m.Instance(), uintptr(color))
}

func (m *ICefTextfield) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Textfield_SetAccessibleName).Call(m.Instance(), api.PascalStr(name))
}
