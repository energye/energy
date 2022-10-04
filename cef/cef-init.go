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

func SetOnException(exception ExceptionCallback) {
	exceptionCallback = exception
}

//初始化
func GlobalCEFInit(libs *embed.FS, resources *embed.FS) {
	inits.Init(libs, resources)
	if Args.IsRender() {
		netIpcPort := Args.Args(MAINARGS_NETIPCPORT)
		if netIpcPort != Empty {
			ipc.IPC.SetPort(int(StrToInt32(netIpcPort)))
		}
	}
	ipc.IPCChannelChooseInit()
	setMacOSXCommandLine(GoStrToDStr(Args.CommandLine()))
	applicationQueueAsyncCallInit()
	CommonInstanceInit()
	cefV8WindowBindFuncEventsInit()
	cefIPCInit()
	//应用低层出错异常捕获
	lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
		if exceptionCallback != nil {
			exceptionCallback(sender, e)
		} else {
			logger.Logger.Error("Exception:", e.Message())
		}
	})
}
