//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// ipc 通道 render 进程(或客户端)
package channel

import (
	"fmt"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"net"
	"sync"
)

// renderChannel 渲染进程
type renderChannel struct {
	channel *channel
	mutex   sync.Mutex
	handler IPCCallback
}

// NewRender 创建渲染进程通道
//
// 参数: channelId 唯一通道ID标识
func NewRender(channelId int64, memoryAddresses ...string) IRenderChannel {
	useNetIPCChannel = isUseNetIPC()
	render := &renderChannel{
		mutex: sync.Mutex{},
	}
	if useNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", Port())
		conn, err := net.Dial("tcp", address)
		if err != nil {
			panic("Client failed to channel to IPC service Error: " + err.Error())
		}
		render.channel = &channel{conn: conn, channelId: channelId, ipcType: IPCT_NET, channelType: Ct_Client}
	} else {
		if len(memoryAddresses) > 0 {
			ipcSock = memoryAddresses[0]
		}
		logger.Debug("new render channel for IPC Sock", ipcSock)
		unixAddr, err := net.ResolveUnixAddr(MemoryNetwork, ipcSock)
		if err != nil {
			panic("Client failed to channel to IPC service Error: " + err.Error())
		}
		unixConn, err := net.DialUnix(MemoryNetwork, nil, unixAddr)
		if err != nil {
			panic("Client failed to channel to IPC service Error: " + err.Error())
		}
		render.channel = &channel{conn: unixConn, channelId: channelId, ipcType: IPCT_UNIX, channelType: Ct_Client}
	}
	go render.receive()
	render.onChannelConnect()
	return render
}

// Channel 返回当前通道
func (m *renderChannel) Channel() IChannel {
	return m.channel
}

// onChannelConnect 建立通道链接
func (m *renderChannel) onChannelConnect() {
	m.sendMessage(mt_connection, m.channel.channelId, m.channel.channelId, []byte{uint8(mt_connection)})
}

// Send 发送数据
func (m *renderChannel) Send(data []byte) {
	if m.channel != nil && m.channel.IsConnect() {
		m.sendMessage(mt_common, m.channel.channelId, m.channel.channelId, data)
	}
}

// SendToChannel 发送到指定通道
func (m *renderChannel) SendToChannel(toChannelId int64, data []byte) {
	if m.channel != nil && m.channel.IsConnect() {
		m.sendMessage(mt_relay, m.channel.channelId, toChannelId, data)
	}
}

// sendMessage 发送消息
func (m *renderChannel) sendMessage(messageType mt, channelId, toChannelId int64, data []byte) {
	_, _ = m.channel.write(messageType, channelId, toChannelId, data)
}

// Handler 设置自定义处理回调函数
func (m *renderChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

// Close 关闭通道链接
func (m *renderChannel) Close() {
	if m.channel != nil {
		m.channel.Close()
		m.channel = nil
	}
}

// receive 接收数据
func (m *renderChannel) receive() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("IPC Render Channel Recover:", err)
		}
		fmt.Println("close")
		m.channel.isConnect = false
		m.Close()
	}()
	m.channel.handler = func(context IIPCContext) {
		if context.Message().Type() == mt_connectd {
			m.channel.isConnect = true
		} else {
			if m.handler != nil {
				go m.handler(context)
			}
		}
	}
	m.channel.ipcRead()
}
