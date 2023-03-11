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
	"fmt"
	"github.com/energye/energy/consts"
	jsoniter "github.com/json-iterator/go"
	"reflect"
)

// callback IPC 监听回调函数
type callback struct {
	context  *contextCallback  //上下文参数回调
	argument *argumentCallback //参数回调
}

// contextCallback 带上下文的回调函数
type contextCallback struct {
	callback EmitContextCallback
}

// argumentCallback 带参数的回调函数
type argumentCallback struct {
	callback        *reflect.Value         //回调函数
	inArgumentType  []consts.GO_VALUE_TYPE //入参类型
	outArgumentType []consts.GO_VALUE_TYPE //出参类型
}

// ContextCallback 返回上下文参数回调
func (m *callback) ContextCallback() EmitContextCallback {
	if m.context == nil {
		return nil
	}
	return m.context.callback
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
	rv := m.callback
	var size = len(m.inArgumentType)
	var inArgsValues = make([]reflect.Value, size)
	for i := 0; i < size; i++ {
		inType := m.inArgumentType[i]
		if i < argsSize {
			argsValue := argsList.GetByIndex(i)
			if argsValue == nil {
				inArgsValues[i] = reflect.New(rv.Type().In(i)).Elem()
				continue
			}
			switch inType {
			case consts.GO_VALUE_STRING:
				inArgsValues[i] = reflect.ValueOf(argsValue.String())
			case consts.GO_VALUE_INT:
				inArgsValues[i] = reflect.ValueOf(argsValue.Int())
			case consts.GO_VALUE_INT8:
				inArgsValues[i] = reflect.ValueOf(int8(argsValue.Int()))
			case consts.GO_VALUE_INT16:
				inArgsValues[i] = reflect.ValueOf(int16(argsValue.Int()))
			case consts.GO_VALUE_INT32:
				inArgsValues[i] = reflect.ValueOf(int32(argsValue.Int()))
			case consts.GO_VALUE_INT64:
				inArgsValues[i] = reflect.ValueOf(int64(argsValue.Int()))
			case consts.GO_VALUE_UINT:
				inArgsValues[i] = reflect.ValueOf(uint(argsValue.Int()))
			case consts.GO_VALUE_UINT8:
				inArgsValues[i] = reflect.ValueOf(uint8(argsValue.Int()))
			case consts.GO_VALUE_UINT16:
				inArgsValues[i] = reflect.ValueOf(uint16(argsValue.Int()))
			case consts.GO_VALUE_UINT32:
				inArgsValues[i] = reflect.ValueOf(uint32(argsValue.Int()))
			case consts.GO_VALUE_UINT64:
				inArgsValues[i] = reflect.ValueOf(uint64(argsValue.Int()))
			case consts.GO_VALUE_UINTPTR:
				inArgsValues[i] = reflect.ValueOf(uintptr(argsValue.Int()))
			case consts.GO_VALUE_FLOAT32:
				inArgsValues[i] = reflect.ValueOf(float32(argsValue.Float()))
			case consts.GO_VALUE_FLOAT64:
				inArgsValues[i] = reflect.ValueOf(argsValue.Float())
			case consts.GO_VALUE_BOOL:
				inArgsValues[i] = reflect.ValueOf(argsValue.Bool())
			case consts.GO_VALUE_STRUCT:
				if argsValue.IsObject() {
					// struct
					if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
						v := reflect.New(rv.Type().In(i))
						if err := jsoniter.Unmarshal(jsonBytes, v.Interface()); err == nil {
							inArgsValues[i] = v.Elem()
							continue
						}
					}
				}
				inArgsValues[i] = reflect.New(rv.Type().In(i)).Elem()
			case consts.GO_VALUE_MAP:
				if argsValue.IsObject() {
					inArgsType := rv.Type().In(i)
					// map key=string : value != interface
					inArgsKind := inArgsType.Elem().Kind()
					if inArgsKind != reflect.Interface {
						if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
							vv := reflect.New(rv.Type().In(i))
							if err := jsoniter.Unmarshal(jsonBytes, vv.Interface()); err == nil {
								inArgsValues[i] = vv.Elem()
								continue
							}
						}
						inArgsValues[i] = reflect.New(inArgsType).Elem()
					} else {
						inArgsValues[i] = reflect.ValueOf(argsValue.Data())
					}
				} else {
					inArgsValues[i] = reflect.New(rv.Type().In(i)).Elem()
				}
			case consts.GO_VALUE_SLICE:
				if argsValue.IsArray() {
					inArgsType := rv.Type().In(i)
					// slick value != interface
					inArgsKind := inArgsType.Elem().Kind()
					if inArgsKind != reflect.Interface {
						if jsonBytes := argsValue.Bytes(); jsonBytes != nil {
							vv := reflect.New(rv.Type().In(i))
							if err := jsoniter.Unmarshal(jsonBytes, vv.Interface()); err == nil {
								inArgsValues[i] = vv.Elem()
								continue
							}
						}
						inArgsValues[i] = reflect.New(inArgsType).Elem()
					} else {
						inArgsValues[i] = reflect.ValueOf(argsValue.Data())
					}
				} else {
					inArgsValues[i] = reflect.New(rv.Type().In(i)).Elem()
				}
			default:
				inArgsValues[i] = reflect.New(rv.Type().In(i)).Elem()
			}
		} else {
			inArgsValues[i] = reflect.New(rv.Type().In(i)).Elem()
		}
	}
	// call
	resultValues := rv.Call(inArgsValues)
	// call result
	fmt.Println("resultValues:", resultValues)
}
