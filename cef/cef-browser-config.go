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
	"github.com/energye/energy/common"
)

type browserWindowOnEventCallback func(event *BrowserEvent, window IBrowserWindow)
type browserWindowAfterOnEventCallback func(window IBrowserWindow)

// 创建主窗口指定的一些快捷配置属性
type browserConfig struct {
	WindowProperty
	chromiumConfig                    *tCefChromiumConfig               //主窗体浏览器配置
	browserWindowOnEventCallback      browserWindowOnEventCallback      //主窗口初始化回调 - 基于LCL窗口
	browserWindowAfterOnEventCallback browserWindowAfterOnEventCallback //主窗口初始化之后回调
}

// 设置chromium配置
func (m *browserConfig) SetChromiumConfig(chromiumConfig *tCefChromiumConfig) {
	if chromiumConfig != nil && common.Args.IsMain() {
		m.chromiumConfig = chromiumConfig
	}
}

// 获取/创建 CEF Chromium Config
func (m *browserConfig) ChromiumConfig() *tCefChromiumConfig {
	if m.chromiumConfig == nil {
		m.chromiumConfig = NewChromiumConfig()
	}
	return m.chromiumConfig
}

// 主窗口初始化回调
//
// 该回调函数和基于CEF窗口回调是互斥的，默认情况只有一个会被回调
func (m *browserConfig) setBrowserWindowInitOnEvent(fn browserWindowOnEventCallback) {
	if fn != nil && common.Args.IsMain() {
		m.browserWindowOnEventCallback = fn
	}
}

// 主窗口初始化回调 - 基于LCL窗口
func (m *browserConfig) setBrowserWindowInitAfterOnEvent(fn browserWindowAfterOnEventCallback) {
	if fn != nil && common.Args.IsMain() {
		m.browserWindowAfterOnEventCallback = fn
	}
}
