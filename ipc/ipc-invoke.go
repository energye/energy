//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ipc

import (
	"github.com/energye/energy/pkgs/json"
	jsoniter "github.com/json-iterator/go"
	"reflect"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// callback IPC 监听回调函数
type callback struct {
	context  *contextCallback  //上下文参数回调
	argument *argumentCallback //参数回调
}

// contextCallback 带上下文的回调函数
type contextCallback struct {
	callback emitContextCallback
}

// argumentCallback 带参数的回调函数
type argumentCallback struct {
	callback *reflect.Value //回调函数
}

// ContextCallback 返回上下文参数回调
func (m *callback) ContextCallback() *contextCallback {
	if m.context == nil {
		return nil
	}
	return m.context
}

// Invoke 调用函数
func (m *contextCallback) Invoke(context IContext) {
	// call
	m.callback(context)
	resultValues := context.Replay().Result()
	if len(resultValues) > 0 {
		// call result
		resultArgument := json.NewJSONArray(nil)
		for _, result := range resultValues {
			switch result.(type) {
			case error:
				resultArgument.Add(result.(error).Error())
			default:
				resultArgument.Add(result)
			}
		}
		// result bytes
		context.Result(resultArgument.Bytes())
		return
	}
	context.Result(nil)
}

// ArgumentCallback 参数回调
func (m *callback) ArgumentCallback() *argumentCallback {
	if m.argument == nil {
		return nil
	}
	return m.argument
}

// Invoke 调用函数
func (m *argumentCallback) Invoke(context IContext) {
	argsList := context.ArgumentList()
	argsSize := argsList.Size()
	rt := m.callback.Type()
	inArgsCount := rt.NumIn()
	var inArgsValues = make([]reflect.Value, inArgsCount)
	for i := 0; i < inArgsCount; i++ {
		inType := rt.In(i)
		if i < argsSize {
			argsValue := argsList.GetByIndex(i)
			if argsValue == nil {
				inArgsValues[i] = reflect.New(inType).Elem()
				continue
			}
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
							continue
						}
					}
				}
				inArgsValues[i] = reflect.New(inType).Elem()
			case reflect.Map:
				if argsValue.IsObject() {
					// map key=string : value != interface
					if inType.Elem().Kind() != reflect.Interface {
						if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
							vv := reflect.New(inType)
							if err := jsoniter.Unmarshal(jsonBytes, vv.Interface()); err == nil {
								inArgsValues[i] = vv.Elem()
								continue
							}
						}
						inArgsValues[i] = reflect.New(inType).Elem()
					} else {
						inArgsValues[i] = reflect.ValueOf(argsValue.Data())
					}
				} else {
					inArgsValues[i] = reflect.New(inType).Elem()
				}
			case reflect.Slice:
				if argsValue.IsArray() {
					// slick value != interface
					if inType.Elem().Kind() != reflect.Interface {
						if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
							vv := reflect.New(inType)
							if err := jsoniter.Unmarshal(jsonBytes, vv.Interface()); err == nil {
								inArgsValues[i] = vv.Elem()
								continue
							}
						}
						inArgsValues[i] = reflect.New(inType).Elem()
					} else {
						inArgsValues[i] = reflect.ValueOf(argsValue.Data())
					}
				} else {
					inArgsValues[i] = reflect.New(inType).Elem()
				}
			default:
				inArgsValues[i] = reflect.New(inType).Elem()
			}
		} else {
			inArgsValues[i] = reflect.New(inType).Elem()
		}
	}
	// call
	resultValues := m.callback.Call(inArgsValues)
	if len(resultValues) > 0 {
		// call result
		resultArgument := json.NewJSONArray(nil)
		for _, result := range resultValues {
			res := result.Interface()
			switch res.(type) {
			case error:
				resultArgument.Add(res.(error).Error())
			default:
				resultArgument.Add(res)
			}
		}
		// result bytes
		context.Result(resultArgument.Bytes())
		return
	}
	// result nil
	context.Result(nil)
}
