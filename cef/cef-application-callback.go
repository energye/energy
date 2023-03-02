//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// application event 默认事件实现
package cef

import (
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
)

const (
	internalIPCKey = "ipc"  // JavaScript -> ipc 事件驱动, 根对象名
	internalEmit   = "emit" // JavaScript -> ipc.emit 在 JavaScript 触发 GO 监听事件函数名
	internalOn     = "on"   // JavaScript -> ipc.on 在 JavaScript 监听事件, 提供给 GO 调用
)

var (
	internalObjectRootName = "energy"     // GO 和 V8Value 绑定根对象名
	ctx                    *contextCreate //
	mRun                   *mainRun       //
)

type contextCreate struct {
	ipc  *ICefV8Value
	bind *ICefV8Value
}

type mainRun struct {
}

func init() {
	if common.Args.IsMain() {
		mRun = &mainRun{}
	} else if common.Args.IsRender() {
		ctx = &contextCreate{}
	}
}

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	fmt.Println("appOnContextCreated-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue(), "browserId:", browser.Identifier(), "frameId:", frame.Identifier())
	ctx.makeCtx(context)
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	fmt.Println("appMainRunCallback-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue())
	//internalBrowserIPCOnEventInit()
	//ipc.IPC.StartBrowserIPC()
	//indGoToJS(nil, nil)
}

// makeCtx ipc 和 bind
func (m *contextCreate) makeCtx(context *ICefV8Context) {
	ctx.makeIPC(context)
	ctx.makeBind(context)
}

// makeIPC ipc
func (m *contextCreate) makeIPC(context *ICefV8Context) {
	// ipc emit
	emitHandler := V8HandlerRef.New()
	emitHandler.Execute(m.ipcEmitExecute)
	// ipc on
	onHandler := V8HandlerRef.New()
	onHandler.Execute(m.ipcOnExecute)
	// ipc object
	m.ipc = V8ValueRef.NewObject(nil)
	m.ipc.setValueByKey(internalEmit, V8ValueRef.newFunction(internalEmit, emitHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.ipc.setValueByKey(internalOn, V8ValueRef.newFunction(internalOn, onHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	// global to v8 ipc key
	context.Global().setValueByKey(internalIPCKey, m.ipc, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}

// ipcEmitExecute ipc.emit 执行
func (m *contextCreate) ipcEmitExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("emit handler name:", name, "arguments-size:", arguments.Size())
	for i := 0; i < arguments.Size(); i++ {
		fmt.Println("\t", arguments.Get(i))
	}
	return false
}

// ipcEmitExecute ipc.on 执行
func (m *contextCreate) ipcOnExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("on handler name:", name, "arguments-size:", arguments.Size())
	frame := V8ContextRef.Current().Frame()
	fmt.Println("frame", frame.Identifier())
	sendBrowserProcessMsg := ProcessMessageRef.New("ipcOnExecute")
	sendBrowserProcessMsg.ArgumentList().SetString(0, "ipcOnExecute测试值")
	frame.SendProcessMessage(consts.PID_BROWSER, sendBrowserProcessMsg)
	sendBrowserProcessMsg.Free()
	return false
}

// makeBind bind object accessor
func (m *contextCreate) makeBind(context *ICefV8Context) {
	var (
		bindRootObjectAccessor *ICefV8Accessor // bind accessor
		bindSubObjectAccessor  *ICefV8Accessor // bind sub object accessor
		bindFuncHandler        *ICefV8Handler  // bind func handler
	)
	bindRootObjectAccessor = V8AccessorRef.New()
	bindRootObjectAccessor.Get(m.bindGet)
	bindRootObjectAccessor.Set(m.bindSet)
	// bind object
	m.bind = V8ValueRef.NewObject(bindRootObjectAccessor)

	fmt.Println("\tbindCount", VariableBind.bindCount())
	binds := VariableBind.binds()
	for name, value := range binds {
		fmt.Println("\t", name, "\t", value.ValueType().ToString())
		var v8value *ICefV8Value
		if value.IsFunction() {
			if bindFuncHandler == nil {
				bindFuncHandler = V8HandlerRef.New()
				bindFuncHandler.Execute(m.bindFuncExecute)
			}
			v8value = V8ValueRef.newFunction(name, bindFuncHandler)
			m.bind.setValueByKey(name, v8value, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		} else {
			if value.IsString() {
				v8value = V8ValueRef.NewString("")
			} else if value.IsInteger() {
				v8value = V8ValueRef.NewInt(0)
			} else if value.IsDouble() {
				v8value = V8ValueRef.NewDouble(0.0)
			} else if value.IsBool() {
				v8value = V8ValueRef.NewBool(false)
			} else if value.IsNull() {
				v8value = V8ValueRef.NewNull()
			} else if value.IsUndefined() {
				v8value = V8ValueRef.NewUndefined()
			} else if value.IsObject() {
				if bindSubObjectAccessor == nil {
					bindSubObjectAccessor = V8AccessorRef.New()
					bindSubObjectAccessor.Get(m.bindSubObjectGet)
					bindSubObjectAccessor.Set(m.bindSubObjectSet)
				}
				v8value = V8ValueRef.NewObject(bindSubObjectAccessor)
			} else if value.IsArray() {
				v8value = V8ValueRef.NewArray(0)
			}
			if v8value != nil {
				m.bind.setValueByAccessor(name, consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE)
				m.bind.setValueByKey(name, v8value, consts.V8_PROPERTY_ATTRIBUTE_NONE)
			}
		}
	}

	// global to v8 bind objectRootName
	context.Global().setValueByKey(internalObjectRootName, m.bind, consts.V8_PROPERTY_ATTRIBUTE_NONE)
}

// bindFuncExecute 绑定函数执行
func (m *contextCreate) bindFuncExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("bindFuncExecute handler name:", name)
	return false
}

// bindObjectGet 绑定对象取值
func (m *contextCreate) bindSubObjectGet(name string, object *ICefV8Value, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("bindSubObjectGet accessor name:", name)
	return false
}

// bindObjectSet 绑定对象赋值
func (m *contextCreate) bindSubObjectSet(name string, object *ICefV8Value, value *ICefV8Value, exception *Exception) bool {
	fmt.Println("bindSubObjectSet accessor name:", name)
	return false
}

// bindGet 绑定字段取值
func (m *contextCreate) bindGet(name string, object *ICefV8Value, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("bindGet accessor name:", name)
	return false
}

// bindSet 绑定字段赋值
func (m *contextCreate) bindSet(name string, object *ICefV8Value, value *ICefV8Value, exception *Exception) bool {
	fmt.Println("bindSet accessor name:", name)
	return false
}
