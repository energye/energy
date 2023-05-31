//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSUndefined 实现

package bind

type JSUndefined interface {
	JSValue
	AsUndefined() JSUndefined
	UndefinedValue() string
}

type jsUndefined struct {
	V8Value
}

func (m *jsUndefined) AsUndefined() JSUndefined {
	if m.IsUndefined() {
		return m
	}
	return nil
}

func (m *jsUndefined) UndefinedValue() string {
	return m.JsonData.String()
}
