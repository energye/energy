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
)

import "syscall"

const (
	ole32 = "ole32.dll"
)

var (
	ole32dll = syscall.NewLazyDLL(ole32)

	_CoInitialize   = ole32dll.NewProc("CoInitialize")
	_CoInitializeEx = ole32dll.NewProc("CoInitializeEx")
	_CoUninitialize = ole32dll.NewProc("CoUninitialize")
)

func CoInitialize(pvReserved uintptr) HRESULT {
	r, _, _ := _CoInitialize.Call(pvReserved)
	return HRESULT(r)
}

func CoInitializeEx(pvReserved uintptr, coInit uint32) HRESULT {
	r, _, _ := _CoInitializeEx.Call(pvReserved, uintptr(coInit))
	return HRESULT(r)
}

func CoUninitialize() {
	_CoUninitialize.Call()
}
