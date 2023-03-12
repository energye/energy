//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package json

import (
	"encoding/json"
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strconv"
	"strings"
)

type BaseJSON interface {
	Size() int
	Type() GO_VALUE_TYPE
	Data() any
	String() string
	Int() int
	UInt() uint
	Bytes() []byte
	Float() float64
	Bool() bool
	JSONObject() JSONObject
	JSONArray() JSONArray
	ToJSONString() string
	IsString() bool
	IsInt() bool
	IsUInt() bool
	IsBytes() bool
	IsFloat() bool
	IsBool() bool
	IsObject() bool
	IsArray() bool
	Clear()
	Free()
}

type JSONArray interface {
	BaseJSON
	Add(value ...any)
	SetByIndex(index int, value any)
	RemoveByIndex(index int)
	GetStringByIndex(index int) string
	GetIntByIndex(index int) int
	GetUIntByIndex(index int) uint
	GetBytesByIndex(index int) []byte
	GetFloatByIndex(index int) float64
	GetBoolByIndex(index int) bool
	GetArrayByIndex(index int) JSONArray
	GetObjectByIndex(index int) JSONObject
	GetByIndex(index int) JSON
}

type JSONObject interface {
	BaseJSON
	Set(key string, value any)
	RemoveByKey(key string)
	GetStringByKey(key string) string
	GetIntByKey(key string) int
	GetUIntByKey(key string) uint
	GetBytesByKey(key string) []byte
	GetFloatByKey(key string) float64
	GetBoolByKey(key string) bool
	GetArrayByKey(key string) JSONArray
	GetObjectByKey(key string) JSONObject
	GetByKey(key string) JSON
	Keys() []string
}

// JSON Object
type JSON interface {
	JSONArray
	JSONObject
}

type jsonData struct {
	T GO_VALUE_TYPE // type
	S int           // size
	V any           // value
}

// NewJSON 返回 jsonData 对象，JSONArray or JSONObject
func NewJSON(data []byte) JSON {
	if data == nil {
		return nil
	}
	var v any
	if err := jsoniter.Unmarshal(data, &v); err == nil {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice:
			if v, ok := v.([]any); ok {
				return &jsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		case reflect.Map:
			if v, ok := v.(map[string]any); ok {
				return &jsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return nil
}

// NewJSONArray 字节JSONArray / 数组 / 切片 转换
func NewJSONArray(value any) JSONArray {
	if value != nil {
		// 如果 []byte 就必须是 字节JSONArray
		switch value.(type) {
		case []byte:
			if v := NewJSON(value.([]byte)); v != nil {
				return v.JSONArray()
			} else {
				return nil
			}
		}
		rv := reflect.ValueOf(value)
		if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
			return nil
		}
		//目的是为了转为any类型
		if byt, err := jsoniter.Marshal(value); err == nil {
			var v []any
			if err = jsoniter.Unmarshal(byt, &v); err == nil {
				return &jsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		}
	}
	return &jsonData{T: GO_VALUE_SLICE, V: make([]any, 0), S: 0}
}

// NewJSONObject 字节JSONObject / 结构 / JSONObject 转换
func NewJSONObject(value any) JSONObject {
	if value != nil {
		// 如果 []byte 就必须是 字节JSONObject
		switch value.(type) {
		case []byte:
			if v := NewJSON(value.([]byte)); v != nil {
				return v.JSONObject()
			} else {
				return nil
			}
		}
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if kind != reflect.Map && kind != reflect.Struct {
			return nil
		}
		//目的是为了转为any类型
		if byt, err := jsoniter.Marshal(value); err == nil {
			var v map[string]any
			if err = jsoniter.Unmarshal(byt, &v); err == nil {
				return &jsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return &jsonData{T: GO_VALUE_MAP, V: make(map[string]any), S: 0}
}

func (m *jsonData) Size() int {
	return m.S
}

func (m *jsonData) Type() GO_VALUE_TYPE {
	return m.T
}

func (m *jsonData) Data() any {
	return m.V
}

func (m *jsonData) Add(value ...any) {
	if m.IsArray() {
		tmp := make([]any, len(value))
		for i, v := range value {
			switch v.(type) {
			case JSON:
				tmp[i] = v.(JSON).Data()
			case JSONObject:
				tmp[i] = v.(JSONObject).Data()
			case JSONArray:
				tmp[i] = v.(JSONArray).Data()
			default:
				tmp[i] = v
			}
		}
		m.V = append(m.V.([]any), tmp...)
		m.S += len(value)
	}
}

func (m *jsonData) SetByIndex(index int, value any) {
	if m.IsArray() && index < m.S {
		switch value.(type) {
		case JSON:
			m.V.([]any)[index] = value.(JSON).Data()
		case JSONObject:
			m.V.([]any)[index] = value.(JSONObject).Data()
		case JSONArray:
			m.V.([]any)[index] = value.(JSONArray).Data()
		default:
			m.V.([]any)[index] = value
		}
	}
}

func (m *jsonData) RemoveByIndex(index int) {
	if m.IsArray() && index >= 0 && index < m.S {
		v := m.V.([]any)
		m.V = append(v[:index], v[index+1:]...)
		m.S--
	}
}

func (m *jsonData) GetStringByIndex(index int) string {
	if m.IsArray() && index < m.S {
		if r, ok := m.V.([]any)[index].(string); ok {
			return r
		}
	}
	return ""
}

func (m *jsonData) GetIntByIndex(index int) int {
	if m.IsArray() && index < m.S {
		return m.toInt(m.V.([]any)[index])
	}
	return 0
}

func (m *jsonData) GetUIntByIndex(index int) uint {
	if m.IsArray() && index < m.S {
		return m.toUInt(m.V.([]any)[index])
	}
	return 0
}

func (m *jsonData) GetBytesByIndex(index int) []byte {
	if m.IsArray() && index < m.S {
		return m.toBytes(m.V.([]any)[index])
	}
	return nil
}

func (m *jsonData) GetFloatByIndex(index int) float64 {
	if m.IsArray() && index < m.S {
		return m.toFloat64(m.V.([]any)[index])
	}
	return 0
}

func (m *jsonData) GetBoolByIndex(index int) bool {
	if m.IsArray() && index < m.S {
		s := m.V.([]any)[index]
		switch s.(type) {
		case bool:
			return s.(bool)
		}
	}
	return false
}

func (m *jsonData) GetArrayByIndex(index int) JSONArray {
	if m.IsArray() && index < m.S {
		if v, ok := m.V.([]any)[index].([]any); ok {
			return &jsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *jsonData) GetObjectByIndex(index int) JSONObject {
	if m.IsArray() && index < m.S {
		if v, ok := m.V.([]any)[index].(map[string]any); ok {
			return &jsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *jsonData) GetByIndex(index int) JSON {
	if m.IsArray() && index < m.S {
		s := m.V.([]any)[index]
		switch s.(type) {
		case json.Number:
			if v, err := s.(json.Number).Int64(); err == nil {
				return &jsonData{T: GO_VALUE_INT, V: v, S: strconv.IntSize}
			}
		case string:
			if r, ok := s.(string); ok {
				return &jsonData{T: GO_VALUE_STRING, V: r, S: len(r)}
			}
		case int, int8, int16, int32, int64:
			if s := m.GetIntByIndex(index); s != 0 {
				return &jsonData{T: GO_VALUE_INT, V: s, S: strconv.IntSize}
			}
		case uint, uint8, uint16, uint32, uint64:
			if s := m.GetUIntByIndex(index); s != 0 {
				return &jsonData{T: GO_VALUE_UINT, V: s, S: strconv.IntSize}
			}
		case []byte:
			if s := m.GetBytesByIndex(index); s != nil {
				return &jsonData{T: GO_VALUE_SLICE_BYTE, V: s, S: len(s)}
			}
		case float32, float64:
			//不带有 . 转为 int 类型
			sv := m.toFloat64(s)
			if strings.Index(strconv.FormatFloat(sv, 'G', -1, 64), ".") != -1 {
				return &jsonData{T: GO_VALUE_FLOAT64, V: sv, S: 8}
			} else {
				return &jsonData{T: GO_VALUE_INT, V: sv, S: strconv.IntSize}
			}
		case bool:
			return &jsonData{T: GO_VALUE_BOOL, V: m.GetBoolByIndex(index), S: 1}
		case []any:
			if index < m.S {
				if v, ok := m.V.([]any)[index].([]any); ok {
					return &jsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
				}
			}
		case map[string]any:
			if v, ok := s.(map[string]any); ok {
				return &jsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return nil
}

func (m *jsonData) Set(key string, value any) {
	if m.IsObject() {
		switch value.(type) {
		case JSON:
			value = value.(JSON).Data()
		case JSONObject:
			value = value.(JSONObject).Data()
		case JSONArray:
			value = value.(JSONArray).Data()
		}
		m.V.(map[string]any)[key] = value
		m.S++
	}
}

func (m *jsonData) RemoveByKey(key string) {
	if m.IsObject() {
		if _, ok := m.V.(map[string]any)[key]; ok {
			delete(m.V.(map[string]any), key)
			m.S--
		}
	}
}

func (m *jsonData) GetStringByKey(key string) string {
	if m.IsObject() {
		if r, ok := m.V.(map[string]any)[key].(string); ok {
			return r
		}
	}
	return ""
}

func (m *jsonData) GetIntByKey(key string) int {
	if m.IsObject() {
		return m.toInt(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *jsonData) GetUIntByKey(key string) uint {
	if m.IsObject() {
		return m.toUInt(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *jsonData) GetBytesByKey(key string) []byte {
	if m.IsObject() {
		return m.toBytes(m.V.(map[string]any)[key])
	}
	return nil
}

func (m *jsonData) GetFloatByKey(key string) float64 {
	if m.IsObject() {
		return m.toFloat64(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *jsonData) GetBoolByKey(key string) bool {
	if m.IsObject() {
		if s, ok := m.V.(map[string]any)[key].(bool); ok {
			return s
		}
	}
	return false
}

func (m *jsonData) GetArrayByKey(key string) JSONArray {
	if m.IsObject() {
		if v, ok := m.V.(map[string]any)[key].([]any); ok {
			return &jsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *jsonData) GetObjectByKey(key string) JSONObject {
	if m.IsObject() {
		if v, ok := m.V.(map[string]any)[key].(map[string]any); ok {
			return &jsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *jsonData) GetByKey(key string) JSON {
	if m.IsObject() {
		s := m.V.(map[string]any)[key]
		switch s.(type) {
		case json.Number:
			if v, err := s.(json.Number).Int64(); err == nil {
				return &jsonData{T: GO_VALUE_INT, V: v, S: strconv.IntSize}
			}
		case string:
			if r, ok := s.(string); ok {
				return &jsonData{T: GO_VALUE_STRING, V: r, S: len(r)}
			}
		case int, int8, int16, int32, int64:
			if s := m.GetIntByKey(key); s != 0 {
				return &jsonData{T: GO_VALUE_INT, V: s, S: strconv.IntSize}
			}
		case uint, uint8, uint16, uint32, uint64:
			if s := m.GetUIntByKey(key); s != 0 {
				return &jsonData{T: GO_VALUE_UINT, V: s, S: strconv.IntSize}
			}
		case []byte:
			if s := m.GetBytesByKey(key); s != nil {
				return &jsonData{T: GO_VALUE_SLICE_BYTE, V: s, S: len(s)}
			}
		case float32, float64:
			//不带有 . 转为 int 类型
			sv := m.toFloat64(s)
			if strings.Index(strconv.FormatFloat(sv, 'G', -1, 64), ".") != -1 {
				return &jsonData{T: GO_VALUE_FLOAT64, V: sv, S: 8}
			} else {
				return &jsonData{T: GO_VALUE_INT, V: int(sv), S: strconv.IntSize}
			}
		case bool:
			return &jsonData{T: GO_VALUE_BOOL, V: m.GetBytesByKey(key), S: 1}
		case []any:
			if v, ok := m.V.(map[string]any)[key].([]any); ok {
				return &jsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		case map[string]any:
			if v, ok := s.(map[string]any); ok {
				return &jsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return nil
}

func (m *jsonData) String() string {
	if m.IsString() {
		return m.V.(string)
	}
	return ""
}

func (m *jsonData) Int() int {
	return m.toInt(m.V)
}

func (m *jsonData) UInt() uint {
	return m.toUInt(m.V)
}

func (m *jsonData) Bytes() []byte {
	return m.toBytes(m.V)
}

func (m *jsonData) Float() float64 {
	return m.toFloat64(m.V)
}

func (m *jsonData) Bool() bool {
	if m.IsBool() {
		return m.V.(bool)
	}
	return false
}

func (m *jsonData) JSONObject() JSONObject {
	if m.IsObject() {
		return m
	}
	return nil
}

func (m *jsonData) JSONArray() JSONArray {
	if m.IsArray() {
		return m
	}
	return nil
}

func (m *jsonData) toBytes(s any) []byte {
	switch s.(type) {
	case []byte:
		return s.([]byte)
	case string:
		return []byte(s.(string))
	case bool:
		return []byte{common.BoolToByte(s.(bool))}
	case float32:
		return common.Float32ToBytes(s.(float32))
	case float64:
		return common.Float64ToBytes(s.(float64))
	case int:
		return common.IntToBytes(s.(int))
	case int8:
		return common.Int8ToBytes(s.(int8))
	case int16:
		return common.Int16ToBytes(s.(int16))
	case int32:
		return common.Int32ToBytes(s.(int32))
	case int64:
		return common.Int64ToBytes(s.(int64))
	case uint:
		return common.UIntToBytes(s.(uint))
	case uint8:
		return common.UInt8ToBytes(s.(uint8))
	case uint16:
		return common.UInt16ToBytes(s.(uint16))
	case uint32:
		return common.UInt32ToBytes(s.(uint32))
	case uint64:
		return common.UInt64ToBytes(s.(uint64))
	default: // slice or map or other
		if r, err := jsoniter.Marshal(s); err == nil {
			return r
		}
	}
	return nil
}

func (m *jsonData) Keys() []string {
	if m.IsObject() {
		var result []string
		for key, _ := range m.V.(map[string]any) {
			result = append(result, key)
		}
		return result
	}
	return nil
}

func (m *jsonData) ToJSONString() string {
	return string(m.Bytes())
}

func (m *jsonData) IsString() bool {
	return m.T == GO_VALUE_STRING
}

func (m *jsonData) IsInt() bool {
	return m.T == GO_VALUE_INT
}

func (m *jsonData) IsUInt() bool {
	return m.T == GO_VALUE_UINT
}

func (m *jsonData) IsBytes() bool {
	return m.T == GO_VALUE_SLICE_BYTE
}

func (m *jsonData) IsFloat() bool {
	return m.T == GO_VALUE_FLOAT64
}

func (m *jsonData) IsBool() bool {
	return m.T == GO_VALUE_BOOL
}

func (m *jsonData) IsObject() bool {
	return m.T == GO_VALUE_MAP
}

func (m *jsonData) IsArray() bool {
	return m.T == GO_VALUE_SLICE
}

func (m *jsonData) Clear() {
	if m.IsObject() {
		m.V = make(map[string]any, 0)
	} else if m.IsArray() {
		m.V = make([]any, 0)
	} else {
		m.V = nil
	}
	m.S = 0
}

func (m *jsonData) Free() {
	m.V = nil
	m.S = 0
	m.T = GO_VALUE_INVALID
}

func (m *jsonData) toFloat64(s any) float64 {
	switch s.(type) {
	case float32:
		return float64(s.(float32))
	case float64:
		return s.(float64)
	case int:
		return float64(s.(int))
	case int8:
		return float64(s.(int8))
	case int16:
		return float64(s.(int16))
	case int32:
		return float64(s.(int32))
	case int64:
		return float64(s.(int64))
	case uint:
		return float64(s.(uint))
	case uint8:
		return float64(s.(uint8))
	case uint16:
		return float64(s.(uint16))
	case uint32:
		return float64(s.(uint32))
	case uint64:
		return float64(s.(uint64))
	}
	return 0
}

func (m *jsonData) toInt(s any) int {
	switch s.(type) {
	case float32:
		return int(s.(float32))
	case float64:
		return int(s.(float64))
	case int:
		return s.(int)
	case int8:
		return int(s.(int8))
	case int16:
		return int(s.(int16))
	case int32:
		return int(s.(int32))
	case int64:
		return int(s.(int64))
	case uint:
		return int(s.(uint))
	case uint8:
		return int(s.(uint8))
	case uint16:
		return int(s.(uint16))
	case uint32:
		return int(s.(uint32))
	case uint64:
		return int(s.(uint64))
	}
	return 0
}

func (m *jsonData) toUInt(s any) uint {
	switch s.(type) {
	case float32:
		return uint(s.(float32))
	case float64:
		return uint(s.(float64))
	case int:
		return uint(s.(int))
	case int8:
		return uint(s.(int8))
	case int16:
		return uint(s.(int16))
	case int32:
		return uint(s.(int32))
	case int64:
		return uint(s.(int64))
	case uint:
		return s.(uint)
	case uint8:
		return uint(s.(uint8))
	case uint16:
		return uint(s.(uint16))
	case uint32:
		return uint(s.(uint32))
	case uint64:
		return uint(s.(uint64))
	}
	return 0
}
