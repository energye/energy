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

type JSInteger interface {
	JSValue
	AsInteger() JSInteger
	Value() int
}

type jsInteger struct {
	V8Value
}

func (m *jsInteger) AsInteger() JSInteger {
	if m.IsInteger() {
		return m
	}
	return nil
}

func (m *jsInteger) Value() int {
	return m.value.Int()
}
