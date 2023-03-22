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
	"github.com/energye/energy/common"
	"sync"
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

// ipc process message key
const (
	ipc_id           = "id"
	ipc_event        = "event"
	ipc_argumentList = "argumentList"
)

var (
	internalObjectRootName = "energy"         // GO 和 V8Value 绑定根对象名
	ipcRender              *ipcRenderProcess  // 渲染进程 IPC
	ipcBrowser             *ipcBrowserProcess // 主进程 IPC
)

// ipcEmitHandler
type ipcEmitHandler struct {
	handler           *ICefV8Handler         // ipc.emit handler
	handlerSync       *ICefV8Handler         // ipc.emitSync handler
	callbackList      map[int32]*ipcCallback // ipc.emit callbackList *list.List
	callbackMessageId int32                  // ipc.emit messageId
	callbackLock      sync.Mutex             // ipc.emit lock
}

// ipcOnHandler
type ipcOnHandler struct {
	handler           *ICefV8Handler          // ipc.on handler
	callbackList      map[string]*ipcCallback // ipc.on callbackList
	callbackMessageId int32                   // ipc.on messageId
	callbackLock      sync.Mutex              // ipc.emit lock
}

// ipcCallback
type ipcCallback struct {
	function *ICefV8Value
}

// isIPCInternalKey IPC 内部 key 不允许使用
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
		ipcBrowser = &ipcBrowserProcess{emitHandler: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)}}
		ipcRender = &ipcRenderProcess{
			emitHandler: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
			onHandler:   &ipcOnHandler{callbackList: make(map[string]*ipcCallback)},
		}
	} else {
		if common.Args.IsMain() {
			ipcBrowser = &ipcBrowserProcess{}
		} else if common.Args.IsRender() {
			ipcRender = &ipcRenderProcess{
				emitHandler: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
				onHandler:   &ipcOnHandler{callbackList: make(map[string]*ipcCallback)},
			}
		}
	}
}

// addCallback
func (m *ipcEmitHandler) addCallback(callback *ipcCallback) int32 {
	//return uintptr(unsafe.Pointer(m.callbackList.PushBack(callback)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	if m.callbackMessageId == -1 {
		m.callbackMessageId = 1
	} else {
		m.callbackMessageId++
	}
	m.callbackList[m.callbackMessageId] = callback
	return m.callbackMessageId
}

// getCallback
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

func (m *ipcEmitHandler) clear() {
	for _, v := range m.callbackList {
		v.free()
	}
	m.callbackList = make(map[int32]*ipcCallback)
}

// addCallback
func (m *ipcOnHandler) addCallback(eventName string, callback *ipcCallback) int32 {
	//return uintptr(unsafe.Pointer(m.callbackList.PushBack(callback)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	m.callbackList[eventName] = callback
	return m.callbackMessageId
}

// removeCallback
func (m *ipcOnHandler) removeCallback(eventName string) {
	//m.callbackList.Remove((*list.Element)(unsafe.Pointer(ptr)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	delete(m.callbackList, eventName)
}

// getCallback
func (m *ipcOnHandler) getCallback(eventName string) *ipcCallback {
	//return (*list.Element)(unsafe.Pointer(ptr)).Value.(*ipcCallback)
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	return m.callbackList[eventName]
}

func (m *ipcOnHandler) clear() {
	for _, v := range m.callbackList {
		v.free()
	}
	m.callbackList = make(map[string]*ipcCallback)
}

func (m *ipcCallback) free() {
	if m.function != nil {
		m.function.Free()
		m.function = nil
	}
}
