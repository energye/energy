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
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
	"reflect"
	"unsafe"
)

var (
	setCefWindowBindCallbackFunc dllimports.ProcAddr
)

func cefV8WindowBindFuncEventsInit() {
	// 初始设置回调函数指针
	setCefWindowBindCallbackFunc = imports.Proc(internale_SetCEFWindowBindCallbackFunc)
	setCefWindowBindCallbackFunc.Call(cefWindowBindEvent)
}

func cefWindowBindCallbackEventProc(f uintptr, args uintptr, argcout int) uintptr {
	getVal := func(i int) uintptr {
		return common.GetParamOf(i, args)
	}
	eventType := BIND_EVENT(getVal(0))
	if BE_FUNC == eventType {
		_cefV8BindFuncCallbackHandler(eventType, f, args, argcout)
	} else {
		_cefV8BindFieldCallbackHandler(eventType, f, args, argcout)
	}
	return 0
}

// 字段处理
func _cefV8BindFieldCallbackHandler(eventType BIND_EVENT, fullNamePtr uintptr, args uintptr, argsLen int) {
	var (
		exceptionPrt *uintptr
		errorMessage = Empty
	)
	defer func() {
		if err := recover(); err != nil {
			if exceptionPrt != nil {
				*exceptionPrt = api.PascalStr((err.(error)).Error())
			}
			logger.Error("V8BindFieldCallbackHandler Error", err)
		}
	}()
	getVal := func(i int) uintptr {
		return common.GetParamOf(i, args)
	}
	getPtr := func(i int) unsafe.Pointer {
		return unsafe.Pointer(getVal(i))
	}
	fullName := api.GoStr(fullNamePtr)
	if eventType == BE_SET { //设置值
		var (
			newValueType   = (V8_JS_VALUE_TYPE)(getVal(1))
			newStringValue string
			newIntValue    int32
			newDoubleValue float64
			newBoolValue   bool
		)
		exceptionPrt = (*uintptr)(getPtr(3))
		args := ipc.NewArgumentList()
		defer args.Clear()
		args.SetString(0, fullName)
		switch newValueType {
		case V8_VALUE_STRING:
			newStringValue = api.GoStr(getVal(2))
			args.SetString(1, newStringValue)
		case V8_VALUE_INT:
			newIntValue = int32(getVal(2))
			args.SetInt32(1, newIntValue)
		case V8_VALUE_DOUBLE:
			newDoubleValue = *(*float64)(getPtr(2))
			args.SetFloat64(1, newDoubleValue)
		case V8_VALUE_BOOLEAN:
			newBoolValue = api.GoBool(getVal(2))
			args.SetBool(1, newBoolValue)
		default:
			errorMessage = cefErrorMessage(CVE_ERROR_TYPE_NOT_SUPPORTED)
		}
		if errorMessage == Empty {
			iCtx := ipc.IPC.Render().EmitAndReturn(ipc.Ln_SET_BIND_FIELD_VALUE, args)
			if iCtx != nil {
				args := iCtx.Arguments()
				isSuccess := args.GetBool(0)
				if !isSuccess {
					exception := CEF_V8_EXCEPTION(args.GetInt8(1))
					errorMessage = cefErrorMessage(exception)
				}
				iCtx.Free()
			} else {
				errorMessage = cefErrorMessage(CVE_ERROR_IPC_GET_BIND_FIELD_VALUE_FAIL)
			}
		}
		*exceptionPrt = api.PascalStr(errorMessage)
	} else if eventType == BE_GET { //获取值和set处理不一样，需要在go里知道实际类型
		var (
			newValueType      = (*V8_JS_VALUE_TYPE)(getPtr(1))
			retStringValuePrt = (*uintptr)(getPtr(2))
			retIntValuePrt    = (*int32)(getPtr(3))
			retDoubleValuePrt = (*float64)(getPtr(4))
			retBoolValuePrt   = (*bool)(getPtr(5))
		)
		exceptionPrt = (*uintptr)(getPtr(6))
		args := ipc.NewArgumentList()
		defer args.Clear()
		args.SetString(0, fullName)
		iCtx := ipc.IPC.Render().EmitAndReturn(ipc.Ln_GET_BIND_FIELD_VALUE, args)
		if iCtx != nil {
			data := iCtx.Message().Data()
			isSuccess := common.ByteToBool(data[0])
			valueType := V8_JS_VALUE_TYPE(common.ByteToInt8(data[1]))
			if isSuccess {
				errorMessage = getPtrValue(valueType, data[2:], retStringValuePrt, retIntValuePrt, retDoubleValuePrt, retBoolValuePrt)
			} else {
				exception := CEF_V8_EXCEPTION(int8(data[1]))
				errorMessage = cefErrorMessage(exception)
			}
			if errorMessage == Empty {
				*newValueType = valueType
			}
			iCtx.Free()
		} else {
			errorMessage = cefErrorMessage(CVE_ERROR_IPC_GET_BIND_FIELD_VALUE_FAIL)
		}
		*exceptionPrt = api.PascalStr(errorMessage)
	}
}

func setPtrValue(jsValue JSValue, newValueType V8_JS_VALUE_TYPE, stringValue string, intValue int32, doubleValue float64, boolValue bool) CEF_V8_EXCEPTION {
	if jsValue.isCommon() {
		//valueType说明给的值类型不相同，需要将 jsValue 转换一下
		switch newValueType { //设置值，这里是通用类型只需要知道js里设置的什么类型即可
		case V8_VALUE_STRING:
			jsValue.setValue(stringValue)
		case V8_VALUE_INT:
			jsValue.setValue(intValue)
		case V8_VALUE_DOUBLE:
			jsValue.setValue(doubleValue)
		case V8_VALUE_BOOLEAN:
			jsValue.setValue(boolValue)
		default:
			return CVE_ERROR_TYPE_NOT_SUPPORTED
		}
		jsValue.setValueType(newValueType)
	} else { //object
		if refValue, ok := jsValue.getValue().(*reflect.Value); !ok {
			return CVE_ERROR_TYPE_INVALID
		} else {
			if newValueType == V8_VALUE_STRING && jsValue.ValueType() == V8_VALUE_STRING {
				refValue.Set(reflect.ValueOf(stringValue))
			} else if newValueType == V8_VALUE_INT && jsValue.ValueType() == V8_VALUE_INT {
				refValue.SetInt(int64(intValue))
			} else if newValueType == V8_VALUE_DOUBLE && jsValue.ValueType() == V8_VALUE_DOUBLE {
				refValue.SetFloat(float64(doubleValue))
			} else if newValueType == V8_VALUE_BOOLEAN && jsValue.ValueType() == V8_VALUE_BOOLEAN {
				refValue.SetBool(boolValue)
			} else {
				return CVE_ERROR_TYPE_NOT_SUPPORTED
			}
		}
	}
	return CVE_ERROR_OK
}

func getPtrValue(valueType V8_JS_VALUE_TYPE, newValue interface{}, stringValuePrt *uintptr, intValuePrt *int32, doubleValuePrt *float64, boolValuePrt *bool) string {
	switch valueType {
	case V8_VALUE_STRING:
		if value, err := common.ValueToString(newValue); err == nil {
			*stringValuePrt = api.PascalStr(value)
		} else {
			return cefErrorMessage(CVE_ERROR_GET_STRING_FAIL)
		}
	case V8_VALUE_INT:
		if value, err := common.ValueToInt32(newValue); err == nil {
			*intValuePrt = value
		} else {
			return cefErrorMessage(CVE_ERROR_GET_INT_FAIL)
		}
	case V8_VALUE_DOUBLE:
		if value, err := common.ValueToFloat64(newValue); err == nil {
			*doubleValuePrt = value
		} else {
			return cefErrorMessage(CVE_ERROR_GET_DOUBLE_FAIL)
		}
	case V8_VALUE_BOOLEAN:
		if value, err := common.ValueToBool(newValue); err == nil {
			*boolValuePrt = value
		} else {
			return cefErrorMessage(CVE_ERROR_GET_BOOL_FAIL)
		}
	case V8_VALUE_NULL:
	case V8_VALUE_UNDEFINED:
	default:
		return cefErrorMessage(CVE_ERROR_TYPE_NOT_SUPPORTED)
	}
	return ""
}

// 函数处理
func _cefV8BindFuncCallbackHandler(eventType BIND_EVENT, fullNamePtr uintptr, args uintptr, argsLen int) {
	var (
		exceptionPrt *uintptr
		errorMessage = Empty
	)
	defer func() {
		if err := recover(); err != nil {
			logger.Error("V8BindFuncCallbackHandler recover:", err)
			if exceptionPrt != nil {
				*exceptionPrt = api.PascalStr(" " + (err.(error)).Error())
			}
		}
	}()
	getVal := func(i int) uintptr {
		return common.GetParamOf(i, args)
	}
	getPtr := func(i int) unsafe.Pointer {
		return unsafe.Pointer(getVal(i))
	}
	fullName := api.GoStr(fullNamePtr)
	exceptionPrt = (*uintptr)(getPtr(1))
	var jsValue, ok = VariableBind.GetValueBind(fullName)
	var fnInfo *funcInfo
	if !ok {
		errorMessage = cefErrorMessage(CVE_ERROR_NOT_FOUND_FUNC)
		return
	}
	fnInfo = jsValue.getFuncInfo()
	if fnInfo != nil {
		var (
			inIdx      = 0
			i          = 0
			argsDefLen = int(2 + fnInfo.OutNum)
		)
		var (
			outParams         []reflect.Value
			ok                bool
			inArgumentList    ipc.IArgumentList
			retStringValuePrt *uintptr
			retIntValuePrt    *int32
			retDoubleValuePrt *float64
			retBoolValuePrt   *bool
			outVType          *vt
		)
		if fnInfo.OutNum == 2 {
			argsDefLen++
			if fnInfo.OutParam[0].Jsv == V8_VALUE_EXCEPTION {
				outVType = fnInfo.OutParam[1]
			} else {
				outVType = fnInfo.OutParam[0]
			}
		} else if fnInfo.OutNum == 1 {
			if fnInfo.OutParam[0].Jsv != V8_VALUE_EXCEPTION {
				argsDefLen++
				outVType = fnInfo.OutParam[0]
			}
		}
		inArgumentList = ipc.NewArgumentList()
		defer inArgumentList.Clear()
		//遍历入参，并校验，默认参数结束下标位置开始
		for inIdx = argsDefLen; inIdx < argsLen && i < len(fnInfo.InParam); inIdx++ {
			i = inIdx - argsDefLen
			inVType := fnInfo.InParam[i]
			switch inVType.Jsv {
			case V8_VALUE_STRING:
				str := api.GoStr(getVal(inIdx))
				inArgumentList.SetString(i, str)
			case V8_VALUE_INT:
				if value := common.NumberUintPtrToInt(getVal(inIdx), inVType.Gov); value != nil {
					inArgumentList.SetIntAuto(i, value, inVType.Gov)
				} else {
					errorMessage = cefErrorMessage(CVE_ERROR_FUNC_GET_IN_PAM_INT_FAIL)
					break
				}
			case V8_VALUE_DOUBLE:
				if value := common.NumberPtrToFloat(getPtr(inIdx), inVType.Gov); value != nil {
					inArgumentList.SetFloatAuto(i, value, inVType.Gov)
				} else {
					errorMessage = cefErrorMessage(CVE_ERROR_FUNC_GET_IN_PAM_DOUBLE_FAIL)
					break
				}
			case V8_VALUE_BOOLEAN:
				inArgumentList.SetBool(i, api.GoBool(getVal(inIdx)))
			default:
				//参数错误停止调用，返回
				errorMessage = cefErrorMessage(CVE_ERROR_FUNC_IN_PAM)
				break
			}
		}
		//有出参时，出参长度是2，并校验
		//出参的固定下标为2
		if errorMessage == Empty && outVType != nil {
			//得到出参的类型指针 出参的固定下标为2
			switch outVType.Jsv {
			case V8_VALUE_STRING:
				retStringValuePrt = (*uintptr)(getPtr(2))
			case V8_VALUE_INT:
				retIntValuePrt = (*int32)(getPtr(2))
			case V8_VALUE_DOUBLE:
				retDoubleValuePrt = (*float64)(getPtr(2))
			case V8_VALUE_BOOLEAN:
				retBoolValuePrt = (*bool)(getPtr(2))
			default:
				errorMessage = cefErrorMessage(CVE_ERROR_FUNC_OUT_PAM)
				break
			}
		}

		if errorMessage == Empty {
			if SingleProcess || common.Args.IsMain() {
				outParams, ok = jsValue.invoke(inArgumentList.ToReflectValue())
			} else {
				ok = true
				inArgumentList.SetString(inArgumentList.Size(), fullName)
				var iCtx = ipc.IPC.Render().EmitAndReturn(ipc.Ln_EXECUTE_BIND_FUNC, inArgumentList)
				if iCtx != nil {
					var outArgument = ipc.NewArgumentList()
					defer outArgument.Clear()
					data := iCtx.Message().Data()
					isSuccess := common.ByteToBool(data[0])
					if isSuccess {
						outArgument.UnPackage(data[1:])
						outParams = outArgument.ToReflectValue()
					} else {
						exception := CEF_V8_EXCEPTION(int8(data[1]))
						errorMessage = cefErrorMessage(exception)
					}
					iCtx.Free()
				}
			}
			if ok {
				if outVType != nil && len(outParams) > 0 &&
					(retStringValuePrt != nil || retIntValuePrt != nil || retDoubleValuePrt != nil || retBoolValuePrt != nil) {
					//出参处理
					outParamValue := outParams[fnInfo.OutParamIdx]
					switch outVType.Jsv {
					case V8_VALUE_STRING:
						if value, err := common.ValueToString(outParamValue.Interface()); err == nil {
							*retStringValuePrt = api.PascalStr(value)
						} else {
							errorMessage = cefErrorMessage(CVE_ERROR_FUNC_GET_OUT_PAM_STRING_FAIL)
						}
					case V8_VALUE_INT:
						if value, err := common.ValueToInt32(outParamValue.Interface()); err == nil {
							*retIntValuePrt = value
						} else {
							errorMessage = cefErrorMessage(CVE_ERROR_FUNC_GET_OUT_PAM_INT_FAIL)
						}
					case V8_VALUE_DOUBLE:
						if value, err := common.ValueToFloat64(outParamValue.Interface()); err == nil {
							*retDoubleValuePrt = value
						} else {
							errorMessage = cefErrorMessage(CVE_ERROR_FUNC_GET_OUT_PAM_DOUBLE_FAIL)
						}
					case V8_VALUE_BOOLEAN:
						if value, err := common.ValueToBool(outParamValue.Interface()); err == nil {
							*retBoolValuePrt = value
						} else {
							errorMessage = cefErrorMessage(CVE_ERROR_FUNC_GET_OUT_PAM_BOOLEAN_FAIL)
						}
					default:
						errorMessage = cefErrorMessage(CVE_ERROR_FUNC_OUT_PAM)
						break
					}
				}
			} else {
				errorMessage = cefErrorMessage(CVE_ERROR_UNKNOWN_ERROR)
			}
		}
	} else {
		errorMessage = cefErrorMessage(CVE_ERROR_NOT_FOUND_FUNC)
	}
	*exceptionPrt = api.PascalStr(errorMessage)
}
