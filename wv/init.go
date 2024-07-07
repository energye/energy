//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import (
	"github.com/energye/lcl/emfs"
	"github.com/energye/wv/wv"
)

// Init 全局初始化, 需手动调用的函数
//
//	参数:
//	   libs 内置到应用程序的类库
//	   resources 内置到应用程序的资源文件
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	wv.Init(libs, resources)
}
