package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/target"
	"github.com/energye/energy/v2/consts"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	var app = cef.NewApplication()
	cef.BrowserWindow.Config.Title = "Energy - mockevent"
	cef.BrowserWindow.Config.Url = "fs://energy" // 设置默认
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         &resources, //静态资源所在的 embed.FS
	}.Build())

	app.SetOnRenderLoadEnd(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
		// 创建 dom visitor
		visitor := cef.DomVisitorRef.New()
		// 监听事件
		// 这个事件在渲染进程中才会执行
		visitor.SetOnVisit(func(document *cef.ICefDomDocument) {
			body := document.GetBody()
			inpText := body.GetDocument().GetElementById("inpText")
			fmt.Println("inpText", inpText.GetElementBounds())
			ipc.EmitTarget("dom", target.NewTargetMain(), 1, 2)
		})
		fmt.Println("visitor:", visitor)
		frame.VisitDom(visitor)
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		var isLoad bool
		// 模拟一个其它任务中执行鼠标事件
		var mockEvent = func() {
			// 页面加载完之后
			for {
				if isLoad {
					break
				}
				time.Sleep(time.Second)         //一秒后执行
				window.RunOnMainThread(func() { //在UI主线程
					chromium := window.Chromium()
					// 鼠标事件
					me := &cef.TCefMouseEvent{}
					/*
					  演示的是点击界面左上角Button按钮, 前提是需要知道按钮在界面的坐标
					*/
					// 设置坐标 相当于窗口内
					me.Y = 12
					me.X = 25
					// 模拟鼠标到指定位置
					chromium.SendMouseMoveEvent(me, false)
					// 模拟鼠标双击事件
					//   左键点击按下1次
					chromium.SendMouseClickEvent(me, consts.MBT_LEFT, false, 1)
					//   左键点击抬起1次
					chromium.SendMouseClickEvent(me, consts.MBT_LEFT, true, 1)

					// 模拟文本框
					//chromium.SendKeyEvent()
				})
			}
		}
		ipc.On("dom", func(s, a int) {
			fmt.Println("dom", s, a)
			mockEvent()
		})
	})
	//运行应用
	cef.Run(app)
}
