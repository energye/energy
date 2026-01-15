//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package window

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "Cocoa/Cocoa.h"
#include "window_darwin.h"

// objective-c log > go println
void LogInfo(NSString* message) {
    const char* msg = [message cStringUsingEncoding:NSUTF8StringEncoding];
	GoLog((char *)msg);
}

void SetFrameless(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
	NSView* contentView = window.contentView; // view := NSView(win.contentView);
	[contentView setWantsLayer:YES]; // view.setWantsLayer(true);
	CALayer* layer = contentView.layer; // contentView.layer
	window.backgroundColor = [NSColor clearColor]; // window.setBackgroundColor(NSColor.clearColor);
	layer.backgroundColor = [NSColor whiteColor].CGColor; // layer.setBackgroundColor(NSColor.whiteColor.CGColor);
	layer.cornerRadius = 8.0; // layer.setCornerRadius(8.0);
	layer.masksToBounds = YES;
}

void SetWindowBackgroundColor(void* nsWindow, int r, int g, int b, int alpha) {
	[(NSWindow*)nsWindow setBackgroundColor:[NSColor colorWithRed:r/255.0 green:g/255.0 blue:b/255.0 alpha:alpha/255.0]];
}

void DragWindow(void* nsWindow) {
    NSEvent *currentMouseEvent = [NSApp currentEvent];

    if (!currentMouseEvent) {
        LogInfo(@"获取当前事件失败：事件为 nil");
        return;
    }
    if (currentMouseEvent.type != NSEventTypeLeftMouseDown) {
        LogInfo(@"获取当前事件失败：非左键按下事件");
        return;
    }
    //NSWindow* window = (NSWindow*)nsWindow;
	NSWindow* window = [currentMouseEvent window];
 	//NSWindow *window = [NSApp keyWindow];
    if (!window) {
        LogInfo(@"获取当前事件窗口失败");
        return;
    }
    [window performWindowDragWithEvent:currentMouseEvent];
}
*/
import "C"

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/pkgs/cocoa/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"unsafe"
)

const (
	NSWindowStyleMaskBorderless     = 0 // 窗口没有标题栏和按钮
	NSWindowStyleMaskTitled         = 1 // 窗口具有标题栏
	NSWindowStyleMaskClosable       = 2 // 窗口具有关闭按钮
	NSWindowStyleMaskMiniaturizable = 4 // 窗口可以被最小化
	NSWindowStyleMaskResizable      = 8 // 窗口可以调整大小
)

type IDarwinWindow interface {
	IWindow
	NSInstance() unsafe.Pointer
	NSWindow() lcl.NSWindow
	DragWindow()
}

type TWindow struct {
	TEnergyWindow
}

func (m *TWindow) DragWindow() {
	C.DragWindow(m.NSInstance())
}

func (m *TWindow) NSInstance() unsafe.Pointer {
	return unsafe.Pointer(m.NSWindow())
}

func (m *TWindow) NSWindow() lcl.NSWindow {
	return lcl.PlatformWindow(m.Instance())
}

func (m *TWindow) SetBackgroundColor(red, green, blue, alpha uint8) {
	C.SetWindowBackgroundColor(m.NSInstance(), C.int(red), C.int(green), C.int(blue), C.int(alpha))
}

func (m *TWindow) Frameless() {
	nsWindow := m.NSWindow()
	nsWindow.SetTitleBarAppearsTransparent(true)
	nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
	mask := uint(NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskResizable)
	options := application.GApplication.Options
	if options.DisableResize {
		mask ^= NSWindowStyleMaskResizable
	}
	nsWindow.SetStyleMask(mask)
	C.SetFrameless(m.NSInstance())
}

func (m *TWindow) _BeforeFormCreate() {

}

// SetOptions 设置webview窗口的选项配置
// 该方法用于配置*TWindow实例的各种选项参数
func (m *TWindow) SetOptions() {
	options := application.GApplication.Options
	if options.Width <= 0 {
		options.Width = m.Width()
	}
	if options.Height <= 0 {
		options.Height = m.Height()
	}
	m.SetCaption(options.Caption)
	m.SetBounds(options.X, options.Y, options.Width, options.Height)
	if options.Frameless {
		m.Frameless()
	}
}

func (m *TWindow) SetWindowState(value types.TWindowState) {
	m.windowsState = value
	switch value {
	case types.WsMaximized:
		window.ExitMinimized(m.NSInstance())
		window.Maximize(m.NSInstance())
	case types.WsNormal:
		window.Restore(m.NSInstance())
	case types.WsMinimized:
		window.Minimized(m.NSInstance())
	}
}

func (m *TWindow) WindowState() types.TWindowState {
	return m.windowsState
}

func (m *TWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		m.SetWindowState(types.WsFullScreen)
		window.EnterFullScreen(m.NSInstance())
	})
}

func (m *TWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.SetWindowState(types.WsNormal)
			window.ExitFullScreen(m.NSInstance())
		})
	}
}
