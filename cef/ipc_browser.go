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
	"github.com/energye/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/json"
)

// ipcBrowserProcess 主进程
type ipcBrowserProcess struct {
}

// 执行 Go 监听函数
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
		argumentListBytes.Free() //立即释放掉
		if c == 0 {
			return
		}
	}
	var messageId int32
	var emitName string
	var argument ipcArgument.IList // json.JSON
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
	argumentListBytes = nil
	browser = BrowserRef.UnWrap(browser) // 必须 Warp
	frame = FrameRef.UnWrap(frame)       // 必须 Warp
	browserID := browser.Identifier()
	frameID := frame.Identifier()
	go func() { // 开启线程, 异步执行
		defer func() {
			if argumentList != nil {
				argumentList.Free()
			}
			if argument != nil {
				argument.Reset()
			}
			frame.Free()
			browser.Free()
		}()
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
			if application.IsSpecVer49() {
				// CEF49
				browser.SendProcessMessageForJSONBytes(internalIPCJSExecuteGoEventReplay, consts.PID_RENDER, replyMessage.Bytes())
			} else {
				frame.SendProcessMessageForJSONBytes(internalIPCJSExecuteGoEventReplay, consts.PID_RENDER, replyMessage.Bytes())
			}
			replyMessage.Reset()
		}
		if ipcContext != nil {
			if ipcContext.ArgumentList() != nil {
				ipcContext.ArgumentList().Free()
			}
			ipcContext.Result(nil)
		}
	}()
	return
}

// 执行Go函数
func (m *ipcBrowserProcess) jsExecuteGoMethod(browserId int32, frameId int64, emitName string, argumentList json.JSONArray) context.IContext {
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

// 执行Go函数回复结果
func (m *ipcBrowserProcess) goExecuteMethodMessageReply(browserId int32, frameId int64, argument ipcArgument.IList) (result bool) {
	var messageId = argument.MessageId()
	var argumentList json.JSONArray
	if argument.JSON() != nil {
		argumentList = argument.JSON().JSONArray()
	}
	if callback := ipc.CheckEmitCallback(messageId); callback != nil {
		ipcContext := context.NewContext(browserId, frameId, false, argumentList)
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
	return
}

// Go IPC 事件监听
func (m *ipcBrowserProcess) registerEvent() {
	ipc.BrowserChan().AddCallback(func(channelId int64, argument ipcArgument.IList) bool {
		if argument != nil {
			name := argument.GetName()
			if name == internalIPCJSExecuteGoSyncEvent { //JS 同步事件
				m.jsExecuteGoSyncMethodMessage(argument.BrowserId(), channelId, argument.MessageId(), argument)
				return true
			} else if name == internalIPCGoExecuteJSEventReplay {
				ipcBrowser.goExecuteMethodMessageReply(argument.BrowserId(), channelId, argument)
				return true
			}
		}
		return false
	})
	// drag
	ipc.BrowserChan().AddCallback(func(channelId int64, argument ipcArgument.IList) bool {
		if argument != nil {
			if argument.GetName() == internalIPCDRAG {
				if wi := BrowserWindow.GetWindowInfo(argument.BrowserId()); wi != nil {
					if wi.IsLCL() {
						dataJSON := argument.JSON()
						if dataJSON != nil {
							bw := wi.AsLCLBrowserWindow().BrowserWindow()
							if bw.drag == nil {
								bw.drag = &drag{}
							}
							object := dataJSON.JSONObject()
							bw.drag.T = int8(object.GetIntByKey("T"))
							bw.drag.X = int32(object.GetIntByKey("X"))
							bw.drag.Y = int32(object.GetIntByKey("Y"))
							bw.drag.window = wi
						}
						RunOnMainThread(func() {
							wi.AsLCLBrowserWindow().BrowserWindow().doDrag()
						})
					}
				}
				return true
			}
		}
		return false
	})
}

// JS执行Go事件 - 同步消息处理
func (m *ipcBrowserProcess) jsExecuteGoSyncMethodMessage(browserId int32, frameId int64, messageId int32, argument ipcArgument.IList) {
	var argumentList json.JSONArray
	if argument.JSON() != nil {
		argumentList = argument.JSON().JSONArray()
	}
	var emitName = argument.GetEventName()
	var ipcContext = m.jsExecuteGoMethod(browserId, frameId, emitName, argumentList)
	message := &ipcArgument.List{
		Id:   messageId,
		Name: internalIPCJSExecuteGoSyncEventReplay,
	}
	// 同步回调函数处理
	if ipcContext != nil {
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			message.Data = replay.Result()
		}
	}
	//回复结果消息
	ipc.BrowserChan().IPC().Send(frameId, message.Bytes())
	message.Reset()
	if ipcContext != nil {
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
}
