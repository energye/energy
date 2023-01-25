//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"energye/systray"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"io/ioutil"
)

//创建系统托盘
func newSysTray() *SysTray {
	tray := &SysTray{
		menu: &SysMenu{
			items: make([]*SysMenuItem, 0, 0),
		},
	}
	return tray
}

func (m *SysTray) onReady() {
	if len(m.icon) > 0 {
		systray.SetIcon(m.icon)
	}
	if m.click != nil {
		systray.SetOnClick(m.click)
	}
	if m.dClick != nil {
		systray.SetOnDClick(m.dClick)
	}
	if m.rClick != nil {
		systray.SetOnRClick(m.rClick)
	}
	if m.title != "" {
		systray.SetTitle(m.title)
	}
	if m.tooltip != "" {
		systray.SetTooltip(m.tooltip)
	}
	systray.ResetMenu()
	m.refreshSystray(m.menu.items, nil)
	go func() {
		for {
			select {}
		}
	}()
}

func (m *SysTray) refreshSystray(items []*SysMenuItem, parent *systray.MenuItem) {
	for _, item := range items {
		mItem := itemForMenuItem(item, parent)
		if mItem == nil {
			continue
		}
		if item.childMenu != nil && len(item.childMenu.items) > 0 {
			m.refreshSystray(item.childMenu.items, mItem)
		}
	}
}

func (m *SysTray) onExit() {
	systray.Quit()
}

func (m *SysTray) AsSysTray() *SysTray {
	return m
}

func (m *SysTray) AsViewsFrameTray() *ViewsFrameTray {
	return nil
}

func (m *SysTray) AsCEFTray() *CEFTray {
	return nil
}

func (m *SysTray) AsLCLTray() *LCLTray {
	return nil
}

func (m *SysTray) Show() {
	if m.start == nil {
		m.start, m.stop = systray.RunWithExternalLoop(m.onReady, m.onExit)
		m.start()
	}
}

func (m *SysTray) Hide() {
}

func (m *SysTray) close() {
	m.onExit()
}

// CreateMenu 创建托盘菜单, 如果托盘菜单是空, 把菜单项添加到托盘
// 该方法主动调用后 如果托盘菜单已创建则添加进去, 之后鼠标事件失效
//
// 仅MacOSX平台
func (m *SysTray) CreateMenu() {
	if m.start == nil {
		systray.CreateMenu()
	}
}

// SetMenuNil 托盘菜单设置为nil, 如果托盘菜单不是空, 把菜单项设置为nil
// 该方法主动调用后 将移除托盘菜单, 之后鼠标事件生效
//
// 仅MacOSX平台
func (m *SysTray) SetMenuNil() {
	if m.start == nil {
		systray.SetMenuNil()
	}
}

//ResetMenu 重置托盘菜单
func (m *SysTray) ResetMenu() {
	if m.start == nil {
		systray.ResetMenu()
	}
}

//Add 添加一个菜单项
func (m *SysTray) Add(menuItem *SysMenuItem) {
	m.menu.Add(menuItem)
}

//AddMenuItem 添加一个菜单项
func (m *SysTray) AddMenuItem(label string, onClick MenuItemClick) *SysMenuItem {
	return m.menu.AddMenuItem(label, onClick)
}

//AddMenuItemSeparator 添加一个分隔线
func (m *SysTray) AddMenuItemSeparator() {
	m.menu.AddMenuItemSeparator()
}

//NewMenuItem 创建一个新菜单项
func (m *SysTray) NewMenuItem(label string, onClick MenuItemClick) *SysMenuItem {
	return &SysMenuItem{label: label, click: onClick}
}

//SetOnDblClick 鼠标双击事件
func (m *SysTray) SetOnDblClick(fn TrayICONClick) {
	m.dClick = fn
}

//SetOnClick 鼠标单击事件
func (m *SysTray) SetOnClick(fn TrayICONClick) {
	m.click = fn
}

//SetOnRClick 鼠标右键
func (m *SysTray) SetOnRClick(fn func(menu systray.IMenu)) {
	m.rClick = fn
}

func (m *SysTray) SetHint(value string) {
	m.tooltip = value
	if m.start != nil {
		systray.SetTooltip(value)
	}
}

func (m *SysTray) SetTitle(title string) {
	m.title = title
	if m.start != nil {
		systray.SetTitle(title)
	}
}

//SetIconFS 设置托盘图标
func (m *SysTray) SetIconFS(iconResourcePath string) {
	if emfs.IsExist(iconResourcePath) {
		data, err := emfs.GetResources(iconResourcePath)
		if err == nil {
			m.icon = data
			if m.start != nil {
				systray.SetIcon(m.icon)
			}
		}
	}
}

//SetIcon 设置托盘图标
func (m *SysTray) SetIcon(iconResourcePath string) {
	if tools.IsExist(iconResourcePath) {
		data, err := ioutil.ReadFile(iconResourcePath)
		if err == nil {
			m.icon = data
			if m.start != nil {
				systray.SetIcon(m.icon)
			}
		}
	}
}

func (m *SysTray) Visible() bool {
	return false
}

func (m *SysTray) SetVisible(v bool) {
}

//设置托盘气泡
//title 气泡标题
//content 气泡内容
//timeout 显示时间(毫秒)
func (m *SysTray) SetBalloon(title, content string, timeout int32) ITray {
	return nil
}

//显示托盘气泡
func (m *SysTray) ShowBalloon() {
}
