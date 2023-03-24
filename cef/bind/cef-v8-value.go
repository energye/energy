//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue 通用类型实现
//
// 通用类型是可变类型，对于JS语言可以动态赋任意类型的值
//
// # Go中针对几种基本类型增加了可变类型实现
//
// 绑定 定义的普通函数 前缀-默认是结构名称
package bind

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

// JSValue
// GO和JS变量类型接口
type JSValue interface {
	SetAnyValue(value interface{}) error //多类型值设置
	IntegerValue() (int32, error)        //返回 Integer 类型值, 成功 error = nil
	DoubleValue() (float64, error)       //返回 Double 类型值, 成功 error = nil
	StringValue() (string, error)        //返回 String 类型值, 成功 error = nil
	BooleanValue() (bool, error)         //返回 Boolean 类型值, 成功 error = nil
	IsInteger() bool                     //是否 Integer
	IsDouble() bool                      //是否 Double
	IsString() bool                      //是否 String
	IsBool() bool                        //是否 Bool
	IsArray() bool                       //是否 JSONArray
	IsObject() bool                      //是否 Object
	IsFunction() bool                    //是否 Function
	IsNull() bool                        //是否 Null
	IsUndefined() bool                   //是否 Undefined
	AsV8Value() *V8Value                 //转换为 V8Value
	AsInteger() *JSInteger               //转换为 Integer 失败返回 nil
	AsDouble() *JSDouble                 //转换为 Double 失败返回 nil
	AsString() *JSString                 //转换为 String 失败返回 nil
	AsBoolean() *JSBoolean               //转换为 Boolean 失败返回 nil
	AsArray() *JSArray                   //转换为 JSONArray 失败返回 nil
	AsFunction() *JSFunction             //转换为 Function 失败返回 nil
	Instance() uintptr                   //当前变量指针
	Name() string                        //当前变量绑定的名称
	ValueType() *VT                      //变量类型
	Bytes() []byte                       //变量值转换为字节
	ValueToPtr() (unsafe.Pointer, error) //值转为指针
	invoke(inParams []reflect.Value) (outParams []reflect.Value, success bool)
	setInstance(instance unsafe.Pointer)
	setName(name string)
	getValue() interface{}
	setValue(value interface{})
	getFuncInfo() *funcInfo
	setEventId(eventId uintptr)
	getEventId() uintptr
	isCommon() bool
	setThat(that JSValue)
}
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
	instance       unsafe.Pointer //
	name           string         //用于字段或函数名
	value          interface{}    //值
	valueType      *VT            //字段类型，go 和 js
	funcInfo       *funcInfo      //普通函数信息
	sfi            *functionInfo  //对象函数信息
	isCommonObject IS_CO          //通用类型或对象类型 默认通用类型
	that           JSValue        //
	rwLock         sync.Mutex     //字段|函数独有锁
}

// CEFv8BindRoot 导出给JavaScript window的根对象名
type CEFv8BindRoot struct {
	commonRootName string //V8Value 通用类型变量属性的所属默认对象名称
	objectRootName string //V8Value 对象类型变量属性的所属默认对象名称
}

// checkFunc 检查函数是否导出, 返回：函数描述信息
func checkFunc(fnOf reflect.Type, fnType FN_TYPE) (*funcInfo, error) {
	numIn := fnOf.NumIn() - int(fnType) // 结构函数的参数-1后才是入参
	numOut := fnOf.NumOut()
	//if numOut > BIND_FUNC_OUT_MAX_SUM {
	//	//出参个数超出
	//	return nil, errors.New(fmt.Sprintf("WindowBind function parameter error: the maximum number of function parameters is %d, but the actual number is %d", BIND_FUNC_OUT_MAX_SUM, numOut))
	//}
	//if numIn > BIND_FUNC_IN_MAX_SUM {
	//	//入参个数超出
	//	return nil, errors.New(fmt.Sprintf("WindowBind function parameter error: up to %d function input parameters, actually %d", BIND_FUNC_IN_MAX_SUM, numIn))
	//}
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
		if gov, jsv := common.FieldReflectType(inTyp); jsv == -1 || gov == -1 {
			//入参参数类型不支持
			return nil, errors.New("input parameter type mismatch")
		} else {
			r.InParam[i] = &VT{Jsv: jsv, Gov: gov}
		}
	}
	r.OutParam = make([]*VT, numOut, numOut)
	for i := 0; i < numOut; i++ {
		outTyp := fnOf.Out(i)
		if gov, jsv := common.FieldReflectType(outTyp.Kind()); jsv == -1 || gov == -1 {
			//出参参数类型不支持
			return nil, errors.New("out parameter type mismatch")
		} else {
			r.OutParam[i] = &VT{Jsv: jsv, Gov: gov}
			r.OutParamIdx = int32(i)
		}
	}
	return r, nil
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
func (m *V8Value) SetAnyValue(value any) error {
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
		return errors.New("invalid type")
	}
	m.setValue(value)
	return nil
}

func (m *V8Value) getFuncInfo() *funcInfo {
	return m.funcInfo
}

func (m *V8Value) Instance() uintptr {
	return uintptr(m.instance)
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

func (m *V8Value) setInstance(instance unsafe.Pointer) {
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

// NewString GO&JS 类型
func (m *V8Value) NewString(name, value string) *JSString {
	return VariableBind.NewString(name, value)
}

// NewInteger GO&JS 类型
func (m *V8Value) NewInteger(name string, value int32) *JSInteger {
	return VariableBind.NewInteger(name, value)
}

// NewDouble GO&JS 类型
func (m *V8Value) NewDouble(name string, value float64) *JSDouble {
	return VariableBind.NewDouble(name, value)
}

// NewBoolean  GO&JS 类型
func (m *V8Value) NewBoolean(name string, value bool) *JSBoolean {
	return VariableBind.NewBoolean(name, value)
}

// NewNull GO&JS 类型
func (m *V8Value) NewNull(name string) *JSNull {
	return VariableBind.NewNull(name)
}

// NewUndefined GO&JS 类型
func (m *V8Value) NewUndefined(name string) *JSUndefined {
	return VariableBind.NewUndefined(name)
}

// NewFunction GO&JS 类型
func (m *V8Value) NewFunction(name string, fn interface{}) error {
	return VariableBind.NewFunction(name, fn)
}

//	NewObjects GO&JS 类型
//
// Go结构类型变量和Go结构函数绑定
func (m *V8Value) NewObjects(objects ...interface{}) {
	VariableBind.NewObjects(objects...)
}

// Bind V8Value
//
// 变量和函数绑定, 在Go中定义的字段绑定到JS字段中, 在Go中定义的函数导出到JS
//
// 支持类型 String = string , Integer = int32 , Double = float64, Boolean = bool, Function = func, Objects = struct | map,  JSONArray = Slice
//
// 主进程和子进程
func (m *V8Value) Bind(name string, bind interface{}) error {
	return VariableBind.Bind(name, bind)
}

func newV8Value(eventId uintptr, fullParentName, name string, value interface{}, sfi *functionInfo, valueType *VT, isCommonObject IS_CO) JSValue {
	jsValueBind := new(V8Value)
	jsValueBind.valueType = new(VT)
	jsValueBind.instance = unsafe.Pointer(&jsValueBind)
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

func (m *V8Value) setEventId(eventId uintptr) {
	m.eventId = eventId
}

func (m *V8Value) getEventId() uintptr {
	return m.eventId
}
