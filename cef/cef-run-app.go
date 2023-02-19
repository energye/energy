//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

var (
	//BrowserWindow 是基于LCL和VF窗口组件的浏览器主窗口
	//
	//可以对窗口的属性设置和事件监听，Chromium的配置和事件监听.
	//
	//该窗口是主窗体，因此初始化时必须第一个初始化完成，如果创建子窗口最好在 SetBrowserInitAfter 回调函数中创建
	//
	//VF窗口组件默认在Linux平台自动开启
	//LCL窗口组件默认在Windows、MacOSX平台自动开启
	BrowserWindow = &browser{
		browserEvent: &BrowserEvent{},
		Config: &browserConfig{
			WindowProperty: NewWindowProperty(),
		},
		windowInfo:   make(map[int32]IBrowserWindow),
		windowSerial: 1, //默认1开始
	}
	browserProcessStartAfterCallback browserProcessStartAfterCallbackFunc
)

type browserProcessStartAfterCallbackFunc func(success bool)

// SetBrowserProcessStartAfterCallback 主进程启动之后回调函数
func SetBrowserProcessStartAfterCallback(callback browserProcessStartAfterCallbackFunc) {
	if common.Args.IsMain() {
		if browserProcessStartAfterCallback == nil {
			browserProcessStartAfterCallback = callback
		}
	}
}

// Run
// 运行应用
//
// 多进程方式，启动主进程然后启动子进程，在MacOS下，需要单独调用启动子进程函数，单进程只启动主进程
//
// 主进程启动成功之后，将创建主窗口 mainBrowserWindow 是一个默认的主窗口
//
// externalMessagePump和multiThreadedMessageLoop是false时启用 ViewsFrameworkBrowserWindow 窗口
//
// 在这里启动浏览器的主进程和子进程
func Run(cefApp *TCEFApplication) {
	defer func() {
		logger.Debug("application process [", common.Args.ProcessType(), "] run end")
		api.EnergyLibRelease()
	}()
	if common.IsDarwin() && !consts.SingleProcess && !common.Args.IsMain() {
		// mac os 启动子进程
		cefApp.StartSubProcess()
		cefApp.Free()
	} else {
		//externalMessagePump 和 multiThreadedMessageLoop 为 false 时启用CEF views framework (ViewsFrameworkBrowserWindow) 窗口
		consts.IsMessageLoop = !api.GoBool(cefApp.cfg.externalMessagePump) && !api.GoBool(cefApp.cfg.multiThreadedMessageLoop)
		if consts.IsMessageLoop {
			BrowserWindow.appContextInitialized(cefApp)
		}
		success := cefApp.StartMainProcess()
		if browserProcessStartAfterCallback != nil {
			browserProcessStartAfterCallback(success)
		}
		if success {
			internalBrowserIPCOnEventInit()
			ipc.IPC.StartBrowserIPC()
			bindGoToJS(nil, nil)
			if consts.IsMessageLoop {
				cefApp.RunMessageLoop()
			} else {
				if BrowserWindow.mainBrowserWindow == nil {
					BrowserWindow.mainBrowserWindow = &browserWindow{}
				}
				lcl.RunApp(&BrowserWindow.mainBrowserWindow)
			}
		}
	}
}
