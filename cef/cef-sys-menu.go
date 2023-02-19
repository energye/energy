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
	"github.com/energye/energy/pkgs/systray"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"io/ioutil"
)

type MenuItemClick func()

type SysMenu struct {
	label string
	items []*SysMenuItem
}

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

func (m *SysMenu) AddMenuItem(label string, action MenuItemClick) *SysMenuItem {
	item := &SysMenuItem{
		label: label,
		click: action,
	}
	m.items = append(m.items, item)
	return item
}

func (m *SysMenu) Add(menuItem *SysMenuItem) {
	m.items = append(m.items, menuItem)
}

func (m *SysMenu) AddSeparator() {
	m.items = append(m.items, &SysMenuItem{isSeparator: true})
}

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
	if emfs.IsExist(iconResourcePath) {
		data, err := emfs.GetResources(iconResourcePath)
		if err == nil {
			m.SetIconBytes(data)
		}
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

func (m *SysMenuItem) Icon() []byte {
	return m.icon
}

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

func (m *SysMenuItem) Checked() bool {
	if m.menuItem != nil {
		m.checked = m.menuItem.Checked()
	}
	return m.checked
}

func (m *SysMenuItem) Check() {
	m.checked = true
	if m.menuItem != nil {
		m.menuItem.Check()
	}
}

func (m *SysMenuItem) Uncheck() {
	m.checked = false
	if m.menuItem != nil {
		m.menuItem.Uncheck()
	}
}

func (m *SysMenuItem) SetDisable(v bool) {
	m.disabled = v
	if m.menuItem != nil {
		m.menuItem.Disable()
	}
}

func (m *SysMenuItem) Disable() {
	if m.menuItem != nil {
		m.menuItem.Disable()
	}
}

func (m *SysMenuItem) Enable() {
	if m.menuItem != nil {
		m.menuItem.Enable()
	}
}

func (m *SysMenuItem) Disabled() bool {
	if m.menuItem != nil {
		m.disabled = m.menuItem.Disabled()
	}
	return m.disabled
}

func (m *SysMenuItem) Show() {
	if m.menuItem != nil {
		m.menuItem.Show()
	}
}

func (m *SysMenuItem) Hide() {
	if m.menuItem != nil {
		m.menuItem.Hide()
	}
}

func (m *SysMenuItem) SetSeparator(v bool) {
	m.isSeparator = v
}

func (m *SysMenuItem) IsSeparator() bool {
	return m.isSeparator
}

func (m *SysMenuItem) SetLabel(v string) {
	m.label = v
	if m.menuItem != nil {
		m.menuItem.SetTitle(v)
	}
}

func (m *SysMenuItem) SetTooltip(v string) {
	if m.menuItem != nil {
		m.menuItem.SetTooltip(v)
	}
}

func (m *SysMenuItem) Label() string {
	return m.label
}

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
