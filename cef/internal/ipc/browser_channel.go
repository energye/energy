//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// GO browser IPC channel

package ipc

import (
	"github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/pkgs/channel"
)

// Main Process IPC
var browserChan *browserIPCChan

type IBrowserIPCChan interface {
	IPC() channel.IBrowserChannel
	AddCallback(callback func(channelId int64, argument argument.IList) bool)
}

// browserIPCChan
//
//	Current browser process IPC channel processing
type browserIPCChan struct {
	ipc      channel.IBrowserChannel
	callback []func(channelId int64, argument argument.IList) bool
}

// CreateBrowserIPC
//
//	Main process IPC creation
func CreateBrowserIPC() IBrowserIPCChan {
	if browserChan == nil {
		browserChan = new(browserIPCChan)
		browserChan.ipc = channel.NewBrowser()
		browserChan.ipc.Handler(func(context channel.IIPCContext) {
			defer func() {
				context.Free()
			}()
			data := context.Message().Data()
			arguments := argument.UnList(data)
			if browserChan.listen(context, arguments) {
				return
			}
			//callback 返回 true ipc 停止遍历,否则继续遍历,直到最后一个
			for _, call := range browserChan.callback {
				if call(context.ChannelId(), arguments) {
					break
				}
			}
		})
	}
	return browserChan
}

// BrowserChan
//
// 返回Browser IPC Channel, 这个是在channel创建的
func BrowserChan() IBrowserIPCChan {
	return browserChan
}

// IPC
//
// 返回Browser IPC Channel
func (m *browserIPCChan) IPC() channel.IBrowserChannel {
	return m.ipc
}

// AddCallback
//
//	Add a callback function
//	callback returns true ipc to stop traversing
//	otherwise continue traversing until the last one
func (m *browserIPCChan) AddCallback(callback func(channelId int64, argument argument.IList) bool) {
	m.callback = append(m.callback, callback)
}
