package main

import (
	"embed"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/example/browser-control/src"
)

//go:embed resources
var resources embed.FS

//go:embed libs
var libs embed.FS

func main() {
	//开发环境中 MacOSX平台必须在"GlobalInit"之前设置CEF
	//设置使用CEF 和 CEF框架目录，生成开发执行应用程序包
	//环境变量 ENERGY_HOME="/app/cefframework" 配置框架所在目录
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(&libs, &resources)
	//可选的应用配置
	cfg := cef.NewApplicationConfig()
	//指定chromium的二进制包框架根目录,
	//不指定为当前程序执行目录
	if common.IsWindows() {
		//SetFrameworkDirPath 或 配置环境变量 ENERGY_HOME
		//cfg.SetFrameworkDirPath("D:\\app.exe\\energy\\chromium64")
	} else if common.IsLinux() {
		//cfg.SetFrameworkDirPath("/home/sxm/app/swt/energy/chromium")
	}
	//创建应用
	cefApp := cef.NewApplication(cfg)
	//主进程窗口
	src.MainBrowserWindow()

	//运行应用
	cef.Run(cefApp)
}
