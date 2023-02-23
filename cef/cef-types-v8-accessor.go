package cef

import (
	"fmt"
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type V8AccessorGet func(name string, object, retVal ICefV8Value, exception string) bool
type V8AccessorSet func(name string, object, value ICefV8Value, exception string) bool

func CreateCefV8Accessor() *ICefV8Accessor {
	var result uintptr
	imports.Proc(internale_CefV8Accessor_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Accessor{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefV8Accessor) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8Accessor) Get(fn V8AccessorGet) {
	imports.Proc(internale_CefV8Accessor_Get).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefV8Accessor) Set(fn V8AccessorSet) {
	imports.Proc(internale_CefV8Accessor_Set).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8AccessorGet:
			fmt.Println("----V8AccessorGet", getPtr)
		case V8AccessorSet:
			fmt.Println("----V8AccessorSet", getPtr)
		default:
			return false
		}
		return true
	})
}
