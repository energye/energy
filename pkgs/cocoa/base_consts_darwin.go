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

package cocoa

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "cocoa.h"
*/
import "C"

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
	NSBezelStyleRegularSquare    = NSBezelStyle(C.NSBezelStyleRegularSquare)
	NSBezelStyleSmallSquare      = NSBezelStyle(C.NSBezelStyleSmallSquare)
	NSBezelStyleDisclosure       = NSBezelStyle(C.NSBezelStyleDisclosure)
	NSBezelStyleShadowlessSquare = NSBezelStyle(C.NSBezelStyleShadowlessSquare)
	NSBezelStyleRoundRect        = NSBezelStyle(C.NSBezelStyleRoundRect)
	NSBezelStyleTexturedSquare   = NSBezelStyle(C.NSBezelStyleTexturedSquare)
	NSBezelStyleTexturedRounded  = NSBezelStyle(C.NSBezelStyleTexturedRounded)
	NSBezelStyleHelpButton       = NSBezelStyle(C.NSBezelStyleHelpButton)
	NSBezelStyleInline           = NSBezelStyle(C.NSBezelStyleInline)
)

type NSControlSize = int

const (
	NSControlSizeMini    = NSControlSize(2) // NSControlSize(C.NSControlSizeMini)
	NSControlSizeSmall   = NSControlSize(1) // NSControlSize(C.NSControlSizeSmall)
	NSControlSizeRegular = NSControlSize(0) // NSControlSize(C.NSControlSizeRegular)
	NSControlSizeLarge   = NSControlSize(3) // NSControlSize(C.NSControlSizeLarge)
)
