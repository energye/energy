//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Visual struct {
	Object
}

func AsVisual(ptr unsafe.Pointer) IVisual {
	if ptr == nil {
		return nil
	}
	m := new(Visual)
	m.instance = ptr
	return m
}
