//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy 扩展定义, 编译版本, CEF版本

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
)

var (
	cef_version       string
	lib_build_version string
)

// CEFVersion 返回CEF版本
func CEFVersion() string {
	if cef_version == "" {
		r1, _, _ := imports.Proc(internale_CEFVersion).Call()
		cef_version = api.GoStr(r1)
	}
	return cef_version
}

// LibBuildVersion 返回lib-lcl构建版本
func LibBuildVersion() string {
	if lib_build_version == "" {
		r1, _, _ := imports.Proc(internale_LibBuildVersion).Call()
		lib_build_version = api.GoStr(r1)
	}
	return lib_build_version
}

// setMacOSXCommandLine
// 针对 MacOSX 设置命令行参数
//
// 没找到什么好的方式，只能这样设置
func setMacOSXCommandLine(commandLine uintptr) {
	imports.Proc(internale_SetMacOSXCommandLine).Call(commandLine)
}
