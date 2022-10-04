//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"bytes"
	"encoding/binary"
)

type ICefDictionaryValue struct {
	iData    []*dictDataItem
	dataLen  int
	dataByte []byte
}

type dictDataItem struct {
	vType   int8
	keyLen  int32
	key     []byte
	dataLen int32
	data    []byte
}

func NewCefDictionaryValue() *ICefDictionaryValue {
	return &ICefDictionaryValue{}
}

func (m *ICefDictionaryValue) Clear() {
	m.iData = []*dictDataItem{}
	m.dataLen = 0
	m.dataByte = []byte{}
}

func (m *ICefDictionaryValue) Package() []byte {
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

func (m *ICefDictionaryValue) Items() []*dictDataItem {
	return m.iData
}

func (m *ICefDictionaryValue) SetDictionary(name string, dict *ICefDictionaryValue) {
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

func (m *ICefDictionaryValue) SetString(name string, v string) {
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

func (m *ICefDictionaryValue) SetInt(name string, v int32) {
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

func (m *ICefDictionaryValue) SetDouble(name string, v float64) {
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

func (m *ICefDictionaryValue) SetBoolean(name string, v bool) {
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
