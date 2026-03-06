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
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/pkgs/cocoa"
	"github.com/energye/energy/v3/pkgs/mime"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	wv "github.com/energye/wv/darwin"
	wvTypes "github.com/energye/wv/types/darwin"
	"strings"
	"unsafe"
)

var (
	frameWidth  = int32(4)
	frameHeight = int32(4)
	frameCorner = frameWidth + frameHeight
)

type _TNSWebview struct {
	nsWebview wvTypes.WKWebView
}

type TWebview struct {
	lcl.ICustomPanel
	TEnergyWebview
	browserId               uint32
	isClose                 bool
	isCreated               bool
	resizeHT                string
	menu                    lcl.IPopupMenu
	window                  window.IWindow
	nsWindow                lcl.NSWindow
	oldBounds               types.TRect
	isAddNSWindowSubview    bool
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
	onDragEnter             TOnDragEnterEvent
	onDragLeave             TOnDragLeaveEvent
	onDragOver              TOnDragOverEvent
}

// NewWebview 创建一个新的浏览器窗口实例
func NewWebview(owner lcl.IComponent) IWebview {
	m := &TWebview{browserId: getNextBrowserID()}
	m.menu = lcl.NewPopupMenu(owner)

	m.ICustomPanel = lcl.NewCustomPanel(owner)
	m.ICustomPanel.SetParentDoubleBuffered(true)
	m.ICustomPanel.SetBevelInner(types.BvNone)
	m.ICustomPanel.SetBevelOuter(types.BvNone)

	m.windowParent = wv.NewWebviewParent(m)
	m.windowParent.SetAlign(types.AlClient)
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
			configuration.SetURLSchemeHandlerForURLScheme(URLSchemeHandler.Data(), scheme.Scheme)
		}
	}
	if gApplication.LocalLoad != nil {
		configuration.SetURLSchemeHandlerForURLScheme(URLSchemeHandler.Data(), gApplication.LocalLoad.Scheme)
	}

	preference := wv.NewPreferences(configuration.Preferences())
	configuration.SetPreferences(preference.Data())
	preference.SetTabFocusesLinks(true)
	preference.SetFraudulentWebsiteWarningEnabled(true)
	m._EnableDevTools(preference)

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
	m.initDefaultDragEvent()
	return m
}

func (m *TWebview) SetWindow(iWindow window.IWindow) {
	m.window = iWindow
	if m.window != nil {
		if m.window.BrowserId() == 0 {
			m.window.SetBrowserId(m.browserId)
		}
	}
	m.window.AddOnWindowStateChange(m.doOnWindowStateChange)
	m.window.AddOnWindowResize(m.doOnWindowResize)
	m.window.AddOnWindowShow(m.doOnWindowShow)
	m.window.AddOnWindowClose(m.doOnWindowClose)
	m.window.AddOnWindowCloseQuery(m.doOnWindowCloseQuery)
}

// UpdateBrowserOptions 更新浏览器配置
func (m *TWebview) UpdateBrowserOptions() {
	// 1. 获取 LocalLoad 全局配置
	if application.GApplication != nil && application.GApplication.LocalLoad != nil {
		newLocalLoad := *application.GApplication.LocalLoad.LocalLoad
		m.SetLocalLoad(newLocalLoad)
	}

	options := m.window.Options()
	if options.DefaultURL != "" {
		m.browser.LoadURL(options.DefaultURL)
	}

	m.SetWebviewTransparent(options.WebviewTransparent)
	m.BecomeFirstResponder()
	m._WebViewRegisterPerformKeyEquivalentMethod()
}

// SetParent 设置浏览器窗口的父控件
// 该方法会同时设置内部面板的父控件和窗口父控件的引用
func (m *TWebview) SetParent(owner lcl.IWinControl) {
	if form, okForm := owner.(window.IWindow); okForm {
		// webview 直接添加到窗口
		m.AddWindowSubviewWebview(form)
	} else {
		// webview 添加到容器组件
		m.windowParent.SetParent(m)
		m.ICustomPanel.SetParent(owner)
	}
}

// 在窗口显示时调用
func (m *TWebview) CreateBrowser() {
	if m.isCreated || m.window == nil {
		return
	}
	m.isCreated = true
	m.UpdateBrowserOptions()
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
func (m *TWebview) doOnWindowShow(sender lcl.IObject) {
	if m.isCreated {
		return
	}
	m.UpdateBounds()
	m.CreateBrowser()
}

// onWindowClose 处理窗口关闭事件的回调函数
// 当窗口接收到关闭信号时，该函数会停止浏览器实例以确保资源被正确释放
func (m *TWebview) doOnWindowClose(sender lcl.IObject, closeAction *types.TCloseAction) {
}

// onWindowCloseQuery 处理窗口关闭查询事件
// 当用户尝试关闭窗口时触发此回调函数
func (m *TWebview) doOnWindowCloseQuery(sender lcl.IObject, canClose *bool) {
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

func (m *TWebview) doOnWindowStateChange(sender lcl.IObject) {
	switch m.window.WindowState() {
	case types.WsNormal, types.WsMinimized, types.WsMaximized:
	case types.WsFullScreen:
		//cocoa.AppSetPresentationOptions(cocoa.NSApplicationPresentationAutoHideToolbar | cocoa.NSApplicationPresentationAutoHideMenuBar | cocoa.NSApplicationPresentationFullScreen)
	}
}

func (m *TWebview) doOnWindowResize(sender lcl.IObject) {
	m.UpdateBounds()
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

func (m *TWebview) SetOnDragEnter(fn TOnDragEnterEvent) {
	m.onDragEnter = fn
}

func (m *TWebview) SetOnDragLeave(fn TOnDragLeaveEvent) {
	m.onDragLeave = fn
}

func (m *TWebview) SetOnDragOver(fn TOnDragOverEvent) {
	m.onDragOver = fn
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
		//println("OnProcessMessage", name, "message:", message, api.MainThreadId() == api.CurrentThreadId())
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
					if currentWindow := getWindow(); currentWindow != nil {
						currentWindow.DragWindow()
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
				case ipc.MT_CONTEXTMENU:
					if m.onContextMenu != nil {
						data := pMessage.Data.(map[string]any)
						x := int32(data["x"].(float64))
						y := int32(data["y"].(float64))
						m.contextMenu(x, y)
					}
				}
			}
		}
		if !handle && m.onProcessMessage != nil {
			m.onProcessMessage(message)
		}
	})
	m.browser.SetOnStartProvisionalNavigation(func(sender lcl.IObject, navigation wvTypes.WKNavigation) {
		if m.onLoadChange != nil {
			nsurl := m.browser.URL()
			wkNsUrl := wv.NewURL(nsurl)
			defer wkNsUrl.Free()
			url := wkNsUrl.AbsoluteString()
			title := m.browser.Title()
			m.onLoadChange(url, title, LcStart)
		}
	})
	m.browser.SetOnCommitNavigation(func(sender lcl.IObject, navigation wvTypes.WKNavigation) {
		if m.onLoadChange != nil {
			nsurl := m.browser.URL()
			wkNsUrl := wv.NewURL(nsurl)
			defer wkNsUrl.Free()
			url := wkNsUrl.AbsoluteString()
			title := m.browser.Title()
			m.onLoadChange(url, title, LcLoading)
		}
	})
	m.browser.SetOnFinishNavigation(func(sender lcl.IObject, navigation wvTypes.WKNavigation) {
		//println("browser-OnFinishNavigation")
		m.createEnergyJavasScript()
		m.listenDarwinContextMenuJavaScript()
		if m.onLoadChange != nil {
			nsurl := m.browser.URL()
			wkNsUrl := wv.NewURL(nsurl)
			defer wkNsUrl.Free()
			url := wkNsUrl.AbsoluteString()
			title := m.browser.Title()
			m.onLoadChange(url, title, LcFinish)
		}
	})
	m.browser.SetOnCreateWebView(func(sender lcl.IObject, configuration wvTypes.WKWebViewConfiguration, navigationAction wvTypes.WKNavigationAction, windowFeatures wvTypes.WKWindowFeatures) wvTypes.WKWebView {
		if m.onPopupWindow != nil {
			wkNavigationAction := wv.NewNavigationAction(navigationAction)
			request := wv.NewURLRequest(wkNavigationAction.Request())
			if request.IsValid() {
				url := wv.NewURL(request.URL())
				defer url.Free()
				targetURL := url.AbsoluteString()
				m.onPopupWindow(targetURL)
			}
		}
		return 0
	})
	m.browser.SetOnStartURLSchemeTask(m.onStartURLSchemeTask)
}

func (m *TWebview) onStartURLSchemeTask(sender lcl.IObject, urlSchemeTask wvTypes.WKURLSchemeTask) {
	schemeTask := wv.NewURLSchemeTask(urlSchemeTask)
	request := wv.NewURLRequest(schemeTask.Request())
	defer request.Free()
	url := wv.NewURL(request.URL())
	defer url.Free()
	var (
		resource    string
		handle      bool
		uri         = url.AbsoluteString()
		path        = url.Path()
		method      = request.HTTPMethod()
		contentType = "Content-Type: " + mime.GetMimeType(path)
	)
	//println("StartURLSchemeTask:", uri, path, method)
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

func (m *TWebview) listenDarwinContextMenuJavaScript() {
	m.ExecuteScript(`window.energy.__listenDarwinContextMenu();`)
}

func (m *TWebview) contextMenu(x, y int32) {
	items := m.menu.Items()
	items.Clear()
	createMenuItem := func(text string, fn func(commandId int32)) (lcl.IMenuItem, int32) {
		commandId := nextContextMenuCommandId()
		item := lcl.NewMenuItem(m)
		item.SetCaption(text)
		if text != "-" && fn != nil {
			item.SetOnClick(func(lcl.IObject) {
				fn(commandId)
			})
		}
		return item, commandId
	}
	//newReloadItem, _ := createMenuItem("重新载入", func(commandId int32) {
	//	m.browser.Reload()
	//})
	//items.Add(newReloadItem)
	//if !gApplication.Options.DisableDevTools {
	//	newDevtoolItem, _ := createMenuItem("开发者工具", func(commandId int32) {
	//		lcl.RunOnMainThreadAsync(func(id uint32) {
	//			m.ExecuteScript("webkit.inspectElement(100, 100);")
	//		})
	//	})
	//	items.Add(newDevtoolItem)
	//}
	menuItemClear := func(menuItems lcl.IMenuItem) {
		if menuItems == nil {
			return
		}
		menuItems.Clear()
	}
	//if gApplication.Options.DisableContextMenu {
	//	menuItemClear(items)
	//	return
	//}
	var (
		add func(text string, kind TContextMenuKind, menuItems lcl.IMenuItem) (*TContextMenuItem, int32)
	)
	add = func(text string, kind TContextMenuKind, menuItems lcl.IMenuItem) (*TContextMenuItem, int32) {
		var (
			newCtxMenuItem lcl.IMenuItem
			newCommandId   int32
		)
		switch kind {
		case CmkCommand:
			newCtxMenuItem, newCommandId = createMenuItem(text, func(commandId int32) {
				if m.onContextMenuCommand != nil {
					m.onContextMenuCommand(commandId)
				}
			})
			menuItems.Add(newCtxMenuItem)
		case CmkSub:
			newCtxMenuItem, newCommandId = createMenuItem(text, func(commandId int32) {
				if m.onContextMenuCommand != nil {
					m.onContextMenuCommand(commandId)
				}
			})
			menuItems.Add(newCtxMenuItem)
		case CmkSeparator:
			newCtxMenuItem, _ = createMenuItem("-", nil)
			menuItems.Add(newCtxMenuItem)
		default:
			return nil, 0
		}
		childContextMenu := &TContextMenuItem{
			clear: func() {
				menuItemClear(newCtxMenuItem)
			},
			add: func(text string, kind TContextMenuKind) (*TContextMenuItem, int32) {
				newMenuItem, newCommandId := add(text, kind, newCtxMenuItem)
				return newMenuItem, newCommandId
			}}
		return childContextMenu, newCommandId
	}
	contextMenuItem := &TContextMenuItem{
		clear: func() {
			items.Clear()
		},
		add: func(text string, kind TContextMenuKind) (*TContextMenuItem, int32) {
			return add(text, kind, items)
		}}
	m.onContextMenu(contextMenuItem)
	m.menu.PopUpWithIntX2(x+1, y+1)
}

//func (m *TWebview) onStopURLSchemeTask(sender lcl.IObject, urlSchemeTask wvTypes.WKURLSchemeTask) {
//	fmt.Println("OnStopURLSchemeTask")
//}

func (m *TWebview) initDefaultDragEvent() {
	leave := false
	m.browser.SetOnDraggingEntered(func(sender wv.NSDraggingInfoProtocol, handle *bool) int32 {
		if m.onDragEnter == nil {
			return cocoa.NSDragOperationNone
		}
		dragInfo := cocoa.WrapNSDraggingInfo(sender)
		dragPasteboard := dragInfo.DraggingPasteboard()
		nsTypes := dragPasteboard.Types()
		isPlainText := nsTypes.In("public.utf8-plain-text")
		isFilename := nsTypes.In("public.file-url")
		if isPlainText || isFilename {
			leave = true
			type_ := DragTypeNo
			if isPlainText {
				type_ = DragTypeData
			} else if isFilename {
				type_ = DragTypeFile
			}
			point := m.ConvertPoint(dragInfo.DraggingLocation())
			m.onDragEnter(type_, point.X, point.Y)
			*handle = true
			return cocoa.NSDragOperationCopy
		}
		return cocoa.NSDragOperationNone
	})
	m.browser.SetOnDraggingUpdated(func(sender wv.NSDraggingInfoProtocol, handle *bool) int32 {
		dragInfo := cocoa.WrapNSDraggingInfo(sender)
		dragPasteboard := dragInfo.DraggingPasteboard()
		nsTypes := dragPasteboard.Types()
		isPlainText := nsTypes.In("public.utf8-plain-text")
		isFilename := nsTypes.In("public.file-url")
		if isPlainText || isFilename {
			if m.onDragLeave != nil && leave {
				leave = false
				m.onDragLeave()
			}
			*handle = true
			return cocoa.NSDragOperationCopy
		}
		return cocoa.NSDragOperationNone
	})
	//m.browser.SetOnDraggingEnded(func(sender wv.NSDraggingInfoProtocol, handle *bool) {
	//	println("SetOnDraggingEnded")
	//})
	//m.browser.SetOnDraggingExited(func(sender wv.NSDraggingInfoProtocol, handle *bool) {
	//	println("SetOnDraggingExited")
	//})
	m.browser.SetOnPrepareForDragOperation(func(sender wv.NSDraggingInfoProtocol) bool {
		return true
	})
	m.browser.SetOnPerformDragOperation(func(sender wv.NSDraggingInfoProtocol) bool {
		if m.onDragOver == nil {
			return false
		}
		dragInfo := cocoa.WrapNSDraggingInfo(sender)
		dragPasteboard := dragInfo.DraggingPasteboard()
		nsTypes := dragPasteboard.Types()
		isPlainText := nsTypes.In("public.utf8-plain-text")
		isFilename := nsTypes.In("public.file-url")
		if isPlainText || isFilename {
			point := m.ConvertPoint(dragInfo.DraggingLocation())
			data := dragPasteboard.PasteboardData()
			type_ := DragTypeNo
			var (
				plainTextData []byte
				filenames     []string
			)
			if isPlainText {
				type_ = DragTypeData
				plainTextData = []byte(strings.Join(data.Texts, "\n"))
			} else if isFilename {
				type_ = DragTypeFile
				filenames = data.FilePaths
			}
			dragData := &TDragData{Type: type_, Data: plainTextData, Filenames: filenames}
			m.onDragOver(dragData, point.X, point.Y)
			return true
		}
		return false
	})
}
