package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/example/mini-browser/src"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/pkgs/macapp"
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
	//开发环境中 MacOSX平台必须在"GlobalCEFInit"之前设置CEF
	//设置使用CEF 和 CEF框架目录，生成开发执行应用程序包
	macapp.MacApp.IsCEF(common.IsDarwin())
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
