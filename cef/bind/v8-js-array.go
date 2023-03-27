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
	"github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/json"
	"reflect"
	"strconv"
)

type JSArray interface {
	JSValue
	AsArray() JSArray
	Data() []JSValue
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

func (m *jsArray) Data() []JSValue {
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
	switch value.(type) {
	case string:
		v := new(jsString)
		v.pName = m.name
		v.name = strconv.Itoa(index)
		v.value = &json.JsonData{T: consts.GO_VALUE_STRING, V: value.(string), S: len(value.(string))}
		bind.Set(v.nameKey(), v)
		return v
	case int, int8, int32, int64:
		v := new(jsInteger)
		v.pName = m.name
		v.name = strconv.Itoa(index)
		v.value = &json.JsonData{T: consts.GO_VALUE_INT, V: value, S: strconv.IntSize}
		bind.Set(v.nameKey(), v)
		return v
	case uint, uint8, uint32, uint64:
		v := new(jsInteger)
		v.pName = m.name
		v.name = strconv.Itoa(index)
		v.value = &json.JsonData{T: consts.GO_VALUE_UINT, V: value, S: strconv.IntSize}
		bind.Set(v.nameKey(), v)
		return v
	case float32, float64:
		v := new(jsDouble)
		v.pName = m.name
		v.name = strconv.Itoa(index)
		v.value = &json.JsonData{T: consts.GO_VALUE_FLOAT64, V: value, S: 8}
		bind.Set(v.nameKey(), v)
		return v
	case bool:
		v := new(jsBoolean)
		v.pName = m.name
		v.name = strconv.Itoa(index)
		v.value = &json.JsonData{T: consts.GO_VALUE_BOOL, V: value, S: 1}
		bind.Set(v.nameKey(), v)
		return v
	case nil:
		v := new(jsNull)
		v.pName = m.name
		v.name = strconv.Itoa(index)
		v.value = &json.JsonData{T: consts.GO_VALUE_BOOL, V: null, S: 0}
		bind.Set(v.nameKey(), v)
		return v
	default:
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			kind = rv.Elem().Kind()
		}
		switch kind {
		case reflect.Struct:
			v := new(jsObject)
			v.pName = m.name
			v.name = strconv.Itoa(index)
			v.value = &json.JsonData{T: consts.GO_VALUE_STRUCT, V: rv.Interface(), S: 0}
			v.rv = &rv
			bind.Set(v.nameKey(), v)
			return v
		case reflect.Map:
			v := new(jsObject)
			v.pName = m.name
			v.name = strconv.Itoa(index)
			v.value = &json.JsonData{T: consts.GO_VALUE_MAP, V: rv.Interface(), S: rv.Len()}
			v.rv = &rv
			bind.Set(v.nameKey(), v)
			return v
		case reflect.Slice:
			v := new(jsArray)
			v.pName = m.name
			v.name = strconv.Itoa(index)
			v.value = &json.JsonData{T: consts.GO_VALUE_SLICE, V: rv.Interface(), S: rv.Len()}
			v.rv = &rv
			bind.Set(v.nameKey(), v)
			return v
		case reflect.Func:
			v := new(jsFunction)
			v.pName = m.name
			v.name = strconv.Itoa(index)
			v.value = &json.JsonData{T: consts.GO_VALUE_FUNC, V: rv.Interface(), S: 0}
			v.rv = &rv
			bind.Set(v.nameKey(), v)
			return v
		}
	}
	return nil
}
