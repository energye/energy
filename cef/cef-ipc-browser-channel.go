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

// 主进程IPC
var browserIPC *browserIPCChan

// browserIPCChan
type browserIPCChan struct {
	ipc      channel.IBrowserChannel                      //channel
	callback []func(channelId int64, data json.JSON) bool //实现回调函数
}

// browserIPCCreate 主进程IPC创建
func browserIPCCreate() {
	if browserIPC == nil {
		browserIPC = new(browserIPCChan)
		browserIPC.ipc = channel.NewBrowser()
		browserIPC.ipc.Handler(func(context channel.IIPCContext) {
			data := context.Message().JSON()
			//callback 返回 true ipc 停止遍历,否则继续遍历,直到最后一个
			for _, call := range browserIPC.callback {
				if call(context.ChannelId(), data) {
					break
				}
			}
			context.Free()
		})
	}
}

// addCallback 添加回调函数, callback 返回 true ipc 停止遍历,否则继续遍历,直到最后一个
func (m *browserIPCChan) addCallback(callback func(channelId int64, data json.JSON) bool) {
	m.callback = append(m.callback, callback)
}
