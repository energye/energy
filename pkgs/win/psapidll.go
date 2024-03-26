//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package win

import (
	. "github.com/energye/energy/v2/types"
	"unsafe"
)

import "syscall"

var (
	psapidll = syscall.NewLazyDLL("psapi.dll")

	_GetModuleFileNameEx = psapidll.NewProc("GetModuleFileNameExW")
)

func GetModuleFileNameEx(hProcess uintptr, hModule HMODULE) (string, uint32) {
	buff := make([]uint16, MAX_PATH)
	r, _, _ := _GetModuleFileNameEx.Call(hProcess, uintptr(hModule), uintptr(unsafe.Pointer(&buff[0])), uintptr(MAX_PATH*2))
	return GoStr(buff), uint32(r)
}
