package wv

import (
	"github.com/energye/energy/v3/internal/assets"
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/wv/wv"
	"sync/atomic"
)

type OnWindowCreate func(window IBrowserWindow)

type IBrowserWindow interface {
	lcl.IForm
	WindowParent() wv.IWVWindowParent
	Browser() wv.IWVBrowser
	// SetOnBrowserAfterCreated Called after a new browser is created and it's ready to navigate to the default URL.
	SetOnBrowserAfterCreated(fn wv.TNotifyEvent)
	// SetOnBrowserMessageReceived
	SetOnBrowserMessageReceived(fn wv.TOnWebMessageReceivedEvent)
	SetOnShow(fn wv.TNotifyEvent)
}

type BrowserWindow struct {
	lcl.TForm
	windowId                uint32
	windowParent            wv.IWVWindowParent
	browser                 wv.IWVBrowser
	options                 Options
	onWindowCreate          OnWindowCreate
	onAfterCreated          wv.TNotifyEvent
	onWebMessageReceived    wv.TOnWebMessageReceivedEvent
	onContextMenuRequested  wv.TOnContextMenuRequestedEvent
	onShow                  wv.TNotifyEvent
	onClose                 lcl.TCloseEvent
	messageReceivedDelegate ipc.IMessageReceivedDelegate
}

var windowID uint32

func getWindowID() uint32 {
	atomic.AddUint32(&windowID, 1)
	return windowID
}

func (m *BrowserWindow) FormCreate(sender lcl.IObject) {
	m.windowId = getWindowID()
	m.SetCaption(m.options.Name)
	m.SetWidth(int32(m.options.Width))
	m.SetHeight(int32(m.options.Height))
	m.ScreenCenter()
	m.SetDoubleBuffered(true)
	m.windowParent = wv.NewWVWindowParent(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)
	m.browser = wv.NewWVBrowser(m)
	m.windowParent.SetBrowser(m.browser)
	if m.options.DefaultURL != "" {
		m.browser.SetDefaultURL(m.options.DefaultURL)
	}
	if m.options.ICON == nil {
		lcl.Application.Icon().LoadFromBytes(assets.ICON.ICO())
	} else {
		lcl.Application.Icon().LoadFromBytes(m.options.ICON)
	}
	ipc.RegisterMessageSend(m)
	m.defaultEvent()
	// call window main form create callback
	if m.onWindowCreate != nil {
		m.onWindowCreate(m)
	}
}

func (m *BrowserWindow) WindowId() uint32 {
	return m.windowId
}

func (m *BrowserWindow) MessageSend() {
	//m.browser.PostWebMessageAsString()
}

func (m *BrowserWindow) defaultEvent() {
	m.messageReceivedDelegate = ipc.NewMessageReceivedDelegate()
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		m.windowParent.UpdateSize()
		settings := m.browser.CoreWebView2Settings()
		settings.SetAreDevToolsEnabled(!m.options.DisableDevTools)
		if m.onAfterCreated != nil {
			m.onAfterCreated(sender)
		}
	})
	m.browser.SetOnContextMenuRequested(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2ContextMenuRequestedEventArgs) {
		if m.options.DisableContextMenu {
			args = wv.NewCoreWebView2ContextMenuRequestedEventArgs(args)
			menuItemCollection := wv.NewCoreWebView2ContextMenuItemCollection(args.MenuItems())
			menuItemCollection.RemoveAllMenuItems()
			menuItemCollection.Free()
			args.Free()
		} else if m.onContextMenuRequested != nil {
			m.onContextMenuRequested(sender, webview, args)
		}
	})
	m.browser.SetOnWebMessageReceived(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebMessageReceivedEventArgs) {
		if m.messageReceivedDelegate != nil {
			args = wv.NewCoreWebView2WebMessageReceivedEventArgs(args)
			message := args.WebMessageAsString()
			args.Free()
			m.messageReceivedDelegate.Received(m.WindowId(), message)
		}
		if m.onWebMessageReceived != nil {
			m.onWebMessageReceived(sender, webview, args)
		}
	})
	m.TForm.SetOnShow(func(sender lcl.IObject) {
		if application.InitializationError() {
			// Log？？
		} else {
			if application.Initialized() {
				m.browser.CreateBrowser(m.windowParent.Handle(), true)
			}
		}
		if m.onShow != nil {
			m.onShow(sender)
		}
	})
	m.TForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		if m.onClose != nil {
			m.onClose(sender, action)
		}
		// 关闭时取消注册
		if *action == types.CaFree {
			ipc.UnRegisterMessageSend(m)
		}
	})
}

func (m *BrowserWindow) WindowParent() wv.IWVWindowParent {
	return m.windowParent
}

func (m *BrowserWindow) Browser() wv.IWVBrowser {
	return m.browser
}

func (m *BrowserWindow) SetOnShow(fn wv.TNotifyEvent) {
	m.onShow = fn
}

func (m *BrowserWindow) SetOnClose(fn lcl.TCloseEvent) {
	m.onClose = fn
}

func (m *BrowserWindow) SetOnBrowserAfterCreated(fn wv.TNotifyEvent) {
	m.onAfterCreated = fn
}

func (m *BrowserWindow) SetOnBrowserMessageReceived(fn wv.TOnWebMessageReceivedEvent) {
	m.onWebMessageReceived = fn
}
