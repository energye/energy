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

import (
	"fmt"
	"github.com/energye/energy/consts"
)

type JSInteger struct {
	V8Value
}

func (m *JSInteger) AsInteger() *JSInteger {
	return m
}

func (m *JSInteger) Float() float64 {
	if val, err := m.IntegerValue(); err == nil {
		return float64(val)
	} else {
		return 0.0
	}
}

func (m *JSInteger) Value() int32 {
	if val, err := m.IntegerValue(); err == nil {
		return val
	} else {
		return 0
	}
}
func (m *JSInteger) SetValue(value int32) {
	m.valueType.Jsv = consts.V8_VALUE_INT
	m.valueType.Gov = consts.GO_VALUE_INT32
	m.value = value
}

func (m *JSInteger) ToString() string {
	if val, err := m.IntegerValue(); err == nil {
		return fmt.Sprintf("%v", val)
	}
	return ""
}
