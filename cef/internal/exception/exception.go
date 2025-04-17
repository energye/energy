//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package exception
//
//	Underlying dynamic link library exception capture
//	Supports: Windows, MacOS.
package exception

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
)

type Callback func(message string)

var (
	exceptionHandler         dllimports.ProcAddr
	exceptionHandlerCallback Callback
)

// HandlerInit
// 底层库异常处理器初始化
func HandlerInit(fn Callback) {
	if exceptionHandlerCallback == nil {
		exceptionHandler = imports.Proc(def.SetExceptionHandlerCallback)
		exceptionHandler.Call(exceptionHandlerProcEventAddr)
		exceptionHandlerCallback = fn
	}
}

func exceptionHandlerProc(message uintptr) uintptr {
	if exceptionHandlerCallback != nil {
		exceptionHandlerCallback(api.GoStr(message))
	}
	return 0
}
