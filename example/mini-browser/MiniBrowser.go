package main

import (
	"embed"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/commons"
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
	if commons.IsDarwin() {
		//libname.LibName = "/Users/zhangli/go/bin/liblcl.dylib"
		macapp.MacApp.IsCEF(true)
		macapp.MacApp.SetBaseCefFrameworksDir("/Users/zhangli/app/swt/energy/chromium")
	}
	cef.GlobalCEFInit(&libs, &resources)
	//Render 子进程一些初始化配置
	cefApp := src.AppRenderInit()
	//Browser 主进程一些初始配置
	src.AppBrowserInit()
	cef.Run(cefApp)
}
