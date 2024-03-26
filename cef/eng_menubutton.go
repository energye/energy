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

// ICefMenuButton Parent: ICefLabelButton
//
//	MenuButton is a button with optional text, icon and/or menu marker that
//	shows a menu when clicked with the left mouse button. All size and position
//	values are in density independent pixels (DIP) unless otherwise indicated.
//	Methods must be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_capi.h">CEF source file: /include/capi/views/cef_menu_button_capi.h (cef_menu_button_t)</a>
type ICefMenuButton interface {
	ICefLabelButton
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
}

// TCefMenuButton Parent: TCefLabelButton
//
//	MenuButton is a button with optional text, icon and/or menu marker that
//	shows a menu when clicked with the left mouse button. All size and position
//	values are in density independent pixels (DIP) unless otherwise indicated.
//	Methods must be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_capi.h">CEF source file: /include/capi/views/cef_menu_button_capi.h (cef_menu_button_t)</a>
type TCefMenuButton struct {
	TCefLabelButton
}

// MenuButtonRef -> ICefMenuButton
var MenuButtonRef menuButton

// menuButton TCefMenuButton Ref
type menuButton uintptr

// UnWrap
//
//	Returns a ICefMenuButton instance using a PCefMenuButton data pointer.
func (m *menuButton) UnWrap(data uintptr) ICefMenuButton {
	var resultCefMenuButton uintptr
	CEF().SysCallN(1082, uintptr(data), uintptr(unsafePointer(&resultCefMenuButton)))
	return AsCefMenuButton(resultCefMenuButton)
}

// CreateMenuButton
//
//	Create a new MenuButton. A |delegate| must be provided to call show_menu()
//	when the button is clicked. |text| will be shown on the MenuButton and used
//	as the default accessible name. If |with_frame| is true(1) the button will
//	have a visible frame at all times, center alignment, additional padding and
//	a default minimum size of 70x33 DIP. If |with_frame| is false(0) the button
//	will only have a visible frame on hover/press, left alignment, less padding
//	and no default minimum size.
func (m *menuButton) CreateMenuButton(delegate ICefMenuButtonDelegate, text string) ICefMenuButton {
	var resultCefMenuButton uintptr
	CEF().SysCallN(1079, GetObjectUintptr(delegate), PascalStr(text), uintptr(unsafePointer(&resultCefMenuButton)))
	return AsCefMenuButton(resultCefMenuButton)
}

func (m *TCefMenuButton) ShowMenu(menumodel ICefMenuModel, screenpoint *TCefPoint, anchorposition TCefMenuAnchorPosition) {
	CEF().SysCallN(1080, m.Instance(), GetObjectUintptr(menumodel), uintptr(unsafePointer(screenpoint)), uintptr(anchorposition))
}

func (m *TCefMenuButton) TriggerMenu() {
	CEF().SysCallN(1081, m.Instance())
}
