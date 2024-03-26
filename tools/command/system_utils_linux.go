//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux
// +build linux

package command

import "syscall"

func HideWindow(bool bool) *syscall.SysProcAttr {
	return nil
}
