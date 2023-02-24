package cef

import (
	"github.com/energye/energy/common/imports"
	"unsafe"
)

func NewCefValue() *ICefValue {
	var result uintptr
	imports.Proc(internale_CefValueRef_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefValue{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefValue) IsValid() {

}

func (m *ICefValue) IsOwned() {

}

func (m *ICefValue) IsReadOnly() {

}

func (m *ICefValue) Copy() {

}

func (m *ICefValue) GetType() {

}

func (m *ICefValue) GetBool() {

}

func (m *ICefValue) GetInt() {

}

func (m *ICefValue) GetDouble() {

}

func (m *ICefValue) GetString() {

}

func (m *ICefValue) GetBinary() {

}

func (m *ICefValue) GetDictionary() {

}

func (m *ICefValue) GetList() {

}

func (m *ICefValue) SetNull() {

}

func (m *ICefValue) SetBool() {

}

func (m *ICefValue) SetInt() {

}

func (m *ICefValue) SetDouble() {

}

func (m *ICefValue) SetString() {

}

func (m *ICefValue) SetBinary() {

}

func (m *ICefValue) SetDictionary() {

}

func (m *ICefValue) SetList() {

}
