package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

// 资源目录，内置到执行程序中
//
//go:embed resources
var resources embed.FS

// 这个示例使用了几个事件来演示下载文件
// 在cef.BrowserWindow.SetBrowserInit初始化函数中设置event.SetOnBeforeDownload，用于设置保存目录
// 并且设置event.SetOnDownloadUpdated获取下载进度信息
func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/download.html"

	//在主窗口初始化回调函数里设置浏览器事件
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow cef.IBrowserWindow) {
		//下载更新事件
		//1. 返回下载进度
		//2. downloadItem 下载项
		//3. callback 下载状态的控制, 下载暂停，开始、取消
		//4. 将下载进度通过事件机制发送到html中展示
		event.SetOnDownloadUpdated(func(sender lcl.IObject, browser *cef.ICefBrowser, downloadItem *cef.ICefDownloadItem, callback *cef.ICefDownloadItemCallback) {
			//传递数据参数到html中
			//这些参数按下标顺序对应到js函数参数位置
			//演示只传递了几个参数
			fmt.Println("DownloadUpdated frameId", browser.MainFrame().Identifier(), "BeforeDownload Id:", downloadItem.Id(), "originalUrl:", downloadItem.OriginalUrl(), "url:", downloadItem.Url())
			fmt.Println("\t", downloadItem.State(), downloadItem.TotalBytes(), "/", downloadItem.ReceivedBytes(), "speed:", downloadItem.CurrentSpeed(), "fullPath", downloadItem.FullPath())
			ipc.Emit("downloadUpdateDemo", downloadItem.Id(), downloadItem.FullPath(), downloadItem.ReceivedBytes(), downloadItem.TotalBytes())
		})
		// CEF 下载貌似有些问题，下载过程中右键下载文件会现文件浏览窗口卡死问题
		// 测试 cefclient.exe 也有这个问题
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
