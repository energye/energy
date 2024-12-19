// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

// Package Application environment initialization loads dynamic libraries

package initialize

import (
	"github.com/energye/golcl/energy/emfs"
)

// Initialize
// 初始化，运行时加载 LibLCL
func Initialize(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	loadLibLCL(libs, resources)
}
