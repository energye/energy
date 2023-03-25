//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSObject 实现
package bind

import "github.com/energye/energy/pkgs/json"

type JSObject interface {
	JSValue
	AsObject() JSObject
	Value() json.JSONObject
}

type jsObject struct {
	V8Value
}

func (m *jsObject) AsObject() JSObject {
	return m
}

func (m *jsObject) Value() json.JSONObject {
	return m.value.JSONObject()
}
