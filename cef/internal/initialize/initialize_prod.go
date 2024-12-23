//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build prod
// +build prod

package initialize

import (
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/pkgs/libname"
	"path"
)

// 发布环境加载 libLCL，不再依赖 .energy 配置文件
// 优先级: 1. 自定义, 2. 当前执行目录, 3. 相对目录
// 不同操作系统加载方式也不同
func loadLibLCL(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	// LCL 初始化时回调， 返回 lib 地址
	api.SetLoadLibCallback(func() (liblcl dllimports.DLL, err error) {
		libPath := libname.LibName
		if libPath != "" {
			// 自定义加载目录
			liblcl, err = dllimports.NewDLL(libPath)
		} else if common.IsDarwin() {
			// MacOS 固定加载目录
			libPath = "@executable_path/../Frameworks/" + libname.GetDLLName()
		} else {
			// Windows, Linux
			// 优先当前执行目录
			currentPathLibPath := path.Join(consts.ExeDir, libname.GetDLLName())
			if tools.IsExist(currentPathLibPath) {
				libPath = currentPathLibPath
			} else {
				// 最后尝试相对目录
				libPath = libname.GetDLLName()
			}
		}
		// 加载 LibLCL
		if libPath != "" {
			libname.LibName = libPath
			liblcl, err = dllimports.NewDLL(libPath)
		}
		if liblcl == 0 {
			if err != nil {
				println("Load LibLCL Error:", err.Error())
			}
			println("LibLCL Path:", libname.LibName)
			panic(`Failed initialize LibLCL`)
		} else {
			imports.LibEnergy().SetOk(true)
			imports.LibEnergy().SetDll(liblcl)
			imports.LibLCLExt().SetOk(true)
			imports.LibLCLExt().SetDll(liblcl)
		}
		return
	})
	emfs.SetEMFS(libs, resources)
}
