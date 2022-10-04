package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"strings"
)

func AppRenderInit() *cef.TCEFApplication {
	//Cef应用的配置
	cfg := cef.NewApplicationConfig()
	var env = cef.Args.Args("env")
	if cef.IsWindows() {
		if env == "32" {
			cfg.SetFrameworkDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-32")
			cfg.SetResourcesDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-32")
			cfg.SetLocalesDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-32\\locales")
			cfg.SetUserDataPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-32\\locales")
			cfg.SetCache("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-32\\UserData")
			cfg.SetLogFile("E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\debug.log")
		} else {
			cfg.SetFrameworkDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-64")
			cfg.SetResourcesDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-64")
			cfg.SetLocalesDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-64\\locales")
			cfg.SetUserDataPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-64\\locales")
			cfg.SetCache("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-64\\UserData")
			cfg.SetLogFile("E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\debug.log")
		}
	} else if cef.IsLinux() {
		cfg.SetFrameworkDirPath("/home/sxm/app/swt/CEF4Delphi-Libs-105.3.39/chromium")
		cfg.SetResourcesDirPath("/home/sxm/app/swt/CEF4Delphi-Libs-105.3.39/chromium")
		cfg.SetLocalesDirPath("/home/sxm/app/swt/CEF4Delphi-Libs-105.3.39/chromium/locales")
		cfg.SetUserDataPath("/home/sxm/app/swt/CEF4Delphi-Libs-105.3.39/chromium/locales")
		cfg.SetCache("/home/sxm/app/swt/CEF4Delphi-Libs-105.3.39/chromium/UserData")
		cfg.SetLogFile("/home/sxm/app/swt/gopath/src/swt-lazarus/demo17-dll-load/debug.log")
	}
	//cfg.SetLogSeverity(cef.LOGSEVERITY_DEBUG)
	cfg.SetLogSeverity(cef.LOGSEVERITY_DISABLE)
	cfg.SetLanguage(cef.LANGUAGE_zh_CN)
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
	if cef.Args.IsMain() {
		//cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.TCefCommandLine) {
		//	//主进程 自定义参数
		//	fmt.Println("======================OnBeforeChildProcessLaunch 定义进程参数: ", cef.Args.ProcessType())
		//	commandLine.AppendSwitch("env", env)
		//})
	} else if cef.Args.IsRender() {
		//取出主进程 自定义参数

		fmt.Println("ipc-port", cef.Args.Args("net-ipc-port"), cef.Args.ProcessType())
	}
	cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.TCefCommandLine) {
		//主进程 自定义参数
		fmt.Println("======================OnBeforeChildProcessLaunch 定义进程参数: ", cef.Args.ProcessType())
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
	cefApp.SetOnProcessMessageReceived(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess cef.CefProcessId, message *cef.ICefProcessMessage) bool {
		fmt.Println("======================渲染进程 OnProcessMessageReceived IPC browserId:", browser.Identifier(), "frameId:", frame.Id, "sourceProcess:", sourceProcess, "processMessage.Name:", message.Name)
		fmt.Println("\t|--Args:", cef.Args.ProcessType(), "message:", message.ArgumentList.GetString(0))
		message = cef.NewProcessMessage("test")
		message.ArgumentList.SetString(0, "渲染进程发送数据")
		frame.SendProcessMessage(cef.PID_BROWSER, message)
		message.ArgumentList.Clear()
		return false
	})
	//cefApp.SetOnLoadStart(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType cef.TCefTransitionType) {
	//	fmt.Println("====================== cefApp.SetOnLoadStart ", browser.Identifier(), frame.Id, transitionType)
	//})
	//cefApp.SetOnLoadEnd(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
	//	fmt.Println("====================== cefApp.SetOnLoadEnd ", browser.Identifier(), frame.Id, httpStatusCode)
	//})
	//cefApp.SetOnBrowserDestroyed(func(browser *cef.ICefBrowser) {
	//	fmt.Println("====================== cefApp.SetOnBrowserDestroyed ", browser.Identifier())
	//})

	//渲染进程 IPC事件
	cef.IPC.Render().SetOnEvent(func(event cef.IEventOn) {
		fmt.Println("渲染进程IPC事件注册")
		//渲染进程监听的事件
		event.On("renderOnEventSubWindowIPCOn", func(context cef.IIPCContext) {
			fmt.Println("render renderOnEventSubWindowIPCOn")
			//渲染进程处理程序....
			context.Response([]byte("返回了,可以关闭"))
		})
	})
	return cefApp
}
