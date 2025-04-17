package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/many-browser/form"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"path/filepath"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	// 定义在当前执行目录缓存
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	//创建应用
	app := cef.NewApplication()
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetEnableGPU(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com/"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 仅在LCL窗口时创建
		if window.IsLCL() {
			form.SetRootCachePath(rootCache)
			//动态创建窗口组件
			form.CreateComponent(window)
		}
	})
	//运行应用
	cef.Run(app)
}
