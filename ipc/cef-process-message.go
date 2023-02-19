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
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"reflect"
	"sync"
	"unsafe"
)

type IArgumentList interface {
	Clear()
	Size() int
	Package() []byte
	UnPackageBytePtr(bytePtr uintptr, byteLength int32)
	UnPackage(dataByte []byte)
	ToReflectValue() []reflect.Value
	ReflectValueConvert(ref []reflect.Value)
	capacity(index int)
	SetArguments(index int, argument IArgumentList)
	SetString(index int, v string, isDStr ...bool) //设置字符串参数 isDStr: 在传递参数到JavaScript需要设置为true
	GetString(index int) string
	SetIntAuto(index int, v interface{}, gov GO_VALUE_TYPE)
	SetInt(index int, v int)
	SetInt8(index int, v int8)
	SetData(index int, vType int8, data []byte)
	SetInt16(index int, v int16)
	SetInt32(index int, v int32)
	SetInt64(index int, v int64)
	GetInt(index int) int
	GetInt8(index int) int8
	GetInt16(index int) int16
	GetInt32(index int) int32
	GetBytes(index int) []byte
	GetData(index int) *ArgumentDataItem
	GetInt64(index int) int64
	SetFloatAuto(index int, v interface{}, gov GO_VALUE_TYPE)
	SetFloat32(index int, v float32)
	SetFloat64(index int, v float64)
	GetFloat32(index int) float32
	GetFloat64(index int) float64
	SetBool(index int, v bool)
	GetBool(index int) bool
	Items() []*ArgumentDataItem
	RangeItems(start, end int) []*ArgumentDataItem
	Item(index int) *ArgumentDataItem
	SetItems(items []*ArgumentDataItem)
	AddItems(items []*ArgumentDataItem)
}

type ICefProcessMessage struct {
	Name         string
	ArgumentList IArgumentList
}

type CefProcessMessagePtr struct {
	Name    uintptr
	Data    uintptr
	DataLen uintptr
}

// 进程之间消息
type argumentList struct {
	iData []*ArgumentDataItem
	len   int
	lock  *sync.Mutex
}

type ArgumentDataItem struct {
	vType int8
	//dataLen int32
	data []byte
}

type BuildBytes struct {
	buf *bytes.Buffer
}

func NewProcessMessage(name string) *ICefProcessMessage {
	return &ICefProcessMessage{
		Name:         name,
		ArgumentList: NewArgumentList(),
	}
}

func baseDataLength(gvt GO_VALUE_TYPE) int32 {
	switch gvt {
	case GO_VALUE_INT, GO_VALUE_UINT:
		return IntSize
	case GO_VALUE_INT8, GO_VALUE_UINT8, GO_VALUE_BOOL:
		return 1
	case GO_VALUE_INT16, GO_VALUE_UINT16:
		return 2
	case GO_VALUE_INT32, GO_VALUE_UINT32, GO_VALUE_FLOAT32:
		return 4
	case GO_VALUE_INT64, GO_VALUE_UINT64, GO_VALUE_FLOAT64:
		return 8
	}
	return -1
}

func (m *BuildBytes) Set(data interface{}) {
	if m.buf == nil {
		m.buf = &bytes.Buffer{}
	}
	binary.Write(m.buf, binary.BigEndian, data)
}

func (m *BuildBytes) Bytes() []byte {
	if m.buf != nil {
		return m.buf.Bytes()
	}
	return nil
}

// NewArgumentList
func NewArgumentList() IArgumentList {
	return &argumentList{lock: new(sync.Mutex)}
}

func (m *argumentList) Clear() {
	m.iData = nil
	m.len = 0
	m.lock = nil
}

func (m *argumentList) Package() []byte {
	buf := new(bytes.Buffer)
	defer buf.Reset()
	for _, data := range m.Items() {
		if data == nil {
			continue
		}
		binary.Write(buf, binary.BigEndian, data.vType)
		if data.vType == 0 {
			binary.Write(buf, binary.BigEndian, int32(len(data.data)))
		}
		binary.Write(buf, binary.BigEndian, data.data)
	}
	return buf.Bytes()
}

func (m *argumentList) UnPackageBytePtr(bytePtr uintptr, byteLength int32) {
	var (
		l    = int(byteLength)
		ret  []*ArgumentDataItem
		low  int
		high = 1
	)
	for i := 0; i < l; i++ {
		var vType int8
		var dataLen int32
		var data []byte
		low = i
		high = low + 1
		vType = *(*int8)(unsafe.Pointer(bytePtr + (uintptr(low))))
		if vType == 0 {
			low = high
			high = high + 4
			dataLen = BytesToInt32(CopyBytePtr(bytePtr, low, high))
		} else {
			dataLen = baseDataLength(GO_VALUE_TYPE(vType))
		}
		low = high
		high = high + int(dataLen)
		data = CopyBytePtr(bytePtr, low, high)

		ret = append(ret, &ArgumentDataItem{vType: vType, data: data})
		i = high - 1
	}
	m.SetItems(ret)
}

func (m *argumentList) UnPackage(dataByte []byte) {
	var (
		l    = len(dataByte)
		ret  []*ArgumentDataItem
		low  int
		high = 1
	)
	for i := 0; i < l; i++ {
		var vType int8
		var dataLen int32
		var data []byte
		low = i
		high = low + 1
		binary.Read(bytes.NewReader(dataByte[low:high]), binary.BigEndian, &vType)
		if vType == 0 {
			low = high
			high = high + 4
			binary.Read(bytes.NewReader(dataByte[low:high]), binary.BigEndian, &dataLen)
		} else {
			dataLen = baseDataLength(GO_VALUE_TYPE(vType))
		}
		low = high
		high = high + int(dataLen)
		data = dataByte[low:high]
		ret = append(ret, &ArgumentDataItem{vType: vType, data: data})
		i = high - 1
	}
	m.SetItems(ret)
}

func (m *argumentList) ToReflectValue() []reflect.Value {
	var ref = make([]reflect.Value, len(m.Items()))
	for i, item := range m.Items() {
		switch GO_VALUE_TYPE(item.vType) {
		case GO_VALUE_STRING:
			ref[i] = reflect.ValueOf(item.GetString())
		case GO_VALUE_INT:
			ref[i] = reflect.ValueOf(item.GetInt())
		case GO_VALUE_INT8:
			ref[i] = reflect.ValueOf(item.GetInt8())
		case GO_VALUE_INT16:
			ref[i] = reflect.ValueOf(item.GetInt16())
		case GO_VALUE_INT32:
			ref[i] = reflect.ValueOf(item.GetInt32())
		case GO_VALUE_INT64:
			ref[i] = reflect.ValueOf(item.GetInt64())
		case GO_VALUE_FLOAT32:
			ref[i] = reflect.ValueOf(item.GetFloat32())
		case GO_VALUE_FLOAT64:
			ref[i] = reflect.ValueOf(item.GetFloat64())
		case GO_VALUE_BOOL:
			ref[i] = reflect.ValueOf(item.GetBool())
		}
	}
	return ref
}

func (m *argumentList) ReflectValueConvert(ref []reflect.Value) {
	m.iData = make([]*ArgumentDataItem, len(ref))
	m.len = len(ref)
	for i, r := range ref {
		var data []byte
		var gov int8
		switch r.Kind() {
		case reflect.String:
			data = StringToBytes(r.Interface().(string))
			gov = int8(GO_VALUE_STRING)
		case reflect.Int:
			data = IntToBytes(r.Interface().(int))
			gov = int8(GO_VALUE_INT)
		case reflect.Int8:
			data = Int8ToBytes(r.Interface().(int8))
			gov = int8(GO_VALUE_INT8)
		case reflect.Int16:
			data = Int16ToBytes(r.Interface().(int16))
			gov = int8(GO_VALUE_INT16)
		case reflect.Int32:
			data = Int32ToBytes(r.Interface().(int32))
			gov = int8(GO_VALUE_INT32)
		case reflect.Int64:
			data = Int64ToBytes(r.Interface().(int64))
			gov = int8(GO_VALUE_INT64)
		case reflect.Float32:
			data = Float32ToBytes(r.Interface().(float32))
			gov = int8(GO_VALUE_FLOAT32)
		case reflect.Float64:
			data = Float64ToBytes(r.Interface().(float64))
			gov = int8(GO_VALUE_FLOAT64)
		case reflect.Bool:
			data = []byte{BoolToByte(r.Interface().(bool))}
			gov = int8(GO_VALUE_BOOL)
		default:
			data = StringToBytes("无效类型")
			gov = int8(GO_VALUE_INVALID_TYPE)
		}
		m.iData[i] = &ArgumentDataItem{
			vType: gov,
			data:  data,
		}
	}
}

func (m *argumentList) capacity(index int) {
	if len(m.iData) == 0 {
		m.iData = make([]*ArgumentDataItem, 8)
	} else if index >= m.len {
		var tmp = m.iData
		var c = m.len + 8
		if index > c {
			c = index + c
		}
		m.iData = make([]*ArgumentDataItem, c)
		copy(m.iData, tmp)
	}
	if m.iData[index] == nil {
		m.len++
	}
}

func (m *argumentList) SetArguments(index int, argument IArgumentList) {
	if argument == nil {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = argument.Package()
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_ARGUMENT)
		//d.dataLen = int32(len(vb))
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_ARGUMENT),
			//dataLen: int32(len(vb)),
			data: vb,
		}
	}
}

func (m *argumentList) SetString(index int, v string, isDStr ...bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = StringToBytes(v, isDStr...)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_STRING)
		//d.dataLen = int32(len(vb))
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_STRING),
			//dataLen: int32(len(vb)),
			data: vb,
		}
	}
}

func (m *argumentList) GetString(index int) string {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetString()
	}
	return ""
}

func (m *argumentList) SetIntAuto(index int, v interface{}, gov GO_VALUE_TYPE) {
	switch gov {
	case GO_VALUE_INT:
		m.SetInt(index, m.numberToInt(v))
	case GO_VALUE_INT8:
		m.SetInt8(index, int8(m.numberToInt(v)))
	case GO_VALUE_INT16:
		m.SetInt16(index, int16(m.numberToInt(v)))
	case GO_VALUE_INT32:
		m.SetInt32(index, int32(m.numberToInt(v)))
	case GO_VALUE_INT64:
		m.SetInt64(index, int64(m.numberToInt(v)))
	}
}

func (m *argumentList) numberToInt(v interface{}) int {
	switch v.(type) {
	case int:
		return v.(int)
	case int8:
		return int(v.(int8))
	case int16:
		return int(v.(int16))
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case uint:
		return int(v.(uint))
	case uint8:
		return int(v.(uint8))
	case uint16:
		return int(v.(uint16))
	case uint32:
		return int(v.(uint32))
	case uint64:
		return int(v.(uint64))
	}
	return 0
}

func (m *argumentList) SetInt(index int, v int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = IntToBytes(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_INT)
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_INT),
			data:  vb,
		}
	}
}

func (m *argumentList) SetInt8(index int, v int8) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = Int8ToBytes(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_INT8)
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_INT8),
			data:  vb,
		}
	}
}

func (m *argumentList) SetData(index int, vType int8, data []byte) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	if d := m.iData[index]; d != nil {
		d.vType = vType
		d.data = data
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: vType,
			data:  data,
		}
	}
}

func (m *argumentList) SetInt16(index int, v int16) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = Int16ToBytes(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_INT16)
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_INT16),
			data:  vb,
		}
	}
}

func (m *argumentList) SetInt32(index int, v int32) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = Int32ToBytes(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_INT32)
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_INT32),
			data:  vb,
		}
	}
}

func (m *argumentList) SetInt64(index int, v int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = Int64ToBytes(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_INT64)
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_INT64),
			data:  vb,
		}
	}
}

func (m *argumentList) GetInt(index int) int {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetInt()
	}
	return 0
}
func (m *argumentList) GetInt8(index int) int8 {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetInt8()
	}
	return 0
}
func (m *argumentList) GetInt16(index int) int16 {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetInt16()
	}
	return 0
}
func (m *argumentList) GetInt32(index int) int32 {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetInt32()
	}
	return 0
}

func (m *argumentList) GetBytes(index int) []byte {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.Data()
	}
	return nil
}

func (m *argumentList) GetData(index int) *ArgumentDataItem {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di
	}
	return nil
}

func (m *argumentList) GetInt64(index int) int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetInt64()
	}
	return 0
}
func (m *argumentList) SetFloatAuto(index int, v interface{}, gov GO_VALUE_TYPE) {
	switch gov {
	case GO_VALUE_FLOAT32:
		m.SetFloat32(index, v.(float32))
	case GO_VALUE_FLOAT64:
		m.SetFloat64(index, v.(float64))
	}
}

func (m *argumentList) SetFloat32(index int, v float32) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = Float32ToBytes(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_FLOAT32)
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_FLOAT32),
			data:  vb,
		}
	}
}
func (m *argumentList) SetFloat64(index int, v float64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = Float64ToBytes(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_FLOAT64)
		//d.dataLen = int32(len(vb))
		d.data = vb
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_FLOAT64),
			//dataLen: int32(len(vb)),
			data: vb,
		}
	}
}
func (m *argumentList) GetFloat32(index int) float32 {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetFloat32()
	}
	return 0
}

func (m *argumentList) GetFloat64(index int) float64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetFloat64()
	}
	return 0
}

func (m *argumentList) SetBool(index int, v bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.capacity(index)
	var vb = BoolToByte(v)
	if d := m.iData[index]; d != nil {
		d.vType = int8(GO_VALUE_BOOL)
		d.data = []byte{vb}
	} else {
		m.iData[index] = &ArgumentDataItem{
			vType: int8(GO_VALUE_BOOL),
			data:  []byte{vb},
		}
	}
}

func (m *argumentList) GetBool(index int) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		di := m.iData[index]
		return di.GetBool()
	}
	return false
}

func (m *argumentList) Size() int {
	return m.len
}

func (m *argumentList) Items() []*ArgumentDataItem {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.len < len(m.iData) {
		m.iData = m.iData[:m.len]
	}
	return m.iData
}

func (m *argumentList) RangeItems(start, end int) []*ArgumentDataItem {
	m.lock.Lock()
	defer m.lock.Unlock()
	if start < end && end < len(m.iData) {
		return m.iData[start:end]
	}
	return nil
}

func (m *argumentList) Item(index int) *ArgumentDataItem {
	m.lock.Lock()
	defer m.lock.Unlock()
	if index < len(m.iData) {
		return m.iData[index]
	}
	return nil
}

func (m *argumentList) SetItems(items []*ArgumentDataItem) {
	m.iData = items
	m.len = len(items)
}

func (m *argumentList) AddItems(items []*ArgumentDataItem) {
	m.iData = m.Items()
	m.iData = append(m.iData, items...)
	m.len = len(items) + m.len
}

func (m *ArgumentDataItem) IsString() bool {
	return m.vType == int8(GO_VALUE_STRING)
}

func (m *ArgumentDataItem) IsIntAuto() bool {
	switch GO_VALUE_TYPE(m.vType) {
	case GO_VALUE_INT, GO_VALUE_INT8, GO_VALUE_INT16, GO_VALUE_INT32, GO_VALUE_INT64, GO_VALUE_UINT, GO_VALUE_UINT8, GO_VALUE_UINT16, GO_VALUE_UINT32, GO_VALUE_UINT64:
		return true
	}
	return false
}
func (m *ArgumentDataItem) IsInt() bool {
	return m.vType == int8(GO_VALUE_INT)
}
func (m *ArgumentDataItem) IsInt8() bool {
	return m.vType == int8(GO_VALUE_INT8)
}
func (m *ArgumentDataItem) IsInt16() bool {
	return m.vType == int8(GO_VALUE_INT16)
}
func (m *ArgumentDataItem) IsInt32() bool {
	return m.vType == int8(GO_VALUE_INT32)
}
func (m *ArgumentDataItem) IsInt64() bool {
	return m.vType == int8(GO_VALUE_INT64)
}

func (m *ArgumentDataItem) IsFloatAuto() bool {
	switch GO_VALUE_TYPE(m.vType) {
	case GO_VALUE_FLOAT32, GO_VALUE_FLOAT64:
		return true
	}
	return false
}
func (m *ArgumentDataItem) IsFloat32() bool {
	return m.vType == int8(GO_VALUE_FLOAT32)
}
func (m *ArgumentDataItem) IsFloat64() bool {
	return m.vType == int8(GO_VALUE_FLOAT64)
}

func (m *ArgumentDataItem) IsBoolean() bool {
	return m.vType == int8(GO_VALUE_BOOL)
}

func (m *ArgumentDataItem) Data() []byte {
	return m.data
}

func (m *ArgumentDataItem) VType() int8 {
	return m.vType
}

func (m *ArgumentDataItem) VTypeToJS() V8_JS_VALUE_TYPE {
	switch GO_VALUE_TYPE(m.vType) {
	case GO_VALUE_STRING:
		return V8_VALUE_STRING
	case GO_VALUE_INT, GO_VALUE_INT8, GO_VALUE_INT16, GO_VALUE_INT32, GO_VALUE_INT64, GO_VALUE_UINT, GO_VALUE_UINT8, GO_VALUE_UINT16, GO_VALUE_UINT32, GO_VALUE_UINT64:
		return V8_VALUE_INT
	case GO_VALUE_FLOAT32, GO_VALUE_FLOAT64:
		return V8_VALUE_DOUBLE
	case GO_VALUE_BOOL:
		return V8_VALUE_BOOLEAN
	}
	return V8_NO_OUT_VALUE
}

func (m *ArgumentDataItem) GetString() string {
	if m.IsString() {
		return BytesToString(m.data)
	}
	return ""
}
func (m *ArgumentDataItem) GetNumber() int32 {
	if m.IsInt() {
		return int32(m.GetInt())
	} else if m.IsInt8() {
		return int32(m.GetInt8())
	} else if m.IsInt16() {
		return int32(m.GetInt16())
	} else if m.IsInt32() {
		return m.GetInt32()
	} else if m.IsInt64() {
		return int32(m.GetInt64())
	}
	return 0
}

func (m *ArgumentDataItem) GetInt() int {
	if m.IsInt() {
		return BytesToInt(m.data)
	}
	return 0
}

func (m *ArgumentDataItem) GetInt8() int8 {
	if m.IsInt8() {
		return ByteToInt8(m.data[0])
	}
	return 0
}

func (m *ArgumentDataItem) GetInt16() int16 {
	if m.IsInt16() {
		return BytesToInt16(m.data)
	}
	return 0
}

func (m *ArgumentDataItem) GetInt32() int32 {
	if m.IsInt32() {
		return BytesToInt32(m.data)
	}
	return 0
}

func (m *ArgumentDataItem) GetInt64() int64 {
	if m.IsInt64() {
		return BytesToInt64(m.data)
	}
	return 0
}

func (m *ArgumentDataItem) GetDouble() float64 {
	if m.IsFloat32() {
		return float64(m.GetFloat32())
	} else if m.IsFloat64() {
		return m.GetFloat64()
	}
	return 0
}

func (m *ArgumentDataItem) GetFloat32() float32 {
	if m.IsFloat32() {
		return BytesToFloat32(m.data)
	}
	return 0
}

func (m *ArgumentDataItem) GetFloat64() float64 {
	if m.IsFloat64() {
		return BytesToFloat64(m.data)
	}
	return 0
}

func (m *ArgumentDataItem) GetBool() bool {
	if m.IsBoolean() {
		return ByteToBool(m.data[0])
	}
	return false
}
