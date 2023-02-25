package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type V8AccessorGet func(name string, object *ICefV8Value, retVal *ResultV8Value, exception *Exception) bool
type V8AccessorSet func(name string, object *ICefV8Value, value *ICefV8Value, exception *Exception) bool

//V8AccessorRef -> ICefV8Accessor
var V8AccessorRef cefV8Accessor

// cefV8Accessor
type cefV8Accessor uintptr

func (*cefV8Accessor) New() *ICefV8Accessor {
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

func (m *ICefV8Accessor) Destroy() {
	imports.Proc(internale_CefV8Accessor_Destroy).Call(m.Instance())
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8AccessorGet:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			retValPtr := (*uintptr)(getPtr(2))
			retVal := &ResultV8Value{}
			exceptionPtr := (*uintptr)(getPtr(3))
			exception := &Exception{}
			resultPtr := (*bool)(getPtr(4))
			result := fn.(V8AccessorGet)(name, object, retVal, exception)
			if retVal.v8value != nil {
				*retValPtr = retVal.v8value.Instance()
			}
			if exception.Message() != "" {
				*exceptionPtr = api.PascalStr(exception.Message())
			} else {
				*exceptionPtr = 0
			}
			*resultPtr = result
			object.Free()
		case V8AccessorSet:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			value := &ICefV8Value{instance: getPtr(2)}
			exceptionPtr := (*uintptr)(getPtr(3))
			exception := &Exception{}
			resultPtr := (*bool)(getPtr(4))
			result := fn.(V8AccessorSet)(name, object, value, exception)
			if exception.Message() != "" {
				*exceptionPtr = api.PascalStr(exception.Message())
			} else {
				*exceptionPtr = 0
			}
			*resultPtr = result
			object.Free()
			value.Free()
		default:
			return false
		}
		return true
	})
}
