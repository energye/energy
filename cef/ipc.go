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
	"github.com/cyber-xxm/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/cyber-xxm/energy/v2/cef/ipc/argument"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/consts"
	"sync"
)

// ipc bind event name
const (
	internalIPC             = "ipc"             // JavaScript -> ipc 事件驱动, 根对象名
	internalIPCEmit         = "emit"            // JavaScript -> ipc.emit 在 JavaScript 触发 GO 监听事件函数名, 异步
	internalIPCEmitWait     = "emitWait"        // JavaScript -> ipc.emitWait 在 JavaScript 触发 GO 监听事件函数名, 等待超时返回结果
	internalIPCOn           = "on"              // JavaScript -> ipc.on 在 JavaScript 监听事件, 提供给 GO 调用
	internalIPCDRAG         = "drag"            // JavaScript -> ipc.on drag
	internalEnergyExtension = "energyExtension" // JavaScript -> energyExtension object
)

// ipc message name
const (
	internalIPCJSEmit           = "JSEmit"           // JS 触发 GO事件异步
	internalIPCJSEmitReplay     = "JSEmitReplay"     // JS 触发 GO事件异步 - 返回结果
	internalIPCJSEmitWait       = "JSEmitWait"       // JS 触发 GO事件同步
	internalIPCJSEmitWaitReplay = "JSEmitWaitReplay" // JS 触发 GO事件同步 - 返回结果
	internalIPCGoEmit           = "GoEmit"           // GO 触发 JS事件
	internalIPCGoEmitReplay     = "GoEmitReplay"     // GO 触发 JS事件 - 返回结果
)

// js execute go 返回类型
type result_type int8

const (
	rt_function result_type = iota //回调函数
	rt_variable                    //变量接收
)

var (
	ipcRender  *ipcRenderProcess  // 渲染进程 IPC
	ipcBrowser *ipcBrowserProcess // 主进程 IPC
)

// ipc.emit 处理器
type ipcEmitHandler struct {
	handler           *ICefV8Handler         // ipc.emit handler
	handlerSync       *ICefV8Handler         // ipc.emitSync handler
	callbackList      map[int32]*ipcCallback // ipc.emit callbackList *list.List
	callbackMessageId int32                  // ipc.emit messageId
	callbackLock      sync.Mutex             // ipc.emit lock
}

// ipc.on 处理器
type ipcOnHandler struct {
	handler      *ICefV8Handler          // ipc.on handler
	callbackList map[string]*ipcCallback // ipc.on callbackList
	callbackLock sync.Mutex              // ipc.emit lock
}

// ipc.emit 回调结果
type ipcCallback struct {
	isSync       bool                     //同否同步 true:同步 false:异步, 默认false
	resultType   result_type              //返回值类型 0:function 1:variable 默认:0
	variable     *ICefV8Value             //回调函数, 根据 resultType
	function     *ICefV8Value             //回调函数, 根据 resultType
	name         *ICefV8Value             //事件名称
	mode         types.Mode               //监听模式
	asyncHandler *asyncGoExecuteJSHandler //监听模式为 MAsync
}

// IPC 内部定义使用 key 不允许使用
func isIPCInternalKey(key string) bool {
	return key == internalIPC || key == internalIPCEmit || key == internalIPCOn ||
		key == internalIPCDRAG || key == internalIPCEmitWait ||
		key == internalIPCJSEmit || key == internalIPCJSEmitReplay ||
		key == internalIPCGoEmit || key == internalIPCGoEmitReplay ||
		key == internalIPCJSEmitWait || key == internalIPCJSEmitWaitReplay ||
		key == internalEnergyExtension

}

// 初始化
func ipcInit() {
	isSingleProcess := application.SingleProcess()
	if isSingleProcess {
		ipcBrowser = &ipcBrowserProcess{}
		ipcRender = &ipcRenderProcess{
			waitChan:    &ipc.WaitChan{Pending: new(sync.Map)},
			emitHandler: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
			onHandler:   make(map[string]*ipcOnHandler),
		}
		dragExtension = &dragExtensionHandler{}
	} else {
		if process.Args.IsMain() {
			ipcBrowser = &ipcBrowserProcess{}
			dragExtension = &dragExtensionHandler{}
		} else if process.Args.IsRender() {
			ipcRender = &ipcRenderProcess{
				waitChan:    &ipc.WaitChan{Pending: new(sync.Map)},
				emitHandler: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback)},
				onHandler:   make(map[string]*ipcOnHandler),
			}
			dragExtension = &dragExtensionHandler{}
		}
	}
}

// 添加一个回调函数
func (m *ipcEmitHandler) addCallback(callback *ipcCallback) int32 {
	//return uintptr(unsafe.Pointer(m.callbackList.PushBack(callback)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	m.callbackList[m.nextMessageId()] = callback
	return m.callbackMessageId
}

// 获取下一个消息ID
func (m *ipcEmitHandler) nextMessageId() int32 {
	m.callbackMessageId++
	if m.callbackMessageId == -1 {
		m.callbackMessageId = 1
	}
	return m.callbackMessageId
}

// 返回回调函数
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

// 清空所有回调函数
func (m *ipcEmitHandler) clear() {
	for _, v := range m.callbackList {
		v.function.SetCanNotFree(false)
		v.function.Free()
		v.name.Free()
	}
	m.callbackMessageId = 0
	m.callbackList = make(map[int32]*ipcCallback)
}

// 根据事件名添加回调函数
func (m *ipcOnHandler) addCallback(eventName string, callback *ipcCallback) {
	//return uintptr(unsafe.Pointer(m.callbackList.PushBack(callback)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	if callbk, ok := m.callbackList[eventName]; ok {
		//如果重复将之前释放掉
		callbk.function.SetCanNotFree(false)
		callbk.function.Free()
		callbk.name.Free()
	}
	m.callbackList[eventName] = callback
}

// 根据事件名移除回调函数
func (m *ipcOnHandler) removeCallback(eventName string) {
	//m.callbackList.Remove((*list.Element)(unsafe.Pointer(ptr)))
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	delete(m.callbackList, eventName)
}

// 根据事件名返回回调函数
func (m *ipcOnHandler) getCallback(eventName string) *ipcCallback {
	//return (*list.Element)(unsafe.Pointer(ptr)).Value.(*ipcCallback)
	m.callbackLock.Lock()
	defer m.callbackLock.Unlock()
	return m.callbackList[eventName]
}

// 清空所有回调函数
func (m *ipcOnHandler) clear() {
	for _, v := range m.callbackList {
		v.function.SetCanNotFree(false)
		v.function.Free()
		v.name.Free()
	}
	m.callbackList = make(map[string]*ipcCallback)
}

// JS: ipc.on 异步模式完成事件处理
type asyncGoExecuteJSHandler struct {
	handler  *ICefV8Handler
	callback *ICefV8Value
}

// 返回一个新的 asyncGoExecuteJSHandler
func newAsyncGoExecuteJSHandler() *asyncGoExecuteJSHandler {
	asyncHandler := &asyncGoExecuteJSHandler{
		handler: V8HandlerRef.New(),
	}
	// 创建处理器函数, 在异步监听模式时使用
	asyncHandler.handler.Execute(asyncHandler.asyncHandler)
	asyncHandler.callback = V8ValueRef.NewFunction("callback", asyncHandler.handler)
	asyncHandler.callback.SetCanNotFree(true)
	return asyncHandler
}

func (m *asyncGoExecuteJSHandler) free() {
	if m.callback != nil {
		m.callback.Free()
		m.callback = nil
	}
}

// Js: ipc.on 异步监听模式, 完成处理器
func (m *asyncGoExecuteJSHandler) asyncHandler(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
	idValue := object.GetValueByKey("id")
	defer idValue.Free()
	messageId := idValue.GetIntValue()
	size := arguments.Size()
	result := make([]interface{}, size)
	for i := 0; i < size; i++ {
		tempVal := arguments.Get(i)
		if !tempVal.IsValid() {
			result[i] = nil
			continue
		}
		isObject := tempVal.IsObject() || tempVal.IsArray()
		callbackArgsBytes := ValueConvert.V8ValueToProcessMessageArray(tempVal)
		tempVal.Free()
		if isObject {
			result[i] = callbackArgsBytes
		} else {
			if val, ok := callbackArgsBytes.([]interface{}); ok && len(val) == 1 {
				result[i] = callbackArgsBytes.([]interface{})[0]
			} else {
				result[i] = nil
			}
		}
	}
	v8ctx := V8ContextRef.Current()
	defer v8ctx.Free()
	var processMessage target.IProcessMessage
	if application.Is49() {
		// CEF49
		processMessage = v8ctx.Browser()
	} else {
		processMessage = v8ctx.Frame()
	}
	replayGoExecuteJSEvent(processMessage, messageId, result)
	return true
}

func replayGoExecuteJSEvent(processMessage target.IProcessMessage, messageId int32, callbackArgsBytes interface{}) {
	if messageId != 0 {
		callbackMessage := &ipcArgument.List{
			Id: messageId,
		}
		if callbackArgsBytes != nil {
			callbackMessage.Data = callbackArgsBytes //json.NewJSONArray(callbackArgsBytes).Data()
		}
		//发送消息到主进程
		processMessage.SendProcessMessageForJSONBytes(internalIPCGoEmitReplay, consts.PID_BROWSER, callbackMessage.Bytes())
		callbackMessage.Reset()
	}
}
