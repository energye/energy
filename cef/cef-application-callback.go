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
	internalObjectRootName = "energy" // GO 和 V8Value 绑定根对象名
)

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
	emitHandler := V8HandlerRef.New()
	emitHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
		return false
	})
	onHandler := V8HandlerRef.New()
	onHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
		return false
	})
	ipc := V8ValueRef.NewObject(nil)
	ipc.setValueByKey(internalEmit, V8ValueRef.newFunction(internalEmit, emitHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	ipc.setValueByKey(internalOn, V8ValueRef.newFunction(internalOn, onHandler), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	context.Global().setValueByKey(internalIPCKey, ipc, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	fmt.Println("appMainRunCallback-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue())
	//internalBrowserIPCOnEventInit()
	//ipc.IPC.StartBrowserIPC()
	//indGoToJS(nil, nil)
}

func init() {

}
