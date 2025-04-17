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
// JSONArray

package json

import (
	"encoding/json"
	"github.com/cyber-xxm/energy/v2/consts"
	"reflect"
	"strconv"
	"strings"
)

// JSONArray
//
//	Manipulate data according to subscript, failure to return the default value of the data type
type JSONArray interface {
	JSON
	Add(value ...interface{})                // Add data of any type
	SetByIndex(index int, value interface{}) // Set any type of data at the specified subscript position
	RemoveByIndex(index int)                 // remove data for index
	GetStringByIndex(index int) string       // return string data for index
	GetIntByIndex(index int) int             // return int data for index
	GetInt64ByIndex(index int) int64         // return int64 data for index
	GetUIntByIndex(index int) uint           // return uint data for index
	GetUInt64ByIndex(index int) uint64       // return uint64 data for index
	GetBytesByIndex(index int) []byte        // return []byte data for index
	GetFloatByIndex(index int) float64       // return float64 data for index
	GetBoolByIndex(index int) bool           // return bool data for index
	GetArrayByIndex(index int) JSONArray     // return JSONArray data for index
	GetObjectByIndex(index int) JSONObject   // return JSONObject data for index
	GetByIndex(index int) JSON               // return JSON data for index
}

// NewJSONArray
//
//	byte JSONArray, array & slice convert
//	value:
//	  []byte("[...]")
//	  []slice
func NewJSONArray(value interface{}) JSONArray {
	if value != nil {
		// 如果 []byte 就必须是 字节JSONArray
		switch value.(type) {
		case []byte:
			if v := NewJSON(value.([]byte)); v != nil {
				return v.JSONArray()
			} else {
				return nil
			}
		case string:
			if v := NewJSON([]byte(value.(string))); v != nil {
				return v.JSONArray()
			} else {
				return nil
			}
		}
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			kind = rv.Elem().Kind()
		}
		if kind != reflect.Slice && kind != reflect.Array {
			return nil
		}
		//转为[]interface{}类型
		if byt, err := json.Marshal(value); err == nil {
			var v []interface{}
			if err = json.Unmarshal(byt, &v); err == nil {
				return &JsonData{t: reflect.Slice, v: v, s: len(v)}
			}
		}
	}
	return &JsonData{t: reflect.Slice, v: make([]interface{}, 0), s: 0}
}

func (m *JsonData) Add(value ...interface{}) {
	if m.IsArray() {
		tmp := make([]interface{}, len(value), len(value))
		for i, v := range value {
			switch v.(type) {
			case []byte:
				if vv := NewJSON(v.([]byte)); v != nil {
					tmp[i] = vv.Data()
				} else {
					tmp[i] = value
				}
			case JsonData:
				tmp[i] = v.(JsonData)
			case *JsonData:
				tmp[i] = v.(*JsonData)
			case JSON:
				tmp[i] = v.(JSON).JsonData()
			case JSONObject:
				tmp[i] = v.(JSONObject).JsonData()
			case JSONArray:
				tmp[i] = v.(JSONArray).JsonData()
			default:
				if isBaseType(v) {
					tmp[i] = v
				} else {
					rv := reflect.ValueOf(v)
					kind := rv.Kind()
					if kind == reflect.Ptr {
						kind = rv.Elem().Kind()
					}
					if kind == reflect.Struct {
						// struct -> map
						if d, err := json.Marshal(v); err == nil {
							var vv map[string]interface{}
							if err = json.Unmarshal(d, &vv); err == nil {
								tmp[i] = vv // json object
							}
						}
					} else if kind == reflect.Slice || kind == reflect.Array {
						// slice -> array
						if d, err := json.Marshal(v); err == nil {
							var vv []interface{}
							if err = json.Unmarshal(d, &vv); err == nil {
								tmp[i] = vv // json array
							}
						}
					} else if kind == reflect.Map {
						tmp[i] = v
					} else {
						tmp[i] = nil
					}
				}
			}
		}
		m.v = append(m.v.([]interface{}), tmp...)
		m.s += len(value)
	}
}

func (m *JsonData) SetByIndex(index int, value interface{}) {
	if m.IsArray() && index < m.s {
		switch value.(type) {
		case []byte:
			if vv := NewJSON(value.([]byte)); vv != nil {
				m.v.([]interface{})[index] = vv.JsonData()
			} else {
				m.v.([]interface{})[index] = value
			}
		case JsonData:
			m.v.([]interface{})[index] = value.(JsonData)
		case *JsonData:
			m.v.([]interface{})[index] = value.(*JsonData)
		case JSON:
			m.v.([]interface{})[index] = value.(JSON).JsonData()
		case JSONObject:
			m.v.([]interface{})[index] = value.(JSONObject).JsonData()
		case JSONArray:
			m.v.([]interface{})[index] = value.(JSONArray).JsonData()
		default:
			if !isBaseType(value) {
				rv := reflect.ValueOf(value)
				kind := rv.Kind()
				if kind == reflect.Ptr {
					kind = rv.Elem().Kind()
				}
				if kind == reflect.Struct {
					if d, err := json.Marshal(value); err == nil {
						var v map[string]interface{}
						if err = json.Unmarshal(d, &v); err == nil {
							value = v // json object
						}
					}
				} else if kind == reflect.Slice || kind == reflect.Array {
					if d, err := json.Marshal(value); err == nil {
						var v []interface{}
						if err = json.Unmarshal(d, &v); err == nil {
							value = v // json array
						}
					}
				}
			}
			m.v.([]interface{})[index] = value
		}
	}
}

func (m *JsonData) RemoveByIndex(index int) {
	if m.IsArray() && index >= 0 && index < m.s {
		v := m.v.([]interface{})
		m.v = append(v[:index], v[index+1:]...)
		m.s--
	}
}

func (m *JsonData) GetStringByIndex(index int) string {
	if m.IsArray() && index < m.s {
		v := m.v.([]interface{})[index]
		switch v.(type) {
		case *JsonData:
			return v.(*JsonData).String()
		default:
			if r, ok := v.(string); ok {
				return r
			}
		}
	}
	return ""
}

func (m *JsonData) GetIntByIndex(index int) int {
	if m.IsArray() && index < m.s {
		v := m.v.([]interface{})[index]
		switch v.(type) {
		case *JsonData:
			return v.(*JsonData).Int()
		default:
			r, _ := toInt(v)
			return r
		}
	}
	return 0
}

func (m *JsonData) GetInt64ByIndex(index int) int64 {
	if m.IsArray() && index < m.s {
		v := m.v.([]interface{})[index]
		switch v.(type) {
		case *JsonData:
			return v.(*JsonData).Int64()
		default:
			r, _ := toInt64(v)
			return r
		}
	}
	return 0
}

func (m *JsonData) GetUIntByIndex(index int) uint {
	if m.IsArray() && index < m.s {
		v := m.v.([]interface{})[index]
		switch v.(type) {
		case *JsonData:
			return v.(*JsonData).UInt()
		default:
			r, _ := toUInt(v)
			return r
		}
	}
	return 0
}

func (m *JsonData) GetUInt64ByIndex(index int) uint64 {
	if m.IsArray() && index < m.s {
		v := m.v.([]interface{})[index]
		switch v.(type) {
		case *JsonData:
			return v.(*JsonData).UInt64()
		default:
			r, _ := toUInt64(v)
			return r
		}
	}
	return 0
}

func (m *JsonData) GetBytesByIndex(index int) []byte {
	if m.IsArray() && index < m.s {
		return toBytes(m.Data())
	}
	return nil
}

func (m *JsonData) GetFloatByIndex(index int) float64 {
	if m.IsArray() && index < m.s {
		v := m.v.([]interface{})[index]
		switch v.(type) {
		case *JsonData:
			return v.(*JsonData).Float()
		default:
			r, _ := toFloat64(v)
			return r
		}
	}
	return 0
}

func (m *JsonData) GetBoolByIndex(index int) bool {
	if m.IsArray() && index < m.s {
		v := m.v.([]interface{})[index]
		switch v.(type) {
		case *JsonData:
			return v.(*JsonData).Bool()
		default:
			r, _ := toBool(v)
			return r
		}
	}
	return false
}

func (m *JsonData) GetArrayByIndex(index int) JSONArray {
	return m.GetByIndex(index).JSONArray()
}

func (m *JsonData) GetObjectByIndex(index int) JSONObject {
	return m.GetByIndex(index).JSONObject()
}

func (m *JsonData) GetByIndex(index int) JSON {
	if m.IsArray() && index < m.s {
		value := m.v.([]interface{})[index]
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
				return &JsonData{t: reflect.Int, v: v, s: 8, p: m, pIndex: index}
			}
		case string:
			v := value.(string)
			return &JsonData{t: reflect.String, v: v, s: len(v), p: m, pIndex: index}
		case int, int8, int16, int32, int64:
			v, _ := toInt(value)
			return &JsonData{t: reflect.Int, v: v, s: strconv.IntSize, p: m, pIndex: index}
		case uint, uint8, uint16, uint32, uint64, uintptr:
			v, _ := toUInt(value)
			return &JsonData{t: reflect.Uint, v: v, s: strconv.IntSize, p: m, pIndex: index}
		case []byte:
			if v := m.GetBytesByIndex(index); v != nil {
				return &JsonData{t: consts.SLICE_BYTE, v: v, s: len(v), p: m, pIndex: index}
			}
		case float32, float64:
			//不带有 . 转为 int 类型
			v, _ := toFloat64(value)
			if strings.Index(strconv.FormatFloat(v, 'G', -1, 64), ".") != -1 {
				return &JsonData{t: reflect.Float64, v: v, s: 8, p: m, pIndex: index}
			} else {
				return &JsonData{t: reflect.Int, v: int(v), s: strconv.IntSize, p: m, pIndex: index}
			}
		case bool:
			return &JsonData{t: reflect.Bool, v: value, s: 1, p: m, pIndex: index}
		case []interface{}:
			if v, ok := value.([]interface{}); ok {
				return &JsonData{t: reflect.Slice, v: v, s: len(v), p: m, pIndex: index}
			}
		case map[string]interface{}:
			if v, ok := value.(map[string]interface{}); ok {
				return &JsonData{t: reflect.Map, v: v, s: len(v), p: m, pIndex: index}
			}
		default:
			return &JsonData{t: consts.NIL, v: nil, s: 0, p: m, pIndex: index}
		}
	}
	return nil
}
