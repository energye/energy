package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication(true)
	//
	cefApp.SetWindowlessRenderingEnabled(true)
	cefApp.SetExternalMessagePump(true)
	cefApp.SetMultiThreadedMessageLoop(false)
	// work
	cef.GlobalWorkSchedulerCreate(nil)
	cefApp.SetOnScheduleMessagePumpWork(nil)

	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	//运行应用
	cef.Run(cefApp)
	fmt.Println("run end")
}