//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Go&JS变量绑定
package cef

import (
	"bytes"
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
	"reflect"
	"strconv"
	"unsafe"
)

// JSValue
// GO和JS变量类型接口
type JSValue interface {
	SetAnyValue(value interface{}) error //多类型值设置
	IntegerValue() (int32, error)        //返回 Integer 类型值, 成功 error = nil
	DoubleValue() (float64, error)       //返回 Double 类型值, 成功 error = nil
	StringValue() (string, error)        //返回 String 类型值, 成功 error = nil
	BooleanValue() (bool, error)         //返回 Boolean 类型值, 成功 error = nil
	IsInteger() bool                     //是否 Integer
	IsDouble() bool                      //是否 Double
	IsString() bool                      //是否 String
	IsBool() bool                        //是否 Bool
	IsArray() bool                       //是否 Array
	IsObject() bool                      //是否 Object
	IsFunction() bool                    //是否 Function
	IsNull() bool                        //是否 Null
	IsUndefined() bool                   //是否 Undefined
	AsV8Value() *V8Value                 //转换为 V8Value
	AsInteger() *JSInteger               //转换为 Integer 失败返回 nil
	AsDouble() *JSDouble                 //转换为 Double 失败返回 nil
	AsString() *JSString                 //转换为 String 失败返回 nil
	AsBoolean() *JSBoolean               //转换为 Boolean 失败返回 nil
	AsArray() *JSArray                   //转换为 Array 失败返回 nil
	AsFunction() *JSFunction             //转换为 Function 失败返回 nil
	Instance() uintptr                   //当前变量指针
	Ptr() unsafe.Pointer                 //当前变量指针
	Name() string                        //当前变量绑定的名称
	ValueType() *VT                      //变量类型
	Bytes() []byte                       //变量值转换为字节
	ValueToPtr() (unsafe.Pointer, error) //值转为指针
	Lock()                               //变量自有锁-加锁
	UnLock()                             //变量自有锁-释放锁
	invoke(inParams []reflect.Value) (outParams []reflect.Value, success bool)
	setPtr(ptr unsafe.Pointer)
	setInstance(instance uintptr)
	setName(name string)
	getValue() interface{}
	setValue(value interface{})
	getFuncInfo() *funcInfo
	setEventId(eventId uintptr)
	getEventId() uintptr
	isCommon() bool
	setThat(that JSValue)
}

// ICefV8Context bindGoToJS
//
// 主进程创建完之后和渲染进程每次创建之后调用
//
// 潜在问题，如果函数名包含数字可能会引起函数冲突，入参或出参类型不正确，导致调用失败
func bindGoToJS(browser *ICefBrowser, frame *ICefFrame) {
	objectTI.isBind = true //设置为已绑定
	//通过回调函数绑定到CEF
	VariableBind.callVariableBind(browser, frame)
	//通过直接绑定到CEF
	objectTI.bindToCEF()
	var valueBindInfos []*valueBindInfo
	for _, value := range VariableBind.binds() {
		if !value.isCommon() {
			continue
		}
		jsValue := value.(JSValue)
		var vBind = &valueBindInfo{
			BindType: uintptr(int32(jsValue.ValueType().Jsv)),
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
