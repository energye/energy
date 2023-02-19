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
	"embed"
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/pkgs/macapp"
)

type ExceptionCallback func(sender lcl.IObject, e *lcl.Exception)

var exceptionCallback ExceptionCallback

// 设置cef低层异常捕获回调函数
func SetOnException(exception ExceptionCallback) {
	exceptionCallback = exception
}

// 全局初始化
//
// 需要手动调用的函数,在main函数中调用
//
// libs 内置到应用程序的类库
//
// resources 内置到应用程序的资源文件
//
// MacOSX环境, ide开发环境需命令行参数[energy_env=dev]以保证应用正常运行
func GlobalInit(libs *embed.FS, resources *embed.FS) {
	macapp.MacApp.IsCEF(IsDarwin())
	//MacOSX环境, ide开发环境需命令行参数[energy_env=dev]以保证应用正常运行
	energyEnv := Args.Args("energy_env")
	if energyEnv != "" {
		macapp.MacApp.SetEnergyEnv(macapp.ENERGY_ENV(energyEnv))
	}
	inits.Init(libs, resources)
	if Args.IsRender() {
		netIpcPort := Args.Args(MAINARGS_NETIPCPORT)
		if netIpcPort != Empty {
			ipc.IPC.SetPort(int(StrToInt32(netIpcPort)))
		}
	}
	//IPC通道选择初始化, 在不支持unix的系统中将选择net socket
	ipc.IPCChannelChooseInit()
	//macos的命令行设置
	setMacOSXCommandLine(api.PascalStr(Args.CommandLine()))
	applicationQueueAsyncCallInit()
	//对于go绑定到v8引擎js的事件处理函数
	cefV8WindowBindFuncEventsInit()
	//ipc事件初始化
	cefIPCInit()
	//应用低层出错异常捕获
	lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
		if exceptionCallback != nil {
			exceptionCallback(sender, e)
		} else {
			logger.Error("Exception:", e.Message())
		}
	})
}
