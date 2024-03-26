//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// Package cef (Chromium Embedded Framework) 实现模块
// 包含所有 CEF API 实现
package cef

import (
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"time"
	"unsafe"
)

const intSize = 32 << (^uint(0) >> 63) // 4 或 8

// IObject 根对象
type IObject = lcl.IObject
type TObject = lcl.TObject

// IComponent 根组件
type IComponent = lcl.IComponent
type TComponent = lcl.TComponent

// IWinControl 根Control
type IWinControl = lcl.IWinControl
type TWinControl = lcl.TWinControl

// IUnknown 根接口
//
//	UUID: '{00000000-0000-0000-C000-000000000046}'
type IUnknown = lcl.IUnknown
type Unknown = lcl.Unknown

type TStream = lcl.TStream
type IStream = lcl.IStream

type TBitmap = lcl.TBitmap
type IBitmap = lcl.IBitmap

type TPersistent = lcl.TPersistent
type IPersistent = lcl.IPersistent

type TRasterImage = lcl.TRasterImage
type IRasterImage = lcl.IRasterImage

type TCustomBitmap = lcl.TCustomBitmap
type ICustomBitmap = lcl.ICustomBitmap

type TFPImageBitmap = lcl.TFPImageBitmap
type IFPImageBitmap = lcl.IFPImageBitmap

type TStrings = lcl.TStrings
type IStrings = lcl.IStrings

type TStringList = lcl.TStringList
type IStringList = lcl.IStringList

type TCustomMemoryStream = lcl.TCustomMemoryStream
type ICustomMemoryStream = lcl.ICustomMemoryStream

type ICustomPanel = lcl.ICustomPanel
type TCustomPanel = lcl.TCustomPanel

type ICustomForm = lcl.ICustomForm
type TCustomForm = lcl.TCustomForm

type unsafePointer = unsafe.Pointer

type NativeUInt = types.NativeUInt

type TClass = types.TClass

type TPoint = types.TPoint

type TRect = types.TRect

type DWORD = types.DWORD

type HMENU = types.HMENU

type Cardinal = types.Cardinal

type Integer = types.Integer

type integer = types.Integer

type Pointer = types.Pointer

type Time = time.Time

type THandle = types.THandle

type HWND = types.HWND

type TDragKind = types.TDragKind
type TCursor = types.TCursor
type TDragMode = types.TDragMode

type TThreadPriority = types.TThreadPriority

// PByte = ^Byte
type PByte = types.PByte

// PString = ^PChar
type PString = uintptr

// PSingle = ^Single
type PSingle = *uintptr

// PPSingle = ^PSingle
type PPSingle = *PSingle

type WPARAM = types.WPARAM
type LPARAM = types.LPARAM
type LRESULT = types.LRESULT

type TMessage = types.TMessage
type TDragState = int32

type Single = types.Single

// TDateTime => Double
//
//	DateTimeToUnix 转换 float64 => int64
//	UnixToDateTime 转换 int64 => float64
//	DateTimeToGoTime 转换 float64 => time.Time
//	DateTimeToDTime 转换 time.Time => float64
type TDateTime = float64

func getPointer(ptr uintptr) unsafePointer {
	return unsafePointer(ptr)
}

// GetInstance As操作的简化。
//
// Simplification of As operation.
//
//go:noinline
func GetInstance(value interface{}) unsafePointer {
	return lcl.GetInstance(value)
}

// SetObjectInstance 设置对你指针实例值, 用于外部组件创建
func SetObjectInstance(object interface{}, instance unsafePointer) {
	if object == nil {
		return
	}
	switch object.(type) {
	case IObject:
		lcl.SetObjectInstance(object.(IObject), instance)
	case IUnknown:
		lcl.SetUnknownInstance(object.(IUnknown), instance)
	}
}

// GetObjectUintptr 获取对象指针地址值
func GetObjectUintptr(object interface{}) uintptr {
	if object == nil {
		return 0
	}
	switch object.(type) {
	case IObject:
		return lcl.GetObjectUintptr(object.(IObject))
	case IUnknown:
		return lcl.GetUnknownUintptr(object.(IUnknown))
	}
	return 0
}

// DateTimeToUnix float64格式时间转换, 仅支持到秒转换
func DateTimeToUnix(dateTime float64) int64 {
	return api.DDateTimeToUnix(dateTime)
}

// UnixToDateTime int64格式时间转换，仅支持到秒转换
func UnixToDateTime(dateTime int64) float64 {
	return api.DUnixToDateTime(dateTime)
}

// DateTimeToGoTime float64格式时间转换, 仅支持到秒转换
func DateTimeToGoTime(dateTime float64) time.Time {
	return time.UnixMilli(DateTimeToUnix(dateTime))
}

// DateTimeToDTime time.Time 格式时间转换, 仅支持到秒转换
func DateTimeToDTime(dateTime time.Time) float64 {
	return UnixToDateTime(dateTime.Unix())
}

// SetCommandLine
//
//	针对 MacOS 设置命令行参数
//	没找到什么好的方式
func SetCommandLine(argsList lcl.IStringList) {
	api.CEFPreDef().SysCallN(10, argsList.Instance())
}

// AddCrDelegate 针对 MacOS
func AddCrDelegate() {
	api.CEFPreDef().SysCallN(11)
}
