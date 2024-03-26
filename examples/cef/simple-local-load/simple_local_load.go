package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
)

//go:embed resources
var resources embed.FS

func main() {
	// 全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	// 创建应用
	app := cef.NewApplication()
	// 本地加载资源方式, 直接读取本地或内置执行文件资源
	// 该模块不使用 http server
	// 默认访问地址fs://energy/index.html, 仅能在应用内访问
	//   fs: 默认的自定义协议名, 你可以任意设置
	//   energy: 默认的自定义域, 你可以任意设置
	//   index.html: 默认打开的页面名，你可以任意设置
	// 页面ajax xhr数据获取
	// xhr数据获取通过Proxy配置, 支持http, https证书配置
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         resources,
	}.Build())
	// 运行应用
	cef.Run(app)
}
