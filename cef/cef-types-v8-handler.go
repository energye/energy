package cef

import (
	"fmt"
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type V8HandlerExecute func(name string, object ICefV8Value, arguments []*ICefV8Value, retVal ICefV8Value, Exception string) bool
type V8HandlerDestroy func()

func CreateCefV8Handler() *ICefV8Handler {
	var result uintptr
	imports.Proc(internale_CefV8Handler_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Handler{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefV8Handler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8Handler) Execute(fn V8HandlerExecute) {
	imports.Proc(internale_CefV8Handler_Execute).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefV8Handler) Destroy(fn V8HandlerDestroy) {
	imports.Proc(internale_CefV8Handler_Destroy).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8HandlerExecute:
			fmt.Println("----V8HandlerExecute", getPtr)
		case V8HandlerDestroy:
			fmt.Println("----V8HandlerDestroy", getPtr)
		default:
			return false
		}
		return true
	})
}
