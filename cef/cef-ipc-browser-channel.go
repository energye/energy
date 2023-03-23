//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// GO browser IPC通道
package cef

import (
	"github.com/energye/energy/pkgs/channel"
	"github.com/energye/energy/pkgs/json"
)

// browserIPCChan
type browserIPCChan struct {
	ipc channel.IBrowserChannel
}

// ipcChannelBrowser Go IPC 主进程监听
func (m *ipcBrowserProcess) ipcChannelBrowser() {
	if m.ipcChannel == nil {
		m.ipcChannel = new(browserIPCChan)
		m.ipcChannel.ipc = channel.NewBrowser()
		m.ipcChannel.ipc.Handler(func(context channel.IIPCContext) {
			messageJSON := context.Message().JSON().JSONObject()
			messageId := messageJSON.GetIntByKey(ipc_id)
			emitName := messageJSON.GetStringByKey(ipc_event)
			name := messageJSON.GetStringByKey(ipc_name)
			browserId := messageJSON.GetIntByKey(ipc_browser_id)
			argumentList := messageJSON.GetArrayByKey(ipc_argumentList)
			if name == internalIPCJSExecuteGoSyncEvent {
				m.jsExecuteGoSyncMethodMessage(int32(messageId), int32(browserId), context.ChannelId(), emitName, argumentList)
			}
			context.Free()
		})
	}
}

// jsExecuteGoSyncMethodMessage JS执行Go事件 - 同步消息处理
func (m *ipcBrowserProcess) jsExecuteGoSyncMethodMessage(messageId, browserId int32, frameId int64, emitName string, argumentList json.JSONArray) {
	var ipcContext = m.jsExecuteGoMethod(browserId, frameId, emitName, argumentList)
	if messageId != 0 { // 同步回调函数处理
		message := json.NewJSONObject(nil)
		message.Set(ipc_id, messageId)
		message.Set(ipc_name, internalIPCJSExecuteGoSyncEventReplay)
		message.Set(ipc_argumentList, nil)
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
		m.ipcChannel.ipc.Send(frameId, message.Bytes())
		message.Free()
	}
	if ipcContext != nil {
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
}
