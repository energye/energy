package main

import (
	"fmt"
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/cef"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/messages"
	"os"
	"path/filepath"
)

type BrowserWindow struct {
	lcl.TForm
	mainWindowId int32 // 主窗口ID
	timer        lcl.ITimer
	windowParent cef.ICEFWindowParent
	chromium     cef.IChromium
	canClose     bool
	ChildForm    lcl.IForm
}

var BW BrowserWindow

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	app := cef.NewCefApplication()
	cef.SetGlobalCEFApp(app)
	// 指定 CEF Framework
	frameworkDir := os.Getenv("ENERGY_HOME")
	app.SetFrameworkDirPath(frameworkDir)
	app.SetResourcesDirPath(frameworkDir)
	app.SetLocalesDirPath(filepath.Join(frameworkDir, "locales"))
	if app.StartMainProcess() {
		// 结束应用后释放资源
		api.SetReleaseCallback(func() {
			fmt.Println("Release")
			app.Free()
		})
		// LCL窗口
		lcl.Application.Initialize()
		lcl.Application.SetMainFormOnTaskBar(true)
		lcl.Application.CreateForm(&BW)
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
	m.chromium = cef.NewChromium(m)
	m.chromium.SetDefaultUrl("https://www.baidu.com")
	m.windowParent = cef.NewCEFWindowParent(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)
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

	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser cef.ICefBrowser) {
		fmt.Println("SetOnAfterCreated 1")
		lcl.RunOnMainThreadAsync(func(id uint32) {
			fmt.Println("SetOnAfterCreated 2")
			//wp := cef.NewWindowProperty()
			//m.ChildForm = cef.NewLCLWindow(wp, m)
		})
		fmt.Println("SetOnAfterCreated 3")
		if m.mainWindowId == 0 {
			m.mainWindowId = browser.GetIdentifier()
		}
	})
	m.chromium.SetOnBeforeBrowse(func(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
		userGesture, isRedirect bool, result *bool) {
		fmt.Println("SetOnBeforeBrowser 1")
	})

	m.chromium.SetOnBeforePopup(func(sender cef.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, beforePopup cef.TBeforePopup, popupFeatures cef.TCefPopupFeatures,
		windowInfo *cef.TCefWindowInfo, settings *cef.TCefBrowserSettings) (
		client cef.ICefClient, extraInfo cef.ICefDictionaryValue, noJavascriptAccess, result bool) {
		fmt.Printf("beforePopup: %+v\n", beforePopup)
		fmt.Printf("popupFeatures: %+v\n", popupFeatures)
		fmt.Println(browser.GetIdentifier())
		fmt.Println(frame.GetIdentifier(), frame.GetUrl())
		v8ctx := frame.GetV8Context()
		if v8ctx != nil {
			fmt.Println(frame.GetV8Context())
			fmt.Println(frame.GetV8Context().GetFrame().GetUrl())
		}
		lcl.RunOnMainThreadAsync(func(id uint32) {
			//m.ChildForm.ChromiumCreate(nil, beforePopupInfo.TargetUrl)
			//m.ChildForm.Show()
		})
		//windowInfo.Y = 100
		settings.DefaultFontSize = 36
		settings.StandardFontFamily = "微软雅黑"
		windowInfo.X = 400
		windowInfo.Y = 10
		windowInfo.Width = 400
		windowInfo.Height = 400
		windowInfo.WindowName = "杨杨红红岩岩"
		//result = true
		return
	})

	m.chromium.SetOnRenderCompMsg(func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, aHandled *bool) {
		//fmt.Println("SetOnRenderCompMsg", *lResult, *aHandled)
		//*aHandled = true
	})

	m.chromium.SetOnBeforeContextMenu(func(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams, model cef.ICefMenuModel) {
		fmt.Println("SetOnBeforeContextMenu")
	})
	m.chromium.SetOnContextMenuCommand(func(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams, commandId cef.MenuId, eventFlags uint32, result *bool) {
		fmt.Println("SetOnContextMenuCommand")
	})
	m.chromium.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest, callback cef.ICefCallback, result *cef.TCefReturnValue) {
		fmt.Println("SetOnBeforeResourceLoad")
	})
}

func (m *BrowserWindow) OnMessages() {
	//m.SetOnWMPaint(func(message *messages.TPaint) {
	//	if m.Chromium() != nil {
	//		m.Chromium().NotifyMoveOrResizeStarted()
	//	}
	//})
	//m.SetOnWMMove(func(message *messages.TMove) {
	//	if m.Chromium() != nil {
	//		m.Chromium().NotifyMoveOrResizeStarted()
	//	}
	//})
	//m.SetOnWMSize(func(message *messages.TSize) {
	//	if m.Chromium() != nil {
	//		m.Chromium().NotifyMoveOrResizeStarted()
	//	}
	//})
	//m.SetOnWMWindowPosChanged(func(message *messages.TWindowPosChanged) {
	//	if m.Chromium() != nil {
	//		m.Chromium().NotifyMoveOrResizeStarted()
	//	}
	//})
}

func (m *BrowserWindow) active(sender lcl.IObject) {
	m.timer.SetEnabled(false)
	rect := m.ClientRect()
	fmt.Println("active", rect)
	if !m.chromium.CreateBrowser(m.windowParent.Handle(), &rect, "", nil, nil, false) &&
		!m.chromium.Initialized() {
		m.timer.SetEnabled(true)
	}
}
func (m *BrowserWindow) show(sender lcl.IObject) {
	m.timer.SetEnabled(false)
	rect := m.ClientRect()
	fmt.Println("show", rect)
	if !m.chromium.CreateBrowser(m.windowParent.Handle(), &rect, "", nil, nil, false) &&
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

func (m *BrowserWindow) chromiumClose(sender lcl.IObject, browser cef.ICefBrowser, aAction *cef.TCefCloseBrowserAction) {
	fmt.Println("chromiumClose id:", browser.GetIdentifier(), "mainWindowId:", m.mainWindowId)
	if browser.GetIdentifier() == m.mainWindowId {
		*aAction = cef.CbaDelay
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.windowParent.Free()
		})
	}
}

func (m *BrowserWindow) chromiumBeforeClose(sender lcl.IObject, browser cef.ICefBrowser) {
	fmt.Println("chromiumBeforeClose id:", browser.GetIdentifier(), "mainWindowId:", m.mainWindowId)
	if browser.GetIdentifier() == m.mainWindowId {
		m.canClose = true
		rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
	}
}
