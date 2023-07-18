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
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/cef/ipc/target"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/example/dev-test/traydemo"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"strings"
	"time"
)

func AppBrowserInit() {

	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)
	config.SetEnableViewSource(true)
	config.SetEnableDevTools(true)
	config.SetEnableWindowPopup(true)
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//默认加载的URL 这个示例启动了一个内置http服务
	cef.BrowserWindow.Config.Url = "http://localhost:22022/demo-misc.html"
	//主进程 IPC事件
	fmt.Println("主进程IPC事件注册")
	ipc.On("ZoomPct", func(context context.IContext) {
		fmt.Println("ZoomPct")
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().SetZoomPct(150.2)
		fmt.Println(bw.Chromium().ZoomPct())
		bw.Chromium().SetDefaultEncoding("UTF-8")
		fmt.Println(bw.Chromium().DefaultEncoding())
	})
	ipc.On("ZoomLevel", func(context context.IContext) {
		fmt.Println("ZoomLevel")
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().SetZoomLevel(150.2)
		fmt.Println(bw.Chromium().ZoomLevel())
	})
	ipc.On("SetZoomStep", func(context context.IContext) {
		fmt.Println("SetZoomStep")
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().SetZoomStep(1)
		fmt.Println(bw.Chromium().ZoomStep())
	})
	//这个事件监听演示了几个示例
	//1.
	//ipc.On("subWindowIPCOn", func(context context.IContext) {
	//	//函数内的 context 返回给调用方
	//	//取入参
	//	fmt.Println("browser renderOnEventSubWindowIPCOn:", context.ArgumentList().GetStringByIndex(0), "channelId:", context.FrameId())
	//	fmt.Println("\t|--", len(cef.BrowserWindow.GetWindowInfos()))
	//	//调用指定渲染进程监听
	//	ipc.EmitTarget("renderOnEventSubWindowIPCOn", context.FrameId(), nil)
	//	fmt.Println("\t|-- 111")
	//
	//	//调用指定渲染进程监听 回调的方式
	//	ipc.EmitTargetAndCallback("renderOnEventSubWindowIPCOn", context.FrameId(), nil, func(context context.IContext) {
	//		fmt.Println("browser renderOnEventSubWindowIPCOn 回调的方式:", string(context.Message().Data()))
	//		//在这里 不能使用 context 返回给页面
	//		//函数内的 context 返回给调用方
	//	})
	//	fmt.Println("\t|-- 222")
	//})
	ipc.On("close", func(context context.IContext) {
		fmt.Println("close browserId:", context.BrowserId())
		winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		winInfo.CloseBrowserWindow()
	})
	ipc.On("minWindow", func(context context.IContext) {
		winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		winInfo.Minimize()
	})
	ipc.On("maxWindow", func(context context.IContext) {
		winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		winInfo.Maximize()
	})
	ipc.On("closeWindow", func(context context.IContext) {
		fmt.Println("closeWindow", context.BrowserId())
		winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		winInfo.CloseBrowserWindow()
	})
	ipc.On("window-list", func(context context.IContext) {
		var ids []int32
		for id, _ := range cef.BrowserWindow.GetWindowInfos() {
			ids = append(ids, id)
		}
		idsStr, err := json.Marshal(ids)
		fmt.Println("获得window-id-list", ids, idsStr, err)
		context.Result(string(idsStr))
	})
	ipc.On("find", func(context context.IContext) {
		args := context.ArgumentList().GetStringByIndex(0)
		fmt.Println("find-args:", args)
		winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		winInfo.Browser().Find(args, false, false, true)
	})
	ipc.On("find-stop", func(context context.IContext) {
		winInfo := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		winInfo.Browser().StopFinding(true)
	})
	var newForm *cef.LCLBrowserWindow
	ipc.On("js-new-window", func(context context.IContext) {
		fmt.Println("创建新窗口 ProcessType:", process.Args.ProcessType())
		if newForm == nil {
			newForm = cef.NewLCLWindow(cef.NewWindowProperty())
			newForm.SetTitle("新窗口标题")
			newForm.SetWidth(600)
			newForm.SetHeight(400)
			//newForm.HideTitle()
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
				newForm.ShowModal()
			}
		})
	})
	var browserWindow *cef.LCLBrowserWindow
	ipc.On("js-new-browser-window", func(context context.IContext) {
		fmt.Println("通过 js ipc emit 事件创建新Browser窗口 ProcessType:", process.Args.ProcessType())

		if browserWindow == nil || browserWindow.IsClosing() {
			wp := cef.NewWindowProperty()
			wp.Url = "https://www.baidu.com"
			wp.Title = "Browser新窗口标题"
			browserWindow = cef.NewLCLBrowserWindow(nil, wp)
			browserWindow.SetWidth(800)
			browserWindow.SetHeight(600)
			//browserWindow.HideTitle()
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
				browserWindow.ShowModal()
			}
		})
	})
	ipc.On("go-call-js-code", func(context context.IContext) {
		fmt.Println("通过 js ipc emit go-call-js-code ProcessType:", process.Args.ProcessType())
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		info.Chromium().ExecuteJavaScript("jsCode('值值值')", "", 0)
		//触发js ipc.on 的函数，通过回调函数接收返回值
		//这个方式不适合于ui线程
		//ctx, err := info.Chromium().EmitAndReturn("js-ipc-on", list, nil)
		//if err != consts.PME_OK {
		//
		//}
		//fmt.Println("等待阻塞的方式返回值:", ctx.Arguments().GetString(0))
	})

	//主窗口初始化回调函数
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow cef.IBrowserWindow) {
		//return
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
		event.SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string, window cef.IBrowserWindow) {
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
		event.SetOnBeforeDownload(func(sender lcl.IObject, browser *cef.ICefBrowser, beforeDownloadItem *cef.ICefDownloadItem, suggestedName string, callback *cef.ICefBeforeDownloadCallback) {
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
		event.SetOnDownloadUpdated(func(sender lcl.IObject, browser *cef.ICefBrowser, downloadItem *cef.ICefDownloadItem, callback *cef.ICefDownloadItemCallback) {
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
					//var tag = target.NewTarget(browser.Identifier(), browser.MainFrame().Identifier())
					//ipc.EmitTarget("window-resize", tag, window.Left(), window.Top(), window.Width(), window.Height())
				}
				return false
			})
			//windows下可以使用这个函数，实时触发
			window.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
				//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
				if popupWindow.Browser() != nil {
					//var tag = target.NewTarget(browser.Identifier(), browser.MainFrame().Identifier())
					//ipc.EmitTarget("window-resize", tag, window.Left(), window.Top(), window.Width(), window.Height())
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
			fmt.Println("OnLoadingStateChange-ProcessType:", process.Args.ProcessType(), "sender.Instance:", sender.Instance(), "browserId:", browser.Identifier(), "isLoading:", isLoading, "canGoBack:", canGoBack, "canGoForward:", canGoForward)
			if isSendEmit {
				var tag = target.NewTarget(browser.Identifier(), browser.MainFrame().Identifier())
				//加載时的状态，对于非html 例如 image pdf 在线预览 需要自行处理
				ipc.EmitTarget("OnLoadingStateChange", tag, isLoading, canGoBack, canGoForward)
			}
		})
		event.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
			//当刷新的是一个完整的浏览器时，如果打开的新页面不是html dom，这里的 emit 消息
			fmt.Println("OnLoadingProgressChange-ProcessType:", process.Args.ProcessType(), "browserId:", browser.Identifier(), "progress:", progress)
			if isSendEmit {
				var tag = target.NewTarget(browser.Identifier(), browser.MainFrame().Identifier())
				ipc.EmitTarget("OnLoadingProgressChange", tag, progress)
			}
		})
		event.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, callback *cef.ICefCallback, result *consts.TCefReturnValue) {
			fmt.Println("SetOnBeforeResourceLoad")
			fmt.Println("SetOnBeforeResourceLoad:", request.Method())
			fmt.Println("SetOnBeforeResourceLoad:", request.URL())
			fmt.Println("headerMap:", request.GetHeaderMap().GetSize())
			headerMap := request.GetHeaderMap()
			fmt.Println("\t", request.GetHeaderByName("energy"), "size:", headerMap.GetSize())
			for i := 0; i < int(headerMap.GetSize()); i++ {
				fmt.Println("\ti", i)
				fmt.Println("\tkey:", headerMap.GetKey(uint32(i)))
				fmt.Println("\tvalue:", headerMap.GetValue(uint32(i)))
			}
			return
			multiMap := cef.StringMultiMapRef.New()
			fmt.Println("multiMap.GetSize()", multiMap.GetSize())
			multiMap.Append("key1", "value1")
			fmt.Println("multiMap.GetSize()", multiMap.GetSize())
			//postData := cef.PostDataRef.New()
			//postData.
			fmt.Println("GetPostData().GetElementCount", request.GetPostData().IsValid())
			if !request.GetPostData().IsValid() {
				data := cef.PostDataRef.New()
				postDataElement := cef.PostDataElementRef.New()
				postDataElement.SetToFile("/Users/zhangli/energy.log")
				data.AddElement(postDataElement)
				postDataElement = cef.PostDataElementRef.New()
				bytes := make([]byte, 256, 256)
				for i := 0; i < len(bytes); i++ {
					bytes[i] = byte(i)
				}
				fmt.Println("postDataElement.SetToBytes", bytes)
				postDataElement.SetToBytes(bytes)
				data.AddElement(postDataElement)
				request.SetPostData(data)
				fmt.Println("\tGetPostData GetElementCount:", request.GetPostData().IsValid(), request.GetPostData().GetElementCount(), data.IsReadOnly())
				fmt.Println("\tGetElements Size:", request.GetPostData().GetElements().Size())
				fmt.Println("\tGetElements GetFile:", request.GetPostData().GetElements().Get(0).GetFile())
				fmt.Println("\tGetElements GetBytesCount:", request.GetPostData().GetElements().Get(0).GetBytesCount())
				postDataElement = request.GetPostData().GetElements().Get(1)
				fmt.Println("\tGetElements GetBytesCount:", postDataElement.GetBytesCount())
				bytes, count := postDataElement.GetBytes()
				fmt.Println("\tGetElements bytes,count:", bytes, count)
			}
		})
		//在这里创建 一些子窗口 子组件 等
		if window.IsLCL() {
			if common.IsLinux() {
				traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用, 推荐linux中使用
			} else {
				//支持 windows
				//traydemo.LCLCefTrayDemo(window) //对于LCL+CEF web端技术托盘实现无法在VF中使用
				//支持 windows or macosx
				traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
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
