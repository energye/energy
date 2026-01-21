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

#import "Cocoa/Cocoa.h"
#import <WebKit/WebKit.h>
*/
import "C"
import (
	"fmt"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/pkgs/cocoa"
	"github.com/energye/lcl/types"
)

func (m *TWindow) DragWindow() {
	cocoa.DragWindow(m.NSInstance())
}

func (m *TWindow) SetBackgroundColor(red, green, blue, alpha uint8) {
	cocoa.SetWindowBackgroundColor(m.NSInstance(), red, green, blue, alpha)
}

// SetWindowTransparent 设置窗口为透明效果
func (m *TWindow) SetWindowTransparent() {
	m.nsFrostedView = cocoa.SetWindowTransparent(m.NSWindowInstance())
}

// SwitchFrostedMaterial 切换窗口的磨砂材质外观
// 该方法会根据指定的外观名称来更改窗口的磨砂视图材质效果
func (m *TWindow) SwitchFrostedMaterial(appearanceName application.AppearanceNamed) {
	cocoa.WindowSwitchFrostedMaterial(m.nsFrostedView, m.NSWindowInstance(), appearanceName)
}

// SetWindowRadius 设置窗口圆角半径
// 该函数仅在无边框模式下生效，根据应用选项中的MacOS配置设置窗口圆角
func (m *TWindow) SetWindowRadius() {
	options := application.GApplication.Options
	if options.Frameless {
		if options.MacOS.WindowRadius > 0.0 {
			cocoa.SetWindowRadius(m.NSInstance(), options.MacOS.WindowRadius)
		}
	}
}

// TitleBar 配置窗口标题栏的外观和行为
// 该方法根据应用程序选项设置窗口样式掩码、标题栏透明度、可见性等属性
func (m *TWindow) TitleBar() {
	options := application.GApplication.Options
	nsWindow := m.NSWindow()
	mask := nsWindow.StyleMask()
	if options.DisableSystemMenu {
		mask ^= C.NSWindowStyleMaskClosable
	}
	if options.DisableMinimize {
		mask ^= C.NSWindowStyleMaskMiniaturizable
	}
	if options.DisableResize || options.DisableMaximize {
		mask ^= C.NSWindowStyleMaskResizable
	}
	if options.MacOS.FullSizeContent {
		mask |= C.NSWindowStyleMaskFullSizeContentView
	}
	nsWindow.SetStyleMask(mask)
	nsWindow.SetTitleBarAppearsTransparent(options.MacOS.TitleTransparent)
	if options.MacOS.TitleHideText {
		nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
	}
	toolBar := options.MacOS.ToolBar
	if toolBar != nil {
		cocoa.NewToolBar(m.NSInstance(), cocoa.ToolbarConfiguration{ShowSeparator: toolBar.ShowSeparator})
	}
}

// Frameless 设置窗口为无边框模式
func (m *TWindow) Frameless() {
	options := application.GApplication.Options
	if options.Frameless {
		nsWindow := m.NSWindow()
		mask := nsWindow.StyleMask()
		mask ^= C.NSWindowStyleMaskTitled
		nsWindow.SetStyleMask(mask)
	}
}

func (m *TWindow) _InitEvent() {
	nsWindow := m.NSInstance()
	baseEventID := fmt.Sprintf("%v", nsWindow)
	EnterFullScreen := fmt.Sprintf("%d_%v", cocoa.TWindowEventEnterFullScreen, baseEventID)
	cocoa.RegisterEvent(EnterFullScreen, cocoa.MakeNotifyEvent(func(identifier string, owner cocoa.Pointer, sender cocoa.Pointer) *cocoa.GoArguments {
		fmt.Println("EnterFullScreen", m.BrowserId(), sender)
		return nil
	}))
	ExitFullScreen := fmt.Sprintf("%d_%v", cocoa.TWindowEventExitFullScreen, baseEventID)
	cocoa.RegisterEvent(ExitFullScreen, cocoa.MakeNotifyEvent(func(identifier string, owner cocoa.Pointer, sender cocoa.Pointer) *cocoa.GoArguments {
		fmt.Println("ExitFullScreen", m.BrowserId(), sender)
		return nil
	}))
	UseFullScreenPresentationOptions := fmt.Sprintf("%d_%v", cocoa.TWindowEventWillUseFullScreenPresentationOptions, baseEventID)
	cocoa.RegisterEvent(UseFullScreenPresentationOptions, cocoa.MakeNotifyEvent(func(identifier string, owner cocoa.Pointer, sender cocoa.Pointer) *cocoa.GoArguments {
		fmt.Println("UseFullScreenPresentationOptions", m.BrowserId(), sender)
		return nil
	}))
}
