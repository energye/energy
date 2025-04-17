//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// ipc Channel render process (or client)

package channel

import (
	"fmt"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"net"
)

// renderChannel renderer process
type renderChannel struct {
	channel *channel
	handler IPCCallback
}

// NewRender Create the renderer process channel
//
// param: channelId Unique channel ID identifier
func NewRender(channelId string, addresses ...string) IRenderChannel {
	useNetIPCChannel = IsUseNetIPC()
	render := &renderChannel{}
	if useNetIPCChannel {
		// 监听端口
		address := fmt.Sprintf("localhost:%d", Port())
		conn, err := net.Dial("tcp", address)
		if err != nil {
			panic("NewRender IPC channel Error: " + err.Error())
		}
		render.channel = &channel{conn: conn, channelId: channelId, ipcType: IPCT_NET, channelType: Ct_Client}
	} else {
		if len(addresses) > 0 {
			ipcSock = addresses[0]
		}
		logger.Debug("new render channel for IPC Sock", ipcSock)
		unixAddr, err := net.ResolveUnixAddr(MemoryNetwork, ipcSock)
		if err != nil {
			panic("NewRender IPC channel Error: " + err.Error())
		}
		conn, err := net.DialUnix(MemoryNetwork, nil, unixAddr)
		if err != nil {
			panic("NewRender IPC channel Error: " + err.Error())
		}
		render.channel = &channel{conn: conn, channelId: channelId, ipcType: IPCT_UNIX, channelType: Ct_Client}
	}
	go render.receive()
	render.onChannelConnect()
	return render
}

// Channel Return to current channel
func (m *renderChannel) Channel() IChannel {
	return m.channel
}

// onChannelConnect Establishing channel connection
func (m *renderChannel) onChannelConnect() {
	m.sendMessage(mt_connection, m.channel.channelId, m.channel.channelId, []byte{uint8(mt_connection)})
}

// Send data
func (m *renderChannel) Send(data []byte) {
	if m.channel != nil && m.channel.IsConnect() {
		m.sendMessage(mt_common, m.channel.channelId, m.channel.channelId, data)
	}
}

// SendToChannel Send to specified channel
func (m *renderChannel) SendToChannel(toChannelId string, data []byte) {
	if m.channel != nil && m.channel.IsConnect() {
		m.sendMessage(mt_relay, m.channel.channelId, toChannelId, data)
	}
}

// UpdateChannelId
//
//	Update channel ID
//	The original channel ID is invalid after updating
func (m *renderChannel) UpdateChannelId(newChannelId string) {
	if m.channel.channelId != newChannelId {
		m.sendMessage(mt_update_channel_id, m.channel.channelId, newChannelId, []byte{uint8(mt_update_channel_id)})
		m.channel.channelId = newChannelId
	}
}

// sendMessage
//
//	Send data to the specified channel
func (m *renderChannel) sendMessage(messageType mt, channelId, toChannelId string, data []byte) {
	_, _ = m.channel.write(messageType, channelId, toChannelId, data)
}

// Handler
//
//	Set custom processing callback function
func (m *renderChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

// Close channel
func (m *renderChannel) Close() {
	if m.channel != nil {
		m.channel.Close()
		m.channel = nil
	}
}

// receive Data
func (m *renderChannel) receive() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("IPC Render Channel Recover:", err)
		}
		m.channel.isConnect = false
		m.Close()
	}()
	// handler
	m.channel.handler = func(context IIPCContext) {
		if context.Message().Type() == mt_connectd {
			m.channel.isConnect = true
		} else {
			// default handler
			if m.handler != nil {
				m.handler(context)
			}
		}
	}
	m.channel.ipcRead()
}
