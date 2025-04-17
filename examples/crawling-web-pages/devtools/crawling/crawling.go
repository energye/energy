package crawling

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/energye/erod"
	"strings"
	"time"
)

// 在这里维护创建过的窗口
var windows = make(map[int]*WindowInfo)

type WindowInfo struct {
	energy *erod.Energy
	url    string
	typ    int
}

type Info struct {
	WindowId int
	URL      string
	Typ      int
}

// WindowIds 返回所有windowId
func WindowIds() (result []*Info) {
	for id, info := range windows {
		result = append(result, &Info{
			WindowId: id,
			URL:      info.url,
			Typ:      info.typ,
		})
	}
	return
}

// Create 创建一个浏览器窗口
func Create(url string, testType int) int {
	windowId := time.Now().Nanosecond()
	wp := cef.NewWindowProperty()
	wp.Url = url // 创建时指定一个URL
	// 创建一个 energy 扩展 rod 的窗口
	energyWindow := erod.NewEnergyWindow(nil, wp, nil)
	windows[windowId] = &WindowInfo{energy: energyWindow, typ: testType}
	createHandle(windowId, energyWindow)
	return windowId
}

// 弹出或创建窗口处理，主要一些事件
func createHandle(newWindowId int, energy *erod.Energy) {
	//注册处理弹出窗口
	energy.SetOnBeforePopup(func(energyWindow *erod.Energy) {
		// 创建新窗口ID
		windowId := time.Now().Nanosecond()
		url := energyWindow.BrowserWindow().WindowProperty().Url
		windows[windowId] = &WindowInfo{energy: energyWindow, url: url}
		createHandle(windowId, energyWindow)
		//通知主应用页面有新窗口创建
		ipc.Emit("create-window", windowId, url)

		// 采集一些东西
		page := energyWindow.Page().MustWaitLoad()
		fmt.Println("OnBeforePopup TargetID:", page.TargetID)
		elements := page.MustElements("a")
		fmt.Println("A tag - count:", len(elements))
		fmt.Println("title:", page.MustElement("title").MustText())

	})
	energy.SetOnLoadingProgressChange(func(energy *erod.Energy, progress float64) {
		ipc.Emit("window-loading-progress", newWindowId, int(progress*100))
	})
	// 窗口关闭时调用，通知主窗口，有窗口关闭
	energy.SetOnClose(func(energy *erod.Energy) {
		ipc.Emit("close-window", newWindowId)
		delete(windows, newWindowId)
	})
}

// Show 显示窗口，在energy中不显示窗口无法使用rod功能
func Show(windowId int, url string) {
	if window, ok := windows[windowId]; ok {
		window.energy.CreateBrowser() // CreateBrowser 创建成功后，不会重复创建
		if url != "" {
			window.url = url
			window.energy.ChromiumBrowser().Chromium().LoadUrl(url)
		}
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
		page := window.energy.Page().MustWaitLoad()
		fmt.Println("TargetID:", page.TargetID)
		head := page.MustElement(`ul[class="mb-0 flex items-center"]`) //清空文本框 .MustSelectAllText().MustInput("")
		li2 := head.MustElements("li")[1]
		openSource := li2.MustElement("a")
		openSource.MustClick() // 开源点击
		queryForm := page.MustWaitLoad().MustElement(`form[class="ui form custom js-form-control"]`)
		queryInp := queryForm.MustElement("#q")
		queryInp.MustSelectAllText().MustInput("") // 清空文本框
		queryInp.MustInput("energy")               // 输入搜索内容
		queryBtn := queryForm.MustElement(`button`)
		queryBtn.MustClick() //点击后 跳转页面
		hitsList := page.MustWaitLoad().MustElement("#hits-list")
		titles := hitsList.MustElements(`div[class="title"]`)
		for _, title := range titles {
			a := title.MustElement("a")
			href := a.MustAttribute("href")
			fmt.Println("gitee energy href:", *href)
			if strings.Index(*href, "gitee.com/energye/energy") != -1 {
				a.MustClick() //点击后 跳转页面
				break
			}
		}
	}
}
