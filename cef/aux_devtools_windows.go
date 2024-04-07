//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

// 辅助工具-开发者工具 windows

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// 创建开发者工具窗口，在 UI 线程中创建
// 在显示开发者工具时，需要在显示的窗口中初始化 CEFWindow
// 步骤
//  1. 创建窗口, 设置窗口宽高为0, 展示位置在屏幕之外
//  2. 创建 CEFWindow
//  3. show & hide, 先显示窗口让CEF初始化CEFWindow, 紧跟着隐藏掉
//  4. 设置默认的窗口宽高、居中显示在桌面并显示在任务栏
func createDevtoolsWindow(owner *LCLBrowserWindow, currentBrowser ICEFChromiumBrowser) *devToolsWindow {
	window := &devToolsWindow{}
	window.TForm = lcl.NewForm(owner)
	window.SetCaption(devToolsName)
	window.SetIcon(owner.Icon())
	window.SetWindowParent(NewCEFWindowParent(window))
	window.WindowParent().SetParent(window)
	window.WindowParent().SetAlign(types.AlClient)
	window.SetWidth(1024)
	window.SetHeight(768)
	window.ScreenCenter()
	window.SetShowInTaskBar(types.StAlways)
	window.SetVisible(false)
	window.TForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		*action = types.CaFree                                         // 释放掉窗口
		currentBrowser.Chromium().CloseDevTools(window.WindowParent()) // close devtools
		// 关闭窗口后
		// 将父窗口开发者工具(当前自己)设置为空, 以便在下次正确创建
		owner.GetAuxTools().SetDevTools(nil)
	})
	window.TForm.SetOnShow(func(sender lcl.IObject) {
		// 显示开发者工具
		currentBrowser.Chromium().ShowDevTools(window.WindowParent())
	})
	return window
}
