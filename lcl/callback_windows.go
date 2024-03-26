//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"syscall"
)

var (
	eventCallback                   = syscall.NewCallback(eventCallbackProc)
	removeEventCallback             = syscall.NewCallback(removeEventCallbackProc)
	messageCallback                 = syscall.NewCallback(messageCallbackProc)
	threadSyncCallback              = syscall.NewCallback(threadSyncCallbackProc)
	requestCallCreateParamsCallback = syscall.NewCallback(requestCallCreateParamsCallbackProc)
)
