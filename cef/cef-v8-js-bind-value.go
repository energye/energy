//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"errors"
	"fmt"
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl/api"
	"reflect"
	"sync"
	"unsafe"
)

type valueBindInfo struct {
	Name           uintptr //string 字段名称
	EventId        uintptr //
	BindType       uintptr //int32 变量类型 0:string 1:int 2:double 3:bool 4:null 5:undefined 6:object 7:array 8:function
	FnInNum        uintptr //int32 入参个数
	FnInParamType  uintptr //string 入参变量类型 0,1,2,3,1
	FnOutNum       uintptr //int32 出参个数
	FnOutParamType uintptr //string 出参变量类型
}

//绑定到Js的字段
type ICEFv8Value struct {
	eventId        uintptr
	instance       uintptr
	ptr            unsafe.Pointer
	name           string           //用于字段或函数名
	value          interface{}      //值
	valueType      V8_JS_VALUE_TYPE //0:string 1:int 2:double 3:bool 4:null 5:undefined 6:object 7:array 8:function
	funcInfo       *funcInfo        //普通函数信息
	sfi            *structFuncInfo  //对象函数信息
	isCommonObject IS_CO            //通用类型或对象类型 默认通用类型
	that           JSValue
	rwLock         *sync.Mutex
}

type CEFv8BindRoot struct {
	commonRootName string //ICEFv8Value 通用类型变量属性的所属默认对象名称
	objectRootName string //ICEFv8Value 对象类型变量属性的所属默认对象名称
}

//检查函数是否符合 返回：函数详情
func checkFunc(fnOf reflect.Type, fnType FN_TYPE) (*funcInfo, error) {
	numIn := fnOf.NumIn() - int(fnType)
	numOut := fnOf.NumOut()
	if numOut > BIND_FUNC_OUT_MAX_SUM {
		//出参个数超出
		return nil, errors.New(fmt.Sprintf("WindowBind function parameter error: the maximum number of function parameters is %d, but the actual number is %d", BIND_FUNC_OUT_MAX_SUM, numOut))
	}
	if numIn > BIND_FUNC_IN_MAX_SUM {
		//入参个数超出
		return nil, errors.New(fmt.Sprintf("WindowBind function parameter error: up to %d function input parameters, actually %d", BIND_FUNC_IN_MAX_SUM, numIn))
	}
	r := new(funcInfo)
	r.InNum = int32(numIn)
	r.OutNum = int32(numOut)
	r.InParam = make([]*vt, numIn, numIn)
	r.FnType = fnType
	var idx = 0
	for i := 0; i < numIn; i++ {
		idx = i + int(fnType)
		inTyp := fnOf.In(idx).Kind().String()
		if jsv, gov := common.ParamType(inTyp); jsv == -1 && gov == -1 {
			//入参参数类型不正确
			return nil, errors.New("WindowBind function parameter error: input parameter type can only be [string int float bool]")
		} else {
			r.InParam[i] = &vt{Jsv: jsv, Gov: gov}
		}
	}
	if numOut > 0 && numOut < 3 {
		r.OutParam = make([]*vt, numOut, numOut)
		for i := 0; i < numOut; i++ {
			outTyp := fnOf.Out(i)
			if jsv, gov := common.ParamType(outTyp.Kind().String()); jsv == -1 && gov == -1 {
				//出参类型错误
				return nil, errors.New("WindowBind function parameter error: output parameter type can only be [string int float bool]")
			} else {
				r.OutParam[i] = &vt{Jsv: jsv, Gov: gov}
				r.OutParamIdx = int32(i)
			}
		}
	}
	return r, nil
}

func (m *ICEFv8Value) Lock() {
	m.rwLock.Lock()
}

func (m *ICEFv8Value) UnLock() {
	m.rwLock.Unlock()
}

//ICEFv8Value isCommon
func (m *ICEFv8Value) isCommon() bool {
	return m.isCommonObject == IS_COMMON
}

func (m *ICEFv8Value) setThat(that JSValue) {
	m.that = that
}

func (m *ICEFv8Value) Bytes() []byte {
	var iValue interface{}
	if m.isCommon() {
		iValue = m.value
	} else { //object
		if refValue, ok := m.value.(*reflect.Value); ok {
			iValue = refValue.Interface()
		} else {
			return nil
		}
	}
	switch m.valueType {
	case V8_VALUE_STRING:
		return common.StringToBytes(iValue.(string))
	case V8_VALUE_INT:
		if v, err := common.ValueToInt32(iValue); err == nil {
			return common.Int32ToBytes(v)
		} else {
			return nil
		}
	case V8_VALUE_DOUBLE:
		if v, err := common.ValueToFloat64(iValue); err == nil {
			return common.Float64ToBytes(v)
		} else {
			return nil
		}
	case V8_VALUE_BOOLEAN:
		if v, err := common.ValueToBool(iValue); err == nil {
			return []byte{common.BoolToByte(v)}
		} else {
			return nil
		}
	default:
		return nil
	}
}

func (m *ICEFv8Value) ValueToPtr() (unsafe.Pointer, error) {
	var iValue interface{}
	if m.isCommon() {
		iValue = m.value
	} else { //object
		if refValue, ok := m.value.(*reflect.Value); ok {
			iValue = refValue.Interface()
		} else {
			return nil, errors.New("object转换reflect.Value失败")
		}
	}
	switch m.valueType {
	case V8_VALUE_STRING:
		return unsafe.Pointer(api.GoStrToDStr(iValue.(string))), nil
		//return GoStrToDStrPointer(iValue.(string)), nil
	case V8_VALUE_INT:
		if v, err := common.ValueToInt32(iValue); err == nil {
			return unsafe.Pointer(uintptr(v)), nil
		} else {
			return nil, err
		}
	case V8_VALUE_DOUBLE:
		if v, err := common.ValueToFloat64(iValue); err == nil {
			return unsafe.Pointer(&v), nil
		} else {
			return nil, err
		}
	case V8_VALUE_BOOLEAN:
		if v, err := common.ValueToBool(iValue); err == nil {
			return unsafe.Pointer(api.GoBoolToDBool(v)), nil
		} else {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported type")
	}
}

func (m *ICEFv8Value) SetAnyValue(value interface{}) error {
	switch common.JSValueAssertType(value) {
	case V8_VALUE_STRING:
		m.setValueType(V8_VALUE_STRING)
	case V8_VALUE_INT:
		m.setValueType(V8_VALUE_INT)
	case V8_VALUE_DOUBLE:
		m.setValueType(V8_VALUE_DOUBLE)
	case V8_VALUE_BOOLEAN:
		m.setValueType(V8_VALUE_BOOLEAN)
	default:
		return errors.New(cefErrorMessage(CVE_ERROR_TYPE_NOT_SUPPORTED))
	}
	m.setValue(value)
	return nil
}

func (m *ICEFv8Value) getFuncInfo() *funcInfo {
	return m.funcInfo
}

func (m *ICEFv8Value) Instance() uintptr {
	return m.instance
}

func (m *ICEFv8Value) Ptr() unsafe.Pointer {
	return m.ptr
}

func (m *ICEFv8Value) Name() string {
	return m.name
}

func (m *ICEFv8Value) setName(name string) {
	m.name = name
}

func (m *ICEFv8Value) getValue() interface{} {
	return m.value
}

func (m *ICEFv8Value) setValue(value interface{}) {
	m.value = value
}

func (m *ICEFv8Value) ValueType() V8_JS_VALUE_TYPE {
	return m.valueType
}

func (m *ICEFv8Value) setValueType(vType V8_JS_VALUE_TYPE) {
	m.valueType = vType
}

func (m *ICEFv8Value) setPtr(ptr unsafe.Pointer) {
	m.ptr = ptr
}

func (m *ICEFv8Value) setInstance(instance uintptr) {
	m.instance = instance
}

func (m *ICEFv8Value) StringValue() (string, error) {
	if m.IsString() {
		return common.ValueToString(m.value)
	}
	return "", errors.New("failed to get a string value")
}

func (m *ICEFv8Value) IntegerValue() (int32, error) {
	if m.IsInteger() {
		return common.ValueToInt32(m.value)
	}
	return 0, errors.New("failed to get a integer value")
}

func (m *ICEFv8Value) DoubleValue() (float64, error) {
	if m.IsDouble() {
		return common.ValueToFloat64(m.value)
	}
	return 0, errors.New("failed to get a double value")
}

func (m *ICEFv8Value) BooleanValue() (bool, error) {
	if m.IsBool() {
		return common.ValueToBool(m.value)
	}
	return false, errors.New("failed to get a boolean value")
}

func (m *ICEFv8Value) IsString() bool {
	return m.valueType == V8_VALUE_STRING
}

func (m *ICEFv8Value) IsInteger() bool {
	return m.valueType == V8_VALUE_INT
}

func (m *ICEFv8Value) IsDouble() bool {
	return m.valueType == V8_VALUE_DOUBLE
}

func (m *ICEFv8Value) IsBool() bool {
	return m.valueType == V8_VALUE_BOOLEAN
}

func (m *ICEFv8Value) IsArray() bool {
	return m.valueType == V8_VALUE_ARRAY
}

func (m *ICEFv8Value) IsObject() bool {
	return m.valueType == V8_VALUE_OBJECT
}

func (m *ICEFv8Value) IsFunction() bool {
	return m.valueType == V8_VALUE_FUNCTION
}

func (m *ICEFv8Value) IsNull() bool {
	return m.valueType == V8_VALUE_NULL
}

func (m *ICEFv8Value) IsUndefined() bool {
	return m.valueType == V8_VALUE_UNDEFINED
}

func (m *ICEFv8Value) AsString() (*JSString, error) {
	if m.that != nil {
		if v, ok := m.that.(*JSString); ok {
			return v, nil
		}
	}
	return nil, errors.New("unable to cast to String")
}

func (m *ICEFv8Value) AsInteger() (*JSInteger, error) {
	if m.that != nil {
		if v, ok := m.that.(*JSInteger); ok {
			return v, nil
		}
	}
	return nil, errors.New("unable to cast to Integer")
}

func (m *ICEFv8Value) AsDouble() (*JSDouble, error) {
	if m.that != nil {
		if v, ok := m.that.(*JSDouble); ok {
			return v, nil
		}
	}
	return nil, errors.New("unable to cast to Double")
}

func (m *ICEFv8Value) AsBoolean() (*JSBoolean, error) {
	if m.that != nil {
		if v, ok := m.that.(*JSBoolean); ok {
			return v, nil
		}
	}
	return nil, errors.New("unable to cast to Boolean")
}

func (m *ICEFv8Value) AsArray() (*JSArray, error) {
	if m.that != nil {
		if v, ok := m.that.(*JSArray); ok {
			return v, nil
		}
	}
	return nil, errors.New("unable to cast to Array")
}

func (m *ICEFv8Value) AsFunction() (*JSFunction, error) {
	if m.that != nil {
		if v, ok := m.that.(*JSFunction); ok {
			return v, nil
		}
	}
	return nil, errors.New("unable to cast to Function")
}

func (m *ICEFv8Value) invoke(inParams []reflect.Value) (outParams []reflect.Value, success bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("V8BindFuncCallbackHandler recover:", err)
			outParams = []reflect.Value{reflect.ValueOf(err.(error).Error())}
			success = false
		}
	}()
	if m.isCommon() {
		outParams = reflect.ValueOf(m.value).Call(inParams)
	} else {
		outParams = m.sfi.Method.Call(inParams)
	}
	return outParams, true
}

//ICEFv8Value newICEFv8Value
func newICEFv8Value(eventId uintptr, fullParentName, name string, value interface{}, sfi *structFuncInfo, valueType V8_JS_VALUE_TYPE, isCommonObject IS_CO) JSValue {
	jsValueBind := new(ICEFv8Value)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.ptr = unsafe.Pointer(&jsValueBind)
	jsValueBind.instance = uintptr(jsValueBind.ptr)
	jsValueBind.eventId = eventId
	jsValueBind.name = name
	jsValueBind.value = value
	if sfi != nil {
		jsValueBind.funcInfo = sfi.funcInfo
	}
	jsValueBind.sfi = sfi
	jsValueBind.valueType = valueType
	jsValueBind.isCommonObject = isCommonObject

	//putValueBind(jsValueBind.eventId, jsValueBind)
	VariableBind.putValueBind(fmt.Sprintf("%s.%s", fullParentName, name), jsValueBind)
	return jsValueBind
}

//ICEFv8Value Pointer
func (m *ICEFv8Value) Pointer() unsafe.Pointer {
	return m.ptr
}

//ICEFv8Value setEventId
func (m *ICEFv8Value) setEventId(eventId uintptr) {
	m.eventId = eventId
}

//ICEFv8Value getEventId
func (m *ICEFv8Value) getEventId() uintptr {
	return m.eventId
}
