package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"os"
	"path/filepath"
	"time"
)

// 主窗口
type MainWindow struct {
	*lcl.TForm
	headless            *HeadlessBrowser
	pageLoadProgressLbl *lcl.TLabel
}

var BW *MainWindow

// 此示例只在windows下测试, 有能力可以修改MacOS，Linux

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	// 开始创建CEF应用程序
	app := cef.CreateApplication()
	app.SetEnableGPU(true)
	app.SetExternalMessagePump(false)
	app.SetMultiThreadedMessageLoop(true)
	app.SetWindowlessRenderingEnabled(true) // 设不设置感觉没什么区别
	// 指定 CEF Framework
	app.SetFrameworkDirPath(config.Get().FrameworkPath())
	if app.StartMainProcess() {
		// 结束应用后释放资源
		api.SetReleaseCallback(func() {
			fmt.Println("Release")
			app.Destroy()
			app.Free()
		})
		// LCL窗口
		lcl.Application.Initialize()
		// 应用图标
		icon, _ := common.ResourcesFS().ReadFile("resources/icon.ico")
		lcl.Application.Icon().LoadFromBytes(icon)
		lcl.Application.CreateForm(&BW, true)
		lcl.Application.Run()
	}
}

func (m *MainWindow) OnFormCreate(sender lcl.IObject) {
	m.SetHeight(100)
	m.SetWidth(600)
	m.EnabledMaximize(false)
	m.WorkAreaCenter()
	m.SetCaption("Energy - headless")

	defaultURL := "https://gitee.com/energye/energy"
	//保存目录
	savePath, _ := os.Getwd()

	// ----- UI控件功能 begin
	pdfPathEit := lcl.NewEdit(m)
	pdfPathEit.SetParent(m)
	pdfPathEit.SetTop(60)
	pdfPathEit.Font().SetColor(colors.ClBlue)
	pdfPathEit.Font().SetStyle(types.FsBold)
	pdfPathEit.Font().SetSize(14)
	pdfPathEit.SetWidth(m.Width())
	pdfPathEit.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	pdfPathEit.SetOnClick(func(sender lcl.IObject) {
		//点击复制到剪切板
		lcl.Clipboard.SetTextBuf(pdfPathEit.Text())
		fmt.Println("已复制到剪切板:", pdfPathEit.Text())
	})

	// 标签输入框
	urlAddr := lcl.NewLabeledEdit(m)
	urlAddr.SetParent(m)
	urlAddr.EditLabel().SetCaption("URL:")
	urlAddr.SetBounds(30, 15, 300, 25)
	urlAddr.SetLabelPosition(types.LpLeft)
	urlAddr.SetText(defaultURL)

	// 加载url按钮
	loadURLBtn := lcl.NewButton(m)
	loadURLBtn.SetParent(m)
	loadURLBtn.SetBounds(330, 10, 100, 35)
	loadURLBtn.SetCaption("加载 URL")
	loadURLBtn.SetOnClick(func(sender lcl.IObject) {
		if m.headless.browserCreated {
			m.headless.chromium.LoadUrl(urlAddr.Text())
		}
	})
	// 打印PDF按钮
	printPDFBtn := lcl.NewButton(m)
	printPDFBtn.SetParent(m)
	printPDFBtn.SetBounds(430, 10, 100, 35)
	printPDFBtn.SetCaption("打印页面 > PDF")
	printPDFBtn.SetOnClick(func(sender lcl.IObject) {
		if m.headless.browserCreated {
			name := fmt.Sprintf("page_print_%v.pdf", time.Now().Nanosecond())
			saveFilePath := filepath.Join(savePath, name)
			m.headless.chromium.PrintToPDF(saveFilePath)
			if m.headless.printPDFCallback != nil {
				printPDFBtn.SetEnabled(false)
				go func() {
					ok := <-m.headless.printPDFCallback
					fmt.Println("打印完成:", ok, "保存目录:", saveFilePath)
					cef.QueueAsyncCall(func(id int) { // 在UI线程修改控件状态
						printPDFBtn.SetEnabled(true)
						pdfPathEit.SetText(saveFilePath)
					})
				}()
			}
		}
	})

	// 当前页面加载进度显示标签
	m.pageLoadProgressLbl = lcl.NewLabel(m)
	m.pageLoadProgressLbl.SetParent(m)
	m.pageLoadProgressLbl.SetCaption("---")
	m.pageLoadProgressLbl.SetBounds(530, 10, 50, 35)
	m.pageLoadProgressLbl.Font().SetColor(colors.ClBlue)
	m.pageLoadProgressLbl.Font().SetStyle(types.FsBold)
	m.pageLoadProgressLbl.Font().SetSize(18)
	// ----- UI控件功能 end

	/*
		这个示例使用一个窗口, 因为有一些按钮功能, 简单控制浏览器, 当然没写那么多.
		也可以隐藏掉这个主窗口，但是以下的创建方式需要改变
	*/

	// 1. 创建无窗口的浏览器, 完全手动创建, 没浏览器那么方便一个命令完事
	m.headless = CreateHeadlessBrowser(m, defaultURL)
	// 2. 必须在Activate或Show事件内创建浏览器
	m.SetOnActivate(func(sender lcl.IObject) {
		m.headless.CreateBrowser(nil)
	})
	// 3. 在关闭窗口时关闭这个chromium
	m.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		m.headless.chromium.CloseBrowser(true)
	})
}

// HeadlessBrowser 定义一个无头浏览器结构体
type HeadlessBrowser struct {
	timer            *lcl.TTimer
	chromium         cef.IChromium
	browserCreated   bool
	printPDFCallback chan bool
}

// CreateHeadlessBrowser 创建一个无头浏览器，不在UI界面渲染
func CreateHeadlessBrowser(owner *MainWindow, url string) *HeadlessBrowser {
	/**
	这个示例里的浏览器未指定父组件，也就无法呈现界面
	*/
	m := new(HeadlessBrowser)
	m.printPDFCallback = make(chan bool)
	// 定时器，创建浏览器
	m.timer = lcl.NewTimer(owner)       // 创建定时器
	m.timer.SetInterval(100)            // 100 毫秒执行一次
	m.timer.SetOnTimer(m.CreateBrowser) // 定时器执行事件

	// 创建Chromium对象, 参数都为空了
	m.chromium = cef.NewChromium(nil, nil)
	m.chromium.SetDefaultURL(url)
	// 页面加载事件
	m.chromium.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
		val := int(progress * 100)
		fmt.Println("页面加载 browserId:", browser.Identifier(), "进度:", val)
		cef.QueueAsyncCall(func(id int) { // 在UI线程修改控件状态
			owner.pageLoadProgressLbl.SetCaption(fmt.Sprintf("%v", val))
		})
	})
	// PDF 打印完成事件
	m.chromium.SetOnPdfPrintFinished(func(sender lcl.IObject, ok bool) {
		fmt.Println("OnPdfPrintFinished:", ok)
		if m.printPDFCallback != nil {
			m.printPDFCallback <- ok
		}
	})
	return m
}

func (m *HeadlessBrowser) CreateBrowser(sender lcl.IObject) {
	if !m.browserCreated {
		m.timer.SetEnabled(false)
		m.chromium.Initialized()
		if !m.chromium.CreateBrowser(nil, "", nil, nil) {
			m.timer.SetEnabled(true)
		} else {
			m.browserCreated = true
		}
	}
}
