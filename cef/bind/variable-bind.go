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
package bind

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
	isMainProcess, isRenderProcess bool //
)

// VariableBind 变量绑定
var VariableBind IProvisionalBindStorage

type IProvisionalBindStorage interface {
	NewString(name, value string) *JSString         //通用类型 - 默认: string
	NewInteger(name string, value int32) *JSInteger //通用类型 - 默认: integer
	NewDouble(name string, value float64) *JSDouble //通用类型 - 默认: double
	NewBoolean(name string, value bool) *JSBoolean  //通用类型 - 默认: boolean
	NewNull(name string) *JSNull                    //通用类型 - 默认: null
	NewUndefined(name string) *JSUndefined          //通用类型 - 默认: undefined
	NewFunction(name string, fn interface{}) error  //固定类型 - function
	NewObjects(objects ...interface{})              //固定类型 - object - struct
	Bind(name string, bind interface{}) error       //固定类型 - 所有支持类型
	getBindValue(fullName string) (JSValue, bool)   //
	binds() map[string]JSValue                      //
	bindCount() int                                 //
	addBind(fullName string, value JSValue)         //
}

type variableBind struct {
	bindMapping map[string]JSValue //所有绑定变量属性或函数集合
	lock        sync.Mutex         //add bind, remove bind lock
}

// VT 字段类型
type VT struct {
	Jsv V8_JS_VALUE_TYPE `json:"jsv"`
	Gov GO_VALUE_TYPE    `json:"gov"`
}

// 函数详情 1.参数个数 2.每个参数类型 3.返回参数类型
type funcInfo struct {
	InNum          int32   `json:"inNum"`          //入参个数
	InParam        []*VT   `json:"inParam"`        //入参类型
	OutNum         int32   `json:"outNum"`         //出参个数
	OutParam       []*VT   `json:"outParam"`       //出参类型
	OutParamErrIdx int32   `json:"outParamErrIdx"` //出参错误位置
	OutParamIdx    int32   `json:"outParamIdx"`    //出参位置
	FnType         FN_TYPE `json:"fnType"`         //函数类型, 直接定义函数或对象函数
}

// functionInfo 结构的函数描述
type functionInfo struct {
	EventId uintptr `json:"event_id"`
	*funcInfo
	Method reflect.Value `json:"-"`
}

func init() {
	isMainProcess = common.Args.IsMain()
	isRenderProcess = common.Args.IsRender()
	VariableBind = &variableBind{bindMapping: make(map[string]JSValue)}
	fmt.Println("isMainProcess:", isMainProcess, "isRenderProcess:", isRenderProcess)
}

func (m *VT) ToString() string {
	gov := common.FuncParamGoTypeStr(m.Gov)
	jsv := common.FuncParamJsTypeStr(m.Jsv)
	return fmt.Sprintf("GO=%s JS=%s", gov, jsv)
}

// IsGoIntAuto 判断Go 所有 int 类型
func (m *VT) IsGoIntAuto() bool {
	switch m.Gov {
	case GO_VALUE_INT, GO_VALUE_INT8, GO_VALUE_INT16, GO_VALUE_INT32, GO_VALUE_INT64, GO_VALUE_UINT, GO_VALUE_UINT8, GO_VALUE_UINT16, GO_VALUE_UINT32, GO_VALUE_UINT64:
		return true
	}
	return false
}

// IsGoFloatAuto 判断Go 所有 float 类型
func (m *VT) IsGoFloatAuto() bool {
	switch m.Gov {
	case GO_VALUE_FLOAT32, GO_VALUE_FLOAT64:
		return true
	}
	return false
}

func (m *variableBind) addBind(fullName string, value JSValue) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.bindMapping[fullName] = value
}

func (m *variableBind) getBindValue(fullName string) (value JSValue, ok bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	value, ok = m.bindMapping[fullName]
	return
}

func (m *variableBind) binds() map[string]JSValue {
	return m.bindMapping
}

func (m *variableBind) bindCount() int {
	return len(m.bindMapping)
}

// NewString V8Value 通用类型, 默认 string
func (m *variableBind) NewString(name, value string) *JSString {
	jsValueBind := new(JSString)
	jsValueBind.valueType = new(VT)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_STRING
	jsValueBind.valueType.Gov = GO_VALUE_STRING
	m.bindValueHandle(jsValueBind)
	return jsValueBind
}

// NewInteger V8Value 通用类型, 默认 integer
func (m *variableBind) NewInteger(name string, value int32) *JSInteger {
	jsValueBind := new(JSInteger)
	jsValueBind.valueType = new(VT)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_INT
	jsValueBind.valueType.Gov = GO_VALUE_INT32
	m.bindValueHandle(jsValueBind)
	return jsValueBind
}

// NewDouble V8Value 通用类型, 默认 double
func (m *variableBind) NewDouble(name string, value float64) *JSDouble {
	jsValueBind := new(JSDouble)
	jsValueBind.valueType = new(VT)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_DOUBLE
	jsValueBind.valueType.Gov = GO_VALUE_FLOAT64
	m.bindValueHandle(jsValueBind)
	return jsValueBind
}

// NewBoolean V8Value 通用类型, 默认 boolean
func (m *variableBind) NewBoolean(name string, value bool) *JSBoolean {
	jsValueBind := new(JSBoolean)
	jsValueBind.valueType = new(VT)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType.Jsv = V8_VALUE_BOOLEAN
	jsValueBind.valueType.Gov = GO_VALUE_BOOL
	m.bindValueHandle(jsValueBind)
	return jsValueBind
}

// NewNull V8Value 通用类型, 默认 null
func (m *variableBind) NewNull(name string) *JSNull {
	jsValueBind := new(JSNull)
	jsValueBind.valueType = new(VT)
	jsValueBind.name = name
	jsValueBind.value = "null"
	jsValueBind.valueType.Jsv = V8_VALUE_NULL
	jsValueBind.valueType.Gov = GO_VALUE_NIL
	m.bindValueHandle(jsValueBind)
	return jsValueBind
}

// NewUndefined V8Value 通用类型, 默认 undefined
func (m *variableBind) NewUndefined(name string) *JSUndefined {
	jsValueBind := new(JSUndefined)
	jsValueBind.valueType = new(VT)
	jsValueBind.name = name
	jsValueBind.value = "undefined"
	jsValueBind.valueType.Jsv = V8_VALUE_UNDEFINED
	jsValueBind.valueType.Gov = GO_VALUE_STRING
	m.bindValueHandle(jsValueBind)
	return jsValueBind
}

// NewFunction V8Value
func (m *variableBind) NewFunction(name string, fn interface{}) error {
	if gov, _ := common.FieldReflectType(fn); gov == GO_VALUE_FUNC {
		if info, err := checkFunc(reflect.TypeOf(fn), FN_TYPE_COMMON); err == nil {
			jsValueBind := new(JSFunction)
			jsValueBind.valueType = new(VT)
			jsValueBind.name = name
			jsValueBind.value = fn
			jsValueBind.valueType.Jsv = V8_VALUE_FUNCTION
			jsValueBind.valueType.Gov = GO_VALUE_FUNC
			jsValueBind.funcInfo = info
			m.bindValueHandle(jsValueBind)
			return nil
		} else {
			return err
		}
	}
	return errors.New("创建的函数不是函数类型")
}

// V8Value bindValueHandle
func (m *variableBind) bindValueHandle(jsValue JSValue) {
	jsValue.setInstance(unsafe.Pointer(&jsValue))
	jsValue.setEventId(jsValue.Instance())
	jsValue.setThat(jsValue)
	m.addBind(jsValue.Name(), jsValue)
}

// NewObjects V8Value
//
// 对象类型变量和对象函数绑定
func (m *variableBind) NewObjects(objects ...interface{}) {

}

// Bind V8Value
//
// 变量和函数绑定
func (m *variableBind) Bind(name string, bind interface{}) error {
	if isMainProcess || isRenderProcess {
		typ := reflect.TypeOf(bind)
		kind := typ.Kind()
		if kind != reflect.Ptr && kind != reflect.Func {
			return errors.New(fmt.Sprintf("bind parameter %s needs to pass pointer", name))
		}
		var (
			gov   GO_VALUE_TYPE
			jsv   V8_JS_VALUE_TYPE
			value *V8Value
		)
		if kind == reflect.Func {
			gov, jsv = GO_VALUE_FUNC, V8_VALUE_FUNCTION
		} else {
			gov, jsv = common.FieldReflectType(typ.Elem())
		}
		fmt.Println("FieldReflectType:", gov, jsv)
		if gov == -1 || jsv == -1 {
			return errors.New("parameter type mismatch")
		}
		value = &V8Value{
			name:      name,
			valueType: &VT{Jsv: jsv, Gov: gov},
			//isCommonObject: IS_OBJECT,
			//eventId:        uintptr(__bind_id()),
		}
		if isMainProcess {
			//类型信息
			if jsv == V8_VALUE_FUNCTION { //function
				if info, err := checkFunc(typ, FN_TYPE_OBJECT); err == nil {
					value.funcInfo = info
					value.instance = unsafe.Pointer(&bind)
					value.value = reflect.ValueOf(bind)
				} else {
					return err
				}
			} else if jsv == V8_VALUE_OBJECT { //object
				if gov == GO_VALUE_STRUCT { // object - struct

				} else if gov == GO_VALUE_MAP { // object - map

				}
			} else if jsv == V8_VALUE_ARRAY { //array

			} else { //field
				switch jsv {
				case V8_VALUE_STRING:
					value.instance = unsafe.Pointer(bind.(*string))
				case V8_VALUE_INT:
					switch gov {
					case GO_VALUE_INT:
						value.instance = unsafe.Pointer(bind.(*int))
					case GO_VALUE_INT8:
						value.instance = unsafe.Pointer(bind.(*int8))
					case GO_VALUE_INT16:
						value.instance = unsafe.Pointer(bind.(*int16))
					case GO_VALUE_INT32:
						value.instance = unsafe.Pointer(bind.(*int32))
					case GO_VALUE_INT64:
						value.instance = unsafe.Pointer(bind.(*int64))
					case GO_VALUE_UINT:
						value.instance = unsafe.Pointer(bind.(*uint))
					case GO_VALUE_UINT8:
						value.instance = unsafe.Pointer(bind.(*uint8))
					case GO_VALUE_UINT16:
						value.instance = unsafe.Pointer(bind.(*uint16))
					case GO_VALUE_UINT32:
						value.instance = unsafe.Pointer(bind.(*uint32))
					case GO_VALUE_UINT64:
						value.instance = unsafe.Pointer(bind.(*uint64))
					case GO_VALUE_UINTPTR:
						value.instance = unsafe.Pointer(bind.(*uintptr))
					}
				case V8_VALUE_DOUBLE:
					if gov == GO_VALUE_FLOAT32 {
						value.instance = unsafe.Pointer(bind.(*float32))
					} else {
						value.instance = unsafe.Pointer(bind.(*float64))
					}
				case V8_VALUE_BOOLEAN:
					value.instance = unsafe.Pointer(bind.(*bool))
				default:
					return errors.New("parameter type mismatch")
				}
				value.value = bind
			}
			value.that = value
		}
		m.addBind(name, value)
	}
	return nil
}
