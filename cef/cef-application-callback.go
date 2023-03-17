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

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	fmt.Println("appOnContextCreated-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue(), "browserId:", browser.Identifier(), "frameId:", frame.Identifier())
	ipcRender.makeCtx(context)
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	fmt.Println("appMainRunCallback-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue())
}

// renderProcessMessageReceived 渲染进程消息 - 默认实现
func renderProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	fmt.Println("render name", message.Name())
	if message.Name() == internalProcessMessageIPCEmitReply {
		result = ipcRender.ipcJSExecuteGoEventMessageReply(browser, frame, sourceProcess, message)
	} else if message.Name() == internalProcessMessageIPCOn {
		result = ipcRender.ipcGoExecuteJSEvent(browser, frame, sourceProcess, message)
	}
	return
}

// browserProcessMessageReceived 主进程消息 - 默认实现
func browserProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	fmt.Println("browser name", message.Name())
	if message.Name() == internalProcessMessageIPCEmit {
		result = ipcBrowser.ipcGoExecuteMethodMessage(browser, frame, sourceProcess, message)
	} else if message.Name() == internalProcessMessageIPCOn {
		result = ipcBrowser.ipcOnMessage(browser, frame, sourceProcess, message)
	} else if message.Name() == internalProcessMessageIPCEmitReply {
		result = ipcBrowser.ipcGoExecuteMethodMessageReply(browser, frame, sourceProcess, message)
	}
	return
}
