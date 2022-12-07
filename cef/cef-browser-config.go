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

//创建主窗口指定的一些快捷配置属性
type browserConfig struct {
	DefaultUrl                        string                                                   //默认URL地址
	Title                             string                                                   //窗口标题
	Icon                              string                                                   //窗口图标
	Width                             int32                                                    //窗口宽
	Height                            int32                                                    //窗口高
	chromiumConfig                    *tCefChromiumConfig                                      //主窗体浏览器配置
	browserWindowOnEventCallback      func(browserEvent *BrowserEvent, window *TCefWindowInfo) //主窗口初始化回调
	browserWindowAfterOnEventCallback func(window *TCefWindowInfo)                             //主窗口初始化之后回调
}

//设置chromium配置
func (m *browserConfig) SetChromiumConfig(chromiumConfig *tCefChromiumConfig) {
	if chromiumConfig != nil && common.Args.IsMain() {
		m.chromiumConfig = chromiumConfig
	}
}

func (m *browserConfig) setBrowserWindowInitOnEvent(fn func(event *BrowserEvent, browserWindow *TCefWindowInfo)) {
	if fn != nil && common.Args.IsMain() {
		m.browserWindowOnEventCallback = fn
	}
}

func (m *browserConfig) setBrowserWindowInitAfterOnEvent(fn func(browserWindow *TCefWindowInfo)) {
	if fn != nil && common.Args.IsMain() {
		m.browserWindowAfterOnEventCallback = fn
	}
}
