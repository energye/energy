//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ipc

import (
	"fmt"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/energy/pkgs/json"
	"net"
	"sync"
)

type renderChannel struct {
	channelId int64
	ipcType   IPC_TYPE
	connect   net.Conn
	mutex     sync.Mutex
	isConnect bool
	handler   IPCCallback
}

func (m *ipcChannel) NewRenderChannel(channelId int64, memoryAddresses ...string) *renderChannel {
	if useNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", IPC.Port())
		conn, err := net.Dial("tcp", address)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		m.render.ipcType = IPCT_NET
		m.render.connect = conn
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
		m.render.ipcType = IPCT_UNIX
		m.render.connect = unixConn
	}
	m.render.channelId = channelId
	go m.render.receive()
	m.render.onConnection()
	return m.render
}

func (m *renderChannel) onConnection() {
	message := json.NewJSONObject(nil)
	message.Set(channelId, m.channelId)
	m.sendMessage(mt_connection, message.Bytes())
	message.Free()
}

func (m *renderChannel) Send(data []byte) {
	m.sendMessage(mt_common, data)
}

func (m *renderChannel) sendMessage(messageType mt, data []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, _ = ipcWrite(messageType, data, m.conn())
}

func (m *renderChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

func (m *renderChannel) Close() {
	if m.connect != nil {
		m.connect.Close()
		m.connect = nil
	}
}

func (m *renderChannel) conn() net.Conn {
	return m.connect
}

func (m *renderChannel) receive() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("IPC Render Channel Recover:", err)
		}
		m.Close()
	}()
	var readHandler = &ipcReadHandler{
		ct:      Ct_Client,
		ipcType: m.ipcType,
		connect: m.connect,
		handler: func(context IIPCContext) {
			if m.handler != nil {
				m.handler(context)
			}
		},
	}
	ipcRead(readHandler)
}
