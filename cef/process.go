//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"github.com/energye/energy/consts"
)

// process
const (
	internalProcess          = "process"
	internalProcessBrowserId = "browserId"
	internalProcessFrameId   = "frameId"
)

// makeProcess 进程扩展变量
func makeProcess(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	// process
	process := V8ValueRef.NewObject(nil)
	process.setValueByKey(internalProcessBrowserId, V8ValueRef.NewInt(browser.Identifier()), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	process.setValueByKey(internalProcessFrameId, V8ValueRef.NewString(fmt.Sprintf("%d", frame.Identifier())), consts.V8_PROPERTY_ATTRIBUTE_READONLY)

	// process key to v8 global
	context.Global().setValueByKey(internalProcess, process, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}
