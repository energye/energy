package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/ipc"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/js-to-go.html"
	cef.BrowserWindow.Config.Title = "Energy - js on event - go emit event"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"

	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
		//定时执行web js
		go timeTask()
	})
	//运行应用
	cef.Run(cefApp)
}

//定时执行web js
func timeTask() {
	//这里模拟go中触发js监听的事件
	var param0 = 0
	for {
		//每1秒钟执行一次
		time.Sleep(time.Second)
		param0++
		//在go中触发js事件需要通过chromium对象
		//获得主窗口对象
		info := cef.BrowserWindow.MainWindow()
		//虽然 Arguments 结构支持多个数据类型，但在js和go的对应中，只保留了 string, integer, double, boolean 的对应关系，其它类型在 go 和 js数据传递时不支持
		args := ipc.NewArgumentList()
		//在给js发送数据时，string类型的，需要设置第3个参数为true, 底层在处理字符串时需要
		args.SetString(0, fmt.Sprintf("Go发送的数据: %d", param0), true)
		args.SetFloat64(1, float64(param0+10))
		//将数据发送出去
		info.Chromium().Emit("js-on-event-demo", args, info.Browser())
		//触发js监听函数，这个函数带有返回值到go中
		info.Chromium().EmitAndCallback("js-on-event-demo-return", args, info.Browser(), func(context ipc.IIPCContext) {
			//因为js和go不一样，返回值只能有1个，且是通过 Arguments 获取
			arguments := context.Arguments()
			//需要正确的获取类型，否则会失败
			fmt.Println("JS返回数据:", arguments.GetString(0))
		})
		fmt.Println(ipc.IPC.Browser().ChannelIds())
	}
}
