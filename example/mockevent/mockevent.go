package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/target"
	"github.com/energye/energy/v2/consts"
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

	/*
		1. 仅在主进程中使用ipc监听renderLoadEnd事件
		2. 在渲染进程中的app.SetOnRenderLoadEnd监听渲染进程页面加载结束事件，在这个事件里获取页面html元素位置和大小
		3. html元素位置和大小获取之后触发主进程监听事件, 把要获取到的html元素位置和大小发送到主进程renderLoadEnd事件
		4. 主进程renderLoadEnd事件被触发后使用渲染进程传递的元素数据模拟事件操作
			模拟事件使用当前窗口chromium或者browser提供的函数
			chromium.SendXXX 只能在主进程中使用
			browser.SendXXX  可在主进程或渲染进程中直接使用
	*/

	// 在渲染进程中处理结束事件
	// 通过VisitDom获取html元素的位置和大小
	// SetOnVisit 函数只能在渲染进程中执行
	app.SetOnRenderLoadEnd(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
		// 创建 dom visitor
		visitor := cef.DomVisitorRef.New()
		// 监听事件
		// 这个事件在渲染进程中才会执行
		visitor.SetOnVisit(func(document *cef.ICefDomDocument) {
			body := document.GetBody()
			btn1 := body.GetDocument().GetElementById("btn1")
			btn2 := body.GetDocument().GetElementById("btn2")
			btn3 := body.GetDocument().GetElementById("btn3")
			inpText := body.GetDocument().GetElementById("inpText")
			fmt.Println("inpText", inpText.GetElementBounds())
			var doms = make(map[string]cef.TCefRect)
			doms["btn1"] = btn1.GetElementBounds()
			doms["btn2"] = btn2.GetElementBounds()
			doms["btn3"] = btn3.GetElementBounds()
			doms["inpText"] = inpText.GetElementBounds()
			ipc.EmitTarget("renderLoadEnd", target.NewTargetMain(), doms)
		})
		frame.VisitDom(visitor)
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		var domXYCenter = func(bound cef.TCefRect) (int32, int32) {
			return bound.X + bound.Width/2, bound.Y + bound.Height/2
		}
		// 模拟按钮点击事件
		var buttonClickEvent = func(domRect cef.TCefRect) {
			// 页面加载完之后
			window.RunOnMainThread(func() { //在UI主线程
				chromium := window.Chromium()
				// 鼠标事件
				me := &cef.TCefMouseEvent{}
				// 设置元素坐标，元素坐标相对于窗口，这里取元素中间位置
				me.X, me.Y = domXYCenter(domRect)
				fmt.Println("buttonClickEvent", me)
				// 模拟鼠标到指定位置
				chromium.SendMouseMoveEvent(me, false)
				// 模拟鼠标双击事件
				//   左键点击按下1次
				chromium.SendMouseClickEvent(me, consts.MBT_LEFT, false, 1)
				//   左键点击抬起1次
				chromium.SendMouseClickEvent(me, consts.MBT_LEFT, true, 1)

			})
		}
		ipc.On("renderLoadEnd", func(doms map[string]cef.TCefRect) {
			fmt.Println("doms", doms)
			buttonClickEvent(doms["btn1"])
			buttonClickEvent(doms["btn2"])
			buttonClickEvent(doms["btn3"])
		})
	})
	//运行应用
	cef.Run(app)
}
