//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"github.com/energye/energy/consts"
)

type JSInteger struct {
	ICEFv8Value
}

func (m *JSInteger) AsInteger() (*JSInteger, error) {
	return m, nil
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
	m.valueType = consts.V8_VALUE_INT
	m.value = value
}

func (m *JSInteger) ToString() string {
	if val, err := m.IntegerValue(); err == nil {
		return fmt.Sprintf("%v", val)
	}
	return ""
}
