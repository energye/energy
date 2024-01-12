package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/consts/messages"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
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
	app := cef.CreateApplication()
	cef.SetApplication(app)
	// 指定 CEF Framework
	app.SetFrameworkDirPath("E:\\app\\energy\\EnergyFramework\\")
	if app.StartMainProcess() {
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
	fmt.Println("app free")
}

func (m *BrowserWindow) OnFormCreate(sender lcl.IObject) {
	m.ScreenCenter()
	m.OnMessages()
	m.SetOnWndProc(func(msg *types.TMessage) {
		m.InheritedWndProc(msg)
		//fmt.Println("SetOnWndProc")
	})
	m.chromium = cef.NewChromium(m, nil)
	m.chromium.SetDefaultURL("https://www.baidu.com")
	m.windowParent = cef.NewCEFWindowParent(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)
	m.windowParent.SetChromium(m.chromium, 0)
	// 创建一个定时器, 用来createBrowser
	m.timer = lcl.NewTimer(m)
	m.timer.SetOnTimer(m.show)
	// 在show时创建chromium browser
	m.TForm.SetOnShow(m.show)
	m.TForm.SetOnActivate(m.active)
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
		//fmt.Println("SetOnBeforeResourceLoad", frame.Url())
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

func (m *BrowserWindow) active(sender lcl.IObject) {
	fmt.Println("active")
	m.timer.SetEnabled(false)
	if !m.chromium.CreateBrowser(m.windowParent, "", nil, nil) &&
		!m.chromium.Initialized() {
		m.timer.SetEnabled(true)
	}
}
func (m *BrowserWindow) show(sender lcl.IObject) {
	fmt.Println("show")
	m.timer.SetEnabled(false)
	if !m.chromium.CreateBrowser(m.windowParent, "", nil, nil) &&
		!m.chromium.Initialized() {
		m.timer.SetEnabled(true)
	}
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
	rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
}
