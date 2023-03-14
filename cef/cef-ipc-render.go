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
	if m.onHandler != nil {
		m.onHandler.clear()
	}
}

// makeCtx ipc 和 bind
func (m *ipcRenderProcess) makeCtx(context *ICefV8Context) {
	m.makeIPC(context)
	m.makeBind(context)
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
		onName      *ICefV8Value //事件名
		onNameValue string       // 事件名
		onCallback  *ICefV8Value //事件回调函数
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
	onNameValue = onName.GetStringValue()

	//ipc on
	m.onHandler.addCallback(onNameValue, &ipcCallback{arguments: arguments, context: V8ContextRef.Current(), function: V8ValueRef.UnWrap(onCallback)})
	result = true
	return
}

// ipcGoExecuteJSEvent Go ipc.emit 执行JS事件
func (m *ipcRenderProcess) ipcGoExecuteJSEvent(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	argument := message.ArgumentList()
	name := argument.GetString(1)
	if callback := ipcRender.onHandler.getCallback(name); callback != nil {
		messageId := argument.GetInt(0)
		args := argument.GetBinary(2)
		var argsBytes []byte
		var count uint32
		var argsV8ValueArray *TCefV8ValueArray
		if args.IsValid() {
			size := args.GetSize()
			argsBytes = make([]byte, size)
			count = args.GetData(argsBytes, 0)
			args.Free()
		}
		if callback.context.Enter() {
			if count > 0 {
				argsV8ValueArray, _ = ipcValueConvert.BytesToV8ArrayValue(argsBytes)
			}
			ret := callback.function.ExecuteFunctionWithContext(callback.context, nil, argsV8ValueArray)
			if argsV8ValueArray != nil {
				argsV8ValueArray.Free()
			}
			ret.Free()
			callback.context.Exit()
		}
		argsBytes = nil
		if messageId != 0 { // callback

		}
		result = true
	}
	return
}

// ipcJSExecuteGoEvent JS ipc.emit 执行Go事件
func (m *ipcRenderProcess) ipcJSExecuteGoEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
	var (
		emitName       *ICefV8Value //事件名
		emitNameValue  string       //事件名
		emitArgs       *ICefV8Value //事件参数
		emitCallback   *ICefV8Value //事件回调函数
		ipcEmitMessage *ICefProcessMessage
		binaryValue    *ICefBinaryValue
		freeV8Value    = func(value *ICefV8Value) {
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
			arguments.Free()
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
			if ipcEmitMessage != nil {
				ipcEmitMessage.Free()
			}
			if binaryValue != nil {
				binaryValue.Free()
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
			args := ipcValueConvert.V8ValueToProcessMessageBytes(emitArgs)
			if args == nil {
				exception.SetValue("ipc.emit convert parameter to value value error")
				isFree = true
				return
			}
			binaryValue = BinaryValueRef.New(args)
			args = nil
		}
		context := V8ContextRef.Current()
		var messageId int32 = 0
		//回调函数
		if emitCallback != nil {
			//回调函数临时存放到缓存中
			messageId = ipcRender.emitHandler.addCallback(&ipcCallback{
				arguments: arguments,
				context:   context,
				function:  V8ValueRef.UnWrap(emitCallback),
			})
		}
		ipcEmitMessage = ProcessMessageRef.new(internalProcessMessageIPCEmit)
		argument := ipcEmitMessage.ArgumentList()
		argument.SetInt(0, messageId)        // 消息id
		argument.SetString(1, emitNameValue) // 事件名
		argument.SetBinary(2, binaryValue)   // args
		frame := context.Frame()
		frame.SendProcessMessage(consts.PID_BROWSER, ipcEmitMessage)
		frame.Free()
		retVal.SetResult(V8ValueRef.NewBool(true))
	}
	return
}

// ipcJSExecuteGoEventMessageReply JS执行Go监听，Go的消息回复
func (m *ipcRenderProcess) ipcJSExecuteGoEventMessageReply(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	result = true
	messageId := message.ArgumentList().GetInt(0)
	if callback := ipcRender.emitHandler.getCallback(messageId); callback != nil {
		//第二个参数 true 有返回参数
		if isReturn := message.ArgumentList().GetBool(1); isReturn {
			//[]byte
			binaryValue := message.ArgumentList().GetBinary(2)
			var count uint32
			var resultArgsBytes []byte
			if binaryValue.IsValid() {
				size := binaryValue.GetSize()
				resultArgsBytes = make([]byte, size)
				count = binaryValue.GetData(resultArgsBytes, 0)
				binaryValue.Free()
			}
			if count > 0 {
				//解析 '[]byte' 参数
				if callback.context.Enter() {
					if args, err := ipcValueConvert.BytesToV8ArrayValue(resultArgsBytes); err == nil {
						callback.function.ExecuteFunctionWithContext(callback.context, nil, args).Free()
						args.Free()
					} else {
						// parsing error
						callback.function.ExecuteFunctionWithContext(callback.context, nil, nil).Free()
					}
				}
				callback.context.Exit()
			}
			resultArgsBytes = nil
		} else { //无返回参数
			if callback.context.Enter() {
				callback.function.ExecuteFunctionWithContext(callback.context, nil, nil).Free()
			}
			callback.context.Exit()
		}
		//remove
		callback.free()
		ipcRender.emitHandler.removeCallback(messageId)
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
