package wv

import (
	"fmt"
	"github.com/energye/energy/v3/internal/assets"
	"github.com/energye/energy/v3/ipc"
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
	id                   uint32
	windowParent         wv.IWVWindowParent
	browser              wv.IWVBrowser
	options              Options
	onWindowCreate       OnWindowCreate
	onAfterCreated       wv.TNotifyEvent
	onWebMessageReceived wv.TOnWebMessageReceivedEvent
	onShow               wv.TNotifyEvent
	messageReceived      ipc.IMessageReceivedDelegate
}

var windowID uint32

func getWindowID() uint32 {
	atomic.AddUint32(&windowID, 1)
	return windowID
}

func (m *BrowserWindow) FormCreate(sender lcl.IObject) {
	m.id = getWindowID()
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
	m.defaultEvent()
	// call window main form create callback
	if m.onWindowCreate != nil {
		m.onWindowCreate(m)
	}
}

func (m *BrowserWindow) MessageSend() {
	//m.browser.PostWebMessageAsString()
}

func (m *BrowserWindow) defaultEvent() {
	m.messageReceived = ipc.NewMessageReceivedDelegate()
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		m.windowParent.UpdateSize()
		if m.onAfterCreated != nil {
			m.onAfterCreated(sender)
		}
	})
	m.browser.SetOnWebMessageReceived(func(sender wv.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebMessageReceivedEventArgs) {
		if m.messageReceived != nil {
			args = wv.NewCoreWebView2WebMessageReceivedEventArgs(args)
			message := args.WebMessageAsString()
			args.Free()
			m.messageReceived.Received(m.id, message)
		}
		if m.onWebMessageReceived != nil {
			m.onWebMessageReceived(sender, webview, args)
		}
	})
	m.TForm.SetOnShow(func(sender lcl.IObject) {
		if application.InitializationError() {
			fmt.Println("回调函数 => SetOnShow 初始化失败")
		} else {
			if application.Initialized() {
				m.browser.CreateBrowser(m.windowParent.Handle(), true)
			}
		}
		if m.onShow != nil {
			m.onShow(sender)
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

func (m *BrowserWindow) SetOnBrowserAfterCreated(fn wv.TNotifyEvent) {
	m.onAfterCreated = fn
}

func (m *BrowserWindow) SetOnBrowserMessageReceived(fn wv.TOnWebMessageReceivedEvent) {
	m.onWebMessageReceived = fn
}
