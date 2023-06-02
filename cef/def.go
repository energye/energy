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
	"github.com/energye/energy/v2/cef/internal/version"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
)

// defInit
//	GlobalInit
func defInit() {
	// Set pascal lib build version
	r1, _, _ := imports.Proc(internale_LibBuildVersion).Call()
	version.SetLibBuildVersion(api.GoStr(r1))
}

// EnergyVersion 返回 Energy 版本
func EnergyVersion() string {
	return version.Version()
}

// CEFVersion 返回CEF版本
func CEFVersion() string {
	return version.CEFVersion()
}

// LibBuildVersion 返回lib-lcl构建版本
func LibBuildVersion() string {
	return version.LibBuildVersion()
}

// setMacOSXCommandLine
// 针对 MacOSX 设置命令行参数
//
// 没找到什么好的方式，只能这样设置
func setMacOSXCommandLine(commandLine uintptr) {
	imports.Proc(internale_SetMacOSXCommandLine).Call(commandLine)
}
