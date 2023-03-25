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
package bind

import (
	"github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/json"
)

// JSValue
//
// GO和JS通用变量类型
type JSValue interface {
	Name() string             //当前变量绑定的名称
	Bytes() []byte            //变量值转换为字节
	SetValue(value any)       //设置新值
	Integer() int             //返回 Integer 类型
	Double() float64          //返回 Double 类型
	String() string           //返回 String 类型
	Boolean() bool            //返回 Boolean 类型
	Object() JSObject         //返回 Object 类型
	Array() JSArray           //返回 Array 类型
	Null() JSNull             //返回 Null 类型
	Undefined() JSUndefined   //返回 Undefined
	Function() JSFunction     //返回 Function
	IsInteger() bool          //是否 Integer
	IsDouble() bool           //是否 Double
	IsString() bool           //是否 String
	IsBoolean() bool          //是否 Boolean
	IsObject() bool           //是否 Object
	IsArray() bool            //是否 Array
	IsNull() bool             //是否 Null
	IsUndefined() bool        //是否 Undefined
	IsFunction() bool         //是否 Function
	AsInteger() JSInteger     //转换为 Integer 失败返回 nil
	AsDouble() JSDouble       //转换为 Double 失败返回 nil
	AsString() JSString       //转换为 String 失败返回 nil
	AsBoolean() JSBoolean     //转换为 Boolean 失败返回 nil
	AsObject() JSObject       //转换为 Object 失败返回 nil
	AsArray() JSArray         //转换为 Array 失败返回 nil
	AsNull() JSNull           //转换为 Null 失败返回 nil
	AsUndefined() JSUndefined //转换为 Undefined 失败返回 nil
	AsFunction() JSFunction   //转换为 Function 失败返回 nil
	AsV8Value() JSValue       //转换为 JSValue
}

// V8Value 绑定到JS的字段
type V8Value struct {
	name  string
	value *json.JsonData
}

// Bytes 值转换为字节
func (m *V8Value) Bytes() []byte {
	return nil
}

// SetValue 设置值
func (m *V8Value) SetValue(value any) {

}

func (m *V8Value) Name() string {
	return m.name
}

func (m *V8Value) String() string {
	if m.IsString() {
		return m.value.String()
	}
	return ""
}

func (m *V8Value) Integer() int {
	return m.value.Int()
}

func (m *V8Value) Double() float64 {
	return m.value.Float()
}

func (m *V8Value) Boolean() bool {
	return m.value.Bool()
}

func (m *V8Value) Object() JSObject {
	if m.IsArray() {
	}
	return nil
}

func (m *V8Value) Array() JSArray {
	if m.IsArray() {
	}
	return nil
}

func (m *V8Value) Null() JSNull {

	return nil
}

func (m *V8Value) Undefined() JSUndefined {

	return nil
}
func (m *V8Value) Function() JSFunction {

	return nil
}

func (m *V8Value) IsString() bool {
	return m.value.IsString()
}

func (m *V8Value) IsInteger() bool {
	return m.value.IsInt() || m.value.IsUInt()
}

func (m *V8Value) IsDouble() bool {
	return m.value.IsFloat()
}

func (m *V8Value) IsBoolean() bool {
	return m.value.IsBool()
}

func (m *V8Value) IsArray() bool {
	return m.value.IsArray()
}

func (m *V8Value) IsObject() bool {
	return m.value.IsObject()
}

func (m *V8Value) IsFunction() bool {
	return false
}

func (m *V8Value) IsNull() bool {
	return m.value == nil
}

func (m *V8Value) IsUndefined() bool {
	return m.value == nil
}

func (m *V8Value) AsString() JSString {
	return nil
}

func (m *V8Value) AsInteger() JSInteger {
	return nil
}

func (m *V8Value) AsDouble() JSDouble {
	return nil
}

func (m *V8Value) AsBoolean() JSBoolean {
	return nil
}

func (m *V8Value) AsNull() JSNull {
	return nil
}

func (m *V8Value) AsUndefined() JSUndefined {
	return nil
}

func (m *V8Value) AsFunction() JSFunction {
	return nil
}

func (m *V8Value) AsObject() JSObject {
	return nil
}

func (m *V8Value) AsArray() JSArray {
	return nil
}

func (m *V8Value) AsV8Value() JSValue {
	return m
}

// NewInteger GO&JS 数字类型
func NewInteger(name string, value int) JSInteger {
	v := new(jsInteger)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_INT, V: value, S: 4}
	bind.Add(name, v)
	return v
}

// NewString GO&JS 字符串类型
func NewString(name, value string) JSString {
	v := new(jsString)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_STRING, V: value, S: len(value)}
	bind.Add(name, v)
	return v
}

// NewDouble GO&JS 浮点类型
func NewDouble(name string, value float64) JSDouble {
	v := new(jsDouble)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_FLOAT64, V: value, S: 8}
	bind.Add(name, v)
	return v
}

// NewBoolean  GO&JS 布尔类型
func NewBoolean(name string, value bool) JSBoolean {
	v := new(jsBoolean)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_BOOL, V: value, S: 1}
	bind.Add(name, v)
	return v
}

// NewNull GO&JS 空类型
func NewNull(name string) JSNull {
	v := new(jsNull)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_NIL, V: "null", S: 4}
	bind.Add(name, v)
	return v
}

// NewUndefined GO&JS 未定义类型
func NewUndefined(name string) JSUndefined {
	v := new(jsUndefined)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_NIL, V: "undefined", S: 9}
	bind.Add(name, v)
	return v
}

// NewFunction GO&JS 函数类型
func NewFunction(name string, fn any) JSFunction {
	v := new(jsFunction)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_FUNC, V: fn, S: 0}
	bind.Add(name, v)
	return v
}

// NewObject GO&JS 对象类型
func NewObject(name string, object any) JSObject {
	if vv := json.NewJSONObject(object); vv != nil {
		v := new(jsObject)
		v.name = name
		v.value = vv.JsonData()
		bind.Add(name, v)
		return v
	}
	return nil
}

// NewArray GO&JS 数组类型
func NewArray(name string, array any) JSArray {
	if vv := json.NewJSONArray(array); vv != nil {
		v := new(jsArray)
		v.name = name
		v.value = vv.JsonData()
		bind.Add(name, v)
		return v
	}
	return nil
}
