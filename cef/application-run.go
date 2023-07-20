//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy application run

package cef

import (
	"github.com/energye/energy/v2/cef/lclwidget"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/common"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

var (
	//BrowserWindow 是基于LCL和VF窗口组件的浏览器主窗口
	//  可以对窗口的属性设置和事件监听，Chromium的配置和事件监听.
	//  该窗口是主窗体，因此初始化时必须第一个初始化完成，如果创建子窗口最好在 SetBrowserInitAfter 回调函数中创建
	//  VF 窗口组件默认在Linux平台
	//  LCL窗口组件默认在Windows、MacOSX平台
	BrowserWindow = &browserWindow{
		browserEvent: &BrowserEvent{},
		Config: &browserConfig{
			WindowProperty: NewWindowProperty(),
		},
		windowInfo: make(map[int32]IBrowserWindow),
	}
	browserProcessStartAfterCallback browserProcessStartAfterCallbackFunc
)

type browserProcessStartAfterCallbackFunc func(success bool)

// SetBrowserProcessStartAfterCallback 主进程启动之后回调函数
func SetBrowserProcessStartAfterCallback(callback browserProcessStartAfterCallbackFunc) {
	if process.Args.IsMain() {
		if browserProcessStartAfterCallback == nil {
			browserProcessStartAfterCallback = callback
		}
	}
}

// Run 运行应用
//  在这里启动浏览器的主进程和子进程
func Run(app *TCEFApplication) {
	defer func() {
		api.EnergyLibRelease()
	}()
	//MacOSX 多进程时，需要调用StartSubProcess来启动子进程
	if common.IsDarwin() && !app.SingleProcess() && !process.Args.IsMain() {
		// 启动子进程
		app.StartSubProcess()
		app.Free()
	} else {
		//externalMessagePump 和 multiThreadedMessageLoop 为 false 时, 启用 VF (ViewsFrameworkBrowserWindow) 窗口组件
		if app.IsMessageLoop() {
			// 启用VF窗口组件
			BrowserWindow.appContextInitialized(app)
		}
		// 启动主进程
		success := app.StartMainProcess()
		if success {
			//LCL -> Linux 必须在主进程启动之后初始化组件
			lclwidget.CustomWidgetSetInitialization()
			// 主进程启动成功之后回调
			if browserProcessStartAfterCallback != nil {
				browserProcessStartAfterCallback(success)
			}
			appMainRunCallback()
			if app.IsMessageLoop() {
				lcl.Application.Initialize()
				// VF窗口
				app.RunMessageLoop()
			} else {
				// 创建LCL窗口组件
				if BrowserWindow.mainBrowserWindow == nil {
					BrowserWindow.mainBrowserWindow = new(lclBrowserWindow)
				}
				// LCL窗口
				lcl.RunApp(&BrowserWindow.mainBrowserWindow)
				lclwidget.CustomWidgetSetFinalization()
			}
		}
	}
}
