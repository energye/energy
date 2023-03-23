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
	"fmt"
	"github.com/energye/energy/pkgs/channel"
	"github.com/energye/energy/pkgs/json"
)

// renderIPCChan
type renderIPCChan struct {
	browserId int32
	frameId   int64
	ipc       channel.IRenderChannel //channel
}

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
			fmt.Println(messageId, name, argumentList)
			if name == internalIPCJSExecuteGoSyncEventReplay {
				m.ipcJSExecuteGoSyncEventMessageReply(int32(messageId), argumentList)
			}
			context.Free()
		})
	}
}

func (m *ipcRenderProcess) ipcJSExecuteGoSyncEventMessageReply(messageId int32, argumentList json.JSONArray) {
	if callback := m.emitHandler.getCallback(messageId); callback != nil {
		if callback.isSync {
			if argumentList != nil {
				callback.resultSyncChan <- argumentList.Bytes()
			} else {
				callback.resultSyncChan <- nil
			}
		}
		callback.free()
	}
}
