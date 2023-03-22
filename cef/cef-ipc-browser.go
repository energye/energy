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
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/ipc/channel"
	"github.com/energye/energy/pkgs/json"
)

// ipcBrowserProcess 主进程
type ipcBrowserProcess struct {
	ipcObject   *ICefV8Value    // ipc object
	emitHandler *ipcEmitHandler // ipc.emit handler
	onHandler   *ipcOnHandler   // ipc.on handler
	ipcChannel  channel.IBrowserChannel
}

//func (m *ipcBrowserProcess) ipcChannelBrowser() {
//	if m.ipcChannel == nil {
//		m.ipcChannel = channel.NewBrowser()
//		m.ipcChannel.Handler(func(context channel.IIPCContext) {
//			var data ipcChannelMessage
//			if err := jsoniter.Unmarshal(context.Message().Data(), &data); err == nil {
//				//messageId := message.GetUIntByKey(ipc_id)
//				//name := message.GetStringByKey(ipc_name)
//				//message.Set(ipc_event, emitNameValue)
//				//message.Set(ipc_argumentList, json.NewJSONArray(args).Data())
//				fmt.Println("ipcChannelBrowser", err)
//				fmt.Println("data", data.Name, data.EventName, json.NewJSONArray(data.Data).GetIntByIndex(1))
//				if data.Name == internalIPCJSExecuteGoEvent {
//					//result = ipcBrowser.ipcGoExecuteMethodMessage(browser, frame, sourceProcess, message)
//				} else if data.Name == internalProcessMessageIPCOn {
//					//result = ipcBrowser.ipcOnMessage(browser, frame, sourceProcess, message)
//				} else if data.Name == internalProcessMessageIPCEmitReply {
//					//result = ipcBrowser.ipcGoExecuteMethodMessageReply(browser, frame, sourceProcess, message)
//				}
//			}
//			context.Free()
//		})
//	}
//}

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
	eventCallback := ipc.CheckOnEvent(emitName)
	if eventCallback == nil {
		return
	}
	ipcContext := ipc.NewContext(browser.Identifier(), frame.Identifier(), true, argumentList)
	//调用监听函数
	if ctxCallback := eventCallback.ContextCallback(); ctxCallback != nil {
		ctxCallback.Invoke(ipcContext)
	} else if argsCallback := eventCallback.ArgumentCallback(); argsCallback != nil {
		argsCallback.Invoke(ipcContext)
	}
	if messageId != 0 { // 回调函数处理
		replyMessage := json.NewJSONArray(nil)
		replyMessage.Add(messageId)
		replyMessage.Add(false)
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			switch replay.Result()[0].(type) {
			case []byte:
				replyMessage.SetByIndex(1, true)
				replyMessage.Add((replay.Result()[0]).([]byte))
			}
		}
		frame.SendProcessMessageForJSONBytes(internalIPCJSExecuteGoEventReplay, consts.PID_RENDER, replyMessage.Bytes())
		replyMessage.Free()
	}
	if ipcContext.ArgumentList() != nil {
		ipcContext.ArgumentList().Free()
	}
	ipcContext.Result(nil)

	return
}

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
