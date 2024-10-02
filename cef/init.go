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
	_ "github.com/energye/energy/v2/cef/internal/def"
	. "github.com/energye/energy/v2/cef/process"
	. "github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/common/imports/tempdll"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/pkgs/libname"
	"github.com/energye/golcl/pkgs/macapp"
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
	if IsDarwin() {
		macapp.MacApp.IsCEF(true)
		//MacOSX环境, ide开发环境需命令行参数[energy_env=dev]以保证应用正常运行
		var env = func() string {
			energyEnv, env := Args.Args("energy_env"), Args.Args("env")
			if energyEnv != "" {
				return energyEnv
			}
			if env != "" {
				return env
			}
			return ""
		}()
		if env != "" {
			macapp.MacApp.SetEnergyEnv(macapp.ENERGY_ENV(env))
		}
	}

	// lcl 初始化时回调，如果设置了该回调函数需要通过该函数返回liblcl库
	api.SetLoadLibCallback(func() (liblcl dllimports.DLL, err error) {
		// load liblcl
		// liblcl name, 不为空时表示自定义加载目录
		if libname.LibName == "" {
			// 如果使用内置dll编译,则通过该方式加载动态库
			path, fullPath, ok := tempdll.CheckAndReleaseDLL(libname.GetDLLName())
			if ok {
				libname.LibName = fullPath
				// 设置到tempDllDir, 使用tempdll将最优先从该目录加载
				libname.SetTempDllDir(path)
			}
		}
		if IsDarwin() { // MacOS固定加载目录
			//MacOSX从Frameworks加载
			libname.LibName = "@executable_path/../Frameworks/" + libname.GetDLLName()
		} else if libname.LibName == "" {
			libname.LibName = libname.LibPath(libname.GetDLLName())
		}
		if libname.LibName != "" {
			liblcl, err = dllimports.NewDLL(libname.LibName)
		}
		if liblcl == 0 {
			if err != nil {
				println("LoadLibrary liblcl ERROR:", err.Error())
			}
			panic(`Hint:
	Golcl dependency library liblcl was not found
	Please check whether liblcl exists locally
	If local liblcl exist, please put it in the specified location, If it does not exist, please download it from the Energy official website.
	Configuration Location:
		1. Current program execution directory
		2. USER_HOME/golcl/
		3. Environment variables LCL_HOME or ENERGY_HOME
			environment variable LCL_HOME is configured preferentially in the non-energy framework
			environment variable ENERGY_HOME takes precedence in the Energy framework
			ENERGY_HOME environment variable is recommended
`)
		}
		// 加载完成设置到libenergy全局
		imports.LibEnergy().SetOk(true)
		imports.LibEnergy().SetDll(liblcl)
		imports.LibLCLExt().SetOk(true)
		imports.LibLCLExt().SetDll(liblcl)
		return
	})
	emfs.SetEMFS(libs, resources)
	// go lcl init
	inits.InitAll()
	// macos command line
	if IsDarwin() {
		argsList := lcl.NewStringList()
		for _, v := range os.Args {
			argsList.Add(v)
		}
		// 手动设置进程命令参数
		SetCommandLine(argsList)
		argsList.Free()
	}
	// main thread run call
	applicationQueueAsyncCallInit()
}

// v8init v8初始化
func v8init() {
	if Args.IsMain() || Args.IsRender() {
		//ipc初始化
		ipcInit()
	}
}
