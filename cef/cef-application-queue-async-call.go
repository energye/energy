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
	"github.com/energye/golcl/lcl"
)

// 1.在UI主进程中执行, 队列异步调用-适用大多场景(包括UI线程和非UI线程)
//
// 2.大多数非UI线程操作都需要使用该函数
//
// 3.在任何变更UI的操作都有可能导致UI线程不一至出现程序错误或程序崩溃, 可以尝试使用该回调函数解决.
//
// 4.在windows linux macos 可同时使用
func QueueAsyncCall(fn lcl.QacFn) int {
	return lcl.QueueAsyncCall(fn)
}

// 1.在UI主进程中执行, 队列异步调用-适用大多场景(非UI线程)
//
// 2.大多数非UI线程操作都需要使用该函数
//
// 3.在任何变更UI的操作都有可能导致UI线程不一至出现程序错误或程序崩溃, 可以尝试使用该回调函数解决.
//
// 4.在windows linux macos 需要注意使用场景, 当非UI线程使用时正常执行, UI线程使用时会造成UI线程锁死, 这种情况建议使用 QueueAsyncCall 自己增加同步锁
func QueueSyncCall(fn lcl.QacFn) int {
	return lcl.QueueSyncCall(fn)
}
