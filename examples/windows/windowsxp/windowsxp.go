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
	m.timer = lcl.NewTimer(m)
	m.timer.SetOnTimer(func(sender lcl.IObject) {
		fmt.Println("SetOnTimer")
		m.timer.SetEnabled(false)
		if !m.chromium.CreateBrowser(m.windowParent, "", nil, nil) &&
			!m.chromium.Initialized() {
			m.timer.SetEnabled(true)
		}
	})

	m.SetOnShow(func(sender lcl.IObject) {
		fmt.Println("SetOnShow")
		m.timer.SetEnabled(false)
		if !m.chromium.CreateBrowser(m.windowParent, "", nil, nil) &&
			!m.chromium.Initialized() {
			m.timer.SetEnabled(true)
		}
	})

	m.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) {
		fmt.Println("SetOnCloseQuery")
		*canClose = m.canClose
		if !m.canClose {
			m.canClose = true
			m.chromium.CloseBrowser(true)
			//m.SetVisible(false)
		}
	})

	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("SetOnAfterCreated")
		m.chromium.LoadUrl("https://www.baidu.com")
	})

	m.chromium.SetOnClose(func(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
		fmt.Println("chromium SetOnClose")
		*aAction = consts.CbaDelay
		cef.RunOnMainThread(func() {
			m.windowParent.Free()
		})
	})

	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("chromium SetOnBeforeClose")
		m.canClose = true
		rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
	})
}
