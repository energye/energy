//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefTaskRunner Parent: ICefBaseRefCounted
//
//	Interface that asynchronously executes tasks on the associated thread. It is safe to call the functions of this interface on any thread. CEF maintains multiple internal threads that are used for handling different types of tasks in different processes. The TCefThreadId definitions in cef_types.h list the common CEF threads. Task runners are also available for other CEF threads as appropriate (for example, V8 WebWorker threads).
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_task_capi.h">CEF source file: /include/capi/cef_task_capi.h (cef_task_runner_t))</a>
type ICefTaskRunner interface {
	ICefBaseRefCounted
	// IsSame
	//  Returns true (1) if this object is pointing to the same task runner as |that| object.
	IsSame(that ICefTaskRunner) bool // function
	// BelongsToCurrentThread
	//  Returns true (1) if this task runner belongs to the current thread.
	BelongsToCurrentThread() bool // function
	// BelongsToThread
	//  Returns true (1) if this task runner is for the specified CEF thread.
	BelongsToThread(threadId TCefThreadId) bool // function
	// PostTask
	//  Post a task for execution on the thread associated with this task runner. Execution will occur asynchronously.
	PostTask(task ICefTask) bool // function
	// PostDelayedTask
	//  Post a task for delayed execution on the thread associated with this task runner. Execution will occur asynchronously. Delayed tasks are not supported on V8 WebWorker threads and will be executed without the specified delay.
	PostDelayedTask(task ICefTask, delayMs int64) bool // function
}

// TCefTaskRunner Parent: TCefBaseRefCounted
//
//	Interface that asynchronously executes tasks on the associated thread. It is safe to call the functions of this interface on any thread. CEF maintains multiple internal threads that are used for handling different types of tasks in different processes. The TCefThreadId definitions in cef_types.h list the common CEF threads. Task runners are also available for other CEF threads as appropriate (for example, V8 WebWorker threads).
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_task_capi.h">CEF source file: /include/capi/cef_task_capi.h (cef_task_runner_t))</a>
type TCefTaskRunner struct {
	TCefBaseRefCounted
}

// TaskRunnerRef -> ICefTaskRunner
var TaskRunnerRef taskRunner

// taskRunner TCefTaskRunner Ref
type taskRunner uintptr

func (m *taskRunner) UnWrap(data uintptr) ICefTaskRunner {
	var resultCefTaskRunner uintptr
	CEF().SysCallN(1421, uintptr(data), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *taskRunner) GetForCurrentThread() ICefTaskRunner {
	var resultCefTaskRunner uintptr
	CEF().SysCallN(1416, uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *taskRunner) GetForThread(threadId TCefThreadId) ICefTaskRunner {
	var resultCefTaskRunner uintptr
	CEF().SysCallN(1417, uintptr(threadId), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *TCefTaskRunner) IsSame(that ICefTaskRunner) bool {
	r1 := CEF().SysCallN(1418, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefTaskRunner) BelongsToCurrentThread() bool {
	r1 := CEF().SysCallN(1414, m.Instance())
	return GoBool(r1)
}

func (m *TCefTaskRunner) BelongsToThread(threadId TCefThreadId) bool {
	r1 := CEF().SysCallN(1415, m.Instance(), uintptr(threadId))
	return GoBool(r1)
}

func (m *TCefTaskRunner) PostTask(task ICefTask) bool {
	r1 := CEF().SysCallN(1420, m.Instance(), GetObjectUintptr(task))
	return GoBool(r1)
}

func (m *TCefTaskRunner) PostDelayedTask(task ICefTask, delayMs int64) bool {
	r1 := CEF().SysCallN(1419, m.Instance(), GetObjectUintptr(task), uintptr(unsafePointer(&delayMs)))
	return GoBool(r1)
}
