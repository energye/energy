//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue js绑定实现
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

// V8Value 绑定到JS的字段
type V8Value struct {
	eventId        uintptr        //事件ID
	instance       uintptr        //
	ptr            unsafe.Pointer //
	name           string         //用于字段或函数名
	value          interface{}    //值
	valueType      *VT            //字段类型，go 和 js
	funcInfo       *funcInfo      //普通函数信息
	sfi            *functionInfo  //对象函数信息
	isCommonObject IS_CO          //通用类型或对象类型 默认通用类型
	that           JSValue        //
	rwLock         *sync.Mutex    //字段|函数独有锁
}

// CEFv8BindRoot 导出给JavaScript window的根对象名
type CEFv8BindRoot struct {
	commonRootName string //V8Value 通用类型变量属性的所属默认对象名称
	objectRootName string //V8Value 对象类型变量属性的所属默认对象名称
}

// checkFunc 检查函数是否导出, 返回：函数描述信息
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
	if numIn < 0 {
		numIn = 0
	}
	r := new(funcInfo)
	r.InNum = int32(numIn)
	r.OutNum = int32(numOut)
	r.InParam = make([]*VT, numIn, numIn)
	r.FnType = fnType
	var idx = 0
	for i := 0; i < numIn; i++ {
		idx = i + int(fnType)
		inTyp := fnOf.In(idx).Kind()
		if gov, jsv := common.FieldReflectType(inTyp); jsv == -1 && gov == -1 {
			//入参参数类型不正确
			return nil, errors.New("WindowBind function parameter error: input parameter type can only be [string int float bool]")
		} else {
			r.InParam[i] = &VT{Jsv: jsv, Gov: gov}
		}
	}
	if numOut > 0 && numOut < 3 {
		r.OutParam = make([]*VT, numOut, numOut)
		for i := 0; i < numOut; i++ {
			outTyp := fnOf.Out(i)
			if gov, jsv := common.FieldReflectType(outTyp.Kind()); jsv == -1 && gov == -1 {
				//出参参数类型不正确
				return nil, errors.New("WindowBind function parameter error: output parameter type can only be [string int float bool]")
			} else {
				r.OutParam[i] = &VT{Jsv: jsv, Gov: gov}
				r.OutParamIdx = int32(i)
			}
		}
	}
	return r, nil
}

// Lock 每一个变量的独有锁 - 加锁
func (m *V8Value) Lock() {
	m.rwLock.Lock()
}

// UnLock 每一个变量的独有锁 - 解锁
func (m *V8Value) UnLock() {
	m.rwLock.Unlock()
}

// V8Value isCommon
func (m *V8Value) isCommon() bool {
	return m.isCommonObject == IS_COMMON
}

func (m *V8Value) setThat(that JSValue) {
	m.that = that
}

// Bytes 值转换为字节
func (m *V8Value) Bytes() []byte {
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
	switch m.valueType.Jsv {
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

// ValueToPtr 转换为指针
func (m *V8Value) ValueToPtr() (unsafe.Pointer, error) {
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
	switch m.valueType.Jsv {
	case V8_VALUE_STRING:
		return unsafe.Pointer(api.PascalStr(iValue.(string))), nil
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
			return unsafe.Pointer(api.PascalBool(v)), nil
		} else {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported type")
	}
}

// SetAnyValue 设置多类型值
func (m *V8Value) SetAnyValue(value interface{}) error {
	switch common.JSValueAssertType(value) {
	case V8_VALUE_STRING:
		m.valueType.Jsv = V8_VALUE_STRING
		m.valueType.Gov = GO_VALUE_STRING
	case V8_VALUE_INT:
		m.valueType.Jsv = V8_VALUE_INT
		m.valueType.Gov = GO_VALUE_INT32
	case V8_VALUE_DOUBLE:
		m.valueType.Jsv = V8_VALUE_DOUBLE
		m.valueType.Gov = GO_VALUE_FLOAT64
	case V8_VALUE_BOOLEAN:
		m.valueType.Jsv = V8_VALUE_BOOLEAN
		m.valueType.Gov = GO_VALUE_BOOL
	default:
		return errors.New(cefErrorMessage(CVE_ERROR_TYPE_NOT_SUPPORTED))
	}
	m.setValue(value)
	return nil
}

func (m *V8Value) getFuncInfo() *funcInfo {
	return m.funcInfo
}

func (m *V8Value) Instance() uintptr {
	return m.instance
}

func (m *V8Value) Ptr() unsafe.Pointer {
	return m.ptr
}

func (m *V8Value) Name() string {
	return m.name
}

func (m *V8Value) setName(name string) {
	m.name = name
}

func (m *V8Value) getValue() interface{} {
	return m.value
}

func (m *V8Value) setValue(value interface{}) {
	m.value = value
}

func (m *V8Value) ValueType() *VT {
	return m.valueType
}

func (m *V8Value) setPtr(ptr unsafe.Pointer) {
	m.ptr = ptr
}

func (m *V8Value) setInstance(instance uintptr) {
	m.instance = instance
}

func (m *V8Value) StringValue() (string, error) {
	if m.IsString() {
		return common.ValueToString(m.value)
	}
	return "", errors.New("failed to get a string value")
}

func (m *V8Value) IntegerValue() (int32, error) {
	if m.IsInteger() {
		return common.ValueToInt32(m.value)
	}
	return 0, errors.New("failed to get a integer value")
}

func (m *V8Value) DoubleValue() (float64, error) {
	if m.IsDouble() {
		return common.ValueToFloat64(m.value)
	}
	return 0, errors.New("failed to get a double value")
}

func (m *V8Value) BooleanValue() (bool, error) {
	if m.IsBool() {
		return common.ValueToBool(m.value)
	}
	return false, errors.New("failed to get a boolean value")
}

func (m *V8Value) IsString() bool {
	return m.valueType.Jsv == V8_VALUE_STRING
}

func (m *V8Value) IsInteger() bool {
	return m.valueType.Jsv == V8_VALUE_INT
}

func (m *V8Value) IsDouble() bool {
	return m.valueType.Jsv == V8_VALUE_DOUBLE
}

func (m *V8Value) IsBool() bool {
	return m.valueType.Jsv == V8_VALUE_BOOLEAN
}

func (m *V8Value) IsArray() bool {
	return m.valueType.Jsv == V8_VALUE_ARRAY
}

func (m *V8Value) IsObject() bool {
	return m.valueType.Jsv == V8_VALUE_OBJECT
}

func (m *V8Value) IsFunction() bool {
	return m.valueType.Jsv == V8_VALUE_FUNCTION
}

func (m *V8Value) IsNull() bool {
	return m.valueType.Jsv == V8_VALUE_NULL
}

func (m *V8Value) IsUndefined() bool {
	return m.valueType.Jsv == V8_VALUE_UNDEFINED
}

func (m *V8Value) AsString() *JSString {
	if m.that != nil {
		if v, ok := m.that.(*JSString); ok {
			return v
		}
	}
	return nil
}

func (m *V8Value) AsInteger() *JSInteger {
	if m.that != nil {
		if v, ok := m.that.(*JSInteger); ok {
			return v
		}
	}
	return nil
}

func (m *V8Value) AsDouble() *JSDouble {
	if m.that != nil {
		if v, ok := m.that.(*JSDouble); ok {
			return v
		}
	}
	return nil
}

func (m *V8Value) AsBoolean() *JSBoolean {
	if m.that != nil {
		if v, ok := m.that.(*JSBoolean); ok {
			return v
		}
	}
	return nil
}

func (m *V8Value) AsV8Value() *V8Value {
	return m
}

func (m *V8Value) AsArray() *JSArray {
	if m.that != nil {
		if v, ok := m.that.(*JSArray); ok {
			return v
		}
	}
	return nil
}

func (m *V8Value) AsFunction() *JSFunction {
	if m.that != nil {
		if v, ok := m.that.(*JSFunction); ok {
			return v
		}
	}
	return nil
}

func (m *V8Value) invoke(inParams []reflect.Value) (outParams []reflect.Value, success bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("invoke recover:", err)
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

func newV8Value(eventId uintptr, fullParentName, name string, value interface{}, sfi *functionInfo, valueType *VT, isCommonObject IS_CO) JSValue {
	jsValueBind := new(V8Value)
	jsValueBind.valueType = new(VT)
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
	VariableBind.addBind(fmt.Sprintf("%s.%s", fullParentName, name), jsValueBind)
	return jsValueBind
}

func (m *V8Value) Pointer() unsafe.Pointer {
	return m.ptr
}

func (m *V8Value) setEventId(eventId uintptr) {
	m.eventId = eventId
}

func (m *V8Value) getEventId() uintptr {
	return m.eventId
}
