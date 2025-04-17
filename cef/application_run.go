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
	"github.com/cyber-xxm/energy/v2/cef/internal/ipc"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

var (
	//BrowserWindow 是基于LCL和VF窗口组件的浏览器主窗口
	//  可以对窗口的属性设置和事件监听，Chromium的配置和事件监听.
	//  该窗口是主窗体，因此初始化时必须第一个初始化完成
	//  VF 窗口组件默认在Linux平台
	//  LCL窗口组件默认在Windows、MacOSX平台
	BrowserWindow = &browserWindow{
		browserEvent: &BrowserEvent{},
		Config: &BrowserConfig{
			WindowProperty: NewWindowProperty(),
		},
		windowInfo: make(map[int32]IBrowserWindow),
	}
	// disabledMainWindow 如果在Config配置中禁用主窗口(EnableMainWindow=false)时, 使用该窗口替代主窗口
	disabledMainWindow *disableMainWindow
	// LCL 主窗口实例
	enableMainWindow                 *lclBrowserWindow
	browserProcessStartAfterCallback browserProcessStartAfterCallbackFunc
)

func init() {
	if process.Args.IsMain() || process.Args.IsRender() {
		// 设置BrowserWindow到IPC实现
		ipc.SetBrowserWindow(BrowserWindow)
	}
}

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
//
//	在这里启动浏览器的主进程和子进程
func Run(app *TCEFApplication) {
	if application == nil {
		application = app
	}
	//MacOSX 多进程时，需要调用StartSubProcess来启动子进程
	if common.IsDarwin() && !application.SingleProcess() && !process.Args.IsMain() {
		// 启动子进程
		application.StartSubProcess()
		application.Free()
	} else {
		//externalMessagePump 和 multiThreadedMessageLoop 为 false 时, 启用 VF (ViewsFrameworkBrowserWindow) 窗口组件
		if application.IsMessageLoop() {
			// 启用VFMessageLoop
			// 初始化窗口组件
			appContextInitialized()
		}
		// Linux Gtk3
		widgetInit := common.IsLinux() && api.WidgetUI().IsGTK3()
		if widgetInit {
			// linux widget init
			// TODO:
			//  Linux Gtk3 自定义初始化组件库，实际应放在 StartMainProcess 成功之后再执行，但不行，不能输入英文字符。
			//  放在 StartMainProcess 之前执行又会有警告。先加在这里。
			//  初始化组件库，主要为了在Gtk3时能使用部分LCL组件库的功能。
			api.CustomWidgetSetInitialization()
		}
		// 窗口消息方式是true(VF)时初始化一次Application，以使用一些LCL功能, 例如UI线程执行函数
		if application.IsMessageLoop() {
			lcl.Application.Initialize()
		}
		// 启动主进程
		success := application.StartMainProcess()
		if success {
			api.SetReleaseCallback(func() {
				if widgetInit {
					// linux widget free
					api.CustomWidgetSetFinalization()
				}
				app.Destroy()
				app.Free()
			})
			// 主进程启动成功之后回调
			if browserProcessStartAfterCallback != nil {
				browserProcessStartAfterCallback(success)
			}
			if application.IsMessageLoop() {
				// VF窗口 MessageLoop
				application.RunMessageLoop()
			} else {
				// LCL窗口 创建并运行应用
				BrowserWindow.createFormAndRun()
				//lclwidget.CustomWidgetSetFinalization()
			}
		}
	}
}
