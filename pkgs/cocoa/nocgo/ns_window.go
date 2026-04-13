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
	. "github.com/energye/energy/v3/pkgs/cocoa/types"
	"github.com/energye/lcl/types"
	"unsafe"
)

type NSWindow struct {
	NSResponder
}

func AsNSWindow(ptr unsafe.Pointer) *NSWindow {
	if ptr == nil {
		return nil
	}
	m := new(NSWindow)
	m.SetInstance(ptr)
	return m
}

func (m *NSWindow) Restore() {
	if m == nil {
		return
	}
	nsWindowID := objc.ID(m.Instance())
	nsWindowID.Send(objc.RegisterName("zoom:"))
}

func (m *NSWindow) Maximize() {
	if m == nil {
		return
	}
	nsWindowID := objc.ID(m.Instance())
	nsWindowID.Send(objc.RegisterName("zoom:"))
}

func (m *NSWindow) Minimized() {
	if m == nil {
		return
	}
	nsWindowID := objc.ID(m.Instance())
	isMiniaturized := nsWindowID.Send(objc.RegisterName("isMiniaturized"))
	if isMiniaturized.Send(objc.RegisterName("boolValue")) == 0 {
		nsWindowID.Send(objc.RegisterName("miniaturize:"))
	}
}

func (m *NSWindow) ExitMinimized() {
	if m == nil {
		return
	}
	nsWindowID := objc.ID(m.Instance())
	isMiniaturized := nsWindowID.Send(objc.RegisterName("isMiniaturized"))
	if isMiniaturized.Send(objc.RegisterName("boolValue")) != 0 {
		nsWindowID.Send(objc.RegisterName("deminiaturize:"))
	}
}

func (m *NSWindow) EnterFullScreen() {
	if m == nil {
		return
	}
	nsWindowID := objc.ID(m.Instance())
	nsWindowID.Send(objc.RegisterName("toggleFullScreen:"))
}

func (m *NSWindow) ExitFullScreen() {
	if m == nil {
		return
	}
	nsWindowID := objc.ID(m.Instance())
	nsWindowID.Send(objc.RegisterName("toggleFullScreen:"))
}

func (m *NSWindow) Drag() {
	if m == nil {
		return
	}
	nsWindowID := objc.ID(m.Instance())
	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))
	currentEvent := nsApp.Send(objc.RegisterName("currentEvent"))
	if currentEvent == 0 {
		return
	}
	eventType := currentEvent.Send(objc.RegisterName("type"))
	const NSEventTypeLeftMouseDown = 1
	if eventType != NSEventTypeLeftMouseDown {
		return
	}
	nsWindowID.Send(objc.RegisterName("performWindowDragWithEvent:"), currentEvent)
}

func (m *NSWindow) SetBackgroundColor(r, g, b, alpha uint8) {
	if m == nil {
		return
	}
	nsWindow := objc.ID(m.Instance())
	nsColorClass := objc.GetClass("NSColor")

	color := objc.ID(nsColorClass).Send(
		objc.RegisterName("colorWithCalibratedRed:green:blue:alpha:"),
		float64(r)/255.0,
		float64(g)/255.0,
		float64(b)/255.0,
		float64(alpha)/255.0,
	)
	nsWindow.Send(objc.RegisterName("setBackgroundColor:"), color)
}

// 设置无边框窗口
func (m *NSWindow) SetRadius(radius float32) {
	if m == nil {
		return
	}
	nsWindow := objc.ID(m.Instance())
	contentView := nsWindow.Send(objc.RegisterName("contentView"))
	if contentView == 0 {
		return
	}
	contentView.Send(objc.RegisterName("setWantsLayer:"), true)
	layer := contentView.Send(objc.RegisterName("layer"))
	if layer == 0 {
		return
	}
	nsColorClass := objc.GetClass("NSColor")
	nsWindow.Send(objc.RegisterName("setBackgroundColor:"),
		objc.ID(nsColorClass).Send(objc.RegisterName("clearColor")))
	layer.Send(objc.RegisterName("setBackgroundColor:"),
		objc.ID(nsColorClass).Send(objc.RegisterName("whiteColor")).Send(objc.RegisterName("CGColor")))
	layer.Send(objc.RegisterName("setCornerRadius:"), float64(radius))
	layer.Send(objc.RegisterName("setMasksToBounds:"), true)
}

// SetTransparent 设置窗口透明（磨砂效果）
// 返回 frostedView
func (m *NSWindow) SetTransparent() *NSVisualEffectView {
	if m == nil {
		return nil
	}

	nsWindow := objc.ID(m.Instance())
	contentView := nsWindow.Send(objc.RegisterName("contentView"))
	if contentView == 0 {
		return nil
	}

	bounds := objc.Send[CGRect](contentView, objc.RegisterName("bounds"))

	frostedView := NewNSVisualEffectView(bounds)
	frostedView.SetAutoresizingMask(NSViewWidthSizable | NSViewHeightSizable)
	frostedView.SetBlendingMode(NSVisualEffectBlendingModeBehindWindow)
	frostedView.SetState(NSVisualEffectStateActive)

	contentView.Send(objc.RegisterName("addSubview:positioned:relativeTo:"),
		frostedView, NSWindowBelow, 0)
	return frostedView
}

// SwitchFrostedMaterial 切换窗口磨砂材质外观
func (m *NSWindow) SwitchFrostedMaterial(frostedView *NSVisualEffectView, appearanceName string) {
	if m == nil || frostedView == nil || frostedView.Instance() == 0 {
		return
	}
	nsWindow := objc.ID(m.Instance())
	nsAppearanceClass := objc.GetClass("NSAppearance")
	// 创建 NSAppearance 对象
	appearance := objc.ID(nsAppearanceClass).Send(objc.RegisterName("appearanceNamed:"), appearanceName)
	if appearance != 0 {
		// 设置窗口外观
		nsWindow.Send(objc.RegisterName("setAppearance:"), appearance)
	}
}

// AddSubview 添加子视图到窗口
func (m *NSWindow) AddSubview(view *NSView, x, y, width, height float32) {
	if m == nil || view == nil {
		return
	}

	nsWindow := objc.ID(m.Instance())
	contentView := nsWindow.Send(objc.RegisterName("contentView"))
	if contentView == 0 {
		return
	}

	nsView := view.Self()

	// 添加到 contentView
	contentView.Send(objc.RegisterName("addSubview:"), nsView)

	// 设置自动调整大小
	const NSViewWidthSizable = 2
	const NSViewHeightSizable = 16
	nsView.Send(objc.RegisterName("setAutoresizingMask:"),
		uintptr(NSViewWidthSizable|NSViewHeightSizable))

	// 设置 frame
	frame := CGRect{
		X:      float64(x),
		Y:      float64(y),
		Width:  float64(width),
		Height: float64(height),
	}

	nsView.Send(objc.RegisterName("setFrame:"), frame)
}

// ContentViewFrame 获取窗口内容视图的帧矩形
func (m *NSWindow) ContentViewFrame() (rect types.TRect) {
	if m == nil {
		return
	}
	nsWindow := objc.ID(m.Instance())
	contentView := nsWindow.Send(objc.RegisterName("contentView"))
	if contentView == 0 {
		return
	}
	originalFrame := objc.Send[CGRect](contentView, objc.RegisterName("frame"))
	// 获取 superview 的高度（用于坐标转换）
	superview := contentView.Send(objc.RegisterName("superview"))
	if superview == 0 {
		rect.Left = int32(originalFrame.X)
		rect.Top = int32(originalFrame.Y)
		rect.Right = int32(originalFrame.X + originalFrame.Width)
		rect.Bottom = int32(originalFrame.Y + originalFrame.Height)
		return rect
	}
	superviewFrame := objc.Send[CGRect](superview, objc.RegisterName("frame"))
	superViewHeight := superviewFrame.Height
	// 转换坐标系（macOS 原点在左下角，需要转换为左上角）
	topLeftX := originalFrame.X
	topLeftY := superViewHeight - originalFrame.Y - originalFrame.Height
	rect.Left = int32(topLeftX)
	rect.Top = int32(topLeftY)
	rect.Right = int32(topLeftX + originalFrame.Width)
	rect.Bottom = int32(topLeftY + originalFrame.Height)
	return rect
}

func (m *NSWindow) doWindowDidResie() {

}

func (m *NSWindow) doWindowWillEnterFullScreen() {

}

func (m *NSWindow) doWindowDidExitFullScreen() {

}

func (m *NSWindow) doWindowWillUseFullScreenPresentationOptions(options NSApplicationPresentationOptions) NSApplicationPresentationOptions {
	return options
}
