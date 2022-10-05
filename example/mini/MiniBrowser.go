package main

import (
	"embed"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/commons"
	"github.com/energye/energy/example/mini/src"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/pkgs/macapp"
)

//go:embed libs
var libs embed.FS

//go:embed resources
var resources embed.FS

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
	if commons.IsDarwin() {
		//libname.LibName = "/Users/zhangli/go/bin/liblcl.dylib"
		macapp.MacApp.IsCEF(true)
		macapp.MacApp.SetBaseCefFrameworksDir("/Users/zhangli/app/swt-lazarus/CEF4Delphi-Libs-105.3.39/chromium")
	}
	cef.GlobalCEFInit(&libs, &resources)
	//Render 子进程一些初始化配置
	cefApp := src.AppRenderInit()
	//Browser 主进程一些初始配置
	src.AppBrowserInit()
	cef.Run(cefApp)
}
