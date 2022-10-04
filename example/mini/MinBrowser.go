package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/example/mini/src"
	"github.com/energye/golcl/pkgs/macapp"
)

//go:embed libs
var libs embed.FS

//go:embed resources
var resources embed.FS

func main() {
	cef.Logger.SetLevel(cef.CefLog_Debug)
	if cef.IsDarwin() {
		//libname.LibName = "/Users/zhangli/go/bin/liblcl.dylib"
		macapp.MacApp.IsCEF(true)
		macapp.MacApp.SetBaseCefFrameworksDir("/Users/zhangli/app/swt-lazarus/CEF4Delphi-Libs-105.3.39/chromium")
	}
	fmt.Println("process- main:", cef.Args.IsMain(), "render:", cef.Args.IsRender())
	cef.GlobalCEFInit(&libs, &resources)

	//手动选择 IPC 通信通道
	cef.UseNetIPCChannel = true

	//Render 子进程一些初始化配置
	cefApp := src.AppRenderInit()
	//Browser 主进程一些初始配置
	src.AppBrowserInit()
	cef.Run(cefApp)
}
