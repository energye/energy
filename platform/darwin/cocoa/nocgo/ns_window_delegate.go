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
	"reflect"
	"runtime"
	"unsafe"
)

var (
	windowDelegateClass                            objc.Class
	sel_windowDidResize                            objc.SEL
	sel_windowWillEnterFullScreen                  objc.SEL
	sel_windowDidExitFullScreen                    objc.SEL
	sel_windowWillUseFullScreenPresentationOptions objc.SEL
)

type NSWindowDelegate struct {
	NSObject
}

func AsNSWindowDelegate(ptr unsafe.Pointer) INSWindowDelegate {
	if ptr == nil {
		return nil
	}
	m := new(NSWindowDelegate)
	m.SetInstance(ptr)
	return m
}

// NewWindowDelegate 为指定的 NSWindow 创建并设置窗口代理（Delegate）。
// 该函数会创建一个 TWindowDelegate 实例，将其与 Go 层的 NSWindow 对象关联，
//
// 参数:
//   - window: 需要设置代理的 NSWindow 对象指针。如果为 nil，则直接返回 nil。
//
// 返回值:
//   - unsafe.Pointer: 指向创建的 Objective-C Delegate 实例的指针。
//     调用者应保存此指针，以便在窗口销毁时调用 DestroyWindowDelegate 进行清理。
func NewWindowDelegate(window INSWindow) INSWindowDelegate {
	if window == nil {
		return nil
	}
	// 创建 windowDelegate 实例
	delegate := objc.ID(windowDelegateClass).Send(objc.RegisterName("new"))

	// 设置 window 属性
	nsWindow := window.(*NSWindow)
	delegate.Send(objc.RegisterName("setWindow:"), uintptr(unsafe.Pointer(nsWindow)))

	// 设置为窗口的 delegate
	nsWindowID := objc.ID(window.Instance())
	nsWindowID.Send(objc.RegisterName("setDelegate:"), delegate)

	return AsNSWindowDelegate(unsafe.Pointer(delegate))
}

func init() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	initSelectors()
	registerWindowDelegateClass()
}

func initSelectors() {
	sel_windowDidResize = objc.RegisterName("windowDidResize:")
	sel_windowWillEnterFullScreen = objc.RegisterName("windowWillEnterFullScreen:")
	sel_windowDidExitFullScreen = objc.RegisterName("windowDidExitFullScreen:")
	sel_windowWillUseFullScreenPresentationOptions = objc.RegisterName("window:willUseFullScreenPresentationOptions:")
}

func registerWindowDelegateClass() {
	var err error
	// 获取 NSWindowDelegate 和 NSToolbarDelegate 协议
	protoWindowDelegate := objc.GetProtocol("NSWindowDelegate")
	protoToolbarDelegate := objc.GetProtocol("NSToolbarDelegate")
	windowDelegateClass, err = objc.RegisterClass(
		"TWindowDelegate",
		objc.GetClass("NSObject"),
		[]*objc.Protocol{protoWindowDelegate, protoToolbarDelegate},
		[]objc.FieldDef{
			{
				Name:      "window",
				Type:      reflect.TypeOf(uintptr(0)),
				Attribute: objc.ReadWrite,
			},
		},
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("windowDidResize:"),
				Fn:  windowDidResize,
			},
			{
				Cmd: objc.RegisterName("windowWillEnterFullScreen:"),
				Fn:  windowWillEnterFullScreen,
			},
			{
				Cmd: objc.RegisterName("windowDidExitFullScreen:"),
				Fn:  windowDidExitFullScreen,
			},
			{
				Cmd: objc.RegisterName("window:willUseFullScreenPresentationOptions:"),
				Fn:  windowWillUseFullScreenPresentationOptions,
			},
		},
	)

	if err != nil {
		panic(err)
	}
}

// windowDidResize 处理窗口调整大小事件
func windowDidResize(self objc.ID, _cmd objc.SEL, notification objc.ID) {
	windowID := self.Send(objc.RegisterName("window"))
	if windowID == 0 {
		return
	}
	nsWindow := (*NSWindow)(unsafe.Pointer(windowID))
	nsWindow.doWindowDidResie()
}

// windowWillEnterFullScreen 处理即将进入全屏
func windowWillEnterFullScreen(self objc.ID, _cmd objc.SEL, notification objc.ID) {
	windowID := self.Send(objc.RegisterName("window"))
	if windowID == 0 {
		return
	}
	nsWindow := (*NSWindow)(unsafe.Pointer(windowID))
	nsWindow.doWindowWillEnterFullScreen()
}

// windowDidExitFullScreen 处理退出全屏
func windowDidExitFullScreen(self objc.ID, _cmd objc.SEL, notification objc.ID) {
	windowID := self.Send(objc.RegisterName("window"))
	if windowID == 0 {
		return
	}
	nsWindow := (*NSWindow)(unsafe.Pointer(windowID))
	nsWindow.doWindowDidExitFullScreen()
}

// windowWillUseFullScreenPresentationOptions 处理全屏选项
func windowWillUseFullScreenPresentationOptions(self objc.ID, _cmd objc.SEL,
	window objc.ID, options NSApplicationPresentationOptions) NSApplicationPresentationOptions {
	windowID := self.Send(objc.RegisterName("window"))
	if windowID == 0 {
		return options
	}
	nsWindow := (*NSWindow)(unsafe.Pointer(windowID))
	options = nsWindow.doWindowWillUseFullScreenPresentationOptions(options)
	return options
}
