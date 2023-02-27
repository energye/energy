//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package src

import (
	"encoding/json"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/example/dev-test/traydemo"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"strings"
	"time"
)

func AppBrowserInit() {
	/*
		通用类型变量和对象类型变量创建的回调函数
		该回调函数-在主进程和渲染进程创建时调用
	*/
	//ar integer cef.JSValue
	cef.VariableBind.VariableCreateCallback(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, bind cef.IProvisionalBindStorage) {
		ObjDemoVar = &ObjDemo{SubObjDemoField: &SubObjDemo{}}
		fmt.Println("GO变量和函数绑定回调", common.Args.ProcessType())
		bind.NewString("stringv", "这是一个字符串变量")
		bind.NewInteger("integerv", 1211111)
		bind.NewDouble("doublev", 11.0505)
		bind.NewBoolean("booleanv", true)
		bind.NewNull("nullv")
		bind.NewUndefined("undefinedv")
		bind.NewObjects(ObjDemoVar)
	})

	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)
	config.SetEnableViewSource(true)
	config.SetEnableDevTools(true)
	config.SetEnableWindowPopup(true)
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//默认加载的URL 这个示例启动了一个内置http服务
	cef.BrowserWindow.Config.Url = "http://localhost:22022/demo-misc.html"
	cef.BrowserWindow.Config.EnableWebkitAppRegion = true
	//cef.BrowserWindow.Config.EnableResize = false
	//cef.BrowserWindow.Config.EnableCenterWindow = false
	//主进程 IPC事件
	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		fmt.Println("主进程IPC事件注册")
		event.On("ZoomPct", func(context ipc.IIPCContext) {
			fmt.Println("ZoomPct")
			bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			bw.Chromium().SetZoomPct(150.2)
			fmt.Println(bw.Chromium().ZoomPct())
			bw.Chromium().SetDefaultEncoding("UTF-8")
			fmt.Println(bw.Chromium().DefaultEncoding())
		})
		event.On("ZoomLevel", func(context ipc.IIPCContext) {
			fmt.Println("ZoomLevel")
			bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			bw.Chromium().SetZoomLevel(150.2)
			fmt.Println(bw.Chromium().ZoomLevel())
		})
		event.On("SetZoomStep", func(context ipc.IIPCContext) {
			fmt.Println("SetZoomStep")
			bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			bw.Chromium().SetZoomStep(1)
			fmt.Println(bw.Chromium().ZoomStep())
		})
		//这个事件监听演示了几个示例
		//1.
		event.On("subWindowIPCOn", func(context ipc.IIPCContext) {
			//函数内的 context 返回给调用方
			//取入参
			fmt.Println("browser renderOnEventSubWindowIPCOn:", context.Arguments().GetString(0), "channelId:", context.ChannelId())
			fmt.Println("\t|--", len(cef.BrowserWindow.GetWindowInfos()))
			//调用指定渲染进程监听
			ipc.IPC.Browser().EmitChannelId("renderOnEventSubWindowIPCOn", context.ChannelId(), nil)
			fmt.Println("\t|-- 111")

			//调用指定渲染进程监听 回调的方式
			ipc.IPC.Browser().EmitChannelIdAndCallback("renderOnEventSubWindowIPCOn", context.ChannelId(), nil, func(context ipc.IIPCContext) {
				fmt.Println("browser renderOnEventSubWindowIPCOn 回调的方式:", string(context.Message().Data()))
				//在这里 不能使用 context 返回给页面
				//函数内的 context 返回给调用方
			})
			fmt.Println("\t|-- 222")
			//调用指定渲染进程监听 同步的方式
			ctx := ipc.IPC.Browser().EmitChannelIdAndReturn("renderOnEventSubWindowIPCOn", context.ChannelId(), nil)
			context.Result().SetString("成功返回, " + string(ctx.Message().Data()))
			fmt.Println("\t|-- 333")
			//返回给页面数据
			//context.Result().SetString("成功返回 ")
		})
		event.On("close", func(context ipc.IIPCContext) {
			fmt.Println("close browserId:", context.BrowserId())
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.CloseBrowserWindow()
		})
		event.On("minWindow", func(context ipc.IIPCContext) {
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Minimize()
		})
		event.On("maxWindow", func(context ipc.IIPCContext) {
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Maximize()
		})
		event.On("closeWindow", func(context ipc.IIPCContext) {
			fmt.Println("closeWindow", context.BrowserId())
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.CloseBrowserWindow()
		})
		event.On("window-list", func(context ipc.IIPCContext) {
			var ids []int32
			for id, _ := range cef.BrowserWindow.GetWindowInfos() {
				ids = append(ids, id)
			}
			idsStr, err := json.Marshal(ids)
			fmt.Println("获得window-id-list", ids, idsStr, err)
			context.Result().SetString(string(idsStr))
		})
		event.On("find", func(context ipc.IIPCContext) {
			args := context.Arguments().GetString(0)
			fmt.Println("find-args:", args)
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Browser().Find(args, false, false, true)
		})
		event.On("find-stop", func(context ipc.IIPCContext) {
			winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			winInfo.Browser().StopFinding(true)
		})
		var newForm *cef.LCLBrowserWindow
		event.On("js-new-window", func(context ipc.IIPCContext) {
			fmt.Println("创建新窗口 ProcessType:", common.Args.ProcessType())
			if newForm == nil {
				newForm = cef.NewLCLWindow(cef.NewWindowProperty())
				newForm.SetTitle("新窗口标题")
				newForm.HideTitle()
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
		var browserWindow *cef.LCLBrowserWindow
		event.On("js-new-browser-window", func(context ipc.IIPCContext) {
			fmt.Println("通过 js ipc emit 事件创建新Browser窗口 ProcessType:", common.Args.ProcessType())

			if browserWindow == nil || browserWindow.IsClosing() {
				wp := cef.NewWindowProperty()
				wp.Url = "https://www.baidu.com"
				wp.Title = "Browser新窗口标题"
				browserWindow = cef.NewLCLBrowserWindow(nil, wp)
				browserWindow.SetWidth(800)
				browserWindow.SetHeight(600)
				browserWindow.HideTitle()
				browserWindow.SetShowInTaskBar()
				browserWindow.EnableDefaultCloseEvent()
				browserWindow.Chromium().SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
					fmt.Println("SetOnTitleChange", wp.Title, title)
				})
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
		event.On("go-call-js-code", func(context ipc.IIPCContext) {
			fmt.Println("通过 js ipc emit go-call-js-code ProcessType:", common.Args.ProcessType())
			info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			info.Chromium().ExecuteJavaScript("jsCode('值值值')", "", 0)
			list := ipc.NewArgumentList()
			list.SetString(0, "js-ipc-on参数", true)
			list.SetFloat64(1, 99999111.0121212)
			//只触发js ipc.on 的函数，忽略返回值
			info.Chromium().Emit("js-ipc-on", list, nil)
			list.SetFloat64(1, 8888888111.0121212)
			//触发js ipc.on 的函数，通过回调函数接收返回值
			info.Chromium().EmitAndCallback("js-ipc-on", list, nil, func(context ipc.IIPCContext) {
				fmt.Println("回调函数方式返回值:", context.Arguments().GetString(0))
			})
			//这个方式不适合于ui线程
			//ctx, err := info.Chromium().EmitAndReturn("js-ipc-on", list, nil)
			//if err != consts.PME_OK {
			//
			//}
			//fmt.Println("等待阻塞的方式返回值:", ctx.Arguments().GetString(0))
		})
	})

	//主窗口初始化回调函数
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow cef.IBrowserWindow) {
		lcl.Application.SetOnMinimize(func(sender lcl.IObject) {
			fmt.Println("minimize")
		})
		fmt.Println("主窗口初始化回调函数")
		fmt.Println("LCL", browserWindow.AsLCLBrowserWindow(), "VF", browserWindow.AsViewsFrameworkBrowserWindow())
		lcl.Application.Icon().LoadFromFSFile("resources/icon.ico") //设置应用图标
		browserWindow.SetTitle("这里设置应用标题")
		browserWindow.HideTitle()
		//browserWindow.EnableTransparent(100) //窗口透明
		//设置窗口样式，无标题 ，最大化按钮等
		window := browserWindow.AsLCLBrowserWindow()
		//window.BrowserWindow().SetBorderStyle(types.BsSizeable)
		//browserWindow.DisableResize()
		//browserWindow.HideTitle()
		//window.BrowserWindow().SetBorderStyle(types.BsNone)
		//window.BrowserWindow().SetFormStyle(types.FsNormal)
		//window.BrowserWindow().SetFormStyle(types.FsSystemStayOnTop)
		//设置窗口大小
		//browserWindow.SetWidth(1600)
		//browserWindow.SetHeight(900)
		//限制窗口大小
		//browserWindow.Window.Constraints().SetMinHeight(300)
		//browserWindow.Window.Constraints().SetMinWidth(300)
		//browserWindow.Window.Constraints().SetMaxWidth(1600)
		//browserWindow.Window.Constraints().SetMaxHeight(900)
		window.BrowserWindow().Constraints().SetOnChange(func(sender lcl.IObject) {
			fmt.Println("browserWindow SetOnChange")
		})
		//添加事件，add不会覆盖默认的事件 set会覆盖默认的事件
		window.BrowserWindow().SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
			fmt.Println("添加onclose事件")
			return false
		})
		event.SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
			fmt.Println("SetOnTitleChange", title)
		})

		//窗口大小改变后触发
		window.BrowserWindow().SetOnResize(func(sender lcl.IObject) bool {
			//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
			if browserWindow.Browser() != nil {
				for {
					if window.Chromium().Initialized() {
						fmt.Println("Initialized", window.Chromium().Initialized())
						break
					}
					fmt.Println("Initialized", window.Chromium().Initialized())
					time.Sleep(time.Second)
				}
				fmt.Println("Identifier", browserWindow.Chromium().Browser().Identifier())
				fmt.Println("SetOnConstrainedResize Identifier", browserWindow.Browser().Identifier())
				fmt.Println("SetOnConstrainedResize MainFrame", browserWindow.Browser().MainFrame())
				fmt.Println("SetOnConstrainedResize MainFrame", browserWindow.Browser().MainFrame())
				//var argumentList = ipc.NewArgumentList()
				//argumentList.SetInt32(0, window.BrowserWindow().Left())
				//argumentList.SetInt32(1, window.BrowserWindow().Top())
				//argumentList.SetInt32(2, window.BrowserWindow().Width())
				//argumentList.SetInt32(3, window.BrowserWindow().Height())
				//browserWindow.Chromium().Emit("window-resize", argumentList, target)
				//browserWindow.Chromium().EmitAndCallback("window-resize", argumentList, target, func(context ipc.IIPCContext) {
				//	fmt.Println("EmitAndCallback AddOnResize")
				//})
				//使用EmitAndReturn函数会锁死
				//ctx, _ := browserWindow.Chromium.EmitAndReturn("window-resize", argumentList, target)
				//fmt.Println("EmitAndReturn AddOnResize", ctx)
			}
			return false
		})
		//windows下可以使用这个函数，实时触发
		window.BrowserWindow().SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
			//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
			if browserWindow.Browser() != nil {
				//var target = &cef.EmitTarget{
				//	BrowseId: browserWindow.Browser().Identifier(),
				//	FrameId:  browserWindow.Browser().MainFrame().Id,
				//}
				fmt.Println("SetOnConstrainedResize Identifier", browserWindow.Browser().Identifier())
				fmt.Println("SetOnConstrainedResize MainFrame", browserWindow.Browser().MainFrame())
				//var argumentList = ipc.NewArgumentList()
				//argumentList.SetInt32(0, window.BrowserWindow().Left())
				//argumentList.SetInt32(1, window.BrowserWindow().Top())
				//argumentList.SetInt32(2, window.BrowserWindow().Width())
				//argumentList.SetInt32(3, window.BrowserWindow().Height())
				//browserWindow.Chromium().Emit("window-resize", argumentList, target)
				//使用EmitAndReturn函数会锁死
				//	browserWindow.Chromium.EmitAndCallback("window-resize", argumentList, target, func(context cef.IIPCContext) {
				//		fmt.Println("EmitAndCallback OnConstrainedResize")
				//	})
			}
		})
		//自定义browser窗体
		addressBar(window.BrowserWindow())
		//下载文件
		//linux系统弹出保存对话框不启作用
		//自己调用系统的保存对话框获得保存路径
		dlSave := lcl.NewSaveDialog(window.BrowserWindow())
		dlSave.SetTitle("保存对话框标题")
		event.SetOnBeforeDownload(func(sender lcl.IObject, browser *cef.ICefBrowser, beforeDownloadItem *cef.DownloadItem, suggestedName string, callback *cef.ICefBeforeDownloadCallback) {
			fmt.Println("OnBeforeDownload:", beforeDownloadItem, suggestedName)
			//linux下 需要这样使用 Sync
			if common.IsLinux() {
				cef.QueueSyncCall(func(id int) {
					dlSave.SetFileName(suggestedName)
					if dlSave.Execute() {
						fmt.Println("OnBeforeDownload filename: ", dlSave.FileName())
						callback.Cont(dlSave.FileName(), false)
					}
				})
			} else {
				callback.Cont(consts.ExePath+consts.Separator+suggestedName, true)
			}
		})
		event.SetOnDownloadUpdated(func(sender lcl.IObject, browser *cef.ICefBrowser, downloadItem *cef.DownloadItem, callback *cef.ICefDownloadItemCallback) {
			fmt.Println("OnDownloadUpdated:", downloadItem)
		})

		event.SetOnFindResult(func(sender lcl.IObject, browser *cef.ICefBrowser, identifier, count int32, selectionRect *cef.TCefRect, activeMatchOrdinal int32, finalUpdate bool) {
			fmt.Println("OnFindResult:", browser.Identifier(), identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
		})
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup: "+beforePopupInfo.TargetUrl, "isLCL:", popupWindow.IsLCL())

			popupWindow.SetShowInTaskBar()
			popupWindow.HideTitle()
			popupWindow.SetTitle("改变了标题 - " + beforePopupInfo.TargetUrl)
			popupWindow.SetSize(800, 600)
			window := popupWindow.AsLCLBrowserWindow().BrowserWindow()
			//窗口弹出之前可自定义系统组件
			//窗口大小改变后触发
			//windows下窗口调整后触发一次
			window.SetOnResize(func(sender lcl.IObject) bool {
				//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
				if popupWindow.Browser() != nil {
					var target = &cef.EmitTarget{
						BrowseId: popupWindow.Browser().Identifier(),
						FrameId:  popupWindow.Browser().MainFrame().Identifier(),
					}
					var argumentList = ipc.NewArgumentList()
					argumentList.SetInt32(0, window.Left())
					argumentList.SetInt32(1, window.Top())
					argumentList.SetInt32(2, window.Width())
					argumentList.SetInt32(3, window.Height())
					popupWindow.Chromium().Emit("window-resize", argumentList, target)
				}
				return false
			})
			//windows下可以使用这个函数，实时触发
			window.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
				//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
				if popupWindow.Browser() != nil {
					var target = &cef.EmitTarget{
						BrowseId: popupWindow.Browser().Identifier(),
						FrameId:  popupWindow.Browser().MainFrame().Identifier(),
					}
					var argumentList = ipc.NewArgumentList()
					argumentList.SetInt32(0, window.Left())
					argumentList.SetInt32(1, window.Top())
					argumentList.SetInt32(2, window.Width())
					argumentList.SetInt32(3, window.Height())
					popupWindow.Chromium().Emit("window-resize", argumentList, target)
				}
			})
			return false
		})
		//加载页面之前
		var isSendEmit bool
		event.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
			fmt.Println("OnLoadStart:", frame.Url())
			//判断一下哪些页面不发送emit消息
			if strings.LastIndex(strings.ToLower(frame.Url()), ".pdf") > 0 || strings.Index(frame.Url(), "about:blank") != -1 {
				isSendEmit = false
			} else {
				isSendEmit = true
			}
		})
		event.SetOnFrameCreated(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame) {
			fmt.Println("OnFrameCreated:", frame.Url())
		})
		event.SetOnLoadingStateChange(func(sender lcl.IObject, browser *cef.ICefBrowser, isLoading, canGoBack, canGoForward bool) {
			//当刷新的是一个完整的浏览器时，如果打开的新页面不是html dom，这里的 emit 消息 将会失败
			fmt.Println("OnLoadingStateChange-ProcessType:", common.Args.ProcessType(), "sender.Instance:", sender.Instance(), "browserId:", browser.Identifier(), "isLoading:", isLoading, "canGoBack:", canGoBack, "canGoForward:", canGoForward)
			if isSendEmit {
				info := cef.BrowserWindow.GetWindowInfo(browser.Identifier())
				var target = &cef.EmitTarget{
					BrowseId: browser.Identifier(),
					FrameId:  browser.MainFrame().Identifier(),
				}
				fmt.Println("browseEmitJsOnEvent 1 browseId:", browser.Identifier(), "info-browserId:", info.Chromium().BrowserId(), "GetFrameById:", browser.GetFrameById(target.FrameId))
				var argumentList = ipc.NewArgumentList()
				argumentList.SetBool(0, isLoading)
				argumentList.SetBool(1, canGoBack)
				argumentList.SetBool(2, canGoForward)
				//加載时的状态，对于非html 例如 image pdf 在线预览 需要自行处理
				fmt.Println("OnLoadingStateChange-Emit:", browserWindow.Chromium().Emit("OnLoadingStateChange", argumentList, target))
			}
		})
		event.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
			//当刷新的是一个完整的浏览器时，如果打开的新页面不是html dom，这里的 emit 消息
			fmt.Println("OnLoadingProgressChange-ProcessType:", common.Args.ProcessType(), "browserId:", browser.Identifier(), "progress:", progress)
			if isSendEmit {
				var target = &cef.EmitTarget{
					BrowseId: browser.Identifier(),
					FrameId:  browser.MainFrame().Identifier(),
				}
				var argumentList = ipc.NewArgumentList()
				argumentList.SetFloat64(0, progress)
				fmt.Println("OnLoadingProgressChange-Emit:", browserWindow.Chromium().Emit("OnLoadingProgressChange", argumentList, target))
			}
		})
		event.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, callback *cef.ICefCallback, result *consts.TCefReturnValue) {
			fmt.Println("SetOnBeforeResourceLoad:", request.Url, request.Method, "headerMap:", request.GetHeaderMap().GetSize())
			headerMap := request.GetHeaderMap()
			fmt.Println("\t", request.GetHeaderByName("energy"), headerMap.GetEnumerate("energy", 1), "size:", headerMap.GetSize())
			for i := 0; i < int(headerMap.GetSize()); i++ {
				fmt.Println("\tkey:", headerMap.GetKey(int32(i)), "value:", headerMap.GetValue(int32(i)))
			}
			multiMap := cef.StringMultiMapRef.New()
			fmt.Println("multiMap.GetSize()", multiMap.GetSize())
			multiMap.Append("key1", "value1")
			fmt.Println("multiMap.GetSize()", multiMap.GetSize())
		})
	})
	//添加子窗口初始化
	cef.BrowserWindow.SetBrowserInitAfter(func(window cef.IBrowserWindow) {
		//在这里创建 一些子窗口 子组件 等
		if window.IsLCL() {
			if common.IsLinux() {
				traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用, 推荐linux中使用
			} else {
				//支持 windows
				traydemo.LCLCefTrayDemo(window) //对于LCL+CEF web端技术托盘实现无法在VF中使用
				//支持 windows or macosx
				//traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
			}
		} else if window.IsViewsFramework() {
			if common.IsLinux() || common.IsDarwin() {
				//在VF窗口组件中, 推荐linux和macosx中使用
				traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用,
			} else {
				//不支持windows VF窗口组件中无法创建或使用LCL组件
				//traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
				//traydemo.LCLCefTrayDemo(window) //对于LCL+CEF web端技术托盘实现无法在VF中使用
				//支持windows
				traydemo.LCLVFTrayDemo(window) //对于LCL+VF web端技术托盘实现
			}
		}
		println("browser init after end")
	})
}

// 自定义组件
func addressBar(window *cef.LCLBrowserWindow) {
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
	//fmt.Println("m.WindowState():", window.WindowState())
}
