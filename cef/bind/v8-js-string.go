//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSString 实现
package bind

type JSString interface {
	JSValue
	AsString() JSString
	Value() string
}

type jsString struct {
	V8Value
}

func (m *jsString) AsString() JSString {
	if m.IsString() {
		return m
	}
	return nil
}

func (m *jsString) Value() string {
	return m.value.String()
}
