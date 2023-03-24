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

type JSString struct {
	V8Value
}

func (m *JSString) AsString() *JSString {
	return m
}

func (m *JSString) Value() string {
	return m.StringValue()
}

func (m *JSString) SetValue(value string) {

}
