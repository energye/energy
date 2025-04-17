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

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api/dllimports"
	"math"
	"sync"
)

var (
	qac                           = &queueAsyncCall{id: 0, calls: sync.Map{}}
	applicationQueueAsyncCallFunc dllimports.ProcAddr
)

func applicationQueueAsyncCallInit() {
	applicationQueueAsyncCallFunc = imports.Proc(def.SetApplicationQueueAsyncCallFunc)
	applicationQueueAsyncCallFunc.Call(applicationQueueAsyncCallEvent)
}

func applicationQueueAsyncCallProc(id uintptr) uintptr {
	qac.call(id)
	return 0
}

// qacFn 队列异步调用函数 id:事件id
type qacFn func(id int)

type queueCall struct {
	IsSync bool
	Fn     qacFn
	Wg     *sync.WaitGroup
}

type queueAsyncCall struct {
	id    uintptr
	calls sync.Map
	lock  sync.Mutex
}

// QueueAsyncCall 仅LCL，在主进程中异步调用
//
// 在UI主进程中执行, 异步执行
//
// 非主进程的多线程操作可使用该函数包裹
//
// 在任何变更UI的操作都有可能因非主线程出现不一至, 而出现程序错误或程序崩溃, 可以尝试使用该回调函数解决.
//
// 提示: CEF事件或函数中不应使用该函数包裹
func QueueAsyncCall(fn qacFn) int {
	qac.lock.Lock()
	defer qac.lock.Unlock()
	id := qac.set(&queueCall{
		IsSync: false,
		Fn:     fn,
	})
	imports.Proc(def.CEFApplication_QueueAsyncCall).Call(id)
	return int(id)
}

// QueueSyncCall 同 QueueAsyncCall
//
// 同步执行 - 阻塞UI
func QueueSyncCall(fn qacFn) int {
	qc := &queueCall{
		IsSync: true,
		Fn:     fn,
		Wg:     &sync.WaitGroup{},
	}
	qc.Wg.Add(1)
	id := qac.set(qc)
	imports.Proc(def.CEFApplication_QueueAsyncCall).Call(id)
	qc.Wg.Wait()
	qc.Fn = nil
	qc.Wg = nil
	qc = nil
	return int(id)
}

func (m *queueAsyncCall) call(id uintptr) {
	if call, ok := m.calls.Load(id); ok {
		m.calls.Delete(id)
		qc := call.(*queueCall)
		if qc.IsSync {
			qc.Fn(int(id))
			qc.Wg.Done()
		} else {
			qc.Fn(int(id))
		}
	}
}
func (m *queueAsyncCall) set(fn *queueCall) uintptr {
	if m.id >= math.MaxUint32 {
		m.id = 0
	}
	m.id++
	m.calls.Store(m.id, fn)
	return m.id
}
