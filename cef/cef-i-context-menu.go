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
	. "github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"strings"
	"unsafe"
)

// 快捷键
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

type iCefContextMenuParams struct {
	XCoord            uintptr
	YCoord            uintptr
	TypeFlags         uintptr
	LinkUrl           uintptr
	UnfilteredLinkUrl uintptr
	SourceUrl         uintptr
	TitleText         uintptr
	PageUrl           uintptr
	FrameUrl          uintptr
	FrameCharset      uintptr
	MediaType         uintptr
	MediaStateFlags   uintptr
	SelectionText     uintptr
	EditStateFlags    uintptr
}

type ICefContextMenuParams struct {
	XCoord            int32
	YCoord            int32
	TypeFlags         TCefContextMenuTypeFlags
	LinkUrl           string
	UnfilteredLinkUrl string
	SourceUrl         string
	TitleText         string
	PageUrl           string
	FrameUrl          string
	FrameCharset      string
	MediaType         TCefContextMenuMediaType
	MediaStateFlags   TCefContextMenuMediaStateFlags
	SelectionText     string
	EditStateFlags    TCefContextMenuEditStateFlags
}

type ICefMenuModel struct {
	instance unsafe.Pointer
	CefMis   *keyEventAccelerator
}

type MenuItem struct {
	CommandId   MenuId // >= 26500 and <= 28500
	Accelerator string // 快捷键 shift ctrl alt【按键VK..】(shift+ctrl+alt+n)
	Text        string //显示文本
	Label       string
	GroupId     int32 //分组 配合 MenuType 使用
	MenuType    TCefContextMenuType
	Callback    FuncCallback //点击 或 快捷键触发的回调
}

// 添加自定义快捷键
func (m *keyEventAccelerator) AddAcceleratorCustom(accelerator *AcceleratorCustom) {
	if accelerator == nil {
		return
	}
	accelerator.Accelerator = strings.ReplaceAll(strings.ToUpper(accelerator.Accelerator), " ", "")
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

// 生成并返回下一个快捷键ID
func (m *keyEventAccelerator) NextCommandId() MenuId {
	m.menuId++
	return m.menuId
}

// 当前快捷键ID
func (m *keyEventAccelerator) CommandId() MenuId {
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

// 添加一个菜单项, 选项类型, 名称, commandId, 快捷键和触发函数
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
		item.Accelerator = strings.ReplaceAll(strings.ToUpper(item.Accelerator), " ", "")
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

func (m *ICefMenuModel) AddSeparator() bool {
	return cefMenuModel_AddSeparator(uintptr(m.instance))
}
func (m *ICefMenuModel) Clear() bool {
	return cefMenuModel_Clear(uintptr(m.instance))
}
func (m *ICefMenuModel) IsSubMenu() bool {
	return cefMenuModel_IsSubMenu(uintptr(m.instance))
}
func (m *ICefMenuModel) GetCount() int32 {
	return cefMenuModel_GetCount(uintptr(m.instance))
}
func (m *ICefMenuModel) AddItem(commandId MenuId, text string) bool {
	return cefMenuModel_AddItem(uintptr(m.instance), commandId, text)
}
func (m *ICefMenuModel) AddCheckItem(commandId MenuId, text string) bool {
	return cefMenuModel_AddCheckItem(uintptr(m.instance), commandId, text)
}
func (m *ICefMenuModel) AddRadioItem(commandId MenuId, text string, groupId int32) bool {
	return cefMenuModel_AddRadioItem(uintptr(m.instance), commandId, text, groupId)
}
func (m *ICefMenuModel) AddSubMenu(commandId MenuId, text string) *ICefMenuModel {
	return cefMenuModel_AddSubMenu(uintptr(m.instance), commandId, text)
}
func (m *ICefMenuModel) Remove(commandId MenuId) bool {
	return cefMenuModel_Remove(uintptr(m.instance), commandId)
}
func (m *ICefMenuModel) RemoveAt(index int32) bool {
	return cefMenuModel_RemoveAt(uintptr(m.instance), index)
}
func (m *ICefMenuModel) SetChecked(commandId MenuId, check bool) bool {
	return cefMenuModel_SetChecked(uintptr(m.instance), commandId, check)
}
func (m *ICefMenuModel) IsChecked(commandId MenuId) bool {
	return cefMenuModel_IsChecked(uintptr(m.instance), commandId)
}
func (m *ICefMenuModel) SetColor(commandId MenuId, colorType TCefMenuColorType, color *TCefARGB) bool {
	return cefMenuModel_SetColor(uintptr(m.instance), commandId, colorType, color)
}

//	func (m *ICefMenuModel) SetFontList(commandId MenuId, fontList string) bool {
//		return cefMenuModel_SetFontList(uintptr(m.instance), commandId, fontList)
//	}
func (m *ICefMenuModel) HasAccelerator(commandId MenuId) bool {
	return cefMenuModel_HasAccelerator(uintptr(m.instance), commandId)
}
func (m *ICefMenuModel) SetAccelerator(commandId MenuId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	return cefMenuModel_SetAccelerator(uintptr(m.instance), commandId, keyCode, shiftPressed, ctrlPressed, altPressed)
}
func (m *ICefMenuModel) RemoveAccelerator(commandId MenuId) bool {
	return cefMenuModel_RemoveAccelerator(uintptr(m.instance), commandId)
}
func (m *ICefMenuModel) IsVisible(commandId MenuId) bool {
	return cefMenuModel_IsVisible(uintptr(m.instance), commandId)
}
func (m *ICefMenuModel) SetVisible(commandId MenuId, visible bool) bool {
	return cefMenuModel_SetVisible(uintptr(m.instance), commandId, visible)
}
func (m *ICefMenuModel) IsEnabled(commandId MenuId) bool {
	return cefMenuModel_IsEnabled(uintptr(m.instance), commandId)
}
func (m *ICefMenuModel) SetEnabled(commandId MenuId, enabled bool) bool {
	return cefMenuModel_SetEnabled(uintptr(m.instance), commandId, enabled)
}
func (m *ICefMenuModel) SetLabel(commandId MenuId, text string) bool {
	return cefMenuModel_SetLabel(uintptr(m.instance), commandId, text)
}
func (m *ICefMenuModel) GetIndexOf(commandId MenuId) int32 {
	return cefMenuModel_GetIndexOf(uintptr(m.instance), commandId)
}

// ------------------------------------ PROC
// ICefMenuModel cefMenuModel_AddSeparator
func cefMenuModel_AddSeparator(instance uintptr) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_AddSeparator).Call(instance)
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_Clear
func cefMenuModel_Clear(instance uintptr) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_Clear).Call(instance)
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_IsSubMenu
func cefMenuModel_IsSubMenu(instance uintptr) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_IsSubMenu).Call(instance)
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_GetCount
func cefMenuModel_GetCount(instance uintptr) int32 {
	r1, _, _ := imports.Proc(internale_cefMenuModel_GetCount).Call(instance)
	return int32(r1)
}

// ICefMenuModel cefMenuModel_AddItem
func cefMenuModel_AddItem(instance uintptr, commandId MenuId, text string) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_AddItem).Call(instance, uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_AddCheckItem
func cefMenuModel_AddCheckItem(instance uintptr, commandId MenuId, text string) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_AddCheckItem).Call(instance, uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_AddRadioItem
func cefMenuModel_AddRadioItem(instance uintptr, commandId MenuId, text string, groupId int32) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_AddRadioItem).Call(instance, uintptr(commandId), api.PascalStr(text), uintptr(groupId))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_AddSubMenu
func cefMenuModel_AddSubMenu(instance uintptr, commandId MenuId, text string) *ICefMenuModel {
	var ret uintptr
	imports.Proc(internale_cefMenuModel_AddSubMenu).Call(instance, uintptr(commandId), api.PascalStr(text), uintptr(unsafe.Pointer(&ret)))
	return &ICefMenuModel{
		instance: unsafe.Pointer(ret),
	}
}

// ICefMenuModel cefMenuModel_Remove
func cefMenuModel_Remove(instance uintptr, commandId MenuId) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_Remove).Call(instance, uintptr(commandId))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_Remove
func cefMenuModel_RemoveAt(instance uintptr, index int32) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_RemoveAt).Call(instance, uintptr(index))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_SetChecked
func cefMenuModel_SetChecked(instance uintptr, commandId MenuId, check bool) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_SetChecked).Call(instance, uintptr(commandId), api.PascalBool(check))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_IsChecked
func cefMenuModel_IsChecked(instance uintptr, commandId MenuId) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_IsChecked).Call(instance, uintptr(commandId))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_SetColor
func cefMenuModel_SetColor(instance uintptr, commandId MenuId, colorType TCefMenuColorType, color *TCefARGB) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_SetColor).Call(instance, uintptr(commandId), uintptr(colorType), uintptr(color.ARGB()))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_SetFontList
func cefMenuModel_SetFontList(instance uintptr, commandId MenuId, fontList string) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_SetFontList).Call(instance, uintptr(commandId), api.PascalStr(fontList))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_HasAccelerator
func cefMenuModel_HasAccelerator(instance uintptr, commandId MenuId) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_HasAccelerator).Call(instance, uintptr(commandId))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_SetAccelerator
func cefMenuModel_SetAccelerator(instance uintptr, commandId MenuId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_SetAccelerator).Call(instance, uintptr(commandId), uintptr(keyCode),
		api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_RemoveAccelerator
func cefMenuModel_RemoveAccelerator(instance uintptr, commandId MenuId) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_RemoveAccelerator).Call(instance, uintptr(commandId))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_IsVisible
func cefMenuModel_IsVisible(instance uintptr, commandId MenuId) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_IsVisible).Call(instance, uintptr(commandId))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_SetVisible
func cefMenuModel_SetVisible(instance uintptr, commandId MenuId, visible bool) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_SetVisible).Call(instance, uintptr(commandId), api.PascalBool(visible))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_IsEnabled
func cefMenuModel_IsEnabled(instance uintptr, commandId MenuId) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_IsEnabled).Call(instance, uintptr(commandId))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_SetEnabled
func cefMenuModel_SetEnabled(instance uintptr, commandId MenuId, enabled bool) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_SetEnabled).Call(instance, uintptr(commandId), api.PascalBool(enabled))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_SetLabel
func cefMenuModel_SetLabel(instance uintptr, commandId MenuId, text string) bool {
	r1, _, _ := imports.Proc(internale_cefMenuModel_SetLabel).Call(instance, uintptr(commandId), api.PascalStr(text))
	return api.GoBool(r1)
}

// ICefMenuModel cefMenuModel_GetIndexOf
func cefMenuModel_GetIndexOf(instance uintptr, commandId MenuId) int32 {
	r1, _, _ := imports.Proc(internale_cefMenuModel_GetIndexOf).Call(instance, uintptr(commandId))
	return int32(r1)
}
