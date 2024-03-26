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

const (
	// CN: 要求最小liblcl二进制版本
	// EN: Requires a minimum liblcl binary version.
	requireMinBinaryVersion = 0x02040000
)

var (
	// 几个实例类，不需要Create即可访问，同时也不需要手动Free

	Application IApplication // 应用程序管理
	Screen      IScreen      // 屏幕
	Mouse       IMouse       // 鼠标
	Clipboard   IClipboard   // 剪切板
	Printer     IPrinter     // 打印机
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
	// 这个似乎得默认加上，锁定主线程，防止中间被改变
	runtime.LockOSThread()
	// 设置事件的回调函数，因go中callback数量有限，只好折中处理
	SetEventCallback(eventCallback)
	// 清除事件回调
	SetRemoveEventCallback(removeEventCallback)
	// 消息回调
	SetMessageCallback(messageCallback)
	// 调求回调CreateParams方法
	SetRequestCallCreateParamsCallback(requestCallCreateParamsCallback)
	// 主线程回调 异步
	SetThreadAsyncCallback(threadAsyncCallback)
	// 主线程同步 回调
	SetThreadSyncCallback(threadSyncCallback)

	// 导入几个实例类
	Application = AsApplication(Application_Instance())
	Screen = AsScreen(Screen_Instance())
	Mouse = AsMouse(Mouse_Instance())
	Clipboard = AsClipboard(Clipboard_Instance())
	Printer = AsPrinter(Printer_Instance())
}
