//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 二进制类型 BinaryValueRef.New() or Create()
//
// ICefBinaryValue
// 示例：
//  Go: 创建并绑定 ICefBinaryValue 对象名称: myobj
//  JavaScript: let obj = new Uint8Array(window.myobj);
//  		    console.log(obj)

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefBinaryValue -> ArgumentList
type ICefBinaryValue struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// BinaryValueRef -> ICefBinaryValue
var BinaryValueRef cefBinaryValue

// cefBinaryValue
type cefBinaryValue uintptr

// New 创建一个带有数据的二进对象
func (*cefBinaryValue) New(data []byte) *ICefBinaryValue {
	var result uintptr
	imports.Proc(def.CefBinaryValueRef_New).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

// Create 创建一个二进对象
func (*cefBinaryValue) Create() *ICefBinaryValue {
	var result uintptr
	imports.Proc(def.CefBinaryValueRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

// UnWrap
func (*cefBinaryValue) UnWrap(data *ICefBinaryValue) *ICefBinaryValue {
	var result uintptr
	imports.Proc(def.CefBinaryValueRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

// Instance 实例
func (m *ICefBinaryValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefBinaryValue) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefBinaryValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefBinaryValue) IsOwned() bool {
	r1, _, _ := imports.Proc(def.CefBinaryValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefBinaryValue) Copy() *ICefBinaryValue {
	var result uintptr
	imports.Proc(def.CefBinaryValue_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefBinaryValue) GetSize() uint32 {
	r1, _, _ := imports.Proc(def.CefBinaryValue_GetSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefBinaryValue) GetData(buffer []byte, dataOffset uint32) uint32 {
	r1, _, _ := imports.Proc(def.CefBinaryValue_GetData).Call(m.Instance(), uintptr(unsafe.Pointer(&buffer[0])), uintptr(uint32(len(buffer))), uintptr(dataOffset))
	return uint32(r1)
}

func (m *ICefBinaryValue) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// TCefBinaryValueArray
//
//	[]ICefBinaryValue
type TCefBinaryValueArray struct {
	instance     unsafe.Pointer
	binaryValues []*ICefBinaryValue
	count        uint32
}

func (m *TCefBinaryValueArray) Get(index uint32) *ICefBinaryValue {
	if index < m.count {
		if m.binaryValues == nil {
			m.binaryValues = make([]*ICefBinaryValue, m.count, m.count)
		}
		var result uintptr
		imports.Proc(def.CefBinaryValueArray_Get).Call(uintptr(m.instance), uintptr(index), uintptr(unsafe.Pointer(&result)))
		if result != 0 {
			m.binaryValues[index] = &ICefBinaryValue{instance: unsafe.Pointer(result)}
			return m.binaryValues[index]
		}
	}
	return nil
}

func (m *TCefBinaryValueArray) Count() uint32 {
	return m.count
}

func (m *TCefBinaryValueArray) Free() {
	if m.instance != nil {
		if m.binaryValues != nil {
			for _, cert := range m.binaryValues {
				cert.Free()
			}
			m.binaryValues = nil
		}
		m.instance = nil
		m.count = 0
	}
}
