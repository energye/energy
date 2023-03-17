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
	"fmt"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/pkgs/json"
)

// ipcBrowserProcess 主进程
type ipcBrowserProcess struct {
	ipc         *ICefV8Value    // ipc object
	emitHandler *ipcEmitHandler // ipc.emit handler
	onHandler   *ipcOnHandler   // ipc.on handler
}

func (m *ipcBrowserProcess) ipcGoExecuteMethodMessageReply(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	//messageId := message.ArgumentList().GetInt(0)
	//if callback := ipc.CheckEmitCallback(messageId); callback != nil {
	//	//第二个参数 true 有返回参数
	//	if isReturn := message.ArgumentList().GetBool(1); isReturn {
	//		//[]byte
	//		binaryValue := message.ArgumentList().GetBinary(2)
	//		var (
	//			count           uint32
	//			resultArgsBytes []byte
	//		)
	//		if binaryValue.IsValid() {
	//			size := binaryValue.GetSize()
	//			resultArgsBytes = make([]byte, size)
	//			count = binaryValue.GetData(resultArgsBytes, 0)
	//			binaryValue.Free()
	//		}
	//		if count > 0 {
	//			ipcContext := ipc.NewContext(browser.Identifier(), frame.Identifier(), false, resultArgsBytes)
	//			if ctxCallback := callback.ContextCallback(); ctxCallback != nil {
	//				ctxCallback.Invoke(ipcContext)
	//			} else if argsCallback := callback.ArgumentCallback(); argsCallback != nil {
	//				argsCallback.Invoke(ipcContext)
	//			}
	//			if ipcContext.ArgumentList() != nil {
	//				ipcContext.ArgumentList().Free()
	//			}
	//			ipcContext.Result(nil)
	//		}
	//		resultArgsBytes = nil
	//	}
	//}
	return
}

// ipcGoExecuteMethodMessage 执行 Go 监听函数
func (m *ipcBrowserProcess) ipcGoExecuteMethodMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	result = true
	argumentListBytes := message.ArgumentList().GetBinary(0)
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
	}
	defer func() {
		if argument != nil {
			argument.Free()
		}
		if argumentList != nil {
			argumentList.Free()
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
		//replyMessage := ProcessMessageRef.new(internalProcessMessageIPCEmitReply)
		//replyMessage.ArgumentList().SetInt(0, messageId)
		replyMessage := json.NewJSONArray(nil)
		replyMessage.Add(messageId)
		replyMessage.Add(false)
		//处理回复消息
		replay := ipcContext.Replay()
		if replay.Result() != nil && len(replay.Result()) > 0 {
			switch replay.Result()[0].(type) {
			case []byte:
				//binaryValue := BinaryValueRef.New((replay.Result()[0]).([]byte))
				//replyMessage.ArgumentList().SetBool(1, true)          //有返回值
				//replyMessage.ArgumentList().SetBinary(2, binaryValue) //result []byte
				replyMessage.SetByIndex(1, true)
				replyMessage.Add((replay.Result()[0]).([]byte))
			default:
				//replyMessage.ArgumentList().SetBool(1, false) //无返回值
			}
		} else {
			//replyMessage.ArgumentList().SetBool(1, false) //无返回值
		}
		//frame.SendProcessMessage(consts.PID_RENDER, replyMessage)
		frame.SendProcessMessageForJSONBytes(internalProcessMessageIPCEmitReply, consts.PID_RENDER, replyMessage.Bytes())
		replyMessage.Free()
		replay.Clear()
		//replyMessage.Free()
	}
	if ipcContext.ArgumentList() != nil {
		ipcContext.ArgumentList().Free()
	}
	ipcContext.Result(nil)
	return
}

// ipcOnMessage 监听事件
func (m *ipcBrowserProcess) ipcOnMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	fmt.Println("ipcOnMessage", message.Name(), message.ArgumentList().Size())
	return false
}
