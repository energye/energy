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
	return &SysTray{
		menu: &SysMenu{
			Items: make([]*SysMenuItem, 0, 0),
		},
	}
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
}

func (m *SysTray) Hide() {
}

func (m *SysTray) close() {
}

func (m *SysTray) AddMenuItem(label string, onClick MenuItemClick) *SysMenuItem {
	return m.menu.AddMenuItem(label, onClick)
}

func (m *SysTray) AddMenuItemSeparator() {
	m.menu.AddMenuItemSeparator()
}

func (m *SysTray) NewMenuItem(label string, onClick MenuItemClick) *SysMenuItem {
	return NewMenuItem(label, onClick)
}

func (m *SysTray) SetOnDblClick(fn TrayICONClick) {
	systray.SetOnDClick(fn)
}

func (m *SysTray) SetOnClick(fn TrayICONClick) {
	systray.SetOnClick(fn)
}

func (m *SysTray) SetOnRClick(fn func(menu systray.IMenu)) {
	systray.SetOnRClick(fn)
}

func (m *SysTray) Visible() bool {
	return false
}

func (m *SysTray) SetVisible(v bool) {
}

func (m *SysTray) SetHint(value string) {
	systray.SetTooltip(value)
}

func (m *SysTray) SetTitle(title string) {
	systray.SetTitle(title)
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

//设置托盘图标
func (m *SysTray) SetIconFS(iconResourcePath string) {
	if emfs.IsExist(iconResourcePath) {
		file, err := emfs.GetResources(iconResourcePath)
		if err == nil {
			systray.SetTemplateIcon(file, file)
		}
	}
}

//设置托盘图标
func (m *SysTray) SetIcon(iconResourcePath string) {
	if tools.IsExist(iconResourcePath) {
		file, err := ioutil.ReadFile(iconResourcePath)
		if err == nil {
			systray.SetTemplateIcon(file, file)
		}
	}
}
