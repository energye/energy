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

import (
	"reflect"
	"strconv"
)

type JSArray interface {
	JSValue
	AsArray() JSArray
	Items() []JSValue
	Get(index int) JSValue
	Add(value any) JSValue
	Set(index int, value any) JSValue
	Clear()
}

// JSArray 类型 先保留 未添加
type jsArray struct {
	V8Value
	items []JSValue
}

func (m *jsArray) AsArray() JSArray {
	if m.IsArray() {
		return m
	}
	return nil
}

func (m *jsArray) Items() []JSValue {
	return m.items
}

func (m *jsArray) Get(index int) JSValue {
	if index >= 0 && index < len(m.items) {
		return m.items[index]
	}
	return nil
}

func (m *jsArray) Add(value any) JSValue {
	item := m.createItem(len(m.items), value)
	m.items = append(m.items, item)
	return item
}

func (m *jsArray) Set(index int, value any) JSValue {
	if index >= 0 && index < len(m.items) {
		if value == nil {
			m.items[index] = nil
		} else {
			m.items[index] = m.createItem(index, value)
		}
		return m.items[index]
	}
	return nil
}

func (m *jsArray) Clear() {
	m.items = make([]JSValue, 0, 0)
}

func (m *jsArray) createItem(index int, value any) JSValue {
	rv := reflect.ValueOf(value)
	return m.createJSValue(strconv.Itoa(index), &rv)
}
