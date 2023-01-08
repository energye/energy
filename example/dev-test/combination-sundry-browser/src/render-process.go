package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"strings"
)

func AppRenderInit() *cef.TCEFApplication {
	//Cef应用的配置
	cfg := cef.NewApplicationConfig()
	var env = common.Args.Args("env")
	if common.IsWindows() {
		//SetFrameworkDirPath 或 配置环境变量 ENERGY_HOME
	} else if common.IsLinux() {
		//cfg.SetFrameworkDirPath("/home/sxm/app/swt/energy/chromium")  或 配置环境变量 ENERGY_HOME
	}
	cfg.SetLogSeverity(consts.LOGSEVERITY_DEBUG)
	//cfg.SetLogSeverity(consts.LOGSEVERITY_DISABLE)
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
	cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.TCefCommandLine) {
		//主进程 自定义参数
		fmt.Println("======================OnBeforeChildProcessLaunch 定义进程参数: ", common.Args.ProcessType())
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
		fmt.Println("\t|--Args:", common.Args.ProcessType(), "message:", message.ArgumentList.GetString(0))
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
			fmt.Println("渲染进程监听事件-执行 renderOnEventSubWindowIPCOn", common.Args.ProcessType())
			//渲染进程处理程序....
			context.Response([]byte("返回了,可以关闭"))
		})
	})
	return cefApp
}
