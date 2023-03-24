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

// 渲染(子)进程IPC
var renderIPC *renderIPCChan

// renderIPCChan
type renderIPCChan struct {
	browserId int32                                        //浏览器ID
	channelId int64                                        //通道ID 使用FrameId
	ipc       channel.IRenderChannel                       //channel
	callback  []func(channelId int64, data json.JSON) bool //实现回调函数
}

// renderIPCCreate 渲染进程IPC创建
func renderIPCCreate(browser *ICefBrowser, frame *ICefFrame) {
	if renderIPC == nil {
		renderIPC = new(renderIPCChan)
		renderIPC.browserId = browser.Identifier()
		renderIPC.channelId = frame.Identifier()
		renderIPC.ipc = channel.NewRender(renderIPC.channelId)
		renderIPC.ipc.Handler(func(context channel.IIPCContext) {
			data := context.Message().JSON()
			//callback 返回 true ipc 停止遍历,否则继续遍历,直到最后一个
			for _, call := range renderIPC.callback {
				if call(context.ChannelId(), data) {
					break
				}
			}
			context.Free()
		})
	}
}

// addCallback 添加回调函数, callback 返回 true ipc 停止遍历,否则继续遍历,直到最后一个
func (m *renderIPCChan) addCallback(callback func(channelId int64, data json.JSON) bool) {
	m.callback = append(m.callback, callback)
}
