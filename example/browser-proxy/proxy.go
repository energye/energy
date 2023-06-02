package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"strings"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication()
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 使用右键菜单切换需要代理的 url s
		var (
			loadEnergyUrl consts.MenuId
			loadBaiduUrl  consts.MenuId
		)
		event.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, model *cef.ICefMenuModel) {
			model.AddSeparator()
			loadEnergyUrl = model.CefMis.NextCommandId()
			model.AddCheckItem(loadEnergyUrl, "load-energy")
			loadBaiduUrl = model.CefMis.NextCommandId()
			model.AddCheckItem(loadBaiduUrl, "load-baidu")
		})
		event.SetOnContextMenuCommand(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
			if commandId == loadEnergyUrl {
				window.Chromium().LoadUrl("https://energy.yanghy.cn")
			} else if commandId == loadBaiduUrl {
				window.Chromium().LoadUrl("https://www.baidu.com")
			}
		})
		// 使用 load start 和 load end 回调观察当前加载的url
		event.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
			fmt.Println("load-start:", frame.Url())
		})
		event.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
			fmt.Println("load-end:", frame.Url())
		})
		// 在 on before browser 配置代理
		event.SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
			requestContext := browser.GetRequestContext()
			if strings.Index(request.URL(), "www.baidu.com") != -1 {
				// 需要配置代理的 url
				proxyDict := cef.DictionaryValueRef.New()
				// 配置字典参数参考
				//&cef.TCefProxy{
				//	ProxyType:              consts.PtAutodetect,
				//	ProxyScheme:            consts.PsSOCKS4,
				//	ProxyServer:            "192.168.1.100",
				//	ProxyPort:              8888,
				//	ProxyUsername:          "username",
				//	ProxyPassword:          "password",
				//	ProxyScriptURL:         "scriptURL",
				//	ProxyByPassList:        "aaa,bbb,ddd",
				//	MaxConnectionsPerProxy: 100,
				//}
				proxyDict.SetString("mode", "fixed_servers")
				proxyDict.SetString("server", "192.168.1.100:8888")
				// 最后通过 Value 将代理字典值设置到 requestContext.SetPreference("proxy", proxy)
				proxy := cef.ValueRef.New()
				proxy.SetDictionary(proxyDict)
				errMsg, ok := requestContext.SetPreference("proxy", proxy)
				fmt.Println("\tproxy errMsg:", errMsg, "ok:", ok)
			} else {
				// 不需要代理的 url value 设置 nil
				errMsg, ok := requestContext.SetPreference("proxy", nil)
				fmt.Println("\tproxy errMsg:", errMsg, "ok:", ok)
			}
			return false
		})
		// 如果代理需要用户名和密码
		window.Chromium().SetOnGetAuthCredentials(func(sender lcl.IObject, browser *cef.ICefBrowser, originUrl string, isProxy bool, host string, port int32, realm, scheme string, callback *cef.ICefAuthCallback) bool {
			fmt.Println("AuthCredentials:", originUrl, "isProxy:", isProxy)
			if isProxy {
				//callback.Cont("username","password")
				return true
			}
			return false
		})
	})
	//运行应用
	cef.Run(cefApp)
}
