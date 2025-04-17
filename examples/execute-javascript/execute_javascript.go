package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/execute-javascript.html"
	cef.BrowserWindow.Config.Title = "Energy - execute-javascript"
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
		//定时执行web js
		go timeTask()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		window.Chromium().SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
			var jsCode = `
(function(){
	console.log("执行");
})();
`
			window.Chromium().ExecuteJavaScript(jsCode, "", frame, 0)
		})
	})
	//运行应用
	cef.Run(cefApp)
}

// 定时执行web js
func timeTask() {
	var param0 = 0
	for {
		//每1秒钟执行一次
		time.Sleep(time.Second)
		info := cef.BrowserWindow.MainWindow()
		param0++
		//调用js中定义的函数GoExecuteJSFunc,并传递参数，但没有返回值
		var jsFunc = fmt.Sprintf("GoExecuteJSFunc(%d, '%d')", param0, time.Now().Second())
		fmt.Println("GoExecuteJSFunc:", jsFunc)
		info.Chromium().ExecuteJavaScript(jsFunc, "", info.Chromium().Browser().MainFrame(), 0)
	}
}
