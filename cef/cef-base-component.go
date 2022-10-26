//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type BaseComponent struct {
	lcl.TComponent
	procName string
	instance uintptr
	ptr      unsafe.Pointer
}

func (m *BaseComponent) Instance() uintptr {
	return m.instance
}

func (m *BaseComponent) Handle() types.HWND {
	return GetHandle(m.procName, m.instance)
}

func (m *BaseComponent) IsValid() bool {
	return m.instance != 0
}
