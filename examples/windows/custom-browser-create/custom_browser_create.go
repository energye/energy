//go:build windows
// +build windows

package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/consts/messages"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type MainWindowDemo struct {
	*lcl.TForm
}

//自定义浏览器创建
//示例演示将CEF做为一个LCL组件使用
//在适当的情况动态创建一个或多个浏览器，实际可做为子组件使用或弹出窗口，根据需求而定
//需要自己正确处理关闭流程
//可以自己定制出CEF的使用, 不同平台有些区别.
//MacOS下使用方式和Windows基本一至
//Linux目前需要在Gtk2下并且使用CEF106版本, 默认Linux大于CEF106版本使用的是GTK3,
//Linux可以使用VF窗口来自定义创建

func main() {
	cef.GlobalInit(nil, nil)
	defer func() {
		api.EnergyLibRelease()
	}()
	//创建应用
	app := cef.NewApplication()
	// 启动主进程
	success := app.StartMainProcess()
	if success {
		// 创建窗口并运行应用
		var window = &MainWindowDemo{}
		lcl.RunApp(&window)
	}
}

func (m *MainWindowDemo) OnFormCreate(sender lcl.IObject) {
	fmt.Println("MainWindowDemo OnFormCreate")

	m.SetCaption("主窗口")
	m.SetWidth(1000)
	m.SetHeight(900)
	m.ScreenCenter()

	// 普通Panel
	panel := lcl.NewPanel(m)
	panel.SetParent(m)
	panel.SetWidth(m.Width())
	panel.SetHeight(800)
	panel.SetTop(100)
	panel.SetColor(colors.ClCyan)
	panel.SetAnchors(types.NewSet(types.AkTop, types.AkRight, types.AkBottom, types.AkLeft))
	labTxt := lcl.NewLabel(m)
	labTxt.SetParent(panel)
	labTxt.SetCaption("普通Panel") //标识一下确定是没有浏览器的panel

	// 带有浏览器的panel
	browserPanel := lcl.NewPanel(m)
	browserPanel.SetParent(m)
	browserPanel.SetWidth(m.Width())
	browserPanel.SetHeight(800)
	browserPanel.SetTop(100)
	browserPanel.SetVisible(false)
	browserPanel.SetAnchors(types.NewSet(types.AkTop, types.AkRight, types.AkBottom, types.AkLeft))

	// 普通Panel 按钮
	panelBtn := lcl.NewButton(m)
	panelBtn.SetParent(m)
	panelBtn.SetCaption("显示普通Panel")
	panelBtn.SetWidth(200)
	panelBtn.SetOnClick(func(sender lcl.IObject) {
		panel.SetVisible(true)
		browserPanel.SetVisible(false)
	})

	// 显示浏览器Panel 按钮
	browserPanelBtn := lcl.NewButton(m)
	browserPanelBtn.SetParent(m)
	browserPanelBtn.SetCaption("显示浏览器Panel")
	browserPanelBtn.SetWidth(200)
	browserPanelBtn.SetLeft(250)
	//在按钮事件中创建浏览器，并设置到指定的panel中
	//需要注意的是，浏览器有一个正确的关闭流程。
	//在窗口的关闭事件中 close, closeQuery
	var chromiumBrowser cef.ICEFChromiumBrowser
	//窗口关闭流程标识
	var (
		canClose  bool
		isClosing bool
	)
	browserPanelBtn.SetOnClick(func(sender lcl.IObject) {
		panel.SetVisible(false)
		browserPanel.SetVisible(true)
		if chromiumBrowser == nil {
			chromiumBrowser = cef.NewChromiumBrowser(browserPanel, nil)
			chromiumBrowser.Chromium().SetEnableMultiBrowserMode(true)
			chromiumBrowser.Chromium().SetDefaultURL("https://www.baidu.com")
			//windowParent
			chromiumBrowser.WindowParent().DefaultAnchors()
			chromiumBrowser.WindowParent().SetOnEnter(func(sender lcl.IObject) {
				chromiumBrowser.Chromium().Initialized()
				chromiumBrowser.Chromium().FrameIsFocused()
				chromiumBrowser.Chromium().SetFocus(true)
			})
			chromiumBrowser.WindowParent().SetOnExit(func(sender lcl.IObject) {
				chromiumBrowser.Chromium().SendCaptureLostEvent()
			})
			//禁用右键菜单
			chromiumBrowser.Chromium().SetOnBeforeContextMenu(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, model *cef.ICefMenuModel) {
				model.Clear()
			})
			//禁止弹出新窗口
			chromiumBrowser.Chromium().SetOnOpenUrlFromTab(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool {
				fmt.Println("OnOpenUrlFromTab")
				return true
			})
			//禁止弹出新窗口
			chromiumBrowser.Chromium().SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
				fmt.Println("OnBeforePopup")
				return true
			})

			// 以下是窗口关闭时 正确关闭浏览器的流程
			// windows的关闭示例
			chromiumBrowser.Chromium().SetOnClose(func(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
				*aAction = consts.CbaDelay
				cef.QueueAsyncCall(func(id int) { //run in main thread
					chromiumBrowser.WindowParent().Free()
				})
			})
			chromiumBrowser.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
				canClose = true                   // 允许关闭窗口
				cef.QueueAsyncCall(func(id int) { // main thread run
					rtl.PostMessage(m.Handle(), messages.WM_CLOSE, 0, 0) //发送关闭消息
				})
			})
			chromiumBrowser.CreateBrowser()
		}
	})
	// 以下是窗口关闭时 正确关闭浏览器的流程
	// windows的关闭示例
	m.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		*action = types.CaFree
	})
	m.SetOnCloseQuery(func(sender lcl.IObject, close *bool) {
		*close = canClose // 是否允许关闭窗口标识
		cef.QueueAsyncCall(func(id int) {
			if !isClosing {
				isClosing = true
				if chromiumBrowser != nil {
					chromiumBrowser.Chromium().CloseBrowser(true) // 关闭浏览器
				}
				m.Hide()
			}
		})
	})
}
