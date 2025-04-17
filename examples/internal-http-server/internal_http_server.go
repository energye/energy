package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
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
	//主窗口的配置
	cef.BrowserWindow.Config.Title = "Energy - 内置资源和内置服务示例"

	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/internal-http-server.html"
	//内置静态资源服务的安全key和value设置
	//通过设置AssetsServerHeaderKeyName和AssetsServerHeaderKeyValue在一定程度上保证资源只能在应用内访问，即使在应用外使用正确的IP和端口号也无法访问到资源
	assetserve.AssetsServerHeaderKeyName = "energy"
	assetserve.AssetsServerHeaderKeyValue = "energy"
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		//LocalAssets 指定本地资源支持热更新 - 适用开发或web端源码可以查看
		//server.LocalAssets = fmt.Sprintf("%s/example/browser-internal-http-server/resources", consts.ExeDir)
		//Assets 内置资源不支持热更新 - 适用应用发布
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
