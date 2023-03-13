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

const (
	internalIPCKey = "ipc"  // JavaScript -> ipc 事件驱动, 根对象名
	internalEmit   = "emit" // JavaScript -> ipc.emit 在 JavaScript 触发 GO 监听事件函数名
	internalOn     = "on"   // JavaScript -> ipc.on 在 JavaScript 监听事件, 提供给 GO 调用
)
const (
	internalProcessMessageIPCEmit      = "ipcEmit"      // 进程消息 ipcEmit
	internalProcessMessageIPCEmitReply = "ipcEmitReply" // 进程消息 ipcEmitReply
	internalProcessMessageIPCOn        = "ipcOn"        // 进程消息 ipcOn
)

var (
	internalObjectRootName = "energy"         // GO 和 V8Value 绑定根对象名
	ipcRender              *ipcRenderProcess  //渲染进程 IPC
	ipcBrowser             *ipcBrowserProcess //主进程 IPC
)

// ipcEmitHandler
type ipcEmitHandler struct {
	handler           *ICefV8Handler         // ipc.emit handler
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
	arguments *TCefV8ValueArray
	context   *ICefV8Context
	function  *ICefV8Value
}

// ipcInit 初始化
func ipcInit() {
	isSingleProcess := application.SingleProcess()
	if isSingleProcess {
		ipcBrowser = &ipcBrowserProcess{}
		ipcRender = &ipcRenderProcess{
			ipcEmit: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
			ipcOn:   &ipcOnHandler{callbackList: make(map[string]*ipcCallback)},
		}
	} else {
		if common.Args.IsMain() {
			ipcBrowser = &ipcBrowserProcess{}
		} else if common.Args.IsRender() {
			ipcRender = &ipcRenderProcess{
				ipcEmit: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
				ipcOn:   &ipcOnHandler{callbackList: make(map[string]*ipcCallback)},
			}
		}
	}
}

// isIPCInternalKey IPC 内部 key 不允许使用
func isIPCInternalKey(key string) bool {
	return key == internalIPCKey || key == internalEmit || key == internalOn ||
		key == internalProcessMessageIPCEmit || key == internalProcessMessageIPCOn || key == internalProcessMessageIPCEmitReply
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

// removeCallback
func (m *ipcEmitHandler) removeCallback(messageId int32) {
	//m.callbackList.Remove((*list.Element)(unsafe.Pointer(ptr)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	delete(m.callbackList, messageId)
}

// getCallback
func (m *ipcEmitHandler) getCallback(messageId int32) *ipcCallback {
	//return (*list.Element)(unsafe.Pointer(ptr)).Value.(*ipcCallback)
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	return m.callbackList[messageId]
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

func (m *ipcCallback) free() {
	if m.context != nil {
		m.context.Free()
		m.context = nil
	}
	if m.function != nil {
		m.function.Free()
		m.function = nil
	}
	if m.arguments != nil {
		m.arguments.Free()
		m.arguments = nil
	}
}
