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

type TMouseEvent func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) bool

type ITray interface {
	SetTitle(title string)                                 //设置标题
	SetVisible(v bool)                                     //显示和隐藏托盘图标
	Visible() bool                                         //
	Show()                                                 //显示托盘菜单窗口 windows有效
	Hide()                                                 //隐藏托盘菜单窗口 windows有效
	close()                                                //关闭托盘菜单窗口 windows有效
	SetOnDblClick(fn lcl.TNotifyEvent)                     //双击事件 linux 和 macos 可能不启作用
	SetOnClick(fn lcl.TNotifyEvent)                        //单击事件
	SetOnMouseUp(fn TMouseEvent)                           //up事件 linux 和 macos 可能不启作用
	SetOnMouseDown(fn lcl.TMouseEvent)                     //down事件 linux 和 macos 可能不启作用
	SetOnMouseMove(fn lcl.TMouseMoveEvent)                 //move事件 linux 和 macos 可能不启作用
	SetIconFS(iconResourcePath string)                     //设置托盘图标
	SetIcon(iconResourcePath string)                       //设置托盘图标
	SetHint(value string)                                  //设置托盘hint(鼠标移动到托盘图标显示的文字)
	ShowBalloon()                                          //显示托盘气泡
	SetBalloon(title, content string, timeout int32) ITray //设置托盘气泡内容
	Tray() ITray                                           //获得 LCLTray, ( CefTray 返回 nil )
	AsViewsFrameTray() *ViewsFrameTray
	AsCEFTray() *CEFTray
	AsLCLTray() *LCLTray
}

//系统托盘
type LCLTray struct {
	owner     lcl.IComponent
	trayIcon  *lcl.TTrayIcon
	popupMenu *lcl.TPopupMenu
}

//创建系统托盘
func newTray(owner lcl.IComponent) *LCLTray {
	trayIcon := lcl.NewTrayIcon(owner)
	popupMenu := lcl.NewPopupMenu(trayIcon)
	trayIcon.SetPopupMenu(popupMenu)
	trayIcon.SetVisible(true)
	return &LCLTray{
		owner:     owner,
		trayIcon:  trayIcon,
		popupMenu: popupMenu,
	}
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

func (m *LCLTray) Tray() ITray {
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

func (m *LCLTray) SetOnDblClick(fn lcl.TNotifyEvent) {
	m.trayIcon.SetOnDblClick(fn)
}

func (m *LCLTray) SetOnClick(fn lcl.TNotifyEvent) {
	m.trayIcon.SetOnClick(fn)
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

//返回托盘根菜单 PopupMenu
func (m *LCLTray) TrayMenu() *lcl.TPopupMenu {
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

func (m *LCLTray) SetTitle(title string) {
	m.trayIcon.SetHint(title)
}

//设置托盘气泡
//
//title	气泡标题
//
//content	气泡内容
//
//timeout	显示时间(毫秒)
func (m *LCLTray) SetBalloon(title, content string, timeout int32) ITray {
	m.trayIcon.SetBalloonTitle(title)
	m.trayIcon.SetBalloonHint(content)
	m.trayIcon.SetBalloonTimeout(timeout)
	return m
}

//显示托盘气泡
func (m *LCLTray) ShowBalloon() {
	m.trayIcon.ShowBalloonHint()
}

//创建一个菜单，还未添加到托盘
func (m *LCLTray) NewMenuItem(caption string, onClick func(lcl.IObject)) *lcl.TMenuItem {
	item := lcl.NewMenuItem(m.trayIcon)
	item.SetCaption(caption)
	if onClick != nil {
		item.SetOnClick(onClick)
	}
	return item
}

//添加一个托盘菜单
func (m *LCLTray) AddMenuItem(caption string, onClick func(lcl.IObject)) *lcl.TMenuItem {
	item := m.NewMenuItem(caption, onClick)
	m.popupMenu.Items().Add(item)
	return item
}
