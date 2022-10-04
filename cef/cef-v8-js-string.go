//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

type JSString struct {
	ICEFv8Value
}

func (m *JSString) AsString() (*JSString, error) {
	return m, nil
}

func (m *JSString) Value() string {
	if val, err := m.StringValue(); err == nil {
		return val
	} else {
		return ""
	}
}

func (m *JSString) SetValue(value string) {
	m.valueType = V8_VALUE_STRING
	m.value = value
}
func (m *JSString) ToString() string {
	if val, err := m.StringValue(); err == nil {
		return val
	}
	return ""
}
