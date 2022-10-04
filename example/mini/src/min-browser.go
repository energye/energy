package src

import (
	"encoding/json"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/tools"
	"strings"
)

func AppBrowserInit() {
	/*
		通用类型变量和对象类型变量创建的回调函数
		该回调函数-在主进程和渲染进程创建时调用
	*/
	//ar integer cef.JSValue
	cef.VariableBind.VariableCreateCallback(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, bind cef.IProvisionalBindStorage) {
		fmt.Println("GO变量和函数绑定回调", cef.Args.ProcessType())
		bind.NewInteger("integerv", 1211111)
		bind.NewDouble("doublev", 11.0505)
		bind.NewBoolean("booleanv", true)
		bind.NewNull("nullv")
		bind.NewUndefined("undefinedv")
	})

	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)
	config.SetEnableViewSource(true)
	config.SetEnableDevTools(true)
	config.SetEnableWindowPopup(true)
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//默认加载的URL
	if cef.IsWindows() {
		cef.BrowserWindow.Config.DefaultUrl = "E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\demo-golang-dll-01-chromium\\demos\\min-browser\\resources\\demo-misc.html"
		if !tools.IsExist(cef.BrowserWindow.Config.DefaultUrl) {
			cef.BrowserWindow.Config.DefaultUrl = cef.ExePath + cef.Separator + "demo-misc.html"
		}
	} else if cef.IsLinux() {
		cef.BrowserWindow.Config.DefaultUrl = "file:///home/sxm/app/swt/gopath/src/github.com/energye/energy/demos/min-browser/resources/demo-misc.html"
	} else if cef.IsDarwin() {
		cef.BrowserWindow.Config.DefaultUrl = "file:///Users/zhangli/go/src/github.com/energye/energy/demos/min-browser/resources/demo-misc.html"
	}
	//主进程 IPC事件
	cef.IPC.Browser().SetOnEvent(func(event cef.IEventOn) {
		fmt.Println("主进程IPC事件注册")
		//这个事件监听演示了几个示例
		//1.
		event.On("subWindowIPCOn", func(context cef.IIPCContext) {
			//函数内的 context 返回给调用方
			//取入参
			fmt.Println("browser renderOnEventSubWindowIPCOn:", context.Arguments().GetString(0), "channelId:", context.ChannelId())
			fmt.Println("\t|--", len(cef.BrowserWindow.GetWindowsInfo()))
			//调用指定渲染进程监听
			cef.IPC.Browser().EmitChannelId("renderOnEventSubWindowIPCOn", context.ChannelId(), nil)
			context.Result().SetString("\t|-- 111")

			//调用指定渲染进程监听 回调的方式
			cef.IPC.Browser().EmitChannelIdAndCallback("renderOnEventSubWindowIPCOn", context.ChannelId(), nil, func(context cef.IIPCContext) {
				fmt.Println("browser renderOnEventSubWindowIPCOn ret:", string(context.Message().Data()))
				//在这里 不能使用 context 返回给页面
				//函数内的 context 返回给调用方
			})
			context.Result().SetString("\t|-- 222")
			//调用指定渲染进程监听 同步的方式
			ctx := cef.IPC.Browser().EmitChannelIdAndReturn("renderOnEventSubWindowIPCOn", context.ChannelId(), nil)
			context.Result().SetString("\t|-- 333")
			//返回给页面数据
			context.Result().SetString("成功返回, " + string(ctx.Message().Data()))

			//返回给页面数据
			//context.Result().SetString("成功返回,")
		})
		event.On("close", func(context cef.IIPCContext) {
			bsr := cef.BrowserWindow.GetBrowser(context.BrowserId())
			fmt.Println("close browserId:", context.BrowserId())
			//发送给渲染进程消息
			bsr.CloseBrowser(true)
		})
		event.On("minWindow", func(context cef.IIPCContext) {
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Minimize()
		})
		event.On("maxWindow", func(context cef.IIPCContext) {
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Maximize()
		})
		event.On("closeWindow", func(context cef.IIPCContext) {
			fmt.Println("closeWindow", context.BrowserId())
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Close()
		})
		event.On("window-list", func(context cef.IIPCContext) {
			var ids []int32
			for id, _ := range cef.BrowserWindow.GetWindowsInfo() {
				ids = append(ids, id)
			}
			idsStr, err := json.Marshal(ids)
			fmt.Println("获得window-id-list", ids, idsStr, err)
			context.Result().SetString(string(idsStr))
		})
		event.On("find", func(context cef.IIPCContext) {
			args := context.Arguments().GetString(0)
			fmt.Println("find-args:", args)
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Browser.Find(args, false, false, true)
		})
		event.On("find-stop", func(context cef.IIPCContext) {
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Browser.StopFinding(true)
		})
		var newForm *cef.Window
		event.On("js-new-window", func(context cef.IIPCContext) {
			fmt.Println("创建新窗口 ProcessType:", cef.Args.ProcessType())
			if newForm == nil {
				newForm = cef.NewWindow()
				newForm.SetCaption("新窗口标题")
				btn := lcl.NewButton(newForm)
				btn.SetParent(newForm)
				btn.SetCaption("点击我有提示")
				btn.SetWidth(100)
				btn.SetOnClick(func(sender lcl.IObject) {
					lcl.ShowMessage("新窗口的按钮事件提示")
				})
			}
			//linux 下必须使用 QueueAsyncCall
			//windows 也可以这样使用
			//macos 也可以这样使用
			//窗口设置或控制时需要在UI线程中操作
			cef.QueueAsyncCall(func(id int) {
				if newForm.Visible() {
					fmt.Println("隐藏")
					newForm.Hide()
				} else {
					fmt.Println("显示")
					newForm.Show()
				}
			})
		})
		var browserWindow *cef.Window
		event.On("js-new-browser-window", func(context cef.IIPCContext) {
			fmt.Println("通过 js ipc emit 事件创建新Browser窗口 ProcessType:", cef.Args.ProcessType())
			if browserWindow == nil {
				browserWindow = cef.NewBrowserWindow(nil, "https://www.baidu.com")
				browserWindow.Chromium().EnableIndependentEvent()
				browserWindow.SetCaption("Browser新窗口标题")
				browserWindow.SetWidth(800)
				browserWindow.SetHeight(600)
			}
			fmt.Println("\t|--", browserWindow.IsClosing())
			cef.QueueAsyncCall(func(id int) {
				if browserWindow.Visible() {
					browserWindow.Hide()
				} else {
					browserWindow.Show()
				}
			})
		})
	})

	//主进程初始化回调函数
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		lcl.Application.SetOnMinimize(func(sender lcl.IObject) {
			fmt.Println("minimize")
		})
		fmt.Println("主进程初始化回调函数")
		lcl.Application.Icon().LoadFromFSFile("resources/icon.ico") //设置应用图标
		browserWindow.Window.SetCaption("这里设置应用标题")
		browserWindow.Window.SetPosition(types.PoScreenCenter) //窗口局中显示
		//设置窗口样式，无标题 ，最大化按钮等
		//window.Window.SetBorderStyle(types.BsNone)
		//window.Form.SetFormStyle(types.FsNormal)
		//window.Form.SetFormStyle(types.FsSystemStayOnTop)
		//设置窗口大小
		browserWindow.Window.SetWidth(1600)
		browserWindow.Window.SetHeight(900)
		//限制窗口大小 linux 下不是很友好
		//browserWindow.Window.Constraints().SetMinHeight(300)
		//browserWindow.Window.Constraints().SetMinWidth(300)
		//browserWindow.Window.Constraints().SetMaxWidth(1600)
		//browserWindow.Window.Constraints().SetMaxHeight(900)
		browserWindow.Window.Constraints().SetOnChange(func(sender lcl.IObject) {
			fmt.Println("browserWindow SetOnChange")
		})
		//添加事件，add不会覆盖默认的事件 set会覆盖默认的事件
		browserWindow.Window.AddOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
			fmt.Println("添加onclose事件")
			return false
		})
		//窗口大小改变后触发
		browserWindow.Window.AddOnResize(func(sender lcl.IObject) bool {
			//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
			if browserWindow.Browser != nil {
				var target = &cef.GoEmitTarget{
					BrowseId: browserWindow.Browser.Identifier(),
					FrameId:  browserWindow.Browser.MainFrame().Id,
				}
				var argumentList = cef.NewArgumentList()
				argumentList.SetInt32(0, browserWindow.Window.Left())
				argumentList.SetInt32(1, browserWindow.Window.Top())
				argumentList.SetInt32(2, browserWindow.Window.Width())
				argumentList.SetInt32(3, browserWindow.Window.Height())
				browserWindow.Chromium().Emit("window-resize", argumentList, target)
				browserWindow.Chromium().EmitAndCallback("window-resize", argumentList, target, func(context cef.IIPCContext) {
					fmt.Println("EmitAndCallback AddOnResize")
				})
				//使用EmitAndReturn函数会锁死
				//ctx, _ := browserWindow.Chromium.EmitAndReturn("window-resize", argumentList, target)
				//fmt.Println("EmitAndReturn AddOnResize", ctx)
			}
			return false
		})
		//windows下可以使用这个函数，实时触发
		browserWindow.Window.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
			//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
			if browserWindow.Browser != nil {
				var target = &cef.GoEmitTarget{
					BrowseId: browserWindow.Browser.Identifier(),
					FrameId:  browserWindow.Browser.MainFrame().Id,
				}
				var argumentList = cef.NewArgumentList()
				argumentList.SetInt32(0, browserWindow.Window.Left())
				argumentList.SetInt32(1, browserWindow.Window.Top())
				argumentList.SetInt32(2, browserWindow.Window.Width())
				argumentList.SetInt32(3, browserWindow.Window.Height())
				fmt.Println("window-resize:", target, "count:", browserWindow.Browser.FrameCount(), "\n\t", browserWindow.Browser.GetFrameNames())
				browserWindow.Chromium().Emit("window-resize", argumentList, target)

				//使用EmitAndReturn函数会锁死
				//	browserWindow.Chromium.EmitAndCallback("window-resize", argumentList, target, func(context cef.IIPCContext) {
				//		fmt.Println("EmitAndCallback OnConstrainedResize")
				//	})

			}
		})
		//自定义browser窗体
		addressBar(browserWindow.Window)
		//下载文件
		//linux系统弹出保存对话框不启作用
		//自己调用系统的保存对话框获得保存路径
		dlSave := lcl.NewSaveDialog(browserWindow.Window)
		dlSave.SetTitle("保存对话框标题")
		event.SetOnBeforeDownload(func(sender lcl.IObject, browser *cef.ICefBrowser, beforeDownloadItem *cef.DownloadItem, suggestedName string, callback *cef.ICefBeforeDownloadCallback) {
			fmt.Println("OnBeforeDownload:", beforeDownloadItem, suggestedName)
			//linux下 需要这样使用 Sync
			if cef.IsLinux() {
				cef.QueueSyncCall(func(id int) {
					dlSave.SetFileName(suggestedName)
					if dlSave.Execute() {
						fmt.Println("OnBeforeDownload filename: ", dlSave.FileName())
						callback.Cont(dlSave.FileName(), false)
					}
				})
			} else {
				callback.Cont(cef.ExePath+cef.Separator+suggestedName, true)
			}
		})
		event.SetOnDownloadUpdated(func(sender lcl.IObject, browser *cef.ICefBrowser, downloadItem *cef.DownloadItem, callback *cef.ICefDownloadItemCallback) {
			fmt.Println("OnDownloadUpdated:", downloadItem)
		})

		event.SetOnFindResult(func(sender lcl.IObject, browser *cef.ICefBrowser, identifier, count int32, selectionRect *cef.TCefRect, activeMatchOrdinal int32, finalUpdate bool) {
			fmt.Println("OnFindResult:", browser.Identifier(), identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
		})
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow *cef.TCefWindowInfo, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup: " + beforePopupInfo.TargetUrl)
			popupWindow.Window.SetCaption("改变了标题 - " + beforePopupInfo.TargetUrl)
			//popupWindow.Form.SetBorderStyle(types.BsNone)
			//popupWindow.Form.SetFormStyle(types.FsNormal)
			popupWindow.Window.SetWidth(800)
			popupWindow.Window.SetHeight(600)
			//窗口弹出之前可自定义系统组件
			//窗口大小改变后触发
			//windows下窗口调整后触发一次
			popupWindow.Window.AddOnResize(func(sender lcl.IObject) bool {
				//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
				if popupWindow.Browser != nil {
					var target = &cef.GoEmitTarget{
						BrowseId: popupWindow.Browser.Identifier(),
						FrameId:  popupWindow.Browser.MainFrame().Id,
					}
					var argumentList = cef.NewArgumentList()
					argumentList.SetInt32(0, popupWindow.Window.Left())
					argumentList.SetInt32(1, popupWindow.Window.Top())
					argumentList.SetInt32(2, popupWindow.Window.Width())
					argumentList.SetInt32(3, popupWindow.Window.Height())
					popupWindow.Chromium().Emit("window-resize", argumentList, target)
				}
				return false
			})
			//windows下可以使用这个函数，实时触发
			popupWindow.Window.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
				//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
				if popupWindow.Browser != nil {
					var target = &cef.GoEmitTarget{
						BrowseId: popupWindow.Browser.Identifier(),
						FrameId:  popupWindow.Browser.MainFrame().Id,
					}
					var argumentList = cef.NewArgumentList()
					argumentList.SetInt32(0, popupWindow.Window.Left())
					argumentList.SetInt32(1, popupWindow.Window.Top())
					argumentList.SetInt32(2, popupWindow.Window.Width())
					argumentList.SetInt32(3, popupWindow.Window.Height())
					popupWindow.Chromium().Emit("window-resize", argumentList, target)
				}
			})
			return false
		})
		//加载页面之前
		var isSendEmit bool
		event.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame) {
			fmt.Println("OnLoadStart:", frame.Url)
			//判断一下哪些页面不发送emit消息
			if strings.LastIndex(strings.ToLower(frame.Url), ".pdf") > 0 || strings.Index(frame.Url, "about:blank") != -1 {
				isSendEmit = false
			} else {
				isSendEmit = true
			}
		})
		event.SetOnFrameCreated(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame) {
			fmt.Println("OnFrameCreated:", frame.Url)
		})
		event.SetOnLoadingStateChange(func(sender lcl.IObject, browser *cef.ICefBrowser, isLoading, canGoBack, canGoForward bool) {
			//当刷新的是一个完整的浏览器时，如果打开的新页面不是html dom，这里的 emit 消息 将会失败
			fmt.Println("OnLoadingStateChange-ProcessType:", cef.Args.ProcessType(), "sender.Instance:", sender.Instance(), "browserId:", browser.Identifier(), "isLoading:", isLoading, "canGoBack:", canGoBack, "canGoForward:", canGoForward)
			if isSendEmit {
				info := cef.BrowserWindow.GetWindowInfo(browser.Identifier())
				var target = &cef.GoEmitTarget{
					BrowseId: browser.Identifier(),
					FrameId:  browser.MainFrame().Id,
				}
				fmt.Println("browseEmitJsOnEvent 1 browseId:", browser.Identifier(), "info-browserId:", info.Chromium().BrowserId(), "GetFrameById:", browser.GetFrameById(target.FrameId))
				var argumentList = cef.NewArgumentList()
				argumentList.SetBool(0, isLoading)
				argumentList.SetBool(1, canGoBack)
				argumentList.SetBool(2, canGoForward)
				//加載时的状态，对于非html 例如 image pdf 在线预览 需要自行处理
				fmt.Println("OnLoadingStateChange-Emit:", browserWindow.Chromium().Emit("OnLoadingStateChange", argumentList, target))
			}
		})
		event.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
			//当刷新的是一个完整的浏览器时，如果打开的新页面不是html dom，这里的 emit 消息
			fmt.Println("OnLoadingProgressChange-ProcessType:", cef.Args.ProcessType(), "browserId:", browser.Identifier(), "progress:", progress)
			if isSendEmit {
				var target = &cef.GoEmitTarget{
					BrowseId: browser.Identifier(),
					FrameId:  browser.MainFrame().Id,
				}
				var argumentList = cef.NewArgumentList()
				argumentList.SetFloat64(0, progress)
				fmt.Println("OnLoadingProgressChange-Emit:", browserWindow.Chromium().Emit("OnLoadingProgressChange", argumentList, target), " frame:", cef.BrowserWindow.GetFrames(browser.Identifier()))
			}
		})
	})
	//添加子窗口初始化
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		//在这里创建 一些子窗口 子组件 等
		//托盘
		if cef.IsWindows() {
			cefTray(browserWindow)
		} else {
			tray(browserWindow)
		}
	})
}

// 托盘 只适用 windows 的系统托盘, 基于html 和 ipc 实现功能
func cefTray(browserWindow *cef.TCefWindowInfo) {
	var url = "E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\demo-golang-dll-01-chromium\\demos\\min-browser\\resources\\min-browser-tray.html"
	if !tools.IsExist(url) {
		url = cef.ExePath + cef.Separator + "min-browser-tray.html"
	}
	tray := browserWindow.NewCefTray(250, 300, url)
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIcon("resources/icon.ico")
	tray.SetOnClick(func(sender lcl.IObject) {
		browserWindow.Window.SetVisible(!browserWindow.Window.Visible())
	})
	tray.SetBalloon("气泡标题", "气泡内容", 2000)
	cef.IPC.Browser().On("tray-show-balloon", func(context cef.IIPCContext) {
		fmt.Println("tray-show-balloon")
		tray.ShowBalloon()
		tray.Hide()
	})
	cef.IPC.Browser().On("tray-show-main-window", func(context cef.IIPCContext) {
		vb := !browserWindow.Window.Visible()
		browserWindow.Window.SetVisible(vb)
		if vb {
			if browserWindow.Window.WindowState() == types.WsMinimized {
				browserWindow.Window.SetWindowState(types.WsNormal)
			}
			browserWindow.Window.Focused()
		}
		tray.Hide()
	})
	cef.IPC.Browser().On("tray-close-main-window", func(context cef.IIPCContext) {
		browserWindow.Window.Close()
	})
	//托盘 end
}

// 托盘 系统原生 windows linux macos
func tray(browserWindow *cef.TCefWindowInfo) {
	//托盘 windows linux macos 系统托盘
	newTray := browserWindow.NewTray()
	tray := newTray.Tray()
	tray.SetIcon("resources/icon.ico")
	menu1 := tray.AddMenuItem("父菜单", nil)
	menu1.Add(tray.NewMenuItem("子菜单", func(object lcl.IObject) {
		lcl.ShowMessage("子菜单点击 提示消息")
	}))
	tray.AddMenuItem("显示气泡", func(object lcl.IObject) {
		newTray.ShowBalloon()
	})
	tray.AddMenuItem("显示/隐藏", func(object lcl.IObject) {
		vis := browserWindow.Window.Visible()
		cef.BrowserWindow.GetWindowInfo(1)
		browserWindow.Window.SetVisible(!vis)
	})
	tray.AddMenuItem("退出", func(object lcl.IObject) {
		browserWindow.Close()
	})
	//linux下有些问题
	tray.SetBalloon("气泡标题", "气泡内容", 2000)
	//托盘 end
}

// 自定义组件
func addressBar(window *cef.BaseWindow) {
	window.WindowParent().SetAlign(types.AlNone)
	window.WindowParent().SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	window.WindowParent().SetTop(0)
	window.WindowParent().SetWidth(window.Width())
	window.WindowParent().SetHeight(window.Height())
	//addrBar := lcl.NewPanel(window)
	//addrBar.SetParent(window)
	//addrBar.SetTop(0)
	//addrBar.SetHeight(35)
	//addrBar.SetWidth(window.Width())
	//addrBar.SetBorderStyle(types.BsNone)
	//addrBar.SetParentBackground(true)
	//addrBar.SetColor(colors.ClHotpink)
	//addrBar.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight))
	//var (
	//	isDown       bool
	//	downX, downY int32 //鼠标 点击
	//	winX, winY   int32 //窗体 坐标
	//)
	//addrBar.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	//	fmt.Println("鼠标 down", x, y)
	//	isDown = true
	//	downX, downY = x, y
	//	winX, winY = window.Left(), window.Top()
	//})
	//addrBar.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
	//	if isDown {
	//		fmt.Println("鼠标 down", x, y)
	//		mex := x - (downX - winX)
	//		mey := y - (downY - winY)
	//		winX = mex
	//		winY = mey
	//		window.SetLeft(winX)
	//		window.SetTop(winY)
	//	}
	//})
	//addrBar.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	//	isDown = false
	//	fmt.Println("鼠标 up", x, y)
	//})
	//monitor := window.Monitor()
	//
	//fmt.Println("monitor WorkareaRect:", monitor.WorkareaRect())
	//fmt.Println("monitor Height:", monitor.Height(), "Width:", monitor.Width())
	//fmt.Println("m.Window.WindowState():", window.WindowState())
}
