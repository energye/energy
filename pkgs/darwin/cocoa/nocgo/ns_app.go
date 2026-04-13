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
	. "github.com/energye/energy/v3/pkgs/darwin/types"
	"unsafe"
)

type NSApp struct{}

func AsNSApp() INSApp {
	return &NSApp{}
}

// AppDockHide 隐藏 Dock 图标
// 设置应用激活策略为 Prohibited，适用于后台应用或菜单栏应用
func (m *NSApp) AppDockHide() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("setActivationPolicy:"), uintptr(NSApplicationActivationPolicyProhibited))
}

// AppDockShow 显示 Dock 图标
// 设置应用激活策略为 Regular，恢复正常的应用行为
func (m *NSApp) AppDockShow() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("setActivationPolicy:"), uintptr(NSApplicationActivationPolicyRegular))
}

// AppSetPresentationOptions 设置应用的全屏展示选项
// 用于控制全屏时的 UI 元素显示行为（如菜单栏、Dock 等）
//
//   - options: 展示选项位掩码，可以是多个选项的组合
//     例如: NSApplicationPresentationAutoHideMenuBar | NSApplicationPresentationFullScreen
func (m *NSApp) AppSetPresentationOptions(options NSApplicationPresentationOptions) {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("setPresentationOptions:"), uintptr(options))
}

// AppSetMainMenu 设置应用的主菜单
//
//   - nsMenu: NSMenu 对象的指针（unsafe.Pointer）
func (m *NSApp) AppSetMainMenu(nsMenu unsafe.Pointer) {
	if nsMenu == nil || uintptr(nsMenu) == 0 {
		return
	}
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))

	menuID := objc.ID(uintptr(nsMenu))
	nsApp.Send(objc.RegisterName("setMainMenu:"), menuID)
}

// AppGetActivationPolicy 获取当前应用的激活策略
//
//   - 0: Regular（正常应用，显示 Dock 图标）
//   - 1: Accessory（辅助应用，不显示在 Dock 但可激活）
//   - 2: Prohibited（禁止激活，不显示 Dock 图标）
func (m *NSApp) AppGetActivationPolicy() int {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	return int(objc.Send[uintptr](nsApp, objc.RegisterName("activationPolicy")))
}

// AppGetPresentationOptions 获取当前的全屏展示选项
//
//	当前的展示选项位掩码
func (m *NSApp) AppGetPresentationOptions() NSApplicationPresentationOptions {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	options := objc.Send[uintptr](nsApp, objc.RegisterName("presentationOptions"))
	return NSApplicationPresentationOptions(options)
}

// AppActivate 激活应用并带到前台
func (m *NSApp) AppActivate() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)
}

// AppDeactivate 取消激活应用
func (m *NSApp) AppDeactivate() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("deactivate"))
}

// AppHide 隐藏应用
func (m *NSApp) AppHide() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))

	nsApp.Send(objc.RegisterName("hide:"), nil)
}

// AppUnHide 取消隐藏应用
func (m *NSApp) AppUnHide() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("unhide:"), nil)
}

// AppTerminate 终止应用
func (m *NSApp) AppTerminate() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("terminate:"), nil)
}
