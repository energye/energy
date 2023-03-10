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
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strconv"
)

type IBaseJSON interface {
	Size() int
	Type() GO_VALUE_TYPE
	Data() any
	String() string
	Int() int
	UInt() uint
	Bytes() []byte
	Float() float64
	Bool() bool
	Map() IJSONMap
	Array() IJSONArray
	IsString() bool
	IsInt() bool
	IsUInt() bool
	IsBytes() bool
	IsFloat() bool
	IsBool() bool
	IsMap() bool
	IsArray() bool
}

type IJSONArray interface {
	IBaseJSON
	Add(value any)
	RemoveByIndex(index int)
	StringByIndex(index int) string
	IntByIndex(index int) int
	UIntByIndex(index int) uint
	BytesByIndex(index int) []byte
	FloatByIndex(index int) float64
	BoolByIndex(index int) bool
	ArrayByIndex(index int) IJSONArray
	MapByIndex(index int) IJSONMap
	GetByIndex(index int) IJSON
}

type IJSONMap interface {
	IBaseJSON
	Set(key, value string)
	RemoveByKey(key string)
	StringByKey(key string) string
	IntByKey(key string) int
	UIntByKey(key string) uint
	BytesByKey(key string) []byte
	FloatByKey(key string) float64
	BoolByKey(key string) bool
	ArrayByKey(key string) IJSONArray
	MapByKey(key string) IJSONMap
	GetByKey(key string) IJSON
	Keys() []string
}

// IJSON IPC参数 数组形式
type IJSON interface {
	IJSONArray
	IJSONMap
}

type JSON struct {
	T GO_VALUE_TYPE // type
	S int           // size
	V any           // value
}

// NewJSON 返回 JSON 对象，IJSONArray or IJSONMap
func NewJSON(data []byte) IJSON {
	if data == nil {
		return nil
	}
	var v any
	if err := jsoniter.Unmarshal(data, &v); err == nil {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice:
			if v, ok := v.([]any); ok {
				return &JSON{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		case reflect.Map:
			if v, ok := v.(map[string]any); ok {
				return &JSON{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return nil
}

// NewJSONArray 字节JSONArray/数组/切片 转换
func NewJSONArray(value any) IJSONArray {
	if value != nil {
		// 如果 []byte 就必须是 字节JSONArray
		switch value.(type) {
		case []byte:
			if v := NewJSON(value.([]byte)); v != nil {
				return v.Array()
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
				return &JSON{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		}
	}
	return nil
}

// NewJSONMap 字节JSONObject/结构/Map 转换
func NewJSONMap(value any) IJSONMap {
	if value != nil {
		// 如果 []byte 就必须是 字节JSONObject
		switch value.(type) {
		case []byte:
			if v := NewJSON(value.([]byte)); v != nil {
				return v.Map()
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
				return &JSON{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return nil
}

func (m *JSON) Size() int {
	return m.S
}

func (m *JSON) Type() GO_VALUE_TYPE {
	return m.T
}

func (m *JSON) Data() any {
	return m.V
}

func (m *JSON) Add(value any) {
	if m.IsArray() {
		m.V = append(m.V.([]any), value)
		m.S++
	}
}

func (m *JSON) RemoveByIndex(index int) {
	if m.IsArray() && index >= 0 && index < m.S {
		v := m.V.([]any)
		m.V = append(v[:index], v[index+1:]...)
		m.S--
	}
}

func (m *JSON) StringByIndex(index int) string {
	if m.IsArray() && index < m.S {
		if r, ok := m.V.([]any)[index].(string); ok {
			return r
		}
	}
	return ""
}

func (m *JSON) IntByIndex(index int) int {
	if m.IsArray() && index < m.S {
		return m.toInt(m.V.([]any)[index])
	}
	return 0
}

func (m *JSON) UIntByIndex(index int) uint {
	if m.IsArray() && index < m.S {
		return m.toUInt(m.V.([]any)[index])
	}
	return 0
}

func (m *JSON) BytesByIndex(index int) []byte {
	if m.IsArray() && index < m.S {
		return m.toBytes(m.V.([]any)[index])
	}
	return nil
}

func (m *JSON) FloatByIndex(index int) float64 {
	if m.IsArray() && index < m.S {
		return m.toFloat64(m.V.([]any)[index])
	}
	return 0
}

func (m *JSON) BoolByIndex(index int) bool {
	if m.IsArray() && index < m.S {
		s := m.V.([]any)[index]
		switch s.(type) {
		case bool:
			return s.(bool)
		}
	}
	return false
}

func (m *JSON) ArrayByIndex(index int) IJSONArray {
	if m.IsArray() && index < m.S {
		if v, ok := m.V.([]any)[index].([]any); ok {
			return &JSON{T: GO_VALUE_SLICE, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *JSON) MapByIndex(index int) IJSONMap {
	if m.IsArray() && index < m.S {
		if v, ok := m.V.([]any)[index].(map[string]any); ok {
			return &JSON{T: GO_VALUE_MAP, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *JSON) GetByIndex(index int) IJSON {
	if m.IsArray() && index < m.S {
		s := m.V.([]any)[index]
		switch s.(type) {
		case string:
			if r, ok := s.(string); ok {
				return &JSON{T: GO_VALUE_STRING, V: r, S: len(r)}
			}
		case int, int8, int16, int32, int64:
			if s := m.IntByIndex(index); s != 0 {
				return &JSON{T: GO_VALUE_INT, V: s, S: strconv.IntSize}
			}
		case uint, uint8, uint16, uint32, uint64:
			if s := m.UIntByIndex(index); s != 0 {
				return &JSON{T: GO_VALUE_UINT, V: s, S: strconv.IntSize}
			}
		case []byte:
			if s := m.BytesByIndex(index); s != nil {
				return &JSON{T: GO_VALUE_SLICE_BYTE, V: s, S: len(s)}
			}
		case float32, float64:
			if s := m.FloatByIndex(index); s != 0 {
				return &JSON{T: GO_VALUE_FLOAT64, V: s, S: 8}
			}
		case bool:
			return &JSON{T: GO_VALUE_BOOL, V: m.BoolByIndex(index), S: 1}
		case []any:
			if index < m.S {
				if v, ok := m.V.([]any)[index].([]any); ok {
					return &JSON{T: GO_VALUE_SLICE, V: v, S: len(v)}
				}
			}
		case map[string]any:
			if v, ok := s.(map[string]any); ok {
				return &JSON{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return nil
}

func (m *JSON) Set(key, value string) {
	if m.IsMap() {
		m.V.(map[string]any)[key] = value
		m.S++
	}
}

func (m *JSON) RemoveByKey(key string) {
	if m.IsMap() {
		if _, ok := m.V.(map[string]any)[key]; ok {
			delete(m.V.(map[string]any), key)
			m.S--
		}
	}
}

func (m *JSON) StringByKey(key string) string {
	if m.IsMap() {
		if r, ok := m.V.(map[string]any)[key].(string); ok {
			return r
		}
	}
	return ""
}

func (m *JSON) IntByKey(key string) int {
	if m.IsMap() {
		return m.toInt(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JSON) UIntByKey(key string) uint {
	if m.IsMap() {
		return m.toUInt(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JSON) BytesByKey(key string) []byte {
	if m.IsMap() {
		return m.toBytes(m.V.(map[string]any)[key])
	}
	return nil
}

func (m *JSON) FloatByKey(key string) float64 {
	if m.IsMap() {
		return m.toFloat64(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JSON) BoolByKey(key string) bool {
	if m.IsMap() {
		if s, ok := m.V.(map[string]any)[key].(bool); ok {
			return s
		}
	}
	return false
}

func (m *JSON) ArrayByKey(key string) IJSONArray {
	if m.IsMap() {
		if v, ok := m.V.(map[string]any)[key].([]any); ok {
			return &JSON{T: GO_VALUE_SLICE, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *JSON) MapByKey(key string) IJSONMap {
	if m.IsMap() {
		if v, ok := m.V.(map[string]any)[key].(map[string]any); ok {
			return &JSON{T: GO_VALUE_MAP, V: v, S: len(v)}
		}
	}
	return nil
}

func (m *JSON) GetByKey(key string) IJSON {
	if m.IsMap() {
		s := m.V.(map[string]any)[key]
		switch s.(type) {
		case string:
			if r, ok := s.(string); ok {
				return &JSON{T: GO_VALUE_STRING, V: r, S: len(r)}
			}
		case int, int8, int16, int32, int64:
			if s := m.IntByKey(key); s != 0 {
				return &JSON{T: GO_VALUE_INT, V: s, S: strconv.IntSize}
			}
		case uint, uint8, uint16, uint32, uint64:
			if s := m.UIntByKey(key); s != 0 {
				return &JSON{T: GO_VALUE_UINT, V: s, S: strconv.IntSize}
			}
		case []byte:
			if s := m.BytesByKey(key); s != nil {
				return &JSON{T: GO_VALUE_SLICE_BYTE, V: s, S: len(s)}
			}
		case float32, float64:
			if s := m.FloatByKey(key); s != 0 {
				return &JSON{T: GO_VALUE_FLOAT64, V: s, S: 8}
			}
		case bool:
			return &JSON{T: GO_VALUE_BOOL, V: m.BytesByKey(key), S: 1}
		case []any:
			if v, ok := m.V.(map[string]any)[key].([]any); ok {
				return &JSON{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		case map[string]any:
			if v, ok := s.(map[string]any); ok {
				return &JSON{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return nil
}

func (m *JSON) String() string {
	if m.IsString() {
		return m.V.(string)
	}
	return ""
}

func (m *JSON) Int() int {
	return m.toInt(m.V)
}

func (m *JSON) UInt() uint {
	return m.toUInt(m.V)
}

func (m *JSON) Bytes() []byte {
	if m.IsBytes() {
		return m.V.([]byte)
	}
	return nil
}

func (m *JSON) Float() float64 {
	return m.toFloat64(m.V)
}

func (m *JSON) Bool() bool {
	if m.IsBool() {
		return m.V.(bool)
	}
	return false
}

func (m *JSON) Map() IJSONMap {
	if m.IsMap() {
		return m
	}
	return nil
}

func (m *JSON) Array() IJSONArray {
	if m.IsArray() {
		return m
	}
	return nil
}

func (m *JSON) toBytes(s any) []byte {
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

func (m *JSON) Keys() []string {
	if m.IsMap() {
		var result []string
		for key, _ := range m.V.(map[string]any) {
			result = append(result, key)
		}
		return result
	}
	return nil
}

func (m *JSON) IsString() bool {
	return m.T == GO_VALUE_STRING
}

func (m *JSON) IsInt() bool {
	return m.T == GO_VALUE_INT
}

func (m *JSON) IsUInt() bool {
	return m.T == GO_VALUE_UINT
}

func (m *JSON) IsBytes() bool {
	return m.T == GO_VALUE_SLICE_BYTE
}

func (m *JSON) IsFloat() bool {
	return m.T == GO_VALUE_FLOAT64
}

func (m *JSON) IsBool() bool {
	return m.T == GO_VALUE_BOOL
}

func (m *JSON) IsMap() bool {
	return m.T == GO_VALUE_MAP
}

func (m *JSON) IsArray() bool {
	return m.T == GO_VALUE_SLICE
}

func (m *JSON) toFloat64(s any) float64 {
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

func (m *JSON) toInt(s any) int {
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

func (m *JSON) toUInt(s any) uint {
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
