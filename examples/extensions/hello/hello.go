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
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/examples/extensions/hello/layout"
	_ "github.com/energye/energy/v2/examples/syso"
	"path/filepath"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	//创建应用
	app := cef.NewApplication()
	app.SetEnableGPU(true)
	app.SetLogFile("debug.log")
	//app.SetLogSeverity(consts.LOGSEVERITY_VERBOSE)
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	cef.BrowserWindow.Config.Url = "https://gitee.com/energye/energy"

	//就目前CEF版本，这是一个过时的示例，扩展插件加载无法正常工作
	//按CEF官方说明，扩展插件在ChromeRuntime模式下工作
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			layout.WindowLayout(window)
		}
	})
	//运行应用
	cef.Run(app)
}
