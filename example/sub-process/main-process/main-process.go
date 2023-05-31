package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/cef/process"
	"github.com/energye/energy/common"
	"github.com/energye/energy/example/sub-process/main-process/src"
	"github.com/energye/energy/logger"
	"github.com/energye/energy/pkgs/assetserve"
	"github.com/energye/golcl/pkgs/macapp"
	"os"
	"path"
)

//go:embed resources
var resources embed.FS

/*
主进程
这个示例演示了 主进程和 子进程相互独立出来，
步骤
 1. 先编译好子进程程序
    cd sub-process
    go build
    sub-process.exe
 2. 将子进程执行文件（sub-process.exe）在主进程SetBrowseSubprocessPath配置，如果在 FrameworkDirPath 可以直接写文件名
 3. 运行主程序
*/
func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
	//MacOS通过指定 IsCEF ，在开发环境中自动生成可运行的程序包
	//MacOS配置要在 GlobalInit 它之前
	//特别说明MacOS：子进程不需要配置
	wd, _ := os.Getwd()
	if common.IsDarwin() {
		//主进程中 主子进程方式，在这里指定子进程的执行文件
		subExePath := path.Join(wd, "example", "sub-process", "sub-process", "sub-process")
		macapp.MacApp.SetBrowseSubprocessPath(subExePath)
	}
	process.Args.ProcessType()
	//CEF全局初始化
	cef.GlobalInit(nil, &resources)
	//Cef应用的配置 执行程序如果在 chromium 目录中可不配置
	//子进程执行程序如果在 chromium 目录中可不配置
	var subExePath string
	if common.IsWindows() {
		subExePath = path.Join(wd, "example", "sub-process", "sub-process", "sub-process.exe")
	} else if common.IsLinux() {
		subExePath = path.Join(wd, "example", "sub-process", "sub-process", "sub-process")
	} else if common.IsDarwin() {
		//MacOS SetBrowseSubprocessPath 将不起任何作用。
		//独立的子程序包需要在 macapp.MacApp.SetBrowseSubprocessPath 配置
	}
	//subExePath = path.Join(wd, "sub-process.exe")
	println("subExePath", subExePath)
	//创建Cef应用
	cefApp := cef.NewApplication()
	//cefApp.SetSingleProcess(false) //单进程 或 多进程 ,单进程上面的子进程配置就不起作用了
	cefApp.SetBrowserSubprocessPath(subExePath)
	//主进程初始化
	src.MainBrowserInit()
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
