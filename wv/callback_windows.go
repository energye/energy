//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	"syscall"
)

var (
	eventCallback       = syscall.NewCallback(eventCallbackProc)
	removeEventCallback = syscall.NewCallback(removeEventCallbackProc)
)
