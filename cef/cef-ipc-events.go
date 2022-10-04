//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"bytes"
	"encoding/binary"
	. "github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/dylib"
	"github.com/energye/golcl/lcl/api"
	"net"
	"reflect"
	"strings"
	"unsafe"
)

const (
	IPC_FN_TYPE_IPCJSEmitGo    = 1 // JS fires the Go registerEvent
	IPC_FN_TYPE_IPCGoEmitJSRet = 2 // Go fires the JS registerEvent
)

var (
	setCefIPCCallbackFunc *dylib.LazyProc
)

type rGoResult struct {
	value       uintptr //Pointer
	valueType   uintptr //PInteger
	valueLength uintptr //PInteger
	bindType    uintptr //PInteger
	exception   uintptr //PInteger
}

func (m *rGoResult) set(value, valueType, valueLength, bindType, exception uintptr) {
	m.value = value
	m.valueType = valueType
	m.valueLength = valueLength
	m.bindType = bindType
	m.exception = exception
}

type rIPCEventParam struct {
	BrowserId       int32
	IPCId           int32
	FullName        uintptr //string
	FrameId         uintptr //string
	ParamType       int32   //js中ipc.emit参数类型: 1:没有入参和回调函数 2:有入参 没有回调函数 3:没有入参 有回调函数 4:有入参 有回调函数
	ValueTypeArrLen int32   //参数个数
	ValueTypeArr    uintptr //ValueTypeArr: array of byte;
}

func cefIPCInit() {
	setCefIPCCallbackFunc = api.GetLibLCL().NewProc("SetCEFIPCCallbackFunc")
	setCefIPCCallbackFunc.Call(cefIPCCallbackFuncEvent)
}

func cefIPCEventProc(fnType uintptr, args uintptr, argsLen int) uintptr {
	//defer func() {
	//	if err := recover(); err != nil {
	//		Logger.Error("IPCEventProc Error:", err)
	//	}
	//}()
	getVal := func(i int) uintptr {
		return GetParamOf(i, args)
	}
	getPtr := func(i int) unsafe.Pointer {
		return unsafe.Pointer(getVal(i))
	}
	if fnType == IPC_FN_TYPE_IPCJSEmitGo {
		eventParam := (*rIPCEventParam)(getPtr(0))
		result := (*rGoResult)(getPtr(1))
		if eventParam == nil || result == nil {
			return 0
		}
		ipcJSEmitGo(eventParam, result, args)
	} else if fnType == IPC_FN_TYPE_IPCGoEmitJSRet {
		ipcId := *(*int32)(getPtr(0))
		tmType := *(*int32)(getPtr(1))
		result := (*rGoResult)(getPtr(2))
		ipcGoEmitJS(ipcId, TriggerMode(tmType), result, args)
		result.set(0, 0, 0, 0, 0)
		result = nil
	}
	return 0
}

//ipc - go emit js on event
func ipcGoEmitJS(ipcId int32, triggerMode TriggerMode, result *rGoResult, args uintptr) {
	inArgument := ipc.NewArgumentList()
	inArgument.SetBool(1, true)
	if CEF_V8_EXCEPTION(result.exception) == CVE_ERROR_OK {
		switch V8_JS_VALUE_TYPE(result.valueType) {
		case V8_VALUE_STRING:
			inArgument.SetString(0, api.DStrToGoStr(result.value))
		case V8_VALUE_INT:
			inArgument.SetInt32(0, int32(result.value))
		case V8_VALUE_DOUBLE:
			inArgument.SetFloat64(0, *(*float64)(GetParamPtr(result.value, 0)))
		case V8_VALUE_BOOLEAN:
			inArgument.SetBool(0, *(*bool)(GetParamPtr(result.value, 0)))
		default:
			inArgument.SetBool(1, false)
			inArgument.SetString(0, "不支持的数据类型")
		}
	} else {
		inArgument.SetBool(1, false)
		inArgument.SetString(0, api.DStrToGoStr(result.value))
	}
	if triggerMode == Tm_Callback { //回调函数
		if callback, ok := executeJS.emitCallback.EmitCollection.Load(ipcId); ok {
			executeJS.emitCallback.EmitCollection.Delete(ipcId)
			ctx := ipc.NewIPCContext("", 0, 0, 0, nil, nil, nil, &ipc.IPCContextResult{}, inArgument)
			callback.(ipc.IPCCallback)(ctx)
			ctx.Free()
		}
	} else if triggerMode == Tm_Sync { //同步调用
		if chn, ok := executeJS.emitSync.EmitCollection.Load(ipcId); ok {
			ctx := ipc.NewIPCContext("", 0, 0, 0, nil, nil, nil, &ipc.IPCContextResult{}, inArgument)
			var c = chn.(chan ipc.IIPCContext)
			c <- ctx
			close(c)
			executeJS.emitSync.EmitCollection.Delete(ctx.EventId())
		}
	}
}

//ipc - js emit go on event
func ipcJSEmitGo(eventParam *rIPCEventParam, result *rGoResult, args uintptr) {
	getVal := func(i int) uintptr {
		return GetParamOf(i, args)
	}
	getPtr := func(i int) unsafe.Pointer {
		return unsafe.Pointer(getVal(i))
	}
	var (
		inArgument   = ipc.NewArgumentList()
		fullName     = api.DStrToGoStr(eventParam.FullName)
		valueTypeLen = eventParam.ValueTypeArrLen //入参类型数组
		accIdx       = 2                          //占用的下标
	)
	if valueTypeLen > 0 {
		//取入参
		for i := 0; i < int(valueTypeLen); i++ {
			valueType := V8_JS_VALUE_TYPE(*(*byte)(GetParamPtr(eventParam.ValueTypeArr, i)))
			switch valueType {
			case V8_VALUE_STRING:
				inArgument.SetString(i, api.DStrToGoStr(getVal(i+accIdx)))
			case V8_VALUE_INT:
				inArgument.SetInt32(i, int32(getVal(i+accIdx)))
			case V8_VALUE_DOUBLE:
				inArgument.SetFloat64(i, *(*float64)(getPtr(i + accIdx)))
			case V8_VALUE_BOOLEAN:
				inArgument.SetBool(i, api.DBoolToGoBool(getVal(i+accIdx)))
			}
		}
	}

	//检索v8绑定
	bindType, name, vType, _, exception := searchBindV8Value(fullName)
	if exception == CVE_ERROR_OK {
		if vType == V8_VALUE_OBJECT || vType == V8_VALUE_ROOT_OBJECT {
			result.set(api.GoStrToDStr(name), uintptr(vType), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
		} else {
			if bindType == IS_OBJECT {
				name = name[len(objectRootName)+1:]
			}
			if jsValue, ok := VariableBind.GetValueBind(name); ok {
				jsValue.Lock()
				defer jsValue.UnLock()
				if jsValue.ValueType() == V8_VALUE_FUNCTION { //func
					var fnInfo = jsValue.getFuncInfo()
					if fnInfo != nil {
						inArgumentAdapter(0, inArgument, fnInfo)
						var outParams, ok = jsValue.invoke(inArgument.ToReflectValue())
						if ok {
							if eventParam.ParamType > 2 {
								var tmpOutParams = make([]reflect.Value, 1)
								if fnInfo.OutNum > 0 {
									tmpOutParams[0] = outParams[fnInfo.OutParamIdx]
									inArgument.ReflectValueConvert(tmpOutParams) //出参转换
									data := inArgument.GetData(0)
									switch GO_VALUE_TYPE(data.VType()) {
									case GO_VALUE_STRING:
										result.set(api.GoStrToDStr(data.GetString()), uintptr(V8_VALUE_STRING), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
									case GO_VALUE_INT, GO_VALUE_INT8, GO_VALUE_INT16, GO_VALUE_INT32, GO_VALUE_INT64, GO_VALUE_UINT, GO_VALUE_UINT8, GO_VALUE_UINT16, GO_VALUE_UINT32, GO_VALUE_UINT64, GO_VALUE_UINTPTR:
										result.set(uintptr(data.GetInt32()), uintptr(V8_VALUE_INT), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
									case GO_VALUE_FLOAT32, GO_VALUE_FLOAT64:
										var ret = data.GetDouble()
										result.set(uintptr(unsafe.Pointer(&ret)), uintptr(V8_VALUE_DOUBLE), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
									case GO_VALUE_BOOL:
										result.set(api.GoBoolToDBool(data.GetBool()), uintptr(V8_VALUE_BOOLEAN), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
									}
								} else {
									//没有出参
									result.set(0, uintptr(V8_NO_OUT_VALUE), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
								}
							}
						} else {
							result.set(api.GoStrToDStr(outParams[0].Interface().(string)), uintptr(V8_VALUE_STRING), 0, uintptr(bindType), uintptr(CVE_ERROR_UNKNOWN_ERROR))
						}
					}
				} else { //field
					inArgs := inArgument.Items()
					if inArgs != nil && len(inArgs) > 0 {
						var (
							errorMsg    string
							oldValuePtr unsafe.Pointer
							err         error
						)
						//字段赋值
						fieldNewValue := inArgs[0]
						oldValuePtr, err = jsValue.ValueToPtr()
						if err != nil {
							errorMsg = cefErrorMessage(CVE_ERROR_TYPE_NOT_SUPPORTED)
						} else {
							if fieldNewValue.IsString() {
								errorMsg = cefErrorMessage(SetPtrValue(jsValue, V8_VALUE_STRING, fieldNewValue.GetString(), 0, 0, false))
							} else if fieldNewValue.IsIntAuto() {
								errorMsg = cefErrorMessage(SetPtrValue(jsValue, V8_VALUE_INT, Empty, fieldNewValue.GetInt32(), 0, false))
							} else if fieldNewValue.IsFloatAuto() {
								errorMsg = cefErrorMessage(SetPtrValue(jsValue, V8_VALUE_DOUBLE, Empty, 0, fieldNewValue.GetFloat64(), false))
							} else if fieldNewValue.IsBoolean() {
								errorMsg = cefErrorMessage(SetPtrValue(jsValue, V8_VALUE_BOOLEAN, Empty, 0, 0, fieldNewValue.GetBool()))
							} else {
								errorMsg = cefErrorMessage(CVE_ERROR_TYPE_NOT_SUPPORTED)
							}
						}
						if Empty == errorMsg {
							result.set(uintptr(oldValuePtr), uintptr(jsValue.ValueType()), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
						} else {
							result.set(api.GoStrToDStr(errorMsg), uintptr(V8_VALUE_STRING), 0, uintptr(bindType), uintptr(CVE_ERROR_TYPE_NOT_SUPPORTED))
						}
					} else {
						//字段取值
						switch jsValue.ValueType() {
						case V8_VALUE_STRING, V8_VALUE_INT, V8_VALUE_DOUBLE, V8_VALUE_BOOLEAN:
							var ret, err = jsValue.ValueToPtr()
							if err != nil {
								result.set(0, 0, 0, uintptr(bindType), uintptr(CVE_ERROR_TYPE_NOT_SUPPORTED))
							} else {
								result.set(uintptr(ret), uintptr(jsValue.ValueType()), 0, uintptr(bindType), uintptr(CVE_ERROR_OK))
							}
						default:
							result.set(0, 0, 0, uintptr(bindType), uintptr(CVE_ERROR_TYPE_NOT_SUPPORTED))
						}
					}
				}
			} else {
				result.set(0, 0, 0, 0, uintptr(CVE_ERROR_NOT_FOUND_FIELD))
			}
		}
	} else {
		//尝试触发用户自定义 ipc fullName emit
		if callback := ipc.IPC.Browser().Events().Get(fullName); callback != nil {
			var (
				channelId = StrToInt64(api.DStrToGoStr(eventParam.FrameId))
				ipcType   IPC_TYPE
				unixConn  *net.UnixConn
				netConn   net.Conn
			)
			if chn := ipc.IPC.Browser().Channel(channelId); chn != nil {
				ipcType = chn.IPCType
				unixConn = chn.UnixConn
				netConn = chn.NetConn
			}
			ctx := ipc.NewIPCContext(fullName, eventParam.BrowserId, channelId, ipcType, unixConn, netConn, &ipc.IPCEventMessage{}, &ipc.IPCContextResult{}, inArgument)
			callback(ctx)
			result.set(uintptr(ctx.Result().Result()), uintptr(ctx.Result().VType()), 0, uintptr(IS_COMMON), uintptr(CVE_ERROR_OK))
		} else {
			result.set(0, 0, 0, 0, uintptr(CVE_ERROR_NOT_FOUND_FIELD))
		}
	}
}

//检索绑定的JsValue
//返回值
//	bindType 绑定类型 0通用类型，1对象类型
//  fullName 绑定全名称
//	vType	 js绑定值类型
//	eventID  事件ID
//  exception 错误码
func searchBindV8Value(fullName string) (IS_CO, string, V8_JS_VALUE_TYPE, int32, CEF_V8_EXCEPTION) {
	if fullName == "" {
		return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
	}
	var fnArr = strings.Split(fullName, ".")
	var fnArrLen = len(fnArr)
	if len(fnArr) > 0 && fnArr[0] == "window" {
		fnArr = fnArr[1:]
		fnArrLen = len(fnArr)
	}
	if fnArrLen == 0 {
		return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
	}
	//对象类型
	if fnArr[0] == commonRootName {
		if fnArrLen == 1 { //commonRootObject 对象
			fullName = strings.Join(fnArr, ".")
			return IS_COMMON, fullName, V8_VALUE_ROOT_OBJECT, 0, CVE_ERROR_OK
		} else if fnArrLen == 2 {
			if jsValue, ok := VariableBind.GetValueBind(fnArr[1]); ok { //通用类型 fnArr[1] 1是因为 obj.field ,这个field 就是1
				return IS_COMMON, jsValue.Name(), jsValue.ValueType(), int32(jsValue.getEventId()), CVE_ERROR_OK
			} else {
				//不存在
				return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
			}
		} else {
			//不正确的变量名或不存在
			return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
		}
	} else if fnArr[0] == objectRootName {
		if fnArrLen == 1 {
			//objectRootObject对象
			fullName = strings.Join(fnArr, ".")
			return IS_OBJECT, fullName, V8_VALUE_ROOT_OBJECT, 0, CVE_ERROR_OK
		}
		//对象类型查找关系
		var preObj *structObjectInfo
		var subObj = objectSti.StructsObject
		var faLen = len(fnArr)
		for i := 0; i < faLen; i++ {
			if obj, ok := subObj[fnArr[i]]; ok {
				subObj = obj.SubStructObjectInfo
				preObj = obj
			} else if i > 0 && i+1 < faLen {
				return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
			}
		}
		if preObj != nil {
			var fieldName = fnArr[fnArrLen-1]
			fullName = strings.Join(fnArr, ".")
			if preObj.ObjName == fieldName {
				//返回object Obj对象
				return IS_OBJECT, fullName, V8_VALUE_OBJECT, 0, CVE_ERROR_OK
			} else {
				//字段 或 函数
				if fieldInfo, ok := preObj.FieldsInfo[fieldName]; ok {
					return IS_OBJECT, fullName, fieldInfo.ValueType.Jsv, int32(fieldInfo.EventId), CVE_ERROR_OK
				} else if fnInfo, ok := preObj.FuncsInfo[fieldName]; ok {
					return IS_OBJECT, fullName, V8_VALUE_FUNCTION, int32(fnInfo.EventId), CVE_ERROR_OK
				}
				return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
			}
		} else {
			return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
		}
	} else {
		return 0, "", 0, 0, CVE_ERROR_NOT_FOUND_FIELD
	}

}

func inArgumentAdapter(argsIdxOffset int, inArgument ipc.IArgumentList, fnInfo *funcInfo) {
	items := inArgument.Items()
	inParamType := fnInfo.InParam
	for i := argsIdxOffset; i < inArgument.Size(); i++ {
		if items[i].IsIntAuto() && inParamType[i-argsIdxOffset].IsGoIntAuto() {
			inArgument.SetIntAuto(i, items[i].GetInt32(), fnInfo.InParam[i-argsIdxOffset].Gov)
		}
	}
}

func outArgumentAdapter(argsIdxOffset int, outArgument ipc.IArgumentList, fnInfo *funcInfo) {
	items := outArgument.Items()
	outParamType := fnInfo.OutParam
	for i := argsIdxOffset; i < outArgument.Size(); i++ {
		if items[i].IsIntAuto() && outParamType[i-argsIdxOffset].IsGoIntAuto() {
			outArgument.SetIntAuto(i, items[i].GetInt32(), fnInfo.InParam[i-argsIdxOffset].Gov)
		}
	}
}

func internalBrowserIPCOnEventInit() {
	var ipcJsBindGetValueSuccess = func(buf *bytes.Buffer, valueType int8, data []byte) {
		//rec
		//	eventID 	[]byte		4bit
		//ret
		//	isSuccess 	[]byte 		1bit
		//	valueType 	[]byte 		1bit
		//	value		[]byte 		MaxInt32bit
		binary.Write(buf, binary.BigEndian, true)      //成功
		binary.Write(buf, binary.BigEndian, valueType) //字段类型
		binary.Write(buf, binary.BigEndian, data)      //字段值
	}

	var ipcJsBindGetValueFail = func(buf *bytes.Buffer, err CEF_V8_EXCEPTION) {
		binary.Write(buf, binary.BigEndian, false) //失败
		binary.Write(buf, binary.BigEndian, int8(err))
	}
	//var occIdx = ipcAccIdx
	//主进程监听回调
	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		//var fieldReadWriteLock = new(sync.Mutex)
		//获取字段值-同步
		event.On(ipc.Ln_GET_BIND_FIELD_VALUE, func(context ipc.IIPCContext) {
			args := context.Arguments()
			defer args.Clear()
			fullName := args.GetString(0)
			var jsValue, ok = VariableBind.GetValueBind(fullName)
			buf := &bytes.Buffer{}
			defer buf.Reset()
			if ok {
				jsValue.Lock()
				defer jsValue.UnLock()
				var valueType = int8(jsValue.ValueType())
				ipcJsBindGetValueSuccess(buf, valueType, jsValue.Bytes())
			} else {
				ipcJsBindGetValueFail(buf, CVE_ERROR_NOT_FOUND_FIELD)
			}
			context.Response(buf.Bytes())
			buf = nil
			context.Free()
		})
		//设置字段值-同步
		event.On(ipc.Ln_SET_BIND_FIELD_VALUE, func(context ipc.IIPCContext) {
			args := context.Arguments()
			defer args.Clear()
			fullName := args.GetString(0)
			item := args.Item(1)
			newValueType := item.VTypeToJS()
			jsValue, ok := VariableBind.GetValueBind(fullName)
			isSuccess := false
			retArgs := ipc.NewArgumentList()
			defer retArgs.Clear()
			if ok {
				jsValue.Lock()
				defer jsValue.UnLock()
				switch newValueType { //设置值，这里是通用类型只需要知道js里设置的什么类型即可
				case V8_VALUE_STRING:
					isSuccess = Empty == cefErrorMessage(SetPtrValue(jsValue, newValueType, item.GetString(), 0, 0, false))
				case V8_VALUE_INT:
					isSuccess = Empty == cefErrorMessage(SetPtrValue(jsValue, newValueType, Empty, item.GetInt32(), 0, false))
				case V8_VALUE_DOUBLE:
					isSuccess = Empty == cefErrorMessage(SetPtrValue(jsValue, newValueType, Empty, 0, item.GetFloat64(), false))
				case V8_VALUE_BOOLEAN:
					isSuccess = Empty == cefErrorMessage(SetPtrValue(jsValue, newValueType, Empty, 0, 0, item.GetBool()))
				}
			} else {
				retArgs.SetInt8(1, int8(CVE_ERROR_NOT_FOUND_FIELD))
			}
			retArgs.SetBool(0, isSuccess)
			context.Response(retArgs.Package())
			context.Free()
		})
		//执行函数-同步
		event.On(ipc.Ln_EXECUTE_BIND_FUNC, func(context ipc.IIPCContext) {
			args := context.Arguments()
			fullName := args.GetString(args.Size() - 1)
			dataItems := args.RangeItems(0, args.Size()-1)
			var inArgument = ipc.NewArgumentList()
			inArgument.SetItems(dataItems)
			var jsValue, ok = VariableBind.GetValueBind(fullName)
			var (
				outParams []reflect.Value
				isSuccess bool
			)
			var buf = &bytes.Buffer{}
			defer buf.Reset()
			if ok {
				jsValue.Lock()
				defer jsValue.UnLock()
				var fnInfo = jsValue.getFuncInfo()
				if fnInfo != nil {
					outParams, isSuccess = jsValue.invoke(inArgument.ToReflectValue())
					inArgument.Clear()
					var outArguments = ipc.NewArgumentList()
					outArguments.ReflectValueConvert(outParams)
					binary.Write(buf, binary.BigEndian, isSuccess)
					binary.Write(buf, binary.BigEndian, outArguments.Package())
					outArguments.Clear()
				} else {
					binary.Write(buf, binary.BigEndian, isSuccess)
					binary.Write(buf, binary.BigEndian, int8(CVE_ERROR_NOT_FOUND_FUNC))
				}
			} else {
				binary.Write(buf, binary.BigEndian, isSuccess)
				binary.Write(buf, binary.BigEndian, int8(CVE_ERROR_NOT_FOUND_FUNC))
			}
			context.Response(buf.Bytes())
			buf = nil
			args.Clear()
			context.Free()
		})
	})
}
