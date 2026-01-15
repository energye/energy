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

package toolbar

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "config.h"
*/
import "C"

// TccType 事件类型, 用于区分普通通知事件, 还是特殊事件
type TccType = C.TccType

const (
	TCCNotify             = TccType(C.TCCNotify)
	TCCClicked            = TccType(C.TCCClicked)
	TCCTextDidChange      = TccType(C.TCCTextDidChange)
	TCCTextDidEndEditing  = TccType(C.TCCTextDidEndEditing)
	TCCSelectionChanged   = TccType(C.TCCSelectionChanged)
	TCCSelectionDidChange = TccType(C.TCCSelectionDidChange)
)

type ToolbarDisplayMode = int

const (
	NSToolbarDisplayModeDefault      ToolbarDisplayMode = 0
	NSToolbarDisplayModeIconAndLabel ToolbarDisplayMode = 1
	NSToolbarDisplayModeIconOnly     ToolbarDisplayMode = 2
	NSToolbarDisplayModeLabelOnly    ToolbarDisplayMode = 3
)

type ToolbarStyle = int

const (
	NSWindowToolbarStyleAutomatic      ToolbarStyle = 0
	NSWindowToolbarStyleExpanded       ToolbarStyle = 1
	NSWindowToolbarStylePreference     ToolbarStyle = 2
	NSWindowToolbarStyleUnified        ToolbarStyle = 3
	NSWindowToolbarStyleUnifiedCompact ToolbarStyle = 4
)

// 边框样式
const (
	BezelStyleRounded           NSBezelStyle = C.NSBezelStyleRounded
	BezelStyleRegularSquare     NSBezelStyle = C.NSBezelStyleRegularSquare
	BezelStyleDisclosure        NSBezelStyle = C.NSBezelStyleDisclosure
	BezelStyleShadowlessSquare  NSBezelStyle = C.NSBezelStyleShadowlessSquare
	BezelStyleCircular          NSBezelStyle = C.NSBezelStyleCircular
	BezelStyleTexturedSquare    NSBezelStyle = C.NSBezelStyleTexturedSquare
	BezelStyleHelpButton        NSBezelStyle = C.NSBezelStyleHelpButton
	BezelStyleSmallSquare       NSBezelStyle = C.NSBezelStyleSmallSquare
	BezelStyleTexturedRounded   NSBezelStyle = C.NSBezelStyleTexturedRounded
	BezelStyleRoundRect         NSBezelStyle = C.NSBezelStyleRoundRect
	BezelStyleRecessed          NSBezelStyle = C.NSBezelStyleRecessed
	BezelStyleRoundedDisclosure NSBezelStyle = C.NSBezelStyleRoundedDisclosure
	BezelStyleInline            NSBezelStyle = C.NSBezelStyleInline
)

// 控件大小
const (
	ControlSizeMini    NSControlSize = C.NSControlSizeMini
	ControlSizeSmall   NSControlSize = C.NSControlSizeSmall
	ControlSizeRegular NSControlSize = C.NSControlSizeRegular
	ControlSizeLarge   NSControlSize = C.NSControlSizeLarge
)

type ItemVisibilityPriority = int

const (
	NSToolbarItemVisibilityPriorityStandard ItemVisibilityPriority = 0
	NSToolbarItemVisibilityPriorityLow      ItemVisibilityPriority = -1000
	NSToolbarItemVisibilityPriorityHigh     ItemVisibilityPriority = 1000
	NSToolbarItemVisibilityPriorityUser     ItemVisibilityPriority = 2000
)

type TitlebarSeparatorStyle = int

const (
	NSTitlebarSeparatorStyleAutomatic TitlebarSeparatorStyle = 0
	NSTitlebarSeparatorStyleNone      TitlebarSeparatorStyle = 1
	NSTitlebarSeparatorStyleLine      TitlebarSeparatorStyle = 2
	NSTitlebarSeparatorStyleShadow    TitlebarSeparatorStyle = 3
)

type ToolbarSizeMode = int

const (
	NSToolbarSizeModeDefault ToolbarSizeMode = 0
	NSToolbarSizeModeRegular ToolbarSizeMode = 1
	NSToolbarSizeModeSmall   ToolbarSizeMode = 2
)

type NSBezelStyle = int

const (
	NSBezelStyleRegularSquare    NSBezelStyle = NSBezelStyle(C.NSBezelStyleRegularSquare)
	NSBezelStyleSmallSquare      NSBezelStyle = NSBezelStyle(C.NSBezelStyleSmallSquare)
	NSBezelStyleDisclosure       NSBezelStyle = NSBezelStyle(C.NSBezelStyleDisclosure)
	NSBezelStyleShadowlessSquare NSBezelStyle = NSBezelStyle(C.NSBezelStyleShadowlessSquare)
	NSBezelStyleRoundRect        NSBezelStyle = NSBezelStyle(C.NSBezelStyleRoundRect)
	NSBezelStyleTexturedSquare   NSBezelStyle = NSBezelStyle(C.NSBezelStyleTexturedSquare)
	NSBezelStyleTexturedRounded  NSBezelStyle = NSBezelStyle(C.NSBezelStyleTexturedRounded)
	NSBezelStyleHelpButton       NSBezelStyle = NSBezelStyle(C.NSBezelStyleHelpButton)
	NSBezelStyleInline           NSBezelStyle = NSBezelStyle(C.NSBezelStyleInline)
)

type NSControlSize = int

const (
	NSControlSizeMini    NSControlSize = NSControlSize(C.NSControlSizeMini)
	NSControlSizeSmall   NSControlSize = NSControlSize(C.NSControlSizeSmall)
	NSControlSizeRegular NSControlSize = NSControlSize(C.NSControlSizeRegular)
	NSControlSizeLarge   NSControlSize = NSControlSize(C.NSControlSizeLarge)
)
