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

func (*cefBinaryValue) New(data []byte) *ICefBinaryValue {
	var result uintptr
	imports.Proc(internale_CefBinaryValueRef_New).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefBinaryValue{
		instance: unsafe.Pointer(result),
	}
}

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
