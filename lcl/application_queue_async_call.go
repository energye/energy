//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 在应用主线程中执行(非主线程使用)异步&同步执行包裹函数

package lcl

import (
	"github.com/energye/energy/v2/api"
	"math"
	"runtime"
	"sync"
)

var (
	qac = &queueAsyncCall{id: 0, calls: sync.Map{}}
)

func threadAsyncCallbackProc(id uintptr) uintptr {
	qac.call(id)
	return 0
}

type queueAsyncCall struct {
	id    uintptr
	calls sync.Map
}

// RunOnMainThreadAsync
//
//	主线程中执行, 异步
func RunOnMainThreadAsync(callback TMainThreadAsyncProc) int {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if api.DMainThreadId() == api.DCurrentThreadId() {
		callback(0)
		return 0
	}
	id := qac.set(callback)
	api.DRunMainAsyncCall(id)
	return int(id)
}

func (m *queueAsyncCall) call(id uintptr) {
	if call, ok := m.calls.Load(id); ok {
		m.calls.Delete(id)
		call.(TMainThreadAsyncProc)(uint32(id))
	}
}
func (m *queueAsyncCall) set(fn TMainThreadAsyncProc) uintptr {
	if m.id >= math.MaxUint32 {
		m.id = 0
	}
	m.id++
	m.calls.Store(m.id, fn)
	return m.id
}
