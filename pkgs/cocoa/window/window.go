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

// 最大化

void WindowMaximize(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口不可用或不可调整大小，无法执行最大化");
        return;
    }
    //if (!window.isZoomed) {
        [window zoom:nil];
    //}
}

void WindowRestore(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    //if (window.isZoomed) {
        [window zoom:nil];
    //}
}

BOOL WindowIsMaximize(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return NO;
    }
    return window.isZoomed;
}

// 最小化

void WindowMinimized(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    if (![window isMiniaturized]) {
        [window miniaturize:nil];
    }
}

void WindowExitMinimized(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    if ([window isMiniaturized]) {
        [window deminiaturize:nil];
    }
}

// 全屏

void WindowEnterFullScreen(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    [window toggleFullScreen:nil];
}

void WindowExitFullScreen(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    [window toggleFullScreen:nil];
}

*/
import "C"
import "unsafe"

func IsMaximize(window unsafe.Pointer) bool {
	if window == nil {
		return false
	}
	return bool(C.WindowIsMaximize(window))
}

func Restore(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowRestore(window)
}

func Minimized(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowMinimized(window)
}

func ExitMinimized(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowExitMinimized(window)
}

func Maximize(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowMaximize(window)
}

func EnterFullScreen(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowEnterFullScreen(window)
}

func ExitFullScreen(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowExitFullScreen(window)
}
