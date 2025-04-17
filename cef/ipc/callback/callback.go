//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package callback

import (
	goJSON "encoding/json"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
	"reflect"
)

var argumentChannelType = reflect.TypeOf(new(IChannel)).Elem()

// EmitContextCallback IPC context callback
type EmitContextCallback func(context context.IContext)

// IChannel
//
//	The channel ID of the parameter type callback function
//	Used for listening to events and receiving parameters from the event channel source
type IChannel interface {
	BrowserId() int32  //Receive Browser Window ID
	ChannelId() string //Receive Channel ID
}

// Callback IPC Listening callback function
//  1. Callback Function - Context Mode
//  2. Callback Function - Argument Mode
type Callback struct {
	IsAsync  bool              // JS ipc.emit mode
	Context  *ContextCallback  // 1 Context
	Argument *ArgumentCallback // 2 Argument
}

// ContextCallback
//
//	Callback function with context
type ContextCallback struct {
	Callback EmitContextCallback
}

// ArgumentCallback
//
//	Callback function with parameters
type ArgumentCallback struct {
	Callback *reflect.Value
}

type argumentChannel struct {
	browserId int32
	channelId string
}

func (m *argumentChannel) BrowserId() int32 {
	return m.browserId
}

func (m *argumentChannel) ChannelId() string {
	return m.channelId
}

// ContextCallback
//
//	Return context parameter callback
func (m *Callback) ContextCallback() *ContextCallback {
	if m.Context == nil {
		return nil
	}
	return m.Context
}

// Invoke context function
func (m *ContextCallback) Invoke(context context.IContext) {
	// call
	m.Callback(context)
	resultValues := context.Replay().Result()
	if len(resultValues) > 0 {
		// call result
		resultArgument := make([]interface{}, len(resultValues), len(resultValues))
		for i, result := range resultValues {
			switch result.(type) {
			case error:
				resultArgument[i] = result.(error).Error()
			default:
				resultArgument[i] = result
			}
		}
		// result
		context.Result(resultArgument...)
	} else {
		context.Result(nil)
	}
}

// ArgumentCallback
//
//	return argument list callback function
func (m *Callback) ArgumentCallback() *ArgumentCallback {
	if m.Argument == nil {
		return nil
	}
	return m.Argument
}

// Invoke argument list function
func (m *ArgumentCallback) Invoke(context context.IContext) {
	var (
		argsList     json.JSONArray
		argsSize     int
		inArgsValues []reflect.Value
	)
	argsList = context.ArgumentList()
	if argsList != nil {
		argsSize = argsList.Size()
	}
	rt := m.Callback.Type()
	inArgsCount := rt.NumIn()
	inArgsValues = make([]reflect.Value, inArgsCount)
	var inIdx = 0
	for i := 0; i < inArgsCount; i++ {
		inType := rt.In(i)
		if inIdx < argsSize {
			argsValue := argsList.GetByIndex(inIdx)
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
							if err := goJSON.Unmarshal(jsonBytes, v.Interface()); err == nil {
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
								if err := goJSON.Unmarshal(jsonBytes, vv.Interface()); err == nil {
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
								if err := goJSON.Unmarshal(jsonBytes, vv.Interface()); err == nil {
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
			newValue := reflect.New(inType).Elem()
			if newValue.Type().Implements(argumentChannelType) {
				newValue.Set(reflect.ValueOf(&argumentChannel{browserId: context.BrowserId(), channelId: context.FrameId()}))
			}
			inArgsValues[i] = newValue
		} else {
			inIdx++
		}
	}
	// call
	resultValues := m.Callback.Call(inArgsValues)
	if len(resultValues) > 0 {
		// call result
		resultArgument := make([]interface{}, len(resultValues), len(resultValues))
		for i, result := range resultValues {
			res := result.Interface()
			switch res.(type) {
			case error:
				resultArgument[i] = res.(error).Error()
			default:
				resultArgument[i] = res
			}
		}
		// result
		context.Result(resultArgument...)
	} else {
		// result nil
		context.Result(nil)
	}
}
