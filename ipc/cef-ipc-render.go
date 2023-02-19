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
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"net"
	"sync"
)

type renderChannel struct {
	msgID              *MsgID
	browserId          int32
	channelId          int64 //render channel Id
	ipcType            IPC_TYPE
	connect            net.Conn
	mutex              sync.Mutex
	events             *event
	emitCallback       *EmitCallbackCollection
	emitSync           map[string]*EmitSyncCollection
	renderOnEvents     []func(browseProcess IEventOn)
	renderEmitCallback []func(renderProcess IEventEmit)
	isConnect          bool
}

// 触发事件回调函数集合
type EmitCallbackCollection struct {
	EmitCollection sync.Map
}

// 触发同步事件集合
type EmitSyncCollection struct {
	Mutex          *sync.Mutex
	EmitCollection sync.Map
}

func (m *ipcChannel) newRenderChannel(memoryAddresses ...string) {
	if UseNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", IPC.Port())
		conn, err := net.Dial("tcp", address)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		m.render.ipcType = IPCT_NET
		m.render.connect = conn
	} else {
		memoryAddr := ipcSock
		logger.Debug("new render channel for IPC Sock", ipcSock)
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
	go m.render.receive()
}

func (m *EmitCallbackCollection) remove(key int32) {
	m.EmitCollection.Delete(key)
}

func (m *renderChannel) Events() *event {
	return m.events
}

func (m *renderChannel) Channel(channelId int64) *channel {
	return nil
}

func (m *renderChannel) Close() {
	if m.connect != nil {
		m.connect.Close()
		m.connect = nil
	}
	m.isConnect = false
}

// IPC render 设置监听初始化回调
func (m *renderChannel) SetOnEvent(callback func(event IEventOn)) {
	if Args.IsRender() || SingleProcess {
		m.renderOnEvents = append(m.renderOnEvents, callback)
	}
}

func (m *renderChannel) call(name string, context IIPCContext) bool {
	callBack := m.events.Get(name)
	if callBack != nil {
		callBack(context)
		return true
	}
	return false
}
func (m *renderChannel) On(name string, eventCallback EventCallback) {
	m.events.add(name, eventCallback)
}

func (m *renderChannel) Emit(eventName string, arguments IArgumentList) {
	if m.conn() == nil {
		return
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if arguments == nil {
		arguments = NewArgumentList()
	}
	ipcWrite(Tm_Async, m.channelId, m.msgID.New(), []byte(eventName), arguments.Package(), m.conn())
}

func (m *renderChannel) EmitAndCallback(eventName string, arguments IArgumentList, callback IPCCallback) {
	if m.conn() == nil {
		return
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if arguments == nil {
		arguments = NewArgumentList()
	}
	eventId := m.msgID.New()
	m.emitCallback.EmitCollection.Store(eventId, callback)
	ipcWrite(Tm_Callback, m.channelId, eventId, []byte(eventName), arguments.Package(), m.conn())
}

func (m *renderChannel) EmitAndReturn(eventName string, arguments IArgumentList) IIPCContext {
	if m.conn() == nil {
		return nil
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var emit = func(emitAsync *EmitSyncCollection) IIPCContext {
		emitAsync.Mutex.Lock()
		defer emitAsync.Mutex.Unlock()
		eventId := m.msgID.New()
		var chn = make(chan IIPCContext)
		emitAsync.EmitCollection.Store(eventId, chn)
		ipcWrite(Tm_Sync, m.channelId, eventId, []byte(eventName), arguments.Package(), m.conn())
		return <-chn
	}
	if arguments == nil {
		arguments = NewArgumentList()
	}
	if emitAsync, ok := m.emitSync[eventName]; ok {
		return emit(emitAsync)
	} else {
		m.emitSync[eventName] = &EmitSyncCollection{Mutex: new(sync.Mutex), EmitCollection: sync.Map{}}
		return emit(m.emitSync[eventName])
	}
}

func (m *renderChannel) conn() net.Conn {
	return m.connect
}

func (m *renderChannel) emitConnect() {
	args := NewArgumentList()
	args.SetString(0, "-connecting")
	m.Emit(Ln_onConnectEvent, args)
	m.isConnect = true
}

func (m *renderChannel) receive() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("IPC Render Channel Recover:", err)
		}
		m.Close()
	}()
	var readHandler = &ipcReadHandler{
		browserId: m.browserId,
		channelId: m.channelId,
		ct:        Ct_Client,
		ipcType:   m.ipcType,
		connect:   m.connect,
		handler: func(ctx *IPCContext) {
			if m.call(ctx.eventName, ctx) {
				if (ctx.triggerMode == Tm_Callback || ctx.triggerMode == Tm_Sync) && !ctx.isReply {
					ctx.Response([]byte{})
				}
			} else {
				if ctx.triggerMode == Tm_Callback { //回调函数
					m.mutex.Lock()
					defer m.mutex.Unlock()
					if callback, ok := m.emitCallback.EmitCollection.Load(ctx.eventId); ok {
						callback.(IPCCallback)(ctx)
						m.emitCallback.EmitCollection.Delete(ctx.eventId)
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
}
