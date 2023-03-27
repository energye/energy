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

type JSBoolean interface {
	JSValue
	AsBoolean() JSBoolean
	Value() bool
}

type jsBoolean struct {
	V8Value
}

func (m *jsBoolean) AsBoolean() JSBoolean {
	if m.IsBoolean() {
		return m
	}
	return nil
}

func (m *jsBoolean) Value() bool {
	return m.value.Bool()
}
