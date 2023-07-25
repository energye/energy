package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
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
		fmt.Println("name:", screen.Name())
		fmt.Println("default-monitor width:", screen.Width(), "height:", screen.Height())
		fmt.Println("MonitorCount", screen.MonitorCount())
		for i := 0; i < int(screen.MonitorCount()); i++ {
			monitor := screen.Monitors(int32(i))
			fmt.Println("\tmonitor", i, "workarea-rect:", monitor.WorkareaRect())
		}
	})
	//运行应用
	cef.Run(app)
}
