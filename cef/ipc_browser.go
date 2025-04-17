//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy 主进程 IPC

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/cyber-xxm/energy/v2/cef/ipc/argument"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
)

// ipcBrowserProcess 主进程
type ipcBrowserProcess struct{}

// 主进程消息 - 默认实现
func browserProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
	name := message.Name()
	if name == internalIPCJSEmit || name == internalIPCGoEmit {
		result = ipcBrowser.jsExecuteGoMethodMessage(browser, frame, message)
	} else if name == internalEnergyExtension {
		dragExtension.drag(browser, frame, message)
	} else if name == internalIPCJSEmitWait {
		ipcBrowser.jsExecuteGoWaitMethodMessage(browser, frame, message)
	} else if name == internalIPCGoEmitReplay {
		ipcBrowser.goExecuteMethodMessageReply(browser, frame, message)
	}
	return
}

// JS 执行 Go 监听函数
func (m *ipcBrowserProcess) jsExecuteGoMethodMessage(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
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
		argumentListBytes.Free() //释放掉
		if c == 0 {
			return
		}
	}
	var (
		messageId    int32
		emitName     string
		argument     ipcArgument.IList // json.JSON
		argumentList json.JSONArray    // []
		browserID    = browser.Identifier()
		frameID      = frame.Identifier()
	)
	if messageDataBytes != nil {
		argument = ipcArgument.UnList(messageDataBytes)
		messageId = argument.MessageId()
		emitName = argument.GetEventName()
		if argument.JSON() != nil {
			argumentList = argument.JSON().JSONArray()
		}
		messageDataBytes = nil
	}
	argumentListBytes = nil
	// 进程消息之后在合适的时机释放掉
	var free = func() {
		if argumentList != nil {
			argumentList.Free()
		}
		if argument != nil {
			argument.Reset()
		}
	}
	// 获取 Go 回调事件函数配置
	eventCallback := m.getJSExecuteGoEventCallback(emitName)
	if eventCallback != nil {
		// 执行函数
		var execute = func() {
			defer func() {
				free()
				if eventCallback.IsAsync {
					// free UnWarp
					frame.Free()
					browser.Free()
				}
			}()
			// 执行Go函数
			var ipcContext = m.jsExecuteGoMethod(browserID, frameID, emitName, argumentList)
			if messageId != 0 { // 异步回调函数处理
				replyMessage := &ipcArgument.List{
					Id: messageId,
				}
				if ipcContext != nil {
					//处理回复消息
					replay := ipcContext.Replay()
					if replay.Result() != nil && len(replay.Result()) > 0 {
						replyMessage.Data = replay.Result()
					}
				}
				if application.Is49() {
					// CEF49
					browser.SendProcessMessageForJSONBytes(internalIPCJSEmitReplay, consts.PID_RENDER, replyMessage.Bytes())
				} else {
					frame.SendProcessMessageForJSONBytes(internalIPCJSEmitReplay, consts.PID_RENDER, replyMessage.Bytes())
				}
				replyMessage.Reset()
			}
			if ipcContext != nil {
				if ipcContext.ArgumentList() != nil {
					ipcContext.ArgumentList().Free()
				}
				ipcContext.Result(nil)
			}
		}
		// 当前监听事件是异步，开启协程执行，但是CEF模式不能Debug协程（IDE无响应）
		if eventCallback.IsAsync {
			browser = BrowserRef.UnWrap(browser) // 必须 Warp
			frame = FrameRef.UnWrap(frame)       // 必须 Warp
			go execute()
		} else {
			// 同步执行, 默认这是使用CEF正确的方式
			execute()
		}
	} else {
		free()
	}
	return
}

// 获取 JS 执行 Go 回调函数
func (m *ipcBrowserProcess) getJSExecuteGoEventCallback(emitName string) *callback.Callback {
	return ipc.CheckOnEvent(emitName)
}

// JS 执行 Go 函数
func (m *ipcBrowserProcess) jsExecuteGoMethod(browserId int32, frameId string, emitName string, argumentList json.JSONArray) context.IContext {
	eventCallback := ipc.CheckOnEvent(emitName)
	var ipcContext context.IContext
	if eventCallback != nil {
		ipcContext = context.NewContext(browserId, frameId, true, argumentList)
		//调用监听函数
		if ctxCallback := eventCallback.ContextCallback(); ctxCallback != nil {
			ctxCallback.Invoke(ipcContext)
		} else if argsCallback := eventCallback.ArgumentCallback(); argsCallback != nil {
			argsCallback.Invoke(ipcContext)
		}
	}
	return ipcContext
}

// Go 执行 JS 函数, 处理 JS 回复结果
func (m *ipcBrowserProcess) goExecuteMethodMessageReply(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) {
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
	if callback := ipc.CheckEmitCallback(messageId); callback != nil {
		ipcContext := context.NewContext(browser.BrowserId(), frame.Identifier(), false, argumentList)
		if ctxCallback := callback.ContextCallback(); ctxCallback != nil {
			ctxCallback.Invoke(ipcContext)
		} else if argsCallback := callback.ArgumentCallback(); argsCallback != nil {
			argsCallback.Invoke(ipcContext)
		}
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
}

// JS 执行 Go 事件 - JS(渲染进程)等待 Go 并同时返回结果
func (m *ipcBrowserProcess) jsExecuteGoWaitMethodMessage(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) {
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
	emitName := argument.GetEventName()
	ipcContext := m.jsExecuteGoMethod(browser.BrowserId(), frame.Identifier(), emitName, argumentList)
	messageData := &ipcArgument.List{Id: argument.MessageId()}
	// 同步回调函数处理
	if ipcContext != nil {
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			messageData.Data = replay.Result()
		}
	}
	//回复结果消息
	var processMessage target.IProcessMessage
	if application.Is49() {
		// CEF49
		processMessage = browser
	} else {
		processMessage = frame
	}
	processMessage.SendProcessMessageForJSONBytes(internalIPCJSEmitWaitReplay, consts.PID_RENDER, messageData.Bytes())
	messageData.Reset()
	if ipcContext != nil {
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
}
