//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

//应用主线程异步回调
import (
	"math"
	"sync"
)

var (
	Qac = &queueAsyncCall{id: 0, calls: sync.Map{}}
	//TODO 注释后面改
	//applicationQueueAsyncCallFunc *dylib.LazyProc
)

func ApplicationQueueAsyncCallInit() {
	//TODO 注释后面改
	//applicationQueueAsyncCallFunc = api.GetLibLCL().NewProc("SetApplicationQueueAsyncCallFunc")
	//applicationQueueAsyncCallFunc.Call(applicationQueueAsyncCallEvent)
}

func applicationQueueAsyncCallProc(id uintptr) uintptr {
	Qac.call(id)
	return 0
}

// 队列异步调用函数 id:事件id
type QacFn func(id int)

type QueueCall struct {
	IsSync bool
	Fn     QacFn
	Wg     *sync.WaitGroup
}

type queueAsyncCall struct {
	id    uintptr
	calls sync.Map
}

// 1.在UI主进程中执行, 队列异步调用-适用大多场景(包括UI线程和非UI线程)
//
// 2.大多数非UI线程操作都需要使用该函数
//
// 3.在任何变更UI的操作都有可能导致UI线程不一至出现程序错误或程序崩溃, 可以尝试使用该回调函数解决.
//
// 4.在windows linux macos 可同时使用
func QueueAsyncCall(fn QacFn) int {
	id := Qac.Set(&QueueCall{
		IsSync: false,
		Fn:     fn,
	})
	//TODO 注释后面改
	//api.GetLazyProc("CEFApplication_QueueAsyncCall").Call(id)
	return int(id)
}

// 1.在UI主进程中执行, 队列异步调用-适用大多场景(非UI线程)
//
// 2.大多数非UI线程操作都需要使用该函数
//
// 3.在任何变更UI的操作都有可能导致UI线程不一至出现程序错误或程序崩溃, 可以尝试使用该回调函数解决.
//
// 4.在windows linux macos 需要注意使用场景, 当非UI线程使用时正常执行, UI线程使用时会造成UI线程锁死, 这种情况建议使用 QueueAsyncCall 自己增加同步锁
func QueueSyncCall(fn QacFn) int {
	qc := &QueueCall{
		IsSync: true,
		Fn:     fn,
		Wg:     &sync.WaitGroup{},
	}
	qc.Wg.Add(1)
	id := Qac.Set(qc)
	//TODO 注释后面改
	//api.GetLazyProc("CEFApplication_QueueAsyncCall").Call(id)
	qc.Wg.Wait()
	qc.Fn = nil
	qc.Wg = nil
	qc = nil
	return int(id)
}

func (m *queueAsyncCall) call(id uintptr) {
	if call, ok := m.calls.LoadAndDelete(id); ok {
		qc := call.(*QueueCall)
		if qc.IsSync {
			qc.Fn(int(id))
			qc.Wg.Done()
		} else {
			qc.Fn(int(id))
		}
	}
}
func (m *queueAsyncCall) Set(fn *QueueCall) uintptr {
	if m.id >= math.MaxUint {
		m.id = 0
	}
	m.id++
	m.calls.Store(m.id, fn)
	return m.id
}
