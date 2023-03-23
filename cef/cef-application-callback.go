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
	"github.com/energye/energy/consts"
)

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	ipcRender.ipcChannelRender(browser, frame)
	ipcRender.makeIPC(browser, frame, context)
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	ipcBrowser.ipcChannelBrowser()
}

// appWebKitInitialized
func appWebKitInitialized() {
	fmt.Println("SetOnWebKitInitialized")
	v8Handler := V8HandlerRef.New()
	v8Handler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		fmt.Println("v8Handler.Execute", name)
		return true
	})
	//注册js
	var jsCode = `
            let bind;
            if (!bind) {
                bind = {};
            }
            (function () {
				Object.defineProperty(bind, 'myparam', {
					get(){
						native function GetMyParam();
						//return ipc.emitSync("testEmitSync", ["同步参数", 1, 2, 3, ["aaaa", "bbb", 6666]]);
						return GetMyParam();
					},
					set(v){
                    	native function SetMyParam();
						SetMyParam(v);
					}
				});
            })();
`
	RegisterExtension("v8/bind", jsCode, v8Handler)
}

// renderProcessMessageReceived 渲染进程消息 - 默认实现
func renderProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	if message.Name() == internalIPCJSExecuteGoEventReplay {
		result = ipcRender.ipcJSExecuteGoEventMessageReply(browser, frame, sourceProcess, message)
	} else if message.Name() == internalIPCGoExecuteJSEvent {
		result = ipcRender.ipcGoExecuteJSEvent(browser, frame, sourceProcess, message)
	}
	return
}

// browserProcessMessageReceived 主进程消息 - 默认实现
func browserProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
	if message.Name() == internalIPCJSExecuteGoEvent {
		result = ipcBrowser.jsExecuteGoMethodMessage(browser, frame, message)
	} else if message.Name() == internalIPCGoExecuteJSEventReplay {
		result = ipcBrowser.ipcGoExecuteMethodMessageReply(browser, frame, message)
	}
	return
}
