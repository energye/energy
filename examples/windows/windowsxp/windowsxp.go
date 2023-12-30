package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/consts/messages"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
)

type BrowserWindow struct {
	*lcl.TForm
	timer        *lcl.TTimer
	windowParent *cef.TCEFWindowParent
	chromium     cef.IChromium
	canClose     bool
}

var BW *BrowserWindow

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	app := cef.CreateApplication()
	app.SetFrameworkDirPath("E:\\app\\energy\\EnergyFramework\\")
	if app.StartMainProcess() {
		// 结束应用后释放资源
		api.SetReleaseCallback(func() {
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
	m.SetOnShow(m.show)
	// 1. 关闭之前先调用chromium.CloseBrowser(true)，然后触发 chromium.SetOnClose
	m.SetOnCloseQuery(m.closeQuery)
	// 2. 触发后控制延迟关闭, 在UI线程中调用 windowParent.Free() 释放对象，然后触发 chromium.SetOnBeforeClose
	m.chromium.SetOnClose(m.chromiumClose)
	// 3. 触发后将canClose设置为true, 发送消息到主窗口关闭，触发 m.SetOnCloseQuery
	m.chromium.SetOnBeforeClose(m.chromiumBeforeClose)

	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, client *cef.ICefClient, browserSettings *cef.TCefBrowserSettings, noJavascriptAccess *bool) bool {
		fmt.Println("beforePopupInfo:", beforePopupInfo.TargetUrl, beforePopupInfo.TargetDisposition, beforePopupInfo.TargetFrameName, beforePopupInfo.UserGesture)
		fmt.Println(*noJavascriptAccess)
		fmt.Println(browser.BrowserId(), frame.Identifier(), frame.Url(), frame.V8Context().Frame().Url())
		fmt.Printf("windowInfo: %+v\n", windowInfo)
		fmt.Printf("browserSettings: %+v\n", browserSettings)
		fmt.Printf("popupFeatures: %+v\n", popupFeatures)
		return true
	})
	//m.chromium.SetOnOpenUrlFromTab(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool {
	//	return true
	//})
}

func (m *BrowserWindow) show(sender lcl.IObject) {
	fmt.Println("show")
	m.timer.SetEnabled(false)
	if !m.chromium.CreateBrowser(m.windowParent, "", nil, nil) &&
		!m.chromium.Initialized() {
		m.timer.SetEnabled(true)
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
