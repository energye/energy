package crawling

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"time"
)

// 在这里维护创建过的窗口
var windows = make(map[int]*WindowInfo)

type WindowInfo struct {
	energy *rod.Energy
	url    string
}
type Info struct {
	WindowId int
	URL      string
}

// WindowIds 返回所有windowId
func WindowIds() (result []*Info) {
	for id, info := range windows {
		result = append(result, &Info{
			WindowId: id,
			URL:      info.url,
		})
	}
	return
}

// Create 创建一个浏览器窗口
func Create() int {
	windowId := time.Now().Nanosecond()
	wp := cef.NewWindowProperty()
	// 创建一个 energy 扩展 rod 的窗口
	energyWindow := rod.NewEnergyWindow(nil, wp, nil)
	windows[windowId] = &WindowInfo{energy: energyWindow}
	createHandle(windowId, energyWindow)
	return windowId
}

func createHandle(newWindowId int, energy *rod.Energy) {
	//注册处理弹出窗口
	energy.SetOnBeforePopup(func(energyWindow *rod.Energy) {
		// 创建新窗口ID
		windowId := time.Now().Nanosecond()
		url := energyWindow.BrowserWindow().WindowProperty().Url
		windows[windowId] = &WindowInfo{energy: energyWindow, url: url}
		createHandle(windowId, energyWindow)
		//通知主应用页面有新窗口创建
		ipc.Emit("create-window", windowId, url)

		// 采集一些东西
		page := energyWindow.Page()
		fmt.Println("OnBeforePopup TargetID:", page.TargetID)
		elements := page.MustElements("a")
		fmt.Println("A tag - count:", len(elements))
		fmt.Println("title:", page.MustElement("title").MustText())

	})
	window := energy.BrowserWindow().AsLCLBrowserWindow().BrowserWindow()
	window.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
		ipc.Emit("close-window", newWindowId)
		return false
	})
}

// Show 显示窗口，在energy中不显示窗口无法使用rod功能
func Show(windowId int, url string) {
	if window, ok := windows[windowId]; ok {
		window.energy.Chromium().SetDefaultURL(url) //在这里设置个url
		window.url = url
		// UI线程中创建浏览器
		cef.RunOnMainThread(func() {
			window.energy.CreateBrowser()
		})
	}
}

// Close 关闭一个窗口
func Close(windowId int) bool {
	if window, ok := windows[windowId]; ok {
		window.energy.BrowserWindow().CloseBrowserWindow()
		delete(windows, windowId)
		return true
	}
	return false
}

// Crawling 抓取一些内容测试
func Crawling(windowId int) {
	if window, ok := windows[windowId]; ok {
		page := window.energy.Page()
		fmt.Println("TargetID:", page.TargetID)
		page.MustElement("#kw").MustSelectAllText().MustInput("") //清空文本框
		page.MustElement("#kw").MustInput("go energy")            //输入内容
		page.MustElement("#su").MustClick()                       //点击按钮
		wrapper := page.MustElement("#wrapper_wrapper")           //根据id获取标签
		containers := wrapper.MustElements(".c-container")        //根据class样式获取所有标签
		for len(containers) == 0 {                                // 返回0个继续获取
			containers = wrapper.MustElements(".c-container")
		}
		fmt.Println("containers:", len(containers))
		for _, container := range containers {
			a := container.MustElement("a")
			fmt.Println("a:", a.MustText())
		}
		fmt.Println(page.MustHTML()[0:100])
	}
}
