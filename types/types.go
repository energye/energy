package types

import (
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Int int
type UInt8 uint8
type UInt16 uint16
type UInt32 uint32
type UInt64 uint64
type UInt uint
type UIntptr uintptr
type String string
type Boolean bool
type Float32 float32
type Float64 float64
type Single = Float32
type PChar = String
type TCefColor = UInt16
type Integer = Int32
type LongInt = Int32
type NativeUInt = UInt32
type TCefString = String
type Cardinal = UInt32
type LongBool = Boolean

type TString struct {
	value string
}

func (m *TString) SetValue(v string) {
	m.value = v
}

func (m *TString) GetValue() string {
	return m.value
}

func (m *TString) ToPtr() uintptr {
	return api.PascalStr(m.value)
}

func (m Int8) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int16) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int32) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int64) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt8) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt16) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt32) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt64) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt) ToPtr() uintptr {
	return uintptr(m)
}

func (m UIntptr) ToPtr() uintptr {
	return uintptr(m)
}

func (m String) ToPtr() uintptr {
	return api.PascalStr(string(m))
}

func (m Boolean) ToPtr() uintptr {
	return api.PascalBool(bool(m))
}

func (m Float32) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}

func (m Float64) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}
