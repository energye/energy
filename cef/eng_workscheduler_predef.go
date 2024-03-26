//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/api"
)

var globalWorkScheduler ICEFWorkScheduler

// SetGlobalWorkScheduler
//
//	设置全局WorkScheduler实例, 该实例应在 NewCEFWorkScheduler 函数返回后设置
func SetGlobalWorkScheduler(scheduler ICEFWorkScheduler) {
	api.CEFPreDef().SysCallN(15, scheduler.Instance())
	globalWorkScheduler = scheduler
}

// DestroyGlobalCEFWorkScheduler
//
//	销毁全局WorkScheduler实例
func DestroyGlobalCEFWorkScheduler() {
	if globalWorkScheduler != nil {
		api.CEFPreDef().SysCallN(16)
		globalWorkScheduler.SetInstance(nil)
		globalWorkScheduler = nil
	}
}

// GlobalWorkSchedulerCreate
//
//	创建全局WorkScheduler实例
func GlobalWorkSchedulerCreate(aOwner IComponent) {
	if globalWorkScheduler == nil {
		scheduler := NewCEFWorkScheduler(aOwner)
		SetGlobalWorkScheduler(scheduler)
	}
}
