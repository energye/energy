package wv

import (
	"fmt"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/wv/wv"
)

type OnWindowCreate func(window IBrowserWindow)

type IBrowserWindow interface {
	lcl.IForm
	WindowParent() wv.IWVWindowParent
	Browser() wv.IWVBrowser
	SetOnAfterCreated(fn wv.TNotifyEvent)
	SetOnWebMessageReceived(fn wv.TOnWebMessageReceivedEvent)
}

type BrowserWindow struct {
	lcl.TForm
	windowParent         wv.IWVWindowParent
	browser              wv.IWVBrowser
	options              Options
	onWindowCreate       OnWindowCreate
	onAfterCreated       wv.TNotifyEvent
	onWebMessageReceived wv.TOnWebMessageReceivedEvent
}

func (m *BrowserWindow) FormCreate(sender lcl.IObject) {
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
	m.defaultEvent()
	// call window main form create callback
	if m.onWindowCreate != nil {
		m.onWindowCreate(m)
	}
}

func (m *BrowserWindow) defaultEvent() {
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		m.windowParent.UpdateSize()
		if m.onAfterCreated != nil {
			m.onAfterCreated(sender)
		}
	})

	m.SetOnShow(func(sender lcl.IObject) {
		if application.InitializationError() {
			fmt.Println("回调函数 => SetOnShow 初始化失败")
		} else {
			if application.Initialized() {
				fmt.Println("回调函数 => SetOnShow 初始化成功")
				m.browser.CreateBrowser(m.windowParent.Handle(), true)
			}
		}
	})
}

func (m *BrowserWindow) WindowParent() wv.IWVWindowParent {
	return m.windowParent
}

func (m *BrowserWindow) Browser() wv.IWVBrowser {
	return m.browser
}

func (m *BrowserWindow) SetOnAfterCreated(fn wv.TNotifyEvent) {
	m.onAfterCreated = fn
}

func (m *BrowserWindow) SetOnWebMessageReceived(fn wv.TOnWebMessageReceivedEvent) {
	m.onWebMessageReceived = fn
}
