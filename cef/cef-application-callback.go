//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// application event 默认事件实现
package cef

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/types"
	"strconv"
	"unsafe"
)

const (
	internalIPCKey = "ipc"  // JavaScript -> ipc 事件驱动, 根对象名
	internalEmit   = "emit" // JavaScript -> ipc.emit 在 JavaScript 触发 GO 监听事件函数名
	internalOn     = "on"   // JavaScript -> ipc.on 在 JavaScript 监听事件, 提供给 GO 调用
)
const (
	internalProcessMessageIPCEmit = "ipcEmit" // 进程消息 ipcEmit
	internalProcessMessageIPCOn   = "ipcOn"   // 进程消息 ipcOn
)

var (
	internalObjectRootName = "energy"     // GO 和 V8Value 绑定根对象名, 可修改
	ctx                    *contextCreate //
	mRun                   *mainRun       //
)

type contextCreate struct {
	ipc             *ICefV8Value
	bind            *ICefV8Value
	ipcCallbackList *list.List
}

type ipcCallback struct {
	context  *ICefV8Context
	function *ICefV8Value
}

type mainRun struct {
}

func init() {
	if common.Args.IsMain() {
		mRun = &mainRun{}
	} else if common.Args.IsRender() {
		ctx = &contextCreate{ipcCallbackList: list.New()}
	}
}

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	fmt.Println("appOnContextCreated-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue(), "browserId:", browser.Identifier(), "frameId:", frame.Identifier())
	ctx.makeCtx(context)
}

// renderProcessMessageReceived 渲染进程消息 - 默认实现
func renderProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	fmt.Println("renderProcessMessageReceived", message.Name(), message.ArgumentList().Size())
	messageId := message.ArgumentList().GetString(0)
	fmt.Println("messageId", messageId)
	if v, err := strconv.Atoi(messageId); err == nil {
		callback := ctx.getIPCCallback(uintptr(v))
		if callback.context.Enter() {
			args := V8ValueArrayRef.New()
			args.Add(V8ValueRef.NewString("结果"))
			args.Add(V8ValueRef.NewDouble(10000.999))
			args.Add(V8ValueRef.NewString("结果"))
			args.Add(V8ValueRef.NewBool(true))
			fmt.Println("返回结果", args.Size())
			callback.function.ExecuteFunctionWithContext(callback.context, nil, args)
			callback.context.Exit()
		}
		ctx.removeIPCCallback(uintptr(v))
	}
	return false
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	fmt.Println("appMainRunCallback-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue())
	//internalBrowserIPCOnEventInit()
	//ipc.IPC.StartBrowserIPC()
	//indGoToJS(nil, nil)
}

// mainProcessMessageReceived 主进程消息 - 默认实现
func mainProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	if message.Name() == internalProcessMessageIPCEmit {
		return mRun.ipcEmitMessage(browser, frame, sourceProcess, message)
	} else if message.Name() == internalProcessMessageIPCOn {
		mRun.ipcOnMessage(browser, frame, sourceProcess, message)
	}
	return false
}

// ipcEmitMessage 触发事件
func (m *mainRun) ipcEmitMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	fmt.Println("ipcEmitMessage browserId:", browser.Identifier(), frame.Identifier(), "name:", message.Name(), message.ArgumentList().Size())
	argument := message.ArgumentList()
	messageId := argument.GetString(0)
	emitName := argument.GetString(1)
	callback := ipc.CheckOnEvent(emitName)
	if callback == nil {
		return false
	}
	args := argument.GetList(2)
	frame.Identifier()
	callback(ipc.NewContext(browser.Identifier(), frame.Identifier(), args))
	fmt.Println("\t", messageId, emitName)
	if messageId != "0" {
		frame.SendProcessMessage(consts.PID_RENDER, message)
	}
	return true
}

// ipcOnMessage 监听事件
func (m *mainRun) ipcOnMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	fmt.Println("ipcOnMessage", message.Name(), message.ArgumentList().Size())
	return false
}

// makeCtx ipc 和 bind
func (m *contextCreate) makeCtx(context *ICefV8Context) {
	ctx.makeIPC(context)
	ctx.makeBind(context)
}

// makeIPC ipc
func (m *contextCreate) makeIPC(context *ICefV8Context) {
	// ipc emit
	emitHandler := V8HandlerRef.New()
	emitHandler.Execute(m.ipcEmitExecute)
	// ipc on
	onHandler := V8HandlerRef.New()
	onHandler.Execute(m.ipcOnExecute)
	// ipc object
	m.ipc = V8ValueRef.NewObject(nil)
	m.ipc.setValueByKey(internalEmit, V8ValueRef.newFunction(internalEmit, emitHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipc.setValueByKey(internalOn, V8ValueRef.newFunction(internalOn, onHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	// global to v8 ipc key
	context.Global().setValueByKey(internalIPCKey, m.ipc, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}

// ipcEmitExecute ipc.emit 执行
func (m *contextCreate) ipcEmitExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("emit handler name:", name, "arguments-size:", arguments.Size())
	keyList := object.GetKeys()
	for i := 0; i < keyList.Count(); i++ {
		fmt.Println("\tkey:", i, keyList.Get(i))
	}
	if name != internalEmit {
		return false
	}
	if arguments.Size() >= 1 { // 1 ~ 3 个参数
		var (
			emitName      *ICefV8Value //事件名
			emitNameValue string       //
			emitArgs      *ICefV8Value //事件参数
			emitCallback  *ICefV8Value //事件回调函数
		)
		emitName = arguments.Get(0)
		if !emitName.IsString() {
			exception.SetMessage("ipc emit event parameter error. Parameter 1 can only be the event name")
			return false
		}
		emitNameValue = emitName.GetStringValue()
		if arguments.Size() == 2 {
			args2 := arguments.Get(1)
			if args2.IsArray() {
				emitArgs = args2
			} else if args2.IsFunction() {
				emitCallback = args2
			} else {
				exception.SetMessage("ipc emit event parameter error. Parameter 2 can only be an input parameter or callback function")
				return false
			}
		} else if arguments.Size() == 3 {
			emitArgs = arguments.Get(1)
			emitCallback = arguments.Get(2)
			if !emitArgs.IsArray() || !emitCallback.IsFunction() {
				exception.SetMessage("ipc emit event parameter error. Parameter 2 can only be an input parameter, Parameter 3 can only be a callback function")
				return false
			}
		}
		//入参
		if emitArgs != nil {
			ipcEmitMessage := ProcessMessageRef.new(internalProcessMessageIPCEmit)
			argument := ipcEmitMessage.ArgumentList()
			args := ListValueRef.New()
			defer func() {
				ipcEmitMessage.Free()
			}()
			if err := m.buildProcessMessageByV8ValueArray(args, emitArgs); err != nil {
				exception.SetMessage(fmt.Sprintf("ipc emit event parameter error.\n%v", err.Error()))
				return false
			}
			v8ctx := V8ContextRef.Current()
			var msgId uintptr = 0
			//回调函数
			if emitCallback != nil {
				//回调函数临时存放到缓存中 list 队列
				msgId = ctx.addIPCCallback(&ipcCallback{
					context:  v8ctx,
					function: V8ValueRef.UnWrap(emitCallback),
				})
			}
			argument.SetString(0, strconv.Itoa(int(msgId))) // 消息id
			argument.SetString(1, emitNameValue)            // 事件名
			argument.SetList(2, args)                       // args
			frame := v8ctx.Frame()
			frame.SendProcessMessage(consts.PID_BROWSER, ipcEmitMessage)
		}
		return true
	}
	exception.SetMessage("ipc emit event parameter error. Parameter 1 can only be the event name")
	return false
}

//buildProcessMessageByV8ValueArray 使V8ValueArray构建进程消息
func (m *contextCreate) buildProcessMessageByV8ValueArray(argumentList *ICefListValue, v8valueArray *ICefV8Value) error {
	if !v8valueArray.IsArray() {
		return errors.New("build process message error. Please pass in the array type")
	}
	argsLen := v8valueArray.GetArrayLength()
	for i := 0; i < argsLen; i++ {
		args := v8valueArray.GetValueByIndex(i)
		if args.IsString() {
			argumentList.SetString(types.NativeUInt(i), args.GetStringValue())
		} else if args.IsInt() {
			argumentList.SetInt(types.NativeUInt(i), args.GetIntValue())
		} else if args.IsUInt() {
			argumentList.SetInt(types.NativeUInt(i), int32(args.GetUIntValue()))
		} else if args.IsDouble() {
			argumentList.SetDouble(types.NativeUInt(i), args.GetDoubleValue())
		} else if args.IsBool() {
			argumentList.SetBool(types.NativeUInt(i), args.GetBoolValue())
		} else if args.IsNull() || args.IsUndefined() {
			argumentList.SetNull(types.NativeUInt(i))
		} else if args.IsArray() {
			arrayValue := ListValueRef.New()
			m.buildProcessMessageByV8ValueArray(arrayValue, args)
			argumentList.SetList(types.NativeUInt(i), arrayValue)
		} else if args.IsObject() {
			objectValue := DictionaryValueRef.New()
			m.buildProcessMessageByV8ValueObject(objectValue, args)
			argumentList.SetDictionary(types.NativeUInt(i), objectValue)
		} else {
			argumentList.SetNull(types.NativeUInt(i))
		}
	}
	return nil
}

//buildProcessMessageByV8ValueObject 使V8ValueObject构建进程消息
func (m *contextCreate) buildProcessMessageByV8ValueObject(object *ICefDictionaryValue, v8valueObject *ICefV8Value) error {
	if !v8valueObject.IsObject() {
		return errors.New("build process message error. Please pass in the object type")
	}
	keys := v8valueObject.GetKeys()
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		args := v8valueObject.GetValueByKey(key)
		if args.IsString() {
			object.SetString(key, args.GetStringValue())
		} else if args.IsInt() {
			object.SetInt(key, args.GetIntValue())
		} else if args.IsUInt() {
			object.SetInt(key, int32(args.GetUIntValue()))
		} else if args.IsDouble() {
			object.SetDouble(key, args.GetDoubleValue())
		} else if args.IsBool() {
			object.SetBool(key, args.GetBoolValue())
		} else if args.IsNull() || args.IsUndefined() {
			object.SetNull(key)
		} else if args.IsArray() {
			arrayValue := ListValueRef.New()
			m.buildProcessMessageByV8ValueArray(arrayValue, args)
			object.SetList(key, arrayValue)
		} else if args.IsObject() {
			objectValue := DictionaryValueRef.New()
			m.buildProcessMessageByV8ValueObject(objectValue, args)
			object.SetDictionary(key, objectValue)
		} else {
			object.SetNull(key)
			continue
		}
	}
	return nil
}

// ipcEmitExecute ipc.on 执行
func (m *contextCreate) ipcOnExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("on handler name:", name, "arguments-size:", arguments.Size())
	frame := V8ContextRef.Current().Frame()
	fmt.Println("frame", frame.Identifier())
	sendBrowserProcessMsg := ProcessMessageRef.New("ipcOnExecute")
	sendBrowserProcessMsg.ArgumentList().SetString(0, "ipcOnExecute测试值")
	frame.SendProcessMessage(consts.PID_BROWSER, sendBrowserProcessMsg)
	sendBrowserProcessMsg.Free()
	return false
}

// makeBind bind object accessor
func (m *contextCreate) makeBind(context *ICefV8Context) {
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
func (m *contextCreate) bindFuncExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("bindFuncExecute handler name:", name)
	return false
}

// bindSubObjectGet 绑定对象取值
func (m *contextCreate) bindSubObjectGet(name string, object *ICefV8Value, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("bindSubObjectGet accessor name:", name)
	return false
}

// bindSubObjectSet 绑定对象赋值
func (m *contextCreate) bindSubObjectSet(name string, object *ICefV8Value, value *ICefV8Value, exception *Exception) bool {
	fmt.Println("bindSubObjectSet accessor name:", name)
	return false
}

// bindGet 绑定字段取值
func (m *contextCreate) bindGet(name string, object *ICefV8Value, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("bindGet accessor name:", name)
	return false
}

// bindSet 绑定字段赋值
func (m *contextCreate) bindSet(name string, object *ICefV8Value, value *ICefV8Value, exception *Exception) bool {
	fmt.Println("bindSet accessor name:", name)
	return false
}

// addIPCCallback
func (m *contextCreate) addIPCCallback(callback *ipcCallback) uintptr {
	return uintptr(unsafe.Pointer(m.ipcCallbackList.PushBack(callback)))
}

// removeIPCCallback
func (m *contextCreate) removeIPCCallback(ptr uintptr) {
	m.ipcCallbackList.Remove((*list.Element)(unsafe.Pointer(ptr)))
}

// getIPCCallback
func (m *contextCreate) getIPCCallback(ptr uintptr) *ipcCallback {
	return (*list.Element)(unsafe.Pointer(ptr)).Value.(*ipcCallback)
}
