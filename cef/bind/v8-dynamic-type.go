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

const (
	null      = "null"
	undefined = "Undefined"
)

// JSValue
//
// GO和JS动态变量类型
type JSValue interface {
	Name() string             //当前变量绑定的名称
	Bytes() []byte            //变量值转换为字节
	SetValue(value any)       //设置新值
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
	value json.JSON
}

// Bytes 值转换为字节
func (m *V8Value) Bytes() []byte {
	return m.value.Bytes()
}

// SetValue 设置值
func (m *V8Value) SetValue(value any) {

}

func (m *V8Value) Name() string {
	return m.name
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
	return m.value.Type() == consts.GO_VALUE_FUNC
}

func (m *V8Value) IsNull() bool {
	return m.value.Type() == consts.GO_VALUE_NIL
}

func (m *V8Value) IsUndefined() bool {
	return m.value.Type() == consts.GO_VALUE_NIL
}

func (m *V8Value) AsString() JSString {
	if m.IsString() {
		v := new(jsString)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsInteger() JSInteger {
	if m.IsInteger() {
		v := new(jsInteger)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsDouble() JSDouble {
	if m.IsDouble() {
		v := new(jsDouble)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsBoolean() JSBoolean {
	if m.IsBoolean() {
		v := new(jsBoolean)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsNull() JSNull {
	if m.IsNull() {
		v := new(jsNull)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsUndefined() JSUndefined {
	if m.IsUndefined() {
		v := new(jsUndefined)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsFunction() JSFunction {
	if m.IsFunction() {
		v := new(jsFunction)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsObject() JSObject {
	if m.IsObject() {
		v := new(jsObject)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsArray() JSArray {
	if m.IsArray() {
		v := new(jsArray)
		v.name = m.name
		v.value = m.value.JsonData()
		bind.set(m.name, v)
		m.free()
		return v
	}
	return nil // default
}

func (m *V8Value) AsV8Value() JSValue {
	return m
}

func (m *V8Value) free() {
	if m.value != nil {
		m.name = ""
		m.value.Free()
		m.value = nil
	}
}

// NewInteger GO&JS 数字类型
func NewInteger(name string, value int) JSInteger {
	if name == "" {
		return nil
	}
	v := new(jsInteger)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_INT, V: value, S: 4}
	bind.set(name, v)
	return v
}

// NewString GO&JS 字符串类型
func NewString(name, value string) JSString {
	if name == "" {
		return nil
	}
	v := new(jsString)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_STRING, V: value, S: len(value)}
	bind.set(name, v)
	return v
}

// NewDouble GO&JS 浮点类型
func NewDouble(name string, value float64) JSDouble {
	if name == "" {
		return nil
	}
	v := new(jsDouble)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_FLOAT64, V: value, S: 8}
	bind.set(name, v)
	return v
}

// NewBoolean  GO&JS 布尔类型
func NewBoolean(name string, value bool) JSBoolean {
	if name == "" {
		return nil
	}
	v := new(jsBoolean)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_BOOL, V: value, S: 1}
	bind.set(name, v)
	return v
}

// NewNull GO&JS 空类型
func NewNull(name string) JSNull {
	if name == "" {
		return nil
	}
	v := new(jsNull)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_NIL, V: null, S: len(null)}
	bind.set(name, v)
	return v
}

// NewUndefined GO&JS 未定义类型
func NewUndefined(name string) JSUndefined {
	if name == "" {
		return nil
	}
	v := new(jsUndefined)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_NIL, V: undefined, S: len(undefined)}
	bind.set(name, v)
	return v
}

// NewFunction GO&JS 函数类型
func NewFunction(name string, fn any) JSFunction {
	if name == "" {
		return nil
	}
	v := new(jsFunction)
	v.name = name
	v.value = &json.JsonData{T: consts.GO_VALUE_FUNC, V: fn, S: 0}
	bind.set(name, v)
	return v
}

// NewObject GO&JS 对象类型
func NewObject(name string, object any) JSObject {
	if name == "" {
		return nil
	}
	if vv := json.NewJSONObject(object); vv != nil {
		v := new(jsObject)
		v.name = name
		v.value = vv.JsonData()
		bind.set(name, v)
		return v
	}
	return nil
}

// NewArray GO&JS 数组类型
func NewArray(name string, array any) JSArray {
	if name == "" {
		return nil
	}
	if vv := json.NewJSONArray(array); vv != nil {
		v := new(jsArray)
		v.name = name
		v.value = vv.JsonData()
		bind.set(name, v)
		return v
	}
	return nil
}
