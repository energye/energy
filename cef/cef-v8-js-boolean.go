//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import "github.com/energye/energy/consts"

type JSBoolean struct {
	ICEFv8Value
}

func (m *JSBoolean) AsBoolean() (*JSBoolean, error) {
	return m, nil
}

func (m *JSBoolean) Value() bool {
	if val, err := m.BooleanValue(); err == nil {
		return val
	}
	return false
}

func (m *JSBoolean) SetValue(value bool) {
	m.valueType = consts.V8_VALUE_BOOLEAN
	m.value = value
}

func (m *JSBoolean) ToString() string {
	if val, err := m.BooleanValue(); err == nil {
		if val {
			return "true"
		}
		return "false"
	}
	return ""
}
