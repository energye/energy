//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"unsafe"
)

type TCEFWorkScheduler struct {
	instance unsafe.Pointer
}

// Instance 实例
func (m *TCEFWorkScheduler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCEFWorkScheduler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func GlobalWorkSchedulerStop() {
	imports.Proc(def.CEFWorkScheduler_Stop).Call()
}

func GlobalWorkSchedulerCreate(owner lcl.IComponent) *TCEFWorkScheduler {
	var aOwner uintptr
	if owner != nil {
		aOwner = owner.Instance()
	}
	var result uintptr
	imports.Proc(def.CEFWorkScheduler_Create).Call(aOwner, uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFWorkScheduler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func GlobalWorkSchedulerDestroy() {
	imports.Proc(def.CEFWorkScheduler_Destroy).Call()
}
