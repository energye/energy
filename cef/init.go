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
	"github.com/energye/lcl/process"
)

// v8init v8初始化
func v8init() {
	if process.Args.IsMain() || process.Args.IsRender() {
		//ipc初始化
		ipcInit()
	}
}
