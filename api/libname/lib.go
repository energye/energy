//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package libname

import (
	"github.com/energye/energy/v2/tools"
	"github.com/energye/energy/v2/types"
	"os"
	"path"
	"runtime"
)

const libName = "liblcl"

var (
	LibName          = ""
	tempDllDir       []string
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

// SetTempDllDir 设置内置dll释放后加载目录
func SetTempDllDir(dir string) {
	tempDllDir = append(tempDllDir, dir)
}

func GetDLLName() string {
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}

// LibPath
//	获取 lib 动态库目录
//  优先级
//	  tempDllDir > 当前目录 > 用户目录 > LCL_HOME > ENERGY_HOME
func LibPath(dllName string) string {
	//tempdll内置目录
	for _, dir := range tempDllDir {
		var tempdllPathLibName = path.Join(dir, dllName)
		if tools.IsExist(tempdllPathLibName) {
			return tempdllPathLibName
		}
	}
	//当前执行文件目录
	var currentPathLibName = path.Join(types.ExeDir, dllName)
	if tools.IsExist(currentPathLibName) {
		return currentPathLibName
	}
	//当前用户golcl目录
	var homePathLibName = path.Join(types.HomeGoLCLDir, dllName)
	if tools.IsExist(homePathLibName) {
		return homePathLibName
	}
	//环境变量 LCL_HOME - 为区分单独使用GO LCL
	var envPathLibName = path.Join(os.Getenv("LCL_HOME"), dllName)
	if tools.IsExist(envPathLibName) {
		return envPathLibName
	}
	//环境变量 ENERGY_HOME -  在使用Energy框架
	var energyPathLibName = path.Join(os.Getenv("ENERGY_HOME"), dllName)
	if tools.IsExist(energyPathLibName) {
		return energyPathLibName
	}
	return ""
}
