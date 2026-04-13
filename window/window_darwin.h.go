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

import (
	"github.com/energye/energy/v3/pkgs/darwin/cocoa"
	. "github.com/energye/energy/v3/pkgs/darwin/types"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

func (m *TWindow) SetBackgroundColor(red, green, blue, alpha uint8) {
	m.nsWindow.SetBackgroundColor(red, green, blue, alpha)
}

// SetWindowTransparent 设置窗口为透明效果
func (m *TWindow) SetWindowTransparent() {
	m.nsFrostedView = m.nsWindow.SetTransparent()
}

// SwitchFrostedMaterial 切换窗口的磨砂材质外观
// 该方法会根据指定的外观名称来更改窗口的磨砂视图材质效果
func (m *TWindow) SwitchFrostedMaterial(appearanceName AppearanceName) {
	m.nsWindow.SwitchFrostedMaterial(string(appearanceName))
}

// SetWindowRadius 设置窗口圆角半径
// 该函数仅在无边框模式下生效，根据应用选项中的MacOS配置设置窗口圆角
func (m *TWindow) SetWindowRadius() {
	if m.options != nil {
		if m.options.Frameless {
			if m.options.MacOS.WindowRadius > 0.0 {
				m.nsWindow.SetRadius(m.options.MacOS.WindowRadius)
			}
		}
	}
}

// TitleBar 配置窗口标题栏的外观和行为
// 该方法根据应用程序选项设置窗口样式掩码、标题栏透明度、可见性等属性
func (m *TWindow) TitleBar() {
	if m.options != nil {
		nsWindow := lcl.PlatformWindow(m.Instance())
		mask := nsWindow.StyleMask()
		if m.options.DisableSystemMenu {
			mask ^= NSWindowStyleMaskClosable
		}
		if m.options.DisableMinimize {
			mask ^= NSWindowStyleMaskMiniaturizable
		}
		if m.options.DisableResize || m.options.DisableMaximize {
			mask ^= NSWindowStyleMaskResizable
		}
		if m.options.MacOS.FullSizeContent {
			mask |= NSWindowStyleMaskFullSizeContentView
		}
		nsWindow.SetStyleMask(mask)
		nsWindow.SetTitleBarAppearsTransparent(m.options.MacOS.TitleTransparent)
		if m.options.MacOS.TitleHideText {
			nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
		}
		toolBar := m.options.MacOS.ToolBar
		if toolBar != nil {
			cocoa.NewToolBar(m.nsWindow, m.nsDelegate, ToolbarConfiguration{ShowSeparator: toolBar.ShowSeparator})
		}
	}
}

// Frameless 设置窗口为无边框模式
func (m *TWindow) Frameless() {
	if m.options != nil {
		if m.options.Frameless {
			nsWindow := lcl.PlatformWindow(m.Instance())
			mask := nsWindow.StyleMask()
			mask ^= NSWindowStyleMaskTitled
			nsWindow.SetStyleMask(mask)
		}
	}
}

func (m *TWindow) _InitEvent() {
	//nsWindow := m.NSWindow().Instance()
	//windowResizeEventId := fmt.Sprintf("%d_%v", TWindowEventDidResize, nsWindow)
	//cgo.RegisterEvent(windowResizeEventId,
	//	cgo.MakeNotifyEvent(func(identifier string, owner cgo.Pointer, sender cgo.Pointer) *cgo.GoArguments {
	//		return nil
	//	}))
}
