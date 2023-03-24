//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 基于IPC的字段数据绑定 - 主进程
package cef

import "github.com/energye/energy/pkgs/json"

var bindBrowser *bindBrowserProcess

type bindBrowserProcess struct {
}

func (m *bindBrowserProcess) initBrowserIPC() {
	browserIPC.addCallback(func(channelId int64, data json.JSON) bool {
		if data != nil && data.IsObject() {
			//messageJSON := data.JSONObject()
			//messageId := messageJSON.GetIntByKey(ipc_id)// messageId: 同步永远是1
			//emitName := messageJSON.GetStringByKey(ipc_event)
			//name := messageJSON.GetStringByKey(ipc_name)
			//browserId := messageJSON.GetIntByKey(ipc_browser_id)
			//argumentList := messageJSON.GetArrayByKey(ipc_argumentList)
		}
		return false
	})
}
