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
	"strings"
)

// 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	process.Current.SetBrowserId(browser.Identifier())                           // 当前进程 browserID
	process.Current.SetFrameId(frame.Identifier())                               // 当前进程 frameId
	ipc.RenderChan().SetRealityChannel(browser.Identifier(), frame.Identifier()) // 设置并更新真实的通道ID
	ipcRender.registerGoSyncReplayEvent()                                        // render ipc
	ipcRender.makeIPC(context)                                                   // render ipc make
	makeProcess(browser, frame, context)                                         // process make
}

// 应用运行 - 默认实现
func appMainRunCallback() {
	ipcBrowser.registerEvent() // browser ipc
}

// webkit - 默认实现
func appWebKitInitialized() {
	dragExtensionHandler() // drag extension handler
}

// 渲染进程消息 - 默认实现
func renderProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	if message.Name() == internalIPCJSExecuteGoEventReplay {
		result = ipcRender.ipcJSExecuteGoEventMessageReply(browser, frame, sourceProcess, message)
	} else if message.Name() == internalIPCGoExecuteJSEvent {
		result = ipcRender.ipcGoExecuteJSEvent(browser, frame, sourceProcess, message)
	}
	return
}

// 注册自定义协议 - 默认实现
func regCustomSchemes(registrar *TCefSchemeRegistrarRef) {
	if localLoadRes.enable() {
		// 以下几种默认的协议不去注册
		switch strings.ToUpper(localLoadRes.Scheme) {
		case "HTTP", "HTTPS", "FILE", "FTP", "ABOUT", "DATA":
			return
		}
		if application.IsSpecVer49() {
			registrar.AddCustomScheme(localLoadRes.Scheme, consts.CEF_SCHEME_OPTION_STANDARD|consts.CEF_SCHEME_OPTION_LOCAL)
		} else {
			registrar.AddCustomScheme(localLoadRes.Scheme,
				consts.CEF_SCHEME_OPTION_STANDARD|consts.CEF_SCHEME_OPTION_CORS_ENABLED|consts.CEF_SCHEME_OPTION_SECURE|consts.CEF_SCHEME_OPTION_FETCH_ENABLED)
		}
	}
}

// 主进程消息 - 默认实现
func browserProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
	if message.Name() == internalIPCJSExecuteGoEvent {
		result = ipcBrowser.jsExecuteGoMethodMessage(browser, frame, message)
	}
	return
}
