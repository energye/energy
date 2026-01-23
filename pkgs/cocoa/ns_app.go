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

package cocoa

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>

void AppDockHide() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyProhibited];
}

void AppDockShow() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyRegular];
}

void AppSetPresentationOptions(NSApplicationPresentationOptions opts) {
	[[NSApplication sharedApplication] setPresentationOptions:opts];
}

void AppSetMainMenu(void* nsMenu) {
	NSMenu* menu = (NSMenu*)nsMenu;
    if (!menu) {
        NSLog(@"AppSetMainMenu menu nil");
        return ;
    }
    NSApplication *app = [NSApplication sharedApplication];
    [app setMainMenu:menu];
}

*/
import "C"
import "unsafe"

const (
	NSApplicationPresentationDefault                   = uint(C.NSApplicationPresentationDefault)                   //默认行为（隐藏菜单栏、Dock 等）
	NSApplicationPresentationAutoHideMenuBar           = uint(C.NSApplicationPresentationAutoHideMenuBar)           //自动隐藏菜单栏
	NSApplicationPresentationHideMenuBar               = uint(C.NSApplicationPresentationHideMenuBar)               //永久隐藏菜单栏（即使鼠标移到顶部也不显示）
	NSApplicationPresentationAutoHideDock              = uint(C.NSApplicationPresentationAutoHideDock)              //自动隐藏 Dock
	NSApplicationPresentationHideDock                  = uint(C.NSApplicationPresentationHideDock)                  //永久隐藏 Dock
	NSApplicationPresentationDisableProcessSwitching   = uint(C.NSApplicationPresentationDisableProcessSwitching)   //禁用 Cmd+Tab 切换应用
	NSApplicationPresentationDisableForceQuit          = uint(C.NSApplicationPresentationDisableForceQuit)          //禁用强制退出（Cmd+Opt+Esc）
	NSApplicationPresentationDisableSessionTermination = uint(C.NSApplicationPresentationDisableSessionTermination) //禁用注销/关机提示
	NSApplicationPresentationFullScreen                = uint(C.NSApplicationPresentationFullScreen)
	NSApplicationPresentationAutoHideToolbar           = uint(C.NSApplicationPresentationAutoHideToolbar)
)

func AppSetPresentationOptions(options uint) {
	C.AppSetPresentationOptions(C.NSApplicationPresentationOptions(C.uint(options)))
}

func AppDockHide() {
	C.AppDockHide()
}

func AppDockShow() {
	C.AppDockShow()
}

func AppSetMainMenu(nsMenu unsafe.Pointer) {
	C.AppSetMainMenu(nsMenu)
}
