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
// JSON

package json

import (
	"encoding/json"
	"github.com/cyber-xxm/energy/v2/common"
	. "github.com/cyber-xxm/energy/v2/consts"
	"reflect"
	"strconv"
)

// JSON object
type JSON interface {
	Size() int                  //返回数据数量
	Type() reflect.Kind         //当前对象数量类型
	Data() interface{}          //返回原始数据
	SetValue(value interface{}) //设置值
	String() string             //返回 string 类型值
	Int() int                   //返回 int 类型值, 把所有数字类型都转换 int 返回
	Int64() int64               //返回 uint 类型值, 把所有数字类型都转换 int64 返回
	UInt() uint                 //返回 uint 类型值, 把所有数字类型都转换 uint 返回
	UInt64() uint64             //返回 uint 类型值, 把所有数字类型都转换 uint64 返回
	Bytes() []byte              //转换为 '[]byte' 并返回, 任意类型都将转换
	Float() float64             //返回 float64 类型值 把所有数字类型都转换 float64 返回
	Bool() bool                 //返回 bool 类型值
	JSONObject() JSONObject     //返回 JSONObject 对象类型
	JSONArray() JSONArray       //返回 JSONArray 对象类型
	JsonData() *JsonData        //JSON Data
	ToJSONString() string       //转换为JSON字符串并返回
	IsString() bool             //当前对象是否为 string
	IsInt() bool                //当前对象是否为 int
	IsUInt() bool               //当前对象是否为 uint
	IsBytes() bool              //当前对象是否为 []byte
	IsFloat() bool              //当前对象是否为 float64
	IsBool() bool               //当前对象是否为 bool
	IsObject() bool             //当前对象是否为 JSONObject
	IsArray() bool              //当前对象是否为 JSONArray
	Clear()                     //清空所有数据，保留原始数据类型
	Free()                      //释放数据空间，且类型失效，当前对象不可用
}

// JsonData
//
//	Data structure
type JsonData struct {
	t      reflect.Kind // type
	s      int          // size
	v      interface{}  // value
	pKey   string       // object parent key
	pIndex int          // array parent index
	p      *JsonData    // parent
	d      bool         // is recursion
}

// NewJsonData
//
//	create Json
//	t: data type
//	s: data size
//	v: data value
func NewJsonData(t reflect.Kind, s int, v interface{}) *JsonData {
	return &JsonData{t: t, s: s, v: v}
}

// NewJSON
//
//		return JSON Object, JSONArray or JSONObject
//	 data:
//	   []byte("{...}") object
//	   []byte("[...]") array
func NewJSON(data []byte) JSON {
	if data == nil {
		return nil
	}
	var v interface{}
	if err := json.Unmarshal(data, &v); err == nil {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice:
			if v, ok := v.([]interface{}); ok {
				return &JsonData{t: reflect.Slice, v: v, s: len(v)}
			}
		case reflect.Map:
			if v, ok := v.(map[string]interface{}); ok {
				return &JsonData{t: reflect.Map, v: v, s: len(v)}
			}
		}
	}
	return nil
}

func (m *JsonData) Size() int {
	return m.s
}

func (m *JsonData) Type() reflect.Kind {
	return m.t
}

func (m *JsonData) Data() interface{} {
	return m.v
}

func (m *JsonData) JsonData() *JsonData {
	return m
}

func (m *JsonData) String() string {
	if m.IsString() {
		return m.v.(string)
	}
	return ""
}

func (m *JsonData) Int() (r int) {
	r, _ = toInt(m.v)
	return
}

func (m *JsonData) Int64() (r int64) {
	r, _ = toInt64(m.v)
	return
}

func (m *JsonData) UInt() (r uint) {
	r, _ = toUInt(m.v)
	return
}

func (m *JsonData) UInt64() (r uint64) {
	r, _ = toUInt64(m.v)
	return
}

func (m *JsonData) Bytes() []byte {
	// TODO 需要改进, 直接返回 [] byte
	return toBytes(m.ConvertToData())
	//return m.toBytes(m.v)
}

func (m *JsonData) Float() (r float64) {
	r, _ = toFloat64(m.v)
	return
}

func (m *JsonData) Bool() (r bool) {
	r, _ = toBool(m.v)
	return
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

func (m *JsonData) modifyParentValue() {
	if m.p != nil {
		if m.p.IsArray() {
			m.p.SetByIndex(m.pIndex, m.v)
		} else if m.p.IsObject() {
			m.p.Set(m.pKey, m.v)
		}
	}
}

func (m *JsonData) SetValue(value interface{}) {
	switch value.(type) {
	case JsonData:
		v := value.(JsonData)
		m.t = v.t
		m.v = v.v
		m.s = v.s
		m.p = v.p
		m.pIndex = v.pIndex
		m.pKey = v.pKey
		m.modifyParentValue()
	case *JsonData:
		v := value.(*JsonData)
		m.t = v.t
		m.v = v.v
		m.s = v.s
		m.p = v.p
		m.pIndex = v.pIndex
		m.pKey = v.pKey
		m.modifyParentValue()
	case JSONObject:
		v := value.(JSONObject).JsonData()
		m.t = v.t
		m.v = v.v
		m.s = v.s
		m.p = v.p
		m.pIndex = v.pIndex
		m.pKey = v.pKey
		m.modifyParentValue()
	case JSONArray:
		v := value.(JSONArray).JsonData()
		m.t = v.t
		m.v = v.v
		m.s = v.s
		m.p = v.p
		m.pIndex = v.pIndex
		m.pKey = v.pKey
		m.modifyParentValue()
	case json.Number:
		if v, err := value.(json.Number).Int64(); err == nil {
			m.t = reflect.Int
			m.v = v
			m.s = 8
			m.modifyParentValue()
		}
	case string:
		v := value.(string)
		m.t = reflect.String
		m.v = v
		m.s = len(v)
		m.modifyParentValue()
	case int, int8, int16, int32, int64:
		m.t = reflect.Int
		m.v, _ = toInt(value)
		m.s = strconv.IntSize
		m.modifyParentValue()
	case uint, uint8, uint16, uint32, uint64, uintptr:
		m.t = reflect.Uint
		m.v, _ = toUInt(value)
		m.s = strconv.IntSize
		m.modifyParentValue()
	case []byte:
		m.t = SLICE_BYTE
		m.v = value.([]byte)
		m.s = len(value.([]byte))
		m.modifyParentValue()
	case float32, float64:
		m.t = reflect.Float64
		m.v, _ = toFloat64(value)
		m.s = 8
		m.modifyParentValue()
	case bool:
		m.t = reflect.Bool
		m.v = value
		m.s = 1
		m.modifyParentValue()
	case []interface{}:
		m.t = reflect.Slice
		m.v = value
		m.s = len(value.([]interface{}))
		m.modifyParentValue()
	case map[string]interface{}:
		m.t = reflect.Map
		m.v = value
		m.s = len(value.(map[string]interface{}))
		m.modifyParentValue()
	default:
		if v := NewJSONArray(value); v != nil {
			m.t = v.Type()
			m.v = v.Data()
			m.s = v.Size()
			m.modifyParentValue()
		} else if v := NewJSONObject(value); v != nil {
			m.t = v.Type()
			m.v = v.Data()
			m.s = v.Size()
			m.modifyParentValue()
		}
	}
}

func (m *JsonData) ToJSONString() string {
	return string(m.Bytes())
}

// ConvertToData to map / slice / value
func (m *JsonData) ConvertToData() interface{} {
	if m.IsObject() {
		result := make(map[string]interface{}, m.s)
		for k, _ := range m.v.(map[string]interface{}) {
			v := m.GetByKey(k)
			if v == nil {
				result[k] = nil
				continue
			}
			if v.IsObject() || v.IsArray() {
				if v.JsonData().d {
					result[k] = nil
					continue
				}
				v.JsonData().d = true
				result[k] = v.JsonData().ConvertToData()
				v.JsonData().d = false
			} else {
				result[k] = v.JsonData().v
			}
		}
		return result
	} else if m.IsArray() {
		result := make([]interface{}, m.s, m.s)
		for i, _ := range m.v.([]interface{}) {
			v := m.GetByIndex(i)
			if v == nil {
				result[i] = nil
				continue
			}
			if v.IsObject() || v.IsArray() {
				if v.JsonData().d {
					result[i] = nil
					continue
				}
				v.JsonData().d = true
				result[i] = v.JsonData().ConvertToData()
				v.JsonData().d = false
			} else {
				result[i] = v.JsonData().v
			}
		}
		return result
	} else {
		return m.v
	}
}

func (m *JsonData) IsString() bool {
	return m.t == reflect.String
}

func (m *JsonData) IsInt() bool {
	return m.t == reflect.Int
}

func (m *JsonData) IsUInt() bool {
	return m.t == reflect.Uint
}

func (m *JsonData) IsBytes() bool {
	return m.t == SLICE_BYTE
}

func (m *JsonData) IsFloat() bool {
	return m.t == reflect.Float64
}

func (m *JsonData) IsBool() bool {
	return m.t == reflect.Bool
}

func (m *JsonData) IsObject() bool {
	return m.t == reflect.Map
}

func (m *JsonData) IsArray() bool {
	return m.t == reflect.Slice
}

func (m *JsonData) Clear() {
	if m.IsObject() {
		m.v = make(map[string]interface{}, 0)
	} else if m.IsArray() {
		m.v = make([]interface{}, 0)
	} else {
		m.v = nil
	}
	m.s = 0
}

func (m *JsonData) Free() {
	if m == nil {
		return
	}
	m.v = nil
	m.s = 0
	m.t = reflect.Invalid
}

func toBytes(s interface{}) []byte {
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
	default:
		if r, err := json.Marshal(s); err == nil {
			return r
		}
	}
	return nil
}

func toFloat64(s interface{}) (result float64, ok bool) {
	ok = true
	switch s.(type) {
	case float32:
		result = float64(s.(float32))
	case float64:
		result = s.(float64)
	case int:
		result = float64(s.(int))
	case int8:
		result = float64(s.(int8))
	case int16:
		result = float64(s.(int16))
	case int32:
		result = float64(s.(int32))
	case int64:
		result = float64(s.(int64))
	case uint:
		result = float64(s.(uint))
	case uint8:
		result = float64(s.(uint8))
	case uint16:
		result = float64(s.(uint16))
	case uint32:
		result = float64(s.(uint32))
	case uint64:
		result = float64(s.(uint64))
	default:
		ok = false
	}
	return
}

func toInt(s interface{}) (result int, ok bool) {
	ok = true
	switch s.(type) {
	case float32:
		result = int(s.(float32))
	case float64:
		result = int(s.(float64))
	case int:
		result = s.(int)
	case int8:
		result = int(s.(int8))
	case int16:
		result = int(s.(int16))
	case int32:
		result = int(s.(int32))
	case int64:
		result = int(s.(int64))
	case uint:
		result = int(s.(uint))
	case uint8:
		result = int(s.(uint8))
	case uint16:
		result = int(s.(uint16))
	case uint32:
		result = int(s.(uint32))
	case uint64:
		result = int(s.(uint64))
	default:
		ok = false
	}
	return
}

func toInt64(s interface{}) (result int64, ok bool) {
	ok = true
	switch s.(type) {
	case float32:
		result = int64(s.(float32))
	case float64:
		result = int64(s.(float64))
	case int:
		result = int64(s.(int))
	case int8:
		result = int64(s.(int8))
	case int16:
		result = int64(s.(int16))
	case int32:
		result = int64(s.(int32))
	case int64:
		result = s.(int64)
	case uint:
		result = int64(s.(uint))
	case uint8:
		result = int64(s.(uint8))
	case uint16:
		result = int64(s.(uint16))
	case uint32:
		result = int64(s.(uint32))
	case uint64:
		result = int64(s.(uint64))
	default:
		ok = false
	}
	return
}

func toUInt(s interface{}) (result uint, ok bool) {
	ok = true
	switch s.(type) {
	case float32:
		result = uint(s.(float32))
	case float64:
		result = uint(s.(float64))
	case int:
		result = uint(s.(int))
	case int8:
		result = uint(s.(int8))
	case int16:
		result = uint(s.(int16))
	case int32:
		result = uint(s.(int32))
	case int64:
		result = uint(s.(int64))
	case uint:
		result = s.(uint)
	case uint8:
		result = uint(s.(uint8))
	case uint16:
		result = uint(s.(uint16))
	case uint32:
		result = uint(s.(uint32))
	case uint64:
		result = uint(s.(uint64))
	default:
		ok = false
	}
	return
}

func toUInt64(s interface{}) (result uint64, ok bool) {
	ok = true
	switch s.(type) {
	case float32:
		result = uint64(s.(float32))
	case float64:
		result = uint64(s.(float64))
	case int:
		result = uint64(s.(int))
	case int8:
		result = uint64(s.(int8))
	case int16:
		result = uint64(s.(int16))
	case int32:
		result = uint64(s.(int32))
	case int64:
		result = uint64(s.(int64))
	case uint:
		result = uint64(s.(uint))
	case uint8:
		result = uint64(s.(uint8))
	case uint16:
		result = uint64(s.(uint16))
	case uint32:
		result = uint64(s.(uint32))
	case uint64:
		result = s.(uint64)
	default:
		ok = false
	}
	return
}

func toBool(s interface{}) (result, ok bool) {
	ok = true
	switch s.(type) {
	case bool:
		result = s.(bool)
	case []uint8:
		if v := s.([]uint8); len(v) > 0 {
			result = v[0] != 0
		}
	case string:
		result = s.(string) != ""
	default:
		if v, vok := toFloat64(s); vok {
			result = v > 0
		} else {
			ok = false
		}
	}
	return
}

func isBaseType(v interface{}) bool {
	switch v.(type) {
	case string, float32, float64, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return true
	}
	return false
}
