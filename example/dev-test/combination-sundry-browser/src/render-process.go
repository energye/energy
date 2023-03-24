//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/cef/ipc"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"strings"
)

func AppRenderInit() *cef.TCEFApplication {
	//Cef应用的配置
	var env = common.Args.Args("env")
	if common.IsWindows() {
		//SetFrameworkDirPath 或 配置环境变量 ENERGY_HOME
	} else if common.IsLinux() {
		//cfg.SetFrameworkDirPath("/home/sxm/app/swt/energy/chromium")  或 配置环境变量 ENERGY_HOME
	}
	//创建Cef应用
	cefApp := cef.NewApplication()
	//fmt.Printf("cefApp:%+v %s\n", cefApp, runtime.GOOS)
	if common.Args.IsMain() {
		//cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.TCefCommandLine) {
		//	//主进程 自定义参数
		//	fmt.Println("======================OnBeforeChildProcessLaunch 定义进程参数: ", cef.Args.ProcessType())
		//	commandLine.AppendSwitch("env", env)
		//})
	} else if common.Args.IsRender() {
		//取出主进程 自定义参数

		fmt.Println("ipc-port", common.Args.Args("net-ipc-port"), common.Args.ProcessType())
	}
	cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.ICefCommandLine) {
		//主进程 自定义参数
		fmt.Println("======================OnBeforeChildProcessLaunch 定义进程参数: ", common.Args.ProcessType())
		commandLine.AppendSwitchWithValue("env", env)
		commandLine.AppendArgument("--test")
	})
	//上下文回调
	cefApp.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		fmt.Println("SetOnContextCreated browserId:", browser.Identifier())
		//判断url地址 运行IPC 和 变量绑定
		if strings.LastIndex(strings.ToLower(frame.Url()), ".pdf") > 0 || strings.Index(frame.Url(), "about:blank") != -1 {
			return true //返回 true 时，不运行IPC 和 变量绑定
		}
		return false //返回 false 时，运行IPC 和 变量绑定
	})
	//渲染进程的消息处理
	cefApp.SetOnProcessMessageReceived(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess consts.CefProcessId, message *cef.ICefProcessMessage) bool {
		fmt.Println("======================渲染进程 OnProcessMessageReceived IPC browserId:", browser.Identifier(), "frameId:", frame.Identifier(), "sourceProcess:", sourceProcess, "processMessage.Name:", message.Name)

		return false
	})
	//渲染进程 IPC事件
	ipc.On("renderOnEventSubWindowIPCOn", func(context ipc.IContext) {
		fmt.Println("渲染进程监听事件-执行 renderOnEventSubWindowIPCOn", common.Args.ProcessType())
		//渲染进程处理程序....
		context.Result("返回了,可以关闭")
	})
	//渲染进程 IPC事件
	return cefApp
}
