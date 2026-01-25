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
	"errors"
	"fmt"
	"github.com/energye/energy/v3/ipc/callback"
	"reflect"
	"strconv"
)

// BindEvent 将指定对象绑定到应用程序的事件系统中
// 该方法会将对象注册到事件处理器中，使其能够接收和处理事件
//
//	obj - 需要绑定事件的对象
func BindEvent(obj any) {
	BindEventPrefix("", obj)
}

// BindEventPrefix 将对象的方法绑定到事件系统中，使用指定前缀作为事件名称的一部分
//   - prefix: 事件名称前缀，如果为空则使用对象类型的名称
//   - obj: 需要绑定的对象指针，该对象的方法将被注册为事件处理器
func BindEventPrefix(prefix string, obj any) {
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() != reflect.Ptr {
		return
	}
	objType := objVal.Type().Elem()
	if prefix == "" {
		prefix = objType.Name()
	}
	if prefix == "" {
		return
	}
	// 遍历对象的所有方法并将其注册为事件处理器
	for i := 0; i < objVal.NumMethod(); i++ {
		method := objVal.Type().Method(i)
		methodName := method.Name
		eventName := fmt.Sprintf("%s.%s", prefix, methodName)
		methodVal := objVal.MethodByName(methodName)
		// 创建事件处理函数，负责参数验证、类型转换和方法调用
		handler := func(args []any, err *error) []any {
			if !methodVal.IsValid() {
				*err = errors.New(fmt.Sprintf("方法%s不存在", methodName))
				return nil
			}
			methodType := methodVal.Type()
			paramCount := methodType.NumIn()
			if len(args) != paramCount {
				*err = errors.New(fmt.Sprintf("参数数量不匹配：需要%d个，实际传入%d个", paramCount, len(args)))
				return nil
			}
			// 处理输入参数的类型转换
			in := make([]reflect.Value, paramCount)
			for j := 0; j < paramCount; j++ {
				argVal := reflect.ValueOf(args[j])
				targetType := methodType.In(j)

				if convertValue, ok := srcConvertTargetType(argVal, targetType); ok {
					in[j] = convertValue
				} else {
					*err = errors.New(fmt.Sprintf("第%d个参数类型不匹配：需要%s，实际传入%s", j+1, methodType.In(j), argVal.Type()))
					return nil
				}
			}
			// 调用目标方法并处理返回值
			out := methodVal.Call(in)
			result := make([]any, len(out))
			for k := 0; k < len(out); k++ {
				result[k] = out[k].Interface()
			}
			return result
		}
		On(eventName, func(context callback.IContext) {
			var err error
			result := handler(context.Data().([]any), &err)
			if err != nil {
				context.Result(err.Error())
			} else {
				context.Result(result...)
			}
		})
	}
}

func srcConvertTargetType(srcVal reflect.Value, targetType reflect.Type) (reflect.Value, bool) {
	if !srcVal.IsValid() || targetType == nil {
		return reflect.Value{}, false
	}
	if srcVal.Type() == targetType {
		return srcVal, true
	}
	srcKind := srcVal.Kind()
	targetKind := targetType.Kind()
	switch targetKind {
	case reflect.Int:
		return convertToInt(srcVal, srcKind)
	case reflect.Float64:
		return convertToFloat64(srcVal, srcKind)
	case reflect.Float32:
		return convertToFloat32(srcVal, srcKind)
	case reflect.Bool:
		return convertToBool(srcVal, srcKind)
	case reflect.String:
		return convertToString(srcVal, srcKind)
	default:
		// 不支持的目标类型（如complex/array/struct等）
		return reflect.Value{}, false
	}
}

const (
	minInt = -(1 << (strconv.IntSize - 1))
	maxInt = (1 << (strconv.IntSize - 1)) - 1
)

// convertToInt 转换为int类型
func convertToInt(srcVal reflect.Value, srcKind reflect.Kind) (reflect.Value, bool) {
	switch srcKind {
	case reflect.Float64, reflect.Float32:
		val := srcVal.Float()
		if val < float64(minInt) || val > float64(maxInt) {
			return reflect.Value{}, false
		}
		return reflect.ValueOf(int(val)), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := srcVal.Int()
		if val < int64(minInt) || val > int64(maxInt) {
			return reflect.Value{}, false
		}
		return reflect.ValueOf(int(val)), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := srcVal.Uint()
		if val > uint64(maxInt) {
			return reflect.Value{}, false
		}
		return reflect.ValueOf(int(val)), true
	case reflect.Bool:
		if srcVal.Bool() {
			return reflect.ValueOf(1), true
		}
		return reflect.ValueOf(0), true
	case reflect.String:
		valStr := srcVal.String()
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return reflect.Value{}, false
		}
		return reflect.ValueOf(val), true
	default:
		return reflect.Value{}, false
	}
}

// convertToFloat64 转换为float64类型
func convertToFloat64(srcVal reflect.Value, srcKind reflect.Kind) (reflect.Value, bool) {
	switch srcKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(float64(srcVal.Int())), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(float64(srcVal.Uint())), true
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(srcVal.Float()), true
	case reflect.Bool:
		if srcVal.Bool() {
			return reflect.ValueOf(1.0), true
		}
		return reflect.ValueOf(0.0), true
	case reflect.String:
		valStr := srcVal.String()
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return reflect.Value{}, false
		}
		return reflect.ValueOf(val), true
	default:
		return reflect.Value{}, false
	}
}

// convertToFloat32 转换为float32类型
func convertToFloat32(srcVal reflect.Value, srcKind reflect.Kind) (reflect.Value, bool) {
	switch srcKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(float32(srcVal.Int())), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(float32(srcVal.Uint())), true
	case reflect.Float32:
		return srcVal, true
	case reflect.Float64:
		return reflect.ValueOf(float32(srcVal.Float())), true
	case reflect.Bool:
		if srcVal.Bool() {
			return reflect.ValueOf(float32(1.0)), true
		}
		return reflect.ValueOf(float32(0.0)), true
	case reflect.String:
		valStr := srcVal.String()
		val, err := strconv.ParseFloat(valStr, 32)
		if err != nil {
			return reflect.Value{}, false
		}
		return reflect.ValueOf(float32(val)), true
	default:
		return reflect.Value{}, false
	}
}

// convertToBool 转换为bool类型
func convertToBool(srcVal reflect.Value, srcKind reflect.Kind) (reflect.Value, bool) {
	switch srcKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(srcVal.Int() != 0), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(srcVal.Uint() != 0), true
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(srcVal.Float() != 0), true
	case reflect.Bool:
		return srcVal, true
	case reflect.String:
		valStr := srcVal.String()
		val, err := strconv.ParseBool(valStr)
		if err != nil {
			switch valStr {
			case "1":
				return reflect.ValueOf(true), true
			case "0":
				return reflect.ValueOf(false), true
			default:
				return reflect.Value{}, false
			}
		}
		return reflect.ValueOf(val), true
	default:
		return reflect.Value{}, false
	}
}

// convertToString 转换为string类型
func convertToString(srcVal reflect.Value, srcKind reflect.Kind) (reflect.Value, bool) {
	switch srcKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(strconv.FormatInt(srcVal.Int(), 10)), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(strconv.FormatUint(srcVal.Uint(), 10)), true
	case reflect.Float32:
		return reflect.ValueOf(strconv.FormatFloat(srcVal.Float(), 'f', -1, 32)), true
	case reflect.Float64:
		return reflect.ValueOf(strconv.FormatFloat(srcVal.Float(), 'f', -1, 64)), true
	case reflect.Bool:
		return reflect.ValueOf(strconv.FormatBool(srcVal.Bool())), true
	case reflect.String:
		return srcVal, true
	default:
		return reflect.Value{}, false
	}
}
