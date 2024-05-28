package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/api/libname"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/messages"
	"github.com/energye/energy/v2/wv"
	"path/filepath"
)

type TMainForm struct {
	lcl.TForm
	windowParent wv.IWVWindowParent
	browser      wv.IWVBrowser
}

var MainForm TMainForm
var load wv.IWVLoader
var scheme = "myscheme"

//go:embed resources
var resources embed.FS

func main() {
	fmt.Println("Go ENERGY Run Main")
	wv.Init(nil, nil)
	// GlobalWebView2Loader
	load = wv.GlobalWebView2Loader(nil)
	liblcl := libname.LibName
	webView2Loader, _ := filepath.Split(liblcl)
	webView2Loader = filepath.Join(webView2Loader, "WebView2Loader.dll")
	fmt.Println("当前目录:", types.CurrentExecuteDir)
	fmt.Println("liblcl.dll目录:", liblcl)
	fmt.Println("WebView2Loader.dll目录:", webView2Loader)
	fmt.Println("用户缓存目录:", filepath.Join(types.CurrentExecuteDir, "EnergyCache"))
	fmt.Println("自定义URL协议头:", scheme)
	load.SetUserDataFolder(filepath.Join(types.CurrentExecuteDir, "EnergyCache"))
	load.SetLoaderDllPath(webView2Loader)
	load.SetOnGetCustomSchemes(func(sender wv.IObject, customSchemes *wv.TWVCustomSchemeInfoArray) {
		fmt.Println("回调函数 WebView2Loader => SetOnGetCustomSchemes")
		*customSchemes = append(*customSchemes, &wv.TWVCustomSchemeInfo{
			SchemeName:            scheme,
			TreatAsSecure:         true,
			AllowedDomains:        "https://*.baidu.com,https://*.yanghy.cn",
			HasAuthorityComponent: true,
		})
	})
	r := load.StartWebView2()
	fmt.Println("StartWebView2", r)

	lcl.Application.SetOnException(func(sender lcl.IObject, e lcl.IException) {
		fmt.Println("底层库异常:", e.ToString())
	})
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.Run()
}

func (m *TMainForm) FormCreate(sender lcl.IObject) {
	m.SetCaption("Energy3.0 - webview2 simple")
	m.SetPosition(types.PoScreenCenter)
	m.SetWidth(1024)
	m.SetHeight(768)
	m.SetDoubleBuffered(true)
	back, forward, stop, refresh, addr := controlUI(m)

	m.windowParent = wv.NewWVWindowParent(m)
	m.windowParent.SetParent(m)
	//m.windowParent.SetWidth(200)
	//m.windowParent.SetHeight(200)
	//重新调整browser窗口的Parent属性
	//重新设置了上边距，宽，高
	m.windowParent.SetAlign(types.AlCustom) //重置对齐,默认是整个客户端
	m.windowParent.SetTop(30)
	m.windowParent.SetHeight(m.Height() - 25)
	m.windowParent.SetWidth(m.Width())
	m.windowParent.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))
	//m.windowParent.SetAlign(types.AlClient)

	m.browser = wv.NewWVBrowser(m)
	m.browser.SetDefaultURL("myscheme://domain/CustomScheme.html")
	//m.browser.SetDefaultURL("https://www.baidu.com")
	m.browser.SetTargetCompatibleBrowserVersion("95.0.1020.44")
	fmt.Println("TargetCompatibleBrowserVersion:", m.browser.TargetCompatibleBrowserVersion())
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		fmt.Println("回调函数 WVBrowser => SetOnAfterCreated")
		m.windowParent.UpdateSize()
		m.browser.AddWebResourceRequestedFilter(scheme+"*", wv.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
	})
	var navBtns = func(aIsNavigating bool) {
		back.SetEnabled(m.browser.CanGoBack())
		forward.SetEnabled(m.browser.CanGoForward())
		refresh.SetEnabled(!aIsNavigating)
		stop.SetEnabled(aIsNavigating)
	}
	m.browser.SetOnNavigationCompleted(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2NavigationCompletedEventArgs) {
		fmt.Println("回调函数 WVBrowser => SetOnNavigationCompleted")
		navBtns(false)
	})
	m.browser.SetOnNavigationStarting(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2NavigationStartingEventArgs) {
		fmt.Println("回调函数 WVBrowser => SetOnNavigationStarting")
		//args = wv.NewCoreWebView2NavigationStartingEventArgs(args)
		navBtns(true)
		//args.Free()
	})
	m.browser.SetOnWebMessageReceived(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2WebMessageReceivedEventArgs) {
		fmt.Println("回调函数 WVBrowser => SetOnWebMessageReceived")
		args = wv.NewCoreWebView2WebMessageReceivedEventArgs(args)
		fmt.Println(args.WebMessageAsString())
		args.Free()
	})
	m.browser.SetOnSourceChanged(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2SourceChangedEventArgs) {
		fmt.Println("回调函数 WVBrowser => SetOnSourceChanged")
		addr.SetText(m.browser.Source())
	})
	m.browser.SetOnContentLoading(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2ContentLoadingEventArgs) {
		fmt.Println("回调函数 WVBrowser => SetOnContentLoading")
	})
	m.browser.SetOnContextMenuRequested(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2ContextMenuRequestedEventArgs) {
		webView = wv.NewCoreWebView2(webView)
		args = wv.NewCoreWebView2ContextMenuRequestedEventArgs(args)
		menuItems := wv.NewCoreWebView2ContextMenuItemCollection(args.MenuItems())
		contextMenuTarge := wv.NewCoreWebView2ContextMenuTarget(args.ContextMenuTarget())
		fmt.Println("回调函数 WVBrowser => SetOnContextMenuRequested:", menuItems.Count(), contextMenuTarge.PageUri(), webView.BrowserProcessID(), webView.FrameId())
		fmt.Println("回调函数 WVBrowser => SelectedCommandId:", args.SelectedCommandId())
		menuItems.Free()
		contextMenuTarge.Free()
		args.Free()
	})
	m.browser.SetOnDocumentTitleChanged(func(sender lcl.IObject) {
		fmt.Println("回调函数 WVBrowser => SetOnDocumentTitleChanged:", m.browser.DocumentTitle())
	})
	m.browser.SetOnFrameCreated(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2FrameCreatedEventArgs) {
		fmt.Println("回调函数 WVBrowser => SetOnFrameCreated")
		args = wv.NewCoreWebView2FrameCreatedEventArgs(args)
		webView = wv.NewCoreWebView2(webView)
		frame := wv.NewCoreWebView2Frame(args.Frame())
		fmt.Println("回调函数 WVBrowser => SetOnFrameCreated", webView.FrameId(), frame.FrameID())
		args.Free()
		frame.Free()
		webView.Free()
	})
	m.browser.SetOnGetCustomSchemes(func(sender wv.IObject, customSchemes wv.TWVCustomSchemeInfoArray) {
		fmt.Println("回调函数 WVBrowser => SetOnGetCustomSchemes")
	})
	m.browser.SetOnWebResourceRequested(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2WebResourceRequestedEventArgs) {
		args = wv.NewCoreWebView2WebResourceRequestedEventArgs(args)
		request := wv.NewCoreWebView2WebResourceRequestRef(args.Request())
		fmt.Println("回调函数 WVBrowser => SetOnWebResourceRequested")
		fmt.Println("回调函数 WVBrowser => TempURI:", request.URI(), request.Method())
		fmt.Println("回调函数 WVBrowser => 内置exe读取 index.html ")
		data, _ := resources.ReadFile("resources/index.html")
		stream := lcl.NewMemoryStream()
		stream.LoadFromBytes(data)
		fmt.Println("回调函数 WVBrowser => stream", stream.Size())
		adapter := lcl.NewStreamAdapter(stream, types.SoOwned)
		fmt.Println("回调函数 WVBrowser => adapter:", adapter.StreamOwnership(), adapter.Stream().Size())

		var response wv.ICoreWebView2WebResourceResponse
		environment := m.browser.CoreWebView2Environment()
		fmt.Println("回调函数 WVBrowser => Initialized():", environment.Initialized(), environment.BrowserVersionInfo())
		environment.CreateWebResourceResponse(adapter, 200, "OK", "Content-Type: text/html", &response)
		args.SetResponse(response)

		// 需要释放掉
		request.Free()
		args.Free()
	})
	m.windowParent.SetBrowser(m.browser)

	m.SetOnShow(func(sender lcl.IObject) {
		if load.InitializationError() {
			fmt.Println("回调函数 => SetOnShow 初始化失败")
		} else {
			if load.Initialized() {
				fmt.Println("回调函数 => SetOnShow 初始化成功")
				m.browser.CreateBrowser(m.windowParent.Handle(), true)
			}
		}
	})
	m.SetOnWndProc(func(msg *types.TMessage) {
		m.InheritedWndProc(msg)
		switch msg.Msg {
		case messages.WM_SIZE, messages.WM_MOVE:
			m.browser.NotifyParentWindowPositionChanged()
		}
	})
}

// 控制组件UI
// 地址栏和控制按钮创建
func controlUI(window *TMainForm) (goBack lcl.IButton, goForward lcl.IButton, stop lcl.IButton, refresh lcl.IButton, addrBox lcl.IComboBox) {
	//这里使用系统UI组件
	//创建panel做为地址栏的父组件
	addrPanel := lcl.NewPanel(window) //设置父组件
	addrPanel.SetParent(window)
	addrPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小
	addrPanel.SetHeight(30)
	addrPanel.SetWidth(window.Width())
	//创建 按钮-后退
	goBack = lcl.NewButton(addrPanel) //设置父组件
	goBack.SetParent(addrPanel)
	goBack.SetCaption("后退")
	goBack.SetBounds(5, 3, 35, 25)
	goForward = lcl.NewButton(addrPanel) //设置父组件
	goForward.SetParent(addrPanel)
	goForward.SetCaption("前进")
	goForward.SetBounds(45, 3, 35, 25)
	stop = lcl.NewButton(addrPanel) //设置父组件
	stop.SetParent(addrPanel)
	stop.SetCaption("停止")
	stop.SetBounds(90, 3, 35, 25)
	refresh = lcl.NewButton(addrPanel) //设置父组件
	refresh.SetParent(addrPanel)
	refresh.SetCaption("刷新")
	refresh.SetBounds(135, 3, 35, 25)

	//创建下拉框
	addrBox = lcl.NewComboBox(addrPanel)
	addrBox.SetParent(addrPanel)
	addrBox.SetLeft(180)                                                       //这里是设置左边距 上面按钮的宽度
	addrBox.SetTop(3)                                                          //
	addrBox.SetWidth(window.Width() - (230))                                   //宽度 减按钮的宽度
	addrBox.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小
	addrBox.Items().Add("myscheme://domain/CustomScheme.html")
	addrBox.Items().Add("https://gitee.com/energye/energy")
	addrBox.Items().Add("https://github.com/energye/energy")
	addrBox.Items().Add("https://www.baidu.com")
	addrBox.Items().Add("https://energy.yanghy.cn")
	addrBox.SetText("myscheme://domain/CustomScheme.html")

	goUrl := lcl.NewButton(addrPanel) //设置父组件
	goUrl.SetParent(addrPanel)
	goUrl.SetCaption("GO")
	goUrl.SetBounds(window.Width()-45, 3, 40, 25)
	goUrl.SetAnchors(types.NewSet(types.AkTop, types.AkRight)) //设置锚点定位，让宽高自动根据窗口调整大小

	//给按钮增加事件
	goBack.SetOnClick(func(sender lcl.IObject) {
		window.browser.GoBack()
	})
	goForward.SetOnClick(func(sender lcl.IObject) {
		window.browser.GoForward()
	})
	stop.SetOnClick(func(sender lcl.IObject) {
		window.browser.Stop()
	})
	refresh.SetOnClick(func(sender lcl.IObject) {
		window.browser.Refresh()
	})
	goUrl.SetOnClick(func(sender lcl.IObject) {
		var url = addrBox.Text()
		if url != "" {
			window.browser.Navigate(url)
		}
	})
	return
}
