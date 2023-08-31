package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	var app = cef.NewApplication()
	if common.IsLinux() && app.IsUIGtk3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	cef.BrowserWindow.Config.Title = "Energy - Local load"
	// 本地加载资源方式, 直接读取本地或内置执行文件资源
	// 该模块不使用 http server
	// 默认访问地址fs://energy/index.html, 仅能在应用内访问
	//   fs: 默认的自定义协议名, 你可以任意设置
	//   energy: 默认的自定义域, 你可以任意设置
	//   index.html: 默认打开的页面名，你可以任意设置
	// 页面ajax xhr数据获取
	// xhr数据获取通过Proxy配置, 支持http, https证书配置
	cef.BrowserWindow.Config.Url = "fs://energy" // 设置默认
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		Scheme:     "fs",             // 自定义协议名
		Domain:     "energy",         // 自定义域名
		ResRootDir: "resources/dist", // 资源存放目录, FS不为空时是内置资源目录名, 空时当前文件执行目录, @/to/path @开头表示当前目录下开始
		FS:         &resources,       //静态资源所在的 embed.FS
		Proxy: &cef.XHRProxy{ // 页面Ajax XHR请求接口代理转发配置
			Scheme: consts.LpsHttps,  // http's 支持ssl配置
			IP:     "your.domain.cn", //http服务ip或domain
			SSL: cef.XHRProxySSL{ // ssl 证书配置,如果使用https但未配置该选项，默认跳过ssl检查
				FS:      &resources,         //如果证书内置到exe中,需要在此配置 embed.FS
				RootDir: "resources/ssl",    //证书存放目录
				Cert:    "demo.pem",         //ssl cert
				Key:     "demo.key",         //ssl key
				CARoots: []string{"ca.cer"}, // ssl ca root
			},
		},
	}.Build())
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {

	})
	//运行应用
	cef.Run(app)
}
