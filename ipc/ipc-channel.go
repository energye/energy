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
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/energy/pkgs/json"
	"github.com/energye/golcl/lcl/rtl/version"
	"math"
	"net"
	"os"
	"path/filepath"
	"sync"
)

var (
	protocolHeader       = []byte{0x01, 0x09, 0x08, 0x07, 0x00, 0x08, 0x02, 0x02}   //协议头
	protocolHeaderLength = int32(len(protocolHeader))                               //协议头长度
	messageType          = int32(1)                                                 //消息类型 int8
	dataByteLength       = int32(4)                                                 //数据长度 int32
	headerLength         = int(protocolHeaderLength + messageType + dataByteLength) //协议头长度
)

var (
	memoryAddress    = "energy.sock"
	useNetIPCChannel = false
	ipcSock          string
	IPC              = &ipcChannel{
		browser: &browserChannel{
			channel: sync.Map{},
			mutex:   sync.Mutex{},
		},
		render: &renderChannel{
			mutex: sync.Mutex{},
		},
	}
)

//消息类型
type mt int8

const (
	mt_connection mt = iota //建立链接消息
	mt_common               //普通消息
)

const (
	channelId = "channelId"
)

func init() {
	//useNetIPCChannel = isUseNetIPC()
	ipcSock = filepath.Join(os.TempDir(), memoryAddress)
}

type ipcChannel struct {
	port    int
	browser *browserChannel
	render  *renderChannel
}

func (m *ipcChannel) Port() int {
	if m.port != 0 {
		return m.port
	}
	//主进程获取端口号
	if common.Args.IsMain() {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			panic("Failed to Get unused Port number Error: " + err.Error())
		}
		listen, err := net.ListenTCP("tcp", addr)
		if err != nil {
			panic("Failed to Get unused Port number Error: " + err.Error())
		}
		defer listen.Close()
		m.port = listen.Addr().(*net.TCPAddr).Port
	}
	return m.port
}

func isUseNetIPC() bool {
	if common.IsDarwin() || common.IsLinux() {
		return false
	}
	ov := version.OSVersion
	if (ov.Major > 10) || (ov.Major == 10 && ov.Build >= 17063) {
		//不支持UnixSocket
		return false
	}
	return true
}

func (m *ipcChannel) Browser() *browserChannel {
	return m.browser
}

func (m *ipcChannel) Render() *renderChannel {
	return m.render
}

// 主进程事件emit
type IIPCChannel interface {
	Close()
	Channel(channelId int64) *channel //IPC 获取指定的通道
	ChannelIds() (result []int64)     //IPC 获取所有通道
}

type IPCCallback func(context IIPCContext)
type messageCallback func(context IMessage)

// 进程间IPC通信回调上下文
type IIPCContext interface {
	Connect() net.Conn //IPC 链接
	ChannelId() int64  //render channel channelId
	BrowserId() int32  //render channel browserId
	Message() IMessage //
	Free()             //
}

type IMessage interface {
	Type() mt
	Length() int32
	Data() []byte
	JSON() json.JSON
	clear()
}

type ipcReadHandler struct {
	ipcType IPC_TYPE
	ct      ChannelType
	connect net.Conn
	handler IPCCallback
}

type ipcMessage struct {
	t mt
	s int32
	v []byte
}

// IPC 上下文
type IPCContext struct {
	browserId int32    //render channel browserId
	frameId   int64    //render channel frameId
	ipcType   IPC_TYPE //
	connect   net.Conn //
	message   IMessage //
}

func (m *ipcReadHandler) Close() {
	if m.connect != nil {
		m.connect.Close()
		m.connect = nil
	}
}

func (m *ipcReadHandler) Read(b []byte) (n int, err error) {
	if m.ipcType == IPCT_NET {
		return m.connect.Read(b)
	} else {
		n, _, err := m.connect.(*net.UnixConn).ReadFromUnix(b)
		return n, err
	}
}

func (m *IPCContext) Free() {
	if m.message != nil {
		m.message.clear()
		m.message = nil
	}
}

func (m *IPCContext) ChannelId() int64 {
	return m.frameId
}

func (m *IPCContext) BrowserId() int32 {
	return m.browserId
}

func (m *IPCContext) Message() IMessage {
	return m.message
}

func (m *IPCContext) Connect() net.Conn {
	return m.connect
}

func (m *ipcMessage) Type() mt {
	return m.t
}

func (m *ipcMessage) Data() []byte {
	return m.v
}

func (m *ipcMessage) Length() int32 {
	return m.s
}

func (m *ipcMessage) JSON() json.JSON {
	return json.NewJSON(m.v)
}

func (m *ipcMessage) clear() {
	m.v = nil
	m.s = 0
}

func ipcWrite(messageType mt, data []byte, conn net.Conn) (n int, err error) {
	defer func() {
		data = nil
	}()
	if conn == nil {
		return 0, errors.New("链接未建立成功")
	}
	var (
		dataByteLen = len(data)
	)
	if dataByteLen > math.MaxInt32 {
		return 0, errors.New("超出最大消息长度")
	}
	var ipcWriteBuf = new(bytes.Buffer)
	binary.Write(ipcWriteBuf, binary.BigEndian, protocolHeader)     //协议头
	binary.Write(ipcWriteBuf, binary.BigEndian, int8(messageType))  //消息类型
	binary.Write(ipcWriteBuf, binary.BigEndian, int32(dataByteLen)) //数据长度
	binary.Write(ipcWriteBuf, binary.BigEndian, data)               //数据
	n, err = conn.Write(ipcWriteBuf.Bytes())
	ipcWriteBuf.Reset()
	return n, err
}

func ipcRead(handler *ipcReadHandler) {
	var ipcType, chnType string
	if handler.ipcType == IPCT_NET {
		ipcType = "[net]"
	} else {
		ipcType = "[unix]"
	}
	if handler.ct == Ct_Server {
		chnType = "[server]"
	} else {
		chnType = "[client]"
	}
	defer func() {
		logger.Debug("IPC Read Disconnect type:", ipcType, "ChannelType:", chnType, "processType:", common.Args.ProcessType())
		handler.Close()
	}()
	for {
		header := make([]byte, headerLength)
		size, err := handler.Read(header)
		if err != nil {
			logger.Debug("IPC Read【Error】IPCType:", ipcType, "ChannelType:", chnType, "Error:", err)
			return
		} else if size == 0 {
			logger.Debug("IPC Read【Size == 0】IPCType:", ipcType, "ChannelType:", chnType, "header:", header, "Error:", err)
			return
		}
		if size == headerLength {
			for i, protocol := range protocolHeader {
				if header[i] != protocol {
					return
				}
			}
			var (
				t         int8  //消息类型
				dataLen   int32 //数据长度
				low, high int32 //
			)
			//消息类型
			low = protocolHeaderLength
			high = protocolHeaderLength + messageType
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &t)
			if err != nil {
				logger.Debug("binary.Read.length: ", err)
				return
			}

			//数据长度
			low = high
			high = high + dataByteLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &dataLen)
			if err != nil {
				logger.Debug("binary.Read.length: ", err)
				return
			}
			//数据
			dataByte := make([]byte, dataLen)
			if dataLen > 0 {
				size, err = handler.Read(dataByte)
			}
			if err != nil {
				logger.Debug("binary.Read.data: ", err)
				return
			}
			handler.handler(&IPCContext{
				ipcType: handler.ipcType,
				connect: handler.connect,
				message: &ipcMessage{
					t: mt(t),
					s: dataLen,
					v: dataByte,
				},
			})
		} else {
			logger.Debug("无效的 != headerLength")
			break
		}
	}
}
