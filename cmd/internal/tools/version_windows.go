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

package tools

import (
	"syscall"
	"unsafe"
)

func VersionNumber() (majorVersion, minorVersion, buildNumber uint32) {
	ntdll := syscall.NewLazyDLL("ntdll.dll")
	procRtlGetNtVersionNumbers := ntdll.NewProc("RtlGetNtVersionNumbers")
	procRtlGetNtVersionNumbers.Call(uintptr(unsafe.Pointer(&majorVersion)), uintptr(unsafe.Pointer(&minorVersion)), uintptr(unsafe.Pointer(&buildNumber)))
	buildNumber &= 0xffff
	return
}
