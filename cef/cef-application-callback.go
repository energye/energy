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

// contextCreate
type contextCreate struct {
	ipc     *ICefV8Value    // ipc object
	ipcEmit *ipcEmitHandler // ipc.emit handler
	ipcOn   *ipcOnHandler   // ipc.on handler
	bind    *ICefV8Value    // go bind
}

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

type mainRun struct {
}

func init() {
	if common.Args.IsMain() {
	} else if common.Args.IsRender() {
	}
	mRun = &mainRun{}
	ctx = &contextCreate{
		ipcEmit: &ipcEmitHandler{callbackList: make(map[int32]*ipcCallback, 256)},
		ipcOn:   &ipcOnHandler{callbackList: make(map[string]*ipcCallback)},
	}
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
		if callback := ctx.ipcEmit.getCallback(messageId); callback != nil {
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
			ctx.ipcEmit.removeCallback(messageId)
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
	if messageId != 0 { // 回调函数处理
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
	m.ipcEmit.handler = V8HandlerRef.New()
	m.ipcEmit.handler.Execute(m.ipcEmitExecute)
	// ipc on
	m.ipcOn.handler = V8HandlerRef.New()
	m.ipcOn.handler.Execute(m.ipcOnExecute)
	// ipc object
	m.ipc = V8ValueRef.NewObject(nil)
	m.ipc.setValueByKey(internalEmit, V8ValueRef.newFunction(internalEmit, m.ipcEmit.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipc.setValueByKey(internalOn, V8ValueRef.newFunction(internalOn, m.ipcOn.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	// global to v8 ipc key
	context.Global().setValueByKey(internalIPCKey, m.ipc, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}

// ipcEmitExecute ipc.on 执行
func (m *contextCreate) ipcOnExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) (result bool) {
	result = true
	if name != internalOn {
		retVal.SetResult(V8ValueRef.NewBool(false))
		return
	} else if arguments.Size() != 2 { //必须是2个参数
		retVal.SetResult(V8ValueRef.NewBool(false))
		exception.SetMessage("ipc.on parameter should be 2 quantity")
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
		exception.SetMessage("ipc.on event name should be a string")
		return
	}
	onCallback = arguments.Get(1)
	//第二个参数必须是函数
	if !onCallback.IsFunction() {
		retVal.SetResult(V8ValueRef.NewBool(false))
		exception.SetMessage("ipc.on event callback should be a function")
		return
	}
	onNameValue = onName.GetStringValue()
	if callback := m.ipcOn.getCallback(onNameValue); callback != nil {
		callback.free()
	}
	//ipc on
	m.ipcOn.addCallback(onNameValue, &ipcCallback{arguments: arguments, context: V8ContextRef.Current(), function: V8ValueRef.UnWrap(onCallback)})

	fmt.Println("on handler name:", name, "arguments-size:", arguments.Size())
	frame := V8ContextRef.Current().Frame()
	fmt.Println("frame", frame.Identifier())
	sendBrowserProcessMsg := ProcessMessageRef.New("ipcOnExecute")
	sendBrowserProcessMsg.ArgumentList().SetString(0, "ipcOnExecute测试值")
	frame.SendProcessMessage(consts.PID_BROWSER, sendBrowserProcessMsg)
	sendBrowserProcessMsg.Free()
	return
}

// ipcEmitExecute ipc.emit 执行
func (m *contextCreate) ipcEmitExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) (result bool) {
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
			exception.SetMessage("ipc.emit event name should be a string")
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
				exception.SetMessage("ipc.emit second argument can only be a parameter or callback function")
				isFree = true
				return
			}
		} else if arguments.Size() == 3 {
			emitArgs = arguments.Get(1)
			emitCallback = arguments.Get(2)
			if !emitArgs.IsArray() || !emitCallback.IsFunction() {
				exception.SetMessage("ipc.emit second argument can only be a input parameter. third parameter can only be a callback function")
				isFree = true
				return
			}
		}
		emitNameValue = emitName.GetStringValue()
		//入参
		if emitArgs != nil {
			//V8Value 转换
			args := V8ValueConvert.V8ValueToProcessMessageBytes(emitArgs)
			if args == nil {
				exception.SetMessage("ipc.emit convert parameter to message value error.")
				isFree = true
				return
			}
			context := V8ContextRef.Current()
			var messageId int32 = 0
			//回调函数
			if emitCallback != nil {
				//回调函数临时存放到缓存中
				messageId = ctx.ipcEmit.addCallback(&ipcCallback{
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
