//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

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

// New CefChromiumConfig 默认禁用相关功能
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

// 设置启用右键菜单
func (m *tCefChromiumConfig) SetEnableMenu(value bool) *tCefChromiumConfig {
	m.enableMenu = api.PascalBool(value)
	return m
}

// 设置启用查询源文件
func (m *tCefChromiumConfig) SetEnableViewSource(value bool) *tCefChromiumConfig {
	m.enableViewSource = api.PascalBool(value)
	return m
}

// 设置启用开发者工具
func (m *tCefChromiumConfig) SetEnableDevTools(value bool) *tCefChromiumConfig {
	m.enableDevTools = api.PascalBool(value)
	return m
}

// 设置启用打开链接弹出新窗口
func (m *tCefChromiumConfig) SetEnableWindowPopup(value bool) *tCefChromiumConfig {
	m.enableWindowPopup = api.PascalBool(value)
	return m
}

// 设置启用打开链接打开新tab
func (m *tCefChromiumConfig) SetEnableOpenUrlTab(value bool) *tCefChromiumConfig {
	m.enableOpenUrlTab = api.PascalBool(value)
	return m
}

// 设置启用Javascript
func (m *tCefChromiumConfig) SetEnabledJavascript(value bool) *tCefChromiumConfig {
	m.enabledJavascript = api.PascalBool(value)
	return m
}
