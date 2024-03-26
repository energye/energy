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

// ICEFLabelButtonComponent Parent: ICEFButtonComponent
type ICEFLabelButtonComponent interface {
	ICEFButtonComponent
	// Text
	//  Gets and sets the text shown on the LabelButton. By default |text| will also be
	//  used as the accessible name.
	Text() string // property
	// SetText Set Text
	SetText(AValue string) // property
	// Image
	//  Returns the image shown for |button_state|. If no image exists for that
	//  state then the image for CEF_BUTTON_STATE_NORMAL will be returned.
	Image(buttonstate TCefButtonState) ICefImage // property
	// SetImage Set Image
	SetImage(buttonstate TCefButtonState, AValue ICefImage) // property
	// AsMenuButton
	//  Returns this LabelButton as a MenuButton or NULL if this is not a
	//  MenuButton.
	AsMenuButton() ICefMenuButton // property
	// CreateLabelButton
	//  Create a new LabelButton. |aText| will be shown on the LabelButton and used as the default
	//  accessible name.
	CreateLabelButton(aText string) // procedure
	// SetTextColor
	//  Sets the text color shown for the specified button |for_state| to |color|.
	SetTextColor(forstate TCefButtonState, color TCefColor) // procedure
	// SetEnabledTextColors
	//  Sets the text colors shown for the non-disabled states to |color|.
	SetEnabledTextColors(color TCefColor) // procedure
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
	// SetHorizontalAlignment
	//  Sets the horizontal alignment; reversed in RTL. Default is
	//  CEF_HORIZONTAL_ALIGNMENT_CENTER.
	SetHorizontalAlignment(alignment TCefHorizontalAlignment) // procedure
	// SetMinimumSizeForCefSize
	//  Reset the minimum size of this LabelButton to |size|.
	SetMinimumSizeForCefSize(size *TCefSize) // procedure
	// SetMaximumSizeForCefSize
	//  Reset the maximum size of this LabelButton to |size|.
	SetMaximumSizeForCefSize(size *TCefSize) // procedure
}

// TCEFLabelButtonComponent Parent: TCEFButtonComponent
type TCEFLabelButtonComponent struct {
	TCEFButtonComponent
}

func NewCEFLabelButtonComponent(aOwner IComponent) ICEFLabelButtonComponent {
	r1 := CEF().SysCallN(145, GetObjectUintptr(aOwner))
	return AsCEFLabelButtonComponent(r1)
}

func (m *TCEFLabelButtonComponent) Text() string {
	r1 := CEF().SysCallN(154, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFLabelButtonComponent) SetText(AValue string) {
	CEF().SysCallN(154, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFLabelButtonComponent) Image(buttonstate TCefButtonState) ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(147, 0, m.Instance(), uintptr(buttonstate), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCEFLabelButtonComponent) SetImage(buttonstate TCefButtonState, AValue ICefImage) {
	CEF().SysCallN(147, 1, m.Instance(), uintptr(buttonstate), GetObjectUintptr(AValue), uintptr(buttonstate), GetObjectUintptr(AValue))
}

func (m *TCEFLabelButtonComponent) AsMenuButton() ICefMenuButton {
	var resultCefMenuButton uintptr
	CEF().SysCallN(143, m.Instance(), uintptr(unsafePointer(&resultCefMenuButton)))
	return AsCefMenuButton(resultCefMenuButton)
}

func CEFLabelButtonComponentClass() TClass {
	ret := CEF().SysCallN(144)
	return TClass(ret)
}

func (m *TCEFLabelButtonComponent) CreateLabelButton(aText string) {
	CEF().SysCallN(146, m.Instance(), PascalStr(aText))
}

func (m *TCEFLabelButtonComponent) SetTextColor(forstate TCefButtonState, color TCefColor) {
	CEF().SysCallN(153, m.Instance(), uintptr(forstate), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetEnabledTextColors(color TCefColor) {
	CEF().SysCallN(148, m.Instance(), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetFontList(fontlist string) {
	CEF().SysCallN(149, m.Instance(), PascalStr(fontlist))
}

func (m *TCEFLabelButtonComponent) SetHorizontalAlignment(alignment TCefHorizontalAlignment) {
	CEF().SysCallN(150, m.Instance(), uintptr(alignment))
}

func (m *TCEFLabelButtonComponent) SetMinimumSizeForCefSize(size *TCefSize) {
	CEF().SysCallN(152, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCEFLabelButtonComponent) SetMaximumSizeForCefSize(size *TCefSize) {
	CEF().SysCallN(151, m.Instance(), uintptr(unsafePointer(size)))
}
