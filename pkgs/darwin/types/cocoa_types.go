//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

// TWindowEvent CGO 窗口事件, 在 cocoa.h 同步配置
type TWindowEvent = int

const (
	TWindowEventEnterFullScreen                      TWindowEvent = 10000
	TWindowEventExitFullScreen                       TWindowEvent = 10001
	TWindowEventWillUseFullScreenPresentationOptions TWindowEvent = 10002
	TWindowEventDidResize                            TWindowEvent = 10003
)

type NSAutoresizingMaskOptions = int

const (
	NSViewNotSizable    NSAutoresizingMaskOptions = 0x00
	NSViewMinXMargin    NSAutoresizingMaskOptions = 0x01
	NSViewWidthSizable  NSAutoresizingMaskOptions = 0x02
	NSViewMaxXMargin    NSAutoresizingMaskOptions = 0x04
	NSViewMinYMargin    NSAutoresizingMaskOptions = 0x08
	NSViewHeightSizable NSAutoresizingMaskOptions = 0x10
	NSViewMaxYMargin    NSAutoresizingMaskOptions = 0x20
)

type NSVisualEffectBlendingMode = int

const (
	NSVisualEffectBlendingModeBehindWindow NSVisualEffectBlendingMode = 0
	NSVisualEffectBlendingModeWithinWindow NSVisualEffectBlendingMode = 1
)

type NSVisualEffectState = int

const (
	NSVisualEffectStateFollowsWindowActiveState NSVisualEffectState = 0
	NSVisualEffectStateActive                   NSVisualEffectState = 1
	NSVisualEffectStateInactive                 NSVisualEffectState = 2
)

type NSWindowOrderingMode = int

const (
	NSWindowBelow NSWindowOrderingMode = -1
	NSWindowOut   NSWindowOrderingMode = 0
	NSWindowAbove NSWindowOrderingMode = 1
)

type NSApplicationPresentationOptions = int

const (
	NSApplicationPresentationDefault                   NSApplicationPresentationOptions = 0x00000000
	NSApplicationPresentationAutoHideDock              NSApplicationPresentationOptions = 0x00000001
	NSApplicationPresentationHideDock                  NSApplicationPresentationOptions = 0x00000002
	NSApplicationPresentationAutoHideMenuBar           NSApplicationPresentationOptions = 0x00000004
	NSApplicationPresentationHideMenuBar               NSApplicationPresentationOptions = 0x00000008
	NSApplicationPresentationDisableAppleMenu          NSApplicationPresentationOptions = 0x00000010
	NSApplicationPresentationDisableProcessSwitching   NSApplicationPresentationOptions = 0x00000020
	NSApplicationPresentationDisableForceQuit          NSApplicationPresentationOptions = 0x00000040
	NSApplicationPresentationDisableSessionTermination NSApplicationPresentationOptions = 0x00000080
	NSApplicationPresentationFullScreen                NSApplicationPresentationOptions = 0x00000400
	NSApplicationPresentationAutoHideToolbar           NSApplicationPresentationOptions = 0x00000800
)

type NSWindowStyleMask = uint

const (
	NSWindowStyleMaskBorderless          NSWindowStyleMask = 0
	NSWindowStyleMaskTitled              NSWindowStyleMask = 1 << 0
	NSWindowStyleMaskClosable            NSWindowStyleMask = 1 << 1
	NSWindowStyleMaskMiniaturizable      NSWindowStyleMask = 1 << 2
	NSWindowStyleMaskResizable           NSWindowStyleMask = 1 << 3
	NSWindowStyleMaskFullSizeContentView NSWindowStyleMask = 1 << 15
)

type AppearanceName string

const (
	// NSAppearanceNameAqua - 标准浅色系统外观
	NSAppearanceNameAqua AppearanceName = "NSAppearanceNameAqua"
	// NSAppearanceNameDarkAqua - 标准深色系统外观
	NSAppearanceNameDarkAqua AppearanceName = "NSAppearanceNameDarkAqua"
	// NSAppearanceNameVibrantLight - 浅色生动外观
	NSAppearanceNameVibrantLight AppearanceName = "NSAppearanceNameVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastAqua - 标准浅色系统外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastAqua AppearanceName = "NSAppearanceNameAccessibilityHighContrastAqua"
	// NSAppearanceNameAccessibilityHighContrastDarkAqua - 标准深色系统外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastDarkAqua AppearanceName = "NSAppearanceNameAccessibilityHighContrastDarkAqua"
	// NSAppearanceNameAccessibilityHighContrastVibrantLight - 浅色生动外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastVibrantLight AppearanceName = "NSAppearanceNameAccessibilityHighContrastVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastVibrantDark - 深色生动外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastVibrantDark AppearanceName = "NSAppearanceNameAccessibilityHighContrastVibrantDark"
)

type CGRect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// ToolbarConfiguration 工具栏配置选项
type ToolbarConfiguration struct {
	ShowSeparator bool // 是否显示基线分隔符
}
