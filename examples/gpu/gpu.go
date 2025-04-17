//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	// 启用GPU加速, 默认 energy 未开启GPU加速
	// 使用测试网址 www.antutu.com/html5 在任务管理器中观察 CPU 和 GPU 使用情况
	app.SetEnableGPU(true)
	cef.BrowserWindow.Config.Url = "https://www.antutu.com/html5"
	//cef.BrowserWindow.Config.Url = "https://threejs.org" // WebGL
	//cef.BrowserWindow.Config.Url = "https://ice-gl.gitee.io/icegl-three-vue-tres"
	cef.BrowserWindow.Config.Title = "ENERGY GPU Test"
	//运行应用
	cef.Run(app)
}
