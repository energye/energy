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

// ICEFMenuButtonComponent Parent: ICEFLabelButtonComponent
type ICEFMenuButtonComponent interface {
	ICEFLabelButtonComponent
	// CreateMenuButton
	//  Create a new MenuButton.
	CreateMenuButton(aText string) // procedure
	// ShowMenu
	//  Show a menu with contents |menu_model|. |screen_point| specifies the menu
	//  position in screen coordinates. |anchor_position| specifies how the menu
	//  will be anchored relative to |screen_point|. This function should be
	//  called from ICefMenuButtonDelegate.OnMenuButtonPressed().
	ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) // procedure
	// TriggerMenu
	//  Show the menu for this button. Results in a call to
	//  ICefMenuButtonDelegate.OnMenuButtonPressed().
	TriggerMenu() // procedure
	// SetOnMenuButtonPressed
	//  Called when |button| is pressed. Call ICefMenuButton.ShowMenu() to
	//  show a popup menu at |screen_point|. When showing a custom popup such as a
	//  window keep a reference to |button_pressed_lock| until the popup is hidden
	//  to maintain the pressed button state.
	SetOnMenuButtonPressed(fn TOnMenuButtonPressed) // property event
}

// TCEFMenuButtonComponent Parent: TCEFLabelButtonComponent
type TCEFMenuButtonComponent struct {
	TCEFLabelButtonComponent
	menuButtonPressedPtr uintptr
}

func NewCEFMenuButtonComponent(aOwner IComponent) ICEFMenuButtonComponent {
	r1 := CEF().SysCallN(161, GetObjectUintptr(aOwner))
	return AsCEFMenuButtonComponent(r1)
}

func CEFMenuButtonComponentClass() TClass {
	ret := CEF().SysCallN(160)
	return TClass(ret)
}

func (m *TCEFMenuButtonComponent) CreateMenuButton(aText string) {
	CEF().SysCallN(162, m.Instance(), PascalStr(aText))
}

func (m *TCEFMenuButtonComponent) ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) {
	CEF().SysCallN(164, m.Instance(), GetObjectUintptr(menumodel), uintptr(unsafePointer(screenpoint)), uintptr(anchorposition))
}

func (m *TCEFMenuButtonComponent) TriggerMenu() {
	CEF().SysCallN(165, m.Instance())
}

func (m *TCEFMenuButtonComponent) SetOnMenuButtonPressed(fn TOnMenuButtonPressed) {
	if m.menuButtonPressedPtr != 0 {
		RemoveEventElement(m.menuButtonPressedPtr)
	}
	m.menuButtonPressedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(163, m.Instance(), m.menuButtonPressedPtr)
}
