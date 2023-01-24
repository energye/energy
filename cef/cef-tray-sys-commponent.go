//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import "energye/systray"

//创建系统托盘
func newSysTray() *SysTray {
	return &SysTray{}
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

func (m *SysTray) SetOnDblClick(fn TrayICONClick) {

}

func (m *SysTray) SetOnClick(fn TrayICONClick) {

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
}

//设置托盘图标
func (m *SysTray) SetIcon(iconResourcePath string) {
}
