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

package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "cocoa.h"
*/
import "C"
import (
	. "github.com/energye/energy/v3/platform/darwin/types"
	"strconv"
	"unsafe"
)

type Pointer = unsafe.Pointer

type TNotifyEvent func(identifier string, owner Pointer, sender Pointer) *GoArguments
type TTextEvent func(identifier string, value string, owner Pointer, sender Pointer) *GoArguments
type TDelegateEvent func(arguments *OCGoArguments, owner Pointer, sender Pointer) *GoArguments

type TColor struct {
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

func (m *TColor) ToOC() C.Color {
	return C.Color{Red: C.CGFloat(m.Red / 255.0), Green: C.CGFloat(m.Green / 255.0), Blue: C.CGFloat(m.Blue / 255.0), Alpha: C.CGFloat(m.Alpha / 255.0)}
}

type TItemBase struct {
	Identifier   string
	Priority     ItemVisibilityPriority
	Navigational bool
}

type TItemUI struct {
	TItemBase
	IconName string
	Title    string
	Tips     string
	Bordered bool
}

type TButtonItem struct {
	TItemUI
}

type TControlTextField struct {
	TItemUI
	SendWhole         bool
	SendImmediately   bool
	ResignsWithCancel bool
	PreferredWidth    float32
	Placeholder       string
}

type TControlComboBox struct {
	TItemUI
	Editable bool
	Items    []string
}

type TCallbackContext struct {
	Identifier string         // 控件唯一标识
	Value      string         // 控件值
	Index      int            // 值索引
	Owner      Pointer        // 所属对象
	Sender     Pointer        // 控件对象
	Arguments  *OCGoArguments // 参数
}

func CCallbackContextToGo(cContext *C.TCallbackContext) *TCallbackContext {
	ctx := &TCallbackContext{
		Identifier: C.GoString(cContext.identifier),
		Value:      C.GoString(cContext.value),
		Index:      int(cContext.index),
		Owner:      cContext.owner,
		Sender:     cContext.sender,
	}
	cArguments := cContext.arguments
	if cArguments != nil {
		ctx.Arguments = &OCGoArguments{arguments: Pointer(cArguments), count: int(cArguments.Count)}
	}
	return ctx
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
