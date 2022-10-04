//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/commons"
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

// Create 创建
func Create(name string, owner uintptr) uintptr {
	name = Proc_Concat_Name(name, "Create")
	r1, _, _ := Proc(name).Call(owner)
	return r1
}

// GetHandle 获取组件句柄
func GetHandle(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "GetHandle")
	ret, _, _ := Proc(name).Call(instance)
	return ret
}

func DestroyChildWindow(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "DestroyChildWindow")
	ret, _, _ := Proc(name).Call(instance)
	return ret
}

// _HandleAllocated
func HandleAllocated(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "HandleAllocated")
	r1, _, _ := Proc(name).Call(instance)
	return r1
}

// CreateHandle
func CreateHandle(name string, instance uintptr) {
	name = Proc_Concat_Name(name, "CreateHandle")
	Proc(name).Call(instance)
}

// Free 释放
func Free(name string, instance uintptr) uintptr {
	name = Proc_Concat_Name(name, "Free")
	r1, _, _ := Proc(name).Call(instance)
	return r1
}

// SetParent 设置父组件
func SetParent(name string, instance, owner uintptr) {
	name = Proc_Concat_Name(name, "SetParent")
	Proc(name).Call(instance, owner)
}

// Align 获取控件自动调整。
func Align(name string, instance uintptr) types.TAlign {
	name = Proc_Concat_Name(name, "GetAlign")
	r1, _, _ := Proc(name).Call(instance)
	return types.TAlign(r1)
}

// SetAlign 设置获取控件自动调整。
func SetAlign(name string, instance uintptr, value types.TAlign) {
	name = Proc_Concat_Name(name, "SetAlign")
	Proc(name).Call(instance, uintptr(value))
}

// Anchors 获取四个角位置的锚点。
func Anchors(name string, instance uintptr) types.TAnchors {
	name = Proc_Concat_Name(name, "GetAnchors")
	r1, _, _ := Proc(name).Call(instance)
	return types.TAnchors(r1)
}

// SetAnchors 设置四个角位置的锚点。
func SetAnchors(name string, instance uintptr, value types.TAnchors) {
	name = Proc_Concat_Name(name, "SetAnchors")
	Proc(name).Call(instance, uintptr(value))
}

// GetVisible 获取控件可视。
func GetVisible(name string, instance uintptr) bool {
	name = Proc_Concat_Name(name, "GetVisible")
	ret, _, _ := Proc(name).Call(instance)
	return api.DBoolToGoBool(ret)
}

// SetVisible 设置控件可视。
func SetVisible(name string, instance uintptr, value bool) {
	name = Proc_Concat_Name(name, "SetVisible")
	Proc(name).Call(instance, api.GoBoolToDBool(value))
}

// GetEnabled 获取是否启用
func GetEnabled(name string, instance uintptr) bool {
	name = Proc_Concat_Name(name, "GetEnabled")
	ret, _, _ := Proc(name).Call(instance)
	return api.DBoolToGoBool(ret)
}

// SetEnabled 设置是否启用
func SetEnabled(name string, instance uintptr, value bool) {
	name = Proc_Concat_Name(name, "SetEnabled")
	Proc(name).Call(instance, api.GoBoolToDBool(value))
}

// GetLeft 获取左边距
func GetLeft(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetLeft")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// SetLeft 设置左边距
func SetLeft(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetLeft")
	Proc(name).Call(instance, uintptr(value))
}

// Top 获取上边距
func GetTop(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetTop")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// SetTop 设置上边距
func SetTop(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetTop")
	Proc(name).Call(instance, uintptr(value))
}

// GetWidth 获取宽度
func GetWidth(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetWidth")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// SetWidth 设置宽度
func SetWidth(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetWidth")
	Proc(name).Call(instance, uintptr(value))
}

// GetHeight 获取高度
func GetHeight(name string, instance uintptr) int32 {
	name = Proc_Concat_Name(name, "GetHeight")
	ret, _, _ := Proc(name).Call(instance)
	return int32(ret)
}

// SetHeight 设置高度
func SetHeight(name string, instance uintptr, value int32) {
	name = Proc_Concat_Name(name, "SetHeight")
	Proc(name).Call(instance, uintptr(value))
}

// GetBoundsRect
func GetBoundsRect(name string, instance uintptr) types.TRect {
	var ret types.TRect
	name = Proc_Concat_Name(name, "GetBoundsRect")
	Proc(name).Call(instance, uintptr(unsafe.Pointer(&ret)))
	return ret
}

// SetBoundsRect
func SetBoundsRect(name string, instance uintptr, value types.TRect) {
	name = Proc_Concat_Name(name, "SetBoundsRect")
	Proc(name).Call(instance, uintptr(unsafe.Pointer(&value)))
}

// GetName 获取组件名称。
func GetName(name string, instance uintptr) string {
	name = Proc_Concat_Name(name, "GetName")
	ret, _, _ := Proc(name).Call(instance)
	return api.DStrToGoStr(ret)
}

// SetName 设置组件名称。
func SetName(name string, instance uintptr, value string) {
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
