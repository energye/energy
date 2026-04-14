//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/ebitengine/purego/objc"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"sync"
	"unsafe"
)

var (
	gEvaluateScriptEventID     int
	gNextEvaluateScriptEventID = func() int {
		gEvaluateScriptEventID++
		return gEvaluateScriptEventID
	}
	gEvaluateScriptEventCallback = sync.Map{}
)

func (m *WkWebView) ExecuteScriptCallback(script string, callback TEvaluateScriptCallbackEvent) {
	if script == "" || callback == nil || m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	eventID := gNextEvaluateScriptEventID()
	jsString := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"), script,
	)
	blockFunc := func(block objc.Block, result objc.ID, nsError objc.ID) {
		var resultStr, errStr string
		if result != 0 {
			resultStr = goStringFromObjcString(result)
		}
		if nsError != 0 {
			localizedDesc := nsError.Send(objc.RegisterName("localizedDescription"))
			if localizedDesc != 0 {
				errStr = goStringFromObjcString(localizedDesc)
			}
		}
		if cb, ok := gEvaluateScriptEventCallback.Load(eventID); ok {
			gEvaluateScriptEventCallback.Delete(eventID)
			cb.(TEvaluateScriptCallbackEvent)(resultStr, errStr)
		}
	}
	block := objc.NewBlock(blockFunc)
	gEvaluateScriptEventCallback.Store(eventID, callback)
	webview.Send(objc.RegisterName("evaluateJavaScript:completionHandler:"),
		jsString, block,
	)
}

func goStringFromObjcString(objcStr objc.ID) string {
	if objcStr == 0 {
		return ""
	}
	utf8Ptr := objcStr.Send(objc.RegisterName("UTF8String"))
	if utf8Ptr == 0 {
		return ""
	}
	return goStringFromPtr(uintptr(utf8Ptr))
}

func goStringFromPtr(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&ptr))
}
