package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"path/filepath"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache", "jstogo")
	app := cef.NewApplication()
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetUseMockKeyChain(true)
	app.SetEnableGPU(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/js-to-go.html"
	cef.BrowserWindow.Config.Title = "Energy - js on event - go emit event"

	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
		// 在这里模拟传递参数在主进程触发JS监听的事件
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 定时执行web js
		go timeTask(window)
	})
	//运行应用
	cef.Run(app)
}

// 定时执行web js
func timeTask(window cef.IBrowserWindow) {
	//这里模拟go中触发js监听的事件
	var param0 = 0
	for {
		//每1秒钟执行一次
		time.Sleep(time.Second * 2)
		fmt.Println("timeTask", param0)
		param0++
		//将数据发送出去
		ipc.Emit("js-on-event-demo", fmt.Sprintf("Go发送的数据: %d", param0), float64(param0+10))
		// 如果JS返回结果, 需要通过回调函数入参方式接收返回值
		ipc.EmitAndCallback("js-on-event-demo-return", []interface{}{fmt.Sprintf("Go发送的数据: %d", param0), float64(param0 + 10)}, func(r1 string) {
			//需要正确的获取类型，否则会失败
			fmt.Println("JS返回数据:", r1)
		})
		browser := window.Browser()
		frameNames := browser.GetFrameNames()
		for i := 0; i < browser.FrameCount(); i++ {
			frame := browser.GetFrameByName(frameNames[i].Value)
			if !frame.IsMain() {
				fmt.Println("\tname:", frameNames[i].Name, "value:", frameNames[i].Value, "frameId:", frame.Identifier())
				targetFrame := target.NewTarget(frame, window.Browser().BrowserId(), frame.Identifier())
				ipc.EmitTarget("js-on-event-demo", targetFrame, fmt.Sprintf("当前FrameId: %v Go发送的数据: %d", frame.Identifier(), param0), float64(param0+10))
			}
		}
	}
}
