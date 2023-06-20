//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 函数工具

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	t "github.com/energye/golcl/lcl/types"
	"unsafe"
)

// WindowInfoAsChild BrowserWindow 设置到指定窗口做为子窗口
func WindowInfoAsChild(windowInfo, windowHandle uintptr, windowName string) {
	imports.Proc(def.Misc_WindowInfoAsChild).Call(windowInfo, windowHandle, api.PascalStr(windowName))
}

// WindowInfoAsPopUp BrowserWindow 设置到做为弹出窗口
func WindowInfoAsPopUp(windowInfo, windowParent uintptr, windowName string) {
	imports.Proc(def.Misc_WindowInfoAsPopUp).Call(windowInfo, windowParent, api.PascalStr(windowName))
}

// WindowInfoAsWindowless BrowserWindow 设置到做为无窗口
func WindowInfoAsWindowless(windowInfo, windowParent uintptr, windowName string) {
	imports.Proc(def.Misc_WindowInfoAsWindowless).Call(windowInfo, windowParent, api.PascalStr(windowName))
}

// RegisterExtension 注册JS扩展
//  将自定义JS代码植入到当前浏览器
//	在 WebKitInitialized 回调函数中使用
//	参数:
//		name: 根对象名, 不允许使用默认的内部名称, 参阅 isInternalBind 函数
//		code: js code
//		handler: 处理器, 根据本地函数名回调该处理器
func RegisterExtension(name, code string, handler *ICefV8Handler) {
	registerExtension(name, code, handler)
}

func registerExtension(name, code string, handler *ICefV8Handler) {
	imports.Proc(def.Misc_CefRegisterExtension).Call(api.PascalStr(name), api.PascalStr(code), handler.Instance())
}

func CheckSubprocessPath(subprocessPath string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := imports.Proc(def.Misc_CheckSubprocessPath).Call(api.PascalStr(subprocessPath), uintptr(unsafe.Pointer(&missingFiles)))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func CheckLocales(localesDirPath, localesRequired string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := imports.Proc(def.Misc_CheckLocales).Call(api.PascalStr(localesDirPath), uintptr(unsafe.Pointer(&missingFiles)), api.PascalStr(localesRequired))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func CheckResources(resourcesDirPath string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := imports.Proc(def.Misc_CheckResources).Call(api.PascalStr(resourcesDirPath), uintptr(unsafe.Pointer(&missingFiles)))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func CheckDLLs(frameworkDirPath string) (missingFiles string, result bool) {
	var missingFilesPtr uintptr
	r1, _, _ := imports.Proc(def.Misc_CheckDLLs).Call(api.PascalStr(frameworkDirPath), uintptr(unsafe.Pointer(&missingFiles)))
	missingFiles = api.GoStr(missingFilesPtr)
	result = api.GoBool(r1)
	return
}

func RegisterSchemeHandlerFactory(schemeName, domainName string, handler TCefResourceHandlerClass) bool {
	r1, _, _ := imports.Proc(def.Misc_CefRegisterSchemeHandlerFactory).Call(api.PascalStr(schemeName), api.PascalStr(domainName), uintptr(handler))
	return api.GoBool(r1)
}

func ClearSchemeHandlerFactories() bool {
	r1, _, _ := imports.Proc(def.Misc_CefClearSchemeHandlerFactories).Call()
	return api.GoBool(r1)
}

func GetMimeType(extension string) string {
	r1, _, _ := imports.Proc(def.Misc_CefGetMimeType).Call(api.PascalStr(extension))
	return api.GoStr(r1)
}

func DeviceToLogicalInt32(value int32, deviceScaleFactor float64) int32 {
	r1, _, _ := imports.Proc(def.Misc_DeviceToLogicalInt32).Call(uintptr(value), uintptr(unsafe.Pointer(&deviceScaleFactor)))
	return int32(r1)
}

func DeviceToLogicalFloat32(value float32, deviceScaleFactor float64) (result float32) {
	imports.Proc(def.Misc_DeviceToLogicalFloat32).Call(uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&deviceScaleFactor)), uintptr(unsafe.Pointer(&result)))
	return
}

func DeviceToLogicalMouse(event *TCefMouseEvent, deviceScaleFactor float64) {
	imports.Proc(def.Misc_DeviceToLogicalMouse).Call(uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func DeviceToLogicalTouch(event *TCefTouchEvent, deviceScaleFactor float64) {
	imports.Proc(def.Misc_DeviceToLogicalTouch).Call(uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func DeviceToLogicalPoint(point *t.TPoint, deviceScaleFactor float64) {
	imports.Proc(def.Misc_DeviceToLogicalPoint).Call(uintptr(unsafe.Pointer(point)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func LogicalToDeviceInt32(value int32, deviceScaleFactor float64) int32 {
	r1, _, _ := imports.Proc(def.Misc_LogicalToDeviceInt32).Call(uintptr(value), uintptr(unsafe.Pointer(&deviceScaleFactor)))
	return int32(r1)
}

func LogicalToDeviceRect(rect *TCefRect, deviceScaleFactor float64) {
	imports.Proc(def.Misc_LogicalToDeviceRect).Call(uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(&deviceScaleFactor)))
}

func CefCursorToWindowsCursor(cefCursor consts.TCefCursorType) (result t.TCursor) {
	switch cefCursor {
	case consts.CT_POINTER:
		result = t.CrArrow
	case consts.CT_CROSS:
		result = t.CrCross
	case consts.CT_HAND:
		result = t.CrHandPoint
	case consts.CT_IBEAM:
		result = t.CrIBeam
	case consts.CT_WAIT:
		result = t.CrHourGlass
	case consts.CT_HELP:
		result = t.CrHelp
	case consts.CT_EASTRESIZE:
		result = t.CrSizeWE
	case consts.CT_NORTHRESIZE:
		result = t.CrSizeNS
	case consts.CT_NORTHEASTRESIZE:
		result = t.CrSizeNESW
	case consts.CT_NORTHWESTRESIZE:
		result = t.CrSizeNWSE
	case consts.CT_SOUTHRESIZE:
		result = t.CrSizeNS
	case consts.CT_SOUTHEASTRESIZE:
		result = t.CrSizeNWSE
	case consts.CT_SOUTHWESTRESIZE:
		result = t.CrSizeNESW
	case consts.CT_WESTRESIZE:
		result = t.CrSizeWE
	case consts.CT_NORTHSOUTHRESIZE:
		result = t.CrSizeNS
	case consts.CT_EASTWESTRESIZE:
		result = t.CrSizeWE
	case consts.CT_NORTHEASTSOUTHWESTRESIZE:
		result = t.CrSizeNESW
	case consts.CT_NORTHWESTSOUTHEASTRESIZE:
		result = t.CrSizeNWSE
	case consts.CT_COLUMNRESIZE:
		result = t.CrHSplit
	case consts.CT_ROWRESIZE:
		result = t.CrVSplit
	case consts.CT_MOVE:
		result = t.CrSizeAll
	case consts.CT_PROGRESS:
		result = t.CrAppStart
	case consts.CT_NONE:
		result = t.CrNone
	case consts.CT_NODROP, consts.CT_NOTALLOWED:
		result = t.CrNo
	case consts.CT_GRAB, consts.CT_GRABBING:
		result = t.CrDrag
	default:
		result = t.CrDefault
	}
	return
}

func CefColorGetA(color types.TCefColor) uint8 {
	return uint8(color>>24) & 0xFF
}

func CefColorGetR(color types.TCefColor) uint8 {
	return uint8(color>>16) & 0xFF
}

func CefColorGetG(color types.TCefColor) uint8 {
	return uint8(color>>8) & 0xFF
}

func CefColorGetB(color types.TCefColor) uint8 {
	return uint8(color) & 0xFF
}

func CefColorSetARGB(a, r, g, b byte) types.TCefColor {
	return types.TCefColor((uint32(a) << 24) | (uint32(r) << 16) | (uint32(g) << 8) | (uint32(b)))
}
