//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 依赖于 json-iterator 或 encoding/json 的JSON字节数组序列化
package json

import (
	"encoding/json"
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	jsoniter "github.com/json-iterator/go" //json-iterator or encoding/json
	"reflect"
	"strconv"
	"strings"
)

// BaseJSON JSON基础对象
//
// 当前对象直接操作
type BaseJSON interface {
	Size() int              //返回数据数量
	Type() GO_VALUE_TYPE    //当前对象数量类型
	Data() any              //返回原始数据
	JsonData() *JsonData    //返回原始JsonData数据结构
	SetValue(value any)     //设置值
	String() string         //返回 string 类型值
	Int() int               //返回 int 类型值, 把所有数字类型都转换 int 返回
	Int64() int64           //返回 uint 类型值, 把所有数字类型都转换 int64 返回
	UInt() uint             //返回 uint 类型值, 把所有数字类型都转换 uint 返回
	UInt64() uint64         //返回 uint 类型值, 把所有数字类型都转换 uint64 返回
	Bytes() []byte          //转换为 '[]byte' 并返回, 任意类型都将转换
	Float() float64         //返回 float64 类型值 把所有数字类型都转换 float64 返回
	Bool() bool             //返回 bool 类型值
	JSONObject() JSONObject //返回 JSONObject 对象类型
	JSONArray() JSONArray   //返回 JSONArray 对象类型
	ToJSONString() string   //转换为JSON字符串并返回
	IsString() bool         //当前对象是否为 string
	IsInt() bool            //当前对象是否为 int
	IsUInt() bool           //当前对象是否为 uint
	IsBytes() bool          //当前对象是否为 []byte
	IsFloat() bool          //当前对象是否为 float64
	IsBool() bool           //当前对象是否为 bool
	IsObject() bool         //当前对象是否为 JSONObject
	IsArray() bool          //当前对象是否为 JSONArray
	Clear()                 //清空所有数据，保留原始数据类型
	Free()                  //释放数据空间，且类型失效，当前对象不可用
}

// JSONArray 数组类型 index
//
// 根据下标操作数据，失败返回数据类型的默认值
type JSONArray interface {
	BaseJSON
	Add(value ...any)                      //添加任意类型数据
	SetByIndex(index int, value any)       //设置指定下标位置任意类型数据
	RemoveByIndex(index int)               //删除指定下标数据
	GetStringByIndex(index int) string     //根据下标返回 string 类型值
	GetIntByIndex(index int) int           //根据下标返回 int 类型值
	GetInt64ByIndex(index int) int64       //根据下标返回 int64 类型值
	GetUIntByIndex(index int) uint         //根据下标返回 uint 类型值
	GetUInt64ByIndex(index int) uint64     //根据下标返回 uint64 类型值
	GetBytesByIndex(index int) []byte      //根据下标返回 []byte 类型值
	GetFloatByIndex(index int) float64     //根据下标返回 float64 类型值
	GetBoolByIndex(index int) bool         //根据下标返回 bool 类型值
	GetArrayByIndex(index int) JSONArray   //根据下标返回 JSONArray 类型值
	GetObjectByIndex(index int) JSONObject //根据下标返回 JSONObject 类型值
	GetByIndex(index int) JSON             //根据下标返回 JSON 类型值
}

// JSONObject 对象类型 key=value
//
// 根据key操作数据，失败返回数据类型的默认值
type JSONObject interface {
	BaseJSON
	Set(key string, value any)            //设置或覆盖指定key，并设置新任意类型值
	RemoveByKey(key string)               //删除指定key数据
	GetStringByKey(key string) string     //根据 key 返回 string 类型值
	GetIntByKey(key string) int           //根据 key 返回 int 类型值
	GetInt64ByKey(key string) int64       //根据 key 返回 int64 类型值
	GetUIntByKey(key string) uint         //根据 key 返回 uint 类型值
	GetUInt64ByKey(key string) uint64     //根据 key 返回 uint64 类型值
	GetBytesByKey(key string) []byte      //根据 key 返回 []byte 类型值
	GetFloatByKey(key string) float64     //根据 key 返回 float64 类型值
	GetBoolByKey(key string) bool         //根据 key 返回 bool 类型值
	GetArrayByKey(key string) JSONArray   //根据 key 返回 JSONArray 类型值
	GetObjectByKey(key string) JSONObject //根据 key 返回 JSONObject 类型值
	GetByKey(key string) JSON             //根据 key 返回 JSON 类型值
	Keys() []string                       //返回当前对象所有 key
}

// JSON Object
type JSON interface {
	JSONArray
	JSONObject
}

// JsonData JSON 数据结构
type JsonData struct {
	T      GO_VALUE_TYPE // type
	S      int           // size
	V      any           // value
	pKey   string        //
	pIndex int           //
	p      *JsonData     // parent
}

// NewJSON 返回 JsonData 对象，JSONArray or JSONObject
//
//    []byte("{...}")
//    []byte("[...]")
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
				return &JsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		case reflect.Map:
			if v, ok := v.(map[string]any); ok {
				return &JsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	} else {
		println("NewJSON error:", err.Error())
	}
	return nil
}

// NewJSONArray 字节JSONArray / 数组 / 切片 转换
//
//    []byte("[...]")
//    []slice
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
		//转为[]any类型
		if byt, err := jsoniter.Marshal(value); err == nil {
			var v []any
			if err = jsoniter.Unmarshal(byt, &v); err == nil {
				return &JsonData{T: GO_VALUE_SLICE, V: v, S: len(v)}
			}
		}
	}
	return &JsonData{T: GO_VALUE_SLICE, V: make([]any, 0), S: 0}
}

// NewJSONObject 字节JSONObject / 结构 / JSONObject 转换
//
//    []byte("{...}")
//    struct
//    map[string][type]
func NewJSONObject(value any) JSONObject {
	if value != nil {
		switch value.(type) {
		// 如果 []byte 就必须是 字节JSONObject
		case []byte:
			if v := NewJSON(value.([]byte)); v != nil {
				return v.JSONObject()
			} else {
				return nil
			}
		// 如果 string 就必须是 字符串JSONObject
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
		//转为map[string]any类型
		if byt, err := jsoniter.Marshal(value); err == nil {
			var v map[string]any
			if err = jsoniter.Unmarshal(byt, &v); err == nil {
				return &JsonData{T: GO_VALUE_MAP, V: v, S: len(v)}
			}
		}
	}
	return &JsonData{T: GO_VALUE_MAP, V: make(map[string]any), S: 0}
}

func (m *JsonData) Size() int {
	return m.S
}

func (m *JsonData) Type() GO_VALUE_TYPE {
	return m.T
}

func (m *JsonData) Data() any {
	return m.V
}

func (m *JsonData) JsonData() *JsonData {
	return m
}

func (m *JsonData) Add(value ...any) {
	if m.IsArray() {
		tmp := make([]any, len(value))
		for i, v := range value {
			switch v.(type) {
			case []byte:
				if vv := NewJSON(v.([]byte)); v != nil {
					tmp[i] = vv.Data()
				} else {
					tmp[i] = value
				}
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

func (m *JsonData) SetByIndex(index int, value any) {
	if m.IsArray() && index < m.S {
		switch value.(type) {
		case []byte:
			if vv := NewJSON(value.([]byte)); vv != nil {
				m.V.([]any)[index] = vv.Data()
			} else {
				m.V.([]any)[index] = value
			}
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

func (m *JsonData) RemoveByIndex(index int) {
	if m.IsArray() && index >= 0 && index < m.S {
		v := m.V.([]any)
		m.V = append(v[:index], v[index+1:]...)
		m.S--
	}
}

func (m *JsonData) GetStringByIndex(index int) string {
	if m.IsArray() && index < m.S {
		if r, ok := m.V.([]any)[index].(string); ok {
			return r
		}
	}
	return ""
}

func (m *JsonData) GetIntByIndex(index int) int {
	if m.IsArray() && index < m.S {
		return m.toInt(m.V.([]any)[index])
	}
	return 0
}

func (m *JsonData) GetInt64ByIndex(index int) int64 {
	if m.IsArray() && index < m.S {
		return m.toInt64(m.V.([]any)[index])
	}
	return 0
}

func (m *JsonData) GetUIntByIndex(index int) uint {
	if m.IsArray() && index < m.S {
		return m.toUInt(m.V.([]any)[index])
	}
	return 0
}

func (m *JsonData) GetUInt64ByIndex(index int) uint64 {
	if m.IsArray() && index < m.S {
		return m.toUInt64(m.V.([]any)[index])
	}
	return 0
}

func (m *JsonData) GetBytesByIndex(index int) []byte {
	if m.IsArray() && index < m.S {
		return m.toBytes(m.V.([]any)[index])
	}
	return nil
}

func (m *JsonData) GetFloatByIndex(index int) float64 {
	if m.IsArray() && index < m.S {
		return m.toFloat64(m.V.([]any)[index])
	}
	return 0
}

func (m *JsonData) GetBoolByIndex(index int) bool {
	if m.IsArray() && index < m.S {
		s := m.V.([]any)[index]
		switch s.(type) {
		case bool:
			return s.(bool)
		}
	}
	return false
}

func (m *JsonData) GetArrayByIndex(index int) JSONArray {
	if m.IsArray() && index < m.S {
		if v, ok := m.V.([]any)[index].([]any); ok {
			return &JsonData{T: GO_VALUE_SLICE, V: v, S: len(v), p: m, pIndex: index}
		}
	}
	return nil
}

func (m *JsonData) GetObjectByIndex(index int) JSONObject {
	if m.IsArray() && index < m.S {
		if v, ok := m.V.([]any)[index].(map[string]any); ok {
			return &JsonData{T: GO_VALUE_MAP, V: v, S: len(v), p: m, pIndex: index}
		}
	}
	return nil
}

func (m *JsonData) GetByIndex(index int) JSON {
	if m.IsArray() && index < m.S {
		value := m.V.([]any)[index]
		switch value.(type) {
		case json.Number:
			if v, err := value.(json.Number).Int64(); err == nil {
				return &JsonData{T: GO_VALUE_INT, V: v, S: strconv.IntSize, p: m, pIndex: index}
			}
		case string:
			v := value.(string)
			return &JsonData{T: GO_VALUE_STRING, V: v, S: len(v), p: m, pIndex: index}
		case int, int8, int16, int32, int64:
			if v := m.GetIntByIndex(index); v != 0 {
				return &JsonData{T: GO_VALUE_INT, V: v, S: strconv.IntSize, p: m, pIndex: index}
			}
		case uint, uint8, uint16, uint32, uint64:
			if v := m.GetUIntByIndex(index); v != 0 {
				return &JsonData{T: GO_VALUE_UINT, V: v, S: strconv.IntSize, p: m, pIndex: index}
			}
		case []byte:
			if v := m.GetBytesByIndex(index); v != nil {
				return &JsonData{T: GO_VALUE_SLICE_BYTE, V: v, S: len(v), p: m, pIndex: index}
			}
		case float32, float64:
			//不带有 . 转为 int 类型
			v := m.toFloat64(value)
			if strings.Index(strconv.FormatFloat(v, 'G', -1, 64), ".") != -1 {
				return &JsonData{T: GO_VALUE_FLOAT64, V: v, S: 8, p: m, pIndex: index}
			} else {
				return &JsonData{T: GO_VALUE_INT, V: int(v), S: strconv.IntSize, p: m, pIndex: index}
			}
		case bool:
			return &JsonData{T: GO_VALUE_BOOL, V: m.GetBoolByIndex(index), S: 1, p: m, pIndex: index}
		case []any:
			if index < m.S {
				if v, ok := m.V.([]any)[index].([]any); ok {
					return &JsonData{T: GO_VALUE_SLICE, V: v, S: len(v), p: m, pIndex: index}
				}
			}
		case map[string]any:
			if v, ok := value.(map[string]any); ok {
				return &JsonData{T: GO_VALUE_MAP, V: v, S: len(v), p: m, pIndex: index}
			}
		}
	}
	return nil
}

func (m *JsonData) Set(key string, value any) {
	if m.IsObject() {
		switch value.(type) {
		case []byte:
			if vv := NewJSON(value.([]byte)); vv != nil {
				value = vv.Data()
			}
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

func (m *JsonData) modifyParentValue() {
	if m.p != nil {
		if m.p.IsArray() {
			m.p.SetByIndex(m.pIndex, m.V)
		} else if m.p.IsObject() {
			m.p.Set(m.pKey, m.V)
		}
	}
}

func (m *JsonData) SetValue(value any) {
	switch value.(type) {
	case json.Number:
		if v, err := value.(json.Number).Int64(); err == nil {
			m.T = GO_VALUE_INT
			m.V = v
			m.S = 8
			m.modifyParentValue()
		}
	case string:
		v := value.(string)
		m.T = GO_VALUE_STRING
		m.V = v
		m.S = len(v)
		m.modifyParentValue()
	case int, int8, int16, int32, int64:
		v := m.toInt(value)
		m.T = GO_VALUE_INT
		m.V = v
		m.S = strconv.IntSize
		m.modifyParentValue()
	case uint, uint8, uint16, uint32, uint64:
		v := m.toUInt(value)
		m.T = GO_VALUE_UINT
		m.V = v
		m.S = strconv.IntSize
		m.modifyParentValue()
	case []byte:
		m.T = GO_VALUE_SLICE_BYTE
		m.V = value.([]byte)
		m.S = len(value.([]byte))
		m.modifyParentValue()
	case float32, float64:
		sv := m.toFloat64(value)
		m.T = GO_VALUE_FLOAT64
		m.V = sv
		m.S = 8
		m.modifyParentValue()
	case bool:
		m.T = GO_VALUE_BOOL
		m.V = value
		m.S = 1
		m.modifyParentValue()
	case []any:
		m.T = GO_VALUE_SLICE
		m.V = value
		m.S = len(value.([]any))
		m.modifyParentValue()
	case map[string]any:
		m.T = GO_VALUE_MAP
		m.V = value
		m.S = len(value.(map[string]any))
		m.modifyParentValue()
	default:
		if v := NewJSONArray(value); v != nil {
			m.T = v.Type()
			m.V = v.Data()
			m.S = v.Size()
			m.modifyParentValue()
		} else if v := NewJSONObject(value); v != nil {
			m.T = v.Type()
			m.V = v.Data()
			m.S = v.Size()
			m.modifyParentValue()
		}
	}
}

func (m *JsonData) RemoveByKey(key string) {
	if m.IsObject() {
		if _, ok := m.V.(map[string]any)[key]; ok {
			delete(m.V.(map[string]any), key)
			m.S--
		}
	}
}

func (m *JsonData) GetStringByKey(key string) string {
	if m.IsObject() {
		if r, ok := m.V.(map[string]any)[key].(string); ok {
			return r
		}
	}
	return ""
}

func (m *JsonData) GetIntByKey(key string) int {
	if m.IsObject() {
		return m.toInt(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JsonData) GetInt64ByKey(key string) int64 {
	if m.IsObject() {
		return m.toInt64(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JsonData) GetUIntByKey(key string) uint {
	if m.IsObject() {
		return m.toUInt(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JsonData) GetUInt64ByKey(key string) uint64 {
	if m.IsObject() {
		return m.toUInt64(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JsonData) GetBytesByKey(key string) []byte {
	if m.IsObject() {
		return m.toBytes(m.V.(map[string]any)[key])
	}
	return nil
}

func (m *JsonData) GetFloatByKey(key string) float64 {
	if m.IsObject() {
		return m.toFloat64(m.V.(map[string]any)[key])
	}
	return 0
}

func (m *JsonData) GetBoolByKey(key string) bool {
	if m.IsObject() {
		if s, ok := m.V.(map[string]any)[key].(bool); ok {
			return s
		}
	}
	return false
}

func (m *JsonData) GetArrayByKey(key string) JSONArray {
	if m.IsObject() {
		if v, ok := m.V.(map[string]any)[key].([]any); ok {
			return &JsonData{T: GO_VALUE_SLICE, V: v, S: len(v), p: m, pKey: key}
		}
	}
	return nil
}

func (m *JsonData) GetObjectByKey(key string) JSONObject {
	if m.IsObject() {
		if v, ok := m.V.(map[string]any)[key].(map[string]any); ok {
			return &JsonData{T: GO_VALUE_MAP, V: v, S: len(v), p: m, pKey: key}
		}
	}
	return nil
}

func (m *JsonData) GetByKey(key string) JSON {
	if m.IsObject() {
		value := m.V.(map[string]any)[key]
		switch value.(type) {
		case json.Number:
			if v, err := value.(json.Number).Int64(); err == nil {
				return &JsonData{T: GO_VALUE_INT, V: v, S: 8, p: m, pKey: key}
			}
		case string:
			if v, ok := value.(string); ok {
				return &JsonData{T: GO_VALUE_STRING, V: v, S: len(v), p: m, pKey: key}
			}
		case int, int8, int16, int32, int64:
			if v := m.GetIntByKey(key); v != 0 {
				return &JsonData{T: GO_VALUE_INT, V: v, S: strconv.IntSize, p: m, pKey: key}
			}
		case uint, uint8, uint16, uint32, uint64:
			if v := m.GetUIntByKey(key); v != 0 {
				return &JsonData{T: GO_VALUE_UINT, V: v, S: strconv.IntSize, p: m, pKey: key}
			}
		case []byte:
			if v := m.GetBytesByKey(key); v != nil {
				return &JsonData{T: GO_VALUE_SLICE_BYTE, V: v, S: len(v), p: m, pKey: key}
			}
		case float32, float64:
			//不带有 . 转为 int 类型
			sv := m.toFloat64(value)
			if strings.Index(strconv.FormatFloat(sv, 'G', -1, 64), ".") != -1 {
				return &JsonData{T: GO_VALUE_FLOAT64, V: sv, S: 8, p: m, pKey: key}
			} else {
				return &JsonData{T: GO_VALUE_INT, V: int(sv), S: strconv.IntSize, p: m, pKey: key}
			}
		case bool:
			return &JsonData{T: GO_VALUE_BOOL, V: m.GetBytesByKey(key), S: 1, p: m, pKey: key}
		case []any:
			if v, ok := value.([]any); ok {
				return &JsonData{T: GO_VALUE_SLICE, V: v, S: len(v), p: m, pKey: key}
			}
		case map[string]any:
			if v, ok := value.(map[string]any); ok {
				return &JsonData{T: GO_VALUE_MAP, V: v, S: len(v), p: m, pKey: key}
			}
		}
	}
	return nil
}

func (m *JsonData) String() string {
	if m.IsString() {
		return m.V.(string)
	}
	return ""
}

func (m *JsonData) Int() int {
	return m.toInt(m.V)
}

func (m *JsonData) Int64() int64 {
	return m.toInt64(m.V)
}

func (m *JsonData) UInt() uint {
	return m.toUInt(m.V)
}

func (m *JsonData) UInt64() uint64 {
	return m.toUInt64(m.V)
}

func (m *JsonData) Bytes() []byte {
	return m.toBytes(m.V)
}

func (m *JsonData) Float() float64 {
	return m.toFloat64(m.V)
}

func (m *JsonData) Bool() bool {
	if m.IsBool() {
		switch m.V.(type) {
		case bool:
			return m.V.(bool)
		case []uint8:
			if v := m.V.([]uint8); len(v) > 0 {
				return v[0] != 0
			}
		}
	}
	return false
}

func (m *JsonData) JSONObject() JSONObject {
	if m.IsObject() {
		return m
	}
	return nil
}

func (m *JsonData) JSONArray() JSONArray {
	if m.IsArray() {
		return m
	}
	return nil
}

func (m *JsonData) Keys() []string {
	if m.IsObject() {
		var result []string
		for key, _ := range m.V.(map[string]any) {
			result = append(result, key)
		}
		return result
	}
	return nil
}

func (m *JsonData) ToJSONString() string {
	return string(m.Bytes())
}

func (m *JsonData) IsString() bool {
	return m.T == GO_VALUE_STRING
}

func (m *JsonData) IsInt() bool {
	return m.T == GO_VALUE_INT
}

func (m *JsonData) IsUInt() bool {
	return m.T == GO_VALUE_UINT
}

func (m *JsonData) IsBytes() bool {
	return m.T == GO_VALUE_SLICE_BYTE
}

func (m *JsonData) IsFloat() bool {
	return m.T == GO_VALUE_FLOAT64
}

func (m *JsonData) IsBool() bool {
	return m.T == GO_VALUE_BOOL
}

func (m *JsonData) IsObject() bool {
	return m.T == GO_VALUE_MAP
}

func (m *JsonData) IsArray() bool {
	return m.T == GO_VALUE_SLICE
}

func (m *JsonData) Clear() {
	if m.IsObject() {
		m.V = make(map[string]any, 0)
	} else if m.IsArray() {
		m.V = make([]any, 0)
	} else {
		m.V = nil
	}
	m.S = 0
}

func (m *JsonData) Free() {
	if m == nil {
		return
	}
	m.V = nil
	m.S = 0
	m.T = GO_VALUE_INVALID
}

func (m *JsonData) toBytes(s any) []byte {
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

func (m *JsonData) toFloat64(s any) float64 {
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

func (m *JsonData) toInt(s any) int {
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

func (m *JsonData) toUInt(s any) uint {
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

func (m *JsonData) toUInt64(s any) uint64 {
	switch s.(type) {
	case float32:
		return uint64(s.(float32))
	case float64:
		return uint64(s.(float64))
	case int:
		return uint64(s.(int))
	case int8:
		return uint64(s.(int8))
	case int16:
		return uint64(s.(int16))
	case int32:
		return uint64(s.(int32))
	case int64:
		return uint64(s.(int64))
	case uint:
		return uint64(s.(uint))
	case uint8:
		return uint64(s.(uint8))
	case uint16:
		return uint64(s.(uint16))
	case uint32:
		return uint64(s.(uint32))
	case uint64:
		return s.(uint64)
	}
	return 0
}

func (m *JsonData) toInt64(s any) int64 {
	switch s.(type) {
	case float32:
		return int64(s.(float32))
	case float64:
		return int64(s.(float64))
	case int:
		return int64(s.(int))
	case int8:
		return int64(s.(int8))
	case int16:
		return int64(s.(int16))
	case int32:
		return int64(s.(int32))
	case int64:
		return s.(int64)
	case uint:
		return int64(s.(uint))
	case uint8:
		return int64(s.(uint8))
	case uint16:
		return int64(s.(uint16))
	case uint32:
		return int64(s.(uint32))
	case uint64:
		return int64(s.(uint64))
	}
	return 0
}
