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
import (
	"strconv"
	"unsafe"
)

type Pointer = unsafe.Pointer

// NotifyEvent 通用事件通知
type NotifyEvent func(identifier string, owner Pointer, sender Pointer) *GoArguments
type TextEvent func(identifier string, value string, owner Pointer, sender Pointer) *GoArguments
type DelegateToolbarEvent func(arguments *OCGoArguments, owner Pointer, sender Pointer) *GoArguments

type Color struct {
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

func (m *Color) ToOC() C.Color {
	return C.Color{Red: C.CGFloat(m.Red / 255.0), Green: C.CGFloat(m.Green / 255.0), Blue: C.CGFloat(m.Blue / 255.0), Alpha: C.CGFloat(m.Alpha / 255.0)}
}

type ItemBase struct {
	Identifier   string
	Priority     ItemVisibilityPriority
	Navigational bool
}

type ItemUI struct {
	ItemBase
	IconName string
	Title    string
	Tips     string
	Bordered bool
}

type ButtonItem struct {
	ItemUI
}

type ControlTextField struct {
	ItemUI
	SendWhole         bool
	SendImmediately   bool
	ResignsWithCancel bool
	PreferredWidth    float32
	Placeholder       string
}

type ControlComboBox struct {
	ItemUI
	Editable bool
	Items    []string
}

// ControlProperty 的Go包装
type ControlProperty struct {
	Width              float64
	Height             float64
	MinWidth           float64
	MaxWidth           float64
	BezelStyle         NSBezelStyle
	ControlSize        NSControlSize
	Font               Pointer
	IsNavigational     bool
	IsCenteredItem     bool
	VisibilityPriority ItemVisibilityPriority
}

func (m *ControlProperty) ToOC() C.ControlProperty {
	cProperty := C.ControlProperty{
		width:              C.CGFloat(m.Width),
		height:             C.CGFloat(m.Height),
		minWidth:           C.CGFloat(m.MinWidth),
		maxWidth:           C.CGFloat(m.MaxWidth),
		bezelStyle:         C.NSBezelStyle(m.BezelStyle),
		controlSize:        C.NSControlSize(m.ControlSize),
		font:               (*C.NSFont)(m.Font),
		IsNavigational:     C.BOOL(m.IsNavigational),
		IsCenteredItem:     C.BOOL(m.IsCenteredItem),
		VisibilityPriority: C.NSInteger(m.VisibilityPriority),
	}
	return cProperty
}

func (m *ControlProperty) ToOCMalloc() *C.ControlProperty {
	cPropertyPtr := (*C.ControlProperty)(C.malloc(C.sizeof_ControlProperty))
	if cPropertyPtr == nil {
		panic("malloc failed for C.ControlProperty") // 处理内存分配失败
	}
	cPropertyPtr.width = C.CGFloat(m.Width)
	cPropertyPtr.height = C.CGFloat(m.Height)
	cPropertyPtr.minWidth = C.CGFloat(m.MinWidth)
	cPropertyPtr.maxWidth = C.CGFloat(m.MaxWidth)
	cPropertyPtr.bezelStyle = C.NSBezelStyle(m.BezelStyle)
	cPropertyPtr.controlSize = C.NSControlSize(m.ControlSize)
	cPropertyPtr.font = (*C.NSFont)(m.Font) // 注意：m.Font需确保是C侧合法指针（如ObjC传过来的NSFont*）
	cPropertyPtr.IsNavigational = C.BOOL(m.IsNavigational)
	cPropertyPtr.IsCenteredItem = C.BOOL(m.IsCenteredItem)
	cPropertyPtr.VisibilityPriority = C.NSInteger(m.VisibilityPriority)
	return cPropertyPtr
}

type ToolbarCallbackContext struct {
	Type       TccType        // 事件类型
	Identifier string         // 控件唯一标识
	Value      string         // 控件值
	Index      int            // 值索引
	Owner      Pointer        // 所属对象
	Sender     Pointer        // 控件对象
	Arguments  *OCGoArguments // 参数
}

// ToolbarConfiguration 的Go包装
type ToolbarConfiguration struct {
	IsAllowsUserCustomization bool
	IsAutoSavesConfiguration  bool
	Transparent               bool
	ShowsToolbarButton        bool // 隐藏工具栏默认的"显示/隐藏"按钮（右侧）
	SeparatorStyle            TitlebarSeparatorStyle
	DisplayMode               ToolbarDisplayMode
	SizeMode                  ToolbarSizeMode
	Style                     ToolbarStyle
}

func ToolbarConfigurationToOC(config ToolbarConfiguration) C.ToolbarConfiguration {
	cConfig := C.ToolbarConfiguration{
		IsAllowsUserCustomization: C.BOOL(config.IsAllowsUserCustomization),
		IsAutoSavesConfiguration:  C.BOOL(config.IsAutoSavesConfiguration),
		Transparent:               C.BOOL(config.Transparent),
		ShowsToolbarButton:        C.BOOL(config.ShowsToolbarButton),
		SeparatorStyle:            C.NSUInteger(config.SeparatorStyle),
		DisplayMode:               C.NSUInteger(config.DisplayMode),
		SizeMode:                  C.NSUInteger(config.SizeMode),
		Style:                     C.NSUInteger(config.Style),
	}
	return cConfig
}

var serialNumber = make(map[string]int)

func nextSerialNumber(type_ string) string {
	var r int
	if sn, ok := serialNumber[type_]; ok {
		r = sn
	} else {
		r = 0
	}
	r++
	serialNumber[type_] = r
	return type_ + strconv.Itoa(r)
}
