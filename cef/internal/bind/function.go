//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSFunction 类型实现

package bind

import (
	"github.com/energye/energy/pkgs/json"
	jsoniter "github.com/json-iterator/go"
	"reflect"
)

type JSFunction interface {
	JSValue
	AsFunction() JSFunction
	// Invoke 调用函数
	//	入参: 以参数列表形式传入参数
	//		入参如果参数类型或数量不匹配这些参数将以类型的默认值传入
	//		nil 无入参
	//	出参: 以参数列表形式返回参数
	//		无返回值返回nil
	Invoke(argumentList json.JSONArray) (resultArgument json.JSONArray)
}

type jsFunction struct {
	V8Value
	rv *reflect.Value
}

func (m *jsFunction) AsFunction() JSFunction {
	if m.IsFunction() {
		return m
	}
	return nil
}

func (m *jsFunction) Invoke(argumentList json.JSONArray) (resultArgument json.JSONArray) {
	resultArgument = nil
	if m.IsFunction() {
		var (
			argsSize     int
			inArgsValues []reflect.Value
		)
		if argumentList != nil {
			argsSize = argumentList.Size()
		}
		rt := m.rv.Type()
		inArgsCount := rt.NumIn()
		inArgsValues = make([]reflect.Value, inArgsCount)
		for i := 0; i < inArgsCount; i++ {
			inType := rt.In(i)
			if i < argsSize {
				argsValue := argumentList.GetByIndex(i)
				if argsValue != nil {
					switch inType.Kind() {
					case reflect.String:
						inArgsValues[i] = reflect.ValueOf(argsValue.String())
					case reflect.Int:
						inArgsValues[i] = reflect.ValueOf(argsValue.Int())
					case reflect.Int8:
						inArgsValues[i] = reflect.ValueOf(int8(argsValue.Int()))
					case reflect.Int16:
						inArgsValues[i] = reflect.ValueOf(int16(argsValue.Int()))
					case reflect.Int32:
						inArgsValues[i] = reflect.ValueOf(int32(argsValue.Int()))
					case reflect.Int64:
						inArgsValues[i] = reflect.ValueOf(int64(argsValue.Int()))
					case reflect.Uint:
						inArgsValues[i] = reflect.ValueOf(uint(argsValue.Int()))
					case reflect.Uint8:
						inArgsValues[i] = reflect.ValueOf(uint8(argsValue.Int()))
					case reflect.Uint16:
						inArgsValues[i] = reflect.ValueOf(uint16(argsValue.Int()))
					case reflect.Uint32:
						inArgsValues[i] = reflect.ValueOf(uint32(argsValue.Int()))
					case reflect.Uint64:
						inArgsValues[i] = reflect.ValueOf(uint64(argsValue.Int()))
					case reflect.Uintptr:
						inArgsValues[i] = reflect.ValueOf(uintptr(argsValue.Int()))
					case reflect.Float32:
						inArgsValues[i] = reflect.ValueOf(float32(argsValue.Float()))
					case reflect.Float64:
						inArgsValues[i] = reflect.ValueOf(argsValue.Float())
					case reflect.Bool:
						inArgsValues[i] = reflect.ValueOf(argsValue.Bool())
					case reflect.Struct:
						if argsValue.IsObject() {
							// struct
							if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
								v := reflect.New(inType)
								if err := jsoniter.Unmarshal(jsonBytes, v.Interface()); err == nil {
									inArgsValues[i] = v.Elem()
								}
							}
						}
					case reflect.Map:
						if argsValue.IsObject() {
							// map key=string : value != interface
							if inType.Elem().Kind() != reflect.Interface {
								if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
									vv := reflect.New(inType)
									if err := jsoniter.Unmarshal(jsonBytes, vv.Interface()); err == nil {
										inArgsValues[i] = vv.Elem()
									}
								}
							} else {
								inArgsValues[i] = reflect.ValueOf(argsValue.Data())
							}
						}
					case reflect.Slice:
						if argsValue.IsArray() {
							// slice value != interface
							if inType.Elem().Kind() != reflect.Interface {
								if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
									vv := reflect.New(inType)
									if err := jsoniter.Unmarshal(jsonBytes, vv.Interface()); err == nil {
										inArgsValues[i] = vv.Elem()
									}
								}
							} else {
								inArgsValues[i] = reflect.ValueOf(argsValue.Data())
							}
						}
					}
				}
			}
			if !inArgsValues[i].IsValid() {
				inArgsValues[i] = reflect.New(inType).Elem()
			}
		}
		// call
		resultValues := m.rv.Call(inArgsValues)
		if len(resultValues) > 0 {
			// call result
			resultArgument = json.NewJSONArray(nil)
			for _, result := range resultValues {
				res := result.Interface()
				switch res.(type) {
				case error:
					resultArgument.Add(res.(error).Error())
				default:
					resultArgument.Add(res)
				}
			}
		}
	}
	return
}
