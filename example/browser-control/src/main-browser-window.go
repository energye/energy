package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"os"
)

//主浏览器窗口
func MainBrowserWindow() {
	//只有启动主进程才会继续执行
	if !common.Args.IsMain() {
		return
	}
	fmt.Println("os.Args", os.Args)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "https://www.baidu.com"
	//窗口的标题
	cef.BrowserWindow.Config.Title = "Energy - 浏览器控制"
	//窗口宽高
	cef.BrowserWindow.Config.Width = 1024
	cef.BrowserWindow.Config.Height = 768
	//chromium配置
	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)
	config.SetEnableDevTools(true)
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//创建窗口时的回调函数 对浏览器事件设置，和窗口属性组件等创建和修改
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		//设置应用图标 这里加载的图标是内置到执行程序里的资源文件
		lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")
		//在窗体初始化时创建窗口内的组件
		back, forward, stop, refresh, progressLabel, addr := controlUI(browserWindow)
		//页面加载处理进度
		event.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
			//参数-进度
			progressLabel.SetCaption(fmt.Sprintf("%v", progress*100))
		})
		//页面加载状态，根据状态判断是否加载完成，和是否可前进后退
		event.SetOnLoadingStateChange(func(sender lcl.IObject, browser *cef.ICefBrowser, isLoading, canGoBack, canGoForward bool) {
			//控制按钮状态
			stop.SetEnabled(isLoading)
			refresh.SetEnabled(!isLoading)
			back.SetEnabled(canGoBack)
			forward.SetEnabled(canGoForward)
		})
		event.SetOnAddressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, url string) {
			addr.SetText(url)
		})
	})
	//创建窗口之后对对主窗口的属性、组件或子窗口的创建
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInitAfter")
	})
}

//控制组件UI
//地址栏和控制按钮创建
func controlUI(browserWindow *cef.TCefWindowInfo) (goBack *lcl.TButton, goForward *lcl.TButton, stop *lcl.TButton, refresh *lcl.TButton, progressLabel *lcl.TLabel, addrBox *lcl.TComboBox) {
	window := browserWindow.Window
	//这里使用系统UI组件
	//创建panel做为地址栏的父组件
	addrPanel := lcl.NewPanel(window) //设置父组件
	addrPanel.SetParent(window)
	addrPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小
	addrPanel.SetHeight(25)
	addrPanel.SetWidth(window.Width())
	//创建 按钮-后退
	goBack = lcl.NewButton(addrPanel) //设置父组件
	goBack.SetParent(addrPanel)
	goBack.SetCaption("后退")
	goBack.SetBounds(0, 0, 35, 25)
	goForward = lcl.NewButton(addrPanel) //设置父组件
	goForward.SetParent(addrPanel)
	goForward.SetCaption("前进")
	goForward.SetBounds(35, 0, 35, 25)
	stop = lcl.NewButton(addrPanel) //设置父组件
	stop.SetParent(addrPanel)
	stop.SetCaption("停止")
	stop.SetBounds(35+35, 0, 35, 25)
	refresh = lcl.NewButton(addrPanel) //设置父组件
	refresh.SetParent(addrPanel)
	refresh.SetCaption("刷新")
	refresh.SetBounds(35+35+35, 0, 35, 25)

	//创建下拉框
	addrBox = lcl.NewComboBox(addrPanel)
	addrBox.SetParent(addrPanel)
	addrBox.SetLeft(35 + 35 + 35 + 35)                                         //这里是设置左边距 上面按钮的宽度
	addrBox.SetWidth(window.Width() - (35 + 35 + 35 + 35 + 35 + 35))           //宽度 减按钮的宽度
	addrBox.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小
	addrBox.Items().Add("https://energy.yanghy.cn")
	addrBox.Items().Add("https://www.baidu.com")

	//显示加载进度
	progressLabel = lcl.NewLabel(addrPanel) //设置父组件
	progressLabel.SetParent(addrPanel)
	progressLabel.SetCaption("0")
	progressLabel.SetBounds(window.Width()-(35+35)-10, 3, 35, 25)
	progressLabel.SetAnchors(types.NewSet(types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小

	goUrl := lcl.NewButton(addrPanel) //设置父组件
	goUrl.SetParent(addrPanel)
	goUrl.SetCaption("GO")
	goUrl.SetBounds(window.Width()-35, 0, 35, 25)
	goUrl.SetAnchors(types.NewSet(types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小

	//重新调整browser窗口的Parent属性
	//重新设置了上边距，宽，高
	window.WindowParent().SetAlign(types.AlNone) //重置对齐,默认是整个客户端
	window.WindowParent().SetTop(25)
	window.WindowParent().SetHeight(window.Height() - 25)
	window.WindowParent().SetWidth(window.Width())
	//设置锚点定位，让宽高自动根据窗口调整大小
	//因为窗口大小已调整，这里不能使用 SetAlign 了
	window.WindowParent().SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))

	//给按钮增加事件
	goBack.SetOnClick(func(sender lcl.IObject) {
		browserWindow.Chromium().GoBack()
	})
	goForward.SetOnClick(func(sender lcl.IObject) {
		browserWindow.Chromium().GoForward()
	})
	stop.SetOnClick(func(sender lcl.IObject) {
		browserWindow.Chromium().StopLoad()
	})
	refresh.SetOnClick(func(sender lcl.IObject) {
		browserWindow.Chromium().Reload()
	})
	goUrl.SetOnClick(func(sender lcl.IObject) {
		var url = addrBox.Text()
		if url != "" {
			browserWindow.Chromium().LoadUrl(url)
		}
	})
	return
}
