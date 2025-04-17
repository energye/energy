package form

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"path/filepath"
	"sync"
)

var (
	rootCachePath string
)

func SetRootCachePath(v string) {
	rootCachePath = v
}

// CreateComponent 动态创建组件
// 在此示例中，运用了 LCL + CEF 创建的丰富窗口
// LCL 原生系统UI组件，注意目前它在Linux下仅支持GTK2，但是GTK2目前还无法输入中文
// 在创建browser过程windows和其它系统有些区别，非Windows要在browser的父组件show 或 active事件中创建, 同时要在Resize实现
// 如果操作任何控制UI方面组件除CEF页面内(html渲染), 正确使用都要在主线程中执行（RunOnMainThread）函数
func CreateComponent(window cef.IBrowserWindow) {
	// 在这里改变主窗口的默认属性
	bw := window.AsLCLBrowserWindow().BrowserWindow()
	page := windowBottomLayout(bw)
	windowTopLayout(bw, page)
}

type tabBrowser struct {
	tab      *lcl.TTabSheet
	chromium cef.ICEFChromiumBrowser
	isClose  bool
	activeFn func()
}

var (
	tabs            = map[string]*tabBrowser{}
	topHeight int32 = 100
	lock      sync.Mutex
	closing   bool // 有正在关闭的
)
var (
	switchCacheRadTrue  *lcl.TRadioButton
	switchCacheRadFalse *lcl.TRadioButton
)

// 窗口上面
func windowTopLayout(window *cef.LCLBrowserWindow, page *lcl.TPageControl) {
	// 主窗口创建一个顶部panel
	topPanel := lcl.NewPanel(window)
	topPanel.SetParent(window)
	topPanel.SetHeight(topHeight)
	topPanel.SetWidth(window.Width())
	// 使panel自动根据窗口调整大小
	topPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	// 创建一些功能组件
	//
	linkLabel := lcl.NewLinkLabel(window)
	linkLabel.SetParent(topPanel)
	linkLabel.SetAlign(types.AlRight)
	linkLabel.SetCaption("<a href=\"https://forum.yanghy.cn/\">点击在浏览器打开链接</a>")
	linkLabel.SetOnLinkClick(func(sender lcl.IObject, link string, linktype types.TSysLinkType) {
		rtl.SysOpen(link)
	})
	// add browser
	newTabBtn := lcl.NewButton(window)
	newTabBtn.SetParent(topPanel)
	newTabBtn.SetCaption(" + AddTab ")
	newTabBtn.SetTop(5)
	newTabBtn.SetLeft(10)
	newTabBtn.SetWidth(100)
	newTabBtn.SetOnClick(func(sender lcl.IObject) {
		newTabBrowser(window, page)
	})
	// remove browser
	removeTabBtn := lcl.NewButton(window)
	removeTabBtn.SetParent(topPanel)
	removeTabBtn.SetCaption(" − RemoveTab ")
	removeTabBtn.SetTop(5)
	removeTabBtn.SetLeft(120)
	removeTabBtn.SetWidth(100)
	removeTabBtn.SetOnClick(func(sender lcl.IObject) {
		active := page.ActivePage()
		name := active.Name()
		if tab, ok := tabs[name]; ok && tab != nil {
			tab.isClose = true
			closing = true
			// 关闭当前激活的tab browser
			tab.chromium.Chromium().CloseBrowser(false)
			delete(tabs, name)
		}
	})
	//
	editUrl := lcl.NewEdit(window)
	editUrl.SetParent(topPanel)
	editUrl.SetTop(5)
	editUrl.SetLeft(230)
	editUrl.SetWidth(250)
	editUrl.SetText("https://gitee.com/energye/energy") // 写个固定的吧
	allURL := lcl.NewButton(window)
	allURL.SetParent(topPanel)
	allURL.SetCaption("所有Tab同一地址")
	allURL.SetTop(35)
	allURL.SetLeft(230)
	allURL.SetWidth(130)
	allURL.SetOnClick(func(sender lcl.IObject) {
		if closing {
			return
		}
		for _, tab := range tabs {
			if tab != nil && !tab.isClose {
				tab.chromium.Chromium().LoadUrl(editUrl.Text())
			}
		}
	})
	// 切换共享分离缓存目录
	switchCacheRadTrue = lcl.NewRadioButton(window)
	switchCacheRadTrue.SetParent(topPanel)
	switchCacheRadTrue.SetLeft(550)
	switchCacheRadTrue.SetCaption("启用共享缓存")
	switchCacheRadFalse = lcl.NewRadioButton(window)
	switchCacheRadFalse.SetParent(topPanel)
	switchCacheRadFalse.SetLeft(550)
	switchCacheRadFalse.SetTop(20)
	switchCacheRadFalse.SetCaption("关闭共享缓存")
	switchCacheRadFalse.SetChecked(true)
}

// 创建一个tabSheet, 并做为浏览器的父组件
func newTabBrowser(window *cef.LCLBrowserWindow, page *lcl.TPageControl) {
	lock.Lock()
	defer lock.Unlock()
	if closing { //有正在关闭的不去创建
		return
	}
	tabSheet := lcl.NewTabSheet(window)
	tabSheet.SetPageControl(page)
	tabSheet.SetCaption("[New Browser]")
	// 使用tab+index做为browser目录
	name := fmt.Sprintf("tab%v", tabSheet.TabIndex())
	tabSheet.SetName(name)
	tab := &tabBrowser{tab: tabSheet}
	// 创建 RequestContextSettings，在这里设置CachePath缓存路径
	// 根据下一个窗口序号ID, CEF默认会给每个browser分配一个唯一的ID，我们应该使用它，但是只有创建之后才能获得，所以此时还拿不到, 使用 tabIndex代替
	var context *cef.ICefRequestContext
	if switchCacheRadFalse.Checked() {
		settings := cef.TCefRequestContextSettings{}
		settings.CachePath = filepath.Join(rootCachePath, fmt.Sprintf("sing_%v", name))
		settings.PersistSessionCookies = 1 //bool: 0 或 1
		context = cef.RequestContextRef.New(settings, nil)
	}
	// 创建功能按钮
	boxTop := lcl.NewPanel(window)
	boxTop.SetParent(tabSheet)
	boxTop.SetHeight(35)
	boxTop.SetWidth(window.Width())
	boxTop.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	// 创建工具栏按钮
	back, forward, stop, refresh, goUrl, progressLabel, addr := toolBar(window, boxTop)

	// 浏览器部分
	boxBottom := lcl.NewPanel(window)
	boxBottom.SetParent(tabSheet)
	boxBottom.SetTop(35)
	boxBottom.SetHeight(tabSheet.Height() - 35)
	boxBottom.SetWidth(window.Width())
	boxBottom.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))
	boxBottom.SetBorderStyle(types.BsNone)

	// 创建Chromium Browser， 将 tabSheet 做为父组件
	chromium := cef.NewChromiumBrowser(boxBottom, nil)
	chromium.SetSelfWindow(window)
	chromium.RegisterDefaultEvent()
	tab.chromium = chromium
	chromium.WindowParent().SetAlign(types.AlClient)
	// 请求上下文配置
	if context != nil {
		chromium.SetCreateBrowserExtraInfo("", context, nil)
	}
	chromium.Chromium().SetDefaultURL("https://www.baidu.com") // 设置加载的页面地址
	// 通过chromium获取回调事件
	chromium.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
		// 在这里更新 WindowParent 大小, 以保证渲染到窗口中
		cef.RunOnMainThread(func() {
			chromium.WindowParent().UpdateSize()
		})
		return false
	})
	chromium.Chromium().SetOnRequestContextInitialized(func(sender lcl.IObject, requestContext *cef.ICefRequestContext) {
		fmt.Println("当前缓存目录:", requestContext.GetCachePath())

	})
	chromium.Chromium().SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
		if !tab.isClose {
			cef.RunOnMainThread(func() {
				tabSheet.SetCaption(title) // 页面中title改变， 同时改变tab的标题
			})
		}
	})
	// 在popup事件里，打开新地址, 不在新窗口，而是在当前页面打开新地址
	chromium.Chromium().SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
		if !tab.isClose {
			frame.LoadUrl(beforePopupInfo.TargetUrl)
		}
		return true // 不创建窗口
	})
	chromium.Chromium().SetOnClose(func(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
		*aAction = consts.CbaDelay
		cef.RunOnMainThread(func() { //run in main thread
			chromium.WindowParent().Free()
			tabSheet.Free()
			closing = false
		})
	}) //页面加载处理进度
	chromium.Chromium().SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
		//linux 更新UI组件必须使用 QueueAsyncCall 主线程异步同步
		cef.QueueAsyncCall(func(id int) {
			//参数-进度
			progressLabel.SetCaption(fmt.Sprintf("%v", progress*100))
		})
	})
	//页面加载状态，根据状态判断是否加载完成，和是否可前进后退
	chromium.Chromium().SetOnLoadingStateChange(func(sender lcl.IObject, browser *cef.ICefBrowser, isLoading, canGoBack, canGoForward bool) {
		//linux 更新UI组件必须使用 QueueAsyncCall 主线程异步同步
		cef.QueueAsyncCall(func(id int) {
			//控制按钮状态
			stop.SetEnabled(isLoading)
			refresh.SetEnabled(!isLoading)
			back.SetEnabled(canGoBack)
			forward.SetEnabled(canGoForward)
		})
	})
	chromium.Chromium().SetOnAddressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, url string) {
		cef.QueueAsyncCall(func(id int) {
			addr.SetText(url)
		})
	})
	//给按钮增加事件
	back.SetOnClick(func(sender lcl.IObject) {
		chromium.Chromium().GoBack()
	})
	forward.SetOnClick(func(sender lcl.IObject) {
		chromium.Chromium().GoForward()
	})
	stop.SetOnClick(func(sender lcl.IObject) {
		chromium.Chromium().StopLoad()
	})
	refresh.SetOnClick(func(sender lcl.IObject) {
		chromium.Chromium().Reload()
	})
	goUrl.SetOnClick(func(sender lcl.IObject) {
		var url = addr.Text()
		if url != "" && !tab.isClose {
			chromium.Chromium().LoadUrl(url)
		}
	})
	// 设置完chromium事件后再创建浏览器
	if common.IsWindows() {
		// windows 可以直接在这里创建
		chromium.CreateBrowser() // 创建浏览器
	}
	tabSheet.SetOnResize(func(sender lcl.IObject) {
		chromium.Chromium().NotifyMoveOrResizeStarted()
		if chromium.WindowParent() != nil {
			chromium.WindowParent().UpdateSize()
		}
	})
	// CreateBrowser For Linux, MacOS
	tabSheet.SetOnShow(func(sender lcl.IObject) {
		// 创建过不会重复创建
		chromium.CreateBrowser()
	})
	// 最后 激活当前这个 tab
	page.SetActivePage(tabSheet)
	// 将对象缓存到map
	tabs[name] = tab // 这个示例，主窗口不在这里关闭
}

// 照搬 control-widget 示例
func toolBar(window *cef.LCLBrowserWindow, toolPanel *lcl.TPanel) (goBack *lcl.TButton, goForward *lcl.TButton, stop *lcl.TButton, refresh *lcl.TButton, goUrl *lcl.TButton, progressLabel *lcl.TLabel, addrBox *lcl.TComboBox) {
	//创建 按钮-后退
	goBack = lcl.NewButton(toolPanel) //设置父组件
	goBack.SetParent(toolPanel)
	goBack.SetCaption("后退")
	goBack.SetBounds(5, 3, 35, 25)
	goForward = lcl.NewButton(toolPanel) //设置父组件
	goForward.SetParent(toolPanel)
	goForward.SetCaption("前进")
	goForward.SetBounds(45, 3, 35, 25)
	stop = lcl.NewButton(toolPanel) //设置父组件
	stop.SetParent(toolPanel)
	stop.SetCaption("停止")
	stop.SetBounds(90, 3, 35, 25)
	refresh = lcl.NewButton(toolPanel) //设置父组件
	refresh.SetParent(toolPanel)
	refresh.SetCaption("刷新")
	refresh.SetBounds(135, 3, 35, 25)

	//创建下拉框
	addrBox = lcl.NewComboBox(toolPanel)
	addrBox.SetParent(toolPanel)
	addrBox.SetLeft(180)                                                       //这里是设置左边距 上面按钮的宽度
	addrBox.SetTop(3)                                                          //
	addrBox.SetWidth(window.Width() - (280))                                   //宽度 减按钮的宽度
	addrBox.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小
	addrBox.Items().Add("https://forum.yanghy.cn/")
	addrBox.Items().Add("https://energy.yanghy.cn")
	addrBox.Items().Add("https://www.csdn.net")
	addrBox.Items().Add("https://www.baidu.com")

	//显示加载进度
	progressLabel = lcl.NewLabel(toolPanel) //设置父组件
	progressLabel.SetParent(toolPanel)
	progressLabel.SetCaption("0")
	progressLabel.SetBounds(window.Width()-100, 5, 35, 25)
	progressLabel.SetAnchors(types.NewSet(types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小

	goUrl = lcl.NewButton(toolPanel) //设置父组件
	goUrl.SetParent(toolPanel)
	goUrl.SetCaption("GO")
	goUrl.SetBounds(window.Width()-50, 3, 35, 25)
	goUrl.SetAnchors(types.NewSet(types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小
	return
}

// window窗口下面, 特殊一点吧，放两个browser
func windowBottomLayout(window *cef.LCLBrowserWindow) *lcl.TPageControl {
	lock.Lock()
	defer lock.Unlock()
	bottomBoxPanel := lcl.NewPanel(window)
	bottomBoxPanel.SetParent(window)
	// 设置窗口下面的panel位置和宽高，要把上面panel高算上
	bottomBoxPanel.SetTop(topHeight)
	bottomBoxPanel.SetHeight(window.Height() - topHeight)
	bottomBoxPanel.SetWidth(window.Width())
	// 使panel自动根据窗口调整大小
	bottomBoxPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))

	// 创建Tab Page, 因为是多浏览器示例所以这里使用tab sheet
	page := lcl.NewPageControl(bottomBoxPanel)
	page.SetParent(bottomBoxPanel)
	page.SetAlign(types.AlClient) //大小自动调整和 bottomBoxPanel 一样
	// 创建一个tab sheet, 做为主窗口浏览器的父组件
	tabSheet := lcl.NewTabSheet(page)
	tabSheet.SetPageControl(page)
	tabSheet.SetCaption("主窗口")
	//改变主窗口浏览器组件的父组件
	//window.WindowParent().RevertCustomAnchors()

	// 分隔两个组件 添加顺序 从左到右
	splitter := lcl.NewSplitter(window)
	splitter.SetParent(tabSheet)
	splitter.SetAlign(types.AlLeft)
	splitter.SetCursor(types.CrHSplit)
	splitter.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkBottom))
	splitter.SetWidth(5)
	// 左半部分
	left := lcl.NewPanel(window)
	left.SetParent(tabSheet)
	left.SetAlign(types.AlLeft)
	left.SetWidth(window.Width() / 2)
	window.WindowParent().SetParent(left)
	// 右半部分
	right := lcl.NewPanel(window)
	right.SetParent(tabSheet)
	right.SetAlign(types.AlClient)
	right.SetColor(colors.ClBackground)
	// 创建右半部分浏览器
	chromium := cef.NewChromiumBrowser(right, nil)
	chromium.SetSelfWindow(window)  // 设置当前窗口，让默认事件全部生效
	chromium.RegisterDefaultEvent() // 注册默认实现事件
	chromium.WindowParent().SetAlign(types.AlClient)
	chromium.Chromium().SetDefaultURL("https://www.baidu.com") // 设置加载的页面地址
	chromium.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
		// 在这里更新 WindowParent 大小, 以保证渲染到窗口中
		chromium.WindowParent().UpdateSize()
		fmt.Println("SetOnBeforeBrowser")
		return false
	})
	if common.IsWindows() {
		// windows 可以直接在这里创建
		chromium.CreateBrowser() // 创建浏览器
	}
	// 改变窗口大小时
	window.SetOnResize(func(sender lcl.IObject) bool {
		chromium.Chromium().NotifyMoveOrResizeStarted()
		if chromium.WindowParent() != nil {
			chromium.WindowParent().UpdateSize()
		}
		return false
	})
	// CreateBrowser For Linux, MacOS
	window.SetOnActivateAfter(func(sender lcl.IObject) {
		// 创建过不会重复创建
		chromium.CreateBrowser()
	})
	return page
}
