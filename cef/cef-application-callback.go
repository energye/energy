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
	"github.com/energye/energy/consts"
)

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	renderIPCCreate(browser.Identifier(), frame.Identifier()) // Go IPC render
	ipcRender.initRenderIPC()                                 // render ipc
	ipcRender.makeIPC(browser, frame, context)                // render ipc make
	bindRender.initBindIPC()                                  // render bind
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	browserIPCCreate()           // Go IPC browser
	ipcBrowser.initBrowserIPC()  // browser ipc
	bindBrowser.initBrowserIPC() // browser bind
}

// appWebKitInitialized - webkit - 默认实现
func appWebKitInitialized() {
	bindRender.webKitMakeBind() // render webkit bind make
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
