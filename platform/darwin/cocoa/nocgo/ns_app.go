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
	"unsafe"
)

type NSApp struct{}

func AsNSApp() INSApp {
	return &NSApp{}
}

// DockHide 隐藏 Dock 图标
// 设置应用激活策略为 Prohibited，适用于后台应用或菜单栏应用
func (m *NSApp) DockHide() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("setActivationPolicy:"), uintptr(NSApplicationActivationPolicyProhibited))
}

// DockShow 显示 Dock 图标
// 设置应用激活策略为 Regular，恢复正常的应用行为
func (m *NSApp) DockShow() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("setActivationPolicy:"), uintptr(NSApplicationActivationPolicyRegular))
}

// SetPresentationOptions 设置应用的全屏展示选项
// 用于控制全屏时的 UI 元素显示行为（如菜单栏、Dock 等）
//
//   - options: 展示选项位掩码，可以是多个选项的组合
//     例如: NSApplicationPresentationAutoHideMenuBar | NSApplicationPresentationFullScreen
func (m *NSApp) SetPresentationOptions(options NSApplicationPresentationOptions) {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("setPresentationOptions:"), uintptr(options))
}

// SetMainMenu 设置应用的主菜单
//
//   - nsMenu: NSMenu 对象的指针（unsafe.Pointer）
func (m *NSApp) SetMainMenu(nsMenu unsafe.Pointer) {
	if nsMenu == nil || uintptr(nsMenu) == 0 {
		return
	}
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))

	menuID := objc.ID(uintptr(nsMenu))
	nsApp.Send(objc.RegisterName("setMainMenu:"), menuID)
}

// GetActivationPolicy 获取当前应用的激活策略
//
//   - 0: Regular（正常应用，显示 Dock 图标）
//   - 1: Accessory（辅助应用，不显示在 Dock 但可激活）
//   - 2: Prohibited（禁止激活，不显示 Dock 图标）
func (m *NSApp) GetActivationPolicy() int {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	return int(objc.Send[uintptr](nsApp, objc.RegisterName("activationPolicy")))
}

// GetPresentationOptions 获取当前的全屏展示选项
//
//	当前的展示选项位掩码
func (m *NSApp) GetPresentationOptions() NSApplicationPresentationOptions {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	options := objc.Send[uintptr](nsApp, objc.RegisterName("presentationOptions"))
	return NSApplicationPresentationOptions(options)
}

// Activate 激活应用并带到前台
func (m *NSApp) Activate() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("activateIgnoringOtherApps:"), true)
}

// Deactivate 取消激活应用
func (m *NSApp) Deactivate() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("deactivate"))
}

// Hide 隐藏应用
func (m *NSApp) Hide() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))

	nsApp.Send(objc.RegisterName("hide:"), nil)
}

// UnHide 取消隐藏应用
func (m *NSApp) UnHide() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("unhide:"), nil)
}

// Terminate 终止应用
func (m *NSApp) Terminate() {
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	nsApp.Send(objc.RegisterName("terminate:"), nil)
}
