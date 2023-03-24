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
	"github.com/energye/energy/pkgs/json"
)

// JSValue
//
// GO和JS变量类型接口
type JSValue interface {
	SetValue(value any) error //多类型值设置
	IntegerValue() int32      //返回 Integer 类型值, 成功 error = nil
	DoubleValue() float64     //返回 Double 类型值, 成功 error = nil
	StringValue() string      //返回 String 类型值, 成功 error = nil
	BooleanValue() bool       //返回 Boolean 类型值, 成功 error = nil
	IsInteger() bool          //是否 Integer
	IsDouble() bool           //是否 Double
	IsString() bool           //是否 String
	IsBool() bool             //是否 Bool
	IsArray() bool            //是否 JSONArray
	IsObject() bool           //是否 Object
	IsFunction() bool         //是否 Function
	IsNull() bool             //是否 Null
	IsUndefined() bool        //是否 Undefined
	AsV8Value() *V8Value      //转换为 V8Value
	AsInteger() *JSInteger    //转换为 Integer 失败返回 nil
	AsDouble() *JSDouble      //转换为 Double 失败返回 nil
	AsString() *JSString      //转换为 String 失败返回 nil
	AsBoolean() *JSBoolean    //转换为 Boolean 失败返回 nil
	AsArray() *JSArray        //转换为 JSONArray 失败返回 nil
	AsFunction() *JSFunction  //转换为 Function 失败返回 nil
	Name() string             //当前变量绑定的名称
	Bytes() []byte            //变量值转换为字节
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

// SetAnyValue 设置多类型值
func (m *V8Value) SetValue(value any) error {

	return nil
}

func (m *V8Value) Name() string {
	return m.name
}

func (m *V8Value) StringValue() string {
	if m.IsString() {
		return m.value.String()
	}
	return ""
}

func (m *V8Value) IntegerValue() int32 {
	if m.IsInteger() {
		return int32(m.value.Int())
	}
	return 0
}

func (m *V8Value) DoubleValue() float64 {
	if m.IsDouble() {
		return m.value.Float()
	}
	return 0
}

func (m *V8Value) BooleanValue() bool {
	if m.IsBool() {
		return m.value.Bool()
	}
	return false
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

func (m *V8Value) IsBool() bool {
	return m.value.IsBool()
}

func (m *V8Value) IsArray() bool {
	return m.value.IsArray()
}

func (m *V8Value) IsObject() bool {
	return m.value.IsObject()
}

func (m *V8Value) IsFunction() bool {
	//return m.value
	return false
}

func (m *V8Value) IsNull() bool {
	return m.value == nil
}

func (m *V8Value) IsUndefined() bool {
	return m.value == nil
}

func (m *V8Value) AsString() *JSString {

	return nil
}

func (m *V8Value) AsInteger() *JSInteger {

	return nil
}

func (m *V8Value) AsDouble() *JSDouble {

	return nil
}

func (m *V8Value) AsBoolean() *JSBoolean {

	return nil
}

func (m *V8Value) AsV8Value() *V8Value {
	return m
}

func (m *V8Value) AsArray() *JSArray {

	return nil
}

func (m *V8Value) AsObject() *JSObject {

	return nil
}

func (m *V8Value) AsFunction() *JSFunction {

	return nil
}

// NewString GO&JS 类型
func NewString(name, value string) *JSString {
	return nil
}

// NewInteger GO&JS 类型
func NewInteger(name string, value int32) *JSInteger {
	return nil
}

// NewDouble GO&JS 类型
func NewDouble(name string, value float64) *JSDouble {
	return nil
}

// NewBoolean  GO&JS 类型
func NewBoolean(name string, value bool) *JSBoolean {
	return nil
}

// NewNull GO&JS 类型
func NewNull(name string) *JSNull {
	return nil
}

// NewUndefined GO&JS 类型
func NewUndefined(name string) *JSUndefined {
	return nil
}

// NewFunction GO&JS 类型
func NewFunction(name string, fn interface{}) error {
	return nil
}

//	NewObjects GO&JS 类型
//
// Go结构类型变量和Go结构函数绑定
func NewObjects(objects ...interface{}) {
}

//	NewObjects GO&JS 类型
//
// Go结构类型变量和Go结构函数绑定
func NewArrays(objects ...interface{}) {
}

// Bind V8Value
//
// 变量和函数绑定, 在Go中定义的字段绑定到JS字段中, 在Go中定义的函数导出到JS
//
// 支持类型 String = string , Integer = int32 , Double = float64, Boolean = bool, Function = func, Objects = struct | map,  JSONArray = Slice
//
// 主进程和子进程
func Bind(name string, bind interface{}) error {
	return nil
}
