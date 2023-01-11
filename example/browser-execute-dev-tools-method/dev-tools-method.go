package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/ipc"
)

//资源目录，内置到执行程序中
//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/execute-dev-tool-method.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	//chromium配置
	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)     //启用右键菜单
	config.SetEnableDevTools(true) //启用开发者工具
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//这里演示使用ipc通信实现js和go互相调用，在go监听事件中执行开发者工具方法
	//使用内置http服务和自定义页面
	//这里执行的方法是仿真移动端
	//1. js使用ipc.emit触发 go事件
	//2. go中"execute-dev-method"事件执行，通过context获得browserId
	//3. 通过browserId获得chromium
	//4. 使用字典对象传递方法参数
	//5. 点击Note链接
	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		event.On("execute-dev-method", func(context ipc.IIPCContext) {
			//获得当前窗口信息
			info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			//字典对象
			var dict = cef.NewCefDictionaryValue()
			//根据chromium字典设置
			dict.SetInt("width", 500)
			dict.SetInt("height", 768)
			dict.SetInt("x", 100)
			dict.SetInt("y", 100)
			dict.SetBoolean("mobile", true)
			dict.SetDouble("deviceScaleFactor", 1)
			TempDict := cef.NewCefDictionaryValue()
			TempDict.SetString("type", "portraitPrimary")
			TempDict.SetInt("angle", 0)
			dict.SetDictionary("screenOrientation", TempDict)
			info.Chromium().ExecuteDevToolsMethod(0, "Emulation.setDeviceMetricsOverride", dict)
			//设置浏览器 userAgent
			dict = cef.NewCefDictionaryValue()
			dict.SetString("userAgent", "Mozilla/5.0 (Linux; Android 11; M2102K1G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Mobile Safari/537.36")
			info.Chromium().ExecuteDevToolsMethod(0, "Emulation.setUserAgentOverride", dict)
		})
	})
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
