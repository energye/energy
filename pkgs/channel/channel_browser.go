//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// ipc Channel browser process (or server)

package channel

import (
	"fmt"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"net"
	"sync"
)

// browserChannel main(browser) process
type browserChannel struct {
	ipcType      IPC_TYPE
	unixAddr     *net.UnixAddr
	unixListener *net.UnixListener
	netListener  net.Listener
	channel      sync.Map
	handler      IPCCallback
}

// NewBrowser Create main(browser) process channel
func NewBrowser(addresses ...string) IBrowserChannel {
	useNetIPCChannel = IsUseNetIPC()
	browser := &browserChannel{
		channel: sync.Map{},
	}
	if useNetIPCChannel {
		// 监听并绑定端口
		address := fmt.Sprintf("localhost:%d", Port())
		listener, err := net.Listen("tcp", address)
		if err != nil {
			panic("NewBrowser IPC channel Error: " + err.Error())
		}
		browser.ipcType = IPCT_NET
		browser.netListener = listener
	} else {
		if len(addresses) > 0 {
			ipcSock = addresses[0]
		}
		removeMemory()
		logger.Debug("new browser channel for IPC Sock", ipcSock)
		unixAddr, err := net.ResolveUnixAddr(MemoryNetwork, ipcSock)
		if err != nil {
			panic("NewBrowser IPC channel Error: " + err.Error())
		}
		unixListener, err := net.ListenUnix(MemoryNetwork, unixAddr)
		if err != nil {
			panic("NewBrowser IPC channel Error: " + err.Error())
		}
		unixListener.SetUnlinkOnClose(true)
		browser.ipcType = IPCT_UNIX
		browser.unixAddr = unixAddr
		browser.unixListener = unixListener
	}
	go browser.accept()
	return browser
}

// Channel Return to the specified channel connection
func (m *browserChannel) Channel(channelId string) IChannel {
	if value, ok := m.channel.Load(channelId); ok {
		return value.(*channel)
	}
	return nil
}

// ChannelIds Return all connected channel IDs
func (m *browserChannel) ChannelIds() (result []int64) {
	m.channel.Range(func(key, value interface{}) bool {
		result = append(result, key.(int64))
		return true
	})
	return
}

// Close Close channel connection
func (m *browserChannel) Close() {
	if m.unixListener != nil {
		m.unixListener.Close()
	}
	if m.netListener != nil {
		m.netListener.Close()
	}
}

// onChannelConnect Establishing channel connection
func (m *browserChannel) onChannelConnect(conn *channel) {
	logger.Info("IPC browser on channel channelId:", conn.channelId)
	m.channel.Store(conn.channelId, conn)
}

// removeChannel Delete specified channel
//
//	When the channel is closed
func (m *browserChannel) removeChannel(channelId string) {
	logger.Debug("IPC browser channel remove channelId:", channelId)
	m.channel.Delete(channelId)
}

// Send Specify channel to send data
func (m *browserChannel) Send(channelId string, data []byte) {
	m.sendMessage(mt_common, channelId, channelId, data)
}

// Send Specify the channel to send messages
func (m *browserChannel) sendMessage(messageType mt, channelId, toChannelId string, data []byte) {
	if chn := m.Channel(toChannelId); chn != nil {
		_, _ = chn.write(messageType, channelId, toChannelId, data)
	}
}

// Handler
//
//	Set custom processing callback function
func (m *browserChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

// accept
//
//	Receive new connection
func (m *browserChannel) accept() {
	for {
		var (
			err  error
			conn net.Conn
		)
		if m.ipcType == IPCT_UNIX {
			conn, err = m.unixListener.AcceptUnix()
		} else {
			conn, err = m.netListener.Accept()
		}
		if err != nil {
			logger.Info("browser channel accept Error:", err.Error())
			continue
		}
		go m.newConnection(conn)
	}
}

// newConnection
//
//	new connection
func (m *browserChannel) newConnection(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("IPC Server Accept Recover:", err)
		}
	}()
	var newChannel *channel
	defer func() {
		if newChannel != nil {
			m.removeChannel(newChannel.channelId)
			newChannel.conn = nil
			newChannel.isConnect = false
		}
	}()
	// create channel
	newChannel = &channel{
		channelType: Ct_Server,
		ipcType:     m.ipcType,
		conn:        conn,
	}
	// handler
	newChannel.handler = func(context IIPCContext) {
		if context.Message().Type() == mt_connection { // new connection
			newChannel.channelId = context.ChannelId()
			m.onChannelConnect(newChannel)
			m.sendMessage(mt_connectd, context.ChannelId(), context.ToChannelId(), []byte{uint8(mt_connectd)})
			newChannel.isConnect = true
		} else if context.Message().Type() == mt_update_channel_id { //update channel id
			var (
				oldChannelId = context.ChannelId()   // old channel id
				newChannelId = context.ToChannelId() // new channel id
			)
			if oldChannelId != newChannelId {
				newChannel.channelId = newChannelId // set new channel id
				m.onChannelConnect(newChannel)      // add new channel id
				m.removeChannel(oldChannelId)       // delete old channel id
			}
		} else if context.Message().Type() == mt_relay { // relay
			m.sendMessage(mt_common, context.ChannelId(), context.ToChannelId(), context.Message().Data())
		} else {
			// default handler
			if m.handler != nil {
				m.handler(context)
			}
		}
	}
	newChannel.ipcRead()
}
