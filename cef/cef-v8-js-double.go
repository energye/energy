//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSDouble 类型实现
package cef

import (
	"fmt"
	"github.com/energye/energy/consts"
)

type JSDouble struct {
	ICEFv8Value
}

func (m *JSDouble) AsDouble() (*JSDouble, error) {
	return m, nil
}

func (m *JSDouble) Value() float64 {
	if val, err := m.DoubleValue(); err == nil {
		return val
	} else {
		return 0.0
	}
}

func (m *JSDouble) SetValue(value float64) {
	m.valueType = consts.V8_VALUE_DOUBLE
	m.value = value
}
func (m *JSDouble) ToString() string {
	if val, err := m.DoubleValue(); err == nil {
		return fmt.Sprintf("%v", val)
	}
	return ""
}
