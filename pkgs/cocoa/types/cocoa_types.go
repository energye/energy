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
