//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy 变量相关
package cef

import (
	"errors"
	"fmt"
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"reflect"
	"sync"
	"unsafe"
)

var (
	nullptr                        unsafe.Pointer     = nil      //
	commonRootName                                    = "gocobj" //V8Value 通用类型变量属性的所属默认对象名称
	objectRootName                                    = "goobj"  //V8Value 对象类型变量属性的所属默认对象名称
	DownloadsDir                   string                        //下载目录
	enableGPU                      = false                       //启用GPU true启用 false不启用
	processName                    common.PRCESS_TYPE            //进程名称
	isMainProcess, isRenderProcess bool                          //
)

type initBindVariableCallback func(browser *ICefBrowser, frame *ICefFrame, bind IProvisionalBindStorage)

// 变量绑定
var VariableBind = &variableBind{valuesBind: make(map[string]JSValue)}

type IProvisionalBindStorage interface {
	NewString(name, value string) *JSString
	NewInteger(name string, value int32) *JSInteger
	NewDouble(name string, value float64) *JSDouble
	NewBoolean(name string, value bool) *JSBoolean
	NewNull(name string) *JSNull
	NewUndefined(name string) *JSUndefined
	NewFunction(name string, fn interface{}) error
	NewObjects(objects ...interface{})
	Bind(name string, bind interface{}) error
}

type variableBind struct {
	initBindVariableCallback initBindVariableCallback //
	valuesBind               map[string]JSValue       //所有绑定变量属性或函数集合
}

func init() {
	isMainProcess = common.Args.IsMain()
	isRenderProcess = common.Args.IsRender()
}

func (m *variableBind) putValueBind(fullName string, value JSValue) {
	if _, ok := m.valuesBind[fullName]; !ok {
		m.valuesBind[fullName] = value
	}
}

func (m *variableBind) ValueBindCount() int {
	return len(m.valuesBind)
}

func (m *variableBind) GetValueBind(fullName string) (JSValue, bool) {
	value, ok := m.valuesBind[fullName]
	return value, ok
}

func clearValueBind() {
	//valuesBind = make(map[uintptr]JSValue)
}

// VariableCreateCallback
//
// # Go 和 javaScript 的函数或变量绑定声明初始函数
//
// 在 javaScript 中调用Go中的（函数,变量）需要在此回调函数中绑定
//
// 主进程和子进程
func (m *variableBind) VariableCreateCallback(callback func(browser *ICefBrowser, frame *ICefFrame, bind IProvisionalBindStorage)) {
	m.initBindVariableCallback = callback
}

// 调用变量绑定回调函数
//
// 在主进程和渲染进程创建时调用
func (m *variableBind) callVariableBind(browser *ICefBrowser, frame *ICefFrame) {
	if m.initBindVariableCallback != nil {
		m.initBindVariableCallback(browser, frame, m)
	}
}

// NewString V8Value
func (m *variableBind) NewString(name, value string) *JSString {
	jsValueBind := new(JSString)
	jsValueBind.valueType = new(VT)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_STRING
	jsValueBind.valueType.Gov = GO_VALUE_STRING
	m.valueHandle(jsValueBind)
	return jsValueBind
}

// NewInteger V8Value
func (m *variableBind) NewInteger(name string, value int32) *JSInteger {
	jsValueBind := new(JSInteger)
	jsValueBind.valueType = new(VT)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_INT
	jsValueBind.valueType.Gov = GO_VALUE_INT32
	m.valueHandle(jsValueBind)
	return jsValueBind
}

// NewDouble V8Value
func (m *variableBind) NewDouble(name string, value float64) *JSDouble {
	jsValueBind := new(JSDouble)
	jsValueBind.valueType = new(VT)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_DOUBLE
	jsValueBind.valueType.Gov = GO_VALUE_FLOAT64
	m.valueHandle(jsValueBind)
	return jsValueBind
}

// NewBool V8Value
func (m *variableBind) NewBoolean(name string, value bool) *JSBoolean {
	jsValueBind := new(JSBoolean)
	jsValueBind.valueType = new(VT)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_BOOLEAN
	jsValueBind.valueType.Gov = GO_VALUE_BOOL
	m.valueHandle(jsValueBind)
	return jsValueBind
}

// NewNull V8Value
func (m *variableBind) NewNull(name string) *JSNull {
	jsValueBind := new(JSNull)
	jsValueBind.valueType = new(VT)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = "null"
	jsValueBind.valueType.Jsv = V8_VALUE_NULL
	jsValueBind.valueType.Gov = GO_VALUE_NIL
	m.valueHandle(jsValueBind)
	return jsValueBind
}

// NewUndefined V8Value
func (m *variableBind) NewUndefined(name string) *JSUndefined {
	jsValueBind := new(JSUndefined)
	jsValueBind.valueType = new(VT)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = "undefined"
	jsValueBind.valueType.Jsv = V8_VALUE_UNDEFINED
	jsValueBind.valueType.Gov = GO_VALUE_STRING
	m.valueHandle(jsValueBind)
	return jsValueBind
}

// NewFunction V8Value
//
// 绑定 定义的普通函数 prefix default struct name
func (m *variableBind) NewFunction(name string, fn interface{}) error {
	if common.GOValueReflectType(fn) == GO_VALUE_FUNC {
		if info, err := checkFunc(reflect.TypeOf(fn), FN_TYPE_COMMON); err == nil {
			jsValueBind := new(JSFunction)
			jsValueBind.valueType = new(VT)
			jsValueBind.rwLock = new(sync.Mutex)
			jsValueBind.name = name
			jsValueBind.value = fn
			jsValueBind.valueType.Jsv = V8_VALUE_FUNCTION
			jsValueBind.valueType.Gov = GO_VALUE_FUNC
			jsValueBind.funcInfo = info
			m.valueHandle(jsValueBind)
			return nil
		} else {
			return err
		}
	}
	return errors.New("创建的函数不是函数类型")
}

// V8Value valueHandle
func (m *variableBind) valueHandle(jsValue JSValue) {
	jsValue.setPtr(unsafe.Pointer(&jsValue))
	jsValue.setInstance(uintptr(jsValue.Ptr()))
	jsValue.setEventId(uintptr(__bind_id()))
	jsValue.setThat(jsValue)
	m.putValueBind(jsValue.Name(), jsValue)
}

// NewObjects V8Value
//
// 对象类型变量和对象函数绑定
func (m *variableBind) NewObjects(objects ...interface{}) {
	bindObject(objects...)
}

// Bind V8Value
//
// 变量和函数绑定, 在Go中定义字段绑定到JS字段中, 在Go中定义的函数导出到JS
//
// 支持类型 String = string , Integer = int32 , Double = float64, Boolean = bool, Function = func, ObjectInfos = struct | map,  Array = Slice
//
// 主进程和子进程
//
// 返回nil表示绑定成功
func (m *variableBind) Bind(name string, bind interface{}) error {
	if isMainProcess || isRenderProcess {
		typ := reflect.TypeOf(bind)
		kind := typ.Kind()
		//直接绑定变量地址
		if kind != reflect.Ptr && kind != reflect.Func {
			return errors.New(fmt.Sprintf("绑定字段 %s: 应传递绑定变量指针", name))
		}
		var refType GO_VALUE_TYPE
		var bindType V8_JS_VALUE_TYPE = -1
		if kind == reflect.Func {
			bindType = V8_VALUE_FUNCTION
		} else {
			refType = common.GOValueReflectType(typ.Elem())
			switch refType {
			case GO_VALUE_INT, GO_VALUE_INT8, GO_VALUE_INT16, GO_VALUE_INT32, GO_VALUE_INT64, GO_VALUE_UINT, GO_VALUE_UINT8, GO_VALUE_UINT16, GO_VALUE_UINT32, GO_VALUE_UINT64, GO_VALUE_UINTPTR:
				bindType = V8_VALUE_INT
			case GO_VALUE_FLOAT32, GO_VALUE_FLOAT64:
				bindType = V8_VALUE_DOUBLE
			case GO_VALUE_STRING:
				bindType = V8_VALUE_STRING
			case GO_VALUE_BOOL:
				bindType = V8_VALUE_BOOLEAN
			case GO_VALUE_NIL:
				bindType = V8_VALUE_NULL
			case GO_VALUE_FUNC:
				bindType = V8_VALUE_FUNCTION
			case GO_VALUE_STRUCT, GO_VALUE_MAP:
				bindType = V8_VALUE_OBJECT
			case GO_VALUE_SLICE:
				bindType = V8_VALUE_ARRAY
			}
		}
		if bindType == -1 {
			return errors.New("类型错误, 支持类型: string, int32, float64, bool, func, struct, map, slice")
		}
		fmt.Println("name", name, "refType", bindType)
		value := &V8Value{
			name:      name,
			rwLock:    new(sync.Mutex),
			valueType: new(VT),
		}
		value.valueType.Jsv = bindType
		value.valueType.Gov = refType
		if bindType == V8_VALUE_FUNCTION {
			if info, err := checkFunc(reflect.TypeOf(bind), FN_TYPE_COMMON); err == nil {
				value.funcInfo = info
				value.ptr = unsafe.Pointer(&bind)
			} else {
				return err
			}
		} else if bindType == V8_VALUE_OBJECT {
			if refType == GO_VALUE_STRUCT {

			} else if refType == GO_VALUE_MAP {

			}
		} else if bindType == V8_VALUE_ARRAY {

		} else {
			switch bindType {
			case V8_VALUE_STRING:
				value.ptr = unsafe.Pointer(bind.(*string))
			case V8_VALUE_INT:
				switch refType {
				case GO_VALUE_INT:
					value.ptr = unsafe.Pointer(bind.(*int))
				case GO_VALUE_INT8:
					value.ptr = unsafe.Pointer(bind.(*int8))
				case GO_VALUE_INT16:
					value.ptr = unsafe.Pointer(bind.(*int16))
				case GO_VALUE_INT32:
					value.ptr = unsafe.Pointer(bind.(*int32))
				case GO_VALUE_INT64:
					value.ptr = unsafe.Pointer(bind.(*int64))
				case GO_VALUE_UINT:
					value.ptr = unsafe.Pointer(bind.(*uint))
				case GO_VALUE_UINT8:
					value.ptr = unsafe.Pointer(bind.(*uint8))
				case GO_VALUE_UINT16:
					value.ptr = unsafe.Pointer(bind.(*uint16))
				case GO_VALUE_UINT32:
					value.ptr = unsafe.Pointer(bind.(*uint32))
				case GO_VALUE_UINT64:
					value.ptr = unsafe.Pointer(bind.(*uint64))
				case GO_VALUE_UINTPTR:
					value.ptr = unsafe.Pointer(bind.(*uintptr))
				}
			case V8_VALUE_DOUBLE:
				if refType == GO_VALUE_FLOAT32 {
					value.ptr = unsafe.Pointer(bind.(*float32))
				} else {
					value.ptr = unsafe.Pointer(bind.(*float64))
				}
			case V8_VALUE_BOOLEAN:
				value.ptr = unsafe.Pointer(bind.(*bool))
			default:
				return errors.New("字段 " + name + ": 不支持的绑定类型")
			}
		}
		value.instance = uintptr(value.ptr)
		value.value = bind
		value.that = value
		value.eventId = uintptr(__bind_id())
		value.isCommonObject = IS_OBJECT
		m.putValueBind(name, value)
		objectTI.bind(value)
		//fieldInfo := &fieldInfo{
		//	EventId: uintptr(__bind_id()),
		//	ValueType: &VT{
		//		Jsv: bindType,
		//		Gov: refType,
		//	},
		//	FieldValue: &filedValue,
		//}
	}
	return nil
}
