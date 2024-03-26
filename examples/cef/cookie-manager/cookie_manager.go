package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"os"
	"path/filepath"
	//_ "net/http/pprof"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	var app = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.Config.Title = "Energy - cookie-manager"
	envPath := os.Getenv("ENERGY_HOME")
	if envPath == "" {
		// 未配置 ENERGY_HOME 环境变量， 在当前目录保存cache
		envPath = types.ExeDir
		envPath = filepath.Join(envPath, "cache")
	} else {
		envPath = filepath.Join(envPath, "cache")
	}
	println("cache-path:", envPath)
	// 设置cache目录, 持久化 localStorage 和 cookie
	app.SetCache(envPath)
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		window.Chromium().SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
			manager := browser.GetRequestContext().GetCookieManager(nil)
			url := frame.Url()
			fmt.Println("manager:", manager, url)
			visitor := cef.CookieVisitorRef.New()
			visitor.SetOnVisit(func(cookie *cef.ICefCookie) (deleteCookie, result bool) {
				fmt.Printf("cookie: %+v\n", cookie)
				window.Chromium().SetCookie(url, cookie.Name, cookie.Value, cookie.Domain, cookie.Path, cookie.Secure, cookie.Httponly, cookie.HasExpires, cookie.Creation, cookie.LastAccess, cookie.Expires, cookie.SameSite, cookie.Priority, cookie.SetImmediately, cookie.ID)
				fmt.Println("\tFlushCookieStore:", window.Chromium().FlushCookieStore(true))
				return false, true
			})
			manager.VisitAllCookies(visitor)
		})
	})
	//运行应用
	cef.Run(app)
}
