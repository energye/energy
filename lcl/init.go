//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"fmt"
	. "github.com/energye/energy/v2/api"
	"os"
	"runtime"
)

var (
	// Application
	//
	//  The TApplication singleton.
	//  is a unit global variable with the
	//  instance representing the currently executing application.
	//  The value for the variable is assigned in the initialization section for the unit when the application is started.
	//  It is freed in the finalization section when the application is terminated.
	//  Use the
	//  TApplicationProperties
	//  component to provide design-time configuration settings for properties (including event handlers) which are automatically applied to Application at run-time.
	Application IApplication
	Screen      IScreen    // 屏幕
	Mouse       IMouse     // 鼠标
	Clipboard   IClipboard // 剪切板
	Printer     IPrinter   // 打印机
)

func toVersionString(ver uint32) string {
	if byte(ver) == 0 {
		return fmt.Sprintf("%d.%d.%d", byte(ver>>24), byte(ver>>16), byte(ver>>8))
	}
	return fmt.Sprintf("%d.%d.%d.%d", byte(ver>>24), byte(ver>>16), byte(ver>>8), byte(ver))
}

func LCLInit() {
	defer func() {
		if err := recover(); err != nil {
			showError(err)
			os.Exit(1)
		}
	}()
	runtime.LockOSThread()
	SetEventCallback(eventCallback)
	SetRemoveEventCallback(removeEventCallback)
	SetMessageCallback(messageCallback)
	SetRequestCallCreateParamsCallback(requestCallCreateParamsCallback)
	SetRequestCallFormCreateCallback(requestCallFormCreateCallback)
	// 主线程回调 异步
	SetThreadAsyncCallback(threadAsyncCallback)
	// 主线程同步 回调
	SetThreadSyncCallback(threadSyncCallback)

	Application = AsApplication(Application_Instance())
	Screen = AsScreen(Screen_Instance())
	Mouse = AsMouse(Mouse_Instance())
	Clipboard = AsClipboard(Clipboard_Instance())
	Printer = AsPrinter(Printer_Instance())
}
