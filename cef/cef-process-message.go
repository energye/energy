package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCefProcessMessageRef -> ICefProcessMessage
type TCefProcessMessageRef uintptr

// ProcessMessageRef -> ICefProcessMessage
var ProcessMessageRef TCefProcessMessageRef

func (*TCefProcessMessageRef) New(name string) *ICefProcessMessage {
	var result uintptr
	imports.Proc(internale_CefProcessMessageRef_New).Call(api.PascalStr(name), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{
		instance: unsafe.Pointer(result),
	}
}
func (m *ICefProcessMessage) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefProcessMessage) ArgumentList() *ICefListValue {
	var result uintptr
	imports.Proc(internale_CefProcessMessage_ArgumentList).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefProcessMessage) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefProcessMessage_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefProcessMessage) Copy() *ICefProcessMessage {
	var result uintptr
	imports.Proc(internale_CefProcessMessage_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefProcessMessage) Name() string {
	r1, _, _ := imports.Proc(internale_CefProcessMessage_Name).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefProcessMessage) Free() {
	imports.Proc(internale_CefProcessMessage_Free).Call(m.Instance())
	m.instance = nil
}
