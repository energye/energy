package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
)

//资源目录，内置到执行程序中
//go:embed resources
var resources embed.FS

//这个示例使用了几个事件来演示下载文件
//在cef.BrowserWindow.SetBrowserInit初始化函数中设置event.SetOnBeforeDownload，用于设置保存目录
//并且设置event.SetOnDownloadUpdated获取下载进度信息
func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/download.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	//在主窗口初始化回调函数里设置浏览器事件
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow cef.IBrowserWindow) {
		//linux 下载文件 系统弹出保存对话框不启作用
		//所以 自己调用系统的保存对话框获得保存路径
		linuxDlSave := lcl.NewSaveDialog(browserWindow.AsLCLBrowserWindow().BrowserWindow())
		linuxDlSave.SetTitle("保存对话框标题")

		//下载之前事件
		event.SetOnBeforeDownload(func(sender lcl.IObject, browser *cef.ICefBrowser, beforeDownloadItem *cef.DownloadItem, suggestedName string, callback *cef.ICefBeforeDownloadCallback) {
			fmt.Println("下载之前事件")
			//设置下载目录, 和弹出保存窗口
			if common.IsLinux() {
				//linux 在大多数据情况操作UI相关的需要使用 QueueSyncCall 函数包起来
				cef.QueueSyncCall(func(id int) {
					linuxDlSave.SetFileName(suggestedName)
					if linuxDlSave.Execute() {
						// showDialog = false 不显示保存对话框
						callback.Cont(linuxDlSave.FileName(), false)
					}
				})
			} else {
				//windows macosx
				callback.Cont(consts.ExePath+consts.Separator+suggestedName, true)
			}
		})
		//下载更新事件
		//1. 返回下载进度
		//2. downloadItem 下载项
		//3. callback 下载状态的控制, 下载暂停，开始、取消
		//4. 将下载进度通过事件机制发送到html中展示
		event.SetOnDownloadUpdated(func(sender lcl.IObject, browser *cef.ICefBrowser, downloadItem *cef.DownloadItem, callback *cef.ICefDownloadItemCallback) {
			//传递数据参数到html中
			//这些参数按下标顺序对应到js函数参数位置
			//演示只传递了几个参数
			var argumentList = ipc.NewArgumentList()
			//第1个参数下载的ID
			argumentList.SetInt32(0, downloadItem.Id)
			//第2个参数文件名
			argumentList.SetString(1, downloadItem.FullPath, true)
			//第3个参数 接收的字节数 - 这里需要注意的是，IPC消息 数字类型只支持32位的，如果大于32位，需要转换成string类型发送, 否则直接使用int64类型会导致消息接收不到
			argumentList.SetInt32(2, int32(downloadItem.ReceivedBytes))
			//第4个参数 文件总大小字节数 - 例如: int64转成string做为参数
			argumentList.SetString(3, fmt.Sprintf("%d", downloadItem.TotalBytes), true)
			browserWindow.Chromium().Emit("downloadUpdateDemo", argumentList, browser)
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
		server.Assets = &resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
