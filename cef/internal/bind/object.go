//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSObject 实现

package bind

import (
	"reflect"
	"unsafe"
)

type JSObject interface {
	JSValue
	ObjectValue() any
	Get(fieldName string) JSValue
	Set(fieldName string, value any)
	fieldToBind()
}

type jsObject struct {
	V8Value
}

func (m *jsObject) ObjectValue() any {
	if m.IsObject() {
		return m.JsonData.Data()
	}
	return nil
}

func (m *jsObject) Get(fieldName string) JSValue {
	if fieldName == "" {
		return nil
	}
	if m.IsObject() {
		rv := m.Data().(*reflect.Value).Elem().FieldByName(fieldName)
		return m.createJSValue(fieldName, &rv)
	}
	return nil
}

// fieldToBind 对象字段绑定
func (m *jsObject) fieldToBind() {
	if m.IsObject() {
		rv := *m.Data().(*reflect.Value)
		rt := rv.Type()
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
			rt = rt.Elem()
		}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			name := rt.Field(i).Name
			value := m.createJSValue(name, &field)
			if value.IsObject() {
				value.(JSObject).fieldToBind()
			} else if value.IsArray() {
				for _, item := range value.(JSArray).Items() {
					if item.IsObject() {
						item.(JSObject).fieldToBind()
					} else if item.IsArray() {
						item.(JSArray).fieldToBind()
					}
				}
			}
		}
	}
}

func (m *jsObject) Set(fieldName string, value any) {
	if m.IsObject() {
		field := m.Data().(*reflect.Value).Elem().FieldByName(fieldName)
		switch field.Kind() {
		case reflect.String:
			if v, ok := value.(string); ok {
				*(*string)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Int:
			if v, ok := value.(int); ok {
				*(*int)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Int8:
			if v, ok := value.(int8); ok {
				*(*int8)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Int16:
			if v, ok := value.(int16); ok {
				*(*int16)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Int32:
			if v, ok := value.(int32); ok {
				*(*int32)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Int64:
			if v, ok := value.(int64); ok {
				*(*int64)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Uint:
			if v, ok := value.(uint); ok {
				*(*uint)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Uint8:
			if v, ok := value.(uint8); ok {
				*(*uint8)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Uint16:
			if v, ok := value.(uint16); ok {
				*(*uint16)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Uint32:
			if v, ok := value.(uint32); ok {
				*(*uint32)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Uint64:
			if v, ok := value.(uint64); ok {
				*(*uint64)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Float32:
			if v, ok := value.(float32); ok {
				*(*float32)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Float64:
			if v, ok := value.(float64); ok {
				*(*float64)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Bool:
			if v, ok := value.(bool); ok {
				*(*bool)(unsafe.Pointer(field.Addr().Pointer())) = v
			}
		case reflect.Ptr: //指针仅支持 struct slice map
			valRv := reflect.ValueOf(value)
			fieldType := field.Type().Elem()
			valRvKind := valRv.Kind()
			if valRvKind == reflect.Ptr {
				valRvKind = valRv.Elem().Kind()
			}
			if fieldType.Kind() == reflect.Struct && valRvKind == reflect.Struct {
				field.Set(valRv)
			} else if field.Kind() == reflect.Slice && valRvKind == reflect.Slice {
				field.Set(valRv)
			} else if field.Kind() == reflect.Map && valRvKind == reflect.Map {
				field.Set(valRv)
			}
		case reflect.Struct:
			valRv := reflect.ValueOf(value)
			field.Set(valRv)
		case reflect.Slice:
			valRv := reflect.ValueOf(value)
			field.Set(valRv)
		case reflect.Map:
			valRv := reflect.ValueOf(value)
			field.Set(valRv)
		}
	}
}
