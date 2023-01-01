//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common"
)

type viewsFrameBrowserWindowOnEventCallback func(event *BrowserEvent, window *ViewsFrameworkBrowserWindow)
type browserWindowOnEventCallback func(event *BrowserEvent, window *TCefWindowInfo)
type browserWindowAfterOnEventCallback func(window *TCefWindowInfo)

//创建主窗口指定的一些快捷配置属性
type browserConfig struct {
	DefaultUrl                             string                                 //默认URL地址
	Title                                  string                                 //窗口标题
	Icon                                   string                                 //窗口图标
	Width                                  int32                                  //窗口宽
	Height                                 int32                                  //窗口高
	chromiumConfig                         *tCefChromiumConfig                    //主窗体浏览器配置
	viewsFrameBrowserWindowOnEventCallback viewsFrameBrowserWindowOnEventCallback //主窗口初始化回调 - 基于CEF views framework窗口
	browserWindowOnEventCallback           browserWindowOnEventCallback           //主窗口初始化回调 - 基于LCL窗口
	browserWindowAfterOnEventCallback      browserWindowAfterOnEventCallback      //主窗口初始化之后回调
}

//设置chromium配置
func (m *browserConfig) SetChromiumConfig(chromiumConfig *tCefChromiumConfig) {
	if chromiumConfig != nil && common.Args.IsMain() {
		m.chromiumConfig = chromiumConfig
	}
}

//主窗口初始化回调 - 基于CEF views framework窗口
//
//该回调函数和基于LCL窗口回调是互斥的，默认情况只有一个会被回调
func (m *browserConfig) setViewsFrameBrowserWindowOnEventCallback(fn viewsFrameBrowserWindowOnEventCallback) {
	if fn != nil && common.Args.IsMain() {
		m.viewsFrameBrowserWindowOnEventCallback = fn
	}
}

//主窗口初始化回调 - 基于LCL窗口
//
//该回调函数和基于CEF窗口回调是互斥的，默认情况只有一个会被回调
func (m *browserConfig) setBrowserWindowInitOnEvent(fn browserWindowOnEventCallback) {
	if fn != nil && common.Args.IsMain() {
		m.browserWindowOnEventCallback = fn
	}
}

//主窗口初始化回调 - 基于LCL窗口
func (m *browserConfig) setBrowserWindowInitAfterOnEvent(fn browserWindowAfterOnEventCallback) {
	if fn != nil && common.Args.IsMain() {
		m.browserWindowAfterOnEventCallback = fn
	}
}
