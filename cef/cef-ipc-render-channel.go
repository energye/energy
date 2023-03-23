//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// GO render IPC通道
package cef

import (
	"github.com/energye/energy/pkgs/channel"
	"github.com/energye/energy/pkgs/json"
)

// renderIPCChan
type renderIPCChan struct {
	browserId int32
	frameId   int64
	ipc       channel.IRenderChannel //channel
}

// ipcChannelRender Go IPC 渲染进程监听
func (m *ipcRenderProcess) ipcChannelRender(browser *ICefBrowser, frame *ICefFrame) {
	if m.ipcChannel == nil {
		m.ipcChannel = new(renderIPCChan)
		m.ipcChannel.browserId = browser.Identifier()
		m.ipcChannel.frameId = frame.Identifier()
		m.ipcChannel.ipc = channel.NewRender(m.ipcChannel.frameId)
		m.ipcChannel.ipc.Handler(func(context channel.IIPCContext) {
			messageJSON := context.Message().JSON().JSONObject()
			messageId := messageJSON.GetIntByKey(ipc_id)
			name := messageJSON.GetStringByKey(ipc_name)
			argumentList := messageJSON.GetArrayByKey(ipc_argumentList)
			if name == internalIPCJSExecuteGoSyncEventReplay {
				m.ipcJSExecuteGoSyncEventMessageReply(int32(messageId), argumentList)
			}
			context.Free()
		})
	}
}

// ipcJSExecuteGoSyncEventMessageReply JS执行Go事件 - 同步回复接收
func (m *ipcRenderProcess) ipcJSExecuteGoSyncEventMessageReply(messageId int32, argumentList json.JSONArray) {
	if argumentList != nil {
		m.syncChan.resultSyncChan <- argumentList
	} else {
		m.syncChan.resultSyncChan <- nil
	}
}
