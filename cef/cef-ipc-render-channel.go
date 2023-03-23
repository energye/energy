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

import "github.com/energye/energy/pkgs/channel"

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
			context.Free()
		})
	}
}
