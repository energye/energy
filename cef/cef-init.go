//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"embed"
	. "github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
)

type ExceptionCallback func(sender lcl.IObject, e *lcl.Exception)

var exceptionCallback ExceptionCallback

//设置cef低层异常捕获回调函数
func SetOnException(exception ExceptionCallback) {
	exceptionCallback = exception
}

//全局初始化
//
//需要手动调用的函数,在main函数中调用
//
//libs 内置到应用程序的类库
//
//resources 内置到应用程序的资源文件
func GlobalCEFInit(libs *embed.FS, resources *embed.FS) {
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
	setMacOSXCommandLine(GoStrToDStr(Args.CommandLine()))
	//在主应用UI线程中执行，以队列异步方式调用的UI线程执行回调函数
	//applicationQueueAsyncCallInit()
	//通用对象初始化
	CommonInstanceInit()
	//对于go绑定到v8引擎js的事件处理函数
	cefV8WindowBindFuncEventsInit()
	//ipc初始化
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
