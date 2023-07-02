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
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/internal/process"
	"github.com/energye/energy/v2/consts"
)

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	process.Current.SetBrowserId(browser.Identifier())                           // 当前进程 browserID
	process.Current.SetFrameId(frame.Identifier())                               // 当前进程 frameId
	ipc.RenderChan().SetRealityChannel(browser.Identifier(), frame.Identifier()) // 设置并更新真实的通道ID
	ipcRender.registerGoSyncReplayEvent()                                        // render ipc
	ipcRender.makeIPC(context)                                                   // render ipc make
	makeProcess(browser, frame, context)                                         // process make
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	ipcBrowser.registerEvent() // browser ipc
}

// appWebKitInitialized - webkit - 默认实现
func appWebKitInitialized() {
	//	var myparamValue string
	//	v8Handler := V8HandlerRef.New()
	//	v8Handler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
	//		fmt.Println("v8Handler.Execute", name)
	//		var result bool
	//		if name == "GetMyParam" {
	//			result = true
	//			retVal.SetResult(V8ValueRef.NewString(myparamValue))
	//		} else if name == "SetMyParam" {
	//			if arguments.Size() > 0 {
	//				newValue := arguments.Get(0)
	//				fmt.Println("value is string:", newValue.IsString())
	//				fmt.Println("value:", newValue.GetStringValue())
	//				myparamValue = newValue.GetStringValue()
	//				newValue.Free()
	//			}
	//			result = true
	//		}
	//		return result
	//	})
	//	//注册js
	//	var jsCode = `
	//            let energyExtension;
	//            if (!energyExtension) {
	//                energyExtension = {};
	//            }
	//            (function () {
	//                test.__defineGetter__('mouseover', function (e) {
	//                    native function mouseover();
	//                    return mouseover(e);
	//                });
	//                test.__defineSetter__('mousemove', function (e) {
	//                    native function mousemove();
	//                    mousemove(e);
	//                });
	//            })();
	//`
	//	// 注册JS 和v8处理器
	//	RegisterExtension("v8/test", jsCode, v8Handler)
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
	}
	return
}
