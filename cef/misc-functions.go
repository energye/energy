//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF的复杂函数导出

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
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
