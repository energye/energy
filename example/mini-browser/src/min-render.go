package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/commons"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"strings"
)

func AppRenderInit() *cef.TCEFApplication {
	//Cef应用的配置
	cfg := cef.NewApplicationConfig()
	var env = commons.Args.Args("env")
	if commons.IsWindows() {
		if env == "32" {
			cfg.SetFrameworkDirPath("D:\\app.exe\\energy\\105.0.5195.127\\dev\\chromium-32")
		} else {
			cfg.SetFrameworkDirPath("D:\\app.exe\\energy\\105.0.5195.127\\dev\\chromium-64")
		}
	} else if commons.IsLinux() {
		cfg.SetFrameworkDirPath("/home/sxm/app/swt/energy/chromium")
	}
	//cfg.SetLogSeverity(cef.LOGSEVERITY_DEBUG)
	cfg.SetLogSeverity(consts.LOGSEVERITY_DISABLE)
	cfg.SetLanguage(consts.LANGUAGE_zh_CN)
	cfg.SetEnableGPU(false)
	cfg.SetSingleProcess(false)
	//cfg.SetDisableZygote(true)
	//cfg.SetNoSandbox(true)
	cfg.SetRemoteDebuggingPort(0)
	cfg.SetCommonRootName("v8obj")
	//cef.AddCrDelegate()
	//创建Cef应用
	cefApp := cef.NewApplication(cfg)
	//fmt.Printf("cefApp:%+v %s\n", cefApp, runtime.GOOS)
	if commons.Args.IsMain() {
		//cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.TCefCommandLine) {
		//	//主进程 自定义参数
		//	fmt.Println("======================OnBeforeChildProcessLaunch 定义进程参数: ", cef.Args.ProcessType())
		//	commandLine.AppendSwitch("env", env)
		//})
	} else if commons.Args.IsRender() {
		//取出主进程 自定义参数

		fmt.Println("ipc-port", commons.Args.Args("net-ipc-port"), commons.Args.ProcessType())
	}
	cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.TCefCommandLine) {
		//主进程 自定义参数
		fmt.Println("======================OnBeforeChildProcessLaunch 定义进程参数: ", commons.Args.ProcessType())
		commandLine.AppendSwitch("env", env)
		commandLine.AppendArgument("--test")
	})
	//上下文回调
	cefApp.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		fmt.Println("======================OnContextCreated 渲染进程启动 Args.IsMain:", frame.IsMain(), "browserId:", browser.Identifier(), "frameId:", frame.Id, "frameName:", frame.Name, "frameUrl:", frame.Url)
		//判断url地址 运行IPC 和 变量绑定
		if strings.LastIndex(strings.ToLower(frame.Url), ".pdf") > 0 || strings.Index(frame.Url, "about:blank") != -1 {
			return true //返回 true 时，不运行IPC 和 变量绑定
		}
		return false //返回 false 时，运行IPC 和 变量绑定
	})
	//渲染进程的消息处理
	cefApp.SetOnProcessMessageReceived(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess consts.CefProcessId, message *ipc.ICefProcessMessage) bool {
		fmt.Println("======================渲染进程 OnProcessMessageReceived IPC browserId:", browser.Identifier(), "frameId:", frame.Id, "sourceProcess:", sourceProcess, "processMessage.Name:", message.Name)
		fmt.Println("\t|--Args:", commons.Args.ProcessType(), "message:", message.ArgumentList.GetString(0))
		message = ipc.NewProcessMessage("test")
		message.ArgumentList.SetString(0, "渲染进程发送数据")
		frame.SendProcessMessage(consts.PID_BROWSER, message)
		message.ArgumentList.Clear()
		return false
	})

	//渲染进程 IPC事件
	ipc.IPC.Render().SetOnEvent(func(event ipc.IEventOn) {
		fmt.Println("渲染进程IPC事件注册")
		//渲染进程监听的事件
		event.On("renderOnEventSubWindowIPCOn", func(context ipc.IIPCContext) {
			fmt.Println("render renderOnEventSubWindowIPCOn")
			//渲染进程处理程序....
			context.Response([]byte("返回了,可以关闭"))
		})
	})
	return cefApp
}
