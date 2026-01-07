//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
)

var beforeRun []func()

func SetOnBeforeRun(fn func()) {
	beforeRun = append(beforeRun, fn)
}

func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	lcl.Init(libs, resources)
}

func Run(forms ...lcl.IEngForm) {
	// 初始化应用程序实例
	lcl.Application.Initialize()
	// 配置应用程序设置，使主窗体在Windows任务栏上显示
	lcl.Application.SetMainFormOnTaskBar(true)
	// 启用自动缩放功能以支持高DPI显示器
	lcl.Application.SetScaled(true)
	for _, fn := range beforeRun {
		fn()
	}
	// 创建所有窗体
	lcl.Application.NewForms(forms...)
	// 启动应用程序消息循环
	lcl.Application.Run()
}

func Application() lcl.IApp {
	return lcl.Application
}
