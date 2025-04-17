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
	"github.com/cyber-xxm/energy/v2/cef/internal/process"
	"github.com/cyber-xxm/energy/v2/consts"
	"strings"
)

// 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	process.Current.SetBrowserId(browser.Identifier()) // 当前进程 browserID
	process.Current.SetFrameId(frame.Identifier())     // 当前进程 frameId
	ipcRender.makeIPC(frame.Identifier(), context)     // render ipc make
	dragExtension.make(frame.Identifier(), context)    // 拖拽JS扩展
	makeProcess(browser, frame, context)               // process make
}

// webkit - 默认实现
func appWebKitInitialized() {
	registerDragExtensionHandler() // drag extension handler
}

// 注册自定义协议 - 默认实现
func regCustomSchemes(registrar *TCefSchemeRegistrarRef) {
	if localLoadRes.enable() {
		// 以下几种默认的协议不去注册
		switch strings.ToUpper(localLoadRes.Scheme) {
		case "HTTP", "HTTPS", "FILE", "FTP", "ABOUT", "DATA":
			return
		}
		if application.Is49() {
			registrar.AddCustomScheme(localLoadRes.Scheme, consts.CEF_SCHEME_OPTION_STANDARD|consts.CEF_SCHEME_OPTION_LOCAL)
		} else {
			registrar.AddCustomScheme(localLoadRes.Scheme,
				consts.CEF_SCHEME_OPTION_STANDARD|consts.CEF_SCHEME_OPTION_CORS_ENABLED|consts.CEF_SCHEME_OPTION_SECURE|consts.CEF_SCHEME_OPTION_FETCH_ENABLED)
		}
	}
}
