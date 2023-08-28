package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	var app = cef.NewApplication()
	//app.SetDisableWebSecurity(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Title = "Energy - Local load"
	var e = true
	if e {
		cef.BrowserWindow.Config.Url = "http://energy"
		cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
			Enable: true,
			//ResRootDir: "@/dist",
			Domain:     "energy",
			ResRootDir: "resources/dist",
			FS:         &resources,
			Proxy: &cef.XHRProxy{
				Scheme: consts.LpsHttps,
				IP:     "energy.yanghy.cn",
				SSL: cef.XHRProxySSL{
					FS:      &resources,
					RootDir: "resources/ssl",
					Cert:    "demo.energy.pem",
					Key:     "demo.energy.key",
					CARoots: []string{"root.cer"},
				},
			},
		}.Build())
	} else {
		cef.BrowserWindow.Config.Url = "http://192.168.0.108:8080/"
	}
	if common.IsLinux() && app.IsUIGtk3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {

	})
	//运行应用
	cef.Run(app)
}
