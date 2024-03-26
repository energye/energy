//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefTextfield Parent: ICefView
//
//	A Textfield supports editing of text. This control is custom rendered with
//	no platform-specific code. Methods must be called on the browser process UI
//	thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_textfield_capi.h">CEF source file: /include/capi/views/cef_textfield_capi.h (cef_textfield_t)</a>
type ICefTextfield interface {
	ICefView
	// IsPasswordInput
	//  Returns true(1) if the text will be displayed as asterisks.
	IsPasswordInput() bool // function
	// IsReadOnly
	//  Returns true(1) if the text is read-only.
	IsReadOnly() bool // function
	// GetText
	//  Returns the currently displayed text.
	GetText() string // function
	// HasSelection
	//  Returns true(1) if there is any selected text.
	HasSelection() bool // function
	// GetSelectedText
	//  Returns the currently selected text.
	GetSelectedText() string // function
	// GetSelectedRange
	//  Returns the selected logical text range.
	GetSelectedRange() (resultCefRange TCefRange) // function
	// GetCursorPosition
	//  Returns the current cursor position.
	GetCursorPosition() NativeUInt // function
	// GetTextColor
	//  Returns the text color.
	GetTextColor() TCefColor // function
	// GetSelectionTextColor
	//  Returns the selection text color.
	GetSelectionTextColor() TCefColor // function
	// GetSelectionBackgroundColor
	//  Returns the selection background color.
	GetSelectionBackgroundColor() TCefColor // function
	// IsCommandEnabled
	//  Returns true(1) if the action associated with the specified command id is
	//  enabled. See additional comments on execute_command().
	IsCommandEnabled(commandid TCefTextFieldCommands) bool // function
	// GetPlaceholderText
	//  Returns the placeholder text that will be displayed when the Textfield is
	//  NULL.
	GetPlaceholderText() string // function
	// SetPasswordInput
	//  Sets whether the text will be displayed as asterisks.
	SetPasswordInput(passwordinput bool) // procedure
	// SetReadOnly
	//  Sets whether the text will read-only.
	SetReadOnly(readonly bool) // procedure
	// SetText
	//  Sets the contents to |text|. The cursor will be moved to end of the text
	//  if the current position is outside of the text range.
	SetText(text string) // procedure
	// AppendText
	//  Appends |text| to the previously-existing text.
	AppendText(text string) // procedure
	// InsertOrReplaceText
	//  Inserts |text| at the current cursor position replacing any selected text.
	InsertOrReplaceText(text string) // procedure
	// SelectAll
	//  Selects all text. If |reversed| is true(1) the range will end at the
	//  logical beginning of the text; this generally shows the leading portion of
	//  text that overflows its display area.
	SelectAll(reversed bool) // procedure
	// ClearSelection
	//  Clears the text selection and sets the caret to the end.
	ClearSelection() // procedure
	// SelectRange
	//  Selects the specified logical text range.
	SelectRange(range_ *TCefRange) // procedure
	// SetTextColor
	//  Sets the text color.
	SetTextColor(color TCefColor) // procedure
	// SetSelectionTextColor
	//  Sets the selection text color.
	SetSelectionTextColor(color TCefColor) // procedure
	// SetSelectionBackgroundColor
	//  Sets the selection background color.
	SetSelectionBackgroundColor(color TCefColor) // procedure
	// SetFontList
	//  Sets the font list. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>",
	//  where: - FONT_FAMILY_LIST is a comma-separated list of font family names,
	//  - STYLES is an optional space-separated list of style names(case-
	//  sensitive
	//  "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//  Here are examples of valid font description strings: - "Arial, Helvetica,
	//  Bold Italic 14px" - "Arial, 14px"
	SetFontList(fontlist string) // procedure
	// ApplyTextColor
	//  Applies |color| to the specified |range| without changing the default
	//  color. If |range| is NULL the color will be set on the complete text
	//  contents.
	ApplyTextColor(color TCefColor, range_ *TCefRange) // procedure
	// ApplyTextStyle
	//  Applies |style| to the specified |range| without changing the default
	//  style. If |add| is true(1) the style will be added, otherwise the style
	//  will be removed. If |range| is NULL the style will be set on the complete
	//  text contents.
	ApplyTextStyle(style TCefTextStyle, add bool, range_ *TCefRange) // procedure
	// ExecuteCommand
	//  Performs the action associated with the specified command id.
	ExecuteCommand(commandid TCefTextFieldCommands) // procedure
	// ClearEditHistory
	//  Clears Edit history.
	ClearEditHistory() // procedure
	// SetPlaceholderText
	//  Sets the placeholder text that will be displayed when the Textfield is
	//  NULL.
	SetPlaceholderText(text string) // procedure
	// SetPlaceholderTextColor
	//  Sets the placeholder text color.
	SetPlaceholderTextColor(color TCefColor) // procedure
	// SetAccessibleName
	//  Set the accessible name that will be exposed to assistive technology(AT).
	SetAccessibleName(name string) // procedure
}

// TCefTextfield Parent: TCefView
//
//	A Textfield supports editing of text. This control is custom rendered with
//	no platform-specific code. Methods must be called on the browser process UI
//	thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_textfield_capi.h">CEF source file: /include/capi/views/cef_textfield_capi.h (cef_textfield_t)</a>
type TCefTextfield struct {
	TCefView
}

// TextfieldRef -> ICefTextfield
var TextfieldRef textfield

// textfield TCefTextfield Ref
type textfield uintptr

// UnWrap
//
//	Returns a ICefTextfield instance using a PCefTextfield data pointer.
func (m *textfield) UnWrap(data uintptr) ICefTextfield {
	var resultCefTextfield uintptr
	CEF().SysCallN(1454, uintptr(data), uintptr(unsafePointer(&resultCefTextfield)))
	return AsCefTextfield(resultCefTextfield)
}

// CreateTextField
//
//	Create a new Textfield.
func (m *textfield) CreateTextField(delegate ICefTextfieldDelegate) ICefTextfield {
	var resultCefTextfield uintptr
	CEF().SysCallN(1427, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefTextfield)))
	return AsCefTextfield(resultCefTextfield)
}

func (m *TCefTextfield) IsPasswordInput() bool {
	r1 := CEF().SysCallN(1440, m.Instance())
	return GoBool(r1)
}

func (m *TCefTextfield) IsReadOnly() bool {
	r1 := CEF().SysCallN(1441, m.Instance())
	return GoBool(r1)
}

func (m *TCefTextfield) GetText() string {
	r1 := CEF().SysCallN(1435, m.Instance())
	return GoStr(r1)
}

func (m *TCefTextfield) HasSelection() bool {
	r1 := CEF().SysCallN(1437, m.Instance())
	return GoBool(r1)
}

func (m *TCefTextfield) GetSelectedText() string {
	r1 := CEF().SysCallN(1432, m.Instance())
	return GoStr(r1)
}

func (m *TCefTextfield) GetSelectedRange() (resultCefRange TCefRange) {
	CEF().SysCallN(1431, m.Instance(), uintptr(unsafePointer(&resultCefRange)))
	return
}

func (m *TCefTextfield) GetCursorPosition() NativeUInt {
	r1 := CEF().SysCallN(1429, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefTextfield) GetTextColor() TCefColor {
	r1 := CEF().SysCallN(1436, m.Instance())
	return TCefColor(r1)
}

func (m *TCefTextfield) GetSelectionTextColor() TCefColor {
	r1 := CEF().SysCallN(1434, m.Instance())
	return TCefColor(r1)
}

func (m *TCefTextfield) GetSelectionBackgroundColor() TCefColor {
	r1 := CEF().SysCallN(1433, m.Instance())
	return TCefColor(r1)
}

func (m *TCefTextfield) IsCommandEnabled(commandid TCefTextFieldCommands) bool {
	r1 := CEF().SysCallN(1439, m.Instance(), uintptr(commandid))
	return GoBool(r1)
}

func (m *TCefTextfield) GetPlaceholderText() string {
	r1 := CEF().SysCallN(1430, m.Instance())
	return GoStr(r1)
}

func (m *TCefTextfield) SetPasswordInput(passwordinput bool) {
	CEF().SysCallN(1446, m.Instance(), PascalBool(passwordinput))
}

func (m *TCefTextfield) SetReadOnly(readonly bool) {
	CEF().SysCallN(1449, m.Instance(), PascalBool(readonly))
}

func (m *TCefTextfield) SetText(text string) {
	CEF().SysCallN(1452, m.Instance(), PascalStr(text))
}

func (m *TCefTextfield) AppendText(text string) {
	CEF().SysCallN(1422, m.Instance(), PascalStr(text))
}

func (m *TCefTextfield) InsertOrReplaceText(text string) {
	CEF().SysCallN(1438, m.Instance(), PascalStr(text))
}

func (m *TCefTextfield) SelectAll(reversed bool) {
	CEF().SysCallN(1442, m.Instance(), PascalBool(reversed))
}

func (m *TCefTextfield) ClearSelection() {
	CEF().SysCallN(1426, m.Instance())
}

func (m *TCefTextfield) SelectRange(range_ *TCefRange) {
	CEF().SysCallN(1443, m.Instance(), uintptr(unsafePointer(range_)))
}

func (m *TCefTextfield) SetTextColor(color TCefColor) {
	CEF().SysCallN(1453, m.Instance(), uintptr(color))
}

func (m *TCefTextfield) SetSelectionTextColor(color TCefColor) {
	CEF().SysCallN(1451, m.Instance(), uintptr(color))
}

func (m *TCefTextfield) SetSelectionBackgroundColor(color TCefColor) {
	CEF().SysCallN(1450, m.Instance(), uintptr(color))
}

func (m *TCefTextfield) SetFontList(fontlist string) {
	CEF().SysCallN(1445, m.Instance(), PascalStr(fontlist))
}

func (m *TCefTextfield) ApplyTextColor(color TCefColor, range_ *TCefRange) {
	CEF().SysCallN(1423, m.Instance(), uintptr(color), uintptr(unsafePointer(range_)))
}

func (m *TCefTextfield) ApplyTextStyle(style TCefTextStyle, add bool, range_ *TCefRange) {
	CEF().SysCallN(1424, m.Instance(), uintptr(style), PascalBool(add), uintptr(unsafePointer(range_)))
}

func (m *TCefTextfield) ExecuteCommand(commandid TCefTextFieldCommands) {
	CEF().SysCallN(1428, m.Instance(), uintptr(commandid))
}

func (m *TCefTextfield) ClearEditHistory() {
	CEF().SysCallN(1425, m.Instance())
}

func (m *TCefTextfield) SetPlaceholderText(text string) {
	CEF().SysCallN(1447, m.Instance(), PascalStr(text))
}

func (m *TCefTextfield) SetPlaceholderTextColor(color TCefColor) {
	CEF().SysCallN(1448, m.Instance(), uintptr(color))
}

func (m *TCefTextfield) SetAccessibleName(name string) {
	CEF().SysCallN(1444, m.Instance(), PascalStr(name))
}
