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
	//if !objectTI.isBind {
	//__idReset()
	//clearValueBind()
	//bindGoToJS(browser, frame)
	//}
	//ipc.IPC.CreateRenderIPC(browser.Identifier(), frame.Identifier())
	fmt.Println("\tbindCount", VariableBind.bindCount())
	binds := VariableBind.binds()
	for name, value := range binds {
		fmt.Println("\t", name, value.ValueType())
	}
	ctx.makeCtx(context)
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	fmt.Println("appMainRunCallback-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue())
	//internalBrowserIPCOnEventInit()
	//ipc.IPC.StartBrowserIPC()
	//indGoToJS(nil, nil)
}

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

func (m *contextCreate) ipcEmitExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("emit handler name:", name)
	return false
}

func (m *contextCreate) ipcOnExecute(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("on handler name:", name, "object", object.IsObject())
	fmt.Println("on handler arguments:", arguments.Size())
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
	// bind accessor
	objectAccessor := V8AccessorRef.New()
	objectAccessor.Get(m.bindGet)
	objectAccessor.Set(m.bindSet)
	// bind object
	m.bind = V8ValueRef.NewObject(objectAccessor)
	// global to v8 bind objectRootName
	context.Global().setValueByKey(internalObjectRootName, m.bind, consts.V8_PROPERTY_ATTRIBUTE_NONE)
}

func (m *contextCreate) bindGet(name string, object *ICefV8Value, retVal *ResultV8Value, exception *Exception) bool {
	fmt.Println("get accessor name:", name)
	return false
}

func (m *contextCreate) bindSet(name string, object *ICefV8Value, value *ICefV8Value, exception *Exception) bool {
	fmt.Println("set accessor name:", name)
	return false
}
