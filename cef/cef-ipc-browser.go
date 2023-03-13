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
}

// ipcEmitMessage 触发事件
func (m *ipcBrowserProcess) ipcEmitMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
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
func (m *ipcBrowserProcess) ipcOnMessage(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	fmt.Println("ipcOnMessage", message.Name(), message.ArgumentList().Size())
	return false
}
