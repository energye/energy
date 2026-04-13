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

package cgo

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

int AppGetActivationPolicy() {
    return (int)[[NSApplication sharedApplication] activationPolicy];
}

NSApplicationPresentationOptions AppGetPresentationOptions() {
    return [[NSApplication sharedApplication] presentationOptions];
}

void AppActivate() {
    [[NSApplication sharedApplication] activateIgnoringOtherApps:YES];
}

void AppDeactivate() {
    [[NSApplication sharedApplication] deactivate];
}

void AppHide() {
    [[NSApplication sharedApplication] hide:nil];
}

void AppUnhide() {
    [[NSApplication sharedApplication] unhide:nil];
}

void AppTerminate() {
    [[NSApplication sharedApplication] terminate:nil];
}


*/
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/darwin/types"
	"unsafe"
)

type NSApp struct{}

func AsNSApp() INSApp {
	return &NSApp{}
}
func (m *NSApp) AppSetPresentationOptions(options NSApplicationPresentationOptions) {
	C.AppSetPresentationOptions(C.NSApplicationPresentationOptions(C.uint(options)))
}

func (m *NSApp) AppDockHide() {
	C.AppDockHide()
}

func (m *NSApp) AppDockShow() {
	C.AppDockShow()
}

func (m *NSApp) AppSetMainMenu(nsMenu unsafe.Pointer) {
	C.AppSetMainMenu(nsMenu)
}

// AppGetActivationPolicy 获取当前应用的激活策略
//
//   - 0: Regular（正常应用，显示 Dock 图标）
//   - 1: Accessory（辅助应用，不显示在 Dock 但可激活）
//   - 2: Prohibited（禁止激活，不显示 Dock 图标）
func (m *NSApp) AppGetActivationPolicy() int {
	return int(C.AppGetActivationPolicy())
}

// AppGetPresentationOptions 获取当前的全屏展示选项
//
// 返回值: 当前的展示选项位掩码
func (m *NSApp) AppGetPresentationOptions() NSApplicationPresentationOptions {
	return NSApplicationPresentationOptions(C.AppGetPresentationOptions())
}

// AppActivate 激活应用并带到前台
func (m *NSApp) AppActivate() {
	C.AppActivate()
}

// AppDeactivate 取消激活应用
func (m *NSApp) AppDeactivate() {
	C.AppDeactivate()
}

// AppHide 隐藏应用（等同于 Cmd+H）
func (m *NSApp) AppHide() {
	C.AppHide()
}

// AppUnHide 取消隐藏应用
func (m *NSApp) AppUnHide() {
	C.AppUnhide()
}

// AppTerminate 终止应用
func (m *NSApp) AppTerminate() {
	C.AppTerminate()
}
