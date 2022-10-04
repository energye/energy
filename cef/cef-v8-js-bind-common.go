//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

//ICEFv8Value NewString
func (m *ICEFv8Value) NewString(name, value string) *JSString {
	return VariableBind.bindStorage.NewString(name, value)
}

//ICEFv8Value NewInteger
func (m *ICEFv8Value) NewInteger(name string, value int32) *JSInteger {
	return VariableBind.bindStorage.NewInteger(name, value)
}

//ICEFv8Value NewDouble
func (m *ICEFv8Value) NewDouble(name string, value float64) *JSDouble {
	return VariableBind.bindStorage.NewDouble(name, value)
}

//ICEFv8Value NewBool
func (m *ICEFv8Value) NewBoolean(name string, value bool) *JSBoolean {
	return VariableBind.bindStorage.NewBoolean(name, value)
}

//ICEFv8Value NewNull
func (m *ICEFv8Value) NewNull(name string) *JSNull {
	return VariableBind.bindStorage.NewNull(name)
}

//ICEFv8Value NewUndefined
func (m *ICEFv8Value) NewUndefined(name string) *JSUndefined {
	return VariableBind.bindStorage.NewUndefined(name)
}

//ICEFv8Value NewFunction
//绑定 定义的普通函数 prefix default struct name
func (m *ICEFv8Value) NewFunction(name string, fn interface{}) error {
	return VariableBind.bindStorage.NewFunction(name, fn)
}

//ICEFv8Value NewObjects
//对象类型变量和对象函数绑定
func (m *ICEFv8Value) NewObjects(objects ...interface{}) {
	bindObject(objects...)
}
