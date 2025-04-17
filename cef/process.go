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
	"github.com/cyber-xxm/energy/v2/consts"
)

// process
const (
	internalProcess          = "process"
	internalProcessBrowserId = "browserId"
	internalProcessFrameId   = "frameId"
)

var _processObject *ICefV8Value

// makeProcess 进程扩展变量
func makeProcess(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	if _processObject != nil {
		// 刷新时释放掉
		_processObject.Free()
	}
	// process
	_processObject = V8ValueRef.NewObject(nil, nil)
	_processObject.setValueByKey(internalProcessBrowserId, V8ValueRef.NewInt(browser.Identifier()), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	_processObject.setValueByKey(internalProcessFrameId, V8ValueRef.NewString(frame.Identifier()), consts.V8_PROPERTY_ATTRIBUTE_READONLY)

	// process key to v8 global
	context.Global().setValueByKey(internalProcess, _processObject, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
}
