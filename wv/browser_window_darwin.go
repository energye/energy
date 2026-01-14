//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package wv

import (
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/pkgs/mime"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	wv "github.com/energye/wv/darwin"
	wvTypes "github.com/energye/wv/types/darwin"
	"unsafe"
)

var (
	frameWidth  = int32(4)
	frameHeight = int32(4)
	frameCorner = frameWidth + frameHeight
)

type TWebview struct {
	lcl.ICustomPanel
	browserId               uint32
	isClose                 bool
	isCreated               bool
	resizeHT                string
	window                  window.IWindow
	windowParent            wv.IWkWebviewParent
	browser                 wv.IWkWebview
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	onBrowserAfterCreated   lcl.TNotifyEvent
	onProcessMessage        TOnProcessMessageEvent
	onResourceRequest       TOnResourceRequestEvent
	onLoadChange            TOnLoadChangeEvent
	onContextMenu           TOnContextMenuEvent
	onContextMenuCommand    TOnContextMenuCommandEvent
	onPopupWindow           TOnPopupWindowEvent
}

// NewWebview 创建一个新的浏览器窗口实例
func NewWebview(owner lcl.IComponent) IWebview {
	m := &TWebview{browserId: getNextBrowserID()}
	m.ICustomPanel = lcl.NewPanel(owner)
	m.ICustomPanel.SetParentDoubleBuffered(true)
	m.ICustomPanel.SetBevelInner(types.BvNone)
	m.ICustomPanel.SetBevelOuter(types.BvNone)

	m.windowParent = wv.NewWebviewParent(m)
	m.windowParent.SetWidth(m.Width())
	m.windowParent.SetHeight(m.Height())
	m.windowParent.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))
	m.windowParent.SetParentDoubleBuffered(true)

	m.browser = wv.NewWebview(owner)

	userContentController := wv.UserContentController.New()
	scriptMessageHandler := wv.NewScriptMessageHandler(m.browser.AsReceiveScriptMessageDelegate())
	userContentController.AddScriptMessageHandlerName(scriptMessageHandler, energyProcessMessage)
	userScript := wv.UserScript.InitWithSourceInjectionTimeForMainFrameOnly(string(ipcJS), 0, false)
	userContentController.AddUserScript(userScript.Data())

	configuration := wv.WebViewConfiguration.New()
	configuration.SetUserContentController(userContentController.Data())

	URLSchemeHandler := wv.NewURLSchemeHandler(m.browser.AsWKURLSchemeHandlerDelegate())

	configuration.SetSuppressesIncrementalRendering(true)
	configuration.SetApplicationNameForUserAgent(energyApplicationName)

	if gApplication.onCustomSchemes != nil {
		customSchemes := &TCustomSchemes{}
		gApplication.onCustomSchemes(customSchemes)
		for _, scheme := range customSchemes.schemes {
			if !setRegisterSchemeCache(scheme.Scheme) {
				configuration.SetURLSchemeHandlerForURLScheme(URLSchemeHandler.Data(), scheme.Scheme)
			}
		}
	}
	if gApplication.LocalLoad != nil {
		if !setRegisterSchemeCache(gApplication.LocalLoad.Scheme) {
			configuration.SetURLSchemeHandlerForURLScheme(URLSchemeHandler.Data(), gApplication.LocalLoad.Scheme)
		}
	}

	preference := wv.NewPreferences(configuration.Preferences())
	configuration.SetPreferences(preference.Data())
	preference.SetTabFocusesLinks(true)
	preference.SetFraudulentWebsiteWarningEnabled(true)
	if !gApplication.Options.DisableDevTools {
		preference.SetBoolValueForKey(true, "developerExtrasEnabled")
	}
	//preference.EnableDevtools()

	navigationDelegate := wv.NewNavigationDelegate(m.browser.AsWKNavigationDelegate())
	uiDelegate := wv.NewUIDelegate(m.browser.AsWKUIDelegate())

	frame := types.TRect{}
	frame.SetWidth(m.Width())
	frame.SetHeight(m.Height())
	m.browser.InitWithFrameConfiguration(frame, configuration.Data())

	m.browser.SetNavigationDelegate(navigationDelegate.Data())
	m.browser.SetUIDelegate(uiDelegate.Data())

	// ipc message received
	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	ipc.RegisterProcessMessage(m)
	m.initDefaultEvent()
	m.SetBrowserOptions()
	return m
}

func (m *TWebview) SetWindow(window window.IWindow) {
	m.window = window
	if m.window != nil {
		if m.window.BrowserId() == 0 {
			m.window.SetBrowserId(m.browserId)
		}
		m.window.SetOptions()
	}
	window.SetOnWindowShow(m.onWindowShow)
	window.SetOnWindowClose(m.onWindowClose)
	window.SetOnWindowCloseQuery(m.onWindowCloseQuery)
}

// SetBrowserOptions 设置浏览器窗口的选项配置
func (m *TWebview) SetBrowserOptions() {
	options := gApplication.Options
	if options.DefaultURL != "" {
		m.browser.LoadURL(options.DefaultURL)
	}
}

// SetParent 设置浏览器窗口的父控件
// 该方法会同时设置内部面板的父控件和窗口父控件的引用
func (m *TWebview) SetParent(window lcl.IWinControl) {
	m.ICustomPanel.SetParent(window)
	m.windowParent.SetParent(m)
}

// 在窗口显示时调用
func (m *TWebview) CreateBrowser() {
	if m.isCreated || m.window == nil {
		return
	}
	m.isCreated = true
	m.windowParent.SetWebview(m.browser.Data())
	if m.onBrowserAfterCreated != nil {
		m.onBrowserAfterCreated(m.browser)
	}
}

// BrowserId Window ID, generated by accumulating sequence numbers
func (m *TWebview) BrowserId() uint32 {
	return m.browserId
}

// SendMessage 发送消息到webview浏览器
func (m *TWebview) SendMessage(payload []byte) {
	if m.isClose {
		return
	}
	js := m.evalExecuteEventJS(payload)
	m.ExecuteScript(js)
}

func (m *TWebview) evalExecuteEventJS(js []byte) string {
	evalJS := "if (typeof window.energy !== 'undefined' && typeof window.energy.__executeEvent === 'function') {window.energy.__executeEvent('" + string(js) + "');}"
	return evalJS
}

// Close 关闭webview窗口并清理相关资源
func (m *TWebview) Close() {
	if m.isClose {
		return
	}
	m.isClose = true
	m.windowParent.Free()
	ipc.UnRegisterProcessMessage(m)
}

// SetDefaultURL 设置WebView的默认URL
func (m *TWebview) SetDefaultURL(url string) {
	m.browser.LoadURL(url)
}

// LoadURL 加载指定的URL地址到webview中
func (m *TWebview) LoadURL(url string) {
	m.browser.LoadURL(url)
}

// Browser 返回TWebview实例关联的浏览器对象
func (m *TWebview) Browser() wv.IWkWebview {
	return m.browser
}

// onWindowShow 是窗口显示事件的回调函数
// 当窗口显示时触发此函数，用于创建浏览器实例
func (m *TWebview) onWindowShow(sender lcl.IObject) {
	m.CreateBrowser()
}

// onWindowClose 处理窗口关闭事件的回调函数
// 当窗口接收到关闭信号时，该函数会停止浏览器实例以确保资源被正确释放
func (m *TWebview) onWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {
}

// onWindowCloseQuery 处理窗口关闭查询事件
// 当用户尝试关闭窗口时触发此回调函数
func (m *TWebview) onWindowCloseQuery(sender lcl.IObject, canClose *bool) {
	//*canClose = m.window.IsClose()
	if !m.window.IsClose() {
		m.window.SetClose(true)
		m.browser.StopLoading()
		m.browser.RemoveFromSuperview()
		m.browser.Release()
		m.windowParent.Free()
	}
	//if *canClose && m.isMainWindow {
	//	os.Exit(0)
	//}
}

// SetOnBrowserAfterCreated 设置浏览器创建后的回调事件处理函数
func (m *TWebview) SetOnBrowserAfterCreated(fn lcl.TNotifyEvent) {
	m.onBrowserAfterCreated = fn
}

// SetOnResourceRequest 设置资源请求事件处理函数
// 该方法用于注册一个回调函数，当webview发起资源请求时会触发此回调
func (m *TWebview) SetOnResourceRequest(fn TOnResourceRequestEvent) {
	m.onResourceRequest = fn
}

// SetOnProcessMessage 设置处理进程消息的回调函数
// 该方法用于注册一个回调函数，当接收到进程消息时会触发该回调
func (m *TWebview) SetOnProcessMessage(fn TOnProcessMessageEvent) {
	m.onProcessMessage = fn
}

func (m *TWebview) SetOnLoadChange(fn TOnLoadChangeEvent) {
	m.onLoadChange = fn
}

func (m *TWebview) SetOnContextMenu(fn TOnContextMenuEvent) {
	m.onContextMenu = fn
}

func (m *TWebview) SetOnContextMenuCommand(fn TOnContextMenuCommandEvent) {
	m.onContextMenuCommand = fn
}

func (m *TWebview) SetOnPopupWindow(fn TOnPopupWindowEvent) {
	m.onPopupWindow = fn
}

func (m *TWebview) ExecuteScript(javaScript string) {
	m.browser.EvaluateJavaScript(javaScript)
}

func (m *TWebview) initDefaultEvent() {
	getWindow := func() window.IDarwinWindow {
		if m.window == nil {
			return nil
		}
		return m.window.(window.IDarwinWindow)
	}
	m.browser.SetOnProcessMessage(func(sender lcl.IObject, userContentController wvTypes.WKUserContentController, name string, message string) {
		fmt.Println("OnProcessMessage", name, "message:", message)
		var handle bool
		if m.messageReceivedDelegate != nil {
			// ipc message
			var pMessage ipc.ProcessMessage
			err := json.Unmarshal([]byte(message), &pMessage)
			if err == nil {
				switch pMessage.Type {
				case ipc.MT_READY:
					// ipc ready
					handle = true
				case ipc.MT_EVENT_GO_EMIT, ipc.MT_EVENT_JS_EMIT, ipc.MT_EVENT_GO_EMIT_CALLBACK, ipc.MT_EVENT_JS_EMIT_CALLBACK:
					// ipc on, emit event
					handle = m.messageReceivedDelegate.Received(m.BrowserId(), &pMessage)
				case ipc.MT_DRAG_MOVE, ipc.MT_DRAG_DOWN, ipc.MT_DRAG_UP, ipc.MT_DRAG_DBLCLICK:
					// ipc drag window
					if window := getWindow(); window != nil {
						handle = true
					}
				case ipc.MT_DRAG_RESIZE:
					// border drag resize
					if m.window != nil {
						//m.resizeHT = pMessage.Data.(string)
						//m.window.Resize(ht)
						handle = true
					}
				case ipc.MT_DRAG_BORDER_WMSZ:
					//fmt.Println("pMessage.Data", pMessage.Data)
					//m._SetCursor(17)
				}
			}
		}
		if !handle && m.onProcessMessage != nil {
			m.onProcessMessage(message)
		}
	})
	m.browser.SetOnStartProvisionalNavigation(func(sender lcl.IObject, navigation wvTypes.WKNavigation) {
		fmt.Println("OnStartProvisionalNavigation")
	})
	m.browser.SetOnFinishNavigation(func(sender lcl.IObject, navigation wvTypes.WKNavigation) {
		fmt.Println("OnFinishNavigation")
		m.createEnergyJavasScript()
	})
	m.browser.SetOnDecidePolicyForNavigationActionPreferences(func(sender lcl.IObject, navigationAction wvTypes.WKNavigationAction, actionPolicy *wvTypes.WKNavigationActionPolicy, preferences *wvTypes.WKWebpagePreferences) {
		fmt.Println("OnDecidePolicyForNavigationActionPreferences")
	})
	m.browser.SetOnCreateWebView(func(sender lcl.IObject, configuration wvTypes.WKWebViewConfiguration, navigationAction wvTypes.WKNavigationAction, windowFeatures wvTypes.WKWindowFeatures) wvTypes.WKWebView {
		fmt.Println("OnCreateWebView")
		return 0
	})
	m.browser.SetOnWebViewDidClose(func(sender lcl.IObject) {
		fmt.Println("OnWebViewDidClose")
	})
	m.browser.SetOnWebContentProcessDidTerminate(func(sender lcl.IObject) {
		fmt.Println("OnWebContentProcessDidTerminate")
	})
	m.browser.SetOnStartURLSchemeTask(m.onStartURLSchemeTask)
	m.browser.SetOnStopURLSchemeTask(m.onStopURLSchemeTask)
}

func (m *TWebview) onStartURLSchemeTask(sender lcl.IObject, urlSchemeTask wvTypes.WKURLSchemeTask) {
	schemeTask := wv.NewURLSchemeTask(urlSchemeTask)
	request := wv.NewURLRequest(schemeTask.Request())
	url := wv.NewURL(request.URL())
	var (
		resource    string
		handle      bool
		uri         = url.AbsoluteString()
		path        = url.Path()
		method      = request.HTTPMethod()
		contentType = "Content-Type: " + mime.GetMimeType(path)
	)
	if m.onResourceRequest != nil {
		header := make(map[string]string)
		headers := request.AllHTTPHeaderFields()
		_ = json.Unmarshal([]byte(headers), &header)
		resource, handle = m.onResourceRequest(uri, path, method, header)
	}
	response := wv.URLResponse.New()
	finish := func(data []byte, length int32, err bool) {
		if err {
			schemeTask.FailWithError(404, string(data))
		} else {
			response.InitWithURLMIMETypeExpectedContentLengthTextEncodingName(url.Data(), contentType, length, "utf-8")
			schemeTask.ReceiveResponse(response.Data())
			schemeTask.ReceiveData(uintptr(unsafe.Pointer(&data[0])), length)
			schemeTask.Finish()
		}
	}
	if !handle {
		data, err := gApplication.LocalLoad.Read(path)
		if err != nil {
			// 404
			resource = "Not Found"
			finish([]byte(resource), int32(len(resource)), true)
			return
		}
		resource = string(data)
	} else if handle && resource == "" {
		// 404
		resource = "Not Found"
		finish([]byte(resource), int32(len(resource)), true)
		return
	}
	finish([]byte(resource), int32(len(resource)), false)
	schemeTask.Free()
	response.Free()
}

func (m *TWebview) onStopURLSchemeTask(sender lcl.IObject, urlSchemeTask wvTypes.WKURLSchemeTask) {
	fmt.Println("OnStopURLSchemeTask")
}
