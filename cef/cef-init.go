//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"embed"
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
)

type ExceptionCallback func(sender lcl.IObject, e *lcl.Exception)

var exceptionCallback ExceptionCallback

func SetOnException(exception ExceptionCallback) {
	exceptionCallback = exception
}

//初始化
func GlobalCEFInit(libs *embed.FS, resources *embed.FS) {
	inits.Init(libs, resources)
	if Args.IsRender() {
		netIpcPort := Args.Args(MAINARGS_NETIPCPORT)
		if netIpcPort != empty {
			IPC.port = int(StrToInt32(netIpcPort))
		}
	}
	ipcChannelChooseInit()
	setMacOSXCommandLine(GoStrToDStr(Args.commandLine))
	applicationQueueAsyncCallInit()
	commonInstanceInit()
	cefV8WindowBindFuncEventsInit()
	cefIPCInit()
	//应用低层出错异常捕获
	lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
		if exceptionCallback != nil {
			exceptionCallback(sender, e)
		} else {
			Logger.Error("Exception:", e.Message())
		}
	})
}
