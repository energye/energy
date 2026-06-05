//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/cef"
	_ "github.com/energye/cef/cef"
	"github.com/energye/lcl/lcl"
)

// Init 全局初始化, 需手动调用的函数
func Init() *Application {
	lcl.Init()
	cef.Init()
	return NewApplication()
}
