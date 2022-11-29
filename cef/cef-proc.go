//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

//--------TCEFApplication proc begin--------

// TCEFApplication AddCrDelegate
func AddCrDelegate() uintptr {
	r1, _, _ := Proc("CEF_AddCrDelegate").Call()
	return r1
}

// TCEFApplication _CEFApplication_Create
func _CEFApplication_Create(config uintptr) uintptr {
	r1, _, _ := Proc("CEFApplication_Create").Call(config)
	return r1
}

// TCEFApplication _CEFApplication_Free
func _CEFApplication_Free() uintptr {
	r1, _, _ := Proc("CEFApplication_Free").Call()
	return r1
}

// TCEFApplication _CEFStartMainProcess
func _CEFStartMainProcess(instance uintptr) uintptr {
	r1, _, _ := Proc("CEFStartMainProcess").Call(instance)
	return r1
}

// TCEFApplication _CEFStartSubProcess
func _CEFStartSubProcess(instance uintptr) uintptr {
	r1, _, _ := Proc("CEFStartSubProcess").Call(instance)
	return r1
}

func _AddCustomCommandLine(commandLine, value uintptr) {
	Proc("AddCustomCommandLine").Call(commandLine, value)
}

func _CEFApplication_ExecuteJS(browseId int32, jsCode string) {
	Proc("CEFApplication_ExecuteJS").Call(uintptr(browseId), api.GoStrToDStr(jsCode))
}

//--------TCEFApplication proc end--------

//--------TCEFWindowParent proc begin--------

// TCEFWindowParent _CEFWindow_UpdateSize
func _CEFWindow_UpdateSize(instance uintptr) {
	Proc("CEFWindow_UpdateSize").Call(instance)
}

func _CEFWindow_OnEnter(instance uintptr, fn interface{}) {
	Proc("CEFWindow_OnEnter").Call(instance, api.GetAddEventToMapFn()(instance, fn))
}

func _CEFWindow_OnExit(instance uintptr, fn interface{}) {
	Proc("CEFWindow_OnExit").Call(instance, api.GetAddEventToMapFn()(instance, fn))
}

// TCEFWindowParent _CEFLinkedWindow_UpdateSize
func _CEFLinkedWindow_UpdateSize(instance uintptr) {
	Proc("CEFLinkedWindow_UpdateSize").Call(instance)
}

func _CEFLinkedWindow_SetChromium(instance uintptr, chromium *TCEFChromium, tag int32) {
	Proc("CEFLinkedWindow_SetChromium").Call(instance, chromium.instance, uintptr(tag))
}

func _CEFLinkedWindow_OnEnter(instance uintptr, fn interface{}) {
	Proc("CEFLinkedWindow_OnEnter").Call(instance, api.GetAddEventToMapFn()(instance, fn))
}

func _CEFLinkedWindow_OnExit(instance uintptr, fn interface{}) {
	Proc("CEFLinkedWindow_OnExit").Call(instance, api.GetAddEventToMapFn()(instance, fn))
}

//--------TCEFWindowParent proc end--------

/* 通用 PROC --- begin ---*/

// _Create 创建
func _Create(name string, owner uintptr) uintptr {
	name = Proc_Concat_Name(name, "Create")
	r1, _, _ := Proc(name).Call(owner)
	return r1
}

// _GetHandle 获取组件句柄
func _GetHandle(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "GetHandle")
	ret, _, _ := Proc(name).Call(instance)
	return ret
}

func _DestroyChildWindow(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "DestroyChildWindow")
	ret, _, _ := Proc(name).Call(instance)
	return ret
}

// _HandleAllocated
func _HandleAllocated(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "HandleAllocated")
	r1, _, _ := Proc(name).Call(instance)
	return r1
}

// _CreateHandle
func _CreateHandle(name string, instance uintptr) {
	name = Proc_Concat_Name(name, "CreateHandle")
	Proc(name).Call(instance)
}

// _Free 释放
func _Free(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "Free")
	r1, _, _ := Proc(name).Call(instance)
	return r1
}

// _SetParent 设置父组件
func _SetParent(name string, instance, owner uintptr) {
	name = Proc_Concat_Name(name, "SetParent")
	Proc(name).Call(instance, owner)
}

// _Align 获取控件自动调整。
func _Align(name string, instance uintptr) types.TAlign {
	name = Proc_Concat_Name(name, "GetAlign")
	r1, _, _ := Proc(name).Call(instance)
	return types.TAlign(r1)
}

// _SetAlign 设置获取控件自动调整。
func _SetAlign(name string, instance uintptr, value types.TAlign) {
	name = Proc_Concat_Name(name, "SetAlign")
	Proc(name).Call(instance, uintptr(value))
}

// _Anchors 获取四个角位置的锚点。
func _Anchors(name string, instance uintptr) types.TAnchors {
	name = Proc_Concat_Name(name, "GetAnchors")
	r1, _, _ := Proc(name).Call(instance)
	return types.TAnchors(r1)
}

// _SetAnchors 设置四个角位置的锚点。
func _SetAnchors(name string, instance uintptr, value types.TAnchors) {
	name = Proc_Concat_Name(name, "SetAnchors")
	Proc(name).Call(instance, uintptr(value))
}

// _GetVisible 获取控件可视。
func _GetVisible(name string, instance uintptr) bool {
	name = Proc_Concat_Name(name, "GetVisible")
	ret, _, _ := Proc(name).Call(instance)
	return api.DBoolToGoBool(ret)
}

// _SetVisible 设置控件可视。
func _SetVisible(name string, instance uintptr, value bool) {
	name = Proc_Concat_Name(name, "SetVisible")
	Proc(name).Call(instance, api.GoBoolToDBool(value))
}

// _GetEnabled 获取是否启用
func _GetEnabled(name string, instance uintptr) bool {
	name = Proc_Concat_Name(name, "GetEnabled")
	ret, _, _ := Proc(name).Call(instance)
	return api.DBoolToGoBool(ret)
}

// _SetEnabled 设置是否启用
func _SetEnabled(name string, instance uintptr, value bool) {
	name = Proc_Concat_Name(name, "SetEnabled")
	Proc(name).Call(instance, api.GoBoolToDBool(value))
}

// _GetLeft 获取左边距
func _GetLeft(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetLeft")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// _SetLeft 设置左边距
func _SetLeft(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetLeft")
	Proc(name).Call(instance, uintptr(value))
}

// _Top 获取上边距
func _GetTop(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetTop")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// _SetTop 设置上边距
func _SetTop(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetTop")
	Proc(name).Call(instance, uintptr(value))
}

// _GetWidth 获取宽度
func _GetWidth(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetWidth")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// _SetWidth 设置宽度
func _SetWidth(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetWidth")
	Proc(name).Call(instance, uintptr(value))
}

// _GetHeight 获取高度
func _GetHeight(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetHeight")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// _SetHeight 设置高度
func _SetHeight(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetHeight")
	Proc(name).Call(instance, uintptr(value))
}

// _GetBoundsRect
func _GetBoundsRect(name string, instance uintptr) types.TRect {
	var ret types.TRect
	name = Proc_Concat_Name(name, "GetBoundsRect")
	Proc(name).Call(instance, uintptr(unsafe.Pointer(&ret)))
	return ret
}

// _SetBoundsRect
func _SetBoundsRect(name string, instance uintptr, value types.TRect) {
	name = Proc_Concat_Name(name, "SetBoundsRect")
	Proc(name).Call(instance, uintptr(unsafe.Pointer(&value)))
}

// _GetName 获取组件名称。
func _GetName(name string, instance uintptr) string {
	name = Proc_Concat_Name(name, "GetName")
	ret, _, _ := Proc(name).Call(instance)
	return api.DStrToGoStr(ret)
}

// _SetName 设置组件名称。
func _SetName(name string, instance uintptr, value string) {
	name = Proc_Concat_Name(name, "SetName")
	Proc(name).Call(instance, api.GoStrToDStr(value))
}

/* 通用 PROC --- end ---*/

// other

// 针对 MacOSX 设置命令行参数
//
//没找到什么好的方式，只能这样设置
func setMacOSXCommandLine(commandLine uintptr) {
	Proc("SetMacOSXCommandLine").Call(commandLine)
}
