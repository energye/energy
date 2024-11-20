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
package ipc

import (
	"github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/pkgs/channel"
)

// Rendering (sub) process IPC
var renderChan *renderIPCChan

type IRenderIPCChan interface {
	IPC() channel.IRenderChannel
	SetRealityChannel(browserId int32, channelId string)
	AddCallback(callback func(channelId string, argumentList argument.IList) bool)
	BrowserId() int32
	ChannelId() string
}

// renderIPCChan
//
//	Current renderer process IPC channel processing
type renderIPCChan struct {
	browserId int32
	channelId string
	ipc       channel.IRenderChannel
	callback  []func(channelId string, argumentList argument.IList) bool
}

// CreateRenderIPC Rendering process IPC creation
func CreateRenderIPC(browserId int32, channelId string) IRenderIPCChan {
	if renderChan == nil {
		renderChan = new(renderIPCChan)
		renderChan.browserId = browserId
		renderChan.channelId = channelId
		renderChan.ipc = channel.NewRender(channelId)
		renderChan.ipc.Handler(func(context channel.IIPCContext) {
			defer func() {
				context.Free()
			}()
			data := context.Message().Data()
			argumentList := argument.UnList(data)
			if renderChan.listen(context, argumentList) {
				return
			}
			//callback Return true to stop traversing, otherwise continue traversing until the last one
			for _, call := range renderChan.callback {
				if call(context.ChannelId(), argumentList) {
					break
				}
			}
		})
	}
	return renderChan
}

func RenderChan() IRenderIPCChan {
	return renderChan
}

func (m *renderIPCChan) IPC() channel.IRenderChannel {
	return m.ipc
}

// SetRealityChannel Set the actual channel ID
func (m *renderIPCChan) SetRealityChannel(browserId int32, channelId string) {
	if m == nil {
		return
	}
	m.browserId = browserId
	m.channelId = channelId
	m.ipc.UpdateChannelId(channelId)
}

func (m *renderIPCChan) BrowserId() int32 {
	return m.browserId
}

func (m *renderIPCChan) ChannelId() string {
	return m.channelId
}

// AddCallback
//
//	Add a callback function
//	callback returns true ipc to stop traversing
//	otherwise continue traversing until the last one
func (m *renderIPCChan) AddCallback(callback func(channelId string, argumentList argument.IList) bool) {
	m.callback = append(m.callback, callback)
}
