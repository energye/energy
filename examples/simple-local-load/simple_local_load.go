package main

import (
	"embed"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"path/filepath"
)

//go:embed resources
var resources embed.FS

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.LDebug)
	// 全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)

	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache", "simple-local-load")
	logger.Debug("rootCache:", rootCache)
	// 创建应用
	app := cef.NewApplication()
	app.SetUseMockKeyChain(true)
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
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
