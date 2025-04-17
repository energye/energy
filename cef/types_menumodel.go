//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	. "github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	"strings"
	"unsafe"
)

// KeyAccelerator 自定义实现快捷键
var KeyAccelerator = &keyEventAccelerator{
	menuId:            MENU_ID_USER_FIRST,                  //menuId的启始位置
	commandItems:      make(map[MenuId]*MenuItem),          //右键菜单命令ID快捷键
	acceleratorItems:  make(map[string]*MenuItem),          //右键菜单命令名称快捷键
	acceleratorCustom: make(map[string]*AcceleratorCustom), //浏览器界面自定义快捷键
}

type FuncCallback func(browser *ICefBrowser, commandId MenuId, params *ICefContextMenuParams, menuType TCefContextMenuType, eventFlags uint32, result *bool)

type keyEventAccelerator struct {
	shift, ctrl, alt  bool
	keyCode           int32
	menuId            MenuId
	commandItems      map[MenuId]*MenuItem
	acceleratorItems  map[string]*MenuItem
	acceleratorCustom map[string]*AcceleratorCustom
}

// MenuItem 菜单项
type MenuItem struct {
	CommandId   MenuId // >= 26500 and <= 28500
	Accelerator string // 快捷键 shift ctrl alt【按键VK..】(shift+ctrl+alt+n)
	Text        string //显示文本
	Label       string
	GroupId     int32 //分组 配合 MenuType 使用
	MenuType    TCefContextMenuType
	Callback    FuncCallback //点击 或 快捷键触发的回调
}

// AddAcceleratorCustom 添加自定义快捷键
func (m *keyEventAccelerator) AddAcceleratorCustom(accelerator *AcceleratorCustom) {
	if accelerator == nil {
		return
	}
	accelerator.Accelerator = strings.Replace(strings.ToUpper(accelerator.Accelerator), " ", "", -1)
	as := strings.Split(accelerator.Accelerator, "+")
	if len(as) > 0 && len(as) <= 4 {
		var shift = ArrayIndexOf(as, MA_Shift) != -1
		var ctrl = ArrayIndexOf(as, MA_Ctrl) != -1
		var alt = ArrayIndexOf(as, MA_Alt) != -1
		var keyCode = rune(strings.ToUpper(as[len(as)-1])[0])
		accelerator.Accelerator = acceleratorCode(shift, ctrl, alt, keyCode)
		m.acceleratorCustom[accelerator.Accelerator] = accelerator
	}
}

func (m *keyEventAccelerator) acceleratorCustomCallback(accelerator string, browse *ICefBrowser, event *TCefKeyEvent, result *bool) bool {
	if item, ok := m.acceleratorCustom[accelerator]; ok {
		if item.Callback != nil {
			item.Callback(browse, event, result)
			return true
		}
	}
	return false
}

// NextCommandId 生成并返回下一个快捷键ID
func (m *keyEventAccelerator) NextCommandId() MenuId {
	m.menuId++
	return m.menuId
}

func (m *keyEventAccelerator) commandIdEventCallback(browse *ICefBrowser, commandId MenuId, params *ICefContextMenuParams, eventFlags uint32, result *bool) bool {
	if item, ok := m.commandItems[commandId]; ok {
		if item.Callback != nil {
			item.Callback(browse, commandId, params, item.MenuType, eventFlags, result)
			return true
		}
	}
	return false
}

func (m *keyEventAccelerator) accelerator(browse *ICefBrowser, event *TCefKeyEvent, result *bool) bool {
	if event.KeyDown() {
		if event.WindowsKeyCode == MA_Shift_Code {
			m.shift = true
		} else if event.WindowsKeyCode == MA_Ctrl_Code {
			m.ctrl = true
		} else if event.WindowsKeyCode == MA_Alt_Code {
			m.alt = true
		} else {
			var accelerator = acceleratorCode(m.shift, m.ctrl, m.alt, rune(event.WindowsKeyCode))
			//m.shift = false
			//m.ctrl = false
			//m.alt = false
			m.keyCode = -1
			if m.acceleratorEventCallback(browse, accelerator, result) {
				return true
			} else if m.acceleratorCustomCallback(accelerator, browse, event, result) {
				return true
			}
		}
	} else {
		m.shift = false
		m.ctrl = false
		m.alt = false
		m.keyCode = -1
	}
	return false
}

func (m *keyEventAccelerator) acceleratorEventCallback(browse *ICefBrowser, accelerator string, result *bool) bool {
	if item, ok := m.acceleratorItems[accelerator]; ok {
		if item.Callback != nil {
			item.Callback(browse, item.CommandId, nil, item.MenuType, 0, result)
			return true
		}
	}
	return false
}

func (m *keyEventAccelerator) clear() {
	m.menuId = MENU_ID_USER_FIRST
	m.commandItems = make(map[MenuId]*MenuItem)
	m.acceleratorItems = make(map[string]*MenuItem)
}

// AddMenuItem 添加一个菜单项 MenuItem
func (m *ICefMenuModel) AddMenuItem(item *MenuItem) bool {
	if item == nil {
		return false
	}
	if item.CommandId == 0 {
		item.CommandId = KeyAccelerator.NextCommandId()
	}
	//if item.CommandId >= MENU_ID_USER_FIRST && item.CommandId <= MENU_ID_USER_LAST {
	if item.MenuType == CMT_NONE {
		m.AddItem(item.CommandId, item.Text)
	} else if item.MenuType == CMT_CHECK {
		m.AddCheckItem(item.CommandId, item.Text)
	} else if item.MenuType == CMT_RADIO {
		m.AddRadioItem(item.CommandId, item.Text, item.GroupId)
	} else {
		return false
	}
	if item.Label != "" {
		m.SetLabel(item.CommandId, item.Label)
	}
	if item.Accelerator != "" {
		item.Accelerator = strings.Replace(strings.ToUpper(item.Accelerator), " ", "", -1)
		as := strings.Split(item.Accelerator, "+")
		if len(as) > 0 && len(as) <= 4 {
			var shift = ArrayIndexOf(as, MA_Shift) != -1
			var ctrl = ArrayIndexOf(as, MA_Ctrl) != -1
			var alt = ArrayIndexOf(as, MA_Alt) != -1
			var keyCode = rune(strings.ToUpper(as[len(as)-1])[0])
			item.Accelerator = acceleratorCode(shift, ctrl, alt, keyCode)
			m.SetAccelerator(item.CommandId, keyCode, shift, ctrl, alt)
			KeyAccelerator.acceleratorItems[item.Accelerator] = item
		}
	}
	KeyAccelerator.commandItems[item.CommandId] = item
	return true
}

// ICefMenuModel
// Supports creation and modification of menus. See TCefMenuId (cef_menu_id_t) for the
// command ids that have default implementations. All user-defined command ids
// should be between MENU_ID_USER_FIRST and MENU_ID_USER_LAST. The functions of
// this structure can only be accessed on the browser process the UI thread.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_capi.h">CEF source file: /include/capi/cef_menu_model_capi.h (cef_menu_model_t)</see></para>
type ICefMenuModel struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	CefMis   *keyEventAccelerator
}

// MenuModelRef -> ICefMenuModel
var MenuModelRef menuModel

type menuModel uintptr

func (*menuModel) New(delegate *ICefMenuModelDelegate) *ICefMenuModel {
	var result uintptr
	imports.Proc(def.CefMenuModelRef_New).Call(delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuModel{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefMenuModel) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMenuModel) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMenuModel) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// AddSeparator Add a separator to the menu. Returns true (1) on success.
func (m *ICefMenuModel) AddSeparator() bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_AddSeparator).Call(m.Instance())
	return api.GoBool(r1)
}

// Clear Clears the menu. Returns true (1) on success.
func (m *ICefMenuModel) Clear() bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_Clear).Call(m.Instance())
	return api.GoBool(r1)
}

// IsSubMenu
// Returns true (1) if this menu is a submenu.
func (m *ICefMenuModel) IsSubMenu() bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_IsSubMenu).Call(m.Instance())
	return api.GoBool(r1)
}

// GetCount Returns the number of items in this menu.
func (m *ICefMenuModel) GetCount() uint32 {
	r1, _, _ := imports.Proc(def.CEFMenuModel_GetCount).Call(m.Instance())
	return uint32(r1)
}

// AddItem Add an item to the menu. Returns true (1) on success.
func (m *ICefMenuModel) AddItem(commandId MenuId, text string) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_AddItem).Call(m.Instance(), uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// AddCheckItem Add a check item to the menu. Returns true (1) on success.
func (m *ICefMenuModel) AddCheckItem(commandId MenuId, text string) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_AddCheckItem).Call(m.Instance(), uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// AddRadioItem Add a radio item to the menu. Only a single item with the specified
//
//	|group_id| can be checked at a time. Returns true (1) on success.
func (m *ICefMenuModel) AddRadioItem(commandId MenuId, text string, groupId int32) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_AddRadioItem).Call(m.Instance(), uintptr(commandId), api.PascalStr(text), uintptr(groupId))
	return api.GoBool(r1)
}

// AddSubMenu Add a sub-menu to the menu. The new sub-menu is returned.
func (m *ICefMenuModel) AddSubMenu(commandId MenuId, text string) *ICefMenuModel {
	var result uintptr
	imports.Proc(def.CEFMenuModel_AddSubMenu).Call(m.Instance(), uintptr(commandId), api.PascalStr(text), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuModel{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Remove
//
//	Removes the item with the specified |command_id|. Returns true (1) on
//	success.
func (m *ICefMenuModel) Remove(commandId MenuId) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_Remove).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

// RemoveAt Removes the item at the specified |index|. Returns true (1) on success.
func (m *ICefMenuModel) RemoveAt(index int32) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_RemoveAt).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// SetChecked
//
//	Check the specified |command_id|. Only applies to check and radio items.
//	Returns true (1) on success.
func (m *ICefMenuModel) SetChecked(commandId MenuId, check bool) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_SetChecked).Call(m.Instance(), uintptr(commandId), api.PascalBool(check))
	return api.GoBool(r1)
}

// IsChecked
//
//	Returns in |color| the color that was explicitly set for |command_id| and
//	|color_type|. If a color was not set then 0 will be returned in |color|.
//	Returns true (1) on success.
func (m *ICefMenuModel) IsChecked(commandId MenuId) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_IsChecked).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

// SetColor 设置可用菜单项 cef.NewCefARGB(a, r, g, b), 颜色根据 consts.TCefMenuColorType
func (m *ICefMenuModel) SetColor(commandId MenuId, colorType TCefMenuColorType, color types.TCefColor) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_SetColor).Call(m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(color))
	return api.GoBool(r1)
}

// SetFontList
//
//	Sets the font list for the specified |command_id|. If |font_list| is NULL
//	the system font will be used. Returns true (1) on success. The format is
//	"<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where:
//	- FONT_FAMILY_LIST is a comma-separated list of font family names,
//	- STYLES is an optional space-separated list of style names
//	  (case-sensitive "Bold" and "Italic" are supported), and
//	- SIZE is an integer font size in pixels with the suffix "px".
//
//	Here are examples of valid font description strings:
//	- "Arial, Helvetica, Bold Italic 14px"
//	- "Arial, 14px"
func (m *ICefMenuModel) SetFontList(commandId MenuId, fontList string) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_SetFontList).Call(m.Instance(), uintptr(commandId), api.PascalStr(fontList))
	return api.GoBool(r1)
}

// HasAccelerator
//
//	Returns true (1) if the specified |command_id| has a keyboard accelerator
//	assigned.
func (m *ICefMenuModel) HasAccelerator(commandId MenuId) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_HasAccelerator).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

// SetAccelerator
//
//	Set the keyboard accelerator for the specified |command_id|. |key_code|
//	can be any virtual key or character value. Returns true (1) on success.
func (m *ICefMenuModel) SetAccelerator(commandId MenuId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_SetAccelerator).Call(m.Instance(), uintptr(commandId), uintptr(keyCode),
		api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
	return api.GoBool(r1)
}

// RemoveAccelerator 删除快捷键
//
//	Remove the keyboard accelerator for the specified |command_id|. Returns
//	true (1) on success.
func (m *ICefMenuModel) RemoveAccelerator(commandId MenuId) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_RemoveAccelerator).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

// IsVisible
//
//	Change the visibility of the specified |command_id|. Returns true (1) on
//	success.
func (m *ICefMenuModel) IsVisible(commandId MenuId) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_IsVisible).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

// SetVisible
//
//	Sets whether this overlay is visible. Overlays are hidden by default. If
//	this overlay is hidden then it and any child Views will not be drawn and,
//	if any of those Views currently have focus, then focus will also be
//	cleared. Painting is scheduled as needed.
func (m *ICefMenuModel) SetVisible(commandId MenuId, visible bool) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_SetVisible).Call(m.Instance(), uintptr(commandId), api.PascalBool(visible))
	return api.GoBool(r1)
}

// IsEnabled
//
//	Returns true (1) if the specified |command_id| is enabled.
func (m *ICefMenuModel) IsEnabled(commandId MenuId) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_IsEnabled).Call(m.Instance(), uintptr(commandId))
	return api.GoBool(r1)
}

// SetEnabled
//
//	Change the enabled status of the specified |command_id|. Returns true (1)
//	on success.
func (m *ICefMenuModel) SetEnabled(commandId MenuId, enabled bool) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_SetEnabled).Call(m.Instance(), uintptr(commandId), api.PascalBool(enabled))
	return api.GoBool(r1)
}

// SetLabel
//
//	Sets the label for the specified |command_id|. Returns true (1) on success.
func (m *ICefMenuModel) SetLabel(commandId MenuId, text string) bool {
	r1, _, _ := imports.Proc(def.CEFMenuModel_SetLabel).Call(m.Instance(), uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// GetIndexOf
//
//	 Returns the index associated with the specified |command_id| or -1 if not
//		found due to the command id not existing in the menu.
func (m *ICefMenuModel) GetIndexOf(commandId MenuId) int32 {
	r1, _, _ := imports.Proc(def.CEFMenuModel_GetIndexOf).Call(m.Instance(), uintptr(commandId))
	return int32(r1)
}

// GetType
//
//	Returns the item type for the specified |command_id|.
func (m *ICefMenuModel) GetType(commandId MenuId) TCefMenuItemType {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetType).Call(m.Instance(), uintptr(commandId))
	return TCefMenuItemType(r1)
}

// GetLabel
//
//	Sets the label for the specified |command_id|. Returns true (1) on success.
func (m *ICefMenuModel) GetLabel(commandId MenuId) string {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetLabel).Call(m.Instance(), uintptr(commandId))
	return api.GoStr(r1)
}

// GetGroupId
//
//	Returns the group id for the specified |command_id| or -1 if invalid.
func (m *ICefMenuModel) GetGroupId(commandId MenuId) int32 {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetGroupId).Call(m.Instance(), uintptr(commandId))
	return int32(r1)
}

// SetGroupId
//
//	Sets the group id for the specified |command_id|. Returns true (1) on success.
func (m *ICefMenuModel) SetGroupId(commandId MenuId, groupId int32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetGroupId).Call(m.Instance(), uintptr(commandId), uintptr(groupId))
	return api.GoBool(r1)
}

// GetSubMenu
//
//	Returns the submenu for the specified |command_id| or NULL if invalid.
func (m *ICefMenuModel) GetSubMenu(commandId MenuId) *ICefMenuModel {
	var result uintptr
	imports.Proc(def.CefMenuModel_GetSubMenu).Call(m.Instance(), uintptr(commandId), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuModel{instance: unsafe.Pointer(result)}
	}
	return nil
}

// GetColor
//
//	Returns in |color| the color that was explicitly set for |command_id| and
//	|color_type|. If a color was not set then 0 will be returned in |color|.
//	Returns true (1) on success.
func (m *ICefMenuModel) GetColor(commandId MenuId, colorType TCefMenuColorType) (color types.TCefColor, result bool) {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetColor).Call(m.Instance(), uintptr(commandId), uintptr(colorType), uintptr(unsafe.Pointer(&color)))
	return color, api.GoBool(r1)
}

// InsertSeparatorAt
//
//	Insert a separator in the menu at the specified |index|. Returns true (1) on success.
func (m *ICefMenuModel) InsertSeparatorAt(index uint32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_InsertSeparatorAt).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// InsertItemAt
// Insert an item in the menu at the specified |index|. Returns true (1) on
// success.
func (m *ICefMenuModel) InsertItemAt(index uint32, commandId int32, text string) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_InsertItemAt).Call(m.Instance(), uintptr(index), uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// InsertCheckItemAt
// Insert a check item in the menu at the specified |index|. Returns true (1)
// on success.
func (m *ICefMenuModel) InsertCheckItemAt(index uint32, commandId int32, text string) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_InsertCheckItemAt).Call(m.Instance(), uintptr(index), uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// InsertRadioItemAt
// Insert a radio item in the menu at the specified |index|. Only a single
// item with the specified |group_id| can be checked at a time. Returns true
// (1) on success.
func (m *ICefMenuModel) InsertRadioItemAt(index uint32, commandId int32, text string, groupId int32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_InsertRadioItemAt).Call(m.Instance(), uintptr(index), uintptr(commandId),
		api.PascalStr(text), uintptr(groupId))
	return api.GoBool(r1)
}

// InsertSubMenuAt
// Insert a sub-menu in the menu at the specified |index|. The new sub-menu
// is returned.
func (m *ICefMenuModel) InsertSubMenuAt(index uint32, commandId int32, text string) *ICefMenuModel {
	var result uintptr
	imports.Proc(def.CefMenuModel_InsertSubMenuAt).Call(m.Instance(), uintptr(index), uintptr(commandId), api.PascalStr(text),
		uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuModel{instance: unsafe.Pointer(result)}
	}
	return nil
}

// GetCommandIdAt Returns the command id at the specified |index| or -1 if not found due to
// invalid range or the index being a separator.
func (m *ICefMenuModel) GetCommandIdAt(index uint32) int32 {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetCommandIdAt).Call(m.Instance(), uintptr(index))
	return int32(r1)
}

// SetCommandIdAt
// Sets the command id at the specified |index|. Returns true (1) on success.
func (m *ICefMenuModel) SetCommandIdAt(index uint32, commandId int32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetCommandIdAt).Call(m.Instance(), uintptr(index), uintptr(commandId))
	return api.GoBool(r1)
}

// GetLabelAt
//
//	Returns the label at the specified |index| or NULL if not found due to
//	invalid range or the index being a separator.
func (m *ICefMenuModel) GetLabelAt(index uint32) string {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetLabelAt).Call(m.Instance(), uintptr(index))
	return api.GoStr(r1)
}

// SetLabelAt
//
//	Set the label at the specified |index|. Returns true (1) on success.
func (m *ICefMenuModel) SetLabelAt(index uint32, text string) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetLabelAt).Call(m.Instance(), uintptr(index), api.PascalStr(text))
	return api.GoBool(r1)
}

// GetTypeAt
//
//	Returns the item type at the specified |index|.
func (m *ICefMenuModel) GetTypeAt(index uint32) TCefMenuItemType {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetTypeAt).Call(m.Instance(), uintptr(index))
	return TCefMenuItemType(r1)
}

// GetGroupIdAt
//
//	Returns the group id at the specified |index| or -1 if invalid.
func (m *ICefMenuModel) GetGroupIdAt(index uint32) int32 {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetGroupIdAt).Call(m.Instance(), uintptr(index))
	return int32(r1)
}

// SetGroupIdAt
//
//	Sets the group id at the specified |index|. Returns true (1) on success.
func (m *ICefMenuModel) SetGroupIdAt(index uint32, groupId int32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetGroupIdAt).Call(m.Instance(), uintptr(index), uintptr(groupId))
	return api.GoBool(r1)
}

// GetSubMenuAt
//
//	Returns the submenu at the specified |index| or NULL if invalid.
func (m *ICefMenuModel) GetSubMenuAt(index uint32) *ICefMenuModel {
	var result uintptr
	imports.Proc(def.CefMenuModel_GetSubMenuAt).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuModel{instance: unsafe.Pointer(result)}
	}
	return nil
}

// GetColorAt
//
//	Returns in |color| the color that was explicitly set for |command_id| and
//	|color_type|. Specify an |index| value of -1 to return the default color
//	in |color|. If a color was not set then 0 will be returned in |color|.
//	Returns true (1) on success.
func (m *ICefMenuModel) GetColorAt(index uint32, colorType TCefMenuColorType) (color types.TCefColor, ok bool) {
	r1, _, _ := imports.Proc(def.CefMenuModel_GetColorAt).Call(m.Instance(), uintptr(index), uintptr(colorType), uintptr(unsafe.Pointer(&color)))
	ok = api.GoBool(r1)
	return
}

// SetFontListAt
//
//	Sets the font list for the specified |index|. Specify an |index| value of
//	-1 to set the default font. If |font_list| is NULL the system font will be
//	used. Returns true (1) on success. The format is
//	"<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where:
//	- FONT_FAMILY_LIST is a comma-separated list of font family names,
//	- STYLES is an optional space-separated list of style names
//	  (case-sensitive "Bold" and "Italic" are supported), and
//	- SIZE is an integer font size in pixels with the suffix "px".
//
//	Here are examples of valid font description strings:
//	- "Arial, Helvetica, Bold Italic 14px"
//	- "Arial, 14px"
func (m *ICefMenuModel) SetFontListAt(index int32, fontList string) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetFontListAt).Call(m.Instance(), uintptr(index), api.PascalStr(fontList))
	return api.GoBool(r1)
}

// RemoveAcceleratorAt
//
//	Remove the keyboard accelerator at the specified |index|. Returns true (1) on success.
func (m *ICefMenuModel) RemoveAcceleratorAt(index uint32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_RemoveAcceleratorAt).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// SetAcceleratorAt
//
//	Set the keyboard accelerator at the specified |index|. |key_code| can be
//	any virtual key or character value. Returns true (1) on success.
func (m *ICefMenuModel) SetAcceleratorAt(index uint32, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetAcceleratorAt).Call(m.Instance(), uintptr(index), uintptr(keyCode),
		api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
	return api.GoBool(r1)
}

// SetColorAt
//
//	Set the explicit color for |command_id| and |index| to |color|. Specify a
//	|color| value of 0 to remove the explicit color. Specify an |index| value
//	of -1 to set the default color for items that do not have an explicit
//	color set. If no explicit color or default color is set for |color_type|
//	then the system color will be used. Returns true (1) on success.
func (m *ICefMenuModel) SetColorAt(index int32, colorType TCefMenuColorType, color types.TCefColor) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetColorAt).Call(m.Instance(), uintptr(index), uintptr(colorType), uintptr(color))
	return api.GoBool(r1)
}

// IsEnabledAt
//
//	Returns true (1) if the specified |index| is enabled.
func (m *ICefMenuModel) IsEnabledAt(index uint32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_IsEnabledAt).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// SetVisibleAt
//
//	Change the visibility at the specified |index|. Returns true (1) on
//	success.
func (m *ICefMenuModel) SetVisibleAt(index uint32, visible bool) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetVisibleAt).Call(m.Instance(), uintptr(index), api.PascalBool(visible))
	return api.GoBool(r1)
}

// IsVisibleAt
//
//	Returns true (1) if the specified |index| is checked. Only applies to
//	check and radio items.
func (m *ICefMenuModel) IsVisibleAt(index uint32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_IsVisibleAt).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// IsCheckedAt
//
//	Returns true (1) if the specified |index| is checked. Only applies to
//	check and radio items.
func (m *ICefMenuModel) IsCheckedAt(index uint32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_IsCheckedAt).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// SetEnabledAt
//
//	Change the enabled status at the specified |index|. Returns true (1) on
//	success.
func (m *ICefMenuModel) SetEnabledAt(index uint32, enabled bool) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetEnabledAt).Call(m.Instance(), uintptr(index), api.PascalBool(enabled))
	return api.GoBool(r1)
}

// HasAcceleratorAt
//
//	Returns true (1) if the specified |index| has a keyboard accelerator
//	assigned.
func (m *ICefMenuModel) HasAcceleratorAt(index uint32) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_HasAcceleratorAt).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// SetCheckedAt
//
//	Check the specified |index|. Only applies to check and radio items.
//	Returns true (1) on success.
func (m *ICefMenuModel) SetCheckedAt(index uint32, checked bool) bool {
	r1, _, _ := imports.Proc(def.CefMenuModel_SetCheckedAt).Call(m.Instance(), uintptr(index), api.PascalBool(checked))
	return api.GoBool(r1)
}
