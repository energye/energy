//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// ipc 通道 browser 进程(或服务端)
package channel

import (
	"fmt"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"net"
	"sync"
)

// browserChannel browser进程
type browserChannel struct {
	ipcType      IPC_TYPE
	unixAddr     *net.UnixAddr
	unixListener *net.UnixListener
	netListener  net.Listener
	mutex        sync.Mutex
	channel      sync.Map
	handler      IPCCallback
}

// NewBrowser 创建主进程通道
func NewBrowser(memoryAddresses ...string) IBrowserChannel {
	useNetIPCChannel = isUseNetIPC()
	if useNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", Port())
		listener, err := net.Listen("tcp", address)
		if err != nil {
			panic("Description Failed to create the IPC service Error: " + err.Error())
		}
		browser.ipcType = IPCT_NET
		browser.netListener = listener
	} else {
		removeMemory()
		memoryAddr := ipcSock
		logger.Debug("new browser channel for IPC Sock", memoryAddr)
		if len(memoryAddresses) > 0 {
			memoryAddr = memoryAddresses[0]
		}
		unixAddr, err := net.ResolveUnixAddr(MemoryNetwork, memoryAddr)
		if err != nil {
			panic("Description Failed to create the IPC service Error: " + err.Error())
		}
		unixListener, err := net.ListenUnix(MemoryNetwork, unixAddr)
		if err != nil {
			panic("Description Failed to create the IPC service Error: " + err.Error())
		}
		unixListener.SetUnlinkOnClose(true)
		browser.ipcType = IPCT_UNIX
		browser.unixAddr = unixAddr
		browser.unixListener = unixListener
	}
	go browser.accept()
	return browser
}

// Channel 返回指定通道链接
func (m *browserChannel) Channel(channelId int64) IChannel {
	if value, ok := m.channel.Load(channelId); ok {
		return value.(*channel)
	}
	return nil
}

// ChannelIds 返回所有已链接通道ID
func (m *browserChannel) ChannelIds() (result []int64) {
	m.channel.Range(func(key, value any) bool {
		result = append(result, key.(int64))
		return true
	})
	return
}

// Close 关闭通道链接
func (m *browserChannel) Close() {
	if m.unixListener != nil {
		m.unixListener.Close()
	}
	if m.netListener != nil {
		m.netListener.Close()
	}
}

// onChannelConnect 建立通道链接
func (m *browserChannel) onChannelConnect(conn *channel) {
	logger.Info("IPC browser on channel key_channelId:", conn.channelId)
	m.channel.Store(conn.channelId, conn)
}

// removeChannel 删除指定通道
func (m *browserChannel) removeChannel(channelId int64) {
	logger.Debug("IPC browser channel remove channelId:", channelId)
	m.channel.Delete(channelId)
}

// singleProcessChannelId 单进程进程通道获取
func (m *browserChannel) singleProcessChannelId() (int64, bool) {
	if SingleProcess {
		var channelId int64 = 0
		//单进程，只有一个IPC连接，直接取出来就好
		m.channel.Range(func(key, value any) bool {
			channelId = key.(int64)
			return false
		})
		if channelId != 0 {
			return channelId, true
		}
	}
	return 0, false
}

// Send 指定通道发送数据
func (m *browserChannel) Send(channelId int64, data []byte) {
	m.sendMessage(mt_common, channelId, channelId, data)
}

// Send 指定通道发送消息
func (m *browserChannel) sendMessage(messageType mt, channelId, toChannelId int64, data []byte) {
	if id, ok := m.singleProcessChannelId(); ok {
		channelId = id
	}
	if chn := m.Channel(toChannelId); chn != nil {
		_, _ = chn.write(messageType, channelId, toChannelId, data)
	}
}

// Handler 设置自定义处理回调函数
func (m *browserChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

// accept 接收新链接
func (m *browserChannel) accept() {
	logger.Info("IPC Server Accept")
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

// newConnection 新链接
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
	newChannel = &channel{
		channelType: Ct_Server,
		ipcType:     m.ipcType,
		conn:        conn,
	}
	newChannel.handler = func(context IIPCContext) {
		if context.Message().Type() == mt_connection {
			newChannel.channelId = context.ChannelId()
			m.onChannelConnect(newChannel)
			m.sendMessage(mt_connectd, context.ChannelId(), context.ToChannelId(), []byte{uint8(mt_connectd)})
			newChannel.isConnect = true
		} else if context.Message().Type() == mt_relay {
			m.sendMessage(mt_common, context.ChannelId(), context.ToChannelId(), context.Message().Data())
		} else {
			if m.handler != nil {
				go m.handler(context)
			}
		}
	}
	newChannel.ipcRead()
}
