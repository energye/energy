//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Window Parent 父组件

package cef

import (
	"github.com/energye/golcl/lcl"
	"unsafe"
)

// BaseWinControl TCEFWindowParent 和 TCEFLinkedWindowParent 父结构体
type BaseWinControl struct {
	lcl.IWinControl
	instance unsafe.Pointer
}

// IsValid 是否有效
func (m *BaseWinControl) IsValid() bool {
	return m != nil && m.instance != nil
}

// Instance 当前实例
func (m *BaseWinControl) Instance() uintptr {
	return uintptr(m.instance)
}
