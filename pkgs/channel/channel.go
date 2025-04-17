//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC Channel
// Communication between multiple processes

package channel

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/common"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
	"github.com/energye/golcl/lcl/rtl/version"
	"math"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	MemoryNetwork = "unix"
)

var (
	protocolHeader       = []byte{0x01, 0x09, 0x08, 0x07, 0x00, 0x08, 0x02, 0x02}                                                                 // 协议头
	protocolHeaderLength = int32(len(protocolHeader))                                                                                             // 协议头长度
	messageTypeLength    = int32(1)                                                                                                               // 消息类型 int8
	processIdLength      = int32(1)                                                                                                               // 消息来源 int8
	channelIdLength      = int32(8)                                                                                                               // 发送通道 int64
	toChannelIdLength    = int32(8)                                                                                                               // 接收通道 int64
	dataByteLength       = int32(4)                                                                                                               // 数据长度 int32
	headerLength         = int(protocolHeaderLength + messageTypeLength + processIdLength + channelIdLength + toChannelIdLength + dataByteLength) // 协议头长度
)

var (
	memoryAddress    = "energy.sock" //
	ipcSock          string          // sock path
	useNetIPCChannel = false         //
	port             = 0             // net ipc default: 0
)

// IPCNetSocketPortKey IPC 监听端口号的Key名, 在初始化之前该值可被改变
var IPCNetSocketPortKey = "--energy-ipc-net-socket-port"

// mt 消息类型
type mt int8

const (
	mt_invalid           mt = iota - 1 // 无效类型
	mt_connection                      // 建立链接消息
	mt_connectd                        // 已链接消息
	mt_update_channel_id               // 更新通道ID消息
	mt_common                          // 普通消息
	mt_relay                           // 转发消息
)

// IPCCallback 回调
type IPCCallback func(context IIPCContext)

func init() {
	ipcSock = filepath.Join(os.TempDir(), memoryAddress)
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

// IsUseNetIPC
//
// 当前IPC使用的通道类型

// MacOS, Linux, Windows10 && Build >= 17063 时使用 unix socket
// Windows10 以下 && Windows10 Build < 17063 时使用 net socket
func IsUseNetIPC() bool {
	if common.IsDarwin() || common.IsLinux() {
		return false
	}
	goVersion := strings.Replace(runtime.Version(), "go", "", -1)
	vers := strings.Split(goVersion, ".")
	supoortWindows := common.StrToInt32(vers[0]) >= 1 && common.StrToInt32(vers[1]) >= 14 // Go 版本 >= 1.14 支持Windows
	ov := version.OSVersion
	if ((ov.Major > 10) || (ov.Major == 10 && ov.Build >= 17063)) && supoortWindows {
		return false
	}
	return true
}

// SetPort 设置 net socket 端口号, 如参数 "v" 非指定范围内端口则获取随机未使用端口号
//
// v 1024 ~ 65535
func SetPort(v int) {
	if v >= 1024 && v < 65535 {
		port = v
	} else {
		port = 0
		port = Port()
	}
}

// Port 获取并返回未使用的net socket端口
func Port() int {
	if port != 0 {
		return port
	}
	//主进程获取端口号
	if process.Args.IsMain() {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			panic("Failed to Get unused Port number Error: " + err.Error())
		}
		listen, err := net.ListenTCP("tcp", addr)
		if err != nil {
			panic("Failed to Get unused Port number Error: " + err.Error())
		}
		defer listen.Close()
		port = listen.Addr().(*net.TCPAddr).Port
		// 主进程(浏览器进程)环境变量以使子进程获取这个端口
		os.Setenv(IPCNetSocketPortKey, strconv.Itoa(port))
	} else if process.Args.IsRender() {
		// 子进程(渲染进程)获取这个端口号
		nsp := os.Getenv(IPCNetSocketPortKey)
		port, _ = strconv.Atoi(nsp)
	}
	return port
}

// IIPCContext IPC通信回调上下文
type IIPCContext interface {
	Connect() net.Conn        // IPC 通道链接
	ChannelId() string        // 返回 发送通道ID
	ToChannelId() string      // 返回 接收发送通道ID
	ChannelType() ChannelType // 返回 当前通道类型
	ProcessId() CefProcessId  // 返回 通道消息来源
	Message() IMessage        // 消息
	Free()                    //
}

// IMessage 消息内容接口
type IMessage interface {
	Type() mt        // 消息类型
	Length() int32   // 数据长度
	Data() []byte    // 数据
	JSON() json.JSON // 转为 JSON 对象并返回
	clear()          // 清空
}

// IChannel 通道链接
type IChannel interface {
	IsConnect() bool
	Close()
	read(b []byte) (n int, err error)
	write(messageType mt, channelId, toChannelId string, data []byte) (n int, err error)
}

type IBrowserChannel interface {
	Channel(channelId string) IChannel
	ChannelIds() (result []int64)
	Send(channelId string, data []byte)
	Handler(handler IPCCallback)
	Close()
}

type IRenderChannel interface {
	Channel() IChannel
	Send(data []byte)
	SendToChannel(toChannelId string, data []byte)
	UpdateChannelId(toChannelId string)
	Handler(handler IPCCallback)
	Close()
}

// ipcMessage 消息内容
type ipcMessage struct {
	t mt     // type
	s int32  // size
	v []byte // data
}

// IPCContext IPC 上下文
type IPCContext struct {
	channelId   string       //render channelId
	toChannelId string       //
	ipcType     IPC_TYPE     // ipc type
	channelType ChannelType  // ipc channel type
	processId   CefProcessId // ipc msg source, browser or render
	connect     net.Conn     // connect
	message     IMessage     // message
}

// Free 释放消息内存空间
func (m *IPCContext) Free() {
	if m.message != nil {
		m.message.clear()
		m.message = nil
	}
}

// ChannelId 返回发送通道ID
func (m *IPCContext) ChannelId() string {
	return m.channelId
}

// ToChannelId 返回接收通道ID
func (m *IPCContext) ToChannelId() string {
	return m.toChannelId
}

// ChannelType 返回当前通道类型
func (m *IPCContext) ChannelType() ChannelType {
	return m.channelType
}

func (m *IPCContext) ProcessId() CefProcessId {
	return m.processId
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
func (m *ipcMessage) Length() int32 {
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
	channelId   string
	isConnect   bool
	conn        net.Conn
	ipcType     IPC_TYPE
	channelType ChannelType
	handler     IPCCallback
}

// IsConnect return is connect success
func (m *channel) IsConnect() bool {
	if m == nil {
		return false
	}
	return m.isConnect
}

// Close the current IPC channel connect
func (m *channel) Close() {
	if m.conn != nil {
		m.conn.Close()
		m.conn = nil
	}
}

// read data
func (m *channel) read(b []byte) (n int, err error) {
	if m.ipcType == IPCT_NET {
		return m.conn.Read(b)
	} else {
		n, _, err := m.conn.(*net.UnixConn).ReadFromUnix(b)
		return n, err
	}
}

// write data
func (m *channel) write(messageType mt, channelId, toChannelId string, data []byte) (n int, err error) {
	defer func() {
		data = nil
	}()
	if m.conn == nil {
		return 0, errors.New("channel link not established successfully")
	}
	var (
		dataByteLen = len(data)
	)
	if dataByteLen > math.MaxInt32 {
		return 0, errors.New("exceeded maximum message length")
	}
	var processId CefProcessId
	if m.channelType == Ct_Server {
		processId = PID_BROWSER
	} else {
		processId = PID_RENDER
	}
	var writeBuf = new(bytes.Buffer)
	_ = binary.Write(writeBuf, binary.BigEndian, protocolHeader)     //protocol header
	_ = binary.Write(writeBuf, binary.BigEndian, int8(messageType))  //message type
	_ = binary.Write(writeBuf, binary.BigEndian, int8(processId))    //source of information
	_ = binary.Write(writeBuf, binary.BigEndian, channelId)          //source channel Id
	_ = binary.Write(writeBuf, binary.BigEndian, toChannelId)        //to     channel Id
	_ = binary.Write(writeBuf, binary.BigEndian, int32(dataByteLen)) //data length
	_ = binary.Write(writeBuf, binary.BigEndian, data)               //data bytes
	n, err = m.conn.Write(writeBuf.Bytes())
	writeBuf.Reset()
	writeBuf = nil
	return n, err
}

// ipcRead Read channel messages
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
		logger.Debug("IPC Read Disconnect type:", ipcType, "ChannelType:", chnType, "processType:", process.Args.ProcessType())
		m.Close()
	}()
	for {
		header := make([]byte, headerLength)
		size, err := m.read(header)
		if err != nil {
			logger.Debug("IPC Read [Error] type:", ipcType, "ChannelType:", chnType, "Error:", err)
			return
		} else if size == 0 {
			logger.Debug("IPC Read [Size == 0] type:", ipcType, "ChannelType:", chnType, "header:", header, "Error:", err)
			return
		}
		if size == headerLength {
			for i, protocol := range protocolHeader {
				if header[i] != protocol {
					logger.Debug("check header protocol error", i, header[i], protocol)
					return
				}
			}
			var (
				t, proId               int8   //
				channelId, toChannelId string //
				dataLen                int32  //数据长度
				low, high              int32  //
			)
			//message type
			low = protocolHeaderLength
			high = protocolHeaderLength + messageTypeLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &t)
			if err != nil {
				logger.Debug("binary.Read.t: ", err)
				return
			}
			//message source
			low = high
			high = high + processIdLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &proId)
			if err != nil {
				logger.Debug("binary.Read.t: ", err)
				return
			}
			//send channel id
			low = high
			high = high + channelIdLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &channelId)
			if err != nil {
				logger.Debug("binary.Read.channelId: ", err)
				return
			}

			//receive channel id
			low = high
			high = high + toChannelIdLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &toChannelId)
			if err != nil {
				logger.Debug("binary.Read.toChannelId: ", err)
				return
			}

			//data length
			low = high
			high = high + dataByteLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &dataLen)
			if err != nil {
				logger.Debug("binary.Read.dataLen: ", err)
				return
			}
			//data
			dataByte := make([]byte, dataLen)
			if dataLen > 0 {
				size, err = m.read(dataByte)
			}
			if err != nil {
				logger.Debug("binary.Read.dataByte: ", err)
				return
			}
			// call handler
			m.handler(&IPCContext{
				channelId:   channelId,
				toChannelId: toChannelId,
				ipcType:     m.ipcType,
				connect:     m.conn,
				channelType: m.channelType,
				processId:   CefProcessId(proId),
				message: &ipcMessage{ // message data
					t: mt(t),
					s: dataLen,
					v: dataByte,
				},
			})
		} else {
			logger.Debug("invalid != headerLength")
			break
		}
	}
}
