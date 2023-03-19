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
	protocolHeader       = []byte{0x01, 0x09, 0x08, 0x07, 0x00, 0x08, 0x02, 0x02}                           //协议头
	protocolHeaderLength = int32(len(protocolHeader))                                                       //协议头长度
	messageTypeLength    = int32(1)                                                                         //消息类型 int8
	channelIdLength      = int32(8)                                                                         //通道Id int64
	dataByteLength       = int32(4)                                                                         //数据长度 int32
	headerLength         = int(protocolHeaderLength + messageTypeLength + channelIdLength + dataByteLength) //协议头长度
)

var (
	memoryAddress    = "energy.sock"
	ipcSock          string
	useNetIPCChannel = false
	Channel          = &ipcChannel{
		browser: &browserChannel{
			channel: sync.Map{},
			mutex:   sync.Mutex{},
		},
		render: &renderChannel{
			mutex: sync.Mutex{},
		},
	}
)

//mt 消息类型
type mt int8

const (
	mt_invalid    mt = iota - 1 //无效类型
	mt_connection               //建立链接消息
	mt_common                   //普通消息
)

// channel key
const (
	key_channelId = "key_channelId"
)

type IPCCallback func(context IIPCContext)

func init() {
	ipcSock = filepath.Join(os.TempDir(), memoryAddress)
}

type ipcChannel struct {
	port    int
	browser *browserChannel
	render  *renderChannel
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

func isUseNetIPC() bool {
	if common.IsDarwin() || common.IsLinux() {
		return false
	}
	ov := version.OSVersion
	if (ov.Major > 10) || (ov.Major == 10 && ov.Build >= 17063) {
		return false
	}
	return true
}

// Port 获取并返回net socket端口
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

// Browser 返回 browser 通道
func (m *ipcChannel) Browser() *browserChannel {
	return m.browser
}

// Render 返回 render 通道
func (m *ipcChannel) Render() *renderChannel {
	return m.render
}

// IIPCContext IPC通信回调上下文
type IIPCContext interface {
	Connect() net.Conn // IPC 通道链接
	ChannelId() int64  // 通道ID
	Message() IMessage // 消息
	Free()             //
}

// IMessage 消息内容接口
type IMessage interface {
	Type() mt
	Length() uint32
	Data() []byte
	JSON() json.JSON
	clear()
}

// ipcMessage 消息内容
type ipcMessage struct {
	t mt     // type
	s uint32 // size
	v []byte // data
}

// IPCContext IPC 上下文
type IPCContext struct {
	channelId int64    //render channelId
	ipcType   IPC_TYPE // ipcType
	connect   net.Conn // connect
	message   IMessage // message
}

// Free 释放消息内存空间
func (m *IPCContext) Free() {
	if m.message != nil {
		m.message.clear()
		m.message = nil
	}
}

// ChannelId 返回通道ID
func (m *IPCContext) ChannelId() int64 {
	return m.channelId
}

// Message 返回消息内容
func (m *IPCContext) Message() IMessage {
	return m.message
}

// Connect 返回当前通道链接
func (m *IPCContext) Connect() net.Conn {
	return m.connect
}

// Type 消息类型
func (m *ipcMessage) Type() mt {
	return m.t
}

// Data 消息[]byte数据
func (m *ipcMessage) Data() []byte {
	return m.v
}

// Length 消息[]byte长度
func (m *ipcMessage) Length() uint32 {
	return m.s
}

// JSON 消息转为JSON对象
func (m *ipcMessage) JSON() json.JSON {
	return json.NewJSON(m.v)
}

// clear 清空内容
func (m *ipcMessage) clear() {
	m.t = mt_invalid
	m.v = nil
	m.s = 0
}

// channel 通道
type channel struct {
	writeBuf    *bytes.Buffer //
	channelId   int64         //通道ID
	conn        net.Conn      //通道链接 net or unix
	ipcType     IPC_TYPE      //IPC类型
	channelType ChannelType   //链接通道类型
	handler     IPCCallback   //
}

// Close 关闭当前ipc通道链接
func (m *channel) Close() {
	if m.conn != nil {
		m.conn.Close()
		m.conn = nil
	}
}

// Read 读取内容
func (m *channel) Read(b []byte) (n int, err error) {
	if m.ipcType == IPCT_NET {
		return m.conn.Read(b)
	} else {
		n, _, err := m.conn.(*net.UnixConn).ReadFromUnix(b)
		return n, err
	}
}

// ipcWrite 写入消息
func (m *channel) ipcWrite(messageType mt, channelId int64, data []byte) (n int, err error) {
	defer func() {
		data = nil
	}()
	if m.conn == nil {
		return 0, errors.New("通道链接未建立成功")
	}
	var (
		dataByteLen = len(data)
	)
	if dataByteLen > math.MaxUint32 {
		return 0, errors.New("超出最大消息长度")
	}
	_ = binary.Write(m.writeBuf, binary.BigEndian, protocolHeader)      //协议头
	_ = binary.Write(m.writeBuf, binary.BigEndian, int8(messageType))   //消息类型
	_ = binary.Write(m.writeBuf, binary.BigEndian, channelId)           //通道Id
	_ = binary.Write(m.writeBuf, binary.BigEndian, uint32(dataByteLen)) //数据长度
	_ = binary.Write(m.writeBuf, binary.BigEndian, data)                //数据
	n, err = m.conn.Write(m.writeBuf.Bytes())
	m.writeBuf.Reset()
	return n, err
}

// ipcRead 读取消息
func (m *channel) ipcRead() {
	var ipcType, chnType string
	if m.ipcType == IPCT_NET {
		ipcType = "[net]"
	} else {
		ipcType = "[unix]"
	}
	if m.channelType == Ct_Server {
		chnType = "[server]"
	} else {
		chnType = "[client]"
	}
	defer func() {
		logger.Debug("IPC Read Disconnect type:", ipcType, "ChannelType:", chnType, "processType:", common.Args.ProcessType())
		m.Close()
	}()
	for {
		header := make([]byte, headerLength)
		size, err := m.Read(header)
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
				t         int8 //消息类型
				channelId int64
				dataLen   uint32 //数据长度
				low, high int32  //
			)
			//消息类型
			low = protocolHeaderLength
			high = protocolHeaderLength + messageTypeLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &t)
			if err != nil {
				logger.Debug("binary.Read.length: ", err)
				return
			}
			//通道ID
			low = high
			high = high + channelIdLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &channelId)
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
				size, err = m.Read(dataByte)
			}
			if err != nil {
				logger.Debug("binary.Read.data: ", err)
				return
			}
			m.handler(&IPCContext{
				channelId: channelId,
				ipcType:   m.ipcType,
				connect:   m.conn,
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
