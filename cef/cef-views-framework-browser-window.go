//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
)

//基于CEF views framework窗口
//
//该窗口使用CEF内部实现，在linux下107.xx以后版本默认使用GTK3，但无法使用lcl组件集成到窗口中
//
//当创建应用配置时 MultiThreadedMessageLoop 和 ExternalMessagePump 属性同时为false(linux系统默认强制false)时启用ViewsFramework窗口
type ViewsFrameworkBrowserWindow struct {
	component            lcl.IComponent            //
	windowComponent      *TCEFWindowComponent      //
	browserViewComponent *TCEFBrowserViewComponent //
}

func (m *ViewsFrameworkBrowserWindow) Component() lcl.IComponent {
	return m.component
}

func (m *ViewsFrameworkBrowserWindow) WindowComponent() *TCEFWindowComponent {
	return m.windowComponent
}

func (m *ViewsFrameworkBrowserWindow) BrowserViewComponent() *TCEFBrowserViewComponent {
	return m.browserViewComponent
}

func (m *browserWindow) appContextInitialized(app *TCEFApplication) {
	if !common.Args.IsMain() {
		return
	}
	app.SetOnContextInitialized(func() {
		m.vFrameBrowserWindow = &ViewsFrameworkBrowserWindow{}
		m.vFrameBrowserWindow.component = lcl.NewComponent(nil)
		if BrowserWindow.Config.chromiumConfig == nil {
			BrowserWindow.Config.chromiumConfig = NewChromiumConfig()
			BrowserWindow.Config.chromiumConfig.SetEnableMenu(true)
			BrowserWindow.Config.chromiumConfig.SetEnableDevTools(true)
			BrowserWindow.Config.chromiumConfig.SetEnableOpenUrlTab(true)
			BrowserWindow.Config.chromiumConfig.SetEnableWindowPopup(true)
		}
		m.chromium = NewChromium(m.vFrameBrowserWindow.component, BrowserWindow.Config.chromiumConfig)
		m.chromium.SetEnableMultiBrowserMode(true)
		m.vFrameBrowserWindow.browserViewComponent = NewBrowserViewComponent(m.vFrameBrowserWindow.component)
		m.vFrameBrowserWindow.windowComponent = NewWindowComponent(m.vFrameBrowserWindow.component)

		m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup TargetUrl:", beforePopupInfo.TargetUrl)

			return false
		})
		m.vFrameBrowserWindow.windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *ICefWindow) {
			if m.chromium.CreateBrowserByBrowserViewComponent(BrowserWindow.Config.DefaultUrl, m.vFrameBrowserWindow.browserViewComponent) {
				m.vFrameBrowserWindow.windowComponent.AddChildView(m.vFrameBrowserWindow.browserViewComponent)
				m.vFrameBrowserWindow.windowComponent.SetTitle(BrowserWindow.Config.Title)
				window.CenterWindow(NewCefSize(BrowserWindow.Config.Width, BrowserWindow.Config.Height))
				m.vFrameBrowserWindow.browserViewComponent.RequestFocus()
				if BrowserWindow.Config.Icon != "" {
					window.SetWindowAppIconFS(1, BrowserWindow.Config.Icon)
				}
				if BrowserWindow.Config.viewsFrameBrowserWindowOnEventCallback != nil {
					BrowserWindow.browserEvent.chromium = m.chromium
					BrowserWindow.Config.viewsFrameBrowserWindowOnEventCallback(BrowserWindow.browserEvent, m.vFrameBrowserWindow)
				}
				window.Show()
			}
		})
		m.vFrameBrowserWindow.windowComponent.SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
			fmt.Println("OnCanClose")
			*aResult = true
			app.QuitMessageLoop()
		})
		m.vFrameBrowserWindow.windowComponent.SetOnGetInitialBounds(func(sender lcl.IObject, window *ICefWindow, aResult *TCefRect) {
			fmt.Println("OnGetInitialBounds")
		})
		m.vFrameBrowserWindow.windowComponent.SetOnGetInitialShowState(func(sender lcl.IObject, window *ICefWindow, aResult *consts.TCefShowState) {
			fmt.Println("OnGetInitialShowState", *aResult)
		})
		m.vFrameBrowserWindow.windowComponent.CreateTopLevelWindow()
	})
}
