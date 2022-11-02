//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
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
	"github.com/energye/golcl/lcl/api"
	"math"
	"net"
	"os"
	"unsafe"
)

var (
	protocolHeader                  = []byte{0x01, 0x09, 0x08, 0x07, 0x00, 0x08, 0x02, 0x02}
	protocolHeaderLength            = int32(len(protocolHeader))
	triggerModeByteLength     int32 = 1
	renderChannelIdByteLength int32 = 8
	eventIdByteLength         int32 = 4
	eventNameByteLength       int32 = 2
	dataByteLength            int32 = 4
	headerLength                    = int(protocolHeaderLength + triggerModeByteLength + renderChannelIdByteLength + eventIdByteLength + eventNameByteLength + dataByteLength)
)

type IPCCallback func(context IIPCContext)

// 进程间IPC通信回调上下文
type IIPCContext interface {
	setArguments(argument IArgumentList)
	Connect() net.Conn
	EventId() int32
	ChannelId() int64          //render channel channelId
	BrowserId() int32          //render channel browserId
	Message() *IPCEventMessage //接收的消息数据 一搬配合Response函数
	Response(data []byte)      //通过响应应达的方式返回结果,在data符合[argumentList]规范可使用Arguments() 一搬配合Message函数
	Free()                     //释放内存
	Arguments() IArgumentList  //用于 ipc 进程之间通信, 带有参数类型
	Result() *IPCContextResult //用于 render(js)进程触发 browser(go)进程监听返回值, 带有固定几个参数类型
}

type IEventMessage interface {
	Data() []byte
	DataLen() int32
}

type EventCallback func(context IIPCContext)

type ipcReadHandler struct {
	browserId int32
	channelId int64
	ipcType   IPC_TYPE
	ct        ChannelType
	connect   net.Conn
	handler   func(ctx *IPCContext)
}

type IPCEventMessage struct {
	dataLen int32
	data    []byte
}

type IPCContextResult struct {
	vType  V8_JS_VALUE_TYPE
	result unsafe.Pointer //[]byte
}

func (m *IPCContextResult) Result() unsafe.Pointer {
	return m.result
}
func (m *IPCContextResult) VType() V8_JS_VALUE_TYPE {
	return m.vType
}

func (m *IPCContextResult) SetString(ret string) {
	m.result = unsafe.Pointer(api.GoStrToDStr(ret)) //StringToBytes(ret)
	m.vType = V8_VALUE_STRING
}

func (m *IPCContextResult) SetInt(ret int32) {
	m.result = unsafe.Pointer(uintptr(ret)) //Int32ToBytes(ret)
	m.vType = V8_VALUE_INT
}

func (m *IPCContextResult) SetFloat64(ret float64) {
	m.result = unsafe.Pointer(&ret) //Float64ToBytes(ret)
	m.vType = V8_VALUE_DOUBLE
}

func (m *IPCContextResult) SetBool(ret bool) {
	m.result = unsafe.Pointer(api.GoBoolToDBool(ret)) //[]byte{BoolToByte(ret)
	m.vType = V8_VALUE_BOOLEAN
}

func (m *IPCContextResult) clear() {
	m.vType = 0
	m.result = nil
}

// IPC 上下文
type IPCContext struct {
	browserId    int32             //render channel browserId
	channelId    int64             //render channel channelId
	eventId      int32             //
	ipcType      IPC_TYPE          //
	isReply      bool              //回复状态 true已回复 false未回复
	triggerMode  TriggerMode       //触发模式 0异步 1回调 2同步
	eventName    string            //
	connect      net.Conn          //
	eventMessage *IPCEventMessage  //
	arguments    IArgumentList     //
	result       *IPCContextResult //用于 render(js)进程触发 browser(go)进程监听返回值
}

// 事件集合
type event struct {
	event map[string]EventCallback
}

func NewIPCContext(eventName string, browserId int32, channelId int64, ipcType IPC_TYPE, conn net.Conn, eventMessage *IPCEventMessage, result *IPCContextResult, arguments IArgumentList) IIPCContext {
	return &IPCContext{
		eventName:    eventName,
		browserId:    browserId,
		channelId:    channelId,
		ipcType:      ipcType,
		connect:      conn,
		eventMessage: eventMessage,
		result:       result,
		arguments:    arguments,
	}
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

func (m *IPCContext) Result() *IPCContextResult {
	return m.result
}

func (m *IPCContext) Free() {
	if m.arguments != nil {
		m.arguments.Clear()
		m.arguments = nil
	}
	if m.eventMessage != nil {
		m.eventMessage.clear()
		m.eventMessage = nil
	}
	if m.result != nil {
		m.result.clear()
		m.result = nil
	}
}
func (m *IPCContext) EventId() int32 {
	return m.eventId
}

func (m *IPCContext) Arguments() IArgumentList {
	if m.arguments == nil {
		m.arguments = NewArgumentList()
		m.arguments.UnPackage(m.Message().Data())
	}
	return m.arguments
}

func (m *IPCContext) setArguments(argument IArgumentList) {
	m.arguments = argument
}

func (m *IPCContext) Response(data []byte) {
	_, _ = ipcWrite(m.triggerMode, m.channelId, m.eventId, []byte(m.eventName), data, m.Connect())
	m.isReply = true
	data = nil
}

func (m *IPCContext) ChannelId() int64 {
	return m.channelId
}
func (m *IPCContext) BrowserId() int32 {
	return m.browserId
}

func (m *IPCContext) Message() *IPCEventMessage {
	return m.eventMessage
}

func (m *IPCContext) Connect() net.Conn {
	return m.connect
}

func (m *IPCEventMessage) Data() []byte {
	return m.data
}

func (m *IPCEventMessage) DataLen() int32 {
	return m.dataLen
}
func (m *IPCEventMessage) clear() {
	m.data = nil
	m.dataLen = 0
}

func removeMemory() {
	os.Remove(ipcSock)
}

func ipcWrite(triggerMode TriggerMode, channelId int64, eventId int32, eventName, data []byte, conn net.Conn) (n int, err error) {
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
	binary.Write(ipcWriteBuf, binary.BigEndian, protocolHeader)        //协议头
	binary.Write(ipcWriteBuf, binary.BigEndian, triggerMode)           //触发模式 0异步 1回调 2同步
	binary.Write(ipcWriteBuf, binary.BigEndian, channelId)             //render channel Id
	binary.Write(ipcWriteBuf, binary.BigEndian, eventId)               //事件ID
	binary.Write(ipcWriteBuf, binary.BigEndian, int16(len(eventName))) //监听事件名长度
	binary.Write(ipcWriteBuf, binary.BigEndian, int32(dataByteLen))    //数据长度
	binary.Write(ipcWriteBuf, binary.BigEndian, eventName)             //监听事件名
	binary.Write(ipcWriteBuf, binary.BigEndian, data)                  //数据
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
		logger.Debug("IPC Read Disconnect type:", ipcType, "ChannelType:", chnType, "channelId:", handler.channelId, "processType:", common.Args.ProcessType())
		handler.Close()
	}()
	for {
		header := make([]byte, headerLength)
		size, err := handler.Read(header)
		if err != nil {
			logger.Debug("IPC Read【Error】IPCType:", ipcType, "ChannelType:", chnType, "channelId:", handler.channelId, "Error:", err)
			return
		} else if size == 0 {
			logger.Debug("IPC Read【Size == 0】IPCType:", ipcType, "ChannelType:", chnType, "channelId:", handler.channelId, "header:", header, "Error:", err)
			return
		}
		if size == headerLength {
			for i, protocol := range protocolHeader {
				if header[i] != protocol {
					return
				}
			}
			//Logger.Debug("IPC Read protocolHeaderLength:", header)
			var (
				triggerMode TriggerMode
				channelId   int64 //render channel Id
				eventId     int32
				eventLen    int16
				dataLen     int32
				eventName   string
				low, high   int32
			)
			low = protocolHeaderLength
			high = protocolHeaderLength + triggerModeByteLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &triggerMode)
			if err != nil {
				logger.Debug("binary.Read.triggerMode: ", err)
				return
			}
			low = high
			high = high + renderChannelIdByteLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &channelId)
			if err != nil {
				logger.Debug("binary.Read.channelId: ", err)
				return
			}
			low = high
			high = high + eventIdByteLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &eventId)
			if err != nil {
				logger.Debug("binary.Read.eventIdByteLength: ", err)
				return
			}
			low = high
			high = high + eventNameByteLength
			err = binary.Read(bytes.NewReader(header[low:high]), binary.BigEndian, &eventLen)
			if err != nil {
				logger.Debug("binary.Read.eventLen: ", err)
				return
			}
			low = high
			err = binary.Read(bytes.NewReader(header[low:headerLength]), binary.BigEndian, &dataLen)
			if err != nil {
				logger.Debug("binary.Read.dataLen: ", err)
				return
			}

			eventNameByte := make([]byte, eventLen)
			size, err = handler.Read(eventNameByte)
			if err != nil {
				logger.Debug("binary.Read.eventNameByte: ", err)
				return
			}
			eventName = string(eventNameByte[:size])

			dataByte := make([]byte, dataLen)
			if dataLen > 0 {
				size, err = handler.Read(dataByte)
			}
			if err != nil {
				logger.Debug("binary.Read.data: ", err)
				return
			}
			ctx := &IPCContext{
				ipcType:     handler.ipcType,
				triggerMode: triggerMode,
				channelId:   channelId,
				eventId:     eventId,
				eventName:   eventName,
				connect:     handler.connect,
				eventMessage: &IPCEventMessage{
					dataLen: dataLen,
					data:    dataByte,
				},
			}
			if handler.channelId == 0 {
				handler.channelId = ctx.channelId
			}
			handler.handler(ctx)
		} else {
			logger.Debug("无效的 != headerLength")
			break
		}
	}
}
