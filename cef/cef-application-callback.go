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
	"github.com/energye/energy/ipc"
)

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame) {
	fmt.Println("appOnContextCreated-ProcessTypeValue:", common.Args.ProcessType(), application.ProcessTypeValue(), "id", browser.Identifier(), frame.Identifier())
	if !objectTI.isBind {
		__idReset()
		clearValueBind()
		bindGoToJS(browser, frame)
	}
	ipc.IPC.CreateRenderIPC(browser.Identifier(), frame.Identifier())
}
