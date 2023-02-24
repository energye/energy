package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp = cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/process-message.html"
	cef.BrowserWindow.Config.Title = "Energy - execute-javascript"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})

	cefApp.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		handler := cef.CreateCefV8Handler()
		fmt.Println("handler:", handler)
		handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.Exception) bool {
			fmt.Println("handler.Execute", name)
			retVal.SetResult(cef.V8ValueRef.NewString("函数返回值？"))
			return true
		})
		object := cef.V8ValueRef.NewObject(nil, nil)
		function := cef.V8ValueRef.NewFunction("testfn", handler)
		object.SetValueByKey("testfn", function, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		context.Global.SetValueByKey("testset", object, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
		return false
	})

	//运行应用
	cef.Run(cefApp)
}
