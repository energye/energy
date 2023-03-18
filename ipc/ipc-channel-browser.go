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
	"os"
	"sync"
)

type browserChannel struct {
	ipcType      IPC_TYPE
	unixAddr     *net.UnixAddr
	unixListener *net.UnixListener
	netListener  net.Listener
	mutex        sync.Mutex
	channel      sync.Map
	handler      IPCCallback
}

type channel struct {
	IPCType IPC_TYPE
	Conn    net.Conn
}

func removeMemory() {
	os.Remove(ipcSock)
}

func UseNetIPCChannel() bool {
	return useNetIPCChannel
}

func MemoryAddress() string {
	return memoryAddress
}

func (m *ipcChannel) NewBrowseChannel(memoryAddresses ...string) *browserChannel {
	if useNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", IPC.Port())
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

func (m *channel) conn() net.Conn {
	return m.Conn
}

func (m *browserChannel) Channel(channelId int64) *channel {
	if value, ok := m.channel.Load(channelId); ok {
		return value.(*channel)
	}
	return nil
}

func (m *browserChannel) ChannelIds() (result []int64) {
	m.channel.Range(func(key, value any) bool {
		result = append(result, key.(int64))
		return true
	})
	return
}

func (m *browserChannel) putChannel(channelId int64, value *channel) {
	m.channel.Store(channelId, value)
}

func (m *browserChannel) Close() {
	if m.unixListener != nil {
		m.unixListener.Close()
	}
}

func (m *browserChannel) onConnect(context IIPCContext) {
	logger.Info("IPC browser on connect key_channelId:", context.ChannelId())
	if chn := m.Channel(context.ChannelId()); chn != nil {
		chn.IPCType = m.ipcType
		chn.Conn = context.Connect()
	} else {
		m.putChannel(context.ChannelId(), &channel{
			IPCType: m.ipcType,
			Conn:    context.Connect(),
		})
	}
}

func (m *browserChannel) removeChannel(id int64) {
	logger.Debug("IPC browser channel remove key_channelId:", id)
	m.channel.Delete(id)
}

// 单进程进程通道获取
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

func (m *browserChannel) Send(channelId int64, data []byte) {
	m.sendMessage(mt_common, channelId, data)
}

func (m *browserChannel) sendMessage(messageType mt, channelId int64, data []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if id, ok := m.singleProcessChannelId(); ok {
		channelId = id
	}
	if chn := m.Channel(channelId); chn != nil {
		_, _ = ipcWrite(messageType, channelId, data, chn.conn())
	}
}

func (m *browserChannel) Handler(handler IPCCallback) {
	m.handler = handler
}

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
		go m.ipcReadHandler(conn)
	}
}

func (m *browserChannel) ipcReadHandler(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("IPC Server Accept Recover:", err)
		}
	}()
	var id int64 //render channel key_channelId
	defer func() {
		m.removeChannel(id)
	}()
	var readHandler = &ipcReadHandler{
		ct:      Ct_Server,
		ipcType: m.ipcType,
		connect: conn,
		handler: func(context IIPCContext) {
			if context.Message().Type() == mt_connection {
				message := json.NewJSONObject(context.Message().Data())
				id = int64(message.GetIntByKey(key_channelId))
				m.onConnect(context)
			} else {
				if m.handler != nil {
					m.handler(context)
				}
			}
		},
	}
	ipcRead(readHandler)
}
