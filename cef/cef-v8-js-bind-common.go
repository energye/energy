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
// 绑定 定义的普通函数 前缀-默认是结构名称, prefix default struct name
package cef

// NewString GO&JS 类型
func (m *ICEFv8Value) NewString(name, value string) *JSString {
	return VariableBind.bindStorage.NewString(name, value)
}

// NewInteger GO&JS 类型
func (m *ICEFv8Value) NewInteger(name string, value int32) *JSInteger {
	return VariableBind.bindStorage.NewInteger(name, value)
}

// NewDouble GO&JS 类型
func (m *ICEFv8Value) NewDouble(name string, value float64) *JSDouble {
	return VariableBind.bindStorage.NewDouble(name, value)
}

// NewBoolean  GO&JS 类型
func (m *ICEFv8Value) NewBoolean(name string, value bool) *JSBoolean {
	return VariableBind.bindStorage.NewBoolean(name, value)
}

// NewNull GO&JS 类型
func (m *ICEFv8Value) NewNull(name string) *JSNull {
	return VariableBind.bindStorage.NewNull(name)
}

// NewUndefined GO&JS 类型
func (m *ICEFv8Value) NewUndefined(name string) *JSUndefined {
	return VariableBind.bindStorage.NewUndefined(name)
}

// NewFunction GO&JS 类型
func (m *ICEFv8Value) NewFunction(name string, fn interface{}) error {
	return VariableBind.bindStorage.NewFunction(name, fn)
}

//	NewObjects GO&JS 类型
//
// Go结构类型变量和Go结构函数绑定
func (m *ICEFv8Value) NewObjects(objects ...interface{}) {
	bindObject(objects...)
}
