package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/example/mini-browser/src"
	"github.com/energye/energy/logger"
)

//go:embed libs
var libs embed.FS

//go:embed resources
var resources embed.FS

//GOOS=windows;GOARCH=386;
//env=32
func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
	//环境变量 ENERGY_HOME="/app/cefframework" 配置框架所在目录
	//全局初始化
	cef.GlobalCEFInit(&libs, &resources)
	//Render 子进程一些初始化配置
	cefApp := src.AppRenderInit()
	//Browser 主进程一些初始配置
	src.AppBrowserInit()
	//内置http服务链接安全配置
	assetserve.AssetsServerHeaderKeyName = "energy"
	assetserve.AssetsServerHeaderKeyValue = "energy"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	cef.Run(cefApp)
}
