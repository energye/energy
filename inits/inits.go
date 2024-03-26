//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package inits

import (
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/api/libname"
	"github.com/energye/energy/v2/emfs"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/locales/zh_CN"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/lcl/rtl/version"
	"github.com/energye/energy/v2/pkgs/i18n"
	"github.com/energye/energy/v2/pkgs/macapp"
	"github.com/energye/energy/v2/tools"
	"github.com/energye/energy/v2/types"
	"os"
	"path"
	"runtime"
)

const (
	emfsLibsPath = "libs"
)

// Init 全局初始化
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	emfs.SetEMFS(libs, resources)
	if libname.LibName == "" {
		if runtime.GOOS == "darwin" {
			//MacOSX从Frameworks加载
			libname.LibName = "@executable_path/../Frameworks/" + libname.GetDLLName()
		} else {
			libname.LibName = libname.LibPath(libname.GetDLLName())
		}
		if libname.LibName == "" {
			libname.LibName = path.Join(types.HomeGoLCLDir, libname.GetDLLName())
			//liblcl都没有的情况, 最后尝试在内置libs中获取-并释放到用户目录
			releaseLib(path.Join(emfsLibsPath, libname.GetDLLName()), libname.LibName)
			if !tools.IsExist(libname.LibName) {
				println(`Hint:
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
		}
	}
	InitAll()
}

// 释放文件
//
//	如果liblcl动态库内置到EXE中, 在EXE中把liblcl释放到out目录
func releaseLib(fsPath, out string) {
	if emfs.GetLibsFS() != nil {
		// 尝试创建目录, 如果目录已存在则不会创建
		tools.MkdirAll(types.HomeGoLCLDir)
		var liblcl, err = emfs.GetLibsFS().ReadFile(fsPath)
		if err == nil {
			var file *os.File
			file, err = os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			file.Write(liblcl)
		}
	}
}

func InitAll() {
	//macapp
	macapp.MacApp.Init()
	//api
	api.APIInit()
	//rtl
	rtl.RtlInit()
	//version
	version.VersionInit()
	//zh_cn
	zh_CN.ZH_CNInit()
	//win
	winInit()
	//lcl
	lcl.LCLInit()
	//i18n
	i18n.I18NInit()

}
