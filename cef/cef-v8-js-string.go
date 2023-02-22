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
package cef

import "github.com/energye/energy/consts"

type JSString struct {
	V8Value
}

func (m *JSString) AsString() *JSString {
	return m
}

func (m *JSString) Value() string {
	if val, err := m.StringValue(); err == nil {
		return val
	} else {
		return ""
	}
}

func (m *JSString) SetValue(value string) {
	m.valueType.Jsv = consts.V8_VALUE_STRING
	m.valueType.Gov = consts.GO_VALUE_STRING
	m.value = value
}
func (m *JSString) ToString() string {
	if val, err := m.StringValue(); err == nil {
		return val
	}
	return ""
}
