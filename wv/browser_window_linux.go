//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

package wv

import (
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/pkgs/mime"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	wv "github.com/energye/wv/linux"
	wvTypes "github.com/energye/wv/types/linux"
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
	settings                wv.IWkSettings
	messageReceivedDelegate ipc.IMessageReceivedDelegate
	onBrowserAfterCreated   lcl.TNotifyEvent
	onProcessMessage        TOnProcessMessageEvent
	onResourceRequest       TOnResourceRequestEvent
	onLoadChange            TOnLoadChangeEvent
	onContextMenu           TOnContextMenuEvent
	onContextMenuCommand    TOnContextMenuCommandEvent
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
	if gWk2Context == nil {
		gWk2Context = wv.WebContext.Default()
	}
	if gApplication.onCustomSchemes != nil {
		customSchemes := &TCustomSchemes{}
		gApplication.onCustomSchemes(customSchemes)
		for _, scheme := range customSchemes.schemes {
			gWk2Context.RegisterURIScheme(scheme.Scheme, m.browser.AsSchemeRequestDelegate())
		}
	}
	if gApplication.LocalLoad != nil {
		gWk2Context.RegisterURIScheme(gApplication.LocalLoad.Scheme, m.browser.AsSchemeRequestDelegate())
	}
	m.browser.RegisterScriptCode(string(ipcJS))
	m.browser.RegisterScriptMessageHandler("processMessage")

	m.settings = wv.NewSettings()
	m.settings.SetEnableDeveloperExtras(true)
	m.settings.SetUserAgentWithApplicationDetails("energy", "3.0")
	m.settings.SetEnablePageCache(true)
	// 需要动态判断当前系统环境是否支持？
	m.settings.SetHardwareAccelerationPolicy(wvTypes.WEBKIT_HARDWARE_ACCELERATION_POLICY_NEVER)
	//setting.SetHardwareAccelerationPolicy(wvTypes.WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS)
	m.browser.SetSettings(m.settings)
	// ipc message received
	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	ipc.RegisterProcessMessage(m)
	m.initDefaultEvent()
	m.SetBrowserOptions()
	return m
}

//func (m *TWebview) SetAlign(v types.TAlign) {
//	switch v {
//	case types.AlClient:
//		part := m.Parent()
//		m.SetWidth(part.Width())
//		m.SetHeight(part.Height())
//		m.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))
//	default:
//		m.ICustomPanel.SetAlign(v)
//	}
//}

func (m *TWebview) SetWidth(v int32) {
	m.ICustomPanel.SetWidth(v + 1) // Gtk3 box width + 1
	m.windowParent.SetWidth(v)
}

func (m *TWebview) SetHeight(v int32) {
	m.ICustomPanel.SetHeight(v + 1) // Gtk3 box height + 1
	m.windowParent.SetHeight(v)
}

func (m *TWebview) SetBoundsRect(value types.TRect) {
	m.SetBounds(value.Left, value.Top, value.Width(), value.Height())
}

func (m *TWebview) SetBounds(left int32, top int32, width int32, height int32) {
	m.ICustomPanel.SetBounds(left, top, width+1, height+1) // Gtk3 box width + 1 height + 1
	m.windowParent.SetBounds(0, 0, width, height)
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
	m.browser.CreateBrowser()
	m.windowParent.SetWebview(m.browser)
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
	m.browser.ExecuteScript(js)
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
	*canClose = m.window.IsClose()
	if !m.window.IsClose() {
		m.window.SetClose(true)
		m.browser.Stop()
		m.browser.TerminateWebProcess()
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

func (m *TWebview) ExecuteScript(javaScript string) {
	m.browser.ExecuteScript(javaScript)
}

func (m *TWebview) initDefaultEvent() {
	m.browser.SetOnContextMenu(func(sender lcl.IObject, contextMenu wvTypes.WebKitContextMenu, defaultAction wvTypes.PWkAction) bool {
		rootContextMenu := wv.NewContextMenu(contextMenu)
		defer rootContextMenu.Free()
		menuItemClear := func(menuItems wv.IWkContextMenu) {
			if menuItems == nil {
				return
			}
			menuItems.RemoveAll()
		}
		if gApplication.Options.DisableContextMenu {
			menuItemClear(rootContextMenu)
			return true
		}
		if m.onContextMenu != nil {
			var (
				add                      func(text string, kind TContextMenuKind, menuItems wv.IWkContextMenu) (*TContextMenuItem, int32)
				tempFreeContextMenus     []wv.IWkContextMenu
				tempFreeContextMenuItems []wv.IWkContextMenuItem
			)
			add = func(text string, kind TContextMenuKind, menuItems wv.IWkContextMenu) (*TContextMenuItem, int32) {
				var (
					newMenuItem    wv.IWkContextMenuItem
					subContextMenu wv.IWkContextMenu
				)
				commandId := nextContextMenuCommandId()
				switch kind {
				case CmkCommand:
					newMenuItem = wv.ContextMenuItem.NewFromAction(defaultAction, text, commandId)
				case CmkSub:
					newMenuItem = wv.ContextMenuItem.NewFromAction(defaultAction, text, commandId)
					subContextMenu = wv.ContextMenu.New()
					newMenuItem.SetSubmenu(subContextMenu.Data())
					tempFreeContextMenus = append(tempFreeContextMenus, subContextMenu)
				case CmkSeparator:
					newMenuItem = wv.ContextMenuItem.NewSeparator()
				default:
					return nil, 0
				}
				tempFreeContextMenuItems = append(tempFreeContextMenuItems, newMenuItem)

				menuItems.Append(newMenuItem.Data())
				childContextMenu := &TContextMenuItem{
					clear: func() {
						menuItemClear(subContextMenu)
					},
					add: func(text string, kind TContextMenuKind) (*TContextMenuItem, int32) {
						newMenuItem, newCommandId := add(text, kind, subContextMenu)
						return newMenuItem, newCommandId
					}}
				return childContextMenu, commandId
			}
			contextMenuItem := &TContextMenuItem{
				clear: func() {
					menuItemClear(rootContextMenu)
				},
				add: func(text string, kind TContextMenuKind) (*TContextMenuItem, int32) {
					return add(text, kind, rootContextMenu)
				}}
			m.onContextMenu(contextMenuItem)
			for _, item := range tempFreeContextMenuItems {
				item.Free()
			}
			for _, item := range tempFreeContextMenus {
				item.Free()
			}
		}
		return false
	})
	m.browser.SetOnContextMenuCommand(func(sender lcl.IObject, menuID int32) {
		if m.onContextMenuCommand != nil {
			m.onContextMenuCommand(menuID)
		}
	})
	m.browser.SetOnDecidePolicy(func(sender lcl.IObject, wkDecision wvTypes.WebKitPolicyDecision, type_ wvTypes.WebKitPolicyDecisionType) bool {
		fmt.Println("OnDecidePolicy type_:", type_)
		tempDecision := wv.NewNavigationPolicyDecision(wkDecision)
		defer tempDecision.Free()
		if type_ == wvTypes.WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION || type_ == wvTypes.WEBKIT_POLICY_DECISION_TYPE_NAVIGATION_ACTION {
			tempNavigationAction := wv.NewNavigationAction(tempDecision.GetNavigationAction())
			defer tempNavigationAction.Free()
			tempURIRequest := wv.NewURIRequest(tempNavigationAction.GetRequest())
			defer tempURIRequest.Free()
			newWindowURL := tempURIRequest.URI()
			fmt.Println("NewWindow URL:", newWindowURL)
			// new window
			if type_ == wvTypes.WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION {

			}
		} else {
			tempResponsePolicyDecision := wv.NewResponsePolicyDecision(wkDecision)
			defer tempResponsePolicyDecision.Free()
			tempURIRequest := wv.NewURIRequest(tempResponsePolicyDecision.GetRequest())
			defer tempURIRequest.Free()
		}
		return true
	})
	m.browser.SetOnLoadChange(func(sender lcl.IObject, loadEvent wvTypes.WebKitLoadEvent) {
		switch loadEvent {
		case wvTypes.WEBKIT_LOAD_FINISHED:
			m.navigationStarting()
		}
		if m.onLoadChange != nil {
			uri := m.browser.GetURI()
			title := m.browser.GetTitle()
			switch loadEvent {
			case wvTypes.WEBKIT_LOAD_STARTED:
				m.onLoadChange(uri, title, LcStart)
			case wvTypes.WEBKIT_LOAD_REDIRECTED, wvTypes.WEBKIT_LOAD_COMMITTED:
				m.onLoadChange(uri, title, LcLoading)
			case wvTypes.WEBKIT_LOAD_FINISHED:
				m.onLoadChange(uri, title, LcFinish)
			}
		}
	})
	m.browser.SetOnWebProcessTerminated(func(sender lcl.IObject, reason wvTypes.WebKitWebProcessTerminationReason) {
		if reason == wvTypes.WEBKIT_WEB_PROCESS_TERMINATED_BY_API { //  call m.webview.TerminateWebProcess()
			lcl.RunOnMainThreadAsync(func(id uint32) {
				m.window.Close()
			})
		}
	})
	m.browser.SetOnProcessMessage(func(sender lcl.IObject, jsValue wv.IWkJSValue, processId wvTypes.TWkProcessId) {
		var handle bool
		message := jsValue.StringValue()
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
					if m.window != nil {
						m.browser.StartDrag(m.window)
						handle = true
					}
				case ipc.MT_DRAG_RESIZE:
					// border drag resize
					if m.window != nil {
						m.resizeHT = pMessage.Data.(string)
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
	m.browser.SetOnURISchemeRequest(func(sender lcl.IObject, wkURISchemeRequest wvTypes.WebKitURISchemeRequest) {
		uriSchemeRequest := wv.NewURISchemeRequest(wkURISchemeRequest)
		defer uriSchemeRequest.Free()
		var (
			resource    string
			handle      bool
			uri         = uriSchemeRequest.Uri()
			path        = uriSchemeRequest.Path()
			method      = uriSchemeRequest.Method()
			contentType = "Content-Type: " + mime.GetMimeType(path)
		)
		if m.onResourceRequest != nil {
			header := make(map[string]string)
			headers := wv.NewHeaders(uriSchemeRequest.Headers())
			defer headers.Free()
			headList := headers.List()
			if headList != nil {
				defer headList.Free()
				count := int(headList.Count())
				for i := 0; i < count; i++ {
					key := headList.Names(int32(i))
					val := headList.Values(key)
					header[key] = val
				}
			}
			resource, handle = m.onResourceRequest(uri, path, method, header)
		}
		if !handle {
			data, err := gApplication.LocalLoad.Read(path)
			if err != nil {
				// 404
				uriSchemeRequest.FinishError(3, 404, err.Error())
				return
			}
			ins := wv.InputStream.New(uintptr(unsafe.Pointer(&data[0])), int64(len(data)))
			uriSchemeRequest.Finish(ins.Data(), int64(len(data)), contentType)
		} else if handle && resource != "" {
			data := []byte(resource)
			ins := wv.InputStream.New(uintptr(unsafe.Pointer(&data[0])), int64(len(data)))
			uriSchemeRequest.Finish(ins.Data(), int64(len(data)), contentType)
		} else {
			// 404
			uriSchemeRequest.FinishError(3, 404, "Not Found")
		}
	})
	var (
		mouseCursor types.TCursor
		setCursor   = func(value types.TCursor, ht string) {
			m.resizeHT = ht
			mouseCursor = value
			lcl.Screen.SetCursor(value)
		}
		isDown                 bool
		mouseDownX, mouseDownY int32
		windowBr               types.TRect
	)
	m.browser.SetOnMouseMove(func(sender lcl.IObject, event wv.TWkButtonEvent) bool {
		if !isDown {
			br := m.BoundsRect()
			w, h := br.Width(), br.Height()
			x, y := event.X, event.Y
			if (w-x) < (frameWidth+frameCorner) && (h-y) < (frameHeight+frameCorner) {
				setCursor(types.CrSizeSE, "se-resize")
			} else if x < (frameWidth+frameCorner) && (h-y) < (frameHeight+frameCorner) {
				setCursor(types.CrSizeSW, "sw-resize")
			} else if x < (frameWidth+frameCorner) && y < (frameHeight+frameCorner) {
				setCursor(types.CrSizeNW, "nw-resize")
			} else if (w-x) < (frameWidth+frameCorner) && y < (frameHeight+frameCorner) {
				setCursor(types.CrSizeNE, "ne-resize")
			} else if x < frameWidth {
				setCursor(types.CrSizeW, "w-resize")
			} else if y < frameHeight {
				setCursor(types.CrSizeN, "n-resize")
			} else if (h - y) < frameHeight {
				setCursor(types.CrSizeS, "s-resize")
			} else if (w - x) < frameWidth {
				setCursor(types.CrSizeE, "e-resize")
			} else {
				setCursor(types.CrDefault, "")
			}
		} else if isDown && mouseCursor != types.CrDefault {
			return m.resize(isDown, windowBr, mouseDownX, mouseDownY)
		} else {
			setCursor(types.CrDefault, "")
		}
		return false
	})
	m.browser.SetOnMousePress(func(sender lcl.IObject, event wv.TWkButtonEvent) bool {
		isDown = true
		if mouseCursor != types.CrDefault && m.window != nil {
			pos := lcl.Mouse.CursorPos()
			mouseDownX, mouseDownY = pos.X, pos.Y
			windowBr = m.window.BoundsRect()
			return true
		}
		return false
	})
	m.browser.SetOnMouseRelease(func(sender lcl.IObject, event wv.TWkButtonEvent) bool {
		isDown = false
		return false
	})
}

func (m *TWebview) resize(isDown bool, windowBr types.TRect, mouseDownX, mouseDownY int32) bool {
	if !isDown || m.resizeHT == "" || m.window == nil {
		return false
	}
	currentX := windowBr.Left
	currentY := windowBr.Top
	currentW := windowBr.Width()
	currentH := windowBr.Height()

	pos := lcl.Mouse.CursorPos()
	mouseMoveX, mouseMoveY := pos.X, pos.Y
	dx := mouseMoveX - mouseDownX
	dy := mouseMoveY - mouseDownY

	newX, newY := currentX, currentY
	newW, newH := currentW, currentH

	switch m.resizeHT {
	case "se-resize": // 右下角
		newW = currentW + dx
		newH = currentH + dy
	case "sw-resize": // 左下角
		newW = currentW - dx
		newH = currentH + dy
		newX = currentX + dx
	case "nw-resize": // 左上角
		newW = currentW - dx
		newH = currentH - dy
		newX = currentX + dx
		newY = currentY + dy
	case "ne-resize": // 右上角
		newW = currentW + dx
		newH = currentH - dy
		newY = currentY + dy
	case "w-resize": // 左
		newW = currentW - dx
		newX = currentX + dx
	case "n-resize": // 上
		newH = currentH - dy
		newY = currentY + dy
	case "s-resize": // 下
		newH = currentH + dy
	case "e-resize": // 右
		newW = currentW + dx
	}
	m.window.SetBounds(newX, newY, newW, newH)
	return true
}
