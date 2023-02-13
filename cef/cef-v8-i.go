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
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl/api"
	"reflect"
	"strconv"
	"unsafe"
)

type JSValue interface {
	SetAnyValue(value interface{}) error
	IntegerValue() (int32, error)
	DoubleValue() (float64, error)
	StringValue() (string, error)
	BooleanValue() (bool, error)
	IsInteger() bool
	IsDouble() bool
	IsString() bool
	IsBool() bool
	IsArray() bool
	IsObject() bool
	IsFunction() bool
	IsNull() bool
	IsUndefined() bool
	AsInteger() (*JSInteger, error)
	AsDouble() (*JSDouble, error)
	AsString() (*JSString, error)
	AsBoolean() (*JSBoolean, error)
	AsArray() (*JSArray, error)
	AsFunction() (*JSFunction, error)
	Instance() uintptr
	Ptr() unsafe.Pointer
	Name() string
	ValueType() V8_JS_VALUE_TYPE
	invoke(inParams []reflect.Value) (outParams []reflect.Value, success bool)
	setPtr(ptr unsafe.Pointer)
	setInstance(instance uintptr)
	setName(name string)
	getValue() interface{}
	setValue(value interface{})
	setValueType(vType V8_JS_VALUE_TYPE)
	getFuncInfo() *funcInfo
	setEventId(eventId uintptr)
	getEventId() uintptr
	isCommon() bool
	setThat(that JSValue)
	Bytes() []byte
	ValueToPtr() (unsafe.Pointer, error)
	Lock()
	UnLock()
}

// ICefV8Context bindGoToJS
//
// 主进程创建完之后和渲染进程每次创建之后调用
//
// 潜在问题，如果函数名包含数字可能会引起函数冲突，入参或出参类型不正确，导致调用失败
func bindGoToJS(browser *ICefBrowser, frame *ICefFrame) {
	//变量绑定回调函数
	VariableBind.callVariableBind(browser, frame)
	var valueBindInfos []*valueBindInfo
	logger.Debug("Total number of bindings：", VariableBind.ValueBindCount())
	for _, value := range VariableBind.valuesBind {
		if !value.isCommon() {
			continue
		}
		jsValue := value.(JSValue)
		var valueType = int32(jsValue.ValueType())
		var vBind = &valueBindInfo{
			BindType: uintptr(valueType),
		}
		vBind.Name = api.PascalStr(jsValue.Name())
		vBind.EventId = jsValue.getEventId()
		valueBindInfos = append(valueBindInfos, vBind)
		if jsValue.IsFunction() {
			var inParamBuf bytes.Buffer
			var outParamBuf bytes.Buffer
			fnInfo := jsValue.getFuncInfo()
			fnInNum := len(fnInfo.InParam)
			fnOutNum := len(fnInfo.OutParam)
			vBind.FnInNum = uintptr(fnInNum)
			vBind.FnOutNum = uintptr(fnOutNum)
			for i, inParamType := range fnInfo.InParam {
				if i > 0 {
					inParamBuf.WriteString(",")
				}
				inParamBuf.WriteString(strconv.Itoa(int(inParamType.Jsv)))
			}
			vBind.FnInParamType = api.PascalStr(inParamBuf.String())
			for i, outParamType := range fnInfo.OutParam {
				if i > 0 {
					outParamBuf.WriteString(",")
				}
				outParamBuf.WriteString(strconv.Itoa(int(outParamType.Jsv)))
			}
			vBind.FnOutParamType = api.PascalStr(outParamBuf.String())
		}
	}
	if len(valueBindInfos) > 0 {
		for i := 0; i < len(valueBindInfos); i++ {
			imports.Proc(internale_CEFV8ValueRef_CommonValueBindInfo).Call(uintptr(unsafe.Pointer(valueBindInfos[i])))
		}
	}
}
