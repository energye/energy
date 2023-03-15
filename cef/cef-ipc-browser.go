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
)

// ipcBrowserProcess 主进程
type ipcBrowserProcess struct {
	ipc         *ICefV8Value    // ipc object
	emitHandler *ipcEmitHandler // ipc.emit handler
	onHandler   *ipcOnHandler   // ipc.on handler
}

func (m *ipcBrowserProcess) ipcGoExecuteMethodMessageReply(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	messageId := message.ArgumentList().GetInt(0)
	if callback := ipc.CheckEmitCallback(messageId); callback != nil {
		//第二个参数 true 有返回参数
		if isReturn := message.ArgumentList().GetBool(1); isReturn {
			//[]byte
			binaryValue := message.ArgumentList().GetBinary(2)
			var (
				count           uint32
				resultArgsBytes []byte
			)
			if binaryValue.IsValid() {
				size := binaryValue.GetSize()
				resultArgsBytes = make([]byte, size)
				count = binaryValue.GetData(resultArgsBytes, 0)
				binaryValue.Free()
			}
			if count > 0 {
				ipcContext := ipc.NewContext(browser.Identifier(), frame.Identifier(), false, resultArgsBytes)
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
			resultArgsBytes = nil
		}
	}
	return
}

// ipcGoExecuteMethodMessage 执行 Go 监听函数
func (m *ipcBrowserProcess) ipcGoExecuteMethodMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	result = true
	argument := message.ArgumentList()
	messageId := argument.GetInt(0)
	emitName := argument.GetString(1)
	eventCallback := ipc.CheckOnEvent(emitName)
	if eventCallback == nil {
		return
	}
	var argsBytes []byte
	//参数 字节数组
	args := argument.GetBinary(2)
	if args.IsValid() {
		size := args.GetSize()
		argsBytes = make([]byte, size)
		args.GetData(argsBytes, 0)
		args.Free() //立即释放掉
	}
	ipcContext := ipc.NewContext(browser.Identifier(), frame.Identifier(), true, argsBytes)
	argsBytes = nil
	//调用监听函数
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
