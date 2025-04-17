//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 原生系统托盘

package cef

import (
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/pkgs/notice"
	"github.com/cyber-xxm/energy/v2/pkgs/systray"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"io/ioutil"
	"sync"
)

// SysTray 系统原生
type SysTray struct {
	lock           sync.Mutex
	menu           *SysMenu
	icon           []byte
	title, tooltip string
	click          TrayICONClick
	dClick         TrayICONClick
	rClick         func(menu systray.IMenu)
	start, stop    func()
}

// 创建系统托盘
func newSysTray() *SysTray {
	notice.SetUniqueID(BrowserWindow.Config.WindowProperty.Title)
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
	//重置菜单
	systray.ResetMenu()
	//刷新并生成菜单
	m.refreshSystray(m.menu.items, nil)
}

// refreshSystray 刷新并生成菜单
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

// onExit 退出回调事件
func (m *SysTray) onExit() {
	if !common.IsDarwin() {
		systray.Quit()
	}
}

// AsSysTray 尝试转换为 SysTray 组件托盘，如果创建的是其它类型托盘返回nil
func (m *SysTray) AsSysTray() *SysTray {
	return m
}

// AsViewsFrameTray 尝试转换为 views framework 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *SysTray) AsViewsFrameTray() *ViewsFrameTray {
	return nil
}

// AsCEFTray 尝试转换为 LCL+CEF 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *SysTray) AsCEFTray() *CEFTray {
	return nil
}

// AsLCLTray 尝试转换为 LCL 组件托盘, 如果创建的是其它类型托盘返回nil
func (m *SysTray) AsLCLTray() *LCLTray {
	return nil
}

// Show 显示/启动 托盘
func (m *SysTray) Show() {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.start == nil {
		var runLoop = func() {
			m.start, m.stop = systray.RunWithExternalLoop(m.onReady, m.onExit)
			m.start()
		}
		if common.IsDarwin() {
			// view framework
			if application.IsMessageLoop() {
				runLoop()
			} else {
				//LCL
				QueueAsyncCall(func(id int) {
					runLoop()
				})
			}
		} else {
			//windows linux
			go runLoop()
		}
	}
}

func (m *SysTray) close() {
	if !common.IsDarwin() {
		if m.stop != nil {
			m.stop()
			m.start = nil
			m.stop = nil
		}
	}
}

// SetOnDblClick 鼠标双击事件
func (m *SysTray) SetOnDblClick(fn TrayICONClick) {
	m.dClick = fn
}

// SetOnClick 鼠标单击事件
func (m *SysTray) SetOnClick(fn TrayICONClick) {
	m.click = fn
}

// SetOnRClick 鼠标右键
func (m *SysTray) SetOnRClick(fn func(menu systray.IMenu)) {
	m.rClick = fn
}

// SetHint 设置托盘提示
func (m *SysTray) SetHint(value string) {
	m.tooltip = value
	if m.start != nil {
		systray.SetTooltip(value)
	}
}

// SetTitle 设置托盘标题
func (m *SysTray) SetTitle(title string) {
	m.title = title
	if m.start != nil {
		systray.SetTitle(title)
	}
}

// Notice
// 显示系统通知
//
// title 标题
//
// content 内容
//
// timeout 显示时间(毫秒)
func (m *SysTray) Notice(title, content string, timeout int32) {
	notification(nil, title, content, timeout)
}

// SetIconFS 设置托盘图标
func (m *SysTray) SetIconFS(iconResourcePath string) {
	data, err := emfs.GetResources(iconResourcePath)
	if err == nil {
		m.icon = data
		if m.start != nil {
			systray.SetIcon(m.icon)
		}
	}
}

// SetIcon 设置托盘图标
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

// ResetMenu 重置托盘菜单
func (m *SysTray) ResetMenu() {
	if m.start == nil {
		systray.ResetMenu()
	}
}

// Add 添加一个菜单项
func (m *SysTray) Add(menuItem *SysMenuItem) {
	m.menu.Add(menuItem)
}

// AddMenuItem 添加一个菜单项
func (m *SysTray) AddMenuItem(label string, click ...MenuItemClick) *SysMenuItem {
	if len(click) > 0 {
		return m.menu.AddMenuItem(label, click[0])
	}
	return m.menu.AddMenuItem(label, nil)
}

// AddSeparator 添加一个分隔线
func (m *SysTray) AddSeparator() {
	m.menu.AddSeparator()
}

// NewMenuItem 创建一个新菜单项
func (m *SysTray) NewMenuItem(label string, click ...MenuItemClick) *SysMenuItem {
	if len(click) > 0 {
		return &SysMenuItem{label: label, click: click[0]}
	}
	return &SysMenuItem{label: label}
}
