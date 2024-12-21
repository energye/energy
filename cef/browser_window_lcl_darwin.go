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
// +build darwin

package cef

/*
#cgo CFLAGS: -mmacosx-version-min=10.12 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.12 -framework Cocoa

#include "Cocoa/Cocoa.h"
#include "browser_window_drag_darwin.h"

// objective-c log > go println
void LogInfo(NSString* message) {
    const char* msg = [message cStringUsingEncoding:NSUTF8StringEncoding];
	GoLog((char *)msg);
}

void setFrameless(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
	NSView* contentView = window.contentView; // view := NSView(win.contentView);
	[contentView setWantsLayer:YES]; // view.setWantsLayer(true);
	CALayer* layer = contentView.layer; // contentView.layer
	window.backgroundColor = [NSColor clearColor]; // window.setBackgroundColor(NSColor.clearColor);
	layer.backgroundColor = [NSColor whiteColor].CGColor; // layer.setBackgroundColor(NSColor.whiteColor.CGColor);
	layer.cornerRadius = 8.0; // layer.setCornerRadius(8.0);
	layer.masksToBounds = YES;
}

void setWindowBackgroundColor(void* nsWindow, int r, int g, int b, int alpha) {
	[(NSWindow*)nsWindow setBackgroundColor:[NSColor colorWithRed:r/255.0 green:g/255.0 blue:b/255.0 alpha:alpha/255.0]];
}

void initDragEventListeners() {
	[NSEvent addLocalMonitorForEventsMatchingMask:NSEventMaskLeftMouseDown handler:^NSEvent * _Nullable(NSEvent * _Nonnull event) {
		NSWindow* eventWindow = [event window];
		if (eventWindow == nil) {
			return event;
		}
		//LogInfo(@"LeftMouseDown");
		BOOL flag = CanDrag(eventWindow);
		int32_t titleBarHeight = GetTitlebarHeight(eventWindow);
		if (flag) {
			[eventWindow performWindowDragWithEvent:event];
		} else if (titleBarHeight > 0) {
			NSPoint location = [event locationInWindow];
			NSRect frame = [eventWindow frame];
			int32_t titleBarHeight = GetTitlebarHeight(eventWindow);
			if(location.y > frame.size.height - titleBarHeight) {
			  [eventWindow performWindowDragWithEvent:event];
			}
		}
		return event;
	}];

	[NSEvent addLocalMonitorForEventsMatchingMask:NSEventMaskMouseMoved handler:^NSEvent * _Nullable(NSEvent * _Nonnull event) {
		NSWindow* window = [event window];
		if (window == nil) {
			return event;
		}
        NSRect windowFrame = [window frame];
		NSPoint locationInWindow = [event locationInWindow];
		// 将左下角坐标转换为左上角坐标
		CGFloat newY = NSHeight(windowFrame) - locationInWindow.y;
		NSPoint adjustedLocation = NSMakePoint(locationInWindow.x, newY);
		CheckDraggableRegions(window, (int32_t)adjustedLocation.x, (int32_t)adjustedLocation.y);
		return event;
	}];

	[NSEvent addLocalMonitorForEventsMatchingMask:NSEventMaskLeftMouseUp handler:^NSEvent * _Nullable(NSEvent * _Nonnull event) {
		NSWindow* eventWindow = [event window];
		if (eventWindow == nil) {
			return event;
        }
		//LogInfo(@"LeftMouseUp");
		SetCanDrag(eventWindow, false);
		return event;
	}];
}
*/
import "C"

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

const (
	NSWindowStyleMaskBorderless     = 0 // 窗口没有标题栏和按钮
	NSWindowStyleMaskTitled         = 1 // 窗口具有标题栏
	NSWindowStyleMaskClosable       = 2 // 窗口具有关闭按钮
	NSWindowStyleMaskMiniaturizable = 4 // 窗口可以被最小化
	NSWindowStyleMaskResizable      = 8 // 窗口可以调整大小
)

type PlatformWindow struct {
	lcl.NSWindow
}

func (m *PlatformWindow) Instance() unsafe.Pointer {
	return unsafe.Pointer(m.NSWindow)
}

func (m *PlatformWindow) SetBackgroundColor(red, green, blue, alpha uint8) {
	C.setWindowBackgroundColor(m.Instance(), C.int(red), C.int(green), C.int(blue), C.int(alpha))
}

func (m *LCLBrowserWindow) PlatformWindow() *PlatformWindow {
	return &PlatformWindow{NSWindow: m.TForm.PlatformWindow()}
}

func (m *LCLBrowserWindow) initDragEventListeners() {
	C.initDragEventListeners()
}

func (m *LCLBrowserWindow) frameless() {
	nsWindow := m.PlatformWindow()
	nsWindow.SetTitleBarAppearsTransparent(true)
	nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
	mask := uint(NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskResizable)
	wp := m.WindowProperty()
	if !wp.EnableResize {
		mask ^= NSWindowStyleMaskResizable
	}
	nsWindow.SetStyleMask(mask)

	C.setFrameless(nsWindow.Instance())
}

func (m *LCLBrowserWindow) SetRoundRectRgn(rgn int) {
	if m.rgn == 0 && rgn > 0 {
		m.rgn = rgn
	}
}

// FullScreen 窗口全屏
func (m *LCLBrowserWindow) FullScreen() {
	if m.IsFullScreen() {
		return
	}
	RunOnMainThread(func() {
		m.WindowProperty().current.windowState = types.WsFullScreen
		m.WindowProperty().current.previousWindowPlacement = m.BoundsRect()
		m.SetWindowState(types.WsFullScreen)
		m.SetBoundsRect(m.Monitor().BoundsRect())
	})
}

// ExitFullScreen 窗口退出全屏
func (m *LCLBrowserWindow) ExitFullScreen() {
	if m.IsFullScreen() {
		wp := m.WindowProperty()
		RunOnMainThread(func() {
			wp.current.windowState = types.WsNormal
			m.SetBoundsRect(m.WindowProperty().current.previousWindowPlacement)
			m.SetWindowState(types.WsNormal)
		})
	}
}
