//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Chromium 字典
package cef

import (
	"bytes"
	"encoding/binary"
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
)

// DictionaryValue 字典数据结构
type DictionaryValue struct {
	iData    []*dictDataItem
	dataLen  int
	dataByte []byte
}

// dictDataItem 字典数据项
type dictDataItem struct {
	vType   int8
	keyLen  int32
	key     []byte
	dataLen int32
	data    []byte
}

// NewCefDictionaryValue 创建一个字典
func NewCefDictionaryValue() *DictionaryValue {
	return &DictionaryValue{}
}

// Clear 清空字典内容
func (m *DictionaryValue) Clear() {
	m.iData = []*dictDataItem{}
	m.dataLen = 0
	m.dataByte = []byte{}
}

// Package 将字典数据打包为字节
func (m *DictionaryValue) Package() []byte {
	buf := &bytes.Buffer{}
	for _, data := range m.Items() {
		binary.Write(buf, binary.BigEndian, data.vType)
		binary.Write(buf, binary.BigEndian, data.keyLen)
		binary.Write(buf, binary.BigEndian, data.key)
		binary.Write(buf, binary.BigEndian, data.dataLen)
		binary.Write(buf, binary.BigEndian, data.data)
	}
	return buf.Bytes()
}

// Items 返回所有字典项
func (m *DictionaryValue) Items() []*dictDataItem {
	return m.iData
}

// SetDictionary 设置 DictionaryValue 类型值
func (m *DictionaryValue) SetDictionary(name string, dict *DictionaryValue) {
	if name == "" || dict == nil {
		return
	}
	var nameByte = StringToBytes(name)
	var dataByte = dict.Package()
	ddi := &dictDataItem{
		vType:   int8(GO_VALUE_DICTVALUE),
		keyLen:  int32(len(nameByte)),
		key:     nameByte,
		dataLen: int32(len(dataByte)),
		data:    dataByte,
	}
	m.iData = append(m.iData, ddi)
	m.dataLen = len(m.iData)
}

// SetString 设置 string 类型值
func (m *DictionaryValue) SetString(name string, v string) {
	if name == "" || v == "" {
		return
	}
	var nameByte = StringToBytes(name)
	var dataByte = StringToBytes(v)
	ddi := &dictDataItem{
		vType:   int8(GO_VALUE_STRING),
		keyLen:  int32(len(nameByte)),
		key:     nameByte,
		dataLen: int32(len(dataByte)),
		data:    dataByte,
	}
	m.iData = append(m.iData, ddi)
	m.dataLen = len(m.iData)
}

// SetInt 设置 int32 类型值
func (m *DictionaryValue) SetInt(name string, v int32) {
	if name == "" || v < 0 {
		return
	}
	var nameByte = StringToBytes(name)
	var dataByte = Int32ToBytes(v)
	ddi := &dictDataItem{
		vType:   int8(GO_VALUE_INT32),
		keyLen:  int32(len(nameByte)),
		key:     nameByte,
		dataLen: int32(len(dataByte)),
		data:    dataByte,
	}
	m.iData = append(m.iData, ddi)
	m.dataLen = len(m.iData)
}

// SetDouble 设置 float64 类型值
func (m *DictionaryValue) SetDouble(name string, v float64) {
	if name == "" || v < 0 {
		return
	}
	var nameByte = StringToBytes(name)
	var dataByte = Float64ToBytes(v)
	ddi := &dictDataItem{
		vType:   int8(GO_VALUE_FLOAT64),
		keyLen:  int32(len(nameByte)),
		key:     nameByte,
		dataLen: int32(len(dataByte)),
		data:    dataByte,
	}
	m.iData = append(m.iData, ddi)
	m.dataLen = len(m.iData)
}

// SetBoolean 设置 bool 类型值
func (m *DictionaryValue) SetBoolean(name string, v bool) {
	if name == "" {
		return
	}
	var nameByte = StringToBytes(name)
	var dataByte = []byte{BoolToByte(v)}
	ddi := &dictDataItem{
		vType:   int8(GO_VALUE_BOOL),
		keyLen:  int32(len(nameByte)),
		key:     nameByte,
		dataLen: int32(len(dataByte)),
		data:    dataByte,
	}
	m.iData = append(m.iData, ddi)
	m.dataLen = len(m.iData)
}
