//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC

package cef

import (
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/process"
	"sync"
	"time"
)

// ipc bind event name
const (
	internalIPC         = "ipc"      // JavaScript -> ipc 事件驱动, 根对象名
	internalIPCEmit     = "emit"     // JavaScript -> ipc.emit 在 JavaScript 触发 GO 监听事件函数名, 异步
	internalIPCEmitSync = "emitSync" // JavaScript -> ipc.emitSync 在 JavaScript 触发 GO 监听事件函数名, 同步
	internalIPCOn       = "on"       // JavaScript -> ipc.on 在 JavaScript 监听事件, 提供给 GO 调用
)

// ipc message name
const (
	internalIPCJSExecuteGoEvent           = "JSEmitGo"           // JS 触发 GO事件异步
	internalIPCJSExecuteGoEventReplay     = "JSEmitGoReplay"     // JS 触发 GO事件异步 - 返回结果
	internalIPCJSExecuteGoSyncEvent       = "JSEmitSyncGo"       // JS 触发 GO事件同步
	internalIPCJSExecuteGoSyncEventReplay = "JSEmitSyncGoReplay" // JS 触发 GO事件同步 - 返回结果
	internalIPCGoExecuteJSEvent           = "GoEmitJS"           // GO 触发 JS事件
	internalIPCGoExecuteJSEventReplay     = "GoEmitJSReplay"     // GO 触发 JS事件 - 返回结果
)

// js execute go 返回类型
type result_type int8

const (
	rt_function result_type = iota //回调函数
	rt_variable                    //变量接收
)

var (
	internalObjectRootName = "energy"         // GO 和 V8Value 绑定根对象名
	ipcRender              *ipcRenderProcess  // 渲染进程 IPC
	ipcBrowser             *ipcBrowserProcess // 主进程 IPC
)

// ipcEmitHandler ipc.emit 处理器
type ipcEmitHandler struct {
	handler           *ICefV8Handler         // ipc.emit handler
	handlerSync       *ICefV8Handler         // ipc.emitSync handler
	callbackList      map[int32]*ipcCallback // ipc.emit callbackList *list.List
	callbackMessageId int32                  // ipc.emit messageId
	callbackLock      sync.Mutex             // ipc.emit lock
}

// ipcOnHandler ipc.on 处理器
type ipcOnHandler struct {
	handler      *ICefV8Handler          // ipc.on handler
	callbackList map[string]*ipcCallback // ipc.on callbackList
	callbackLock sync.Mutex              // ipc.emit lock
}

// ipcCallback ipc.emit 回调结果
type ipcCallback struct {
	isSync     bool         //同否同步 true:同步 false:异步, 默认false
	resultType result_type  //返回值类型 0:function 1:variable 默认:0
	variable   *ICefV8Value //回调函数, 根据 resultType
	function   *ICefV8Value //回调函数, 根据 resultType
}

// isIPCInternalKey IPC 内部定义使用 key 不允许使用
func isIPCInternalKey(key string) bool {
	return key == internalIPC || key == internalIPCEmit || key == internalIPCOn || key == internalIPCEmitSync ||
		key == internalIPCJSExecuteGoEvent || key == internalIPCJSExecuteGoEventReplay ||
		key == internalIPCGoExecuteJSEvent || key == internalIPCGoExecuteJSEventReplay ||
		key == internalIPCJSExecuteGoSyncEvent || key == internalIPCJSExecuteGoSyncEventReplay

}

// ipcInit 初始化
func ipcInit() {
	isSingleProcess := application.SingleProcess()
	if isSingleProcess {
		ipcBrowser = &ipcBrowserProcess{}
		ipcRender = &ipcRenderProcess{
			syncChan:    &ipc.SyncChan{},
			emitHandler: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
			onHandler:   &ipcOnHandler{callbackList: make(map[string]*ipcCallback)},
		}
		ipc.CreateBrowserIPC()                         // Go IPC browser
		ipc.CreateRenderIPC(0, time.Now().UnixMicro()) // Go IPC render
	} else {
		if process.Args.IsMain() {
			ipcBrowser = &ipcBrowserProcess{}
			ipc.CreateBrowserIPC() // Go IPC browser
		} else if process.Args.IsRender() {
			ipcRender = &ipcRenderProcess{
				syncChan:    &ipc.SyncChan{},
				emitHandler: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
				onHandler:   &ipcOnHandler{callbackList: make(map[string]*ipcCallback)},
			}
			ipc.CreateRenderIPC(0, time.Now().UnixMicro()) // Go IPC render
		}
	}
}

// addCallback
func (m *ipcEmitHandler) addCallback(callback *ipcCallback) int32 {
	//return uintptr(unsafe.Pointer(m.callbackList.PushBack(callback)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	m.callbackList[m.nextMessageId()] = callback
	return m.callbackMessageId
}

// nextMessageId 获取下一个消息ID
func (m *ipcEmitHandler) nextMessageId() int32 {
	m.callbackMessageId++
	if m.callbackMessageId == -1 {
		m.callbackMessageId = 1
	}
	return m.callbackMessageId
}

// getCallback 返回回调函数
func (m *ipcEmitHandler) getCallback(messageId int32) *ipcCallback {
	//return (*list.Element)(unsafe.Pointer(ptr)).Value.(*ipcCallback)
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	if callback, ok := m.callbackList[messageId]; ok {
		delete(m.callbackList, messageId)
		return callback
	}
	return nil
}

//clear 清空所有回调函数
func (m *ipcEmitHandler) clear() {
	for _, v := range m.callbackList {
		v.free()
	}
	m.callbackMessageId = 0
	m.callbackList = make(map[int32]*ipcCallback)
}

// addCallback 根据事件名添加回调函数
func (m *ipcOnHandler) addCallback(eventName string, callback *ipcCallback) {
	//return uintptr(unsafe.Pointer(m.callbackList.PushBack(callback)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	m.callbackList[eventName] = callback
}

// removeCallback 根据事件名移除回调函数
func (m *ipcOnHandler) removeCallback(eventName string) {
	//m.callbackList.Remove((*list.Element)(unsafe.Pointer(ptr)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	delete(m.callbackList, eventName)
}

// getCallback 根据事件名返回回调函数
func (m *ipcOnHandler) getCallback(eventName string) *ipcCallback {
	//return (*list.Element)(unsafe.Pointer(ptr)).Value.(*ipcCallback)
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	return m.callbackList[eventName]
}

//clear 清空所有回调函数
func (m *ipcOnHandler) clear() {
	for _, v := range m.callbackList {
		v.free()
	}
	m.callbackList = make(map[string]*ipcCallback)
}

//free 清空所有回调函数
func (m *ipcCallback) free() {
	if m.function != nil {
		m.function.Free()
		m.function = nil
	}
}
