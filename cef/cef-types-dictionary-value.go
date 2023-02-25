package cef

import (
	"unsafe"
)

// DictionaryValueRef -> ICefDictionaryValue
var DictionaryValueRef cefDictionaryValue

//cefDictionaryValue
type cefDictionaryValue uintptr

func (*cefDictionaryValue) New() *ICefDictionaryValue {
	var result uintptr
	//imports.Proc(internale_CefBinaryValueRef_New).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefDictionaryValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}
