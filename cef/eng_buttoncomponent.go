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

// ICEFButtonComponent Parent: ICEFViewComponent
type ICEFButtonComponent interface {
	ICEFViewComponent
	// AsLabelButton
	//  Returns this Button as a LabelButton or NULL if this is not a LabelButton.
	AsLabelButton() ICefLabelButton // property
	// State
	//  Returns the current display state of the Button.
	State() TCefButtonState // property
	// SetState Set State
	SetState(AValue TCefButtonState) // property
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
	// SetOnButtonPressed
	//  Called when |button| is pressed.
	SetOnButtonPressed(fn TOnButtonPressed) // property event
	// SetOnButtonStateChanged
	//  Called when the state of |button| changes.
	SetOnButtonStateChanged(fn TOnButtonStateChanged) // property event
}

// TCEFButtonComponent Parent: TCEFViewComponent
type TCEFButtonComponent struct {
	TCEFViewComponent
	buttonPressedPtr      uintptr
	buttonStateChangedPtr uintptr
}

func NewCEFButtonComponent(aOwner IComponent) ICEFButtonComponent {
	r1 := CEF().SysCallN(102, GetObjectUintptr(aOwner))
	return AsCEFButtonComponent(r1)
}

func (m *TCEFButtonComponent) AsLabelButton() ICefLabelButton {
	var resultCefLabelButton uintptr
	CEF().SysCallN(100, m.Instance(), uintptr(unsafePointer(&resultCefLabelButton)))
	return AsCefLabelButton(resultCefLabelButton)
}

func (m *TCEFButtonComponent) State() TCefButtonState {
	r1 := CEF().SysCallN(108, 0, m.Instance(), 0)
	return TCefButtonState(r1)
}

func (m *TCEFButtonComponent) SetState(AValue TCefButtonState) {
	CEF().SysCallN(108, 1, m.Instance(), uintptr(AValue))
}

func CEFButtonComponentClass() TClass {
	ret := CEF().SysCallN(101)
	return TClass(ret)
}

func (m *TCEFButtonComponent) SetInkDropEnabled(enabled bool) {
	CEF().SysCallN(104, m.Instance(), PascalBool(enabled))
}

func (m *TCEFButtonComponent) SetTooltipText(tooltiptext string) {
	CEF().SysCallN(107, m.Instance(), PascalStr(tooltiptext))
}

func (m *TCEFButtonComponent) SetAccessibleName(name string) {
	CEF().SysCallN(103, m.Instance(), PascalStr(name))
}

func (m *TCEFButtonComponent) SetOnButtonPressed(fn TOnButtonPressed) {
	if m.buttonPressedPtr != 0 {
		RemoveEventElement(m.buttonPressedPtr)
	}
	m.buttonPressedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(105, m.Instance(), m.buttonPressedPtr)
}

func (m *TCEFButtonComponent) SetOnButtonStateChanged(fn TOnButtonStateChanged) {
	if m.buttonStateChangedPtr != 0 {
		RemoveEventElement(m.buttonStateChangedPtr)
	}
	m.buttonStateChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(106, m.Instance(), m.buttonStateChangedPtr)
}
