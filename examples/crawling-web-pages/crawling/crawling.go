package crawling

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod"
	"time"
)

// 在这里维护创建过的窗口
var windows = make(map[int]*rod.Chromium)

// WindowIds 返回所有windowId
func WindowIds() (result []int) {
	for id, _ := range windows {
		result = append(result, id)
	}
	return
}

// Create 创建一个浏览器窗口
func Create(url string) int {
	windowId := time.Now().Nanosecond()
	wp := cef.NewWindowProperty()
	wp.Url = url
	// 创建一个 energy 扩展 rod 的窗口
	rodWindow := rod.NewWindow(nil, wp, nil)
	windows[windowId] = rodWindow
	// 在UI线程中创建或显示这个窗口
	cef.RunOnMainThread(func() {
		rodWindow.CreateBrowser()
	})
	return windowId
}

// Close 关闭一个窗口
func Close(windowId int) {
	if window, ok := windows[windowId]; ok {
		window.BrowserWindow().CloseBrowserWindow()
	}
}

// Crawling 抓取一些内容测试
func Crawling(windowId int) {
	if window, ok := windows[windowId]; ok {
		page := window.Page()
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
