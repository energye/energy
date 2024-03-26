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
//  Underlying dynamic link library exception capture
//  Supports: Windows, MacOS.
package exception

import (
	"github.com/energye/energy/v2/api"
)

type callback func(message string)

var exceptionHandlerCallback callback

// SetOnException
// 底层库异常
func SetOnException(fn callback) {
	if exceptionHandlerCallback == nil {
		api.SetExceptionHandlerCallback(exceptionHandlerProcEventAddr)
		exceptionHandlerCallback = fn
	}
}

func exceptionHandlerProc(message uintptr) uintptr {
	if exceptionHandlerCallback != nil {
		exceptionHandlerCallback(api.GoStr(message))
	}
	return 0
}
