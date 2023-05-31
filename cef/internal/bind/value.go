//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue 动态类型实现
//
// 动态类型是可变类型，对于JS语言可以动态赋任意类型的值
//
// 字段值仅主进程有效, 非主进程字段值为默认值

package bind

import (
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/json"
	"reflect"
	"strconv"
	"strings"
)

const (
	null      = "null"
	undefined = "undefined"
)

// BindType 绑定值类型
//	动态类型是可变类型, NewDynamic创建绑定变量, 函数传值方式
//	静态类型是不可变类型, NewStatic创建绑定变量, 直接传入变量指针, 非 interface{} 类型
type BindType int8

const (
	BtDynamic BindType = iota // 动态类型
	BtStatic                  // 静态类型
)

// JSValue
//
// GO和JS动态变量类型
//
// 在主进程有效
type JSValue interface {
	Name() string             //当前变量绑定的名称
	JSON() json.JSON          //
	Value() any               //
	SetValue(value any)       //设置新值
	IsInteger() bool          //是否 Integer
	IsDouble() bool           //是否 Double
	IsString() bool           //是否 String
	IsBoolean() bool          //是否 Boolean
	IsObject() bool           //是否 Object
	IsArray() bool            //是否 Array
	IsUndefined() bool        //是否 Undefined
	IsFunction() bool         //是否 Function
	AsInteger() JSInteger     //转换为 Integer 失败返回 nil
	AsDouble() JSDouble       //转换为 Double 失败返回 nil
	AsString() JSString       //转换为 String 失败返回 nil
	AsBoolean() JSBoolean     //转换为 Boolean 失败返回 nil
	AsUndefined() JSUndefined //转换为 Undefined 失败返回 nil
	AsFunction() JSFunction   //转换为 Function 失败返回 nil
	//AsArray() JSArray

	AsV8Value() JSValue //转换为 JSValue
	StringValue() string
	BooleanValue() bool
	DoubleValue() float64
	IntegerValue() int
	UndefinedValue() string
	// Invoke 调用函数
	//	入参: 以参数列表形式传入参数
	//		入参如果参数类型或数量不匹配这些参数将以类型的默认值传入
	//		nil 无入参
	//	出参: 以参数列表形式返回参数
	//		无返回值返回nil
	Invoke(argumentList json.JSONArray) (resultArgument json.JSONArray)
	Id() int            //指针ID
	Type() reflect.Kind //值类型
	BindType() BindType //绑定类型
	setId(id uintptr)
	free()
}

// V8Value 绑定到JS的字段
type V8Value struct {
	*json.JsonData
	id   uintptr
	name string
	bt   BindType
}

//BindType 绑定类型
func (m *V8Value) BindType() BindType {
	return m.bt
}

func (m *V8Value) free() {
	bind.Remove(uintptr(m.Id()))
	m.id = 0
	m.name = ""
}

// nameKey
func (m *V8Value) nameKey(pName, name string) {
	var concat = strings.Builder{}
	concat.WriteString(pName)
	concat.WriteString(".")
	concat.WriteString(name)
	m.name = concat.String()
}

// createJSValue 创建 JSValue
func (m *V8Value) createJSValue(name string, rv *reflect.Value) JSValue {
	kind := rv.Kind()
	if kind == reflect.Ptr {
		kind = rv.Elem().Kind()
	}
	switch kind {
	case reflect.String:
		v := new(jsString)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.String, rv.Len(), rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := new(jsInteger)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Int, strconv.IntSize, rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := new(jsInteger)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Uint, strconv.IntSize, rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Float32, reflect.Float64:
		v := new(jsDouble)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Float64, 8, rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Bool:
		v := new(jsBoolean)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Bool, 1, rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Struct:
		v := new(jsObject)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Map, 0, rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Map:
		v := new(jsObject)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Map, rv.Len(), rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Slice:
		v := new(jsArray)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Slice, rv.Len(), rv)
		v.bt = m.bt
		bind.Set(v)
		return v
	case reflect.Func:
		v := new(jsFunction)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(reflect.Func, 0, rv)
		v.rv = rv
		v.bt = m.bt
		bind.Set(v)
		return v
	default:
		v := new(jsUndefined)
		v.nameKey(m.name, name)
		v.JsonData = json.NewJsonData(consts.NIL, 0, null)
		v.bt = m.bt
		bind.Set(v)
		return v
	}
}

func (m *V8Value) JSON() json.JSON {
	return m.JsonData
}

func (m *V8Value) Value() any {
	if m.JsonData != nil {
		return m.JsonData.Data()
	}
	return nil
}

// SetValue 设置值
func (m *V8Value) SetValue(value any) {
	if isMainProcess {
		if m.bt == BtDynamic {
			m.JsonData.SetValue(value)
		} else if m.bt == BtStatic {
			if m.JsonData.Data() != nil {
				rv := reflect.ValueOf(value)
				kind := rv.Kind()
				if kind == reflect.Ptr {
					kind = rv.Elem().Kind()
				}
				m.JsonData.Data().(*reflect.Value).Set(rv)
			}
		}

	}
}

func (m *V8Value) setId(id uintptr) {
	m.id = id
}

func (m *V8Value) Id() int {
	return int(m.id)
}

func (m *V8Value) Name() string {
	return m.name
}

func (m *V8Value) IsString() bool {
	return m.JsonData.IsString()
}

func (m *V8Value) IsInteger() bool {
	return m.JsonData.IsInt() || m.JsonData.IsUInt()
}

func (m *V8Value) IsDouble() bool {
	return m.JsonData.IsFloat()
}

func (m *V8Value) IsBoolean() bool {
	return m.JsonData.IsBool()
}

func (m *V8Value) IsArray() bool {
	return m.JsonData.IsArray()
}

func (m *V8Value) IsObject() bool {
	return m.JsonData.IsObject()
}

func (m *V8Value) IsFunction() bool {
	return m.JsonData.Type() == reflect.Func
}

func (m *V8Value) IsUndefined() bool {
	return m.JsonData.Type() == consts.NIL
}

func (m *V8Value) AsString() JSString {
	if m.IsString() {
		return m
	}
	return nil // default
}

func (m *V8Value) StringValue() string {
	return m.JsonData.String()
}

func (m *V8Value) AsInteger() JSInteger {
	if m.IsInteger() {
		return m
	}
	return nil // default
}

func (m *V8Value) IntegerValue() int {
	return m.JsonData.Int()
}

func (m *V8Value) AsDouble() JSDouble {
	if m.IsDouble() {
		return m
	}
	return nil // default
}

func (m *V8Value) DoubleValue() float64 {
	return m.JsonData.Float()
}

func (m *V8Value) AsBoolean() JSBoolean {
	if m.IsBoolean() {
		return m
	}
	return nil // default
}

func (m *V8Value) BooleanValue() bool {
	return m.JsonData.Bool()
}

func (m *V8Value) AsUndefined() JSUndefined {
	if m.IsUndefined() {
		return m
	}
	return nil // default
}

func (m *V8Value) UndefinedValue() string {
	return m.JsonData.String()
}

func (m *V8Value) AsFunction() JSFunction {
	if m.IsFunction() {
		return m
	}
	return nil // default
}

func (m *V8Value) Invoke(argumentList json.JSONArray) (resultArgument json.JSONArray) {
	if f := m.AsFunction(); f != nil {
		return f.Invoke(argumentList)
	}
	return nil
}

//func (m *V8Value) AsArray() JSArray {
//	if m.IsArray() {
//		v := new(jsArray)
//		v.name = m.name
//		v.value = m.value.JsonData()
//		v.bt = m.bt
//		return v
//	}
//	return nil // default
//}

func (m *V8Value) AsV8Value() JSValue {
	return m
}

func (m *V8Value) Type() reflect.Kind {
	if m == nil {
		return reflect.Invalid
	}
	return m.JsonData.Type()
}

// NewInteger GO&JS 数字类型
func NewInteger(name string, value int) JSInteger {
	if name == "" || (!isMainProcess && !isSubProcess) {
		return nil
	}
	v := new(jsInteger)
	v.name = name
	v.JsonData = json.NewJsonData(reflect.Int, 4, value)
	v.bt = BtDynamic
	bind.Set(v)
	return v
}

// NewString GO&JS 字符串类型
func NewString(name, value string) JSString {
	if name == "" || (!isMainProcess && !isSubProcess) {
		return nil
	}
	v := new(jsString)
	v.name = name
	v.JsonData = json.NewJsonData(reflect.String, len(value), value)
	v.bt = BtDynamic
	bind.Set(v)
	return v
}

// NewDouble GO&JS 浮点类型
func NewDouble(name string, value float64) JSDouble {
	if name == "" || (!isMainProcess && !isSubProcess) {
		return nil
	}
	v := new(jsDouble)
	v.name = name
	v.JsonData = json.NewJsonData(reflect.Float64, 8, value)
	v.bt = BtDynamic
	bind.Set(v)
	return v
}

// NewBoolean  GO&JS 布尔类型
func NewBoolean(name string, value bool) JSBoolean {
	if name == "" || (!isMainProcess && !isSubProcess) {
		return nil
	}
	v := new(jsBoolean)
	v.name = name
	v.JsonData = json.NewJsonData(reflect.Bool, 1, value)
	v.bt = BtDynamic
	bind.Set(v)
	return v
}

// NewUndefined GO&JS 未定义类型
func NewUndefined(name string) JSUndefined {
	if name == "" || (!isMainProcess && !isSubProcess) {
		return nil
	}
	v := new(jsUndefined)
	v.name = name
	v.JsonData = json.NewJsonData(consts.NIL, 0, undefined)
	v.bt = BtDynamic
	bind.Set(v)
	return v
}

// NewFunction GO&JS 函数类型
func NewFunction(name string, fn any) JSFunction {
	if name == "" || (!isMainProcess && !isSubProcess) {
		return nil
	}
	rv := reflect.ValueOf(fn)
	if rv.Kind() != reflect.Func {
		return nil
	}
	v := new(jsFunction)
	v.name = name
	v.JsonData = json.NewJsonData(reflect.Func, 0, &rv)
	v.rv = &rv
	v.bt = BtStatic
	bind.Set(v)
	return v
}

// NewObject GO&JS 对象类型 &struct{} 仅结构
func NewObject(object any) JSObject {
	if object == nil || (!isMainProcess && !isSubProcess) {
		return nil
	}
	rv := reflect.ValueOf(object)
	kind := rv.Kind()
	//必须是指针
	if kind != reflect.Ptr {
		return nil
	}
	kind = rv.Elem().Kind()
	//必须是结构
	if kind != reflect.Struct {
		return nil
	}
	v := new(jsObject)
	v.name = rv.Type().Elem().Name()
	v.JsonData = json.NewJsonData(reflect.Map, 0, &rv)
	v.bt = BtStatic
	bind.Set(v)
	v.fieldToBind()
	return v
}

// NewArray GO&JS 数组类型
func NewArray(name string, values []any) JSArray {
	if name == "" || values == nil || (!isMainProcess && !isSubProcess) {
		return nil
	}
	v := new(jsArray)
	v.name = name
	v.JsonData = json.NewJsonData(reflect.Slice, len(values), values)
	v.bt = BtStatic
	bind.Set(v)
	for _, value := range values {
		v.Add(value)
	}
	return v
}
