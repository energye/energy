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

import "github.com/energye/energy/pkgs/json"

type JSArray interface {
	JSValue
	AsArray() JSArray
	Value() json.JSONArray
}

// JSArray 类型 先保留 未添加
type jsArray struct {
	V8Value
}

func (m *jsArray) AsArray() JSArray {
	return m
}

func (m *jsArray) Value() json.JSONArray {
	return m.value.JSONArray()
}
