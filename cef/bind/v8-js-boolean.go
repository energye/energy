//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSBoolean 类型实现
package bind

type JSBoolean struct {
	V8Value
}

func (m *JSBoolean) AsBoolean() *JSBoolean {
	return m
}

func (m *JSBoolean) Value() bool {
	return m.BooleanValue()
}

func (m *JSBoolean) SetValue(value bool) {
}
