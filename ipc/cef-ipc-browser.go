//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package ipc

import (
	"fmt"
	. "github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"net"
	"sync"
)

type browserChannel struct {
	msgID              *MsgID
	cliID              *CliID
	ipcType            IPC_TYPE
	unixAddr           *net.UnixAddr
	unixListener       *net.UnixListener
	netListener        net.Listener
	mutex              sync.Mutex
	events             *event
	channel            map[int64]*channel
	emitCallback       *EmitCallbackCollection
	emitSync           map[string]*EmitSyncCollection
	browseOnEvents     []func(browseProcess IEventOn)
	browseEmitCallback []func(renderProcess IEventEmit)
}

type channel struct {
	IPCType  IPC_TYPE
	UnixConn *net.UnixConn
	NetConn  net.Conn
}

func (m *ipcChannel) newBrowseChannel(memoryAddresses ...string) {
	if UseNetIPCChannel {
		IPC.SetPort()
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
	m.browser.onConnect()
}

func (m *channel) conn() net.Conn {
	if m.IPCType == IPCT_NET {
		return m.NetConn
	}
	return m.UnixConn
}

func (m *event) check(name string) bool {
	if _, ok := m.event[name]; ok {
		return true
	}
	return false
}

func (m *event) add(name string, eventCallback EventCallback) {
	if !m.check(name) {
		m.event[name] = eventCallback
	}
}

func (m *event) removeOnEvent(name string) {
	delete(m.event, name)
}

func (m *event) Get(name string) EventCallback {
	if call, ok := m.event[name]; ok {
		return call
	}
	return nil
}

func (m *browserChannel) Channel(channelId int64) *channel {
	return m.channel[channelId]
}
func (m *browserChannel) Close() {
	if m.unixListener != nil {
		m.unixListener.Close()
	}
}

func (m *browserChannel) onConnect() {
	m.On(Ln_onConnectEvent, func(context IIPCContext) {
		logger.Info("IPC browser on connect channelId:", context.ChannelId())
		if context.ChannelId() > 0 {
			if chn, ok := m.channel[context.ChannelId()]; ok {
				chn.IPCType = m.ipcType
				chn.UnixConn = context.uConn()
				chn.NetConn = context.conn()
			} else {
				m.channel[context.ChannelId()] = &channel{
					IPCType:  m.ipcType,
					UnixConn: context.uConn(),
					NetConn:  context.conn(),
				}
			}
			//context.setChannelId(context.ChannelId())
			//context.Response(Int16ToBytes(context.ChannelId()))
		}
	})
}

func (m *browserChannel) removeChannel(id int64) {
	logger.Debug("IPC browser channel remove channelId:", id)
	delete(m.channel, id)
}

func (m *browserChannel) Events() *event {
	return m.events
}

func (m *browserChannel) call(name string, context IIPCContext) bool {
	callBack := m.events.Get(name)
	if callBack != nil {
		callBack(context)
		return true
	}
	return false
}

// IPC browser 设置监听初始化回调
func (m *browserChannel) SetOnEvent(callback func(event IEventOn)) {
	if Args.IsMain() {
		m.browseOnEvents = append(m.browseOnEvents, callback)
	}
}

func (m *browserChannel) On(name string, eventCallback EventCallback) {
	m.events.add(name, eventCallback)
}

func (m *browserChannel) RemoveOn(name string) {
	m.events.removeOnEvent(name)
}

//单进程进程通道获取
func (m *browserChannel) singleProcessChannelId() (int64, bool) {
	if SingleProcess {
		//单进程，只有一个IPC连接，直接取出来就好
		for channelId, _ := range IPC.browser.channel {
			return channelId, true
		}
	}
	return 0, false
}

func (m *browserChannel) EmitChannelId(eventName string, chnId int64, arguments IArgumentList) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if channelId, ok := m.singleProcessChannelId(); ok {
		chnId = channelId
	}
	if chn, ok := m.channel[chnId]; ok {
		if arguments == nil {
			arguments = NewArgumentList()
		}
		_, _ = ipcWrite(Tm_Async, chnId, m.msgID.New(), []byte(eventName), arguments.Package(), chn.conn())
		arguments.Clear()
		arguments = nil
	}

}

func (m *browserChannel) Emit(eventName string, arguments IArgumentList) {
	m.EmitChannelId(eventName, int64(PID_RENDERER), arguments)
}

func (m *browserChannel) EmitChannelIdAndCallback(eventName string, chnId int64, arguments IArgumentList, callback IPCCallback) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if channelId, ok := m.singleProcessChannelId(); ok {
		chnId = channelId
	}
	if chn, ok := m.channel[chnId]; ok {
		if arguments == nil {
			arguments = NewArgumentList()
		}
		eventId := m.msgID.New()
		m.emitCallback.EmitCollection.Store(eventId, callback)
		_, _ = ipcWrite(Tm_Callback, chnId, eventId, []byte(eventName), arguments.Package(), chn.conn())
	}

}
func (m *browserChannel) EmitAndCallback(eventName string, arguments IArgumentList, callback IPCCallback) {
	m.EmitChannelIdAndCallback(eventName, int64(PID_RENDERER), arguments, callback)
}

func (m *browserChannel) EmitChannelIdAndReturn(eventName string, chnId int64, arguments IArgumentList) IIPCContext {
	if channelId, ok := m.singleProcessChannelId(); ok {
		chnId = channelId
	}
	if chn, ok := m.channel[chnId]; ok {
		var emit = func(emitAsync *EmitSyncCollection, arguments IArgumentList, conn net.Conn) IIPCContext {
			emitAsync.Mutex.Lock()
			defer emitAsync.Mutex.Unlock()
			if arguments == nil {
				arguments = NewArgumentList()
			}
			eventId := m.msgID.New()
			var chn = make(chan IIPCContext)
			emitAsync.EmitCollection.Store(eventId, chn)
			_, _ = ipcWrite(Tm_Sync, chnId, eventId, []byte(eventName), arguments.Package(), conn)
			return <-chn
		}
		if emitAsync, ok := m.emitSync[eventName]; ok {
			return emit(emitAsync, arguments, chn.conn())
		} else {
			m.emitSync[eventName] = &EmitSyncCollection{Mutex: new(sync.Mutex), EmitCollection: sync.Map{}}
			return emit(m.emitSync[eventName], arguments, chn.conn())
		}
	}
	return nil
}

func (m *browserChannel) EmitAndReturn(eventName string, arguments IArgumentList) IIPCContext {
	return m.EmitChannelIdAndReturn(eventName, int64(PID_RENDERER), arguments)
}

func (m *browserChannel) accept() {
	logger.Info("IPC Server Accept")
	for {
		var (
			err      error
			unixConn *net.UnixConn
			netConn  net.Conn
		)
		if m.ipcType == IPCT_UNIX {
			unixConn, err = m.unixListener.AcceptUnix()
		} else {
			netConn, err = m.netListener.Accept()
		}
		if err != nil {
			logger.Info("browser channel accept Error:", err.Error())
			continue
		}
		go func() {
			var id int64 //render channel channelId
			defer func() {
				m.removeChannel(id)
				if unixConn != nil {
					unixConn.Close()
				} else if netConn != nil {
					netConn.Close()
				}
			}()
			var readHandler = &ipcReadHandler{
				ct:       Ct_Server,
				ipcType:  m.ipcType,
				unixConn: unixConn,
				unixAddr: m.unixAddr,
				netConn:  netConn,
				handler: func(ctx *IPCContext) {
					if m.call(ctx.eventName, ctx) {
						if id == 0 && ctx.channelId > 0 {
							id = ctx.channelId
						}
						if (ctx.triggerMode == Tm_Callback || ctx.triggerMode == Tm_Sync) && !ctx.isReply {
							ctx.Response([]byte{})
						}
					} else {
						if ctx.triggerMode == Tm_Callback { //回调函数
							if callback, ok := m.emitCallback.EmitCollection.Load(ctx.eventId); ok {
								m.emitCallback.EmitCollection.Delete(ctx.eventId)
								callback.(IPCCallback)(ctx)
							}
						} else if ctx.triggerMode == Tm_Sync { //同步调用
							if emitAsync, ok := m.emitSync[ctx.eventName]; ok {
								if chn, ok := emitAsync.EmitCollection.Load(ctx.eventId); ok {
									var c = chn.(chan IIPCContext)
									c <- ctx
									close(c)
									emitAsync.EmitCollection.Delete(ctx.eventId)
								}
							}
						}
					}
				},
			}
			ipcRead(readHandler)
		}()
	}
}

func (m *browserChannel) free() {
	m.unixListener.Close()
}
