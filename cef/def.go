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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
)

// SetCommandLine
// 针对 MacOS 设置命令行参数
//
// 没找到什么好的方式，只能这样设置
func SetCommandLine(argsList *lcl.TStringList) {
	imports.Proc(def.SetCommandLine).Call(argsList.Instance())
}
