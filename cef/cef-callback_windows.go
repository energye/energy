//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "syscall"

var (
	cefIPCCallbackFuncEvent        = syscall.NewCallback(cefIPCEventProc)
	cefWindowBindEvent             = syscall.NewCallback(cefWindowBindCallbackEventProc)
	applicationQueueAsyncCallEvent = syscall.NewCallback(applicationQueueAsyncCallProc)
)
