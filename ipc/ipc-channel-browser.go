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
func (m *ipcChannel) NewBrowser(memoryAddresses ...string) *browserChannel {
	useNetIPCChannel = isUseNetIPC()
	if useNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", Channel.Port())
		listener, err := net.Listen("tcp", address)
		if err != nil {
			panic("Description Failed to create the IPC service Error: " + err.Error())
		}
		m.browser.ipcType = IPCT_NET
		m.browser.netListener = listener
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
		m.browser.ipcType = IPCT_UNIX
		m.browser.unixAddr = unixAddr
		m.browser.unixListener = unixListener
	}
	go m.browser.accept()
	return m.browser
}

// Channel 返回指定通道链接
func (m *browserChannel) Channel(channelId int64) *connect {
	if value, ok := m.channel.Load(channelId); ok {
		return value.(*connect)
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

// onConnect 建立链接
func (m *browserChannel) onConnect(conn *connect) {
	logger.Info("IPC browser on connect key_channelId:", conn.channelId)
	m.channel.Store(conn.channelId, conn)
}

// removeChannel 删除指定通道
func (m *browserChannel) removeChannel(id int64) {
	logger.Debug("IPC browser channel remove key_channelId:", id)
	m.channel.Delete(id)
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
	m.sendMessage(mt_common, channelId, data)
}

// Send 指定通道发送消息
func (m *browserChannel) sendMessage(messageType mt, channelId int64, data []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if id, ok := m.singleProcessChannelId(); ok {
		channelId = id
	}
	if chn := m.Channel(channelId); chn != nil {
		_, _ = chn.ipcWrite(messageType, channelId, data)
	}
}

// Handler 设置自定义处理回调函数
func (m *browserChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

// accept 接收链接的链接
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
	var newConn *connect
	defer func() {
		if newConn != nil {
			m.removeChannel(newConn.channelId)
		}
	}()
	newConn = &connect{
		writeBuf: new(bytes.Buffer),
		ct:       Ct_Server,
		ipcType:  m.ipcType,
		conn:     conn,
	}
	newConn.handler = func(context IIPCContext) {
		if context.Message().Type() == mt_connection {
			message := json.NewJSONObject(context.Message().Data())
			newConn.channelId = int64(message.GetIntByKey(key_channelId))
			message.Free()
			m.onConnect(newConn)
		} else {
			if m.handler != nil {
				m.handler(context)
			}
		}
	}
	newConn.ipcRead()
}
