package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/pkgs/macapp"
)

//go:embed resources
var resources embed.FS

//go:embed libs
var libs embed.FS

//这是一个简单的窗口创建示例
func main() {
	//开发环境中 MacOSX平台必须在"GlobalCEFInit"之前设置CEF
	//设置使用CEF 和 CEF框架目录，生成开发执行应用程序包
	if common.IsDarwin() {
		macapp.MacApp.IsCEF(true)
		macapp.MacApp.SetBaseCefFrameworksDir("/Users/zhangli/app/swt/energy/chromium")
	}
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(&libs, &resources)
	//可选的应用配置
	cfg := cef.NewApplicationConfig() //指定chromium的二进制包框架根目录,
	//不指定为当前程序执行目录
	if common.IsWindows() {
		cfg.SetFrameworkDirPath("D:\\app.exe\\energy\\105.0.5195.127\\dev\\chromium-64")
	} else if common.IsLinux() {
		cfg.SetFrameworkDirPath("/home/sxm/app/swt/energy/chromium")
	}
	//创建应用
	cefApp := cef.NewApplication(cfg)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "https://www.baidu.com/"
	//窗口的标题
	cef.BrowserWindow.Config.Title = "energy - 这是一个简单的窗口示例"
	//窗口宽高
	cef.BrowserWindow.Config.Width = 1024
	cef.BrowserWindow.Config.Height = 768
	//chromium配置
	cef.BrowserWindow.Config.SetChromiumConfig(cef.NewChromiumConfig())
	//创建窗口时的回调函数 对浏览器事件设置，和窗口属性组件等创建和修改
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		//设置应用图标 这里加载的图标是内置到执行程序里的资源文件
		lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")
		fmt.Println("SetBrowserInit")
	})
	//创建窗口之后对对主窗口的属性、组件或子窗口的创建
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInitAfter")
	})
	//内置http服务链接安全配置
	assetserve.AssetsServerHeaderKeyName = "energy"
	assetserve.AssetsServerHeaderKeyValue = "energy"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务 然后使用 http://localhost/资源目录或资源名称")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})

	//运行应用
	cef.Run(cefApp)
}
