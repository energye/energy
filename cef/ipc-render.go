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
	"errors"
	"github.com/energye/energy/cef/internal/ipc"
	ipcArgument "github.com/energye/energy/cef/ipc/argument"
	"github.com/energye/energy/cef/ipc/context"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/json"
)

// ipcRenderProcess 渲染进程
type ipcRenderProcess struct {
	isInitRenderIPC bool
	ipcObject       *ICefV8Value    // ipc object
	emitHandler     *ipcEmitHandler // ipc.emit handler
	onHandler       *ipcOnHandler   // ipc.on handler
	syncChan        *ipc.SyncChan
	v8Context       *ICefV8Context
}

func (m *ipcRenderProcess) clear() {
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
	result = true
	argumentListBytes := message.ArgumentList().GetBinary(0)
	//接收二进制数据失败
	if argumentListBytes == nil {
		return
	}
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
	var argument ipcArgument.IList
	var argumentList json.JSONArray
	if messageDataBytes != nil {
		argument = ipcArgument.UnList(messageDataBytes)
		messageId = argument.MessageId()
		emitName = argument.GetEventName()
		if argument.JSON() != nil {
			argumentList = argument.JSON().JSONArray()
		}
		messageDataBytes = nil
	}
	defer func() {
		if argument != nil {
			argument.Reset()
		}
		if argumentList != nil {
			argumentList.Free()
		}
	}()

	if callback := ipcRender.onHandler.getCallback(emitName); callback != nil {
		var callbackArgsBytes any
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
				callbackArgsBytes = ValueConvert.V8ValueToProcessMessageArray(ret)
				ret.Free()
			} else if ret != nil {
				ret.Free()
			}
			//exit v8context
			m.v8Context.Exit()
		}
		if messageId != 0 { //messageId != 0 callback func
			callbackMessage := &ipcArgument.List{
				Id:   messageId,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCGoExecuteJSEventReplay,
			}
			if callbackArgsBytes != nil {
				callbackMessage.Data = callbackArgsBytes //json.NewJSONArray(callbackArgsBytes).Data()
			}
			// send ipc message
			// send bytes data to browser ipc
			ipc.RenderChan().IPC().Send(callbackMessage.Bytes())
			callbackMessage.Reset()
		}
	}
	return
}

// 规则选项名定义
const (
	ruleName      = "name"
	ruleArguments = "arguments"
	ruleCallback  = "callback"
	ruleMode      = "mode"
	ruleTarget    = "target"
)

// 触发模式
type modelRule int8

const (
	modelAsync modelRule = iota // 异步
	modelSync                   // 同步
)

// 触发接收目标
type targetRule int8

const (
	targetMain    targetRule = iota //主进程
	targetCurrent                   //当前进程
	targetOther                     //其它子进程
)

// 执行规则
type executeRule int8

const (
	exerOption executeRule = iota //选项
	exerArgs                      //参数
)

// jsExecuteRule
//	JS 执行规则
type jsExecuteRule struct {
	rule          executeRule  //规则
	model         modelRule    //模式 默认 modelAsync
	target        targetRule   //目标 默认 targetMain
	first         *ICefV8Value //
	emitName      *ICefV8Value //事件名
	emitNameValue string       //事件名值
	emitArgs      *ICefV8Value //事件参数
	emitCallback  *ICefV8Value //事件回调函数
}

// freeV8Value
func (*jsExecuteRule) freeV8Value(value *ICefV8Value) {
	if value != nil {
		value.Free()
	}
}

// freeAll
func (m *jsExecuteRule) freeAll() {
	m.freeV8Value(m.first)
	m.freeV8Value(m.emitName)
	m.freeV8Value(m.emitArgs)
	m.freeV8Value(m.emitCallback)
	m.emitNameValue = ""
}

// jsExecuteGoRule 执行规则
func (*ipcRenderProcess) jsExecuteGoRule(name string, arguments *TCefV8ValueArray) (*jsExecuteRule, error) {
	if arguments.Size() >= 1 {
		m := new(jsExecuteRule)
		m.first = arguments.Get(0)
		if m.first.IsString() { // args
			m.emitName = m.first
			m.emitNameValue = m.emitName.GetStringValue()
			if arguments.Size() == 2 {
				args2 := arguments.Get(1)
				if args2.IsArray() {
					m.emitArgs = args2
				} else if args2.IsFunction() {
					m.emitCallback = args2
				} else {
					m.freeAll()
					return nil, errors.New("ipc.emit second argument can only be a parameter or callback function")
				}
			} else if arguments.Size() == 3 {
				m.emitArgs = arguments.Get(1)
				m.emitCallback = arguments.Get(2)
				if !m.emitArgs.IsArray() || !m.emitCallback.IsFunction() {
					m.freeAll()
					return nil, errors.New("ipc.emit second argument can only be a input parameter. third parameter can only be a callback function")
				}
			}
			if name == internalIPCEmitSync { //同步
				m.model = modelSync
			} else {
				m.model = modelAsync // default
			}
			m.rule = exerArgs
			m.target = targetMain // default
			return m, nil
		} else if m.first.IsObject() { // options
			keys := m.first.GetKeys()
			keyMap := make(map[string]bool, keys.Count())
			for i := 0; i < keys.Count(); i++ {
				keyMap[keys.Get(i)] = true
			}
			keys.Free()
			if _, ok := keyMap[ruleName]; ok { // 事件名称
				m.emitName = m.first.getValueByKey(ruleName)
				if !m.emitName.IsString() {
					m.freeAll()
					return nil, errors.New("ipc.emit event name is incorrect, Pass as an string")
				}
				m.emitNameValue = m.emitName.GetStringValue()
			} else {
				m.freeAll()
				return nil, errors.New("ipc.emit event name is incorrect, Pass as an string")
			}
			if _, ok := keyMap[ruleArguments]; ok { // 参数列表
				m.emitArgs = m.first.getValueByKey(ruleArguments) // array
				if !m.emitArgs.IsArray() {
					m.freeAll()
					return nil, errors.New("ipc.emit event arguments is incorrect, Pass as an array")
				}
			}
			if _, ok := keyMap[ruleCallback]; ok { // 回调函数
				m.emitCallback = m.first.getValueByKey(ruleCallback) // function
				if !m.emitCallback.IsFunction() {
					m.freeAll()
					return nil, errors.New("ipc.emit event callback function is incorrect, Pass as an function")
				}
			}
			if _, ok := keyMap[ruleMode]; ok { // 触发模式
				mode := m.first.getValueByKey(ruleMode) // int 0:async or 1:sync, default 0:async
				if !mode.IsInt() {
					m.freeAll()
					return nil, errors.New("ipc.emit event mode is incorrect, Pass as an integer")
				}
				modeValue := modelRule(mode.GetIntValue())
				mode.Free()
				if modeValue == modelAsync || modeValue == modelSync {
					m.model = modeValue
				} else {
					m.model = modelAsync // default
				}
			} else {
				m.model = modelAsync // default
			}
			if _, ok := keyMap[ruleTarget]; ok { // 触发目标
				target := m.first.getValueByKey(ruleTarget) // int 0:main 1:current 2:other
				if !target.IsInt() {
					m.freeAll()
					return nil, errors.New("ipc.emit event target is incorrect, Pass as an integer")
				}
				targetValue := targetRule(target.GetIntValue())
				target.Free()
				if targetValue == targetMain || targetValue == targetCurrent || targetValue == targetOther {
					m.target = targetValue
				} else {
					m.target = targetMain // default
				}
			} else {
				m.target = targetMain // default
			}
			m.rule = exerOption
			return m, nil
		}
	}
	return nil, errors.New("ipc.emit first argument must be the event name or option rule")
}

// jsExecuteGoEvent
//
//  JS 执行 GO 监听事件
func (m *ipcRenderProcess) jsExecuteGoEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
	m.initEventGlobal()
	result = true
	var args any
	defer func() {
		if args != nil {
			args = nil
		}
	}()
	rule, err := m.jsExecuteGoRule(name, arguments)
	if err == nil {
		defer rule.freeAll()
		if rule.emitArgs != nil { //入参
			//V8Value 转换
			args = ValueConvert.V8ValueToProcessMessageArray(rule.emitArgs)
			if args == nil {
				exception.SetValue("ipc.emit convert parameter to value value error")
				return
			}
		}
		var isSync = rule.model == modelSync //同步
		//单进程 不通过进程消息, 全是同步
		if application.SingleProcess() {
			callback := &ipcCallback{}
			if rule.emitCallback != nil { //callback function
				callback.resultType = rt_function
				callback.function = rule.emitCallback
			} else { //variable
				callback.resultType = rt_variable
			}
			m.singleProcess(rule.emitNameValue, callback, args)
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
			// 同步 或 当前子进程
			if isSync || rule.target == targetCurrent {
				callback := &ipcCallback{isSync: true}
				// 回调函数或变量方式接收返回值, 优先回调函数
				if rule.emitCallback != nil { //callback function
					callback.resultType = rt_function
					callback.function = rule.emitCallback
				} else { //variable
					callback.resultType = rt_variable
				}
				//当前子进程
				if rule.target == targetCurrent {
					m.multiProcessCurrentProcess(rule.emitNameValue, callback, args)
				} else { //同步 - 主进程
					m.multiProcessSync(1, rule.emitNameValue, callback, args)
				}
				if callback.resultType == rt_variable {
					if callback.variable != nil {
						retVal.SetResult(callback.variable)
					} else {
						retVal.SetResult(V8ValueRef.NewBool(true))
					}
				} else {
					retVal.SetResult(V8ValueRef.NewBool(true))
				}
			} else { // 异步 - 主进程
				var messageId int32 = 0
				if rule.emitCallback != nil {
					rule.emitCallback.SetCanNotFree(true)
					callback := &ipcCallback{resultType: rt_function, function: V8ValueRef.UnWrap(rule.emitCallback)}
					messageId = m.emitHandler.addCallback(callback)
				}
				if success := m.multiProcessAsync(m.v8Context.Frame(), messageId, rule.emitNameValue, args); !success {
					//失败，释放回调函数
					rule.emitCallback.SetCanNotFree(false)
				}
				retVal.SetResult(V8ValueRef.NewBool(true))
			}
		}
	} else {
		// emit rule error
		println("js execute go event error:", err.Error())
	}
	return
}

// multiProcessCurrentProcess 多进程消息 -  当前进程
func (m *ipcRenderProcess) multiProcessCurrentProcess(emitName string, callback *ipcCallback, data any) {
	// 主进程
	eventCallback := ipc.CheckOnEvent(emitName)
	var ipcContext context.IContext
	if eventCallback != nil {
		ipcContext = context.NewContext(m.v8Context.Browser().Identifier(), m.v8Context.Frame().Identifier(), true, json.NewJSONArray(data))
		//调用监听函数
		if ctxCallback := eventCallback.ContextCallback(); ctxCallback != nil {
			ctxCallback.Invoke(ipcContext)
		} else if argsCallback := eventCallback.ArgumentCallback(); argsCallback != nil {
			argsCallback.Invoke(ipcContext)
		}
	}
	if ipcContext != nil && callback != nil {
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			m.executeCallbackFunction(true, callback, json.NewJSONArray(replay.Result()))
			return
		}
	}
	m.executeCallbackFunction(false, callback, nil)
}

// singleProcess 单进程
func (m *ipcRenderProcess) singleProcess(emitName string, callback *ipcCallback, data any) {
	if ipcBrowser == nil {
		return
	}
	// 主进程
	ipcContext := ipcBrowser.jsExecuteGoMethod(m.v8Context.Browser().Identifier(), m.v8Context.Frame().Identifier(), emitName, json.NewJSONArray(data))
	if ipcContext != nil && callback != nil {
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			m.executeCallbackFunction(true, callback, json.NewJSONArray(replay.Result()))
			return
		}
	}
	m.executeCallbackFunction(false, callback, nil)
}

// multiProcessSync 多进程消息 - 同步
func (m *ipcRenderProcess) multiProcessSync(messageId int32, emitName string, callback *ipcCallback, data any) {
	//延迟等待接收结果，默认5秒
	m.syncChan.DelayWaiting()
	message := &ipcArgument.List{
		Id:        messageId,
		BId:       ipc.RenderChan().BrowserId(),
		Name:      internalIPCJSExecuteGoSyncEvent,
		EventName: emitName,
		Data:      data,
	}
	//发送数据到主进程
	ipc.RenderChan().IPC().Send(message.Bytes())
	message.Reset()
	//同步等待结果 delayWaiting 自动结束
	resultData := <-m.syncChan.ResultSyncChan
	//接收成功，停止
	m.syncChan.Stop()
	var argumentList json.JSONArray
	if resultData != nil {
		argumentList = resultData.(json.JSONArray)
	}
	if argumentList != nil {
		m.executeCallbackFunction(true, callback, argumentList)
		argumentList.Free()
	} else {
		m.executeCallbackFunction(false, callback, nil)
	}
}

// multiProcessAsync 多进程消息 - 异步
func (m *ipcRenderProcess) multiProcessAsync(frame *ICefFrame, messageId int32, emitName string, data any) bool {
	if frame != nil {
		message := &ipcArgument.List{
			Id:        messageId,
			EventName: emitName,
			Data:      data,
		}
		frame.SendProcessMessageForJSONBytes(internalIPCJSExecuteGoEvent, consts.PID_BROWSER, message.Bytes())
		message.Reset()
		return true
	}
	return false
}

// ipcJSExecuteGoEventMessageReply JS执行Go监听，Go的消息回复
func (m *ipcRenderProcess) ipcJSExecuteGoEventMessageReply(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	result = true
	argumentListBytes := message.ArgumentList().GetBinary(0)
	if argumentListBytes == nil {
		return
	}
	var messageDataBytes []byte
	if argumentListBytes.IsValid() {
		size := argumentListBytes.GetSize()
		messageDataBytes = make([]byte, size)
		c := argumentListBytes.GetData(messageDataBytes, 0)
		argumentListBytes.Free()
		if c == 0 {
			return
		}
	}
	var (
		messageId    int32
		isReturnArgs bool
		argumentList ipcArgument.IList // json.JSONArray
	)
	if messageDataBytes != nil {
		argumentList = ipcArgument.UnList(messageDataBytes)
		//argumentList = json.NewJSONArray(messageDataBytes)
		messageId = argumentList.MessageId()
		isReturnArgs = argumentList.JSON() != nil // argumentList.GetBoolByIndex(1)
		messageDataBytes = nil
	}
	defer func() {
		if argumentList != nil {
			argumentList.Reset()
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
		if callback.function != nil {
			//设置允许释放
			callback.function.SetCanNotFree(false)
		}
		//[]byte
		returnArgs = argumentList.JSON().JSONArray()
		if returnArgs != nil {
			m.executeCallbackFunction(isReturnArgs, callback, returnArgs)
		} else {
			m.executeCallbackFunction(isReturnArgs, callback, nil)
		}
		callback.free()
	}
	return
}

// executeCallbackFunction 执行 v8 function 回调函数
func (m *ipcRenderProcess) executeCallbackFunction(isReturnArgs bool, callback *ipcCallback, returnArgs json.JSONArray) {
	if isReturnArgs { //有返回参数
		//enter v8context
		if m.v8Context.Enter() {
			var argsArray *TCefV8ValueArray
			var err error
			if callback.resultType == rt_function {
				if returnArgs != nil {
					//bytes to v8array value
					argsArray, err = ValueConvert.JSONArrayToV8ArrayValue(returnArgs)
				}
				if argsArray != nil && err == nil {
					// parse v8array success
					callback.function.ExecuteFunctionWithContext(m.v8Context, nil, argsArray).Free()
					argsArray.Free()
				} else {
					// parse v8array fail
					callback.function.ExecuteFunctionWithContext(m.v8Context, nil, nil).Free()
				}
			} else if callback.resultType == rt_variable {
				if returnArgs != nil && returnArgs.Size() > 0 {
					//bytes to v8array value
					argsArray, err = ValueConvert.JSONArrayToV8ArrayValue(returnArgs)
					if argsArray != nil && err == nil {
						if argsArray.Size() == 1 {
							callback.variable = argsArray.Get(0)
							callback.variable.SetCanNotFree(true)
						} else {
							callback.variable = V8ValueRef.NewArray(int32(argsArray.Size()))
							for i := 0; i < argsArray.Size(); i++ {
								callback.variable.SetValueByIndex(int32(i), argsArray.Get(i))
							}
						}
						argsArray.Free()
					} else {
						callback.variable = V8ValueRef.NewBool(true)
					}
				} else {
					callback.variable = V8ValueRef.NewBool(true)
				}
			}
			//exit v8context
			m.v8Context.Exit()
		}
	} else { //无返回参数
		if callback.resultType == rt_function {
			if m.v8Context.Enter() {
				callback.function.ExecuteFunctionWithContext(m.v8Context, nil, nil).Free()
				m.v8Context.Exit()
			}
		} else if callback.resultType == rt_variable {
			// null
		}
	}
}

// registerGoSyncReplayEvent Go IPC 渲染进程监听
func (m *ipcRenderProcess) registerGoSyncReplayEvent() {
	if m.isInitRenderIPC {
		return
	}
	m.isInitRenderIPC = true
	ipc.RenderChan().AddCallback(func(channelId int64, argument ipcArgument.IList) bool {
		if argument != nil {
			name := argument.GetName()
			if name == internalIPCJSExecuteGoSyncEventReplay {
				var argumentList json.JSONArray
				if argument.JSON() != nil {
					argumentList = argument.JSON().JSONArray()
				}
				m.ipcJSExecuteGoSyncEventMessageReply(argumentList)
				return true
			}
		}
		return false
	})

}

// ipcJSExecuteGoSyncEventMessageReply JS执行Go事件 - 同步回复接收
func (m *ipcRenderProcess) ipcJSExecuteGoSyncEventMessageReply(argumentList json.JSONArray) {
	if argumentList != nil {
		m.syncChan.ResultSyncChan <- argumentList
	} else {
		m.syncChan.ResultSyncChan <- nil
	}
}

// makeIPC ipc
func (m *ipcRenderProcess) makeIPC(context *ICefV8Context) {
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
