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

// TCefChromiumConfig 提供Chromium的基础快捷配置
type TCefChromiumConfig struct {
	enableMenu        bool //启用右键菜单
	enableViewSource  bool //启用查看源代码
	enableDevTools    bool //启用开发者工具
	enableWindowPopup bool //启用弹出新窗口
	enableOpenUrlTab  bool //启用tab签打开新窗口(需自定义实现)
	enabledJavascript bool //启用Javascript
}

// NewChromiumConfig 创建默认启用相关功能
func NewChromiumConfig() *TCefChromiumConfig {
	return &TCefChromiumConfig{
		enableMenu:        true,
		enableViewSource:  true,
		enableDevTools:    true,
		enableWindowPopup: true,
		enableOpenUrlTab:  false,
		enabledJavascript: false,
	}
}

// SetEnableMenu 设置启用右键菜单
func (m *TCefChromiumConfig) SetEnableMenu(value bool) *TCefChromiumConfig {
	m.enableMenu = value
	return m
}

func (m *TCefChromiumConfig) EnableMenu() bool {
	return m.enableMenu
}

// SetEnableViewSource 设置启用查看源文件
func (m *TCefChromiumConfig) SetEnableViewSource(value bool) *TCefChromiumConfig {
	m.enableViewSource = value
	return m
}

func (m *TCefChromiumConfig) EnableViewSource() bool {
	return m.enableViewSource
}

// SetEnableDevTools 设置启用开发者工具
func (m *TCefChromiumConfig) SetEnableDevTools(value bool) *TCefChromiumConfig {
	m.enableDevTools = value
	return m
}

func (m *TCefChromiumConfig) EnableDevTools() bool {
	return m.enableDevTools
}

// SetEnableWindowPopup 设置启用弹出新窗口
//
//	与tab互斥
func (m *TCefChromiumConfig) SetEnableWindowPopup(value bool) *TCefChromiumConfig {
	m.enableWindowPopup = value
	if value {
		m.SetEnableOpenUrlTab(false)
	}
	return m
}

func (m *TCefChromiumConfig) EnableWindowPopup() bool {
	return m.enableWindowPopup
}

// SetEnableOpenUrlTab 设置启用打开新tab
//
//	与popup互斥
func (m *TCefChromiumConfig) SetEnableOpenUrlTab(value bool) *TCefChromiumConfig {
	m.enableOpenUrlTab = value
	if value {
		m.SetEnableWindowPopup(false)
	}
	return m
}

func (m *TCefChromiumConfig) EnableOpenUrlTab() bool {
	return m.enableOpenUrlTab
}

// SetEnabledJavascript 设置启用Javascript
func (m *TCefChromiumConfig) SetEnabledJavascript(value bool) *TCefChromiumConfig {
	m.enabledJavascript = value
	return m
}

func (m *TCefChromiumConfig) EnabledJavascript() bool {
	return m.enabledJavascript
}
