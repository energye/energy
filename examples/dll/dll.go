package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/consts/messages"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
)

type TMainForm struct {
	*lcl.TForm
	btn *lcl.TButton
}

var (
	MainForm       *TMainForm
	hostWindowHwnd uintptr
)

func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	Println("Libenergy.dll OnFormCreate")
	m.ScreenCenter()
	m.SetCaption("energy-window")
	m.btn = lcl.NewButton(m)
	m.btn.SetParent(m)
	m.btn.SetCaption("TestBtn")
	m.btn.SetOnClick(func(sender lcl.IObject) {
		Println("Libenergy.dll OnFormCreate")
	})
	m.SetOnDestroy(func(sender lcl.IObject) {
		Println("Libenergy.dll SetOnDestroy")
	})
}

func main() {
}

func Println(s ...interface{}) {
	fmt.Println(s...)
}

type BrowserWindow struct {
	*lcl.TForm
	mainWindowId int32 // 主窗口ID
	timer        *lcl.TTimer
	windowParent cef.ICEFWindowParent
	chromium     cef.IChromium
	canClose     bool
}

var BW *BrowserWindow

func (m *BrowserWindow) OnFormCreate(sender lcl.IObject) {
	m.ScreenCenter()
	m.SetCaption("Energy DLL - CEF simple")
	m.chromium = cef.NewChromium(m, nil)
	m.chromium.SetDefaultURL(cef.BrowserWindow.Config.Url)
	m.windowParent = cef.NewCEFWindowParent(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)
	m.windowParent.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
	// 创建一个定时器, 用来createBrowser
	m.timer = lcl.NewTimer(m)
	m.timer.SetEnabled(false)
	m.timer.SetInterval(200)
	m.timer.SetOnTimer(m.createBrowser)
	// 在show时创建chromium browser
	m.TForm.SetOnShow(m.show)
	m.TForm.SetOnActivate(m.active)
	m.TForm.SetOnResize(m.resize)
	m.windowParent.SetOnEnter(func(sender lcl.IObject) {
		m.chromium.Initialized()
		m.chromium.FrameIsFocused()
		m.chromium.SetFocus(true)
	})
	m.windowParent.SetOnExit(func(sender lcl.IObject) {
		m.chromium.SendCaptureLostEvent()
	})
	// 1. 关闭之前先调用chromium.CloseBrowser(true)，然后触发 chromium.SetOnClose
	m.TForm.SetOnCloseQuery(m.closeQuery)
	// 2. 触发后控制延迟关闭, 在UI线程中调用 windowParent.Free() 释放对象，然后触发 chromium.SetOnBeforeClose
	m.chromium.SetOnClose(m.chromiumClose)
	// 3. 触发后将canClose设置为true, 发送消息到主窗口关闭，触发 m.SetOnCloseQuery
	m.chromium.SetOnBeforeClose(m.chromiumBeforeClose)
	m.chromium.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
		Println("OnLoadingProgressChange:", progress)
	})
	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		Println("SetOnAfterCreated 1")
		cef.RunOnMainThread(func() {
			Println("SetOnAfterCreated 2")
		})
		Println("SetOnAfterCreated 3")
		if m.mainWindowId == 0 {
			m.mainWindowId = browser.Identifier()
		}
		m.windowParent.UpdateSize()
	})
	m.chromium.SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
		Println("SetOnBeforeBrowser")
		m.windowParent.UpdateSize()
		return false
	})
}

func (m *BrowserWindow) createBrowser(sender lcl.IObject) {
	if m.timer == nil {
		return
	}
	m.timer.SetEnabled(false)
	rect := m.ClientRect()
	init := m.chromium.Initialized()
	created := m.chromium.CreateBrowserByWindowHandle(consts.TCefWindowHandle(m.windowParent.Handle()), rect, "", nil, nil, false)
	Println("createBrowser rect:", rect, "init:", init, "create:", created)
	if !created {
		m.timer.SetEnabled(true)
	} else {
		m.windowParent.UpdateSize()
		m.timer.Free()
		m.timer = nil
	}
}

func (m *BrowserWindow) active(sender lcl.IObject) {
	Println("window active")
	m.createBrowser(sender)
}

func (m *BrowserWindow) show(sender lcl.IObject) {
	Println("window show")
	m.createBrowser(sender)
}

func (m *BrowserWindow) resize(sender lcl.IObject) {
	Println("resize")
	if m.chromium != nil {
		m.chromium.NotifyMoveOrResizeStarted()
		if m.windowParent != nil {
			m.windowParent.UpdateSize()
		}
	}
}
func (m *BrowserWindow) closeQuery(sender lcl.IObject, canClose *bool) {
	Println("closeQuery canClose:", m.canClose)
	*canClose = m.canClose
	if !m.canClose {
		m.canClose = true
		m.chromium.CloseBrowser(true)
		//m.SetVisible(false)
	}
}

func (m *BrowserWindow) chromiumClose(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
	Println("chromiumClose id:", browser.Identifier(), "mainWindowId:", m.mainWindowId)
	if browser.Identifier() == m.mainWindowId {
		*aAction = consts.CbaDelay
		cef.RunOnMainThread(func() {
			Println("chromiumClose windowParent.Free()")
			m.windowParent.Free()
		})
	}
}

func (m *BrowserWindow) chromiumBeforeClose(sender lcl.IObject, browser *cef.ICefBrowser) {
	Println("chromiumBeforeClose id:", browser.Identifier(), "mainWindowId:", m.mainWindowId)
	if browser.Identifier() == m.mainWindowId {
		m.canClose = true
		rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0)
		if hostWindowHwnd != 0 {
			//rtl.PostMessage(hostWindowHwnd, messages.WM_CLOSE, 0, 0)
		}
	}
}
