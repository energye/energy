package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/examples/helper-process/app"
)

/*
helper 子进程
子进程需要先编译好得到执行文件, 提供给主进程(SetBrowseSubprocessPath)配置
*/
func main() {
	//全局配置初始化
	cef.GlobalInit(nil, nil)
	//创建Cef应用
	application := app.GetApplication()
	// 渲染进程监听IPC事件
	helperLisIPC()
	//启动子进程
	application.StartSubProcess()
	application.Free()
}

// 渲染进程监听IPC事件
func helperLisIPC() {
	fmt.Println("渲染进程IPC事件注册 ProcessType:", process.Args.ProcessType())
	//渲染进程监听的事件
	ipc.On("sub-process-on-event", func(context context.IContext) {
		fmt.Println("sub-process-on-event")
		//渲染进程处理程序....
		context.Result("返回结果")
	})
}
