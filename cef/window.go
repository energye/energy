//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import "github.com/energye/golcl/lcl"

// 自定义窗口组件
type Window struct {
	BaseWindow
	defaultUrl string
	config     *tCefChromiumConfig
}

//创建一个新window窗口
func NewWindow() *Window {
	var window = &Window{}
	//window.TForm = lcl.NewForm(owner)
	lcl.Application.CreateForm(&window)
	window.ParentDoubleBuffered()
	window.FormCreate()
	window.SetNotInTaskBar()
	window.defaultWindowEvent()
	return window
}

//返回完整的chromium对象
func (m *Window) Chromium() IChromium {
	return m.chromium
}

//启用默认关闭事件行为-该窗口将被关闭
func (m *Window) EnableDefaultClose() {
	m.defaultWindowCloseEvent()
	m.registerDefaultChromiumCloseEvent()
}

//启用所有默认事件行为
func (m *Window) EnableAllDefaultEvent() {
	m.defaultWindowCloseEvent()
	m.defaultChromiumEvent()
}
