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
	"unsafe"
)

type BaseWinControl struct {
	lcl.IWinControl
	instance unsafe.Pointer
}

func (m *BaseWinControl) IsValid() bool {
	return m.instance != nullptr
}

//Instance 当前实例
func (m *BaseWinControl) Instance() unsafe.Pointer {
	return m.instance
}
