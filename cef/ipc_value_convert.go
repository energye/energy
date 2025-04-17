//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// ValueConvert -> v8ValueProcessMessageConvert
//
// IPC 和 ICefV8Value 数据序列化转换

package cef

import (
	goJSON "encoding/json"
	"errors"
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
	"github.com/energye/golcl/lcl/api"
	"reflect"
	"unsafe"
)

// ValueConvert
var ValueConvert v8ValueProcessMessageConvert

// v8ValueProcessMessageConvert ICefV8Value 和 ICefProcessMessage 转换
type v8ValueProcessMessageConvert uintptr

// ListValueToV8Value ICefListValue 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) ListValueToV8Value(list *ICefListValue) (*ICefV8Value, error) {
	if list == nil {
		return nil, errors.New("build v8 value error. Parameter null")
	}
	size := int(list.Size())
	result := V8ValueRef.NewArray(int32(size))
	if !result.IsValid() {
		return nil, errors.New("create v8 array error")
	}
	for i := 0; i < size; i++ {
		var newValue *ICefV8Value
		value := list.GetValue(uint32(i))
		if value.IsValid() {
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
				if v, err := m.DictionaryValueToV8Value(value.GetDictionary()); err == nil {
					newValue = v
				}
			case consts.VTYPE_LIST: // JSONArray
				if v, err := m.ListValueToV8Value(value.GetList()); err == nil {
					newValue = v
				}
			}
		}
		if newValue == nil {
			newValue = V8ValueRef.NewNull()
		}
		result.SetValueByIndex(int32(i), newValue)

	}
	return result, nil
}

// DictionaryValueToV8Value ICefDictionaryValue 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) DictionaryValueToV8Value(dictionary *ICefDictionaryValue) (*ICefV8Value, error) {
	if dictionary == nil {
		return nil, errors.New("build v8 value error. Parameter null")
	}
	keys := dictionary.GetKeys()
	if keys == nil || keys.keys == nil || !keys.keys.IsValid() || keys.Count() == 0 {
		return nil, errors.New("get dict keys error")
	}

	result := V8ValueRef.NewObject(nil, nil)
	if !result.IsValid() {
		return nil, errors.New("create v8 object error")
	}
	for i := 0; i < keys.Count(); i++ {
		var newValue *ICefV8Value
		key := keys.Get(i)
		value := dictionary.GetValue(key)
		if value.IsValid() {
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
				if v, err := m.DictionaryValueToV8Value(value.GetDictionary()); err == nil {
					newValue = v
				}
			case consts.VTYPE_LIST: // JSONArray
				if v, err := m.ListValueToV8Value(value.GetList()); err == nil {
					newValue = v
				}
			}
		}
		if newValue == nil {
			newValue = V8ValueRef.NewNull()
		}
		result.setValueByAccessor(key, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		result.setValueByKey(key, newValue, consts.V8_PROPERTY_ATTRIBUTE_NONE)
	}
	return result, nil
}

// BytesToV8ArrayValue JSONArray 字节数组转换 TCefV8ValueArray
func (m *v8ValueProcessMessageConvert) BytesToV8ArrayValue(resultArgsBytes []byte) (*TCefV8ValueArray, error) {
	//只能是 JSONArray 对象类型
	jsonArray := json.NewJSONArray(resultArgsBytes)
	if jsonArray == nil {
		return nil, errors.New("parsing parameter failure")
	}
	defer func() {
		jsonArray.Free()
	}()
	return m.JSONArrayToV8ArrayValue(jsonArray)
}

func (m *v8ValueProcessMessageConvert) JSONArrayToV8ArrayValue(jsonArray json.JSONArray) (*TCefV8ValueArray, error) {
	//只能是 JSONArray 对象类型
	if jsonArray == nil {
		return nil, errors.New("parsing parameter failure")
	}
	size := jsonArray.Size()
	resultArgs := V8ValueArrayRef.New()
	if resultArgs == nil {
		return nil, errors.New("create v8 value array error")
	}
	for i := 0; i < size; i++ {
		value := jsonArray.GetByIndex(i)
		if value == nil {
			resultArgs.Add(V8ValueRef.NewNull())
			continue
		}
		switch value.Type() {
		case reflect.String:
			resultArgs.Add(V8ValueRef.NewString(value.String()))
		case reflect.Int:
			resultArgs.Add(V8ValueRef.NewInt(int32(value.Int())))
		case reflect.Uint:
			resultArgs.Add(V8ValueRef.NewUInt(uint32(value.UInt())))
		case reflect.Float64:
			resultArgs.Add(V8ValueRef.NewDouble(value.Float()))
		case reflect.Bool:
			resultArgs.Add(V8ValueRef.NewBool(value.Bool()))
		case reflect.Slice:
			if v := m.JSONArrayToV8Value(value.JSONArray()); v != nil {
				resultArgs.Add(v)
			} else {
				resultArgs.Add(V8ValueRef.NewNull())
			}
		case reflect.Map:
			if v := m.JSONObjectToV8Value(value.JSONObject()); v != nil {
				resultArgs.Add(v)
			} else {
				resultArgs.Add(V8ValueRef.NewNull())
			}
		}
	}
	return resultArgs, nil
}

// JSONArrayToV8Value JSONArray 转 ICefV8Value
func (m *v8ValueProcessMessageConvert) JSONArrayToV8Value(array json.JSONArray) *ICefV8Value {
	if array == nil || !array.IsArray() {
		return nil
	}
	size := array.Size()
	result := V8ValueRef.NewArray(int32(size))
	if !result.IsValid() {
		return nil
	}
	for i := 0; i < size; i++ {
		value := array.GetByIndex(i)
		if value == nil {
			result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
			continue
		}
		switch value.Type() {
		case reflect.String:
			result.SetValueByIndex(int32(i), V8ValueRef.NewString(value.String()))
		case reflect.Int:
			result.SetValueByIndex(int32(i), V8ValueRef.NewInt(int32(value.Int())))
		case reflect.Uint:
			result.SetValueByIndex(int32(i), V8ValueRef.NewUInt(uint32(value.UInt())))
		case reflect.Float64:
			result.SetValueByIndex(int32(i), V8ValueRef.NewDouble(value.Float()))
		case reflect.Bool:
			result.SetValueByIndex(int32(i), V8ValueRef.NewBool(value.Bool()))
		case reflect.Slice:
			if v := m.JSONArrayToV8Value(value.JSONArray()); v != nil {
				result.SetValueByIndex(int32(i), v)
			} else {
				result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
			}
		case reflect.Map:
			if v := m.JSONObjectToV8Value(value.JSONObject()); v != nil {
				result.SetValueByIndex(int32(i), v)
			} else {
				result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
			}
		default:
			result.SetValueByIndex(int32(i), V8ValueRef.NewNull())
		}
	}
	return result
}

// BindDataToV8Value  转 ICefV8Value
func (m *v8ValueProcessMessageConvert) BindDataToV8Value(data json.JSONObject) *ICefV8Value {
	if data == nil || data.Size() == 0 {
		return nil
	}
	t := reflect.Kind(data.GetIntByKey("T")) // type
	// 根据类型取值
	switch t {
	case reflect.String:
		return V8ValueRef.NewString(data.GetStringByKey("V"))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return V8ValueRef.NewInt(int32(data.GetIntByKey("V")))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return V8ValueRef.NewUInt(uint32(data.GetUIntByKey("V")))
	case reflect.Float32, reflect.Float64:
		return V8ValueRef.NewDouble(data.GetFloatByKey("V"))
	case reflect.Bool:
		return V8ValueRef.NewBool(data.GetBoolByKey("V"))
	case reflect.Struct, reflect.Map:
		return m.JSONObjectToV8Value(data.GetObjectByKey("V"))
	case reflect.Slice:
		return m.JSONArrayToV8Value(data.GetArrayByKey("V"))
	}
	return nil
}

// JSONObjectToV8Value JSONObject 转 ICefV8Value
func (m *v8ValueProcessMessageConvert) JSONObjectToV8Value(object json.JSONObject) *ICefV8Value {
	if object == nil || !object.IsObject() {
		return nil
	}
	result := V8ValueRef.NewObject(nil, nil)
	if !result.IsValid() {
		return nil
	}
	size := object.Size()
	keys := object.Keys()
	for i := 0; i < size; i++ {
		key := keys[i]
		value := object.GetByKey(key)
		if value == nil {
			result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
			continue
		}
		switch value.Type() {
		case reflect.String:
			result.setValueByKey(key, V8ValueRef.NewString(value.String()), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case reflect.Int:
			result.setValueByKey(key, V8ValueRef.NewInt(int32(value.Int())), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case reflect.Uint:
			result.setValueByKey(key, V8ValueRef.NewUInt(uint32(value.UInt())), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case reflect.Float64:
			result.setValueByKey(key, V8ValueRef.NewDouble(value.Float()), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case reflect.Bool:
			result.setValueByKey(key, V8ValueRef.NewBool(value.Bool()), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		case reflect.Slice:
			if v := m.JSONArrayToV8Value(value.JSONArray()); v != nil {
				result.setValueByKey(key, v, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			} else {
				result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
			}
		case reflect.Map:
			if v := m.JSONObjectToV8Value(value.JSONObject()); v != nil {
				result.setValueByKey(key, v, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			} else {
				result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
			}
		default:
			result.setValueByKey(key, V8ValueRef.NewNull(), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		}
	}
	return result
}

// V8ValueToProcessMessageBytes ICefV8Value 转换 []byte 进程消息
func (m *v8ValueProcessMessageConvert) V8ValueToProcessMessageBytes(v8value *ICefV8Value) []byte {
	if !v8value.IsValid() {
		return nil
	}
	if v := m.V8ValueToProcessMessageArray(v8value); v != nil {
		if v, err := goJSON.Marshal(v); err == nil {
			return v
		}
	}
	return nil
}

// V8ValueToProcessMessageArray 转换 []byte 进程消息
func (m *v8ValueProcessMessageConvert) V8ValueToProcessMessageArray(v8value *ICefV8Value) interface{} {
	if !v8value.IsValid() {
		return nil
	}
	if v8value.IsArray() {
		if result, err := m.V8valueArrayToSlice(v8value); err == nil {
			return result
		}
	} else if v8value.IsObject() {
		if result, err := m.V8valueObjectToMap(v8value); err == nil {
			return result
		}
	} else {
		result := make([]interface{}, 1)
		if v8value.IsString() {
			result[0] = v8value.GetStringValue()
		} else if v8value.IsInt() {
			result[0] = int(v8value.GetIntValue())
		} else if v8value.IsUInt() {
			result[0] = uint(v8value.GetUIntValue())
		} else if v8value.IsDouble() {
			result[0] = v8value.GetDoubleValue()
		} else if v8value.IsBool() {
			result[0] = v8value.GetBoolValue()
		} else if v8value.IsDate() {
			result[0] = v8value.GetDateValue()
		} else if v8value.IsNull() {
			result[0] = "null"
		} else if v8value.IsUndefined() {
			result[0] = "undefined"
		} else if v8value.IsArrayBuffer() {
			result[0] = ""
		} else {
			result[0] = "" // function/byte/buffer
		}
		return result
	}
	return nil
}

// V8valueArrayToSlice ICefV8Value 转换 Slice
func (m *v8ValueProcessMessageConvert) V8valueArrayToSlice(v8value *ICefV8Value) ([]interface{}, error) {
	if !v8value.IsArray() {
		return nil, errors.New("convert list value error. Please pass in the array type")
	}
	argsLen := v8value.GetArrayLength()
	result := make([]interface{}, argsLen)
	for i := 0; i < argsLen; i++ {
		args := v8value.GetValueByIndex(i)
		if !args.IsValid() {
			result[i] = "null"
			continue
		}
		if args.IsString() {
			result[i] = args.GetStringValue()
		} else if args.IsInt() {
			result[i] = int(args.GetIntValue())
		} else if args.IsUInt() {
			result[i] = uint(args.GetUIntValue())
		} else if args.IsDouble() {
			result[i] = args.GetDoubleValue()
		} else if args.IsBool() {
			result[i] = args.GetBoolValue()
		} else if v8value.IsDate() {
			result[i] = v8value.GetDateValue()
		} else if args.IsNull() {
			result[i] = "null"
		} else if args.IsUndefined() {
			result[i] = "undefined"
		} else if args.IsArray() {
			if v, err := m.V8valueArrayToSlice(args); err == nil {
				result[i] = v
			} else {
				result[i] = nil
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToMap(args); err == nil {
				result[i] = v
			} else {
				result[i] = nil
			}
		} else if args.IsArrayBuffer() {
			result[i] = ""
		} else {
			result[i] = ""
		}
		args.Free()
	}
	return result, nil
}

// V8valueObjectToMap ICefV8Value 转换 Maps
func (m *v8ValueProcessMessageConvert) V8valueObjectToMap(v8value *ICefV8Value) (map[string]interface{}, error) {
	if !v8value.IsObject() {
		return nil, errors.New("convert dictionary value error. Please pass in the object type")
	}
	keys := v8value.GetKeys()
	if keys == nil || keys.keys == nil || !keys.keys.IsValid() || keys.Count() == 0 {
		return nil, errors.New("get dict keys error")
	}
	result := make(map[string]interface{}, keys.Count())
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		args := v8value.getValueByKey(key)
		if !args.IsValid() {
			result[key] = "null"
			continue
		}
		if args.IsString() {
			result[key] = args.GetStringValue()
		} else if args.IsInt() {
			result[key] = int(args.GetIntValue())
		} else if args.IsUInt() {
			result[key] = uint(args.GetUIntValue())
		} else if args.IsDouble() {
			result[key] = args.GetDoubleValue()
		} else if args.IsBool() {
			result[key] = args.GetBoolValue()
		} else if v8value.IsDate() {
			result[key] = v8value.GetDateValue()
		} else if args.IsNull() {
			result[key] = "null"
		} else if args.IsUndefined() {
			result[key] = "undefined"
		} else if args.IsArray() {
			if v, err := m.V8valueArrayToSlice(args); err == nil {
				result[key] = v
			} else {
				result[key] = nil
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToMap(args); err == nil {
				result[key] = v
			} else {
				result[key] = nil
			}
		} else if args.IsArrayBuffer() {
			//arrayValue.SetBinary()
			result[key] = ""
		} else {
			result[key] = ""
		}
		args.Free()
	}
	keys.Free()
	return result, nil
}

// V8ValueToProcessMessage ICefV8Value 转换 进程消息
func (m *v8ValueProcessMessageConvert) V8ValueToProcessMessage(v8value *ICefV8Value) (*ICefListValue, error) {
	if v8value == nil || !v8value.IsValid() {
		return nil, errors.New("build process value error. Parameter null")
	}
	if v8value.IsArray() {
		return m.V8valueArrayToListValue(v8value)
	} else if v8value.IsObject() {
		if v, err := m.V8valueObjectToDictionaryValue(v8value); err == nil {
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
		} else if v8value.IsArrayBuffer() {
			//arrayValue.SetBinary()
		} else {
			arrayValue.SetNull(uint32(0))
		}
		return arrayValue, nil
	}
}

// V8valueArrayToListValue ICefV8Value 转换 ICefListValue
func (m *v8ValueProcessMessageConvert) V8valueArrayToListValue(v8value *ICefV8Value) (*ICefListValue, error) {
	if !v8value.IsArray() {
		return nil, errors.New("convert list value error. Please pass in the array type")
	}
	arrayValue := ListValueRef.New()
	if !arrayValue.IsValid() {
		return nil, errors.New("create list value error")
	}
	argsLen := v8value.GetArrayLength()
	for i := 0; i < argsLen; i++ {
		args := v8value.GetValueByIndex(i)
		if !args.IsValid() {
			arrayValue.SetNull(uint32(i))
			continue
		}
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
			if v, err := m.V8valueArrayToListValue(args); err == nil {
				arrayValue.SetList(uint32(i), v)
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToDictionaryValue(args); err == nil {
				arrayValue.SetDictionary(uint32(i), v)
			}
		} else if args.IsArrayBuffer() {
			//arrayValue.SetBinary()
		} else {
			arrayValue.SetNull(uint32(i))
		}
		args.Free()
	}
	return arrayValue, nil
}

// V8valueObjectToDictionaryValue ICefV8Value 转换 ICefDictionaryValue
func (m *v8ValueProcessMessageConvert) V8valueObjectToDictionaryValue(v8value *ICefV8Value) (*ICefDictionaryValue, error) {
	if !v8value.IsObject() {
		return nil, errors.New("convert dictionary value error. Please pass in the object type")
	}
	keys := v8value.GetKeys()
	if keys == nil || keys.keys == nil || !keys.keys.IsValid() || keys.Count() == 0 {
		return nil, errors.New("get dict keys error")
	}
	dictionaryValue := DictionaryValueRef.New()
	if !dictionaryValue.IsValid() {
		return nil, errors.New("create dict value error")
	}
	for i := 0; i < keys.Count(); i++ {
		key := keys.Get(i)
		args := v8value.getValueByKey(key)
		if !args.IsValid() {
			dictionaryValue.SetNull(key)
			continue
		}
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
			if v, err := m.V8valueArrayToListValue(args); err == nil {
				dictionaryValue.SetList(key, v)
			}
		} else if args.IsObject() {
			if v, err := m.V8valueObjectToDictionaryValue(args); err == nil {
				dictionaryValue.SetDictionary(key, v)
			}
		} else {
			dictionaryValue.SetNull(key)
		}
		args.Free()
	}
	keys.Free()
	return dictionaryValue, nil
}

// BytesToProcessMessage JSONArray []byte 转换 进程消息 ICefProcessMessage
func (m *v8ValueProcessMessageConvert) BytesToProcessMessage(name string, data []byte) *ICefProcessMessage {
	var result uintptr
	imports.Proc(def.ValueConvert_BytesToProcessMessage).Call(api.PascalStr(name), uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefProcessMessage{instance: unsafe.Pointer(result)}
}

// JSONArrayToProcessMessage 转换 进程消息 ICefProcessMessage
func (m *v8ValueProcessMessageConvert) JSONArrayToProcessMessage(name string, array json.JSONArray) *ICefProcessMessage {
	if array == nil || !array.IsArray() {
		return nil
	}
	return m.BytesToProcessMessage(name, array.Bytes())
}

// BytesToListValue []byte 转换 ICefListValue
func (m *v8ValueProcessMessageConvert) BytesToListValue(data []byte) *ICefListValue {
	var result uintptr
	imports.Proc(def.ValueConvert_BytesToListValue).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefListValue{instance: unsafe.Pointer(result)}
}

// JSONArrayToListValue 转换 ICefListValue
func (m *v8ValueProcessMessageConvert) JSONArrayToListValue(array json.JSONArray) *ICefListValue {
	if array == nil || !array.IsArray() {
		return nil
	}
	return m.BytesToListValue(array.Bytes())
}

// BytesToDictionaryValue 转换 ICefDictionaryValue
func (m *v8ValueProcessMessageConvert) BytesToDictionaryValue(data []byte) *ICefDictionaryValue {
	var result uintptr
	imports.Proc(def.ValueConvert_BytesToDictionaryValue).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefDictionaryValue{instance: unsafe.Pointer(result)}
}

// JSONObjectToDictionaryValue 转换 ICefDictionaryValue
func (m *v8ValueProcessMessageConvert) JSONObjectToDictionaryValue(object json.JSONObject) *ICefDictionaryValue {
	if object == nil || !object.IsObject() {
		return nil
	}
	return m.BytesToDictionaryValue(object.Bytes())
}

// BytesToV8ValueArray 转换 TCefV8ValueArray
func (m *v8ValueProcessMessageConvert) BytesToV8ValueArray(data []byte) *TCefV8ValueArray {
	var result uintptr
	var resultLength uintptr
	imports.Proc(def.ValueConvert_BytesToV8ValueArray).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&resultLength)))
	if result == 0 || resultLength == 0 {
		return nil
	}
	return &TCefV8ValueArray{instance: unsafe.Pointer(result), arguments: result, argumentsLength: int(resultLength) /*, argumentsCollect: make([]*ICefV8Value, int(resultLength))*/}
}

// JSONArrayToV8ValueArray 转换 TCefV8ValueArray
func (m *v8ValueProcessMessageConvert) JSONArrayToV8ValueArray(array json.JSONArray) *TCefV8ValueArray {
	if array == nil || !array.IsArray() {
		return nil
	}
	return m.BytesToV8ValueArray(array.Bytes())
}

// BytesToV8Array 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) BytesToV8Array(data []byte) *ICefV8Value {
	var result uintptr
	imports.Proc(def.ValueConvert_BytesToV8Array).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{instance: unsafe.Pointer(result)}
}

// JSONArrayToV8Array 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) JSONArrayToV8Array(array json.JSONArray) *ICefV8Value {
	if array == nil || !array.IsArray() {
		return nil
	}
	return m.BytesToV8Array(array.Bytes())
}

// BytesToV8Object 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) BytesToV8Object(data []byte) *ICefV8Value {
	var result uintptr
	imports.Proc(def.ValueConvert_BytesToV8Object).Call(uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{instance: unsafe.Pointer(result)}
}

// JSONObjectToV8Object 转换 ICefV8Value
func (m *v8ValueProcessMessageConvert) JSONObjectToV8Object(object json.JSONObject) *ICefV8Value {
	if object == nil || !object.IsObject() {
		return nil
	}
	return m.BytesToV8Object(object.Bytes())
}
