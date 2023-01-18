//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

import "github.com/energye/golcl/lcl/types"

//显示标题栏
func (m *LCLBrowserWindow) ShowTitle() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetBorderStyle(types.BsSizeable)
}

//隐藏标题栏
func (m *LCLBrowserWindow) HideTitle() {
	if m.TForm == nil {
		return
	}
	m.TForm.SetBorderStyle(types.BsSingle)
}

// 默认事件注册 windows 消息事件
func (m *LCLBrowserWindow) registerWindowsCompMsgEvent() {

}
