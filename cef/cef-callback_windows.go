//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import "syscall"

var (
	cefIPCCallbackFuncEvent = syscall.NewCallback(cefIPCEventProc)
	cefWindowBindEvent      = syscall.NewCallback(cefWindowBindCallbackEventProc)
)
