package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
)

// CEF 49是最后一个支持Windows XP的CEF版本
// 在energy中使用与新的CEF版本没太多区别，因版本跨度太大不支持新的CEF API, 使用时需注意且会报出相关错误, 虽然energy提供了。
// 仅在Windows XP以及其它Windows系统中运行, 不支持Linux、MacOS
// 如果必须在Windows XP环境中运行CEF，需要接受旧版本的限制和潜在的安全风险。

/*
  由于Golang最后一个支持WindowsXP版本是1.10, 经测试Golang1.11版本编译的执行文件也可在 Windows XP SP3 中运行, 同时Golang1.11也支持go.mod模块管理, 方便很多.
  此时我们应使用Golang1.11，也可能有未知的不支持或错误还未发现.
  注意Golang版本大于1.11编译出的执行文件无法在Windows XP SP3中运行.
    Golang1.10, 经测试目前无法正常运行energy构建的程序
  内嵌资源, Golang1.11不支持 embed.FS 该功能是Golang1.16才开始支持
    在energy最新命令行工具中扩展了bindata命令参数, 通过go:generate来生成静态资源“字节数据”进行绑定
    以达到静态资源内嵌
*/

// 以下命令生成静态资源“字节数据”进行绑定，使用 go 命令: go generate
/*
 说明:
  固定 //go:generate energy bindata --fs
  可变 --o=输出目录, --pkg=包名,  --paths=打包目录可多个‘,‘豆号分隔
*/
// 静态资源
//go:generate energy bindata --fs --o=pkg/assets/assets.go --pkg=assets --paths=./resources/...
// 动态链接库
//go:generate energy bindata --fs --o=pkg/libs/libs.go --pkg=libs --paths=./libs

func main() {
	//libsFS := libs.AssetFile()
	//resourceFS := assets.AssetFile()
	//全局初始化 每个应用都必须调用的
	//cef.GlobalInit(libsFS, resourceFS)
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	// 主进程回调函数
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		//server.Assets = resourceFS
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
