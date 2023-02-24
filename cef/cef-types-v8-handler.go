package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type V8HandlerExecute func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool

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

func (m *ICefV8Handler) Destroy() {
	imports.Proc(internale_CefV8Handler_Destroy).Call(m.Instance())
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8HandlerExecute:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			argumentsPtr := getVal(2)
			argumentsLength := int32(getVal(3))
			arguments := &TCefV8ValueArray{arguments: argumentsPtr, argumentsLength: int(argumentsLength)}
			retValPtr := (*uintptr)(getPtr(4))
			retVal := &ResultV8Value{}
			exceptionPtr := (*uintptr)(getPtr(5))
			exception := &Exception{}
			resultPtr := (*bool)(getPtr(6))
			result := fn.(V8HandlerExecute)(name, object, arguments, retVal, exception)
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
		default:
			return false
		}
		return true
	})
}
