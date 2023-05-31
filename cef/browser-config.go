//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Browser Window 配置

package cef

import (
	"github.com/energye/energy/v2/cef/process"
)

type browserWindowOnEventCallback func(event *BrowserEvent, window IBrowserWindow)

// 创建主窗口指定的一些快捷配置属性
type browserConfig struct {
	WindowProperty
	chromiumConfig               *tCefChromiumConfig          //主窗体浏览器配置
	browserWindowOnEventCallback browserWindowOnEventCallback //主窗口初始化回调
}

// SetChromiumConfig 设置chromium配置
func (m *browserConfig) SetChromiumConfig(chromiumConfig *tCefChromiumConfig) {
	if chromiumConfig != nil && process.Args.IsMain() {
		m.chromiumConfig = chromiumConfig
	}
}

// ChromiumConfig 扩展配置
//  获取/创建 CEF Chromium Config
func (m *browserConfig) ChromiumConfig() *tCefChromiumConfig {
	if m.chromiumConfig == nil {
		m.chromiumConfig = NewChromiumConfig()
	}
	return m.chromiumConfig
}

// 主窗口初始化回调
//	创建主窗口后,显示之前执行
func (m *browserConfig) setBrowserWindowInitOnEvent(fn browserWindowOnEventCallback) {
	if fn != nil && process.Args.IsMain() {
		m.browserWindowOnEventCallback = fn
	}
}
