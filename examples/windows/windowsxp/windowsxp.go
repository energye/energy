package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/pkgs/assetserve"
)

// windows xp 使用 CEF 49, 是最后一个支持 windows xp 的CEF版本, 使用-与比较新的CEF版本没太多区别，但注意的是不支持新的CEF API 使用时需注意, 虽然energy提供了。
// 仅在 windows xp, 以及其它 windows 系统中运行, 不支持 Linux, MacOS

/*
  由于Golang最后一个支持WindowsXP版本是1.10, 但经测试Golang1.11版本编译的执行文件也可在 Windows XP SP3 中运行, 同时Golang1.11也支持go.mod模块管理, 方便很多.
  此时我们应使用Golang1.11，也可能有未知的不支持或错误还未发现.
    Golang1.10, 经测试目前无法正常运行构建energy的程序
  内嵌资源, Golang1.11不支持 embed.FS
    在energy最新命令行工具中扩展了bindata命令参数, 通过go:generate来生成静态资源“字节数据”进行绑定
    以达到静态资源内嵌
*/
// 以下命令生成静态资源“字节数据”进行绑定，使用 go 命令: go generate
//go:generate energy bindata -fs -o=assets.go -pkg=assets ./assets

func main() {

	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	// 主进程回调函数
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
