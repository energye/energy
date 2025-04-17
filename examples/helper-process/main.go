package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/helper-process/app"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/pkgs/macapp"
	"path/filepath"
)

//go:embed resources
var resources embed.FS

/*
主进程
这个示例演示了 主进程和 子进程相互独立出来，
步骤
 1. 先编译好子进程程序
    helper
    go build
 2. 将helper子进程执行文件（helper）在主进程SetBrowseSubprocessPath配置，如果在 CEF框架目录 可直接写文件名
 3. 运行主程序
*/
func main() {
	//MacOS通过指定 IsCEF ，在开发环境中自动生成可运行的程序包
	//MacOS配置要在调用 GlobalInit 它之前设置
	//MacOS：子(helper)进程不需要配置 app.SetBrowseSubprocessPath(),
	//       MacOS如需要单独的 helper 进程通过 macapp.MacApp.SetBrowseSubprocessPath 生成 Mac App特定的helper进程
	wd := consts.CurrentExecuteDir
	if common.IsDarwin() {
		// 主进程中 主子进程方式，在这里指定子进程的执行文件
		subExePath := filepath.Join(wd, "examples", "helper-process", "helper", "helper")
		macapp.MacApp.SetBrowseSubprocessPath(subExePath)
	}
	//CEF全局初始化
	cef.GlobalInit(nil, resources)
	if !common.IsDarwin() {
		// 使用 SetBrowserSubprocessPath 设置子进程执行文件目录
		var subName = subExeName()
		// 子进程执行文件完整目录
		var subExePath = filepath.Join(wd, "helper", subName)
		println("subExePath", subExePath)
		// 设置子进程执行文件目录
		// 子进程执行文件如果在同一目录并且在CEF框架目录可只写文件名
		app.GetApplication().SetBrowserSubprocessPath(subExePath)
	}
	//主进程初始化
	mainBrowserInit()
	//启动内置http服务
	startHttpServer()

	cef.Run(app.GetApplication())
}

func subExeName() (name string) {
	if common.IsWindows() {
		name = "helper.exe"
	} else if common.IsLinux() {
		name = "helper"
	} else if common.IsDarwin() {
		//MacOS SetBrowseSubprocessPath 不起任何作用。
		//MacOS 独立的子程序包需要在 macapp.MacApp.SetBrowseSubprocessPath 配置
	}
	return
}

// 启动内置http服务
func startHttpServer() {
	fmt.Println("主进程启动 创建一个内置http服务")
	//通过内置http服务加载资源
	server := assetserve.NewAssetsHttpServer()
	server.PORT = 22022
	server.AssetsFSName = "resources" //必须设置目录名
	server.Assets = resources
	go server.StartHttpServer()
}

// 主进程浏览器初始化
func mainBrowserInit() {
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.Title = "ENERGY 区分主/子进程执行文件"
	//主窗口初始化回调函数
	//在这个函数里，主进程浏览初始化之前创建窗口之后
	//在这里可以设置窗口的属性和事件监听
	//SetOnBeforePopup是子进程弹出窗口时触发，可以改变子进程窗口属性
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		fmt.Println("主窗口初始化回调函数")
		window.SetCenterWindow(true)

		event.SetOnBeforePopup(func(sender lcl.IObject, popupWindow cef.IBrowserWindow, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup: " + beforePopupInfo.TargetUrl)
			popupWindow.SetTitle("改变了标题 - " + beforePopupInfo.TargetUrl)
			return false
		})
	})

}
