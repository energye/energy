//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy 系统菜单 -> 系统托盘

package cef

import (
	"github.com/cyber-xxm/energy/v2/pkgs/systray"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"io/ioutil"
)

type MenuItemClick func()

// SysMenu 系统菜单
type SysMenu struct {
	label string
	items []*SysMenuItem
}

// SysMenuItem 菜单项
type SysMenuItem struct {
	childMenu   *SysMenu
	menuItem    *systray.MenuItem
	label       string
	click       MenuItemClick
	isSeparator bool
	disabled    bool
	checked     bool
	icon        []byte
	shortcut    string
}

// AddMenuItem 添加菜单项
func (m *SysMenu) AddMenuItem(label string, action MenuItemClick) *SysMenuItem {
	item := &SysMenuItem{
		label: label,
		click: action,
	}
	m.items = append(m.items, item)
	return item
}

// Add 添加菜单项
func (m *SysMenu) Add(menuItem *SysMenuItem) {
	m.items = append(m.items, menuItem)
}

// AddSeparator 添加分隔线
func (m *SysMenu) AddSeparator() {
	m.items = append(m.items, &SysMenuItem{isSeparator: true})
}

// AddSubMenu 添加并创建子菜单，返回菜单项
func (m *SysMenuItem) AddSubMenu(label string, click ...MenuItemClick) *SysMenuItem {
	if m.childMenu == nil {
		m.childMenu = &SysMenu{
			items: make([]*SysMenuItem, 0, 0),
		}
	}
	if len(click) > 0 {
		return m.childMenu.AddMenuItem(label, click[0])
	}
	return m.childMenu.AddMenuItem(label, nil)
}

// SetIconFS windows推荐使用ico图标, linux推荐使用png图标, macosx使用ico和png都可
func (m *SysMenuItem) SetIconFS(iconResourcePath string) {
	data, err := emfs.GetResources(iconResourcePath)
	if err == nil {
		m.SetIconBytes(data)
	}
}

// SetIcon windows推荐使用ico图标, linux推荐使用png图标, macosx使用ico和png都可
func (m *SysMenuItem) SetIcon(iconResourcePath string) {
	if tools.IsExist(iconResourcePath) {
		data, err := ioutil.ReadFile(iconResourcePath)
		if err == nil {
			m.SetIconBytes(data)
		}
	}
}

// SetIcon windows推荐使用ico图标, linux推荐使用png图标, macosx使用ico和png都可
func (m *SysMenuItem) SetIconBytes(v []byte) {
	m.icon = v
	if m.menuItem != nil {
		m.menuItem.SetIcon(v)
	}
}

// Icon 获取图标
func (m *SysMenuItem) Icon() []byte {
	return m.icon
}

// SetChecked 设置选中
func (m *SysMenuItem) SetChecked(v bool) {
	m.checked = v
	if m.menuItem != nil {
		if v {
			m.menuItem.Checked()
		} else {
			m.menuItem.Uncheck()
		}
	}
}

// Checked 选中状态
func (m *SysMenuItem) Checked() bool {
	if m.menuItem != nil {
		m.checked = m.menuItem.Checked()
	}
	return m.checked
}

// Check
func (m *SysMenuItem) Check() {
	m.checked = true
	if m.menuItem != nil {
		m.menuItem.Check()
	}
}

// Uncheck
func (m *SysMenuItem) Uncheck() {
	m.checked = false
	if m.menuItem != nil {
		m.menuItem.Uncheck()
	}
}

// SetDisable 设置禁用/启用
func (m *SysMenuItem) SetDisable(v bool) {
	m.disabled = v
	if m.menuItem != nil {
		m.menuItem.Disable()
	}
}

// Disable 禁用状态
func (m *SysMenuItem) Disable() {
	if m.menuItem != nil {
		m.menuItem.Disable()
	}
}

// Enable 启用
func (m *SysMenuItem) Enable() {
	if m.menuItem != nil {
		m.menuItem.Enable()
	}
}

// Disabled 禁用
func (m *SysMenuItem) Disabled() bool {
	if m.menuItem != nil {
		m.disabled = m.menuItem.Disabled()
	}
	return m.disabled
}

// Show 显示
func (m *SysMenuItem) Show() {
	if m.menuItem != nil {
		m.menuItem.Show()
	}
}

// Remove 移除
func (m *SysMenuItem) Remove() {
	if m.menuItem != nil {
		m.menuItem.Remove()
	}
}

// Hide 隐藏
func (m *SysMenuItem) Hide() {
	if m.menuItem != nil {
		m.menuItem.Hide()
	}
}

// SetSeparator 设置是否分隔线
func (m *SysMenuItem) SetSeparator(v bool) {
	m.isSeparator = v
}

// IsSeparator 是否分隔线
func (m *SysMenuItem) IsSeparator() bool {
	return m.isSeparator
}

// SetLabel 设置Label
func (m *SysMenuItem) SetLabel(v string) {
	m.label = v
	if m.menuItem != nil {
		m.menuItem.SetTitle(v)
	}
}

// SetTooltip 设置提示
func (m *SysMenuItem) SetTooltip(v string) {
	if m.menuItem != nil {
		m.menuItem.SetTooltip(v)
	}
}

// Label 获取Label
func (m *SysMenuItem) Label() string {
	return m.label
}

// Click 设置菜单项事件
func (m *SysMenuItem) Click(fn MenuItemClick) {
	m.click = fn
}

func itemForMenuItem(item *SysMenuItem, parent *systray.MenuItem) *systray.MenuItem {
	if item == nil || item.isSeparator {
		systray.AddSeparator()
		return nil
	}
	var mItem *systray.MenuItem
	if item.checked {
		if parent != nil {
			mItem = parent.AddSubMenuItemCheckbox(item.label, item.label, true)
		} else {
			mItem = systray.AddMenuItemCheckbox(item.label, item.label, true)
		}
	} else {
		if parent != nil {
			mItem = parent.AddSubMenuItem(item.label, item.label)
		} else {
			mItem = systray.AddMenuItem(item.label, item.label)
		}
	}
	if item.disabled {
		mItem.Disable()
	}
	if item.icon != nil {
		mItem.SetIcon(item.icon)
	}
	item.menuItem = mItem
	if item.click != nil {
		item.menuItem.Click(item.click)
	}
	return mItem
}
