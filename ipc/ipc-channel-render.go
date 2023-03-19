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
package ipc

import (
	"bytes"
	"fmt"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/energy/pkgs/json"
	"net"
	"sync"
)

// renderChannel 渲染进程
type renderChannel struct {
	connect   *connect
	mutex     sync.Mutex
	isConnect bool
	handler   IPCCallback
}

// NewRender 创建渲染进程通道
//
// channelId 唯一通道ID标识
func (m *ipcChannel) NewRender(channelId int64, memoryAddresses ...string) *renderChannel {
	useNetIPCChannel = isUseNetIPC()
	if useNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", Channel.Port())
		conn, err := net.Dial("tcp", address)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		m.render.connect = &connect{writeBuf: new(bytes.Buffer), conn: conn, channelId: channelId, ipcType: IPCT_NET, ct: Ct_Client}
	} else {
		memoryAddr := ipcSock
		logger.Debug("new render channel for IPC Sock", memoryAddr)
		if len(memoryAddresses) > 0 {
			memoryAddr = memoryAddresses[0]
		}
		unixAddr, err := net.ResolveUnixAddr(MemoryNetwork, memoryAddr)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		unixConn, err := net.DialUnix(MemoryNetwork, nil, unixAddr)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		m.render.connect = &connect{writeBuf: new(bytes.Buffer), conn: unixConn, channelId: channelId, ipcType: IPCT_UNIX, ct: Ct_Client}
	}
	go m.render.receive()
	m.render.onConnection()
	return m.render
}

// onConnection 建立链接
func (m *renderChannel) onConnection() {
	message := json.NewJSONObject(nil)
	message.Set(key_channelId, m.connect.channelId)
	m.sendMessage(mt_connection, message.Bytes())
	message.Free()
}

// Send 发送数据
func (m *renderChannel) Send(data []byte) {
	m.sendMessage(mt_common, data)
}

// sendMessage 发送消息
func (m *renderChannel) sendMessage(messageType mt, data []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, _ = m.connect.ipcWrite(messageType, m.connect.channelId, data)
}

// Handler 设置自定义处理回调函数
func (m *renderChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

// Close 关闭通道链接
func (m *renderChannel) Close() {
	if m.connect != nil {
		m.connect.Close()
		m.connect = nil
	}
}

// receive 接收数据
func (m *renderChannel) receive() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("IPC Render Channel Recover:", err)
		}
		m.Close()
	}()
	m.connect.handler = func(context IIPCContext) {
		if m.handler != nil {
			m.handler(context)
		}
	}
	m.connect.ipcRead()
}
