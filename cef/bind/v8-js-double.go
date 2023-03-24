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
package bind

import (
	"fmt"
	"github.com/energye/energy/consts"
)

type JSDouble struct {
	V8Value
}

func (m *JSDouble) AsDouble() *JSDouble {
	return m
}

func (m *JSDouble) Value() float64 {
	if val, err := m.DoubleValue(); err == nil {
		return val
	} else {
		return 0.0
	}
}

func (m *JSDouble) SetValue(value float64) {
	m.valueType.Jsv = consts.V8_VALUE_DOUBLE
	m.valueType.Gov = consts.GO_VALUE_FLOAT64
	m.value = value
}
func (m *JSDouble) ToString() string {
	if val, err := m.DoubleValue(); err == nil {
		return fmt.Sprintf("%v", val)
	}
	return ""
}
