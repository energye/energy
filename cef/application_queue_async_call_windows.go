//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows && cgo
// +build windows,cgo

package cef

import "syscall"

var (
	applicationQueueAsyncCallEvent = syscall.NewCallback(applicationQueueAsyncCallProc)
)
