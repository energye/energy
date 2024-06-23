//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Energy Global initialization

package cef

import (
	"github.com/energye/cef/cef"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/inits"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/pkgs/macapp"
	"github.com/energye/lcl/process"
	"github.com/energye/lcl/tools"
	"os"
)

// GlobalInit 全局初始化
//
//	 需要手动调用的函数,在main函数中调用
//		参数:
//	   libs 内置到应用程序的类库
//	   resources 内置到应用程序的资源文件
//	 MacOSX环境, goland、ide等开发环境需配置命令行参数[energy_env=dev]以保证应用正常运行
func GlobalInit(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	if tools.IsDarwin() {
		macapp.MacApp.IsCEF(true)
		//MacOSX环境, ide开发环境需命令行参数[energy_env=dev]以保证应用正常运行
		var env = func() string {
			energyEnv, env := process.Args.Args("energy_env"), process.Args.Args("env")
			if energyEnv != "" {
				return energyEnv
			}
			if env != "" {
				return env
			}
			return ""
		}()
		if env != "" {
			macapp.MacApp.SetEnergyEnv(env)
		}
	}

	emfs.SetEMFS(libs, resources)
	// go lcl init
	inits.InitAll()
	// macos command line
	if tools.IsDarwin() {
		argsList := lcl.NewStringList()
		for _, v := range os.Args {
			argsList.Add(v)
		}
		// 手动设置进程命令参数
		cef.SetCommandLine(argsList)
		argsList.Free()
	}
}

// v8init v8初始化
func v8init() {
	if process.Args.IsMain() || process.Args.IsRender() {
		//ipc初始化
		ipcInit()
	}
}
