package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		screen := window.Screen()
		pr := screen.Primary()
		fmt.Println("Primary:", pr)
		fmt.Println("MonitorCount", screen.Count())
		for i := 0; i < screen.Count(); i++ {
			monitor := screen.Get(i)
			fmt.Println("\tmonitor", i, "work-rect:", monitor)
		}
	})
	//运行应用
	cef.Run(app)
}
