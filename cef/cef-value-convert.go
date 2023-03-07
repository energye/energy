//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//ValueConvert 值转换
// 进程消息、V8Value、Go类型
package cef

import (
	"errors"
	"fmt"
	"github.com/energye/energy/consts"
	"reflect"
	"unsafe"
)

//go 基本类型值转换
func goValueConvert(result *ICefListValue, i uint32, value interface{}) bool {
	switch value.(type) {
	case bool:
		result.SetBool(i, value.(bool))
	case float32:
		result.SetDouble(i, float64(value.(float32)))
	case float64:
		result.SetDouble(i, value.(float64))
	case string:
		result.SetString(i, value.(string))
	case int:
		result.SetInt(i, int32(value.(int)))
	case int8:
		result.SetInt(i, int32(value.(int8)))
	case int16:
		result.SetInt(i, int32(value.(int16)))
	case int32:
		result.SetInt(i, value.(int32))
	case int64:
		result.SetString(i, fmt.Sprintf("%d", value.(int64)))
	case uint:
		result.SetInt(i, int32(value.(uint)))
	case uint8:
		result.SetInt(i, int32(value.(uint8)))
	case uint16:
		result.SetInt(i, int32(value.(uint16)))
	case uint32:
		result.SetInt(i, int32(value.(uint32)))
	case uint64:
		result.SetString(i, fmt.Sprintf("%d", value.(uint64)))
	default:
		return false
	}
	return true
}

//go 数组基本类型值转换
func goArrayValueConvert(result *ICefListValue, i uint32, value interface{}) bool {
	switch value.(type) {
	case []bool:
		var list = ListValueRef.New()
		for j, v := range value.([]bool) {
			list.SetBool(uint32(j), v)
		}
		result.SetList(i, list)
	case []float32:
		var list = ListValueRef.New()
		for j, v := range value.([]float32) {
			list.SetDouble(uint32(j), float64(v))
		}
		result.SetList(i, list)
	case []float64:
		var list = ListValueRef.New()
		for j, v := range value.([]float64) {
			list.SetDouble(uint32(j), v)
		}
		result.SetList(i, list)
	case []string:
		var list = ListValueRef.New()
		for j, v := range value.([]string) {
			list.SetString(uint32(j), v)
		}
		result.SetList(i, list)
	case []int:
		var list = ListValueRef.New()
		for j, v := range value.([]int) {
			list.SetInt(uint32(j), int32(v))
		}
		result.SetList(i, list)
	case []int8:
		var list = ListValueRef.New()
		for j, v := range value.([]int8) {
			list.SetInt(uint32(j), int32(v))
		}
		result.SetList(i, list)
	case []int16:
		var list = ListValueRef.New()
		for j, v := range value.([]int16) {
			list.SetInt(uint32(j), int32(v))
		}
		result.SetList(i, list)
	case []int32:
		var list = ListValueRef.New()
		for j, v := range value.([]int32) {
			list.SetInt(uint32(j), v)
		}
		result.SetList(i, list)
	case []int64:
		var list = ListValueRef.New()
		for j, v := range value.([]int64) {
			list.SetString(uint32(j), fmt.Sprintf("%d", v))
		}
		result.SetList(i, list)
	case []uint:
		var list = ListValueRef.New()
		for j, v := range value.([]uint) {
			list.SetInt(uint32(j), int32(v))
		}
		result.SetList(i, list)
	case []uint8:
		BinaryValueRef.New(value.([]byte))
		result.SetBinary(i, BinaryValueRef.New(value.([]byte)))
	case []uint16:
		var list = ListValueRef.New()
		for j, v := range value.([]uint16) {
			list.SetInt(uint32(j), int32(v))
		}
		result.SetList(i, list)
	case []uint32:
		var list = ListValueRef.New()
		for j, v := range value.([]uint32) {
			list.SetInt(uint32(j), int32(v))
		}
		result.SetList(i, list)
	case []uint64:
		var list = ListValueRef.New()
		for j, v := range value.([]uint64) {
			list.SetString(uint32(j), fmt.Sprintf("%d", v))
		}
		result.SetList(i, list)
	case []interface{}:
		if v, err := goValueToListValue(value.([]interface{})); err == nil {
			result.SetList(i, v)
		} else {
			result.SetList(i, ListValueRef.New())
		}
	default:
		return false
	}
	return true
}

//goValueToListValue GoValue 转换 ICefListValue
func goValueToListValue(data []interface{}) (*ICefListValue, error) {
	if data == nil {
		return nil, errors.New("build process message error. Parameter null")
	}
	var result = ListValueRef.New()

	for i, value := range data {
		if value == nil {
			result.SetNull(uint32(i))
			continue
		}
		var rv = reflect.ValueOf(value)
		kind := rv.Kind()
		switch kind {
		case reflect.Ptr: //单指针, 不允许(切片 | 数组)指针, 包含基本类型、结构、Map
			if rv.IsNil() {
				result.SetNull(uint32(i))
				continue
			}
			//基本类型
			if ok := goValueConvert(result, uint32(i), rv.Elem().Interface()); !ok {
				//非基本类型
				switch rv.Elem().Kind() {
				case reflect.Struct:
					if v := goStructValueToDictionaryValue(rv); v != nil {
						result.SetDictionary(uint32(i), v)
					} else {
						result.SetNull(uint32(i))
					}
				case reflect.Map:
				default:
					result.SetNull(uint32(i))
				}
			}
		case reflect.Slice, reflect.Array: //切片 | 数组, 包含基本类型、结构、Map
			//基本类型
			if ok := goArrayValueConvert(result, uint32(i), value); !ok {
				//非基本类型
				var sliceType = rv.Type().Elem()
				fmt.Println("sliceType 1", sliceType, sliceType.Kind())
				if sliceType.Kind() == reflect.Ptr {
					//result.SetList(uint32(i), ListValueRef.New())
					//continue
					sliceType = sliceType.Elem()
				}
				fmt.Println("sliceType 2", sliceType, sliceType.Kind())
				switch sliceType.Kind() {
				case reflect.Struct:
					sliceListValue := ListValueRef.New()
					for j := 0; j < rv.Len(); j++ {
						if v := goStructValueToDictionaryValue(rv.Index(j)); v != nil {
							sliceListValue.SetDictionary(uint32(j), v)
						} else {
							sliceListValue.SetNull(uint32(j))
						}
					}
					result.SetList(uint32(i), sliceListValue)
				case reflect.Map:
				default:
					result.SetList(uint32(i), ListValueRef.New())
				}

			}
		case reflect.Map: // 单Map
			fmt.Println("goValueToListValue Map")
		default: //默认 基本类型
			if ok := goValueConvert(result, uint32(i), value); !ok {
				result.SetNull(uint32(i))
			}
		}

	}
	return result, nil
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

// goStructValueToDictionaryValue Go结构转字典
func goStructValueToDictionaryValue(rv reflect.Value) *ICefDictionaryValue {
	var rt reflect.Type
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rt = rv.Type().Elem()
	} else {
		rt = rv.Type()
	}
	result := DictionaryValueRef.New()
	nf := rt.NumField()
	ei := rv.Interface()
	ptrOffset := uintptr((*emptyInterface)(unsafe.Pointer(&ei)).word)
	for i := 0; i < nf; i++ {
		fieldRt := rt.Field(i)
		name := fieldRt.Name
		if fieldRt.IsExported() { //导出的字段
			ptr := unsafe.Pointer(fieldRt.Offset + ptrOffset)
			switch fieldRt.Type.Kind() {
			case reflect.String:
				result.SetString(name, *(*string)(ptr))
			case reflect.Bool:
				result.SetBool(name, *(*bool)(ptr))
			case reflect.Int:
				result.SetInt(name, int32(*(*int)(ptr)))
			case reflect.Int8:
				result.SetInt(name, int32(*(*int8)(ptr)))
			case reflect.Int16:
				result.SetInt(name, int32(*(*int16)(ptr)))
			case reflect.Int32:
				result.SetInt(name, *(*int32)(ptr))
			case reflect.Int64:
				result.SetInt(name, int32(*(*int64)(ptr)))
			case reflect.Uint:
				result.SetInt(name, int32(*(*uint)(ptr)))
			case reflect.Uint8:
				result.SetInt(name, int32(*(*uint8)(ptr)))
			case reflect.Uint16:
				result.SetInt(name, int32(*(*uint16)(ptr)))
			case reflect.Uint32:
				result.SetInt(name, int32(*(*uint32)(ptr)))
			case reflect.Uint64:
				result.SetInt(name, int32(*(*uint64)(ptr)))
			case reflect.Uintptr:
				result.SetInt(name, int32(*(*uintptr)(ptr)))
			case reflect.Float32:
				result.SetDouble(name, float64(*(*float32)(ptr)))
			case reflect.Float64:
				result.SetDouble(name, *(*float64)(ptr))
			default:
				result.SetNull(name)
			}
		} else {
			result.SetNull(name)
		}
	}
	return result
}

//listValueToV8Value ICefListValue 转换 ICefV8Value
func listValueToV8Value(list *ICefListValue) (*ICefV8Value, error) {
	if list == nil {
		return nil, errors.New("build v8 value error. Parameter null")
	}
	size := int(list.Size())
	result := V8ValueRef.NewArray(int32(size))
	for i := 0; i < size; i++ {
		value := list.GetValue(uint32(i))
		var newValue *ICefV8Value
		switch value.GetType() {
		case consts.VTYPE_NULL:
			newValue = V8ValueRef.NewNull()
		case consts.VTYPE_BOOL:
			newValue = V8ValueRef.NewBool(value.GetBool())
		case consts.VTYPE_INT:
			newValue = V8ValueRef.NewInt(value.GetInt())
		case consts.VTYPE_DOUBLE:
			newValue = V8ValueRef.NewDouble(value.GetDouble())
		case consts.VTYPE_STRING:
			newValue = V8ValueRef.NewString(value.GetString())
		case consts.VTYPE_BINARY: // []byte
			binaryValue := value.GetBinary()
			byteSize := binaryValue.GetSize()
			if byteSize > 0 {
				dataByte := make([]byte, binaryValue.GetSize())
				if c := binaryValue.GetData(dataByte, 0); c > 0 {
					newValue = V8ValueRef.NewArrayBuffer(dataByte, nil)
				}
			}
		case consts.VTYPE_DICTIONARY: // Object
			if v, err := dictionaryValueToV8Value(value.GetDictionary()); err == nil {
				newValue = v
			}
		case consts.VTYPE_LIST: // Array
			if v, err := listValueToV8Value(value.GetList()); err == nil {
				newValue = v
			}
		}
		if newValue == nil {
			newValue = V8ValueRef.NewNull()
		}
		result.SetValueByIndex(int32(i), newValue)
	}
	return result, nil
}

//dictionaryValueToV8Value ICefDictionaryValue 转换 ICefV8Value
func dictionaryValueToV8Value(dictionary *ICefDictionaryValue) (*ICefV8Value, error) {
	if dictionary == nil {
		return nil, errors.New("build v8 value error. Parameter null")
	}
	keys := dictionary.GetKeys()
	//bindSubObjectAccessor := V8AccessorRef.New()
	//bindSubObjectAccessor.Get(ctx.bindSubObjectGet)
	//bindSubObjectAccessor.Set(ctx.bindSubObjectSet)
	result := V8ValueRef.NewObject(nil)
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		value := dictionary.GetValue(key)
		var newValue *ICefV8Value
		switch value.GetType() {
		case consts.VTYPE_NULL:
			newValue = V8ValueRef.NewNull()
		case consts.VTYPE_BOOL:
			newValue = V8ValueRef.NewBool(value.GetBool())
		case consts.VTYPE_INT:
			newValue = V8ValueRef.NewInt(value.GetInt())
		case consts.VTYPE_DOUBLE:
			newValue = V8ValueRef.NewDouble(value.GetDouble())
		case consts.VTYPE_STRING:
			newValue = V8ValueRef.NewString(value.GetString())
		case consts.VTYPE_BINARY: // []byte
		case consts.VTYPE_DICTIONARY: // Object
			if v, err := dictionaryValueToV8Value(value.GetDictionary()); err == nil {
				newValue = v
			}
		case consts.VTYPE_LIST: // Array
			if v, err := listValueToV8Value(value.GetList()); err == nil {
				newValue = v
			}
		}
		if newValue != nil {
			result.setValueByAccessor(key, consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			result.setValueByKey(key, newValue, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		}
	}
	return result, nil
}

// v8ValueToProcessMessage ICefV8Value 转换 ICefListValue
func v8ValueToProcessMessage(v8value *ICefV8Value) (*ICefListValue, error) {
	if v8value == nil {
		return nil, errors.New("build process message error. Parameter null")
	}
	if v8value.IsArray() {
		return v8valueArrayEncode(v8value)
	} else if v8value.IsObject() {
		if v, err := v8valueObjectEncode(v8value); err == nil {
			arrayValue := ListValueRef.New()
			arrayValue.SetDictionary(uint32(0), v)
			return arrayValue, nil
		} else {
			return nil, err
		}
	} else {
		arrayValue := ListValueRef.New()
		if v8value.IsString() {
			arrayValue.SetString(uint32(0), v8value.GetStringValue())
		} else if v8value.IsInt() {
			arrayValue.SetInt(uint32(0), v8value.GetIntValue())
		} else if v8value.IsUInt() {
			arrayValue.SetInt(uint32(0), int32(v8value.GetUIntValue()))
		} else if v8value.IsDouble() {
			arrayValue.SetDouble(uint32(0), v8value.GetDoubleValue())
		} else if v8value.IsBool() {
			arrayValue.SetBool(uint32(0), v8value.GetBoolValue())
		} else {
			arrayValue.SetNull(uint32(0))
		}
		return arrayValue, nil
	}
}

// v8valueArrayEncode ICefV8Value 转换 ICefListValue
func v8valueArrayEncode(v8value *ICefV8Value) (*ICefListValue, error) {
	if !v8value.IsArray() {
		return nil, errors.New("build process message error. Please pass in the array type")
	}
	arrayValue := ListValueRef.New()
	argsLen := v8value.GetArrayLength()
	for i := 0; i < argsLen; i++ {
		args := v8value.GetValueByIndex(i)
		if args.IsString() {
			arrayValue.SetString(uint32(i), args.GetStringValue())
		} else if args.IsInt() {
			arrayValue.SetInt(uint32(i), args.GetIntValue())
		} else if args.IsUInt() {
			arrayValue.SetInt(uint32(i), int32(args.GetUIntValue()))
		} else if args.IsDouble() {
			arrayValue.SetDouble(uint32(i), args.GetDoubleValue())
		} else if args.IsBool() {
			arrayValue.SetBool(uint32(i), args.GetBoolValue())
		} else if args.IsNull() || args.IsUndefined() {
			arrayValue.SetNull(uint32(i))
		} else if args.IsArray() {
			if v, err := v8valueArrayEncode(args); err == nil {
				arrayValue.SetList(uint32(i), v)
			}
		} else if args.IsObject() {
			if v, err := v8valueObjectEncode(args); err == nil {
				arrayValue.SetDictionary(uint32(i), v)
			}
		} else {
			arrayValue.SetNull(uint32(i))
		}
	}
	return arrayValue, nil
}

// v8valueObjectEncode ICefV8Value 转换 ICefDictionaryValue
func v8valueObjectEncode(v8value *ICefV8Value) (*ICefDictionaryValue, error) {
	if !v8value.IsObject() {
		return nil, errors.New("build process message error. Please pass in the object type")
	}
	dictionaryValue := DictionaryValueRef.New()
	keys := v8value.GetKeys()
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		args := v8value.GetValueByKey(key)
		if args.IsString() {
			dictionaryValue.SetString(key, args.GetStringValue())
		} else if args.IsInt() {
			dictionaryValue.SetInt(key, args.GetIntValue())
		} else if args.IsUInt() {
			dictionaryValue.SetInt(key, int32(args.GetUIntValue()))
		} else if args.IsDouble() {
			dictionaryValue.SetDouble(key, args.GetDoubleValue())
		} else if args.IsBool() {
			dictionaryValue.SetBool(key, args.GetBoolValue())
		} else if args.IsNull() || args.IsUndefined() {
			dictionaryValue.SetNull(key)
		} else if args.IsArray() {
			if v, err := v8valueArrayEncode(args); err == nil {
				dictionaryValue.SetList(key, v)
			}
		} else if args.IsObject() {
			if v, err := v8valueObjectEncode(args); err == nil {
				dictionaryValue.SetDictionary(key, v)
			}
		} else {
			dictionaryValue.SetNull(key)
		}
	}
	return dictionaryValue, nil
}
