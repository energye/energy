//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSArray 类型实现
package bind

type JSArray interface {
	JSValue
	AsArray() JSArray
	Value() any
}

// JSArray 类型 先保留 未添加
type jsArray struct {
	V8Value
}

func (m *jsArray) AsArray() JSArray {
	if m.IsArray() {
		return m
	}
	return nil
}

func (m *jsArray) Value() any {
	if m.IsArray() {
		return m.value.Data()
	}
	return nil
}
