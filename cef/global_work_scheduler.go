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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

var GlobalWorkScheduler *TCEFWorkScheduler

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

func GlobalWorkSchedulerCreate(owner lcl.IComponent) *TCEFWorkScheduler {
	if GlobalWorkScheduler != nil {
		return GlobalWorkScheduler
	}
	var aOwner uintptr
	if owner != nil {
		aOwner = owner.Instance()
	}
	var result uintptr
	imports.Proc(def.CEFWorkScheduler_Create).Call(aOwner, uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		GlobalWorkScheduler = &TCEFWorkScheduler{instance: unsafe.Pointer(result)}
	}
	return GlobalWorkScheduler
}

func GlobalWorkSchedulerCreateDelayed() *TCEFWorkScheduler {
	if GlobalWorkScheduler != nil {
		return GlobalWorkScheduler
	}
	var result uintptr
	imports.Proc(def.CEFWorkScheduler_CreateDelayed).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		GlobalWorkScheduler = &TCEFWorkScheduler{instance: unsafe.Pointer(result)}
	}
	return GlobalWorkScheduler
}

func (m *TCEFWorkScheduler) StopScheduler() {
	imports.Proc(def.CEFWorkScheduler_StopScheduler).Call()
}

func (m *TCEFWorkScheduler) CreateThread() {
	imports.Proc(def.CEFWorkScheduler_CreateThread).Call()
}

func (m *TCEFWorkScheduler) Destroy() {
	imports.Proc(def.CEFWorkScheduler_Destroy).Call()
}

// GetPriority Windows
func (m *TCEFWorkScheduler) GetPriority() consts.TThreadPriority {
	r := imports.SysCallN(def.CEFWorkScheduler_GetPriority)
	return consts.TThreadPriority(r)
}

// GetPriority Windows
func (m *TCEFWorkScheduler) SetPriority(value consts.TThreadPriority) {
	imports.SysCallN(def.CEFWorkScheduler_SetPriority, uintptr(value))
}

func (m *TCEFWorkScheduler) GetDefaultInterval() int32 {
	r := imports.SysCallN(def.CEFWorkScheduler_GetDefaultInterval)
	return int32(r)
}

func (m *TCEFWorkScheduler) GetDepleteWorkCycles() uint32 {
	r := imports.SysCallN(def.CEFWorkScheduler_GetDepleteWorkCycles)
	return uint32(r)
}

func (m *TCEFWorkScheduler) GetDepleteWorkDelay() uint32 {
	r := imports.SysCallN(def.CEFWorkScheduler_GetDepleteWorkDelay)
	return uint32(r)
}

func (m *TCEFWorkScheduler) GetUseQueueThread() bool {
	r := imports.SysCallN(def.CEFWorkScheduler_GetUseQueueThread)
	return api.GoBool(r)
}

func (m *TCEFWorkScheduler) SetDefaultInterval(value int32) {
	imports.SysCallN(def.CEFWorkScheduler_SetDefaultInterval, uintptr(value))
}

func (m *TCEFWorkScheduler) SetDepleteWorkCycles(value uint32) {
	imports.SysCallN(def.CEFWorkScheduler_SetDepleteWorkCycles, uintptr(value))
}

func (m *TCEFWorkScheduler) SetDepleteWorkDelay(value uint32) {
	imports.SysCallN(def.CEFWorkScheduler_SetDepleteWorkDelay, uintptr(value))
}

func (m *TCEFWorkScheduler) SetUseQueueThread(value bool) {
	imports.SysCallN(def.CEFWorkScheduler_SetUseQueueThread, api.PascalBool(value))
}
