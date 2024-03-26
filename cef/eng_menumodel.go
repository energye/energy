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

// ICefMenuModel Parent: ICefBaseRefCounted
//
//	Supports creation and modification of menus. See TCefMenuId for the command ids that have default implementations. All user-defined command ids should be between MENU_ID_USER_FIRST and MENU_ID_USER_LAST. The functions of this interface can only be accessed on the browser process the UI thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_capi.h">CEF source file: /include/capi/cef_menu_model_capi.h (cef_menu_model_t))
type ICefMenuModel interface {
	ICefBaseRefCounted
	// IsSubMenu
	//  Returns true (1) if this menu is a submenu.
	IsSubMenu() bool // function
	// Clear
	//  Clears the menu. Returns true (1) on success.
	Clear() bool // function
	// GetCount
	//  Returns the number of items in this menu.
	GetCount() NativeUInt // function
	// AddSeparator
	//  Add a separator to the menu. Returns true (1) on success.
	AddSeparator() bool // function
	// AddItem
	//  Add an item to the menu. Returns true (1) on success.
	AddItem(commandId int32, text string) bool // function
	// AddCheckItem
	//  Add a check item to the menu. Returns true (1) on success.
	AddCheckItem(commandId int32, text string) bool // function
	// AddRadioItem
	//  Add a radio item to the menu. Only a single item with the specified |group_id| can be checked at a time. Returns true (1) on success.
	AddRadioItem(commandId int32, text string, groupId int32) bool // function
	// AddSubMenu
	//  Add a sub-menu to the menu. The new sub-menu is returned.
	AddSubMenu(commandId int32, text string) ICefMenuModel // function
	// InsertSeparatorAt
	//  Insert a separator in the menu at the specified |index|. Returns true (1) on success.
	InsertSeparatorAt(index NativeUInt) bool // function
	// InsertItemAt
	//  Insert an item in the menu at the specified |index|. Returns true (1) on success.
	InsertItemAt(index NativeUInt, commandId int32, text string) bool // function
	// InsertCheckItemAt
	//  Insert a check item in the menu at the specified |index|. Returns true (1) on success.
	InsertCheckItemAt(index NativeUInt, commandId int32, text string) bool // function
	// InsertRadioItemAt
	//  Insert a radio item in the menu at the specified |index|. Only a single item with the specified |group_id| can be checked at a time. Returns true (1) on success.
	InsertRadioItemAt(index NativeUInt, commandId int32, text string, groupId int32) bool // function
	// InsertSubMenuAt
	//  Insert a sub-menu in the menu at the specified |index|. The new sub-menu is returned.
	InsertSubMenuAt(index NativeUInt, commandId int32, text string) ICefMenuModel // function
	// Remove
	//  Removes the item with the specified |command_id|. Returns true (1) on success.
	Remove(commandId int32) bool // function
	// RemoveAt
	//  Removes the item at the specified |index|. Returns true (1) on success.
	RemoveAt(index NativeUInt) bool // function
	// GetIndexOf
	//  Returns the index associated with the specified |command_id| or -1 if not found due to the command id not existing in the menu.
	GetIndexOf(commandId int32) int32 // function
	// GetCommandIdAt
	//  Returns the command id at the specified |index| or -1 if not found due to invalid range or the index being a separator.
	GetCommandIdAt(index NativeUInt) int32 // function
	// SetCommandIdAt
	//  Sets the command id at the specified |index|. Returns true (1) on success.
	SetCommandIdAt(index NativeUInt, commandId int32) bool // function
	// GetLabel
	//  Returns the label for the specified |command_id| or NULL if not found.
	GetLabel(commandId int32) string // function
	// GetLabelAt
	//  Returns the label at the specified |index| or NULL if not found due to invalid range or the index being a separator.
	GetLabelAt(index NativeUInt) string // function
	// SetLabel
	//  Sets the label for the specified |command_id|. Returns true (1) on success.
	SetLabel(commandId int32, text string) bool // function
	// SetLabelAt
	//  Set the label at the specified |index|. Returns true (1) on success.
	SetLabelAt(index NativeUInt, text string) bool // function
	// GetType
	//  Returns the item type for the specified |command_id|.
	GetType(commandId int32) TCefMenuItemType // function
	// GetTypeAt
	//  Returns the item type at the specified |index|.
	GetTypeAt(index NativeUInt) TCefMenuItemType // function
	// GetGroupId
	//  Returns the group id for the specified |command_id| or -1 if invalid.
	GetGroupId(commandId int32) int32 // function
	// GetGroupIdAt
	//  Returns the group id at the specified |index| or -1 if invalid.
	GetGroupIdAt(index NativeUInt) int32 // function
	// SetGroupId
	//  Sets the group id for the specified |command_id|. Returns true (1) on success.
	SetGroupId(commandId, groupId int32) bool // function
	// SetGroupIdAt
	//  Sets the group id at the specified |index|. Returns true (1) on success.
	SetGroupIdAt(index NativeUInt, groupId int32) bool // function
	// GetSubMenu
	//  Returns the submenu for the specified |command_id| or NULL if invalid.
	GetSubMenu(commandId int32) ICefMenuModel // function
	// GetSubMenuAt
	//  Returns the submenu at the specified |index| or NULL if invalid.
	GetSubMenuAt(index NativeUInt) ICefMenuModel // function
	// IsVisible
	//  Returns true (1) if the specified |command_id| is visible.
	IsVisible(commandId int32) bool // function
	// IsVisibleAt
	//  Returns true (1) if the specified |index| is visible.
	IsVisibleAt(index NativeUInt) bool // function
	// SetVisible
	//  Change the visibility of the specified |command_id|. Returns true (1) on success.
	SetVisible(commandId int32, visible bool) bool // function
	// SetVisibleAt
	//  Change the visibility at the specified |index|. Returns true (1) on success.
	SetVisibleAt(index NativeUInt, visible bool) bool // function
	// IsEnabled
	//  Returns true (1) if the specified |command_id| is enabled.
	IsEnabled(commandId int32) bool // function
	// IsEnabledAt
	//  Returns true (1) if the specified |index| is enabled.
	IsEnabledAt(index NativeUInt) bool // function
	// SetEnabled
	//  Change the enabled status of the specified |command_id|. Returns true (1) on success.
	SetEnabled(commandId int32, enabled bool) bool // function
	// SetEnabledAt
	//  Change the enabled status at the specified |index|. Returns true (1) on success.
	SetEnabledAt(index NativeUInt, enabled bool) bool // function
	// IsChecked
	//  Returns true (1) if the specified |command_id| is checked. Only applies to check and radio items.
	IsChecked(commandId int32) bool // function
	// IsCheckedAt
	//  Returns true (1) if the specified |index| is checked. Only applies to check and radio items.
	IsCheckedAt(index NativeUInt) bool // function
	// SetChecked
	//  Check the specified |command_id|. Only applies to check and radio items. Returns true (1) on success.
	SetChecked(commandId int32, checked bool) bool // function
	// SetCheckedAt
	//  Check the specified |index|. Only applies to check and radio items. Returns true (1) on success.
	SetCheckedAt(index NativeUInt, checked bool) bool // function
	// HasAccelerator
	//  Returns true (1) if the specified |command_id| has a keyboard accelerator assigned.
	HasAccelerator(commandId int32) bool // function
	// HasAcceleratorAt
	//  Returns true (1) if the specified |index| has a keyboard accelerator assigned.
	HasAcceleratorAt(index NativeUInt) bool // function
	// SetAccelerator
	//  Set the keyboard accelerator for the specified |command_id|. |key_code| can be any virtual key or character value. Returns true (1) on success.
	SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool // function
	// SetAcceleratorAt
	//  Set the keyboard accelerator at the specified |index|. |key_code| can be any virtual key or character value. Returns true (1) on success.
	SetAcceleratorAt(index NativeUInt, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool // function
	// RemoveAccelerator
	//  Remove the keyboard accelerator for the specified |command_id|. Returns true (1) on success.
	RemoveAccelerator(commandId int32) bool // function
	// RemoveAcceleratorAt
	//  Remove the keyboard accelerator at the specified |index|. Returns true (1) on success.
	RemoveAcceleratorAt(index NativeUInt) bool // function
	// GetAccelerator
	//  Retrieves the keyboard accelerator for the specified |command_id|. Returns true (1) on success.
	GetAccelerator(commandId int32, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool // function
	// GetAcceleratorAt
	//  Retrieves the keyboard accelerator for the specified |index|. Returns true (1) on success.
	GetAcceleratorAt(index NativeUInt, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool // function
	// SetColor
	//  Set the explicit color for |command_id| and |color_type| to |color|. Specify a |color| value of 0 to remove the explicit color. If no explicit color or default color is set for |color_type| then the system color will be used. Returns true (1) on success.
	SetColor(commandId int32, colorType TCefMenuColorType, color TCefColor) bool // function
	// SetColorAt
	//  Set the explicit color for |command_id| and |index| to |color|. Specify a |color| value of 0 to remove the explicit color. Specify an |index| value of -1 to set the default color for items that do not have an explicit color set. If no explicit color or default color is set for |color_type| then the system color will be used. Returns true (1) on success.
	SetColorAt(index int32, colorType TCefMenuColorType, color TCefColor) bool // function
	// GetColor
	//  Returns in |color| the color that was explicitly set for |command_id| and |color_type|. If a color was not set then 0 will be returned in |color|. Returns true (1) on success.
	GetColor(commandId int32, colorType TCefMenuColorType, outColor *TCefColor) bool // function
	// GetColorAt
	//  Returns in |color| the color that was explicitly set for |command_id| and |color_type|. Specify an |index| value of -1 to return the default color in |color|. If a color was not set then 0 will be returned in |color|. Returns true (1) on success.
	GetColorAt(index int32, colorType TCefMenuColorType, outColor *TCefColor) bool // function
	// SetFontList
	//  Sets the font list for the specified |command_id|. If |font_list| is NULL the system font will be used. Returns true (1) on success. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where: - FONT_FAMILY_LIST is a comma-separated list of font family names, - STYLES is an optional space- separated list of style names (case-sensitive "Bold" and "Italic" are supported), and - SIZE is an integer font size in pixels with the suffix "px". Here are examples of valid font description strings: - "Arial, Helvetica, Bold Italic 14px" - "Arial, 14px"
	SetFontList(commandId int32, fontList string) bool // function
	// SetFontListAt
	//  Sets the font list for the specified |index|. Specify an |index| value of -1 to set the default font. If |font_list| is NULL the system font will be used. Returns true (1) on success. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where: - FONT_FAMILY_LIST is a comma-separated list of font family names, - STYLES is an optional space- separated list of style names (case-sensitive "Bold" and "Italic" are supported), and - SIZE is an integer font size in pixels with the suffix "px". Here are examples of valid font description strings: - "Arial, Helvetica, Bold Italic 14px" - "Arial, 14px"
	SetFontListAt(index int32, fontList string) bool // function
}

// TCefMenuModel Parent: TCefBaseRefCounted
//
//	Supports creation and modification of menus. See TCefMenuId for the command ids that have default implementations. All user-defined command ids should be between MENU_ID_USER_FIRST and MENU_ID_USER_LAST. The functions of this interface can only be accessed on the browser process the UI thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_capi.h">CEF source file: /include/capi/cef_menu_model_capi.h (cef_menu_model_t))
type TCefMenuModel struct {
	TCefBaseRefCounted
}

// MenuModelRef -> ICefMenuModel
var MenuModelRef menuModel

// menuModel TCefMenuModel Ref
type menuModel uintptr

func (m *menuModel) UnWrap(data uintptr) ICefMenuModel {
	var resultCefMenuModel uintptr
	CEF().SysCallN(1140, uintptr(data), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *menuModel) New(delegate ICefMenuModelDelegate) ICefMenuModel {
	var resultCefMenuModel uintptr
	CEF().SysCallN(1118, GetObjectUintptr(delegate), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) IsSubMenu() bool {
	r1 := CEF().SysCallN(1115, m.Instance())
	return GoBool(r1)
}

func (m *TCefMenuModel) Clear() bool {
	r1 := CEF().SysCallN(1088, m.Instance())
	return GoBool(r1)
}

func (m *TCefMenuModel) GetCount() NativeUInt {
	r1 := CEF().SysCallN(1094, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefMenuModel) AddSeparator() bool {
	r1 := CEF().SysCallN(1086, m.Instance())
	return GoBool(r1)
}

func (m *TCefMenuModel) AddItem(commandId int32, text string) bool {
	r1 := CEF().SysCallN(1084, m.Instance(), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) AddCheckItem(commandId int32, text string) bool {
	r1 := CEF().SysCallN(1083, m.Instance(), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) AddRadioItem(commandId int32, text string, groupId int32) bool {
	r1 := CEF().SysCallN(1085, m.Instance(), uintptr(commandId), PascalStr(text), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) AddSubMenu(commandId int32, text string) ICefMenuModel {
	var resultCefMenuModel uintptr
	CEF().SysCallN(1087, m.Instance(), uintptr(commandId), PascalStr(text), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) InsertSeparatorAt(index NativeUInt) bool {
	r1 := CEF().SysCallN(1109, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertItemAt(index NativeUInt, commandId int32, text string) bool {
	r1 := CEF().SysCallN(1107, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertCheckItemAt(index NativeUInt, commandId int32, text string) bool {
	r1 := CEF().SysCallN(1106, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertRadioItemAt(index NativeUInt, commandId int32, text string, groupId int32) bool {
	r1 := CEF().SysCallN(1108, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) InsertSubMenuAt(index NativeUInt, commandId int32, text string) ICefMenuModel {
	var resultCefMenuModel uintptr
	CEF().SysCallN(1110, m.Instance(), uintptr(index), uintptr(commandId), PascalStr(text), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) Remove(commandId int32) bool {
	r1 := CEF().SysCallN(1119, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) RemoveAt(index NativeUInt) bool {
	r1 := CEF().SysCallN(1122, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetIndexOf(commandId int32) int32 {
	r1 := CEF().SysCallN(1097, m.Instance(), uintptr(commandId))
	return int32(r1)
}

func (m *TCefMenuModel) GetCommandIdAt(index NativeUInt) int32 {
	r1 := CEF().SysCallN(1093, m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *TCefMenuModel) SetCommandIdAt(index NativeUInt, commandId int32) bool {
	r1 := CEF().SysCallN(1129, m.Instance(), uintptr(index), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetLabel(commandId int32) string {
	r1 := CEF().SysCallN(1098, m.Instance(), uintptr(commandId))
	return GoStr(r1)
}

func (m *TCefMenuModel) GetLabelAt(index NativeUInt) string {
	r1 := CEF().SysCallN(1099, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TCefMenuModel) SetLabel(commandId int32, text string) bool {
	r1 := CEF().SysCallN(1136, m.Instance(), uintptr(commandId), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetLabelAt(index NativeUInt, text string) bool {
	r1 := CEF().SysCallN(1137, m.Instance(), uintptr(index), PascalStr(text))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetType(commandId int32) TCefMenuItemType {
	r1 := CEF().SysCallN(1102, m.Instance(), uintptr(commandId))
	return TCefMenuItemType(r1)
}

func (m *TCefMenuModel) GetTypeAt(index NativeUInt) TCefMenuItemType {
	r1 := CEF().SysCallN(1103, m.Instance(), uintptr(index))
	return TCefMenuItemType(r1)
}

func (m *TCefMenuModel) GetGroupId(commandId int32) int32 {
	r1 := CEF().SysCallN(1095, m.Instance(), uintptr(commandId))
	return int32(r1)
}

func (m *TCefMenuModel) GetGroupIdAt(index NativeUInt) int32 {
	r1 := CEF().SysCallN(1096, m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *TCefMenuModel) SetGroupId(commandId, groupId int32) bool {
	r1 := CEF().SysCallN(1134, m.Instance(), uintptr(commandId), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetGroupIdAt(index NativeUInt, groupId int32) bool {
	r1 := CEF().SysCallN(1135, m.Instance(), uintptr(index), uintptr(groupId))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetSubMenu(commandId int32) ICefMenuModel {
	var resultCefMenuModel uintptr
	CEF().SysCallN(1100, m.Instance(), uintptr(commandId), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) GetSubMenuAt(index NativeUInt) ICefMenuModel {
	var resultCefMenuModel uintptr
	CEF().SysCallN(1101, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefMenuModel)))
	return AsCefMenuModel(resultCefMenuModel)
}

func (m *TCefMenuModel) IsVisible(commandId int32) bool {
	r1 := CEF().SysCallN(1116, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsVisibleAt(index NativeUInt) bool {
	r1 := CEF().SysCallN(1117, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetVisible(commandId int32, visible bool) bool {
	r1 := CEF().SysCallN(1138, m.Instance(), uintptr(commandId), PascalBool(visible))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetVisibleAt(index NativeUInt, visible bool) bool {
	r1 := CEF().SysCallN(1139, m.Instance(), uintptr(index), PascalBool(visible))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsEnabled(commandId int32) bool {
	r1 := CEF().SysCallN(1113, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsEnabledAt(index NativeUInt) bool {
	r1 := CEF().SysCallN(1114, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetEnabled(commandId int32, enabled bool) bool {
	r1 := CEF().SysCallN(1130, m.Instance(), uintptr(commandId), PascalBool(enabled))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetEnabledAt(index NativeUInt, enabled bool) bool {
	r1 := CEF().SysCallN(1131, m.Instance(), uintptr(index), PascalBool(enabled))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsChecked(commandId int32) bool {
	r1 := CEF().SysCallN(1111, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) IsCheckedAt(index NativeUInt) bool {
	r1 := CEF().SysCallN(1112, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetChecked(commandId int32, checked bool) bool {
	r1 := CEF().SysCallN(1125, m.Instance(), uintptr(commandId), PascalBool(checked))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetCheckedAt(index NativeUInt, checked bool) bool {
	r1 := CEF().SysCallN(1126, m.Instance(), uintptr(index), PascalBool(checked))
	return GoBool(r1)
}

func (m *TCefMenuModel) HasAccelerator(commandId int32) bool {
	r1 := CEF().SysCallN(1104, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) HasAcceleratorAt(index NativeUInt) bool {
	r1 := CEF().SysCallN(1105, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	r1 := CEF().SysCallN(1123, m.Instance(), uintptr(commandId), uintptr(keyCode), PascalBool(shiftPressed), PascalBool(ctrlPressed), PascalBool(altPressed))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetAcceleratorAt(index NativeUInt, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	r1 := CEF().SysCallN(1124, m.Instance(), uintptr(index), uintptr(keyCode), PascalBool(shiftPressed), PascalBool(ctrlPressed), PascalBool(altPressed))
	return GoBool(r1)
}

func (m *TCefMenuModel) RemoveAccelerator(commandId int32) bool {
	r1 := CEF().SysCallN(1120, m.Instance(), uintptr(commandId))
	return GoBool(r1)
}

func (m *TCefMenuModel) RemoveAcceleratorAt(index NativeUInt) bool {
	r1 := CEF().SysCallN(1121, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetAccelerator(commandId int32, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := CEF().SysCallN(1089, m.Instance(), uintptr(commandId), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*outKeyCode = int32(result1)
	*outShiftPressed = GoBool(result2)
	*outCtrlPressed = GoBool(result3)
	*outLtPressed = GoBool(result4)
	return GoBool(r1)
}

func (m *TCefMenuModel) GetAcceleratorAt(index NativeUInt, outKeyCode *int32, outShiftPressed, outCtrlPressed, outLtPressed *bool) bool {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := CEF().SysCallN(1090, m.Instance(), uintptr(index), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*outKeyCode = int32(result1)
	*outShiftPressed = GoBool(result2)
	*outCtrlPressed = GoBool(result3)
	*outLtPressed = GoBool(result4)
	return GoBool(r1)
}

func (m *TCefMenuModel) SetColor(commandId int32, colorType TCefMenuColorType, color TCefColor) bool {
	r1 := CEF().SysCallN(1127, m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(color))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetColorAt(index int32, colorType TCefMenuColorType, color TCefColor) bool {
	r1 := CEF().SysCallN(1128, m.Instance(), uintptr(index), uintptr(colorType), uintptr(color))
	return GoBool(r1)
}

func (m *TCefMenuModel) GetColor(commandId int32, colorType TCefMenuColorType, outColor *TCefColor) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(1091, m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(unsafePointer(&result2)))
	*outColor = TCefColor(result2)
	return GoBool(r1)
}

func (m *TCefMenuModel) GetColorAt(index int32, colorType TCefMenuColorType, outColor *TCefColor) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(1092, m.Instance(), uintptr(index), uintptr(colorType), uintptr(unsafePointer(&result2)))
	*outColor = TCefColor(result2)
	return GoBool(r1)
}

func (m *TCefMenuModel) SetFontList(commandId int32, fontList string) bool {
	r1 := CEF().SysCallN(1132, m.Instance(), uintptr(commandId), PascalStr(fontList))
	return GoBool(r1)
}

func (m *TCefMenuModel) SetFontListAt(index int32, fontList string) bool {
	r1 := CEF().SysCallN(1133, m.Instance(), uintptr(index), PascalStr(fontList))
	return GoBool(r1)
}
