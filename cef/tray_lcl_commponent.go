//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 基于 LCL 系统托盘

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"path/filepath"
	"strings"
)

// LCLTray LCL组件 托盘
type LCLTray struct {
	owner     lcl.IComponent
	trayIcon  *lcl.TTrayIcon
	popupMenu *lcl.TPopupMenu
}

// 创建系统托盘
func newTray(owner lcl.IComponent) *LCLTray {
	trayIcon := lcl.NewTrayIcon(owner)
	return &LCLTray{
		owner:    owner,
		trayIcon: trayIcon,
	}
}

// AsSysTray 尝试转换为 SysTray 组件托盘，如果创建的是其它类型托盘返回nil
func (m *LCLTray) AsSysTray() *SysTray {
	return nil
}

// AsViewsFrameTray 尝试转换为 views framework 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *LCLTray) AsViewsFrameTray() *ViewsFrameTray {
	return nil
}

// AsCEFTray 尝试转换为 LCL+CEF 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *LCLTray) AsCEFTray() *CEFTray {
	return nil
}

// AsLCLTray 尝试转换为 LCL 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *LCLTray) AsLCLTray() *LCLTray {
	return m
}

// SetVisible 设置显示状态
func (m *LCLTray) SetVisible(v bool) {
	m.trayIcon.SetVisible(v)
}

// Visible 显示状态
func (m *LCLTray) Visible() bool {
	return m.trayIcon.Visible()
}

// Show 显示/启动 托盘
func (m *LCLTray) Show() {
	m.SetVisible(true)
}

// Hide 隐藏 托盘
func (m *LCLTray) Hide() {
	m.SetVisible(false)
}

func (m *LCLTray) close() {
	m.Hide()
}

// SetOnDblClick 设置双击事件
func (m *LCLTray) SetOnDblClick(fn TrayICONClick) {
	m.trayIcon.SetOnDblClick(func(sender lcl.IObject) {
		fn()
	})
}

// SetOnClick 设置单击事件
func (m *LCLTray) SetOnClick(fn TrayICONClick) {
	m.trayIcon.SetOnClick(func(sender lcl.IObject) {
		fn()
	})
}

// SetOnMouseUp 鼠标抬起事件
func (m *LCLTray) SetOnMouseUp(fn TMouseEvent) {
	m.trayIcon.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fn(sender, button, shift, x, y)
	})
}

// SetOnMouseDown 鼠标按下事件
func (m *LCLTray) SetOnMouseDown(fn lcl.TMouseEvent) {
	m.trayIcon.SetOnMouseDown(fn)
}

// SetOnMouseMove 鼠标移动事件
func (m *LCLTray) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	m.trayIcon.SetOnMouseMove(fn)
}

// TrayMenu 创建并返回托盘根菜单 PopupMenu
func (m *LCLTray) TrayMenu() *lcl.TPopupMenu {
	if m.popupMenu == nil {
		m.popupMenu = lcl.NewPopupMenu(m.trayIcon)
		m.trayIcon.SetPopupMenu(m.popupMenu)
	}
	return m.popupMenu
}

// SetIconFS 设置托盘图标
func (m *LCLTray) SetIconFS(iconResourcePath string) {
	ext := strings.ToLower(filepath.Ext(iconResourcePath))
	switch ext {
	case ".png":
		png := lcl.NewPngImage()
		png.LoadFromFSFile(iconResourcePath)
		m.trayIcon.Icon().Assign(png)
		png.Free()
	case ".jpeg":
		jpeg := lcl.NewJPEGImage()
		jpeg.LoadFromFSFile(iconResourcePath)
		m.trayIcon.Icon().Assign(jpeg)
		jpeg.Free()
	case ".ico":
		m.trayIcon.Icon().LoadFromFSFile(iconResourcePath)
	}
}

// SetIcon 设置托盘图标
func (m *LCLTray) SetIcon(iconResourcePath string) {
	ext := strings.ToLower(filepath.Ext(iconResourcePath))
	switch ext {
	case ".png":
		png := lcl.NewPngImage()
		png.LoadFromFile(iconResourcePath)
		m.trayIcon.Icon().Assign(png)
		png.Free()
	case ".jpeg":
		jpeg := lcl.NewJPEGImage()
		jpeg.LoadFromFile(iconResourcePath)
		m.trayIcon.Icon().Assign(jpeg)
		jpeg.Free()
	case ".ico":
		m.trayIcon.Icon().LoadFromFile(iconResourcePath)
	}
}

// TrayIcon return TTrayIcon
func (m *LCLTray) TrayIcon() *lcl.TTrayIcon {
	return m.trayIcon
}

// SetHint 设置提示
func (m *LCLTray) SetHint(value string) {
	m.trayIcon.SetHint(value)
}

// SetTitle 设置标题
func (m *LCLTray) SetTitle(title string) {
	m.trayIcon.SetHint(title)
}

// Notice
// 显示系统通知
//
// title 标题
//
// content 内容
//
// timeout 显示时间(毫秒)
func (m *LCLTray) Notice(title, content string, timeout int32) {
	notification(m.trayIcon, title, content, timeout)
}

// NewMenuItem
// 创建一个菜单，还未添加到托盘
func (m *LCLTray) NewMenuItem(caption string, onClick lcl.TNotifyEvent) *lcl.TMenuItem {
	item := lcl.NewMenuItem(m.trayIcon)
	if caption != "" {
		item.SetCaption(caption)
	}
	if onClick != nil {
		item.SetOnClick(onClick)
	}
	return item
}

// AddMenuItem
// 添加一个托盘菜单
func (m *LCLTray) AddMenuItem(caption string, onClick lcl.TNotifyEvent) *lcl.TMenuItem {
	item := m.NewMenuItem(caption, onClick)
	m.TrayMenu().Items().Add(item)
	return item
}
