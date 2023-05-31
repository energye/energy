//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Chromium 配置

package cef

import "github.com/energye/golcl/lcl/api"

// tCefChromiumConfig Chromium的基础配置
type tCefChromiumConfig struct {
	enableMenu        uintptr //bool 启用右键菜单
	enableViewSource  uintptr //bool 启用查看源代码
	enableDevTools    uintptr //bool 启用开发者工具
	enableWindowPopup uintptr //bool 启用弹出新窗口
	enableOpenUrlTab  uintptr //bool 启用tab签打开新窗口(需自定义实现)
	enabledJavascript uintptr //bool 启用Javascript
}

// NewChromiumConfig 创建默认禁用相关功能
func NewChromiumConfig() *tCefChromiumConfig {
	return &tCefChromiumConfig{
		enableMenu:        api.PascalBool(true),
		enableViewSource:  api.PascalBool(true),
		enableDevTools:    api.PascalBool(true),
		enableWindowPopup: api.PascalBool(true),
		enableOpenUrlTab:  api.PascalBool(false),
		enabledJavascript: api.PascalBool(false),
	}
}

// SetEnableMenu 设置启用右键菜单
func (m *tCefChromiumConfig) SetEnableMenu(value bool) *tCefChromiumConfig {
	m.enableMenu = api.PascalBool(value)
	return m
}

func (m *tCefChromiumConfig) EnableMenu() bool {
	return api.GoBool(m.enableMenu)
}

// SetEnableViewSource 设置启用查看源文件
func (m *tCefChromiumConfig) SetEnableViewSource(value bool) *tCefChromiumConfig {
	m.enableViewSource = api.PascalBool(value)
	return m
}

func (m *tCefChromiumConfig) EnableViewSource() bool {
	return api.GoBool(m.enableViewSource)
}

// SetEnableDevTools 设置启用开发者工具
func (m *tCefChromiumConfig) SetEnableDevTools(value bool) *tCefChromiumConfig {
	m.enableDevTools = api.PascalBool(value)
	return m
}

func (m *tCefChromiumConfig) EnableDevTools() bool {
	return api.GoBool(m.enableDevTools)
}

// SetEnableWindowPopup 设置启用弹出新窗口
//	与tab互斥
func (m *tCefChromiumConfig) SetEnableWindowPopup(value bool) *tCefChromiumConfig {
	m.enableWindowPopup = api.PascalBool(value)
	if value {
		m.SetEnableOpenUrlTab(false)
	}
	return m
}

func (m *tCefChromiumConfig) EnableWindowPopup() bool {
	return api.GoBool(m.enableWindowPopup)
}

// SetEnableOpenUrlTab 设置启用打开新tab
//	与popup互斥
func (m *tCefChromiumConfig) SetEnableOpenUrlTab(value bool) *tCefChromiumConfig {
	m.enableOpenUrlTab = api.PascalBool(value)
	if value {
		m.SetEnableWindowPopup(false)
	}
	return m
}

func (m *tCefChromiumConfig) EnableOpenUrlTab() bool {
	return api.GoBool(m.enableOpenUrlTab)
}

// SetEnabledJavascript 设置启用Javascript
func (m *tCefChromiumConfig) SetEnabledJavascript(value bool) *tCefChromiumConfig {
	m.enabledJavascript = api.PascalBool(value)
	return m
}

func (m *tCefChromiumConfig) EnabledJavascript() bool {
	return api.GoBool(m.enabledJavascript)
}
