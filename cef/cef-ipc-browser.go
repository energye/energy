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
	"github.com/energye/energy/cef/ipc"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/json"
)

// ipcBrowserProcess 主进程
type ipcBrowserProcess struct {
}

// ipcGoExecuteMethodMessage 执行 Go 监听函数
func (m *ipcBrowserProcess) jsExecuteGoMethodMessage(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
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
		if argumentList != nil {
			argumentList.Free()
		}
		if argument != nil {
			argument.Free()
		}
	}()
	argumentListBytes = nil
	var ipcContext = m.jsExecuteGoMethod(browser.Identifier(), frame.Identifier(), emitName, argumentList)
	if messageId != 0 { // 异步回调函数处理
		replyMessage := json.NewJSONArray(nil)
		replyMessage.Add(messageId)
		replyMessage.Add(false)
		if ipcContext != nil {
			//处理回复消息
			replay := ipcContext.Replay()
			if replay.Result() != nil && len(replay.Result()) > 0 {
				switch replay.Result()[0].(type) {
				case []byte:
					replyMessage.SetByIndex(1, true)
					replyMessage.Add((replay.Result()[0]).([]byte))
				}
			}
		}
		frame.SendProcessMessageForJSONBytes(internalIPCJSExecuteGoEventReplay, consts.PID_RENDER, replyMessage.Bytes())
		replyMessage.Free()
	}
	if ipcContext != nil {
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
	return
}

// jsExecuteGoMethod 执行Go函数
func (m *ipcBrowserProcess) jsExecuteGoMethod(browserId int32, frameId int64, emitName string, argumentList json.JSONArray) ipc.IContext {
	eventCallback := ipc.CheckOnEvent(emitName)
	var ipcContext ipc.IContext
	if eventCallback != nil {
		ipcContext = ipc.NewContext(browserId, frameId, true, argumentList)
		//调用监听函数
		if ctxCallback := eventCallback.ContextCallback(); ctxCallback != nil {
			ctxCallback.Invoke(ipcContext)
		} else if argsCallback := eventCallback.ArgumentCallback(); argsCallback != nil {
			argsCallback.Invoke(ipcContext)
		}
	}
	return ipcContext
}

// ipcGoExecuteMethodMessageReply 执行Go函数回复结果
func (m *ipcBrowserProcess) ipcGoExecuteMethodMessageReply(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
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
		message.Free()
		if c == 0 {
			return
		}
	}
	var messageId int32
	var argument json.JSON
	var argumentList json.JSONArray
	if messageDataBytes != nil {
		argument = json.NewJSON(messageDataBytes)
		messageId = int32(argument.GetIntByKey(ipc_id))
		argumentList = argument.GetArrayByKey(ipc_argumentList)
		messageDataBytes = nil
	}
	defer func() {
		if argument != nil {
			argument.Free()
		}
	}()
	if callback := ipc.CheckEmitCallback(messageId); callback != nil {
		ipcContext := ipc.NewContext(browser.Identifier(), frame.Identifier(), false, argumentList)
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

// ipcChannelBrowser Go IPC 主进程监听
func (m *ipcBrowserProcess) initBrowserIPC() {
	browserIPC.addCallback(func(channelId int64, data json.JSON) bool {
		if data != nil && data.IsObject() {
			messageJSON := data.JSONObject()
			//messageId := messageJSON.GetIntByKey(ipc_id)// messageId: 同步永远是1
			emitName := messageJSON.GetStringByKey(ipc_event)
			name := messageJSON.GetStringByKey(ipc_name)
			browserId := messageJSON.GetIntByKey(ipc_browser_id)
			argumentList := messageJSON.GetArrayByKey(ipc_argumentList)
			if name == internalIPCJSExecuteGoSyncEvent {
				m.jsExecuteGoSyncMethodMessage(int32(browserId), channelId, emitName, argumentList)
				return true
			}
		}
		return false
	})

}

// jsExecuteGoSyncMethodMessage JS执行Go事件 - 同步消息处理
func (m *ipcBrowserProcess) jsExecuteGoSyncMethodMessage(browserId int32, frameId int64, emitName string, argumentList json.JSONArray) {
	var ipcContext = m.jsExecuteGoMethod(browserId, frameId, emitName, argumentList)
	message := json.NewJSONObject(nil)
	message.Set(ipc_id, 1)
	message.Set(ipc_name, internalIPCJSExecuteGoSyncEventReplay)
	message.Set(ipc_argumentList, nil)
	// 同步回调函数处理
	if ipcContext != nil {
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			switch replay.Result()[0].(type) {
			case []byte:
				message.Set(ipc_argumentList, json.NewJSONArray((replay.Result()[0]).([]byte)).Data())
			}
		}
	}
	//回复结果消息
	browserIPC.ipc.Send(frameId, message.Bytes())
	message.Free()
	if ipcContext != nil {
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
}
