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

func ToolbarConfigurationToOC(config ToolbarConfiguration) C.ToolbarConfiguration {
	cConfig := C.ToolbarConfiguration{
		ShowSeparator: _BoolToCInt(config.ShowSeparator),
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

func _BoolToCInt(value bool) C.int {
	if value {
		return C.int(1)
	}
	return C.int(0)
}
