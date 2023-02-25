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
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/types"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// BinaryValueRef -> ICefBinaryValue
var BinaryValueRef cefBinaryValue

//cefBinaryValue
type cefBinaryValue uintptr

// New 创建一个带有数据的二进对象
func (*cefBinaryValue) New(data []byte) *ICefBinaryValue {
	var result uintptr
	imports.Proc(internale_CefBinaryValueRef_New).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

// Create 创建一个二进对象
func (*cefBinaryValue) Create() *ICefBinaryValue {
	var result uintptr
	imports.Proc(internale_CefBinaryValueRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefBinaryValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefBinaryValue) IsValid() bool {
	r1, _, _ := imports.Proc(internale_CefBinaryValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefBinaryValue) IsOwned() bool {
	r1, _, _ := imports.Proc(internale_CefBinaryValue_IsOwned).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefBinaryValue) Copy() *ICefBinaryValue {
	var result uintptr
	imports.Proc(internale_CefBinaryValue_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefBinaryValue) GetSize() uint32 {
	r1, _, _ := imports.Proc(internale_CefBinaryValue_GetSize).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefBinaryValue) GetData(buffer []byte, dataOffset types.NativeUInt) uint32 {
	r1, _, _ := imports.Proc(internale_CefBinaryValue_GetData).Call(m.Instance(), uintptr(unsafe.Pointer(&buffer[0])), uintptr(uint32(len(buffer))), dataOffset.ToPtr())
	return uint32(r1)
}
