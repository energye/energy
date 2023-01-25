//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

type MenuItemClick func()

type SysMenu struct {
	Label string
	Items []*SysMenuItem
}

type SysMenuItem struct {
	ChildMenu   *SysMenu
	Label       string
	Action      MenuItemClick
	IsSeparator bool
	Disabled    bool
	Checked     bool
	Icon        []byte
	Shortcut    string
}

func (m *SysMenu) AddMenuItem(label string, action MenuItemClick) *SysMenuItem {
	item := &SysMenuItem{
		Label:  label,
		Action: action,
	}
	m.Items = append(m.Items, item)
	return item
}

func (m *SysMenu) Add(menuItem *SysMenuItem) {
	m.Items = append(m.Items, menuItem)
}

func (m *SysMenu) AddMenuItemSeparator() {
	m.Items = append(m.Items, &SysMenuItem{IsSeparator: true})
}

func (m *SysMenuItem) Add(label string, items ...*SysMenuItem) *SysMenu {
	m.ChildMenu = &SysMenu{Label: label, Items: items}
	return m.ChildMenu
}

// NewMenu 创建一个新菜单，给定指定的标签和要显示的项目列表
func NewMenu(label string, items ...*SysMenuItem) *SysMenu {
	return &SysMenu{Label: label, Items: items}
}

// NewMenuItem 根据传递的标签和操作参数创建一个新菜单项
func NewMenuItem(label string, action MenuItemClick) *SysMenuItem {
	return &SysMenuItem{Label: label, Action: action}
}

// NewMenuItemSeparator 创建将用作分隔符的菜单项
func NewMenuItemSeparator() *SysMenuItem {
	return &SysMenuItem{IsSeparator: true}
}
