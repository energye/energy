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
	"github.com/cyber-xxm/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/cyber-xxm/energy/v2/cef/ipc/argument"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
	"time"
)

// 渲染进程消息 - 默认实现
func renderProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
	name := message.Name()
	if name == internalIPCJSEmitReplay {
		result = ipcRender.jsExecuteGoEventMessageReply(browser, frame, message)
	} else if name == internalIPCGoEmit {
		result = ipcRender.goExecuteJSEvent(browser, frame, message)
	} else if name == internalIPCJSEmitWaitReplay {
		ipcRender.jsExecuteGoSyncEventMessageReply(browser, frame, message)
	}
	return
}

// ipcRenderProcess 渲染进程
type ipcRenderProcess struct {
	isInitRenderIPC bool
	ipcObject       *ICefV8Value             // ipc object
	emitHandler     *ipcEmitHandler          // ipc.emit handler
	onHandler       map[string]*ipcOnHandler // ipc.on handler
	waitChan        *ipc.WaitChan
}

// JS ipc.on 监听事件
func (m *ipcOnHandler) jsOnEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
	var size int
	if name != internalIPCOn {
		return
	} else if size = arguments.Size(); !(size >= 2 && size <= 3) { //必须是2 | 3个参数
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
	//事件名，第一个参数必须是字符串
	if !onName.IsString() {
		exception.SetValue("ipc.on event name should be a string")
		arguments.Free()
		return
	}
	onCallback = arguments.Get(1)
	//回调函数，第二个参数必须是函数
	if !onCallback.IsFunction() {
		exception.SetValue("ipc.on event callback should be a function")
		arguments.Free()
		return
	}
	//监听选项, 第三个参数
	options := arguments.Get(2)
	if size == 3 && !options.IsObject() {
		exception.SetValue("ipc.on event options should be a object")
		arguments.Free()
		return
	}
	retVal.SetResult(V8ValueRef.NewBool(true))
	onCallback.SetCanNotFree(true)
	onNameValue = onName.GetStringValue()

	callback := &ipcCallback{
		function: V8ValueRef.UnWrap(onCallback),
		name:     V8ValueRef.UnWrap(onName),
	}

	// 监听模式，JS: ipc.on 监听事件 options 模式(mode)值是 MAsync(1)时，在回调函数的参数列表最后一个参数 complete 对象
	if size == 3 && options.IsValid() {
		modeValue := options.GetValueByKey("mode")
		if modeValue.IsValid() {
			callback.mode = types.Mode(modeValue.GetIntValue())
			modeValue.Free()
		}
		options.Free()
		if callback.mode == types.MAsync {
			callback.asyncHandler = newAsyncGoExecuteJSHandler()
		}
	}

	//ipc on, 添加到维护集合
	m.addCallback(onNameValue, callback)
	result = true
	return
}

// Go ipc.emit 执行JS事件
func (m *ipcRenderProcess) goExecuteJSEvent(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
	on, ok := m.onHandler[frame.Identifier()]
	if !ok {
		return false
	}
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

	if callback := on.getCallback(emitName); callback != nil {
		var callbackArgsBytes interface{}
		//enter v8context
		ctx := frame.V8Context()
		if ctx.Enter() {
			var ret *ICefV8Value
			var argsArray *TCefV8ValueArray
			var err error
			if argumentList != nil {
				//bytes to v8array value
				argsArray, err = ValueConvert.BytesToV8ArrayValue(argumentList.Bytes())
			}
			// MAsync: 异步模式
			if callback.mode == types.MAsync {
				if argsArray == nil {
					argsArray = V8ValueArrayRef.New()
				}
				complete := V8ValueRef.NewObject(nil, nil)
				complete.SetValueByKey("callback", callback.asyncHandler.callback, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
				complete.SetValueByKey("id", V8ValueRef.NewInt(messageId), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
				argsArray.Add(complete)
			}
			// 执行事件函数
			if argsArray != nil && err == nil {
				// parse v8array success
				ret = callback.function.ExecuteFunctionWithContext(ctx, nil, argsArray)
				argsArray.Free()
			} else {
				// parse v8array fail
				ret = callback.function.ExecuteFunctionWithContext(ctx, nil, nil)
			}
			// MSync: 同步模式
			if callback.mode == types.MSync && ret != nil && ret.IsValid() && messageId != 0 { // messageId != 0 callback func args
				// v8value to process message bytes
				callbackArgsBytes = ValueConvert.V8ValueToProcessMessageArray(ret)
				ret.Free()
			} else if ret != nil {
				ret.Free()
			}
			//exit v8context
			ctx.Exit()
		}
		ctx.Free()
		// MSync: 同步模式
		if callback.mode == types.MSync {
			var processMessage target.IProcessMessage
			if application.Is49() {
				// CEF49
				processMessage = browser
			} else {
				processMessage = frame
			}
			replayGoExecuteJSEvent(processMessage, messageId, callbackArgsBytes)
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
	ruleWaitTime  = "waitTime"
	ruleTarget    = "target"
)

// 触发模式
type modelRule int8

const (
	modelAsync modelRule = iota // 异步
	modelWait                   // 等待结果
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

// 触发模式是 modelWait 时的默认等待时间
const defaultEmitWaitTime = 5 * time.Second

// jsExecuteRule
//
//	JS 执行规则
type jsExecuteRule struct {
	rule          executeRule   //规则
	model         modelRule     //模式 默认 modelAsync
	waitTime      time.Duration //等待时间，当触发模式为 modelWait 有效, 默认5000毫秒
	target        targetRule    //目标 默认 targetMain
	first         *ICefV8Value  //
	emitName      *ICefV8Value  //事件名
	emitNameValue string        //事件名值
	emitArgs      *ICefV8Value  //事件参数
	emitCallback  *ICefV8Value  //事件回调函数
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

// 执行规则
func (*ipcRenderProcess) jsExecuteGoRule(name string, arguments *TCefV8ValueArray) (*jsExecuteRule, error) {
	if arguments.Size() >= 1 {
		m := new(jsExecuteRule)
		m.first = arguments.Get(0)
		m.waitTime = defaultEmitWaitTime // default
		if m.first.IsString() {          // args
			m.emitName = m.first
			m.emitNameValue = m.emitName.GetStringValue()
			if name == internalIPCEmitWait { // 等待超时返回结果
				m.model = modelWait
			} else {
				m.model = modelAsync // default
			}
			if arguments.Size() == 2 {
				args2 := arguments.Get(1)
				if args2.IsArray() {
					m.emitArgs = args2
				} else if m.model == modelAsync && args2.IsFunction() {
					m.emitCallback = args2
				} else if m.model == modelWait && args2.IsUInt() {
					// 等待超时触发
					waitTimeValue := args2.GetUIntValue()
					args2.Free()
					if waitTimeValue > 0 {
						m.waitTime = time.Duration(waitTimeValue) * time.Millisecond
					}
				} else {
					m.freeAll()
					if m.model == modelAsync {
						return nil, errors.New("ipc.emit second argument can only be a input parameter. third parameter can only be a callback function")
					} else {
						return nil, errors.New("ipc.emitWait second argument can only be a input parameter. third parameter can only be a timeout time")
					}
				}
			} else if arguments.Size() == 3 {
				m.emitArgs = arguments.Get(1)
				// 第三个参数, 根据触发模式
				args3 := arguments.Get(2)
				// 异步触发
				if m.model == modelAsync && args3.IsFunction() {
					m.emitCallback = args3
				} else if m.model == modelWait && args3.IsUInt() {
					// 等待超时触发
					waitTimeValue := args3.GetUIntValue()
					args3.Free()
					if waitTimeValue > 0 {
						m.waitTime = time.Duration(waitTimeValue) * time.Millisecond
					}
				} else {
					m.freeAll()
					if m.model == modelAsync {
						return nil, errors.New("ipc.emit second argument can only be a input parameter. third parameter can only be a callback function")
					} else {
						return nil, errors.New("ipc.emitWait second argument can only be a input parameter. third parameter can only be a timeout time")
					}
				}
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
			// 事件名称
			if _, ok := keyMap[ruleName]; ok {
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
			// 参数列表
			if _, ok := keyMap[ruleArguments]; ok {
				m.emitArgs = m.first.getValueByKey(ruleArguments) // array
				if !m.emitArgs.IsArray() {
					m.freeAll()
					return nil, errors.New("ipc.emit event arguments is incorrect, Pass as an array")
				}
			}
			// 回调函数
			if _, ok := keyMap[ruleCallback]; ok {
				m.emitCallback = m.first.getValueByKey(ruleCallback) // function
				if !m.emitCallback.IsFunction() {
					m.freeAll()
					return nil, errors.New("ipc.emit event callback function is incorrect, Pass as an function")
				}
			}
			// 触发模式
			if _, ok := keyMap[ruleMode]; ok {
				mode := m.first.getValueByKey(ruleMode) // int 0:async or 1:sync, default 0:async
				if !mode.IsInt() {
					m.freeAll()
					return nil, errors.New("ipc.emit event mode is incorrect, Pass as an integer")
				}
				modeValue := modelRule(mode.GetIntValue())
				mode.Free()
				if modeValue == modelAsync || modeValue == modelWait {
					m.model = modeValue
				} else {
					m.model = modelAsync // default
				}
			} else {
				m.model = modelAsync // default
			}
			// 触发等待时间，触发模式为 modelWait 有效
			if _, ok := keyMap[ruleWaitTime]; ok && m.model == modelWait {
				waitTime := m.first.getValueByKey(ruleWaitTime) // int 0:async or 1:sync, default 0:async
				if !waitTime.IsUInt() {
					m.freeAll()
					return nil, errors.New("ipc.emitWait The timeout waiting time is incorrect, Pass as an uint, unit millisecond")
				}
				// 单位 毫秒
				waitTimeValue := waitTime.GetUIntValue()
				waitTime.Free()
				if waitTimeValue > 0 {
					m.waitTime = time.Duration(waitTimeValue) * time.Millisecond
				}
			}
			// 触发目标
			if _, ok := keyMap[ruleTarget]; ok {
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

// JS 执行 GO 监听事件
func (m *ipcRenderProcess) jsExecuteGoEvent(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) (result bool) {
	result = true
	var args interface{}
	defer func() {
		if args != nil {
			args = nil
		}
	}()
	rule, err := m.jsExecuteGoRule(name, arguments)
	if err == nil {
		defer rule.freeAll()
		v8ctx := V8ContextRef.Current()
		defer v8ctx.Free()
		if rule.emitArgs != nil { //入参
			//V8Value 转换
			args = ValueConvert.V8ValueToProcessMessageArray(rule.emitArgs)
			if args == nil {
				exception.SetValue("ipc.emit convert parameter to value value error")
				return
			}
		}
		var isWait = rule.model == modelWait //等待超时返回结果
		//单进程 不通过进程消息, 全是同步
		if application.SingleProcess() {
			callback := &ipcCallback{}
			if rule.emitCallback != nil { //callback function
				callback.resultType = rt_function
				callback.function = rule.emitCallback
			} else { //variable
				callback.resultType = rt_variable
			}
			m.singleProcess(v8ctx, rule.emitNameValue, callback, args)
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
			// 等待超时 或 当前子进程
			if isWait || rule.target == targetCurrent {
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
					m.multiProcessCurrentProcess(v8ctx, rule.emitNameValue, callback, args)
				} else { //等待 - 主进程
					m.multiProcessWait(v8ctx, rule.emitNameValue, callback, rule.waitTime, args)
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
				var message target.IProcessMessage
				if application.Is49() {
					// CEF49
					message = v8ctx.Browser()
				} else {
					message = v8ctx.Frame()
				}
				// 在多进程时，发送多进程异步消息，如果当前CEF版本是CEF49使用Browser的发送消息函数
				if success := m.multiProcessAsync(message, messageId, rule.emitNameValue, args); !success {
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

// 多进程消息 -  当前进程
func (m *ipcRenderProcess) multiProcessCurrentProcess(v8ctx *ICefV8Context, emitName string, callback *ipcCallback, data interface{}) {
	// 主进程
	eventCallback := ipc.CheckOnEvent(emitName)
	var ipcContext context.IContext
	if eventCallback != nil {
		ipcContext = context.NewContext(v8ctx.Browser().Identifier(), v8ctx.Frame().Identifier(), true, json.NewJSONArray(data))
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
			m.executeCallbackFunction(v8ctx, true, callback, json.NewJSONArray(replay.Result()))
			return
		}
	}
	m.executeCallbackFunction(v8ctx, false, callback, nil)
}

// 单进程
func (m *ipcRenderProcess) singleProcess(v8ctx *ICefV8Context, emitName string, callback *ipcCallback, data interface{}) {
	if ipcBrowser == nil {
		return
	}
	// 当前为主进程
	ipcContext := ipcBrowser.jsExecuteGoMethod(v8ctx.Browser().Identifier(), v8ctx.Frame().Identifier(), emitName, json.NewJSONArray(data))
	if ipcContext != nil && callback != nil {
		// 处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			m.executeCallbackFunction(v8ctx, true, callback, json.NewJSONArray(replay.Result()))
			return
		}
	}
	m.executeCallbackFunction(v8ctx, false, callback, nil)
}

// 多进程消息 - 会阻塞渲染进程，并等待 delay 时间后自动返回
func (m *ipcRenderProcess) multiProcessWait(v8ctx *ICefV8Context, emitName string, callback *ipcCallback, delay time.Duration, data interface{}) {
	//延迟等待接收结果，默认5秒
	messageId := m.waitChan.NextMessageId()
	resultChan := make(chan interface{})
	m.waitChan.Pending.Store(messageId, func(result interface{}) {
		select {
		case resultChan <- result:
		}
	})
	// 组装消息
	message := &ipcArgument.List{
		Id:        messageId,
		EventName: emitName,
		Data:      data,
	}
	//发送消息到主进程
	//ipc.RenderChan().IPC().Send(message.Bytes())
	var processMessage target.IProcessMessage
	if application.Is49() {
		// CEF49
		processMessage = v8ctx.Browser()
	} else {
		processMessage = v8ctx.Frame()
	}
	processMessage.SendProcessMessageForJSONBytes(internalIPCJSEmitWait, consts.PID_BROWSER, message.Bytes())
	message.Reset()
	// 开始计时, 直到 delay 时间结束，或正常返回结果
	delayer := m.waitChan.NewDelayer(messageId, delay)
	select {
	case resultData := <-resultChan:
		//接收成功，停止
		delayer.Stop()
		delayer = nil
		var argumentList json.JSONArray
		if resultData != nil {
			argumentList = resultData.(json.JSONArray)
		}
		if argumentList != nil {
			m.executeCallbackFunction(v8ctx, true, callback, argumentList)
			argumentList.Free()
		} else {
			m.executeCallbackFunction(v8ctx, false, callback, nil)
		}
	}

}

// 多进程消息 - 异步
func (m *ipcRenderProcess) multiProcessAsync(processMessage target.IProcessMessage, messageId int32, emitName string, data interface{}) bool {
	if processMessage != nil {
		message := &ipcArgument.List{
			Id:        messageId,
			EventName: emitName,
			Data:      data,
		}
		processMessage.SendProcessMessageForJSONBytes(internalIPCJSEmit, consts.PID_BROWSER, message.Bytes())
		message.Reset()
		return true
	}
	return false
}

// JS执行Go监听，Go的消息回复
func (m *ipcRenderProcess) jsExecuteGoEventMessageReply(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
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
		messageId = argumentList.MessageId()
		isReturnArgs = argumentList.JSON() != nil
		messageDataBytes = nil
	}
	defer func() {
		if argumentList != nil {
			argumentList.Reset()
		}
	}()
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
		if isReturnArgs {
			//[]byte
			returnArgs = argumentList.JSON().JSONArray()
		}
		v8ctx := frame.V8Context()
		if returnArgs != nil {
			m.executeCallbackFunction(v8ctx, isReturnArgs, callback, returnArgs)
		} else {
			m.executeCallbackFunction(v8ctx, isReturnArgs, callback, nil)
		}
		callback.function.Free()
		callback.name.Free()
		v8ctx.Free()
	}
	return
}

// 执行 v8 function 回调函数
func (m *ipcRenderProcess) executeCallbackFunction(v8ctx *ICefV8Context, isReturnArgs bool, callback *ipcCallback, returnArgs json.JSONArray) {
	if isReturnArgs { //有返回参数
		//enter v8context
		if v8ctx.Enter() {
			var argsArray *TCefV8ValueArray
			var err error
			if callback.resultType == rt_function {
				if returnArgs != nil {
					//bytes to v8array value
					argsArray, err = ValueConvert.JSONArrayToV8ArrayValue(returnArgs)
				}
				if argsArray != nil && err == nil {
					// parse v8array success
					callback.function.ExecuteFunctionWithContext(v8ctx, nil, argsArray).Free()
					argsArray.Free()
				} else {
					// parse v8array fail
					callback.function.ExecuteFunctionWithContext(v8ctx, nil, nil).Free()
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
			v8ctx.Exit()
		}
	} else { //无返回参数
		if callback.resultType == rt_function {
			if v8ctx.Enter() {
				callback.function.ExecuteFunctionWithContext(v8ctx, nil, nil).Free()
				v8ctx.Exit()
			}
		} else if callback.resultType == rt_variable {
			// null
		}
	}
}

// JS执行Go事件 - 同步回复接收
func (m *ipcRenderProcess) jsExecuteGoSyncEventMessageReply(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) {
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
	var argument ipcArgument.IList // json.JSONArray
	if messageDataBytes != nil {
		argument = ipcArgument.UnList(messageDataBytes)
		messageDataBytes = nil
	}
	var argumentList json.JSONArray
	if argument.JSON() != nil {
		argumentList = argument.JSON().JSONArray()
	}
	messageId := argument.MessageId()
	if argument.JSON() != nil {
		argumentList = argument.JSON().JSONArray()
	}
	m.waitChan.Done(messageId, argumentList)
}

// ipc
func (m *ipcRenderProcess) makeIPC(frameId string, context *ICefV8Context) {
	// ipc emit
	m.emitHandler.handler = V8HandlerRef.New()
	m.emitHandler.handler.Execute(m.jsExecuteGoEvent)

	// ipc emit sync
	m.emitHandler.handlerSync = V8HandlerRef.New()
	m.emitHandler.handlerSync.Execute(m.jsExecuteGoEvent)

	// ipc on
	var on *ipcOnHandler
	if handler, ok := m.onHandler[frameId]; ok {
		on = handler
	} else {
		on = &ipcOnHandler{handler: V8HandlerRef.New(), callbackList: make(map[string]*ipcCallback)}
		on.handler.Execute(on.jsOnEvent)
		m.onHandler[frameId] = on
	}

	// ipc object
	m.ipcObject = V8ValueRef.NewObject(nil, nil)
	m.ipcObject.setValueByKey(internalIPCEmit, V8ValueRef.newFunction(internalIPCEmit, m.emitHandler.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipcObject.setValueByKey(internalIPCEmitWait, V8ValueRef.newFunction(internalIPCEmitWait, m.emitHandler.handlerSync), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipcObject.setValueByKey(internalIPCOn, V8ValueRef.newFunction(internalIPCOn, on.handler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)

	// ipc key to v8 global
	context.Global().setValueByKey(internalIPC, m.ipcObject, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}
