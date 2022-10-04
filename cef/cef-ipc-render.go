//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"net"
	"sync"
)

type renderChannel struct {
	msgID              *msgID
	browserId          int32
	channelId          int64 //render channel Id
	ipcType            IPC_TYPE
	unixAddr           *net.UnixAddr
	unixConn           *net.UnixConn
	netConn            net.Conn
	mutex              sync.Mutex
	events             *event
	emitCallback       *emitCallbackCollection
	emitSync           map[string]*emitSyncCollection
	renderOnEvents     []func(browseProcess IEventOn)
	renderEmitCallback []func(renderProcess IEventEmit)
	isConnect          bool
}

// 触发事件回调函数集合
type emitCallbackCollection struct {
	emitCollection sync.Map
}

// 触发同步事件集合
type emitSyncCollection struct {
	mutex          *sync.Mutex
	emitCollection sync.Map
}

func (m *ipcChannel) newRenderChannel(memoryAddresses ...string) {
	if UseNetIPCChannel {
		address := fmt.Sprintf("localhost:%d", IPC.port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		m.render.ipcType = IPCT_NET
		m.render.netConn = conn
	} else {
		memoryAddr := ipcSock
		if len(memoryAddresses) > 0 {
			memoryAddr = memoryAddresses[0]
		}
		unixAddr, err := net.ResolveUnixAddr(memoryNetwork, memoryAddr)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		unixConn, err := net.DialUnix(memoryNetwork, nil, unixAddr)
		if err != nil {
			panic("Client failed to connect to IPC service Error: " + err.Error())
		}
		m.render.ipcType = IPCT_UNIX
		m.render.unixAddr = unixAddr
		m.render.unixConn = unixConn
	}
	go m.render.receive()
}

func (m *emitCallbackCollection) remove(key int32) {
	m.emitCollection.Delete(key)
}

func (m *renderChannel) close() {
	if m.unixConn != nil {
		m.unixConn.Close()
		m.unixConn = nil
	}
	if m.netConn != nil {
		m.netConn.Close()
		m.netConn = nil
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
	callBack := m.events.get(name)
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
	ipcWrite(tm_Async, m.channelId, m.msgID.new(), []byte(eventName), arguments.Package(), m.conn())
}

func (m *renderChannel) EmitAndCallback(eventName string, arguments IArgumentList, callback ipcCallback) {
	if m.conn() == nil {
		return
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if arguments == nil {
		arguments = NewArgumentList()
	}
	eventId := m.msgID.new()
	m.emitCallback.emitCollection.Store(eventId, callback)
	ipcWrite(tm_Callback, m.channelId, eventId, []byte(eventName), arguments.Package(), m.conn())
}

func (m *renderChannel) EmitAndReturn(eventName string, arguments IArgumentList) IIPCContext {
	if m.conn() == nil {
		return nil
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var emit = func(emitAsync *emitSyncCollection) IIPCContext {
		emitAsync.mutex.Lock()
		defer emitAsync.mutex.Unlock()
		eventId := m.msgID.new()
		var chn = make(chan IIPCContext)
		emitAsync.emitCollection.Store(eventId, chn)
		ipcWrite(tm_Sync, m.channelId, eventId, []byte(eventName), arguments.Package(), m.conn())
		return <-chn
	}
	if arguments == nil {
		arguments = NewArgumentList()
	}
	if emitAsync, ok := m.emitSync[eventName]; ok {
		return emit(emitAsync)
	} else {
		m.emitSync[eventName] = &emitSyncCollection{mutex: new(sync.Mutex), emitCollection: sync.Map{}}
		return emit(m.emitSync[eventName])
	}
}

func (m *renderChannel) conn() net.Conn {
	if m.ipcType == IPCT_NET {
		return m.netConn
	}
	return m.unixConn
}

func (m *renderChannel) emitConnect() {
	args := NewArgumentList()
	args.SetString(0, "-connecting")
	m.Emit(ln_onConnectEvent, args)
	m.isConnect = true
}

func (m *renderChannel) receive() {
	defer func() {
		m.close()
	}()
	var readHandler = &ipcReadHandler{
		browserId: m.browserId,
		channelId: m.channelId,
		ct:        ct_Client,
		ipcType:   m.ipcType,
		unixConn:  m.unixConn,
		unixAddr:  m.unixAddr,
		netConn:   m.netConn,
		handler: func(ctx *IPCContext) {
			if m.call(ctx.eventName, ctx) {
				if (ctx.triggerMode == tm_Callback || ctx.triggerMode == tm_Sync) && !ctx.isReply {
					ctx.Response([]byte{})
				}
			} else {
				if ctx.triggerMode == tm_Callback { //回调函数
					m.mutex.Lock()
					defer m.mutex.Unlock()
					if callback, ok := m.emitCallback.emitCollection.Load(ctx.eventId); ok {
						callback.(ipcCallback)(ctx)
						m.emitCallback.emitCollection.Delete(ctx.eventId)
					}
				} else if ctx.triggerMode == tm_Sync { //同步调用
					if emitAsync, ok := m.emitSync[ctx.eventName]; ok {
						if chn, ok := emitAsync.emitCollection.Load(ctx.eventId); ok {
							var c = chn.(chan IIPCContext)
							c <- ctx
							close(c)
							emitAsync.emitCollection.Delete(ctx.eventId)
						}
					}
				}
			}
		},
	}
	ipcRead(readHandler)
}
