package main

import (
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/example/sub-process/main-process/src"
	"github.com/energye/energy/example/sub-process/vars"
	"github.com/energye/golcl/pkgs/macapp"
)

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
	//MacOS通过指定 IsCEF ，在开发环境中自动生成可运行的程序包
	//MacOS配置要在 GlobalCEFInit 它之前
	//特别说明MacOS：子进程不需要配置
	if common.IsDarwin() {
		//自动生成mac app程序包
		macapp.MacApp.IsCEF(true)
		//IsCEF=True 必须指定Frameworks目录
		macapp.MacApp.SetBaseCefFrameworksDir("/Users/zhangli/app/swt-lazarus/cef_binary_105.3.39+g2ec21f9+chromium-105.0.5195.127_macosx64/Release")
		//主进程中 主子进程方式，在这里指定子进程的执行文件
		macapp.MacApp.SetBrowseSubprocessPath("/Users/zhangli/go/src/github.com/energye/energy/demos/demo-sub-process/sub-process/sub-process")
	}
	//CEF全局初始化
	cef.GlobalCEFInit(nil, nil)
	//Cef应用的配置 执行程序如果在 chromium 目录中可不配置
	cfg := cef.NewApplicationConfig()
	//配置chromium frameworks 编译好的二进制目录
	if common.IsWindows() {
		cfg.SetFrameworkDirPath("E:\\SWT\\CEF4Delphi-Libs-103.0.9\\chromium")
	} else if common.IsLinux() {
		cfg.SetFrameworkDirPath("/home/sxm/app/swt/CEFDelphi-Libs-103.09/chromium")
	}
	//子进程执行程序如果在 chromium 目录中可不配置
	if common.IsWindows() {
		cfg.SetBrowseSubprocessPath("E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\demo-golang-dll-01-chromium\\demos\\demo-sub-process\\sub-process\\sub-process.exe")
	} else if common.IsLinux() {
		cfg.SetBrowseSubprocessPath("/home/sxm/app/swt/gopath/src/github.com/energye/energy/demos/demo-sub-process/sub-process/sub-process")
	} else if common.IsDarwin() {
		//MacOS SetBrowseSubprocessPath 将不起任何作用。
		//独立的子程序包需要在 macapp.MacApp.SetBrowseSubprocessPath 配置
	}
	cfg.SetSingleProcess(false) //单进程 或 多进程 ,单进程上面的子进程配置就不起作用了
	//创建Cef应用
	cefApp := cef.NewApplication(cfg)
	//主进程和子进程的变量绑定函数定义
	vars.VariableBind()
	//主进程浏览器初始化
	src.MainBrowserInit()
	cef.Run(cefApp)
}
