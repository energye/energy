package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/consts/messages"
	exampleCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	et "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"path/filepath"
	"time"
)

type BrowserWindow struct {
	cef.LCLBrowserWindow
	timer        *lcl.TTimer
	windowParent cef.ICEFWindowParent
	chromium     cef.IChromium
	canClose     bool
	ChildForm    *cef.LCLBrowserWindow
}

var BW *BrowserWindow

func main() {
	fmt.Println(time.Now().UnixNano()/1e6, time.Now().UnixNano())
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache", "debug_most")
	app := cef.CreateApplication()
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetLogSeverity(consts.LOGSEVERITY_DEBUG)
	app.SetShowMessageDlg(true)
	cef.SetApplication(app)
	// setting
	if common.IsDarwin() {
		if process.Args.IsMain() {
			app.SetUseMockKeyChain(true)
			cef.GlobalWorkSchedulerCreate(nil)
			app.SetOnScheduleMessagePumpWork(nil)
			app.SetExternalMessagePump(true)
			app.SetMultiThreadedMessageLoop(false)
		} else {
			app.InitLibLocationFromArgs()
		}
	} else {
		// 指定 CEF Framework
		app.SetFrameworkDirPath(config.Get().FrameworkPath())
		if common.IsLinux() {
			app.SetDisableZygote(true)
		}
		app.SetExternalMessagePump(false)
		app.SetMultiThreadedMessageLoop(true)
	}
	// run
	if common.IsDarwin() && !process.Args.IsMain() {
		startSub := app.StartSubProcess()
		fmt.Println("start sub:", startSub)
		app.Free()
	} else {
		startMain := app.StartMainProcess()
		fmt.Println("start main:", startMain)
		if startMain {
			fmt.Println("WidgetUI:", api.WidgetUI(), "ChromeVersion:", app.ChromeVersion(), "LibCefVersion:", app.LibCefVersion())
			fmt.Println("LibAbout:", api.DLibAbout())
			startServer()
			// 结束应用后释放资源
			api.SetReleaseCallback(func() {
				fmt.Println("Release")
				app.Destroy()
				app.Free()
			})
			// LCL窗口
			lcl.Application.Initialize()
			lcl.Application.SetMainFormOnTaskBar(true)
			lcl.Application.CreateForm(&BW, true)
			lcl.Application.Run()
		}
	}
}

func (m *BrowserWindow) OnFormCreate(sender lcl.IObject) {
	m.ScreenCenter()
	//m.SetOnWndProc(func(msg *types.TMessage) {
	//	m.InheritedWndProc(msg)
	//	//fmt.Println("SetOnWndProc")
	//})

	m.createMenu()

	m.chromium = cef.NewChromium(m, nil)
	//m.chromium.SetDefaultURL("https://www.baidu.com")
	//m.chromium.SetDefaultURL("https://energye.github.io")
	m.chromium.SetDefaultURL("http://localhost:22022")
	m.windowParent = cef.NewCEFWindowParent(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)
	m.windowParent.SetChromium(m.chromium, 0)
	// 创建一个定时器, 用来createBrowser
	m.timer = lcl.NewTimer(m)
	m.timer.SetOnTimer(m.createBrowser)
	m.timer.SetEnabled(false)
	// 在show时创建chromium browser
	if !common.IsLinux() {
		m.TForm.SetOnShow(m.createBrowser)
	}
	m.TForm.SetOnActivate(m.createBrowser)
	m.TForm.SetOnResize(m.resize)
	// 1. 关闭之前先调用chromium.CloseBrowser(true)，然后触发 chromium.SetOnClose
	m.TForm.SetOnCloseQuery(m.closeQuery)
	// 2. 触发后控制延迟关闭, 在UI线程中调用 windowParent.Free() 释放对象，然后触发 chromium.SetOnBeforeClose
	m.chromium.SetOnClose(m.chromiumClose)
	// 3. 触发后将canClose设置为true, 发送消息到主窗口关闭，触发 m.SetOnCloseQuery
	m.chromium.SetOnBeforeClose(m.chromiumBeforeClose)

	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("SetOnAfterCreated 1")
		cef.RunOnMainThread(func() {
			fmt.Println("SetOnAfterCreated 2")
			wp := cef.NewWindowProperty()
			m.ChildForm = cef.NewLCLWindow(wp, m)
		})
		fmt.Println("SetOnAfterCreated 3")
	})
	m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
		fmt.Println("SetOnBeforeBrowser 1")
		m.windowParent.UpdateSize()
		return false
	})

	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo,
		popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, client *cef.ICefClient, browserSettings *cef.TCefBrowserSettings,
		resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
		fmt.Println("beforePopupInfo:", beforePopupInfo.TargetUrl, beforePopupInfo.TargetDisposition, beforePopupInfo.TargetFrameName, beforePopupInfo.UserGesture)
		fmt.Println(*noJavascriptAccess)
		fmt.Println(browser.BrowserId(), frame.Identifier(), frame.Url(), frame.V8Context().Frame().Url())
		fmt.Printf("windowInfo: %+v\n", windowInfo)
		fmt.Printf("browserSettings: %+v\n", browserSettings)
		fmt.Printf("popupFeatures: %+v\n", popupFeatures)
		cef.RunOnMainThread(func() {
			m.ChildForm.ChromiumCreate(nil, beforePopupInfo.TargetUrl)
			m.ChildForm.Show()
		})
		return true
	})
	m.chromium.SetOnRenderCompMsg(func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
		//fmt.Println("SetOnRenderCompMsg", *lResult, *aHandled)
		//*aHandled = true
	})
	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, model *cef.ICefMenuModel) {
		fmt.Println("SetOnBeforeContextMenu")
	})
	m.chromium.SetOnContextMenuCommand(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32) bool {
		fmt.Println("SetOnContextMenuCommand")
		return false
	})
	m.chromium.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, callback *cef.ICefCallback, result *consts.TCefReturnValue) {
		fmt.Println("SetOnBeforeResourceLoad", frame.Url())
	})
	m.chromium.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
		fmt.Println("OnLoadStart")
	})
	m.chromium.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
		fmt.Println("OnLoadEnd")
	})
	m.chromium.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
		fmt.Println("OnLoadingProgressChange")
	})
	m.chromium.SetOnLoadingStateChange(func(sender lcl.IObject, browser *cef.ICefBrowser, isLoading, canGoBack, canGoForward bool) {
		fmt.Println("OnLoadingStateChange")
	})
	m.chromium.SetOnPreKeyEvent(func(sender lcl.IObject, browser *cef.ICefBrowser, event *cef.TCefKeyEvent, osEvent consts.TCefEventHandle) (isKeyboardShortcut, result bool) {
		fmt.Println("OnPreKeyEvent", event.WindowsKeyCode)
		return
	})
}

func (m *BrowserWindow) OnMessages() {
	m.SetOnWMPaint(func(message *et.TPaint) {
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
	})
	m.SetOnWMMove(func(message *et.TMove) {
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
	})
	m.SetOnWMSize(func(message *et.TSize) {
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
	})
	m.SetOnWMWindowPosChanged(func(message *et.TWindowPosChanged) {
		if m.Chromium() != nil {
			m.Chromium().NotifyMoveOrResizeStarted()
		}
	})
}

func (m *BrowserWindow) createBrowser(sender lcl.IObject) {
	fmt.Println("createBrowser")
	m.timer.SetEnabled(false)
	createBrowserOK := m.chromium.CreateBrowser(m.windowParent, "", nil, nil)
	initializedOK := m.chromium.Initialized()
	if !createBrowserOK && !initializedOK {
		m.timer.SetEnabled(true)
	}
	fmt.Println("CreateBrowser:", createBrowserOK, "Initialized:", initializedOK)
	//m.chromium.LoadUrl("https://www.baidu.com")
	//m.chromium.LoadUrl("https://www.gitee.com")
	//m.chromium.LoadUrl("http://localhost:22022")
}
func (m *BrowserWindow) resize(sender lcl.IObject) {
	fmt.Println("resize")
	if m.chromium != nil {
		m.chromium.NotifyMoveOrResizeStarted()
		if m.windowParent != nil {
			m.windowParent.UpdateSize()
		}
	}
}
func (m *BrowserWindow) closeQuery(sender lcl.IObject, canClose *bool) {
	fmt.Println("closeQuery")
	*canClose = m.canClose
	if !m.canClose {
		m.canClose = true
		m.chromium.CloseBrowser(true)
		//m.SetVisible(false)
	}
}

func (m *BrowserWindow) chromiumClose(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
	fmt.Println("chromiumClose")
	*aAction = consts.CbaDelay
	cef.RunOnMainThread(func() {
		m.windowParent.Free()
	})
}

func (m *BrowserWindow) chromiumBeforeClose(sender lcl.IObject, browser *cef.ICefBrowser) {
	fmt.Println("chromiumBeforeClose")
	m.canClose = true
	if !common.IsWindows() {
		cef.RunOnMainThread(func() {
			m.Close()
		})
	} else {
		rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
	}
}

func (m *BrowserWindow) createMenu() {
	mainMenu := lcl.NewMainMenu(m)
	// 创建一级菜单
	fileClassA := lcl.NewMenuItem(m)
	fileClassA.SetCaption("文件(&F)") //菜单名称 alt + f
	aboutClassA := lcl.NewMenuItem(m)
	aboutClassA.SetCaption("关于(&A)")

	var createMenuItem = func(label, shortCut string, click func(lcl.IObject)) (result *lcl.TMenuItem) {
		result = lcl.NewMenuItem(m)
		result.SetCaption(label)               //菜单项显示的文字
		result.SetShortCutFromString(shortCut) // 快捷键
		result.SetOnClick(click)               // 触发事件，回调函数
		return
	}
	// 给一级菜单添加菜单项
	createItem := createMenuItem("新建(&N)", "Meta+N", func(lcl.IObject) {
		fmt.Println("单击了新建")
	})
	fileClassA.Add(createItem) // 把创建好的菜单项添加到 第一个菜单中
	openItem := createMenuItem("打开(&O)", "Meta+O", func(lcl.IObject) {
		fmt.Println("单击了打开")
	})
	fileClassA.Add(openItem) // 把创建好的菜单项添加到 第一个菜单中
	mainMenu.Items().Add(fileClassA)
	mainMenu.Items().Add(aboutClassA)
	if common.IsDarwin() {
		// https://wiki.lazarus.freepascal.org/Mac_Preferences_and_About_Menu
		// 动态添加的，静态好像是通过设计器将顶级的菜单标题设置为应用程序名，但动态的就是另一种方式
		appMenu := lcl.NewMenuItem(m)
		// 动态添加的，设置一个Unicode Apple logo char
		appMenu.SetCaption(types.AppleLogoChar)
		subItem := lcl.NewMenuItem(m)

		subItem.SetCaption("关于")
		subItem.SetOnClick(func(sender lcl.IObject) {
			lcl.ShowMessage("About")
		})
		appMenu.Add(subItem)

		subItem = lcl.NewMenuItem(m)
		subItem.SetCaption("-")
		appMenu.Add(subItem)

		subItem = lcl.NewMenuItem(m)
		subItem.SetCaption("首选项")
		subItem.SetShortCutFromString("Meta+,")
		subItem.SetOnClick(func(sender lcl.IObject) {
			lcl.ShowMessage("Preferences")
		})
		appMenu.Add(subItem)
		// 添加
		mainMenu.Items().Insert(0, appMenu)
	}
}
func startServer() {
	fmt.Println("主进程启动 创建一个内置http服务")
	server := assetserve.NewAssetsHttpServer()
	server.PORT = 22022
	server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
	server.Assets = exampleCommon.ResourcesFS()
	go server.StartHttpServer()
}
