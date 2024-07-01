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
	"github.com/energye/lcl/process"
)

// Init 全局初始化, 需手动调用的函数
//
//	参数:
//	   libs 内置到应用程序的类库
//	   resources 内置到应用程序的资源文件
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	cef.Init(libs, resources)
}

// v8init v8初始化
func v8init() {
	if process.Args.IsMain() || process.Args.IsRender() {
		//ipc初始化
		ipcInit()
	}
}
