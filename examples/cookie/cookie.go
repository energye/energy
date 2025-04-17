package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"time"
)

// 资源目录，内置到执行程序中
//
//go:embed resources
var resources embed.FS

// 这个示例使用了几个事件来演示下载文件
func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/cookie.html"

	//监听获取cookie事件
	ipc.On("VisitCookie", func(context context.IContext) {
		fmt.Println("VisitCookie")
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		info.Chromium().VisitURLCookies("https://www.baidu.com", true, 1)
		info.Chromium().VisitAllCookies(1)
		context.Result("执行成功，结果将在 SetOnCookiesVisited 事件中获得")
	})
	//监听删除cookie
	ipc.On("DeleteCookie", func(context context.IContext) {
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		info.Chromium().DeleteCookies("", "", false)
		context.Result("执行成功，结果将在 SetOnCookiesDeleted 事件中获得")
	})
	//监听设置cookie
	ipc.On("SetCookie", func(context context.IContext) {
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		info.Chromium().SetCookie("https://www.example.com", "example_cookie_name", "1234", "", "/", true, true, false, time.Now(), time.Now(), time.Now(), consts.Ccss_CEF_COOKIE_SAME_SITE_UNSPECIFIED, consts.CEF_COOKIE_PRIORITY_MEDIUM, false, 0)
		info.Chromium().SetCookie("https://www.example.com", "example_cookie_name2", "123422", "", "/", true, true, false, time.Now(), time.Now(), time.Now(), consts.Ccss_CEF_COOKIE_SAME_SITE_UNSPECIFIED, consts.CEF_COOKIE_PRIORITY_MEDIUM, false, 0)
		info.Chromium().SetCookie("https://www.baidu.com", "demo_name", "4321", "", "/", true, true, false, time.Now(), time.Now(), time.Now(), consts.Ccss_CEF_COOKIE_SAME_SITE_NO_RESTRICTION, consts.CEF_COOKIE_PRIORITY_MEDIUM, false, 1)
		context.Result("执行成功，结果将在 SetOnCookieSet 事件中获得")
	})

	//在SetBrowserInit中设置cookie事件,这些事件将返回操作后的结果
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow cef.IBrowserWindow) {
		//获取cookie时触发
		event.SetOnCookiesVisited(func(sender lcl.IObject, cookie *cef.TCefCookie, deleteCookie, result *bool) {
			fmt.Printf("SetOnCookiesVisited: %+v\n", cookie)
			//将结果返回到html中
			data, _ := json.Marshal(cookie)
			ipc.Emit("VisitCookieResult", string(data))
		})
		//删除cookie时触发
		event.SetOnCookiesDeleted(func(sender lcl.IObject, numDeleted int32) {
			fmt.Printf("SetOnCookiesDeleted: %+v\n", numDeleted)
		})
		//设置cookie时触发
		event.SetOnCookieSet(func(sender lcl.IObject, success bool, ID int32) {
			fmt.Println("SetOnCookieSet: ", success, ID)
		})
		event.SetOnCookiesFlushed(func(sender lcl.IObject) {
			fmt.Println("OnCookiesFlushed")
		})
		event.SetOnCookieVisitorDestroyed(func(sender lcl.IObject, ID int32) {
			fmt.Println("OnCookieVisitorDestroyed")
		})
	})
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
