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
	"unsafe"
)

func (m *TWindow) DragWindow() {
	cocoa.DragWindow(m.NSInstance())
}

func (m *TWindow) SetBackgroundColor(red, green, blue, alpha uint8) {
	cocoa.SetWindowBackgroundColor(m.NSInstance(), red, green, blue, alpha)
}

// SetWindowTransparent 设置窗口为透明效果
func (m *TWindow) SetWindowTransparent() {
	frostedView := cocoa.SetWindowTransparent(m.NSWindowInstance())
	m.frostedView = frostedView
}

// SwitchFrostedMaterial 切换窗口的磨砂材质外观
// 该方法会根据指定的外观名称来更改窗口的磨砂视图材质效果
func (m *TWindow) SwitchFrostedMaterial(appearanceName application.AppearanceNamed) {
	cocoa.WindowSwitchFrostedMaterial(m.frostedView, m.NSWindowInstance(), appearanceName)
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

func (m *TWindow) AddSubview(nsView unsafe.Pointer) {
	fmt.Println("AddSubview", nsView)
	//CGRect init = { 0,0,0,0 };
	//[self.webview initWithFrame:init configuration:config];
	//[contentView addSubview:self.webview];
	//[self.webview setAutoresizingMask: NSViewWidthSizable|NSViewHeightSizable];
	//CGRect contentViewBounds = [contentView bounds];
	////contentViewBounds.origin.x = -10.0;
	////contentViewBounds.origin.y = -20.0;
	////contentViewBounds.size.width = 300.0;
	////contentViewBounds.size.height = 400.0;
	//[self.webview setFrame:contentViewBounds];

}
