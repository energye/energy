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
	_ "github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/cef/internal/initialize"
	. "github.com/cyber-xxm/energy/v2/cef/process"
	. "github.com/cyber-xxm/energy/v2/common"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl"
	"os"
)

// GlobalInit 全局初始化
//
// 需要手动调用的函数,在main函数中调用
//
// 参数:
// libs 内置到应用程序的类库
// resources 内置到应用程序的资源文件
func GlobalInit(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	initialize.Initialize(libs, resources)
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
