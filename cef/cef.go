//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/cef"
	"github.com/energye/lcl/lcl"
	"reflect"
	"time"
)

type tExecuteScriptMessage struct {
	Id     int32  `json:"id"`
	Script string `json:"script"`
}

type tExecuteScriptResultMessage struct {
	Id    int32  `json:"id"`
	Data  string `json:"data"`
	Error string `json:"error"`
}

type tObjectFile struct {
	Name         string `json:"name"`
	LastModified int64  `json:"last_modified"`
	Size         uint32 `json:"size"`
}

// v8ValueToJSON
func v8ValueToJSON(v cef.ICefv8Value) any {
	if v == nil || v.IsUndefined() || v.IsNull() {
		return nil
	} else if v.IsBool() {
		return v.GetBoolValue()
	} else if v.IsInt() {
		return v.GetIntValue()
	} else if v.IsUInt() {
		return v.GetUIntValue()
	} else if v.IsDouble() {
		return v.GetDoubleValue()
	} else if v.IsString() {
		return v.GetStringValue()
	} else if v.IsDate() {
		return v.GetDateValue().ToTime().Format(time.RFC3339)
	} else if v.IsArray() {
		length := v.GetArrayLength()
		arr := make([]any, length)
		for i := int32(0); i < length; i++ {
			item := v.GetValueByIndex(i)
			arr[i] = v8ValueToJSON(item)
			item.Release()
		}
		return arr
	} else if v.IsObject() {
		keys := lcl.NewStringList()
		defer keys.Free()
		if v.GetKeys(keys) == 0 {
			return nil
		}
		obj := make(map[string]any, keys.Count())
		for i := int32(0); i < keys.Count(); i++ {
			key := keys.Strings(i)
			val := v.GetValueByKey(key)
			obj[key] = v8ValueToJSON(val)
			val.Release()
		}
		return obj
	}
	return nil
}

// IsNil Check if any value is nil.
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return rv.IsNil()
	default:
		return false
	}
}
