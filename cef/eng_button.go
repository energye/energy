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

// ICefButton Parent: ICefView
//
//	A View representing a button. Depending on the specific type, the button
//	could be implemented by a native control or custom rendered. Methods must be
//	called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_capi.h">CEF source file: /include/capi/views/cef_button_capi.h (cef_button_t)</a>
type ICefButton interface {
	ICefView
	// AsLabelButton
	//  Returns this Button as a LabelButton or NULL if this is not a LabelButton.
	AsLabelButton() ICefLabelButton // function
	// GetState
	//  Returns the current display state of the Button.
	GetState() TCefButtonState // function
	// SetState
	//  Sets the current display state of the Button.
	SetState(state TCefButtonState) // procedure
	// SetInkDropEnabled
	//  Sets the Button will use an ink drop effect for displaying state changes.
	SetInkDropEnabled(enabled bool) // procedure
	// SetTooltipText
	//  Sets the tooltip text that will be displayed when the user hovers the
	//  mouse cursor over the Button.
	SetTooltipText(tooltiptext string) // procedure
	// SetAccessibleName
	//  Sets the accessible name that will be exposed to assistive technology
	SetAccessibleName(name string) // procedure
}

// TCefButton Parent: TCefView
//
//	A View representing a button. Depending on the specific type, the button
//	could be implemented by a native control or custom rendered. Methods must be
//	called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_capi.h">CEF source file: /include/capi/views/cef_button_capi.h (cef_button_t)</a>
type TCefButton struct {
	TCefView
}

// ButtonRef -> ICefButton
var ButtonRef button

// button TCefButton Ref
type button uintptr

// UnWrap
//
//	Returns a ICefButton instance using a PCefButton data pointer.
func (m *button) UnWrap(data uintptr) ICefButton {
	var resultCefButton uintptr
	CEF().SysCallN(709, uintptr(data), uintptr(unsafePointer(&resultCefButton)))
	return AsCefButton(resultCefButton)
}

func (m *TCefButton) AsLabelButton() ICefLabelButton {
	var resultCefLabelButton uintptr
	CEF().SysCallN(703, m.Instance(), uintptr(unsafePointer(&resultCefLabelButton)))
	return AsCefLabelButton(resultCefLabelButton)
}

func (m *TCefButton) GetState() TCefButtonState {
	r1 := CEF().SysCallN(704, m.Instance())
	return TCefButtonState(r1)
}

func (m *TCefButton) SetState(state TCefButtonState) {
	CEF().SysCallN(707, m.Instance(), uintptr(state))
}

func (m *TCefButton) SetInkDropEnabled(enabled bool) {
	CEF().SysCallN(706, m.Instance(), PascalBool(enabled))
}

func (m *TCefButton) SetTooltipText(tooltiptext string) {
	CEF().SysCallN(708, m.Instance(), PascalStr(tooltiptext))
}

func (m *TCefButton) SetAccessibleName(name string) {
	CEF().SysCallN(705, m.Instance(), PascalStr(name))
}
