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
	"reflect"
)

const (
	ln_IPC_GoEmitJS         = "IPCGoEmitJS"                   //Go执行Js on监听
	ln_GET_BIND_FIELD_VALUE = "internal_GET_BIND_FIELD_VALUE" //browse进程监听获取字段值
	ln_SET_BIND_FIELD_VALUE = "internal_SET_BIND_FIELD_VALUE" //browse进程监听设置字段值
	ln_EXECUTE_BIND_FUNC    = "internal_EXECUTE_BIND_FUNC"    //browse进程监听执行绑定函数
	ln_onConnectEvent       = "connect"
)

func internalIPCNameCheck(name string) bool {
	if name == ln_IPC_GoEmitJS || name == ln_GET_BIND_FIELD_VALUE || name == ln_SET_BIND_FIELD_VALUE || name == ln_EXECUTE_BIND_FUNC || name == ln_onConnectEvent {
		return true
	}
	return false
}

func (m *ipcChannel) internalBrowserIPCOnEventInit() {
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
	m.browser.SetOnEvent(func(event IEventOn) {
		//var fieldReadWriteLock = new(sync.Mutex)
		//获取字段值-同步
		event.On(ln_GET_BIND_FIELD_VALUE, func(context IIPCContext) {
			args := context.Arguments()
			defer args.Clear()
			fullName := args.GetString(0)
			var jsValue, ok = VariableBind.getValueBind(fullName)
			buf := &bytes.Buffer{}
			defer buf.Reset()
			if ok {
				jsValue.lock()
				defer jsValue.unLock()
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
		event.On(ln_SET_BIND_FIELD_VALUE, func(context IIPCContext) {
			args := context.Arguments()
			defer args.Clear()
			fullName := args.GetString(0)
			item := args.Item(1)
			newValueType := item.VTypeToJS()
			jsValue, ok := VariableBind.getValueBind(fullName)
			isSuccess := false
			retArgs := NewArgumentList()
			defer retArgs.Clear()
			if ok {
				jsValue.lock()
				defer jsValue.unLock()
				switch newValueType { //设置值，这里是通用类型只需要知道js里设置的什么类型即可
				case V8_VALUE_STRING:
					isSuccess = empty == cefErrorMessage(_setPtrValue(jsValue, newValueType, item.GetString(), 0, 0, false))
				case V8_VALUE_INT:
					isSuccess = empty == cefErrorMessage(_setPtrValue(jsValue, newValueType, empty, item.GetInt32(), 0, false))
				case V8_VALUE_DOUBLE:
					isSuccess = empty == cefErrorMessage(_setPtrValue(jsValue, newValueType, empty, 0, item.GetFloat64(), false))
				case V8_VALUE_BOOLEAN:
					isSuccess = empty == cefErrorMessage(_setPtrValue(jsValue, newValueType, empty, 0, 0, item.GetBool()))
				}
			} else {
				retArgs.SetInt8(1, int8(CVE_ERROR_NOT_FOUND_FIELD))
			}
			retArgs.SetBool(0, isSuccess)
			context.Response(retArgs.Package())
			context.Free()
		})
		//执行函数-同步
		event.On(ln_EXECUTE_BIND_FUNC, func(context IIPCContext) {
			args := context.Arguments()
			fullName := args.GetString(args.Size() - 1)
			dataItems := args.RangeItems(0, args.Size()-1)
			var inArgument = NewArgumentList()
			inArgument.SetItems(dataItems)
			var jsValue, ok = VariableBind.getValueBind(fullName)
			var (
				outParams []reflect.Value
				isSuccess bool
			)
			var buf = &bytes.Buffer{}
			defer buf.Reset()
			if ok {
				jsValue.lock()
				defer jsValue.unLock()
				var fnInfo = jsValue.getFuncInfo()
				if fnInfo != nil {
					outParams, isSuccess = jsValue.invoke(inArgument.ToReflectValue())
					inArgument.Clear()
					var outArguments = NewArgumentList()
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

func (m *ipcChannel) internalRenderIPCOnEventInit() {
}
