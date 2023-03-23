//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy 渲染进程 IPC
package cef

import (
	"fmt"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/json"
	"time"
)

// ipcRenderProcess 渲染进程
type ipcRenderProcess struct {
	bind        *ICefV8Value    // go bind
	ipcObject   *ICefV8Value    // ipc object
	emitHandler *ipcEmitHandler // ipc.emit handler
	onHandler   *ipcOnHandler   // ipc.on handler
	ipcChannel  *renderIPCChan  // channel
	v8Context   *ICefV8Context
}

func (m *ipcRenderProcess) clear() {
	if m.bind != nil {
		m.bind.Free()
		m.bind = nil
	}
	if m.ipcObject != nil {
		m.ipcObject.Free()
		m.ipcObject = nil
	}
	//if m.onHandler != nil {
	//	m.onHandler.clear()
	//}
	if m.v8Context != nil {
		m.v8Context.Free()
		m.v8Context = nil
	}
}

func (m *ipcRenderProcess) initEventGlobal() {
	if m.v8Context == nil {
		m.v8Context = V8ContextRef.Current()
	}
}

// jsOnEvent JS ipc.on 监听事件
func (m *ipcRenderProcess) jsOnEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
	if name != internalIPCOn {
		return
	} else if arguments.Size() != 2 { //必须是2个参数
		exception.SetValue("ipc.on parameter should be 2 quantity")
		arguments.Free()
		return
	}
	m.initEventGlobal()
	var (
		onName      *ICefV8Value // 事件名
		onNameValue string       // 事件名
		onCallback  *ICefV8Value // 事件回调函数
	)
	onName = arguments.Get(0)
	//事件名 第一个参数必须是字符串
	if !onName.IsString() {
		exception.SetValue("ipc.on event name should be a string")
		return
	}
	onCallback = arguments.Get(1)
	//第二个参数必须是函数
	if !onCallback.IsFunction() {
		exception.SetValue("ipc.on event callback should be a function")
		return
	}
	retVal.SetResult(V8ValueRef.NewBool(true))
	onCallback.SetCanNotFree(true)
	onNameValue = onName.GetStringValue()
	//ipc on
	m.onHandler.addCallback(onNameValue, &ipcCallback{function: V8ValueRef.UnWrap(onCallback)})
	result = true
	return
}

// ipcGoExecuteJSEvent Go ipc.emit 执行JS事件
func (m *ipcRenderProcess) ipcGoExecuteJSEvent(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	argumentListBytes := message.ArgumentList().GetBinary(0)
	//接收二进制数据失败
	if argumentListBytes == nil {
		return
	}
	result = true
	var messageDataBytes []byte
	if argumentListBytes.IsValid() {
		size := argumentListBytes.GetSize()
		messageDataBytes = make([]byte, size)
		c := argumentListBytes.GetData(messageDataBytes, 0)
		argumentListBytes.Free()
		message.Free()
		if c == 0 {
			return
		}
	}
	var messageId int32
	var emitName string
	var argument json.JSON
	var argumentList json.JSONArray
	if messageDataBytes != nil {
		argument = json.NewJSON(messageDataBytes)
		messageId = int32(argument.GetIntByKey(ipc_id))
		emitName = argument.GetStringByKey(ipc_event)
		argumentList = argument.GetArrayByKey(ipc_argumentList)
		messageDataBytes = nil
	}
	defer func() {
		if argument != nil {
			argument.Free()
		}
		if argumentList != nil {
			argumentList.Free()
		}
	}()

	if callback := ipcRender.onHandler.getCallback(emitName); callback != nil {
		var callbackArgsBytes []byte
		//enter v8context
		if m.v8Context.Enter() {
			var ret *ICefV8Value
			var argsArray *TCefV8ValueArray
			var err error
			if argumentList != nil {
				//bytes to v8array value
				argsArray, err = ValueConvert.BytesToV8ArrayValue(argumentList.Bytes())
			}
			if argsArray != nil && err == nil {
				// parse v8array success
				ret = callback.function.ExecuteFunctionWithContext(m.v8Context, nil, argsArray)
				argsArray.Free()
			} else {
				// parse v8array fail
				ret = callback.function.ExecuteFunctionWithContext(m.v8Context, nil, nil)
			}
			if ret != nil && ret.IsValid() && messageId != 0 { // messageId != 0 callback func args
				// v8value to process message bytes
				callbackArgsBytes = ValueConvert.V8ValueToProcessMessageBytes(ret)
				ret.Free()
			} else if ret != nil {
				ret.Free()
			}
			//exit v8context
			m.v8Context.Exit()
		}
		if messageId != 0 { //messageId != 0 callback func
			callbackMessage := json.NewJSONObject(nil)
			callbackMessage.Set(ipc_id, messageId)
			if callbackArgsBytes != nil {
				callbackMessage.Set(ipc_argumentList, json.NewJSONArray(callbackArgsBytes).Data())
			} else {
				callbackMessage.Set(ipc_argumentList, nil)
			}
			// frame v8context send ipc message
			if m.v8Context.Frame() != nil {
				// send bytes data to browser ipc
				m.v8Context.Frame().SendProcessMessageForJSONBytes(internalIPCGoExecuteJSEventReplay, consts.PID_BROWSER, callbackMessage.Bytes())
			}
		}
		result = true
	}
	return
}

// jsExecuteGoEvent JS ipc.emit 执行Go事件
func (m *ipcRenderProcess) jsExecuteGoEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
	if name != internalIPCEmit && name != internalIPCEmitSync {
		return
	}
	m.initEventGlobal()
	var (
		emitName      *ICefV8Value //事件名
		emitNameValue string       //事件名
		emitArgs      *ICefV8Value //事件参数
		emitCallback  *ICefV8Value //事件回调函数
		args          []byte
		freeV8Value   = func(value *ICefV8Value) {
			if value != nil {
				value.Free()
			}
		}
	)
	result = true
	defer func() {
		freeV8Value(emitCallback)
		freeV8Value(emitArgs)
		freeV8Value(emitName)
	}()
	if arguments.Size() >= 1 { // 1 ~ 3 个参数
		emitName = arguments.Get(0)
		if !emitName.IsString() {
			exception.SetValue("ipc.emit event name should be a string")
			return
		}
		if arguments.Size() == 2 {
			args2 := arguments.Get(1)
			if args2.IsArray() {
				emitArgs = args2
			} else if args2.IsFunction() {
				emitCallback = args2
			} else {
				exception.SetValue("ipc.emit second argument can only be a parameter or callback function")
				return
			}
		} else if arguments.Size() == 3 {
			emitArgs = arguments.Get(1)
			emitCallback = arguments.Get(2)
			if !emitArgs.IsArray() || !emitCallback.IsFunction() {
				exception.SetValue("ipc.emit second argument can only be a input parameter. third parameter can only be a callback function")
				return
			}
		}
		emitNameValue = emitName.GetStringValue()
		//入参
		if emitArgs != nil {
			//V8Value 转换
			args = ValueConvert.V8ValueToProcessMessageBytes(emitArgs)
			if args == nil {
				exception.SetValue("ipc.emit convert parameter to value value error")
				return
			}
		}
		var isSync = name == internalIPCEmitSync //同步
		//单进程只有同步
		if consts.SingleProcess {
			//不通过进程消息
			callback := &ipcCallback{isSync: true}
			if emitCallback != nil { //callback function
				callback.resultType = rt_function
				callback.function = emitCallback
			} else { //variable
				callback.resultType = rt_variable
			}
			m.singleProcess(emitNameValue, callback, args)
			if callback.resultType == rt_variable {
				if callback.variable != nil {
					retVal.SetResult(callback.variable)
				} else {
					retVal.SetResult(V8ValueRef.NewBool(true))
				}
			} else {
				retVal.SetResult(V8ValueRef.NewBool(true))
			}
		} else { //多进程
			// 同步
			if isSync {
				callback := &ipcCallback{isSync: true, result: make(chan []byte)}
				if emitCallback != nil {
					callback.resultType = rt_function
					callback.function = emitCallback
				} else {
					callback.resultType = rt_variable
				}
				m.emitHandler.addCallback(callback)
				m.multiProcessSync(callback)
				if callback.resultType == rt_variable {
					if callback.variable != nil {
						retVal.SetResult(callback.variable)
					} else {
						retVal.SetResult(V8ValueRef.NewBool(true))
					}
				} else {
					retVal.SetResult(V8ValueRef.NewBool(true))
				}
				return
			} else {
				//异步
				var messageId int32 = 0
				if emitCallback != nil {
					emitCallback.SetCanNotFree(true)
					callback := &ipcCallback{resultType: rt_function, function: V8ValueRef.UnWrap(emitCallback)}
					messageId = m.emitHandler.addCallback(callback)
				}
				if success := m.multiProcessAsync(m.v8Context.Frame(), messageId, isSync, emitNameValue, args); !success {
					//失败，释放回调函数
					emitCallback.SetCanNotFree(false)
				}
			}
			retVal.SetResult(V8ValueRef.NewBool(true))
		}
		args = nil
	}
	return
}

// singleProcess 单进程
func (m *ipcRenderProcess) singleProcess(emitName string, callback *ipcCallback, data []byte) {
	if ipcBrowser == nil {
		return
	}
	var ipcContext = ipcBrowser.jsExecuteGoMethod(m.v8Context.Browser().Identifier(), m.v8Context.Frame().Identifier(), emitName, json.NewJSONArray(data))
	if ipcContext != nil && callback != nil {
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			switch replay.Result()[0].(type) {
			case []byte:
				m.executeCallbackFunction(true, callback, replay.Result()[0].([]byte))
				return
			}
		}
	}
	m.executeCallbackFunction(false, callback, nil)
}

// multiProcessSync 多进程同步消息
func (m *ipcRenderProcess) multiProcessSync(callback *ipcCallback) {
	var isClose = false
	var closeResultChan = func() {
		if !isClose {
			close(callback.result)
			isClose = true
		}
	}
	fmt.Println("等待结果 1")
	delayedWait := time.AfterFunc(2*time.Second, closeResultChan)
	fmt.Println("等待结果 2")
	data, isOpen := <-callback.result
	fmt.Println("结果 isOpen:", isOpen, "data:", string(data))
	if isOpen {
		delayedWait.Stop()
		closeResultChan()
	}
}

// multiProcessAsync 多进程异步消息
func (m *ipcRenderProcess) multiProcessAsync(frame *ICefFrame, messageId int32, isSync bool, emitName string, data []byte) bool {
	if frame != nil {
		message := json.NewJSONObject(nil)
		message.Set(ipc_id, messageId)
		message.Set(ipc_type, isSync)
		message.Set(ipc_event, emitName)
		if data != nil {
			message.Set(ipc_argumentList, json.NewJSONArray(data).Data())
		} else {
			message.Set(ipc_argumentList, nil)
		}
		frame.SendProcessMessageForJSONBytes(internalIPCJSExecuteGoEvent, consts.PID_BROWSER, message.Bytes())
		message.Free()
		return true
	}
	return false
}

// ipcJSExecuteGoEventMessageReply JS执行Go监听，Go的消息回复
func (m *ipcRenderProcess) ipcJSExecuteGoEventMessageReply(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	argumentListBytes := message.ArgumentList().GetBinary(0)
	if argumentListBytes == nil {
		return
	}
	result = true
	var messageDataBytes []byte
	if argumentListBytes.IsValid() {
		size := argumentListBytes.GetSize()
		messageDataBytes = make([]byte, size)
		c := argumentListBytes.GetData(messageDataBytes, 0)
		argumentListBytes.Free()
		if c == 0 {
			result = false
			return
		}
	}
	var (
		messageId    int32
		isReturnArgs bool
		argumentList json.JSONArray
	)
	if messageDataBytes != nil {
		argumentList = json.NewJSONArray(messageDataBytes)
		messageId = int32(argumentList.GetIntByIndex(0))
		isReturnArgs = argumentList.GetBoolByIndex(1)
		messageDataBytes = nil
	}
	defer func() {
		if argumentList != nil {
			argumentList.Free()
		}
	}()
	if m.v8Context == nil {
		return
	}
	if callback := m.emitHandler.getCallback(messageId); callback != nil {
		var returnArgs json.JSONArray
		defer func() {
			if returnArgs != nil {
				returnArgs.Free()
			}
		}()
		//[]byte
		returnArgs = argumentList.GetArrayByIndex(2)
		if callback.isSync {
			if returnArgs != nil {
				callback.result <- returnArgs.Bytes()
			} else {
				callback.result <- nil
			}
		} else {
			if callback.function != nil {
				//设置允许释放
				callback.function.SetCanNotFree(false)
			}
			if returnArgs != nil {
				m.executeCallbackFunction(isReturnArgs, callback, returnArgs.Bytes())
			} else {
				m.executeCallbackFunction(isReturnArgs, callback, nil)
			}
			callback.free()
		}
	}
	return
}

// executeCallbackFunction 执行 v8 function 回调函数
func (m *ipcRenderProcess) executeCallbackFunction(isReturnArgs bool, callback *ipcCallback, returnArgs []byte) {
	if isReturnArgs { //有返回参数
		//enter v8context
		if m.v8Context.Enter() {
			var argsArray *TCefV8ValueArray
			var err error
			if returnArgs != nil {
				//bytes to v8array value
				argsArray, err = ValueConvert.BytesToV8ArrayValue(returnArgs)
			}
			if callback.resultType == rt_function {
				if argsArray != nil && err == nil {
					// parse v8array success
					callback.function.ExecuteFunctionWithContext(m.v8Context, nil, argsArray).Free()
					argsArray.Free()
				} else {
					// parse v8array fail
					callback.function.ExecuteFunctionWithContext(m.v8Context, nil, nil).Free()
				}
			} else if callback.resultType == rt_variable {
				if argsArray != nil && err == nil {
					callback.variable = V8ValueRef.NewArray(int32(argsArray.Size()))
					for i := 0; i < argsArray.Size(); i++ {
						callback.variable.SetValueByIndex(int32(i), argsArray.Get(i))
					}
					argsArray.Free()
				} else {
					callback.variable = V8ValueRef.NewBool(true)
				}
			}
			//exit v8context
			m.v8Context.Exit()
		}
	} else { //无返回参数
		if m.v8Context.Enter() {
			callback.function.ExecuteFunctionWithContext(m.v8Context, nil, nil).Free()
			m.v8Context.Exit()
		}
	}
}

// makeIPC ipc
func (m *ipcRenderProcess) makeIPC(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	// ipc emit
	m.emitHandler.handler = V8HandlerRef.New()
	m.emitHandler.handler.Execute(m.jsExecuteGoEvent)
	// ipc emit sync
	m.emitHandler.handlerSync = V8HandlerRef.New()
	m.emitHandler.handlerSync.Execute(m.jsExecuteGoEvent)
	// ipc on
	m.onHandler.handler = V8HandlerRef.New()
	m.onHandler.handler.Execute(m.jsOnEvent)
	// ipc object
	m.ipcObject = V8ValueRef.NewObject(nil)
	m.ipcObject.setValueByKey(internalIPCEmit, V8ValueRef.newFunction(internalIPCEmit, m.emitHandler.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipcObject.setValueByKey(internalIPCEmitSync, V8ValueRef.newFunction(internalIPCEmitSync, m.emitHandler.handlerSync), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipcObject.setValueByKey(internalIPCOn, V8ValueRef.newFunction(internalIPCOn, m.onHandler.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	// ipc key to v8 global
	context.Global().setValueByKey(internalIPC, m.ipcObject, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}
