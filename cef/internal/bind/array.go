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
	fieldToBind()
	Add(value any) JSValue
	Set(index int, value any) JSValue
	Clear()
	Remove(index int)
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

func (m *jsArray) fieldToBind() {
	for _, item := range m.Items() {
		if item.IsObject() {
			item.(JSObject).fieldToBind()
		}
	}
}

func (m *jsArray) Add(value any) JSValue {
	item := m.createItem(len(m.items), value)
	m.items = append(m.items, item)
	if item.IsObject() {
		item.(JSObject).fieldToBind()
	} else if item.IsArray() {
		array := item.(JSArray)
		arrayRv := reflect.ValueOf(value)
		for i := 0; i < arrayRv.Len(); i++ {
			rv := arrayRv.Index(i)
			array.Add(rv)
		}
	}
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
	for _, v := range m.items {
		v.free()
	}
	m.items = make([]JSValue, 0, 0)
}

func (m *jsArray) Remove(index int) {
	if v := m.Get(index); v != nil {
		v.free()
		m.items = append(m.items[:index], m.items[index+1:]...)
	}
}

func (m *jsArray) createItem(index int, value any) JSValue {
	var rv reflect.Value
	switch value.(type) {
	case reflect.Value:
		rv = value.(reflect.Value)
	default:
		rv = reflect.ValueOf(value)
	}
	return m.createJSValue(strconv.Itoa(index), &rv)
}