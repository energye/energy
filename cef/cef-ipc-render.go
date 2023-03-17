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
)

// ipcRenderProcess 渲染进程
type ipcRenderProcess struct {
	bind        *ICefV8Value    // go bind
	ipc         *ICefV8Value    // ipc object
	emitHandler *ipcEmitHandler // ipc.emit handler
	onHandler   *ipcOnHandler   // ipc.on handler
}

func (m *ipcRenderProcess) clear() {
	if m.bind != nil {
		m.bind.Free()
		m.bind = nil
	}
	if m.ipc != nil {
		m.ipc.Free()
		m.ipc = nil
	}
	//if m.onHandler != nil {
	//	m.onHandler.clear()
	//}
}

// makeCtx ipc 和 bind
func (m *ipcRenderProcess) makeCtx(context *ICefV8Context) {
	m.makeIPC(context)
	//m.makeBind(context)
}

// makeIPC ipc
func (m *ipcRenderProcess) makeIPC(context *ICefV8Context) {
	// ipc emit
	m.emitHandler.handler = V8HandlerRef.New()
	m.emitHandler.handler.Execute(m.ipcJSExecuteGoEvent)
	// ipc on
	m.onHandler.handler = V8HandlerRef.New()
	m.onHandler.handler.Execute(m.ipcJSOnEvent)
	// ipc object
	m.ipc = V8ValueRef.NewObject(nil)
	m.ipc.setValueByKey(internalEmit, V8ValueRef.newFunction(internalEmit, m.emitHandler.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipc.setValueByKey(internalOn, V8ValueRef.newFunction(internalOn, m.onHandler.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	// global to v8 ipc key
	context.Global().setValueByKey(internalIPCKey, m.ipc, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}

// ipcJSOnEvent JS ipc.on 监听事件
func (m *ipcRenderProcess) ipcJSOnEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
	if name != internalOn {
		retVal.SetResult(V8ValueRef.NewBool(false))
		return
	} else if arguments.Size() != 2 { //必须是2个参数
		retVal.SetResult(V8ValueRef.NewBool(false))
		exception.SetValue("ipc.on parameter should be 2 quantity")
		arguments.Free()
		return
	}
	var (
		onName      *ICefV8Value // 事件名
		onNameValue string       // 事件名
		onCallback  *ICefV8Value // 事件回调函数
	)
	onName = arguments.Get(0)
	//事件名 第一个参数必须是字符串
	if !onName.IsString() {
		retVal.SetResult(V8ValueRef.NewBool(false))
		exception.SetValue("ipc.on event name should be a string")
		return
	}
	onCallback = arguments.Get(1)
	//第二个参数必须是函数
	if !onCallback.IsFunction() {
		retVal.SetResult(V8ValueRef.NewBool(false))
		exception.SetValue("ipc.on event callback should be a function")
		return
	}
	onCallback.SetCanNotFree(true)
	onNameValue = onName.GetStringValue()
	//ipc on
	m.onHandler.addCallback(onNameValue, &ipcCallback{ /*arguments: arguments, */ context: V8ContextRef.Current(), function: V8ValueRef.UnWrap(onCallback)})
	result = true
	return
}

// ipcGoExecuteJSEvent Go ipc.emit 执行JS事件
func (m *ipcRenderProcess) ipcGoExecuteJSEvent(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	//argument := message.ArgumentList()
	argumentListBytes := message.ArgumentList().GetBinary(0)
	var messageDataBytes []byte
	if argumentListBytes.IsValid() {
		size := argumentListBytes.GetSize()
		messageDataBytes = make([]byte, size)
		c := argumentListBytes.GetData(messageDataBytes, 0)
		argumentListBytes.Free() //立即释放掉
		if c == 0 {
			result = false
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
	}()

	if callback := ipcRender.onHandler.getCallback(emitName); callback != nil {
		var (
			resultData   []byte
			ret          *ICefV8Value
			replyMessage *ICefProcessMessage
		)
		defer func() {
			if replyMessage != nil {
				replyMessage.Free()
			}
			if argumentList != nil {
				argumentList.Free()
			}
		}()
		if callback.context.Enter() {
			ret = callback.function.ExecuteFunctionWithContextForArgsBytes(callback.context, nil, argumentList.Bytes())
			if messageId != 0 { // callback
				resultData = ipcValueConvert.V8ValueToProcessMessageBytes(ret)
				returnArgs := json.NewJSONArray(nil)
				returnArgs.Add(messageId) //0 消息ID
				returnArgs.Add(false)     //1 是否有返回值
				if resultData != nil {
					returnArgs.SetByIndex(1, true) //1 有返回值
					returnArgs.Add(resultData)     //2 result []byte
				}
				frame.SendProcessMessageForJSONBytes(internalProcessMessageIPCEmitReply, consts.PID_BROWSER, returnArgs.Bytes())
				returnArgs.Free()
			}
			ret.Free()
			callback.context.Exit()
		}
		result = true
	}
	return
}

// ipcJSExecuteGoEvent JS ipc.emit 执行Go事件
func (m *ipcRenderProcess) ipcJSExecuteGoEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
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
	isFree := false
	defer func() {
		if isFree { //失败时释放掉
			result = false
			freeV8Value(emitCallback)
		}
	}()
	if name != internalEmit {
		isFree = true
		return
	}
	if arguments.Size() >= 1 { // 1 ~ 3 个参数
		defer func() {
			//释放掉这些指针，不然不会自动释放
			freeV8Value(emitName)
			freeV8Value(emitArgs)
			if args != nil {
				args = nil
			}
		}()
		emitName = arguments.Get(0)
		if !emitName.IsString() {
			exception.SetValue("ipc.emit event name should be a string")
			isFree = true
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
				isFree = true
				return
			}
		} else if arguments.Size() == 3 {
			emitArgs = arguments.Get(1)
			emitCallback = arguments.Get(2)
			if !emitArgs.IsArray() || !emitCallback.IsFunction() {
				exception.SetValue("ipc.emit second argument can only be a input parameter. third parameter can only be a callback function")
				isFree = true
				return
			}
		}
		emitNameValue = emitName.GetStringValue()
		//入参
		if emitArgs != nil {
			//V8Value 转换
			args = ipcValueConvert.V8ValueToProcessMessageBytes(emitArgs)
			if args == nil {
				exception.SetValue("ipc.emit convert parameter to value value error")
				isFree = true
				return
			}
		}
		context := V8ContextRef.Current()
		var messageId int32 = 0
		//回调函数
		if emitCallback != nil {
			//回调函数临时存放到缓存中
			emitCallback.SetCanNotFree(true)
			messageId = ipcRender.emitHandler.addCallback(&ipcCallback{
				//arguments: arguments,
				context:  context,
				function: V8ValueRef.UnWrap(emitCallback),
			})
		}

		frame := context.Frame()
		message := json.NewJSONObject(nil)
		message.Set(ipc_id, messageId)
		message.Set(ipc_event, emitNameValue)
		message.Set(ipc_argumentList, json.NewJSONArray(args).Data())
		frame.SendProcessMessageForJSONBytes(internalProcessMessageIPCEmit, consts.PID_BROWSER, message.Bytes())
		message.Free()
		frame.Free()
		args = nil
		//context.Free() // TODO dev
		retVal.SetResult(V8ValueRef.NewBool(true))
	}
	return
}

// ipcJSExecuteGoEventMessageReply JS执行Go监听，Go的消息回复
func (m *ipcRenderProcess) ipcJSExecuteGoEventMessageReply(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	result = true
	argumentListBytes := message.ArgumentList().GetBinary(0)
	var messageDataBytes []byte
	if argumentListBytes.IsValid() {
		size := argumentListBytes.GetSize()
		messageDataBytes = make([]byte, size)
		c := argumentListBytes.GetData(messageDataBytes, 0)
		argumentListBytes.Free() //立即释放掉
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

	if callback := ipcRender.emitHandler.getCallback(messageId); callback != nil {
		callback.function.SetCanNotFree(false)
		//第二个参数 true 有返回参数
		if isReturnArgs {
			//[]byte
			returnArgs := argumentList.GetArrayByIndex(2)
			//解析 '[]byte' 参数
			if callback.context.Enter() {
				callback.function.ExecuteFunctionWithContextForArgsBytes(callback.context, nil, returnArgs.Bytes()).Free()
				callback.context.Exit()
			}
			returnArgs.Free()
		} else { //无返回参数
			if callback.context.Enter() {
				callback.function.ExecuteFunctionWithContext(callback.context, nil, nil).Free()
			}
			callback.context.Exit()
		}
		//remove
		callback.free()
	}
	return
}

// makeBind bind object accessor
func (m *ipcRenderProcess) makeBind(context *ICefV8Context) {
	var (
		bindRootObjectAccessor *ICefV8Accessor // bind accessor
		bindSubObjectAccessor  *ICefV8Accessor // bind sub object accessor
		bindFuncHandler        *ICefV8Handler  // bind func handler
	)
	bindRootObjectAccessor = V8AccessorRef.New()
	bindRootObjectAccessor.Get(m.bindGet)
	bindRootObjectAccessor.Set(m.bindSet)
	// bind object
	m.bind = V8ValueRef.NewObject(bindRootObjectAccessor)

	fmt.Println("\tbindCount", VariableBind.bindCount())
	binds := VariableBind.binds()
	for name, value := range binds {
		fmt.Println("\t", name, "\t", value.ValueType().ToString())
		var v8value *ICefV8Value
		if value.IsFunction() {
			if bindFuncHandler == nil {
				bindFuncHandler = V8HandlerRef.New()
				bindFuncHandler.Execute(m.bindFuncExecute)
			}
			v8value = V8ValueRef.newFunction(name, bindFuncHandler)
			m.bind.setValueByKey(name, v8value, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		} else {
			if value.IsString() {
				v8value = V8ValueRef.NewString("")
			} else if value.IsInteger() {
				v8value = V8ValueRef.NewInt(0)
			} else if value.IsDouble() {
				v8value = V8ValueRef.NewDouble(0.0)
			} else if value.IsBool() {
				v8value = V8ValueRef.NewBool(false)
			} else if value.IsNull() {
				v8value = V8ValueRef.NewNull()
			} else if value.IsUndefined() {
				v8value = V8ValueRef.NewUndefined()
			} else if value.IsObject() {
				if bindSubObjectAccessor == nil {
					bindSubObjectAccessor = V8AccessorRef.New()
					bindSubObjectAccessor.Get(m.bindSubObjectGet)
					bindSubObjectAccessor.Set(m.bindSubObjectSet)
				}
				v8value = V8ValueRef.NewObject(bindSubObjectAccessor)
			} else if value.IsArray() {
				v8value = V8ValueRef.NewArray(0)
			}
			if v8value != nil {
				m.bind.setValueByAccessor(name, consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE)
				m.bind.setValueByKey(name, v8value, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			}
		}
	}

	// global to v8 bind objectRootName
	context.Global().setValueByKey(internalObjectRootName, m.bind, consts.V8_PROPERTY_ATTRIBUTE_NONE)
}

// bindFuncExecute 绑定函数执行
func (m *ipcRenderProcess) bindFuncExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
	fmt.Println("bindFuncExecute handler name:", name)
	return false
}

// bindSubObjectGet 绑定对象取值
func (m *ipcRenderProcess) bindSubObjectGet(name string, object *ICefV8Value, retVal *ResultV8Value, exception *ResultString) bool {
	fmt.Println("bindSubObjectGet accessor name:", name)
	return false
}

// bindSubObjectSet 绑定对象赋值
func (m *ipcRenderProcess) bindSubObjectSet(name string, object *ICefV8Value, value *ICefV8Value, exception *ResultString) bool {
	fmt.Println("bindSubObjectSet accessor name:", name)
	return false
}

// bindGet 绑定字段取值
func (m *ipcRenderProcess) bindGet(name string, object *ICefV8Value, retVal *ResultV8Value, exception *ResultString) bool {
	fmt.Println("bindGet accessor name:", name)
	return false
}

// bindSet 绑定字段赋值
func (m *ipcRenderProcess) bindSet(name string, object *ICefV8Value, value *ICefV8Value, exception *ResultString) bool {
	fmt.Println("bindSet accessor name:", name)
	return false
}
