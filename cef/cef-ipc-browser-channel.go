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

import "github.com/energye/energy/pkgs/channel"

// browserIPCChan
type browserIPCChan struct {
	ipc channel.IBrowserChannel
}

func (m *ipcBrowserProcess) ipcChannelBrowser() {
	if m.ipcChannel == nil {
		m.ipcChannel = new(browserIPCChan)
		m.ipcChannel.ipc = channel.NewBrowser()
		m.ipcChannel.ipc.Handler(func(context channel.IIPCContext) {
			context.Free()
		})
	}
}
