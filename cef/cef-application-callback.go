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
	"errors"
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
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
	internalObjectRootName = "energy"     // GO 和 V8Value 绑定根对象名
	ctx                    *contextCreate //
	mRun                   *mainRun       //
)

type contextCreate struct {
	ipc            *ICefV8Value
	ipcEmitHandler *ICefV8Handler
	ipcOnHandler   *ICefV8Handler
	//ipcCallbackList *list.List
	ipcCallbackList      map[int32]*ipcCallback
	ipcCallbackMessageId int32
	ipcCallbackLock      sync.Mutex
	bind                 *ICefV8Value
}

type ipcCallback struct {
	arguments *TCefV8ValueArray
	context   *ICefV8Context
	function  *ICefV8Value
}

type mainRun struct {
}

func init() {
	if common.Args.IsMain() {
	} else if common.Args.IsRender() {
	}
	mRun = &mainRun{}
	ctx = &contextCreate{ipcCallbackList: make(map[int32]*ipcCallback, 1000)}
}

// isInternalKey 内部key不允许使用
func isInternalKey(key string) bool {
	return key == internalIPCKey || key == internalEmit || key == internalOn ||
		key == internalProcessMessageIPCEmit || key == internalProcessMessageIPCOn || key == internalProcessMessageIPCEmitReply
}

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	fmt.Println("appOnContextCreated-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue(), "browserId:", browser.Identifier(), "frameId:", frame.Identifier())
	ctx.makeCtx(context)
}

// renderProcessMessageReceived 渲染进程消息 - 默认实现
func renderProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	result = true
	if message.Name() == internalProcessMessageIPCEmitReply { //接收回复消息
		messageId := message.ArgumentList().GetInt(0)
		if callback := ctx.getIPCCallback(messageId); callback != nil {
			//第二个参数 true 有返回参数
			if isReturn := message.ArgumentList().GetBool(1); isReturn {
				//[]byte
				binaryValue := message.ArgumentList().GetBinary(2)
				size := binaryValue.GetSize()
				resultArgsBytes := make([]byte, size)
				count := binaryValue.GetData(resultArgsBytes, 0)
				binaryValue.Free()

				if count > 0 {
					//解析 '[]byte' 参数
					if callback.context.Enter() {
						if args, err := V8ValueConvert.BytesToV8ArrayValue(resultArgsBytes); err == nil {
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
			ctx.removeIPCCallback(messageId)
		}
	}
	message.Free()
	return
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
func (m *mainRun) ipcEmitMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	result = true
	argument := message.ArgumentList()
	messageId := argument.GetInt(0)
	emitName := argument.GetString(1)
	eventCallback := ipc.CheckOnEvent(emitName)
	if eventCallback == nil {
		return
	}
	isCallback := messageId != 0 //消息ID
	//参数 字节数组
	args := argument.GetBinary(2)
	size := args.GetSize()
	argsBytes := make([]byte, size)
	args.GetData(argsBytes, 0)
	args.Free() //立即释放掉
	ipcContext := ipc.NewContext(browser.Identifier(), frame.Identifier(), argsBytes)
	argsBytes = nil
	if ctxCallback := eventCallback.ContextCallback(); ctxCallback != nil {
		ctxCallback.Invoke(ipcContext)
	} else if argsCallback := eventCallback.ArgumentCallback(); argsCallback != nil {
		argsCallback.Invoke(ipcContext)
	}
	if isCallback {
		replyMessage := ProcessMessageRef.new(internalProcessMessageIPCEmitReply)
		replyMessage.ArgumentList().SetInt(0, messageId)
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			switch replay.Result()[0].(type) {
			case []byte:
				binaryValue := BinaryValueRef.New((replay.Result()[0]).([]byte))
				replyMessage.ArgumentList().SetBool(1, true)          //有返回值
				replyMessage.ArgumentList().SetBinary(2, binaryValue) //result []byte
			default:
				replyMessage.ArgumentList().SetBool(1, false) //无返回值
			}
		} else {
			replyMessage.ArgumentList().SetBool(1, false) //无返回值
		}
		frame.SendProcessMessage(consts.PID_RENDER, replyMessage)
		replay.Clear()
		replyMessage.Free()
	}
	ipcContext.ArgumentList().Free()
	ipcContext.Result(nil)
	return
}

// ipcOnMessage 监听事件
func (m *mainRun) ipcOnMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	fmt.Println("ipcOnMessage", message.Name(), message.ArgumentList().Size())
	return false
}

// makeCtx ipc 和 bind
func (m *contextCreate) makeCtx(context *ICefV8Context) {
	m.makeIPC(context)
	m.makeBind(context)
}

// makeIPC ipc
func (m *contextCreate) makeIPC(context *ICefV8Context) {
	// ipc emit
	m.ipcEmitHandler = V8HandlerRef.New()
	m.ipcEmitHandler.Execute(m.ipcEmitExecute)
	// ipc on
	m.ipcOnHandler = V8HandlerRef.New()
	m.ipcOnHandler.Execute(m.ipcOnExecute)
	// ipc object
	m.ipc = V8ValueRef.NewObject(nil)
	m.ipc.setValueByKey(internalEmit, V8ValueRef.newFunction(internalEmit, m.ipcEmitHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipc.setValueByKey(internalOn, V8ValueRef.newFunction(internalOn, m.ipcOnHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	// global to v8 ipc key
	context.Global().setValueByKey(internalIPCKey, m.ipc, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}

// ipcEmitExecute ipc.emit 执行
func (m *contextCreate) ipcEmitExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) (result bool) {
	result = true
	if name != internalEmit {
		return
	}
	var freeV8Value = func(value *ICefV8Value) {
		if value != nil {
			value.Free()
		}
	}
	if arguments.Size() >= 1 { // 1 ~ 3 个参数
		var (
			emitName       *ICefV8Value //事件名
			emitNameValue  string       //
			emitArgs       *ICefV8Value //事件参数
			emitCallback   *ICefV8Value //事件回调函数
			ipcEmitMessage *ICefProcessMessage
			binaryValue    *ICefBinaryValue
		)
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
			return
		}
		emitNameValue = emitName.GetStringValue()
		if arguments.Size() == 2 {
			args2 := arguments.Get(1)
			if args2.IsArray() {
				emitArgs = args2
			} else if args2.IsFunction() {
				emitCallback = args2
			} else {
				return
			}
		} else if arguments.Size() == 3 {
			emitArgs = arguments.Get(1)
			emitCallback = arguments.Get(2)
			if !emitArgs.IsArray() || !emitCallback.IsFunction() {
				return
			}
		}
		//入参
		if emitArgs != nil {
			//ipcEmitMessage.Free()
			//freeV8Value(emitCallback)
			//arguments.Free()
			//return
			//V8Value 转换
			args := V8ValueConvert.V8ValueToProcessMessageBytes(emitArgs)
			if args == nil {
				return
			}
			context := V8ContextRef.Current()
			var messageId int32 = 0
			//回调函数
			if emitCallback != nil {
				//回调函数临时存放到缓存中 list 队列
				messageId = ctx.addIPCCallback(&ipcCallback{
					arguments: arguments,
					context:   context,
					function:  V8ValueRef.UnWrap(emitCallback),
				})
			}
			ipcEmitMessage = ProcessMessageRef.new(internalProcessMessageIPCEmit)
			argument := ipcEmitMessage.ArgumentList()
			binaryValue = BinaryValueRef.New(args)
			argument.SetInt(0, messageId)        // 消息id
			argument.SetString(1, emitNameValue) // 事件名
			argument.SetBinary(2, binaryValue)   // args
			frame := context.Frame()
			frame.SendProcessMessage(consts.PID_BROWSER, ipcEmitMessage)
			args = nil
			frame.Free()
			//context.Free()
		}
		retVal.SetResult(V8ValueRef.NewBool(true))
	}
	return
}

//buildV8ArrayValueToListValue V8Value 转换 ListValue
func (m *contextCreate) buildV8ArrayValueToListValue(argumentList *ICefListValue, v8valueArray *ICefV8Value) error {
	if !v8valueArray.IsArray() {
		return errors.New("build process message error. Please pass in the array type")
	}
	argsLen := v8valueArray.GetArrayLength()
	for i := 0; i < argsLen; i++ {
		args := v8valueArray.GetValueByIndex(i)
		if args.IsString() {
			argumentList.SetString(uint32(i), args.GetStringValue())
		} else if args.IsInt() {
			argumentList.SetInt(uint32(i), args.GetIntValue())
		} else if args.IsUInt() {
			argumentList.SetInt(uint32(i), int32(args.GetUIntValue()))
		} else if args.IsDouble() {
			argumentList.SetDouble(uint32(i), args.GetDoubleValue())
		} else if args.IsBool() {
			argumentList.SetBool(uint32(i), args.GetBoolValue())
		} else if args.IsNull() || args.IsUndefined() {
			argumentList.SetNull(uint32(i))
		} else if args.IsArray() {
			arrayValue := ListValueRef.New()
			_ = m.buildV8ArrayValueToListValue(arrayValue, args)
			argumentList.SetList(uint32(i), arrayValue)
		} else if args.IsObject() {
			objectValue := DictionaryValueRef.New()
			_ = m.buildV8ObjectValueToDictionaryValue(objectValue, args)
			argumentList.SetDictionary(uint32(i), objectValue)
		} else {
			argumentList.SetNull(uint32(i))
		}
	}
	return nil
}

//buildV8ObjectValueToDictionaryValue V8Value 转换 DictionaryValue
func (m *contextCreate) buildV8ObjectValueToDictionaryValue(object *ICefDictionaryValue, v8valueObject *ICefV8Value) error {
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
			_ = m.buildV8ArrayValueToListValue(arrayValue, args)
			object.SetList(key, arrayValue)
		} else if args.IsObject() {
			objectValue := DictionaryValueRef.New()
			_ = m.buildV8ObjectValueToDictionaryValue(objectValue, args)
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
func (m *contextCreate) addIPCCallback(callback *ipcCallback) int32 {
	//return uintptr(unsafe.Pointer(m.ipcCallbackList.PushBack(callback)))
	m.ipcCallbackLock.Lock()
	defer m.ipcCallbackLock.Unlock()
	if m.ipcCallbackMessageId == -1 {
		m.ipcCallbackMessageId = 1
	} else {
		m.ipcCallbackMessageId++
	}
	m.ipcCallbackList[m.ipcCallbackMessageId] = callback
	return m.ipcCallbackMessageId
}

// removeIPCCallback
func (m *contextCreate) removeIPCCallback(messageId int32) {
	//m.ipcCallbackList.Remove((*list.Element)(unsafe.Pointer(ptr)))
	m.ipcCallbackLock.Lock()
	defer m.ipcCallbackLock.Unlock()
	delete(m.ipcCallbackList, messageId)
}

// getIPCCallback
func (m *contextCreate) getIPCCallback(messageId int32) *ipcCallback {
	//return (*list.Element)(unsafe.Pointer(ptr)).Value.(*ipcCallback)
	m.ipcCallbackLock.Lock()
	defer m.ipcCallbackLock.Unlock()
	return m.ipcCallbackList[messageId]
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
