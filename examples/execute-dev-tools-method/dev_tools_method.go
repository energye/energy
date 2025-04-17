package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/common/utils"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

// 资源目录，内置到执行程序中
//
//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	cefApp.SetTouchEvents(consts.STATE_ENABLED)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/execute-dev-tool-method.html"
	cef.BrowserWindow.Config.ChromiumConfig().SetEnableDevTools(false)

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//这里演示使用ipc通信实现js和go互相调用，在go监听事件中执行开发者工具方法
		//使用内置http服务和自定义页面
		//这里执行的方法是仿真移动端
		//1. js使用ipc.emit触发 go事件
		//2. go中"execute-dev-method"事件执行，通过context获得browserId
		//3. 通过browserId获得chromium
		//4. 使用字典对象传递方法参数
		//5. 点击Note链接

		// 全局变量，自动递增的消息ID
		var gId int32 = 0
		ipc.On("execute-dev-method", func(context context.IContext) {
			//获得当前窗口信息
			info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			//字典对象
			var dict = cef.DictionaryValueRef.New() // cef.NewCefDictionaryValue()
			//根据chromium字典设置
			//dict.SetInt("width", 500)
			//dict.SetInt("height", 768)
			//dict.SetInt("x", 100)
			//dict.SetInt("y", 100)
			dict.SetBool("mobile", true)
			dict.SetDouble("deviceScaleFactor", 1)
			TempDict := cef.DictionaryValueRef.New()
			TempDict.SetString("type", "portraitPrimary")
			TempDict.SetInt("angle", 0)
			dict.SetDictionary("screenOrientation", TempDict)
			gId = info.Chromium().ExecuteDevToolsMethod(gId, "Emulation.setDeviceMetricsOverride", dict)
			fmt.Println("ExecuteDevToolsMethod - result messageId:", gId)
			//设置浏览器 userAgent
			dict = cef.DictionaryValueRef.New()
			dict.SetString("userAgent", "Mozilla/5.0 (Linux; Android 11; M2102K1G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Mobile Safari/537.36")
			gId = info.Chromium().ExecuteDevToolsMethod(gId, "Emulation.setUserAgentOverride", dict)
			fmt.Println("ExecuteDevToolsMethod - result messageId:", gId)
			fmt.Println()

			// 更多devtools: https://chromedevtools.github.io/devtools-protocol/
		})
		// 使用  DevToolsRawMessage  处理，方便些
		window.Chromium().SetOnDevToolsRawMessage(func(sender lcl.IObject, browser *cef.ICefBrowser, message uintptr, messageSize uint32) (handled bool) {
			fmt.Println("OnDevToolsRawMessage message:", message, messageSize)
			data := utils.ReadData(message, messageSize)
			fmt.Println("data:", string(data))
			return false
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
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
