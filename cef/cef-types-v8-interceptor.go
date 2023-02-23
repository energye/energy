package cef

import (
	"fmt"
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type V8InterceptorGetByName func(name string, object, retVal *ICefV8Value, exception string)
type V8InterceptorGetByIndex func(index int32, object, retVal *ICefV8Value, exception string)
type V8InterceptorSetByName func(name string, object, value, retVal *ICefV8Value, exception string)
type V8InterceptorSetByIndex func(index int32, object, value, retVal *ICefV8Value, exception string)
type V8InterceptorDestroy func()

func CreateCefV8InterceptorRef() *ICefV8Interceptor {
	var result uintptr
	imports.Proc(internale_CefV8InterceptorRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Interceptor{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefV8Interceptor) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8Interceptor) GetByName(fn V8InterceptorGetByName) {
	imports.Proc(internale_CefV8InterceptorRef_GetByName).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefV8Interceptor) GetByIndex(fn V8InterceptorGetByIndex) {
	imports.Proc(internale_CefV8InterceptorRef_GetByIndex).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefV8Interceptor) SetByName(fn V8InterceptorSetByName) {
	imports.Proc(internale_CefV8InterceptorRef_SetByName).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefV8Interceptor) SetByIndex(fn V8InterceptorSetByIndex) {
	imports.Proc(internale_CefV8InterceptorRef_SetByIndex).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefV8Interceptor) Destroy(fn V8InterceptorDestroy) {
	imports.Proc(internale_CefV8InterceptorRef_Destroy).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8InterceptorGetByName:
			fmt.Println("----V8InterceptorGetByName", getPtr)
		case V8InterceptorGetByIndex:
			fmt.Println("----V8InterceptorGetByIndex", getPtr)
		case V8InterceptorSetByName:
			fmt.Println("----V8InterceptorSetByName", getPtr)
		case V8InterceptorSetByIndex:
			fmt.Println("----V8InterceptorSetByIndex", getPtr)
		case V8InterceptorDestroy:
			fmt.Println("----V8InterceptorDestroy", getPtr)
		default:
			return false
		}
		return true
	})
}
