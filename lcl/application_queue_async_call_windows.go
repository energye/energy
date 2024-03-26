//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package lcl

import "syscall"

var (
	threadAsyncCallback = syscall.NewCallback(threadAsyncCallbackProc)
)
