//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/energy/v2/api/imports"
	"github.com/energye/energy/v2/api/libname"
)

var (
	uiLib           imports.DLL                 // 全局导入库
	loadLibCallback func() (imports.DLL, error) // 自定义加载liblcl动态库回调函数
)

// SetLoadLibCallback
//  设置加载liblcl动态库回调函数
//  如果设置该回调函数我们可以自定义加载动态链接库
//  在调用 inits.Init 之前设置
func SetLoadLibCallback(fn func() (imports.DLL, error)) {
	if loadLibCallback == nil {
		loadLibCallback = fn
	}
}

func loadUILib() imports.DLL {
	if loadLibCallback != nil {
		dll, _ := loadLibCallback()
		return dll
	} else {
		lib, err := imports.NewDLL(libname.LibName)
		if lib == 0 && err != nil {
			panic(err)
		}
		return lib
	}
}

func closeLib() {
	if uiLib != 0 {
		uiLib.Release()
		uiLib = 0
	}
}

// 调用预定义导入API
// 由 predefImport 提供
func defSyscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := LCLPreDef().Proc(trap).Call(args...)
	return r1
}
