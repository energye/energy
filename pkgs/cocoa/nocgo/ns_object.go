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
	"github.com/ebitengine/purego/objc"
	"unsafe"
)

type NSObject struct {
	instance unsafe.Pointer
}

func (m *NSObject) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *NSObject) SetInstance(ptr unsafe.Pointer) {
	m.instance = ptr
}

func (m *NSObject) Self() objc.ID {
	return objc.ID(m.instance)
}
