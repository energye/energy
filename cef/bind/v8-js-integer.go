//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSInteger 类型实现
package bind

type JSInteger struct {
	V8Value
}

func (m *JSInteger) AsInteger() *JSInteger {
	return m
}

func (m *JSInteger) Value() int32 {
	return m.IntegerValue()
}
func (m *JSInteger) SetValue(value int32) {

}
