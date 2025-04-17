// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

//go:build !prod
// +build !prod

package initialize

import (
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/cef/internal/initialize/macapp"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/pkgs/libname"
	"path"
	"path/filepath"
)

// 开发环境加载 libLCL，在程序启动后读取用户目录 .energy 配置文件
// 不同操作系统加载方式也不同
// 优先级: 1. 自定义, 2. 当前执行目录, 3. .energy 配置文件， 4. 相对目录
// .energy 配置文件 root 字段是 energy 框架所在根目录, framework 字段是当前正在使用的 CEF Framework
// 获得一个完整目录 [root]+energy+[framework]
func loadLibLCL(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	if common.IsDarwin() {
		// 开发模式 自动生成 xxx.app
		macapp.Init()
	}
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
				// 开发环境配置目录
				if config.Get() != nil {
					libPath = filepath.Join(config.Get().FrameworkPath(), libname.GetDLLName())
				} else {
					// 最后尝试相对目录
					libPath = libname.GetDLLName()
				}
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
			panic(`Hint:
  Failed initialize LibLCL, check the development environment
  Use CLI: 
    [energy env] : Check the configuration of the development environment
    [energy install] : Installation development environment
`)
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
