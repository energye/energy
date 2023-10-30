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
	"embed"
	. "github.com/energye/energy/v2/cef/process"
	. "github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports/tempdll"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/pkgs/libname"
	"github.com/energye/golcl/pkgs/macapp"
	"os"
	"runtime"
)

// ExceptionCallback 异常回调函数
type ExceptionCallback func(sender lcl.IObject, e *lcl.Exception)

var exceptionCallback ExceptionCallback

// SetOnException 设置 lib-lcl -> CEF 低层异常捕获回调函数
func SetOnException(exception ExceptionCallback) {
	exceptionCallback = exception
}

// GlobalInit 全局初始化
//  需要手动调用的函数,在main函数中调用
//	参数:
//    libs 内置到应用程序的类库
//    resources 内置到应用程序的资源文件
//  MacOSX环境, goland、ide等开发环境需配置命令行参数[energy_env=dev]以保证应用正常运行
func GlobalInit(libs *embed.FS, resources *embed.FS) {
	if IsDarwin() {
		macapp.MacApp.IsCEF(true)
		//MacOSX环境, ide开发环境需命令行参数[energy_env=dev]以保证应用正常运行
		var env = func() string {
			energyEnv := Args.Args("energy_env")
			env := Args.Args("env")
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
	// 如果使用 liblclbinres 编译则通过该方式加载动态库
	if dllPath, dllOk := tempdll.CheckAndReleaseDLL(); dllOk {
		api.SetLoadUILibCallback(func() (path string, ok bool) {
			if runtime.GOOS == "darwin" {
				//MacOSX从Frameworks加载
				libname.LibName = "@executable_path/../Frameworks/" + libname.GetDLLName()
			} else {
				libname.LibName = dllPath
			}
			path = dllPath
			ok = dllOk
			return
		})
	}
	// go lcl
	inits.Init(libs, resources)
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
	//应用低层出错异常捕获
	lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
		if exceptionCallback != nil {
			exceptionCallback(sender, e)
		} else {
			logger.Error("ResultString:", e.Message())
		}
	})
}

// v8init v8初始化
func v8init() {
	if Args.IsMain() || Args.IsRender() {
		//ipc初始化
		ipcInit()
		//bind初始化
		//bindInit()
	}
}
