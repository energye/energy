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

// ICEFTextfieldComponent Parent: ICEFViewComponent
type ICEFTextfieldComponent interface {
	ICEFViewComponent
	// PasswordInput
	//  Returns true(1) if the text will be displayed as asterisks.
	PasswordInput() bool // property
	// SetPasswordInput Set PasswordInput
	SetPasswordInput(AValue bool) // property
	// ReadOnly
	//  Returns true(1) if the text is read-only.
	ReadOnly() bool // property
	// SetReadOnly Set ReadOnly
	SetReadOnly(AValue bool) // property
	// Text
	//  Returns the currently displayed text.
	Text() string // property
	// SetText Set Text
	SetText(AValue string) // property
	// SelectedText
	//  Returns the currently selected text.
	SelectedText() string // property
	// SelectedRange
	//  Returns the selected logical text range.
	SelectedRange() (resultCefRange TCefRange) // property
	// SetSelectedRange Set SelectedRange
	SetSelectedRange(AValue *TCefRange) // property
	// CursorPosition
	//  Returns the current cursor position.
	CursorPosition() NativeUInt // property
	// TextColor
	//  Returns the text color.
	TextColor() TCefColor // property
	// SetTextColor Set TextColor
	SetTextColor(AValue TCefColor) // property
	// SelectionTextColor
	//  Returns the selection text color.
	SelectionTextColor() TCefColor // property
	// SetSelectionTextColor Set SelectionTextColor
	SetSelectionTextColor(AValue TCefColor) // property
	// SelectionBackgroundColor
	//  Returns the selection background color.
	SelectionBackgroundColor() TCefColor // property
	// SetSelectionBackgroundColor Set SelectionBackgroundColor
	SetSelectionBackgroundColor(AValue TCefColor) // property
	// PlaceholderText
	//  Returns the placeholder text that will be displayed when the Textfield is
	//  NULL.
	PlaceholderText() string // property
	// SetPlaceholderText Set PlaceholderText
	SetPlaceholderText(AValue string) // property
	// HasSelection
	//  Returns true(1) if there is any selected text.
	HasSelection() bool // property
	// IsCommandEnabled
	//  Returns true(1) if the action associated with the specified command id is
	//  enabled. See additional comments on execute_command().
	IsCommandEnabled(commandid TCefTextFieldCommands) bool // function
	// CreateTextField
	//  Create a new Textfield.
	CreateTextField() // procedure
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
	// SetFontList
	//  Sets the font list. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>",
	//  where:
	//  <code>
	//  - FONT_FAMILY_LIST is a comma-separated list of font family names,
	//  - STYLES is an optional space-separated list of style names(case-sensitive
	//  "Bold" and "Italic" are supported), and
	//  - SIZE is an integer font size in pixels with the suffix "px".
	//  </code>
	//  Here are examples of valid font description strings:
	//  <code>
	//  - "Arial, Helvetica, Bold Italic 14px"
	//  - "Arial, 14px"
	//  </code>
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
	// SetAccessibleName
	//  Set the accessible name that will be exposed to assistive technology(AT).
	SetAccessibleName(name string) // procedure
	// SetPlaceholderTextColor
	//  Sets the placeholder text color.
	SetPlaceholderTextColor(color TCefColor) // procedure
	// SetOnTextfieldKeyEvent
	//  Called when |textfield| recieves a keyboard event. |event| contains
	//  information about the keyboard event. Return true(1) if the keyboard
	//  event was handled or false(0) otherwise for default handling.
	SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEvent) // property event
	// SetOnAfterUserAction
	//  Called after performing a user action that may change |textfield|.
	SetOnAfterUserAction(fn TOnAfterUserAction) // property event
}

// TCEFTextfieldComponent Parent: TCEFViewComponent
type TCEFTextfieldComponent struct {
	TCEFViewComponent
	textfieldKeyEventPtr uintptr
	afterUserActionPtr   uintptr
}

func NewCEFTextfieldComponent(aOwner IComponent) ICEFTextfieldComponent {
	r1 := CEF().SysCallN(245, GetObjectUintptr(aOwner))
	return AsCEFTextfieldComponent(r1)
}

func (m *TCEFTextfieldComponent) PasswordInput() bool {
	r1 := CEF().SysCallN(252, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFTextfieldComponent) SetPasswordInput(AValue bool) {
	CEF().SysCallN(252, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFTextfieldComponent) ReadOnly() bool {
	r1 := CEF().SysCallN(254, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFTextfieldComponent) SetReadOnly(AValue bool) {
	CEF().SysCallN(254, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCEFTextfieldComponent) Text() string {
	r1 := CEF().SysCallN(265, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFTextfieldComponent) SetText(AValue string) {
	CEF().SysCallN(265, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFTextfieldComponent) SelectedText() string {
	r1 := CEF().SysCallN(257, m.Instance())
	return GoStr(r1)
}

func (m *TCEFTextfieldComponent) SelectedRange() (resultCefRange TCefRange) {
	CEF().SysCallN(256, 0, m.Instance(), uintptr(unsafePointer(&resultCefRange)), uintptr(unsafePointer(&resultCefRange)))
	return
}

func (m *TCEFTextfieldComponent) SetSelectedRange(AValue *TCefRange) {
	CEF().SysCallN(256, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCEFTextfieldComponent) CursorPosition() NativeUInt {
	r1 := CEF().SysCallN(247, m.Instance())
	return NativeUInt(r1)
}

func (m *TCEFTextfieldComponent) TextColor() TCefColor {
	r1 := CEF().SysCallN(266, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCEFTextfieldComponent) SetTextColor(AValue TCefColor) {
	CEF().SysCallN(266, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFTextfieldComponent) SelectionTextColor() TCefColor {
	r1 := CEF().SysCallN(259, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCEFTextfieldComponent) SetSelectionTextColor(AValue TCefColor) {
	CEF().SysCallN(259, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFTextfieldComponent) SelectionBackgroundColor() TCefColor {
	r1 := CEF().SysCallN(258, 0, m.Instance(), 0)
	return TCefColor(r1)
}

func (m *TCEFTextfieldComponent) SetSelectionBackgroundColor(AValue TCefColor) {
	CEF().SysCallN(258, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFTextfieldComponent) PlaceholderText() string {
	r1 := CEF().SysCallN(253, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFTextfieldComponent) SetPlaceholderText(AValue string) {
	CEF().SysCallN(253, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFTextfieldComponent) HasSelection() bool {
	r1 := CEF().SysCallN(249, m.Instance())
	return GoBool(r1)
}

func (m *TCEFTextfieldComponent) IsCommandEnabled(commandid TCefTextFieldCommands) bool {
	r1 := CEF().SysCallN(251, m.Instance(), uintptr(commandid))
	return GoBool(r1)
}

func CEFTextfieldComponentClass() TClass {
	ret := CEF().SysCallN(242)
	return TClass(ret)
}

func (m *TCEFTextfieldComponent) CreateTextField() {
	CEF().SysCallN(246, m.Instance())
}

func (m *TCEFTextfieldComponent) AppendText(text string) {
	CEF().SysCallN(239, m.Instance(), PascalStr(text))
}

func (m *TCEFTextfieldComponent) InsertOrReplaceText(text string) {
	CEF().SysCallN(250, m.Instance(), PascalStr(text))
}

func (m *TCEFTextfieldComponent) SelectAll(reversed bool) {
	CEF().SysCallN(255, m.Instance(), PascalBool(reversed))
}

func (m *TCEFTextfieldComponent) ClearSelection() {
	CEF().SysCallN(244, m.Instance())
}

func (m *TCEFTextfieldComponent) SetFontList(fontlist string) {
	CEF().SysCallN(261, m.Instance(), PascalStr(fontlist))
}

func (m *TCEFTextfieldComponent) ApplyTextColor(color TCefColor, range_ *TCefRange) {
	CEF().SysCallN(240, m.Instance(), uintptr(color), uintptr(unsafePointer(range_)))
}

func (m *TCEFTextfieldComponent) ApplyTextStyle(style TCefTextStyle, add bool, range_ *TCefRange) {
	CEF().SysCallN(241, m.Instance(), uintptr(style), PascalBool(add), uintptr(unsafePointer(range_)))
}

func (m *TCEFTextfieldComponent) ExecuteCommand(commandid TCefTextFieldCommands) {
	CEF().SysCallN(248, m.Instance(), uintptr(commandid))
}

func (m *TCEFTextfieldComponent) ClearEditHistory() {
	CEF().SysCallN(243, m.Instance())
}

func (m *TCEFTextfieldComponent) SetAccessibleName(name string) {
	CEF().SysCallN(260, m.Instance(), PascalStr(name))
}

func (m *TCEFTextfieldComponent) SetPlaceholderTextColor(color TCefColor) {
	CEF().SysCallN(264, m.Instance(), uintptr(color))
}

func (m *TCEFTextfieldComponent) SetOnTextfieldKeyEvent(fn TOnTextfieldKeyEvent) {
	if m.textfieldKeyEventPtr != 0 {
		RemoveEventElement(m.textfieldKeyEventPtr)
	}
	m.textfieldKeyEventPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(263, m.Instance(), m.textfieldKeyEventPtr)
}

func (m *TCEFTextfieldComponent) SetOnAfterUserAction(fn TOnAfterUserAction) {
	if m.afterUserActionPtr != 0 {
		RemoveEventElement(m.afterUserActionPtr)
	}
	m.afterUserActionPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(262, m.Instance(), m.afterUserActionPtr)
}
