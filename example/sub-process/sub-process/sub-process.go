package main

import (
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/example/sub-process/sub-process/src"
	"github.com/energye/energy/example/sub-process/vars"
)

/*
子进程
这个示例演示了 主进程和 子进程相互独立出来，
子进程需要先编译好,提供给主进程(SetBrowseSubprocessPath)配置
*/
func main() {
	//全局配置初始化
	cef.GlobalCEFInit(nil, nil)
	//IPC通信
	src.IPCInit()
	//Cef应用的配置 执行程序如果在 chromium 目录中可不配置
	cfg := cef.NewApplicationConfig()
	//配置chromium frameworks 编译好的二进制目录
	if common.IsWindows() {
		cfg.SetFrameworkDirPath("E:\\SWT\\CEF4Delphi-Libs-103.0.9\\chromium")
	} else if common.IsLinux() {
		cfg.SetFrameworkDirPath("/home/sxm/app/swt/CEFDelphi-Libs-103.09/chromium")
	}
	//创建Cef应用
	cefApp := cef.NewApplication(cfg)
	//主进程和子进程的变量绑定函数定义
	vars.VariableBind()
	//启动子进程
	cefApp.StartSubProcess()
	cefApp.Free()
}
