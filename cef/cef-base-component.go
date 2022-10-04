//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"unsafe"
)

type BaseComponent struct {
	lcl.TComponent
	procName string
	instance uintptr
	ptr      unsafe.Pointer
}
