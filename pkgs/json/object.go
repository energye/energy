//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Go implements JSON serialization and JSON deserialization based on map and slice
// key string, value any
// JSONObject

package json

import (
	"encoding/json"
	"github.com/energye/energy/consts"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strconv"
	"strings"
)

// JSONObject
//	object type key : value
//  According to the key operation data, the default value of the data type is returned when a failure occurs
type JSONObject interface {
	BaseJSON
	HasKey(key string) bool               // has key
	Set(key string, value any)            // Sets or overrides the value of the specified key and sets a new arbitrary type value
	RemoveByKey(key string)               // remove data for key
	GetStringByKey(key string) string     // return string data for key
	GetIntByKey(key string) int           // return int data for key
	GetInt64ByKey(key string) int64       // return int64 data for key
	GetUIntByKey(key string) uint         // return uint data for key
	GetUInt64ByKey(key string) uint64     // return uint64 data for key
	GetBytesByKey(key string) []byte      // return []byte data for key
	GetFloatByKey(key string) float64     // return float64 data for key
	GetBoolByKey(key string) bool         // return bool data for key
	GetArrayByKey(key string) JSONArray   // return JSONArray data for key
	GetObjectByKey(key string) JSONObject // return JSONObject data for key
	GetByKey(key string) JSON             // return JSON data for key
	Keys() []string                       // return current object all key
}

// NewJSONObject
//	byte JSONObject, struct convert
//  value:
//    []byte("{...}")
//    struct
//    map[string][type]
func NewJSONObject(value any) JSONObject {
	if value != nil {
		switch value.(type) {
		// if []byte must byte JSONObject
		case []byte:
			if v := NewJSON(value.([]byte)); v != nil {
				return v.JSONObject()
			} else {
				return nil
			}
		// if string must string JSONObject
		case string:
			if v := NewJSON([]byte(value.(string))); v != nil {
				return v.JSONObject()
			} else {
				return nil
			}
		}
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			kind = rv.Elem().Kind()
		}
		if kind != reflect.Map && kind != reflect.Struct {
			return nil
		}
		//convert map[string]any type
		if byt, err := jsoniter.Marshal(value); err == nil {
			var v map[string]any
			if err = jsoniter.Unmarshal(byt, &v); err == nil {
				return &JsonData{t: reflect.Map, v: v, s: len(v)}
			}
		}
	}
	return &JsonData{t: reflect.Map, v: make(map[string]any), s: 0}
}

func (m *JsonData) HasKey(key string) bool {
	if m.IsObject() {
		_, ok := m.v.(map[string]any)[key]
		return ok
	}
	return false
}

func (m *JsonData) Set(key string, value any) {
	if m.IsObject() {
		switch value.(type) {
		case []byte:
			if vv := NewJSON(value.([]byte)); vv != nil {
				value = vv.Data()
			}
		case JsonData:
			value = value.(JsonData)
		case *JsonData:
			value = value.(*JsonData)
		case JSON:
			value = value.(JSON).JsonData()
		case JSONObject:
			value = value.(JSONObject).JsonData()
		case JSONArray:
			value = value.(JSONArray).JsonData()
		default:
			if !isBaseType(value) {
				rv := reflect.ValueOf(value)
				kind := rv.Kind()
				if kind == reflect.Ptr {
					kind = rv.Elem().Kind()
				}
				if kind == reflect.Struct {
					// struct -> map
					if d, err := jsoniter.Marshal(value); err == nil {
						var v map[string]any
						if err = jsoniter.Unmarshal(d, &v); err == nil {
							value = v // json object
						}
					}
				} else if kind == reflect.Slice || kind == reflect.Array {
					// slice -> array
					if d, err := jsoniter.Marshal(value); err == nil {
						var v []any
						if err = jsoniter.Unmarshal(d, &v); err == nil {
							value = v // json array
						}
					}
				}
			}
		}
		if _, ok := m.v.(map[string]any)[key]; !ok {
			m.s++
		}
		m.v.(map[string]any)[key] = value // default base type
	}
}

func (m *JsonData) RemoveByKey(key string) {
	if m.IsObject() {
		if _, ok := m.v.(map[string]any)[key]; ok {
			delete(m.v.(map[string]any), key)
			m.s--
		}
	}
}

func (m *JsonData) GetStringByKey(key string) string {
	if m.IsObject() {
		if value, ok := m.v.(map[string]any)[key]; ok {
			switch value.(type) {
			case *JsonData:
				return value.(*JsonData).String()
			default:
				if r, rok := value.(string); rok {
					return r
				}
			}
		}
	}
	return ""
}

func (m *JsonData) GetIntByKey(key string) (r int) {
	if m.IsObject() {
		if value, ok := m.v.(map[string]any)[key]; ok {
			switch value.(type) {
			case *JsonData:
				return value.(*JsonData).Int()
			default:
				r, _ = toInt(value)
			}
		}
	}
	return
}

func (m *JsonData) GetInt64ByKey(key string) (r int64) {
	if m.IsObject() {
		if value, ok := m.v.(map[string]any)[key]; ok {
			switch value.(type) {
			case *JsonData:
				return value.(*JsonData).Int64()
			default:
				r, _ = toInt64(value)
			}
		}
	}
	return
}

func (m *JsonData) GetUIntByKey(key string) (r uint) {
	if m.IsObject() {
		if value, ok := m.v.(map[string]any)[key]; ok {
			switch value.(type) {
			case *JsonData:
				return value.(*JsonData).UInt()
			default:
				r, _ = toUInt(value)
			}
		}
	}
	return
}

func (m *JsonData) GetUInt64ByKey(key string) (r uint64) {
	if m.IsObject() {
		if value, ok := m.v.(map[string]any)[key]; ok {
			switch value.(type) {
			case *JsonData:
				return value.(*JsonData).UInt64()
			default:
				r, _ = toUInt64(value)
			}
		}
	}
	return
}

func (m *JsonData) GetBytesByKey(key string) []byte {
	if m.IsObject() {
		return toBytes(m.GetObjectByKey(key).JsonData().ConvertToData())
	}
	return nil
}

func (m *JsonData) GetFloatByKey(key string) (r float64) {
	if m.IsObject() {
		if value, ok := m.v.(map[string]any)[key]; ok {
			switch value.(type) {
			case *JsonData:
				return value.(*JsonData).Float()
			default:
				r, _ = toFloat64(value)
			}
		}
	}
	return
}

func (m *JsonData) GetBoolByKey(key string) (r bool) {
	if m.IsObject() {
		if value, ok := m.v.(map[string]any)[key]; ok {
			switch value.(type) {
			case *JsonData:
				return value.(*JsonData).Bool()
			default:
				r, _ = toBool(value)
			}
		}
	}
	return
}

func (m *JsonData) GetArrayByKey(key string) JSONArray {
	return m.GetByKey(key)
}

func (m *JsonData) GetObjectByKey(key string) JSONObject {
	return m.GetByKey(key)
}

func (m *JsonData) Keys() []string {
	if m.IsObject() {
		var result []string
		for key, _ := range m.v.(map[string]any) {
			result = append(result, key)
		}
		return result
	}
	return nil
}

func (m *JsonData) GetByKey(key string) JSON {
	if m.IsObject() {
		value, ok := m.v.(map[string]any)[key]
		if !ok {
			return nil
		}
		switch value.(type) {
		case JsonData:
			v := value.(JsonData)
			return &v
		case *JsonData:
			return value.(*JsonData)
		case JSON:
			return value.(JSON).JsonData()
		case JSONObject:
			return value.(JSONObject).JsonData()
		case json.Number:
			if v, err := value.(json.Number).Int64(); err == nil {
				return &JsonData{t: reflect.Int, v: v, s: 8, p: m, pKey: key}
			}
		case string:
			v := value.(string)
			return &JsonData{t: reflect.String, v: v, s: len(v), p: m, pKey: key}
		case int, int8, int16, int32, int64: // to int
			v, _ := toInt(value)
			return &JsonData{t: reflect.Int, v: v, s: strconv.IntSize, p: m, pKey: key}
		case uint, uint8, uint16, uint32, uint64, uintptr: // to uint
			v, _ := toUInt(value)
			return &JsonData{t: reflect.Uint, v: v, s: strconv.IntSize, p: m, pKey: key}
		case []byte:
			v := value.([]byte)
			return &JsonData{t: consts.SLICE_BYTE, v: v, s: len(v), p: m, pKey: key}
		case float32, float64: // to float64
			sv, _ := toFloat64(value)
			if strings.Index(strconv.FormatFloat(sv, 'G', -1, 64), ".") != -1 {
				return &JsonData{t: reflect.Float64, v: sv, s: 8, p: m, pKey: key} // float64
			} else {
				return &JsonData{t: reflect.Int, v: int(sv), s: strconv.IntSize, p: m, pKey: key} // int
			}
		case bool:
			return &JsonData{t: reflect.Bool, v: value, s: 1, p: m, pKey: key}
		case []any:
			if v, ok := value.([]any); ok {
				return &JsonData{t: reflect.Slice, v: v, s: len(v), p: m, pKey: key}
			}
		case map[string]any:
			if v, ok := value.(map[string]any); ok {
				return &JsonData{t: reflect.Map, v: v, s: len(v), p: m, pKey: key}
			}
		default:
			return &JsonData{t: consts.NIL, v: nil, s: 0, p: m, pKey: key}
		}
	}
	return nil
}
