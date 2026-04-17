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
	"encoding/json"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

type NSApp struct {
	initializationAppDelegate bool
	onOpenURLs                TOpenURLsEvent
	onUniversalLink           TUniversalLinkEvent
}

func AsNSApp() INSApp {
	m := &NSApp{}
	m.registerEvents()
	return m
}

func (m *NSApp) registerEvents() {
	println("[DEBUG] registerEvents")
	RegisterEvent("openURLs", &TCallback{cb: func(ctx *TCallbackContext) *GoArguments {
		if ctx.Arguments != nil {
			urls := ctx.Arguments.GetString(0)
			m.doOpenURLs(urls)
		}
		return nil
	}})
	RegisterEvent("universalLink", &TCallback{cb: func(ctx *TCallbackContext) *GoArguments {
		if ctx.Arguments != nil {
			universalLink := ctx.Arguments.GetString(0)
			m.doUniversalLink(universalLink)
		}
		return nil
	}})
}

// SetPresentationOptions 设置应用的全屏展示选项
// 用于控制全屏时的 UI 元素显示行为（如菜单栏、Dock 等）
//
//   - options: 展示选项位掩码，可以是多个选项的组合
//     例如: NSApplicationPresentationAutoHideMenuBar | NSApplicationPresentationFullScreen
func (m *NSApp) SetPresentationOptions(options NSApplicationPresentationOptions) {
	C.AppSetPresentationOptions(C.NSApplicationPresentationOptions(C.uint(options)))
}

// DockHide 隐藏 Dock 图标
// 设置应用激活策略为 Prohibited，适用于后台应用或菜单栏应用
func (m *NSApp) DockHide() {
	C.AppDockHide()
}

// DockShow 显示 Dock 图标
// 设置应用激活策略为 Regular，恢复正常的应用行为
func (m *NSApp) DockShow() {
	C.AppDockShow()
}

// SetMainMenu 设置应用的主菜单
//
//   - nsMenu: NSMenu 对象的指针（unsafe.Pointer）
func (m *NSApp) SetMainMenu(nsMenu unsafe.Pointer) {
	C.AppSetMainMenu(nsMenu)
}

// GetActivationPolicy 获取当前应用的激活策略
//
//   - 0: Regular（正常应用，显示 Dock 图标）
//   - 1: Accessory（辅助应用，不显示在 Dock 但可激活）
//   - 2: Prohibited（禁止激活，不显示 Dock 图标）
func (m *NSApp) GetActivationPolicy() int {
	return int(C.AppGetActivationPolicy())
}

// GetPresentationOptions 获取当前的全屏展示选项
//
// 返回值: 当前的展示选项位掩码
func (m *NSApp) GetPresentationOptions() NSApplicationPresentationOptions {
	return NSApplicationPresentationOptions(C.AppGetPresentationOptions())
}

// Activate 激活应用并带到前台
func (m *NSApp) Activate() {
	C.AppActivate()
}

// Deactivate 取消激活应用
func (m *NSApp) Deactivate() {
	C.AppDeactivate()
}

// Hide 隐藏应用（等同于 Cmd+H）
func (m *NSApp) Hide() {
	C.AppHide()
}

// UnHide 取消隐藏应用
func (m *NSApp) UnHide() {
	C.AppUnhide()
}

// Terminate 终止应用
func (m *NSApp) Terminate() {
	C.AppTerminate()
}

func (m *NSApp) doOpenURLs(urls string) {
	if m.onOpenURLs != nil {
		var items []string
		_ = json.Unmarshal([]byte(urls), &items)
		m.onOpenURLs(items)
	}
}

func (m *NSApp) doUniversalLink(universalLink string) {
	if m.onUniversalLink != nil {
		m.onUniversalLink(universalLink)
	}
}
