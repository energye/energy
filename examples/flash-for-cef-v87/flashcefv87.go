package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"path/filepath"
)

// 参考并查注释说明
// 该示例运行在CEF版本号87.1.14, liblcl branch 89.0.18
func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://www.ultrasounds.com" // flash 在线测试地址
	eh := config.Get().FrameworkPath()
	if eh == "" {
		eh = consts.ExeDir
	}
	// Energy CEF Flash Demo, 支持 CEF-87.1.14
	// flash 下载地址: https://www.flash.cn/download-wins 选择Adobe Flash Player PPAPI版
	// 在 SetOnBeforeChildProcessLaunch 设置命令行参数
	cefApp.AddCustomCommandLine("ppapi-flash-path", filepath.Join(eh, "pepflashplayer64_34_0_0_289.dll")) // 设置ppapi flash目录
	cefApp.AddCustomCommandLine("ppapi-flash-version", "34.0.0.289")                                      // 版本号和flash dll一样
	cefApp.AddCustomCommandLine("allow-outdated-plugins", "allow")
	cefApp.SetOnBeforeChildProcessLaunch(func(commandLine *cef.ICefCommandLine) {
		eh := config.Get().FrameworkPath()
		if eh == "" {
			eh = consts.ExeDir
		}
		commandLine.AppendSwitch("--allow-outdated-plugins")
		commandLine.AppendSwitch("--disable-web-security")
	})
	// 在浏览器窗口初始化时设置扩展参数
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		intValue := cef.ValueRef.New()
		intValue.SetInt(1)
		boolValue := cef.ValueRef.New()
		boolValue.SetBool(true)
		ch := cef.RequestContextHandlerRef.New()
		// 设置请求上下文插件设置
		ch.SetOnRequestContextInitialized(func(requestContext *cef.ICefRequestContext) {
			err, ok := requestContext.SetPreference("profile.default_content_setting_values.plugins", intValue)
			println("profile.default_content_setting_values.plugins:", err, ok)
			err, ok = requestContext.SetPreference("plugins.run_all_flash_in_allow_mode", boolValue)
			println("plugins.run_all_flash_in_allow_mode:", err, ok)
			err, ok = requestContext.SetPreference("plugins.allow_outdated", boolValue)
			println("plugins.allow_outdated:", err, ok)
		})
		context := cef.RequestContextRef.New(cef.TCefRequestContextSettings{}, ch)
		// 主窗口的扩展参数设置
		window.SetCreateBrowserExtraInfo("", context, nil)
		// 弹出窗口
		event.SetOnBeforePopup(func(sender lcl.IObject, popupWindow cef.IBrowserWindow, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
			// 弹出窗口的扩展参数设置
			popupWindow.SetCreateBrowserExtraInfo("", context, nil)
			return false
		})
	})
	//运行应用
	cef.Run(cefApp)
}
