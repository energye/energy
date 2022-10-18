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
	"github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl/rtl/version"
	"net"
	"sync"
)

var (
	//GO IPC 通道选择条件
	//
	//默认值=IsDarwin() true:使用net socket, false:使用unix socket
	UseNetIPCChannel = true
	ipcSock          string
	//主进程Browser启动 IPC，Render进程创建 IPC
	IPC = &ipcChannel{
		browser: &browserChannel{
			msgID:              &MsgID{},
			cliID:              &CliID{},
			events:             &event{event: make(map[string]EventCallback)},
			channel:            make(map[int64]*channel),
			emitSync:           make(map[string]*EmitSyncCollection),
			mutex:              sync.Mutex{},
			emitCallback:       &EmitCallbackCollection{EmitCollection: sync.Map{}},
			browseOnEvents:     make([]func(browseProcess IEventOn), 0),
			browseEmitCallback: make([]func(renderProcess IEventEmit), 0)},
		render: &renderChannel{
			msgID:              &MsgID{},
			mutex:              sync.Mutex{},
			emitCallback:       &EmitCallbackCollection{EmitCollection: sync.Map{}},
			emitSync:           make(map[string]*EmitSyncCollection),
			events:             &event{event: make(map[string]EventCallback)},
			renderOnEvents:     make([]func(browseProcess IEventOn), 0),
			renderEmitCallback: make([]func(renderProcess IEventEmit), 0),
		},
	}
)

func init() {
	ipcSock = fmt.Sprintf("%s%sgolcl%s%s", consts.HomeDir, consts.Separator, consts.Separator, consts.MemoryAddress)
}

func IPCChannelChooseInit() {
	UseNetIPCChannel = isUseNetIPC()
}

func isUseNetIPC() bool {
	if IsDarwin() || IsLinux() {
		return false
	}
	ov := version.OSVersion
	if (ov.Major > 10) || (ov.Major == 10 && ov.Build >= 17063) {
		//不支持UnixSocket
		return false
	}
	//使用net socket
	return true
}

// 主Browser进程和Render进程事件on
type IEventOn interface {
	On(name string, eventCallback EventCallback) //IPC 事件监听
	Close()
}

// 主Browser进程和Render进程事件emit
type IEventEmit interface {
	IEventOn
	Events() *event
	Channel(channelId int64) *channel
	SetOnEvent(callback func(event IEventOn))                                        //IPC 事件监听
	Emit(eventName string, arguments IArgumentList)                                  //IPC 异步事件触发
	EmitAndCallback(eventName string, arguments IArgumentList, callback IPCCallback) //IPC 回调事件触发
	EmitAndReturn(eventName string, arguments IArgumentList) IIPCContext             //IPC 返回值事件触发(处理时间复杂操作尽量不使用，容易造成UI进程锁死)
}

// 主进程事件emit
type IBrowseEventEmit interface {
	IEventOn
	IEventEmit
	EmitChannelId(eventName string, channelId int64, arguments IArgumentList)                                  //IPC 异步事件触发-指定通道ID
	EmitChannelIdAndCallback(eventName string, channelId int64, arguments IArgumentList, callback IPCCallback) //IPC 回调事件触发-指定通道ID
	EmitChannelIdAndReturn(eventName string, channelId int64, arguments IArgumentList) IIPCContext             //IPC 返回值事件触发(处理时间复杂操作尽量不使用，容易造成UI进程锁死)-指定通道ID
}

func getFreePort() int {
	//主进程获取端口号
	if Args.IsMain() {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			panic("Failed to Get unused Port number Error: " + err.Error())
		}
		listen, err := net.ListenTCP("tcp", addr)
		if err != nil {
			panic("Failed to Get unused Port number Error: " + err.Error())
		}
		defer listen.Close()
		return listen.Addr().(*net.TCPAddr).Port
	}
	return 0
}

type ipcChannel struct {
	serverIsStart chan int
	port          int
	browser       *browserChannel
	render        *renderChannel
}

func (m *ipcChannel) closeClient() {
	if m.render != nil {
		m.render.Close()
	}
}
func (m *ipcChannel) Port() int {
	return m.port
}
func (m *ipcChannel) SetPort(port ...int) {
	if len(port) > 0 {
		m.port = port[0]
	} else {
		if m.port == 0 {
			m.port = getFreePort()
		}
	}
}

func (m *ipcChannel) Browser() IBrowseEventEmit {
	return m.browser
}

func (m *ipcChannel) Render() IEventEmit {
	return m.render
}

// 启动IPC服务
func (m *ipcChannel) StartBrowserIPC() {
	logger.Info("Create IPC browser")
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		m.SetPort()
		m.newBrowseChannel()
		defer m.browser.Close()
		if m.browser.browseOnEvents != nil {
			for _, cb := range m.browser.browseOnEvents {
				if cb != nil {
					cb(m.browser)
				}
			}
		}
		if m.browser.browseEmitCallback != nil {
			for _, cb := range m.browser.browseEmitCallback {
				if cb != nil {
					cb(m.browser)
				}
			}
		}
		group.Done()
		m.browser.accept()
	}()
	group.Wait()
}

// 创建IPC客户端服务 基于unix socket, windows <= 10.17063 基于net socket
//
//单进程，只创建一个连接
//
//多进程，每个渲染进程创建一个连接
func (m *ipcChannel) CreateRenderIPC(browserId int32, channelId int64) *renderChannel {
	logger.Info("Create IPC render isConnect:", m.render.isConnect, "channelId:", channelId)
	if m.render.isConnect {
		return m.render
	}
	//m.closeClient()
	m.newRenderChannel()
	if m.render.renderOnEvents != nil {
		for _, cb := range m.render.renderOnEvents {
			if cb != nil {
				cb(m.render)
			}
		}
	}
	if m.render.renderEmitCallback != nil {
		for _, cb := range m.render.renderEmitCallback {
			if cb != nil {
				cb(m.render)
			}
		}
	}
	m.render.browserId = browserId
	m.render.channelId = channelId
	m.render.emitConnect()
	return m.render
}
