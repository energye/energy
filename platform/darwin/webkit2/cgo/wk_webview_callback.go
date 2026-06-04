//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cgo

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa

#include "wk_webview.h"
*/
import "C"
import (
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

//export evaluateScriptCallback
func evaluateScriptCallback(cCallbackID C.int, resC *C.char, errC *C.char) {
	var (
		result, err string
		callbackID  int
	)
	callbackID = int(cCallbackID)
	if resC != nil {
		result = C.GoString(resC)
	}
	if errC != nil {
		err = C.GoString(errC)
	}
	if callback, ok := gEvaluateScriptEventCallback.Load(callbackID); ok {
		gEvaluateScriptEventCallback.Delete(callbackID)
		callback.(TEvaluateScriptCallbackEvent)(result, err)
	}
}

func (m *WkWebView) ExecuteScriptCallback(script string, callback TEvaluateScriptCallbackEvent) {
	if script == "" || callback == nil {
		return
	}
	webview := unsafe.Pointer(m.Instance())
	cScript := C.CString(script)
	defer C.free(unsafe.Pointer(cScript))
	eventID := gNextEvaluateScriptEventID()
	cEventID := C.int(eventID)
	cCallback := (C.CGoEvaluateScriptCallback)(C.evaluateScriptCallback)
	var goTmpCallback TEvaluateScriptCallbackEvent
	goTmpCallback = func(result string, err string) {
		callback(result, err)
	}
	gEvaluateScriptEventCallback.Store(eventID, goTmpCallback)
	C.WebViewEvaluateScriptCallback(webview, cEventID, cScript, cCallback)
}
