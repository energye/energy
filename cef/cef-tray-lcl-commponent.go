//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//创建系统托盘
func newTray(owner lcl.IComponent) *LCLTray {
	trayIcon := lcl.NewTrayIcon(owner)
	return &LCLTray{
		owner:    owner,
		trayIcon: trayIcon,
	}
}

func (m *LCLTray) AsSysTray() *SysTray {
	return nil
}

func (m *LCLTray) AsViewsFrameTray() *ViewsFrameTray {
	return nil
}

func (m *LCLTray) AsCEFTray() *CEFTray {
	return nil
}

func (m *LCLTray) AsLCLTray() *LCLTray {
	return m
}

func (m *LCLTray) SetVisible(v bool) {
	m.trayIcon.SetVisible(v)
}

func (m *LCLTray) Visible() bool {
	return m.trayIcon.Visible()
}

func (m *LCLTray) Show() {
	m.SetVisible(true)
}

func (m *LCLTray) Hide() {
	m.SetVisible(false)
}

func (m *LCLTray) close() {
	m.Hide()
}

func (m *LCLTray) SetOnDblClick(fn TrayICONClick) {
	m.trayIcon.SetOnDblClick(func(sender lcl.IObject) {
		fn()
	})
}

func (m *LCLTray) SetOnClick(fn TrayICONClick) {
	m.trayIcon.SetOnClick(func(sender lcl.IObject) {
		fn()
	})
}

func (m *LCLTray) SetOnMouseUp(fn TMouseEvent) {
	m.trayIcon.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fn(sender, button, shift, x, y)
	})
}

func (m *LCLTray) SetOnMouseDown(fn lcl.TMouseEvent) {
	m.trayIcon.SetOnMouseDown(fn)
}

func (m *LCLTray) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	m.trayIcon.SetOnMouseMove(fn)
}

//创建并返回托盘根菜单 PopupMenu
func (m *LCLTray) TrayMenu() *lcl.TPopupMenu {
	if m.popupMenu == nil {
		m.popupMenu = lcl.NewPopupMenu(m.trayIcon)
		m.trayIcon.SetPopupMenu(m.popupMenu)
	}
	return m.popupMenu
}

//设置托盘图标
func (m *LCLTray) SetIconFS(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFSFile(iconResourcePath)
}

//设置托盘图标
func (m *LCLTray) SetIcon(iconResourcePath string) {
	m.trayIcon.Icon().LoadFromFile(iconResourcePath)
}

func (m *LCLTray) SetHint(value string) {
	m.trayIcon.SetHint(value)
}

//SetTitle 设置标题
func (m *LCLTray) SetTitle(title string) {
	m.trayIcon.SetHint(title)
}

//显示系统通知
//
//title 标题
//
//content 内容
//
//timeout 显示时间(毫秒)
func (m *LCLTray) Notice(title, content string, timeout int32) {
	notification(m.trayIcon, title, content, timeout)
}

//创建一个菜单，还未添加到托盘
func (m *LCLTray) NewMenuItem(caption string, onClick MenuItemClick) *lcl.TMenuItem {
	item := lcl.NewMenuItem(m.trayIcon)
	item.SetCaption(caption)
	if onClick != nil {
		item.SetOnClick(func(sender lcl.IObject) {
			onClick()
		})
	}
	return item
}

//添加一个托盘菜单
func (m *LCLTray) AddMenuItem(caption string, onClick MenuItemClick) *lcl.TMenuItem {
	item := m.NewMenuItem(caption, onClick)
	m.TrayMenu().Items().Add(item)
	return item
}
