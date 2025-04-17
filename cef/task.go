//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	. "github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ITask Parent: TCefTask
//
//	Implement this interface for asynchronous task execution. If the task is
//	posted successfully and if the associated message loop is still running then
//	the execute() function will be called on the target thread. If the task
//	fails to post then the task object may be destroyed on the source thread
//	instead of the target thread. For this reason be cautious when performing
//	work in the task object destructor.
//	<a cref="uCEFTypes|TCefTask">Implements TCefTask</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_task_capi.h">CEF source file: /include/capi/cef_task_capi.h (cef_task_t)</a>
type ITask struct {
	base       TCefBaseRefCounted
	instance   unsafe.Pointer
	executePtr uintptr
}

func NewTask() *ITask {
	r1, _, _ := imports.Proc(def.Task_Create).Call()
	return &ITask{
		instance: unsafe.Pointer(r1),
	}
}

func (m *ITask) OnlyPostTask(threadId TCefThreadId) bool {
	r1, _, _ := imports.Proc(def.Task_OnlyPostTask).Call(m.Instance(), uintptr(threadId))
	return GoBool(r1)
}

func (m *ITask) OnlyPostDelayedTask(threadId TCefThreadId, delay int64) bool {
	r1, _, _ := imports.Proc(def.Task_OnlyPostDelayedTask).Call(m.Instance(), uintptr(threadId), uintptr(unsafe.Pointer(&delay)))
	return GoBool(r1)
}

func (m *ITask) SetOnExecute(fn TTaskExecute) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	imports.Proc(def.Task_SetOnExecute).Call(m.Instance(), m.executePtr)
}

func (m *ITask) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ITask) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

type TTaskExecute func()

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case TTaskExecute:
			fn.(TTaskExecute)()
		default:
			return false
		}
		return true
	})
}
